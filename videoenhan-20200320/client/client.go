// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type GetAsyncJobResultRequest struct {
	Async *bool   `json:"Async,omitempty" xml:"Async,omitempty"`
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAsyncJobResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultRequest) SetAsync(v bool) *GetAsyncJobResultRequest {
	s.Async = &v
	return s
}

func (s *GetAsyncJobResultRequest) SetJobId(v string) *GetAsyncJobResultRequest {
	s.JobId = &v
	return s
}

type GetAsyncJobResultResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetAsyncJobResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetAsyncJobResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBody) SetRequestId(v string) *GetAsyncJobResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsyncJobResultResponseBody) SetData(v *GetAsyncJobResultResponseBodyData) *GetAsyncJobResultResponseBody {
	s.Data = v
	return s
}

type GetAsyncJobResultResponseBodyData struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAsyncJobResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBodyData) SetStatus(v string) *GetAsyncJobResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorMessage(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetResult(v string) *GetAsyncJobResultResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorCode(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetJobId(v string) *GetAsyncJobResultResponseBodyData {
	s.JobId = &v
	return s
}

type GetAsyncJobResultResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAsyncJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetAsyncJobResultResponse) SetBody(v *GetAsyncJobResultResponseBody) *GetAsyncJobResultResponse {
	s.Body = v
	return s
}

type EnhanceVideoQualityRequest struct {
	VideoURL       *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
	Async          *bool   `json:"Async,omitempty" xml:"Async,omitempty"`
	OutPutWidth    *int32  `json:"OutPutWidth,omitempty" xml:"OutPutWidth,omitempty"`
	OutPutHeight   *int32  `json:"OutPutHeight,omitempty" xml:"OutPutHeight,omitempty"`
	FrameRate      *int32  `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	HDRFormat      *string `json:"HDRFormat,omitempty" xml:"HDRFormat,omitempty"`
	MaxIlluminance *int32  `json:"MaxIlluminance,omitempty" xml:"MaxIlluminance,omitempty"`
	Bitrate        *int32  `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
}

func (s EnhanceVideoQualityRequest) String() string {
	return tea.Prettify(s)
}

func (s EnhanceVideoQualityRequest) GoString() string {
	return s.String()
}

func (s *EnhanceVideoQualityRequest) SetVideoURL(v string) *EnhanceVideoQualityRequest {
	s.VideoURL = &v
	return s
}

func (s *EnhanceVideoQualityRequest) SetAsync(v bool) *EnhanceVideoQualityRequest {
	s.Async = &v
	return s
}

func (s *EnhanceVideoQualityRequest) SetOutPutWidth(v int32) *EnhanceVideoQualityRequest {
	s.OutPutWidth = &v
	return s
}

func (s *EnhanceVideoQualityRequest) SetOutPutHeight(v int32) *EnhanceVideoQualityRequest {
	s.OutPutHeight = &v
	return s
}

func (s *EnhanceVideoQualityRequest) SetFrameRate(v int32) *EnhanceVideoQualityRequest {
	s.FrameRate = &v
	return s
}

func (s *EnhanceVideoQualityRequest) SetHDRFormat(v string) *EnhanceVideoQualityRequest {
	s.HDRFormat = &v
	return s
}

func (s *EnhanceVideoQualityRequest) SetMaxIlluminance(v int32) *EnhanceVideoQualityRequest {
	s.MaxIlluminance = &v
	return s
}

func (s *EnhanceVideoQualityRequest) SetBitrate(v int32) *EnhanceVideoQualityRequest {
	s.Bitrate = &v
	return s
}

type EnhanceVideoQualityAdvanceRequest struct {
	VideoURLObject io.Reader `json:"VideoURLObject,omitempty" xml:"VideoURLObject,omitempty" require:"true"`
	Async          *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
	OutPutWidth    *int32    `json:"OutPutWidth,omitempty" xml:"OutPutWidth,omitempty"`
	OutPutHeight   *int32    `json:"OutPutHeight,omitempty" xml:"OutPutHeight,omitempty"`
	FrameRate      *int32    `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	HDRFormat      *string   `json:"HDRFormat,omitempty" xml:"HDRFormat,omitempty"`
	MaxIlluminance *int32    `json:"MaxIlluminance,omitempty" xml:"MaxIlluminance,omitempty"`
	Bitrate        *int32    `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
}

func (s EnhanceVideoQualityAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s EnhanceVideoQualityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *EnhanceVideoQualityAdvanceRequest) SetVideoURLObject(v io.Reader) *EnhanceVideoQualityAdvanceRequest {
	s.VideoURLObject = v
	return s
}

func (s *EnhanceVideoQualityAdvanceRequest) SetAsync(v bool) *EnhanceVideoQualityAdvanceRequest {
	s.Async = &v
	return s
}

func (s *EnhanceVideoQualityAdvanceRequest) SetOutPutWidth(v int32) *EnhanceVideoQualityAdvanceRequest {
	s.OutPutWidth = &v
	return s
}

func (s *EnhanceVideoQualityAdvanceRequest) SetOutPutHeight(v int32) *EnhanceVideoQualityAdvanceRequest {
	s.OutPutHeight = &v
	return s
}

func (s *EnhanceVideoQualityAdvanceRequest) SetFrameRate(v int32) *EnhanceVideoQualityAdvanceRequest {
	s.FrameRate = &v
	return s
}

func (s *EnhanceVideoQualityAdvanceRequest) SetHDRFormat(v string) *EnhanceVideoQualityAdvanceRequest {
	s.HDRFormat = &v
	return s
}

func (s *EnhanceVideoQualityAdvanceRequest) SetMaxIlluminance(v int32) *EnhanceVideoQualityAdvanceRequest {
	s.MaxIlluminance = &v
	return s
}

func (s *EnhanceVideoQualityAdvanceRequest) SetBitrate(v int32) *EnhanceVideoQualityAdvanceRequest {
	s.Bitrate = &v
	return s
}

type EnhanceVideoQualityResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *EnhanceVideoQualityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s EnhanceVideoQualityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnhanceVideoQualityResponseBody) GoString() string {
	return s.String()
}

func (s *EnhanceVideoQualityResponseBody) SetRequestId(v string) *EnhanceVideoQualityResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnhanceVideoQualityResponseBody) SetData(v *EnhanceVideoQualityResponseBodyData) *EnhanceVideoQualityResponseBody {
	s.Data = v
	return s
}

type EnhanceVideoQualityResponseBodyData struct {
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s EnhanceVideoQualityResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnhanceVideoQualityResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnhanceVideoQualityResponseBodyData) SetVideoURL(v string) *EnhanceVideoQualityResponseBodyData {
	s.VideoURL = &v
	return s
}

type EnhanceVideoQualityResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnhanceVideoQualityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnhanceVideoQualityResponse) String() string {
	return tea.Prettify(s)
}

func (s EnhanceVideoQualityResponse) GoString() string {
	return s.String()
}

func (s *EnhanceVideoQualityResponse) SetHeaders(v map[string]*string) *EnhanceVideoQualityResponse {
	s.Headers = v
	return s
}

func (s *EnhanceVideoQualityResponse) SetBody(v *EnhanceVideoQualityResponseBody) *EnhanceVideoQualityResponse {
	s.Body = v
	return s
}

type GenerateVideoRequest struct {
	Async            *bool                           `json:"Async,omitempty" xml:"Async,omitempty"`
	Scene            *string                         `json:"Scene,omitempty" xml:"Scene,omitempty"`
	Width            *int32                          `json:"Width,omitempty" xml:"Width,omitempty"`
	Height           *int32                          `json:"Height,omitempty" xml:"Height,omitempty"`
	Style            *string                         `json:"Style,omitempty" xml:"Style,omitempty"`
	Duration         *float32                        `json:"Duration,omitempty" xml:"Duration,omitempty"`
	DurationAdaption *bool                           `json:"DurationAdaption,omitempty" xml:"DurationAdaption,omitempty"`
	TransitionStyle  *string                         `json:"TransitionStyle,omitempty" xml:"TransitionStyle,omitempty"`
	SmartEffect      *bool                           `json:"SmartEffect,omitempty" xml:"SmartEffect,omitempty"`
	PuzzleEffect     *bool                           `json:"PuzzleEffect,omitempty" xml:"PuzzleEffect,omitempty"`
	Mute             *bool                           `json:"Mute,omitempty" xml:"Mute,omitempty"`
	FileList         []*GenerateVideoRequestFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
}

func (s GenerateVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoRequest) GoString() string {
	return s.String()
}

func (s *GenerateVideoRequest) SetAsync(v bool) *GenerateVideoRequest {
	s.Async = &v
	return s
}

func (s *GenerateVideoRequest) SetScene(v string) *GenerateVideoRequest {
	s.Scene = &v
	return s
}

func (s *GenerateVideoRequest) SetWidth(v int32) *GenerateVideoRequest {
	s.Width = &v
	return s
}

func (s *GenerateVideoRequest) SetHeight(v int32) *GenerateVideoRequest {
	s.Height = &v
	return s
}

func (s *GenerateVideoRequest) SetStyle(v string) *GenerateVideoRequest {
	s.Style = &v
	return s
}

func (s *GenerateVideoRequest) SetDuration(v float32) *GenerateVideoRequest {
	s.Duration = &v
	return s
}

func (s *GenerateVideoRequest) SetDurationAdaption(v bool) *GenerateVideoRequest {
	s.DurationAdaption = &v
	return s
}

func (s *GenerateVideoRequest) SetTransitionStyle(v string) *GenerateVideoRequest {
	s.TransitionStyle = &v
	return s
}

func (s *GenerateVideoRequest) SetSmartEffect(v bool) *GenerateVideoRequest {
	s.SmartEffect = &v
	return s
}

func (s *GenerateVideoRequest) SetPuzzleEffect(v bool) *GenerateVideoRequest {
	s.PuzzleEffect = &v
	return s
}

func (s *GenerateVideoRequest) SetMute(v bool) *GenerateVideoRequest {
	s.Mute = &v
	return s
}

func (s *GenerateVideoRequest) SetFileList(v []*GenerateVideoRequestFileList) *GenerateVideoRequest {
	s.FileList = v
	return s
}

type GenerateVideoRequestFileList struct {
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	FileUrl  *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GenerateVideoRequestFileList) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoRequestFileList) GoString() string {
	return s.String()
}

func (s *GenerateVideoRequestFileList) SetType(v string) *GenerateVideoRequestFileList {
	s.Type = &v
	return s
}

func (s *GenerateVideoRequestFileList) SetFileUrl(v string) *GenerateVideoRequestFileList {
	s.FileUrl = &v
	return s
}

func (s *GenerateVideoRequestFileList) SetFileName(v string) *GenerateVideoRequestFileList {
	s.FileName = &v
	return s
}

type GenerateVideoResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GenerateVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GenerateVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateVideoResponseBody) SetRequestId(v string) *GenerateVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateVideoResponseBody) SetData(v *GenerateVideoResponseBodyData) *GenerateVideoResponseBody {
	s.Data = v
	return s
}

type GenerateVideoResponseBodyData struct {
	VideoCoverUrl *string `json:"VideoCoverUrl,omitempty" xml:"VideoCoverUrl,omitempty"`
	VideoUrl      *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GenerateVideoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateVideoResponseBodyData) SetVideoCoverUrl(v string) *GenerateVideoResponseBodyData {
	s.VideoCoverUrl = &v
	return s
}

func (s *GenerateVideoResponseBodyData) SetVideoUrl(v string) *GenerateVideoResponseBodyData {
	s.VideoUrl = &v
	return s
}

type GenerateVideoResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateVideoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoResponse) GoString() string {
	return s.String()
}

func (s *GenerateVideoResponse) SetHeaders(v map[string]*string) *GenerateVideoResponse {
	s.Headers = v
	return s
}

func (s *GenerateVideoResponse) SetBody(v *GenerateVideoResponseBody) *GenerateVideoResponse {
	s.Body = v
	return s
}

type EraseVideoLogoRequest struct {
	VideoUrl *string                       `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	Async    *bool                         `json:"Async,omitempty" xml:"Async,omitempty"`
	Boxes    []*EraseVideoLogoRequestBoxes `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
}

func (s EraseVideoLogoRequest) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoLogoRequest) GoString() string {
	return s.String()
}

func (s *EraseVideoLogoRequest) SetVideoUrl(v string) *EraseVideoLogoRequest {
	s.VideoUrl = &v
	return s
}

func (s *EraseVideoLogoRequest) SetAsync(v bool) *EraseVideoLogoRequest {
	s.Async = &v
	return s
}

func (s *EraseVideoLogoRequest) SetBoxes(v []*EraseVideoLogoRequestBoxes) *EraseVideoLogoRequest {
	s.Boxes = v
	return s
}

