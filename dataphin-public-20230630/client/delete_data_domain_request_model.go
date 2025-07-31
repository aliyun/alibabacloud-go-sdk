// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizUnitId(v int64) *DeleteDataDomainRequest
	GetBizUnitId() *int64
	SetId(v int64) *DeleteDataDomainRequest
	GetId() *int64
	SetOpTenantId(v int64) *DeleteDataDomainRequest
	GetOpTenantId() *int64
}

type DeleteDataDomainRequest struct {
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
	// 1241844456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteDataDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataDomainRequest) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *DeleteDataDomainRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteDataDomainRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteDataDomainRequest) SetBizUnitId(v int64) *DeleteDataDomainRequest {
	s.BizUnitId = &v
	return s
}

func (s *DeleteDataDomainRequest) SetId(v int64) *DeleteDataDomainRequest {
	s.Id = &v
	return s
}

func (s *DeleteDataDomainRequest) SetOpTenantId(v int64) *DeleteDataDomainRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteDataDomainRequest) Validate() error {
	return dara.Validate(s)
}
