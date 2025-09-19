// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupStructRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeGroupStructRequest
	GetLang() *string
}

type DescribeGroupStructRequest struct {
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

func (s DescribeGroupStructRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupStructRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupStructRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGroupStructRequest) SetLang(v string) *DescribeGroupStructRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGroupStructRequest) Validate() error {
	return dara.Validate(s)
}
