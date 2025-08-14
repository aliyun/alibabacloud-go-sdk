// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthEventNameListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAuthEventNameListRequest
	GetLang() *string
	SetRegId(v string) *DescribeAuthEventNameListRequest
	GetRegId() *string
}

type DescribeAuthEventNameListRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
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
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeAuthEventNameListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthEventNameListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuthEventNameListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAuthEventNameListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeAuthEventNameListRequest) SetLang(v string) *DescribeAuthEventNameListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAuthEventNameListRequest) SetRegId(v string) *DescribeAuthEventNameListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeAuthEventNameListRequest) Validate() error {
	return dara.Validate(s)
}
