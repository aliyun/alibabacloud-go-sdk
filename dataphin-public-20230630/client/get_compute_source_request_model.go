// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetComputeSourceRequest
	GetId() *int64
	SetOpTenantId(v int64) *GetComputeSourceRequest
	GetOpTenantId() *int64
}

type GetComputeSourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12356
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetComputeSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetComputeSourceRequest) GoString() string {
	return s.String()
}

func (s *GetComputeSourceRequest) GetId() *int64 {
	return s.Id
}

func (s *GetComputeSourceRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetComputeSourceRequest) SetId(v int64) *GetComputeSourceRequest {
	s.Id = &v
	return s
}

func (s *GetComputeSourceRequest) SetOpTenantId(v int64) *GetComputeSourceRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetComputeSourceRequest) Validate() error {
	return dara.Validate(s)
}
