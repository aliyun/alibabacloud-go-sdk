// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoModerationResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVideoModerationResultResponseBody
	GetCode() *string
	SetEndTime(v string) *GetVideoModerationResultResponseBody
	GetEndTime() *string
	SetEventId(v string) *GetVideoModerationResultResponseBody
	GetEventId() *string
	SetMessage(v string) *GetVideoModerationResultResponseBody
	GetMessage() *string
	SetModerationResult(v *GetVideoModerationResultResponseBodyModerationResult) *GetVideoModerationResultResponseBody
	GetModerationResult() *GetVideoModerationResultResponseBodyModerationResult
	SetProjectName(v string) *GetVideoModerationResultResponseBody
	GetProjectName() *string
	SetRequestId(v string) *GetVideoModerationResultResponseBody
	GetRequestId() *string
	SetStartTime(v string) *GetVideoModerationResultResponseBody
	GetStartTime() *string
	SetStatus(v string) *GetVideoModerationResultResponseBody
	GetStatus() *string
	SetTaskId(v string) *GetVideoModerationResultResponseBody
	GetTaskId() *string
	SetTaskType(v string) *GetVideoModerationResultResponseBody
	GetTaskType() *string
	SetUserData(v string) *GetVideoModerationResultResponseBody
	GetUserData() *string
}

type GetVideoModerationResultResponseBody struct {
	// example:
	//
	// ResourceNotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2023-04-03T10:20:56.87Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 05C-1XBQvsG2Tn5kBx2dUWo43******
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// The specified resource TaskId is not found.
	Message          *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	ModerationResult *GetVideoModerationResultResponseBodyModerationResult `json:"ModerationResult,omitempty" xml:"ModerationResult,omitempty" type:"Struct"`
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// VideoModeration-d0f0df1d-531d-4ab4-b353-e7f475******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2023-04-03T10:20:41.432Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// VideoModeration-d0f0df1d-531d-4ab4-b353-e7f4750******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// VideoModeration
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetVideoModerationResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoModerationResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoModerationResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVideoModerationResultResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetVideoModerationResultResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *GetVideoModerationResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVideoModerationResultResponseBody) GetModerationResult() *GetVideoModerationResultResponseBodyModerationResult {
	return s.ModerationResult
}

func (s *GetVideoModerationResultResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetVideoModerationResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoModerationResultResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetVideoModerationResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetVideoModerationResultResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVideoModerationResultResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *GetVideoModerationResultResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *GetVideoModerationResultResponseBody) SetCode(v string) *GetVideoModerationResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetEndTime(v string) *GetVideoModerationResultResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetEventId(v string) *GetVideoModerationResultResponseBody {
	s.EventId = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetMessage(v string) *GetVideoModerationResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetModerationResult(v *GetVideoModerationResultResponseBodyModerationResult) *GetVideoModerationResultResponseBody {
	s.ModerationResult = v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetProjectName(v string) *GetVideoModerationResultResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetRequestId(v string) *GetVideoModerationResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetStartTime(v string) *GetVideoModerationResultResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetStatus(v string) *GetVideoModerationResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetTaskId(v string) *GetVideoModerationResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetTaskType(v string) *GetVideoModerationResultResponseBody {
	s.TaskType = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) SetUserData(v string) *GetVideoModerationResultResponseBody {
	s.UserData = &v
	return s
}

func (s *GetVideoModerationResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetVideoModerationResultResponseBodyModerationResult struct {
	Categories []*string                                                   `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	Frames     *GetVideoModerationResultResponseBodyModerationResultFrames `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Struct"`
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// example:
	//
	// oss://test-bucket/test-object
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s GetVideoModerationResultResponseBodyModerationResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoModerationResultResponseBodyModerationResult) GoString() string {
	return s.String()
}

func (s *GetVideoModerationResultResponseBodyModerationResult) GetCategories() []*string {
	return s.Categories
}

func (s *GetVideoModerationResultResponseBodyModerationResult) GetFrames() *GetVideoModerationResultResponseBodyModerationResultFrames {
	return s.Frames
}

func (s *GetVideoModerationResultResponseBodyModerationResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetVideoModerationResultResponseBodyModerationResult) GetURI() *string {
	return s.URI
}

func (s *GetVideoModerationResultResponseBodyModerationResult) SetCategories(v []*string) *GetVideoModerationResultResponseBodyModerationResult {
	s.Categories = v
	return s
}

func (s *GetVideoModerationResultResponseBodyModerationResult) SetFrames(v *GetVideoModerationResultResponseBodyModerationResultFrames) *GetVideoModerationResultResponseBodyModerationResult {
	s.Frames = v
	return s
}

func (s *GetVideoModerationResultResponseBodyModerationResult) SetSuggestion(v string) *GetVideoModerationResultResponseBodyModerationResult {
	s.Suggestion = &v
	return s
}

func (s *GetVideoModerationResultResponseBodyModerationResult) SetURI(v string) *GetVideoModerationResultResponseBodyModerationResult {
	s.URI = &v
	return s
}

func (s *GetVideoModerationResultResponseBodyModerationResult) Validate() error {
	return dara.Validate(s)
}

type GetVideoModerationResultResponseBodyModerationResultFrames struct {
	BlockFrames []*GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames `json:"BlockFrames,omitempty" xml:"BlockFrames,omitempty" type:"Repeated"`
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetVideoModerationResultResponseBodyModerationResultFrames) String() string {
	return dara.Prettify(s)
}

func (s GetVideoModerationResultResponseBodyModerationResultFrames) GoString() string {
	return s.String()
}

func (s *GetVideoModerationResultResponseBodyModerationResultFrames) GetBlockFrames() []*GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames {
	return s.BlockFrames
}

func (s *GetVideoModerationResultResponseBodyModerationResultFrames) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetVideoModerationResultResponseBodyModerationResultFrames) SetBlockFrames(v []*GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames) *GetVideoModerationResultResponseBodyModerationResultFrames {
	s.BlockFrames = v
	return s
}

func (s *GetVideoModerationResultResponseBodyModerationResultFrames) SetTotalCount(v int32) *GetVideoModerationResultResponseBodyModerationResultFrames {
	s.TotalCount = &v
	return s
}

func (s *GetVideoModerationResultResponseBodyModerationResultFrames) Validate() error {
	return dara.Validate(s)
}

type GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames struct {
	// example:
	//
	// {"teat":"val"}
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// 1
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// example:
	//
	// 10
	Rate *float64 `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames) String() string {
	return dara.Prettify(s)
}

func (s GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames) GoString() string {
	return s.String()
}

func (s *GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames) GetLabel() *string {
	return s.Label
}

func (s *GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames) GetOffset() *int32 {
	return s.Offset
}

func (s *GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames) GetRate() *float64 {
	return s.Rate
}

func (s *GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames) SetLabel(v string) *GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames {
	s.Label = &v
	return s
}

func (s *GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames) SetOffset(v int32) *GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames {
	s.Offset = &v
	return s
}

func (s *GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames) SetRate(v float64) *GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames {
	s.Rate = &v
	return s
}

func (s *GetVideoModerationResultResponseBodyModerationResultFramesBlockFrames) Validate() error {
	return dara.Validate(s)
}
