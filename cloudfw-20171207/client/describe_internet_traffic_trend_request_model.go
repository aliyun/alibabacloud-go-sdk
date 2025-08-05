// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetTrafficTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *DescribeInternetTrafficTrendRequest
	GetDirection() *string
	SetEndTime(v string) *DescribeInternetTrafficTrendRequest
	GetEndTime() *string
	SetLang(v string) *DescribeInternetTrafficTrendRequest
	GetLang() *string
	SetSourceCode(v string) *DescribeInternetTrafficTrendRequest
	GetSourceCode() *string
	SetSourceIp(v string) *DescribeInternetTrafficTrendRequest
	GetSourceIp() *string
	SetSrcPrivateIP(v string) *DescribeInternetTrafficTrendRequest
	GetSrcPrivateIP() *string
	SetSrcPublicIP(v string) *DescribeInternetTrafficTrendRequest
	GetSrcPublicIP() *string
	SetStartTime(v string) *DescribeInternetTrafficTrendRequest
	GetStartTime() *string
	SetTrafficType(v string) *DescribeInternetTrafficTrendRequest
	GetTrafficType() *string
}

type DescribeInternetTrafficTrendRequest struct {
	// The direction of the internet traffic.
	//
	// Valid values:
	//
	// 	- **in**: inbound traffic
	//
	// 	- **out**: outbound traffic
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1674958929
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content in the request and response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source code.
	//
	// This parameter is required.
	//
	// example:
	//
	// yundun
	SourceCode *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	// Deprecated
	//
	// The IP address of the access source.
	//
	// example:
	//
	// 101.80.171.196
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The private IP address of the source.
	//
	// example:
	//
	// 10.100.134.60
	SrcPrivateIP *string `json:"SrcPrivateIP,omitempty" xml:"SrcPrivateIP,omitempty"`
	// The public IP address of the source.
	//
	// example:
	//
	// 47.112.210.136
	SrcPublicIP *string `json:"SrcPublicIP,omitempty" xml:"SrcPublicIP,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1670307484
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the traffic that is captured. Valid values:
	//
	// 	- **max*	- (default): peak traffic
	//
	// 	- **avg**: average traffic
	//
	// example:
	//
	// max
	TrafficType *string `json:"TrafficType,omitempty" xml:"TrafficType,omitempty"`
}

func (s DescribeInternetTrafficTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetTrafficTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeInternetTrafficTrendRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeInternetTrafficTrendRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInternetTrafficTrendRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInternetTrafficTrendRequest) GetSourceCode() *string {
	return s.SourceCode
}

func (s *DescribeInternetTrafficTrendRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInternetTrafficTrendRequest) GetSrcPrivateIP() *string {
	return s.SrcPrivateIP
}

func (s *DescribeInternetTrafficTrendRequest) GetSrcPublicIP() *string {
	return s.SrcPublicIP
}

func (s *DescribeInternetTrafficTrendRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInternetTrafficTrendRequest) GetTrafficType() *string {
	return s.TrafficType
}

func (s *DescribeInternetTrafficTrendRequest) SetDirection(v string) *DescribeInternetTrafficTrendRequest {
	s.Direction = &v
	return s
}

func (s *DescribeInternetTrafficTrendRequest) SetEndTime(v string) *DescribeInternetTrafficTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInternetTrafficTrendRequest) SetLang(v string) *DescribeInternetTrafficTrendRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInternetTrafficTrendRequest) SetSourceCode(v string) *DescribeInternetTrafficTrendRequest {
	s.SourceCode = &v
	return s
}

func (s *DescribeInternetTrafficTrendRequest) SetSourceIp(v string) *DescribeInternetTrafficTrendRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInternetTrafficTrendRequest) SetSrcPrivateIP(v string) *DescribeInternetTrafficTrendRequest {
	s.SrcPrivateIP = &v
	return s
}

func (s *DescribeInternetTrafficTrendRequest) SetSrcPublicIP(v string) *DescribeInternetTrafficTrendRequest {
	s.SrcPublicIP = &v
	return s
}

func (s *DescribeInternetTrafficTrendRequest) SetStartTime(v string) *DescribeInternetTrafficTrendRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInternetTrafficTrendRequest) SetTrafficType(v string) *DescribeInternetTrafficTrendRequest {
	s.TrafficType = &v
	return s
}

func (s *DescribeInternetTrafficTrendRequest) Validate() error {
	return dara.Validate(s)
}
