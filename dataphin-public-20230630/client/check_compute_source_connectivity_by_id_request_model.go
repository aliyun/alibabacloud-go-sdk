// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckComputeSourceConnectivityByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CheckComputeSourceConnectivityByIdRequest
	GetId() *int64
	SetOpTenantId(v int64) *CheckComputeSourceConnectivityByIdRequest
	GetOpTenantId() *int64
}

type CheckComputeSourceConnectivityByIdRequest struct {
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

func (s CheckComputeSourceConnectivityByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckComputeSourceConnectivityByIdRequest) GoString() string {
	return s.String()
}

func (s *CheckComputeSourceConnectivityByIdRequest) GetId() *int64 {
	return s.Id
}

func (s *CheckComputeSourceConnectivityByIdRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CheckComputeSourceConnectivityByIdRequest) SetId(v int64) *CheckComputeSourceConnectivityByIdRequest {
	s.Id = &v
	return s
}

func (s *CheckComputeSourceConnectivityByIdRequest) SetOpTenantId(v int64) *CheckComputeSourceConnectivityByIdRequest {
	s.OpTenantId = &v
	return s
}

func (s *CheckComputeSourceConnectivityByIdRequest) Validate() error {
	return dara.Validate(s)
}
