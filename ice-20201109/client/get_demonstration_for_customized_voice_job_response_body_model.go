// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDemonstrationForCustomizedVoiceJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDemonstrationForCustomizedVoiceJobResponseBodyData) *GetDemonstrationForCustomizedVoiceJobResponseBody
	GetData() *GetDemonstrationForCustomizedVoiceJobResponseBodyData
	SetRequestId(v string) *GetDemonstrationForCustomizedVoiceJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDemonstrationForCustomizedVoiceJobResponseBody
	GetSuccess() *bool
}

type GetDemonstrationForCustomizedVoiceJobResponseBody struct {
	// The data returned.
	Data *GetDemonstrationForCustomizedVoiceJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
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

func (s GetDemonstrationForCustomizedVoiceJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDemonstrationForCustomizedVoiceJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetDemonstrationForCustomizedVoiceJobResponseBody) GetData() *GetDemonstrationForCustomizedVoiceJobResponseBodyData {
	return s.Data
}

func (s *GetDemonstrationForCustomizedVoiceJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDemonstrationForCustomizedVoiceJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDemonstrationForCustomizedVoiceJobResponseBody) SetData(v *GetDemonstrationForCustomizedVoiceJobResponseBodyData) *GetDemonstrationForCustomizedVoiceJobResponseBody {
	s.Data = v
	return s
}

func (s *GetDemonstrationForCustomizedVoiceJobResponseBody) SetRequestId(v string) *GetDemonstrationForCustomizedVoiceJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDemonstrationForCustomizedVoiceJobResponseBody) SetSuccess(v bool) *GetDemonstrationForCustomizedVoiceJobResponseBody {
	s.Success = &v
	return s
}

func (s *GetDemonstrationForCustomizedVoiceJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDemonstrationForCustomizedVoiceJobResponseBodyData struct {
	// A list of 20 text entries to be read and the corresponding sample audio.
	DemonstrationList []*GetDemonstrationForCustomizedVoiceJobResponseBodyDataDemonstrationList `json:"DemonstrationList,omitempty" xml:"DemonstrationList,omitempty" type:"Repeated"`
}

func (s GetDemonstrationForCustomizedVoiceJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDemonstrationForCustomizedVoiceJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDemonstrationForCustomizedVoiceJobResponseBodyData) GetDemonstrationList() []*GetDemonstrationForCustomizedVoiceJobResponseBodyDataDemonstrationList {
	return s.DemonstrationList
}

func (s *GetDemonstrationForCustomizedVoiceJobResponseBodyData) SetDemonstrationList(v []*GetDemonstrationForCustomizedVoiceJobResponseBodyDataDemonstrationList) *GetDemonstrationForCustomizedVoiceJobResponseBodyData {
	s.DemonstrationList = v
	return s
}

func (s *GetDemonstrationForCustomizedVoiceJobResponseBodyData) Validate() error {
	if s.DemonstrationList != nil {
		for _, item := range s.DemonstrationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDemonstrationForCustomizedVoiceJobResponseBodyDataDemonstrationList struct {
	// The sequence number of the text, which corresponds to the AduioRecordId parameter to be passed during audio check.
	//
	// example:
	//
	// 2
	AudioId *int32 `json:"AudioId,omitempty" xml:"AudioId,omitempty"`
	// The URL of the sample audio.
	//
	// 	- The value is an Object Storage Service (OSS) URL.
	//
	//     **
	//
	//     **Note**: The URL expires in 12 hours.
	//
	// example:
	//
	// http://bucket.oss-cn-shanghai.aliyuncs.com/1.wav
	DemoAudio *string `json:"DemoAudio,omitempty" xml:"DemoAudio,omitempty"`
	// The text content to be read.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetDemonstrationForCustomizedVoiceJobResponseBodyDataDemonstrationList) String() string {
	return dara.Prettify(s)
}

func (s GetDemonstrationForCustomizedVoiceJobResponseBodyDataDemonstrationList) GoString() string {
	return s.String()
}

func (s *GetDemonstrationForCustomizedVoiceJobResponseBodyDataDemonstrationList) GetAudioId() *int32 {
	return s.AudioId
}

func (s *GetDemonstrationForCustomizedVoiceJobResponseBodyDataDemonstrationList) GetDemoAudio() *string {
	return s.DemoAudio
}

func (s *GetDemonstrationForCustomizedVoiceJobResponseBodyDataDemonstrationList) GetText() *string {
	return s.Text
}

func (s *GetDemonstrationForCustomizedVoiceJobResponseBodyDataDemonstrationList) SetAudioId(v int32) *GetDemonstrationForCustomizedVoiceJobResponseBodyDataDemonstrationList {
	s.AudioId = &v
	return s
}

func (s *GetDemonstrationForCustomizedVoiceJobResponseBodyDataDemonstrationList) SetDemoAudio(v string) *GetDemonstrationForCustomizedVoiceJobResponseBodyDataDemonstrationList {
	s.DemoAudio = &v
	return s
}

func (s *GetDemonstrationForCustomizedVoiceJobResponseBodyDataDemonstrationList) SetText(v string) *GetDemonstrationForCustomizedVoiceJobResponseBodyDataDemonstrationList {
	s.Text = &v
	return s
}

func (s *GetDemonstrationForCustomizedVoiceJobResponseBodyDataDemonstrationList) Validate() error {
	return dara.Validate(s)
}
