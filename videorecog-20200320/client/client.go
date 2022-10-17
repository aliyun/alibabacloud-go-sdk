// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/v2/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type DetectVideoShotRequest struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s DetectVideoShotRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoShotRequest) GoString() string {
	return s.String()
}

func (s *DetectVideoShotRequest) SetVideoUrl(v string) *DetectVideoShotRequest {
	s.VideoUrl = &v
	return s
}

type DetectVideoShotAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s DetectVideoShotAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoShotAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectVideoShotAdvanceRequest) SetVideoUrlObject(v io.Reader) *DetectVideoShotAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

type DetectVideoShotResponseBody struct {
	Data      *DetectVideoShotResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectVideoShotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoShotResponseBody) GoString() string {
	return s.String()
}

func (s *DetectVideoShotResponseBody) SetData(v *DetectVideoShotResponseBodyData) *DetectVideoShotResponseBody {
	s.Data = v
	return s
}

func (s *DetectVideoShotResponseBody) SetMessage(v string) *DetectVideoShotResponseBody {
	s.Message = &v
	return s
}

func (s *DetectVideoShotResponseBody) SetRequestId(v string) *DetectVideoShotResponseBody {
	s.RequestId = &v
	return s
}

type DetectVideoShotResponseBodyData struct {
	ShotFrameIds []*int32 `json:"ShotFrameIds,omitempty" xml:"ShotFrameIds,omitempty" type:"Repeated"`
}

func (s DetectVideoShotResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoShotResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectVideoShotResponseBodyData) SetShotFrameIds(v []*int32) *DetectVideoShotResponseBodyData {
	s.ShotFrameIds = v
	return s
}

type DetectVideoShotResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectVideoShotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectVideoShotResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoShotResponse) GoString() string {
	return s.String()
}

func (s *DetectVideoShotResponse) SetHeaders(v map[string]*string) *DetectVideoShotResponse {
	s.Headers = v
	return s
}

func (s *DetectVideoShotResponse) SetStatusCode(v int32) *DetectVideoShotResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectVideoShotResponse) SetBody(v *DetectVideoShotResponseBody) *DetectVideoShotResponse {
	s.Body = v
	return s
}

type GenerateVideoCoverRequest struct {
	IsGif    *bool   `json:"IsGif,omitempty" xml:"IsGif,omitempty"`
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GenerateVideoCoverRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoCoverRequest) GoString() string {
	return s.String()
}

func (s *GenerateVideoCoverRequest) SetIsGif(v bool) *GenerateVideoCoverRequest {
	s.IsGif = &v
	return s
}

func (s *GenerateVideoCoverRequest) SetVideoUrl(v string) *GenerateVideoCoverRequest {
	s.VideoUrl = &v
	return s
}

type GenerateVideoCoverAdvanceRequest struct {
	IsGif          *bool     `json:"IsGif,omitempty" xml:"IsGif,omitempty"`
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GenerateVideoCoverAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoCoverAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateVideoCoverAdvanceRequest) SetIsGif(v bool) *GenerateVideoCoverAdvanceRequest {
	s.IsGif = &v
	return s
}

func (s *GenerateVideoCoverAdvanceRequest) SetVideoUrlObject(v io.Reader) *GenerateVideoCoverAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

type GenerateVideoCoverResponseBody struct {
	Data      *GenerateVideoCoverResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateVideoCoverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoCoverResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateVideoCoverResponseBody) SetData(v *GenerateVideoCoverResponseBodyData) *GenerateVideoCoverResponseBody {
	s.Data = v
	return s
}

func (s *GenerateVideoCoverResponseBody) SetMessage(v string) *GenerateVideoCoverResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateVideoCoverResponseBody) SetRequestId(v string) *GenerateVideoCoverResponseBody {
	s.RequestId = &v
	return s
}

type GenerateVideoCoverResponseBodyData struct {
	Outputs []*GenerateVideoCoverResponseBodyDataOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
}

