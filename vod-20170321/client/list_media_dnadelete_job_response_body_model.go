// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaDNADeleteJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIJobList(v *ListMediaDNADeleteJobResponseBodyAIJobList) *ListMediaDNADeleteJobResponseBody
	GetAIJobList() *ListMediaDNADeleteJobResponseBodyAIJobList
	SetNonExistAIJobIds(v *ListMediaDNADeleteJobResponseBodyNonExistAIJobIds) *ListMediaDNADeleteJobResponseBody
	GetNonExistAIJobIds() *ListMediaDNADeleteJobResponseBodyNonExistAIJobIds
	SetRequestId(v string) *ListMediaDNADeleteJobResponseBody
	GetRequestId() *string
}

type ListMediaDNADeleteJobResponseBody struct {
	AIJobList        *ListMediaDNADeleteJobResponseBodyAIJobList        `json:"AIJobList,omitempty" xml:"AIJobList,omitempty" type:"Struct"`
	NonExistAIJobIds *ListMediaDNADeleteJobResponseBodyNonExistAIJobIds `json:"NonExistAIJobIds,omitempty" xml:"NonExistAIJobIds,omitempty" type:"Struct"`
	RequestId        *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMediaDNADeleteJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMediaDNADeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListMediaDNADeleteJobResponseBody) GetAIJobList() *ListMediaDNADeleteJobResponseBodyAIJobList {
	return s.AIJobList
}

func (s *ListMediaDNADeleteJobResponseBody) GetNonExistAIJobIds() *ListMediaDNADeleteJobResponseBodyNonExistAIJobIds {
	return s.NonExistAIJobIds
}

func (s *ListMediaDNADeleteJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMediaDNADeleteJobResponseBody) SetAIJobList(v *ListMediaDNADeleteJobResponseBodyAIJobList) *ListMediaDNADeleteJobResponseBody {
	s.AIJobList = v
	return s
}

func (s *ListMediaDNADeleteJobResponseBody) SetNonExistAIJobIds(v *ListMediaDNADeleteJobResponseBodyNonExistAIJobIds) *ListMediaDNADeleteJobResponseBody {
	s.NonExistAIJobIds = v
	return s
}

func (s *ListMediaDNADeleteJobResponseBody) SetRequestId(v string) *ListMediaDNADeleteJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMediaDNADeleteJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMediaDNADeleteJobResponseBodyAIJobList struct {
	AIJob []*ListMediaDNADeleteJobResponseBodyAIJobListAIJob `json:"AIJob,omitempty" xml:"AIJob,omitempty" type:"Repeated"`
}

func (s ListMediaDNADeleteJobResponseBodyAIJobList) String() string {
	return dara.Prettify(s)
}

func (s ListMediaDNADeleteJobResponseBodyAIJobList) GoString() string {
	return s.String()
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobList) GetAIJob() []*ListMediaDNADeleteJobResponseBodyAIJobListAIJob {
	return s.AIJob
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobList) SetAIJob(v []*ListMediaDNADeleteJobResponseBodyAIJobListAIJob) *ListMediaDNADeleteJobResponseBodyAIJobList {
	s.AIJob = v
	return s
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobList) Validate() error {
	return dara.Validate(s)
}

type ListMediaDNADeleteJobResponseBodyAIJobListAIJob struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	FpDBId  *string `json:"FpDBId,omitempty" xml:"FpDBId,omitempty"`
	JobId   *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMediaDNADeleteJobResponseBodyAIJobListAIJob) String() string {
	return dara.Prettify(s)
}

func (s ListMediaDNADeleteJobResponseBodyAIJobListAIJob) GoString() string {
	return s.String()
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobListAIJob) GetCode() *string {
	return s.Code
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobListAIJob) GetFpDBId() *string {
	return s.FpDBId
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobListAIJob) GetJobId() *string {
	return s.JobId
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobListAIJob) GetMediaId() *string {
	return s.MediaId
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobListAIJob) GetMessage() *string {
	return s.Message
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobListAIJob) GetStatus() *string {
	return s.Status
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobListAIJob) SetCode(v string) *ListMediaDNADeleteJobResponseBodyAIJobListAIJob {
	s.Code = &v
	return s
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobListAIJob) SetFpDBId(v string) *ListMediaDNADeleteJobResponseBodyAIJobListAIJob {
	s.FpDBId = &v
	return s
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobListAIJob) SetJobId(v string) *ListMediaDNADeleteJobResponseBodyAIJobListAIJob {
	s.JobId = &v
	return s
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobListAIJob) SetMediaId(v string) *ListMediaDNADeleteJobResponseBodyAIJobListAIJob {
	s.MediaId = &v
	return s
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobListAIJob) SetMessage(v string) *ListMediaDNADeleteJobResponseBodyAIJobListAIJob {
	s.Message = &v
	return s
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobListAIJob) SetStatus(v string) *ListMediaDNADeleteJobResponseBodyAIJobListAIJob {
	s.Status = &v
	return s
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobListAIJob) Validate() error {
	return dara.Validate(s)
}

type ListMediaDNADeleteJobResponseBodyNonExistAIJobIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s ListMediaDNADeleteJobResponseBodyNonExistAIJobIds) String() string {
	return dara.Prettify(s)
}

func (s ListMediaDNADeleteJobResponseBodyNonExistAIJobIds) GoString() string {
	return s.String()
}

func (s *ListMediaDNADeleteJobResponseBodyNonExistAIJobIds) GetString_() []*string {
	return s.String_
}

func (s *ListMediaDNADeleteJobResponseBodyNonExistAIJobIds) SetString_(v []*string) *ListMediaDNADeleteJobResponseBodyNonExistAIJobIds {
	s.String_ = v
	return s
}

func (s *ListMediaDNADeleteJobResponseBodyNonExistAIJobIds) Validate() error {
	return dara.Validate(s)
}
