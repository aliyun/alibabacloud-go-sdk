// This file is auto-generated, don't edit it. Thanks.
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
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videorecog/DetectVideoShot/DetectVideoShot2.mp4
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
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videorecog/DetectVideoShot/DetectVideoShot2.mp4
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
	Data    *DetectVideoShotResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0033B795-09C7-4EB9-A33C-EBA325192B0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// 1
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectVideoShotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type EvaluateVideoQualityRequest struct {
	// example:
	//
	// vqa_plus
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://public-vigen-video.oss-cn-shanghai.aliyuncs.com/Common/xxx/dont_delete/decaption/123.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EvaluateVideoQualityRequest) String() string {
	return tea.Prettify(s)
}

func (s EvaluateVideoQualityRequest) GoString() string {
	return s.String()
}

func (s *EvaluateVideoQualityRequest) SetMode(v string) *EvaluateVideoQualityRequest {
	s.Mode = &v
	return s
}

func (s *EvaluateVideoQualityRequest) SetVideoUrl(v string) *EvaluateVideoQualityRequest {
	s.VideoUrl = &v
	return s
}

type EvaluateVideoQualityAdvanceRequest struct {
	// example:
	//
	// vqa_plus
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://public-vigen-video.oss-cn-shanghai.aliyuncs.com/Common/xxx/dont_delete/decaption/123.mp4
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EvaluateVideoQualityAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s EvaluateVideoQualityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *EvaluateVideoQualityAdvanceRequest) SetMode(v string) *EvaluateVideoQualityAdvanceRequest {
	s.Mode = &v
	return s
}

func (s *EvaluateVideoQualityAdvanceRequest) SetVideoUrlObject(v io.Reader) *EvaluateVideoQualityAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

type EvaluateVideoQualityResponseBody struct {
	Data    *EvaluateVideoQualityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1d33e538-c949-4fcd-83f6-4d57e4b31527
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EvaluateVideoQualityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EvaluateVideoQualityResponseBody) GoString() string {
	return s.String()
}

func (s *EvaluateVideoQualityResponseBody) SetData(v *EvaluateVideoQualityResponseBodyData) *EvaluateVideoQualityResponseBody {
	s.Data = v
	return s
}

func (s *EvaluateVideoQualityResponseBody) SetMessage(v string) *EvaluateVideoQualityResponseBody {
	s.Message = &v
	return s
}

func (s *EvaluateVideoQualityResponseBody) SetRequestId(v string) *EvaluateVideoQualityResponseBody {
	s.RequestId = &v
	return s
}

type EvaluateVideoQualityResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-damo-eas-cn-shanghai.oss-cn-shanghai.aliyuncs.com/eas-video-quality-assessment/2023-01-13-10/31%3A08-cVeN9ZQlzIPfGqsa.json?Expires=1673578869&amp;OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&amp;Signature=AiSsOsZ7rYfhf9w3Mxn%2Fq4GKKy****
	JsonUrl *string `json:"JsonUrl,omitempty" xml:"JsonUrl,omitempty"`
	// example:
	//
	// http://vibktprfx-prod-prod-damo-eas-cn-shanghai.oss-cn-shanghai.aliyuncs.com/eas-video-quality-assessment/2023-01-13-10/31%3A08-cVeN9ZQlzIPfGqsa.pdf?Expires=1673578869&amp;OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&amp;Signature=xULlZzVuhoYWAXRbp9A4EzzZcS****
	PdfUrl           *string                                               `json:"PdfUrl,omitempty" xml:"PdfUrl,omitempty"`
	VideoQualityInfo *EvaluateVideoQualityResponseBodyDataVideoQualityInfo `json:"VideoQualityInfo,omitempty" xml:"VideoQualityInfo,omitempty" type:"Struct"`
}

func (s EvaluateVideoQualityResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EvaluateVideoQualityResponseBodyData) GoString() string {
	return s.String()
}

