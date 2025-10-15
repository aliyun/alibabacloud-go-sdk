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
	SetApplicationVisibility(v []*string) *UpdateApplicationInfoRequest
	GetApplicationVisibility() []*string
	SetClientToken(v string) *UpdateApplicationInfoRequest
	GetClientToken() *string
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
	// This parameter is required.
	//
	// example:
	//
	// Ram Account SSO
	ApplicationName       *string   `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	ApplicationVisibility []*string `json:"ApplicationVisibility,omitempty" xml:"ApplicationVisibility,omitempty" type:"Repeated"`
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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

func (s *UpdateApplicationInfoRequest) GetApplicationVisibility() []*string {
	return s.ApplicationVisibility
}

func (s *UpdateApplicationInfoRequest) GetClientToken() *string {
	return s.ClientToken
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

func (s *UpdateApplicationInfoRequest) SetApplicationVisibility(v []*string) *UpdateApplicationInfoRequest {
	s.ApplicationVisibility = v
	return s
}

func (s *UpdateApplicationInfoRequest) SetClientToken(v string) *UpdateApplicationInfoRequest {
	s.ClientToken = &v
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
	return dara.Validate(s)
}
