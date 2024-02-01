// Copyright 2024 Apache Software Foundation (ASF)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package services

import (
	"context"

	"github.com/apache/incubator-kie-kogito-serverless-operator/controllers/knative"

	"github.com/apache/incubator-kie-kogito-serverless-operator/container-builder/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	eventingv1 "knative.dev/eventing/pkg/apis/eventing/v1"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func IsKnativeAvailableForPlatformService(ctx context.Context, c client.Client, brokerName *string, namespace string) bool {
	if knativeAvail, err := getKnativeAvailability(c); err != nil || knativeAvail == nil || !knativeAvail.Eventing || brokerName == nil {
		return false
	}
	broker := &eventingv1.Broker{
		ObjectMeta: metav1.ObjectMeta{
			Name:      *brokerName,
			Namespace: namespace,
		},
	}
	if err := c.Get(ctx, k8sclient.ObjectKeyFromObject(broker), broker); err != nil {
		return false
	}
	return true
}

func getKnativeAvailability(c client.Client) (*knative.Availability, error) {
	result := new(knative.Availability)

	if c.Scheme().IsGroupRegistered("serving.knative.dev") {
		result.Serving = true
	}
	if c.Scheme().IsGroupRegistered("eventing.knative.dev") {
		result.Eventing = true
	}
	return result, nil

}
