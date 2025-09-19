// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *CreateAlertStrategyRequest
	GetEnabled() *bool
	SetName(v string) *CreateAlertStrategyRequest
	GetName() *string
	SetStrategy(v *CreateAlertStrategyRequestStrategy) *CreateAlertStrategyRequest
	GetStrategy() *CreateAlertStrategyRequestStrategy
}

type CreateAlertStrategyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// strategy1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Strategy *CreateAlertStrategyRequestStrategy `json:"strategy,omitempty" xml:"strategy,omitempty" type:"Struct"`
}

func (s CreateAlertStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertStrategyRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertStrategyRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateAlertStrategyRequest) GetName() *string {
	return s.Name
}

func (s *CreateAlertStrategyRequest) GetStrategy() *CreateAlertStrategyRequestStrategy {
	return s.Strategy
}

func (s *CreateAlertStrategyRequest) SetEnabled(v bool) *CreateAlertStrategyRequest {
	s.Enabled = &v
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

func (s *CreateAlertStrategyRequest) Validate() error {
	return dara.Validate(s)
}

type CreateAlertStrategyRequestStrategy struct {
	Clusters []*string `json:"clusters,omitempty" xml:"clusters,omitempty" type:"Repeated"`
	Items    []*string `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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

func (s *CreateAlertStrategyRequestStrategy) GetItems() []*string {
	return s.Items
}

func (s *CreateAlertStrategyRequestStrategy) SetClusters(v []*string) *CreateAlertStrategyRequestStrategy {
	s.Clusters = v
	return s
}

func (s *CreateAlertStrategyRequestStrategy) SetItems(v []*string) *CreateAlertStrategyRequestStrategy {
	s.Items = v
	return s
}

func (s *CreateAlertStrategyRequestStrategy) Validate() error {
	return dara.Validate(s)
}
