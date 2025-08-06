// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoCoverJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoCoverJobList(v *ListAIVideoCoverJobResponseBodyAIVideoCoverJobList) *ListAIVideoCoverJobResponseBody
	GetAIVideoCoverJobList() *ListAIVideoCoverJobResponseBodyAIVideoCoverJobList
	SetNonExistAIVideoCoverJobIds(v *ListAIVideoCoverJobResponseBodyNonExistAIVideoCoverJobIds) *ListAIVideoCoverJobResponseBody
	GetNonExistAIVideoCoverJobIds() *ListAIVideoCoverJobResponseBodyNonExistAIVideoCoverJobIds
	SetRequestId(v string) *ListAIVideoCoverJobResponseBody
	GetRequestId() *string
}

type ListAIVideoCoverJobResponseBody struct {
	AIVideoCoverJobList        *ListAIVideoCoverJobResponseBodyAIVideoCoverJobList        `json:"AIVideoCoverJobList,omitempty" xml:"AIVideoCoverJobList,omitempty" type:"Struct"`
	NonExistAIVideoCoverJobIds *ListAIVideoCoverJobResponseBodyNonExistAIVideoCoverJobIds `json:"NonExistAIVideoCoverJobIds,omitempty" xml:"NonExistAIVideoCoverJobIds,omitempty" type:"Struct"`
	RequestId                  *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAIVideoCoverJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoCoverJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIVideoCoverJobResponseBody) GetAIVideoCoverJobList() *ListAIVideoCoverJobResponseBodyAIVideoCoverJobList {
	return s.AIVideoCoverJobList
}

func (s *ListAIVideoCoverJobResponseBody) GetNonExistAIVideoCoverJobIds() *ListAIVideoCoverJobResponseBodyNonExistAIVideoCoverJobIds {
	return s.NonExistAIVideoCoverJobIds
}

func (s *ListAIVideoCoverJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIVideoCoverJobResponseBody) SetAIVideoCoverJobList(v *ListAIVideoCoverJobResponseBodyAIVideoCoverJobList) *ListAIVideoCoverJobResponseBody {
	s.AIVideoCoverJobList = v
	return s
}

func (s *ListAIVideoCoverJobResponseBody) SetNonExistAIVideoCoverJobIds(v *ListAIVideoCoverJobResponseBodyNonExistAIVideoCoverJobIds) *ListAIVideoCoverJobResponseBody {
	s.NonExistAIVideoCoverJobIds = v
	return s
}

func (s *ListAIVideoCoverJobResponseBody) SetRequestId(v string) *ListAIVideoCoverJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIVideoCoverJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoCoverJobResponseBodyAIVideoCoverJobList struct {
	AIVideoCoverJob []*ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob `json:"AIVideoCoverJob,omitempty" xml:"AIVideoCoverJob,omitempty" type:"Repeated"`
}

func (s ListAIVideoCoverJobResponseBodyAIVideoCoverJobList) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoCoverJobResponseBodyAIVideoCoverJobList) GoString() string {
	return s.String()
}

func (s *ListAIVideoCoverJobResponseBodyAIVideoCoverJobList) GetAIVideoCoverJob() []*ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob {
	return s.AIVideoCoverJob
}

func (s *ListAIVideoCoverJobResponseBodyAIVideoCoverJobList) SetAIVideoCoverJob(v []*ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob) *ListAIVideoCoverJobResponseBodyAIVideoCoverJobList {
	s.AIVideoCoverJob = v
	return s
}

func (s *ListAIVideoCoverJobResponseBodyAIVideoCoverJobList) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob) GoString() string {
	return s.String()
}

func (s *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob) GetCode() *string {
	return s.Code
}

func (s *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob) GetData() *string {
	return s.Data
}

func (s *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob) GetJobId() *string {
	return s.JobId
}

func (s *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob) GetMediaId() *string {
	return s.MediaId
}

func (s *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob) GetMessage() *string {
	return s.Message
}

func (s *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob) GetStatus() *string {
	return s.Status
}

func (s *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob) SetCode(v string) *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob {
	s.Code = &v
	return s
}

func (s *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob) SetCreationTime(v string) *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob {
	s.CreationTime = &v
	return s
}

func (s *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob) SetData(v string) *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob {
	s.Data = &v
	return s
}

func (s *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob) SetJobId(v string) *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob {
	s.JobId = &v
	return s
}

func (s *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob) SetMediaId(v string) *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob {
	s.MediaId = &v
	return s
}

func (s *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob) SetMessage(v string) *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob {
	s.Message = &v
	return s
}

func (s *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob) SetStatus(v string) *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob {
	s.Status = &v
	return s
}

func (s *ListAIVideoCoverJobResponseBodyAIVideoCoverJobListAIVideoCoverJob) Validate() error {
	return dara.Validate(s)
}

type ListAIVideoCoverJobResponseBodyNonExistAIVideoCoverJobIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s ListAIVideoCoverJobResponseBodyNonExistAIVideoCoverJobIds) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoCoverJobResponseBodyNonExistAIVideoCoverJobIds) GoString() string {
	return s.String()
}

func (s *ListAIVideoCoverJobResponseBodyNonExistAIVideoCoverJobIds) GetString_() []*string {
	return s.String_
}

func (s *ListAIVideoCoverJobResponseBodyNonExistAIVideoCoverJobIds) SetString_(v []*string) *ListAIVideoCoverJobResponseBodyNonExistAIVideoCoverJobIds {
	s.String_ = v
	return s
}

func (s *ListAIVideoCoverJobResponseBodyNonExistAIVideoCoverJobIds) Validate() error {
	return dara.Validate(s)
}