func (s GenerateVideoCoverResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoCoverResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateVideoCoverResponseBodyData) SetOutputs(v []*GenerateVideoCoverResponseBodyDataOutputs) *GenerateVideoCoverResponseBodyData {
	s.Outputs = v
	return s
}

type GenerateVideoCoverResponseBodyDataOutputs struct {
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	ImageURL   *string  `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s GenerateVideoCoverResponseBodyDataOutputs) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoCoverResponseBodyDataOutputs) GoString() string {
	return s.String()
}

func (s *GenerateVideoCoverResponseBodyDataOutputs) SetConfidence(v float32) *GenerateVideoCoverResponseBodyDataOutputs {
	s.Confidence = &v
	return s
}

func (s *GenerateVideoCoverResponseBodyDataOutputs) SetImageURL(v string) *GenerateVideoCoverResponseBodyDataOutputs {
	s.ImageURL = &v
	return s
}

type GenerateVideoCoverResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateVideoCoverResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateVideoCoverResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoCoverResponse) GoString() string {
	return s.String()
}

func (s *GenerateVideoCoverResponse) SetHeaders(v map[string]*string) *GenerateVideoCoverResponse {
	s.Headers = v
	return s
}

func (s *GenerateVideoCoverResponse) SetStatusCode(v int32) *GenerateVideoCoverResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateVideoCoverResponse) SetBody(v *GenerateVideoCoverResponseBody) *GenerateVideoCoverResponse {
	s.Body = v
	return s
}

type GetAsyncJobResultRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAsyncJobResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultRequest) SetJobId(v string) *GetAsyncJobResultRequest {
	s.JobId = &v
	return s
}

type GetAsyncJobResultResponseBody struct {
	Data      *GetAsyncJobResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAsyncJobResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBody) SetData(v *GetAsyncJobResultResponseBodyData) *GetAsyncJobResultResponseBody {
	s.Data = v
	return s
}

func (s *GetAsyncJobResultResponseBody) SetRequestId(v string) *GetAsyncJobResultResponseBody {
	s.RequestId = &v
	return s
}

type GetAsyncJobResultResponseBodyData struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAsyncJobResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorCode(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorMessage(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetJobId(v string) *GetAsyncJobResultResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetResult(v string) *GetAsyncJobResultResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetStatus(v string) *GetAsyncJobResultResponseBodyData {
	s.Status = &v
	return s
}

type GetAsyncJobResultResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAsyncJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAsyncJobResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponse) SetHeaders(v map[string]*string) *GetAsyncJobResultResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncJobResultResponse) SetStatusCode(v int32) *GetAsyncJobResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncJobResultResponse) SetBody(v *GetAsyncJobResultResponseBody) *GetAsyncJobResultResponse {
	s.Body = v
	return s
}

type RecognizeVideoCastCrewListRequest struct {
	Params      []*RecognizeVideoCastCrewListRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	RegisterUrl *string                                    `json:"RegisterUrl,omitempty" xml:"RegisterUrl,omitempty"`
	VideoUrl    *string                                    `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s RecognizeVideoCastCrewListRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListRequest) SetParams(v []*RecognizeVideoCastCrewListRequestParams) *RecognizeVideoCastCrewListRequest {
	s.Params = v
	return s
}

func (s *RecognizeVideoCastCrewListRequest) SetRegisterUrl(v string) *RecognizeVideoCastCrewListRequest {
	s.RegisterUrl = &v
	return s
}

func (s *RecognizeVideoCastCrewListRequest) SetVideoUrl(v string) *RecognizeVideoCastCrewListRequest {
	s.VideoUrl = &v
	return s
}

type RecognizeVideoCastCrewListRequestParams struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RecognizeVideoCastCrewListRequestParams) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListRequestParams) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListRequestParams) SetType(v string) *RecognizeVideoCastCrewListRequestParams {
	s.Type = &v
	return s
}

