// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomizedVoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetCustomizedVoiceResponseBodyData) *GetCustomizedVoiceResponseBody
	GetData() *GetCustomizedVoiceResponseBodyData
	SetRequestId(v string) *GetCustomizedVoiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCustomizedVoiceResponseBody
	GetSuccess() *bool
}

type GetCustomizedVoiceResponseBody struct {
	// The data returned.
	Data *GetCustomizedVoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCustomizedVoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomizedVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomizedVoiceResponseBody) GetData() *GetCustomizedVoiceResponseBodyData {
	return s.Data
}

func (s *GetCustomizedVoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomizedVoiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCustomizedVoiceResponseBody) SetData(v *GetCustomizedVoiceResponseBodyData) *GetCustomizedVoiceResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomizedVoiceResponseBody) SetRequestId(v string) *GetCustomizedVoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomizedVoiceResponseBody) SetSuccess(v bool) *GetCustomizedVoiceResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomizedVoiceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomizedVoiceResponseBodyData struct {
	// The personalized human voice.
	CustomizedVoice *GetCustomizedVoiceResponseBodyDataCustomizedVoice `json:"CustomizedVoice,omitempty" xml:"CustomizedVoice,omitempty" type:"Struct"`
}

func (s GetCustomizedVoiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCustomizedVoiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomizedVoiceResponseBodyData) GetCustomizedVoice() *GetCustomizedVoiceResponseBodyDataCustomizedVoice {
	return s.CustomizedVoice
}

func (s *GetCustomizedVoiceResponseBodyData) SetCustomizedVoice(v *GetCustomizedVoiceResponseBodyDataCustomizedVoice) *GetCustomizedVoiceResponseBodyData {
	s.CustomizedVoice = v
	return s
}

func (s *GetCustomizedVoiceResponseBodyData) Validate() error {
	if s.CustomizedVoice != nil {
		if err := s.CustomizedVoice.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomizedVoiceResponseBodyDataCustomizedVoice struct {
	// The media asset ID of the sample audio file.
	//
	// example:
	//
	// ****42d3c312402982be65975f5b****
	DemoAudioMediaId *string `json:"DemoAudioMediaId,omitempty" xml:"DemoAudioMediaId,omitempty"`
	// The gender. Valid values:
	//
	// 	- female
	//
	// 	- male
	//
	// example:
	//
	// female
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// The demonstration scenario.
	//
	// Valid values:
	//
	// 	- **story**
	//
	// 	- **interaction**
	//
	// 	- **navigation**
	//
	// example:
	//
	// interaction
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// The voice description.
	VoiceDesc *string `json:"VoiceDesc,omitempty" xml:"VoiceDesc,omitempty"`
	// The voice ID.
	//
	// example:
	//
	// xiaozhuan
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
	// The voice name.
	VoiceName *string `json:"VoiceName,omitempty" xml:"VoiceName,omitempty"`
}

func (s GetCustomizedVoiceResponseBodyDataCustomizedVoice) String() string {
	return dara.Prettify(s)
}

func (s GetCustomizedVoiceResponseBodyDataCustomizedVoice) GoString() string {
	return s.String()
}

func (s *GetCustomizedVoiceResponseBodyDataCustomizedVoice) GetDemoAudioMediaId() *string {
	return s.DemoAudioMediaId
}

func (s *GetCustomizedVoiceResponseBodyDataCustomizedVoice) GetGender() *string {
	return s.Gender
}

func (s *GetCustomizedVoiceResponseBodyDataCustomizedVoice) GetScenario() *string {
	return s.Scenario
}

func (s *GetCustomizedVoiceResponseBodyDataCustomizedVoice) GetVoiceDesc() *string {
	return s.VoiceDesc
}

func (s *GetCustomizedVoiceResponseBodyDataCustomizedVoice) GetVoiceId() *string {
	return s.VoiceId
}

func (s *GetCustomizedVoiceResponseBodyDataCustomizedVoice) GetVoiceName() *string {
	return s.VoiceName
}

func (s *GetCustomizedVoiceResponseBodyDataCustomizedVoice) SetDemoAudioMediaId(v string) *GetCustomizedVoiceResponseBodyDataCustomizedVoice {
	s.DemoAudioMediaId = &v
	return s
}

func (s *GetCustomizedVoiceResponseBodyDataCustomizedVoice) SetGender(v string) *GetCustomizedVoiceResponseBodyDataCustomizedVoice {
	s.Gender = &v
	return s
}

func (s *GetCustomizedVoiceResponseBodyDataCustomizedVoice) SetScenario(v string) *GetCustomizedVoiceResponseBodyDataCustomizedVoice {
	s.Scenario = &v
	return s
}

func (s *GetCustomizedVoiceResponseBodyDataCustomizedVoice) SetVoiceDesc(v string) *GetCustomizedVoiceResponseBodyDataCustomizedVoice {
	s.VoiceDesc = &v
	return s
}

func (s *GetCustomizedVoiceResponseBodyDataCustomizedVoice) SetVoiceId(v string) *GetCustomizedVoiceResponseBodyDataCustomizedVoice {
	s.VoiceId = &v
	return s
}

func (s *GetCustomizedVoiceResponseBodyDataCustomizedVoice) SetVoiceName(v string) *GetCustomizedVoiceResponseBodyDataCustomizedVoice {
	s.VoiceName = &v
	return s
}

func (s *GetCustomizedVoiceResponseBodyDataCustomizedVoice) Validate() error {
	return dara.Validate(s)
}
