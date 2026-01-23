// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndex(v int64) *GetSecurityLevelRequest
	GetIndex() *int64
	SetOpTenantId(v int64) *GetSecurityLevelRequest
	GetOpTenantId() *int64
}

type GetSecurityLevelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetSecurityLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityLevelRequest) GoString() string {
	return s.String()
}

func (s *GetSecurityLevelRequest) GetIndex() *int64 {
	return s.Index
}

func (s *GetSecurityLevelRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetSecurityLevelRequest) SetIndex(v int64) *GetSecurityLevelRequest {
	s.Index = &v
	return s
}

func (s *GetSecurityLevelRequest) SetOpTenantId(v int64) *GetSecurityLevelRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetSecurityLevelRequest) Validate() error {
	return dara.Validate(s)
}
