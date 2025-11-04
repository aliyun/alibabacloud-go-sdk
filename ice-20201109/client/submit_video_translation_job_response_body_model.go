// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoTranslationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitVideoTranslationJobResponseBodyData) *SubmitVideoTranslationJobResponseBody
	GetData() *SubmitVideoTranslationJobResponseBodyData
	SetRequestId(v string) *SubmitVideoTranslationJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitVideoTranslationJobResponseBody
	GetSuccess() *bool
}

type SubmitVideoTranslationJobResponseBody struct {
	// The data returned.
	Data *SubmitVideoTranslationJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// Valid values:
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

func (s SubmitVideoTranslationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoTranslationJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitVideoTranslationJobResponseBody) GetData() *SubmitVideoTranslationJobResponseBodyData {
	return s.Data
}

func (s *SubmitVideoTranslationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitVideoTranslationJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitVideoTranslationJobResponseBody) SetData(v *SubmitVideoTranslationJobResponseBodyData) *SubmitVideoTranslationJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitVideoTranslationJobResponseBody) SetRequestId(v string) *SubmitVideoTranslationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitVideoTranslationJobResponseBody) SetSuccess(v bool) *SubmitVideoTranslationJobResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitVideoTranslationJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitVideoTranslationJobResponseBodyData struct {
	// The ID of the video translation job.
	//
	// example:
	//
	// ****d718e2ff4f018ccf419a7b71****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitVideoTranslationJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoTranslationJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitVideoTranslationJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *SubmitVideoTranslationJobResponseBodyData) SetJobId(v string) *SubmitVideoTranslationJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *SubmitVideoTranslationJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
