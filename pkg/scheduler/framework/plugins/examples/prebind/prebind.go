/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package prebind

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	fwk "k8s.io/kube-scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

// StatelessPreBindExample is an example of a simple plugin that has no state
// and implements only one hook for prebind.
type StatelessPreBindExample struct{}

var _ framework.PreBindPlugin = StatelessPreBindExample{}

// Name is the name of the plugin used in Registry and configurations.
const Name = "stateless-prebind-plugin-example"

// Name returns name of the plugin. It is used in logs, etc.
func (sr StatelessPreBindExample) Name() string {
	return Name
}

func (mc StatelessPreBindExample) PreBindPreFlight(ctx context.Context, state fwk.CycleState, pod *v1.Pod, nodeName string) *fwk.Status {
	return nil
}

// PreBind is the functions invoked by the framework at "prebind" extension point.
func (sr StatelessPreBindExample) PreBind(ctx context.Context, state fwk.CycleState, pod *v1.Pod, nodeName string) *fwk.Status {
	if pod == nil {
		return fwk.NewStatus(fwk.Error, "pod cannot be nil")
	}
	if pod.Namespace != "foo" {
		return fwk.NewStatus(fwk.Unschedulable, "only pods from 'foo' namespace are allowed")
	}
	return nil
}

// New initializes a new plugin and returns it.
func New(_ context.Context, _ *runtime.Unknown, _ framework.Handle) (framework.Plugin, error) {
	return &StatelessPreBindExample{}, nil
}
