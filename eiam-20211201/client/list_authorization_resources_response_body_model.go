// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizationResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationResources(v []*ListAuthorizationResourcesResponseBodyAuthorizationResources) *ListAuthorizationResourcesResponseBody
	GetAuthorizationResources() []*ListAuthorizationResourcesResponseBodyAuthorizationResources
	SetMaxResults(v int32) *ListAuthorizationResourcesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAuthorizationResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAuthorizationResourcesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListAuthorizationResourcesResponseBody
	GetTotalCount() *int64
}

type ListAuthorizationResourcesResponseBody struct {
	AuthorizationResources []*ListAuthorizationResourcesResponseBodyAuthorizationResources `json:"AuthorizationResources,omitempty" xml:"AuthorizationResources,omitempty" type:"Repeated"`
	// 分页查询时每页行数。
	//
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

func (s ListAuthorizationResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizationResourcesResponseBody) GetAuthorizationResources() []*ListAuthorizationResourcesResponseBodyAuthorizationResources {
	return s.AuthorizationResources
}

func (s *ListAuthorizationResourcesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAuthorizationResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAuthorizationResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorizationResourcesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAuthorizationResourcesResponseBody) SetAuthorizationResources(v []*ListAuthorizationResourcesResponseBodyAuthorizationResources) *ListAuthorizationResourcesResponseBody {
	s.AuthorizationResources = v
	return s
}

func (s *ListAuthorizationResourcesResponseBody) SetMaxResults(v int32) *ListAuthorizationResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAuthorizationResourcesResponseBody) SetNextToken(v string) *ListAuthorizationResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAuthorizationResourcesResponseBody) SetRequestId(v string) *ListAuthorizationResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizationResourcesResponseBody) SetTotalCount(v int64) *ListAuthorizationResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAuthorizationResourcesResponseBody) Validate() error {
	if s.AuthorizationResources != nil {
		for _, item := range s.AuthorizationResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAuthorizationResourcesResponseBodyAuthorizationResources struct {
	// 资源实体标识
	//
	// example:
	//
	// carole_01kmek49aqxxxx
	AuthorizationResourceEntityId *string `json:"AuthorizationResourceEntityId,omitempty" xml:"AuthorizationResourceEntityId,omitempty"`
	// 资源实体类型，枚举类型：asset（资产）、credential（凭据）、cloud_identity_role（云账号角色）
	//
	// example:
	//
	// cloud_account_role
	AuthorizationResourceEntityType *string `json:"AuthorizationResourceEntityType,omitempty" xml:"AuthorizationResourceEntityType,omitempty"`
	// 授权资源标识
	//
	// example:
	//
	// arres_01kgh3jvt7pk093rv6giu0c0qxxxx
	AuthorizationResourceId *string `json:"AuthorizationResourceId,omitempty" xml:"AuthorizationResourceId,omitempty"`
	// 授权规则标识
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// 云账号ID。
	//
	// example:
	//
	// ca_01kmegjc11qa1txxxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// 实例ID
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListAuthorizationResourcesResponseBodyAuthorizationResources) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationResourcesResponseBodyAuthorizationResources) GoString() string {
	return s.String()
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) GetAuthorizationResourceEntityId() *string {
	return s.AuthorizationResourceEntityId
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) GetAuthorizationResourceEntityType() *string {
	return s.AuthorizationResourceEntityType
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) GetAuthorizationResourceId() *string {
	return s.AuthorizationResourceId
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) SetAuthorizationResourceEntityId(v string) *ListAuthorizationResourcesResponseBodyAuthorizationResources {
	s.AuthorizationResourceEntityId = &v
	return s
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) SetAuthorizationResourceEntityType(v string) *ListAuthorizationResourcesResponseBodyAuthorizationResources {
	s.AuthorizationResourceEntityType = &v
	return s
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) SetAuthorizationResourceId(v string) *ListAuthorizationResourcesResponseBodyAuthorizationResources {
	s.AuthorizationResourceId = &v
	return s
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) SetAuthorizationRuleId(v string) *ListAuthorizationResourcesResponseBodyAuthorizationResources {
	s.AuthorizationRuleId = &v
	return s
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) SetCloudAccountId(v string) *ListAuthorizationResourcesResponseBodyAuthorizationResources {
	s.CloudAccountId = &v
	return s
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) SetInstanceId(v string) *ListAuthorizationResourcesResponseBodyAuthorizationResources {
	s.InstanceId = &v
	return s
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) Validate() error {
	return dara.Validate(s)
}
