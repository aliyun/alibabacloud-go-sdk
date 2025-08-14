// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitStandardCustomizedVoiceJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitStandardCustomizedVoiceJobResponseBodyData) *SubmitStandardCustomizedVoiceJobResponseBody
	GetData() *SubmitStandardCustomizedVoiceJobResponseBodyData
	SetRequestId(v string) *SubmitStandardCustomizedVoiceJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitStandardCustomizedVoiceJobResponseBody
	GetSuccess() *bool
}

type SubmitStandardCustomizedVoiceJobResponseBody struct {
	// The data returned.
	Data *SubmitStandardCustomizedVoiceJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SubmitStandardCustomizedVoiceJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitStandardCustomizedVoiceJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitStandardCustomizedVoiceJobResponseBody) GetData() *SubmitStandardCustomizedVoiceJobResponseBodyData {
	return s.Data
}

func (s *SubmitStandardCustomizedVoiceJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitStandardCustomizedVoiceJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitStandardCustomizedVoiceJobResponseBody) SetData(v *SubmitStandardCustomizedVoiceJobResponseBodyData) *SubmitStandardCustomizedVoiceJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitStandardCustomizedVoiceJobResponseBody) SetRequestId(v string) *SubmitStandardCustomizedVoiceJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitStandardCustomizedVoiceJobResponseBody) SetSuccess(v bool) *SubmitStandardCustomizedVoiceJobResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitStandardCustomizedVoiceJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitStandardCustomizedVoiceJobResponseBodyData struct {
	// The ID of the human voice cloning job.
	//
	// example:
	//
	// ****d718e2ff4f018ccf419a7b71****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitStandardCustomizedVoiceJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitStandardCustomizedVoiceJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitStandardCustomizedVoiceJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *SubmitStandardCustomizedVoiceJobResponseBodyData) SetJobId(v string) *SubmitStandardCustomizedVoiceJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *SubmitStandardCustomizedVoiceJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
