// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectAudioForCustomizedVoiceJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetectAudioForCustomizedVoiceJobResponseBodyData) *DetectAudioForCustomizedVoiceJobResponseBody
	GetData() *DetectAudioForCustomizedVoiceJobResponseBodyData
	SetRequestId(v string) *DetectAudioForCustomizedVoiceJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DetectAudioForCustomizedVoiceJobResponseBody
	GetSuccess() *bool
}

type DetectAudioForCustomizedVoiceJobResponseBody struct {
	// The data returned.
	Data *DetectAudioForCustomizedVoiceJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetectAudioForCustomizedVoiceJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectAudioForCustomizedVoiceJobResponseBody) GoString() string {
	return s.String()
}

func (s *DetectAudioForCustomizedVoiceJobResponseBody) GetData() *DetectAudioForCustomizedVoiceJobResponseBodyData {
	return s.Data
}

func (s *DetectAudioForCustomizedVoiceJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectAudioForCustomizedVoiceJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DetectAudioForCustomizedVoiceJobResponseBody) SetData(v *DetectAudioForCustomizedVoiceJobResponseBodyData) *DetectAudioForCustomizedVoiceJobResponseBody {
	s.Data = v
	return s
}

func (s *DetectAudioForCustomizedVoiceJobResponseBody) SetRequestId(v string) *DetectAudioForCustomizedVoiceJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectAudioForCustomizedVoiceJobResponseBody) SetSuccess(v bool) *DetectAudioForCustomizedVoiceJobResponseBody {
	s.Success = &v
	return s
}

func (s *DetectAudioForCustomizedVoiceJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type DetectAudioForCustomizedVoiceJobResponseBodyData struct {
	// Indicates whether the audio file passes the check. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Pass *bool `json:"Pass,omitempty" xml:"Pass,omitempty"`
	// The reason returned if the audio file failed to pass the check.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s DetectAudioForCustomizedVoiceJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetectAudioForCustomizedVoiceJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectAudioForCustomizedVoiceJobResponseBodyData) GetPass() *bool {
	return s.Pass
}

func (s *DetectAudioForCustomizedVoiceJobResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *DetectAudioForCustomizedVoiceJobResponseBodyData) SetPass(v bool) *DetectAudioForCustomizedVoiceJobResponseBodyData {
	s.Pass = &v
	return s
}

func (s *DetectAudioForCustomizedVoiceJobResponseBodyData) SetReason(v string) *DetectAudioForCustomizedVoiceJobResponseBodyData {
	s.Reason = &v
	return s
}

func (s *DetectAudioForCustomizedVoiceJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
