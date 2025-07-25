// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeCustomLineRequest
	GetLang() *string
	SetLineId(v int64) *DescribeCustomLineRequest
	GetLineId() *int64
}

type DescribeCustomLineRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The unique ID of the custom line. You can call [DescribeCustomLines](https://www.alibabacloud.com/help/en/dns/api-alidns-2015-01-09-describecustomlines?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the ID
	//
	// example:
	//
	// 597
	LineId *int64 `json:"LineId,omitempty" xml:"LineId,omitempty"`
}

func (s DescribeCustomLineRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomLineRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomLineRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCustomLineRequest) GetLineId() *int64 {
	return s.LineId
}

func (s *DescribeCustomLineRequest) SetLang(v string) *DescribeCustomLineRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCustomLineRequest) SetLineId(v int64) *DescribeCustomLineRequest {
	s.LineId = &v
	return s
}

func (s *DescribeCustomLineRequest) Validate() error {
	return dara.Validate(s)
}
