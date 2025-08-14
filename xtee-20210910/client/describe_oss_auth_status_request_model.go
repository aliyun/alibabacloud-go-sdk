// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssAuthStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeOssAuthStatusRequest
	GetLang() *string
	SetRegId(v string) *DescribeOssAuthStatusRequest
	GetRegId() *string
}

type DescribeOssAuthStatusRequest struct {
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

func (s DescribeOssAuthStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssAuthStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssAuthStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOssAuthStatusRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeOssAuthStatusRequest) SetLang(v string) *DescribeOssAuthStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOssAuthStatusRequest) SetRegId(v string) *DescribeOssAuthStatusRequest {
	s.RegId = &v
	return s
}

func (s *DescribeOssAuthStatusRequest) Validate() error {
	return dara.Validate(s)
}
