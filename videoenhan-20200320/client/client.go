// This file is auto-generated, don't edit it. Thanks.
package client

import (
	number "github.com/alibabacloud-go/darabonba-number/client"
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

type AbstractEcommerceVideoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 480
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/AbstractEcommerceVideo/AbstractEcommerceVideo1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	// example:
	//
	// 480
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s AbstractEcommerceVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s AbstractEcommerceVideoRequest) GoString() string {
	return s.String()
}

func (s *AbstractEcommerceVideoRequest) SetDuration(v float32) *AbstractEcommerceVideoRequest {
	s.Duration = &v
	return s
}

func (s *AbstractEcommerceVideoRequest) SetHeight(v int32) *AbstractEcommerceVideoRequest {
	s.Height = &v
	return s
}

func (s *AbstractEcommerceVideoRequest) SetVideoUrl(v string) *AbstractEcommerceVideoRequest {
	s.VideoUrl = &v
	return s
}

func (s *AbstractEcommerceVideoRequest) SetWidth(v int32) *AbstractEcommerceVideoRequest {
	s.Width = &v
	return s
}

type AbstractEcommerceVideoAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 480
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/AbstractEcommerceVideo/AbstractEcommerceVideo1.mp4
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	// example:
	//
	// 480
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s AbstractEcommerceVideoAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AbstractEcommerceVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AbstractEcommerceVideoAdvanceRequest) SetDuration(v float32) *AbstractEcommerceVideoAdvanceRequest {
	s.Duration = &v
	return s
}

func (s *AbstractEcommerceVideoAdvanceRequest) SetHeight(v int32) *AbstractEcommerceVideoAdvanceRequest {
	s.Height = &v
	return s
}

func (s *AbstractEcommerceVideoAdvanceRequest) SetVideoUrlObject(v io.Reader) *AbstractEcommerceVideoAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *AbstractEcommerceVideoAdvanceRequest) SetWidth(v int32) *AbstractEcommerceVideoAdvanceRequest {
	s.Width = &v
	return s
}

type AbstractEcommerceVideoResponseBody struct {
	Data    *AbstractEcommerceVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 50B33B81-CCB8-42BC-8A73-AC838618936E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AbstractEcommerceVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbstractEcommerceVideoResponseBody) GoString() string {
	return s.String()
}

func (s *AbstractEcommerceVideoResponseBody) SetData(v *AbstractEcommerceVideoResponseBodyData) *AbstractEcommerceVideoResponseBody {
	s.Data = v
	return s
}

func (s *AbstractEcommerceVideoResponseBody) SetMessage(v string) *AbstractEcommerceVideoResponseBody {
	s.Message = &v
	return s
}

func (s *AbstractEcommerceVideoResponseBody) SetRequestId(v string) *AbstractEcommerceVideoResponseBody {
	s.RequestId = &v
	return s
}

type AbstractEcommerceVideoResponseBodyData struct {
	// example:
	//
	// http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/shop-video-abs/2020-03-20-19/YVgDynxB.jpg?Expires=1584707249&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=e5Q3O%2ByA6H7UhYJeMZxz4p70de****
	VideoCoverUrl *string `json:"VideoCoverUrl,omitempty" xml:"VideoCoverUrl,omitempty"`
	// example:
	//
	// http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/shop-video-abs/2020-03-20-19/YVgDynxB.mp4?Expires=1584707249&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=KErufmbHvTUYYLRj6i42wY7Tew****
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbstractEcommerceVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AbstractEcommerceVideoResponse) SetStatusCode(v int32) *AbstractEcommerceVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *AbstractEcommerceVideoResponse) SetBody(v *AbstractEcommerceVideoResponseBody) *AbstractEcommerceVideoResponse {
	s.Body = v
	return s
}

type AbstractFilmVideoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/AbstractFilmVideo/AbstractFilmVideo1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s AbstractFilmVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s AbstractFilmVideoRequest) GoString() string {
	return s.String()
}

func (s *AbstractFilmVideoRequest) SetLength(v int32) *AbstractFilmVideoRequest {
	s.Length = &v
	return s
}

func (s *AbstractFilmVideoRequest) SetVideoUrl(v string) *AbstractFilmVideoRequest {
	s.VideoUrl = &v
	return s
}

type AbstractFilmVideoAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/AbstractFilmVideo/AbstractFilmVideo1.mp4
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s AbstractFilmVideoAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AbstractFilmVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AbstractFilmVideoAdvanceRequest) SetLength(v int32) *AbstractFilmVideoAdvanceRequest {
	s.Length = &v
	return s
}

func (s *AbstractFilmVideoAdvanceRequest) SetVideoUrlObject(v io.Reader) *AbstractFilmVideoAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

type AbstractFilmVideoResponseBody struct {
	Data    *AbstractFilmVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9A5B872B-1BF0-4D84-90DA-A2EE1F072B82
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AbstractFilmVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbstractFilmVideoResponseBody) GoString() string {
	return s.String()
}

func (s *AbstractFilmVideoResponseBody) SetData(v *AbstractFilmVideoResponseBodyData) *AbstractFilmVideoResponseBody {
	s.Data = v
	return s
}

func (s *AbstractFilmVideoResponseBody) SetMessage(v string) *AbstractFilmVideoResponseBody {
	s.Message = &v
	return s
}

func (s *AbstractFilmVideoResponseBody) SetRequestId(v string) *AbstractFilmVideoResponseBody {
	s.RequestId = &v
	return s
}

type AbstractFilmVideoResponseBodyData struct {
	// example:
	//
	// http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/film-summary/EA61D012-5F89-4102-931E-419158BE2ADA_gb27k00.mp4?Expires=1584707613&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=weTexlBR1wmQlAhuU2JXaE7AyJ****
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbstractFilmVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AbstractFilmVideoResponse) SetStatusCode(v int32) *AbstractFilmVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *AbstractFilmVideoResponse) SetBody(v *AbstractFilmVideoResponseBody) *AbstractFilmVideoResponse {
	s.Body = v
	return s
}

type AddFaceVideoTemplateRequest struct {
	VideoScene *string `json:"VideoScene,omitempty" xml:"VideoScene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://invi-label.oss-cn-shanghai.aliyuncs.com/labl/temp/faceswap/test_for_api/xxxx.mp4
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s AddFaceVideoTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceVideoTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddFaceVideoTemplateRequest) SetVideoScene(v string) *AddFaceVideoTemplateRequest {
	s.VideoScene = &v
	return s
}

func (s *AddFaceVideoTemplateRequest) SetVideoURL(v string) *AddFaceVideoTemplateRequest {
	s.VideoURL = &v
	return s
}

type AddFaceVideoTemplateAdvanceRequest struct {
	VideoScene *string `json:"VideoScene,omitempty" xml:"VideoScene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://invi-label.oss-cn-shanghai.aliyuncs.com/labl/temp/faceswap/test_for_api/xxxx.mp4
	VideoURLObject io.Reader `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s AddFaceVideoTemplateAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceVideoTemplateAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AddFaceVideoTemplateAdvanceRequest) SetVideoScene(v string) *AddFaceVideoTemplateAdvanceRequest {
	s.VideoScene = &v
	return s
}

func (s *AddFaceVideoTemplateAdvanceRequest) SetVideoURLObject(v io.Reader) *AddFaceVideoTemplateAdvanceRequest {
	s.VideoURLObject = v
	return s
}

type AddFaceVideoTemplateResponseBody struct {
	Date    *AddFaceVideoTemplateResponseBodyDate `json:"Date,omitempty" xml:"Date,omitempty" type:"Struct"`
	Message *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F61182AE-515B-5B0A-A877-1C9AE908FF15
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddFaceVideoTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceVideoTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceVideoTemplateResponseBody) SetDate(v *AddFaceVideoTemplateResponseBodyDate) *AddFaceVideoTemplateResponseBody {
	s.Date = v
	return s
}

func (s *AddFaceVideoTemplateResponseBody) SetMessage(v string) *AddFaceVideoTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *AddFaceVideoTemplateResponseBody) SetRequestId(v string) *AddFaceVideoTemplateResponseBody {
	s.RequestId = &v
	return s
}

type AddFaceVideoTemplateResponseBodyDate struct {
	FaceInfos []*AddFaceVideoTemplateResponseBodyDateFaceInfos `json:"FaceInfos,omitempty" xml:"FaceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 6cd509ea-54fa-4730-8e9d-c94cadcda048
	TemplateId  *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TransResult *string `json:"TransResult,omitempty" xml:"TransResult,omitempty"`
}

func (s AddFaceVideoTemplateResponseBodyDate) String() string {
	return tea.Prettify(s)
}

func (s AddFaceVideoTemplateResponseBodyDate) GoString() string {
	return s.String()
}

func (s *AddFaceVideoTemplateResponseBodyDate) SetFaceInfos(v []*AddFaceVideoTemplateResponseBodyDateFaceInfos) *AddFaceVideoTemplateResponseBodyDate {
	s.FaceInfos = v
	return s
}

func (s *AddFaceVideoTemplateResponseBodyDate) SetTemplateId(v string) *AddFaceVideoTemplateResponseBodyDate {
	s.TemplateId = &v
	return s
}

func (s *AddFaceVideoTemplateResponseBodyDate) SetTransResult(v string) *AddFaceVideoTemplateResponseBodyDate {
	s.TransResult = &v
	return s
}

type AddFaceVideoTemplateResponseBodyDateFaceInfos struct {
	TemplateFaceID  *string `json:"TemplateFaceID,omitempty" xml:"TemplateFaceID,omitempty"`
	TemplateFaceURL *string `json:"TemplateFaceURL,omitempty" xml:"TemplateFaceURL,omitempty"`
}

func (s AddFaceVideoTemplateResponseBodyDateFaceInfos) String() string {
	return tea.Prettify(s)
}

func (s AddFaceVideoTemplateResponseBodyDateFaceInfos) GoString() string {
	return s.String()
}

func (s *AddFaceVideoTemplateResponseBodyDateFaceInfos) SetTemplateFaceID(v string) *AddFaceVideoTemplateResponseBodyDateFaceInfos {
	s.TemplateFaceID = &v
	return s
}

func (s *AddFaceVideoTemplateResponseBodyDateFaceInfos) SetTemplateFaceURL(v string) *AddFaceVideoTemplateResponseBodyDateFaceInfos {
	s.TemplateFaceURL = &v
	return s
}

type AddFaceVideoTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFaceVideoTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFaceVideoTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFaceVideoTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddFaceVideoTemplateResponse) SetHeaders(v map[string]*string) *AddFaceVideoTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddFaceVideoTemplateResponse) SetStatusCode(v int32) *AddFaceVideoTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceVideoTemplateResponse) SetBody(v *AddFaceVideoTemplateResponseBody) *AddFaceVideoTemplateResponse {
	s.Body = v
	return s
}

type AdjustVideoColorRequest struct {
	// example:
	//
	// LogC
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// 20
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// example:
	//
	// h264
	VideoCodec *string `json:"VideoCodec,omitempty" xml:"VideoCodec,omitempty"`
	// example:
	//
	// mp4
	VideoFormat *string `json:"VideoFormat,omitempty" xml:"VideoFormat,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/AdjustVideoColor/AdjustVideoColor1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s AdjustVideoColorRequest) String() string {
	return tea.Prettify(s)
}

func (s AdjustVideoColorRequest) GoString() string {
	return s.String()
}

func (s *AdjustVideoColorRequest) SetMode(v string) *AdjustVideoColorRequest {
	s.Mode = &v
	return s
}

func (s *AdjustVideoColorRequest) SetVideoBitrate(v int64) *AdjustVideoColorRequest {
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

func (s *AdjustVideoColorRequest) SetVideoUrl(v string) *AdjustVideoColorRequest {
	s.VideoUrl = &v
	return s
}

type AdjustVideoColorAdvanceRequest struct {
	// example:
	//
	// LogC
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// 20
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// example:
	//
	// h264
	VideoCodec *string `json:"VideoCodec,omitempty" xml:"VideoCodec,omitempty"`
	// example:
	//
	// mp4
	VideoFormat *string `json:"VideoFormat,omitempty" xml:"VideoFormat,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/AdjustVideoColor/AdjustVideoColor1.mp4
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s AdjustVideoColorAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AdjustVideoColorAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AdjustVideoColorAdvanceRequest) SetMode(v string) *AdjustVideoColorAdvanceRequest {
	s.Mode = &v
	return s
}

