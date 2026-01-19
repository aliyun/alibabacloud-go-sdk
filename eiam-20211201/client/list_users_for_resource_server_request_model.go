// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersForResourceServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListUsersForResourceServerRequest
	GetApplicationId() *string
	SetFilter(v []*ListUsersForResourceServerRequestFilter) *ListUsersForResourceServerRequest
	GetFilter() []*ListUsersForResourceServerRequestFilter
	SetInstanceId(v string) *ListUsersForResourceServerRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListUsersForResourceServerRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListUsersForResourceServerRequest
	GetNextToken() *string
}

type ListUsersForResourceServerRequest struct {
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string                                    `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Filter        []*ListUsersForResourceServerRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
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

func (s ListUsersForResourceServerRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForResourceServerRequest) GoString() string {
	return s.String()
}

func (s *ListUsersForResourceServerRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListUsersForResourceServerRequest) GetFilter() []*ListUsersForResourceServerRequestFilter {
	return s.Filter
}

func (s *ListUsersForResourceServerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUsersForResourceServerRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUsersForResourceServerRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUsersForResourceServerRequest) SetApplicationId(v string) *ListUsersForResourceServerRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListUsersForResourceServerRequest) SetFilter(v []*ListUsersForResourceServerRequestFilter) *ListUsersForResourceServerRequest {
	s.Filter = v
	return s
}

func (s *ListUsersForResourceServerRequest) SetInstanceId(v string) *ListUsersForResourceServerRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUsersForResourceServerRequest) SetMaxResults(v int32) *ListUsersForResourceServerRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUsersForResourceServerRequest) SetNextToken(v string) *ListUsersForResourceServerRequest {
	s.NextToken = &v
	return s
}

func (s *ListUsersForResourceServerRequest) Validate() error {
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

type ListUsersForResourceServerRequestFilter struct {
	// example:
	//
	// UserIds
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListUsersForResourceServerRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForResourceServerRequestFilter) GoString() string {
	return s.String()
}

func (s *ListUsersForResourceServerRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListUsersForResourceServerRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListUsersForResourceServerRequestFilter) SetName(v string) *ListUsersForResourceServerRequestFilter {
	s.Name = &v
	return s
}

func (s *ListUsersForResourceServerRequestFilter) SetValue(v []*string) *ListUsersForResourceServerRequestFilter {
	s.Value = v
	return s
}

func (s *ListUsersForResourceServerRequestFilter) Validate() error {
	return dara.Validate(s)
}
