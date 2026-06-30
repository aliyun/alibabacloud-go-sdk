// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoModerationResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *VideoModerationResultResponseBody
	GetCode() *int32
	SetData(v *VideoModerationResultResponseBodyData) *VideoModerationResultResponseBody
	GetData() *VideoModerationResultResponseBodyData
	SetMessage(v string) *VideoModerationResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *VideoModerationResultResponseBody
	GetRequestId() *string
}

type VideoModerationResultResponseBody struct {
	// The response code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The moderation result data.
	Data *VideoModerationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success finished
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6CF2815C-C8C7-4A01-B52E-FF6E24F53492
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VideoModerationResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBody) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *VideoModerationResultResponseBody) GetData() *VideoModerationResultResponseBodyData {
	return s.Data
}

func (s *VideoModerationResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VideoModerationResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VideoModerationResultResponseBody) SetCode(v int32) *VideoModerationResultResponseBody {
	s.Code = &v
	return s
}

func (s *VideoModerationResultResponseBody) SetData(v *VideoModerationResultResponseBodyData) *VideoModerationResultResponseBody {
	s.Data = v
	return s
}

func (s *VideoModerationResultResponseBody) SetMessage(v string) *VideoModerationResultResponseBody {
	s.Message = &v
	return s
}

func (s *VideoModerationResultResponseBody) SetRequestId(v string) *VideoModerationResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *VideoModerationResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VideoModerationResultResponseBodyData struct {
	// The segmented results of video audio moderation.
	AudioResult *VideoModerationResultResponseBodyDataAudioResult `json:"AudioResult,omitempty" xml:"AudioResult,omitempty" type:"Struct"`
	// The value of dataId passed in the API request. This field is not returned if dataId was not specified in the request.
	//
	// example:
	//
	// product_content-2055763
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The extended information.
	Ext *VideoModerationResultResponseBodyDataExt `json:"Ext,omitempty" xml:"Ext,omitempty" type:"Struct"`
	// The list of video frame capture results.
	FrameResult *VideoModerationResultResponseBodyDataFrameResult `json:"FrameResult,omitempty" xml:"FrameResult,omitempty" type:"Struct"`
	// The unique ID of the live stream.
	//
	// example:
	//
	// liveId
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// The manual review task ID.
	//
	// example:
	//
	// xxxxx-xxxxx
	ManualTaskId *string `json:"ManualTaskId,omitempty" xml:"ManualTaskId,omitempty"`
	// The risk level, returned based on the configured high and low risk score thresholds. Valid values:
	//
	// - high: High risk.
	//
	// - medium: Medium risk.
	//
	//
	//
	// - low: Low risk.
	//
	// - none: No risk detected.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The task ID.
	//
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s VideoModerationResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyData) GetAudioResult() *VideoModerationResultResponseBodyDataAudioResult {
	return s.AudioResult
}

func (s *VideoModerationResultResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *VideoModerationResultResponseBodyData) GetExt() *VideoModerationResultResponseBodyDataExt {
	return s.Ext
}

func (s *VideoModerationResultResponseBodyData) GetFrameResult() *VideoModerationResultResponseBodyDataFrameResult {
	return s.FrameResult
}

func (s *VideoModerationResultResponseBodyData) GetLiveId() *string {
	return s.LiveId
}

func (s *VideoModerationResultResponseBodyData) GetManualTaskId() *string {
	return s.ManualTaskId
}

func (s *VideoModerationResultResponseBodyData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *VideoModerationResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *VideoModerationResultResponseBodyData) SetAudioResult(v *VideoModerationResultResponseBodyDataAudioResult) *VideoModerationResultResponseBodyData {
	s.AudioResult = v
	return s
}

func (s *VideoModerationResultResponseBodyData) SetDataId(v string) *VideoModerationResultResponseBodyData {
	s.DataId = &v
	return s
}

func (s *VideoModerationResultResponseBodyData) SetExt(v *VideoModerationResultResponseBodyDataExt) *VideoModerationResultResponseBodyData {
	s.Ext = v
	return s
}

func (s *VideoModerationResultResponseBodyData) SetFrameResult(v *VideoModerationResultResponseBodyDataFrameResult) *VideoModerationResultResponseBodyData {
	s.FrameResult = v
	return s
}

func (s *VideoModerationResultResponseBodyData) SetLiveId(v string) *VideoModerationResultResponseBodyData {
	s.LiveId = &v
	return s
}

