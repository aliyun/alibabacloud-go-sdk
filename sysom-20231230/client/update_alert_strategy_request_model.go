// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *UpdateAlertStrategyRequest
	GetEnabled() *bool
	SetId(v int64) *UpdateAlertStrategyRequest
	GetId() *int64
	SetK8sLabel(v bool) *UpdateAlertStrategyRequest
	GetK8sLabel() *bool
	SetName(v string) *UpdateAlertStrategyRequest
	GetName() *string
	SetStrategy(v *UpdateAlertStrategyRequestStrategy) *UpdateAlertStrategyRequest
	GetStrategy() *UpdateAlertStrategyRequestStrategy
}

type UpdateAlertStrategyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id       *int64 `json:"id,omitempty" xml:"id,omitempty"`
	K8sLabel *bool  `json:"k8sLabel,omitempty" xml:"k8sLabel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// strategy1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Strategy *UpdateAlertStrategyRequestStrategy `json:"strategy,omitempty" xml:"strategy,omitempty" type:"Struct"`
}

func (s UpdateAlertStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertStrategyRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertStrategyRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateAlertStrategyRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateAlertStrategyRequest) GetK8sLabel() *bool {
	return s.K8sLabel
}

func (s *UpdateAlertStrategyRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAlertStrategyRequest) GetStrategy() *UpdateAlertStrategyRequestStrategy {
	return s.Strategy
}

func (s *UpdateAlertStrategyRequest) SetEnabled(v bool) *UpdateAlertStrategyRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateAlertStrategyRequest) SetId(v int64) *UpdateAlertStrategyRequest {
	s.Id = &v
	return s
}

func (s *UpdateAlertStrategyRequest) SetK8sLabel(v bool) *UpdateAlertStrategyRequest {
	s.K8sLabel = &v
	return s
}

func (s *UpdateAlertStrategyRequest) SetName(v string) *UpdateAlertStrategyRequest {
	s.Name = &v
	return s
}

func (s *UpdateAlertStrategyRequest) SetStrategy(v *UpdateAlertStrategyRequestStrategy) *UpdateAlertStrategyRequest {
	s.Strategy = v
	return s
}

func (s *UpdateAlertStrategyRequest) Validate() error {
	if s.Strategy != nil {
		if err := s.Strategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAlertStrategyRequestStrategy struct {
	Clusters []*string `json:"clusters,omitempty" xml:"clusters,omitempty" type:"Repeated"`
	Items    []*string `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s UpdateAlertStrategyRequestStrategy) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertStrategyRequestStrategy) GoString() string {
	return s.String()
}

func (s *UpdateAlertStrategyRequestStrategy) GetClusters() []*string {
	return s.Clusters
}

func (s *UpdateAlertStrategyRequestStrategy) GetItems() []*string {
	return s.Items
}

func (s *UpdateAlertStrategyRequestStrategy) SetClusters(v []*string) *UpdateAlertStrategyRequestStrategy {
	s.Clusters = v
	return s
}

func (s *UpdateAlertStrategyRequestStrategy) SetItems(v []*string) *UpdateAlertStrategyRequestStrategy {
	s.Items = v
	return s
}

func (s *UpdateAlertStrategyRequestStrategy) Validate() error {
	return dara.Validate(s)
}
