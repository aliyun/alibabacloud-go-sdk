// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerId(v string) *DescribeListenerRequest
	GetListenerId() *string
	SetRegionId(v string) *DescribeListenerRequest
	GetRegionId() *string
}

type DescribeListenerRequest struct {
	// The ID of the listener that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeListenerRequest) GoString() string {
	return s.String()
}

func (s *DescribeListenerRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *DescribeListenerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeListenerRequest) SetListenerId(v string) *DescribeListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *DescribeListenerRequest) SetRegionId(v string) *DescribeListenerRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeListenerRequest) Validate() error {
	return dara.Validate(s)
}
