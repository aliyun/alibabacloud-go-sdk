// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceServersForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListResourceServersForUserRequestFilter) *ListResourceServersForUserRequest
	GetFilter() []*ListResourceServersForUserRequestFilter
	SetInstanceId(v string) *ListResourceServersForUserRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListResourceServersForUserRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceServersForUserRequest
	GetNextToken() *string
	SetUserId(v string) *ListResourceServersForUserRequest
	GetUserId() *string
}

type ListResourceServersForUserRequest struct {
	Filter []*ListResourceServersForUserRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
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
	// 用户ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// user_mkv7rgt4d7i4u7zqtzev2mxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListResourceServersForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceServersForUserRequest) GoString() string {
	return s.String()
}

func (s *ListResourceServersForUserRequest) GetFilter() []*ListResourceServersForUserRequestFilter {
	return s.Filter
}

func (s *ListResourceServersForUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListResourceServersForUserRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceServersForUserRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceServersForUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListResourceServersForUserRequest) SetFilter(v []*ListResourceServersForUserRequestFilter) *ListResourceServersForUserRequest {
	s.Filter = v
	return s
}

func (s *ListResourceServersForUserRequest) SetInstanceId(v string) *ListResourceServersForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListResourceServersForUserRequest) SetMaxResults(v int32) *ListResourceServersForUserRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceServersForUserRequest) SetNextToken(v string) *ListResourceServersForUserRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceServersForUserRequest) SetUserId(v string) *ListResourceServersForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListResourceServersForUserRequest) Validate() error {
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

type ListResourceServersForUserRequestFilter struct {
	// example:
	//
	// ApplicationIds
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListResourceServersForUserRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListResourceServersForUserRequestFilter) GoString() string {
	return s.String()
}

func (s *ListResourceServersForUserRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListResourceServersForUserRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListResourceServersForUserRequestFilter) SetName(v string) *ListResourceServersForUserRequestFilter {
	s.Name = &v
	return s
}

func (s *ListResourceServersForUserRequestFilter) SetValue(v []*string) *ListResourceServersForUserRequestFilter {
	s.Value = v
	return s
}

func (s *ListResourceServersForUserRequestFilter) Validate() error {
	return dara.Validate(s)
}
