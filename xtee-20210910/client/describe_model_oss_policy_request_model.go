// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelOssPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeModelOssPolicyRequest
	GetLang() *string
	SetRegId(v string) *DescribeModelOssPolicyRequest
	GetRegId() *string
}

type DescribeModelOssPolicyRequest struct {
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
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeModelOssPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOssPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeModelOssPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeModelOssPolicyRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeModelOssPolicyRequest) SetLang(v string) *DescribeModelOssPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeModelOssPolicyRequest) SetRegId(v string) *DescribeModelOssPolicyRequest {
	s.RegId = &v
	return s
}

func (s *DescribeModelOssPolicyRequest) Validate() error {
	return dara.Validate(s)
}
