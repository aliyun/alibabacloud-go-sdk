// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortAttackMaxFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBps(v int64) *DescribePortAttackMaxFlowResponseBody
	GetBps() *int64
	SetPps(v int64) *DescribePortAttackMaxFlowResponseBody
	GetPps() *int64
	SetRequestId(v string) *DescribePortAttackMaxFlowResponseBody
	GetRequestId() *string
}

type DescribePortAttackMaxFlowResponseBody struct {
	// The peak bandwidth of attack traffic. Unit: bit/s.
	//
	// example:
	//
	// 149559
	Bps *int64 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The peak packet rate of attack traffic . Unit: packets per second (pps).
	//
	// example:
	//
	// 23
	Pps *int64 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePortAttackMaxFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePortAttackMaxFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortAttackMaxFlowResponseBody) GetBps() *int64 {
	return s.Bps
}

func (s *DescribePortAttackMaxFlowResponseBody) GetPps() *int64 {
	return s.Pps
}

func (s *DescribePortAttackMaxFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePortAttackMaxFlowResponseBody) SetBps(v int64) *DescribePortAttackMaxFlowResponseBody {
	s.Bps = &v
	return s
}

func (s *DescribePortAttackMaxFlowResponseBody) SetPps(v int64) *DescribePortAttackMaxFlowResponseBody {
	s.Pps = &v
	return s
}

func (s *DescribePortAttackMaxFlowResponseBody) SetRequestId(v string) *DescribePortAttackMaxFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortAttackMaxFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
