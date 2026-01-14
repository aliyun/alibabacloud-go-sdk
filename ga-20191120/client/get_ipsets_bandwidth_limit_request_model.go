// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpsetsBandwidthLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *GetIpsetsBandwidthLimitRequest
	GetAcceleratorId() *string
	SetRegionId(v string) *GetIpsetsBandwidthLimitRequest
	GetRegionId() *string
}

type GetIpsetsBandwidthLimitRequest struct {
	// The ID of the GA instance to be queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetIpsetsBandwidthLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIpsetsBandwidthLimitRequest) GoString() string {
	return s.String()
}

func (s *GetIpsetsBandwidthLimitRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *GetIpsetsBandwidthLimitRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetIpsetsBandwidthLimitRequest) SetAcceleratorId(v string) *GetIpsetsBandwidthLimitRequest {
	s.AcceleratorId = &v
	return s
}

func (s *GetIpsetsBandwidthLimitRequest) SetRegionId(v string) *GetIpsetsBandwidthLimitRequest {
	s.RegionId = &v
	return s
}

func (s *GetIpsetsBandwidthLimitRequest) Validate() error {
	return dara.Validate(s)
}
