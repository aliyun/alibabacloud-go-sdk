// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetOpenPortRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeInternetOpenPortRequest
	GetCurrentPage() *string
	SetEndTime(v string) *DescribeInternetOpenPortRequest
	GetEndTime() *string
	SetLang(v string) *DescribeInternetOpenPortRequest
	GetLang() *string
	SetPageSize(v string) *DescribeInternetOpenPortRequest
	GetPageSize() *string
	SetPort(v string) *DescribeInternetOpenPortRequest
	GetPort() *string
	SetRiskLevel(v string) *DescribeInternetOpenPortRequest
	GetRiskLevel() *string
	SetServiceName(v string) *DescribeInternetOpenPortRequest
	GetServiceName() *string
	SetServiceNameFuzzy(v string) *DescribeInternetOpenPortRequest
	GetServiceNameFuzzy() *string
	SetSourceIp(v string) *DescribeInternetOpenPortRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeInternetOpenPortRequest
	GetStartTime() *string
	SetSuggestLevel(v string) *DescribeInternetOpenPortRequest
	GetSuggestLevel() *string
}

type DescribeInternetOpenPortRequest struct {
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 1748358644
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 3
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// SMB
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// example:
	//
	// SMB
	ServiceNameFuzzy *string `json:"ServiceNameFuzzy,omitempty" xml:"ServiceNameFuzzy,omitempty"`
	// example:
	//
	// 114.242.33.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// example:
	//
	// 1735264800
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 10
	SuggestLevel *string `json:"SuggestLevel,omitempty" xml:"SuggestLevel,omitempty"`
}

func (s DescribeInternetOpenPortRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenPortRequest) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenPortRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeInternetOpenPortRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInternetOpenPortRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInternetOpenPortRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeInternetOpenPortRequest) GetPort() *string {
	return s.Port
}

func (s *DescribeInternetOpenPortRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeInternetOpenPortRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeInternetOpenPortRequest) GetServiceNameFuzzy() *string {
	return s.ServiceNameFuzzy
}

func (s *DescribeInternetOpenPortRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInternetOpenPortRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInternetOpenPortRequest) GetSuggestLevel() *string {
	return s.SuggestLevel
}

func (s *DescribeInternetOpenPortRequest) SetCurrentPage(v string) *DescribeInternetOpenPortRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInternetOpenPortRequest) SetEndTime(v string) *DescribeInternetOpenPortRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInternetOpenPortRequest) SetLang(v string) *DescribeInternetOpenPortRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInternetOpenPortRequest) SetPageSize(v string) *DescribeInternetOpenPortRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInternetOpenPortRequest) SetPort(v string) *DescribeInternetOpenPortRequest {
	s.Port = &v
	return s
}

func (s *DescribeInternetOpenPortRequest) SetRiskLevel(v string) *DescribeInternetOpenPortRequest {
	s.RiskLevel = &v
	return s
}

func (s *DescribeInternetOpenPortRequest) SetServiceName(v string) *DescribeInternetOpenPortRequest {
	s.ServiceName = &v
	return s
}

func (s *DescribeInternetOpenPortRequest) SetServiceNameFuzzy(v string) *DescribeInternetOpenPortRequest {
	s.ServiceNameFuzzy = &v
	return s
}

func (s *DescribeInternetOpenPortRequest) SetSourceIp(v string) *DescribeInternetOpenPortRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInternetOpenPortRequest) SetStartTime(v string) *DescribeInternetOpenPortRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInternetOpenPortRequest) SetSuggestLevel(v string) *DescribeInternetOpenPortRequest {
	s.SuggestLevel = &v
	return s
}

func (s *DescribeInternetOpenPortRequest) Validate() error {
	return dara.Validate(s)
}
