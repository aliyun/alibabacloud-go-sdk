// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloneVoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCloneVoiceResponseBody
	GetCode() *string
	SetData(v *CreateCloneVoiceResponseBodyData) *CreateCloneVoiceResponseBody
	GetData() *CreateCloneVoiceResponseBodyData
	SetHttpStatusCode(v int32) *CreateCloneVoiceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateCloneVoiceResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateCloneVoiceResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateCloneVoiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCloneVoiceResponseBody
	GetSuccess() *bool
}

type CreateCloneVoiceResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *CreateCloneVoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Instance does not exist. Instance=outb003.
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

func (s CreateCloneVoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloneVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloneVoiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCloneVoiceResponseBody) GetData() *CreateCloneVoiceResponseBodyData {
	return s.Data
}

func (s *CreateCloneVoiceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateCloneVoiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCloneVoiceResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateCloneVoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloneVoiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCloneVoiceResponseBody) SetCode(v string) *CreateCloneVoiceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCloneVoiceResponseBody) SetData(v *CreateCloneVoiceResponseBodyData) *CreateCloneVoiceResponseBody {
	s.Data = v
	return s
}

func (s *CreateCloneVoiceResponseBody) SetHttpStatusCode(v int32) *CreateCloneVoiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCloneVoiceResponseBody) SetMessage(v string) *CreateCloneVoiceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCloneVoiceResponseBody) SetParams(v []*string) *CreateCloneVoiceResponseBody {
	s.Params = v
	return s
}

func (s *CreateCloneVoiceResponseBody) SetRequestId(v string) *CreateCloneVoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloneVoiceResponseBody) SetSuccess(v bool) *CreateCloneVoiceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCloneVoiceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCloneVoiceResponseBodyData struct {
	// The UUID of the cloned voice.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b15
	CloneVoiceId *string `json:"CloneVoiceId,omitempty" xml:"CloneVoiceId,omitempty"`
	// The voice ID.
	//
	// example:
	//
	// cosyvoice-v3-flash-voicebot2-8aa485413eba42089c873eec1f901d64
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
}

func (s CreateCloneVoiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCloneVoiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCloneVoiceResponseBodyData) GetCloneVoiceId() *string {
	return s.CloneVoiceId
}

func (s *CreateCloneVoiceResponseBodyData) GetVoice() *string {
	return s.Voice
}

func (s *CreateCloneVoiceResponseBodyData) SetCloneVoiceId(v string) *CreateCloneVoiceResponseBodyData {
	s.CloneVoiceId = &v
	return s
}

func (s *CreateCloneVoiceResponseBodyData) SetVoice(v string) *CreateCloneVoiceResponseBodyData {
	s.Voice = &v
	return s
}

func (s *CreateCloneVoiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
