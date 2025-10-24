// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFpShotImportJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFpShotImportJobList(v []*ListFpShotImportJobResponseBodyFpShotImportJobList) *ListFpShotImportJobResponseBody
	GetFpShotImportJobList() []*ListFpShotImportJobResponseBodyFpShotImportJobList
	SetNonExistIds(v []*string) *ListFpShotImportJobResponseBody
	GetNonExistIds() []*string
	SetRequestId(v string) *ListFpShotImportJobResponseBody
	GetRequestId() *string
}

type ListFpShotImportJobResponseBody struct {
	// The jobs of importing text files to a text fingerprint library.
	FpShotImportJobList []*ListFpShotImportJobResponseBodyFpShotImportJobList `json:"FpShotImportJobList,omitempty" xml:"FpShotImportJobList,omitempty" type:"Repeated"`
	// The job IDs that do not exist. This parameter is not returned if all specified job IDs exist.
	NonExistIds []*string `json:"NonExistIds,omitempty" xml:"NonExistIds,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFpShotImportJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFpShotImportJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListFpShotImportJobResponseBody) GetFpShotImportJobList() []*ListFpShotImportJobResponseBodyFpShotImportJobList {
	return s.FpShotImportJobList
}

func (s *ListFpShotImportJobResponseBody) GetNonExistIds() []*string {
	return s.NonExistIds
}

func (s *ListFpShotImportJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFpShotImportJobResponseBody) SetFpShotImportJobList(v []*ListFpShotImportJobResponseBodyFpShotImportJobList) *ListFpShotImportJobResponseBody {
	s.FpShotImportJobList = v
	return s
}

func (s *ListFpShotImportJobResponseBody) SetNonExistIds(v []*string) *ListFpShotImportJobResponseBody {
	s.NonExistIds = v
	return s
}

func (s *ListFpShotImportJobResponseBody) SetRequestId(v string) *ListFpShotImportJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFpShotImportJobResponseBody) Validate() error {
	if s.FpShotImportJobList != nil {
		for _, item := range s.FpShotImportJobList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFpShotImportJobResponseBodyFpShotImportJobList struct {
	// The error code returned when the job fails.
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the job was created.
	//
	// example:
	//
	// 2020-06-30T00:33:18Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the job was completed.
	//
	// example:
	//
	// 2020-06-30T00:34:02Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the text fingerprint library.
	//
	// example:
	//
	// 2288c6ca184c0e47098a5b665e2a12****
	FpDBId *string `json:"FpDBId,omitempty" xml:"FpDBId,omitempty"`
	// The import configuration.
	//
	// example:
	//
	// ""
	FpImportConfig *string `json:"FpImportConfig,omitempty" xml:"FpImportConfig,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 25bacf2824614bcf9273dc0744db****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The input file.
	//
	// example:
	//
	// {\\"Bucket\\":\\"mts-example****\\",\\"Location\\":\\"oss-cn-shanghai\\",\\"Object\\":\\"test-0828/video/test.mp4\\"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The error message returned when the job fails.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the ApsaraVideo Media Processing (MPS) queue to which the job is submitted.
	//
	// example:
	//
	// ebb51ee30f0b49aba959823fa991****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The processing information of the job.
	//
	// example:
	//
	// http://testbucket.oss-cn-shanghai.aliyuncs.com/932ajjw***32ssoj_importResult.txt
	ProcessMessage *string `json:"ProcessMessage,omitempty" xml:"ProcessMessage,omitempty"`
	// The status of the job. Valid values:
	//
	// 	- Processing: The job is in progress.
	//
	// 	- Fail: The job fails.
	//
	// 	- Success: The job is successful.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The user-defined data.
	//
	// example:
	//
	// 001
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ListFpShotImportJobResponseBodyFpShotImportJobList) String() string {
	return dara.Prettify(s)
}

func (s ListFpShotImportJobResponseBodyFpShotImportJobList) GoString() string {
	return s.String()
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) GetCode() *string {
	return s.Code
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) GetFpDBId() *string {
	return s.FpDBId
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) GetFpImportConfig() *string {
	return s.FpImportConfig
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) GetId() *string {
	return s.Id
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) GetInput() *string {
	return s.Input
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) GetMessage() *string {
	return s.Message
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) GetPipelineId() *string {
	return s.PipelineId
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) GetProcessMessage() *string {
	return s.ProcessMessage
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) GetStatus() *string {
	return s.Status
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) GetUserData() *string {
	return s.UserData
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) SetCode(v string) *ListFpShotImportJobResponseBodyFpShotImportJobList {
	s.Code = &v
	return s
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) SetCreateTime(v string) *ListFpShotImportJobResponseBodyFpShotImportJobList {
	s.CreateTime = &v
	return s
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) SetFinishTime(v string) *ListFpShotImportJobResponseBodyFpShotImportJobList {
	s.FinishTime = &v
	return s
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) SetFpDBId(v string) *ListFpShotImportJobResponseBodyFpShotImportJobList {
	s.FpDBId = &v
	return s
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) SetFpImportConfig(v string) *ListFpShotImportJobResponseBodyFpShotImportJobList {
	s.FpImportConfig = &v
	return s
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) SetId(v string) *ListFpShotImportJobResponseBodyFpShotImportJobList {
	s.Id = &v
	return s
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) SetInput(v string) *ListFpShotImportJobResponseBodyFpShotImportJobList {
	s.Input = &v
	return s
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) SetMessage(v string) *ListFpShotImportJobResponseBodyFpShotImportJobList {
	s.Message = &v
	return s
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) SetPipelineId(v string) *ListFpShotImportJobResponseBodyFpShotImportJobList {
	s.PipelineId = &v
	return s
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) SetProcessMessage(v string) *ListFpShotImportJobResponseBodyFpShotImportJobList {
	s.ProcessMessage = &v
	return s
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) SetStatus(v string) *ListFpShotImportJobResponseBodyFpShotImportJobList {
	s.Status = &v
	return s
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) SetUserData(v string) *ListFpShotImportJobResponseBodyFpShotImportJobList {
	s.UserData = &v
	return s
}

func (s *ListFpShotImportJobResponseBodyFpShotImportJobList) Validate() error {
	return dara.Validate(s)
}
