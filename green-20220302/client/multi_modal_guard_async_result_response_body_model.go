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
	// example:
	//
	// 200
	Code *int32                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *MultiModalGuardAsyncResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	AudioResult *MultiModalGuardAsyncResultResponseBodyDataAudioResult `json:"AudioResult,omitempty" xml:"AudioResult,omitempty" type:"Struct"`
	// example:
	//
	// data1234
	DataId      *string                                                `json:"DataId,omitempty" xml:"DataId,omitempty"`
	FrameResult *MultiModalGuardAsyncResultResponseBodyDataFrameResult `json:"FrameResult,omitempty" xml:"FrameResult,omitempty" type:"Struct"`
	// example:
	//
	// liveId
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
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
	return nil
}

type MultiModalGuardAsyncResultResponseBodyDataAudioResult struct {
	SliceDetails []*MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetails `json:"SliceDetails,omitempty" xml:"SliceDetails,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	SliceNum *int32 `json:"SliceNum,omitempty" xml:"SliceNum,omitempty"`
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
	Detail []*MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// example:
	//
	// 30
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 0
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	Text       *string `json:"Text,omitempty" xml:"Text,omitempty"`
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
	// example:
	//
	// high
	Level  *string                                                                          `json:"Level,omitempty" xml:"Level,omitempty"`
	Result []*MultiModalGuardAsyncResultResponseBodyDataAudioResultSliceDetailsDetailResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
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
	// example:
	//
	// 90
	Confidence  *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Description *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// {}
	Ext interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// drug
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
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
	Frames []*MultiModalGuardAsyncResultResponseBodyDataFrameResultFrames `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	SliceNum *int32 `json:"SliceNum,omitempty" xml:"SliceNum,omitempty"`
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
	Detail []*MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// example:
	//
	// 1.5
	Offset *float32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// example:
	//
	// 1684559739000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
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
	// example:
	//
	// low
	Level  *string                                                                    `json:"Level,omitempty" xml:"Level,omitempty"`
	Result []*MultiModalGuardAsyncResultResponseBodyDataFrameResultFramesDetailResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// watch
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
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
	// example:
	//
	// 80
	Confidence  *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Description *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// {}
	Ext interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// ad
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
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
