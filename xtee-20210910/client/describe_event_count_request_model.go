// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeEventCountRequest
	GetLang() *string
	SetCreateType(v string) *DescribeEventCountRequest
	GetCreateType() *string
	SetRegId(v string) *DescribeEventCountRequest
	GetRegId() *string
}

type DescribeEventCountRequest struct {
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

func (s DescribeEventCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventCountRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEventCountRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeEventCountRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeEventCountRequest) SetLang(v string) *DescribeEventCountRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventCountRequest) SetCreateType(v string) *DescribeEventCountRequest {
	s.CreateType = &v
	return s
}

func (s *DescribeEventCountRequest) SetRegId(v string) *DescribeEventCountRequest {
	s.RegId = &v
	return s
}

func (s *DescribeEventCountRequest) Validate() error {
	return dara.Validate(s)
}
