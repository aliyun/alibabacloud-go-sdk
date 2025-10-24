// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFpDBDeleteJobListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFpDBDeleteJobList(v *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobList) *QueryFpDBDeleteJobListResponseBody
	GetFpDBDeleteJobList() *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobList
	SetNonExistIds(v *QueryFpDBDeleteJobListResponseBodyNonExistIds) *QueryFpDBDeleteJobListResponseBody
	GetNonExistIds() *QueryFpDBDeleteJobListResponseBodyNonExistIds
	SetRequestId(v string) *QueryFpDBDeleteJobListResponseBody
	GetRequestId() *string
}

type QueryFpDBDeleteJobListResponseBody struct {
	// The jobs of deleting a media fingerprint library. For more information, see the "FpDBDeleteJob" section of the [Data types](https://www.alibabacloud.com/help/en/apsaravideo-for-media-processing/latest/datatypes) topic.
	FpDBDeleteJobList *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobList `json:"FpDBDeleteJobList,omitempty" xml:"FpDBDeleteJobList,omitempty" type:"Struct"`
	// The IDs of the jobs that do not exist.
	NonExistIds *QueryFpDBDeleteJobListResponseBodyNonExistIds `json:"NonExistIds,omitempty" xml:"NonExistIds,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 4247B23C-26DE-529F-8D9F-FD6811AE979B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryFpDBDeleteJobListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFpDBDeleteJobListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFpDBDeleteJobListResponseBody) GetFpDBDeleteJobList() *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobList {
	return s.FpDBDeleteJobList
}

func (s *QueryFpDBDeleteJobListResponseBody) GetNonExistIds() *QueryFpDBDeleteJobListResponseBodyNonExistIds {
	return s.NonExistIds
}

func (s *QueryFpDBDeleteJobListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFpDBDeleteJobListResponseBody) SetFpDBDeleteJobList(v *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobList) *QueryFpDBDeleteJobListResponseBody {
	s.FpDBDeleteJobList = v
	return s
}

func (s *QueryFpDBDeleteJobListResponseBody) SetNonExistIds(v *QueryFpDBDeleteJobListResponseBodyNonExistIds) *QueryFpDBDeleteJobListResponseBody {
	s.NonExistIds = v
	return s
}

func (s *QueryFpDBDeleteJobListResponseBody) SetRequestId(v string) *QueryFpDBDeleteJobListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFpDBDeleteJobListResponseBody) Validate() error {
	if s.FpDBDeleteJobList != nil {
		if err := s.FpDBDeleteJobList.Validate(); err != nil {
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

type QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobList struct {
	FpDBDeleteJob []*QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob `json:"FpDBDeleteJob,omitempty" xml:"FpDBDeleteJob,omitempty" type:"Repeated"`
}

func (s QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobList) String() string {
	return dara.Prettify(s)
}

func (s QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobList) GoString() string {
	return s.String()
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobList) GetFpDBDeleteJob() []*QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob {
	return s.FpDBDeleteJob
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobList) SetFpDBDeleteJob(v []*QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobList {
	s.FpDBDeleteJob = v
	return s
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobList) Validate() error {
	if s.FpDBDeleteJob != nil {
		for _, item := range s.FpDBDeleteJob {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob struct {
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
	// The type of the operation.
	//
	// example:
	//
	// Purge
	DelType *string `json:"DelType,omitempty" xml:"DelType,omitempty"`
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
	// The status of the job. Valid values:
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

func (s QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) String() string {
	return dara.Prettify(s)
}

func (s QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) GoString() string {
	return s.String()
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) GetCode() *string {
	return s.Code
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) GetDelType() *string {
	return s.DelType
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) GetFpDBId() *string {
	return s.FpDBId
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) GetId() *string {
	return s.Id
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) GetMessage() *string {
	return s.Message
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) GetPipelineId() *string {
	return s.PipelineId
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) GetStatus() *string {
	return s.Status
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) GetUserData() *string {
	return s.UserData
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) SetCode(v string) *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob {
	s.Code = &v
	return s
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) SetCreationTime(v string) *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob {
	s.CreationTime = &v
	return s
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) SetDelType(v string) *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob {
	s.DelType = &v
	return s
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) SetFinishTime(v string) *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob {
	s.FinishTime = &v
	return s
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) SetFpDBId(v string) *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob {
	s.FpDBId = &v
	return s
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) SetId(v string) *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob {
	s.Id = &v
	return s
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) SetMessage(v string) *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob {
	s.Message = &v
	return s
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) SetPipelineId(v string) *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob {
	s.PipelineId = &v
	return s
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) SetStatus(v string) *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob {
	s.Status = &v
	return s
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) SetUserData(v string) *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob {
	s.UserData = &v
	return s
}

func (s *QueryFpDBDeleteJobListResponseBodyFpDBDeleteJobListFpDBDeleteJob) Validate() error {
	return dara.Validate(s)
}

type QueryFpDBDeleteJobListResponseBodyNonExistIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s QueryFpDBDeleteJobListResponseBodyNonExistIds) String() string {
	return dara.Prettify(s)
}

func (s QueryFpDBDeleteJobListResponseBodyNonExistIds) GoString() string {
	return s.String()
}

func (s *QueryFpDBDeleteJobListResponseBodyNonExistIds) GetString_() []*string {
	return s.String_
}

func (s *QueryFpDBDeleteJobListResponseBodyNonExistIds) SetString_(v []*string) *QueryFpDBDeleteJobListResponseBodyNonExistIds {
	s.String_ = v
	return s
}

func (s *QueryFpDBDeleteJobListResponseBodyNonExistIds) Validate() error {
	return dara.Validate(s)
}
