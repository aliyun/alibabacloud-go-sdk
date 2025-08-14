// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNameListTypeListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeNameListTypeListRequest
	GetLang() *string
	SetRegId(v string) *DescribeNameListTypeListRequest
	GetRegId() *string
}

type DescribeNameListTypeListRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
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

func (s DescribeNameListTypeListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListTypeListRequest) GoString() string {
	return s.String()
}

func (s *DescribeNameListTypeListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeNameListTypeListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeNameListTypeListRequest) SetLang(v string) *DescribeNameListTypeListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNameListTypeListRequest) SetRegId(v string) *DescribeNameListTypeListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeNameListTypeListRequest) Validate() error {
	return dara.Validate(s)
}