func (s *VideoModerationResultResponseBodyData) SetManualTaskId(v string) *VideoModerationResultResponseBodyData {
	s.ManualTaskId = &v
	return s
}

func (s *VideoModerationResultResponseBodyData) SetRiskLevel(v string) *VideoModerationResultResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *VideoModerationResultResponseBodyData) SetTaskId(v string) *VideoModerationResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *VideoModerationResultResponseBodyData) Validate() error {
	if s.AudioResult != nil {
		if err := s.AudioResult.Validate(); err != nil {
			return err
		}
	}
	if s.Ext != nil {
		if err := s.Ext.Validate(); err != nil {
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

type VideoModerationResultResponseBodyDataAudioResult struct {
	// The audio label summary.
	AudioSummarys []*VideoModerationResultResponseBodyDataAudioResultAudioSummarys `json:"AudioSummarys,omitempty" xml:"AudioSummarys,omitempty" type:"Repeated"`
	// The risk level, returned based on the configured high and low risk score thresholds. Valid values:
	//
	// - high: High risk.
	//
	// - medium: Medium risk.
	//
	//
	//
	// - low: Low risk.
	//
	// - none: No risk detected.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The list of audio segments.
	SliceDetails []*VideoModerationResultResponseBodyDataAudioResultSliceDetails `json:"SliceDetails,omitempty" xml:"SliceDetails,omitempty" type:"Repeated"`
}

func (s VideoModerationResultResponseBodyDataAudioResult) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataAudioResult) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataAudioResult) GetAudioSummarys() []*VideoModerationResultResponseBodyDataAudioResultAudioSummarys {
	return s.AudioSummarys
}

func (s *VideoModerationResultResponseBodyDataAudioResult) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *VideoModerationResultResponseBodyDataAudioResult) GetSliceDetails() []*VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	return s.SliceDetails
}

func (s *VideoModerationResultResponseBodyDataAudioResult) SetAudioSummarys(v []*VideoModerationResultResponseBodyDataAudioResultAudioSummarys) *VideoModerationResultResponseBodyDataAudioResult {
	s.AudioSummarys = v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResult) SetRiskLevel(v string) *VideoModerationResultResponseBodyDataAudioResult {
	s.RiskLevel = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResult) SetSliceDetails(v []*VideoModerationResultResponseBodyDataAudioResultSliceDetails) *VideoModerationResultResponseBodyDataAudioResult {
	s.SliceDetails = v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResult) Validate() error {
	if s.AudioSummarys != nil {
		for _, item := range s.AudioSummarys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type VideoModerationResultResponseBodyDataAudioResultAudioSummarys struct {
	// The label descriptions.
	//
	// example:
	//
	// 疑似违禁内容
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The video audio label.
	//
	// example:
	//
	// profanity
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The number of times the label appears.
	//
	// example:
	//
	// 8
	LabelSum *int32 `json:"LabelSum,omitempty" xml:"LabelSum,omitempty"`
}

func (s VideoModerationResultResponseBodyDataAudioResultAudioSummarys) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataAudioResultAudioSummarys) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataAudioResultAudioSummarys) GetDescription() *string {
	return s.Description
}

func (s *VideoModerationResultResponseBodyDataAudioResultAudioSummarys) GetLabel() *string {
	return s.Label
}

func (s *VideoModerationResultResponseBodyDataAudioResultAudioSummarys) GetLabelSum() *int32 {
	return s.LabelSum
}

func (s *VideoModerationResultResponseBodyDataAudioResultAudioSummarys) SetDescription(v string) *VideoModerationResultResponseBodyDataAudioResultAudioSummarys {
	s.Description = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultAudioSummarys) SetLabel(v string) *VideoModerationResultResponseBodyDataAudioResultAudioSummarys {
	s.Label = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultAudioSummarys) SetLabelSum(v int32) *VideoModerationResultResponseBodyDataAudioResultAudioSummarys {
	s.LabelSum = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultAudioSummarys) Validate() error {
	return dara.Validate(s)
}

