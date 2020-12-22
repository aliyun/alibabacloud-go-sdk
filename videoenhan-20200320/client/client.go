// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	rpcutil "github.com/alibabacloud-go/tea-rpc-utils/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type ConvertHdrVideoRequest struct {
	VideoURL       *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty" require:"true"`
	HDRFormat      *string `json:"HDRFormat,omitempty" xml:"HDRFormat,omitempty"`
	MaxIlluminance *int    `json:"MaxIlluminance,omitempty" xml:"MaxIlluminance,omitempty"`
	Bitrate        *int    `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
}

func (s ConvertHdrVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertHdrVideoRequest) GoString() string {
	return s.String()
}

func (s *ConvertHdrVideoRequest) SetVideoURL(v string) *ConvertHdrVideoRequest {
	s.VideoURL = &v
	return s
}

func (s *ConvertHdrVideoRequest) SetHDRFormat(v string) *ConvertHdrVideoRequest {
	s.HDRFormat = &v
	return s
}

func (s *ConvertHdrVideoRequest) SetMaxIlluminance(v int) *ConvertHdrVideoRequest {
	s.MaxIlluminance = &v
	return s
}

func (s *ConvertHdrVideoRequest) SetBitrate(v int) *ConvertHdrVideoRequest {
	s.Bitrate = &v
	return s
}

type ConvertHdrVideoResponse struct {
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ConvertHdrVideoResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ConvertHdrVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s ConvertHdrVideoResponse) GoString() string {
	return s.String()
}

func (s *ConvertHdrVideoResponse) SetRequestId(v string) *ConvertHdrVideoResponse {
	s.RequestId = &v
	return s
}

func (s *ConvertHdrVideoResponse) SetData(v *ConvertHdrVideoResponseData) *ConvertHdrVideoResponse {
	s.Data = v
	return s
}

type ConvertHdrVideoResponseData struct {
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty" require:"true"`
}

func (s ConvertHdrVideoResponseData) String() string {
	return tea.Prettify(s)
}

func (s ConvertHdrVideoResponseData) GoString() string {
	return s.String()
}

func (s *ConvertHdrVideoResponseData) SetVideoURL(v string) *ConvertHdrVideoResponseData {
	s.VideoURL = &v
	return s
}

type ConvertHdrVideoAdvanceRequest struct {
	VideoURLObject io.Reader `json:"VideoURLObject,omitempty" xml:"VideoURLObject,omitempty" require:"true"`
	HDRFormat      *string   `json:"HDRFormat,omitempty" xml:"HDRFormat,omitempty"`
	MaxIlluminance *int      `json:"MaxIlluminance,omitempty" xml:"MaxIlluminance,omitempty"`
	Bitrate        *int      `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
}

func (s ConvertHdrVideoAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertHdrVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ConvertHdrVideoAdvanceRequest) SetVideoURLObject(v io.Reader) *ConvertHdrVideoAdvanceRequest {
	s.VideoURLObject = v
	return s
}

func (s *ConvertHdrVideoAdvanceRequest) SetHDRFormat(v string) *ConvertHdrVideoAdvanceRequest {
	s.HDRFormat = &v
	return s
}

func (s *ConvertHdrVideoAdvanceRequest) SetMaxIlluminance(v int) *ConvertHdrVideoAdvanceRequest {
	s.MaxIlluminance = &v
	return s
}

func (s *ConvertHdrVideoAdvanceRequest) SetBitrate(v int) *ConvertHdrVideoAdvanceRequest {
	s.Bitrate = &v
	return s
}

type InterpolateVideoFrameRequest struct {
	VideoURL  *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty" require:"true"`
	FrameRate *int    `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	Bitrate   *int    `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
}

func (s InterpolateVideoFrameRequest) String() string {
	return tea.Prettify(s)
}

func (s InterpolateVideoFrameRequest) GoString() string {
	return s.String()
}

func (s *InterpolateVideoFrameRequest) SetVideoURL(v string) *InterpolateVideoFrameRequest {
	s.VideoURL = &v
	return s
}

func (s *InterpolateVideoFrameRequest) SetFrameRate(v int) *InterpolateVideoFrameRequest {
	s.FrameRate = &v
	return s
}

func (s *InterpolateVideoFrameRequest) SetBitrate(v int) *InterpolateVideoFrameRequest {
	s.Bitrate = &v
	return s
}

type InterpolateVideoFrameResponse struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *InterpolateVideoFrameResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s InterpolateVideoFrameResponse) String() string {
	return tea.Prettify(s)
}

func (s InterpolateVideoFrameResponse) GoString() string {
	return s.String()
}

func (s *InterpolateVideoFrameResponse) SetRequestId(v string) *InterpolateVideoFrameResponse {
	s.RequestId = &v
	return s
}

func (s *InterpolateVideoFrameResponse) SetData(v *InterpolateVideoFrameResponseData) *InterpolateVideoFrameResponse {
	s.Data = v
	return s
}

type InterpolateVideoFrameResponseData struct {
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty" require:"true"`
}

func (s InterpolateVideoFrameResponseData) String() string {
	return tea.Prettify(s)
}

func (s InterpolateVideoFrameResponseData) GoString() string {
	return s.String()
}

func (s *InterpolateVideoFrameResponseData) SetVideoURL(v string) *InterpolateVideoFrameResponseData {
	s.VideoURL = &v
	return s
}

type InterpolateVideoFrameAdvanceRequest struct {
	VideoURLObject io.Reader `json:"VideoURLObject,omitempty" xml:"VideoURLObject,omitempty" require:"true"`
	FrameRate      *int      `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	Bitrate        *int      `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
}

func (s InterpolateVideoFrameAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s InterpolateVideoFrameAdvanceRequest) GoString() string {
	return s.String()
}

