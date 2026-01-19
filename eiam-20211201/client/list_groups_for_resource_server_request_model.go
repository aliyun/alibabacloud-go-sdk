// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsForResourceServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListGroupsForResourceServerRequest
	GetApplicationId() *string
	SetFilter(v []*ListGroupsForResourceServerRequestFilter) *ListGroupsForResourceServerRequest
	GetFilter() []*ListGroupsForResourceServerRequestFilter
	SetInstanceId(v string) *ListGroupsForResourceServerRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListGroupsForResourceServerRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListGroupsForResourceServerRequest
	GetNextToken() *string
}

type ListGroupsForResourceServerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string                                     `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Filter        []*ListGroupsForResourceServerRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 查询凭证（Token），取值为上一次API调用返回的NextToken参数值。
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListGroupsForResourceServerRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForResourceServerRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsForResourceServerRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListGroupsForResourceServerRequest) GetFilter() []*ListGroupsForResourceServerRequestFilter {
	return s.Filter
}

func (s *ListGroupsForResourceServerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListGroupsForResourceServerRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListGroupsForResourceServerRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGroupsForResourceServerRequest) SetApplicationId(v string) *ListGroupsForResourceServerRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListGroupsForResourceServerRequest) SetFilter(v []*ListGroupsForResourceServerRequestFilter) *ListGroupsForResourceServerRequest {
	s.Filter = v
	return s
}

func (s *ListGroupsForResourceServerRequest) SetInstanceId(v string) *ListGroupsForResourceServerRequest {
	s.InstanceId = &v
	return s
}

func (s *ListGroupsForResourceServerRequest) SetMaxResults(v int32) *ListGroupsForResourceServerRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGroupsForResourceServerRequest) SetNextToken(v string) *ListGroupsForResourceServerRequest {
	s.NextToken = &v
	return s
}

func (s *ListGroupsForResourceServerRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGroupsForResourceServerRequestFilter struct {
	// example:
	//
	// GroupIds
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListGroupsForResourceServerRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForResourceServerRequestFilter) GoString() string {
	return s.String()
}

func (s *ListGroupsForResourceServerRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListGroupsForResourceServerRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListGroupsForResourceServerRequestFilter) SetName(v string) *ListGroupsForResourceServerRequestFilter {
	s.Name = &v
	return s
}

func (s *ListGroupsForResourceServerRequestFilter) SetValue(v []*string) *ListGroupsForResourceServerRequestFilter {
	s.Value = v
	return s
}

func (s *ListGroupsForResourceServerRequestFilter) Validate() error {
	return dara.Validate(s)
}
