// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAcceleratorConfirmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *UpdateAcceleratorConfirmRequest
	GetAcceleratorId() *string
	SetRegionId(v string) *UpdateAcceleratorConfirmRequest
	GetRegionId() *string
}

type UpdateAcceleratorConfirmRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateAcceleratorConfirmRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAcceleratorConfirmRequest) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorConfirmRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *UpdateAcceleratorConfirmRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAcceleratorConfirmRequest) SetAcceleratorId(v string) *UpdateAcceleratorConfirmRequest {
	s.AcceleratorId = &v
	return s
}

func (s *UpdateAcceleratorConfirmRequest) SetRegionId(v string) *UpdateAcceleratorConfirmRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAcceleratorConfirmRequest) Validate() error {
	return dara.Validate(s)
}
