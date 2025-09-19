// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStrictEventNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeStrictEventNameRequest
	GetLang() *string
}

type DescribeStrictEventNameRequest struct {
	// Sets the language type for requests and received messages, default is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeStrictEventNameRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrictEventNameRequest) GoString() string {
	return s.String()
}

func (s *DescribeStrictEventNameRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeStrictEventNameRequest) SetLang(v string) *DescribeStrictEventNameRequest {
	s.Lang = &v
	return s
}

func (s *DescribeStrictEventNameRequest) Validate() error {
	return dara.Validate(s)
}
