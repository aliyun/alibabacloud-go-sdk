// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranscript interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *Transcript
	GetContent() *string
	SetSpeakerId(v string) *Transcript
	GetSpeakerId() *string
	SetTimeRange(v []*int64) *Transcript
	GetTimeRange() []*int64
	SetType(v string) *Transcript
	GetType() *string
}

type Transcript struct {
	Content   *string  `json:"Content,omitempty" xml:"Content,omitempty"`
	SpeakerId *string  `json:"SpeakerId,omitempty" xml:"SpeakerId,omitempty"`
	TimeRange []*int64 `json:"TimeRange,omitempty" xml:"TimeRange,omitempty" type:"Repeated"`
	Type      *string  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s Transcript) String() string {
	return dara.Prettify(s)
}

func (s Transcript) GoString() string {
	return s.String()
}

func (s *Transcript) GetContent() *string {
	return s.Content
}

func (s *Transcript) GetSpeakerId() *string {
	return s.SpeakerId
}

func (s *Transcript) GetTimeRange() []*int64 {
	return s.TimeRange
}

func (s *Transcript) GetType() *string {
	return s.Type
}

func (s *Transcript) SetContent(v string) *Transcript {
	s.Content = &v
	return s
}

func (s *Transcript) SetSpeakerId(v string) *Transcript {
	s.SpeakerId = &v
	return s
}

func (s *Transcript) SetTimeRange(v []*int64) *Transcript {
	s.TimeRange = v
	return s
}

func (s *Transcript) SetType(v string) *Transcript {
	s.Type = &v
	return s
}

func (s *Transcript) Validate() error {
	return dara.Validate(s)
}
