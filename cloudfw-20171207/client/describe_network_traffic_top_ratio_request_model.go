// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkTrafficTopRatioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeNetworkTrafficTopRatioRequest
	GetAppName() *string
	SetAssetIP(v string) *DescribeNetworkTrafficTopRatioRequest
	GetAssetIP() *string
	SetAssetRegion(v string) *DescribeNetworkTrafficTopRatioRequest
	GetAssetRegion() *string
	SetDataType(v string) *DescribeNetworkTrafficTopRatioRequest
	GetDataType() *string
	SetDirection(v string) *DescribeNetworkTrafficTopRatioRequest
	GetDirection() *string
	SetDstIP(v string) *DescribeNetworkTrafficTopRatioRequest
	GetDstIP() *string
	SetDstPort(v string) *DescribeNetworkTrafficTopRatioRequest
	GetDstPort() *string
	SetEndTime(v string) *DescribeNetworkTrafficTopRatioRequest
	GetEndTime() *string
	SetIpProperty(v string) *DescribeNetworkTrafficTopRatioRequest
	GetIpProperty() *string
	SetIsp(v string) *DescribeNetworkTrafficTopRatioRequest
	GetIsp() *string
	SetLang(v string) *DescribeNetworkTrafficTopRatioRequest
	GetLang() *string
	SetLocation(v string) *DescribeNetworkTrafficTopRatioRequest
	GetLocation() *string
	SetRuleResult(v string) *DescribeNetworkTrafficTopRatioRequest
	GetRuleResult() *string
	SetSort(v string) *DescribeNetworkTrafficTopRatioRequest
	GetSort() *string
	SetSourceCode(v string) *DescribeNetworkTrafficTopRatioRequest
	GetSourceCode() *string
	SetSourceIp(v string) *DescribeNetworkTrafficTopRatioRequest
	GetSourceIp() *string
	SetSrcIP(v string) *DescribeNetworkTrafficTopRatioRequest
	GetSrcIP() *string
	SetStartTime(v string) *DescribeNetworkTrafficTopRatioRequest
	GetStartTime() *string
}

type DescribeNetworkTrafficTopRatioRequest struct {
	// example:
	//
	// HTTP
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 192.0.XX.XX
	AssetIP *string `json:"AssetIP,omitempty" xml:"AssetIP,omitempty"`
	// example:
	//
	// cn-beijing
	AssetRegion *string `json:"AssetRegion,omitempty" xml:"AssetRegion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// in_src_ip
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// example:
	//
	// 39.144.124.XXX
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// example:
	//
	// 8080
	DstPort *string `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1757433863
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// proxy
	IpProperty *string `json:"IpProperty,omitempty" xml:"IpProperty,omitempty"`
	Isp        *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// example:
	//
	// zh
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// 1
	RuleResult *string `json:"RuleResult,omitempty" xml:"RuleResult,omitempty"`
	// example:
	//
	// in_bytes
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yundun
	SourceCode *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	// example:
	//
	// 60.12.220.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// example:
	//
	// 172.16.169.XXX
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1749176793
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeNetworkTrafficTopRatioRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkTrafficTopRatioRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkTrafficTopRatioRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeNetworkTrafficTopRatioRequest) GetAssetIP() *string {
	return s.AssetIP
}

func (s *DescribeNetworkTrafficTopRatioRequest) GetAssetRegion() *string {
	return s.AssetRegion
}

func (s *DescribeNetworkTrafficTopRatioRequest) GetDataType() *string {
	return s.DataType
}

func (s *DescribeNetworkTrafficTopRatioRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeNetworkTrafficTopRatioRequest) GetDstIP() *string {
	return s.DstIP
}

func (s *DescribeNetworkTrafficTopRatioRequest) GetDstPort() *string {
	return s.DstPort
}

func (s *DescribeNetworkTrafficTopRatioRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeNetworkTrafficTopRatioRequest) GetIpProperty() *string {
	return s.IpProperty
}

func (s *DescribeNetworkTrafficTopRatioRequest) GetIsp() *string {
	return s.Isp
}

func (s *DescribeNetworkTrafficTopRatioRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeNetworkTrafficTopRatioRequest) GetLocation() *string {
	return s.Location
}

func (s *DescribeNetworkTrafficTopRatioRequest) GetRuleResult() *string {
	return s.RuleResult
}

func (s *DescribeNetworkTrafficTopRatioRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeNetworkTrafficTopRatioRequest) GetSourceCode() *string {
	return s.SourceCode
}

func (s *DescribeNetworkTrafficTopRatioRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeNetworkTrafficTopRatioRequest) GetSrcIP() *string {
	return s.SrcIP
}

func (s *DescribeNetworkTrafficTopRatioRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeNetworkTrafficTopRatioRequest) SetAppName(v string) *DescribeNetworkTrafficTopRatioRequest {
	s.AppName = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioRequest) SetAssetIP(v string) *DescribeNetworkTrafficTopRatioRequest {
	s.AssetIP = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioRequest) SetAssetRegion(v string) *DescribeNetworkTrafficTopRatioRequest {
	s.AssetRegion = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioRequest) SetDataType(v string) *DescribeNetworkTrafficTopRatioRequest {
	s.DataType = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioRequest) SetDirection(v string) *DescribeNetworkTrafficTopRatioRequest {
	s.Direction = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioRequest) SetDstIP(v string) *DescribeNetworkTrafficTopRatioRequest {
	s.DstIP = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioRequest) SetDstPort(v string) *DescribeNetworkTrafficTopRatioRequest {
	s.DstPort = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioRequest) SetEndTime(v string) *DescribeNetworkTrafficTopRatioRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioRequest) SetIpProperty(v string) *DescribeNetworkTrafficTopRatioRequest {
	s.IpProperty = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioRequest) SetIsp(v string) *DescribeNetworkTrafficTopRatioRequest {
	s.Isp = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioRequest) SetLang(v string) *DescribeNetworkTrafficTopRatioRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioRequest) SetLocation(v string) *DescribeNetworkTrafficTopRatioRequest {
	s.Location = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioRequest) SetRuleResult(v string) *DescribeNetworkTrafficTopRatioRequest {
	s.RuleResult = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioRequest) SetSort(v string) *DescribeNetworkTrafficTopRatioRequest {
	s.Sort = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioRequest) SetSourceCode(v string) *DescribeNetworkTrafficTopRatioRequest {
	s.SourceCode = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioRequest) SetSourceIp(v string) *DescribeNetworkTrafficTopRatioRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioRequest) SetSrcIP(v string) *DescribeNetworkTrafficTopRatioRequest {
	s.SrcIP = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioRequest) SetStartTime(v string) *DescribeNetworkTrafficTopRatioRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioRequest) Validate() error {
	return dara.Validate(s)
}
