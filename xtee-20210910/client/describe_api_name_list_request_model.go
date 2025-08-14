// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiNameListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeApiNameListRequest
	GetLang() *string
	SetRegId(v string) *DescribeApiNameListRequest
	GetRegId() *string
}

type DescribeApiNameListRequest struct {
	// Sets the language type for requests and responses, with a default value of **zh**. Values:
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

func (s DescribeApiNameListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiNameListRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiNameListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeApiNameListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeApiNameListRequest) SetLang(v string) *DescribeApiNameListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeApiNameListRequest) SetRegId(v string) *DescribeApiNameListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeApiNameListRequest) Validate() error {
	return dara.Validate(s)
}
