// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventTotalCountReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeEventTotalCountReportRequest
	GetLang() *string
	SetRegId(v string) *DescribeEventTotalCountReportRequest
	GetRegId() *string
}

type DescribeEventTotalCountReportRequest struct {
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

func (s DescribeEventTotalCountReportRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventTotalCountReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventTotalCountReportRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEventTotalCountReportRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeEventTotalCountReportRequest) SetLang(v string) *DescribeEventTotalCountReportRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventTotalCountReportRequest) SetRegId(v string) *DescribeEventTotalCountReportRequest {
	s.RegId = &v
	return s
}

func (s *DescribeEventTotalCountReportRequest) Validate() error {
	return dara.Validate(s)
}
