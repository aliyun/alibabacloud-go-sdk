// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBizUnitsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *ListBizUnitsRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListBizUnitsRequest
	GetOpUserId() *string
}

type ListBizUnitsRequest struct {
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

func (s ListBizUnitsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBizUnitsRequest) GoString() string {
	return s.String()
}

func (s *ListBizUnitsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListBizUnitsRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListBizUnitsRequest) SetOpTenantId(v int64) *ListBizUnitsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListBizUnitsRequest) SetOpUserId(v string) *ListBizUnitsRequest {
	s.OpUserId = &v
	return s
}

func (s *ListBizUnitsRequest) Validate() error {
	return dara.Validate(s)
}