type RecognizeVideoCastCrewListAdvanceRequest struct {
	Params            []*RecognizeVideoCastCrewListAdvanceRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	RegisterUrlObject io.Reader                                         `json:"RegisterUrl,omitempty" xml:"RegisterUrl,omitempty"`
	VideoUrlObject    io.Reader                                         `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s RecognizeVideoCastCrewListAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListAdvanceRequest) SetParams(v []*RecognizeVideoCastCrewListAdvanceRequestParams) *RecognizeVideoCastCrewListAdvanceRequest {
	s.Params = v
	return s
}

func (s *RecognizeVideoCastCrewListAdvanceRequest) SetRegisterUrlObject(v io.Reader) *RecognizeVideoCastCrewListAdvanceRequest {
	s.RegisterUrlObject = v
	return s
}

func (s *RecognizeVideoCastCrewListAdvanceRequest) SetVideoUrlObject(v io.Reader) *RecognizeVideoCastCrewListAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

type RecognizeVideoCastCrewListAdvanceRequestParams struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RecognizeVideoCastCrewListAdvanceRequestParams) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListAdvanceRequestParams) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListAdvanceRequestParams) SetType(v string) *RecognizeVideoCastCrewListAdvanceRequestParams {
	s.Type = &v
	return s
}

type RecognizeVideoCastCrewListShrinkRequest struct {
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
	RegisterUrl  *string `json:"RegisterUrl,omitempty" xml:"RegisterUrl,omitempty"`
	VideoUrl     *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s RecognizeVideoCastCrewListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListShrinkRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListShrinkRequest) SetParamsShrink(v string) *RecognizeVideoCastCrewListShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *RecognizeVideoCastCrewListShrinkRequest) SetRegisterUrl(v string) *RecognizeVideoCastCrewListShrinkRequest {
	s.RegisterUrl = &v
	return s
}

func (s *RecognizeVideoCastCrewListShrinkRequest) SetVideoUrl(v string) *RecognizeVideoCastCrewListShrinkRequest {
	s.VideoUrl = &v
	return s
}

type RecognizeVideoCastCrewListResponseBody struct {
	Data      *RecognizeVideoCastCrewListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBody) GoString() string {
	return s.String()
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

type RecognizeVideoCastCrewListResponseBodyData struct {
	CastResults        []*RecognizeVideoCastCrewListResponseBodyDataCastResults      `json:"CastResults,omitempty" xml:"CastResults,omitempty" type:"Repeated"`
	OcrResults         []*RecognizeVideoCastCrewListResponseBodyDataOcrResults       `json:"OcrResults,omitempty" xml:"OcrResults,omitempty" type:"Repeated"`
	OcrResultsUrl      *string                                                       `json:"OcrResultsUrl,omitempty" xml:"OcrResultsUrl,omitempty"`
	OcrVideoResultsUrl *string                                                       `json:"OcrVideoResultsUrl,omitempty" xml:"OcrVideoResultsUrl,omitempty"`
	SubtitlesResults   []*RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults `json:"SubtitlesResults,omitempty" xml:"SubtitlesResults,omitempty" type:"Repeated"`
	VideoOcrResults    []*RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults  `json:"VideoOcrResults,omitempty" xml:"VideoOcrResults,omitempty" type:"Repeated"`
}

func (s RecognizeVideoCastCrewListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyData) GoString() string {
	return s.String()
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

type RecognizeVideoCastCrewListResponseBodyDataCastResults struct {
	DetailInfo map[string]*string `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty"`
	EndTime    *float32           `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *float32           `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataCastResults) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataCastResults) GoString() string {
	return s.String()
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

type RecognizeVideoCastCrewListResponseBodyDataOcrResults struct {
	DetailInfo []*RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Repeated"`
	EndTime    *float32                                                          `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *float32                                                          `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataOcrResults) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataOcrResults) GoString() string {
	return s.String()
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

type RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo struct {
	Boxes      []*int32                                                                  `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	CharProbs  [][]*float32                                                              `json:"CharProbs,omitempty" xml:"CharProbs,omitempty" type:"Repeated"`
	FrameIndex *int64                                                                    `json:"FrameIndex,omitempty" xml:"FrameIndex,omitempty"`
	Position   []*RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition `json:"Position,omitempty" xml:"Position,omitempty" type:"Repeated"`
	Score      *float32                                                                  `json:"Score,omitempty" xml:"Score,omitempty"`
	Text       *string                                                                   `json:"Text,omitempty" xml:"Text,omitempty"`
	TextProb   *float32                                                                  `json:"TextProb,omitempty" xml:"TextProb,omitempty"`
	TimeStamp  *float32                                                                  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	TrackId    *int64                                                                    `json:"TrackId,omitempty" xml:"TrackId,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) GoString() string {
	return s.String()
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

type RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition struct {
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition) SetX(v int64) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition {
	s.X = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition) SetY(v int64) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition {
	s.Y = &v
	return s
}

type RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults struct {
	SubtitlesAllResults        map[string]*string     `json:"SubtitlesAllResults,omitempty" xml:"SubtitlesAllResults,omitempty"`
	SubtitlesAllResultsUrl     *string                `json:"SubtitlesAllResultsUrl,omitempty" xml:"SubtitlesAllResultsUrl,omitempty"`
	SubtitlesChineseResults    map[string]*string     `json:"SubtitlesChineseResults,omitempty" xml:"SubtitlesChineseResults,omitempty"`
	SubtitlesChineseResultsUrl *string                `json:"SubtitlesChineseResultsUrl,omitempty" xml:"SubtitlesChineseResultsUrl,omitempty"`
	SubtitlesEnglishResults    map[string]interface{} `json:"SubtitlesEnglishResults,omitempty" xml:"SubtitlesEnglishResults,omitempty"`
	SubtitlesEnglishResultsUrl *string                `json:"SubtitlesEnglishResultsUrl,omitempty" xml:"SubtitlesEnglishResultsUrl,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) GoString() string {
	return s.String()
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

type RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults struct {
	DetailInfo []*RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Repeated"`
	EndTime    *float32                                                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *float32                                                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults) GoString() string {
	return s.String()
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

type RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo struct {
	Boxes    []*int64                                                                       `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Position []*RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition `json:"Position,omitempty" xml:"Position,omitempty" type:"Repeated"`
	Score    *float32                                                                       `json:"Score,omitempty" xml:"Score,omitempty"`
	Text     *string                                                                        `json:"Text,omitempty" xml:"Text,omitempty"`
	TextType *int64                                                                         `json:"TextType,omitempty" xml:"TextType,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) GoString() string {
	return s.String()
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

type RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition struct {
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition) SetX(v int64) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition {
	s.X = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition) SetY(v int64) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition {
	s.Y = &v
	return s
}

type RecognizeVideoCastCrewListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeVideoCastCrewListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeVideoCastCrewListResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponse) SetHeaders(v map[string]*string) *RecognizeVideoCastCrewListResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVideoCastCrewListResponse) SetStatusCode(v int32) *RecognizeVideoCastCrewListResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponse) SetBody(v *RecognizeVideoCastCrewListResponseBody) *RecognizeVideoCastCrewListResponse {
	s.Body = v
	return s
}

type SplitVideoPartsRequest struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s SplitVideoPartsRequest) String() string {
	return tea.Prettify(s)
}