func (s *AdjustVideoColorAdvanceRequest) SetVideoBitrate(v int64) *AdjustVideoColorAdvanceRequest {
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

func (s *AdjustVideoColorAdvanceRequest) SetVideoUrlObject(v io.Reader) *AdjustVideoColorAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

type AdjustVideoColorResponseBody struct {
	Data    *AdjustVideoColorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C4EB5E0B-0718-42CC-9B2C-1FB149256874
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AdjustVideoColorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AdjustVideoColorResponseBody) GoString() string {
	return s.String()
}

func (s *AdjustVideoColorResponseBody) SetData(v *AdjustVideoColorResponseBodyData) *AdjustVideoColorResponseBody {
	s.Data = v
	return s
}

func (s *AdjustVideoColorResponseBody) SetMessage(v string) *AdjustVideoColorResponseBody {
	s.Message = &v
	return s
}

func (s *AdjustVideoColorResponseBody) SetRequestId(v string) *AdjustVideoColorResponseBody {
	s.RequestId = &v
	return s
}

type AdjustVideoColorResponseBodyData struct {
	// example:
	//
	// http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/video-recolor/2021-01-21-07/46%3A05-test.mov?Expires=1611216966&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=icKn5gEQ6rNlSHmCi2zAf2tC0L****
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AdjustVideoColorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AdjustVideoColorResponse) SetStatusCode(v int32) *AdjustVideoColorResponse {
	s.StatusCode = &v
	return s
}

func (s *AdjustVideoColorResponse) SetBody(v *AdjustVideoColorResponseBody) *AdjustVideoColorResponse {
	s.Body = v
	return s
}

type ChangeVideoSizeRequest struct {
	// example:
	//
	// 0
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// example:
	//
	// smart
	CropType *string `json:"CropType,omitempty" xml:"CropType,omitempty"`
	// example:
	//
	// image
	FillType *string `json:"FillType,omitempty" xml:"FillType,omitempty"`
	// example:
	//
	// 0
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 0
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
	// example:
	//
	// 0.5
	Tightness *float32 `json:"Tightness,omitempty" xml:"Tightness,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/ChangeVideoSize/ChangeVideoSize1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 600
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ChangeVideoSizeRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeVideoSizeRequest) GoString() string {
	return s.String()
}

func (s *ChangeVideoSizeRequest) SetB(v int32) *ChangeVideoSizeRequest {
	s.B = &v
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

func (s *ChangeVideoSizeRequest) SetG(v int32) *ChangeVideoSizeRequest {
	s.G = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetHeight(v int32) *ChangeVideoSizeRequest {
	s.Height = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetR(v int32) *ChangeVideoSizeRequest {
	s.R = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetTightness(v float32) *ChangeVideoSizeRequest {
	s.Tightness = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetVideoUrl(v string) *ChangeVideoSizeRequest {
	s.VideoUrl = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetWidth(v int32) *ChangeVideoSizeRequest {
	s.Width = &v
	return s
}

type ChangeVideoSizeAdvanceRequest struct {
	// example:
	//
	// 0
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// example:
	//
	// smart
	CropType *string `json:"CropType,omitempty" xml:"CropType,omitempty"`
	// example:
	//
	// image
	FillType *string `json:"FillType,omitempty" xml:"FillType,omitempty"`
	// example:
	//
	// 0
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 0
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
	// example:
	//
	// 0.5
	Tightness *float32 `json:"Tightness,omitempty" xml:"Tightness,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/ChangeVideoSize/ChangeVideoSize1.mp4
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 600
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ChangeVideoSizeAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeVideoSizeAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ChangeVideoSizeAdvanceRequest) SetB(v int32) *ChangeVideoSizeAdvanceRequest {
	s.B = &v
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

func (s *ChangeVideoSizeAdvanceRequest) SetG(v int32) *ChangeVideoSizeAdvanceRequest {
	s.G = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetHeight(v int32) *ChangeVideoSizeAdvanceRequest {
	s.Height = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetR(v int32) *ChangeVideoSizeAdvanceRequest {
	s.R = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetTightness(v float32) *ChangeVideoSizeAdvanceRequest {
	s.Tightness = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetVideoUrlObject(v io.Reader) *ChangeVideoSizeAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetWidth(v int32) *ChangeVideoSizeAdvanceRequest {
	s.Width = &v
	return s
}

type ChangeVideoSizeResponseBody struct {
	Data    *ChangeVideoSizeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C00C5A32-9F54-44F0-9778-0968DD9BF22A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeVideoSizeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeVideoSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeVideoSizeResponseBody) SetData(v *ChangeVideoSizeResponseBodyData) *ChangeVideoSizeResponseBody {
	s.Data = v
	return s
}

func (s *ChangeVideoSizeResponseBody) SetMessage(v string) *ChangeVideoSizeResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeVideoSizeResponseBody) SetRequestId(v string) *ChangeVideoSizeResponseBody {
	s.RequestId = &v
	return s
}

type ChangeVideoSizeResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-vd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/video-crop/2020-07-24-20/ZTZslWcU.jpg?Expires=1595597077&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=5cq1NNBEzS64U7RTXRBGlo7WPy****
	VideoCoverUrl *string `json:"VideoCoverUrl,omitempty" xml:"VideoCoverUrl,omitempty"`
	// example:
	//
	// http://vibktprfx-prod-prod-aic-vd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/video-crop/2020-07-24-20/ZTZslWcU.mp4?Expires=1595597077&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=ZyvD9AXCT2IUFkVJngQdbXMwX6****
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeVideoSizeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ChangeVideoSizeResponse) SetStatusCode(v int32) *ChangeVideoSizeResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeVideoSizeResponse) SetBody(v *ChangeVideoSizeResponseBody) *ChangeVideoSizeResponse {
	s.Body = v
	return s
}

type ConvertHdrVideoRequest struct {
	// example:
	//
	// 30
	Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// HLG
	HDRFormat *string `json:"HDRFormat,omitempty" xml:"HDRFormat,omitempty"`
	// example:
	//
	// 1000
	MaxIlluminance *int32 `json:"MaxIlluminance,omitempty" xml:"MaxIlluminance,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://invi-label.oss-cn-shanghai.aliyuncs.com/label/temp/faceswap/test_for_api/xxxx.mp4
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s ConvertHdrVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertHdrVideoRequest) GoString() string {
	return s.String()
}

func (s *ConvertHdrVideoRequest) SetBitrate(v int32) *ConvertHdrVideoRequest {
	s.Bitrate = &v
	return s
}

func (s *ConvertHdrVideoRequest) SetHDRFormat(v string) *ConvertHdrVideoRequest {
	s.HDRFormat = &v
	return s
}

func (s *ConvertHdrVideoRequest) SetMaxIlluminance(v int32) *ConvertHdrVideoRequest {
	s.MaxIlluminance = &v
	return s
}

func (s *ConvertHdrVideoRequest) SetVideoURL(v string) *ConvertHdrVideoRequest {
	s.VideoURL = &v
	return s
}

type ConvertHdrVideoAdvanceRequest struct {
	// example:
	//
	// 30
	Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// HLG
	HDRFormat *string `json:"HDRFormat,omitempty" xml:"HDRFormat,omitempty"`
	// example:
	//
	// 1000
	MaxIlluminance *int32 `json:"MaxIlluminance,omitempty" xml:"MaxIlluminance,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://invi-label.oss-cn-shanghai.aliyuncs.com/label/temp/faceswap/test_for_api/xxxx.mp4
	VideoURLObject io.Reader `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s ConvertHdrVideoAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertHdrVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ConvertHdrVideoAdvanceRequest) SetBitrate(v int32) *ConvertHdrVideoAdvanceRequest {
	s.Bitrate = &v
	return s
}

func (s *ConvertHdrVideoAdvanceRequest) SetHDRFormat(v string) *ConvertHdrVideoAdvanceRequest {
	s.HDRFormat = &v
	return s
}

func (s *ConvertHdrVideoAdvanceRequest) SetMaxIlluminance(v int32) *ConvertHdrVideoAdvanceRequest {
	s.MaxIlluminance = &v
	return s
}

func (s *ConvertHdrVideoAdvanceRequest) SetVideoURLObject(v io.Reader) *ConvertHdrVideoAdvanceRequest {
	s.VideoURLObject = v
	return s
}

type ConvertHdrVideoResponseBody struct {
	Data    *ConvertHdrVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E0CF495F-E806-4B9C-B204-E1230608239D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConvertHdrVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConvertHdrVideoResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertHdrVideoResponseBody) SetData(v *ConvertHdrVideoResponseBodyData) *ConvertHdrVideoResponseBody {
	s.Data = v
	return s
}

func (s *ConvertHdrVideoResponseBody) SetMessage(v string) *ConvertHdrVideoResponseBody {
	s.Message = &v
	return s
}

func (s *ConvertHdrVideoResponseBody) SetRequestId(v string) *ConvertHdrVideoResponseBody {
	s.RequestId = &v
	return s
}

type ConvertHdrVideoResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-vd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/hdr-enhance/20-12-22/HaKDdTI48i2GQGy7_20-12-22-06-42-45.mp4?Expires=1608621178&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=SWc90T0JHg5eWc64x8GmYHKsvX****
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s ConvertHdrVideoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ConvertHdrVideoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ConvertHdrVideoResponseBodyData) SetVideoURL(v string) *ConvertHdrVideoResponseBodyData {
	s.VideoURL = &v
	return s
}

type ConvertHdrVideoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertHdrVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertHdrVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s ConvertHdrVideoResponse) GoString() string {
	return s.String()
}

func (s *ConvertHdrVideoResponse) SetHeaders(v map[string]*string) *ConvertHdrVideoResponse {
	s.Headers = v
	return s
}

func (s *ConvertHdrVideoResponse) SetStatusCode(v int32) *ConvertHdrVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertHdrVideoResponse) SetBody(v *ConvertHdrVideoResponseBody) *ConvertHdrVideoResponse {
	s.Body = v
	return s
}

type DeleteFaceVideoTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3bf2418c-7adf-4002-a9d6-2f7cf1889c0d
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteFaceVideoTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceVideoTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceVideoTemplateRequest) SetTemplateId(v string) *DeleteFaceVideoTemplateRequest {
	s.TemplateId = &v
	return s
}

type DeleteFaceVideoTemplateResponseBody struct {
	// example:
	//
	// 2337D184-CD63-57B5-96A2-B60AABCB7B58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFaceVideoTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceVideoTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceVideoTemplateResponseBody) SetRequestId(v string) *DeleteFaceVideoTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFaceVideoTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceVideoTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceVideoTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceVideoTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceVideoTemplateResponse) SetHeaders(v map[string]*string) *DeleteFaceVideoTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceVideoTemplateResponse) SetStatusCode(v int32) *DeleteFaceVideoTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceVideoTemplateResponse) SetBody(v *DeleteFaceVideoTemplateResponseBody) *DeleteFaceVideoTemplateResponse {
	s.Body = v
	return s
}

type EnhancePortraitVideoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/xxx/shang/video/SD%289516100%29.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EnhancePortraitVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s EnhancePortraitVideoRequest) GoString() string {
	return s.String()
}

func (s *EnhancePortraitVideoRequest) SetVideoUrl(v string) *EnhancePortraitVideoRequest {
	s.VideoUrl = &v
	return s
}

type EnhancePortraitVideoAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/xxx/shang/video/SD%289516100%29.mp4
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EnhancePortraitVideoAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s EnhancePortraitVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *EnhancePortraitVideoAdvanceRequest) SetVideoUrlObject(v io.Reader) *EnhancePortraitVideoAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

type EnhancePortraitVideoResponseBody struct {
	Data    *EnhancePortraitVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// d21a2afa-4d52-4bca-803b-e65028146603
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnhancePortraitVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnhancePortraitVideoResponseBody) GoString() string {
	return s.String()
}

func (s *EnhancePortraitVideoResponseBody) SetData(v *EnhancePortraitVideoResponseBodyData) *EnhancePortraitVideoResponseBody {
	s.Data = v
	return s
}

func (s *EnhancePortraitVideoResponseBody) SetMessage(v string) *EnhancePortraitVideoResponseBody {
	s.Message = &v
	return s
}

func (s *EnhancePortraitVideoResponseBody) SetRequestId(v string) *EnhancePortraitVideoResponseBody {
	s.RequestId = &v
	return s
}

type EnhancePortraitVideoResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-xstream-cn-shanghai.oss-cn-shanghai.aliyuncs.com/xstream-framework/upload_result_video_2023-02-10_09.45.55.mp4?Expires=1675995564&amp;OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&amp;Signature=aIXTeM4IU4nARjy3SNA3YGhhqj****
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EnhancePortraitVideoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnhancePortraitVideoResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnhancePortraitVideoResponseBodyData) SetVideoUrl(v string) *EnhancePortraitVideoResponseBodyData {
	s.VideoUrl = &v
	return s
}

type EnhancePortraitVideoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnhancePortraitVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnhancePortraitVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s EnhancePortraitVideoResponse) GoString() string {
	return s.String()
}

func (s *EnhancePortraitVideoResponse) SetHeaders(v map[string]*string) *EnhancePortraitVideoResponse {
	s.Headers = v
	return s
}

func (s *EnhancePortraitVideoResponse) SetStatusCode(v int32) *EnhancePortraitVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *EnhancePortraitVideoResponse) SetBody(v *EnhancePortraitVideoResponseBody) *EnhancePortraitVideoResponse {
	s.Body = v
	return s
}

type EnhanceVideoQualityRequest struct {
	// example:
	//
	// 20
	Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// 50
	FrameRate *int32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	// example:
	//
	// PQ
	HDRFormat *string `json:"HDRFormat,omitempty" xml:"HDRFormat,omitempty"`
	// example:
	//
	// 600
	MaxIlluminance *int32 `json:"MaxIlluminance,omitempty" xml:"MaxIlluminance,omitempty"`
	// example:
	//
	// 200
	OutPutHeight *int32 `json:"OutPutHeight,omitempty" xml:"OutPutHeight,omitempty"`
	// example:
	//
	// 200
	OutPutWidth *int32 `json:"OutPutWidth,omitempty" xml:"OutPutWidth,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://invi-label.oss-cn-shanghai.aliyuncs.com/label/temp/faceswap/test_for_api/xxxx.mp4
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s EnhanceVideoQualityRequest) String() string {
	return tea.Prettify(s)
}

func (s EnhanceVideoQualityRequest) GoString() string {
	return s.String()
}

func (s *EnhanceVideoQualityRequest) SetBitrate(v int32) *EnhanceVideoQualityRequest {
	s.Bitrate = &v
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

func (s *EnhanceVideoQualityRequest) SetOutPutHeight(v int32) *EnhanceVideoQualityRequest {
	s.OutPutHeight = &v
	return s
}

func (s *EnhanceVideoQualityRequest) SetOutPutWidth(v int32) *EnhanceVideoQualityRequest {
	s.OutPutWidth = &v
	return s
}

func (s *EnhanceVideoQualityRequest) SetVideoURL(v string) *EnhanceVideoQualityRequest {
	s.VideoURL = &v
	return s
}

type EnhanceVideoQualityAdvanceRequest struct {
	// example:
	//
	// 20
	Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// 50
	FrameRate *int32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	// example:
	//
	// PQ
	HDRFormat *string `json:"HDRFormat,omitempty" xml:"HDRFormat,omitempty"`
	// example:
	//
	// 600
	MaxIlluminance *int32 `json:"MaxIlluminance,omitempty" xml:"MaxIlluminance,omitempty"`
	// example:
	//
	// 200
	OutPutHeight *int32 `json:"OutPutHeight,omitempty" xml:"OutPutHeight,omitempty"`
	// example:
	//
	// 200
	OutPutWidth *int32 `json:"OutPutWidth,omitempty" xml:"OutPutWidth,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://invi-label.oss-cn-shanghai.aliyuncs.com/label/temp/faceswap/test_for_api/xxxx.mp4
	VideoURLObject io.Reader `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s EnhanceVideoQualityAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s EnhanceVideoQualityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *EnhanceVideoQualityAdvanceRequest) SetBitrate(v int32) *EnhanceVideoQualityAdvanceRequest {
	s.Bitrate = &v
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

func (s *EnhanceVideoQualityAdvanceRequest) SetOutPutHeight(v int32) *EnhanceVideoQualityAdvanceRequest {
	s.OutPutHeight = &v
	return s
}

func (s *EnhanceVideoQualityAdvanceRequest) SetOutPutWidth(v int32) *EnhanceVideoQualityAdvanceRequest {
	s.OutPutWidth = &v
	return s
}

func (s *EnhanceVideoQualityAdvanceRequest) SetVideoURLObject(v io.Reader) *EnhanceVideoQualityAdvanceRequest {
	s.VideoURLObject = v
	return s
}

type EnhanceVideoQualityResponseBody struct {
	Data    *EnhanceVideoQualityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 881F39DC-C107-4817-A6D5-000BE833CC2A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnhanceVideoQualityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnhanceVideoQualityResponseBody) GoString() string {
	return s.String()
}

func (s *EnhanceVideoQualityResponseBody) SetData(v *EnhanceVideoQualityResponseBodyData) *EnhanceVideoQualityResponseBody {
	s.Data = v
	return s
}

func (s *EnhanceVideoQualityResponseBody) SetMessage(v string) *EnhanceVideoQualityResponseBody {
	s.Message = &v
	return s
}

func (s *EnhanceVideoQualityResponseBody) SetRequestId(v string) *EnhanceVideoQualityResponseBody {
	s.RequestId = &v
	return s
}

type EnhanceVideoQualityResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-vd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/uhd-enhance/20-11-20/Wwzf9z75GO5XdisS_20-11-20-07-13-48.mp4?Expires=1605858272&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=vvY0D%2Bl5eEzp%2BD7mPOWz0zMU7v****
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnhanceVideoQualityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *EnhanceVideoQualityResponse) SetStatusCode(v int32) *EnhanceVideoQualityResponse {
	s.StatusCode = &v
	return s
}

func (s *EnhanceVideoQualityResponse) SetBody(v *EnhanceVideoQualityResponseBody) *EnhanceVideoQualityResponse {
	s.Body = v
	return s
}

type EraseVideoLogoRequest struct {
	Boxes []*EraseVideoLogoRequestBoxes `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/EraseVideoLogo/EraseVideoLogo1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EraseVideoLogoRequest) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoLogoRequest) GoString() string {
	return s.String()
}

func (s *EraseVideoLogoRequest) SetBoxes(v []*EraseVideoLogoRequestBoxes) *EraseVideoLogoRequest {
	s.Boxes = v
	return s
}

func (s *EraseVideoLogoRequest) SetVideoUrl(v string) *EraseVideoLogoRequest {
	s.VideoUrl = &v
	return s
}

type EraseVideoLogoRequestBoxes struct {
	// example:
	//
	// 1.0
	H *float32 `json:"H,omitempty" xml:"H,omitempty"`
	// example:
	//
	// 1.0
	W *float32 `json:"W,omitempty" xml:"W,omitempty"`
	// example:
	//
	// 0.0
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0.0
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s EraseVideoLogoRequestBoxes) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoLogoRequestBoxes) GoString() string {
	return s.String()
}

func (s *EraseVideoLogoRequestBoxes) SetH(v float32) *EraseVideoLogoRequestBoxes {
	s.H = &v
	return s
}

func (s *EraseVideoLogoRequestBoxes) SetW(v float32) *EraseVideoLogoRequestBoxes {
	s.W = &v
	return s
}

func (s *EraseVideoLogoRequestBoxes) SetX(v float32) *EraseVideoLogoRequestBoxes {
	s.X = &v
	return s
}

func (s *EraseVideoLogoRequestBoxes) SetY(v float32) *EraseVideoLogoRequestBoxes {
	s.Y = &v
	return s
}

type EraseVideoLogoAdvanceRequest struct {
	Boxes []*EraseVideoLogoAdvanceRequestBoxes `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/EraseVideoLogo/EraseVideoLogo1.mp4
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EraseVideoLogoAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoLogoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *EraseVideoLogoAdvanceRequest) SetBoxes(v []*EraseVideoLogoAdvanceRequestBoxes) *EraseVideoLogoAdvanceRequest {
	s.Boxes = v
	return s
}

func (s *EraseVideoLogoAdvanceRequest) SetVideoUrlObject(v io.Reader) *EraseVideoLogoAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

type EraseVideoLogoAdvanceRequestBoxes struct {
	// example:
	//
	// 1.0
	H *float32 `json:"H,omitempty" xml:"H,omitempty"`
	// example:
	//
	// 1.0
	W *float32 `json:"W,omitempty" xml:"W,omitempty"`
	// example:
	//
	// 0.0
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0.0
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s EraseVideoLogoAdvanceRequestBoxes) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoLogoAdvanceRequestBoxes) GoString() string {
	return s.String()
}

func (s *EraseVideoLogoAdvanceRequestBoxes) SetH(v float32) *EraseVideoLogoAdvanceRequestBoxes {
	s.H = &v
	return s
}

func (s *EraseVideoLogoAdvanceRequestBoxes) SetW(v float32) *EraseVideoLogoAdvanceRequestBoxes {
	s.W = &v
	return s
}

func (s *EraseVideoLogoAdvanceRequestBoxes) SetX(v float32) *EraseVideoLogoAdvanceRequestBoxes {
	s.X = &v
	return s
}

func (s *EraseVideoLogoAdvanceRequestBoxes) SetY(v float32) *EraseVideoLogoAdvanceRequestBoxes {
	s.Y = &v
	return s
}

type EraseVideoLogoResponseBody struct {
	Data    *EraseVideoLogoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 95532F36-98FC-4DCD-815C-282BB26D2DA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EraseVideoLogoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoLogoResponseBody) GoString() string {
	return s.String()
}

func (s *EraseVideoLogoResponseBody) SetData(v *EraseVideoLogoResponseBodyData) *EraseVideoLogoResponseBody {
	s.Data = v
	return s
}

func (s *EraseVideoLogoResponseBody) SetMessage(v string) *EraseVideoLogoResponseBody {
	s.Message = &v
	return s
}

func (s *EraseVideoLogoResponseBody) SetRequestId(v string) *EraseVideoLogoResponseBody {
	s.RequestId = &v
	return s
}

type EraseVideoLogoResponseBodyData struct {
	// example:
	//
	// http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/video-delogo/2020-03-20-11/53%3A56-DGNUGG7AcRlAylhr.mp4?Expires=1584707036&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=7CvsX7X1rSU%2B%2FDxnw484lb3LCD****
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EraseVideoLogoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *EraseVideoLogoResponse) SetStatusCode(v int32) *EraseVideoLogoResponse {
	s.StatusCode = &v
	return s
}

func (s *EraseVideoLogoResponse) SetBody(v *EraseVideoLogoResponseBody) *EraseVideoLogoResponse {
	s.Body = v
	return s
}

type EraseVideoSubtitlesRequest struct {
	// example:
	//
	// 0.25
	BH *float32 `json:"BH,omitempty" xml:"BH,omitempty"`
	// example:
	//
	// 1
	BW *float32 `json:"BW,omitempty" xml:"BW,omitempty"`
	// example:
	//
	// 0
	BX *float32 `json:"BX,omitempty" xml:"BX,omitempty"`
	// example:
	//
	// 0.75
	BY *float32 `json:"BY,omitempty" xml:"BY,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/EraseVideoSubtitles/EraseVideoSubtitles1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EraseVideoSubtitlesRequest) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoSubtitlesRequest) GoString() string {
	return s.String()
}

func (s *EraseVideoSubtitlesRequest) SetBH(v float32) *EraseVideoSubtitlesRequest {
	s.BH = &v
	return s
}

func (s *EraseVideoSubtitlesRequest) SetBW(v float32) *EraseVideoSubtitlesRequest {
	s.BW = &v
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

func (s *EraseVideoSubtitlesRequest) SetVideoUrl(v string) *EraseVideoSubtitlesRequest {
	s.VideoUrl = &v
	return s
}

type EraseVideoSubtitlesAdvanceRequest struct {
	// example:
	//
	// 0.25
	BH *float32 `json:"BH,omitempty" xml:"BH,omitempty"`
	// example:
	//
	// 1
	BW *float32 `json:"BW,omitempty" xml:"BW,omitempty"`
	// example:
	//
	// 0
	BX *float32 `json:"BX,omitempty" xml:"BX,omitempty"`
	// example:
	//
	// 0.75
	BY *float32 `json:"BY,omitempty" xml:"BY,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/EraseVideoSubtitles/EraseVideoSubtitles1.mp4
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EraseVideoSubtitlesAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoSubtitlesAdvanceRequest) GoString() string {
	return s.String()
}

func (s *EraseVideoSubtitlesAdvanceRequest) SetBH(v float32) *EraseVideoSubtitlesAdvanceRequest {
	s.BH = &v
	return s
}

func (s *EraseVideoSubtitlesAdvanceRequest) SetBW(v float32) *EraseVideoSubtitlesAdvanceRequest {
	s.BW = &v
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

func (s *EraseVideoSubtitlesAdvanceRequest) SetVideoUrlObject(v io.Reader) *EraseVideoSubtitlesAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

type EraseVideoSubtitlesResponseBody struct {
	Data    *EraseVideoSubtitlesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CCB082BF-A6B1-4C28-9E49-562EEE7DE639
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EraseVideoSubtitlesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoSubtitlesResponseBody) GoString() string {
	return s.String()
}

func (s *EraseVideoSubtitlesResponseBody) SetData(v *EraseVideoSubtitlesResponseBodyData) *EraseVideoSubtitlesResponseBody {
	s.Data = v
	return s
}

func (s *EraseVideoSubtitlesResponseBody) SetMessage(v string) *EraseVideoSubtitlesResponseBody {
	s.Message = &v
	return s
}

func (s *EraseVideoSubtitlesResponseBody) SetRequestId(v string) *EraseVideoSubtitlesResponseBody {
	s.RequestId = &v
	return s
}

type EraseVideoSubtitlesResponseBodyData struct {
	// example:
	//
	// http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/video-desubtitle/2021-04-13-10/41%3A57-TcFd6Zug7gXwbeqs.mp4?Expires=1618312317&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=LZnGSQ8019%2Br5rcR4vKOaaT2UE****
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EraseVideoSubtitlesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *EraseVideoSubtitlesResponse) SetStatusCode(v int32) *EraseVideoSubtitlesResponse {
	s.StatusCode = &v
	return s
}

func (s *EraseVideoSubtitlesResponse) SetBody(v *EraseVideoSubtitlesResponseBody) *EraseVideoSubtitlesResponse {
	s.Body = v
	return s
}

type GenerateHumanAnimeStyleVideoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// anime
	CartoonStyle *string `json:"CartoonStyle,omitempty" xml:"CartoonStyle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test/xxx/eas/EvaluateVideoQuality/123.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GenerateHumanAnimeStyleVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanAnimeStyleVideoRequest) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleVideoRequest) SetCartoonStyle(v string) *GenerateHumanAnimeStyleVideoRequest {
	s.CartoonStyle = &v
	return s
}

func (s *GenerateHumanAnimeStyleVideoRequest) SetVideoUrl(v string) *GenerateHumanAnimeStyleVideoRequest {
	s.VideoUrl = &v
	return s
}

type GenerateHumanAnimeStyleVideoAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// anime
	CartoonStyle *string `json:"CartoonStyle,omitempty" xml:"CartoonStyle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test/xxx/eas/EvaluateVideoQuality/123.mp4
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GenerateHumanAnimeStyleVideoAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanAnimeStyleVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleVideoAdvanceRequest) SetCartoonStyle(v string) *GenerateHumanAnimeStyleVideoAdvanceRequest {
	s.CartoonStyle = &v
	return s
}

func (s *GenerateHumanAnimeStyleVideoAdvanceRequest) SetVideoUrlObject(v io.Reader) *GenerateHumanAnimeStyleVideoAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

type GenerateHumanAnimeStyleVideoResponseBody struct {
	Data    *GenerateHumanAnimeStyleVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// d21a2afa-4d52-4bca-803b-e65028146603
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateHumanAnimeStyleVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanAnimeStyleVideoResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleVideoResponseBody) SetData(v *GenerateHumanAnimeStyleVideoResponseBodyData) *GenerateHumanAnimeStyleVideoResponseBody {
	s.Data = v
	return s
}

func (s *GenerateHumanAnimeStyleVideoResponseBody) SetMessage(v string) *GenerateHumanAnimeStyleVideoResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateHumanAnimeStyleVideoResponseBody) SetRequestId(v string) *GenerateHumanAnimeStyleVideoResponseBody {
	s.RequestId = &v
	return s
}

type GenerateHumanAnimeStyleVideoResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-xstream-cn-shanghai.oss-cn-shanghai.aliyuncs.com/xstream-framework/upload_result_video_2023-02-10_09.45.55.mp4?Expires=1675995564&amp;OSSAccessKeyId=LTAI4FoLmvQ9urWXgSR****&amp;Signature=aIXTeM4IU4nARjy3SNA3YGhhqj****
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GenerateHumanAnimeStyleVideoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanAnimeStyleVideoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleVideoResponseBodyData) SetVideoUrl(v string) *GenerateHumanAnimeStyleVideoResponseBodyData {
	s.VideoUrl = &v
	return s
}

type GenerateHumanAnimeStyleVideoResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateHumanAnimeStyleVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateHumanAnimeStyleVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanAnimeStyleVideoResponse) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleVideoResponse) SetHeaders(v map[string]*string) *GenerateHumanAnimeStyleVideoResponse {
	s.Headers = v
	return s
}

