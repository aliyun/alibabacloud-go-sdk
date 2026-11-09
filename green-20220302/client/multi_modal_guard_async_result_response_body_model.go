// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalGuardAsyncResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *MultiModalGuardAsyncResultResponseBody
	GetCode() *int32
	SetData(v *MultiModalGuardAsyncResultResponseBodyData) *MultiModalGuardAsyncResultResponseBody
	GetData() *MultiModalGuardAsyncResultResponseBodyData
	SetMessage(v string) *MultiModalGuardAsyncResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *MultiModalGuardAsyncResultResponseBody
	GetRequestId() *string
}

type MultiModalGuardAsyncResultResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *MultiModalGuardAsyncResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MultiModalGuardAsyncResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardAsyncResultResponseBody) GoString() string {
	return s.String()
}

func (s *MultiModalGuardAsyncResultResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *MultiModalGuardAsyncResultResponseBody) GetData() *MultiModalGuardAsyncResultResponseBodyData {
	return s.Data
}

func (s *MultiModalGuardAsyncResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MultiModalGuardAsyncResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MultiModalGuardAsyncResultResponseBody) SetCode(v int32) *MultiModalGuardAsyncResultResponseBody {
	s.Code = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBody) SetData(v *MultiModalGuardAsyncResultResponseBodyData) *MultiModalGuardAsyncResultResponseBody {
	s.Data = v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBody) SetMessage(v string) *MultiModalGuardAsyncResultResponseBody {
	s.Message = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBody) SetRequestId(v string) *MultiModalGuardAsyncResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MultiModalGuardAsyncResultResponseBodyData struct {
	// The segmented audio moderation results.
	AudioResult *MultiModalGuardAsyncResultResponseBodyDataAudioResult `json:"AudioResult,omitempty" xml:"AudioResult,omitempty" type:"Struct"`
	// The value of the dataId parameter passed in the API request. If this parameter is not passed in the request, this field is not returned.
	//
	// example:
	//
	// data1234
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The list of video frame capture results.
	FrameResult *MultiModalGuardAsyncResultResponseBodyDataFrameResult `json:"FrameResult,omitempty" xml:"FrameResult,omitempty" type:"Struct"`
	// The unique ID of the live stream.
	//
	// example:
	//
	// liveId
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// The data masking result.
	RedactResult *MultiModalGuardAsyncResultResponseBodyDataRedactResult `json:"RedactResult,omitempty" xml:"RedactResult,omitempty" type:"Struct"`
	// The suggestion. Valid values:
	//
	// - block: Block the content.
	//
	// - pass: Allow the content.
	//
	// - watch: Review the content.
	//
	// - mask: Mask the content.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The task ID.
	//
	// example:
	//
	// vi_f_xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s MultiModalGuardAsyncResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardAsyncResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *MultiModalGuardAsyncResultResponseBodyData) GetAudioResult() *MultiModalGuardAsyncResultResponseBodyDataAudioResult {
	return s.AudioResult
}

func (s *MultiModalGuardAsyncResultResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *MultiModalGuardAsyncResultResponseBodyData) GetFrameResult() *MultiModalGuardAsyncResultResponseBodyDataFrameResult {
	return s.FrameResult
}

func (s *MultiModalGuardAsyncResultResponseBodyData) GetLiveId() *string {
	return s.LiveId
}

func (s *MultiModalGuardAsyncResultResponseBodyData) GetRedactResult() *MultiModalGuardAsyncResultResponseBodyDataRedactResult {
	return s.RedactResult
}

func (s *MultiModalGuardAsyncResultResponseBodyData) GetSuggestion() *string {
	return s.Suggestion
}

func (s *MultiModalGuardAsyncResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *MultiModalGuardAsyncResultResponseBodyData) SetAudioResult(v *MultiModalGuardAsyncResultResponseBodyDataAudioResult) *MultiModalGuardAsyncResultResponseBodyData {
	s.AudioResult = v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyData) SetDataId(v string) *MultiModalGuardAsyncResultResponseBodyData {
	s.DataId = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyData) SetFrameResult(v *MultiModalGuardAsyncResultResponseBodyDataFrameResult) *MultiModalGuardAsyncResultResponseBodyData {
	s.FrameResult = v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyData) SetLiveId(v string) *MultiModalGuardAsyncResultResponseBodyData {
	s.LiveId = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyData) SetRedactResult(v *MultiModalGuardAsyncResultResponseBodyDataRedactResult) *MultiModalGuardAsyncResultResponseBodyData {
	s.RedactResult = v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyData) SetSuggestion(v string) *MultiModalGuardAsyncResultResponseBodyData {
	s.Suggestion = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyData) SetTaskId(v string) *MultiModalGuardAsyncResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyData) Validate() error {
	if s.AudioResult != nil {
		if err := s.AudioResult.Validate(); err != nil {
			return err
		}
	}
	if s.FrameResult != nil {
		if err := s.FrameResult.Validate(); err != nil {
			return err
		}
	}
	if s.RedactResult != nil {
		if err := s.RedactResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MultiModalGuardAsyncResultResponseBodyDataAudioResult struct {
	// The segment details.
	SliceDetails []*MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails `json:"SliceDetails,omitempty" xml:"SliceDetails,omitempty" type:"Repeated"`
	// The number of segments.
	//
	// example:
	//
	// 2
	SliceNum *int32 `json:"SliceNum,omitempty" xml:"SliceNum,omitempty"`
	// The suggestion.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s MultiModalGuardAsyncResultResponseBodyDataAudioResult) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardAsyncResultResponseBodyDataAudioResult) GoString() string {
	return s.String()
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResult) GetSliceDetails() []*MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails {
	return s.SliceDetails
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResult) GetSliceNum() *int32 {
	return s.SliceNum
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResult) SetSliceDetails(v []*MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails) *MultiModalGuardAsyncResultResponseBodyDataAudioResult {
	s.SliceDetails = v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResult) SetSliceNum(v int32) *MultiModalGuardAsyncResultResponseBodyDataAudioResult {
	s.SliceNum = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResult) SetSuggestion(v string) *MultiModalGuardAsyncResultResponseBodyDataAudioResult {
	s.Suggestion = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResult) Validate() error {
	if s.SliceDetails != nil {
		for _, item := range s.SliceDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails struct {
	// The moderation details.
	Detail []*MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// The end time of the segment, in seconds.
	//
	// example:
	//
	// 20
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the segment, in seconds.
	//
	// example:
	//
	// 0
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The moderation suggestion. Valid values:
	//
	// - block: Block the content.
	//
	// - pass: Allow the content.
	//
	// - watch: Review the content.
	//
	// - mask: Mask the content.
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The transcribed text of the audio segment.
	//
	// example:
	//
	// The weather is really nice today
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The temporary URL of the audio segment file.
	//
	// example:
	//
	// http://xxxx.abc.wav
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails) GoString() string {
	return s.String()
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails) GetDetail() []*MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetail {
	return s.Detail
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails) GetEndTime() *int64 {
	return s.EndTime
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails) GetStartTime() *int64 {
	return s.StartTime
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails) GetSuggestion() *string {
	return s.Suggestion
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails) GetText() *string {
	return s.Text
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails) GetUrl() *string {
	return s.Url
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails) SetDetail(v []*MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetail) *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails {
	s.Detail = v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails) SetEndTime(v int64) *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails {
	s.EndTime = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails) SetStartTime(v int64) *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails {
	s.StartTime = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails) SetSuggestion(v string) *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails {
	s.Suggestion = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails) SetText(v string) *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails {
	s.Text = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails) SetUrl(v string) *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails {
	s.Url = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails) Validate() error {
	if s.Detail != nil {
		for _, item := range s.Detail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetail struct {
	// The risk level. Valid values:
	//
	// - high: High risk. If a custom dictionary is hit, the risk level is high by default.
	//
	// - medium: Medium risk.
	//
	// - low: Low risk.
	//
	// - none: No risk detected.
	//
	// example:
	//
	// high
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The list of moderation results.
	Result []*MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The moderation suggestion. Valid values:
	//
	// - block: Block the content.
	//
	// - pass: Allow the content.
	//
	// - watch: Review the content.
	//
	// - mask: Mask the content.
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The moderation type. Valid values:
	//
	// - contentModeration: Content compliance moderation.
	//
	// - promptAttack: Prompt attack detection.
	//
	// - sensitiveData: Sensitive content detection.
	//
	// - modelHallucination: Model hallucination detection.
	//
	// - maliciousFile: Malicious file detection.
	//
	// example:
	//
	// contentModeration
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetail) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetail) GoString() string {
	return s.String()
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetail) GetLevel() *string {
	return s.Level
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetail) GetResult() []*MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult {
	return s.Result
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetail) GetSuggestion() *string {
	return s.Suggestion
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetail) GetType() *string {
	return s.Type
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetail) SetLevel(v string) *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetail {
	s.Level = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetail) SetResult(v []*MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult) *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetail {
	s.Result = v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetail) SetSuggestion(v string) *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetail {
	s.Suggestion = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetail) SetType(v string) *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetail {
	s.Type = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetail) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult struct {
	// The confidence score. Valid values: 0 to 100. The value is accurate to two decimal places.
	//
	// example:
	//
	// 90
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The label description.
	//
	// example:
	//
	// Drugs
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The extended information of the moderation result.
	//
	// example:
	//
	// {}
	Ext interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// The label.
	//
	// example:
	//
	// drug
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The risk level. Valid values:
	//
	// - high: High risk. If a custom dictionary is hit, the risk level is high by default.
	//
	// - medium: Medium risk.
	//
	// - low: Low risk.
	//
	// - none: No risk detected.
	//
	// example:
	//
	// high
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult) GoString() string {
	return s.String()
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult) GetConfidence() *float32 {
	return s.Confidence
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult) GetDescription() *string {
	return s.Description
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult) GetExt() interface{} {
	return s.Ext
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult) GetLabel() *string {
	return s.Label
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult) GetLevel() *string {
	return s.Level
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult) SetConfidence(v float32) *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult {
	s.Confidence = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult) SetDescription(v string) *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult {
	s.Description = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult) SetExt(v interface{}) *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult {
	s.Ext = v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult) SetLabel(v string) *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult {
	s.Label = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult) SetLevel(v string) *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult {
	s.Level = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult) Validate() error {
	return dara.Validate(s)
}

type MultiModalGuardAsyncResultResponseBodyDataFrameResult struct {
	// The frame moderation results.
	Frames []*MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Repeated"`
	// The number of segments.
	//
	// example:
	//
	// 2
	SliceNum *int32 `json:"SliceNum,omitempty" xml:"SliceNum,omitempty"`
	// The moderation suggestion. Valid values:
	//
	// - block: Block the content.
	//
	// - pass: Allow the content.
	//
	// - watch: Review the content.
	//
	// - mask: Mask the content.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s MultiModalGuardAsyncResultResponseBodyDataFrameResult) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardAsyncResultResponseBodyDataFrameResult) GoString() string {
	return s.String()
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResult) GetFrames() []*MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames {
	return s.Frames
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResult) GetSliceNum() *int32 {
	return s.SliceNum
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResult) SetFrames(v []*MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames) *MultiModalGuardAsyncResultResponseBodyDataFrameResult {
	s.Frames = v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResult) SetSliceNum(v int32) *MultiModalGuardAsyncResultResponseBodyDataFrameResult {
	s.SliceNum = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResult) SetSuggestion(v string) *MultiModalGuardAsyncResultResponseBodyDataFrameResult {
	s.Suggestion = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResult) Validate() error {
	if s.Frames != nil {
		for _, item := range s.Frames {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames struct {
	// The list of moderation results.
	Detail []*MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// 401
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// parameter invalid
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The position of the captured frame.
	//
	// example:
	//
	// 1.5
	Offset *float32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The moderation suggestion. Valid values:
	//
	// - block: Block the content.
	//
	// - pass: Allow the content.
	//
	// - watch: Review the content.
	//
	// - mask: Mask the content.
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The absolute timestamp, in milliseconds.
	//
	// example:
	//
	// 1684559739000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The temporary URL of the segment.
	//
	// example:
	//
	// https://xxx.jpeg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames) GoString() string {
	return s.String()
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames) GetDetail() []*MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetail {
	return s.Detail
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames) GetOffset() *float32 {
	return s.Offset
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames) GetSuggestion() *string {
	return s.Suggestion
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames) GetUrl() *string {
	return s.Url
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames) SetDetail(v []*MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetail) *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames {
	s.Detail = v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames) SetErrorCode(v int32) *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames {
	s.ErrorCode = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames) SetErrorMsg(v string) *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames {
	s.ErrorMsg = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames) SetOffset(v float32) *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames {
	s.Offset = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames) SetSuggestion(v string) *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames {
	s.Suggestion = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames) SetTimestamp(v int64) *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames {
	s.Timestamp = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames) SetUrl(v string) *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames {
	s.Url = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames) Validate() error {
	if s.Detail != nil {
		for _, item := range s.Detail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetail struct {
	// The risk level. Valid values:
	//
	// - high: High risk. If a custom dictionary is hit, the risk level is high by default.
	//
	// - medium: Medium risk.
	//
	// - low: Low risk.
	//
	// - none: No risk detected.
	//
	// example:
	//
	// low
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The moderation result.
	Result []*MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The moderation suggestion. Valid values:
	//
	// - block: Block the content.
	//
	// - pass: Allow the content.
	//
	// - watch: Review the content.
	//
	// - mask: Mask the content.
	//
	// example:
	//
	// watch
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The moderation type. Valid values:
	//
	// - contentModeration: Content compliance moderation.
	//
	// - promptAttack: Prompt attack detection.
	//
	// - sensitiveData: Sensitive content detection.
	//
	// - modelHallucination: Model hallucination detection.
	//
	// - maliciousFile: Malicious file detection.
	//
	// example:
	//
	// contentModeration
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetail) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetail) GoString() string {
	return s.String()
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetail) GetLevel() *string {
	return s.Level
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetail) GetResult() []*MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult {
	return s.Result
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetail) GetSuggestion() *string {
	return s.Suggestion
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetail) GetType() *string {
	return s.Type
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetail) SetLevel(v string) *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetail {
	s.Level = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetail) SetResult(v []*MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult) *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetail {
	s.Result = v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetail) SetSuggestion(v string) *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetail {
	s.Suggestion = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetail) SetType(v string) *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetail {
	s.Type = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetail) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult struct {
	// The confidence score. Valid values: 0 to 100. The value is accurate to two decimal places.
	//
	// example:
	//
	// 80
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The label description.
	//
	// example:
	//
	// Advertisement
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The extended information of the moderation result.
	//
	// example:
	//
	// {}
	Ext interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// The label.
	//
	// example:
	//
	// ad
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The risk level. Valid values:
	//
	// - high: High risk. If a custom dictionary is hit, the risk level is high by default.
	//
	// - medium: Medium risk.
	//
	// - low: Low risk.
	//
	// - none: No risk detected.
	//
	// example:
	//
	// loose
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult) GoString() string {
	return s.String()
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult) GetConfidence() *float32 {
	return s.Confidence
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult) GetDescription() *string {
	return s.Description
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult) GetExt() interface{} {
	return s.Ext
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult) GetLabel() *string {
	return s.Label
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult) GetLevel() *string {
	return s.Level
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult) SetConfidence(v float32) *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult {
	s.Confidence = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult) SetDescription(v string) *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult {
	s.Description = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult) SetExt(v interface{}) *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult {
	s.Ext = v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult) SetLabel(v string) *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult {
	s.Label = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult) SetLevel(v string) *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult {
	s.Level = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult) Validate() error {
	return dara.Validate(s)
}

type MultiModalGuardAsyncResultResponseBodyDataRedactResult struct {
	// The processing type.
	//
	// example:
	//
	// faceBlur
	ProcessType *string `json:"ProcessType,omitempty" xml:"ProcessType,omitempty"`
	// The name of the OSS bucket that stores the masked object.
	//
	// example:
	//
	// test_bucket
	ResultOssBucketName *string `json:"ResultOssBucketName,omitempty" xml:"ResultOssBucketName,omitempty"`
	// The name of the OSS object for the masked data.
	//
	// example:
	//
	// video.mp4
	ResultOssObjectName *string `json:"ResultOssObjectName,omitempty" xml:"ResultOssObjectName,omitempty"`
	// The region of the masked OSS object.
	//
	// example:
	//
	// cn-shanghai
	ResultOssRegionId *string `json:"ResultOssRegionId,omitempty" xml:"ResultOssRegionId,omitempty"`
	// The temporary URL of the masked object.
	//
	// example:
	//
	// https://XX
	ResultUrl *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
	// The URL of the object before masking.
	//
	// example:
	//
	// https://XX
	SourceUrl *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	// The status.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s MultiModalGuardAsyncResultResponseBodyDataRedactResult) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardAsyncResultResponseBodyDataRedactResult) GoString() string {
	return s.String()
}

func (s *MultiModalGuardAsyncResultResponseBodyDataRedactResult) GetProcessType() *string {
	return s.ProcessType
}

func (s *MultiModalGuardAsyncResultResponseBodyDataRedactResult) GetResultOssBucketName() *string {
	return s.ResultOssBucketName
}

func (s *MultiModalGuardAsyncResultResponseBodyDataRedactResult) GetResultOssObjectName() *string {
	return s.ResultOssObjectName
}

func (s *MultiModalGuardAsyncResultResponseBodyDataRedactResult) GetResultOssRegionId() *string {
	return s.ResultOssRegionId
}

func (s *MultiModalGuardAsyncResultResponseBodyDataRedactResult) GetResultUrl() *string {
	return s.ResultUrl
}

func (s *MultiModalGuardAsyncResultResponseBodyDataRedactResult) GetSourceUrl() *string {
	return s.SourceUrl
}

func (s *MultiModalGuardAsyncResultResponseBodyDataRedactResult) GetStatus() *string {
	return s.Status
}

func (s *MultiModalGuardAsyncResultResponseBodyDataRedactResult) SetProcessType(v string) *MultiModalGuardAsyncResultResponseBodyDataRedactResult {
	s.ProcessType = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataRedactResult) SetResultOssBucketName(v string) *MultiModalGuardAsyncResultResponseBodyDataRedactResult {
	s.ResultOssBucketName = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataRedactResult) SetResultOssObjectName(v string) *MultiModalGuardAsyncResultResponseBodyDataRedactResult {
	s.ResultOssObjectName = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataRedactResult) SetResultOssRegionId(v string) *MultiModalGuardAsyncResultResponseBodyDataRedactResult {
	s.ResultOssRegionId = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataRedactResult) SetResultUrl(v string) *MultiModalGuardAsyncResultResponseBodyDataRedactResult {
	s.ResultUrl = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataRedactResult) SetSourceUrl(v string) *MultiModalGuardAsyncResultResponseBodyDataRedactResult {
	s.SourceUrl = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataRedactResult) SetStatus(v string) *MultiModalGuardAsyncResultResponseBodyDataRedactResult {
	s.Status = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponseBodyDataRedactResult) Validate() error {
	return dara.Validate(s)
}