func (s *InterpolateVideoFrameAdvanceRequest) SetVideoURLObject(v io.Reader) *InterpolateVideoFrameAdvanceRequest {
	s.VideoURLObject = v
	return s
}

func (s *InterpolateVideoFrameAdvanceRequest) SetFrameRate(v int) *InterpolateVideoFrameAdvanceRequest {
	s.FrameRate = &v
	return s
}

func (s *InterpolateVideoFrameAdvanceRequest) SetBitrate(v int) *InterpolateVideoFrameAdvanceRequest {
	s.Bitrate = &v
	return s
}

type ToneSdrVideoRequest struct {
	VideoURL     *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty" require:"true"`
	Bitrate      *int    `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	RecolorModel *string `json:"RecolorModel,omitempty" xml:"RecolorModel,omitempty"`
}

func (s ToneSdrVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s ToneSdrVideoRequest) GoString() string {
	return s.String()
}

func (s *ToneSdrVideoRequest) SetVideoURL(v string) *ToneSdrVideoRequest {
	s.VideoURL = &v
	return s
}

func (s *ToneSdrVideoRequest) SetBitrate(v int) *ToneSdrVideoRequest {
	s.Bitrate = &v
	return s
}

func (s *ToneSdrVideoRequest) SetRecolorModel(v string) *ToneSdrVideoRequest {
	s.RecolorModel = &v
	return s
}

type ToneSdrVideoResponse struct {
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ToneSdrVideoResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ToneSdrVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s ToneSdrVideoResponse) GoString() string {
	return s.String()
}

func (s *ToneSdrVideoResponse) SetRequestId(v string) *ToneSdrVideoResponse {
	s.RequestId = &v
	return s
}

func (s *ToneSdrVideoResponse) SetData(v *ToneSdrVideoResponseData) *ToneSdrVideoResponse {
	s.Data = v
	return s
}

type ToneSdrVideoResponseData struct {
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty" require:"true"`
}

func (s ToneSdrVideoResponseData) String() string {
	return tea.Prettify(s)
}

func (s ToneSdrVideoResponseData) GoString() string {
	return s.String()
}

func (s *ToneSdrVideoResponseData) SetVideoURL(v string) *ToneSdrVideoResponseData {
	s.VideoURL = &v
	return s
}

type ToneSdrVideoAdvanceRequest struct {
	VideoURLObject io.Reader `json:"VideoURLObject,omitempty" xml:"VideoURLObject,omitempty" require:"true"`
	Bitrate        *int      `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	RecolorModel   *string   `json:"RecolorModel,omitempty" xml:"RecolorModel,omitempty"`
}

func (s ToneSdrVideoAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ToneSdrVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ToneSdrVideoAdvanceRequest) SetVideoURLObject(v io.Reader) *ToneSdrVideoAdvanceRequest {
	s.VideoURLObject = v
	return s
}

func (s *ToneSdrVideoAdvanceRequest) SetBitrate(v int) *ToneSdrVideoAdvanceRequest {
	s.Bitrate = &v
	return s
}

func (s *ToneSdrVideoAdvanceRequest) SetRecolorModel(v string) *ToneSdrVideoAdvanceRequest {
	s.RecolorModel = &v
	return s
}

type EnhanceVideoQualityRequest struct {
	VideoURL       *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty" require:"true"`
	OutPutWidth    *int    `json:"OutPutWidth,omitempty" xml:"OutPutWidth,omitempty"`
	OutPutHeight   *int    `json:"OutPutHeight,omitempty" xml:"OutPutHeight,omitempty"`
	FrameRate      *int    `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	HDRFormat      *string `json:"HDRFormat,omitempty" xml:"HDRFormat,omitempty"`
	MaxIlluminance *int    `json:"MaxIlluminance,omitempty" xml:"MaxIlluminance,omitempty"`
	Bitrate        *int    `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
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

func (s *EnhanceVideoQualityRequest) SetOutPutWidth(v int) *EnhanceVideoQualityRequest {
	s.OutPutWidth = &v
	return s
}

func (s *EnhanceVideoQualityRequest) SetOutPutHeight(v int) *EnhanceVideoQualityRequest {
	s.OutPutHeight = &v
	return s
}

func (s *EnhanceVideoQualityRequest) SetFrameRate(v int) *EnhanceVideoQualityRequest {
	s.FrameRate = &v
	return s
}

func (s *EnhanceVideoQualityRequest) SetHDRFormat(v string) *EnhanceVideoQualityRequest {
	s.HDRFormat = &v
	return s
}

func (s *EnhanceVideoQualityRequest) SetMaxIlluminance(v int) *EnhanceVideoQualityRequest {
	s.MaxIlluminance = &v
	return s
}

func (s *EnhanceVideoQualityRequest) SetBitrate(v int) *EnhanceVideoQualityRequest {
	s.Bitrate = &v
	return s
}

type EnhanceVideoQualityResponse struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *EnhanceVideoQualityResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s EnhanceVideoQualityResponse) String() string {
	return tea.Prettify(s)
}

func (s EnhanceVideoQualityResponse) GoString() string {
	return s.String()
}

func (s *EnhanceVideoQualityResponse) SetRequestId(v string) *EnhanceVideoQualityResponse {
	s.RequestId = &v
	return s
}

func (s *EnhanceVideoQualityResponse) SetData(v *EnhanceVideoQualityResponseData) *EnhanceVideoQualityResponse {
	s.Data = v
	return s
}

type EnhanceVideoQualityResponseData struct {
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty" require:"true"`
}

func (s EnhanceVideoQualityResponseData) String() string {
	return tea.Prettify(s)
}

func (s EnhanceVideoQualityResponseData) GoString() string {
	return s.String()
}

func (s *EnhanceVideoQualityResponseData) SetVideoURL(v string) *EnhanceVideoQualityResponseData {
	s.VideoURL = &v
	return s
}

