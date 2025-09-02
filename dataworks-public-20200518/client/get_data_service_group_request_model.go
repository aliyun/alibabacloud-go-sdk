// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *GetDataServiceGroupRequest
	GetGroupId() *string
	SetProjectId(v int64) *GetDataServiceGroupRequest
	GetProjectId() *int64
	SetTenantId(v int64) *GetDataServiceGroupRequest
	GetTenantId() *int64
}

type GetDataServiceGroupRequest struct {
	// The business process ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ds_123abc
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tenant ID. This parameter is deprecated.
	//
	// example:
	//
	// 10002
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetDataServiceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *GetDataServiceGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetDataServiceGroupRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDataServiceGroupRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetDataServiceGroupRequest) SetGroupId(v string) *GetDataServiceGroupRequest {
	s.GroupId = &v
	return s
}

func (s *GetDataServiceGroupRequest) SetProjectId(v int64) *GetDataServiceGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDataServiceGroupRequest) SetTenantId(v int64) *GetDataServiceGroupRequest {
	s.TenantId = &v
	return s
}

func (s *GetDataServiceGroupRequest) Validate() error {
	return dara.Validate(s)
}
