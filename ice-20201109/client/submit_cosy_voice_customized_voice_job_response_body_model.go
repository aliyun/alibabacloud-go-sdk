// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCosyVoiceCustomizedVoiceJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitCosyVoiceCustomizedVoiceJobResponseBodyData) *SubmitCosyVoiceCustomizedVoiceJobResponseBody
	GetData() *SubmitCosyVoiceCustomizedVoiceJobResponseBodyData
	SetRequestId(v string) *SubmitCosyVoiceCustomizedVoiceJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitCosyVoiceCustomizedVoiceJobResponseBody
	GetSuccess() *bool
}

type SubmitCosyVoiceCustomizedVoiceJobResponseBody struct {
	// The returned result.
	Data *SubmitCosyVoiceCustomizedVoiceJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// C0C02296-113C-5838-8FE9-8F3A32998DDC
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

func (s SubmitCosyVoiceCustomizedVoiceJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitCosyVoiceCustomizedVoiceJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCosyVoiceCustomizedVoiceJobResponseBody) GetData() *SubmitCosyVoiceCustomizedVoiceJobResponseBodyData {
	return s.Data
}

func (s *SubmitCosyVoiceCustomizedVoiceJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitCosyVoiceCustomizedVoiceJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitCosyVoiceCustomizedVoiceJobResponseBody) SetData(v *SubmitCosyVoiceCustomizedVoiceJobResponseBodyData) *SubmitCosyVoiceCustomizedVoiceJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitCosyVoiceCustomizedVoiceJobResponseBody) SetRequestId(v string) *SubmitCosyVoiceCustomizedVoiceJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitCosyVoiceCustomizedVoiceJobResponseBody) SetSuccess(v bool) *SubmitCosyVoiceCustomizedVoiceJobResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitCosyVoiceCustomizedVoiceJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitCosyVoiceCustomizedVoiceJobResponseBodyData struct {
	// The ID of the voice cloning task.
	//
	// example:
	//
	// bfb786c639894f4d80648792021****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The voice ID.
	//
	// example:
	//
	// Voice-l*****
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
}

func (s SubmitCosyVoiceCustomizedVoiceJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitCosyVoiceCustomizedVoiceJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitCosyVoiceCustomizedVoiceJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *SubmitCosyVoiceCustomizedVoiceJobResponseBodyData) GetVoiceId() *string {
	return s.VoiceId
}

func (s *SubmitCosyVoiceCustomizedVoiceJobResponseBodyData) SetJobId(v string) *SubmitCosyVoiceCustomizedVoiceJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *SubmitCosyVoiceCustomizedVoiceJobResponseBodyData) SetVoiceId(v string) *SubmitCosyVoiceCustomizedVoiceJobResponseBodyData {
	s.VoiceId = &v
	return s
}

func (s *SubmitCosyVoiceCustomizedVoiceJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
