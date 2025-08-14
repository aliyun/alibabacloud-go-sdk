// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRequestPeakReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRequestPeakReportRequest
	GetLang() *string
	SetRegId(v string) *DescribeRequestPeakReportRequest
	GetRegId() *string
}

type DescribeRequestPeakReportRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeRequestPeakReportRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRequestPeakReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeRequestPeakReportRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRequestPeakReportRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeRequestPeakReportRequest) SetLang(v string) *DescribeRequestPeakReportRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRequestPeakReportRequest) SetRegId(v string) *DescribeRequestPeakReportRequest {
	s.RegId = &v
	return s
}

func (s *DescribeRequestPeakReportRequest) Validate() error {
	return dara.Validate(s)
}
