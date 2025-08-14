// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckFieldLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CheckFieldLimitRequest
	GetLang() *string
	SetRegId(v string) *CheckFieldLimitRequest
	GetRegId() *string
	SetSource(v string) *CheckFieldLimitRequest
	GetSource() *string
}

type CheckFieldLimitRequest struct {
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
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Source of the field
	//
	// example:
	//
	// DEFINE
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s CheckFieldLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckFieldLimitRequest) GoString() string {
	return s.String()
}

func (s *CheckFieldLimitRequest) GetLang() *string {
	return s.Lang
}

func (s *CheckFieldLimitRequest) GetRegId() *string {
	return s.RegId
}

func (s *CheckFieldLimitRequest) GetSource() *string {
	return s.Source
}

func (s *CheckFieldLimitRequest) SetLang(v string) *CheckFieldLimitRequest {
	s.Lang = &v
	return s
}

func (s *CheckFieldLimitRequest) SetRegId(v string) *CheckFieldLimitRequest {
	s.RegId = &v
	return s
}

func (s *CheckFieldLimitRequest) SetSource(v string) *CheckFieldLimitRequest {
	s.Source = &v
	return s
}

func (s *CheckFieldLimitRequest) Validate() error {
	return dara.Validate(s)
}
