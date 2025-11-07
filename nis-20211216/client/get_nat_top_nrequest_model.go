// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNatTopNRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *GetNatTopNRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *GetNatTopNRequest
	GetEndTime() *int64
	SetIp(v string) *GetNatTopNRequest
	GetIp() *string
	SetNatGatewayId(v string) *GetNatTopNRequest
	GetNatGatewayId() *string
	SetOrderBy(v string) *GetNatTopNRequest
	GetOrderBy() *string
	SetRegionId(v string) *GetNatTopNRequest
	GetRegionId() *string
	SetTopN(v int32) *GetNatTopNRequest
	GetTopN() *int32
}

type GetNatTopNRequest struct {
	// The beginning of the time range to query in milliseconds. If you do not specify **EndTime**, the point in time specified by **BeginTime*	- is queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1638239092000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end of the time range to query in milliseconds. The time range specified by **BeginTime*	- and **EndTime*	- cannot exceed **86400000*	- milliseconds (24 hours).
	//
	// example:
	//
	// 1638239093000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Query ranking statistics for a specific IP address. If you specify this parameter, you do not need to specify **TopN*	- or **OrderBy**.
	//
	// example:
	//
	// 192.168.156.101
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the NAT gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-sample***
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The metric that is used for real-time SNAT performance ranking. Valid values:
	//
	// 	- **InBps**: inbound data transfer. Unit: bit/s.
	//
	// 	- **OutBps**: outbound data transfer. Unit: bit/s.
	//
	// 	- **InPps**: inbound packet forwarding rate. Unit: packets per second.
	//
	// 	- **OutPps**: outbound packet forwarding rate. Unit: packets per second.
	//
	// 	- **NewSessionPerSecond**: new connection creation rate. Unit: connections per second.
	//
	// 	- **ActiveSessionCount**: number of concurrent connections. Unit: connections.
	//
	// example:
	//
	// InBps
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The ID of the region in which the NAT gateway is deployed.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of entries to return for real-time SNAT performance ranking. Valid values: **1 to 100**. Default value: **10**.
	//
	// example:
	//
	// 10
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
}

func (s GetNatTopNRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNatTopNRequest) GoString() string {
	return s.String()
}

func (s *GetNatTopNRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetNatTopNRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetNatTopNRequest) GetIp() *string {
	return s.Ip
}

func (s *GetNatTopNRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *GetNatTopNRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetNatTopNRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetNatTopNRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *GetNatTopNRequest) SetBeginTime(v int64) *GetNatTopNRequest {
	s.BeginTime = &v
	return s
}

func (s *GetNatTopNRequest) SetEndTime(v int64) *GetNatTopNRequest {
	s.EndTime = &v
	return s
}

func (s *GetNatTopNRequest) SetIp(v string) *GetNatTopNRequest {
	s.Ip = &v
	return s
}

func (s *GetNatTopNRequest) SetNatGatewayId(v string) *GetNatTopNRequest {
	s.NatGatewayId = &v
	return s
}

func (s *GetNatTopNRequest) SetOrderBy(v string) *GetNatTopNRequest {
	s.OrderBy = &v
	return s
}

func (s *GetNatTopNRequest) SetRegionId(v string) *GetNatTopNRequest {
	s.RegionId = &v
	return s
}

func (s *GetNatTopNRequest) SetTopN(v int32) *GetNatTopNRequest {
	s.TopN = &v
	return s
}

func (s *GetNatTopNRequest) Validate() error {
	return dara.Validate(s)
}
