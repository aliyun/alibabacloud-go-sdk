// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVoiceCloneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteVoiceCloneResponseBody
	GetCode() *string
	SetData(v string) *DeleteVoiceCloneResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteVoiceCloneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteVoiceCloneResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteVoiceCloneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteVoiceCloneResponseBody
	GetSuccess() *bool
}

type DeleteVoiceCloneResponseBody struct {
	// The API status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the deleted voice.
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
	// The API message.
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

func (s DeleteVoiceCloneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVoiceCloneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVoiceCloneResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteVoiceCloneResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteVoiceCloneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteVoiceCloneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteVoiceCloneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVoiceCloneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteVoiceCloneResponseBody) SetCode(v string) *DeleteVoiceCloneResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVoiceCloneResponseBody) SetData(v string) *DeleteVoiceCloneResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteVoiceCloneResponseBody) SetHttpStatusCode(v int32) *DeleteVoiceCloneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteVoiceCloneResponseBody) SetMessage(v string) *DeleteVoiceCloneResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVoiceCloneResponseBody) SetRequestId(v string) *DeleteVoiceCloneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVoiceCloneResponseBody) SetSuccess(v bool) *DeleteVoiceCloneResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteVoiceCloneResponseBody) Validate() error {
	return dara.Validate(s)
}
