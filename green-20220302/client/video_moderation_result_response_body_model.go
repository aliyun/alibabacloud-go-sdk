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
	// The returned HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The moderation results.
	Data *VideoModerationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
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
	// The voice moderation results. The moderation results contain a structure.
	AudioResult *VideoModerationResultResponseBodyDataAudioResult `json:"AudioResult,omitempty" xml:"AudioResult,omitempty" type:"Struct"`
	// The value of dataId that is specified in the API request. If this parameter is not specified in the API request, the dataId field is not available in the response.
	//
	// example:
	//
	// product_content-2055763
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The image moderation results. If the call is successful, the HTTP status code 200 and moderation results are returned. The moderation results contain a structure.
	FrameResult *VideoModerationResultResponseBodyDataFrameResult `json:"FrameResult,omitempty" xml:"FrameResult,omitempty" type:"Struct"`
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
	if s.FrameResult != nil {
		if err := s.FrameResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VideoModerationResultResponseBodyDataAudioResult struct {
	// Summary of voice labels.
	AudioSummarys []*VideoModerationResultResponseBodyDataAudioResultAudioSummarys `json:"AudioSummarys,omitempty" xml:"AudioSummarys,omitempty" type:"Repeated"`
	// Risk Level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The details about the text in the moderated voice. The value is a JSON array that contains one or more elements. Each element corresponds to a text entry.
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
	// The description of the labels.
	//
	// example:
	//
	// no risk
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The voice label.
	//
	// example:
	//
	// profanity
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The number of times that the label is matched.
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
	// The description of the labels.
	//
	// example:
	//
	// no risk
	Descriptions *string `json:"Descriptions,omitempty" xml:"Descriptions,omitempty"`
	// The end time of the text after voice-to-text conversion. Unit: seconds.
	//
	// example:
	//
	// 30
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The end timestamp of the segment. Unit: milliseconds.
	//
	// example:
	//
	// 1685245261939
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// {\\"consoleProduct\\":\\"slbnext\\"}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The details of the labels.
	//
	// example:
	//
	// porn
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// Risk Level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// Subcategory labels. Multiple labels are separated by commas (,).
	//
	// example:
	//
	// ""
	RiskTips *string `json:"RiskTips,omitempty" xml:"RiskTips,omitempty"`
	// The risk words that are hit. Multiple words are separated by commas (,).
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
	// The start time of the text after voice-to-text conversion. Unit: seconds.
	//
	// example:
	//
	// 0
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The start timestamp of the segment. Unit: milliseconds.
	//
	// example:
	//
	// 1659935002123
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The text converted from voice.
	//
	// example:
	//
	// Disgusting
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// If the moderation object is a voice stream, this parameter indicates the temporary access URL of the voice stream to which the text entry corresponds. The validity period of the URL is 30 minutes. You must prepare another URL to store the voice stream at the earliest opportunity.
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
	return dara.Validate(s)
}

type VideoModerationResultResponseBodyDataFrameResult struct {
	// The number of captured frames that are returned for the video file.
	//
	// example:
	//
	// 10
	FrameNum *int32 `json:"FrameNum,omitempty" xml:"FrameNum,omitempty"`
	// The summary of the labels against which captured frames are matched.
	FrameSummarys []*VideoModerationResultResponseBodyDataFrameResultFrameSummarys `json:"FrameSummarys,omitempty" xml:"FrameSummarys,omitempty" type:"Repeated"`
	// The information about the frames that match the labels.
	Frames []*VideoModerationResultResponseBodyDataFrameResultFrames `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Repeated"`
	// Risk Level.
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
	// The description of the result.
	//
	// example:
	//
	// no risk
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The label against which a captured frame is matched.
	//
	// example:
	//
	// violent_armedForces
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The number of times that the label is matched.
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
	// The interval between the start of the video file and the captured frame. Unit: seconds.
	//
	// example:
	//
	// 338
	Offset *float32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The results of frame moderation parameters such as the label parameter and the confidence parameter.
	Results []*VideoModerationResultResponseBodyDataFrameResultFramesResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// Risk Level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The temporary URL of a captured frame.
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
	// If a custom image library is hit, information about the custom image library is returned.
	CustomImage []*VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage `json:"CustomImage,omitempty" xml:"CustomImage,omitempty" type:"Repeated"`
	// Returns logo information when the video contains a logo.
	LogoData []*VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData `json:"LogoData,omitempty" xml:"LogoData,omitempty" type:"Repeated"`
	// If the video contains a specific figure, the code of the identified figure is returned.
	PublicFigure []*VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure `json:"PublicFigure,omitempty" xml:"PublicFigure,omitempty" type:"Repeated"`
	// The results of frame moderation parameters such as the label parameter and the confidence parameter.
	Result []*VideoModerationResultResponseBodyDataFrameResultFramesResultsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The moderation service that is called.
	//
	// example:
	//
	// tonalityImprove
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The information about the text hit in the image is returned.
	TextInImage map[string]interface{} `json:"TextInImage,omitempty" xml:"TextInImage,omitempty"`
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
	// The location of the logo.
	Location *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// Logo information.
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
	// The distance from the top-left corner of the text area to the y-axis, with the top-left corner of the image as the origin. Unit: pixels.
	//
	// example:
	//
	// 111
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The distance from the top-left corner of the text area to the x-axis, with the top-left corner of the image as the origin. Unit: pixels.
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
	// Confidence score, ranging from 0 to 100, with two decimal places.
	//
	// example:
	//
	// 99.10
	Confidence *int64 `json:"confidence,omitempty" xml:"confidence,omitempty"`
	// label
	//
	// example:
	//
	// pt_logotoSocialNetwork
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// Logo name.
	//
	// example:
	//
	// **logo
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
	// The information about the code of the identified figure.
	//
	// example:
	//
	// xxx001
	FigureId   *string                                                                              `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
	FigureName *string                                                                              `json:"FigureName,omitempty" xml:"FigureName,omitempty"`
	Location   []*VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Repeated"`
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
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
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
	// The score of the confidence level. Valid values: 0 to 100. The value is accurate to two decimal places.
	//
	// example:
	//
	// 50
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The description of the result.
	//
	// example:
	//
	// no risk
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The label returned after a frame is moderated. Multiple risk labels and the corresponding scores of confidence levels may be returned for a frame.
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
