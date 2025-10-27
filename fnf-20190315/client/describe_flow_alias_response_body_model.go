// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowAliasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v *DescribeFlowAliasResponseBodyAlias) *DescribeFlowAliasResponseBody
	GetAlias() *DescribeFlowAliasResponseBodyAlias
	SetRequestId(v string) *DescribeFlowAliasResponseBody
	GetRequestId() *string
}

type DescribeFlowAliasResponseBody struct {
	Alias *DescribeFlowAliasResponseBodyAlias `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 294D68C1-5108-5971-853A-1A9CC87B4816
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFlowAliasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowAliasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowAliasResponseBody) GetAlias() *DescribeFlowAliasResponseBodyAlias {
	return s.Alias
}

func (s *DescribeFlowAliasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFlowAliasResponseBody) SetAlias(v *DescribeFlowAliasResponseBodyAlias) *DescribeFlowAliasResponseBody {
	s.Alias = v
	return s
}

func (s *DescribeFlowAliasResponseBody) SetRequestId(v string) *DescribeFlowAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowAliasResponseBody) Validate() error {
	if s.Alias != nil {
		if err := s.Alias.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFlowAliasResponseBodyAlias struct {
	// example:
	//
	// 2024-04-22T06:09:39.907Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// alias description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// alias-name
	Name                  *string                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	RoutingConfigurations []*DescribeFlowAliasResponseBodyAliasRoutingConfigurations `json:"RoutingConfigurations,omitempty" xml:"RoutingConfigurations,omitempty" type:"Repeated"`
}

func (s DescribeFlowAliasResponseBodyAlias) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowAliasResponseBodyAlias) GoString() string {
	return s.String()
}

func (s *DescribeFlowAliasResponseBodyAlias) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeFlowAliasResponseBodyAlias) GetDescription() *string {
	return s.Description
}

func (s *DescribeFlowAliasResponseBodyAlias) GetName() *string {
	return s.Name
}

func (s *DescribeFlowAliasResponseBodyAlias) GetRoutingConfigurations() []*DescribeFlowAliasResponseBodyAliasRoutingConfigurations {
	return s.RoutingConfigurations
}

func (s *DescribeFlowAliasResponseBodyAlias) SetCreatedTime(v string) *DescribeFlowAliasResponseBodyAlias {
	s.CreatedTime = &v
	return s
}

func (s *DescribeFlowAliasResponseBodyAlias) SetDescription(v string) *DescribeFlowAliasResponseBodyAlias {
	s.Description = &v
	return s
}

func (s *DescribeFlowAliasResponseBodyAlias) SetName(v string) *DescribeFlowAliasResponseBodyAlias {
	s.Name = &v
	return s
}

func (s *DescribeFlowAliasResponseBodyAlias) SetRoutingConfigurations(v []*DescribeFlowAliasResponseBodyAliasRoutingConfigurations) *DescribeFlowAliasResponseBodyAlias {
	s.RoutingConfigurations = v
	return s
}

func (s *DescribeFlowAliasResponseBodyAlias) Validate() error {
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

type DescribeFlowAliasResponseBodyAliasRoutingConfigurations struct {
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// 10
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeFlowAliasResponseBodyAliasRoutingConfigurations) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowAliasResponseBodyAliasRoutingConfigurations) GoString() string {
	return s.String()
}

func (s *DescribeFlowAliasResponseBodyAliasRoutingConfigurations) GetVersion() *string {
	return s.Version
}

func (s *DescribeFlowAliasResponseBodyAliasRoutingConfigurations) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeFlowAliasResponseBodyAliasRoutingConfigurations) SetVersion(v string) *DescribeFlowAliasResponseBodyAliasRoutingConfigurations {
	s.Version = &v
	return s
}

func (s *DescribeFlowAliasResponseBodyAliasRoutingConfigurations) SetWeight(v int32) *DescribeFlowAliasResponseBodyAliasRoutingConfigurations {
	s.Weight = &v
	return s
}

func (s *DescribeFlowAliasResponseBodyAliasRoutingConfigurations) Validate() error {
	return dara.Validate(s)
}
