// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoTerrorismRecogJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoTerrorismRecogJobList(v *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobList) *ListAIVideoTerrorismRecogJobResponseBody
	GetAIVideoTerrorismRecogJobList() *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobList
	SetNonExistTerrorismRecogJobIds(v *ListAIVideoTerrorismRecogJobResponseBodyNonExistTerrorismRecogJobIds) *ListAIVideoTerrorismRecogJobResponseBody
	GetNonExistTerrorismRecogJobIds() *ListAIVideoTerrorismRecogJobResponseBodyNonExistTerrorismRecogJobIds
	SetRequestId(v string) *ListAIVideoTerrorismRecogJobResponseBody
	GetRequestId() *string
}

type ListAIVideoTerrorismRecogJobResponseBody struct {
	AIVideoTerrorismRecogJobList *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobList `json:"AIVideoTerrorismRecogJobList,omitempty" xml:"AIVideoTerrorismRecogJobList,omitempty" type:"Struct"`
	NonExistTerrorismRecogJobIds *ListAIVideoTerrorismRecogJobResponseBodyNonExistTerrorismRecogJobIds `json:"NonExistTerrorismRecogJobIds,omitempty" xml:"NonExistTerrorismRecogJobIds,omitempty" type:"Struct"`
	RequestId                    *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAIVideoTerrorismRecogJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoTerrorismRecogJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIVideoTerrorismRecogJobResponseBody) GetAIVideoTerrorismRecogJobList() *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobList {
	return s.AIVideoTerrorismRecogJobList
}

func (s *ListAIVideoTerrorismRecogJobResponseBody) GetNonExistTerrorismRecogJobIds() *ListAIVideoTerrorismRecogJobResponseBodyNonExistTerrorismRecogJobIds {
	return s.NonExistTerrorismRecogJobIds
}

func (s *ListAIVideoTerrorismRecogJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIVideoTerrorismRecogJobResponseBody) SetAIVideoTerrorismRecogJobList(v *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobList) *ListAIVideoTerrorismRecogJobResponseBody {
	s.AIVideoTerrorismRecogJobList = v
	return s
}

func (s *ListAIVideoTerrorismRecogJobResponseBody) SetNonExistTerrorismRecogJobIds(v *ListAIVideoTerrorismRecogJobResponseBodyNonExistTerrorismRecogJobIds) *ListAIVideoTerrorismRecogJobResponseBody {
	s.NonExistTerrorismRecogJobIds = v
	return s
}

func (s *ListAIVideoTerrorismRecogJobResponseBody) SetRequestId(v string) *ListAIVideoTerrorismRecogJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIVideoTerrorismRecogJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobList struct {
	AIVideoTerrorismRecogJob []*ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob `json:"AIVideoTerrorismRecogJob,omitempty" xml:"AIVideoTerrorismRecogJob,omitempty" type:"Repeated"`
}

func (s ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobList) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobList) GoString() string {
	return s.String()
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobList) GetAIVideoTerrorismRecogJob() []*ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob {
	return s.AIVideoTerrorismRecogJob
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobList) SetAIVideoTerrorismRecogJob(v []*ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob) *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobList {
	s.AIVideoTerrorismRecogJob = v
	return s
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobList) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob) GoString() string {
	return s.String()
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob) GetCode() *string {
	return s.Code
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob) GetData() *string {
	return s.Data
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob) GetJobId() *string {
	return s.JobId
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob) GetMediaId() *string {
	return s.MediaId
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob) GetMessage() *string {
	return s.Message
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob) GetStatus() *string {
	return s.Status
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob) SetCode(v string) *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob {
	s.Code = &v
	return s
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob) SetCreationTime(v string) *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob {
	s.CreationTime = &v
	return s
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob) SetData(v string) *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob {
	s.Data = &v
	return s
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob) SetJobId(v string) *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob {
	s.JobId = &v
	return s
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob) SetMediaId(v string) *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob {
	s.MediaId = &v
	return s
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob) SetMessage(v string) *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob {
	s.Message = &v
	return s
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob) SetStatus(v string) *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob {
	s.Status = &v
	return s
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyAIVideoTerrorismRecogJobListAIVideoTerrorismRecogJob) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoTerrorismRecogJobResponseBodyNonExistTerrorismRecogJobIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s ListAIVideoTerrorismRecogJobResponseBodyNonExistTerrorismRecogJobIds) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoTerrorismRecogJobResponseBodyNonExistTerrorismRecogJobIds) GoString() string {
	return s.String()
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyNonExistTerrorismRecogJobIds) GetString_() []*string {
	return s.String_
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyNonExistTerrorismRecogJobIds) SetString_(v []*string) *ListAIVideoTerrorismRecogJobResponseBodyNonExistTerrorismRecogJobIds {
	s.String_ = v
	return s
}

func (s *ListAIVideoTerrorismRecogJobResponseBodyNonExistTerrorismRecogJobIds) Validate() error {
	return dara.Validate(s)
}
