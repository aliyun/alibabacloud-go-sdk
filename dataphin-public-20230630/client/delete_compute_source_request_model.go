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
	SetOpUserId(v string) *DeleteComputeSourceRequest
	GetOpUserId() *string
}

type DeleteComputeSourceRequest struct {
	// The compute source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12356
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s *DeleteComputeSourceRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *DeleteComputeSourceRequest) SetId(v int64) *DeleteComputeSourceRequest {
	s.Id = &v
	return s
}

func (s *DeleteComputeSourceRequest) SetOpTenantId(v int64) *DeleteComputeSourceRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteComputeSourceRequest) SetOpUserId(v string) *DeleteComputeSourceRequest {
	s.OpUserId = &v
	return s
}

func (s *DeleteComputeSourceRequest) Validate() error {
	return dara.Validate(s)
}
