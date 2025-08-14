// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFieldByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeFieldByIdRequest
	GetLang() *string
	SetId(v int64) *DescribeFieldByIdRequest
	GetId() *int64
	SetRegId(v string) *DescribeFieldByIdRequest
	GetRegId() *string
}

type DescribeFieldByIdRequest struct {
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
	// Primary key ID
	//
	// example:
	//
	// 223
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeFieldByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFieldByIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeFieldByIdRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeFieldByIdRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeFieldByIdRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeFieldByIdRequest) SetLang(v string) *DescribeFieldByIdRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFieldByIdRequest) SetId(v int64) *DescribeFieldByIdRequest {
	s.Id = &v
	return s
}

func (s *DescribeFieldByIdRequest) SetRegId(v string) *DescribeFieldByIdRequest {
	s.RegId = &v
	return s
}

func (s *DescribeFieldByIdRequest) Validate() error {
	return dara.Validate(s)
}
