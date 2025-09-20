// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageModerationResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetImageModerationResultResponseBody
	GetCode() *string
	SetEndTime(v string) *GetImageModerationResultResponseBody
	GetEndTime() *string
	SetEventId(v string) *GetImageModerationResultResponseBody
	GetEventId() *string
	SetMessage(v string) *GetImageModerationResultResponseBody
	GetMessage() *string
	SetModerationResult(v *GetImageModerationResultResponseBodyModerationResult) *GetImageModerationResultResponseBody
	GetModerationResult() *GetImageModerationResultResponseBodyModerationResult
	SetProjectName(v string) *GetImageModerationResultResponseBody
	GetProjectName() *string
	SetRequestId(v string) *GetImageModerationResultResponseBody
	GetRequestId() *string
	SetStartTime(v string) *GetImageModerationResultResponseBody
	GetStartTime() *string
	SetStatus(v string) *GetImageModerationResultResponseBody
	GetStatus() *string
	SetTaskId(v string) *GetImageModerationResultResponseBody
	GetTaskId() *string
	SetTaskType(v string) *GetImageModerationResultResponseBody
	GetTaskType() *string
	SetUserData(v string) *GetImageModerationResultResponseBody
	GetUserData() *string
}

type GetImageModerationResultResponseBody struct {
	// The error code of the task.
	//
	// example:
	//
	// ResourceNotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The end time of the task.
	//
	// example:
	//
	// 2023-04-03T09:44:32Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 1B6-1XBMX3BixLMILvXVGtlkr******
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The error message of the task.
	//
	// example:
	//
	// The specified resource TaskId is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The result of the image compliance detection task.
	ModerationResult *GetImageModerationResultResponseBodyModerationResult `json:"ModerationResult,omitempty" xml:"ModerationResult,omitempty" type:"Struct"`
	// The name of the project.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E6A120B1-BEB3-0F63-A7C2-0783B6******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the task.
	//
	// example:
	//
	// 2023-04-03T09:44:31.029Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
	//
	// 	- Running
	//
	// 	- Succeeded
	//
	// 	- Failed
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ImageModeration-ff207203-3f93-4645-a041-7b8f0f******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the task.
	//
	// example:
	//
	// ImageModeration
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The custom information.
	//
	// example:
	//
	// {
	//
	//       "fileId": "123"
	//
	// }
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetImageModerationResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageModerationResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageModerationResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetImageModerationResultResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetImageModerationResultResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *GetImageModerationResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetImageModerationResultResponseBody) GetModerationResult() *GetImageModerationResultResponseBodyModerationResult {
	return s.ModerationResult
}

func (s *GetImageModerationResultResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetImageModerationResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageModerationResultResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetImageModerationResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetImageModerationResultResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetImageModerationResultResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *GetImageModerationResultResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *GetImageModerationResultResponseBody) SetCode(v string) *GetImageModerationResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageModerationResultResponseBody) SetEndTime(v string) *GetImageModerationResultResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetImageModerationResultResponseBody) SetEventId(v string) *GetImageModerationResultResponseBody {
	s.EventId = &v
	return s
}

func (s *GetImageModerationResultResponseBody) SetMessage(v string) *GetImageModerationResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageModerationResultResponseBody) SetModerationResult(v *GetImageModerationResultResponseBodyModerationResult) *GetImageModerationResultResponseBody {
	s.ModerationResult = v
	return s
}

func (s *GetImageModerationResultResponseBody) SetProjectName(v string) *GetImageModerationResultResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetImageModerationResultResponseBody) SetRequestId(v string) *GetImageModerationResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageModerationResultResponseBody) SetStartTime(v string) *GetImageModerationResultResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetImageModerationResultResponseBody) SetStatus(v string) *GetImageModerationResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetImageModerationResultResponseBody) SetTaskId(v string) *GetImageModerationResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetImageModerationResultResponseBody) SetTaskType(v string) *GetImageModerationResultResponseBody {
	s.TaskType = &v
	return s
}

func (s *GetImageModerationResultResponseBody) SetUserData(v string) *GetImageModerationResultResponseBody {
	s.UserData = &v
	return s
}