func (s SplitVideoPartsRequest) GoString() string {
	return s.String()
}

func (s *SplitVideoPartsRequest) SetVideoUrl(v string) *SplitVideoPartsRequest {
	s.VideoUrl = &v
	return s
}

type SplitVideoPartsAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s SplitVideoPartsAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SplitVideoPartsAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SplitVideoPartsAdvanceRequest) SetVideoUrlObject(v io.Reader) *SplitVideoPartsAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

type SplitVideoPartsResponseBody struct {
	Data      *SplitVideoPartsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SplitVideoPartsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SplitVideoPartsResponseBody) GoString() string {
	return s.String()
}

func (s *SplitVideoPartsResponseBody) SetData(v *SplitVideoPartsResponseBodyData) *SplitVideoPartsResponseBody {
	s.Data = v
	return s
}

func (s *SplitVideoPartsResponseBody) SetMessage(v string) *SplitVideoPartsResponseBody {
	s.Message = &v
	return s
}

func (s *SplitVideoPartsResponseBody) SetRequestId(v string) *SplitVideoPartsResponseBody {
	s.RequestId = &v
	return s
}

type SplitVideoPartsResponseBodyData struct {
	Elements []*SplitVideoPartsResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s SplitVideoPartsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SplitVideoPartsResponseBodyData) GoString() string {
	return s.String()
}

