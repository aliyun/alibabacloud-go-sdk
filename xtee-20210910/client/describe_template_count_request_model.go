// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeTemplateCountRequest
	GetLang() *string
	SetRegId(v string) *DescribeTemplateCountRequest
	GetRegId() *string
}

type DescribeTemplateCountRequest struct {
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

func (s DescribeTemplateCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplateCountRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTemplateCountRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeTemplateCountRequest) SetLang(v string) *DescribeTemplateCountRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTemplateCountRequest) SetRegId(v string) *DescribeTemplateCountRequest {
	s.RegId = &v
	return s
}

func (s *DescribeTemplateCountRequest) Validate() error {
	return dara.Validate(s)
}
