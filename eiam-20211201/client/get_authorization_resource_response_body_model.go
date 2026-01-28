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
	AuthorizationResource *GetAuthorizationResourceResponseBodyAuthorizationResource `json:"AuthorizationResource,omitempty" xml:"AuthorizationResource,omitempty" type:"Struct"`
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

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) GetInstanceId() *string {
	return s.InstanceId
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

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) SetInstanceId(v string) *GetAuthorizationResourceResponseBodyAuthorizationResource {
	s.InstanceId = &v
	return s
}

func (s *GetAuthorizationResourceResponseBodyAuthorizationResource) Validate() error {
	return dara.Validate(s)
}
