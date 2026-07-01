// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCustomizedVoiceJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitCustomizedVoiceJobResponseBodyData) *SubmitCustomizedVoiceJobResponseBody
	GetData() *SubmitCustomizedVoiceJobResponseBodyData
	SetRequestId(v string) *SubmitCustomizedVoiceJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitCustomizedVoiceJobResponseBody
	GetSuccess() *bool
}

type SubmitCustomizedVoiceJobResponseBody struct {
	// The response data.
	Data *SubmitCustomizedVoiceJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique identifier for the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitCustomizedVoiceJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomizedVoiceJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCustomizedVoiceJobResponseBody) GetData() *SubmitCustomizedVoiceJobResponseBodyData {
	return s.Data
}

func (s *SubmitCustomizedVoiceJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitCustomizedVoiceJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitCustomizedVoiceJobResponseBody) SetData(v *SubmitCustomizedVoiceJobResponseBodyData) *SubmitCustomizedVoiceJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitCustomizedVoiceJobResponseBody) SetRequestId(v string) *SubmitCustomizedVoiceJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitCustomizedVoiceJobResponseBody) SetSuccess(v bool) *SubmitCustomizedVoiceJobResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitCustomizedVoiceJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitCustomizedVoiceJobResponseBodyData struct {
	// The unique identifier for the voice cloning job.
	//
	// example:
	//
	// ****d718e2ff4f018ccf419a7b71****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The unique identifier for the voice.
	//
	// example:
	//
	// xiaozhuan
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
}

func (s SubmitCustomizedVoiceJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomizedVoiceJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitCustomizedVoiceJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *SubmitCustomizedVoiceJobResponseBodyData) GetVoiceId() *string {
	return s.VoiceId
}

func (s *SubmitCustomizedVoiceJobResponseBodyData) SetJobId(v string) *SubmitCustomizedVoiceJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *SubmitCustomizedVoiceJobResponseBodyData) SetVoiceId(v string) *SubmitCustomizedVoiceJobResponseBodyData {
	s.VoiceId = &v
	return s
}

func (s *SubmitCustomizedVoiceJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