type EraseVideoLogoRequestBoxes struct {
	W *float32 `json:"W,omitempty" xml:"W,omitempty"`
	H *float32 `json:"H,omitempty" xml:"H,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s EraseVideoLogoRequestBoxes) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoLogoRequestBoxes) GoString() string {
	return s.String()
}

func (s *EraseVideoLogoRequestBoxes) SetW(v float32) *EraseVideoLogoRequestBoxes {
	s.W = &v
	return s
}

func (s *EraseVideoLogoRequestBoxes) SetH(v float32) *EraseVideoLogoRequestBoxes {
	s.H = &v
	return s
}

func (s *EraseVideoLogoRequestBoxes) SetY(v float32) *EraseVideoLogoRequestBoxes {
	s.Y = &v
	return s
}

func (s *EraseVideoLogoRequestBoxes) SetX(v float32) *EraseVideoLogoRequestBoxes {
	s.X = &v
	return s
}

type EraseVideoLogoAdvanceRequest struct {
	VideoUrlObject io.Reader                            `json:"VideoUrlObject,omitempty" xml:"VideoUrlObject,omitempty" require:"true"`
	Async          *bool                                `json:"Async,omitempty" xml:"Async,omitempty"`
	Boxes          []*EraseVideoLogoAdvanceRequestBoxes `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
}

func (s EraseVideoLogoAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoLogoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *EraseVideoLogoAdvanceRequest) SetVideoUrlObject(v io.Reader) *EraseVideoLogoAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *EraseVideoLogoAdvanceRequest) SetAsync(v bool) *EraseVideoLogoAdvanceRequest {
	s.Async = &v
	return s
}

func (s *EraseVideoLogoAdvanceRequest) SetBoxes(v []*EraseVideoLogoAdvanceRequestBoxes) *EraseVideoLogoAdvanceRequest {
	s.Boxes = v
	return s
}

type EraseVideoLogoAdvanceRequestBoxes struct {
	W *float32 `json:"W,omitempty" xml:"W,omitempty"`
	H *float32 `json:"H,omitempty" xml:"H,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s EraseVideoLogoAdvanceRequestBoxes) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoLogoAdvanceRequestBoxes) GoString() string {
	return s.String()
}

func (s *EraseVideoLogoAdvanceRequestBoxes) SetW(v float32) *EraseVideoLogoAdvanceRequestBoxes {
	s.W = &v
	return s
}

func (s *EraseVideoLogoAdvanceRequestBoxes) SetH(v float32) *EraseVideoLogoAdvanceRequestBoxes {
	s.H = &v
	return s
}

func (s *EraseVideoLogoAdvanceRequestBoxes) SetY(v float32) *EraseVideoLogoAdvanceRequestBoxes {
	s.Y = &v
	return s
}

func (s *EraseVideoLogoAdvanceRequestBoxes) SetX(v float32) *EraseVideoLogoAdvanceRequestBoxes {
	s.X = &v
	return s
}

type EraseVideoLogoResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *EraseVideoLogoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s EraseVideoLogoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoLogoResponseBody) GoString() string {
	return s.String()
}

func (s *EraseVideoLogoResponseBody) SetRequestId(v string) *EraseVideoLogoResponseBody {
	s.RequestId = &v
	return s
}

func (s *EraseVideoLogoResponseBody) SetData(v *EraseVideoLogoResponseBodyData) *EraseVideoLogoResponseBody {
	s.Data = v
	return s
}

type EraseVideoLogoResponseBodyData struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EraseVideoLogoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoLogoResponseBodyData) GoString() string {
	return s.String()
}

func (s *EraseVideoLogoResponseBodyData) SetVideoUrl(v string) *EraseVideoLogoResponseBodyData {
	s.VideoUrl = &v
	return s
}

type EraseVideoLogoResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EraseVideoLogoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EraseVideoLogoResponse) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoLogoResponse) GoString() string {
	return s.String()
}

func (s *EraseVideoLogoResponse) SetHeaders(v map[string]*string) *EraseVideoLogoResponse {
	s.Headers = v
	return s
}

func (s *EraseVideoLogoResponse) SetBody(v *EraseVideoLogoResponseBody) *EraseVideoLogoResponse {
	s.Body = v
	return s
}

type AbstractEcommerceVideoRequest struct {
	VideoUrl *string  `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	Async    *bool    `json:"Async,omitempty" xml:"Async,omitempty"`
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Width    *int32   `json:"Width,omitempty" xml:"Width,omitempty"`
	Height   *int32   `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s AbstractEcommerceVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s AbstractEcommerceVideoRequest) GoString() string {
	return s.String()
}

func (s *AbstractEcommerceVideoRequest) SetVideoUrl(v string) *AbstractEcommerceVideoRequest {
	s.VideoUrl = &v
	return s
}

func (s *AbstractEcommerceVideoRequest) SetAsync(v bool) *AbstractEcommerceVideoRequest {
	s.Async = &v
	return s
}

func (s *AbstractEcommerceVideoRequest) SetDuration(v float32) *AbstractEcommerceVideoRequest {
	s.Duration = &v
	return s
}

func (s *AbstractEcommerceVideoRequest) SetWidth(v int32) *AbstractEcommerceVideoRequest {
	s.Width = &v
	return s
}

func (s *AbstractEcommerceVideoRequest) SetHeight(v int32) *AbstractEcommerceVideoRequest {
	s.Height = &v
	return s
}

type AbstractEcommerceVideoAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrlObject,omitempty" xml:"VideoUrlObject,omitempty" require:"true"`
	Async          *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
	Duration       *float32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Width          *int32    `json:"Width,omitempty" xml:"Width,omitempty"`
	Height         *int32    `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s AbstractEcommerceVideoAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AbstractEcommerceVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AbstractEcommerceVideoAdvanceRequest) SetVideoUrlObject(v io.Reader) *AbstractEcommerceVideoAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *AbstractEcommerceVideoAdvanceRequest) SetAsync(v bool) *AbstractEcommerceVideoAdvanceRequest {
	s.Async = &v
	return s
}

func (s *AbstractEcommerceVideoAdvanceRequest) SetDuration(v float32) *AbstractEcommerceVideoAdvanceRequest {
	s.Duration = &v
	return s
}

func (s *AbstractEcommerceVideoAdvanceRequest) SetWidth(v int32) *AbstractEcommerceVideoAdvanceRequest {
	s.Width = &v
	return s
}

func (s *AbstractEcommerceVideoAdvanceRequest) SetHeight(v int32) *AbstractEcommerceVideoAdvanceRequest {
	s.Height = &v
	return s
}

type AbstractEcommerceVideoResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *AbstractEcommerceVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s AbstractEcommerceVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbstractEcommerceVideoResponseBody) GoString() string {
	return s.String()
}

func (s *AbstractEcommerceVideoResponseBody) SetRequestId(v string) *AbstractEcommerceVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbstractEcommerceVideoResponseBody) SetData(v *AbstractEcommerceVideoResponseBodyData) *AbstractEcommerceVideoResponseBody {
	s.Data = v
	return s
}

type AbstractEcommerceVideoResponseBodyData struct {
	VideoCoverUrl *string `json:"VideoCoverUrl,omitempty" xml:"VideoCoverUrl,omitempty"`
	VideoUrl      *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s AbstractEcommerceVideoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AbstractEcommerceVideoResponseBodyData) GoString() string {
	return s.String()
}

func (s *AbstractEcommerceVideoResponseBodyData) SetVideoCoverUrl(v string) *AbstractEcommerceVideoResponseBodyData {
	s.VideoCoverUrl = &v
	return s
}

func (s *AbstractEcommerceVideoResponseBodyData) SetVideoUrl(v string) *AbstractEcommerceVideoResponseBodyData {
	s.VideoUrl = &v
	return s
}

type AbstractEcommerceVideoResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AbstractEcommerceVideoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AbstractEcommerceVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s AbstractEcommerceVideoResponse) GoString() string {
	return s.String()
}

func (s *AbstractEcommerceVideoResponse) SetHeaders(v map[string]*string) *AbstractEcommerceVideoResponse {
	s.Headers = v
	return s
}

func (s *AbstractEcommerceVideoResponse) SetBody(v *AbstractEcommerceVideoResponseBody) *AbstractEcommerceVideoResponse {
	s.Body = v
	return s
}

type AbstractFilmVideoRequest struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	Async    *bool   `json:"Async,omitempty" xml:"Async,omitempty"`
	Length   *int32  `json:"Length,omitempty" xml:"Length,omitempty"`
}

func (s AbstractFilmVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s AbstractFilmVideoRequest) GoString() string {
	return s.String()
}

func (s *AbstractFilmVideoRequest) SetVideoUrl(v string) *AbstractFilmVideoRequest {
	s.VideoUrl = &v
	return s
}

func (s *AbstractFilmVideoRequest) SetAsync(v bool) *AbstractFilmVideoRequest {
	s.Async = &v
	return s
}

func (s *AbstractFilmVideoRequest) SetLength(v int32) *AbstractFilmVideoRequest {
	s.Length = &v
	return s
}

type AbstractFilmVideoAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrlObject,omitempty" xml:"VideoUrlObject,omitempty" require:"true"`
	Async          *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
	Length         *int32    `json:"Length,omitempty" xml:"Length,omitempty"`
}

