// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowAliasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v *UpdateFlowAliasResponseBodyAlias) *UpdateFlowAliasResponseBody
	GetAlias() *UpdateFlowAliasResponseBodyAlias
	SetRequestId(v string) *UpdateFlowAliasResponseBody
	GetRequestId() *string
}

type UpdateFlowAliasResponseBody struct {
	Alias *UpdateFlowAliasResponseBodyAlias `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 294D68C1-5108-5971-853A-1A9CC87B4816
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFlowAliasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowAliasResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlowAliasResponseBody) GetAlias() *UpdateFlowAliasResponseBodyAlias {
	return s.Alias
}

func (s *UpdateFlowAliasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFlowAliasResponseBody) SetAlias(v *UpdateFlowAliasResponseBodyAlias) *UpdateFlowAliasResponseBody {
	s.Alias = v
	return s
}

func (s *UpdateFlowAliasResponseBody) SetRequestId(v string) *UpdateFlowAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlowAliasResponseBody) Validate() error {
	if s.Alias != nil {
		if err := s.Alias.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateFlowAliasResponseBodyAlias struct {
	// example:
	//
	// 2025-10-24T14:11:26+08:00
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// my alias description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// my-alias-name
	Name                  *string                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	RoutingConfigurations []*UpdateFlowAliasResponseBodyAliasRoutingConfigurations `json:"RoutingConfigurations,omitempty" xml:"RoutingConfigurations,omitempty" type:"Repeated"`
}

func (s UpdateFlowAliasResponseBodyAlias) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowAliasResponseBodyAlias) GoString() string {
	return s.String()
}

func (s *UpdateFlowAliasResponseBodyAlias) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *UpdateFlowAliasResponseBodyAlias) GetDescription() *string {
	return s.Description
}

func (s *UpdateFlowAliasResponseBodyAlias) GetName() *string {
	return s.Name
}

func (s *UpdateFlowAliasResponseBodyAlias) GetRoutingConfigurations() []*UpdateFlowAliasResponseBodyAliasRoutingConfigurations {
	return s.RoutingConfigurations
}

func (s *UpdateFlowAliasResponseBodyAlias) SetCreatedTime(v string) *UpdateFlowAliasResponseBodyAlias {
	s.CreatedTime = &v
	return s
}

func (s *UpdateFlowAliasResponseBodyAlias) SetDescription(v string) *UpdateFlowAliasResponseBodyAlias {
	s.Description = &v
	return s
}

func (s *UpdateFlowAliasResponseBodyAlias) SetName(v string) *UpdateFlowAliasResponseBodyAlias {
	s.Name = &v
	return s
}

func (s *UpdateFlowAliasResponseBodyAlias) SetRoutingConfigurations(v []*UpdateFlowAliasResponseBodyAliasRoutingConfigurations) *UpdateFlowAliasResponseBodyAlias {
	s.RoutingConfigurations = v
	return s
}

func (s *UpdateFlowAliasResponseBodyAlias) Validate() error {
	if s.RoutingConfigurations != nil {
		for _, item := range s.RoutingConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateFlowAliasResponseBodyAliasRoutingConfigurations struct {
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// 20
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateFlowAliasResponseBodyAliasRoutingConfigurations) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowAliasResponseBodyAliasRoutingConfigurations) GoString() string {
	return s.String()
}

func (s *UpdateFlowAliasResponseBodyAliasRoutingConfigurations) GetVersion() *string {
	return s.Version
}

func (s *UpdateFlowAliasResponseBodyAliasRoutingConfigurations) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateFlowAliasResponseBodyAliasRoutingConfigurations) SetVersion(v string) *UpdateFlowAliasResponseBodyAliasRoutingConfigurations {
	s.Version = &v
	return s
}

func (s *UpdateFlowAliasResponseBodyAliasRoutingConfigurations) SetWeight(v int32) *UpdateFlowAliasResponseBodyAliasRoutingConfigurations {
	s.Weight = &v
	return s
}

func (s *UpdateFlowAliasResponseBodyAliasRoutingConfigurations) Validate() error {
	return dara.Validate(s)
}
