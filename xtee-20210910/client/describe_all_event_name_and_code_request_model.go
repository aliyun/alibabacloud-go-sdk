// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllEventNameAndCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAllEventNameAndCodeRequest
	GetLang() *string
	SetCreateType(v string) *DescribeAllEventNameAndCodeRequest
	GetCreateType() *string
	SetRegId(v string) *DescribeAllEventNameAndCodeRequest
	GetRegId() *string
}

type DescribeAllEventNameAndCodeRequest struct {
	// Sets the language type for the request and response messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Creation type
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeAllEventNameAndCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllEventNameAndCodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllEventNameAndCodeRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAllEventNameAndCodeRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeAllEventNameAndCodeRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeAllEventNameAndCodeRequest) SetLang(v string) *DescribeAllEventNameAndCodeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAllEventNameAndCodeRequest) SetCreateType(v string) *DescribeAllEventNameAndCodeRequest {
	s.CreateType = &v
	return s
}

func (s *DescribeAllEventNameAndCodeRequest) SetRegId(v string) *DescribeAllEventNameAndCodeRequest {
	s.RegId = &v
	return s
}

func (s *DescribeAllEventNameAndCodeRequest) Validate() error {
	return dara.Validate(s)
}
