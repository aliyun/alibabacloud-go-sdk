// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceProjectAddableUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetDataServiceProjectAddableUsersRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *GetDataServiceProjectAddableUsersRequest
	GetProjectId() *int32
}

type GetDataServiceProjectAddableUsersRequest struct {
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

func (s GetDataServiceProjectAddableUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceProjectAddableUsersRequest) GoString() string {
	return s.String()
}

func (s *GetDataServiceProjectAddableUsersRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDataServiceProjectAddableUsersRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *GetDataServiceProjectAddableUsersRequest) SetOpTenantId(v int64) *GetDataServiceProjectAddableUsersRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDataServiceProjectAddableUsersRequest) SetProjectId(v int32) *GetDataServiceProjectAddableUsersRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDataServiceProjectAddableUsersRequest) Validate() error {
	return dara.Validate(s)
}
