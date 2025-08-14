// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRequestHitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRequestHitRequest
	GetLang() *string
	SetRegId(v string) *DescribeRequestHitRequest
	GetRegId() *string
	SetSRequestId(v string) *DescribeRequestHitRequest
	GetSRequestId() *string
}

type DescribeRequestHitRequest struct {
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
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60C97040-D5D5-4906-9522-B9B413730CAA
	SRequestId *string `json:"sRequestId,omitempty" xml:"sRequestId,omitempty"`
}

func (s DescribeRequestHitRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRequestHitRequest) GoString() string {
	return s.String()
}

func (s *DescribeRequestHitRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRequestHitRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeRequestHitRequest) GetSRequestId() *string {
	return s.SRequestId
}

func (s *DescribeRequestHitRequest) SetLang(v string) *DescribeRequestHitRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRequestHitRequest) SetRegId(v string) *DescribeRequestHitRequest {
	s.RegId = &v
	return s
}

func (s *DescribeRequestHitRequest) SetSRequestId(v string) *DescribeRequestHitRequest {
	s.SRequestId = &v
	return s
}

func (s *DescribeRequestHitRequest) Validate() error {
	return dara.Validate(s)
}
