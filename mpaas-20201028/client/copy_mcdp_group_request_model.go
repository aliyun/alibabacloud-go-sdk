// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyMcdpGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CopyMcdpGroupRequest
	GetAppId() *string
	SetMpaasMappcenterMcdpGroupCopyJsonStr(v string) *CopyMcdpGroupRequest
	GetMpaasMappcenterMcdpGroupCopyJsonStr() *string
	SetTenantId(v string) *CopyMcdpGroupRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *CopyMcdpGroupRequest
	GetWorkspaceId() *string
}

type CopyMcdpGroupRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	MpaasMappcenterMcdpGroupCopyJsonStr *string `json:"MpaasMappcenterMcdpGroupCopyJsonStr,omitempty" xml:"MpaasMappcenterMcdpGroupCopyJsonStr,omitempty"`
	TenantId                            *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId                         *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CopyMcdpGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyMcdpGroupRequest) GoString() string {
	return s.String()
}

func (s *CopyMcdpGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *CopyMcdpGroupRequest) GetMpaasMappcenterMcdpGroupCopyJsonStr() *string {
	return s.MpaasMappcenterMcdpGroupCopyJsonStr
}

func (s *CopyMcdpGroupRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CopyMcdpGroupRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CopyMcdpGroupRequest) SetAppId(v string) *CopyMcdpGroupRequest {
	s.AppId = &v
	return s
}

func (s *CopyMcdpGroupRequest) SetMpaasMappcenterMcdpGroupCopyJsonStr(v string) *CopyMcdpGroupRequest {
	s.MpaasMappcenterMcdpGroupCopyJsonStr = &v
	return s
}

func (s *CopyMcdpGroupRequest) SetTenantId(v string) *CopyMcdpGroupRequest {
	s.TenantId = &v
	return s
}

func (s *CopyMcdpGroupRequest) SetWorkspaceId(v string) *CopyMcdpGroupRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CopyMcdpGroupRequest) Validate() error {
	return dara.Validate(s)
}
