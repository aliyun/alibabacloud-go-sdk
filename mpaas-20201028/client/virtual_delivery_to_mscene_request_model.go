// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVirtualDeliveryToMsceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *VirtualDeliveryToMsceneRequest
	GetAppId() *string
	SetCustomId(v string) *VirtualDeliveryToMsceneRequest
	GetCustomId() *string
	SetMiniProgramId(v string) *VirtualDeliveryToMsceneRequest
	GetMiniProgramId() *string
	SetPlatformId(v string) *VirtualDeliveryToMsceneRequest
	GetPlatformId() *string
	SetTenantId(v string) *VirtualDeliveryToMsceneRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *VirtualDeliveryToMsceneRequest
	GetWorkspaceId() *string
}

type VirtualDeliveryToMsceneRequest struct {
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// test_custom_id
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	// example:
	//
	// 1000001221323232
	MiniProgramId *string `json:"MiniProgramId,omitempty" xml:"MiniProgramId,omitempty"`
	// example:
	//
	// mPaaS_Goosefish
	PlatformId *string `json:"PlatformId,omitempty" xml:"PlatformId,omitempty"`
	// example:
	//
	// IDUKCGEB
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// default
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s VirtualDeliveryToMsceneRequest) String() string {
	return dara.Prettify(s)
}

func (s VirtualDeliveryToMsceneRequest) GoString() string {
	return s.String()
}

func (s *VirtualDeliveryToMsceneRequest) GetAppId() *string {
	return s.AppId
}

func (s *VirtualDeliveryToMsceneRequest) GetCustomId() *string {
	return s.CustomId
}

func (s *VirtualDeliveryToMsceneRequest) GetMiniProgramId() *string {
	return s.MiniProgramId
}

func (s *VirtualDeliveryToMsceneRequest) GetPlatformId() *string {
	return s.PlatformId
}

func (s *VirtualDeliveryToMsceneRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *VirtualDeliveryToMsceneRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *VirtualDeliveryToMsceneRequest) SetAppId(v string) *VirtualDeliveryToMsceneRequest {
	s.AppId = &v
	return s
}

func (s *VirtualDeliveryToMsceneRequest) SetCustomId(v string) *VirtualDeliveryToMsceneRequest {
	s.CustomId = &v
	return s
}

func (s *VirtualDeliveryToMsceneRequest) SetMiniProgramId(v string) *VirtualDeliveryToMsceneRequest {
	s.MiniProgramId = &v
	return s
}

func (s *VirtualDeliveryToMsceneRequest) SetPlatformId(v string) *VirtualDeliveryToMsceneRequest {
	s.PlatformId = &v
	return s
}

func (s *VirtualDeliveryToMsceneRequest) SetTenantId(v string) *VirtualDeliveryToMsceneRequest {
	s.TenantId = &v
	return s
}

func (s *VirtualDeliveryToMsceneRequest) SetWorkspaceId(v string) *VirtualDeliveryToMsceneRequest {
	s.WorkspaceId = &v
	return s
}

func (s *VirtualDeliveryToMsceneRequest) Validate() error {
	return dara.Validate(s)
}
