// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUsageVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CheckUsageVariableRequest
	GetLang() *string
	SetId(v int64) *CheckUsageVariableRequest
	GetId() *int64
	SetRegId(v string) *CheckUsageVariableRequest
	GetRegId() *string
}

type CheckUsageVariableRequest struct {
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
	// Primary Key ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Region ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s CheckUsageVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckUsageVariableRequest) GoString() string {
	return s.String()
}

func (s *CheckUsageVariableRequest) GetLang() *string {
	return s.Lang
}

func (s *CheckUsageVariableRequest) GetId() *int64 {
	return s.Id
}

func (s *CheckUsageVariableRequest) GetRegId() *string {
	return s.RegId
}

func (s *CheckUsageVariableRequest) SetLang(v string) *CheckUsageVariableRequest {
	s.Lang = &v
	return s
}

func (s *CheckUsageVariableRequest) SetId(v int64) *CheckUsageVariableRequest {
	s.Id = &v
	return s
}

func (s *CheckUsageVariableRequest) SetRegId(v string) *CheckUsageVariableRequest {
	s.RegId = &v
	return s
}

func (s *CheckUsageVariableRequest) Validate() error {
	return dara.Validate(s)
}
