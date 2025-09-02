// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceApiAuthorityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v int64) *CreateDataServiceApiAuthorityRequest
	GetApiId() *int64
	SetAuthorizedProjectId(v int64) *CreateDataServiceApiAuthorityRequest
	GetAuthorizedProjectId() *int64
	SetEndTime(v int64) *CreateDataServiceApiAuthorityRequest
	GetEndTime() *int64
	SetProjectId(v int64) *CreateDataServiceApiAuthorityRequest
	GetProjectId() *int64
	SetTenantId(v int64) *CreateDataServiceApiAuthorityRequest
	GetTenantId() *int64
}

type CreateDataServiceApiAuthorityRequest struct {
	// The ID of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The ID of the workspace to which the access permissions on the API are granted.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10002
	AuthorizedProjectId *int64 `json:"AuthorizedProjectId,omitempty" xml:"AuthorizedProjectId,omitempty"`
	// The end time of the validity period of the access permissions. The time must be a UNIX timestamp. Unit: seconds. Example: 1600531564, which indicates 2020-09-20 00:06:04 (UTC+8).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1600531564
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10003
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1004
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CreateDataServiceApiAuthorityRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceApiAuthorityRequest) GoString() string {
	return s.String()
}

func (s *CreateDataServiceApiAuthorityRequest) GetApiId() *int64 {
	return s.ApiId
}

func (s *CreateDataServiceApiAuthorityRequest) GetAuthorizedProjectId() *int64 {
	return s.AuthorizedProjectId
}

func (s *CreateDataServiceApiAuthorityRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateDataServiceApiAuthorityRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataServiceApiAuthorityRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *CreateDataServiceApiAuthorityRequest) SetApiId(v int64) *CreateDataServiceApiAuthorityRequest {
	s.ApiId = &v
	return s
}

func (s *CreateDataServiceApiAuthorityRequest) SetAuthorizedProjectId(v int64) *CreateDataServiceApiAuthorityRequest {
	s.AuthorizedProjectId = &v
	return s
}

func (s *CreateDataServiceApiAuthorityRequest) SetEndTime(v int64) *CreateDataServiceApiAuthorityRequest {
	s.EndTime = &v
	return s
}

func (s *CreateDataServiceApiAuthorityRequest) SetProjectId(v int64) *CreateDataServiceApiAuthorityRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataServiceApiAuthorityRequest) SetTenantId(v int64) *CreateDataServiceApiAuthorityRequest {
	s.TenantId = &v
	return s
}

func (s *CreateDataServiceApiAuthorityRequest) Validate() error {
	return dara.Validate(s)
}