func (s *GenerateHumanAnimeStyleVideoResponse) SetStatusCode(v int32) *GenerateHumanAnimeStyleVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateHumanAnimeStyleVideoResponse) SetBody(v *GenerateHumanAnimeStyleVideoResponseBody) *GenerateHumanAnimeStyleVideoResponse {
	s.Body = v
	return s
}

type GenerateVideoRequest struct {
	// example:
	//
	// 10
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// true
	DurationAdaption *bool `json:"DurationAdaption,omitempty" xml:"DurationAdaption,omitempty"`
	// 1
	//
	// This parameter is required.
	FileList []*GenerateVideoRequestFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	// example:
	//
	// 640
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// true
	Mute *bool `json:"Mute,omitempty" xml:"Mute,omitempty"`
	// example:
	//
	// true
	PuzzleEffect *bool `json:"PuzzleEffect,omitempty" xml:"PuzzleEffect,omitempty"`
	// example:
	//
	// costume
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// example:
	//
	// true
	SmartEffect *bool `json:"SmartEffect,omitempty" xml:"SmartEffect,omitempty"`
	// example:
	//
	// fast
	Style *string `json:"Style,omitempty" xml:"Style,omitempty"`
	// example:
	//
	// brush
	TransitionStyle *string `json:"TransitionStyle,omitempty" xml:"TransitionStyle,omitempty"`
	// example:
	//
	// 640
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GenerateVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoRequest) GoString() string {
	return s.String()
}

