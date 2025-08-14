// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsedServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeUsedServiceRequest
	GetLang() *string
	SetRegId(v string) *DescribeUsedServiceRequest
	GetRegId() *string
}

type DescribeUsedServiceRequest struct {
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

func (s DescribeUsedServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsedServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsedServiceRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeUsedServiceRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeUsedServiceRequest) SetLang(v string) *DescribeUsedServiceRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUsedServiceRequest) SetRegId(v string) *DescribeUsedServiceRequest {
	s.RegId = &v
	return s
}

func (s *DescribeUsedServiceRequest) Validate() error {
	return dara.Validate(s)
}
