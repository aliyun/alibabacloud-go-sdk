// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAuthorizedProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetDataServiceAuthorizedProjectsRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetDataServiceAuthorizedProjectsRequest
	GetOpUserId() *string
}

type GetDataServiceAuthorizedProjectsRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s GetDataServiceAuthorizedProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAuthorizedProjectsRequest) GoString() string {
	return s.String()
}

func (s *GetDataServiceAuthorizedProjectsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDataServiceAuthorizedProjectsRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetDataServiceAuthorizedProjectsRequest) SetOpTenantId(v int64) *GetDataServiceAuthorizedProjectsRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDataServiceAuthorizedProjectsRequest) SetOpUserId(v string) *GetDataServiceAuthorizedProjectsRequest {
	s.OpUserId = &v
	return s
}

func (s *GetDataServiceAuthorizedProjectsRequest) Validate() error {
	return dara.Validate(s)
}
