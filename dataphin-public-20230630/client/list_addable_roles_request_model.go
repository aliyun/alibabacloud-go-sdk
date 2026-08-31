// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddableRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *ListAddableRolesRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListAddableRolesRequest
	GetOpUserId() *string
}

type ListAddableRolesRequest struct {
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

func (s ListAddableRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAddableRolesRequest) GoString() string {
	return s.String()
}

func (s *ListAddableRolesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListAddableRolesRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListAddableRolesRequest) SetOpTenantId(v int64) *ListAddableRolesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListAddableRolesRequest) SetOpUserId(v string) *ListAddableRolesRequest {
	s.OpUserId = &v
	return s
}

func (s *ListAddableRolesRequest) Validate() error {
	return dara.Validate(s)
}
