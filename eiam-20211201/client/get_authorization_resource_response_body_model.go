// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorizationResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationResource(v *GetAuthorizationResourceResponseBodyAuthorizationResource) *GetAuthorizationResourceResponseBody
	GetAuthorizationResource() *GetAuthorizationResourceResponseBodyAuthorizationResource
	SetRequestId(v string) *GetAuthorizationResourceResponseBody
	GetRequestId() *string
}

type GetAuthorizationResourceResponseBody struct {
	// The authorization resource.
	AuthorizationResource *GetAuthorizationResourceResponseBodyAuthorizationResource `json:"AuthorizationResource,omitempty" xml:"AuthorizationResource,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAuthorizationResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthorizationResourceResponseBody) GetAuthorizationResource() *GetAuthorizationResourceResponseBodyAuthorizationResource {
	return s.AuthorizationResource
}

func (s *GetAuthorizationResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuthorizationResourceResponseBody) SetAuthorizationResource(v *GetAuthorizationResourceResponseBodyAuthorizationResource) *GetAuthorizationResourceResponseBody {
	s.AuthorizationResource = v
	return s
}

func (s *GetAuthorizationResourceResponseBody) SetRequestId(v string) *GetAuthorizationResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthorizationResourceResponseBody) Validate() error {
	if s.AuthorizationResource != nil {
		if err := s.AuthorizationResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAuthorizationResourceResponseBodyAuthorizationResource struct {
	// The resource entity ID associated with the authorization resource.
	//
	// example:
	//
	// carole_01kmek49aqxxxx
	AuthorizationResourceEntityId *string `json:"AuthorizationResourceEntityId,omitempty" xml:"AuthorizationResourceEntityId,omitempty"`
	// The resource entity type associated with the authorization resource. Valid values:
	//
	// - cloud_account_role: cloud role.
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
	// The condition restriction.
	Condition *GetAuthorizationResourceResponseBodyAuthorizationResourceCondition `json:"Condition,omitempty" xml:"Condition,omitempty" type:"Struct"`
	// The creation time.
	//
	// example:
	//
	// 1787023451494
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
	// 1787023451494
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetAuthorizationResourceResponseBodyAuthorizationResource) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationResourceResponseBodyAuthorizationResource) GoString() string {
	return s.String()
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) GetAuthorizationResourceEntityId() *string {
	return s.AuthorizationResourceEntityId
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) GetAuthorizationResourceEntityType() *string {
	return s.AuthorizationResourceEntityType
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) GetAuthorizationResourceId() *string {
	return s.AuthorizationResourceId
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) GetCondition() *GetAuthorizationResourceResponseBodyAuthorizationResourceCondition {
	return s.Condition
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) SetAuthorizationResourceEntityId(v string) *GetAuthorizationResourceResponseBodyAuthorizationResource {
	s.AuthorizationResourceEntityId = &v
	return s
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) SetAuthorizationResourceEntityType(v string) *GetAuthorizationResourceResponseBodyAuthorizationResource {
	s.AuthorizationResourceEntityType = &v
	return s
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) SetAuthorizationResourceId(v string) *GetAuthorizationResourceResponseBodyAuthorizationResource {
	s.AuthorizationResourceId = &v
	return s
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) SetAuthorizationRuleId(v string) *GetAuthorizationResourceResponseBodyAuthorizationResource {
	s.AuthorizationRuleId = &v
	return s
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) SetCloudAccountId(v string) *GetAuthorizationResourceResponseBodyAuthorizationResource {
	s.CloudAccountId = &v
	return s
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) SetCondition(v *GetAuthorizationResourceResponseBodyAuthorizationResourceCondition) *GetAuthorizationResourceResponseBodyAuthorizationResource {
	s.Condition = v
	return s
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) SetCreateTime(v int64) *GetAuthorizationResourceResponseBodyAuthorizationResource {
	s.CreateTime = &v
	return s
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) SetInstanceId(v string) *GetAuthorizationResourceResponseBodyAuthorizationResource {
	s.InstanceId = &v
	return s
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) SetUpdateTime(v int64) *GetAuthorizationResourceResponseBodyAuthorizationResource {
	s.UpdateTime = &v
	return s
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) Validate() error {
	if s.Condition != nil {
		if err := s.Condition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAuthorizationResourceResponseBodyAuthorizationResourceCondition struct {
	// The credential condition.
	CredentialCondition *GetAuthorizationResourceResponseBodyAuthorizationResourceConditionCredentialCondition `json:"CredentialCondition,omitempty" xml:"CredentialCondition,omitempty" type:"Struct"`
}

func (s GetAuthorizationResourceResponseBodyAuthorizationResourceCondition) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationResourceResponseBodyAuthorizationResourceCondition) GoString() string {
	return s.String()
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResourceCondition) GetCredentialCondition() *GetAuthorizationResourceResponseBodyAuthorizationResourceConditionCredentialCondition {
	return s.CredentialCondition
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResourceCondition) SetCredentialCondition(v *GetAuthorizationResourceResponseBodyAuthorizationResourceConditionCredentialCondition) *GetAuthorizationResourceResponseBodyAuthorizationResourceCondition {
	s.CredentialCondition = v
	return s
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResourceCondition) Validate() error {
	if s.CredentialCondition != nil {
		if err := s.CredentialCondition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAuthorizationResourceResponseBodyAuthorizationResourceConditionCredentialCondition struct {
	// Specifies whether same-name identity accounts are supported.
	AllowSameNameIdentity *bool `json:"AllowSameNameIdentity,omitempty" xml:"AllowSameNameIdentity,omitempty"`
}

func (s GetAuthorizationResourceResponseBodyAuthorizationResourceConditionCredentialCondition) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationResourceResponseBodyAuthorizationResourceConditionCredentialCondition) GoString() string {
	return s.String()
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResourceConditionCredentialCondition) GetAllowSameNameIdentity() *bool {
	return s.AllowSameNameIdentity
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResourceConditionCredentialCondition) SetAllowSameNameIdentity(v bool) *GetAuthorizationResourceResponseBodyAuthorizationResourceConditionCredentialCondition {
	s.AllowSameNameIdentity = &v
	return s
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResourceConditionCredentialCondition) Validate() error {
	return dara.Validate(s)
}
