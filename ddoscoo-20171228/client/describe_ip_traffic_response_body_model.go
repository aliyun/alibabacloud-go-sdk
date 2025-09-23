// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpTrafficResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvgInBps(v int64) *DescribeIpTrafficResponseBody
	GetAvgInBps() *int64
	SetAvgOutBps(v int64) *DescribeIpTrafficResponseBody
	GetAvgOutBps() *int64
	SetIpTrafficPoints(v []*DescribeIpTrafficResponseBodyIpTrafficPoints) *DescribeIpTrafficResponseBody
	GetIpTrafficPoints() []*DescribeIpTrafficResponseBodyIpTrafficPoints
	SetMaxInBps(v int64) *DescribeIpTrafficResponseBody
	GetMaxInBps() *int64
	SetMaxOutBps(v int64) *DescribeIpTrafficResponseBody
	GetMaxOutBps() *int64
	SetRequestId(v string) *DescribeIpTrafficResponseBody
	GetRequestId() *string
}

type DescribeIpTrafficResponseBody struct {
	// example:
	//
	// 10000
	AvgInBps *int64 `json:"AvgInBps,omitempty" xml:"AvgInBps,omitempty"`
	// example:
	//
	// 10000
	AvgOutBps       *int64                                          `json:"AvgOutBps,omitempty" xml:"AvgOutBps,omitempty"`
	IpTrafficPoints []*DescribeIpTrafficResponseBodyIpTrafficPoints `json:"IpTrafficPoints,omitempty" xml:"IpTrafficPoints,omitempty" type:"Repeated"`
	// example:
	//
	// 10000
	MaxInBps *int64 `json:"MaxInBps,omitempty" xml:"MaxInBps,omitempty"`
	// example:
	//
	// 10000
	MaxOutBps *int64 `json:"MaxOutBps,omitempty" xml:"MaxOutBps,omitempty"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIpTrafficResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpTrafficResponseBody) GetAvgInBps() *int64 {
	return s.AvgInBps
}

func (s *DescribeIpTrafficResponseBody) GetAvgOutBps() *int64 {
	return s.AvgOutBps
}

func (s *DescribeIpTrafficResponseBody) GetIpTrafficPoints() []*DescribeIpTrafficResponseBodyIpTrafficPoints {
	return s.IpTrafficPoints
}

func (s *DescribeIpTrafficResponseBody) GetMaxInBps() *int64 {
	return s.MaxInBps
}

func (s *DescribeIpTrafficResponseBody) GetMaxOutBps() *int64 {
	return s.MaxOutBps
}

func (s *DescribeIpTrafficResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIpTrafficResponseBody) SetAvgInBps(v int64) *DescribeIpTrafficResponseBody {
	s.AvgInBps = &v
	return s
}

func (s *DescribeIpTrafficResponseBody) SetAvgOutBps(v int64) *DescribeIpTrafficResponseBody {
	s.AvgOutBps = &v
	return s
}

func (s *DescribeIpTrafficResponseBody) SetIpTrafficPoints(v []*DescribeIpTrafficResponseBodyIpTrafficPoints) *DescribeIpTrafficResponseBody {
	s.IpTrafficPoints = v
	return s
}

func (s *DescribeIpTrafficResponseBody) SetMaxInBps(v int64) *DescribeIpTrafficResponseBody {
	s.MaxInBps = &v
	return s
}

func (s *DescribeIpTrafficResponseBody) SetMaxOutBps(v int64) *DescribeIpTrafficResponseBody {
	s.MaxOutBps = &v
	return s
}

func (s *DescribeIpTrafficResponseBody) SetRequestId(v string) *DescribeIpTrafficResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpTrafficResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeIpTrafficResponseBodyIpTrafficPoints struct {
	// example:
	//
	// 100
	ActConns *int32 `json:"ActConns,omitempty" xml:"ActConns,omitempty"`
	// example:
	//
	// 100
	Cps *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// example:
	//
	// 100
	InactConns *int32 `json:"InactConns,omitempty" xml:"InactConns,omitempty"`
	// example:
	//
	// 10000
	MaxInbps *int64 `json:"MaxInbps,omitempty" xml:"MaxInbps,omitempty"`
	// example:
	//
	// 10000
	MaxOutbps *int64 `json:"MaxOutbps,omitempty" xml:"MaxOutbps,omitempty"`
	// example:
	//
	// 1536734112
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeIpTrafficResponseBodyIpTrafficPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpTrafficResponseBodyIpTrafficPoints) GoString() string {
	return s.String()
}

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) GetActConns() *int32 {
	return s.ActConns
}

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) GetCps() *int32 {
	return s.Cps
}

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) GetInactConns() *int32 {
	return s.InactConns
}

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) GetMaxInbps() *int64 {
	return s.MaxInbps
}

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) GetMaxOutbps() *int64 {
	return s.MaxOutbps
}

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) GetTime() *int64 {
	return s.Time
}

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) SetActConns(v int32) *DescribeIpTrafficResponseBodyIpTrafficPoints {
	s.ActConns = &v
	return s
}

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) SetCps(v int32) *DescribeIpTrafficResponseBodyIpTrafficPoints {
	s.Cps = &v
	return s
}

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) SetInactConns(v int32) *DescribeIpTrafficResponseBodyIpTrafficPoints {
	s.InactConns = &v
	return s
}

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) SetMaxInbps(v int64) *DescribeIpTrafficResponseBodyIpTrafficPoints {
	s.MaxInbps = &v
	return s
}

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) SetMaxOutbps(v int64) *DescribeIpTrafficResponseBodyIpTrafficPoints {
	s.MaxOutbps = &v
	return s
}

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) SetTime(v int64) *DescribeIpTrafficResponseBodyIpTrafficPoints {
	s.Time = &v
	return s
}

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) Validate() error {
	return dara.Validate(s)
}
