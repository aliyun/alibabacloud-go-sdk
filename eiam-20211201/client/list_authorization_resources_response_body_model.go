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
	// The list of authorization resources.
	AuthorizationResources []*ListAuthorizationResourcesResponseBodyAuthorizationResources `json:"AuthorizationResources,omitempty" xml:"AuthorizationResources,omitempty" type:"Repeated"`
	// The number of rows per page in a paging query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token returned by this call, used for the next page query.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries in the list.
	//
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
	// The resource entity ID associated with the authorization resource.
	//
	// example:
	//
	// carole_01kmek49aqxxxx
	AuthorizationResourceEntityId *string `json:"AuthorizationResourceEntityId,omitempty" xml:"AuthorizationResourceEntityId,omitempty"`
	// The resource entity type associated with the authorization resource. Valid values:
	//
	// - cloud_account_role: cloud role
	//
	// example:
	//
	// cloud_account_role
	AuthorizationResourceEntityType *string `json:"AuthorizationResourceEntityType,omitempty" xml:"AuthorizationResourceEntityType,omitempty"`
	// The authorization resource ID.
	//
	// example:
	//
	// arres_01kgh3jvt7pk093rv6giu0c0qxxxx
	AuthorizationResourceId *string `json:"AuthorizationResourceId,omitempty" xml:"AuthorizationResourceId,omitempty"`
	// The authorization rule ID.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// The cloud account ID to which the resource entity associated with the authorization resource belongs.
	//
	// example:
	//
	// ca_01kmegjc11qa1txxxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// The effective condition.
	Condition *ListAuthorizationResourcesResponseBodyAuthorizationResourcesCondition `json:"Condition,omitempty" xml:"Condition,omitempty" type:"Struct"`
	// The creation time.
	//
	// example:
	//
	// 1768789292000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1768789292000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) GetCondition() *ListAuthorizationResourcesResponseBodyAuthorizationResourcesCondition {
	return s.Condition
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) GetUpdateTime() *int64 {
	return s.UpdateTime
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

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) SetCondition(v *ListAuthorizationResourcesResponseBodyAuthorizationResourcesCondition) *ListAuthorizationResourcesResponseBodyAuthorizationResources {
	s.Condition = v
	return s
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) SetCreateTime(v int64) *ListAuthorizationResourcesResponseBodyAuthorizationResources {
	s.CreateTime = &v
	return s
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) SetInstanceId(v string) *ListAuthorizationResourcesResponseBodyAuthorizationResources {
	s.InstanceId = &v
	return s
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) SetUpdateTime(v int64) *ListAuthorizationResourcesResponseBodyAuthorizationResources {
	s.UpdateTime = &v
	return s
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResources) Validate() error {
	if s.Condition != nil {
		if err := s.Condition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAuthorizationResourcesResponseBodyAuthorizationResourcesCondition struct {
	// The effective condition when used as a credential.
	CredentialCondition *ListAuthorizationResourcesResponseBodyAuthorizationResourcesConditionCredentialCondition `json:"CredentialCondition,omitempty" xml:"CredentialCondition,omitempty" type:"Struct"`
}

func (s ListAuthorizationResourcesResponseBodyAuthorizationResourcesCondition) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationResourcesResponseBodyAuthorizationResourcesCondition) GoString() string {
	return s.String()
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResourcesCondition) GetCredentialCondition() *ListAuthorizationResourcesResponseBodyAuthorizationResourcesConditionCredentialCondition {
	return s.CredentialCondition
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResourcesCondition) SetCredentialCondition(v *ListAuthorizationResourcesResponseBodyAuthorizationResourcesConditionCredentialCondition) *ListAuthorizationResourcesResponseBodyAuthorizationResourcesCondition {
	s.CredentialCondition = v
	return s
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResourcesCondition) Validate() error {
	if s.CredentialCondition != nil {
		if err := s.CredentialCondition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAuthorizationResourcesResponseBodyAuthorizationResourcesConditionCredentialCondition struct {
	// Specifies whether same-name identity accounts are supported.
	//
	// example:
	//
	// true
	AllowSameNameIdentity *bool `json:"AllowSameNameIdentity,omitempty" xml:"AllowSameNameIdentity,omitempty"`
}

func (s ListAuthorizationResourcesResponseBodyAuthorizationResourcesConditionCredentialCondition) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationResourcesResponseBodyAuthorizationResourcesConditionCredentialCondition) GoString() string {
	return s.String()
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResourcesConditionCredentialCondition) GetAllowSameNameIdentity() *bool {
	return s.AllowSameNameIdentity
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResourcesConditionCredentialCondition) SetAllowSameNameIdentity(v bool) *ListAuthorizationResourcesResponseBodyAuthorizationResourcesConditionCredentialCondition {
	s.AllowSameNameIdentity = &v
	return s
}

func (s *ListAuthorizationResourcesResponseBodyAuthorizationResourcesConditionCredentialCondition) Validate() error {
	return dara.Validate(s)
}
