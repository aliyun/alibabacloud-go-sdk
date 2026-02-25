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
	MediaWorkflowExecutionList *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionList `json:"MediaWorkflowExecutionList,omitempty" xml:"MediaWorkflowExecutionList,omitempty" type:"Struct"`
	NonExistRunIds             *QueryMediaWorkflowExecutionListResponseBodyNonExistRunIds             `json:"NonExistRunIds,omitempty" xml:"NonExistRunIds,omitempty" type:"Struct"`
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
	ActivityList    *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityList `json:"ActivityList,omitempty" xml:"ActivityList,omitempty" type:"Struct"`
	CreationTime    *string                                                                                                  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Input           *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInput        `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	MediaId         *string                                                                                                  `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	MediaWorkflowId *string                                                                                                  `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty"`
	Name            *string                                                                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	RunId           *string                                                                                                  `json:"RunId,omitempty" xml:"RunId,omitempty"`
	State           *string                                                                                                  `json:"State,omitempty" xml:"State,omitempty"`
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
	Code             *string                                                                                                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	EndTime          *string                                                                                                                          `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	JobId            *string                                                                                                                          `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MNSMessageResult *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionActivityListActivityMNSMessageResult `json:"MNSMessageResult,omitempty" xml:"MNSMessageResult,omitempty" type:"Struct"`
	Message          *string                                                                                                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Name             *string                                                                                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	StartTime        *string                                                                                                                          `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State            *string                                                                                                                          `json:"State,omitempty" xml:"State,omitempty"`
	Type             *string                                                                                                                          `json:"Type,omitempty" xml:"Type,omitempty"`
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
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	MessageId    *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
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
	InputFile *QueryMediaWorkflowExecutionListResponseBodyMediaWorkflowExecutionListMediaWorkflowExecutionInputInputFile `json:"InputFile,omitempty" xml:"InputFile,omitempty" type:"Struct"`
	UserData  *string                                                                                                    `json:"UserData,omitempty" xml:"UserData,omitempty"`
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
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
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
