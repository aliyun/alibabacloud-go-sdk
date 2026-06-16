// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DescribeTagListRequest
	GetId() *string
	SetLang(v string) *DescribeTagListRequest
	GetLang() *string
	SetRegId(v string) *DescribeTagListRequest
	GetRegId() *string
}

type DescribeTagListRequest struct {
	// The primary key ID.
	//
	// example:
	//
	// 433102
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
}

func (s DescribeTagListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagListRequest) GetId() *string {
	return s.Id
}

func (s *DescribeTagListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTagListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeTagListRequest) SetId(v string) *DescribeTagListRequest {
	s.Id = &v
	return s
}

func (s *DescribeTagListRequest) SetLang(v string) *DescribeTagListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTagListRequest) SetRegId(v string) *DescribeTagListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeTagListRequest) Validate() error {
	return dara.Validate(s)
}
