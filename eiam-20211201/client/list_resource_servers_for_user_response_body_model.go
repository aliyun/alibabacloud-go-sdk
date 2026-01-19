// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceServersForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListResourceServersForUserResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceServersForUserResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListResourceServersForUserResponseBody
	GetRequestId() *string
	SetResourceServers(v []*ListResourceServersForUserResponseBodyResourceServers) *ListResourceServersForUserResponseBody
	GetResourceServers() []*ListResourceServersForUserResponseBodyResourceServers
	SetTotalCount(v int64) *ListResourceServersForUserResponseBody
	GetTotalCount() *int64
}

type ListResourceServersForUserResponseBody struct {
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
	RequestId       *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceServers []*ListResourceServersForUserResponseBodyResourceServers `json:"ResourceServers,omitempty" xml:"ResourceServers,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceServersForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceServersForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceServersForUserResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceServersForUserResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceServersForUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceServersForUserResponseBody) GetResourceServers() []*ListResourceServersForUserResponseBodyResourceServers {
	return s.ResourceServers
}

func (s *ListResourceServersForUserResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListResourceServersForUserResponseBody) SetMaxResults(v int32) *ListResourceServersForUserResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListResourceServersForUserResponseBody) SetNextToken(v string) *ListResourceServersForUserResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceServersForUserResponseBody) SetRequestId(v string) *ListResourceServersForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceServersForUserResponseBody) SetResourceServers(v []*ListResourceServersForUserResponseBodyResourceServers) *ListResourceServersForUserResponseBody {
	s.ResourceServers = v
	return s
}

func (s *ListResourceServersForUserResponseBody) SetTotalCount(v int64) *ListResourceServersForUserResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourceServersForUserResponseBody) Validate() error {
	if s.ResourceServers != nil {
		for _, item := range s.ResourceServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceServersForUserResponseBodyResourceServers struct {
	// 资源服务应用的唯一标识
	//
	// example:
	//
	// app_nbsomva32b6utec3hgi7scxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 实例唯一标识
	//
	// example:
	//
	// idaas_qsw77zl5vrllwzyrrfwbmpxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// urn:idaas:test
	ResourceServerIdentifier *string `json:"ResourceServerIdentifier,omitempty" xml:"ResourceServerIdentifier,omitempty"`
	// 资源服务Scope权限集合
	ResourceServerScopes []*ListResourceServersForUserResponseBodyResourceServersResourceServerScopes `json:"ResourceServerScopes,omitempty" xml:"ResourceServerScopes,omitempty" type:"Repeated"`
}

func (s ListResourceServersForUserResponseBodyResourceServers) String() string {
	return dara.Prettify(s)
}

func (s ListResourceServersForUserResponseBodyResourceServers) GoString() string {
	return s.String()
}

func (s *ListResourceServersForUserResponseBodyResourceServers) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListResourceServersForUserResponseBodyResourceServers) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListResourceServersForUserResponseBodyResourceServers) GetResourceServerIdentifier() *string {
	return s.ResourceServerIdentifier
}

func (s *ListResourceServersForUserResponseBodyResourceServers) GetResourceServerScopes() []*ListResourceServersForUserResponseBodyResourceServersResourceServerScopes {
	return s.ResourceServerScopes
}

func (s *ListResourceServersForUserResponseBodyResourceServers) SetApplicationId(v string) *ListResourceServersForUserResponseBodyResourceServers {
	s.ApplicationId = &v
	return s
}

func (s *ListResourceServersForUserResponseBodyResourceServers) SetInstanceId(v string) *ListResourceServersForUserResponseBodyResourceServers {
	s.InstanceId = &v
	return s
}

func (s *ListResourceServersForUserResponseBodyResourceServers) SetResourceServerIdentifier(v string) *ListResourceServersForUserResponseBodyResourceServers {
	s.ResourceServerIdentifier = &v
	return s
}

func (s *ListResourceServersForUserResponseBodyResourceServers) SetResourceServerScopes(v []*ListResourceServersForUserResponseBodyResourceServersResourceServerScopes) *ListResourceServersForUserResponseBodyResourceServers {
	s.ResourceServerScopes = v
	return s
}

func (s *ListResourceServersForUserResponseBodyResourceServers) Validate() error {
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

type ListResourceServersForUserResponseBodyResourceServersResourceServerScopes struct {
	// 直接分配给当前用户的权限，视为直接授权。
	//
	// example:
	//
	// true
	HasDirectAuthorization *bool `json:"HasDirectAuthorization,omitempty" xml:"HasDirectAuthorization,omitempty"`
	// 通过用户隶属的组织、组获取的权限，视为继承权限。
	//
	// example:
	//
	// false
	HasInheritAuthorization *bool `json:"HasInheritAuthorization,omitempty" xml:"HasInheritAuthorization,omitempty"`
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

func (s ListResourceServersForUserResponseBodyResourceServersResourceServerScopes) String() string {
	return dara.Prettify(s)
}

func (s ListResourceServersForUserResponseBodyResourceServersResourceServerScopes) GoString() string {
	return s.String()
}

func (s *ListResourceServersForUserResponseBodyResourceServersResourceServerScopes) GetHasDirectAuthorization() *bool {
	return s.HasDirectAuthorization
}

func (s *ListResourceServersForUserResponseBodyResourceServersResourceServerScopes) GetHasInheritAuthorization() *bool {
	return s.HasInheritAuthorization
}

func (s *ListResourceServersForUserResponseBodyResourceServersResourceServerScopes) GetResourceServerScopeId() *string {
	return s.ResourceServerScopeId
}

func (s *ListResourceServersForUserResponseBodyResourceServersResourceServerScopes) GetResourceServerScopeName() *string {
	return s.ResourceServerScopeName
}

func (s *ListResourceServersForUserResponseBodyResourceServersResourceServerScopes) SetHasDirectAuthorization(v bool) *ListResourceServersForUserResponseBodyResourceServersResourceServerScopes {
	s.HasDirectAuthorization = &v
	return s
}

func (s *ListResourceServersForUserResponseBodyResourceServersResourceServerScopes) SetHasInheritAuthorization(v bool) *ListResourceServersForUserResponseBodyResourceServersResourceServerScopes {
	s.HasInheritAuthorization = &v
	return s
}

func (s *ListResourceServersForUserResponseBodyResourceServersResourceServerScopes) SetResourceServerScopeId(v string) *ListResourceServersForUserResponseBodyResourceServersResourceServerScopes {
	s.ResourceServerScopeId = &v
	return s
}

func (s *ListResourceServersForUserResponseBodyResourceServersResourceServerScopes) SetResourceServerScopeName(v string) *ListResourceServersForUserResponseBodyResourceServersResourceServerScopes {
	s.ResourceServerScopeName = &v
	return s
}

func (s *ListResourceServersForUserResponseBodyResourceServersResourceServerScopes) Validate() error {
	return dara.Validate(s)
}