func (s *GenerateVideoRequest) SetDuration(v float32) *GenerateVideoRequest {
	s.Duration = &v
	return s
}

func (s *GenerateVideoRequest) SetDurationAdaption(v bool) *GenerateVideoRequest {
	s.DurationAdaption = &v
	return s
}

func (s *GenerateVideoRequest) SetFileList(v []*GenerateVideoRequestFileList) *GenerateVideoRequest {
	s.FileList = v
	return s
}

func (s *GenerateVideoRequest) SetHeight(v int32) *GenerateVideoRequest {
	s.Height = &v
	return s
}

func (s *GenerateVideoRequest) SetMute(v bool) *GenerateVideoRequest {
	s.Mute = &v
	return s
}

func (s *GenerateVideoRequest) SetPuzzleEffect(v bool) *GenerateVideoRequest {
	s.PuzzleEffect = &v
	return s
}

func (s *GenerateVideoRequest) SetScene(v string) *GenerateVideoRequest {
	s.Scene = &v
	return s
}

func (s *GenerateVideoRequest) SetSmartEffect(v bool) *GenerateVideoRequest {
	s.SmartEffect = &v
	return s
}

func (s *GenerateVideoRequest) SetStyle(v string) *GenerateVideoRequest {
	s.Style = &v
	return s
}

func (s *GenerateVideoRequest) SetTransitionStyle(v string) *GenerateVideoRequest {
	s.TransitionStyle = &v
	return s
}

func (s *GenerateVideoRequest) SetWidth(v int32) *GenerateVideoRequest {
	s.Width = &v
	return s
}

type GenerateVideoRequestFileList struct {
	// This parameter is required.
	//
	// example:
	//
	// 1-video1.mp4
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/GenerateVideo/1-video1.mp4
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// video
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GenerateVideoRequestFileList) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoRequestFileList) GoString() string {
	return s.String()
}

func (s *GenerateVideoRequestFileList) SetFileName(v string) *GenerateVideoRequestFileList {
	s.FileName = &v
	return s
}

func (s *GenerateVideoRequestFileList) SetFileUrl(v string) *GenerateVideoRequestFileList {
	s.FileUrl = &v
	return s
}

func (s *GenerateVideoRequestFileList) SetType(v string) *GenerateVideoRequestFileList {
	s.Type = &v
	return s
}

type GenerateVideoAdvanceRequest struct {
	// example:
	//
	// 10
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// true
	DurationAdaption *bool `json:"DurationAdaption,omitempty" xml:"DurationAdaption,omitempty"`
	// 1
	//
	// This parameter is required.
	FileList []*GenerateVideoAdvanceRequestFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	// example:
	//
	// 640
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// true
	Mute *bool `json:"Mute,omitempty" xml:"Mute,omitempty"`
	// example:
	//
	// true
	PuzzleEffect *bool `json:"PuzzleEffect,omitempty" xml:"PuzzleEffect,omitempty"`
	// example:
	//
	// costume
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// example:
	//
	// true
	SmartEffect *bool `json:"SmartEffect,omitempty" xml:"SmartEffect,omitempty"`
	// example:
	//
	// fast
	Style *string `json:"Style,omitempty" xml:"Style,omitempty"`
	// example:
	//
	// brush
	TransitionStyle *string `json:"TransitionStyle,omitempty" xml:"TransitionStyle,omitempty"`
	// example:
	//
	// 640
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GenerateVideoAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateVideoAdvanceRequest) SetDuration(v float32) *GenerateVideoAdvanceRequest {
	s.Duration = &v
	return s
}

func (s *GenerateVideoAdvanceRequest) SetDurationAdaption(v bool) *GenerateVideoAdvanceRequest {
	s.DurationAdaption = &v
	return s
}

func (s *GenerateVideoAdvanceRequest) SetFileList(v []*GenerateVideoAdvanceRequestFileList) *GenerateVideoAdvanceRequest {
	s.FileList = v
	return s
}

func (s *GenerateVideoAdvanceRequest) SetHeight(v int32) *GenerateVideoAdvanceRequest {
	s.Height = &v
	return s
}

func (s *GenerateVideoAdvanceRequest) SetMute(v bool) *GenerateVideoAdvanceRequest {
	s.Mute = &v
	return s
}

func (s *GenerateVideoAdvanceRequest) SetPuzzleEffect(v bool) *GenerateVideoAdvanceRequest {
	s.PuzzleEffect = &v
	return s
}

func (s *GenerateVideoAdvanceRequest) SetScene(v string) *GenerateVideoAdvanceRequest {
	s.Scene = &v
	return s
}

func (s *GenerateVideoAdvanceRequest) SetSmartEffect(v bool) *GenerateVideoAdvanceRequest {
	s.SmartEffect = &v
	return s
}

func (s *GenerateVideoAdvanceRequest) SetStyle(v string) *GenerateVideoAdvanceRequest {
	s.Style = &v
	return s
}

func (s *GenerateVideoAdvanceRequest) SetTransitionStyle(v string) *GenerateVideoAdvanceRequest {
	s.TransitionStyle = &v
	return s
}

func (s *GenerateVideoAdvanceRequest) SetWidth(v int32) *GenerateVideoAdvanceRequest {
	s.Width = &v
	return s
}

type GenerateVideoAdvanceRequestFileList struct {
	// This parameter is required.
	//
	// example:
	//
	// 1-video1.mp4
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/GenerateVideo/1-video1.mp4
	FileUrlObject io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// video
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GenerateVideoAdvanceRequestFileList) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoAdvanceRequestFileList) GoString() string {
	return s.String()
}

func (s *GenerateVideoAdvanceRequestFileList) SetFileName(v string) *GenerateVideoAdvanceRequestFileList {
	s.FileName = &v
	return s
}

func (s *GenerateVideoAdvanceRequestFileList) SetFileUrlObject(v io.Reader) *GenerateVideoAdvanceRequestFileList {
	s.FileUrlObject = v
	return s
}

func (s *GenerateVideoAdvanceRequestFileList) SetType(v string) *GenerateVideoAdvanceRequestFileList {
	s.Type = &v
	return s
}

type GenerateVideoResponseBody struct {
	Data    *GenerateVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7CB9B663-3EF8-4C9C-A464-FDA2B5F1E3A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateVideoResponseBody) SetData(v *GenerateVideoResponseBodyData) *GenerateVideoResponseBody {
	s.Data = v
	return s
}

func (s *GenerateVideoResponseBody) SetMessage(v string) *GenerateVideoResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateVideoResponseBody) SetRequestId(v string) *GenerateVideoResponseBody {
	s.RequestId = &v
	return s
}

type GenerateVideoResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-vd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/video-gen/2021-05-07-15/B9MGfwxu.mp4?Expires=1620372653&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=F9flL1n5GPYaae0dLl%2F8D%2Bn4j6****
	VideoCoverUrl *string `json:"VideoCoverUrl,omitempty" xml:"VideoCoverUrl,omitempty"`
	// example:
	//
	// http://vibktprfx-prod-prod-aic-vd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/video-gen/2021-05-07-15/B9MGfwxu.jpg?Expires=1620372653&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=djBvGvdJu8bd%2FC%2BVHdg1d57U%2Bu****
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GenerateVideoResponse) SetStatusCode(v int32) *GenerateVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateVideoResponse) SetBody(v *GenerateVideoResponseBody) *GenerateVideoResponse {
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
	// 43A0AEB6-45F4-4138-8E89-E1A5D63200E3
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
	// 49E2CC28-ED1D-4CC5-854D-7D0AE2B20976
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// {\\"VideoUrl\\":\\"http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/video-super-resolution/2020-03-26-02/55%3A59-XFLsuEKpPxleHl4n.mp4?Expires=1585193159&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=OqsxMwAKsFVjPO%2Brgw%2BpXHIm6v****\\"}
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

type InterpolateVideoFrameRequest struct {
	// example:
	//
	// 30
	Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// 70
	FrameRate *int32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/InterpolateVideoFrame/InterpolateVideoFrame3.mp4
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s InterpolateVideoFrameRequest) String() string {
	return tea.Prettify(s)
}

func (s InterpolateVideoFrameRequest) GoString() string {
	return s.String()
}

func (s *InterpolateVideoFrameRequest) SetBitrate(v int32) *InterpolateVideoFrameRequest {
	s.Bitrate = &v
	return s
}

func (s *InterpolateVideoFrameRequest) SetFrameRate(v int32) *InterpolateVideoFrameRequest {
	s.FrameRate = &v
	return s
}

func (s *InterpolateVideoFrameRequest) SetVideoURL(v string) *InterpolateVideoFrameRequest {
	s.VideoURL = &v
	return s
}

type InterpolateVideoFrameAdvanceRequest struct {
	// example:
	//
	// 30
	Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// 70
	FrameRate *int32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/InterpolateVideoFrame/InterpolateVideoFrame3.mp4
	VideoURLObject io.Reader `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s InterpolateVideoFrameAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s InterpolateVideoFrameAdvanceRequest) GoString() string {
	return s.String()
}

func (s *InterpolateVideoFrameAdvanceRequest) SetBitrate(v int32) *InterpolateVideoFrameAdvanceRequest {
	s.Bitrate = &v
	return s
}

func (s *InterpolateVideoFrameAdvanceRequest) SetFrameRate(v int32) *InterpolateVideoFrameAdvanceRequest {
	s.FrameRate = &v
	return s
}

func (s *InterpolateVideoFrameAdvanceRequest) SetVideoURLObject(v io.Reader) *InterpolateVideoFrameAdvanceRequest {
	s.VideoURLObject = v
	return s
}

type InterpolateVideoFrameResponseBody struct {
	Data    *InterpolateVideoFrameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7EF85B5B-FB44-4C3E-9B8F-08C6CD912CEB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InterpolateVideoFrameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InterpolateVideoFrameResponseBody) GoString() string {
	return s.String()
}

func (s *InterpolateVideoFrameResponseBody) SetData(v *InterpolateVideoFrameResponseBodyData) *InterpolateVideoFrameResponseBody {
	s.Data = v
	return s
}

func (s *InterpolateVideoFrameResponseBody) SetMessage(v string) *InterpolateVideoFrameResponseBody {
	s.Message = &v
	return s
}

func (s *InterpolateVideoFrameResponseBody) SetRequestId(v string) *InterpolateVideoFrameResponseBody {
	s.RequestId = &v
	return s
}

type InterpolateVideoFrameResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-vd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/video-interp/20-12-22/mRsPNVunG7717nne_20-12-22-07-29-51.mp4?Expires=1608624020&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=NFjSLll8E7E1tkuLPFyTpr6ULi****
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s InterpolateVideoFrameResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InterpolateVideoFrameResponseBodyData) GoString() string {
	return s.String()
}

func (s *InterpolateVideoFrameResponseBodyData) SetVideoURL(v string) *InterpolateVideoFrameResponseBodyData {
	s.VideoURL = &v
	return s
}

type InterpolateVideoFrameResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InterpolateVideoFrameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InterpolateVideoFrameResponse) String() string {
	return tea.Prettify(s)
}

func (s InterpolateVideoFrameResponse) GoString() string {
	return s.String()
}

func (s *InterpolateVideoFrameResponse) SetHeaders(v map[string]*string) *InterpolateVideoFrameResponse {
	s.Headers = v
	return s
}

func (s *InterpolateVideoFrameResponse) SetStatusCode(v int32) *InterpolateVideoFrameResponse {
	s.StatusCode = &v
	return s
}

func (s *InterpolateVideoFrameResponse) SetBody(v *InterpolateVideoFrameResponseBody) *InterpolateVideoFrameResponse {
	s.Body = v
	return s
}

type MergeVideoFaceRequest struct {
	AddWatermark *bool `json:"AddWatermark,omitempty" xml:"AddWatermark,omitempty"`
	Enhance      *bool `json:"Enhance,omitempty" xml:"Enhance,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/MergeVideoFace/MergeVideoFace-pic1.png
	ReferenceURL *string `json:"ReferenceURL,omitempty" xml:"ReferenceURL,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/MergeVideoFace/MergeVideoFace1.mp4
	VideoURL      *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
	WatermarkType *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
}

func (s MergeVideoFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s MergeVideoFaceRequest) GoString() string {
	return s.String()
}

func (s *MergeVideoFaceRequest) SetAddWatermark(v bool) *MergeVideoFaceRequest {
	s.AddWatermark = &v
	return s
}

func (s *MergeVideoFaceRequest) SetEnhance(v bool) *MergeVideoFaceRequest {
	s.Enhance = &v
	return s
}

func (s *MergeVideoFaceRequest) SetReferenceURL(v string) *MergeVideoFaceRequest {
	s.ReferenceURL = &v
	return s
}

func (s *MergeVideoFaceRequest) SetVideoURL(v string) *MergeVideoFaceRequest {
	s.VideoURL = &v
	return s
}

func (s *MergeVideoFaceRequest) SetWatermarkType(v string) *MergeVideoFaceRequest {
	s.WatermarkType = &v
	return s
}

type MergeVideoFaceAdvanceRequest struct {
	AddWatermark *bool `json:"AddWatermark,omitempty" xml:"AddWatermark,omitempty"`
	Enhance      *bool `json:"Enhance,omitempty" xml:"Enhance,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/MergeVideoFace/MergeVideoFace-pic1.png
	ReferenceURLObject io.Reader `json:"ReferenceURL,omitempty" xml:"ReferenceURL,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/MergeVideoFace/MergeVideoFace1.mp4
	VideoURLObject io.Reader `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
	WatermarkType  *string   `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
}

func (s MergeVideoFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s MergeVideoFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *MergeVideoFaceAdvanceRequest) SetAddWatermark(v bool) *MergeVideoFaceAdvanceRequest {
	s.AddWatermark = &v
	return s
}

func (s *MergeVideoFaceAdvanceRequest) SetEnhance(v bool) *MergeVideoFaceAdvanceRequest {
	s.Enhance = &v
	return s
}

func (s *MergeVideoFaceAdvanceRequest) SetReferenceURLObject(v io.Reader) *MergeVideoFaceAdvanceRequest {
	s.ReferenceURLObject = v
	return s
}

func (s *MergeVideoFaceAdvanceRequest) SetVideoURLObject(v io.Reader) *MergeVideoFaceAdvanceRequest {
	s.VideoURLObject = v
	return s
}

func (s *MergeVideoFaceAdvanceRequest) SetWatermarkType(v string) *MergeVideoFaceAdvanceRequest {
	s.WatermarkType = &v
	return s
}

type MergeVideoFaceResponseBody struct {
	Data    *MergeVideoFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DEF90E76-B62D-45EF-8835-CA3C83842B18
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MergeVideoFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MergeVideoFaceResponseBody) GoString() string {
	return s.String()
}

func (s *MergeVideoFaceResponseBody) SetData(v *MergeVideoFaceResponseBodyData) *MergeVideoFaceResponseBody {
	s.Data = v
	return s
}

func (s *MergeVideoFaceResponseBody) SetMessage(v string) *MergeVideoFaceResponseBody {
	s.Message = &v
	return s
}

func (s *MergeVideoFaceResponseBody) SetRequestId(v string) *MergeVideoFaceResponseBody {
	s.RequestId = &v
	return s
}

type MergeVideoFaceResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-gd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/video-face-fusion/A657011C-82B4-4705-A5DB-69B18B7CE89D.mp4?Expires=1606378308&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=Hl3cq5XedTGCscOSr0OGVxAS2o****
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MergeVideoFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *MergeVideoFaceResponse) SetStatusCode(v int32) *MergeVideoFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *MergeVideoFaceResponse) SetBody(v *MergeVideoFaceResponseBody) *MergeVideoFaceResponse {
	s.Body = v
	return s
}

type MergeVideoModelFaceRequest struct {
	AddWatermark *bool `json:"AddWatermark,omitempty" xml:"AddWatermark,omitempty"`
	Enhance      *bool `json:"Enhance,omitempty" xml:"Enhance,omitempty"`
	// example:
	//
	// https://invi-label.oss-cn-shanghai.aliyuncs.com/label/temp/faceswap/ref/ref.jpg
	FaceImageURL *string                                 `json:"FaceImageURL,omitempty" xml:"FaceImageURL,omitempty"`
	MergeInfos   []*MergeVideoModelFaceRequestMergeInfos `json:"MergeInfos,omitempty" xml:"MergeInfos,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 3bf2418c-7adf-4002-a9d6-2f7cf1889c0d
	TemplateId    *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	WatermarkType *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
}

func (s MergeVideoModelFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s MergeVideoModelFaceRequest) GoString() string {
	return s.String()
}

func (s *MergeVideoModelFaceRequest) SetAddWatermark(v bool) *MergeVideoModelFaceRequest {
	s.AddWatermark = &v
	return s
}

func (s *MergeVideoModelFaceRequest) SetEnhance(v bool) *MergeVideoModelFaceRequest {
	s.Enhance = &v
	return s
}

func (s *MergeVideoModelFaceRequest) SetFaceImageURL(v string) *MergeVideoModelFaceRequest {
	s.FaceImageURL = &v
	return s
}

func (s *MergeVideoModelFaceRequest) SetMergeInfos(v []*MergeVideoModelFaceRequestMergeInfos) *MergeVideoModelFaceRequest {
	s.MergeInfos = v
	return s
}

func (s *MergeVideoModelFaceRequest) SetTemplateId(v string) *MergeVideoModelFaceRequest {
	s.TemplateId = &v
	return s
}

func (s *MergeVideoModelFaceRequest) SetWatermarkType(v string) *MergeVideoModelFaceRequest {
	s.WatermarkType = &v
	return s
}

type MergeVideoModelFaceRequestMergeInfos struct {
	ImageURL        *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	TemplateFaceID  *string `json:"TemplateFaceID,omitempty" xml:"TemplateFaceID,omitempty"`
	TemplateFaceURL *string `json:"TemplateFaceURL,omitempty" xml:"TemplateFaceURL,omitempty"`
}

func (s MergeVideoModelFaceRequestMergeInfos) String() string {
	return tea.Prettify(s)
}

func (s MergeVideoModelFaceRequestMergeInfos) GoString() string {
	return s.String()
}

func (s *MergeVideoModelFaceRequestMergeInfos) SetImageURL(v string) *MergeVideoModelFaceRequestMergeInfos {
	s.ImageURL = &v
	return s
}

func (s *MergeVideoModelFaceRequestMergeInfos) SetTemplateFaceID(v string) *MergeVideoModelFaceRequestMergeInfos {
	s.TemplateFaceID = &v
	return s
}

func (s *MergeVideoModelFaceRequestMergeInfos) SetTemplateFaceURL(v string) *MergeVideoModelFaceRequestMergeInfos {
	s.TemplateFaceURL = &v
	return s
}

type MergeVideoModelFaceAdvanceRequest struct {
	AddWatermark *bool `json:"AddWatermark,omitempty" xml:"AddWatermark,omitempty"`
	Enhance      *bool `json:"Enhance,omitempty" xml:"Enhance,omitempty"`
	// example:
	//
	// https://invi-label.oss-cn-shanghai.aliyuncs.com/label/temp/faceswap/ref/ref.jpg
	FaceImageURLObject io.Reader                                      `json:"FaceImageURL,omitempty" xml:"FaceImageURL,omitempty"`
	MergeInfos         []*MergeVideoModelFaceAdvanceRequestMergeInfos `json:"MergeInfos,omitempty" xml:"MergeInfos,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 3bf2418c-7adf-4002-a9d6-2f7cf1889c0d
	TemplateId    *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	WatermarkType *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
}

func (s MergeVideoModelFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s MergeVideoModelFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *MergeVideoModelFaceAdvanceRequest) SetAddWatermark(v bool) *MergeVideoModelFaceAdvanceRequest {
	s.AddWatermark = &v
	return s
}

func (s *MergeVideoModelFaceAdvanceRequest) SetEnhance(v bool) *MergeVideoModelFaceAdvanceRequest {
	s.Enhance = &v
	return s
}

func (s *MergeVideoModelFaceAdvanceRequest) SetFaceImageURLObject(v io.Reader) *MergeVideoModelFaceAdvanceRequest {
	s.FaceImageURLObject = v
	return s
}

func (s *MergeVideoModelFaceAdvanceRequest) SetMergeInfos(v []*MergeVideoModelFaceAdvanceRequestMergeInfos) *MergeVideoModelFaceAdvanceRequest {
	s.MergeInfos = v
	return s
}

func (s *MergeVideoModelFaceAdvanceRequest) SetTemplateId(v string) *MergeVideoModelFaceAdvanceRequest {
	s.TemplateId = &v
	return s
}

func (s *MergeVideoModelFaceAdvanceRequest) SetWatermarkType(v string) *MergeVideoModelFaceAdvanceRequest {
	s.WatermarkType = &v
	return s
}

type MergeVideoModelFaceAdvanceRequestMergeInfos struct {
	ImageURL        *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	TemplateFaceID  *string `json:"TemplateFaceID,omitempty" xml:"TemplateFaceID,omitempty"`
	TemplateFaceURL *string `json:"TemplateFaceURL,omitempty" xml:"TemplateFaceURL,omitempty"`
}

func (s MergeVideoModelFaceAdvanceRequestMergeInfos) String() string {
	return tea.Prettify(s)
}

func (s MergeVideoModelFaceAdvanceRequestMergeInfos) GoString() string {
	return s.String()
}

func (s *MergeVideoModelFaceAdvanceRequestMergeInfos) SetImageURL(v string) *MergeVideoModelFaceAdvanceRequestMergeInfos {
	s.ImageURL = &v
	return s
}

func (s *MergeVideoModelFaceAdvanceRequestMergeInfos) SetTemplateFaceID(v string) *MergeVideoModelFaceAdvanceRequestMergeInfos {
	s.TemplateFaceID = &v
	return s
}

func (s *MergeVideoModelFaceAdvanceRequestMergeInfos) SetTemplateFaceURL(v string) *MergeVideoModelFaceAdvanceRequestMergeInfos {
	s.TemplateFaceURL = &v
	return s
}

type MergeVideoModelFaceResponseBody struct {
	Data    *MergeVideoModelFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 334F180F-3B50-51CB-B4CB-9A86A542D3BC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MergeVideoModelFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MergeVideoModelFaceResponseBody) GoString() string {
	return s.String()
}

func (s *MergeVideoModelFaceResponseBody) SetData(v *MergeVideoModelFaceResponseBodyData) *MergeVideoModelFaceResponseBody {
	s.Data = v
	return s
}

func (s *MergeVideoModelFaceResponseBody) SetMessage(v string) *MergeVideoModelFaceResponseBody {
	s.Message = &v
	return s
}

func (s *MergeVideoModelFaceResponseBody) SetRequestId(v string) *MergeVideoModelFaceResponseBody {
	s.RequestId = &v
	return s
}

type MergeVideoModelFaceResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-gd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/video-face-fusion/334F180F-3B50-51CB-B4CB-9A86A542D3BC-5716-20210906-074905.mp4?Expires=1630916420&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=OEtNqVYxXRwkyO4BrsYVJ8q5bx****
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s MergeVideoModelFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MergeVideoModelFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *MergeVideoModelFaceResponseBodyData) SetVideoURL(v string) *MergeVideoModelFaceResponseBodyData {
	s.VideoURL = &v
	return s
}

type MergeVideoModelFaceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MergeVideoModelFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MergeVideoModelFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s MergeVideoModelFaceResponse) GoString() string {
	return s.String()
}

func (s *MergeVideoModelFaceResponse) SetHeaders(v map[string]*string) *MergeVideoModelFaceResponse {
	s.Headers = v
	return s
}

func (s *MergeVideoModelFaceResponse) SetStatusCode(v int32) *MergeVideoModelFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *MergeVideoModelFaceResponse) SetBody(v *MergeVideoModelFaceResponseBody) *MergeVideoModelFaceResponse {
	s.Body = v
	return s
}

type QueryFaceVideoTemplateRequest struct {
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 3bf2418c-7adf-4002-a9d6-2f7cf1889c0d
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryFaceVideoTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceVideoTemplateRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceVideoTemplateRequest) SetPageNo(v int64) *QueryFaceVideoTemplateRequest {
	s.PageNo = &v
	return s
}

func (s *QueryFaceVideoTemplateRequest) SetPageSize(v int64) *QueryFaceVideoTemplateRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceVideoTemplateRequest) SetTemplateId(v string) *QueryFaceVideoTemplateRequest {
	s.TemplateId = &v
	return s
}

type QueryFaceVideoTemplateResponseBody struct {
	Data *QueryFaceVideoTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// A06E3D21-890D-500B-97DA-D8B99D2DDDC4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryFaceVideoTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceVideoTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceVideoTemplateResponseBody) SetData(v *QueryFaceVideoTemplateResponseBodyData) *QueryFaceVideoTemplateResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceVideoTemplateResponseBody) SetRequestId(v string) *QueryFaceVideoTemplateResponseBody {
	s.RequestId = &v
	return s
}

type QueryFaceVideoTemplateResponseBodyData struct {
	Elements []*QueryFaceVideoTemplateResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	Total    *int64                                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryFaceVideoTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceVideoTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceVideoTemplateResponseBodyData) SetElements(v []*QueryFaceVideoTemplateResponseBodyDataElements) *QueryFaceVideoTemplateResponseBodyData {
	s.Elements = v
	return s
}

func (s *QueryFaceVideoTemplateResponseBodyData) SetTotal(v int64) *QueryFaceVideoTemplateResponseBodyData {
	s.Total = &v
	return s
}

type QueryFaceVideoTemplateResponseBodyDataElements struct {
	// example:
	//
	// 2021-09-06 15:17:19
	CreateTime *string                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FaceInfos  []*QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos `json:"FaceInfos,omitempty" xml:"FaceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 3bf2418c-7adf-4002-a9d6-2f7cf1889c0d
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// http://vibktprfx-prod-prod-aic-gd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/image-face-fusion/OriginUrl/user/2021-09-06/e4d6ecf6-8dc8-4dac-acb5-56a737ccbc06?Expires=1630914551&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=kBcLm66y7%2FZ3eIMgwXJg1zNP7k****
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// example:
	//
	// 2021-09-06 15:18:15
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// zhangsan
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFaceVideoTemplateResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceVideoTemplateResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *QueryFaceVideoTemplateResponseBodyDataElements) SetCreateTime(v string) *QueryFaceVideoTemplateResponseBodyDataElements {
	s.CreateTime = &v
	return s
}

func (s *QueryFaceVideoTemplateResponseBodyDataElements) SetFaceInfos(v []*QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos) *QueryFaceVideoTemplateResponseBodyDataElements {
	s.FaceInfos = v
	return s
}

func (s *QueryFaceVideoTemplateResponseBodyDataElements) SetTemplateId(v string) *QueryFaceVideoTemplateResponseBodyDataElements {
	s.TemplateId = &v
	return s
}

func (s *QueryFaceVideoTemplateResponseBodyDataElements) SetTemplateURL(v string) *QueryFaceVideoTemplateResponseBodyDataElements {
	s.TemplateURL = &v
	return s
}

func (s *QueryFaceVideoTemplateResponseBodyDataElements) SetUpdateTime(v string) *QueryFaceVideoTemplateResponseBodyDataElements {
	s.UpdateTime = &v
	return s
}

func (s *QueryFaceVideoTemplateResponseBodyDataElements) SetUserId(v string) *QueryFaceVideoTemplateResponseBodyDataElements {
	s.UserId = &v
	return s
}

type QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos struct {
	TemplateFaceID  *string `json:"TemplateFaceID,omitempty" xml:"TemplateFaceID,omitempty"`
	TemplateFaceURL *string `json:"TemplateFaceURL,omitempty" xml:"TemplateFaceURL,omitempty"`
}

func (s QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos) GoString() string {
	return s.String()
}

func (s *QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos) SetTemplateFaceID(v string) *QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos {
	s.TemplateFaceID = &v
	return s
}

func (s *QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos) SetTemplateFaceURL(v string) *QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos {
	s.TemplateFaceURL = &v
	return s
}

type QueryFaceVideoTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceVideoTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceVideoTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceVideoTemplateResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceVideoTemplateResponse) SetHeaders(v map[string]*string) *QueryFaceVideoTemplateResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceVideoTemplateResponse) SetStatusCode(v int32) *QueryFaceVideoTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceVideoTemplateResponse) SetBody(v *QueryFaceVideoTemplateResponseBody) *QueryFaceVideoTemplateResponse {
	s.Body = v
	return s
}

type ReduceVideoNoiseRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/xxx/shang/video/SD%289516100%29.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s ReduceVideoNoiseRequest) String() string {
	return tea.Prettify(s)
}

func (s ReduceVideoNoiseRequest) GoString() string {
	return s.String()
}

func (s *ReduceVideoNoiseRequest) SetVideoUrl(v string) *ReduceVideoNoiseRequest {
	s.VideoUrl = &v
	return s
}

type ReduceVideoNoiseAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/xxx/shang/video/SD%289516100%29.mp4
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s ReduceVideoNoiseAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReduceVideoNoiseAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ReduceVideoNoiseAdvanceRequest) SetVideoUrlObject(v io.Reader) *ReduceVideoNoiseAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

type ReduceVideoNoiseResponseBody struct {
	Data    *ReduceVideoNoiseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// d21a2afa-4d52-4bca-803b-e65028146603
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReduceVideoNoiseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReduceVideoNoiseResponseBody) GoString() string {
	return s.String()
}

func (s *ReduceVideoNoiseResponseBody) SetData(v *ReduceVideoNoiseResponseBodyData) *ReduceVideoNoiseResponseBody {
	s.Data = v
	return s
}

func (s *ReduceVideoNoiseResponseBody) SetMessage(v string) *ReduceVideoNoiseResponseBody {
	s.Message = &v
	return s
}

func (s *ReduceVideoNoiseResponseBody) SetRequestId(v string) *ReduceVideoNoiseResponseBody {
	s.RequestId = &v
	return s
}

type ReduceVideoNoiseResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-xstream-cn-shanghai.oss-cn-shanghai.aliyuncs.com/xstream-framework/upload_result_video_2023-02-10_09.45.55.mp4?Expires=1675995564&amp;OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&amp;Signature=aIXTeM4IU4nARjy3SNA3YGhhqj****
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s ReduceVideoNoiseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ReduceVideoNoiseResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReduceVideoNoiseResponseBodyData) SetVideoUrl(v string) *ReduceVideoNoiseResponseBodyData {
	s.VideoUrl = &v
	return s
}

type ReduceVideoNoiseResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReduceVideoNoiseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReduceVideoNoiseResponse) String() string {
	return tea.Prettify(s)
}

func (s ReduceVideoNoiseResponse) GoString() string {
	return s.String()
}

func (s *ReduceVideoNoiseResponse) SetHeaders(v map[string]*string) *ReduceVideoNoiseResponse {
	s.Headers = v
	return s
}

func (s *ReduceVideoNoiseResponse) SetStatusCode(v int32) *ReduceVideoNoiseResponse {
	s.StatusCode = &v
	return s
}

func (s *ReduceVideoNoiseResponse) SetBody(v *ReduceVideoNoiseResponseBody) *ReduceVideoNoiseResponse {
	s.Body = v
	return s
}

type SuperResolveVideoRequest struct {
	// example:
	//
	// 5
	BitRate *int32 `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/SuperResolveVideo/SuperResolveVideo2.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s SuperResolveVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s SuperResolveVideoRequest) GoString() string {
	return s.String()
}

func (s *SuperResolveVideoRequest) SetBitRate(v int32) *SuperResolveVideoRequest {
	s.BitRate = &v
	return s
}

func (s *SuperResolveVideoRequest) SetVideoUrl(v string) *SuperResolveVideoRequest {
	s.VideoUrl = &v
	return s
}

type SuperResolveVideoAdvanceRequest struct {
	// example:
	//
	// 5
	BitRate *int32 `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/SuperResolveVideo/SuperResolveVideo2.mp4
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s SuperResolveVideoAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SuperResolveVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SuperResolveVideoAdvanceRequest) SetBitRate(v int32) *SuperResolveVideoAdvanceRequest {
	s.BitRate = &v
	return s
}

func (s *SuperResolveVideoAdvanceRequest) SetVideoUrlObject(v io.Reader) *SuperResolveVideoAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

type SuperResolveVideoResponseBody struct {
	Data    *SuperResolveVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 186AC396-0EEC-46F1-AAA1-BF3585227427
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SuperResolveVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SuperResolveVideoResponseBody) GoString() string {
	return s.String()
}

func (s *SuperResolveVideoResponseBody) SetData(v *SuperResolveVideoResponseBodyData) *SuperResolveVideoResponseBody {
	s.Data = v
	return s
}

func (s *SuperResolveVideoResponseBody) SetMessage(v string) *SuperResolveVideoResponseBody {
	s.Message = &v
	return s
}

func (s *SuperResolveVideoResponseBody) SetRequestId(v string) *SuperResolveVideoResponseBody {
	s.RequestId = &v
	return s
}

type SuperResolveVideoResponseBodyData struct {
	// example:
	//
	// http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/video-super-resolution/2020-03-20-12/12%3A11-UlLVELFzIy5EAyEh.mp4?Expires=1584708132&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=0V8yKrCVybC4KIPtRuGKJDJaQT****
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuperResolveVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *SuperResolveVideoResponse) SetStatusCode(v int32) *SuperResolveVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *SuperResolveVideoResponse) SetBody(v *SuperResolveVideoResponseBody) *SuperResolveVideoResponse {
	s.Body = v
	return s
}

