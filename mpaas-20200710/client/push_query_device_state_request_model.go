// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushQueryDeviceStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *PushQueryDeviceStateRequest
	GetAppId() *string
	SetTarget(v string) *PushQueryDeviceStateRequest
	GetTarget() *string
	SetTargetType(v int32) *PushQueryDeviceStateRequest
	GetTargetType() *int32
	SetTenantId(v string) *PushQueryDeviceStateRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *PushQueryDeviceStateRequest
	GetWorkspaceId() *string
}

type PushQueryDeviceStateRequest struct {
	// example:
	//
	// ALIPUB9A63274111812
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 8985d1b78d135e10dc26703379369879
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// example:
	//
	// 2
	TargetType *int32 `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// example:
	//
	// BJUVXFNW
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// default
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PushQueryDeviceStateRequest) String() string {
	return dara.Prettify(s)
}

func (s PushQueryDeviceStateRequest) GoString() string {
	return s.String()
}

func (s *PushQueryDeviceStateRequest) GetAppId() *string {
	return s.AppId
}

func (s *PushQueryDeviceStateRequest) GetTarget() *string {
	return s.Target
}

func (s *PushQueryDeviceStateRequest) GetTargetType() *int32 {
	return s.TargetType
}

func (s *PushQueryDeviceStateRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PushQueryDeviceStateRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PushQueryDeviceStateRequest) SetAppId(v string) *PushQueryDeviceStateRequest {
	s.AppId = &v
	return s
}

func (s *PushQueryDeviceStateRequest) SetTarget(v string) *PushQueryDeviceStateRequest {
	s.Target = &v
	return s
}

func (s *PushQueryDeviceStateRequest) SetTargetType(v int32) *PushQueryDeviceStateRequest {
	s.TargetType = &v
	return s
}

func (s *PushQueryDeviceStateRequest) SetTenantId(v string) *PushQueryDeviceStateRequest {
	s.TenantId = &v
	return s
}

func (s *PushQueryDeviceStateRequest) SetWorkspaceId(v string) *PushQueryDeviceStateRequest {
	s.WorkspaceId = &v
	return s
}

func (s *PushQueryDeviceStateRequest) Validate() error {
	return dara.Validate(s)
}