func (s *SplitVideoPartsResponseBodyData) SetElements(v []*SplitVideoPartsResponseBodyDataElements) *SplitVideoPartsResponseBodyData {
	s.Elements = v
	return s
}

type SplitVideoPartsResponseBodyDataElements struct {
	BeginTime *float32 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime   *float32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Index     *int64   `json:"Index,omitempty" xml:"Index,omitempty"`
}

func (s SplitVideoPartsResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s SplitVideoPartsResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *SplitVideoPartsResponseBodyDataElements) SetBeginTime(v float32) *SplitVideoPartsResponseBodyDataElements {
	s.BeginTime = &v
	return s
}

func (s *SplitVideoPartsResponseBodyDataElements) SetEndTime(v float32) *SplitVideoPartsResponseBodyDataElements {
	s.EndTime = &v
	return s
}

func (s *SplitVideoPartsResponseBodyDataElements) SetIndex(v int64) *SplitVideoPartsResponseBodyDataElements {
	s.Index = &v
	return s
}

type SplitVideoPartsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SplitVideoPartsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SplitVideoPartsResponse) String() string {
	return tea.Prettify(s)
}

func (s SplitVideoPartsResponse) GoString() string {
	return s.String()
}

func (s *SplitVideoPartsResponse) SetHeaders(v map[string]*string) *SplitVideoPartsResponse {
	s.Headers = v
	return s
}

func (s *SplitVideoPartsResponse) SetStatusCode(v int32) *SplitVideoPartsResponse {
	s.StatusCode = &v
	return s
}

func (s *SplitVideoPartsResponse) SetBody(v *SplitVideoPartsResponseBody) *SplitVideoPartsResponse {
	s.Body = v
	return s
}

type UnderstandVideoContentRequest struct {
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s UnderstandVideoContentRequest) String() string {
	return tea.Prettify(s)
}

func (s UnderstandVideoContentRequest) GoString() string {
	return s.String()
}

func (s *UnderstandVideoContentRequest) SetVideoURL(v string) *UnderstandVideoContentRequest {
	s.VideoURL = &v
	return s
}

type UnderstandVideoContentAdvanceRequest struct {
	VideoURLObject io.Reader `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s UnderstandVideoContentAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnderstandVideoContentAdvanceRequest) GoString() string {
	return s.String()
}

func (s *UnderstandVideoContentAdvanceRequest) SetVideoURLObject(v io.Reader) *UnderstandVideoContentAdvanceRequest {
	s.VideoURLObject = v
	return s
}

type UnderstandVideoContentResponseBody struct {
	Data      *UnderstandVideoContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnderstandVideoContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnderstandVideoContentResponseBody) GoString() string {
	return s.String()
}

func (s *UnderstandVideoContentResponseBody) SetData(v *UnderstandVideoContentResponseBodyData) *UnderstandVideoContentResponseBody {
	s.Data = v
	return s
}

func (s *UnderstandVideoContentResponseBody) SetMessage(v string) *UnderstandVideoContentResponseBody {
	s.Message = &v
	return s
}

func (s *UnderstandVideoContentResponseBody) SetRequestId(v string) *UnderstandVideoContentResponseBody {
	s.RequestId = &v
	return s
}

type UnderstandVideoContentResponseBodyData struct {
	TagInfo   map[string]interface{}                           `json:"TagInfo,omitempty" xml:"TagInfo,omitempty"`
	VideoInfo *UnderstandVideoContentResponseBodyDataVideoInfo `json:"VideoInfo,omitempty" xml:"VideoInfo,omitempty" type:"Struct"`
}

func (s UnderstandVideoContentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UnderstandVideoContentResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnderstandVideoContentResponseBodyData) SetTagInfo(v map[string]interface{}) *UnderstandVideoContentResponseBodyData {
	s.TagInfo = v
	return s
}

func (s *UnderstandVideoContentResponseBodyData) SetVideoInfo(v *UnderstandVideoContentResponseBodyDataVideoInfo) *UnderstandVideoContentResponseBodyData {
	s.VideoInfo = v
	return s
}

type UnderstandVideoContentResponseBodyDataVideoInfo struct {
	Duration *int64   `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Fps      *float32 `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Height   *int64   `json:"Height,omitempty" xml:"Height,omitempty"`
	Width    *int64   `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s UnderstandVideoContentResponseBodyDataVideoInfo) String() string {
	return tea.Prettify(s)
}

