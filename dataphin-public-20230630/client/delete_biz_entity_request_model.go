// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBizEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizUnitId(v int64) *DeleteBizEntityRequest
	GetBizUnitId() *int64
	SetId(v int64) *DeleteBizEntityRequest
	GetId() *int64
	SetOpTenantId(v int64) *DeleteBizEntityRequest
	GetOpTenantId() *int64
	SetType(v string) *DeleteBizEntityRequest
	GetType() *string
}

type DeleteBizEntityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 6798087749072704
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 101001201
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// BIZ_OBJECT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteBizEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBizEntityRequest) GoString() string {
	return s.String()
}

func (s *DeleteBizEntityRequest) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *DeleteBizEntityRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteBizEntityRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteBizEntityRequest) GetType() *string {
	return s.Type
}

func (s *DeleteBizEntityRequest) SetBizUnitId(v int64) *DeleteBizEntityRequest {
	s.BizUnitId = &v
	return s
}

func (s *DeleteBizEntityRequest) SetId(v int64) *DeleteBizEntityRequest {
	s.Id = &v
	return s
}

func (s *DeleteBizEntityRequest) SetOpTenantId(v int64) *DeleteBizEntityRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteBizEntityRequest) SetType(v string) *DeleteBizEntityRequest {
	s.Type = &v
	return s
}

func (s *DeleteBizEntityRequest) Validate() error {
	return dara.Validate(s)
}