func (s AbstractFilmVideoAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AbstractFilmVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AbstractFilmVideoAdvanceRequest) SetVideoUrlObject(v io.Reader) *AbstractFilmVideoAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *AbstractFilmVideoAdvanceRequest) SetAsync(v bool) *AbstractFilmVideoAdvanceRequest {
	s.Async = &v
	return s
}

func (s *AbstractFilmVideoAdvanceRequest) SetLength(v int32) *AbstractFilmVideoAdvanceRequest {
	s.Length = &v
	return s
}

type AbstractFilmVideoResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *AbstractFilmVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s AbstractFilmVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbstractFilmVideoResponseBody) GoString() string {
	return s.String()
}

func (s *AbstractFilmVideoResponseBody) SetRequestId(v string) *AbstractFilmVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbstractFilmVideoResponseBody) SetData(v *AbstractFilmVideoResponseBodyData) *AbstractFilmVideoResponseBody {
	s.Data = v
	return s
}

type AbstractFilmVideoResponseBodyData struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s AbstractFilmVideoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AbstractFilmVideoResponseBodyData) GoString() string {
	return s.String()
}

func (s *AbstractFilmVideoResponseBodyData) SetVideoUrl(v string) *AbstractFilmVideoResponseBodyData {
	s.VideoUrl = &v
	return s
}

type AbstractFilmVideoResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AbstractFilmVideoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AbstractFilmVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s AbstractFilmVideoResponse) GoString() string {
	return s.String()
}

func (s *AbstractFilmVideoResponse) SetHeaders(v map[string]*string) *AbstractFilmVideoResponse {
	s.Headers = v
	return s
}

func (s *AbstractFilmVideoResponse) SetBody(v *AbstractFilmVideoResponseBody) *AbstractFilmVideoResponse {
	s.Body = v
	return s
}

type AdjustVideoColorRequest struct {
	VideoUrl     *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	Async        *bool   `json:"Async,omitempty" xml:"Async,omitempty"`
	VideoBitrate *string `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	VideoCodec   *string `json:"VideoCodec,omitempty" xml:"VideoCodec,omitempty"`
	VideoFormat  *string `json:"VideoFormat,omitempty" xml:"VideoFormat,omitempty"`
	Mode         *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s AdjustVideoColorRequest) String() string {
	return tea.Prettify(s)
}

func (s AdjustVideoColorRequest) GoString() string {
	return s.String()
}

func (s *AdjustVideoColorRequest) SetVideoUrl(v string) *AdjustVideoColorRequest {
	s.VideoUrl = &v
	return s
}

func (s *AdjustVideoColorRequest) SetAsync(v bool) *AdjustVideoColorRequest {
	s.Async = &v
	return s
}

func (s *AdjustVideoColorRequest) SetVideoBitrate(v string) *AdjustVideoColorRequest {
	s.VideoBitrate = &v
	return s
}

func (s *AdjustVideoColorRequest) SetVideoCodec(v string) *AdjustVideoColorRequest {
	s.VideoCodec = &v
	return s
}

func (s *AdjustVideoColorRequest) SetVideoFormat(v string) *AdjustVideoColorRequest {
	s.VideoFormat = &v
	return s
}

func (s *AdjustVideoColorRequest) SetMode(v string) *AdjustVideoColorRequest {
	s.Mode = &v
	return s
}

type AdjustVideoColorAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrlObject,omitempty" xml:"VideoUrlObject,omitempty" require:"true"`
	Async          *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
	VideoBitrate   *string   `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	VideoCodec     *string   `json:"VideoCodec,omitempty" xml:"VideoCodec,omitempty"`
	VideoFormat    *string   `json:"VideoFormat,omitempty" xml:"VideoFormat,omitempty"`
	Mode           *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s AdjustVideoColorAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AdjustVideoColorAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AdjustVideoColorAdvanceRequest) SetVideoUrlObject(v io.Reader) *AdjustVideoColorAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *AdjustVideoColorAdvanceRequest) SetAsync(v bool) *AdjustVideoColorAdvanceRequest {
	s.Async = &v
	return s
}

func (s *AdjustVideoColorAdvanceRequest) SetVideoBitrate(v string) *AdjustVideoColorAdvanceRequest {
	s.VideoBitrate = &v
	return s
}

func (s *AdjustVideoColorAdvanceRequest) SetVideoCodec(v string) *AdjustVideoColorAdvanceRequest {
	s.VideoCodec = &v
	return s
}

func (s *AdjustVideoColorAdvanceRequest) SetVideoFormat(v string) *AdjustVideoColorAdvanceRequest {
	s.VideoFormat = &v
	return s
}

func (s *AdjustVideoColorAdvanceRequest) SetMode(v string) *AdjustVideoColorAdvanceRequest {
	s.Mode = &v
	return s
}

type AdjustVideoColorResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *AdjustVideoColorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s AdjustVideoColorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AdjustVideoColorResponseBody) GoString() string {
	return s.String()
}

func (s *AdjustVideoColorResponseBody) SetRequestId(v string) *AdjustVideoColorResponseBody {
	s.RequestId = &v
	return s
}

func (s *AdjustVideoColorResponseBody) SetData(v *AdjustVideoColorResponseBodyData) *AdjustVideoColorResponseBody {
	s.Data = v
	return s
}

type AdjustVideoColorResponseBodyData struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s AdjustVideoColorResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AdjustVideoColorResponseBodyData) GoString() string {
	return s.String()
}

func (s *AdjustVideoColorResponseBodyData) SetVideoUrl(v string) *AdjustVideoColorResponseBodyData {
	s.VideoUrl = &v
	return s
}

type AdjustVideoColorResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AdjustVideoColorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AdjustVideoColorResponse) String() string {
	return tea.Prettify(s)
}

func (s AdjustVideoColorResponse) GoString() string {
	return s.String()
}

func (s *AdjustVideoColorResponse) SetHeaders(v map[string]*string) *AdjustVideoColorResponse {
	s.Headers = v
	return s
}

func (s *AdjustVideoColorResponse) SetBody(v *AdjustVideoColorResponseBody) *AdjustVideoColorResponse {
	s.Body = v
	return s
}

type EraseVideoSubtitlesRequest struct {
	VideoUrl *string  `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	Async    *bool    `json:"Async,omitempty" xml:"Async,omitempty"`
	BX       *float32 `json:"BX,omitempty" xml:"BX,omitempty"`
	BY       *float32 `json:"BY,omitempty" xml:"BY,omitempty"`
	BW       *float32 `json:"BW,omitempty" xml:"BW,omitempty"`
	BH       *float32 `json:"BH,omitempty" xml:"BH,omitempty"`
}