func (s UnderstandVideoContentResponseBodyDataVideoInfo) GoString() string {
	return s.String()
}

func (s *UnderstandVideoContentResponseBodyDataVideoInfo) SetDuration(v int64) *UnderstandVideoContentResponseBodyDataVideoInfo {
	s.Duration = &v
	return s
}

func (s *UnderstandVideoContentResponseBodyDataVideoInfo) SetFps(v float32) *UnderstandVideoContentResponseBodyDataVideoInfo {
	s.Fps = &v
	return s
}

func (s *UnderstandVideoContentResponseBodyDataVideoInfo) SetHeight(v int64) *UnderstandVideoContentResponseBodyDataVideoInfo {
	s.Height = &v
	return s
}

func (s *UnderstandVideoContentResponseBodyDataVideoInfo) SetWidth(v int64) *UnderstandVideoContentResponseBodyDataVideoInfo {
	s.Width = &v
	return s
}

type UnderstandVideoContentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnderstandVideoContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnderstandVideoContentResponse) String() string {
	return tea.Prettify(s)
}

func (s UnderstandVideoContentResponse) GoString() string {
	return s.String()
}

func (s *UnderstandVideoContentResponse) SetHeaders(v map[string]*string) *UnderstandVideoContentResponse {
	s.Headers = v
	return s
}

func (s *UnderstandVideoContentResponse) SetStatusCode(v int32) *UnderstandVideoContentResponse {
	s.StatusCode = &v
	return s
}

