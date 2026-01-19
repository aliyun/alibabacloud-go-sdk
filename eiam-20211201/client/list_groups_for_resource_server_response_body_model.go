// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsForResourceServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroups(v []*ListGroupsForResourceServerResponseBodyGroups) *ListGroupsForResourceServerResponseBody
	GetGroups() []*ListGroupsForResourceServerResponseBodyGroups
	SetMaxResults(v int32) *ListGroupsForResourceServerResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListGroupsForResourceServerResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListGroupsForResourceServerResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListGroupsForResourceServerResponseBody
	GetTotalCount() *int64
}

type ListGroupsForResourceServerResponseBody struct {
	Groups []*ListGroupsForResourceServerResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGroupsForResourceServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForResourceServerResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsForResourceServerResponseBody) GetGroups() []*ListGroupsForResourceServerResponseBodyGroups {
	return s.Groups
}

func (s *ListGroupsForResourceServerResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListGroupsForResourceServerResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGroupsForResourceServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupsForResourceServerResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListGroupsForResourceServerResponseBody) SetGroups(v []*ListGroupsForResourceServerResponseBodyGroups) *ListGroupsForResourceServerResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupsForResourceServerResponseBody) SetMaxResults(v int32) *ListGroupsForResourceServerResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListGroupsForResourceServerResponseBody) SetNextToken(v string) *ListGroupsForResourceServerResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGroupsForResourceServerResponseBody) SetRequestId(v string) *ListGroupsForResourceServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsForResourceServerResponseBody) SetTotalCount(v int64) *ListGroupsForResourceServerResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListGroupsForResourceServerResponseBody) Validate() error {
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGroupsForResourceServerResponseBodyGroups struct {
	// 用户组的唯一标识
	//
	// example:
	//
	// group_nbsomva32b6utec3hgi7scxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// 实例唯一标识
	//
	// example:
	//
	// idaas_qsw77zl5vrllwzyrrfwbmpxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 资源服务Scope权限集合
	ResourceServerScopes []*ListGroupsForResourceServerResponseBodyGroupsResourceServerScopes `json:"ResourceServerScopes,omitempty" xml:"ResourceServerScopes,omitempty" type:"Repeated"`
}

func (s ListGroupsForResourceServerResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForResourceServerResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsForResourceServerResponseBodyGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *ListGroupsForResourceServerResponseBodyGroups) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListGroupsForResourceServerResponseBodyGroups) GetResourceServerScopes() []*ListGroupsForResourceServerResponseBodyGroupsResourceServerScopes {
	return s.ResourceServerScopes
}

func (s *ListGroupsForResourceServerResponseBodyGroups) SetGroupId(v string) *ListGroupsForResourceServerResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *ListGroupsForResourceServerResponseBodyGroups) SetInstanceId(v string) *ListGroupsForResourceServerResponseBodyGroups {
	s.InstanceId = &v
	return s
}

func (s *ListGroupsForResourceServerResponseBodyGroups) SetResourceServerScopes(v []*ListGroupsForResourceServerResponseBodyGroupsResourceServerScopes) *ListGroupsForResourceServerResponseBodyGroups {
	s.ResourceServerScopes = v
	return s
}

func (s *ListGroupsForResourceServerResponseBodyGroups) Validate() error {
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

type ListGroupsForResourceServerResponseBodyGroupsResourceServerScopes struct {
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
	// 获取资源信息
	ResourceServerScopeName *string `json:"ResourceServerScopeName,omitempty" xml:"ResourceServerScopeName,omitempty"`
}

func (s ListGroupsForResourceServerResponseBodyGroupsResourceServerScopes) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForResourceServerResponseBodyGroupsResourceServerScopes) GoString() string {
	return s.String()
}

func (s *ListGroupsForResourceServerResponseBodyGroupsResourceServerScopes) GetResourceServerScopeId() *string {
	return s.ResourceServerScopeId
}

func (s *ListGroupsForResourceServerResponseBodyGroupsResourceServerScopes) GetResourceServerScopeName() *string {
	return s.ResourceServerScopeName
}

func (s *ListGroupsForResourceServerResponseBodyGroupsResourceServerScopes) SetResourceServerScopeId(v string) *ListGroupsForResourceServerResponseBodyGroupsResourceServerScopes {
	s.ResourceServerScopeId = &v
	return s
}

func (s *ListGroupsForResourceServerResponseBodyGroupsResourceServerScopes) SetResourceServerScopeName(v string) *ListGroupsForResourceServerResponseBodyGroupsResourceServerScopes {
	s.ResourceServerScopeName = &v
	return s
}

func (s *ListGroupsForResourceServerResponseBodyGroupsResourceServerScopes) Validate() error {
	return dara.Validate(s)
}
