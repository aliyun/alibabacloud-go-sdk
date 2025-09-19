// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllImageBaselineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAllImageBaselineRequest
	GetLang() *string
}

type DescribeAllImageBaselineRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeAllImageBaselineRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllImageBaselineRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllImageBaselineRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAllImageBaselineRequest) SetLang(v string) *DescribeAllImageBaselineRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAllImageBaselineRequest) Validate() error {
	return dara.Validate(s)
}
