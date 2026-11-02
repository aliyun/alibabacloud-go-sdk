// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMiniGameQualificationByAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryMiniGameQualificationByAppRequest
	GetAppId() *string
	SetMiniProgramCode(v string) *QueryMiniGameQualificationByAppRequest
	GetMiniProgramCode() *string
	SetTenantId(v string) *QueryMiniGameQualificationByAppRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *QueryMiniGameQualificationByAppRequest
	GetWorkspaceId() *string
}

type QueryMiniGameQualificationByAppRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	MiniProgramCode *string `json:"MiniProgramCode,omitempty" xml:"MiniProgramCode,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMiniGameQualificationByAppRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMiniGameQualificationByAppRequest) GoString() string {
	return s.String()
}

func (s *QueryMiniGameQualificationByAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryMiniGameQualificationByAppRequest) GetMiniProgramCode() *string {
	return s.MiniProgramCode
}

func (s *QueryMiniGameQualificationByAppRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMiniGameQualificationByAppRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMiniGameQualificationByAppRequest) SetAppId(v string) *QueryMiniGameQualificationByAppRequest {
	s.AppId = &v
	return s
}

func (s *QueryMiniGameQualificationByAppRequest) SetMiniProgramCode(v string) *QueryMiniGameQualificationByAppRequest {
	s.MiniProgramCode = &v
	return s
}

func (s *QueryMiniGameQualificationByAppRequest) SetTenantId(v string) *QueryMiniGameQualificationByAppRequest {
	s.TenantId = &v
	return s
}

func (s *QueryMiniGameQualificationByAppRequest) SetWorkspaceId(v string) *QueryMiniGameQualificationByAppRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMiniGameQualificationByAppRequest) Validate() error {
	return dara.Validate(s)
}
