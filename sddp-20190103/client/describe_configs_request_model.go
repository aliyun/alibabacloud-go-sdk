// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeConfigsRequest
	GetLang() *string
}

type DescribeConfigsRequest struct {
	// The language of the content within the request and response. Valid values:
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

func (s DescribeConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeConfigsRequest) SetLang(v string) *DescribeConfigsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeConfigsRequest) Validate() error {
	return dara.Validate(s)
}
