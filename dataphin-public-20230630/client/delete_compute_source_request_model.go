// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComputeSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteComputeSourceRequest
	GetId() *int64
	SetOpTenantId(v int64) *DeleteComputeSourceRequest
	GetOpTenantId() *int64
}

type DeleteComputeSourceRequest struct {
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

func (s DeleteComputeSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteComputeSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteComputeSourceRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteComputeSourceRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteComputeSourceRequest) SetId(v int64) *DeleteComputeSourceRequest {
	s.Id = &v
	return s
}

func (s *DeleteComputeSourceRequest) SetOpTenantId(v int64) *DeleteComputeSourceRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteComputeSourceRequest) Validate() error {
	return dara.Validate(s)
}
