// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloneVoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateCloneVoiceResponseBody
	GetCode() *string
	SetData(v string) *UpdateCloneVoiceResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateCloneVoiceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateCloneVoiceResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateCloneVoiceResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateCloneVoiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCloneVoiceResponseBody
	GetSuccess() *bool
}

type UpdateCloneVoiceResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The UUID of the cloned voice.
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
	// Instance does not exist. Instance=anchashi.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 019FDAC7-13C5-1B64-A853-999DF105B9EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCloneVoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloneVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloneVoiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateCloneVoiceResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateCloneVoiceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateCloneVoiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCloneVoiceResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateCloneVoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloneVoiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCloneVoiceResponseBody) SetCode(v string) *UpdateCloneVoiceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCloneVoiceResponseBody) SetData(v string) *UpdateCloneVoiceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateCloneVoiceResponseBody) SetHttpStatusCode(v int32) *UpdateCloneVoiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateCloneVoiceResponseBody) SetMessage(v string) *UpdateCloneVoiceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCloneVoiceResponseBody) SetParams(v []*string) *UpdateCloneVoiceResponseBody {
	s.Params = v
	return s
}

func (s *UpdateCloneVoiceResponseBody) SetRequestId(v string) *UpdateCloneVoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloneVoiceResponseBody) SetSuccess(v bool) *UpdateCloneVoiceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCloneVoiceResponseBody) Validate() error {
	return dara.Validate(s)
}
