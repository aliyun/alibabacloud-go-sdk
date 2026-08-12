// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSecurityCheckServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartSecurityCheckServiceResponseBody
	GetCode() *string
	SetMessage(v string) *StartSecurityCheckServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartSecurityCheckServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartSecurityCheckServiceResponseBody
	GetSuccess() *bool
}

type StartSecurityCheckServiceResponseBody struct {
	// The status code.
	//
	// - **200**: Succeeded.
	//
	// - **Others (400, 500)**: Failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BEE90F8C-EDC2-5394-953B-D07A121612B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartSecurityCheckServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartSecurityCheckServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StartSecurityCheckServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartSecurityCheckServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartSecurityCheckServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartSecurityCheckServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartSecurityCheckServiceResponseBody) SetCode(v string) *StartSecurityCheckServiceResponseBody {
	s.Code = &v
	return s
}

func (s *StartSecurityCheckServiceResponseBody) SetMessage(v string) *StartSecurityCheckServiceResponseBody {
	s.Message = &v
	return s
}

func (s *StartSecurityCheckServiceResponseBody) SetRequestId(v string) *StartSecurityCheckServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSecurityCheckServiceResponseBody) SetSuccess(v bool) *StartSecurityCheckServiceResponseBody {
	s.Success = &v
	return s
}

func (s *StartSecurityCheckServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
