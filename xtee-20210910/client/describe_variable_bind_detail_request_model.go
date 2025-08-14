// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableBindDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeVariableBindDetailRequest
	GetLang() *string
	SetDefineId(v int64) *DescribeVariableBindDetailRequest
	GetDefineId() *int64
	SetId(v int64) *DescribeVariableBindDetailRequest
	GetId() *int64
	SetRegId(v string) *DescribeVariableBindDetailRequest
	GetRegId() *string
}

type DescribeVariableBindDetailRequest struct {
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
	// Bound variable definition ID
	//
	// example:
	//
	// 10
	DefineId *int64 `json:"defineId,omitempty" xml:"defineId,omitempty"`
	// Primary key ID of the variable, which is empty if it\\"s a new addition
	//
	// example:
	//
	// 3144
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeVariableBindDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableBindDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeVariableBindDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVariableBindDetailRequest) GetDefineId() *int64 {
	return s.DefineId
}

func (s *DescribeVariableBindDetailRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeVariableBindDetailRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeVariableBindDetailRequest) SetLang(v string) *DescribeVariableBindDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVariableBindDetailRequest) SetDefineId(v int64) *DescribeVariableBindDetailRequest {
	s.DefineId = &v
	return s
}

func (s *DescribeVariableBindDetailRequest) SetId(v int64) *DescribeVariableBindDetailRequest {
	s.Id = &v
	return s
}

func (s *DescribeVariableBindDetailRequest) SetRegId(v string) *DescribeVariableBindDetailRequest {
	s.RegId = &v
	return s
}

func (s *DescribeVariableBindDetailRequest) Validate() error {
	return dara.Validate(s)
}
