// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFpFileDeleteJobListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFpFileDeleteJobList(v *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobList) *QueryFpFileDeleteJobListResponseBody
	GetFpFileDeleteJobList() *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobList
	SetNonExistIds(v *QueryFpFileDeleteJobListResponseBodyNonExistIds) *QueryFpFileDeleteJobListResponseBody
	GetNonExistIds() *QueryFpFileDeleteJobListResponseBodyNonExistIds
	SetRequestId(v string) *QueryFpFileDeleteJobListResponseBody
	GetRequestId() *string
}

type QueryFpFileDeleteJobListResponseBody struct {
	// The jobs of deleting media files from a media fingerprint library. For more information, see the "FpFileDeleteJob" section of the [Data types](https://help.aliyun.com/document_detail/93555.html) topic.
	FpFileDeleteJobList *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobList `json:"FpFileDeleteJobList,omitempty" xml:"FpFileDeleteJobList,omitempty" type:"Struct"`
	// The response parameters.
	NonExistIds *QueryFpFileDeleteJobListResponseBodyNonExistIds `json:"NonExistIds,omitempty" xml:"NonExistIds,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// D127C68E-F1A1-4CE5-A874-8FF724881A12
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryFpFileDeleteJobListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFpFileDeleteJobListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFpFileDeleteJobListResponseBody) GetFpFileDeleteJobList() *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobList {
	return s.FpFileDeleteJobList
}

func (s *QueryFpFileDeleteJobListResponseBody) GetNonExistIds() *QueryFpFileDeleteJobListResponseBodyNonExistIds {
	return s.NonExistIds
}

func (s *QueryFpFileDeleteJobListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFpFileDeleteJobListResponseBody) SetFpFileDeleteJobList(v *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobList) *QueryFpFileDeleteJobListResponseBody {
	s.FpFileDeleteJobList = v
	return s
}

func (s *QueryFpFileDeleteJobListResponseBody) SetNonExistIds(v *QueryFpFileDeleteJobListResponseBodyNonExistIds) *QueryFpFileDeleteJobListResponseBody {
	s.NonExistIds = v
	return s
}

func (s *QueryFpFileDeleteJobListResponseBody) SetRequestId(v string) *QueryFpFileDeleteJobListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFpFileDeleteJobListResponseBody) Validate() error {
	if s.FpFileDeleteJobList != nil {
		if err := s.FpFileDeleteJobList.Validate(); err != nil {
			return err
		}
	}
	if s.NonExistIds != nil {
		if err := s.NonExistIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobList struct {
	FpFileDeleteJob []*QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob `json:"FpFileDeleteJob,omitempty" xml:"FpFileDeleteJob,omitempty" type:"Repeated"`
}

func (s QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobList) String() string {
	return dara.Prettify(s)
}

func (s QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobList) GoString() string {
	return s.String()
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobList) GetFpFileDeleteJob() []*QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob {
	return s.FpFileDeleteJob
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobList) SetFpFileDeleteJob(v []*QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobList {
	s.FpFileDeleteJob = v
	return s
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobList) Validate() error {
	if s.FpFileDeleteJob != nil {
		for _, item := range s.FpFileDeleteJob {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob struct {
	// The error code returned if the job fails. This parameter is not returned if the job is successful.
	//
	// example:
	//
	// ServiceUnavailable
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the job was created.
	//
	// example:
	//
	// 2020-06-30T00:33:18Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the file.
	//
	// example:
	//
	// 41e6536e4f2250e2e9bf26cdea19****
	FileIds *string `json:"FileIds,omitempty" xml:"FileIds,omitempty"`
	// The time when the job was complete.
	//
	// example:
	//
	// 2020-06-30T00:34:02Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the media fingerprint library.
	//
	// example:
	//
	// 88c6ca184c0e47098a5b665e2a12****
	FpDBId *string `json:"FpDBId,omitempty" xml:"FpDBId,omitempty"`
	// The ID of the job.
	//
	// example:
	//
	// 25bacf2824614bcf9273dc0744db****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The error message returned if the job fails. This parameter is not returned if the job is successful.
	//
	// example:
	//
	// The request has failed due to a temporary failure of the server.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the ApsaraVideo Media Processing (MPS) queue to which the job was submitted.
	//
	// example:
	//
	// fb712a6890464059b1b2ea7c8647****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The status of the job. Valid values: Valid values:
	//
	// 	- **Queuing**: The job is waiting in the queue.
	//
	// 	- **Analysing**: The job is in progress.
	//
	// 	- **Success**: The job is successful.
	//
	// 	- **Fail**: The job fails.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The user-defined data.
	//
	// example:
	//
	// example data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) String() string {
	return dara.Prettify(s)
}

func (s QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) GoString() string {
	return s.String()
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) GetCode() *string {
	return s.Code
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) GetFileIds() *string {
	return s.FileIds
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) GetFpDBId() *string {
	return s.FpDBId
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) GetId() *string {
	return s.Id
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) GetMessage() *string {
	return s.Message
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) GetPipelineId() *string {
	return s.PipelineId
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) GetStatus() *string {
	return s.Status
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) GetUserData() *string {
	return s.UserData
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) SetCode(v string) *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob {
	s.Code = &v
	return s
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) SetCreationTime(v string) *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob {
	s.CreationTime = &v
	return s
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) SetFileIds(v string) *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob {
	s.FileIds = &v
	return s
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) SetFinishTime(v string) *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob {
	s.FinishTime = &v
	return s
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) SetFpDBId(v string) *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob {
	s.FpDBId = &v
	return s
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) SetId(v string) *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob {
	s.Id = &v
	return s
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) SetMessage(v string) *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob {
	s.Message = &v
	return s
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) SetPipelineId(v string) *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob {
	s.PipelineId = &v
	return s
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) SetStatus(v string) *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob {
	s.Status = &v
	return s
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) SetUserData(v string) *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob {
	s.UserData = &v
	return s
}

func (s *QueryFpFileDeleteJobListResponseBodyFpFileDeleteJobListFpFileDeleteJob) Validate() error {
	return dara.Validate(s)
}

type QueryFpFileDeleteJobListResponseBodyNonExistIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s QueryFpFileDeleteJobListResponseBodyNonExistIds) String() string {
	return dara.Prettify(s)
}

func (s QueryFpFileDeleteJobListResponseBodyNonExistIds) GoString() string {
	return s.String()
}

func (s *QueryFpFileDeleteJobListResponseBodyNonExistIds) GetString_() []*string {
	return s.String_
}

func (s *QueryFpFileDeleteJobListResponseBodyNonExistIds) SetString_(v []*string) *QueryFpFileDeleteJobListResponseBodyNonExistIds {
	s.String_ = v
	return s
}

func (s *QueryFpFileDeleteJobListResponseBodyNonExistIds) Validate() error {
	return dara.Validate(s)
}
