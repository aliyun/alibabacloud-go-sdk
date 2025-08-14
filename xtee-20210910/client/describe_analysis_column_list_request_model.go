// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnalysisColumnListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAnalysisColumnListRequest
	GetLang() *string
	SetRegId(v string) *DescribeAnalysisColumnListRequest
	GetRegId() *string
}

type DescribeAnalysisColumnListRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeAnalysisColumnListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnalysisColumnListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAnalysisColumnListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAnalysisColumnListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeAnalysisColumnListRequest) SetLang(v string) *DescribeAnalysisColumnListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAnalysisColumnListRequest) SetRegId(v string) *DescribeAnalysisColumnListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeAnalysisColumnListRequest) Validate() error {
	return dara.Validate(s)
}
