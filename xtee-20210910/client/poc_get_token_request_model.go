// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPocGetTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *PocGetTokenRequest
	GetLang() *string
	SetRegId(v string) *PocGetTokenRequest
	GetRegId() *string
	SetServiceCode(v string) *PocGetTokenRequest
	GetServiceCode() *string
}

type PocGetTokenRequest struct {
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
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// Service code.
	//
	// This parameter is required.
	//
	// example:
	//
	// alinlp
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s PocGetTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s PocGetTokenRequest) GoString() string {
	return s.String()
}

func (s *PocGetTokenRequest) GetLang() *string {
	return s.Lang
}

func (s *PocGetTokenRequest) GetRegId() *string {
	return s.RegId
}

func (s *PocGetTokenRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *PocGetTokenRequest) SetLang(v string) *PocGetTokenRequest {
	s.Lang = &v
	return s
}

func (s *PocGetTokenRequest) SetRegId(v string) *PocGetTokenRequest {
	s.RegId = &v
	return s
}

func (s *PocGetTokenRequest) SetServiceCode(v string) *PocGetTokenRequest {
	s.ServiceCode = &v
	return s
}

func (s *PocGetTokenRequest) Validate() error {
	return dara.Validate(s)
}
