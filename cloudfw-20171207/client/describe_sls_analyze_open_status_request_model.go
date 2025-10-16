// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsAnalyzeOpenStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSlsAnalyzeOpenStatusRequest
	GetLang() *string
}

type DescribeSlsAnalyzeOpenStatusRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeSlsAnalyzeOpenStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsAnalyzeOpenStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsAnalyzeOpenStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSlsAnalyzeOpenStatusRequest) SetLang(v string) *DescribeSlsAnalyzeOpenStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsAnalyzeOpenStatusRequest) Validate() error {
	return dara.Validate(s)
}
