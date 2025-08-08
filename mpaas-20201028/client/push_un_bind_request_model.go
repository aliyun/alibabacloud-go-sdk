// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushUnBindRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *PushUnBindRequest
	GetAppId() *string
	SetDeliveryToken(v string) *PushUnBindRequest
	GetDeliveryToken() *string
	SetTenantId(v string) *PushUnBindRequest
	GetTenantId() *string
	SetUserId(v string) *PushUnBindRequest
	GetUserId() *string
	SetWorkspaceId(v string) *PushUnBindRequest
	GetWorkspaceId() *string
}

type PushUnBindRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	DeliveryToken *string `json:"DeliveryToken,omitempty" xml:"DeliveryToken,omitempty"`
	TenantId      *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PushUnBindRequest) String() string {
	return dara.Prettify(s)
}

func (s PushUnBindRequest) GoString() string {
	return s.String()
}

func (s *PushUnBindRequest) GetAppId() *string {
	return s.AppId
}

func (s *PushUnBindRequest) GetDeliveryToken() *string {
	return s.DeliveryToken
}

func (s *PushUnBindRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PushUnBindRequest) GetUserId() *string {
	return s.UserId
}

func (s *PushUnBindRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PushUnBindRequest) SetAppId(v string) *PushUnBindRequest {
	s.AppId = &v
	return s
}

func (s *PushUnBindRequest) SetDeliveryToken(v string) *PushUnBindRequest {
	s.DeliveryToken = &v
	return s
}

func (s *PushUnBindRequest) SetTenantId(v string) *PushUnBindRequest {
	s.TenantId = &v
	return s
}

func (s *PushUnBindRequest) SetUserId(v string) *PushUnBindRequest {
	s.UserId = &v
	return s
}

func (s *PushUnBindRequest) SetWorkspaceId(v string) *PushUnBindRequest {
	s.WorkspaceId = &v
	return s
}

func (s *PushUnBindRequest) Validate() error {
	return dara.Validate(s)
}
