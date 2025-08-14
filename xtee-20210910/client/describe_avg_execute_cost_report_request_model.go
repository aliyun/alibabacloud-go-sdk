// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvgExecuteCostReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAvgExecuteCostReportRequest
	GetLang() *string
	SetRegId(v string) *DescribeAvgExecuteCostReportRequest
	GetRegId() *string
}

type DescribeAvgExecuteCostReportRequest struct {
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

func (s DescribeAvgExecuteCostReportRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvgExecuteCostReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvgExecuteCostReportRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAvgExecuteCostReportRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeAvgExecuteCostReportRequest) SetLang(v string) *DescribeAvgExecuteCostReportRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAvgExecuteCostReportRequest) SetRegId(v string) *DescribeAvgExecuteCostReportRequest {
	s.RegId = &v
	return s
}

func (s *DescribeAvgExecuteCostReportRequest) Validate() error {
	return dara.Validate(s)
}
