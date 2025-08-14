// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckExpressionVariableLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CheckExpressionVariableLimitRequest
	GetLang() *string
	SetRegId(v string) *CheckExpressionVariableLimitRequest
	GetRegId() *string
}

type CheckExpressionVariableLimitRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s CheckExpressionVariableLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckExpressionVariableLimitRequest) GoString() string {
	return s.String()
}

func (s *CheckExpressionVariableLimitRequest) GetLang() *string {
	return s.Lang
}

func (s *CheckExpressionVariableLimitRequest) GetRegId() *string {
	return s.RegId
}

func (s *CheckExpressionVariableLimitRequest) SetLang(v string) *CheckExpressionVariableLimitRequest {
	s.Lang = &v
	return s
}

func (s *CheckExpressionVariableLimitRequest) SetRegId(v string) *CheckExpressionVariableLimitRequest {
	s.RegId = &v
	return s
}

func (s *CheckExpressionVariableLimitRequest) Validate() error {
	return dara.Validate(s)
}
