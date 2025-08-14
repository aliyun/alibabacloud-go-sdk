// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNameListLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeNameListLimitRequest
	GetLang() *string
	SetCreateType(v string) *DescribeNameListLimitRequest
	GetCreateType() *string
	SetRegId(v string) *DescribeNameListLimitRequest
	GetRegId() *string
}

type DescribeNameListLimitRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
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

func (s DescribeNameListLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListLimitRequest) GoString() string {
	return s.String()
}

func (s *DescribeNameListLimitRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeNameListLimitRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeNameListLimitRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeNameListLimitRequest) SetLang(v string) *DescribeNameListLimitRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNameListLimitRequest) SetCreateType(v string) *DescribeNameListLimitRequest {
	s.CreateType = &v
	return s
}

func (s *DescribeNameListLimitRequest) SetRegId(v string) *DescribeNameListLimitRequest {
	s.RegId = &v
	return s
}

func (s *DescribeNameListLimitRequest) Validate() error {
	return dara.Validate(s)
}
