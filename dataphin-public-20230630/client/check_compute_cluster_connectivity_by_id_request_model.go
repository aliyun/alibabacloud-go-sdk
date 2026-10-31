// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckComputeClusterConnectivityByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CheckComputeClusterConnectivityByIdRequest
	GetId() *int64
	SetOpTenantId(v int64) *CheckComputeClusterConnectivityByIdRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CheckComputeClusterConnectivityByIdRequest
	GetOpUserId() *string
}

type CheckComputeClusterConnectivityByIdRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 102311
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The operator user ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s CheckComputeClusterConnectivityByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckComputeClusterConnectivityByIdRequest) GoString() string {
	return s.String()
}

func (s *CheckComputeClusterConnectivityByIdRequest) GetId() *int64 {
	return s.Id
}

func (s *CheckComputeClusterConnectivityByIdRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CheckComputeClusterConnectivityByIdRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CheckComputeClusterConnectivityByIdRequest) SetId(v int64) *CheckComputeClusterConnectivityByIdRequest {
	s.Id = &v
	return s
}

func (s *CheckComputeClusterConnectivityByIdRequest) SetOpTenantId(v int64) *CheckComputeClusterConnectivityByIdRequest {
	s.OpTenantId = &v
	return s
}

func (s *CheckComputeClusterConnectivityByIdRequest) SetOpUserId(v string) *CheckComputeClusterConnectivityByIdRequest {
	s.OpUserId = &v
	return s
}

func (s *CheckComputeClusterConnectivityByIdRequest) Validate() error {
	return dara.Validate(s)
}
