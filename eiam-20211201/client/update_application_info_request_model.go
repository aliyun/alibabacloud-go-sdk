// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpdateApplicationInfoRequest
	GetApplicationId() *string
	SetApplicationName(v string) *UpdateApplicationInfoRequest
	GetApplicationName() *string
	SetApplicationOwner(v *UpdateApplicationInfoRequestApplicationOwner) *UpdateApplicationInfoRequest
	GetApplicationOwner() *UpdateApplicationInfoRequestApplicationOwner
	SetApplicationVisibility(v []*string) *UpdateApplicationInfoRequest
	GetApplicationVisibility() []*string
	SetClientToken(v string) *UpdateApplicationInfoRequest
	GetClientToken() *string
	SetCustomFields(v []*UpdateApplicationInfoRequestCustomFields) *UpdateApplicationInfoRequest
	GetCustomFields() []*UpdateApplicationInfoRequestCustomFields
	SetInstanceId(v string) *UpdateApplicationInfoRequest
	GetInstanceId() *string
	SetLogoUrl(v string) *UpdateApplicationInfoRequest
	GetLogoUrl() *string
}

type UpdateApplicationInfoRequest struct {
	// IDaaS的应用主键id
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 应用的表示名称
	//
	// example:
	//
	// Ram Account SSO
	ApplicationName       *string                                       `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	ApplicationOwner      *UpdateApplicationInfoRequestApplicationOwner `json:"ApplicationOwner,omitempty" xml:"ApplicationOwner,omitempty" type:"Struct"`
	ApplicationVisibility []*string                                     `json:"ApplicationVisibility,omitempty" xml:"ApplicationVisibility,omitempty" type:"Repeated"`
	// example:
	//
	// client-token-example
	ClientToken  *string                                     `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CustomFields []*UpdateApplicationInfoRequestCustomFields `json:"CustomFields,omitempty" xml:"CustomFields,omitempty" type:"Repeated"`
	// IDaaS EIAM的实例id
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 应用Logo地址
	//
	// example:
	//
	// https://example.aliyuncs.com/logo.png
	LogoUrl *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
}

func (s UpdateApplicationInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationInfoRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateApplicationInfoRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *UpdateApplicationInfoRequest) GetApplicationOwner() *UpdateApplicationInfoRequestApplicationOwner {
	return s.ApplicationOwner
}

func (s *UpdateApplicationInfoRequest) GetApplicationVisibility() []*string {
	return s.ApplicationVisibility
}

func (s *UpdateApplicationInfoRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateApplicationInfoRequest) GetCustomFields() []*UpdateApplicationInfoRequestCustomFields {
	return s.CustomFields
}

func (s *UpdateApplicationInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateApplicationInfoRequest) GetLogoUrl() *string {
	return s.LogoUrl
}

func (s *UpdateApplicationInfoRequest) SetApplicationId(v string) *UpdateApplicationInfoRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationInfoRequest) SetApplicationName(v string) *UpdateApplicationInfoRequest {
	s.ApplicationName = &v
	return s
}

func (s *UpdateApplicationInfoRequest) SetApplicationOwner(v *UpdateApplicationInfoRequestApplicationOwner) *UpdateApplicationInfoRequest {
	s.ApplicationOwner = v
	return s
}

func (s *UpdateApplicationInfoRequest) SetApplicationVisibility(v []*string) *UpdateApplicationInfoRequest {
	s.ApplicationVisibility = v
	return s
}

func (s *UpdateApplicationInfoRequest) SetClientToken(v string) *UpdateApplicationInfoRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateApplicationInfoRequest) SetCustomFields(v []*UpdateApplicationInfoRequestCustomFields) *UpdateApplicationInfoRequest {
	s.CustomFields = v
	return s
}

func (s *UpdateApplicationInfoRequest) SetInstanceId(v string) *UpdateApplicationInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateApplicationInfoRequest) SetLogoUrl(v string) *UpdateApplicationInfoRequest {
	s.LogoUrl = &v
	return s
}

func (s *UpdateApplicationInfoRequest) Validate() error {
	if s.ApplicationOwner != nil {
		if err := s.ApplicationOwner.Validate(); err != nil {
			return err
		}
	}
	if s.CustomFields != nil {
		for _, item := range s.CustomFields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApplicationInfoRequestApplicationOwner struct {
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	UserIds  []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s UpdateApplicationInfoRequestApplicationOwner) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationInfoRequestApplicationOwner) GoString() string {
	return s.String()
}

func (s *UpdateApplicationInfoRequestApplicationOwner) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *UpdateApplicationInfoRequestApplicationOwner) GetUserIds() []*string {
	return s.UserIds
}

func (s *UpdateApplicationInfoRequestApplicationOwner) SetGroupIds(v []*string) *UpdateApplicationInfoRequestApplicationOwner {
	s.GroupIds = v
	return s
}

func (s *UpdateApplicationInfoRequestApplicationOwner) SetUserIds(v []*string) *UpdateApplicationInfoRequestApplicationOwner {
	s.UserIds = v
	return s
}

func (s *UpdateApplicationInfoRequestApplicationOwner) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationInfoRequestCustomFields struct {
	FieldName  *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	Operation  *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
}

func (s UpdateApplicationInfoRequestCustomFields) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationInfoRequestCustomFields) GoString() string {
	return s.String()
}

func (s *UpdateApplicationInfoRequestCustomFields) GetFieldName() *string {
	return s.FieldName
}

func (s *UpdateApplicationInfoRequestCustomFields) GetFieldValue() *string {
	return s.FieldValue
}

func (s *UpdateApplicationInfoRequestCustomFields) GetOperation() *string {
	return s.Operation
}

func (s *UpdateApplicationInfoRequestCustomFields) SetFieldName(v string) *UpdateApplicationInfoRequestCustomFields {
	s.FieldName = &v
	return s
}

func (s *UpdateApplicationInfoRequestCustomFields) SetFieldValue(v string) *UpdateApplicationInfoRequestCustomFields {
	s.FieldValue = &v
	return s
}

func (s *UpdateApplicationInfoRequestCustomFields) SetOperation(v string) *UpdateApplicationInfoRequestCustomFields {
	s.Operation = &v
	return s
}

func (s *UpdateApplicationInfoRequestCustomFields) Validate() error {
	return dara.Validate(s)
}