func (s EraseVideoSubtitlesRequest) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoSubtitlesRequest) GoString() string {
	return s.String()
}

func (s *EraseVideoSubtitlesRequest) SetVideoUrl(v string) *EraseVideoSubtitlesRequest {
	s.VideoUrl = &v
	return s
}

func (s *EraseVideoSubtitlesRequest) SetAsync(v bool) *EraseVideoSubtitlesRequest {
	s.Async = &v
	return s
}

func (s *EraseVideoSubtitlesRequest) SetBX(v float32) *EraseVideoSubtitlesRequest {
	s.BX = &v
	return s
}

func (s *EraseVideoSubtitlesRequest) SetBY(v float32) *EraseVideoSubtitlesRequest {
	s.BY = &v
	return s
}

func (s *EraseVideoSubtitlesRequest) SetBW(v float32) *EraseVideoSubtitlesRequest {
	s.BW = &v
	return s
}

func (s *EraseVideoSubtitlesRequest) SetBH(v float32) *EraseVideoSubtitlesRequest {
	s.BH = &v
	return s
}

type EraseVideoSubtitlesAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrlObject,omitempty" xml:"VideoUrlObject,omitempty" require:"true"`
	Async          *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
	BX             *float32  `json:"BX,omitempty" xml:"BX,omitempty"`
	BY             *float32  `json:"BY,omitempty" xml:"BY,omitempty"`
	BW             *float32  `json:"BW,omitempty" xml:"BW,omitempty"`
	BH             *float32  `json:"BH,omitempty" xml:"BH,omitempty"`
}

func (s EraseVideoSubtitlesAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoSubtitlesAdvanceRequest) GoString() string {
	return s.String()
}

func (s *EraseVideoSubtitlesAdvanceRequest) SetVideoUrlObject(v io.Reader) *EraseVideoSubtitlesAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *EraseVideoSubtitlesAdvanceRequest) SetAsync(v bool) *EraseVideoSubtitlesAdvanceRequest {
	s.Async = &v
	return s
}

func (s *EraseVideoSubtitlesAdvanceRequest) SetBX(v float32) *EraseVideoSubtitlesAdvanceRequest {
	s.BX = &v
	return s
}

func (s *EraseVideoSubtitlesAdvanceRequest) SetBY(v float32) *EraseVideoSubtitlesAdvanceRequest {
	s.BY = &v
	return s
}

func (s *EraseVideoSubtitlesAdvanceRequest) SetBW(v float32) *EraseVideoSubtitlesAdvanceRequest {
	s.BW = &v
	return s
}

func (s *EraseVideoSubtitlesAdvanceRequest) SetBH(v float32) *EraseVideoSubtitlesAdvanceRequest {
	s.BH = &v
	return s
}

type EraseVideoSubtitlesResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *EraseVideoSubtitlesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s EraseVideoSubtitlesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoSubtitlesResponseBody) GoString() string {
	return s.String()
}

func (s *EraseVideoSubtitlesResponseBody) SetRequestId(v string) *EraseVideoSubtitlesResponseBody {
	s.RequestId = &v
	return s
}

func (s *EraseVideoSubtitlesResponseBody) SetData(v *EraseVideoSubtitlesResponseBodyData) *EraseVideoSubtitlesResponseBody {
	s.Data = v
	return s
}

type EraseVideoSubtitlesResponseBodyData struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EraseVideoSubtitlesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoSubtitlesResponseBodyData) GoString() string {
	return s.String()
}

func (s *EraseVideoSubtitlesResponseBodyData) SetVideoUrl(v string) *EraseVideoSubtitlesResponseBodyData {
	s.VideoUrl = &v
	return s
}

type EraseVideoSubtitlesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EraseVideoSubtitlesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EraseVideoSubtitlesResponse) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoSubtitlesResponse) GoString() string {
	return s.String()
}

func (s *EraseVideoSubtitlesResponse) SetHeaders(v map[string]*string) *EraseVideoSubtitlesResponse {
	s.Headers = v
	return s
}

func (s *EraseVideoSubtitlesResponse) SetBody(v *EraseVideoSubtitlesResponseBody) *EraseVideoSubtitlesResponse {
	s.Body = v
	return s
}

