// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataDomainInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetDataDomainInfoRequest
	GetId() *int64
	SetOpTenantId(v int64) *GetDataDomainInfoRequest
	GetOpTenantId() *int64
}

type GetDataDomainInfoRequest struct {
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

func (s GetDataDomainInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataDomainInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDataDomainInfoRequest) GetId() *int64 {
	return s.Id
}

func (s *GetDataDomainInfoRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDataDomainInfoRequest) SetId(v int64) *GetDataDomainInfoRequest {
	s.Id = &v
	return s
}

func (s *GetDataDomainInfoRequest) SetOpTenantId(v int64) *GetDataDomainInfoRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDataDomainInfoRequest) Validate() error {
	return dara.Validate(s)
}