type VideoModerationResultResponseBodyDataAudioResultSliceDetails struct {
	// The label descriptions.
	//
	// example:
	//
	// 疑似违禁内容
	Descriptions *string `json:"Descriptions,omitempty" xml:"Descriptions,omitempty"`
	// The end time of the segment, in seconds.
	//
	// example:
	//
	// 30
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The end timestamp.
	//
	// example:
	//
	// 1685245261939
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The extended field.
	//
	// example:
	//
	// {\\"consoleProduct\\":\\"slbnext\\"}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The violation labels that are hit.
	//
	// example:
	//
	// porn
	Labels *string                                                               `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Result []*VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The risk level, returned based on the configured high and low risk score thresholds. Valid values:
	//
	// - high: High risk.
	//
	// - medium: Medium risk.
	//
	//
	//
	// - low: Low risk.
	//
	// - none: No risk detected.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The details of the hit risk.
	//
	// example:
	//
	// ""
	RiskTips *string `json:"RiskTips,omitempty" xml:"RiskTips,omitempty"`
	// The risk keywords that are hit.
	//
	// example:
	//
	// ""
	RiskWords *string `json:"RiskWords,omitempty" xml:"RiskWords,omitempty"`
	// The risk score. Default range: 0 to 99.
	//
	// example:
	//
	// 5
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The start time of the segment, in seconds.
	//
	// example:
	//
	// 0
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The start timestamp, in milliseconds.
	//
	// example:
	//
	// 1659935002123
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The transcribed text of the audio segment.
	//
	// example:
	//
	// 今天天气真不错
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The temporary URL of the audio segment file.
	//
	// example:
	//
	// http://xxxx.abc.img
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s VideoModerationResultResponseBodyDataAudioResultSliceDetails) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataAudioResultSliceDetails) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) GetDescriptions() *string {
	return s.Descriptions
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) GetEndTime() *int64 {
	return s.EndTime
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) GetExtend() *string {
	return s.Extend
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) GetLabels() *string {
	return s.Labels
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) GetResult() []*VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult {
	return s.Result
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) GetRiskTips() *string {
	return s.RiskTips
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) GetRiskWords() *string {
	return s.RiskWords
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) GetScore() *float32 {
	return s.Score
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) GetStartTime() *int64 {
	return s.StartTime
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) GetText() *string {
	return s.Text
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) GetUrl() *string {
	return s.Url
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetDescriptions(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.Descriptions = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetEndTime(v int64) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.EndTime = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetEndTimestamp(v int64) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.EndTimestamp = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetExtend(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.Extend = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetLabels(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.Labels = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetResult(v []*VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.Result = v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetRiskLevel(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.RiskLevel = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetRiskTips(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.RiskTips = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetRiskWords(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.RiskWords = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetScore(v float32) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.Score = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetStartTime(v int64) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.StartTime = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetStartTimestamp(v int64) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.StartTimestamp = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetText(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.Text = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetUrl(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.Url = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) Validate() error {
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

type VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult struct {
	Confidence    *float32                                                                           `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	CustomizedHit []*VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultCustomizedHit `json:"CustomizedHit,omitempty" xml:"CustomizedHit,omitempty" type:"Repeated"`
	Description   *string                                                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	Label         *string                                                                            `json:"Label,omitempty" xml:"Label,omitempty"`
	RiskLevel     *string                                                                            `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	RiskPositions []*VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultRiskPositions `json:"RiskPositions,omitempty" xml:"RiskPositions,omitempty" type:"Repeated"`
	RiskWords     *string                                                                            `json:"RiskWords,omitempty" xml:"RiskWords,omitempty"`
}

func (s VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult) GetConfidence() *float32 {
	return s.Confidence
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult) GetCustomizedHit() []*VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultCustomizedHit {
	return s.CustomizedHit
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult) GetDescription() *string {
	return s.Description
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult) GetLabel() *string {
	return s.Label
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult) GetRiskPositions() []*VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultRiskPositions {
	return s.RiskPositions
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult) GetRiskWords() *string {
	return s.RiskWords
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult) SetConfidence(v float32) *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult {
	s.Confidence = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult) SetCustomizedHit(v []*VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultCustomizedHit) *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult {
	s.CustomizedHit = v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult) SetDescription(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult {
	s.Description = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult) SetLabel(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult {
	s.Label = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult) SetRiskLevel(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult {
	s.RiskLevel = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult) SetRiskPositions(v []*VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultRiskPositions) *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult {
	s.RiskPositions = v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult) SetRiskWords(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult {
	s.RiskWords = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResult) Validate() error {
	if s.CustomizedHit != nil {
		for _, item := range s.CustomizedHit {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RiskPositions != nil {
		for _, item := range s.RiskPositions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultCustomizedHit struct {
	KeyWords *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	LibName  *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
}

func (s VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultCustomizedHit) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultCustomizedHit) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultCustomizedHit) GetKeyWords() *string {
	return s.KeyWords
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultCustomizedHit) GetLibName() *string {
	return s.LibName
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultCustomizedHit) SetKeyWords(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultCustomizedHit {
	s.KeyWords = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultCustomizedHit) SetLibName(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultCustomizedHit {
	s.LibName = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultCustomizedHit) Validate() error {
	return dara.Validate(s)
}

type VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultRiskPositions struct {
	EndPos   *int32  `json:"EndPos,omitempty" xml:"EndPos,omitempty"`
	RiskWord *string `json:"RiskWord,omitempty" xml:"RiskWord,omitempty"`
	StartPos *int32  `json:"StartPos,omitempty" xml:"StartPos,omitempty"`
}

func (s VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultRiskPositions) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultRiskPositions) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultRiskPositions) GetEndPos() *int32 {
	return s.EndPos
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultRiskPositions) GetRiskWord() *string {
	return s.RiskWord
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultRiskPositions) GetStartPos() *int32 {
	return s.StartPos
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultRiskPositions) SetEndPos(v int32) *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultRiskPositions {
	s.EndPos = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultRiskPositions) SetRiskWord(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultRiskPositions {
	s.RiskWord = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultRiskPositions) SetStartPos(v int32) *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultRiskPositions {
	s.StartPos = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetailsResultRiskPositions) Validate() error {
	return dara.Validate(s)
}

type VideoModerationResultResponseBodyDataExt struct {
	// The AIGC metadata detection result.
	AigcData *VideoModerationResultResponseBodyDataExtAigcData `json:"AigcData,omitempty" xml:"AigcData,omitempty" type:"Struct"`
}

func (s VideoModerationResultResponseBodyDataExt) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataExt) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataExt) GetAigcData() *VideoModerationResultResponseBodyDataExtAigcData {
	return s.AigcData
}

func (s *VideoModerationResultResponseBodyDataExt) SetAigcData(v *VideoModerationResultResponseBodyDataExtAigcData) *VideoModerationResultResponseBodyDataExt {
	s.AigcData = v
	return s
}

func (s *VideoModerationResultResponseBodyDataExt) Validate() error {
	if s.AigcData != nil {
		if err := s.AigcData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VideoModerationResultResponseBodyDataExtAigcData struct {
	// The AIGC metadata.
	AIGC *VideoModerationResultResponseBodyDataExtAigcDataAIGC `json:"AIGC,omitempty" xml:"AIGC,omitempty" type:"Struct"`
	// The detection result.
	//
	// example:
	//
	// None
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s VideoModerationResultResponseBodyDataExtAigcData) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataExtAigcData) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataExtAigcData) GetAIGC() *VideoModerationResultResponseBodyDataExtAigcDataAIGC {
	return s.AIGC
}

func (s *VideoModerationResultResponseBodyDataExtAigcData) GetResult() *string {
	return s.Result
}

func (s *VideoModerationResultResponseBodyDataExtAigcData) SetAIGC(v *VideoModerationResultResponseBodyDataExtAigcDataAIGC) *VideoModerationResultResponseBodyDataExtAigcData {
	s.AIGC = v
	return s
}

func (s *VideoModerationResultResponseBodyDataExtAigcData) SetResult(v string) *VideoModerationResultResponseBodyDataExtAigcData {
	s.Result = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataExtAigcData) Validate() error {
	if s.AIGC != nil {
		if err := s.AIGC.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VideoModerationResultResponseBodyDataExtAigcDataAIGC struct {
	// The code or name of the service provider, used to identify the content producer.
	//
	// example:
	//
	// 001191******M000100Y43
	ContentProducer *string `json:"ContentProducer,omitempty" xml:"ContentProducer,omitempty"`
	// The name, ID, or code of the propagation platform. For services that provide artificial intelligence-generated content, this value can be the same as ContentProducer.
	//
	// example:
	//
	// 001191******M000100Y43
	ContentPropagator *string `json:"ContentPropagator,omitempty" xml:"ContentPropagator,omitempty"`
	// Indicates whether the content is generated by artificial intelligence (AI). Valid values:
	//
	// - 1: The content is generated through artificial intelligence content generation.
	//
	// - 2: (Propagation platforms only) The content may be generated through artificial intelligence content generation.
	//
	// - 3: (Propagation platforms only) The content is suspected to be generated through artificial intelligence content generation.
	//
	// example:
	//
	// 1
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The content production ID, which is the unique identifier used by the production platform to trace synthesized content.
	//
	// example:
	//
	// 123******456
	ProduceID *string `json:"ProduceID,omitempty" xml:"ProduceID,omitempty"`
	// The content propagation ID, which is the unique identifier assigned by the propagation platform to the propagated AI-generated content.
	//
	// example:
	//
	// 123******456
	PropagateID *string `json:"PropagateID,omitempty" xml:"PropagateID,omitempty"`
	// The reserved field.
	//
	// This field can store information used by the content generation service provider for self-initiated security protection to safeguard content and identifier integrity. A hashing mechanism based on ContentProducer and ProduceID can be used to securely store and verify critical information.
	//
	// example:
	//
	// d41d**********427e
	ReservedCode1 *string `json:"ReservedCode1,omitempty" xml:"ReservedCode1,omitempty"`
	// The reserved field.
	//
	// This field can be used by the content propagation service provider for self-initiated security protection to safeguard content and identifier integrity. A hashing mechanism based on ContentProducer and ProduceID can be used to securely store and verify critical information.
	//
	// example:
	//
	// d41d**********427e
	ReservedCode2 *string `json:"ReservedCode2,omitempty" xml:"ReservedCode2,omitempty"`
}

func (s VideoModerationResultResponseBodyDataExtAigcDataAIGC) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataExtAigcDataAIGC) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataExtAigcDataAIGC) GetContentProducer() *string {
	return s.ContentProducer
}

func (s *VideoModerationResultResponseBodyDataExtAigcDataAIGC) GetContentPropagator() *string {
	return s.ContentPropagator
}

func (s *VideoModerationResultResponseBodyDataExtAigcDataAIGC) GetLabel() *string {
	return s.Label
}

func (s *VideoModerationResultResponseBodyDataExtAigcDataAIGC) GetProduceID() *string {
	return s.ProduceID
}

func (s *VideoModerationResultResponseBodyDataExtAigcDataAIGC) GetPropagateID() *string {
	return s.PropagateID
}

func (s *VideoModerationResultResponseBodyDataExtAigcDataAIGC) GetReservedCode1() *string {
	return s.ReservedCode1
}

func (s *VideoModerationResultResponseBodyDataExtAigcDataAIGC) GetReservedCode2() *string {
	return s.ReservedCode2
}

func (s *VideoModerationResultResponseBodyDataExtAigcDataAIGC) SetContentProducer(v string) *VideoModerationResultResponseBodyDataExtAigcDataAIGC {
	s.ContentProducer = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataExtAigcDataAIGC) SetContentPropagator(v string) *VideoModerationResultResponseBodyDataExtAigcDataAIGC {
	s.ContentPropagator = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataExtAigcDataAIGC) SetLabel(v string) *VideoModerationResultResponseBodyDataExtAigcDataAIGC {
	s.Label = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataExtAigcDataAIGC) SetProduceID(v string) *VideoModerationResultResponseBodyDataExtAigcDataAIGC {
	s.ProduceID = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataExtAigcDataAIGC) SetPropagateID(v string) *VideoModerationResultResponseBodyDataExtAigcDataAIGC {
	s.PropagateID = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataExtAigcDataAIGC) SetReservedCode1(v string) *VideoModerationResultResponseBodyDataExtAigcDataAIGC {
	s.ReservedCode1 = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataExtAigcDataAIGC) SetReservedCode2(v string) *VideoModerationResultResponseBodyDataExtAigcDataAIGC {
	s.ReservedCode2 = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataExtAigcDataAIGC) Validate() error {
	return dara.Validate(s)
}

type VideoModerationResultResponseBodyDataFrameResult struct {
	// The number of result frames.
	//
	// example:
	//
	// 10
	FrameNum *int32 `json:"FrameNum,omitempty" xml:"FrameNum,omitempty"`
	// The video frame label summary.
	FrameSummarys []*VideoModerationResultResponseBodyDataFrameResultFrameSummarys `json:"FrameSummarys,omitempty" xml:"FrameSummarys,omitempty" type:"Repeated"`
	// The information about video frames that contain hit labels.
	Frames []*VideoModerationResultResponseBodyDataFrameResultFrames `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Repeated"`
	// The risk level, returned based on the configured high and low risk score thresholds. Valid values:
	//
	// - high: High risk.
	//
	// - medium: Medium risk.
	//
	//
	//
	// - low: Low risk.
	//
	// - none: No risk detected.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s VideoModerationResultResponseBodyDataFrameResult) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResult) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResult) GetFrameNum() *int32 {
	return s.FrameNum
}

