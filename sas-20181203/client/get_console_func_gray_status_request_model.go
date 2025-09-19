// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsoleFuncGrayStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCondition(v string) *GetConsoleFuncGrayStatusRequest
	GetCondition() *string
	SetLang(v string) *GetConsoleFuncGrayStatusRequest
	GetLang() *string
}

type GetConsoleFuncGrayStatusRequest struct {
	// Name of the function module.
	//
	// example:
	//
	// vpcConsoleSwitch
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// Set the language type for request and response messages. Default value: **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s GetConsoleFuncGrayStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConsoleFuncGrayStatusRequest) GoString() string {
	return s.String()
}

func (s *GetConsoleFuncGrayStatusRequest) GetCondition() *string {
	return s.Condition
}

func (s *GetConsoleFuncGrayStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *GetConsoleFuncGrayStatusRequest) SetCondition(v string) *GetConsoleFuncGrayStatusRequest {
	s.Condition = &v
	return s
}

func (s *GetConsoleFuncGrayStatusRequest) SetLang(v string) *GetConsoleFuncGrayStatusRequest {
	s.Lang = &v
	return s
}

func (s *GetConsoleFuncGrayStatusRequest) Validate() error {
	return dara.Validate(s)
}
