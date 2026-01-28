// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeVideoCastCrewListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizeVideoCastCrewListResponseBodyData) *RecognizeVideoCastCrewListResponseBody
	GetData() *RecognizeVideoCastCrewListResponseBodyData
	SetMessage(v string) *RecognizeVideoCastCrewListResponseBody
	GetMessage() *string
	SetRequestId(v string) *RecognizeVideoCastCrewListResponseBody
	GetRequestId() *string
}

type RecognizeVideoCastCrewListResponseBody struct {
	Data    *RecognizeVideoCastCrewListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE5B1A95-064F-1C5E-A6FE-FEE0D734A632
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBody) GetData() *RecognizeVideoCastCrewListResponseBodyData {
	return s.Data
}

func (s *RecognizeVideoCastCrewListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RecognizeVideoCastCrewListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeVideoCastCrewListResponseBody) SetData(v *RecognizeVideoCastCrewListResponseBodyData) *RecognizeVideoCastCrewListResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBody) SetMessage(v string) *RecognizeVideoCastCrewListResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBody) SetRequestId(v string) *RecognizeVideoCastCrewListResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RecognizeVideoCastCrewListResponseBodyData struct {
	CastResults []*RecognizeVideoCastCrewListResponseBodyDataCastResults `json:"CastResults,omitempty" xml:"CastResults,omitempty" type:"Repeated"`
	OcrResults  []*RecognizeVideoCastCrewListResponseBodyDataOcrResults  `json:"OcrResults,omitempty" xml:"OcrResults,omitempty" type:"Repeated"`
	// example:
	//
	// http://vibktprfx-prod-prod-media-ai-cn-shanghai.oss-cn-shanghai.aliyuncs.com/video-ocr/1665475907_bGHMygKsFw.json?Expires=1665477707&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=6KQb9OXQldsg30w%2FNurHwAbjiJs%3D
	OcrResultsUrl *string `json:"OcrResultsUrl,omitempty" xml:"OcrResultsUrl,omitempty"`
	// example:
	//
	// http://vibktprfx-prod-prod-media-ai-cn-shanghai.oss-cn-shanghai.aliyuncs.com/video-ocr/1665475907_VSRvetTHon.json?Expires=1665477707&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=wfQviVVSyVRLPVlHDKXi6cTefHY%3D
	OcrVideoResultsUrl *string                                                       `json:"OcrVideoResultsUrl,omitempty" xml:"OcrVideoResultsUrl,omitempty"`
	SubtitlesResults   []*RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults `json:"SubtitlesResults,omitempty" xml:"SubtitlesResults,omitempty" type:"Repeated"`
	VideoOcrResults    []*RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults  `json:"VideoOcrResults,omitempty" xml:"VideoOcrResults,omitempty" type:"Repeated"`
}

func (s RecognizeVideoCastCrewListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBodyData) GetCastResults() []*RecognizeVideoCastCrewListResponseBodyDataCastResults {
	return s.CastResults
}

func (s *RecognizeVideoCastCrewListResponseBodyData) GetOcrResults() []*RecognizeVideoCastCrewListResponseBodyDataOcrResults {
	return s.OcrResults
}

func (s *RecognizeVideoCastCrewListResponseBodyData) GetOcrResultsUrl() *string {
	return s.OcrResultsUrl
}

func (s *RecognizeVideoCastCrewListResponseBodyData) GetOcrVideoResultsUrl() *string {
	return s.OcrVideoResultsUrl
}

func (s *RecognizeVideoCastCrewListResponseBodyData) GetSubtitlesResults() []*RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults {
	return s.SubtitlesResults
}

func (s *RecognizeVideoCastCrewListResponseBodyData) GetVideoOcrResults() []*RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults {
	return s.VideoOcrResults
}

