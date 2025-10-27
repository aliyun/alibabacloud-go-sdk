// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowAliasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliases(v []*ListFlowAliasesResponseBodyAliases) *ListFlowAliasesResponseBody
	GetAliases() []*ListFlowAliasesResponseBodyAliases
	SetNextToken(v string) *ListFlowAliasesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListFlowAliasesResponseBody
	GetRequestId() *string
}

type ListFlowAliasesResponseBody struct {
	Aliases []*ListFlowAliasesResponseBodyAliases `json:"Aliases,omitempty" xml:"Aliases,omitempty" type:"Repeated"`
	// list token
	//
	// example:
	//
	// testNextToken
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3A44E113-9962-5B0B-AB92-14060EFE3164
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFlowAliasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFlowAliasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowAliasesResponseBody) GetAliases() []*ListFlowAliasesResponseBodyAliases {
	return s.Aliases
}

func (s *ListFlowAliasesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFlowAliasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFlowAliasesResponseBody) SetAliases(v []*ListFlowAliasesResponseBodyAliases) *ListFlowAliasesResponseBody {
	s.Aliases = v
	return s
}

func (s *ListFlowAliasesResponseBody) SetNextToken(v string) *ListFlowAliasesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListFlowAliasesResponseBody) SetRequestId(v string) *ListFlowAliasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowAliasesResponseBody) Validate() error {
	if s.Aliases != nil {
		for _, item := range s.Aliases {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFlowAliasesResponseBodyAliases struct {
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
	Name                  *string                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	RoutingConfigurations []*ListFlowAliasesResponseBodyAliasesRoutingConfigurations `json:"RoutingConfigurations,omitempty" xml:"RoutingConfigurations,omitempty" type:"Repeated"`
}

func (s ListFlowAliasesResponseBodyAliases) String() string {
	return dara.Prettify(s)
}

func (s ListFlowAliasesResponseBodyAliases) GoString() string {
	return s.String()
}

func (s *ListFlowAliasesResponseBodyAliases) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListFlowAliasesResponseBodyAliases) GetDescription() *string {
	return s.Description
}

func (s *ListFlowAliasesResponseBodyAliases) GetName() *string {
	return s.Name
}

func (s *ListFlowAliasesResponseBodyAliases) GetRoutingConfigurations() []*ListFlowAliasesResponseBodyAliasesRoutingConfigurations {
	return s.RoutingConfigurations
}

func (s *ListFlowAliasesResponseBodyAliases) SetCreatedTime(v string) *ListFlowAliasesResponseBodyAliases {
	s.CreatedTime = &v
	return s
}

func (s *ListFlowAliasesResponseBodyAliases) SetDescription(v string) *ListFlowAliasesResponseBodyAliases {
	s.Description = &v
	return s
}

func (s *ListFlowAliasesResponseBodyAliases) SetName(v string) *ListFlowAliasesResponseBodyAliases {
	s.Name = &v
	return s
}

func (s *ListFlowAliasesResponseBodyAliases) SetRoutingConfigurations(v []*ListFlowAliasesResponseBodyAliasesRoutingConfigurations) *ListFlowAliasesResponseBodyAliases {
	s.RoutingConfigurations = v
	return s
}

func (s *ListFlowAliasesResponseBodyAliases) Validate() error {
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

type ListFlowAliasesResponseBodyAliasesRoutingConfigurations struct {
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// 20
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListFlowAliasesResponseBodyAliasesRoutingConfigurations) String() string {
	return dara.Prettify(s)
}

func (s ListFlowAliasesResponseBodyAliasesRoutingConfigurations) GoString() string {
	return s.String()
}

func (s *ListFlowAliasesResponseBodyAliasesRoutingConfigurations) GetVersion() *string {
	return s.Version
}

func (s *ListFlowAliasesResponseBodyAliasesRoutingConfigurations) GetWeight() *string {
	return s.Weight
}

func (s *ListFlowAliasesResponseBodyAliasesRoutingConfigurations) SetVersion(v string) *ListFlowAliasesResponseBodyAliasesRoutingConfigurations {
	s.Version = &v
	return s
}

func (s *ListFlowAliasesResponseBodyAliasesRoutingConfigurations) SetWeight(v string) *ListFlowAliasesResponseBodyAliasesRoutingConfigurations {
	s.Weight = &v
	return s
}

func (s *ListFlowAliasesResponseBodyAliasesRoutingConfigurations) Validate() error {
	return dara.Validate(s)
}
