// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnalysisColumnFieldListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAnalysisColumnFieldListRequest
	GetLang() *string
	SetRegId(v string) *DescribeAnalysisColumnFieldListRequest
	GetRegId() *string
}

type DescribeAnalysisColumnFieldListRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
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
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeAnalysisColumnFieldListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnalysisColumnFieldListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAnalysisColumnFieldListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAnalysisColumnFieldListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeAnalysisColumnFieldListRequest) SetLang(v string) *DescribeAnalysisColumnFieldListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAnalysisColumnFieldListRequest) SetRegId(v string) *DescribeAnalysisColumnFieldListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeAnalysisColumnFieldListRequest) Validate() error {
	return dara.Validate(s)
}
