// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableSubscriptionResponseBody
	GetCode() *string
	SetData(v string) *DisableSubscriptionResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DisableSubscriptionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DisableSubscriptionResponseBody
	GetMessage() *string
	SetParams(v []*string) *DisableSubscriptionResponseBody
	GetParams() []*string
	SetRequestId(v string) *DisableSubscriptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableSubscriptionResponseBody
	GetSuccess() *bool
}

type DisableSubscriptionResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the event push configuration.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b15
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Instance does not exist. Instance=ob-1234567890
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSubscriptionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableSubscriptionResponseBody) GetData() *string {
	return s.Data
}

func (s *DisableSubscriptionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DisableSubscriptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableSubscriptionResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DisableSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableSubscriptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableSubscriptionResponseBody) SetCode(v string) *DisableSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *DisableSubscriptionResponseBody) SetData(v string) *DisableSubscriptionResponseBody {
	s.Data = &v
	return s
}

func (s *DisableSubscriptionResponseBody) SetHttpStatusCode(v int32) *DisableSubscriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisableSubscriptionResponseBody) SetMessage(v string) *DisableSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *DisableSubscriptionResponseBody) SetParams(v []*string) *DisableSubscriptionResponseBody {
	s.Params = v
	return s
}

func (s *DisableSubscriptionResponseBody) SetRequestId(v string) *DisableSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSubscriptionResponseBody) SetSuccess(v bool) *DisableSubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *DisableSubscriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
