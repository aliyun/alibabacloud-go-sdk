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
}

type CreateCloneVoiceResponseBody struct {
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateCloneVoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// F132DDBA-66C4-5BD3-BF7E-9642FE877158
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *CreateCloneVoiceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCloneVoiceResponseBodyData struct {
	// example:
	//
	// 8ee1160a-6999-478f-8df6-f33ef21f27d5
	CloneVoiceId *string `json:"CloneVoiceId,omitempty" xml:"CloneVoiceId,omitempty"`
	// example:
	//
	// cosyvoice-v3-plus-voicebot2-3666e4bbb2b94832ac4f4107b5804c34
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
