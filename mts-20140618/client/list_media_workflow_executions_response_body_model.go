// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaWorkflowExecutionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaWorkflowExecutionList(v *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionList) *ListMediaWorkflowExecutionsResponseBody
	GetMediaWorkflowExecutionList() *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionList
	SetNextPageToken(v string) *ListMediaWorkflowExecutionsResponseBody
	GetNextPageToken() *string
	SetRequestId(v string) *ListMediaWorkflowExecutionsResponseBody
	GetRequestId() *string
}

type ListMediaWorkflowExecutionsResponseBody struct {
	// The details of the media workflows.
	MediaWorkflowExecutionList *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionList `json:"MediaWorkflowExecutionList,omitempty" xml:"MediaWorkflowExecutionList,omitempty" type:"Struct"`
	// The returned value of NextPageToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D1D5C080-8E2F-5030-8AB4-13092F17631B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMediaWorkflowExecutionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMediaWorkflowExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMediaWorkflowExecutionsResponseBody) GetMediaWorkflowExecutionList() *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionList {
	return s.MediaWorkflowExecutionList
}

func (s *ListMediaWorkflowExecutionsResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListMediaWorkflowExecutionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMediaWorkflowExecutionsResponseBody) SetMediaWorkflowExecutionList(v *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionList) *ListMediaWorkflowExecutionsResponseBody {
	s.MediaWorkflowExecutionList = v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBody) SetNextPageToken(v string) *ListMediaWorkflowExecutionsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBody) SetRequestId(v string) *ListMediaWorkflowExecutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBody) Validate() error {
	if s.MediaWorkflowExecutionList != nil {
		if err := s.MediaWorkflowExecutionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionList struct {
	MediaWorkflowExecution []*ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution `json:"MediaWorkflowExecution,omitempty" xml:"MediaWorkflowExecution,omitempty" type:"Repeated"`
}

func (s ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionList) String() string {
	return dara.Prettify(s)
}

func (s ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionList) GoString() string {
	return s.String()
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionList) GetMediaWorkflowExecution() []*ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution {
	return s.MediaWorkflowExecution
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionList) SetMediaWorkflowExecution(v []*ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionList {
	s.MediaWorkflowExecution = v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionList) Validate() error {
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

type ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution struct {
	// The activities that are executed in the media workflow.
	ActivityList *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList `json:"ActivityList,omitempty" xml:"ActivityList,omitempty" type:"Struct"`
	// The time when the media workflow was created.
	//
	// example:
	//
	// 2016-04-01T06:53:43Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The custom data of the media workflow.
	Input *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The ID of the media file. A media file contains all the information about a media workflow.
	//
	// example:
	//
	// 512046582a924698a41e0f8b0d2b****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The ID of the media workflow.
	//
	// example:
	//
	// 43b7335a4b1d4fe883670036affb****
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
	// 	- **running**: The execution is in progress.
	//
	// 	- **Completed**: The execution is complete.
	//
	// > A value of Completed indicates that the execution is complete. For the information about whether each activity, such as Transcode or Snapshot, is successful, check the status of the activity.
	//
	// 	- **Fail**: The execution failed.
	//
	// 	- **Success**: The execution was successful.
	//
	// example:
	//
	// Success
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) String() string {
	return dara.Prettify(s)
}

func (s ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) GoString() string {
	return s.String()
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) GetActivityList() *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList {
	return s.ActivityList
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) GetInput() *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput {
	return s.Input
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) GetMediaId() *string {
	return s.MediaId
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) GetName() *string {
	return s.Name
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) GetRunId() *string {
	return s.RunId
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) GetState() *string {
	return s.State
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) SetActivityList(v *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution {
	s.ActivityList = v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) SetCreationTime(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution {
	s.CreationTime = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) SetInput(v *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution {
	s.Input = v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) SetMediaId(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution {
	s.MediaId = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) SetMediaWorkflowId(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution {
	s.MediaWorkflowId = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) SetName(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution {
	s.Name = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) SetRunId(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution {
	s.RunId = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) SetState(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution {
	s.State = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecution) Validate() error {
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

type ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList struct {
	Activity []*ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity `json:"Activity,omitempty" xml:"Activity,omitempty" type:"Repeated"`
}

func (s ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList) String() string {
	return dara.Prettify(s)
}

func (s ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList) GoString() string {
	return s.String()
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList) GetActivity() []*ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity {
	return s.Activity
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList) SetActivity(v []*ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList {
	s.Activity = v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList) Validate() error {
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

type ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity struct {
	// The error code returned if the request failed.
	//
	// 	- The specific error code appears if the state of the activity is **Fail**.
	//
	// 	- This parameter is not returned if the state of the activity is **Success**.
	//
	// example:
	//
	// null
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The end time of the activity.
	//
	// example:
	//
	// 2016-04-01T06:54:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the job generated when the activity is executed. We recommend that you keep this ID for subsequent operation calls.
	//
	// example:
	//
	// 2376030d9d0849399cd20e20c876****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The message sent by Message Service (MNS) to notify the user of the job result.
	MNSMessageResult *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult `json:"MNSMessageResult,omitempty" xml:"MNSMessageResult,omitempty" type:"Struct"`
	// The error message returned if the request failed.
	//
	// 	- The detailed error message appears if the state of the activity is **Fail**.
	//
	// 	- This parameter is not returned if the state of the activity is **Success**.
	//
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the media workflow activity.
	//
	// > The name of an activity in a media workflow is unique.
	//
	// example:
	//
	// Act-2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The start time of the activity.
	//
	// example:
	//
	// 2016-04-01T06:53:45Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the activity. Valid values:
	//
	// 	- **Running**: The activity is being executed.
	//
	// 	- **Fail**: The activity failed to be executed.
	//
	// 	- **Skipped**: The activity was skipped.
	//
	// 	- **Success**: The activity was successfully executed.
	//
	// > For example, the high-definition and standard-definition transcoding activities are to be run after the analysis activity is complete. The system determines the activity to run based on the analysis result. If the definition of the input video content is insufficient, the high-definition transcoding activity may be skipped.
	//
	// example:
	//
	// Success
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The type of the media workflow activity. Valid values: Start, Snapshot, Transcode, Analysis, and Report. For more information, see [Methods supported for media workflows](https://help.aliyun.com/document_detail/68494.html).
	//
	// example:
	//
	// Start
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) String() string {
	return dara.Prettify(s)
}

func (s ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) GoString() string {
	return s.String()
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) GetCode() *string {
	return s.Code
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) GetEndTime() *string {
	return s.EndTime
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) GetJobId() *string {
	return s.JobId
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) GetMNSMessageResult() *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult {
	return s.MNSMessageResult
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) GetMessage() *string {
	return s.Message
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) GetName() *string {
	return s.Name
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) GetStartTime() *string {
	return s.StartTime
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) GetState() *string {
	return s.State
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) GetType() *string {
	return s.Type
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) SetCode(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity {
	s.Code = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) SetEndTime(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity {
	s.EndTime = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) SetJobId(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity {
	s.JobId = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) SetMNSMessageResult(v *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity {
	s.MNSMessageResult = v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) SetMessage(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity {
	s.Message = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) SetName(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity {
	s.Name = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) SetStartTime(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity {
	s.StartTime = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) SetState(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity {
	s.State = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) SetType(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity {
	s.Type = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivity) Validate() error {
	if s.MNSMessageResult != nil {
		if err := s.MNSMessageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult struct {
	// The error code returned if the job failed. If the job was successful, this parameter is not returned.
	//
	// example:
	//
	// The Topic/Queue config is empty, not send message
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the job failed. If the job was successful, this parameter is not returned.
	//
	// example:
	//
	// MessageConfigEmpty
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the success message. If the job failed, this parameter is not returned.
	//
	// example:
	//
	// 4f3bc83233de4e2f81c7dade443e****
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult) String() string {
	return dara.Prettify(s)
}

func (s ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult) GoString() string {
	return s.String()
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult) GetMessageId() *string {
	return s.MessageId
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult) SetErrorCode(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult {
	s.ErrorCode = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult) SetErrorMessage(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult {
	s.ErrorMessage = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult) SetMessageId(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult {
	s.MessageId = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult) Validate() error {
	return dara.Validate(s)
}

type ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput struct {
	// The information about the storage location of the input file of the media workflow in OSS.
	InputFile *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile `json:"InputFile,omitempty" xml:"InputFile,omitempty" type:"Struct"`
	// The custom data.
	//
	// example:
	//
	// example data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput) String() string {
	return dara.Prettify(s)
}

func (s ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput) GoString() string {
	return s.String()
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput) GetInputFile() *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile {
	return s.InputFile
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput) GetUserData() *string {
	return s.UserData
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput) SetInputFile(v *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput {
	s.InputFile = v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput) SetUserData(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput {
	s.UserData = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput) Validate() error {
	if s.InputFile != nil {
		if err := s.InputFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile struct {
	// The name of the OSS bucket in which the input media file is stored.
	//
	// example:
	//
	// example-bucket-****
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The OSS region in which the input file resides.
	//
	// example:
	//
	// cn-shanghai
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The name of the OSS object that is used as the input media file.
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile) String() string {
	return dara.Prettify(s)
}

func (s ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile) GoString() string {
	return s.String()
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile) GetBucket() *string {
	return s.Bucket
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile) GetLocation() *string {
	return s.Location
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile) GetObject() *string {
	return s.Object
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile) SetBucket(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile {
	s.Bucket = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile) SetLocation(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile {
	s.Location = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile) SetObject(v string) *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile {
	s.Object = &v
	return s
}

func (s *ListMediaWorkflowExecutionsResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile) Validate() error {
	return dara.Validate(s)
}