type ChangeVideoSizeRequest struct {
	VideoUrl  *string  `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	Async     *bool    `json:"Async,omitempty" xml:"Async,omitempty"`
	Width     *int32   `json:"Width,omitempty" xml:"Width,omitempty"`
	Height    *int32   `json:"Height,omitempty" xml:"Height,omitempty"`
	CropType  *string  `json:"CropType,omitempty" xml:"CropType,omitempty"`
	FillType  *string  `json:"FillType,omitempty" xml:"FillType,omitempty"`
	Tightness *float32 `json:"Tightness,omitempty" xml:"Tightness,omitempty"`
	R         *int32   `json:"R,omitempty" xml:"R,omitempty"`
	G         *int32   `json:"G,omitempty" xml:"G,omitempty"`
	B         *int32   `json:"B,omitempty" xml:"B,omitempty"`
}

func (s ChangeVideoSizeRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeVideoSizeRequest) GoString() string {
	return s.String()
}

func (s *ChangeVideoSizeRequest) SetVideoUrl(v string) *ChangeVideoSizeRequest {
	s.VideoUrl = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetAsync(v bool) *ChangeVideoSizeRequest {
	s.Async = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetWidth(v int32) *ChangeVideoSizeRequest {
	s.Width = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetHeight(v int32) *ChangeVideoSizeRequest {
	s.Height = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetCropType(v string) *ChangeVideoSizeRequest {
	s.CropType = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetFillType(v string) *ChangeVideoSizeRequest {
	s.FillType = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetTightness(v float32) *ChangeVideoSizeRequest {
	s.Tightness = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetR(v int32) *ChangeVideoSizeRequest {
	s.R = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetG(v int32) *ChangeVideoSizeRequest {
	s.G = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetB(v int32) *ChangeVideoSizeRequest {
	s.B = &v
	return s
}

type ChangeVideoSizeAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrlObject,omitempty" xml:"VideoUrlObject,omitempty" require:"true"`
	Async          *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
	Width          *int32    `json:"Width,omitempty" xml:"Width,omitempty"`
	Height         *int32    `json:"Height,omitempty" xml:"Height,omitempty"`
	CropType       *string   `json:"CropType,omitempty" xml:"CropType,omitempty"`
	FillType       *string   `json:"FillType,omitempty" xml:"FillType,omitempty"`
	Tightness      *float32  `json:"Tightness,omitempty" xml:"Tightness,omitempty"`
	R              *int32    `json:"R,omitempty" xml:"R,omitempty"`
	G              *int32    `json:"G,omitempty" xml:"G,omitempty"`
	B              *int32    `json:"B,omitempty" xml:"B,omitempty"`
}