func (s *RecognizeVideoCastCrewListResponseBodyData) SetCastResults(v []*RecognizeVideoCastCrewListResponseBodyDataCastResults) *RecognizeVideoCastCrewListResponseBodyData {
	s.CastResults = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyData) SetOcrResults(v []*RecognizeVideoCastCrewListResponseBodyDataOcrResults) *RecognizeVideoCastCrewListResponseBodyData {
	s.OcrResults = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyData) SetOcrResultsUrl(v string) *RecognizeVideoCastCrewListResponseBodyData {
	s.OcrResultsUrl = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyData) SetOcrVideoResultsUrl(v string) *RecognizeVideoCastCrewListResponseBodyData {
	s.OcrVideoResultsUrl = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyData) SetSubtitlesResults(v []*RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) *RecognizeVideoCastCrewListResponseBodyData {
	s.SubtitlesResults = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyData) SetVideoOcrResults(v []*RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults) *RecognizeVideoCastCrewListResponseBodyData {
	s.VideoOcrResults = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyData) Validate() error {
	if s.CastResults != nil {
		for _, item := range s.CastResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OcrResults != nil {
		for _, item := range s.OcrResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SubtitlesResults != nil {
		for _, item := range s.SubtitlesResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoOcrResults != nil {
		for _, item := range s.VideoOcrResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RecognizeVideoCastCrewListResponseBodyDataCastResults struct {
	DetailInfo map[string]*string `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty"`
	// example:
	//
	// 0.6
	EndTime *float32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 0.6
	StartTime *float32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataCastResults) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataCastResults) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBodyDataCastResults) GetDetailInfo() map[string]*string {
	return s.DetailInfo
}

func (s *RecognizeVideoCastCrewListResponseBodyDataCastResults) GetEndTime() *float32 {
	return s.EndTime
}

func (s *RecognizeVideoCastCrewListResponseBodyDataCastResults) GetStartTime() *float32 {
	return s.StartTime
}

func (s *RecognizeVideoCastCrewListResponseBodyDataCastResults) SetDetailInfo(v map[string]*string) *RecognizeVideoCastCrewListResponseBodyDataCastResults {
	s.DetailInfo = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataCastResults) SetEndTime(v float32) *RecognizeVideoCastCrewListResponseBodyDataCastResults {
	s.EndTime = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataCastResults) SetStartTime(v float32) *RecognizeVideoCastCrewListResponseBodyDataCastResults {
	s.StartTime = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataCastResults) Validate() error {
	return dara.Validate(s)
}

type RecognizeVideoCastCrewListResponseBodyDataOcrResults struct {
	DetailInfo []*RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Repeated"`
	// example:
	//
	// 0.28
	EndTime *float32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 0.28
	StartTime *float32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataOcrResults) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataOcrResults) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResults) GetDetailInfo() []*RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo {
	return s.DetailInfo
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResults) GetEndTime() *float32 {
	return s.EndTime
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResults) GetStartTime() *float32 {
	return s.StartTime
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResults) SetDetailInfo(v []*RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) *RecognizeVideoCastCrewListResponseBodyDataOcrResults {
	s.DetailInfo = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResults) SetEndTime(v float32) *RecognizeVideoCastCrewListResponseBodyDataOcrResults {
	s.EndTime = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResults) SetStartTime(v float32) *RecognizeVideoCastCrewListResponseBodyDataOcrResults {
	s.StartTime = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResults) Validate() error {
	if s.DetailInfo != nil {
		for _, item := range s.DetailInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo struct {
	Boxes     []*int32     `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	CharProbs [][]*float32 `json:"CharProbs,omitempty" xml:"CharProbs,omitempty" type:"Repeated"`
	// example:
	//
	// 17
	FrameIndex *int64                                                                    `json:"FrameIndex,omitempty" xml:"FrameIndex,omitempty"`
	Position   []*RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition `json:"Position,omitempty" xml:"Position,omitempty" type:"Repeated"`
	// example:
	//
	// 92.07685702563117
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Text  *string  `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// 0.9207685702563116
	TextProb *float32 `json:"TextProb,omitempty" xml:"TextProb,omitempty"`
	// example:
	//
	// 0.28
	TimeStamp *float32 `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// example:
	//
	// 1
	TrackId *int64 `json:"TrackId,omitempty" xml:"TrackId,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) GetBoxes() []*int32 {
	return s.Boxes
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) GetCharProbs() [][]*float32 {
	return s.CharProbs
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) GetFrameIndex() *int64 {
	return s.FrameIndex
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) GetPosition() []*RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition {
	return s.Position
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) GetScore() *float32 {
	return s.Score
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) GetText() *string {
	return s.Text
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) GetTextProb() *float32 {
	return s.TextProb
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) GetTimeStamp() *float32 {
	return s.TimeStamp
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) GetTrackId() *int64 {
	return s.TrackId
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) SetBoxes(v []*int32) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo {
	s.Boxes = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) SetCharProbs(v [][]*float32) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo {
	s.CharProbs = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) SetFrameIndex(v int64) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo {
	s.FrameIndex = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) SetPosition(v []*RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo {
	s.Position = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) SetScore(v float32) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo {
	s.Score = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) SetText(v string) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo {
	s.Text = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) SetTextProb(v float32) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo {
	s.TextProb = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) SetTimeStamp(v float32) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo {
	s.TimeStamp = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) SetTrackId(v int64) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo {
	s.TrackId = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) Validate() error {
	if s.Position != nil {
		for _, item := range s.Position {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition struct {
	// example:
	//
	// 266
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 440
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition) GetX() *int64 {
	return s.X
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition) GetY() *int64 {
	return s.Y
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition) SetX(v int64) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition {
	s.X = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition) SetY(v int64) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition {
	s.Y = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition) Validate() error {
	return dara.Validate(s)
}

type RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults struct {
	SubtitlesAllResults map[string]*string `json:"SubtitlesAllResults,omitempty" xml:"SubtitlesAllResults,omitempty"`
	// example:
	//
	// url
	SubtitlesAllResultsUrl  *string            `json:"SubtitlesAllResultsUrl,omitempty" xml:"SubtitlesAllResultsUrl,omitempty"`
	SubtitlesChineseResults map[string]*string `json:"SubtitlesChineseResults,omitempty" xml:"SubtitlesChineseResults,omitempty"`
	// example:
	//
	// url1
	SubtitlesChineseResultsUrl *string `json:"SubtitlesChineseResultsUrl,omitempty" xml:"SubtitlesChineseResultsUrl,omitempty"`
	// example:
	//
	// hello
	SubtitlesEnglishResults map[string]interface{} `json:"SubtitlesEnglishResults,omitempty" xml:"SubtitlesEnglishResults,omitempty"`
	// example:
	//
	// url2
	SubtitlesEnglishResultsUrl *string `json:"SubtitlesEnglishResultsUrl,omitempty" xml:"SubtitlesEnglishResultsUrl,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) GetSubtitlesAllResults() map[string]*string {
	return s.SubtitlesAllResults
}

func (s *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) GetSubtitlesAllResultsUrl() *string {
	return s.SubtitlesAllResultsUrl
}

func (s *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) GetSubtitlesChineseResults() map[string]*string {
	return s.SubtitlesChineseResults
}

func (s *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) GetSubtitlesChineseResultsUrl() *string {
	return s.SubtitlesChineseResultsUrl
}

func (s *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) GetSubtitlesEnglishResults() map[string]interface{} {
	return s.SubtitlesEnglishResults
}

func (s *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) GetSubtitlesEnglishResultsUrl() *string {
	return s.SubtitlesEnglishResultsUrl
}

func (s *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) SetSubtitlesAllResults(v map[string]*string) *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults {
	s.SubtitlesAllResults = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) SetSubtitlesAllResultsUrl(v string) *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults {
	s.SubtitlesAllResultsUrl = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) SetSubtitlesChineseResults(v map[string]*string) *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults {
	s.SubtitlesChineseResults = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) SetSubtitlesChineseResultsUrl(v string) *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults {
	s.SubtitlesChineseResultsUrl = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) SetSubtitlesEnglishResults(v map[string]interface{}) *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults {
	s.SubtitlesEnglishResults = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) SetSubtitlesEnglishResultsUrl(v string) *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults {
	s.SubtitlesEnglishResultsUrl = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) Validate() error {
	return dara.Validate(s)
}

type RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults struct {
	DetailInfo []*RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Repeated"`
	// example:
	//
	// 0.92
	EndTime *float32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 0.92
	StartTime *float32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults) GetDetailInfo() []*RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo {
	return s.DetailInfo
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults) GetEndTime() *float32 {
	return s.EndTime
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults) GetStartTime() *float32 {
	return s.StartTime
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults) SetDetailInfo(v []*RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults {
	s.DetailInfo = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults) SetEndTime(v float32) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults {
	s.EndTime = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults) SetStartTime(v float32) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults {
	s.StartTime = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults) Validate() error {
	if s.DetailInfo != nil {
		for _, item := range s.DetailInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo struct {
	Boxes    []*int64                                                                       `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Position []*RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition `json:"Position,omitempty" xml:"Position,omitempty" type:"Repeated"`
	// example:
	//
	// 92.07685702563117
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Text  *string  `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// 0
	TextType *int64 `json:"TextType,omitempty" xml:"TextType,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) GetBoxes() []*int64 {
	return s.Boxes
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) GetPosition() []*RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition {
	return s.Position
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) GetScore() *float32 {
	return s.Score
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) GetText() *string {
	return s.Text
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) GetTextType() *int64 {
	return s.TextType
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) SetBoxes(v []*int64) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo {
	s.Boxes = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) SetPosition(v []*RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo {
	s.Position = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) SetScore(v float32) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo {
	s.Score = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) SetText(v string) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo {
	s.Text = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) SetTextType(v int64) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo {
	s.TextType = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) Validate() error {
	if s.Position != nil {
		for _, item := range s.Position {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition struct {
	// example:
	//
	// 269
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 423
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition) GetX() *int64 {
	return s.X
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition) GetY() *int64 {
	return s.Y
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition) SetX(v int64) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition {
	s.X = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition) SetY(v int64) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition {
	s.Y = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition) Validate() error {
	return dara.Validate(s)
}
