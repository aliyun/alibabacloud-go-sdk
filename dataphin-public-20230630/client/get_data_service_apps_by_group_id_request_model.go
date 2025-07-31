// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAppsByGroupIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int32) *GetDataServiceAppsByGroupIdRequest
	GetGroupId() *int32
	SetOpTenantId(v int64) *GetDataServiceAppsByGroupIdRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *GetDataServiceAppsByGroupIdRequest
	GetProjectId() *int32
}

type GetDataServiceAppsByGroupIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 202102
	GroupId *int32 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 102102
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetDataServiceAppsByGroupIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppsByGroupIdRequest) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppsByGroupIdRequest) GetGroupId() *int32 {
	return s.GroupId
}

func (s *GetDataServiceAppsByGroupIdRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDataServiceAppsByGroupIdRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *GetDataServiceAppsByGroupIdRequest) SetGroupId(v int32) *GetDataServiceAppsByGroupIdRequest {
	s.GroupId = &v
	return s
}

func (s *GetDataServiceAppsByGroupIdRequest) SetOpTenantId(v int64) *GetDataServiceAppsByGroupIdRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDataServiceAppsByGroupIdRequest) SetProjectId(v int32) *GetDataServiceAppsByGroupIdRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDataServiceAppsByGroupIdRequest) Validate() error {
	return dara.Validate(s)
}
