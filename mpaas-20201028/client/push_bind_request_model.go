// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushBindRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *PushBindRequest
	GetAppId() *string
	SetDeliveryToken(v string) *PushBindRequest
	GetDeliveryToken() *string
	SetOsType(v int32) *PushBindRequest
	GetOsType() *int32
	SetPhoneNumber(v string) *PushBindRequest
	GetPhoneNumber() *string
	SetTenantId(v string) *PushBindRequest
	GetTenantId() *string
	SetUserId(v string) *PushBindRequest
	GetUserId() *string
	SetWorkspaceId(v string) *PushBindRequest
	GetWorkspaceId() *string
}

type PushBindRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	DeliveryToken *string `json:"DeliveryToken,omitempty" xml:"DeliveryToken,omitempty"`
	// This parameter is required.
	OsType      *int32  `json:"OsType,omitempty" xml:"OsType,omitempty"`
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PushBindRequest) String() string {
	return dara.Prettify(s)
}

func (s PushBindRequest) GoString() string {
	return s.String()
}

func (s *PushBindRequest) GetAppId() *string {
	return s.AppId
}

func (s *PushBindRequest) GetDeliveryToken() *string {
	return s.DeliveryToken
}

func (s *PushBindRequest) GetOsType() *int32 {
	return s.OsType
}

func (s *PushBindRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *PushBindRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PushBindRequest) GetUserId() *string {
	return s.UserId
}

func (s *PushBindRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PushBindRequest) SetAppId(v string) *PushBindRequest {
	s.AppId = &v
	return s
}

func (s *PushBindRequest) SetDeliveryToken(v string) *PushBindRequest {
	s.DeliveryToken = &v
	return s
}

func (s *PushBindRequest) SetOsType(v int32) *PushBindRequest {
	s.OsType = &v
	return s
}

func (s *PushBindRequest) SetPhoneNumber(v string) *PushBindRequest {
	s.PhoneNumber = &v
	return s
}

func (s *PushBindRequest) SetTenantId(v string) *PushBindRequest {
	s.TenantId = &v
	return s
}

func (s *PushBindRequest) SetUserId(v string) *PushBindRequest {
	s.UserId = &v
	return s
}

func (s *PushBindRequest) SetWorkspaceId(v string) *PushBindRequest {
	s.WorkspaceId = &v
	return s
}

func (s *PushBindRequest) Validate() error {
	return dara.Validate(s)
}