func (s *EvaluateVideoQualityResponseBodyData) SetJsonUrl(v string) *EvaluateVideoQualityResponseBodyData {
	s.JsonUrl = &v
	return s
}

func (s *EvaluateVideoQualityResponseBodyData) SetPdfUrl(v string) *EvaluateVideoQualityResponseBodyData {
	s.PdfUrl = &v
	return s
}

func (s *EvaluateVideoQualityResponseBodyData) SetVideoQualityInfo(v *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) *EvaluateVideoQualityResponseBodyData {
	s.VideoQualityInfo = v
	return s
}

type EvaluateVideoQualityResponseBodyDataVideoQualityInfo struct {
	// example:
	//
	// 0.15
	Blurriness *float32 `json:"Blurriness,omitempty" xml:"Blurriness,omitempty"`
	// example:
	//
	// 0.55
	ColorContrast *float32 `json:"ColorContrast,omitempty" xml:"ColorContrast,omitempty"`
	// example:
	//
	// 0.17
	ColorSaturation *float32 `json:"ColorSaturation,omitempty" xml:"ColorSaturation,omitempty"`
	// example:
	//
	// 0.48
	Colorfulness *float32 `json:"Colorfulness,omitempty" xml:"Colorfulness,omitempty"`
	// example:
	//
	// 0.25
	CompressiveStrength *float32 `json:"CompressiveStrength,omitempty" xml:"CompressiveStrength,omitempty"`
	// example:
	//
	// 0.51
	Luminance *float32 `json:"Luminance,omitempty" xml:"Luminance,omitempty"`
	// example:
	//
	// 0.7048
	MosScore *float32 `json:"MosScore,omitempty" xml:"MosScore,omitempty"`
	// example:
	//
	// 0.01
	NoiseIntensity *float32 `json:"NoiseIntensity,omitempty" xml:"NoiseIntensity,omitempty"`
}

func (s EvaluateVideoQualityResponseBodyDataVideoQualityInfo) String() string {
	return tea.Prettify(s)
}

func (s EvaluateVideoQualityResponseBodyDataVideoQualityInfo) GoString() string {
	return s.String()
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) SetBlurriness(v float32) *EvaluateVideoQualityResponseBodyDataVideoQualityInfo {
	s.Blurriness = &v
	return s
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) SetColorContrast(v float32) *EvaluateVideoQualityResponseBodyDataVideoQualityInfo {
	s.ColorContrast = &v
	return s
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) SetColorSaturation(v float32) *EvaluateVideoQualityResponseBodyDataVideoQualityInfo {
	s.ColorSaturation = &v
	return s
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) SetColorfulness(v float32) *EvaluateVideoQualityResponseBodyDataVideoQualityInfo {
	s.Colorfulness = &v
	return s
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) SetCompressiveStrength(v float32) *EvaluateVideoQualityResponseBodyDataVideoQualityInfo {
	s.CompressiveStrength = &v
	return s
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) SetLuminance(v float32) *EvaluateVideoQualityResponseBodyDataVideoQualityInfo {
	s.Luminance = &v
	return s
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) SetMosScore(v float32) *EvaluateVideoQualityResponseBodyDataVideoQualityInfo {
	s.MosScore = &v
	return s
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) SetNoiseIntensity(v float32) *EvaluateVideoQualityResponseBodyDataVideoQualityInfo {
	s.NoiseIntensity = &v
	return s
}

type EvaluateVideoQualityResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EvaluateVideoQualityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EvaluateVideoQualityResponse) String() string {
	return tea.Prettify(s)
}

func (s EvaluateVideoQualityResponse) GoString() string {
	return s.String()
}

func (s *EvaluateVideoQualityResponse) SetHeaders(v map[string]*string) *EvaluateVideoQualityResponse {
	s.Headers = v
	return s
}

func (s *EvaluateVideoQualityResponse) SetStatusCode(v int32) *EvaluateVideoQualityResponse {
	s.StatusCode = &v
	return s
}

