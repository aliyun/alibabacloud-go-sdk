// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKgAuthorizedWorkspacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetKgAuthorizedWorkspacesRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetKgAuthorizedWorkspacesRequest
	GetOpUserId() *string
}

type GetKgAuthorizedWorkspacesRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30012011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s GetKgAuthorizedWorkspacesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKgAuthorizedWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *GetKgAuthorizedWorkspacesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetKgAuthorizedWorkspacesRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetKgAuthorizedWorkspacesRequest) SetOpTenantId(v int64) *GetKgAuthorizedWorkspacesRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetKgAuthorizedWorkspacesRequest) SetOpUserId(v string) *GetKgAuthorizedWorkspacesRequest {
	s.OpUserId = &v
	return s
}

func (s *GetKgAuthorizedWorkspacesRequest) Validate() error {
	return dara.Validate(s)
}
