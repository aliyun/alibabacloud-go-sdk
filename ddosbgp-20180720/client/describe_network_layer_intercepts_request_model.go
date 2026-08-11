// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkLayerInterceptsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationIp(v string) *DescribeNetworkLayerInterceptsRequest
	GetDestinationIp() *string
	SetDestinationPort(v int64) *DescribeNetworkLayerInterceptsRequest
	GetDestinationPort() *int64
	SetEndTime(v int64) *DescribeNetworkLayerInterceptsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeNetworkLayerInterceptsRequest
	GetInstanceId() *string
	SetNetworkProtocol(v string) *DescribeNetworkLayerInterceptsRequest
	GetNetworkProtocol() *string
	SetPage(v int64) *DescribeNetworkLayerInterceptsRequest
	GetPage() *int64
	SetPageSize(v int64) *DescribeNetworkLayerInterceptsRequest
	GetPageSize() *int64
	SetProtocolNumber(v int64) *DescribeNetworkLayerInterceptsRequest
	GetProtocolNumber() *int64
	SetSourcePort(v int64) *DescribeNetworkLayerInterceptsRequest
	GetSourcePort() *int64
	SetSrcIp(v string) *DescribeNetworkLayerInterceptsRequest
	GetSrcIp() *string
	SetStartTime(v int64) *DescribeNetworkLayerInterceptsRequest
	GetStartTime() *int64
}

type DescribeNetworkLayerInterceptsRequest struct {
	// The destination IP address.
	//
	// example:
	//
	// 47.118.170.18
	DestinationIp *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	// The destination port.
	//
	// example:
	//
	// 8080
	DestinationPort *int64 `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	// The end time of the DDoS attack event to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1563445054
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID of the Anti-DDoS Origin instance to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-z2q1qzxb****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network protocol.
	//
	// example:
	//
	// tcp
	NetworkProtocol *string `json:"NetworkProtocol,omitempty" xml:"NetworkProtocol,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// Settings for the number of interception logs to return on each page when you perform a paged query. Paging is used to return results.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The network protocol number. This is a standard network protocol number.
	//
	// example:
	//
	// 6
	ProtocolNumber *int64 `json:"ProtocolNumber,omitempty" xml:"ProtocolNumber,omitempty"`
	// The source port.
	//
	// example:
	//
	// 5432
	SourcePort *int64 `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 37.60.241.154
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
	// The start time of the DDoS attack event to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1557305044
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeNetworkLayerInterceptsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkLayerInterceptsRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkLayerInterceptsRequest) GetDestinationIp() *string {
	return s.DestinationIp
}

func (s *DescribeNetworkLayerInterceptsRequest) GetDestinationPort() *int64 {
	return s.DestinationPort
}

func (s *DescribeNetworkLayerInterceptsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeNetworkLayerInterceptsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNetworkLayerInterceptsRequest) GetNetworkProtocol() *string {
	return s.NetworkProtocol
}

func (s *DescribeNetworkLayerInterceptsRequest) GetPage() *int64 {
	return s.Page
}

func (s *DescribeNetworkLayerInterceptsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeNetworkLayerInterceptsRequest) GetProtocolNumber() *int64 {
	return s.ProtocolNumber
}

func (s *DescribeNetworkLayerInterceptsRequest) GetSourcePort() *int64 {
	return s.SourcePort
}

func (s *DescribeNetworkLayerInterceptsRequest) GetSrcIp() *string {
	return s.SrcIp
}

func (s *DescribeNetworkLayerInterceptsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeNetworkLayerInterceptsRequest) SetDestinationIp(v string) *DescribeNetworkLayerInterceptsRequest {
	s.DestinationIp = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsRequest) SetDestinationPort(v int64) *DescribeNetworkLayerInterceptsRequest {
	s.DestinationPort = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsRequest) SetEndTime(v int64) *DescribeNetworkLayerInterceptsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsRequest) SetInstanceId(v string) *DescribeNetworkLayerInterceptsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsRequest) SetNetworkProtocol(v string) *DescribeNetworkLayerInterceptsRequest {
	s.NetworkProtocol = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsRequest) SetPage(v int64) *DescribeNetworkLayerInterceptsRequest {
	s.Page = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsRequest) SetPageSize(v int64) *DescribeNetworkLayerInterceptsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsRequest) SetProtocolNumber(v int64) *DescribeNetworkLayerInterceptsRequest {
	s.ProtocolNumber = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsRequest) SetSourcePort(v int64) *DescribeNetworkLayerInterceptsRequest {
	s.SourcePort = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsRequest) SetSrcIp(v string) *DescribeNetworkLayerInterceptsRequest {
	s.SrcIp = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsRequest) SetStartTime(v int64) *DescribeNetworkLayerInterceptsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsRequest) Validate() error {
	return dara.Validate(s)
}
