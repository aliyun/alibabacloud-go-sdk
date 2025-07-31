// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApiGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetDataServiceApiGroupsRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *GetDataServiceApiGroupsRequest
	GetProjectId() *int32
}

type GetDataServiceApiGroupsRequest struct {
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

func (s GetDataServiceApiGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiGroupsRequest) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiGroupsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDataServiceApiGroupsRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *GetDataServiceApiGroupsRequest) SetOpTenantId(v int64) *GetDataServiceApiGroupsRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDataServiceApiGroupsRequest) SetProjectId(v int32) *GetDataServiceApiGroupsRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDataServiceApiGroupsRequest) Validate() error {
	return dara.Validate(s)
}
