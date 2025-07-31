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
}

type ListAddableRolesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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

func (s *ListAddableRolesRequest) SetOpTenantId(v int64) *ListAddableRolesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListAddableRolesRequest) Validate() error {
	return dara.Validate(s)
}