func (s *VideoModerationResultResponseBodyDataFrameResult) GetFrameSummarys() []*VideoModerationResultResponseBodyDataFrameResultFrameSummarys {
	return s.FrameSummarys
}

func (s *VideoModerationResultResponseBodyDataFrameResult) GetFrames() []*VideoModerationResultResponseBodyDataFrameResultFrames {
	return s.Frames
}

func (s *VideoModerationResultResponseBodyDataFrameResult) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *VideoModerationResultResponseBodyDataFrameResult) SetFrameNum(v int32) *VideoModerationResultResponseBodyDataFrameResult {
	s.FrameNum = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResult) SetFrameSummarys(v []*VideoModerationResultResponseBodyDataFrameResultFrameSummarys) *VideoModerationResultResponseBodyDataFrameResult {
	s.FrameSummarys = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResult) SetFrames(v []*VideoModerationResultResponseBodyDataFrameResultFrames) *VideoModerationResultResponseBodyDataFrameResult {
	s.Frames = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResult) SetRiskLevel(v string) *VideoModerationResultResponseBodyDataFrameResult {
	s.RiskLevel = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResult) Validate() error {
	if s.FrameSummarys != nil {
		for _, item := range s.FrameSummarys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type VideoModerationResultResponseBodyDataFrameResultFrameSummarys struct {
	// The description of the Label field.
	//
	// example:
	//
	// 未检测出风险
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The video frame label.
	//
	// example:
	//
	// violent_armedForces
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The number of times the label appears.
	//
	// example:
	//
	// 8
	LabelSum *int32 `json:"LabelSum,omitempty" xml:"LabelSum,omitempty"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFrameSummarys) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFrameSummarys) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrameSummarys) GetDescription() *string {
	return s.Description
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrameSummarys) GetLabel() *string {
	return s.Label
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrameSummarys) GetLabelSum() *int32 {
	return s.LabelSum
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrameSummarys) SetDescription(v string) *VideoModerationResultResponseBodyDataFrameResultFrameSummarys {
	s.Description = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrameSummarys) SetLabel(v string) *VideoModerationResultResponseBodyDataFrameResultFrameSummarys {
	s.Label = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrameSummarys) SetLabelSum(v int32) *VideoModerationResultResponseBodyDataFrameResultFrameSummarys {
	s.LabelSum = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrameSummarys) Validate() error {
	return dara.Validate(s)
}

type VideoModerationResultResponseBodyDataFrameResultFrames struct {
	// The offset of the captured frame.
	//
	// example:
	//
	// 338
	Offset *float32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The frame detection result details.
	Results []*VideoModerationResultResponseBodyDataFrameResultFramesResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// The risk level, returned based on the configured high and low risk score thresholds. Valid values:
	//
	// - high: High risk.
	//
	// - medium: Medium risk.
	//
	//
	//
	// - low: Low risk.
	//
	// - none: No risk detected.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The temporary URL of the captured frame image.
	//
	// example:
	//
	// http://xxxx.abc.jpg
	TempUrl *string `json:"TempUrl,omitempty" xml:"TempUrl,omitempty"`
	// The absolute timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1684559739000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFrames) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFrames) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrames) GetOffset() *float32 {
	return s.Offset
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrames) GetResults() []*VideoModerationResultResponseBodyDataFrameResultFramesResults {
	return s.Results
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrames) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrames) GetTempUrl() *string {
	return s.TempUrl
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrames) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrames) SetOffset(v float32) *VideoModerationResultResponseBodyDataFrameResultFrames {
	s.Offset = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrames) SetResults(v []*VideoModerationResultResponseBodyDataFrameResultFramesResults) *VideoModerationResultResponseBodyDataFrameResultFrames {
	s.Results = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrames) SetRiskLevel(v string) *VideoModerationResultResponseBodyDataFrameResultFrames {
	s.RiskLevel = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrames) SetTempUrl(v string) *VideoModerationResultResponseBodyDataFrameResultFrames {
	s.TempUrl = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrames) SetTimestamp(v int64) *VideoModerationResultResponseBodyDataFrameResultFrames {
	s.Timestamp = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrames) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type VideoModerationResultResponseBodyDataFrameResultFramesResults struct {
	// The custom image library information that is hit. This field is returned only when a custom image library is hit.
	CustomImage []*VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage `json:"CustomImage,omitempty" xml:"CustomImage,omitempty" type:"Repeated"`
	// The logo information returned when the video contains a logo.
	LogoData []*VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData `json:"LogoData,omitempty" xml:"LogoData,omitempty" type:"Repeated"`
	// The recognized public figure codes returned when the video contains specific public figures.
	PublicFigure []*VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure `json:"PublicFigure,omitempty" xml:"PublicFigure,omitempty" type:"Repeated"`
	// The hit result details.
	Result []*VideoModerationResultResponseBodyDataFrameResultFramesResultsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The image moderation service type.
	//
	// example:
	//
	// tonalityImprove
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The text information in the image that is hit.
	TextInImage map[string]interface{} `json:"TextInImage,omitempty" xml:"TextInImage,omitempty"`
	// The foundation model result.
	VlContent *VideoModerationResultResponseBodyDataFrameResultFramesResultsVlContent `json:"VlContent,omitempty" xml:"VlContent,omitempty" type:"Struct"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResults) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResults) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) GetCustomImage() []*VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage {
	return s.CustomImage
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) GetLogoData() []*VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData {
	return s.LogoData
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) GetPublicFigure() []*VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure {
	return s.PublicFigure
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) GetResult() []*VideoModerationResultResponseBodyDataFrameResultFramesResultsResult {
	return s.Result
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) GetService() *string {
	return s.Service
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) GetTextInImage() map[string]interface{} {
	return s.TextInImage
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) GetVlContent() *VideoModerationResultResponseBodyDataFrameResultFramesResultsVlContent {
	return s.VlContent
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) SetCustomImage(v []*VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage) *VideoModerationResultResponseBodyDataFrameResultFramesResults {
	s.CustomImage = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) SetLogoData(v []*VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData) *VideoModerationResultResponseBodyDataFrameResultFramesResults {
	s.LogoData = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) SetPublicFigure(v []*VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure) *VideoModerationResultResponseBodyDataFrameResultFramesResults {
	s.PublicFigure = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) SetResult(v []*VideoModerationResultResponseBodyDataFrameResultFramesResultsResult) *VideoModerationResultResponseBodyDataFrameResultFramesResults {
	s.Result = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) SetService(v string) *VideoModerationResultResponseBodyDataFrameResultFramesResults {
	s.Service = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) SetTextInImage(v map[string]interface{}) *VideoModerationResultResponseBodyDataFrameResultFramesResults {
	s.TextInImage = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) SetVlContent(v *VideoModerationResultResponseBodyDataFrameResultFramesResultsVlContent) *VideoModerationResultResponseBodyDataFrameResultFramesResults {
	s.VlContent = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) Validate() error {
	if s.CustomImage != nil {
		for _, item := range s.CustomImage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LogoData != nil {
		for _, item := range s.LogoData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PublicFigure != nil {
		for _, item := range s.PublicFigure {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VlContent != nil {
		if err := s.VlContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage struct {
	// The ID of the custom image that is hit.
	//
	// example:
	//
	// 1234
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the custom image library that is hit.
	//
	// example:
	//
	// 12345678
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage) GetImageId() *string {
	return s.ImageId
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage) GetLibId() *string {
	return s.LibId
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage) SetImageId(v string) *VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage {
	s.ImageId = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage) SetLibId(v string) *VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage {
	s.LibId = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage) Validate() error {
	return dara.Validate(s)
}

type VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData struct {
	// The text line and coordinate information.
	Location *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// The logo identification information.
	Logo []*VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo `json:"Logo,omitempty" xml:"Logo,omitempty" type:"Repeated"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData) GetLocation() *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation {
	return s.Location
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData) GetLogo() []*VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo {
	return s.Logo
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData) SetLocation(v *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation) *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData {
	s.Location = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData) SetLogo(v []*VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo) *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData {
	s.Logo = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData) Validate() error {
	if s.Location != nil {
		if err := s.Location.Validate(); err != nil {
			return err
		}
	}
	if s.Logo != nil {
		for _, item := range s.Logo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation struct {
	// The height of the text area. Unit: pixels.
	//
	// example:
	//
	// 111
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The width of the text area. Unit: pixels.
	//
	// example:
	//
	// 111
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// The distance from the upper-left corner of the text area to the y-axis, with the upper-left corner of the image as the origin. Unit: pixels.
	//
	// example:
	//
	// 111
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The distance from the upper-left corner of the text area to the x-axis, with the upper-left corner of the image as the origin. Unit: pixels.
	//
	// example:
	//
	// 222
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation) GetH() *int32 {
	return s.H
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation) GetW() *int32 {
	return s.W
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation) GetX() *int32 {
	return s.X
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation) GetY() *int32 {
	return s.Y
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation) SetH(v int32) *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation {
	s.H = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation) SetW(v int32) *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation {
	s.W = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation) SetX(v int32) *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation {
	s.X = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation) SetY(v int32) *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation {
	s.Y = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation) Validate() error {
	return dara.Validate(s)
}

type VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo struct {
	// The confidence score, ranging from 0 to 100, rounded to two decimal places.
	//
	// example:
	//
	// 99.1
	Confidence *int64 `json:"confidence,omitempty" xml:"confidence,omitempty"`
	// The hit label.
	//
	// example:
	//
	// pt_logotoSocialNetwork
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// The logo name.
	//
	// example:
	//
	// **卫视
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo) GetConfidence() *int64 {
	return s.Confidence
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo) GetLabel() *string {
	return s.Label
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo) GetName() *string {
	return s.Name
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo) SetConfidence(v int64) *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo {
	s.Confidence = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo) SetLabel(v string) *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo {
	s.Label = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo) SetName(v string) *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo {
	s.Name = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo) Validate() error {
	return dara.Validate(s)
}

type VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure struct {
	// The code of the recognized public figure.
	//
	// example:
	//
	// xxx001
	FigureId *string `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
	// The name of the recognized public figure.
	//
	// example:
	//
	// 张三
	FigureName *string `json:"FigureName,omitempty" xml:"FigureName,omitempty"`
	// The location of the recognized public figure.
	Location []*VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Repeated"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure) GetFigureId() *string {
	return s.FigureId
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure) GetFigureName() *string {
	return s.FigureName
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure) GetLocation() []*VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation {
	return s.Location
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure) SetFigureId(v string) *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure {
	s.FigureId = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure) SetFigureName(v string) *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure {
	s.FigureName = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure) SetLocation(v []*VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation) *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure {
	s.Location = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure) Validate() error {
	if s.Location != nil {
		for _, item := range s.Location {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation struct {
	// The height.
	//
	// example:
	//
	// 222
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The width.
	//
	// example:
	//
	// 111
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// The x-coordinate of the starting point.
	//
	// example:
	//
	// 111
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The y-coordinate of the starting point.
	//
	// example:
	//
	// 222
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation) GetH() *int32 {
	return s.H
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation) GetW() *int32 {
	return s.W
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation) GetX() *int32 {
	return s.X
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation) GetY() *int32 {
	return s.Y
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation) SetH(v int32) *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation {
	s.H = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation) SetW(v int32) *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation {
	s.W = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation) SetX(v int32) *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation {
	s.X = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation) SetY(v int32) *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation {
	s.Y = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation) Validate() error {
	return dara.Validate(s)
}

type VideoModerationResultResponseBodyDataFrameResultFramesResultsResult struct {
	// The confidence score, ranging from 0 to 100, rounded to two decimal places.
	//
	// example:
	//
	// 50
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The description of the Label field.
	//
	// example:
	//
	// 未检测出风险
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The classification of the detection result.
	//
	// example:
	//
	// bloody
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsResult) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsResult) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsResult) GetConfidence() *float32 {
	return s.Confidence
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsResult) GetDescription() *string {
	return s.Description
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsResult) GetLabel() *string {
	return s.Label
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsResult) SetConfidence(v float32) *VideoModerationResultResponseBodyDataFrameResultFramesResultsResult {
	s.Confidence = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsResult) SetDescription(v string) *VideoModerationResultResponseBodyDataFrameResultFramesResultsResult {
	s.Description = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsResult) SetLabel(v string) *VideoModerationResultResponseBodyDataFrameResultFramesResultsResult {
	s.Label = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsResult) Validate() error {
	return dara.Validate(s)
}

type VideoModerationResultResponseBodyDataFrameResultFramesResultsVlContent struct {
	// The output text from the foundation model.
	//
	// example:
	//
	// in the picture XXX
	OutputText *string `json:"OutputText,omitempty" xml:"OutputText,omitempty"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsVlContent) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsVlContent) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsVlContent) GetOutputText() *string {
	return s.OutputText
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsVlContent) SetOutputText(v string) *VideoModerationResultResponseBodyDataFrameResultFramesResultsVlContent {
	s.OutputText = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsVlContent) Validate() error {
	return dara.Validate(s)
}
