// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableTokenResponseBody
	GetCode() *string
	SetDisabled(v bool) *DisableTokenResponseBody
	GetDisabled() *bool
	SetMessage(v string) *DisableTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableTokenResponseBody
	GetRequestId() *string
}

type DisableTokenResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Indicates whether the token is disabled. Valid values:
	//
	// - true: Disabled.
	//
	// - false: Not disabled.
	//
	// example:
	//
	// true
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DisableTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DisableTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableTokenResponseBody) GetDisabled() *bool {
	return s.Disabled
}

func (s *DisableTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableTokenResponseBody) SetCode(v string) *DisableTokenResponseBody {
	s.Code = &v
	return s
}

func (s *DisableTokenResponseBody) SetDisabled(v bool) *DisableTokenResponseBody {
	s.Disabled = &v
	return s
}

func (s *DisableTokenResponseBody) SetMessage(v string) *DisableTokenResponseBody {
	s.Message = &v
	return s
}

func (s *DisableTokenResponseBody) SetRequestId(v string) *DisableTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