func (s *UnderstandVideoContentResponse) SetBody(v *UnderstandVideoContentResponseBody) *UnderstandVideoContentResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("videorecog"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectVideoShotWithOptions(request *DetectVideoShotRequest, runtime *util.RuntimeOptions) (_result *DetectVideoShotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectVideoShot"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectVideoShotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectVideoShot(request *DetectVideoShotRequest) (_result *DetectVideoShotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectVideoShotResponse{}
	_body, _err := client.DetectVideoShotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectVideoShotAdvance(request *DetectVideoShotAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectVideoShotResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("videorecog"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectVideoShotReq := &DetectVideoShotRequest{}
	openapiutil.Convert(request, detectVideoShotReq)
	if !tea.BoolValue(util.IsUnset(request.VideoUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.VideoUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectVideoShotReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectVideoShotResp, _err := client.DetectVideoShotWithOptions(detectVideoShotReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectVideoShotResp
	return _result, _err
}

func (client *Client) GenerateVideoCoverWithOptions(request *GenerateVideoCoverRequest, runtime *util.RuntimeOptions) (_result *GenerateVideoCoverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsGif)) {
		body["IsGif"] = request.IsGif
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateVideoCover"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateVideoCoverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateVideoCover(request *GenerateVideoCoverRequest) (_result *GenerateVideoCoverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateVideoCoverResponse{}
	_body, _err := client.GenerateVideoCoverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateVideoCoverAdvance(request *GenerateVideoCoverAdvanceRequest, runtime *util.RuntimeOptions) (_result *GenerateVideoCoverResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("videorecog"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	generateVideoCoverReq := &GenerateVideoCoverRequest{}
	openapiutil.Convert(request, generateVideoCoverReq)
	if !tea.BoolValue(util.IsUnset(request.VideoUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.VideoUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		generateVideoCoverReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	generateVideoCoverResp, _err := client.GenerateVideoCoverWithOptions(generateVideoCoverReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = generateVideoCoverResp
	return _result, _err
}

func (client *Client) GetAsyncJobResultWithOptions(request *GetAsyncJobResultRequest, runtime *util.RuntimeOptions) (_result *GetAsyncJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsyncJobResult"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncJobResult(request *GetAsyncJobResultRequest) (_result *GetAsyncJobResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.GetAsyncJobResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeVideoCastCrewListWithOptions(tmpReq *RecognizeVideoCastCrewListRequest, runtime *util.RuntimeOptions) (_result *RecognizeVideoCastCrewListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RecognizeVideoCastCrewListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Params)) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, tea.String("Params"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParamsShrink)) {
		body["Params"] = request.ParamsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegisterUrl)) {
		body["RegisterUrl"] = request.RegisterUrl
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeVideoCastCrewList"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeVideoCastCrewListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeVideoCastCrewList(request *RecognizeVideoCastCrewListRequest) (_result *RecognizeVideoCastCrewListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeVideoCastCrewListResponse{}
	_body, _err := client.RecognizeVideoCastCrewListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeVideoCastCrewListAdvance(request *RecognizeVideoCastCrewListAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeVideoCastCrewListResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("videorecog"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeVideoCastCrewListReq := &RecognizeVideoCastCrewListRequest{}
	openapiutil.Convert(request, recognizeVideoCastCrewListReq)
	if !tea.BoolValue(util.IsUnset(request.RegisterUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.RegisterUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeVideoCastCrewListReq.RegisterUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.VideoUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeVideoCastCrewListReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeVideoCastCrewListResp, _err := client.RecognizeVideoCastCrewListWithOptions(recognizeVideoCastCrewListReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeVideoCastCrewListResp
	return _result, _err
}

func (client *Client) SplitVideoPartsWithOptions(request *SplitVideoPartsRequest, runtime *util.RuntimeOptions) (_result *SplitVideoPartsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SplitVideoParts"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SplitVideoPartsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SplitVideoParts(request *SplitVideoPartsRequest) (_result *SplitVideoPartsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SplitVideoPartsResponse{}
	_body, _err := client.SplitVideoPartsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SplitVideoPartsAdvance(request *SplitVideoPartsAdvanceRequest, runtime *util.RuntimeOptions) (_result *SplitVideoPartsResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("videorecog"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	splitVideoPartsReq := &SplitVideoPartsRequest{}
	openapiutil.Convert(request, splitVideoPartsReq)
	if !tea.BoolValue(util.IsUnset(request.VideoUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.VideoUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		splitVideoPartsReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	splitVideoPartsResp, _err := client.SplitVideoPartsWithOptions(splitVideoPartsReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = splitVideoPartsResp
	return _result, _err
}

func (client *Client) UnderstandVideoContentWithOptions(request *UnderstandVideoContentRequest, runtime *util.RuntimeOptions) (_result *UnderstandVideoContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VideoURL)) {
		body["VideoURL"] = request.VideoURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnderstandVideoContent"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnderstandVideoContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnderstandVideoContent(request *UnderstandVideoContentRequest) (_result *UnderstandVideoContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnderstandVideoContentResponse{}
	_body, _err := client.UnderstandVideoContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnderstandVideoContentAdvance(request *UnderstandVideoContentAdvanceRequest, runtime *util.RuntimeOptions) (_result *UnderstandVideoContentResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("videorecog"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	understandVideoContentReq := &UnderstandVideoContentRequest{}
	openapiutil.Convert(request, understandVideoContentReq)
	if !tea.BoolValue(util.IsUnset(request.VideoURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.VideoURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		understandVideoContentReq.VideoURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	understandVideoContentResp, _err := client.UnderstandVideoContentWithOptions(understandVideoContentReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = understandVideoContentResp
	return _result, _err
}
