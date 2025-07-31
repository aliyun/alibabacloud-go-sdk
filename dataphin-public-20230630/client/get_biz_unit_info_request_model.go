// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBizUnitInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetBizUnitInfoRequest
	GetId() *int64
	SetOpTenantId(v int64) *GetBizUnitInfoRequest
	GetOpTenantId() *int64
}

type GetBizUnitInfoRequest struct {
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
}

func (s GetBizUnitInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBizUnitInfoRequest) GoString() string {
	return s.String()
}

func (s *GetBizUnitInfoRequest) GetId() *int64 {
	return s.Id
}

func (s *GetBizUnitInfoRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetBizUnitInfoRequest) SetId(v int64) *GetBizUnitInfoRequest {
	s.Id = &v
	return s
}

func (s *GetBizUnitInfoRequest) SetOpTenantId(v int64) *GetBizUnitInfoRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetBizUnitInfoRequest) Validate() error {
	return dara.Validate(s)
}