func (s *GetImageModerationResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetImageModerationResultResponseBodyModerationResult struct {
	// List of categories.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The information about video and motion detection frames.
	Frames *GetImageModerationResultResponseBodyModerationResultFrames `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Struct"`
	// The recommended operation. Valid values:
	//
	// 	- pass: The image has passed the check. No action is required.
	//
	// 	- review: The image contains suspected violations and requires human review.
	//
	// 	- block: The image contains violations. Further actions, such as deleting or blocking the image, are recommended.
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The OSS URI of the file. The URI follows the oss://${bucketname}/${objectname} format. bucketname indicates the name of an OSS bucket that is in the same region as the current project, and objectname is the file path.
	//
	// example:
	//
	// oss://test-bucket/test-object
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s GetImageModerationResultResponseBodyModerationResult) String() string {
	return dara.Prettify(s)
}

func (s GetImageModerationResultResponseBodyModerationResult) GoString() string {
	return s.String()
}

func (s *GetImageModerationResultResponseBodyModerationResult) GetCategories() []*string {
	return s.Categories
}

func (s *GetImageModerationResultResponseBodyModerationResult) GetFrames() *GetImageModerationResultResponseBodyModerationResultFrames {
	return s.Frames
}

func (s *GetImageModerationResultResponseBodyModerationResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetImageModerationResultResponseBodyModerationResult) GetURI() *string {
	return s.URI
}

func (s *GetImageModerationResultResponseBodyModerationResult) SetCategories(v []*string) *GetImageModerationResultResponseBodyModerationResult {
	s.Categories = v
	return s
}

func (s *GetImageModerationResultResponseBodyModerationResult) SetFrames(v *GetImageModerationResultResponseBodyModerationResultFrames) *GetImageModerationResultResponseBodyModerationResult {
	s.Frames = v
	return s
}

func (s *GetImageModerationResultResponseBodyModerationResult) SetSuggestion(v string) *GetImageModerationResultResponseBodyModerationResult {
	s.Suggestion = &v
	return s
}

func (s *GetImageModerationResultResponseBodyModerationResult) SetURI(v string) *GetImageModerationResultResponseBodyModerationResult {
	s.URI = &v
	return s
}

func (s *GetImageModerationResultResponseBodyModerationResult) Validate() error {
	return dara.Validate(s)
}

type GetImageModerationResultResponseBodyModerationResultFrames struct {
	// The violated frames.
	BlockFrames []*GetImageModerationResultResponseBodyModerationResultFramesBlockFrames `json:"BlockFrames,omitempty" xml:"BlockFrames,omitempty" type:"Repeated"`
	// The total number of detected frames.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetImageModerationResultResponseBodyModerationResultFrames) String() string {
	return dara.Prettify(s)
}

func (s GetImageModerationResultResponseBodyModerationResultFrames) GoString() string {
	return s.String()
}

func (s *GetImageModerationResultResponseBodyModerationResultFrames) GetBlockFrames() []*GetImageModerationResultResponseBodyModerationResultFramesBlockFrames {
	return s.BlockFrames
}

func (s *GetImageModerationResultResponseBodyModerationResultFrames) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetImageModerationResultResponseBodyModerationResultFrames) SetBlockFrames(v []*GetImageModerationResultResponseBodyModerationResultFramesBlockFrames) *GetImageModerationResultResponseBodyModerationResultFrames {
	s.BlockFrames = v
	return s
}

func (s *GetImageModerationResultResponseBodyModerationResultFrames) SetTotalCount(v int32) *GetImageModerationResultResponseBodyModerationResultFrames {
	s.TotalCount = &v
	return s
}

func (s *GetImageModerationResultResponseBodyModerationResultFrames) Validate() error {
	return dara.Validate(s)
}

type GetImageModerationResultResponseBodyModerationResultFramesBlockFrames struct {
	// The label of the violation.
	//
	// example:
	//
	// {
	//
	//       "test": "val"
	//
	// }
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The offset of the frame.
	//
	// example:
	//
	// 2
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The confidence level of the violation.
	//
	// example:
	//
	// 30
	Rate *float64 `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s GetImageModerationResultResponseBodyModerationResultFramesBlockFrames) String() string {
	return dara.Prettify(s)
}

func (s GetImageModerationResultResponseBodyModerationResultFramesBlockFrames) GoString() string {
	return s.String()
}

func (s *GetImageModerationResultResponseBodyModerationResultFramesBlockFrames) GetLabel() *string {
	return s.Label
}

func (s *GetImageModerationResultResponseBodyModerationResultFramesBlockFrames) GetOffset() *int32 {
	return s.Offset
}

func (s *GetImageModerationResultResponseBodyModerationResultFramesBlockFrames) GetRate() *float64 {
	return s.Rate
}

func (s *GetImageModerationResultResponseBodyModerationResultFramesBlockFrames) SetLabel(v string) *GetImageModerationResultResponseBodyModerationResultFramesBlockFrames {
	s.Label = &v
	return s
}

func (s *GetImageModerationResultResponseBodyModerationResultFramesBlockFrames) SetOffset(v int32) *GetImageModerationResultResponseBodyModerationResultFramesBlockFrames {
	s.Offset = &v
	return s
}

func (s *GetImageModerationResultResponseBodyModerationResultFramesBlockFrames) SetRate(v float64) *GetImageModerationResultResponseBodyModerationResultFramesBlockFrames {
	s.Rate = &v
	return s
}

func (s *GetImageModerationResultResponseBodyModerationResultFramesBlockFrames) Validate() error {
	return dara.Validate(s)
}
