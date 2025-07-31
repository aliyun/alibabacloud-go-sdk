// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBizUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteBizUnitRequest
	GetId() *int64
	SetOpTenantId(v int64) *DeleteBizUnitRequest
	GetOpTenantId() *int64
}

type DeleteBizUnitRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 6798087749072704
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteBizUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBizUnitRequest) GoString() string {
	return s.String()
}

func (s *DeleteBizUnitRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteBizUnitRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteBizUnitRequest) SetId(v int64) *DeleteBizUnitRequest {
	s.Id = &v
	return s
}

func (s *DeleteBizUnitRequest) SetOpTenantId(v int64) *DeleteBizUnitRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteBizUnitRequest) Validate() error {
	return dara.Validate(s)
}
