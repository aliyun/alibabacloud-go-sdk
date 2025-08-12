// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyHostInfoResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyHostInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyHostInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyHostInfoResponseBody
	GetSuccess() *bool
}

type ModifyHostInfoResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the call was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EBB5215C-44AB-4000-A2D7-48634FDC4F04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. The value true indicates a success. The value false indicates a failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyHostInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHostInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyHostInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyHostInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHostInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyHostInfoResponseBody) SetCode(v string) *ModifyHostInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyHostInfoResponseBody) SetMessage(v string) *ModifyHostInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyHostInfoResponseBody) SetRequestId(v string) *ModifyHostInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHostInfoResponseBody) SetSuccess(v bool) *ModifyHostInfoResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyHostInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
