// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomizedVoiceJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateCustomizedVoiceJobResponseBodyData) *CreateCustomizedVoiceJobResponseBody
	GetData() *CreateCustomizedVoiceJobResponseBodyData
	SetRequestId(v string) *CreateCustomizedVoiceJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCustomizedVoiceJobResponseBody
	GetSuccess() *bool
}

type CreateCustomizedVoiceJobResponseBody struct {
	// The data returned.
	Data *CreateCustomizedVoiceJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CreateCustomizedVoiceJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomizedVoiceJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomizedVoiceJobResponseBody) GetData() *CreateCustomizedVoiceJobResponseBodyData {
	return s.Data
}

func (s *CreateCustomizedVoiceJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomizedVoiceJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCustomizedVoiceJobResponseBody) SetData(v *CreateCustomizedVoiceJobResponseBodyData) *CreateCustomizedVoiceJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateCustomizedVoiceJobResponseBody) SetRequestId(v string) *CreateCustomizedVoiceJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomizedVoiceJobResponseBody) SetSuccess(v bool) *CreateCustomizedVoiceJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCustomizedVoiceJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCustomizedVoiceJobResponseBodyData struct {
	// The ID of the human voice cloning job.
	//
	// example:
	//
	// ****29faef8144638ba42eb8e037****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The voice ID.
	//
	// example:
	//
	// xiaozhuan
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
}

func (s CreateCustomizedVoiceJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomizedVoiceJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCustomizedVoiceJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *CreateCustomizedVoiceJobResponseBodyData) GetVoiceId() *string {
	return s.VoiceId
}

func (s *CreateCustomizedVoiceJobResponseBodyData) SetJobId(v string) *CreateCustomizedVoiceJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreateCustomizedVoiceJobResponseBodyData) SetVoiceId(v string) *CreateCustomizedVoiceJobResponseBodyData {
	s.VoiceId = &v
	return s
}

func (s *CreateCustomizedVoiceJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