func (s *EvaluateVideoQualityResponse) SetBody(v *EvaluateVideoQualityResponseBody) *EvaluateVideoQualityResponse {
	s.Body = v
	return s
}

type GenerateVideoCoverRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// false
	IsGif *bool `json:"IsGif,omitempty" xml:"IsGif,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videorecog/videorecog/videorecog1.mp4
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
	// This parameter is required.
	//
	// example:
	//
	// false
	IsGif *bool `json:"IsGif,omitempty" xml:"IsGif,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videorecog/videorecog/videorecog1.mp4
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
	Data    *GenerateVideoCoverResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5B95B724-C5B9-4F77-A743-0CA4EA95CC82
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 6.1819260887924425
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// example:
	//
	// http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/video-cover/2020-05-11-07/pic_lOyxGGAqQYSANGxP.mp4_202_544_960_c9f88b2a5f75e17c093d1a65f5edff4d_beautified.png?Expires=1589185385&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=PAalKsfeZC4UQzYDTU%2F3D1G7Xt****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateVideoCoverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// E75FE679-0303-4DD1-8252-1143B4FA8A27
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
	Data *GetAsyncJobResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// A1F44EC4-118D-4A03-B213-F908F36F7DAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// InvalidParameter
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// paramsIllegal
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 7DFDA846-178B-4ADB-B69A-62C641214D81
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// {\\"Outputs\\":[{\\"ImageURL\\":\\"http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/video-cover/2020-03-25-11/pic_4t7zW6R6SUGn4DLF.mp4_2375_1920_1080_96ce5a96b5b16628cd778c035b68356d_beautified.png?Expires=1585136160&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=MDd7BqbivlLRd16MTKbPFQHV3u****\\",\\"Confidence\\":8.426481079120514},{\\"ImageURL\\":\\"http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/video-cover/2020-03-25-11/pic_4t7zW6R6SUGn4DLF.mp4_2996_1920_1080_d5df0556bf420242c84fe6f7a45d01e1_beautified.png?Expires=1585136160&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=htaW5G%2BpqPBW%2BEMTe01ckVoGsQ****\\",\\"Confidence\\":6.225726566341124}]}
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// PROCESS_SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAsyncJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Params []*RecognizeVideoCastCrewListRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// https://shanghai.oss-cn-shanghai.aliyuncs.com/download/xxxx.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
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

func (s *RecognizeVideoCastCrewListRequest) SetVideoUrl(v string) *RecognizeVideoCastCrewListRequest {
	s.VideoUrl = &v
	return s
}

