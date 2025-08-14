// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustVariableDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeCustVariableDetailRequest
	GetLang() *string
	SetId(v int64) *DescribeCustVariableDetailRequest
	GetId() *int64
	SetRegId(v string) *DescribeCustVariableDetailRequest
	GetRegId() *string
}

type DescribeCustVariableDetailRequest struct {
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
	// Primary key ID of the cumulative variable
	//
	// This parameter is required.
	//
	// example:
	//
	// 2793
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeCustVariableDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustVariableDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustVariableDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCustVariableDetailRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeCustVariableDetailRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeCustVariableDetailRequest) SetLang(v string) *DescribeCustVariableDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCustVariableDetailRequest) SetId(v int64) *DescribeCustVariableDetailRequest {
	s.Id = &v
	return s
}

func (s *DescribeCustVariableDetailRequest) SetRegId(v string) *DescribeCustVariableDetailRequest {
	s.RegId = &v
	return s
}

func (s *DescribeCustVariableDetailRequest) Validate() error {
	return dara.Validate(s)
}
