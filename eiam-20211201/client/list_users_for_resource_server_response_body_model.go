// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersForResourceServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListUsersForResourceServerResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListUsersForResourceServerResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListUsersForResourceServerResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListUsersForResourceServerResponseBody
	GetTotalCount() *int64
	SetUsers(v []*ListUsersForResourceServerResponseBodyUsers) *ListUsersForResourceServerResponseBody
	GetUsers() []*ListUsersForResourceServerResponseBodyUsers
}

type ListUsersForResourceServerResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 本次调用返回的查询凭证（Token）值，用于下一次翻页查询。
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Users      []*ListUsersForResourceServerResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListUsersForResourceServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForResourceServerResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersForResourceServerResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUsersForResourceServerResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUsersForResourceServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersForResourceServerResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUsersForResourceServerResponseBody) GetUsers() []*ListUsersForResourceServerResponseBodyUsers {
	return s.Users
}

func (s *ListUsersForResourceServerResponseBody) SetMaxResults(v int32) *ListUsersForResourceServerResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUsersForResourceServerResponseBody) SetNextToken(v string) *ListUsersForResourceServerResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUsersForResourceServerResponseBody) SetRequestId(v string) *ListUsersForResourceServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersForResourceServerResponseBody) SetTotalCount(v int64) *ListUsersForResourceServerResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersForResourceServerResponseBody) SetUsers(v []*ListUsersForResourceServerResponseBodyUsers) *ListUsersForResourceServerResponseBody {
	s.Users = v
	return s
}

func (s *ListUsersForResourceServerResponseBody) Validate() error {
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUsersForResourceServerResponseBodyUsers struct {
	// 实例唯一标识
	//
	// example:
	//
	// idaas_qsw77zl5vrllwzyrrfwbmpxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源服务Scope权限集合
	ResourceServerScopes []*ListUsersForResourceServerResponseBodyUsersResourceServerScopes `json:"ResourceServerScopes,omitempty" xml:"ResourceServerScopes,omitempty" type:"Repeated"`
	// 用户的唯一标识
	//
	// example:
	//
	// user_nbsomva32b6utec3hgi7scxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUsersForResourceServerResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForResourceServerResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersForResourceServerResponseBodyUsers) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUsersForResourceServerResponseBodyUsers) GetResourceServerScopes() []*ListUsersForResourceServerResponseBodyUsersResourceServerScopes {
	return s.ResourceServerScopes
}

func (s *ListUsersForResourceServerResponseBodyUsers) GetUserId() *string {
	return s.UserId
}

func (s *ListUsersForResourceServerResponseBodyUsers) SetInstanceId(v string) *ListUsersForResourceServerResponseBodyUsers {
	s.InstanceId = &v
	return s
}

func (s *ListUsersForResourceServerResponseBodyUsers) SetResourceServerScopes(v []*ListUsersForResourceServerResponseBodyUsersResourceServerScopes) *ListUsersForResourceServerResponseBodyUsers {
	s.ResourceServerScopes = v
	return s
}

func (s *ListUsersForResourceServerResponseBodyUsers) SetUserId(v string) *ListUsersForResourceServerResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *ListUsersForResourceServerResponseBodyUsers) Validate() error {
	if s.ResourceServerScopes != nil {
		for _, item := range s.ResourceServerScopes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUsersForResourceServerResponseBodyUsersResourceServerScopes struct {
	// ResourceServerScope唯一标识
	//
	// example:
	//
	// ress_nbte4bb3qqqnaq73rlmkqixxxx
	ResourceServerScopeId *string `json:"ResourceServerScopeId,omitempty" xml:"ResourceServerScopeId,omitempty"`
	// ResourceServerScope名称
	//
	// example:
	//
	// Read All User
	ResourceServerScopeName *string `json:"ResourceServerScopeName,omitempty" xml:"ResourceServerScopeName,omitempty"`
}

func (s ListUsersForResourceServerResponseBodyUsersResourceServerScopes) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForResourceServerResponseBodyUsersResourceServerScopes) GoString() string {
	return s.String()
}

func (s *ListUsersForResourceServerResponseBodyUsersResourceServerScopes) GetResourceServerScopeId() *string {
	return s.ResourceServerScopeId
}

func (s *ListUsersForResourceServerResponseBodyUsersResourceServerScopes) GetResourceServerScopeName() *string {
	return s.ResourceServerScopeName
}

func (s *ListUsersForResourceServerResponseBodyUsersResourceServerScopes) SetResourceServerScopeId(v string) *ListUsersForResourceServerResponseBodyUsersResourceServerScopes {
	s.ResourceServerScopeId = &v
	return s
}

func (s *ListUsersForResourceServerResponseBodyUsersResourceServerScopes) SetResourceServerScopeName(v string) *ListUsersForResourceServerResponseBodyUsersResourceServerScopes {
	s.ResourceServerScopeName = &v
	return s
}

func (s *ListUsersForResourceServerResponseBodyUsersResourceServerScopes) Validate() error {
	return dara.Validate(s)
}