type RecognizeVideoCastCrewListRequestParams struct {
	// example:
	//
	// cast
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
	Params []*RecognizeVideoCastCrewListAdvanceRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// https://shanghai.oss-cn-shanghai.aliyuncs.com/download/xxxx.mp4
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
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

func (s *RecognizeVideoCastCrewListAdvanceRequest) SetVideoUrlObject(v io.Reader) *RecognizeVideoCastCrewListAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

type RecognizeVideoCastCrewListAdvanceRequestParams struct {
	// example:
	//
	// cast
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
	// This parameter is required.
	//
	// example:
	//
	// https://shanghai.oss-cn-shanghai.aliyuncs.com/download/xxxx.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
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

func (s *RecognizeVideoCastCrewListShrinkRequest) SetVideoUrl(v string) *RecognizeVideoCastCrewListShrinkRequest {
	s.VideoUrl = &v
	return s
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeVideoCastCrewListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	MaxTime  *int32  `json:"MaxTime,omitempty" xml:"MaxTime,omitempty"`
	MinTime  *int32  `json:"MinTime,omitempty" xml:"MinTime,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/ocr/xxxx.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s SplitVideoPartsRequest) String() string {
	return tea.Prettify(s)
}

func (s SplitVideoPartsRequest) GoString() string {
	return s.String()
}

func (s *SplitVideoPartsRequest) SetMaxTime(v int32) *SplitVideoPartsRequest {
	s.MaxTime = &v
	return s
}

func (s *SplitVideoPartsRequest) SetMinTime(v int32) *SplitVideoPartsRequest {
	s.MinTime = &v
	return s
}

func (s *SplitVideoPartsRequest) SetTemplate(v string) *SplitVideoPartsRequest {
	s.Template = &v
	return s
}

func (s *SplitVideoPartsRequest) SetVideoUrl(v string) *SplitVideoPartsRequest {
	s.VideoUrl = &v
	return s
}

type SplitVideoPartsAdvanceRequest struct {
	MaxTime  *int32  `json:"MaxTime,omitempty" xml:"MaxTime,omitempty"`
	MinTime  *int32  `json:"MinTime,omitempty" xml:"MinTime,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/ocr/xxxx.mp4
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s SplitVideoPartsAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SplitVideoPartsAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SplitVideoPartsAdvanceRequest) SetMaxTime(v int32) *SplitVideoPartsAdvanceRequest {
	s.MaxTime = &v
	return s
}

func (s *SplitVideoPartsAdvanceRequest) SetMinTime(v int32) *SplitVideoPartsAdvanceRequest {
	s.MinTime = &v
	return s
}

func (s *SplitVideoPartsAdvanceRequest) SetTemplate(v string) *SplitVideoPartsAdvanceRequest {
	s.Template = &v
	return s
}

func (s *SplitVideoPartsAdvanceRequest) SetVideoUrlObject(v io.Reader) *SplitVideoPartsAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

type SplitVideoPartsResponseBody struct {
	Data    *SplitVideoPartsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A00A3C17-61D5-1489-860D-B709F83A7C40
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Elements              []*SplitVideoPartsResponseBodyDataElements              `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	SplitVideoPartResults []*SplitVideoPartsResponseBodyDataSplitVideoPartResults `json:"SplitVideoPartResults,omitempty" xml:"SplitVideoPartResults,omitempty" type:"Repeated"`
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

func (s *SplitVideoPartsResponseBodyData) SetSplitVideoPartResults(v []*SplitVideoPartsResponseBodyDataSplitVideoPartResults) *SplitVideoPartsResponseBodyData {
	s.SplitVideoPartResults = v
	return s
}

type SplitVideoPartsResponseBodyDataElements struct {
	// example:
	//
	// 10.06
	BeginTime *float32 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 17.3
	EndTime *float32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
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

type SplitVideoPartsResponseBodyDataSplitVideoPartResults struct {
	BeginTime *float32 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	By        *string  `json:"By,omitempty" xml:"By,omitempty"`
	EndTime   *float32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Theme     *string  `json:"Theme,omitempty" xml:"Theme,omitempty"`
	Type      *string  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SplitVideoPartsResponseBodyDataSplitVideoPartResults) String() string {
	return tea.Prettify(s)
}

func (s SplitVideoPartsResponseBodyDataSplitVideoPartResults) GoString() string {
	return s.String()
}

func (s *SplitVideoPartsResponseBodyDataSplitVideoPartResults) SetBeginTime(v float32) *SplitVideoPartsResponseBodyDataSplitVideoPartResults {
	s.BeginTime = &v
	return s
}

func (s *SplitVideoPartsResponseBodyDataSplitVideoPartResults) SetBy(v string) *SplitVideoPartsResponseBodyDataSplitVideoPartResults {
	s.By = &v
	return s
}

func (s *SplitVideoPartsResponseBodyDataSplitVideoPartResults) SetEndTime(v float32) *SplitVideoPartsResponseBodyDataSplitVideoPartResults {
	s.EndTime = &v
	return s
}

func (s *SplitVideoPartsResponseBodyDataSplitVideoPartResults) SetTheme(v string) *SplitVideoPartsResponseBodyDataSplitVideoPartResults {
	s.Theme = &v
	return s
}

func (s *SplitVideoPartsResponseBodyDataSplitVideoPartResults) SetType(v string) *SplitVideoPartsResponseBodyDataSplitVideoPartResults {
	s.Type = &v
	return s
}

type SplitVideoPartsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SplitVideoPartsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videorecog/UnderstandVideoContent/UnderstandVideoContent1.mp4
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
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videorecog/UnderstandVideoContent/UnderstandVideoContent1.mp4
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
	Data    *UnderstandVideoContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 71EC3F13-F0CA-4558-AC7F-A351106F59F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 43380
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 25.0
	Fps *float32 `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// example:
	//
	// 1280
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 720
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnderstandVideoContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

// @param request - DetectVideoShotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectVideoShotResponse
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

// @param request - DetectVideoShotRequest
//
// @return DetectVideoShotResponse
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

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

// Summary:
//
// 视频质量评估
//
// @param request - EvaluateVideoQualityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EvaluateVideoQualityResponse
func (client *Client) EvaluateVideoQualityWithOptions(request *EvaluateVideoQualityRequest, runtime *util.RuntimeOptions) (_result *EvaluateVideoQualityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EvaluateVideoQuality"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EvaluateVideoQualityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频质量评估
//
// @param request - EvaluateVideoQualityRequest
//
// @return EvaluateVideoQualityResponse
func (client *Client) EvaluateVideoQuality(request *EvaluateVideoQualityRequest) (_result *EvaluateVideoQualityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EvaluateVideoQualityResponse{}
	_body, _err := client.EvaluateVideoQualityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EvaluateVideoQualityAdvance(request *EvaluateVideoQualityAdvanceRequest, runtime *util.RuntimeOptions) (_result *EvaluateVideoQualityResponse, _err error) {
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	evaluateVideoQualityReq := &EvaluateVideoQualityRequest{}
	openapiutil.Convert(request, evaluateVideoQualityReq)
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
		evaluateVideoQualityReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	evaluateVideoQualityResp, _err := client.EvaluateVideoQualityWithOptions(evaluateVideoQualityReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = evaluateVideoQualityResp
	return _result, _err
}

// @param request - GenerateVideoCoverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateVideoCoverResponse
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

// @param request - GenerateVideoCoverRequest
//
// @return GenerateVideoCoverResponse
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

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

// @param request - GetAsyncJobResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsyncJobResultResponse
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

// @param request - GetAsyncJobResultRequest
//
// @return GetAsyncJobResultResponse
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

// Summary:
//
// 视频OCR
//
// @param tmpReq - RecognizeVideoCastCrewListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeVideoCastCrewListResponse
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

// Summary:
//
// 视频OCR
//
// @param request - RecognizeVideoCastCrewListRequest
//
// @return RecognizeVideoCastCrewListResponse
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeVideoCastCrewListReq := &RecognizeVideoCastCrewListRequest{}
	openapiutil.Convert(request, recognizeVideoCastCrewListReq)
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

// Summary:
//
// 视频拆条
//
// @param request - SplitVideoPartsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SplitVideoPartsResponse
func (client *Client) SplitVideoPartsWithOptions(request *SplitVideoPartsRequest, runtime *util.RuntimeOptions) (_result *SplitVideoPartsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxTime)) {
		body["MaxTime"] = request.MaxTime
	}

	if !tea.BoolValue(util.IsUnset(request.MinTime)) {
		body["MinTime"] = request.MinTime
	}

	if !tea.BoolValue(util.IsUnset(request.Template)) {
		body["Template"] = request.Template
	}

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

// Summary:
//
// 视频拆条
//
// @param request - SplitVideoPartsRequest
//
// @return SplitVideoPartsResponse
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

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

// Summary:
//
// 视频内容理解
//
// @param request - UnderstandVideoContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnderstandVideoContentResponse
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

// Summary:
//
// 视频内容理解
//
// @param request - UnderstandVideoContentRequest
//
// @return UnderstandVideoContentResponse
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
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
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
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

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
