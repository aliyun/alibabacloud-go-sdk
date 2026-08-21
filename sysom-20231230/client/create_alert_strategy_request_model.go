// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *CreateAlertStrategyRequest
	GetXDebugId() *string
	SetEnabled(v bool) *CreateAlertStrategyRequest
	GetEnabled() *bool
	SetK8sLabel(v bool) *CreateAlertStrategyRequest
	GetK8sLabel() *bool
	SetName(v string) *CreateAlertStrategyRequest
	GetName() *string
	SetStrategy(v *CreateAlertStrategyRequestStrategy) *CreateAlertStrategyRequest
	GetStrategy() *CreateAlertStrategyRequestStrategy
	SetXSysomInvokeSource(v string) *CreateAlertStrategyRequest
	GetXSysomInvokeSource() *string
}

type CreateAlertStrategyRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// Specifies whether the alert policy is enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The Kubernetes label.
	K8sLabel *bool `json:"k8sLabel,omitempty" xml:"k8sLabel,omitempty"`
	// The Policy Name of the alerting policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// strategy1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The details of the alert policy.
	//
	// This parameter is required.
	Strategy           *CreateAlertStrategyRequestStrategy `json:"strategy,omitempty" xml:"strategy,omitempty" type:"Struct"`
	XSysomInvokeSource *string                             `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s CreateAlertStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertStrategyRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertStrategyRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *CreateAlertStrategyRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateAlertStrategyRequest) GetK8sLabel() *bool {
	return s.K8sLabel
}

func (s *CreateAlertStrategyRequest) GetName() *string {
	return s.Name
}

func (s *CreateAlertStrategyRequest) GetStrategy() *CreateAlertStrategyRequestStrategy {
	return s.Strategy
}

func (s *CreateAlertStrategyRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *CreateAlertStrategyRequest) SetXDebugId(v string) *CreateAlertStrategyRequest {
	s.XDebugId = &v
	return s
}

func (s *CreateAlertStrategyRequest) SetEnabled(v bool) *CreateAlertStrategyRequest {
	s.Enabled = &v
	return s
}

func (s *CreateAlertStrategyRequest) SetK8sLabel(v bool) *CreateAlertStrategyRequest {
	s.K8sLabel = &v
	return s
}

func (s *CreateAlertStrategyRequest) SetName(v string) *CreateAlertStrategyRequest {
	s.Name = &v
	return s
}

func (s *CreateAlertStrategyRequest) SetStrategy(v *CreateAlertStrategyRequestStrategy) *CreateAlertStrategyRequest {
	s.Strategy = v
	return s
}

func (s *CreateAlertStrategyRequest) SetXSysomInvokeSource(v string) *CreateAlertStrategyRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *CreateAlertStrategyRequest) Validate() error {
	if s.Strategy != nil {
		if err := s.Strategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAlertStrategyRequestStrategy struct {
	// The collection of clusters for which alerts are received.
	Clusters []*string `json:"clusters,omitempty" xml:"clusters,omitempty" type:"Repeated"`
	// The alert contacts.
	Destinations []*int32 `json:"destinations,omitempty" xml:"destinations,omitempty" type:"Repeated"`
	// The collection of anomaly items for which alerts are received.
	Items []*string `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s CreateAlertStrategyRequestStrategy) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertStrategyRequestStrategy) GoString() string {
	return s.String()
}

func (s *CreateAlertStrategyRequestStrategy) GetClusters() []*string {
	return s.Clusters
}

func (s *CreateAlertStrategyRequestStrategy) GetDestinations() []*int32 {
	return s.Destinations
}

func (s *CreateAlertStrategyRequestStrategy) GetItems() []*string {
	return s.Items
}

func (s *CreateAlertStrategyRequestStrategy) SetClusters(v []*string) *CreateAlertStrategyRequestStrategy {
	s.Clusters = v
	return s
}

func (s *CreateAlertStrategyRequestStrategy) SetDestinations(v []*int32) *CreateAlertStrategyRequestStrategy {
	s.Destinations = v
	return s
}

func (s *CreateAlertStrategyRequestStrategy) SetItems(v []*string) *CreateAlertStrategyRequestStrategy {
	s.Items = v
	return s
}

func (s *CreateAlertStrategyRequestStrategy) Validate() error {
	return dara.Validate(s)
}
