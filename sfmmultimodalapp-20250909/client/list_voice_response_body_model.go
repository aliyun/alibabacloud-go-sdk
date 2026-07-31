// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListVoiceResponseBody
	GetRequestId() *string
	SetVoiceList(v []*ListVoiceResponseBodyVoiceList) *ListVoiceResponseBody
	GetVoiceList() []*ListVoiceResponseBodyVoiceList
}

type ListVoiceResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VoiceList []*ListVoiceResponseBodyVoiceList `json:"VoiceList,omitempty" xml:"VoiceList,omitempty" type:"Repeated"`
}

func (s ListVoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *ListVoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVoiceResponseBody) GetVoiceList() []*ListVoiceResponseBodyVoiceList {
	return s.VoiceList
}

func (s *ListVoiceResponseBody) SetRequestId(v string) *ListVoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVoiceResponseBody) SetVoiceList(v []*ListVoiceResponseBodyVoiceList) *ListVoiceResponseBody {
	s.VoiceList = v
	return s
}

func (s *ListVoiceResponseBody) Validate() error {
	if s.VoiceList != nil {
		for _, item := range s.VoiceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVoiceResponseBodyVoiceList struct {
	Gender            *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Illustration      *string `json:"Illustration,omitempty" xml:"Illustration,omitempty"`
	IllustrationAudio *string `json:"IllustrationAudio,omitempty" xml:"IllustrationAudio,omitempty"`
	Language          *string `json:"Language,omitempty" xml:"Language,omitempty"`
	ModelId           *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Voice             *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
}

func (s ListVoiceResponseBodyVoiceList) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceResponseBodyVoiceList) GoString() string {
	return s.String()
}

func (s *ListVoiceResponseBodyVoiceList) GetGender() *string {
	return s.Gender
}

func (s *ListVoiceResponseBodyVoiceList) GetIllustration() *string {
	return s.Illustration
}

func (s *ListVoiceResponseBodyVoiceList) GetIllustrationAudio() *string {
	return s.IllustrationAudio
}

func (s *ListVoiceResponseBodyVoiceList) GetLanguage() *string {
	return s.Language
}

func (s *ListVoiceResponseBodyVoiceList) GetModelId() *string {
	return s.ModelId
}

func (s *ListVoiceResponseBodyVoiceList) GetName() *string {
	return s.Name
}

func (s *ListVoiceResponseBodyVoiceList) GetVoice() *string {
	return s.Voice
}

func (s *ListVoiceResponseBodyVoiceList) SetGender(v string) *ListVoiceResponseBodyVoiceList {
	s.Gender = &v
	return s
}

func (s *ListVoiceResponseBodyVoiceList) SetIllustration(v string) *ListVoiceResponseBodyVoiceList {
	s.Illustration = &v
	return s
}

func (s *ListVoiceResponseBodyVoiceList) SetIllustrationAudio(v string) *ListVoiceResponseBodyVoiceList {
	s.IllustrationAudio = &v
	return s
}

func (s *ListVoiceResponseBodyVoiceList) SetLanguage(v string) *ListVoiceResponseBodyVoiceList {
	s.Language = &v
	return s
}

func (s *ListVoiceResponseBodyVoiceList) SetModelId(v string) *ListVoiceResponseBodyVoiceList {
	s.ModelId = &v
	return s
}

func (s *ListVoiceResponseBodyVoiceList) SetName(v string) *ListVoiceResponseBodyVoiceList {
	s.Name = &v
	return s
}

func (s *ListVoiceResponseBodyVoiceList) SetVoice(v string) *ListVoiceResponseBodyVoiceList {
	s.Voice = &v
	return s
}

func (s *ListVoiceResponseBodyVoiceList) Validate() error {
	return dara.Validate(s)
}