func (s ChangeVideoSizeAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeVideoSizeAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ChangeVideoSizeAdvanceRequest) SetVideoUrlObject(v io.Reader) *ChangeVideoSizeAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetAsync(v bool) *ChangeVideoSizeAdvanceRequest {
	s.Async = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetWidth(v int32) *ChangeVideoSizeAdvanceRequest {
	s.Width = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetHeight(v int32) *ChangeVideoSizeAdvanceRequest {
	s.Height = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetCropType(v string) *ChangeVideoSizeAdvanceRequest {
	s.CropType = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetFillType(v string) *ChangeVideoSizeAdvanceRequest {
	s.FillType = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetTightness(v float32) *ChangeVideoSizeAdvanceRequest {
	s.Tightness = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetR(v int32) *ChangeVideoSizeAdvanceRequest {
	s.R = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetG(v int32) *ChangeVideoSizeAdvanceRequest {
	s.G = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetB(v int32) *ChangeVideoSizeAdvanceRequest {
	s.B = &v
	return s
}

type ChangeVideoSizeResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ChangeVideoSizeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ChangeVideoSizeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeVideoSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeVideoSizeResponseBody) SetRequestId(v string) *ChangeVideoSizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeVideoSizeResponseBody) SetData(v *ChangeVideoSizeResponseBodyData) *ChangeVideoSizeResponseBody {
	s.Data = v
	return s
}

type ChangeVideoSizeResponseBodyData struct {
	VideoCoverUrl *string `json:"VideoCoverUrl,omitempty" xml:"VideoCoverUrl,omitempty"`
	VideoUrl      *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s ChangeVideoSizeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ChangeVideoSizeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeVideoSizeResponseBodyData) SetVideoCoverUrl(v string) *ChangeVideoSizeResponseBodyData {
	s.VideoCoverUrl = &v
	return s
}

func (s *ChangeVideoSizeResponseBodyData) SetVideoUrl(v string) *ChangeVideoSizeResponseBodyData {
	s.VideoUrl = &v
	return s
}

type ChangeVideoSizeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ChangeVideoSizeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeVideoSizeResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeVideoSizeResponse) GoString() string {
	return s.String()
}

func (s *ChangeVideoSizeResponse) SetHeaders(v map[string]*string) *ChangeVideoSizeResponse {
	s.Headers = v
	return s
}

func (s *ChangeVideoSizeResponse) SetBody(v *ChangeVideoSizeResponseBody) *ChangeVideoSizeResponse {
	s.Body = v
	return s
}

type MergeVideoFaceRequest struct {
	VideoURL     *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
	Async        *bool   `json:"Async,omitempty" xml:"Async,omitempty"`
	PostURL      *string `json:"PostURL,omitempty" xml:"PostURL,omitempty"`
	ReferenceURL *string `json:"ReferenceURL,omitempty" xml:"ReferenceURL,omitempty"`
}

func (s MergeVideoFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s MergeVideoFaceRequest) GoString() string {
	return s.String()
}

func (s *MergeVideoFaceRequest) SetVideoURL(v string) *MergeVideoFaceRequest {
	s.VideoURL = &v
	return s
}

func (s *MergeVideoFaceRequest) SetAsync(v bool) *MergeVideoFaceRequest {
	s.Async = &v
	return s
}

func (s *MergeVideoFaceRequest) SetPostURL(v string) *MergeVideoFaceRequest {
	s.PostURL = &v
	return s
}

func (s *MergeVideoFaceRequest) SetReferenceURL(v string) *MergeVideoFaceRequest {
	s.ReferenceURL = &v
	return s
}

type MergeVideoFaceAdvanceRequest struct {
	VideoURLObject io.Reader `json:"VideoURLObject,omitempty" xml:"VideoURLObject,omitempty" require:"true"`
	Async          *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
	PostURL        *string   `json:"PostURL,omitempty" xml:"PostURL,omitempty"`
	ReferenceURL   *string   `json:"ReferenceURL,omitempty" xml:"ReferenceURL,omitempty"`
}

func (s MergeVideoFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s MergeVideoFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *MergeVideoFaceAdvanceRequest) SetVideoURLObject(v io.Reader) *MergeVideoFaceAdvanceRequest {
	s.VideoURLObject = v
	return s
}

func (s *MergeVideoFaceAdvanceRequest) SetAsync(v bool) *MergeVideoFaceAdvanceRequest {
	s.Async = &v
	return s
}

func (s *MergeVideoFaceAdvanceRequest) SetPostURL(v string) *MergeVideoFaceAdvanceRequest {
	s.PostURL = &v
	return s
}

func (s *MergeVideoFaceAdvanceRequest) SetReferenceURL(v string) *MergeVideoFaceAdvanceRequest {
	s.ReferenceURL = &v
	return s
}

type MergeVideoFaceResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *MergeVideoFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s MergeVideoFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MergeVideoFaceResponseBody) GoString() string {
	return s.String()
}

func (s *MergeVideoFaceResponseBody) SetRequestId(v string) *MergeVideoFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MergeVideoFaceResponseBody) SetData(v *MergeVideoFaceResponseBodyData) *MergeVideoFaceResponseBody {
	s.Data = v
	return s
}

type MergeVideoFaceResponseBodyData struct {
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s MergeVideoFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MergeVideoFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *MergeVideoFaceResponseBodyData) SetVideoURL(v string) *MergeVideoFaceResponseBodyData {
	s.VideoURL = &v
	return s
}

type MergeVideoFaceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MergeVideoFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MergeVideoFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s MergeVideoFaceResponse) GoString() string {
	return s.String()
}

func (s *MergeVideoFaceResponse) SetHeaders(v map[string]*string) *MergeVideoFaceResponse {
	s.Headers = v
	return s
}

func (s *MergeVideoFaceResponse) SetBody(v *MergeVideoFaceResponseBody) *MergeVideoFaceResponse {
	s.Body = v
	return s
}

type SuperResolveVideoRequest struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	Async    *bool   `json:"Async,omitempty" xml:"Async,omitempty"`
	BitRate  *int32  `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
}

func (s SuperResolveVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s SuperResolveVideoRequest) GoString() string {
	return s.String()
}

func (s *SuperResolveVideoRequest) SetVideoUrl(v string) *SuperResolveVideoRequest {
	s.VideoUrl = &v
	return s
}

func (s *SuperResolveVideoRequest) SetAsync(v bool) *SuperResolveVideoRequest {
	s.Async = &v
	return s
}

func (s *SuperResolveVideoRequest) SetBitRate(v int32) *SuperResolveVideoRequest {
	s.BitRate = &v
	return s
}

type SuperResolveVideoAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrlObject,omitempty" xml:"VideoUrlObject,omitempty" require:"true"`
	Async          *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
	BitRate        *int32    `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
}

func (s SuperResolveVideoAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SuperResolveVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SuperResolveVideoAdvanceRequest) SetVideoUrlObject(v io.Reader) *SuperResolveVideoAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *SuperResolveVideoAdvanceRequest) SetAsync(v bool) *SuperResolveVideoAdvanceRequest {
	s.Async = &v
	return s
}

func (s *SuperResolveVideoAdvanceRequest) SetBitRate(v int32) *SuperResolveVideoAdvanceRequest {
	s.BitRate = &v
	return s
}

type SuperResolveVideoResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *SuperResolveVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s SuperResolveVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SuperResolveVideoResponseBody) GoString() string {
	return s.String()
}

func (s *SuperResolveVideoResponseBody) SetRequestId(v string) *SuperResolveVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuperResolveVideoResponseBody) SetData(v *SuperResolveVideoResponseBodyData) *SuperResolveVideoResponseBody {
	s.Data = v
	return s
}

type SuperResolveVideoResponseBodyData struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s SuperResolveVideoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SuperResolveVideoResponseBodyData) GoString() string {
	return s.String()
}

func (s *SuperResolveVideoResponseBodyData) SetVideoUrl(v string) *SuperResolveVideoResponseBodyData {
	s.VideoUrl = &v
	return s
}

type SuperResolveVideoResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SuperResolveVideoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SuperResolveVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s SuperResolveVideoResponse) GoString() string {
	return s.String()
}

func (s *SuperResolveVideoResponse) SetHeaders(v map[string]*string) *SuperResolveVideoResponse {
	s.Headers = v
	return s
}

func (s *SuperResolveVideoResponse) SetBody(v *SuperResolveVideoResponseBody) *SuperResolveVideoResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("videoenhan"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetAsyncJobResultWithOptions(request *GetAsyncJobResultRequest, runtime *util.RuntimeOptions) (_result *GetAsyncJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAsyncJobResult"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) EnhanceVideoQualityWithOptions(request *EnhanceVideoQualityRequest, runtime *util.RuntimeOptions) (_result *EnhanceVideoQualityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnhanceVideoQualityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnhanceVideoQuality"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnhanceVideoQuality(request *EnhanceVideoQualityRequest) (_result *EnhanceVideoQualityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnhanceVideoQualityResponse{}
	_body, _err := client.EnhanceVideoQualityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnhanceVideoQualityAdvance(request *EnhanceVideoQualityAdvanceRequest, runtime *util.RuntimeOptions) (_result *EnhanceVideoQualityResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("videoenhan"),
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
	enhanceVideoQualityReq := &EnhanceVideoQualityRequest{}
	openapiutil.Convert(request, enhanceVideoQualityReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.VideoURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	enhanceVideoQualityReq.VideoURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	enhanceVideoQualityResp, _err := client.EnhanceVideoQualityWithOptions(enhanceVideoQualityReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = enhanceVideoQualityResp
	return _result, _err
}

func (client *Client) GenerateVideoWithOptions(request *GenerateVideoRequest, runtime *util.RuntimeOptions) (_result *GenerateVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateVideoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateVideo"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateVideo(request *GenerateVideoRequest) (_result *GenerateVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateVideoResponse{}
	_body, _err := client.GenerateVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EraseVideoLogoWithOptions(request *EraseVideoLogoRequest, runtime *util.RuntimeOptions) (_result *EraseVideoLogoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EraseVideoLogoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EraseVideoLogo"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EraseVideoLogo(request *EraseVideoLogoRequest) (_result *EraseVideoLogoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EraseVideoLogoResponse{}
	_body, _err := client.EraseVideoLogoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EraseVideoLogoAdvance(request *EraseVideoLogoAdvanceRequest, runtime *util.RuntimeOptions) (_result *EraseVideoLogoResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("videoenhan"),
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
	eraseVideoLogoReq := &EraseVideoLogoRequest{}
	openapiutil.Convert(request, eraseVideoLogoReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.VideoUrlObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	eraseVideoLogoReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	eraseVideoLogoResp, _err := client.EraseVideoLogoWithOptions(eraseVideoLogoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = eraseVideoLogoResp
	return _result, _err
}

func (client *Client) AbstractEcommerceVideoWithOptions(request *AbstractEcommerceVideoRequest, runtime *util.RuntimeOptions) (_result *AbstractEcommerceVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AbstractEcommerceVideoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AbstractEcommerceVideo"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AbstractEcommerceVideo(request *AbstractEcommerceVideoRequest) (_result *AbstractEcommerceVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AbstractEcommerceVideoResponse{}
	_body, _err := client.AbstractEcommerceVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AbstractEcommerceVideoAdvance(request *AbstractEcommerceVideoAdvanceRequest, runtime *util.RuntimeOptions) (_result *AbstractEcommerceVideoResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("videoenhan"),
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
	abstractEcommerceVideoReq := &AbstractEcommerceVideoRequest{}
	openapiutil.Convert(request, abstractEcommerceVideoReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.VideoUrlObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	abstractEcommerceVideoReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	abstractEcommerceVideoResp, _err := client.AbstractEcommerceVideoWithOptions(abstractEcommerceVideoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = abstractEcommerceVideoResp
	return _result, _err
}

func (client *Client) AbstractFilmVideoWithOptions(request *AbstractFilmVideoRequest, runtime *util.RuntimeOptions) (_result *AbstractFilmVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AbstractFilmVideoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AbstractFilmVideo"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AbstractFilmVideo(request *AbstractFilmVideoRequest) (_result *AbstractFilmVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AbstractFilmVideoResponse{}
	_body, _err := client.AbstractFilmVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AbstractFilmVideoAdvance(request *AbstractFilmVideoAdvanceRequest, runtime *util.RuntimeOptions) (_result *AbstractFilmVideoResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("videoenhan"),
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
	abstractFilmVideoReq := &AbstractFilmVideoRequest{}
	openapiutil.Convert(request, abstractFilmVideoReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.VideoUrlObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	abstractFilmVideoReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	abstractFilmVideoResp, _err := client.AbstractFilmVideoWithOptions(abstractFilmVideoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = abstractFilmVideoResp
	return _result, _err
}

func (client *Client) AdjustVideoColorWithOptions(request *AdjustVideoColorRequest, runtime *util.RuntimeOptions) (_result *AdjustVideoColorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AdjustVideoColorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AdjustVideoColor"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AdjustVideoColor(request *AdjustVideoColorRequest) (_result *AdjustVideoColorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AdjustVideoColorResponse{}
	_body, _err := client.AdjustVideoColorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AdjustVideoColorAdvance(request *AdjustVideoColorAdvanceRequest, runtime *util.RuntimeOptions) (_result *AdjustVideoColorResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("videoenhan"),
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
	adjustVideoColorReq := &AdjustVideoColorRequest{}
	openapiutil.Convert(request, adjustVideoColorReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.VideoUrlObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	adjustVideoColorReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	adjustVideoColorResp, _err := client.AdjustVideoColorWithOptions(adjustVideoColorReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = adjustVideoColorResp
	return _result, _err
}

func (client *Client) EraseVideoSubtitlesWithOptions(request *EraseVideoSubtitlesRequest, runtime *util.RuntimeOptions) (_result *EraseVideoSubtitlesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EraseVideoSubtitlesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EraseVideoSubtitles"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EraseVideoSubtitles(request *EraseVideoSubtitlesRequest) (_result *EraseVideoSubtitlesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EraseVideoSubtitlesResponse{}
	_body, _err := client.EraseVideoSubtitlesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EraseVideoSubtitlesAdvance(request *EraseVideoSubtitlesAdvanceRequest, runtime *util.RuntimeOptions) (_result *EraseVideoSubtitlesResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("videoenhan"),
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
	eraseVideoSubtitlesReq := &EraseVideoSubtitlesRequest{}
	openapiutil.Convert(request, eraseVideoSubtitlesReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.VideoUrlObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	eraseVideoSubtitlesReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	eraseVideoSubtitlesResp, _err := client.EraseVideoSubtitlesWithOptions(eraseVideoSubtitlesReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = eraseVideoSubtitlesResp
	return _result, _err
}

func (client *Client) ChangeVideoSizeWithOptions(request *ChangeVideoSizeRequest, runtime *util.RuntimeOptions) (_result *ChangeVideoSizeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ChangeVideoSizeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ChangeVideoSize"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeVideoSize(request *ChangeVideoSizeRequest) (_result *ChangeVideoSizeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeVideoSizeResponse{}
	_body, _err := client.ChangeVideoSizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeVideoSizeAdvance(request *ChangeVideoSizeAdvanceRequest, runtime *util.RuntimeOptions) (_result *ChangeVideoSizeResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("videoenhan"),
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
	changeVideoSizeReq := &ChangeVideoSizeRequest{}
	openapiutil.Convert(request, changeVideoSizeReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.VideoUrlObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	changeVideoSizeReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	changeVideoSizeResp, _err := client.ChangeVideoSizeWithOptions(changeVideoSizeReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = changeVideoSizeResp
	return _result, _err
}

func (client *Client) MergeVideoFaceWithOptions(request *MergeVideoFaceRequest, runtime *util.RuntimeOptions) (_result *MergeVideoFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MergeVideoFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MergeVideoFace"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MergeVideoFace(request *MergeVideoFaceRequest) (_result *MergeVideoFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MergeVideoFaceResponse{}
	_body, _err := client.MergeVideoFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MergeVideoFaceAdvance(request *MergeVideoFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *MergeVideoFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("videoenhan"),
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
	mergeVideoFaceReq := &MergeVideoFaceRequest{}
	openapiutil.Convert(request, mergeVideoFaceReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.VideoURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	mergeVideoFaceReq.VideoURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	mergeVideoFaceResp, _err := client.MergeVideoFaceWithOptions(mergeVideoFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = mergeVideoFaceResp
	return _result, _err
}

func (client *Client) SuperResolveVideoWithOptions(request *SuperResolveVideoRequest, runtime *util.RuntimeOptions) (_result *SuperResolveVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SuperResolveVideoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SuperResolveVideo"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SuperResolveVideo(request *SuperResolveVideoRequest) (_result *SuperResolveVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SuperResolveVideoResponse{}
	_body, _err := client.SuperResolveVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SuperResolveVideoAdvance(request *SuperResolveVideoAdvanceRequest, runtime *util.RuntimeOptions) (_result *SuperResolveVideoResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("videoenhan"),
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
	superResolveVideoReq := &SuperResolveVideoRequest{}
	openapiutil.Convert(request, superResolveVideoReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.VideoUrlObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	superResolveVideoReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	superResolveVideoResp, _err := client.SuperResolveVideoWithOptions(superResolveVideoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = superResolveVideoResp
	return _result, _err
}
