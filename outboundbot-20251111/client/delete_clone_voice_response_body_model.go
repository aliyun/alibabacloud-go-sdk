// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloneVoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCloneVoiceResponseBody
	GetCode() *string
	SetData(v string) *DeleteCloneVoiceResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteCloneVoiceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteCloneVoiceResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteCloneVoiceResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteCloneVoiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCloneVoiceResponseBody
	GetSuccess() *bool
}

type DeleteCloneVoiceResponseBody struct {
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
	// Instance does not exist. Instance=outb003
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

func (s DeleteCloneVoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloneVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloneVoiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCloneVoiceResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteCloneVoiceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteCloneVoiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCloneVoiceResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteCloneVoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloneVoiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCloneVoiceResponseBody) SetCode(v string) *DeleteCloneVoiceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCloneVoiceResponseBody) SetData(v string) *DeleteCloneVoiceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCloneVoiceResponseBody) SetHttpStatusCode(v int32) *DeleteCloneVoiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteCloneVoiceResponseBody) SetMessage(v string) *DeleteCloneVoiceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCloneVoiceResponseBody) SetParams(v []*string) *DeleteCloneVoiceResponseBody {
	s.Params = v
	return s
}

func (s *DeleteCloneVoiceResponseBody) SetRequestId(v string) *DeleteCloneVoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloneVoiceResponseBody) SetSuccess(v bool) *DeleteCloneVoiceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCloneVoiceResponseBody) Validate() error {
	return dara.Validate(s)
}
