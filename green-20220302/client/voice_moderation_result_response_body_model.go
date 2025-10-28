// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVoiceModerationResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *VoiceModerationResultResponseBody
	GetCode() *int32
	SetData(v *VoiceModerationResultResponseBodyData) *VoiceModerationResultResponseBody
	GetData() *VoiceModerationResultResponseBodyData
	SetMessage(v string) *VoiceModerationResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *VoiceModerationResultResponseBody
	GetRequestId() *string
}

type VoiceModerationResultResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *VoiceModerationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2881AD4F-638B-52A3-BA20-F74C5B1CEAE3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VoiceModerationResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VoiceModerationResultResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceModerationResultResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *VoiceModerationResultResponseBody) GetData() *VoiceModerationResultResponseBodyData {
	return s.Data
}

func (s *VoiceModerationResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VoiceModerationResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VoiceModerationResultResponseBody) SetCode(v int32) *VoiceModerationResultResponseBody {
	s.Code = &v
	return s
}

func (s *VoiceModerationResultResponseBody) SetData(v *VoiceModerationResultResponseBodyData) *VoiceModerationResultResponseBody {
	s.Data = v
	return s
}

func (s *VoiceModerationResultResponseBody) SetMessage(v string) *VoiceModerationResultResponseBody {
	s.Message = &v
	return s
}

func (s *VoiceModerationResultResponseBody) SetRequestId(v string) *VoiceModerationResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *VoiceModerationResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VoiceModerationResultResponseBodyData struct {
	// The ID of the moderated object.
	//
	// example:
	//
	// 26769ada6e264e7ba9aa048241e12be9
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The unique ID of the live stream.
	//
	// example:
	//
	// liveId
	LiveId       *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	ManualTaskId *string `json:"ManualTaskId,omitempty" xml:"ManualTaskId,omitempty"`
	// Risk Level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The moderation results of audio segments.
	SliceDetails []*VoiceModerationResultResponseBodyDataSliceDetails `json:"SliceDetails,omitempty" xml:"SliceDetails,omitempty" type:"Repeated"`
	// The task ID.
	//
	// example:
	//
	// kw24ihd0WGkdi5nniVZM@qOj-1x5Ibb
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The URL of the moderated content.
	//
	// example:
	//
	// https://aliyundoc.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s VoiceModerationResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VoiceModerationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *VoiceModerationResultResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *VoiceModerationResultResponseBodyData) GetLiveId() *string {
	return s.LiveId
}

func (s *VoiceModerationResultResponseBodyData) GetManualTaskId() *string {
	return s.ManualTaskId
}

func (s *VoiceModerationResultResponseBodyData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *VoiceModerationResultResponseBodyData) GetSliceDetails() []*VoiceModerationResultResponseBodyDataSliceDetails {
	return s.SliceDetails
}

func (s *VoiceModerationResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *VoiceModerationResultResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *VoiceModerationResultResponseBodyData) SetDataId(v string) *VoiceModerationResultResponseBodyData {
	s.DataId = &v
	return s
}

func (s *VoiceModerationResultResponseBodyData) SetLiveId(v string) *VoiceModerationResultResponseBodyData {
	s.LiveId = &v
	return s
}

func (s *VoiceModerationResultResponseBodyData) SetManualTaskId(v string) *VoiceModerationResultResponseBodyData {
	s.ManualTaskId = &v
	return s
}

func (s *VoiceModerationResultResponseBodyData) SetRiskLevel(v string) *VoiceModerationResultResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *VoiceModerationResultResponseBodyData) SetSliceDetails(v []*VoiceModerationResultResponseBodyDataSliceDetails) *VoiceModerationResultResponseBodyData {
	s.SliceDetails = v
	return s
}

func (s *VoiceModerationResultResponseBodyData) SetTaskId(v string) *VoiceModerationResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *VoiceModerationResultResponseBodyData) SetUrl(v string) *VoiceModerationResultResponseBodyData {
	s.Url = &v
	return s
}

func (s *VoiceModerationResultResponseBodyData) Validate() error {
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

type VoiceModerationResultResponseBodyDataSliceDetails struct {
	// The description of the labels.
	//
	// example:
	//
	// no risk
	Descriptions *string `json:"Descriptions,omitempty" xml:"Descriptions,omitempty"`
	// The end time of the audio segment in seconds.
	//
	// example:
	//
	// 10
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The end timestamp of the segment. Unit: milliseconds.
	//
	// example:
	//
	// 1678854649720
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// Extended fields.
	//
	// example:
	//
	// {\\"riskTips\\":\\"sexuality_Suggestive\\",\\"riskWords\\":\\"pxxxxy\\"}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The details of the labels.
	//
	// example:
	//
	// sexual_sounds
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// {}
	OriginAlgoResult map[string]interface{} `json:"OriginAlgoResult,omitempty" xml:"OriginAlgoResult,omitempty"`
	// Risk Level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The details of the risky content.
	//
	// example:
	//
	// sexuality_Suggestive
	RiskTips *string `json:"RiskTips,omitempty" xml:"RiskTips,omitempty"`
	// The term hit by the risky content.
	//
	// example:
	//
	// AAA,BBB,CCC
	RiskWords *string `json:"RiskWords,omitempty" xml:"RiskWords,omitempty"`
	// The risk score. Default range: 0 to 99.
	//
	// example:
	//
	// 87.01
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The start time of the audio segment in seconds.
	//
	// example:
	//
	// 0
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The start timestamp of the segment. Unit: milliseconds.
	//
	// example:
	//
	// 1678854649720
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The text converted from the audio segment.
	//
	// example:
	//
	// Disgusting
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The temporary URL of the audio segment.
	//
	// example:
	//
	// https://aliyundoc.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s VoiceModerationResultResponseBodyDataSliceDetails) String() string {
	return dara.Prettify(s)
}

func (s VoiceModerationResultResponseBodyDataSliceDetails) GoString() string {
	return s.String()
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) GetDescriptions() *string {
	return s.Descriptions
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) GetEndTime() *int64 {
	return s.EndTime
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) GetExtend() *string {
	return s.Extend
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) GetLabels() *string {
	return s.Labels
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) GetOriginAlgoResult() map[string]interface{} {
	return s.OriginAlgoResult
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) GetRiskTips() *string {
	return s.RiskTips
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) GetRiskWords() *string {
	return s.RiskWords
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) GetScore() *float32 {
	return s.Score
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) GetStartTime() *int64 {
	return s.StartTime
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) GetText() *string {
	return s.Text
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) GetUrl() *string {
	return s.Url
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetDescriptions(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.Descriptions = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetEndTime(v int64) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.EndTime = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetEndTimestamp(v int64) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.EndTimestamp = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetExtend(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.Extend = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetLabels(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.Labels = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetOriginAlgoResult(v map[string]interface{}) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.OriginAlgoResult = v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetRiskLevel(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.RiskLevel = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetRiskTips(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.RiskTips = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetRiskWords(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.RiskWords = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetScore(v float32) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.Score = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetStartTime(v int64) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.StartTime = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetStartTimestamp(v int64) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.StartTimestamp = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetText(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.Text = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetUrl(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.Url = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) Validate() error {
	return dara.Validate(s)
}
