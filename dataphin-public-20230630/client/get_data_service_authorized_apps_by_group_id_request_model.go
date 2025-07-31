// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAuthorizedAppsByGroupIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int32) *GetDataServiceAuthorizedAppsByGroupIdRequest
	GetGroupId() *int32
	SetOpTenantId(v int64) *GetDataServiceAuthorizedAppsByGroupIdRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *GetDataServiceAuthorizedAppsByGroupIdRequest
	GetProjectId() *int32
}

type GetDataServiceAuthorizedAppsByGroupIdRequest struct {
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

func (s GetDataServiceAuthorizedAppsByGroupIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAuthorizedAppsByGroupIdRequest) GoString() string {
	return s.String()
}

func (s *GetDataServiceAuthorizedAppsByGroupIdRequest) GetGroupId() *int32 {
	return s.GroupId
}

func (s *GetDataServiceAuthorizedAppsByGroupIdRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDataServiceAuthorizedAppsByGroupIdRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *GetDataServiceAuthorizedAppsByGroupIdRequest) SetGroupId(v int32) *GetDataServiceAuthorizedAppsByGroupIdRequest {
	s.GroupId = &v
	return s
}

func (s *GetDataServiceAuthorizedAppsByGroupIdRequest) SetOpTenantId(v int64) *GetDataServiceAuthorizedAppsByGroupIdRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDataServiceAuthorizedAppsByGroupIdRequest) SetProjectId(v int32) *GetDataServiceAuthorizedAppsByGroupIdRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDataServiceAuthorizedAppsByGroupIdRequest) Validate() error {
	return dara.Validate(s)
}
