// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetOpenServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeInternetOpenServiceRequest
	GetCurrentPage() *string
	SetEndTime(v string) *DescribeInternetOpenServiceRequest
	GetEndTime() *string
	SetLang(v string) *DescribeInternetOpenServiceRequest
	GetLang() *string
	SetPageSize(v string) *DescribeInternetOpenServiceRequest
	GetPageSize() *string
	SetPort(v string) *DescribeInternetOpenServiceRequest
	GetPort() *string
	SetRiskLevel(v string) *DescribeInternetOpenServiceRequest
	GetRiskLevel() *string
	SetServiceName(v string) *DescribeInternetOpenServiceRequest
	GetServiceName() *string
	SetServiceNameFuzzy(v string) *DescribeInternetOpenServiceRequest
	GetServiceNameFuzzy() *string
	SetSourceIp(v string) *DescribeInternetOpenServiceRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeInternetOpenServiceRequest
	GetStartTime() *string
	SetSuggestLevel(v string) *DescribeInternetOpenServiceRequest
	GetSuggestLevel() *string
}

type DescribeInternetOpenServiceRequest struct {
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 1753804800
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
	// 3389
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 2
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
	// 122.200.64.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// example:
	//
	// 1755742107
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 10
	SuggestLevel *string `json:"SuggestLevel,omitempty" xml:"SuggestLevel,omitempty"`
}

func (s DescribeInternetOpenServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenServiceRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeInternetOpenServiceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInternetOpenServiceRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInternetOpenServiceRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeInternetOpenServiceRequest) GetPort() *string {
	return s.Port
}

func (s *DescribeInternetOpenServiceRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeInternetOpenServiceRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeInternetOpenServiceRequest) GetServiceNameFuzzy() *string {
	return s.ServiceNameFuzzy
}

func (s *DescribeInternetOpenServiceRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInternetOpenServiceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInternetOpenServiceRequest) GetSuggestLevel() *string {
	return s.SuggestLevel
}

func (s *DescribeInternetOpenServiceRequest) SetCurrentPage(v string) *DescribeInternetOpenServiceRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInternetOpenServiceRequest) SetEndTime(v string) *DescribeInternetOpenServiceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInternetOpenServiceRequest) SetLang(v string) *DescribeInternetOpenServiceRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInternetOpenServiceRequest) SetPageSize(v string) *DescribeInternetOpenServiceRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInternetOpenServiceRequest) SetPort(v string) *DescribeInternetOpenServiceRequest {
	s.Port = &v
	return s
}

func (s *DescribeInternetOpenServiceRequest) SetRiskLevel(v string) *DescribeInternetOpenServiceRequest {
	s.RiskLevel = &v
	return s
}

func (s *DescribeInternetOpenServiceRequest) SetServiceName(v string) *DescribeInternetOpenServiceRequest {
	s.ServiceName = &v
	return s
}

func (s *DescribeInternetOpenServiceRequest) SetServiceNameFuzzy(v string) *DescribeInternetOpenServiceRequest {
	s.ServiceNameFuzzy = &v
	return s
}

func (s *DescribeInternetOpenServiceRequest) SetSourceIp(v string) *DescribeInternetOpenServiceRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInternetOpenServiceRequest) SetStartTime(v string) *DescribeInternetOpenServiceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInternetOpenServiceRequest) SetSuggestLevel(v string) *DescribeInternetOpenServiceRequest {
	s.SuggestLevel = &v
	return s
}

func (s *DescribeInternetOpenServiceRequest) Validate() error {
	return dara.Validate(s)
}
