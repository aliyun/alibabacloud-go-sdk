// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIJobList(v *SubmitAIJobResponseBodyAIJobList) *SubmitAIJobResponseBody
	GetAIJobList() *SubmitAIJobResponseBodyAIJobList
	SetRequestId(v string) *SubmitAIJobResponseBody
	GetRequestId() *string
}

type SubmitAIJobResponseBody struct {
	// The information about the AI jobs.
	AIJobList *SubmitAIJobResponseBodyAIJobList `json:"AIJobList,omitempty" xml:"AIJobList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-D73936****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAIJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAIJobResponseBody) GetAIJobList() *SubmitAIJobResponseBodyAIJobList {
	return s.AIJobList
}

func (s *SubmitAIJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAIJobResponseBody) SetAIJobList(v *SubmitAIJobResponseBodyAIJobList) *SubmitAIJobResponseBody {
	s.AIJobList = v
	return s
}

func (s *SubmitAIJobResponseBody) SetRequestId(v string) *SubmitAIJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAIJobResponseBody) Validate() error {
	if s.AIJobList != nil {
		if err := s.AIJobList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitAIJobResponseBodyAIJobList struct {
	AIJob []*SubmitAIJobResponseBodyAIJobListAIJob `json:"AIJob,omitempty" xml:"AIJob,omitempty" type:"Repeated"`
}

func (s SubmitAIJobResponseBodyAIJobList) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIJobResponseBodyAIJobList) GoString() string {
	return s.String()
}

func (s *SubmitAIJobResponseBodyAIJobList) GetAIJob() []*SubmitAIJobResponseBodyAIJobListAIJob {
	return s.AIJob
}

func (s *SubmitAIJobResponseBodyAIJobList) SetAIJob(v []*SubmitAIJobResponseBodyAIJobListAIJob) *SubmitAIJobResponseBodyAIJobList {
	s.AIJob = v
	return s
}

func (s *SubmitAIJobResponseBodyAIJobList) Validate() error {
	if s.AIJob != nil {
		for _, item := range s.AIJob {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitAIJobResponseBodyAIJobListAIJob struct {
	// The ID of the AI job.
	//
	// example:
	//
	// 9e82640c85114bf5af23edfaf****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the video.
	//
	// example:
	//
	// 3D3D12340d92c641401fab46a0b847****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The type of the AI job. Valid values:
	//
	// 	- **AIMediaDNA**: The media fingerprinting job.
	//
	// 	- **AIVideoTag**: The smart tagging job.
	//
	// example:
	//
	// AIVideoTag
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitAIJobResponseBodyAIJobListAIJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIJobResponseBodyAIJobListAIJob) GoString() string {
	return s.String()
}

func (s *SubmitAIJobResponseBodyAIJobListAIJob) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAIJobResponseBodyAIJobListAIJob) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIJobResponseBodyAIJobListAIJob) GetType() *string {
	return s.Type
}

func (s *SubmitAIJobResponseBodyAIJobListAIJob) SetJobId(v string) *SubmitAIJobResponseBodyAIJobListAIJob {
	s.JobId = &v
	return s
}

func (s *SubmitAIJobResponseBodyAIJobListAIJob) SetMediaId(v string) *SubmitAIJobResponseBodyAIJobListAIJob {
	s.MediaId = &v
	return s
}

func (s *SubmitAIJobResponseBodyAIJobListAIJob) SetType(v string) *SubmitAIJobResponseBodyAIJobListAIJob {
	s.Type = &v
	return s
}

func (s *SubmitAIJobResponseBodyAIJobListAIJob) Validate() error {
	return dara.Validate(s)
}
