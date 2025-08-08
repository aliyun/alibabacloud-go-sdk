// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMpaasAppInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateMpaasAppInfoRequest
	GetAppId() *string
	SetAppName(v string) *UpdateMpaasAppInfoRequest
	GetAppName() *string
	SetIconFileUrl(v string) *UpdateMpaasAppInfoRequest
	GetIconFileUrl() *string
	SetIdentifier(v string) *UpdateMpaasAppInfoRequest
	GetIdentifier() *string
	SetOnexFlag(v bool) *UpdateMpaasAppInfoRequest
	GetOnexFlag() *bool
	SetSystemType(v string) *UpdateMpaasAppInfoRequest
	GetSystemType() *string
	SetTenantId(v string) *UpdateMpaasAppInfoRequest
	GetTenantId() *string
}

type UpdateMpaasAppInfoRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	IconFileUrl *string `json:"IconFileUrl,omitempty" xml:"IconFileUrl,omitempty"`
	Identifier  *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	OnexFlag    *bool   `json:"OnexFlag,omitempty" xml:"OnexFlag,omitempty"`
	SystemType  *string `json:"SystemType,omitempty" xml:"SystemType,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s UpdateMpaasAppInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMpaasAppInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateMpaasAppInfoRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMpaasAppInfoRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateMpaasAppInfoRequest) GetIconFileUrl() *string {
	return s.IconFileUrl
}

func (s *UpdateMpaasAppInfoRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *UpdateMpaasAppInfoRequest) GetOnexFlag() *bool {
	return s.OnexFlag
}

func (s *UpdateMpaasAppInfoRequest) GetSystemType() *string {
	return s.SystemType
}

func (s *UpdateMpaasAppInfoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateMpaasAppInfoRequest) SetAppId(v string) *UpdateMpaasAppInfoRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMpaasAppInfoRequest) SetAppName(v string) *UpdateMpaasAppInfoRequest {
	s.AppName = &v
	return s
}

func (s *UpdateMpaasAppInfoRequest) SetIconFileUrl(v string) *UpdateMpaasAppInfoRequest {
	s.IconFileUrl = &v
	return s
}

func (s *UpdateMpaasAppInfoRequest) SetIdentifier(v string) *UpdateMpaasAppInfoRequest {
	s.Identifier = &v
	return s
}

func (s *UpdateMpaasAppInfoRequest) SetOnexFlag(v bool) *UpdateMpaasAppInfoRequest {
	s.OnexFlag = &v
	return s
}

func (s *UpdateMpaasAppInfoRequest) SetSystemType(v string) *UpdateMpaasAppInfoRequest {
	s.SystemType = &v
	return s
}

func (s *UpdateMpaasAppInfoRequest) SetTenantId(v string) *UpdateMpaasAppInfoRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateMpaasAppInfoRequest) Validate() error {
	return dara.Validate(s)
}
