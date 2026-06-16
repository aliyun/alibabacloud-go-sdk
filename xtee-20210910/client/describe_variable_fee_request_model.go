// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableFeeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeVariableFeeRequest
	GetLang() *string
	SetIds(v []*int64) *DescribeVariableFeeRequest
	GetIds() []*int64
	SetRegId(v string) *DescribeVariableFeeRequest
	GetRegId() *string
}

type DescribeVariableFeeRequest struct {
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The variable ID.
	//
	// This parameter is required.
	Ids []*int64 `json:"ids,omitempty" xml:"ids,omitempty" type:"Repeated"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeVariableFeeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableFeeRequest) GoString() string {
	return s.String()
}

func (s *DescribeVariableFeeRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVariableFeeRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *DescribeVariableFeeRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeVariableFeeRequest) SetLang(v string) *DescribeVariableFeeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVariableFeeRequest) SetIds(v []*int64) *DescribeVariableFeeRequest {
	s.Ids = v
	return s
}

func (s *DescribeVariableFeeRequest) SetRegId(v string) *DescribeVariableFeeRequest {
	s.RegId = &v
	return s
}

func (s *DescribeVariableFeeRequest) Validate() error {
	return dara.Validate(s)
}
