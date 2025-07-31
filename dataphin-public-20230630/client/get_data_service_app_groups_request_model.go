// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAppGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetDataServiceAppGroupsRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *GetDataServiceAppGroupsRequest
	GetProjectId() *int32
}

type GetDataServiceAppGroupsRequest struct {
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

func (s GetDataServiceAppGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppGroupsRequest) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppGroupsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDataServiceAppGroupsRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *GetDataServiceAppGroupsRequest) SetOpTenantId(v int64) *GetDataServiceAppGroupsRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDataServiceAppGroupsRequest) SetProjectId(v int32) *GetDataServiceAppGroupsRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDataServiceAppGroupsRequest) Validate() error {
	return dara.Validate(s)
}
