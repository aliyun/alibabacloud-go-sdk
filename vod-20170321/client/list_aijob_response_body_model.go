// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIJobList(v *ListAIJobResponseBodyAIJobList) *ListAIJobResponseBody
	GetAIJobList() *ListAIJobResponseBodyAIJobList
	SetNonExistAIJobIds(v *ListAIJobResponseBodyNonExistAIJobIds) *ListAIJobResponseBody
	GetNonExistAIJobIds() *ListAIJobResponseBodyNonExistAIJobIds
	SetRequestId(v string) *ListAIJobResponseBody
	GetRequestId() *string
}

type ListAIJobResponseBody struct {
	AIJobList        *ListAIJobResponseBodyAIJobList        `json:"AIJobList,omitempty" xml:"AIJobList,omitempty" type:"Struct"`
	NonExistAIJobIds *ListAIJobResponseBodyNonExistAIJobIds `json:"NonExistAIJobIds,omitempty" xml:"NonExistAIJobIds,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 8233A0E4-E112-44*****58-2BCED1B88173
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAIJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIJobResponseBody) GetAIJobList() *ListAIJobResponseBodyAIJobList {
	return s.AIJobList
}

func (s *ListAIJobResponseBody) GetNonExistAIJobIds() *ListAIJobResponseBodyNonExistAIJobIds {
	return s.NonExistAIJobIds
}

func (s *ListAIJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIJobResponseBody) SetAIJobList(v *ListAIJobResponseBodyAIJobList) *ListAIJobResponseBody {
	s.AIJobList = v
	return s
}

func (s *ListAIJobResponseBody) SetNonExistAIJobIds(v *ListAIJobResponseBodyNonExistAIJobIds) *ListAIJobResponseBody {
	s.NonExistAIJobIds = v
	return s
}

func (s *ListAIJobResponseBody) SetRequestId(v string) *ListAIJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIJobResponseBody) Validate() error {
	if s.AIJobList != nil {
		if err := s.AIJobList.Validate(); err != nil {
			return err
		}
	}
	if s.NonExistAIJobIds != nil {
		if err := s.NonExistAIJobIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAIJobResponseBodyAIJobList struct {
	AIJob []*ListAIJobResponseBodyAIJobListAIJob `json:"AIJob,omitempty" xml:"AIJob,omitempty" type:"Repeated"`
}

func (s ListAIJobResponseBodyAIJobList) String() string {
	return dara.Prettify(s)
}

func (s ListAIJobResponseBodyAIJobList) GoString() string {
	return s.String()
}

func (s *ListAIJobResponseBodyAIJobList) GetAIJob() []*ListAIJobResponseBodyAIJobListAIJob {
	return s.AIJob
}

func (s *ListAIJobResponseBodyAIJobList) SetAIJob(v []*ListAIJobResponseBodyAIJobListAIJob) *ListAIJobResponseBodyAIJobList {
	s.AIJob = v
	return s
}

func (s *ListAIJobResponseBodyAIJobList) Validate() error {
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

type ListAIJobResponseBodyAIJobListAIJob struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAIJobResponseBodyAIJobListAIJob) String() string {
	return dara.Prettify(s)
}

func (s ListAIJobResponseBodyAIJobListAIJob) GoString() string {
	return s.String()
}

func (s *ListAIJobResponseBodyAIJobListAIJob) GetCode() *string {
	return s.Code
}

func (s *ListAIJobResponseBodyAIJobListAIJob) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *ListAIJobResponseBodyAIJobListAIJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListAIJobResponseBodyAIJobListAIJob) GetData() *string {
	return s.Data
}

func (s *ListAIJobResponseBodyAIJobListAIJob) GetJobId() *string {
	return s.JobId
}

func (s *ListAIJobResponseBodyAIJobListAIJob) GetMediaId() *string {
	return s.MediaId
}

func (s *ListAIJobResponseBodyAIJobListAIJob) GetMessage() *string {
	return s.Message
}

func (s *ListAIJobResponseBodyAIJobListAIJob) GetStatus() *string {
	return s.Status
}

func (s *ListAIJobResponseBodyAIJobListAIJob) GetType() *string {
	return s.Type
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetCode(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.Code = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetCompleteTime(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.CompleteTime = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetCreationTime(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.CreationTime = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetData(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.Data = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetJobId(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.JobId = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetMediaId(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.MediaId = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetMessage(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.Message = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetStatus(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.Status = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetType(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.Type = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) Validate() error {
	return dara.Validate(s)
}

type ListAIJobResponseBodyNonExistAIJobIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s ListAIJobResponseBodyNonExistAIJobIds) String() string {
	return dara.Prettify(s)
}

func (s ListAIJobResponseBodyNonExistAIJobIds) GoString() string {
	return s.String()
}

func (s *ListAIJobResponseBodyNonExistAIJobIds) GetString_() []*string {
	return s.String_
}

func (s *ListAIJobResponseBodyNonExistAIJobIds) SetString_(v []*string) *ListAIJobResponseBodyNonExistAIJobIds {
	s.String_ = v
	return s
}

func (s *ListAIJobResponseBodyNonExistAIJobIds) Validate() error {
	return dara.Validate(s)
}