type EnhanceVideoQualityAdvanceRequest struct {
	VideoURLObject io.Reader `json:"VideoURLObject,omitempty" xml:"VideoURLObject,omitempty" require:"true"`
	OutPutWidth    *int      `json:"OutPutWidth,omitempty" xml:"OutPutWidth,omitempty"`
	OutPutHeight   *int      `json:"OutPutHeight,omitempty" xml:"OutPutHeight,omitempty"`
	FrameRate      *int      `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	HDRFormat      *string   `json:"HDRFormat,omitempty" xml:"HDRFormat,omitempty"`
	MaxIlluminance *int      `json:"MaxIlluminance,omitempty" xml:"MaxIlluminance,omitempty"`
	Bitrate        *int      `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
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

func (s *EnhanceVideoQualityAdvanceRequest) SetOutPutWidth(v int) *EnhanceVideoQualityAdvanceRequest {
	s.OutPutWidth = &v
	return s
}

func (s *EnhanceVideoQualityAdvanceRequest) SetOutPutHeight(v int) *EnhanceVideoQualityAdvanceRequest {
	s.OutPutHeight = &v
	return s
}

func (s *EnhanceVideoQualityAdvanceRequest) SetFrameRate(v int) *EnhanceVideoQualityAdvanceRequest {
	s.FrameRate = &v
	return s
}

func (s *EnhanceVideoQualityAdvanceRequest) SetHDRFormat(v string) *EnhanceVideoQualityAdvanceRequest {
	s.HDRFormat = &v
	return s
}

func (s *EnhanceVideoQualityAdvanceRequest) SetMaxIlluminance(v int) *EnhanceVideoQualityAdvanceRequest {
	s.MaxIlluminance = &v
	return s
}

func (s *EnhanceVideoQualityAdvanceRequest) SetBitrate(v int) *EnhanceVideoQualityAdvanceRequest {
	s.Bitrate = &v
	return s
}

type MergeVideoFaceRequest struct {
	VideoURL     *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty" require:"true"`
	PostURL      *string `json:"PostURL,omitempty" xml:"PostURL,omitempty" require:"true"`
	ReferenceURL *string `json:"ReferenceURL,omitempty" xml:"ReferenceURL,omitempty" require:"true"`
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

func (s *MergeVideoFaceRequest) SetPostURL(v string) *MergeVideoFaceRequest {
	s.PostURL = &v
	return s
}

func (s *MergeVideoFaceRequest) SetReferenceURL(v string) *MergeVideoFaceRequest {
	s.ReferenceURL = &v
	return s
}

type MergeVideoFaceResponse struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *MergeVideoFaceResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s MergeVideoFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s MergeVideoFaceResponse) GoString() string {
	return s.String()
}

func (s *MergeVideoFaceResponse) SetRequestId(v string) *MergeVideoFaceResponse {
	s.RequestId = &v
	return s
}

func (s *MergeVideoFaceResponse) SetData(v *MergeVideoFaceResponseData) *MergeVideoFaceResponse {
	s.Data = v
	return s
}

type MergeVideoFaceResponseData struct {
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty" require:"true"`
}

func (s MergeVideoFaceResponseData) String() string {
	return tea.Prettify(s)
}

func (s MergeVideoFaceResponseData) GoString() string {
	return s.String()
}

func (s *MergeVideoFaceResponseData) SetVideoURL(v string) *MergeVideoFaceResponseData {
	s.VideoURL = &v
	return s
}

type MergeVideoFaceAdvanceRequest struct {
	VideoURLObject io.Reader `json:"VideoURLObject,omitempty" xml:"VideoURLObject,omitempty" require:"true"`
	PostURL        *string   `json:"PostURL,omitempty" xml:"PostURL,omitempty" require:"true"`
	ReferenceURL   *string   `json:"ReferenceURL,omitempty" xml:"ReferenceURL,omitempty" require:"true"`
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

func (s *MergeVideoFaceAdvanceRequest) SetPostURL(v string) *MergeVideoFaceAdvanceRequest {
	s.PostURL = &v
	return s
}

func (s *MergeVideoFaceAdvanceRequest) SetReferenceURL(v string) *MergeVideoFaceAdvanceRequest {
	s.ReferenceURL = &v
	return s
}

type ChangeVideoSizeRequest struct {
	VideoUrl  *string  `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
	Width     *int     `json:"Width,omitempty" xml:"Width,omitempty" require:"true"`
	Height    *int     `json:"Height,omitempty" xml:"Height,omitempty" require:"true"`
	CropType  *string  `json:"CropType,omitempty" xml:"CropType,omitempty"`
	FillType  *string  `json:"FillType,omitempty" xml:"FillType,omitempty"`
	Tightness *float32 `json:"Tightness,omitempty" xml:"Tightness,omitempty"`
	R         *int     `json:"R,omitempty" xml:"R,omitempty"`
	G         *int     `json:"G,omitempty" xml:"G,omitempty"`
	B         *int     `json:"B,omitempty" xml:"B,omitempty"`
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

func (s *ChangeVideoSizeRequest) SetWidth(v int) *ChangeVideoSizeRequest {
	s.Width = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetHeight(v int) *ChangeVideoSizeRequest {
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

func (s *ChangeVideoSizeRequest) SetR(v int) *ChangeVideoSizeRequest {
	s.R = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetG(v int) *ChangeVideoSizeRequest {
	s.G = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetB(v int) *ChangeVideoSizeRequest {
	s.B = &v
	return s
}

type ChangeVideoSizeResponse struct {
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ChangeVideoSizeResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ChangeVideoSizeResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeVideoSizeResponse) GoString() string {
	return s.String()
}

func (s *ChangeVideoSizeResponse) SetRequestId(v string) *ChangeVideoSizeResponse {
	s.RequestId = &v
	return s
}

func (s *ChangeVideoSizeResponse) SetData(v *ChangeVideoSizeResponseData) *ChangeVideoSizeResponse {
	s.Data = v
	return s
}

type ChangeVideoSizeResponseData struct {
	VideoUrl      *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
	VideoCoverUrl *string `json:"VideoCoverUrl,omitempty" xml:"VideoCoverUrl,omitempty" require:"true"`
}

func (s ChangeVideoSizeResponseData) String() string {
	return tea.Prettify(s)
}

func (s ChangeVideoSizeResponseData) GoString() string {
	return s.String()
}

func (s *ChangeVideoSizeResponseData) SetVideoUrl(v string) *ChangeVideoSizeResponseData {
	s.VideoUrl = &v
	return s
}

func (s *ChangeVideoSizeResponseData) SetVideoCoverUrl(v string) *ChangeVideoSizeResponseData {
	s.VideoCoverUrl = &v
	return s
}

type ChangeVideoSizeAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrlObject,omitempty" xml:"VideoUrlObject,omitempty" require:"true"`
	Width          *int      `json:"Width,omitempty" xml:"Width,omitempty" require:"true"`
	Height         *int      `json:"Height,omitempty" xml:"Height,omitempty" require:"true"`
	CropType       *string   `json:"CropType,omitempty" xml:"CropType,omitempty"`
	FillType       *string   `json:"FillType,omitempty" xml:"FillType,omitempty"`
	Tightness      *float32  `json:"Tightness,omitempty" xml:"Tightness,omitempty"`
	R              *int      `json:"R,omitempty" xml:"R,omitempty"`
	G              *int      `json:"G,omitempty" xml:"G,omitempty"`
	B              *int      `json:"B,omitempty" xml:"B,omitempty"`
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

func (s *ChangeVideoSizeAdvanceRequest) SetWidth(v int) *ChangeVideoSizeAdvanceRequest {
	s.Width = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetHeight(v int) *ChangeVideoSizeAdvanceRequest {
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

func (s *ChangeVideoSizeAdvanceRequest) SetR(v int) *ChangeVideoSizeAdvanceRequest {
	s.R = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetG(v int) *ChangeVideoSizeAdvanceRequest {
	s.G = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetB(v int) *ChangeVideoSizeAdvanceRequest {
	s.B = &v
	return s
}

type GenerateVideoRequest struct {
	FileList         []*GenerateVideoRequestFileList `json:"FileList,omitempty" xml:"FileList,omitempty" require:"true" type:"Repeated"`
	Scene            *string                         `json:"Scene,omitempty" xml:"Scene,omitempty"`
	Width            *int                            `json:"Width,omitempty" xml:"Width,omitempty"`
	Height           *int                            `json:"Height,omitempty" xml:"Height,omitempty"`
	Style            *string                         `json:"Style,omitempty" xml:"Style,omitempty"`
	Duration         *float32                        `json:"Duration,omitempty" xml:"Duration,omitempty"`
	DurationAdaption *bool                           `json:"DurationAdaption,omitempty" xml:"DurationAdaption,omitempty"`
	TransitionStyle  *string                         `json:"TransitionStyle,omitempty" xml:"TransitionStyle,omitempty"`
	SmartEffect      *bool                           `json:"SmartEffect,omitempty" xml:"SmartEffect,omitempty"`
	PuzzleEffect     *bool                           `json:"PuzzleEffect,omitempty" xml:"PuzzleEffect,omitempty"`
	Mute             *bool                           `json:"Mute,omitempty" xml:"Mute,omitempty"`
}

func (s GenerateVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoRequest) GoString() string {
	return s.String()
}

func (s *GenerateVideoRequest) SetFileList(v []*GenerateVideoRequestFileList) *GenerateVideoRequest {
	s.FileList = v
	return s
}

func (s *GenerateVideoRequest) SetScene(v string) *GenerateVideoRequest {
	s.Scene = &v
	return s
}

func (s *GenerateVideoRequest) SetWidth(v int) *GenerateVideoRequest {
	s.Width = &v
	return s
}

func (s *GenerateVideoRequest) SetHeight(v int) *GenerateVideoRequest {
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

type GenerateVideoRequestFileList struct {
	FileUrl  *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty" require:"true"`
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty" require:"true"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
}

func (s GenerateVideoRequestFileList) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoRequestFileList) GoString() string {
	return s.String()
}

func (s *GenerateVideoRequestFileList) SetFileUrl(v string) *GenerateVideoRequestFileList {
	s.FileUrl = &v
	return s
}

func (s *GenerateVideoRequestFileList) SetFileName(v string) *GenerateVideoRequestFileList {
	s.FileName = &v
	return s
}

func (s *GenerateVideoRequestFileList) SetType(v string) *GenerateVideoRequestFileList {
	s.Type = &v
	return s
}

type GenerateVideoResponse struct {
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *GenerateVideoResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GenerateVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoResponse) GoString() string {
	return s.String()
}

func (s *GenerateVideoResponse) SetRequestId(v string) *GenerateVideoResponse {
	s.RequestId = &v
	return s
}

func (s *GenerateVideoResponse) SetData(v *GenerateVideoResponseData) *GenerateVideoResponse {
	s.Data = v
	return s
}

type GenerateVideoResponseData struct {
	VideoUrl      *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
	VideoCoverUrl *string `json:"VideoCoverUrl,omitempty" xml:"VideoCoverUrl,omitempty" require:"true"`
}

func (s GenerateVideoResponseData) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoResponseData) GoString() string {
	return s.String()
}

func (s *GenerateVideoResponseData) SetVideoUrl(v string) *GenerateVideoResponseData {
	s.VideoUrl = &v
	return s
}

func (s *GenerateVideoResponseData) SetVideoCoverUrl(v string) *GenerateVideoResponseData {
	s.VideoCoverUrl = &v
	return s
}

type GetAsyncJobResultRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty" require:"true"`
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

type GetAsyncJobResultResponse struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *GetAsyncJobResultResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetAsyncJobResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponse) SetRequestId(v string) *GetAsyncJobResultResponse {
	s.RequestId = &v
	return s
}

func (s *GetAsyncJobResultResponse) SetData(v *GetAsyncJobResultResponseData) *GetAsyncJobResultResponse {
	s.Data = v
	return s
}

type GetAsyncJobResultResponseData struct {
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty" require:"true"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty" require:"true"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty" require:"true"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty" require:"true"`
}

func (s GetAsyncJobResultResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseData) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseData) SetJobId(v string) *GetAsyncJobResultResponseData {
	s.JobId = &v
	return s
}

func (s *GetAsyncJobResultResponseData) SetStatus(v string) *GetAsyncJobResultResponseData {
	s.Status = &v
	return s
}

func (s *GetAsyncJobResultResponseData) SetResult(v string) *GetAsyncJobResultResponseData {
	s.Result = &v
	return s
}

func (s *GetAsyncJobResultResponseData) SetErrorCode(v string) *GetAsyncJobResultResponseData {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncJobResultResponseData) SetErrorMessage(v string) *GetAsyncJobResultResponseData {
	s.ErrorMessage = &v
	return s
}

type SuperResolveVideoRequest struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
	BitRate  *int    `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
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

func (s *SuperResolveVideoRequest) SetBitRate(v int) *SuperResolveVideoRequest {
	s.BitRate = &v
	return s
}

type SuperResolveVideoResponse struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SuperResolveVideoResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SuperResolveVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s SuperResolveVideoResponse) GoString() string {
	return s.String()
}

func (s *SuperResolveVideoResponse) SetRequestId(v string) *SuperResolveVideoResponse {
	s.RequestId = &v
	return s
}

func (s *SuperResolveVideoResponse) SetData(v *SuperResolveVideoResponseData) *SuperResolveVideoResponse {
	s.Data = v
	return s
}

type SuperResolveVideoResponseData struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
}

func (s SuperResolveVideoResponseData) String() string {
	return tea.Prettify(s)
}

func (s SuperResolveVideoResponseData) GoString() string {
	return s.String()
}

func (s *SuperResolveVideoResponseData) SetVideoUrl(v string) *SuperResolveVideoResponseData {
	s.VideoUrl = &v
	return s
}

type SuperResolveVideoAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrlObject,omitempty" xml:"VideoUrlObject,omitempty" require:"true"`
	BitRate        *int      `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
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

func (s *SuperResolveVideoAdvanceRequest) SetBitRate(v int) *SuperResolveVideoAdvanceRequest {
	s.BitRate = &v
	return s
}

type EraseVideoLogoRequest struct {
	VideoUrl *string                       `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
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

func (s *EraseVideoLogoRequest) SetBoxes(v []*EraseVideoLogoRequestBoxes) *EraseVideoLogoRequest {
	s.Boxes = v
	return s
}

type EraseVideoLogoRequestBoxes struct {
	H *float32 `json:"H,omitempty" xml:"H,omitempty"`
	W *float32 `json:"W,omitempty" xml:"W,omitempty"`
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
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

type EraseVideoLogoResponse struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *EraseVideoLogoResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s EraseVideoLogoResponse) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoLogoResponse) GoString() string {
	return s.String()
}

func (s *EraseVideoLogoResponse) SetRequestId(v string) *EraseVideoLogoResponse {
	s.RequestId = &v
	return s
}

func (s *EraseVideoLogoResponse) SetData(v *EraseVideoLogoResponseData) *EraseVideoLogoResponse {
	s.Data = v
	return s
}

type EraseVideoLogoResponseData struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
}

func (s EraseVideoLogoResponseData) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoLogoResponseData) GoString() string {
	return s.String()
}

func (s *EraseVideoLogoResponseData) SetVideoUrl(v string) *EraseVideoLogoResponseData {
	s.VideoUrl = &v
	return s
}

type EraseVideoLogoAdvanceRequest struct {
	VideoUrlObject io.Reader                            `json:"VideoUrlObject,omitempty" xml:"VideoUrlObject,omitempty" require:"true"`
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

func (s *EraseVideoLogoAdvanceRequest) SetBoxes(v []*EraseVideoLogoAdvanceRequestBoxes) *EraseVideoLogoAdvanceRequest {
	s.Boxes = v
	return s
}

type EraseVideoLogoAdvanceRequestBoxes struct {
	H *float32 `json:"H,omitempty" xml:"H,omitempty"`
	W *float32 `json:"W,omitempty" xml:"W,omitempty"`
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
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

type EraseVideoSubtitlesRequest struct {
	VideoUrl *string  `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
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

type EraseVideoSubtitlesResponse struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *EraseVideoSubtitlesResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s EraseVideoSubtitlesResponse) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoSubtitlesResponse) GoString() string {
	return s.String()
}

func (s *EraseVideoSubtitlesResponse) SetRequestId(v string) *EraseVideoSubtitlesResponse {
	s.RequestId = &v
	return s
}

func (s *EraseVideoSubtitlesResponse) SetData(v *EraseVideoSubtitlesResponseData) *EraseVideoSubtitlesResponse {
	s.Data = v
	return s
}

type EraseVideoSubtitlesResponseData struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
}

func (s EraseVideoSubtitlesResponseData) String() string {
	return tea.Prettify(s)
}

func (s EraseVideoSubtitlesResponseData) GoString() string {
	return s.String()
}

func (s *EraseVideoSubtitlesResponseData) SetVideoUrl(v string) *EraseVideoSubtitlesResponseData {
	s.VideoUrl = &v
	return s
}

type EraseVideoSubtitlesAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrlObject,omitempty" xml:"VideoUrlObject,omitempty" require:"true"`
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

type AbstractEcommerceVideoRequest struct {
	VideoUrl *string  `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty" require:"true"`
	Width    *int     `json:"Width,omitempty" xml:"Width,omitempty"`
	Height   *int     `json:"Height,omitempty" xml:"Height,omitempty"`
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

func (s *AbstractEcommerceVideoRequest) SetDuration(v float32) *AbstractEcommerceVideoRequest {
	s.Duration = &v
	return s
}

func (s *AbstractEcommerceVideoRequest) SetWidth(v int) *AbstractEcommerceVideoRequest {
	s.Width = &v
	return s
}

func (s *AbstractEcommerceVideoRequest) SetHeight(v int) *AbstractEcommerceVideoRequest {
	s.Height = &v
	return s
}

type AbstractEcommerceVideoResponse struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *AbstractEcommerceVideoResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s AbstractEcommerceVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s AbstractEcommerceVideoResponse) GoString() string {
	return s.String()
}

func (s *AbstractEcommerceVideoResponse) SetRequestId(v string) *AbstractEcommerceVideoResponse {
	s.RequestId = &v
	return s
}

func (s *AbstractEcommerceVideoResponse) SetData(v *AbstractEcommerceVideoResponseData) *AbstractEcommerceVideoResponse {
	s.Data = v
	return s
}

type AbstractEcommerceVideoResponseData struct {
	VideoUrl      *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
	VideoCoverUrl *string `json:"VideoCoverUrl,omitempty" xml:"VideoCoverUrl,omitempty" require:"true"`
}

func (s AbstractEcommerceVideoResponseData) String() string {
	return tea.Prettify(s)
}

func (s AbstractEcommerceVideoResponseData) GoString() string {
	return s.String()
}

func (s *AbstractEcommerceVideoResponseData) SetVideoUrl(v string) *AbstractEcommerceVideoResponseData {
	s.VideoUrl = &v
	return s
}

func (s *AbstractEcommerceVideoResponseData) SetVideoCoverUrl(v string) *AbstractEcommerceVideoResponseData {
	s.VideoCoverUrl = &v
	return s
}

type AbstractEcommerceVideoAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrlObject,omitempty" xml:"VideoUrlObject,omitempty" require:"true"`
	Duration       *float32  `json:"Duration,omitempty" xml:"Duration,omitempty" require:"true"`
	Width          *int      `json:"Width,omitempty" xml:"Width,omitempty"`
	Height         *int      `json:"Height,omitempty" xml:"Height,omitempty"`
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

func (s *AbstractEcommerceVideoAdvanceRequest) SetDuration(v float32) *AbstractEcommerceVideoAdvanceRequest {
	s.Duration = &v
	return s
}

func (s *AbstractEcommerceVideoAdvanceRequest) SetWidth(v int) *AbstractEcommerceVideoAdvanceRequest {
	s.Width = &v
	return s
}

func (s *AbstractEcommerceVideoAdvanceRequest) SetHeight(v int) *AbstractEcommerceVideoAdvanceRequest {
	s.Height = &v
	return s
}

type AbstractFilmVideoRequest struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
	Length   *int    `json:"Length,omitempty" xml:"Length,omitempty" require:"true"`
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

func (s *AbstractFilmVideoRequest) SetLength(v int) *AbstractFilmVideoRequest {
	s.Length = &v
	return s
}

type AbstractFilmVideoResponse struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *AbstractFilmVideoResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s AbstractFilmVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s AbstractFilmVideoResponse) GoString() string {
	return s.String()
}

func (s *AbstractFilmVideoResponse) SetRequestId(v string) *AbstractFilmVideoResponse {
	s.RequestId = &v
	return s
}

func (s *AbstractFilmVideoResponse) SetData(v *AbstractFilmVideoResponseData) *AbstractFilmVideoResponse {
	s.Data = v
	return s
}

type AbstractFilmVideoResponseData struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
}

func (s AbstractFilmVideoResponseData) String() string {
	return tea.Prettify(s)
}

func (s AbstractFilmVideoResponseData) GoString() string {
	return s.String()
}

func (s *AbstractFilmVideoResponseData) SetVideoUrl(v string) *AbstractFilmVideoResponseData {
	s.VideoUrl = &v
	return s
}

type AbstractFilmVideoAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrlObject,omitempty" xml:"VideoUrlObject,omitempty" require:"true"`
	Length         *int      `json:"Length,omitempty" xml:"Length,omitempty" require:"true"`
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

func (s *AbstractFilmVideoAdvanceRequest) SetLength(v int) *AbstractFilmVideoAdvanceRequest {
	s.Length = &v
	return s
}

type AdjustVideoColorRequest struct {
	VideoUrl     *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
	VideoBitrate *string `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	VideoCodec   *string `json:"VideoCodec,omitempty" xml:"VideoCodec,omitempty"`
	VideoFormat  *string `json:"VideoFormat,omitempty" xml:"VideoFormat,omitempty"`
	Mode         *string `json:"Mode,omitempty" xml:"Mode,omitempty" require:"true"`
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

type AdjustVideoColorResponse struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *AdjustVideoColorResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s AdjustVideoColorResponse) String() string {
	return tea.Prettify(s)
}

func (s AdjustVideoColorResponse) GoString() string {
	return s.String()
}

func (s *AdjustVideoColorResponse) SetRequestId(v string) *AdjustVideoColorResponse {
	s.RequestId = &v
	return s
}

func (s *AdjustVideoColorResponse) SetData(v *AdjustVideoColorResponseData) *AdjustVideoColorResponse {
	s.Data = v
	return s
}

type AdjustVideoColorResponseData struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
}

func (s AdjustVideoColorResponseData) String() string {
	return tea.Prettify(s)
}

func (s AdjustVideoColorResponseData) GoString() string {
	return s.String()
}

func (s *AdjustVideoColorResponseData) SetVideoUrl(v string) *AdjustVideoColorResponseData {
	s.VideoUrl = &v
	return s
}

type AdjustVideoColorAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrlObject,omitempty" xml:"VideoUrlObject,omitempty" require:"true"`
	VideoBitrate   *string   `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	VideoCodec     *string   `json:"VideoCodec,omitempty" xml:"VideoCodec,omitempty"`
	VideoFormat    *string   `json:"VideoFormat,omitempty" xml:"VideoFormat,omitempty"`
	Mode           *string   `json:"Mode,omitempty" xml:"Mode,omitempty" require:"true"`
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

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
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

func (client *Client) ConvertHdrVideo(request *ConvertHdrVideoRequest, runtime *util.RuntimeOptions) (_result *ConvertHdrVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ConvertHdrVideoResponse{}
	_body, _err := client.DoRequest(tea.String("ConvertHdrVideo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConvertHdrVideoSimply(request *ConvertHdrVideoRequest) (_result *ConvertHdrVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConvertHdrVideoResponse{}
	_body, _err := client.ConvertHdrVideo(request, runtime)
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
	rpcutil.Convert(runtime, ossRuntime)
	convertHdrVideoReq := &ConvertHdrVideoRequest{}
	rpcutil.Convert(request, convertHdrVideoReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
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
	convertHdrVideoReq.VideoURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	convertHdrVideoResp, _err := client.ConvertHdrVideo(convertHdrVideoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = convertHdrVideoResp
	return _result, _err
}

func (client *Client) InterpolateVideoFrame(request *InterpolateVideoFrameRequest, runtime *util.RuntimeOptions) (_result *InterpolateVideoFrameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InterpolateVideoFrameResponse{}
	_body, _err := client.DoRequest(tea.String("InterpolateVideoFrame"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InterpolateVideoFrameSimply(request *InterpolateVideoFrameRequest) (_result *InterpolateVideoFrameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InterpolateVideoFrameResponse{}
	_body, _err := client.InterpolateVideoFrame(request, runtime)
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
	rpcutil.Convert(runtime, ossRuntime)
	interpolateVideoFrameReq := &InterpolateVideoFrameRequest{}
	rpcutil.Convert(request, interpolateVideoFrameReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
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
	interpolateVideoFrameReq.VideoURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	interpolateVideoFrameResp, _err := client.InterpolateVideoFrame(interpolateVideoFrameReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = interpolateVideoFrameResp
	return _result, _err
}

func (client *Client) ToneSdrVideo(request *ToneSdrVideoRequest, runtime *util.RuntimeOptions) (_result *ToneSdrVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ToneSdrVideoResponse{}
	_body, _err := client.DoRequest(tea.String("ToneSdrVideo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ToneSdrVideoSimply(request *ToneSdrVideoRequest) (_result *ToneSdrVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ToneSdrVideoResponse{}
	_body, _err := client.ToneSdrVideo(request, runtime)
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
	rpcutil.Convert(runtime, ossRuntime)
	toneSdrVideoReq := &ToneSdrVideoRequest{}
	rpcutil.Convert(request, toneSdrVideoReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
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
	toneSdrVideoReq.VideoURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	toneSdrVideoResp, _err := client.ToneSdrVideo(toneSdrVideoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = toneSdrVideoResp
	return _result, _err
}

func (client *Client) EnhanceVideoQuality(request *EnhanceVideoQualityRequest, runtime *util.RuntimeOptions) (_result *EnhanceVideoQualityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &EnhanceVideoQualityResponse{}
	_body, _err := client.DoRequest(tea.String("EnhanceVideoQuality"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnhanceVideoQualitySimply(request *EnhanceVideoQualityRequest) (_result *EnhanceVideoQualityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnhanceVideoQualityResponse{}
	_body, _err := client.EnhanceVideoQuality(request, runtime)
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
	rpcutil.Convert(runtime, ossRuntime)
	enhanceVideoQualityReq := &EnhanceVideoQualityRequest{}
	rpcutil.Convert(request, enhanceVideoQualityReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
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
	enhanceVideoQualityResp, _err := client.EnhanceVideoQuality(enhanceVideoQualityReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = enhanceVideoQualityResp
	return _result, _err
}

func (client *Client) MergeVideoFace(request *MergeVideoFaceRequest, runtime *util.RuntimeOptions) (_result *MergeVideoFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &MergeVideoFaceResponse{}
	_body, _err := client.DoRequest(tea.String("MergeVideoFace"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MergeVideoFaceSimply(request *MergeVideoFaceRequest) (_result *MergeVideoFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MergeVideoFaceResponse{}
	_body, _err := client.MergeVideoFace(request, runtime)
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
	rpcutil.Convert(runtime, ossRuntime)
	mergeVideoFaceReq := &MergeVideoFaceRequest{}
	rpcutil.Convert(request, mergeVideoFaceReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
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
	mergeVideoFaceResp, _err := client.MergeVideoFace(mergeVideoFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = mergeVideoFaceResp
	return _result, _err
}

func (client *Client) ChangeVideoSize(request *ChangeVideoSizeRequest, runtime *util.RuntimeOptions) (_result *ChangeVideoSizeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ChangeVideoSizeResponse{}
	_body, _err := client.DoRequest(tea.String("ChangeVideoSize"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeVideoSizeSimply(request *ChangeVideoSizeRequest) (_result *ChangeVideoSizeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeVideoSizeResponse{}
	_body, _err := client.ChangeVideoSize(request, runtime)
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
	rpcutil.Convert(runtime, ossRuntime)
	changeVideoSizeReq := &ChangeVideoSizeRequest{}
	rpcutil.Convert(request, changeVideoSizeReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
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
	changeVideoSizeResp, _err := client.ChangeVideoSize(changeVideoSizeReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = changeVideoSizeResp
	return _result, _err
}

func (client *Client) GenerateVideo(request *GenerateVideoRequest, runtime *util.RuntimeOptions) (_result *GenerateVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GenerateVideoResponse{}
	_body, _err := client.DoRequest(tea.String("GenerateVideo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateVideoSimply(request *GenerateVideoRequest) (_result *GenerateVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateVideoResponse{}
	_body, _err := client.GenerateVideo(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAsyncJobResult(request *GetAsyncJobResultRequest, runtime *util.RuntimeOptions) (_result *GetAsyncJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.DoRequest(tea.String("GetAsyncJobResult"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncJobResultSimply(request *GetAsyncJobResultRequest) (_result *GetAsyncJobResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.GetAsyncJobResult(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SuperResolveVideo(request *SuperResolveVideoRequest, runtime *util.RuntimeOptions) (_result *SuperResolveVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SuperResolveVideoResponse{}
	_body, _err := client.DoRequest(tea.String("SuperResolveVideo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SuperResolveVideoSimply(request *SuperResolveVideoRequest) (_result *SuperResolveVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SuperResolveVideoResponse{}
	_body, _err := client.SuperResolveVideo(request, runtime)
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
	rpcutil.Convert(runtime, ossRuntime)
	superResolveVideoReq := &SuperResolveVideoRequest{}
	rpcutil.Convert(request, superResolveVideoReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
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
	superResolveVideoResp, _err := client.SuperResolveVideo(superResolveVideoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = superResolveVideoResp
	return _result, _err
}

func (client *Client) EraseVideoLogo(request *EraseVideoLogoRequest, runtime *util.RuntimeOptions) (_result *EraseVideoLogoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &EraseVideoLogoResponse{}
	_body, _err := client.DoRequest(tea.String("EraseVideoLogo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EraseVideoLogoSimply(request *EraseVideoLogoRequest) (_result *EraseVideoLogoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EraseVideoLogoResponse{}
	_body, _err := client.EraseVideoLogo(request, runtime)
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
	rpcutil.Convert(runtime, ossRuntime)
	eraseVideoLogoReq := &EraseVideoLogoRequest{}
	rpcutil.Convert(request, eraseVideoLogoReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
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
	eraseVideoLogoResp, _err := client.EraseVideoLogo(eraseVideoLogoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = eraseVideoLogoResp
	return _result, _err
}

func (client *Client) EraseVideoSubtitles(request *EraseVideoSubtitlesRequest, runtime *util.RuntimeOptions) (_result *EraseVideoSubtitlesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &EraseVideoSubtitlesResponse{}
	_body, _err := client.DoRequest(tea.String("EraseVideoSubtitles"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EraseVideoSubtitlesSimply(request *EraseVideoSubtitlesRequest) (_result *EraseVideoSubtitlesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EraseVideoSubtitlesResponse{}
	_body, _err := client.EraseVideoSubtitles(request, runtime)
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
	rpcutil.Convert(runtime, ossRuntime)
	eraseVideoSubtitlesReq := &EraseVideoSubtitlesRequest{}
	rpcutil.Convert(request, eraseVideoSubtitlesReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
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
	eraseVideoSubtitlesResp, _err := client.EraseVideoSubtitles(eraseVideoSubtitlesReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = eraseVideoSubtitlesResp
	return _result, _err
}

func (client *Client) AbstractEcommerceVideo(request *AbstractEcommerceVideoRequest, runtime *util.RuntimeOptions) (_result *AbstractEcommerceVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AbstractEcommerceVideoResponse{}
	_body, _err := client.DoRequest(tea.String("AbstractEcommerceVideo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AbstractEcommerceVideoSimply(request *AbstractEcommerceVideoRequest) (_result *AbstractEcommerceVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AbstractEcommerceVideoResponse{}
	_body, _err := client.AbstractEcommerceVideo(request, runtime)
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
	rpcutil.Convert(runtime, ossRuntime)
	abstractEcommerceVideoReq := &AbstractEcommerceVideoRequest{}
	rpcutil.Convert(request, abstractEcommerceVideoReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
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
	abstractEcommerceVideoResp, _err := client.AbstractEcommerceVideo(abstractEcommerceVideoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = abstractEcommerceVideoResp
	return _result, _err
}

func (client *Client) AbstractFilmVideo(request *AbstractFilmVideoRequest, runtime *util.RuntimeOptions) (_result *AbstractFilmVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AbstractFilmVideoResponse{}
	_body, _err := client.DoRequest(tea.String("AbstractFilmVideo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AbstractFilmVideoSimply(request *AbstractFilmVideoRequest) (_result *AbstractFilmVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AbstractFilmVideoResponse{}
	_body, _err := client.AbstractFilmVideo(request, runtime)
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
	rpcutil.Convert(runtime, ossRuntime)
	abstractFilmVideoReq := &AbstractFilmVideoRequest{}
	rpcutil.Convert(request, abstractFilmVideoReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
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
	abstractFilmVideoResp, _err := client.AbstractFilmVideo(abstractFilmVideoReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = abstractFilmVideoResp
	return _result, _err
}

func (client *Client) AdjustVideoColor(request *AdjustVideoColorRequest, runtime *util.RuntimeOptions) (_result *AdjustVideoColorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AdjustVideoColorResponse{}
	_body, _err := client.DoRequest(tea.String("AdjustVideoColor"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-03-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AdjustVideoColorSimply(request *AdjustVideoColorRequest) (_result *AdjustVideoColorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AdjustVideoColorResponse{}
	_body, _err := client.AdjustVideoColor(request, runtime)
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
	rpcutil.Convert(runtime, ossRuntime)
	adjustVideoColorReq := &AdjustVideoColorRequest{}
	rpcutil.Convert(request, adjustVideoColorReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
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
	adjustVideoColorResp, _err := client.AdjustVideoColor(adjustVideoColorReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = adjustVideoColorResp
	return _result, _err
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
