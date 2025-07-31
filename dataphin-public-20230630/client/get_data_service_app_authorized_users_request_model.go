// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAppAuthorizedUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int32) *GetDataServiceAppAuthorizedUsersRequest
	GetAppId() *int32
	SetOpTenantId(v int64) *GetDataServiceAppAuthorizedUsersRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *GetDataServiceAppAuthorizedUsersRequest
	GetProjectId() *int32
}

type GetDataServiceAppAuthorizedUsersRequest struct {
	// AppId
	//
	// This parameter is required.
	//
	// example:
	//
	// 1022
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
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

func (s GetDataServiceAppAuthorizedUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppAuthorizedUsersRequest) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppAuthorizedUsersRequest) GetAppId() *int32 {
	return s.AppId
}

func (s *GetDataServiceAppAuthorizedUsersRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDataServiceAppAuthorizedUsersRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *GetDataServiceAppAuthorizedUsersRequest) SetAppId(v int32) *GetDataServiceAppAuthorizedUsersRequest {
	s.AppId = &v
	return s
}

func (s *GetDataServiceAppAuthorizedUsersRequest) SetOpTenantId(v int64) *GetDataServiceAppAuthorizedUsersRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDataServiceAppAuthorizedUsersRequest) SetProjectId(v int32) *GetDataServiceAppAuthorizedUsersRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDataServiceAppAuthorizedUsersRequest) Validate() error {
	return dara.Validate(s)
}
