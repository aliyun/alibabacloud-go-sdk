// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityCheckResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSecurityCheckResultResponseBody
	GetCode() *string
	SetData(v bool) *UpdateSecurityCheckResultResponseBody
	GetData() *bool
	SetMessage(v string) *UpdateSecurityCheckResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSecurityCheckResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSecurityCheckResultResponseBody
	GetSuccess() *bool
}

type UpdateSecurityCheckResultResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6B57D35D-9DAC-5393-AE39-07697E37C2E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSecurityCheckResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecurityCheckResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSecurityCheckResultResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateSecurityCheckResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSecurityCheckResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSecurityCheckResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSecurityCheckResultResponseBody) SetCode(v string) *UpdateSecurityCheckResultResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSecurityCheckResultResponseBody) SetData(v bool) *UpdateSecurityCheckResultResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateSecurityCheckResultResponseBody) SetMessage(v string) *UpdateSecurityCheckResultResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSecurityCheckResultResponseBody) SetRequestId(v string) *UpdateSecurityCheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSecurityCheckResultResponseBody) SetSuccess(v bool) *UpdateSecurityCheckResultResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSecurityCheckResultResponseBody) Validate() error {
	return dara.Validate(s)
}
