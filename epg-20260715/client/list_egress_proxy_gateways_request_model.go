// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEgressProxyGatewaysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEgressProxyConfigurationIds(v []*string) *ListEgressProxyGatewaysRequest
	GetEgressProxyConfigurationIds() []*string
	SetEgressProxyGatewayIds(v []*string) *ListEgressProxyGatewaysRequest
	GetEgressProxyGatewayIds() []*string
	SetEgressProxyGatewayNames(v []*string) *ListEgressProxyGatewaysRequest
	GetEgressProxyGatewayNames() []*string
	SetMaxResults(v int32) *ListEgressProxyGatewaysRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListEgressProxyGatewaysRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListEgressProxyGatewaysRequest
	GetResourceGroupId() *string
	SetSkip(v int32) *ListEgressProxyGatewaysRequest
	GetSkip() *int32
	SetStatus(v string) *ListEgressProxyGatewaysRequest
	GetStatus() *string
	SetTag(v []*ListEgressProxyGatewaysRequestTag) *ListEgressProxyGatewaysRequest
	GetTag() []*ListEgressProxyGatewaysRequestTag
	SetVpcIds(v []*string) *ListEgressProxyGatewaysRequest
	GetVpcIds() []*string
}

type ListEgressProxyGatewaysRequest struct {
	EgressProxyConfigurationIds []*string `json:"EgressProxyConfigurationIds,omitempty" xml:"EgressProxyConfigurationIds,omitempty" type:"Repeated"`
	EgressProxyGatewayIds       []*string `json:"EgressProxyGatewayIds,omitempty" xml:"EgressProxyGatewayIds,omitempty" type:"Repeated"`
	EgressProxyGatewayNames     []*string `json:"EgressProxyGatewayNames,omitempty" xml:"EgressProxyGatewayNames,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFmyTO70tTpLG6I4FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// rg-acfm4jcyx5e****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 0
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// example:
	//
	// Active
	Status *string                              `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag    []*ListEgressProxyGatewaysRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VpcIds []*string                            `json:"VpcIds,omitempty" xml:"VpcIds,omitempty" type:"Repeated"`
}

func (s ListEgressProxyGatewaysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyGatewaysRequest) GoString() string {
	return s.String()
}

func (s *ListEgressProxyGatewaysRequest) GetEgressProxyConfigurationIds() []*string {
	return s.EgressProxyConfigurationIds
}

func (s *ListEgressProxyGatewaysRequest) GetEgressProxyGatewayIds() []*string {
	return s.EgressProxyGatewayIds
}

func (s *ListEgressProxyGatewaysRequest) GetEgressProxyGatewayNames() []*string {
	return s.EgressProxyGatewayNames
}

func (s *ListEgressProxyGatewaysRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEgressProxyGatewaysRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEgressProxyGatewaysRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListEgressProxyGatewaysRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListEgressProxyGatewaysRequest) GetStatus() *string {
	return s.Status
}

func (s *ListEgressProxyGatewaysRequest) GetTag() []*ListEgressProxyGatewaysRequestTag {
	return s.Tag
}

func (s *ListEgressProxyGatewaysRequest) GetVpcIds() []*string {
	return s.VpcIds
}

func (s *ListEgressProxyGatewaysRequest) SetEgressProxyConfigurationIds(v []*string) *ListEgressProxyGatewaysRequest {
	s.EgressProxyConfigurationIds = v
	return s
}

func (s *ListEgressProxyGatewaysRequest) SetEgressProxyGatewayIds(v []*string) *ListEgressProxyGatewaysRequest {
	s.EgressProxyGatewayIds = v
	return s
}

func (s *ListEgressProxyGatewaysRequest) SetEgressProxyGatewayNames(v []*string) *ListEgressProxyGatewaysRequest {
	s.EgressProxyGatewayNames = v
	return s
}

func (s *ListEgressProxyGatewaysRequest) SetMaxResults(v int32) *ListEgressProxyGatewaysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEgressProxyGatewaysRequest) SetNextToken(v string) *ListEgressProxyGatewaysRequest {
	s.NextToken = &v
	return s
}

func (s *ListEgressProxyGatewaysRequest) SetResourceGroupId(v string) *ListEgressProxyGatewaysRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListEgressProxyGatewaysRequest) SetSkip(v int32) *ListEgressProxyGatewaysRequest {
	s.Skip = &v
	return s
}

func (s *ListEgressProxyGatewaysRequest) SetStatus(v string) *ListEgressProxyGatewaysRequest {
	s.Status = &v
	return s
}

func (s *ListEgressProxyGatewaysRequest) SetTag(v []*ListEgressProxyGatewaysRequestTag) *ListEgressProxyGatewaysRequest {
	s.Tag = v
	return s
}

func (s *ListEgressProxyGatewaysRequest) SetVpcIds(v []*string) *ListEgressProxyGatewaysRequest {
	s.VpcIds = v
	return s
}

func (s *ListEgressProxyGatewaysRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEgressProxyGatewaysRequestTag struct {
	// example:
	//
	// epg
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEgressProxyGatewaysRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyGatewaysRequestTag) GoString() string {
	return s.String()
}

func (s *ListEgressProxyGatewaysRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListEgressProxyGatewaysRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListEgressProxyGatewaysRequestTag) SetKey(v string) *ListEgressProxyGatewaysRequestTag {
	s.Key = &v
	return s
}

func (s *ListEgressProxyGatewaysRequestTag) SetValue(v string) *ListEgressProxyGatewaysRequestTag {
	s.Value = &v
	return s
}

func (s *ListEgressProxyGatewaysRequestTag) Validate() error {
	return dara.Validate(s)
}
