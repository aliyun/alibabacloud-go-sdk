// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCustVariableLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CheckCustVariableLimitRequest
	GetLang() *string
	SetCreateType(v string) *CheckCustVariableLimitRequest
	GetCreateType() *string
	SetRegId(v string) *CheckCustVariableLimitRequest
	GetRegId() *string
}

type CheckCustVariableLimitRequest struct {
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

func (s CheckCustVariableLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckCustVariableLimitRequest) GoString() string {
	return s.String()
}

func (s *CheckCustVariableLimitRequest) GetLang() *string {
	return s.Lang
}

func (s *CheckCustVariableLimitRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *CheckCustVariableLimitRequest) GetRegId() *string {
	return s.RegId
}

func (s *CheckCustVariableLimitRequest) SetLang(v string) *CheckCustVariableLimitRequest {
	s.Lang = &v
	return s
}

func (s *CheckCustVariableLimitRequest) SetCreateType(v string) *CheckCustVariableLimitRequest {
	s.CreateType = &v
	return s
}

func (s *CheckCustVariableLimitRequest) SetRegId(v string) *CheckCustVariableLimitRequest {
	s.RegId = &v
	return s
}

func (s *CheckCustVariableLimitRequest) Validate() error {
	return dara.Validate(s)
}
