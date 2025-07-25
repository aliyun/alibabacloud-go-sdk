// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomLinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteCustomLinesRequest
	GetLang() *string
	SetLineIds(v string) *DeleteCustomLinesRequest
	GetLineIds() *string
}

type DeleteCustomLinesRequest struct {
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
	// The unique IDs of the custom lines that you want to delete. Separate the unique IDs with commas (,). You can call the [DescribeCustomLines](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describecustomlines?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234,1235
	LineIds *string `json:"LineIds,omitempty" xml:"LineIds,omitempty"`
}

func (s DeleteCustomLinesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomLinesRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomLinesRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteCustomLinesRequest) GetLineIds() *string {
	return s.LineIds
}

func (s *DeleteCustomLinesRequest) SetLang(v string) *DeleteCustomLinesRequest {
	s.Lang = &v
	return s
}

func (s *DeleteCustomLinesRequest) SetLineIds(v string) *DeleteCustomLinesRequest {
	s.LineIds = &v
	return s
}

func (s *DeleteCustomLinesRequest) Validate() error {
	return dara.Validate(s)
}
