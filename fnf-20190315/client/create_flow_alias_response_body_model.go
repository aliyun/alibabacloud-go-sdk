// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowAliasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *CreateFlowAliasResponseBody
	GetCreatedTime() *string
	SetDescription(v string) *CreateFlowAliasResponseBody
	GetDescription() *string
	SetFlowName(v string) *CreateFlowAliasResponseBody
	GetFlowName() *string
	SetName(v string) *CreateFlowAliasResponseBody
	GetName() *string
	SetRequestId(v string) *CreateFlowAliasResponseBody
	GetRequestId() *string
	SetRoutingConfigurations(v []*CreateFlowAliasResponseBodyRoutingConfigurations) *CreateFlowAliasResponseBody
	GetRoutingConfigurations() []*CreateFlowAliasResponseBodyRoutingConfigurations
}

type CreateFlowAliasResponseBody struct {
	// example:
	//
	// 2020-01-01T01:01:01.001Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// example-flow-name
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// example:
	//
	// exampe-alias-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Id of the request
	//
	// example:
	//
	// testRequestID
	RequestId             *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoutingConfigurations []*CreateFlowAliasResponseBodyRoutingConfigurations `json:"RoutingConfigurations,omitempty" xml:"RoutingConfigurations,omitempty" type:"Repeated"`
}

func (s CreateFlowAliasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowAliasResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowAliasResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *CreateFlowAliasResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateFlowAliasResponseBody) GetFlowName() *string {
	return s.FlowName
}

func (s *CreateFlowAliasResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateFlowAliasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFlowAliasResponseBody) GetRoutingConfigurations() []*CreateFlowAliasResponseBodyRoutingConfigurations {
	return s.RoutingConfigurations
}

func (s *CreateFlowAliasResponseBody) SetCreatedTime(v string) *CreateFlowAliasResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateFlowAliasResponseBody) SetDescription(v string) *CreateFlowAliasResponseBody {
	s.Description = &v
	return s
}

func (s *CreateFlowAliasResponseBody) SetFlowName(v string) *CreateFlowAliasResponseBody {
	s.FlowName = &v
	return s
}

func (s *CreateFlowAliasResponseBody) SetName(v string) *CreateFlowAliasResponseBody {
	s.Name = &v
	return s
}

func (s *CreateFlowAliasResponseBody) SetRequestId(v string) *CreateFlowAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowAliasResponseBody) SetRoutingConfigurations(v []*CreateFlowAliasResponseBodyRoutingConfigurations) *CreateFlowAliasResponseBody {
	s.RoutingConfigurations = v
	return s
}

func (s *CreateFlowAliasResponseBody) Validate() error {
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

type CreateFlowAliasResponseBodyRoutingConfigurations struct {
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// 30
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateFlowAliasResponseBodyRoutingConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowAliasResponseBodyRoutingConfigurations) GoString() string {
	return s.String()
}

func (s *CreateFlowAliasResponseBodyRoutingConfigurations) GetVersion() *string {
	return s.Version
}

func (s *CreateFlowAliasResponseBodyRoutingConfigurations) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateFlowAliasResponseBodyRoutingConfigurations) SetVersion(v string) *CreateFlowAliasResponseBodyRoutingConfigurations {
	s.Version = &v
	return s
}

func (s *CreateFlowAliasResponseBodyRoutingConfigurations) SetWeight(v int32) *CreateFlowAliasResponseBodyRoutingConfigurations {
	s.Weight = &v
	return s
}

func (s *CreateFlowAliasResponseBodyRoutingConfigurations) Validate() error {
	return dara.Validate(s)
}
