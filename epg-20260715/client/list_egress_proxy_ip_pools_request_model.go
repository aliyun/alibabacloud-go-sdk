// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEgressProxyIpPoolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEgressProxyGatewayIds(v []*string) *ListEgressProxyIpPoolsRequest
	GetEgressProxyGatewayIds() []*string
	SetEgressProxyIpPoolIds(v []*string) *ListEgressProxyIpPoolsRequest
	GetEgressProxyIpPoolIds() []*string
	SetEgressProxyIpPoolNames(v []*string) *ListEgressProxyIpPoolsRequest
	GetEgressProxyIpPoolNames() []*string
	SetMaxResults(v int32) *ListEgressProxyIpPoolsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListEgressProxyIpPoolsRequest
	GetNextToken() *string
	SetSkip(v int32) *ListEgressProxyIpPoolsRequest
	GetSkip() *int32
	SetStatus(v string) *ListEgressProxyIpPoolsRequest
	GetStatus() *string
	SetTag(v []*ListEgressProxyIpPoolsRequestTag) *ListEgressProxyIpPoolsRequest
	GetTag() []*ListEgressProxyIpPoolsRequestTag
}

type ListEgressProxyIpPoolsRequest struct {
	EgressProxyGatewayIds  []*string `json:"EgressProxyGatewayIds,omitempty" xml:"EgressProxyGatewayIds,omitempty" type:"Repeated"`
	EgressProxyIpPoolIds   []*string `json:"EgressProxyIpPoolIds,omitempty" xml:"EgressProxyIpPoolIds,omitempty" type:"Repeated"`
	EgressProxyIpPoolNames []*string `json:"EgressProxyIpPoolNames,omitempty" xml:"EgressProxyIpPoolNames,omitempty" type:"Repeated"`
	// example:
	//
	// 12
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// a8cfe419acc68a35e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// example:
	//
	// Active
	Status *string                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag    []*ListEgressProxyIpPoolsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListEgressProxyIpPoolsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyIpPoolsRequest) GoString() string {
	return s.String()
}

func (s *ListEgressProxyIpPoolsRequest) GetEgressProxyGatewayIds() []*string {
	return s.EgressProxyGatewayIds
}

func (s *ListEgressProxyIpPoolsRequest) GetEgressProxyIpPoolIds() []*string {
	return s.EgressProxyIpPoolIds
}

func (s *ListEgressProxyIpPoolsRequest) GetEgressProxyIpPoolNames() []*string {
	return s.EgressProxyIpPoolNames
}

func (s *ListEgressProxyIpPoolsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEgressProxyIpPoolsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEgressProxyIpPoolsRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListEgressProxyIpPoolsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListEgressProxyIpPoolsRequest) GetTag() []*ListEgressProxyIpPoolsRequestTag {
	return s.Tag
}

func (s *ListEgressProxyIpPoolsRequest) SetEgressProxyGatewayIds(v []*string) *ListEgressProxyIpPoolsRequest {
	s.EgressProxyGatewayIds = v
	return s
}

func (s *ListEgressProxyIpPoolsRequest) SetEgressProxyIpPoolIds(v []*string) *ListEgressProxyIpPoolsRequest {
	s.EgressProxyIpPoolIds = v
	return s
}

func (s *ListEgressProxyIpPoolsRequest) SetEgressProxyIpPoolNames(v []*string) *ListEgressProxyIpPoolsRequest {
	s.EgressProxyIpPoolNames = v
	return s
}

func (s *ListEgressProxyIpPoolsRequest) SetMaxResults(v int32) *ListEgressProxyIpPoolsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEgressProxyIpPoolsRequest) SetNextToken(v string) *ListEgressProxyIpPoolsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEgressProxyIpPoolsRequest) SetSkip(v int32) *ListEgressProxyIpPoolsRequest {
	s.Skip = &v
	return s
}

func (s *ListEgressProxyIpPoolsRequest) SetStatus(v string) *ListEgressProxyIpPoolsRequest {
	s.Status = &v
	return s
}

func (s *ListEgressProxyIpPoolsRequest) SetTag(v []*ListEgressProxyIpPoolsRequestTag) *ListEgressProxyIpPoolsRequest {
	s.Tag = v
	return s
}

func (s *ListEgressProxyIpPoolsRequest) Validate() error {
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

type ListEgressProxyIpPoolsRequestTag struct {
	// example:
	//
	// apitest_msgkey_20260529101130_key_1780020691018
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 1770104466
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEgressProxyIpPoolsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyIpPoolsRequestTag) GoString() string {
	return s.String()
}

func (s *ListEgressProxyIpPoolsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListEgressProxyIpPoolsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListEgressProxyIpPoolsRequestTag) SetKey(v string) *ListEgressProxyIpPoolsRequestTag {
	s.Key = &v
	return s
}

func (s *ListEgressProxyIpPoolsRequestTag) SetValue(v string) *ListEgressProxyIpPoolsRequestTag {
	s.Value = &v
	return s
}

func (s *ListEgressProxyIpPoolsRequestTag) Validate() error {
	return dara.Validate(s)
}
