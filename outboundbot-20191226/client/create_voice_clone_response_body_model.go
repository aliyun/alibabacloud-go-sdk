// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVoiceCloneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateVoiceCloneResponseBody
	GetCode() *string
	SetData(v string) *CreateVoiceCloneResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateVoiceCloneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateVoiceCloneResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateVoiceCloneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateVoiceCloneResponseBody
	GetSuccess() *bool
}

type CreateVoiceCloneResponseBody struct {
	// The API status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the created cloned voice.
	//
	// example:
	//
	// cosyvoice-v2-067dc4b2e9d34dbeb5de20b5b093d1e3
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned by the operation.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7A573837-3AD3-54CF-930A-07A3287042C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateVoiceCloneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVoiceCloneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVoiceCloneResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateVoiceCloneResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateVoiceCloneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateVoiceCloneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVoiceCloneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVoiceCloneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateVoiceCloneResponseBody) SetCode(v string) *CreateVoiceCloneResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVoiceCloneResponseBody) SetData(v string) *CreateVoiceCloneResponseBody {
	s.Data = &v
	return s
}

func (s *CreateVoiceCloneResponseBody) SetHttpStatusCode(v int32) *CreateVoiceCloneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateVoiceCloneResponseBody) SetMessage(v string) *CreateVoiceCloneResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVoiceCloneResponseBody) SetRequestId(v string) *CreateVoiceCloneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVoiceCloneResponseBody) SetSuccess(v bool) *CreateVoiceCloneResponseBody {
	s.Success = &v
	return s
}

func (s *CreateVoiceCloneResponseBody) Validate() error {
	return dara.Validate(s)
}
