// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMcdpZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryMcdpZoneRequest
	GetAppId() *string
	SetId(v int64) *QueryMcdpZoneRequest
	GetId() *int64
	SetTenantId(v string) *QueryMcdpZoneRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *QueryMcdpZoneRequest
	GetWorkspaceId() *string
}

type QueryMcdpZoneRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMcdpZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMcdpZoneRequest) GoString() string {
	return s.String()
}

func (s *QueryMcdpZoneRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryMcdpZoneRequest) GetId() *int64 {
	return s.Id
}

func (s *QueryMcdpZoneRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMcdpZoneRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMcdpZoneRequest) SetAppId(v string) *QueryMcdpZoneRequest {
	s.AppId = &v
	return s
}

func (s *QueryMcdpZoneRequest) SetId(v int64) *QueryMcdpZoneRequest {
	s.Id = &v
	return s
}

func (s *QueryMcdpZoneRequest) SetTenantId(v string) *QueryMcdpZoneRequest {
	s.TenantId = &v
	return s
}

func (s *QueryMcdpZoneRequest) SetWorkspaceId(v string) *QueryMcdpZoneRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMcdpZoneRequest) Validate() error {
	return dara.Validate(s)
}
