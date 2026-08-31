// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMyRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetMyRolesRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetMyRolesRequest
	GetOpUserId() *string
}

type GetMyRolesRequest struct {
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

func (s GetMyRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMyRolesRequest) GoString() string {
	return s.String()
}

func (s *GetMyRolesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetMyRolesRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetMyRolesRequest) SetOpTenantId(v int64) *GetMyRolesRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetMyRolesRequest) SetOpUserId(v string) *GetMyRolesRequest {
	s.OpUserId = &v
	return s
}

func (s *GetMyRolesRequest) Validate() error {
	return dara.Validate(s)
}
