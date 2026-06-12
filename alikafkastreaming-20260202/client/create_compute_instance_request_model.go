// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPaidType(v int64) *CreateComputeInstanceRequest
	GetPaidType() *int64
	SetRegionId(v string) *CreateComputeInstanceRequest
	GetRegionId() *string
}

type CreateComputeInstanceRequest struct {
	// This parameter is required.
	PaidType *int64 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateComputeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateComputeInstanceRequest) GetPaidType() *int64 {
	return s.PaidType
}

func (s *CreateComputeInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateComputeInstanceRequest) SetPaidType(v int64) *CreateComputeInstanceRequest {
	s.PaidType = &v
	return s
}

func (s *CreateComputeInstanceRequest) SetRegionId(v string) *CreateComputeInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateComputeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
