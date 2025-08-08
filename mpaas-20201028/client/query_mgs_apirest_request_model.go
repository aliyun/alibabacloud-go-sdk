// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMgsApirestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryMgsApirestRequest
	GetAppId() *string
	SetFormat(v string) *QueryMgsApirestRequest
	GetFormat() *string
	SetId(v int64) *QueryMgsApirestRequest
	GetId() *int64
	SetTenantId(v string) *QueryMgsApirestRequest
	GetTenantId() *string
	SetType(v string) *QueryMgsApirestRequest
	GetType() *string
	SetWorkspaceId(v string) *QueryMgsApirestRequest
	GetWorkspaceId() *string
}

type QueryMgsApirestRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Format      *string `json:"Format,omitempty" xml:"Format,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMgsApirestRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApirestRequest) GoString() string {
	return s.String()
}

func (s *QueryMgsApirestRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryMgsApirestRequest) GetFormat() *string {
	return s.Format
}

func (s *QueryMgsApirestRequest) GetId() *int64 {
	return s.Id
}

func (s *QueryMgsApirestRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMgsApirestRequest) GetType() *string {
	return s.Type
}

func (s *QueryMgsApirestRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMgsApirestRequest) SetAppId(v string) *QueryMgsApirestRequest {
	s.AppId = &v
	return s
}

func (s *QueryMgsApirestRequest) SetFormat(v string) *QueryMgsApirestRequest {
	s.Format = &v
	return s
}

func (s *QueryMgsApirestRequest) SetId(v int64) *QueryMgsApirestRequest {
	s.Id = &v
	return s
}

func (s *QueryMgsApirestRequest) SetTenantId(v string) *QueryMgsApirestRequest {
	s.TenantId = &v
	return s
}

func (s *QueryMgsApirestRequest) SetType(v string) *QueryMgsApirestRequest {
	s.Type = &v
	return s
}

func (s *QueryMgsApirestRequest) SetWorkspaceId(v string) *QueryMgsApirestRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMgsApirestRequest) Validate() error {
	return dara.Validate(s)
}
