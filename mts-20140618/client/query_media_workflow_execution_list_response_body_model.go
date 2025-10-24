// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaWorkflowExecutionListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaWorkflowExecutionList(v *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionList) *QueryMediaWorkflowExecutionListResponseBody
	GetMediaWorkflowExecutionList() *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionList
	SetNonExistRunIds(v *QueryMediaWorkflowExecutionListResponseBodyNonExistRunIds) *QueryMediaWorkflowExecutionListResponseBody
	GetNonExistRunIds() *QueryMediaWorkflowExecutionListResponseBodyNonExistRunIds
	SetRequestId(v string) *QueryMediaWorkflowExecutionListResponseBody
	GetRequestId() *string
}

type QueryMediaWorkflowExecutionListResponseBody struct {
	// The details of the media workflows.
	MediaWorkflowExecutionList *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionList `json:"MediaWorkflowExecutionList,omitempty" xml:"MediaWorkflowExecutionList,omitempty" type:"Struct"`
	// The IDs of the execution instances that do not exist.
	NonExistRunIds *QueryMediaWorkflowExecutionListResponseBodyNonExistRunIds `json:"NonExistRunIds,omitempty" xml:"NonExistRunIds,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// D1D5C080-8E2F-5030-8AB4-13092F17631B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryMediaWorkflowExecutionListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaWorkflowExecutionListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMediaWorkflowExecutionListResponseBody) GetMediaWorkflowExecutionList() *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionList {
	return s.MediaWorkflowExecutionList
}

func (s *QueryMediaWorkflowExecutionListResponseBody) GetNonExistRunIds() *QueryMediaWorkflowExecutionListResponseBodyNonExistRunIds {
	return s.NonExistRunIds
}

func (s *QueryMediaWorkflowExecutionListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMediaWorkflowExecutionListResponseBody) SetMediaWorkflowExecutionList(v *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionList) *QueryMediaWorkflowExecutionListResponseBody {
	s.MediaWorkflowExecutionList = v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBody) SetNonExistRunIds(v *QueryMediaWorkflowExecutionListResponseBodyNonExistRunIds) *QueryMediaWorkflowExecutionListResponseBody {
	s.NonExistRunIds = v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBody) SetRequestId(v string) *QueryMediaWorkflowExecutionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBody) Validate() error {
	if s.MediaWorkflowExecutionList != nil {
		if err := s.MediaWorkflowExecutionList.Validate(); err != nil {
			return err
		}
	}
	if s.NonExistRunIds != nil {
		if err := s.NonExistRunIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionList struct {
	MediaWorkflowExecution []*QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution `json:"MediaWorkflowExecution,omitempty" xml:"MediaWorkflowExecution,omitempty" type:"Repeated"`
}

func (s QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionList) GoString() string {
	return s.String()
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionList) GetMediaWorkflowExecution() []*QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution {
	return s.MediaWorkflowExecution
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionList) SetMediaWorkflowExecution(v []*QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionList {
	s.MediaWorkflowExecution = v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionList) Validate() error {
	if s.MediaWorkflowExecution != nil {
		for _, item := range s.MediaWorkflowExecution {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution struct {
	// The methods that are called in the media workflow.
	ActivityList *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList `json:"ActivityList,omitempty" xml:"ActivityList,omitempty" type:"Struct"`
	// The time when the media workflow was created.
	//
	// example:
	//
	// 016-04-01T06:53:43Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The input data of the media workflow.
	Input *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The ID of the media asset. A media file contains all the information about a media workflow.
	//
	// example:
	//
	// 512046582a924698a41e0f8b0d2b****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The ID of the media workflow.
	//
	// example:
	//
	// 93ab850b4f6f44eab54b6e91****81d4
	MediaWorkflowId *string `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty"`
	// The name of the media workflow.
	//
	// example:
	//
	// example-mediaworkflow-****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the execution instance.
	//
	// example:
	//
	// 48e33690ac19445488c706924321****
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// The status of the media workflow. Valid values:
	//
	// 	- Running: The media workflow is running.
	//
	// 	- Completed: The media workflow is complete.
	//
	// > Completed only indicates that the media workflow is complete. View the status of each method in the workflow, such as the transcode and snapshot methods, to check whether the method is called.
	//
	// 	- Fail: The media workflow fails.
	//
	// example:
	//
	// Completed
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) GoString() string {
	return s.String()
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) GetActivityList() *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList {
	return s.ActivityList
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) GetCreationTime() *string {
	return s.CreationTime
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) GetInput() *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput {
	return s.Input
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) GetMediaId() *string {
	return s.MediaId
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) GetName() *string {
	return s.Name
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) GetRunId() *string {
	return s.RunId
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) GetState() *string {
	return s.State
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) SetActivityList(v *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution {
	s.ActivityList = v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) SetCreationTime(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution {
	s.CreationTime = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) SetInput(v *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution {
	s.Input = v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) SetMediaId(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution {
	s.MediaId = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) SetMediaWorkflowId(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution {
	s.MediaWorkflowId = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) SetName(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution {
	s.Name = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) SetRunId(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution {
	s.RunId = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) SetState(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution {
	s.State = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) Validate() error {
	if s.ActivityList != nil {
		if err := s.ActivityList.Validate(); err != nil {
			return err
		}
	}
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList struct {
	Activity []*QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity `json:"Activity,omitempty" xml:"Activity,omitempty" type:"Repeated"`
}

func (s QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList) GoString() string {
	return s.String()
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList) GetActivity() []*QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity {
	return s.Activity
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList) SetActivity(v []*QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList {
	s.Activity = v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList) Validate() error {
	if s.Activity != nil {
		for _, item := range s.Activity {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity struct {
	// The error code.
	//
	// 	- This parameter is returned only if **Fail*	- is returned for the State parameter.
	//
	// 	- This parameter is not returned if the method status is **Success**.
	//
	// example:
	//
	// InvalidParameter.ResourceContentBad
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the method ends.
	//
	// example:
	//
	// 2016-04-01T06:53:44Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IDs of the jobs that are generated when the methods are called, such as the job IDs for the analysis, transcode, and snapshot methods.
	//
	// example:
	//
	// 2376030d9d0849399cd20e20f4f3****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The message sent by Message Service (MNS) to notify the user of the job result.
	MNSMessageResult *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult `json:"MNSMessageResult,omitempty" xml:"MNSMessageResult,omitempty" type:"Struct"`
	// The error message.
	//
	// 	- This parameter is returned only if **Fail*	- is returned for the State parameter.
	//
	// 	- This parameter is not returned if the method status is **Success**.
	//
	// example:
	//
	// The resource operated InputFile is bad
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the method.
	//
	// > The name of each method in a media workflow is unique.
	//
	// example:
	//
	// Start
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the method is called.
	//
	// example:
	//
	// 2016-04-01T06:53:44Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the workflow method. Valid values:
	//
	// 	- Running: The method is being called.
	//
	// 	- Success: The method is called.
	//
	// 	- Fail: The method failed to be called.
	//
	// 	- Skipped: The method is skipped.
	//
	// > For example, after the analysis is complete, the transcode method is called and high-definition and standard-definition transcoding jobs are created. The system determines whether to run the jobs based on the analysis result. If the resolution of the input video is low, the high-definition transcoding job may be skipped.
	//
	// example:
	//
	// Running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The methods that are supported in the media workflow. Valid values: Start, Snapshot, Transcode, Analysis, and Report. For more information, see [Methods supported for media workflows](https://help.aliyun.com/document_detail/68494.html).
	//
	// example:
	//
	// Start
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) GoString() string {
	return s.String()
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) GetCode() *string {
	return s.Code
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) GetJobId() *string {
	return s.JobId
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) GetMNSMessageResult() *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult {
	return s.MNSMessageResult
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) GetMessage() *string {
	return s.Message
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) GetName() *string {
	return s.Name
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) GetState() *string {
	return s.State
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) GetType() *string {
	return s.Type
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) SetCode(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity {
	s.Code = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) SetEndTime(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity {
	s.EndTime = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) SetJobId(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity {
	s.JobId = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) SetMNSMessageResult(v *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity {
	s.MNSMessageResult = v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) SetMessage(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity {
	s.Message = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) SetName(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity {
	s.Name = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) SetStartTime(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity {
	s.StartTime = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) SetState(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity {
	s.State = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) SetType(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity {
	s.Type = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) Validate() error {
	if s.MNSMessageResult != nil {
		if err := s.MNSMessageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult struct {
	// The error code returned if the MNS message fails to be sent. This parameter is not returned if the MNS message is sent.
	//
	// example:
	//
	// The Topic/Queue config is empty, not send message
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the MNS message fails to be sent. This parameter is not returned if the MNS message is sent.
	//
	// example:
	//
	// MessageConfigEmpty
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the message that indicates the MNS message is sent. This parameter is not returned if the MNS message fails to be sent.
	//
	// example:
	//
	// 4f3bc83233de4e2f81c7dade443e****
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult) GoString() string {
	return s.String()
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult) GetMessageId() *string {
	return s.MessageId
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult) SetErrorCode(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult {
	s.ErrorCode = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult) SetErrorMessage(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult {
	s.ErrorMessage = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult) SetMessageId(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult {
	s.MessageId = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult) Validate() error {
	return dara.Validate(s)
}

type QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput struct {
	// The input file of the media workflow.
	InputFile *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile `json:"InputFile,omitempty" xml:"InputFile,omitempty" type:"Struct"`
	// The user-defined data.
	//
	// example:
	//
	// example data ****
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput) GoString() string {
	return s.String()
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput) GetInputFile() *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile {
	return s.InputFile
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput) GetUserData() *string {
	return s.UserData
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput) SetInputFile(v *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput {
	s.InputFile = v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput) SetUserData(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput {
	s.UserData = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput) Validate() error {
	if s.InputFile != nil {
		if err := s.InputFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile struct {
	// The name of the OSS bucket in which the input file is stored.
	//
	// example:
	//
	// example-bucket-****
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The ID of the OSS region in which the input file resides.
	//
	// example:
	//
	// mps-cn-shanghai
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The name of the OSS object that is used as the input file.
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile) GoString() string {
	return s.String()
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile) GetBucket() *string {
	return s.Bucket
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile) GetLocation() *string {
	return s.Location
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile) GetObject() *string {
	return s.Object
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile) SetBucket(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile {
	s.Bucket = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile) SetLocation(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile {
	s.Location = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile) SetObject(v string) *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile {
	s.Object = &v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile) Validate() error {
	return dara.Validate(s)
}

type QueryMediaWorkflowExecutionListResponseBodyNonExistRunIds struct {
	RunId []*string `json:"RunId,omitempty" xml:"RunId,omitempty" type:"Repeated"`
}

func (s QueryMediaWorkflowExecutionListResponseBodyNonExistRunIds) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaWorkflowExecutionListResponseBodyNonExistRunIds) GoString() string {
	return s.String()
}

func (s *QueryMediaWorkflowExecutionListResponseBodyNonExistRunIds) GetRunId() []*string {
	return s.RunId
}

func (s *QueryMediaWorkflowExecutionListResponseBodyNonExistRunIds) SetRunId(v []*string) *QueryMediaWorkflowExecutionListResponseBodyNonExistRunIds {
	s.RunId = v
	return s
}

func (s *QueryMediaWorkflowExecutionListResponseBodyNonExistRunIds) Validate() error {
	return dara.Validate(s)
}