type ToneSdrVideoRequest struct {
	// example:
	//
	// 30
	Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// auto_l2
	RecolorModel *string `json:"RecolorModel,omitempty" xml:"RecolorModel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://invi-label.oss-cn-shanghai.aliyuncs.com/label/temp/faceswap/test_for_api/xxxx.mp4
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s ToneSdrVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s ToneSdrVideoRequest) GoString() string {
	return s.String()
}

func (s *ToneSdrVideoRequest) SetBitrate(v int32) *ToneSdrVideoRequest {
	s.Bitrate = &v
	return s
}

func (s *ToneSdrVideoRequest) SetRecolorModel(v string) *ToneSdrVideoRequest {
	s.RecolorModel = &v
	return s
}

func (s *ToneSdrVideoRequest) SetVideoURL(v string) *ToneSdrVideoRequest {
	s.VideoURL = &v
	return s
}

type ToneSdrVideoAdvanceRequest struct {
	// example:
	//
	// 30
	Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// auto_l2
	RecolorModel *string `json:"RecolorModel,omitempty" xml:"RecolorModel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://invi-label.oss-cn-shanghai.aliyuncs.com/label/temp/faceswap/test_for_api/xxxx.mp4
	VideoURLObject io.Reader `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s ToneSdrVideoAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ToneSdrVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ToneSdrVideoAdvanceRequest) SetBitrate(v int32) *ToneSdrVideoAdvanceRequest {
	s.Bitrate = &v
	return s
}

func (s *ToneSdrVideoAdvanceRequest) SetRecolorModel(v string) *ToneSdrVideoAdvanceRequest {
	s.RecolorModel = &v
	return s
}

func (s *ToneSdrVideoAdvanceRequest) SetVideoURLObject(v io.Reader) *ToneSdrVideoAdvanceRequest {
	s.VideoURLObject = v
	return s
}

type ToneSdrVideoResponseBody struct {
	Data    *ToneSdrVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 89B5AFF1-8A64-4F76-B391-56AD7D22DE35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ToneSdrVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ToneSdrVideoResponseBody) GoString() string {
	return s.String()
}

func (s *ToneSdrVideoResponseBody) SetData(v *ToneSdrVideoResponseBodyData) *ToneSdrVideoResponseBody {
	s.Data = v
	return s
}

func (s *ToneSdrVideoResponseBody) SetMessage(v string) *ToneSdrVideoResponseBody {
	s.Message = &v
	return s
}

func (s *ToneSdrVideoResponseBody) SetRequestId(v string) *ToneSdrVideoResponseBody {
	s.RequestId = &v
	return s
}

type ToneSdrVideoResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-vd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/sdr-color-enhance/20-12-22/SxBKgwBhlObusG20_20-12-22-07-59-45.mp4?Expires=1608625795&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=maoOZ52y7U9ZuL2KqI0IfGq8%2FR****
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s ToneSdrVideoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ToneSdrVideoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ToneSdrVideoResponseBodyData) SetVideoURL(v string) *ToneSdrVideoResponseBodyData {
	s.VideoURL = &v
	return s
}

type ToneSdrVideoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ToneSdrVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ToneSdrVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s ToneSdrVideoResponse) GoString() string {
	return s.String()
}

func (s *ToneSdrVideoResponse) SetHeaders(v map[string]*string) *ToneSdrVideoResponse {
	s.Headers = v
	return s
}

func (s *ToneSdrVideoResponse) SetStatusCode(v int32) *ToneSdrVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *ToneSdrVideoResponse) SetBody(v *ToneSdrVideoResponseBody) *ToneSdrVideoResponse {
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

// @param request - AbstractEcommerceVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbstractEcommerceVideoResponse
func (client *Client) AbstractEcommerceVideoWithOptions(request *AbstractEcommerceVideoRequest, runtime *util.RuntimeOptions) (_result *AbstractEcommerceVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.Height)) {
		body["Height"] = request.Height
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["VideoUrl"] = request.VideoUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Width)) {
		body["Width"] = request.Width
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AbstractEcommerceVideo"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AbstractEcommerceVideoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AbstractEcommerceVideoRequest
//
// @return AbstractEcommerceVideoResponse
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
		Product:  tea.String("videoenhan"),
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
	abstractEcommerceVideoReq := &AbstractEcommerceVideoRequest{}
	openapiutil.Convert(request, abstractEcommerceVideoReq)
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
		abstractEcommerceVideoReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	abstractEcommerceVideoResp, _err := client.AbstractEcommerceVideoWithOptions(abstractEcommerceVideoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = abstractEcommerceVideoResp
	return _result, _err
}

// @param request - AbstractFilmVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbstractFilmVideoResponse
func (client *Client) AbstractFilmVideoWithOptions(request *AbstractFilmVideoRequest, runtime *util.RuntimeOptions) (_result *AbstractFilmVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Length)) {
		body["Length"] = request.Length
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AbstractFilmVideo"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AbstractFilmVideoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AbstractFilmVideoRequest
//
// @return AbstractFilmVideoResponse
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
		Product:  tea.String("videoenhan"),
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
	abstractFilmVideoReq := &AbstractFilmVideoRequest{}
	openapiutil.Convert(request, abstractFilmVideoReq)
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
		abstractFilmVideoReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	abstractFilmVideoResp, _err := client.AbstractFilmVideoWithOptions(abstractFilmVideoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = abstractFilmVideoResp
	return _result, _err
}

// Summary:
//
// 视频人脸融合模板增加
//
// @param request - AddFaceVideoTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFaceVideoTemplateResponse
func (client *Client) AddFaceVideoTemplateWithOptions(request *AddFaceVideoTemplateRequest, runtime *util.RuntimeOptions) (_result *AddFaceVideoTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VideoScene)) {
		body["VideoScene"] = request.VideoScene
	}

	if !tea.BoolValue(util.IsUnset(request.VideoURL)) {
		body["VideoURL"] = request.VideoURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddFaceVideoTemplate"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddFaceVideoTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频人脸融合模板增加
//
// @param request - AddFaceVideoTemplateRequest
//
// @return AddFaceVideoTemplateResponse
func (client *Client) AddFaceVideoTemplate(request *AddFaceVideoTemplateRequest) (_result *AddFaceVideoTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddFaceVideoTemplateResponse{}
	_body, _err := client.AddFaceVideoTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFaceVideoTemplateAdvance(request *AddFaceVideoTemplateAdvanceRequest, runtime *util.RuntimeOptions) (_result *AddFaceVideoTemplateResponse, _err error) {
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
		Product:  tea.String("videoenhan"),
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
	addFaceVideoTemplateReq := &AddFaceVideoTemplateRequest{}
	openapiutil.Convert(request, addFaceVideoTemplateReq)
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
		addFaceVideoTemplateReq.VideoURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	addFaceVideoTemplateResp, _err := client.AddFaceVideoTemplateWithOptions(addFaceVideoTemplateReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = addFaceVideoTemplateResp
	return _result, _err
}

// @param request - AdjustVideoColorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AdjustVideoColorResponse
func (client *Client) AdjustVideoColorWithOptions(request *AdjustVideoColorRequest, runtime *util.RuntimeOptions) (_result *AdjustVideoColorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.VideoBitrate)) {
		body["VideoBitrate"] = request.VideoBitrate
	}

	if !tea.BoolValue(util.IsUnset(request.VideoCodec)) {
		body["VideoCodec"] = request.VideoCodec
	}

	if !tea.BoolValue(util.IsUnset(request.VideoFormat)) {
		body["VideoFormat"] = request.VideoFormat
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AdjustVideoColor"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AdjustVideoColorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AdjustVideoColorRequest
//
// @return AdjustVideoColorResponse
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
		Product:  tea.String("videoenhan"),
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
	adjustVideoColorReq := &AdjustVideoColorRequest{}
	openapiutil.Convert(request, adjustVideoColorReq)
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
		adjustVideoColorReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	adjustVideoColorResp, _err := client.AdjustVideoColorWithOptions(adjustVideoColorReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = adjustVideoColorResp
	return _result, _err
}

// @param request - ChangeVideoSizeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeVideoSizeResponse
func (client *Client) ChangeVideoSizeWithOptions(request *ChangeVideoSizeRequest, runtime *util.RuntimeOptions) (_result *ChangeVideoSizeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.B)) {
		body["B"] = request.B
	}

	if !tea.BoolValue(util.IsUnset(request.CropType)) {
		body["CropType"] = request.CropType
	}

	if !tea.BoolValue(util.IsUnset(request.FillType)) {
		body["FillType"] = request.FillType
	}

	if !tea.BoolValue(util.IsUnset(request.G)) {
		body["G"] = request.G
	}

	if !tea.BoolValue(util.IsUnset(request.Height)) {
		body["Height"] = request.Height
	}

	if !tea.BoolValue(util.IsUnset(request.R)) {
		body["R"] = request.R
	}

	if !tea.BoolValue(util.IsUnset(request.Tightness)) {
		body["Tightness"] = request.Tightness
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["VideoUrl"] = request.VideoUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Width)) {
		body["Width"] = request.Width
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeVideoSize"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeVideoSizeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ChangeVideoSizeRequest
//
// @return ChangeVideoSizeResponse
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
		Product:  tea.String("videoenhan"),
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
	changeVideoSizeReq := &ChangeVideoSizeRequest{}
	openapiutil.Convert(request, changeVideoSizeReq)
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
		changeVideoSizeReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	changeVideoSizeResp, _err := client.ChangeVideoSizeWithOptions(changeVideoSizeReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = changeVideoSizeResp
	return _result, _err
}

// @param request - ConvertHdrVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConvertHdrVideoResponse
func (client *Client) ConvertHdrVideoWithOptions(request *ConvertHdrVideoRequest, runtime *util.RuntimeOptions) (_result *ConvertHdrVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bitrate)) {
		body["Bitrate"] = request.Bitrate
	}

	if !tea.BoolValue(util.IsUnset(request.HDRFormat)) {
		body["HDRFormat"] = request.HDRFormat
	}

	if !tea.BoolValue(util.IsUnset(request.MaxIlluminance)) {
		body["MaxIlluminance"] = request.MaxIlluminance
	}

	if !tea.BoolValue(util.IsUnset(request.VideoURL)) {
		body["VideoURL"] = request.VideoURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConvertHdrVideo"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConvertHdrVideoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ConvertHdrVideoRequest
//
// @return ConvertHdrVideoResponse
func (client *Client) ConvertHdrVideo(request *ConvertHdrVideoRequest) (_result *ConvertHdrVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConvertHdrVideoResponse{}
	_body, _err := client.ConvertHdrVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConvertHdrVideoAdvance(request *ConvertHdrVideoAdvanceRequest, runtime *util.RuntimeOptions) (_result *ConvertHdrVideoResponse, _err error) {
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
		Product:  tea.String("videoenhan"),
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
	convertHdrVideoReq := &ConvertHdrVideoRequest{}
	openapiutil.Convert(request, convertHdrVideoReq)
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
		convertHdrVideoReq.VideoURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	convertHdrVideoResp, _err := client.ConvertHdrVideoWithOptions(convertHdrVideoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = convertHdrVideoResp
	return _result, _err
}

// Summary:
//
// 视频人脸融合模板删除
//
// @param request - DeleteFaceVideoTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFaceVideoTemplateResponse
func (client *Client) DeleteFaceVideoTemplateWithOptions(request *DeleteFaceVideoTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteFaceVideoTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFaceVideoTemplate"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFaceVideoTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频人脸融合模板删除
//
// @param request - DeleteFaceVideoTemplateRequest
//
// @return DeleteFaceVideoTemplateResponse
func (client *Client) DeleteFaceVideoTemplate(request *DeleteFaceVideoTemplateRequest) (_result *DeleteFaceVideoTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFaceVideoTemplateResponse{}
	_body, _err := client.DeleteFaceVideoTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 视频人像增强
//
// @param request - EnhancePortraitVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnhancePortraitVideoResponse
func (client *Client) EnhancePortraitVideoWithOptions(request *EnhancePortraitVideoRequest, runtime *util.RuntimeOptions) (_result *EnhancePortraitVideoResponse, _err error) {
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
		Action:      tea.String("EnhancePortraitVideo"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnhancePortraitVideoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频人像增强
//
// @param request - EnhancePortraitVideoRequest
//
// @return EnhancePortraitVideoResponse
func (client *Client) EnhancePortraitVideo(request *EnhancePortraitVideoRequest) (_result *EnhancePortraitVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnhancePortraitVideoResponse{}
	_body, _err := client.EnhancePortraitVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnhancePortraitVideoAdvance(request *EnhancePortraitVideoAdvanceRequest, runtime *util.RuntimeOptions) (_result *EnhancePortraitVideoResponse, _err error) {
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
		Product:  tea.String("videoenhan"),
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
	enhancePortraitVideoReq := &EnhancePortraitVideoRequest{}
	openapiutil.Convert(request, enhancePortraitVideoReq)
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
		enhancePortraitVideoReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	enhancePortraitVideoResp, _err := client.EnhancePortraitVideoWithOptions(enhancePortraitVideoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = enhancePortraitVideoResp
	return _result, _err
}

// @param request - EnhanceVideoQualityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnhanceVideoQualityResponse
func (client *Client) EnhanceVideoQualityWithOptions(request *EnhanceVideoQualityRequest, runtime *util.RuntimeOptions) (_result *EnhanceVideoQualityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bitrate)) {
		body["Bitrate"] = request.Bitrate
	}

	if !tea.BoolValue(util.IsUnset(request.FrameRate)) {
		body["FrameRate"] = request.FrameRate
	}

	if !tea.BoolValue(util.IsUnset(request.HDRFormat)) {
		body["HDRFormat"] = request.HDRFormat
	}

	if !tea.BoolValue(util.IsUnset(request.MaxIlluminance)) {
		body["MaxIlluminance"] = request.MaxIlluminance
	}

	if !tea.BoolValue(util.IsUnset(request.OutPutHeight)) {
		body["OutPutHeight"] = request.OutPutHeight
	}

	if !tea.BoolValue(util.IsUnset(request.OutPutWidth)) {
		body["OutPutWidth"] = request.OutPutWidth
	}

	if !tea.BoolValue(util.IsUnset(request.VideoURL)) {
		body["VideoURL"] = request.VideoURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnhanceVideoQuality"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnhanceVideoQualityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnhanceVideoQualityRequest
//
// @return EnhanceVideoQualityResponse
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
		Product:  tea.String("videoenhan"),
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
	enhanceVideoQualityReq := &EnhanceVideoQualityRequest{}
	openapiutil.Convert(request, enhanceVideoQualityReq)
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
		enhanceVideoQualityReq.VideoURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	enhanceVideoQualityResp, _err := client.EnhanceVideoQualityWithOptions(enhanceVideoQualityReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = enhanceVideoQualityResp
	return _result, _err
}

// @param request - EraseVideoLogoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EraseVideoLogoResponse
func (client *Client) EraseVideoLogoWithOptions(request *EraseVideoLogoRequest, runtime *util.RuntimeOptions) (_result *EraseVideoLogoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Boxes)) {
		body["Boxes"] = request.Boxes
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EraseVideoLogo"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EraseVideoLogoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - EraseVideoLogoRequest
//
// @return EraseVideoLogoResponse
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
		Product:  tea.String("videoenhan"),
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
	eraseVideoLogoReq := &EraseVideoLogoRequest{}
	openapiutil.Convert(request, eraseVideoLogoReq)
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
		eraseVideoLogoReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	eraseVideoLogoResp, _err := client.EraseVideoLogoWithOptions(eraseVideoLogoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = eraseVideoLogoResp
	return _result, _err
}

// @param request - EraseVideoSubtitlesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EraseVideoSubtitlesResponse
func (client *Client) EraseVideoSubtitlesWithOptions(request *EraseVideoSubtitlesRequest, runtime *util.RuntimeOptions) (_result *EraseVideoSubtitlesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BH)) {
		body["BH"] = request.BH
	}

	if !tea.BoolValue(util.IsUnset(request.BW)) {
		body["BW"] = request.BW
	}

	if !tea.BoolValue(util.IsUnset(request.BX)) {
		body["BX"] = request.BX
	}

	if !tea.BoolValue(util.IsUnset(request.BY)) {
		body["BY"] = request.BY
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EraseVideoSubtitles"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EraseVideoSubtitlesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - EraseVideoSubtitlesRequest
//
// @return EraseVideoSubtitlesResponse
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
		Product:  tea.String("videoenhan"),
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
	eraseVideoSubtitlesReq := &EraseVideoSubtitlesRequest{}
	openapiutil.Convert(request, eraseVideoSubtitlesReq)
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
		eraseVideoSubtitlesReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	eraseVideoSubtitlesResp, _err := client.EraseVideoSubtitlesWithOptions(eraseVideoSubtitlesReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = eraseVideoSubtitlesResp
	return _result, _err
}

// Summary:
//
// 视频人像卡通化
//
// @param request - GenerateHumanAnimeStyleVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateHumanAnimeStyleVideoResponse
func (client *Client) GenerateHumanAnimeStyleVideoWithOptions(request *GenerateHumanAnimeStyleVideoRequest, runtime *util.RuntimeOptions) (_result *GenerateHumanAnimeStyleVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CartoonStyle)) {
		body["CartoonStyle"] = request.CartoonStyle
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateHumanAnimeStyleVideo"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateHumanAnimeStyleVideoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频人像卡通化
//
// @param request - GenerateHumanAnimeStyleVideoRequest
//
// @return GenerateHumanAnimeStyleVideoResponse
func (client *Client) GenerateHumanAnimeStyleVideo(request *GenerateHumanAnimeStyleVideoRequest) (_result *GenerateHumanAnimeStyleVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateHumanAnimeStyleVideoResponse{}
	_body, _err := client.GenerateHumanAnimeStyleVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateHumanAnimeStyleVideoAdvance(request *GenerateHumanAnimeStyleVideoAdvanceRequest, runtime *util.RuntimeOptions) (_result *GenerateHumanAnimeStyleVideoResponse, _err error) {
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
		Product:  tea.String("videoenhan"),
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
	generateHumanAnimeStyleVideoReq := &GenerateHumanAnimeStyleVideoRequest{}
	openapiutil.Convert(request, generateHumanAnimeStyleVideoReq)
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
		generateHumanAnimeStyleVideoReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	generateHumanAnimeStyleVideoResp, _err := client.GenerateHumanAnimeStyleVideoWithOptions(generateHumanAnimeStyleVideoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = generateHumanAnimeStyleVideoResp
	return _result, _err
}

// @param request - GenerateVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateVideoResponse
func (client *Client) GenerateVideoWithOptions(request *GenerateVideoRequest, runtime *util.RuntimeOptions) (_result *GenerateVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.DurationAdaption)) {
		body["DurationAdaption"] = request.DurationAdaption
	}

	if !tea.BoolValue(util.IsUnset(request.FileList)) {
		body["FileList"] = request.FileList
	}

	if !tea.BoolValue(util.IsUnset(request.Height)) {
		body["Height"] = request.Height
	}

	if !tea.BoolValue(util.IsUnset(request.Mute)) {
		body["Mute"] = request.Mute
	}

	if !tea.BoolValue(util.IsUnset(request.PuzzleEffect)) {
		body["PuzzleEffect"] = request.PuzzleEffect
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.SmartEffect)) {
		body["SmartEffect"] = request.SmartEffect
	}

	if !tea.BoolValue(util.IsUnset(request.Style)) {
		body["Style"] = request.Style
	}

	if !tea.BoolValue(util.IsUnset(request.TransitionStyle)) {
		body["TransitionStyle"] = request.TransitionStyle
	}

	if !tea.BoolValue(util.IsUnset(request.Width)) {
		body["Width"] = request.Width
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateVideo"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateVideoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GenerateVideoRequest
//
// @return GenerateVideoResponse
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

func (client *Client) GenerateVideoAdvance(request *GenerateVideoAdvanceRequest, runtime *util.RuntimeOptions) (_result *GenerateVideoResponse, _err error) {
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
		Product:  tea.String("videoenhan"),
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
	generateVideoReq := &GenerateVideoRequest{}
	openapiutil.Convert(request, generateVideoReq)
	if !tea.BoolValue(util.IsUnset(request.FileList)) {
		i0 := tea.Int(0)
		for _, item0 := range request.FileList {
			if !tea.BoolValue(util.IsUnset(item0.FileUrlObject)) {
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
					Content:     item0.FileUrlObject,
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
				tmp := generateVideoReq.FileList[tea.IntValue(i0)]
				tmp.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i0 = number.Ltoi(number.Add(number.Itol(i0), number.Itol(tea.Int(1))))
			}

		}
	}

	generateVideoResp, _err := client.GenerateVideoWithOptions(generateVideoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = generateVideoResp
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

// @param request - InterpolateVideoFrameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InterpolateVideoFrameResponse
func (client *Client) InterpolateVideoFrameWithOptions(request *InterpolateVideoFrameRequest, runtime *util.RuntimeOptions) (_result *InterpolateVideoFrameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bitrate)) {
		body["Bitrate"] = request.Bitrate
	}

	if !tea.BoolValue(util.IsUnset(request.FrameRate)) {
		body["FrameRate"] = request.FrameRate
	}

	if !tea.BoolValue(util.IsUnset(request.VideoURL)) {
		body["VideoURL"] = request.VideoURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InterpolateVideoFrame"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InterpolateVideoFrameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - InterpolateVideoFrameRequest
//
// @return InterpolateVideoFrameResponse
func (client *Client) InterpolateVideoFrame(request *InterpolateVideoFrameRequest) (_result *InterpolateVideoFrameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InterpolateVideoFrameResponse{}
	_body, _err := client.InterpolateVideoFrameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InterpolateVideoFrameAdvance(request *InterpolateVideoFrameAdvanceRequest, runtime *util.RuntimeOptions) (_result *InterpolateVideoFrameResponse, _err error) {
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
		Product:  tea.String("videoenhan"),
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
	interpolateVideoFrameReq := &InterpolateVideoFrameRequest{}
	openapiutil.Convert(request, interpolateVideoFrameReq)
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
		interpolateVideoFrameReq.VideoURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	interpolateVideoFrameResp, _err := client.InterpolateVideoFrameWithOptions(interpolateVideoFrameReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = interpolateVideoFrameResp
	return _result, _err
}

// @param request - MergeVideoFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MergeVideoFaceResponse
func (client *Client) MergeVideoFaceWithOptions(request *MergeVideoFaceRequest, runtime *util.RuntimeOptions) (_result *MergeVideoFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddWatermark)) {
		body["AddWatermark"] = request.AddWatermark
	}

	if !tea.BoolValue(util.IsUnset(request.Enhance)) {
		body["Enhance"] = request.Enhance
	}

	if !tea.BoolValue(util.IsUnset(request.ReferenceURL)) {
		body["ReferenceURL"] = request.ReferenceURL
	}

	if !tea.BoolValue(util.IsUnset(request.VideoURL)) {
		body["VideoURL"] = request.VideoURL
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkType)) {
		body["WatermarkType"] = request.WatermarkType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MergeVideoFace"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MergeVideoFaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - MergeVideoFaceRequest
//
// @return MergeVideoFaceResponse
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
		Product:  tea.String("videoenhan"),
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
	mergeVideoFaceReq := &MergeVideoFaceRequest{}
	openapiutil.Convert(request, mergeVideoFaceReq)
	if !tea.BoolValue(util.IsUnset(request.ReferenceURLObject)) {
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
			Content:     request.ReferenceURLObject,
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
		mergeVideoFaceReq.ReferenceURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

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
		mergeVideoFaceReq.VideoURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	mergeVideoFaceResp, _err := client.MergeVideoFaceWithOptions(mergeVideoFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = mergeVideoFaceResp
	return _result, _err
}

// Summary:
//
// 视频模板融合换脸
//
// @param request - MergeVideoModelFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MergeVideoModelFaceResponse
func (client *Client) MergeVideoModelFaceWithOptions(request *MergeVideoModelFaceRequest, runtime *util.RuntimeOptions) (_result *MergeVideoModelFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddWatermark)) {
		body["AddWatermark"] = request.AddWatermark
	}

	if !tea.BoolValue(util.IsUnset(request.Enhance)) {
		body["Enhance"] = request.Enhance
	}

	if !tea.BoolValue(util.IsUnset(request.FaceImageURL)) {
		body["FaceImageURL"] = request.FaceImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.MergeInfos)) {
		body["MergeInfos"] = request.MergeInfos
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkType)) {
		body["WatermarkType"] = request.WatermarkType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MergeVideoModelFace"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MergeVideoModelFaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频模板融合换脸
//
// @param request - MergeVideoModelFaceRequest
//
// @return MergeVideoModelFaceResponse
func (client *Client) MergeVideoModelFace(request *MergeVideoModelFaceRequest) (_result *MergeVideoModelFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MergeVideoModelFaceResponse{}
	_body, _err := client.MergeVideoModelFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MergeVideoModelFaceAdvance(request *MergeVideoModelFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *MergeVideoModelFaceResponse, _err error) {
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
		Product:  tea.String("videoenhan"),
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
	mergeVideoModelFaceReq := &MergeVideoModelFaceRequest{}
	openapiutil.Convert(request, mergeVideoModelFaceReq)
	if !tea.BoolValue(util.IsUnset(request.FaceImageURLObject)) {
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
			Content:     request.FaceImageURLObject,
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
		mergeVideoModelFaceReq.FaceImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	mergeVideoModelFaceResp, _err := client.MergeVideoModelFaceWithOptions(mergeVideoModelFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = mergeVideoModelFaceResp
	return _result, _err
}

// Summary:
//
// 视频人脸融合模板查询
//
// @param request - QueryFaceVideoTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFaceVideoTemplateResponse
func (client *Client) QueryFaceVideoTemplateWithOptions(request *QueryFaceVideoTemplateRequest, runtime *util.RuntimeOptions) (_result *QueryFaceVideoTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFaceVideoTemplate"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFaceVideoTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频人脸融合模板查询
//
// @param request - QueryFaceVideoTemplateRequest
//
// @return QueryFaceVideoTemplateResponse
func (client *Client) QueryFaceVideoTemplate(request *QueryFaceVideoTemplateRequest) (_result *QueryFaceVideoTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFaceVideoTemplateResponse{}
	_body, _err := client.QueryFaceVideoTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 视频降噪
//
// @param request - ReduceVideoNoiseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReduceVideoNoiseResponse
func (client *Client) ReduceVideoNoiseWithOptions(request *ReduceVideoNoiseRequest, runtime *util.RuntimeOptions) (_result *ReduceVideoNoiseResponse, _err error) {
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
		Action:      tea.String("ReduceVideoNoise"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReduceVideoNoiseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频降噪
//
// @param request - ReduceVideoNoiseRequest
//
// @return ReduceVideoNoiseResponse
func (client *Client) ReduceVideoNoise(request *ReduceVideoNoiseRequest) (_result *ReduceVideoNoiseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReduceVideoNoiseResponse{}
	_body, _err := client.ReduceVideoNoiseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReduceVideoNoiseAdvance(request *ReduceVideoNoiseAdvanceRequest, runtime *util.RuntimeOptions) (_result *ReduceVideoNoiseResponse, _err error) {
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
		Product:  tea.String("videoenhan"),
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
	reduceVideoNoiseReq := &ReduceVideoNoiseRequest{}
	openapiutil.Convert(request, reduceVideoNoiseReq)
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
		reduceVideoNoiseReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	reduceVideoNoiseResp, _err := client.ReduceVideoNoiseWithOptions(reduceVideoNoiseReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = reduceVideoNoiseResp
	return _result, _err
}

// @param request - SuperResolveVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuperResolveVideoResponse
func (client *Client) SuperResolveVideoWithOptions(request *SuperResolveVideoRequest, runtime *util.RuntimeOptions) (_result *SuperResolveVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BitRate)) {
		body["BitRate"] = request.BitRate
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SuperResolveVideo"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SuperResolveVideoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SuperResolveVideoRequest
//
// @return SuperResolveVideoResponse
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
		Product:  tea.String("videoenhan"),
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
	superResolveVideoReq := &SuperResolveVideoRequest{}
	openapiutil.Convert(request, superResolveVideoReq)
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
		superResolveVideoReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	superResolveVideoResp, _err := client.SuperResolveVideoWithOptions(superResolveVideoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = superResolveVideoResp
	return _result, _err
}

// @param request - ToneSdrVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ToneSdrVideoResponse
func (client *Client) ToneSdrVideoWithOptions(request *ToneSdrVideoRequest, runtime *util.RuntimeOptions) (_result *ToneSdrVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bitrate)) {
		body["Bitrate"] = request.Bitrate
	}

	if !tea.BoolValue(util.IsUnset(request.RecolorModel)) {
		body["RecolorModel"] = request.RecolorModel
	}

	if !tea.BoolValue(util.IsUnset(request.VideoURL)) {
		body["VideoURL"] = request.VideoURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ToneSdrVideo"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ToneSdrVideoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ToneSdrVideoRequest
//
// @return ToneSdrVideoResponse
func (client *Client) ToneSdrVideo(request *ToneSdrVideoRequest) (_result *ToneSdrVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ToneSdrVideoResponse{}
	_body, _err := client.ToneSdrVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ToneSdrVideoAdvance(request *ToneSdrVideoAdvanceRequest, runtime *util.RuntimeOptions) (_result *ToneSdrVideoResponse, _err error) {
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
		Product:  tea.String("videoenhan"),
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
	toneSdrVideoReq := &ToneSdrVideoRequest{}
	openapiutil.Convert(request, toneSdrVideoReq)
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
		toneSdrVideoReq.VideoURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	toneSdrVideoResp, _err := client.ToneSdrVideoWithOptions(toneSdrVideoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = toneSdrVideoResp
	return _result, _err
}
