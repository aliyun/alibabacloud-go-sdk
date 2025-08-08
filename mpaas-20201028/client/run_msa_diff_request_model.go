// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunMsaDiffRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RunMsaDiffRequest
	GetAppId() *string
	SetMpaasMappcenterMsaDiffRunJsonStr(v string) *RunMsaDiffRequest
	GetMpaasMappcenterMsaDiffRunJsonStr() *string
	SetTenantId(v string) *RunMsaDiffRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *RunMsaDiffRequest
	GetWorkspaceId() *string
}

type RunMsaDiffRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	MpaasMappcenterMsaDiffRunJsonStr *string `json:"MpaasMappcenterMsaDiffRunJsonStr,omitempty" xml:"MpaasMappcenterMsaDiffRunJsonStr,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunMsaDiffRequest) String() string {
	return dara.Prettify(s)
}

func (s RunMsaDiffRequest) GoString() string {
	return s.String()
}

func (s *RunMsaDiffRequest) GetAppId() *string {
	return s.AppId
}

func (s *RunMsaDiffRequest) GetMpaasMappcenterMsaDiffRunJsonStr() *string {
	return s.MpaasMappcenterMsaDiffRunJsonStr
}

func (s *RunMsaDiffRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RunMsaDiffRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunMsaDiffRequest) SetAppId(v string) *RunMsaDiffRequest {
	s.AppId = &v
	return s
}

func (s *RunMsaDiffRequest) SetMpaasMappcenterMsaDiffRunJsonStr(v string) *RunMsaDiffRequest {
	s.MpaasMappcenterMsaDiffRunJsonStr = &v
	return s
}

func (s *RunMsaDiffRequest) SetTenantId(v string) *RunMsaDiffRequest {
	s.TenantId = &v
	return s
}

func (s *RunMsaDiffRequest) SetWorkspaceId(v string) *RunMsaDiffRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunMsaDiffRequest) Validate() error {
	return dara.Validate(s)
}
