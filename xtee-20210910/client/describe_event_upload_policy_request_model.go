// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventUploadPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeEventUploadPolicyRequest
	GetLang() *string
	SetRegId(v string) *DescribeEventUploadPolicyRequest
	GetRegId() *string
}

type DescribeEventUploadPolicyRequest struct {
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

func (s DescribeEventUploadPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventUploadPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventUploadPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEventUploadPolicyRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeEventUploadPolicyRequest) SetLang(v string) *DescribeEventUploadPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventUploadPolicyRequest) SetRegId(v string) *DescribeEventUploadPolicyRequest {
	s.RegId = &v
	return s
}

func (s *DescribeEventUploadPolicyRequest) Validate() error {
	return dara.Validate(s)
}
