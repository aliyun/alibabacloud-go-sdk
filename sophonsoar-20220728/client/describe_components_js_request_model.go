// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentsJsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeComponentsJsRequest
	GetLang() *string
}

type DescribeComponentsJsRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeComponentsJsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentsJsRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentsJsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeComponentsJsRequest) SetLang(v string) *DescribeComponentsJsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeComponentsJsRequest) Validate() error {
	return dara.Validate(s)
}
