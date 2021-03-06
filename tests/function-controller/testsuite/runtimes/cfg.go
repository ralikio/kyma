package runtimes

import (
	"fmt"
	"math/rand"
	"net/url"

	"k8s.io/client-go/kubernetes"

	"github.com/kyma-project/kyma/tests/function-controller/pkg/apirule"
	"github.com/kyma-project/kyma/tests/function-controller/pkg/function"
	"github.com/kyma-project/kyma/tests/function-controller/pkg/job"
	"github.com/kyma-project/kyma/tests/function-controller/pkg/servicebinding"
	"github.com/kyma-project/kyma/tests/function-controller/pkg/servicebindingusage"
	"github.com/kyma-project/kyma/tests/function-controller/pkg/serviceinstance"
	"github.com/kyma-project/kyma/tests/function-controller/pkg/shared"
	"github.com/kyma-project/kyma/tests/function-controller/pkg/subscription"
)

type FunctionTestConfig struct {
	FnName          string
	APIRuleURL      *url.URL
	Fn              *function.Function
	ApiRule         *apirule.APIRule
	SvcInstance     *serviceinstance.ServiceInstance
	SvcInstanceName string
	SvcBinding      *servicebinding.ServiceBinding
	SvcBindingUsage *servicebindingusage.ServiceBindingUsage
	UsageName       string
	Subscription    *subscription.Subscription
	Job             *job.Job
	InClusterURL    *url.URL
	SinkURL         *url.URL
	SvcBindingName  string
}

func NewFunctionConfig(fnName, usageKindName, domainName string, toolBox shared.Container, clientset *kubernetes.Clientset) (FunctionTestConfig, error) {
	svcInstanceName := fmt.Sprintf("%s-service-instance", fnName)
	svcUsageName := fmt.Sprintf("%s-service-binding-usage", fnName)
	svcBindingName := fmt.Sprintf("%s-service-binding", fnName)

	gatewayURL, err := url.Parse(fmt.Sprintf("https://%s-%d.%s", fnName, rand.Uint32(), domainName))
	if err != nil {
		return FunctionTestConfig{}, err
	}

	inClusterURL, err := url.Parse(fmt.Sprintf("http://%s.%s.svc.cluster.local", fnName, toolBox.Namespace))
	if err != nil {
		return FunctionTestConfig{}, err
	}

	sinkURL, err := url.Parse(fmt.Sprintf("http://%s.%s.svc.cluster.local", fnName, toolBox.Namespace))
	if err != nil {
		return FunctionTestConfig{}, err
	}

	return FunctionTestConfig{
		FnName:          fnName,
		Fn:              function.NewFunction(fnName, toolBox),
		InClusterURL:    inClusterURL,
		ApiRule:         apirule.New(fmt.Sprintf("%s-rule", fnName), toolBox),
		APIRuleURL:      gatewayURL,
		SvcInstance:     serviceinstance.New(svcInstanceName, toolBox),
		SvcInstanceName: svcInstanceName,
		SvcBinding:      servicebinding.New(svcBindingName, toolBox),
		SvcBindingName:  svcBindingName,
		SvcBindingUsage: servicebindingusage.New(svcUsageName, usageKindName, toolBox),
		UsageName:       svcUsageName,
		Subscription:    subscription.New(fmt.Sprintf("%s-subscription", fnName), toolBox),
		Job:             job.New(fnName, clientset.BatchV1(), toolBox),
		SinkURL:         sinkURL,
	}, nil
}

type FunctionSimpleTestConfig struct {
	FnName       string
	Fn           *function.Function
	InClusterURL *url.URL
}

func NewFunctionSimpleConfig(fnName string, toolBox shared.Container) (FunctionSimpleTestConfig, error) {
	inClusterURL, err := url.Parse(fmt.Sprintf("http://%s.%s.svc.cluster.local", fnName, toolBox.Namespace))
	if err != nil {
		return FunctionSimpleTestConfig{}, err
	}

	return FunctionSimpleTestConfig{
		FnName:       fnName,
		Fn:           function.NewFunction(fnName, toolBox),
		InClusterURL: inClusterURL,
	}, nil
}
