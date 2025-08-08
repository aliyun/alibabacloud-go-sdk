// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMcdpAimRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryMcdpAimRequest
	GetAppId() *string
	SetId(v int64) *QueryMcdpAimRequest
	GetId() *int64
	SetTenantId(v string) *QueryMcdpAimRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *QueryMcdpAimRequest
	GetWorkspaceId() *string
}

type QueryMcdpAimRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMcdpAimRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMcdpAimRequest) GoString() string {
	return s.String()
}

func (s *QueryMcdpAimRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryMcdpAimRequest) GetId() *int64 {
	return s.Id
}

func (s *QueryMcdpAimRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMcdpAimRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMcdpAimRequest) SetAppId(v string) *QueryMcdpAimRequest {
	s.AppId = &v
	return s
}

func (s *QueryMcdpAimRequest) SetId(v int64) *QueryMcdpAimRequest {
	s.Id = &v
	return s
}

func (s *QueryMcdpAimRequest) SetTenantId(v string) *QueryMcdpAimRequest {
	s.TenantId = &v
	return s
}

func (s *QueryMcdpAimRequest) SetWorkspaceId(v string) *QueryMcdpAimRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMcdpAimRequest) Validate() error {
	return dara.Validate(s)
}
