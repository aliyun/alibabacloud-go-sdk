// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetTrafficTopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataType(v string) *DescribeInternetTrafficTopRequest
	GetDataType() *string
	SetDirection(v string) *DescribeInternetTrafficTopRequest
	GetDirection() *string
	SetEndTime(v string) *DescribeInternetTrafficTopRequest
	GetEndTime() *string
	SetLang(v string) *DescribeInternetTrafficTopRequest
	GetLang() *string
	SetLimit(v string) *DescribeInternetTrafficTopRequest
	GetLimit() *string
	SetRuleResult(v string) *DescribeInternetTrafficTopRequest
	GetRuleResult() *string
	SetRuleSource(v string) *DescribeInternetTrafficTopRequest
	GetRuleSource() *string
	SetShowCountryName(v string) *DescribeInternetTrafficTopRequest
	GetShowCountryName() *string
	SetSort(v string) *DescribeInternetTrafficTopRequest
	GetSort() *string
	SetSourceCode(v string) *DescribeInternetTrafficTopRequest
	GetSourceCode() *string
	SetSourceIp(v string) *DescribeInternetTrafficTopRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeInternetTrafficTopRequest
	GetStartTime() *string
}

type DescribeInternetTrafficTopRequest struct {
	// The type of traffic data.
	//
	// This parameter is required.
	//
	// example:
	//
	// in_src_ip
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The traffic direction.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end time. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1734055824
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the response.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of top entries to return. The default value is 50. Valid values: 1 to 50.
	//
	// example:
	//
	// 50
	Limit *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The action of the rule.
	//
	// example:
	//
	// 1
	RuleResult *string `json:"RuleResult,omitempty" xml:"RuleResult,omitempty"`
	// The source of the rule.
	//
	// example:
	//
	// 1
	RuleSource *string `json:"RuleSource,omitempty" xml:"RuleSource,omitempty"`
	// The name of the country to display.
	//
	// example:
	//
	// China
	ShowCountryName *string `json:"ShowCountryName,omitempty" xml:"ShowCountryName,omitempty"`
	// The sorting method.
	//
	// example:
	//
	// in_bytes
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The tracing code for the source.
	//
	// This parameter is required.
	//
	// example:
	//
	// yundun
	SourceCode *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 117.82.14.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The start time. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1656664560
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInternetTrafficTopRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetTrafficTopRequest) GoString() string {
	return s.String()
}

func (s *DescribeInternetTrafficTopRequest) GetDataType() *string {
	return s.DataType
}

func (s *DescribeInternetTrafficTopRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeInternetTrafficTopRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInternetTrafficTopRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInternetTrafficTopRequest) GetLimit() *string {
	return s.Limit
}

func (s *DescribeInternetTrafficTopRequest) GetRuleResult() *string {
	return s.RuleResult
}

func (s *DescribeInternetTrafficTopRequest) GetRuleSource() *string {
	return s.RuleSource
}

func (s *DescribeInternetTrafficTopRequest) GetShowCountryName() *string {
	return s.ShowCountryName
}

func (s *DescribeInternetTrafficTopRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeInternetTrafficTopRequest) GetSourceCode() *string {
	return s.SourceCode
}

func (s *DescribeInternetTrafficTopRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInternetTrafficTopRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInternetTrafficTopRequest) SetDataType(v string) *DescribeInternetTrafficTopRequest {
	s.DataType = &v
	return s
}

func (s *DescribeInternetTrafficTopRequest) SetDirection(v string) *DescribeInternetTrafficTopRequest {
	s.Direction = &v
	return s
}

func (s *DescribeInternetTrafficTopRequest) SetEndTime(v string) *DescribeInternetTrafficTopRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInternetTrafficTopRequest) SetLang(v string) *DescribeInternetTrafficTopRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInternetTrafficTopRequest) SetLimit(v string) *DescribeInternetTrafficTopRequest {
	s.Limit = &v
	return s
}

func (s *DescribeInternetTrafficTopRequest) SetRuleResult(v string) *DescribeInternetTrafficTopRequest {
	s.RuleResult = &v
	return s
}

func (s *DescribeInternetTrafficTopRequest) SetRuleSource(v string) *DescribeInternetTrafficTopRequest {
	s.RuleSource = &v
	return s
}

func (s *DescribeInternetTrafficTopRequest) SetShowCountryName(v string) *DescribeInternetTrafficTopRequest {
	s.ShowCountryName = &v
	return s
}

func (s *DescribeInternetTrafficTopRequest) SetSort(v string) *DescribeInternetTrafficTopRequest {
	s.Sort = &v
	return s
}

func (s *DescribeInternetTrafficTopRequest) SetSourceCode(v string) *DescribeInternetTrafficTopRequest {
	s.SourceCode = &v
	return s
}

func (s *DescribeInternetTrafficTopRequest) SetSourceIp(v string) *DescribeInternetTrafficTopRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInternetTrafficTopRequest) SetStartTime(v string) *DescribeInternetTrafficTopRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInternetTrafficTopRequest) Validate() error {
	return dara.Validate(s)
}
