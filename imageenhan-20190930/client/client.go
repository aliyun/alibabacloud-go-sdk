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

type ExtendImageStyleRequest struct {
	StyleUrl *string `json:"StyleUrl,omitempty" xml:"StyleUrl,omitempty"`
	MajorUrl *string `json:"MajorUrl,omitempty" xml:"MajorUrl,omitempty"`
}

func (s ExtendImageStyleRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtendImageStyleRequest) GoString() string {
	return s.String()
}

func (s *ExtendImageStyleRequest) SetStyleUrl(v string) *ExtendImageStyleRequest {
	s.StyleUrl = &v
	return s
}

func (s *ExtendImageStyleRequest) SetMajorUrl(v string) *ExtendImageStyleRequest {
	s.MajorUrl = &v
	return s
}

type ExtendImageStyleResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ExtendImageStyleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ExtendImageStyleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExtendImageStyleResponseBody) GoString() string {
	return s.String()
}

func (s *ExtendImageStyleResponseBody) SetRequestId(v string) *ExtendImageStyleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExtendImageStyleResponseBody) SetData(v *ExtendImageStyleResponseBodyData) *ExtendImageStyleResponseBody {
	s.Data = v
	return s
}

type ExtendImageStyleResponseBodyData struct {
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
	MajorUrl *string `json:"MajorUrl,omitempty" xml:"MajorUrl,omitempty"`
}

func (s ExtendImageStyleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExtendImageStyleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExtendImageStyleResponseBodyData) SetUrl(v string) *ExtendImageStyleResponseBodyData {
	s.Url = &v
	return s
}

func (s *ExtendImageStyleResponseBodyData) SetMajorUrl(v string) *ExtendImageStyleResponseBodyData {
	s.MajorUrl = &v
	return s
}

type ExtendImageStyleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExtendImageStyleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExtendImageStyleResponse) String() string {
	return tea.Prettify(s)
}

func (s ExtendImageStyleResponse) GoString() string {
	return s.String()
}

func (s *ExtendImageStyleResponse) SetHeaders(v map[string]*string) *ExtendImageStyleResponse {
	s.Headers = v
	return s
}

func (s *ExtendImageStyleResponse) SetBody(v *ExtendImageStyleResponseBody) *ExtendImageStyleResponse {
	s.Body = v
	return s
}

type ImageBlindCharacterWatermarkRequest struct {
	FunctionType      *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	Text              *string `json:"Text,omitempty" xml:"Text,omitempty"`
	WatermarkImageURL *string `json:"WatermarkImageURL,omitempty" xml:"WatermarkImageURL,omitempty"`
	OutputFileType    *string `json:"OutputFileType,omitempty" xml:"OutputFileType,omitempty"`
	QualityFactor     *int32  `json:"QualityFactor,omitempty" xml:"QualityFactor,omitempty"`
	OriginImageURL    *string `json:"OriginImageURL,omitempty" xml:"OriginImageURL,omitempty"`
}

func (s ImageBlindCharacterWatermarkRequest) String() string {
	return tea.Prettify(s)
}

func (s ImageBlindCharacterWatermarkRequest) GoString() string {
	return s.String()
}

func (s *ImageBlindCharacterWatermarkRequest) SetFunctionType(v string) *ImageBlindCharacterWatermarkRequest {
	s.FunctionType = &v
	return s
}

func (s *ImageBlindCharacterWatermarkRequest) SetText(v string) *ImageBlindCharacterWatermarkRequest {
	s.Text = &v
	return s
}

func (s *ImageBlindCharacterWatermarkRequest) SetWatermarkImageURL(v string) *ImageBlindCharacterWatermarkRequest {
	s.WatermarkImageURL = &v
	return s
}

func (s *ImageBlindCharacterWatermarkRequest) SetOutputFileType(v string) *ImageBlindCharacterWatermarkRequest {
	s.OutputFileType = &v
	return s
}

func (s *ImageBlindCharacterWatermarkRequest) SetQualityFactor(v int32) *ImageBlindCharacterWatermarkRequest {
	s.QualityFactor = &v
	return s
}

func (s *ImageBlindCharacterWatermarkRequest) SetOriginImageURL(v string) *ImageBlindCharacterWatermarkRequest {
	s.OriginImageURL = &v
	return s
}

type ImageBlindCharacterWatermarkAdvanceRequest struct {
	OriginImageURLObject io.Reader `json:"OriginImageURLObject,omitempty" xml:"OriginImageURLObject,omitempty" require:"true"`
	FunctionType         *string   `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	Text                 *string   `json:"Text,omitempty" xml:"Text,omitempty"`
	WatermarkImageURL    *string   `json:"WatermarkImageURL,omitempty" xml:"WatermarkImageURL,omitempty"`
	OutputFileType       *string   `json:"OutputFileType,omitempty" xml:"OutputFileType,omitempty"`
	QualityFactor        *int32    `json:"QualityFactor,omitempty" xml:"QualityFactor,omitempty"`
}

func (s ImageBlindCharacterWatermarkAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ImageBlindCharacterWatermarkAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) SetOriginImageURLObject(v io.Reader) *ImageBlindCharacterWatermarkAdvanceRequest {
	s.OriginImageURLObject = v
	return s
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) SetFunctionType(v string) *ImageBlindCharacterWatermarkAdvanceRequest {
	s.FunctionType = &v
	return s
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) SetText(v string) *ImageBlindCharacterWatermarkAdvanceRequest {
	s.Text = &v
	return s
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) SetWatermarkImageURL(v string) *ImageBlindCharacterWatermarkAdvanceRequest {
	s.WatermarkImageURL = &v
	return s
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) SetOutputFileType(v string) *ImageBlindCharacterWatermarkAdvanceRequest {
	s.OutputFileType = &v
	return s
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) SetQualityFactor(v int32) *ImageBlindCharacterWatermarkAdvanceRequest {
	s.QualityFactor = &v
	return s
}

type ImageBlindCharacterWatermarkResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ImageBlindCharacterWatermarkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ImageBlindCharacterWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImageBlindCharacterWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *ImageBlindCharacterWatermarkResponseBody) SetRequestId(v string) *ImageBlindCharacterWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageBlindCharacterWatermarkResponseBody) SetData(v *ImageBlindCharacterWatermarkResponseBodyData) *ImageBlindCharacterWatermarkResponseBody {
	s.Data = v
	return s
}

type ImageBlindCharacterWatermarkResponseBodyData struct {
	WatermarkImageURL *string `json:"WatermarkImageURL,omitempty" xml:"WatermarkImageURL,omitempty"`
	TextImageURL      *string `json:"TextImageURL,omitempty" xml:"TextImageURL,omitempty"`
}

func (s ImageBlindCharacterWatermarkResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ImageBlindCharacterWatermarkResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageBlindCharacterWatermarkResponseBodyData) SetWatermarkImageURL(v string) *ImageBlindCharacterWatermarkResponseBodyData {
	s.WatermarkImageURL = &v
	return s
}

func (s *ImageBlindCharacterWatermarkResponseBodyData) SetTextImageURL(v string) *ImageBlindCharacterWatermarkResponseBodyData {
	s.TextImageURL = &v
	return s
}

type ImageBlindCharacterWatermarkResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImageBlindCharacterWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImageBlindCharacterWatermarkResponse) String() string {
	return tea.Prettify(s)
}

func (s ImageBlindCharacterWatermarkResponse) GoString() string {
	return s.String()
}

func (s *ImageBlindCharacterWatermarkResponse) SetHeaders(v map[string]*string) *ImageBlindCharacterWatermarkResponse {
	s.Headers = v
	return s
}

func (s *ImageBlindCharacterWatermarkResponse) SetBody(v *ImageBlindCharacterWatermarkResponseBody) *ImageBlindCharacterWatermarkResponse {
	s.Body = v
	return s
}

type RemoveImageWatermarkRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RemoveImageWatermarkRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageWatermarkRequest) GoString() string {
	return s.String()
}

func (s *RemoveImageWatermarkRequest) SetImageURL(v string) *RemoveImageWatermarkRequest {
	s.ImageURL = &v
	return s
}

type RemoveImageWatermarkAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s RemoveImageWatermarkAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageWatermarkAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RemoveImageWatermarkAdvanceRequest) SetImageURLObject(v io.Reader) *RemoveImageWatermarkAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RemoveImageWatermarkResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RemoveImageWatermarkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RemoveImageWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveImageWatermarkResponseBody) SetRequestId(v string) *RemoveImageWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveImageWatermarkResponseBody) SetData(v *RemoveImageWatermarkResponseBodyData) *RemoveImageWatermarkResponseBody {
	s.Data = v
	return s
}

type RemoveImageWatermarkResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RemoveImageWatermarkResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageWatermarkResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemoveImageWatermarkResponseBodyData) SetImageURL(v string) *RemoveImageWatermarkResponseBodyData {
	s.ImageURL = &v
	return s
}

type RemoveImageWatermarkResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveImageWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveImageWatermarkResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageWatermarkResponse) GoString() string {
	return s.String()
}

func (s *RemoveImageWatermarkResponse) SetHeaders(v map[string]*string) *RemoveImageWatermarkResponse {
	s.Headers = v
	return s
}

func (s *RemoveImageWatermarkResponse) SetBody(v *RemoveImageWatermarkResponseBody) *RemoveImageWatermarkResponse {
	s.Body = v
	return s
}

type GenerateDynamicImageRequest struct {
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
}

func (s GenerateDynamicImageRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDynamicImageRequest) GoString() string {
	return s.String()
}

func (s *GenerateDynamicImageRequest) SetUrl(v string) *GenerateDynamicImageRequest {
	s.Url = &v
	return s
}

func (s *GenerateDynamicImageRequest) SetOperation(v string) *GenerateDynamicImageRequest {
	s.Operation = &v
	return s
}

type GenerateDynamicImageAdvanceRequest struct {
	UrlObject io.Reader `json:"UrlObject,omitempty" xml:"UrlObject,omitempty" require:"true"`
	Operation *string   `json:"Operation,omitempty" xml:"Operation,omitempty"`
}

func (s GenerateDynamicImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDynamicImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateDynamicImageAdvanceRequest) SetUrlObject(v io.Reader) *GenerateDynamicImageAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *GenerateDynamicImageAdvanceRequest) SetOperation(v string) *GenerateDynamicImageAdvanceRequest {
	s.Operation = &v
	return s
}

type GenerateDynamicImageResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GenerateDynamicImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GenerateDynamicImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateDynamicImageResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDynamicImageResponseBody) SetRequestId(v string) *GenerateDynamicImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateDynamicImageResponseBody) SetData(v *GenerateDynamicImageResponseBodyData) *GenerateDynamicImageResponseBody {
	s.Data = v
	return s
}

type GenerateDynamicImageResponseBodyData struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GenerateDynamicImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateDynamicImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateDynamicImageResponseBodyData) SetUrl(v string) *GenerateDynamicImageResponseBodyData {
	s.Url = &v
	return s
}

type GenerateDynamicImageResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateDynamicImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateDynamicImageResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateDynamicImageResponse) GoString() string {
	return s.String()
}

func (s *GenerateDynamicImageResponse) SetHeaders(v map[string]*string) *GenerateDynamicImageResponse {
	s.Headers = v
	return s
}

func (s *GenerateDynamicImageResponse) SetBody(v *GenerateDynamicImageResponseBody) *GenerateDynamicImageResponse {
	s.Body = v
	return s
}

type ImageBlindPicWatermarkRequest struct {
	FunctionType      *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	LogoURL           *string `json:"LogoURL,omitempty" xml:"LogoURL,omitempty"`
	WatermarkImageURL *string `json:"WatermarkImageURL,omitempty" xml:"WatermarkImageURL,omitempty"`
	OutputFileType    *string `json:"OutputFileType,omitempty" xml:"OutputFileType,omitempty"`
	QualityFactor     *int32  `json:"QualityFactor,omitempty" xml:"QualityFactor,omitempty"`
	OriginImageURL    *string `json:"OriginImageURL,omitempty" xml:"OriginImageURL,omitempty"`
}

func (s ImageBlindPicWatermarkRequest) String() string {
	return tea.Prettify(s)
}

func (s ImageBlindPicWatermarkRequest) GoString() string {
	return s.String()
}

func (s *ImageBlindPicWatermarkRequest) SetFunctionType(v string) *ImageBlindPicWatermarkRequest {
	s.FunctionType = &v
	return s
}

func (s *ImageBlindPicWatermarkRequest) SetLogoURL(v string) *ImageBlindPicWatermarkRequest {
	s.LogoURL = &v
	return s
}

func (s *ImageBlindPicWatermarkRequest) SetWatermarkImageURL(v string) *ImageBlindPicWatermarkRequest {
	s.WatermarkImageURL = &v
	return s
}

func (s *ImageBlindPicWatermarkRequest) SetOutputFileType(v string) *ImageBlindPicWatermarkRequest {
	s.OutputFileType = &v
	return s
}

func (s *ImageBlindPicWatermarkRequest) SetQualityFactor(v int32) *ImageBlindPicWatermarkRequest {
	s.QualityFactor = &v
	return s
}

func (s *ImageBlindPicWatermarkRequest) SetOriginImageURL(v string) *ImageBlindPicWatermarkRequest {
	s.OriginImageURL = &v
	return s
}

type ImageBlindPicWatermarkAdvanceRequest struct {
	OriginImageURLObject io.Reader `json:"OriginImageURLObject,omitempty" xml:"OriginImageURLObject,omitempty" require:"true"`
	FunctionType         *string   `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	LogoURL              *string   `json:"LogoURL,omitempty" xml:"LogoURL,omitempty"`
	WatermarkImageURL    *string   `json:"WatermarkImageURL,omitempty" xml:"WatermarkImageURL,omitempty"`
	OutputFileType       *string   `json:"OutputFileType,omitempty" xml:"OutputFileType,omitempty"`
	QualityFactor        *int32    `json:"QualityFactor,omitempty" xml:"QualityFactor,omitempty"`
}

func (s ImageBlindPicWatermarkAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ImageBlindPicWatermarkAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ImageBlindPicWatermarkAdvanceRequest) SetOriginImageURLObject(v io.Reader) *ImageBlindPicWatermarkAdvanceRequest {
	s.OriginImageURLObject = v
	return s
}

func (s *ImageBlindPicWatermarkAdvanceRequest) SetFunctionType(v string) *ImageBlindPicWatermarkAdvanceRequest {
	s.FunctionType = &v
	return s
}

func (s *ImageBlindPicWatermarkAdvanceRequest) SetLogoURL(v string) *ImageBlindPicWatermarkAdvanceRequest {
	s.LogoURL = &v
	return s
}

func (s *ImageBlindPicWatermarkAdvanceRequest) SetWatermarkImageURL(v string) *ImageBlindPicWatermarkAdvanceRequest {
	s.WatermarkImageURL = &v
	return s
}

func (s *ImageBlindPicWatermarkAdvanceRequest) SetOutputFileType(v string) *ImageBlindPicWatermarkAdvanceRequest {
	s.OutputFileType = &v
	return s
}

func (s *ImageBlindPicWatermarkAdvanceRequest) SetQualityFactor(v int32) *ImageBlindPicWatermarkAdvanceRequest {
	s.QualityFactor = &v
	return s
}

type ImageBlindPicWatermarkResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ImageBlindPicWatermarkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ImageBlindPicWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImageBlindPicWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *ImageBlindPicWatermarkResponseBody) SetRequestId(v string) *ImageBlindPicWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageBlindPicWatermarkResponseBody) SetData(v *ImageBlindPicWatermarkResponseBodyData) *ImageBlindPicWatermarkResponseBody {
	s.Data = v
	return s
}

type ImageBlindPicWatermarkResponseBodyData struct {
	WatermarkImageURL *string `json:"WatermarkImageURL,omitempty" xml:"WatermarkImageURL,omitempty"`
	LogoURL           *string `json:"LogoURL,omitempty" xml:"LogoURL,omitempty"`
}

func (s ImageBlindPicWatermarkResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ImageBlindPicWatermarkResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageBlindPicWatermarkResponseBodyData) SetWatermarkImageURL(v string) *ImageBlindPicWatermarkResponseBodyData {
	s.WatermarkImageURL = &v
	return s
}

func (s *ImageBlindPicWatermarkResponseBodyData) SetLogoURL(v string) *ImageBlindPicWatermarkResponseBodyData {
	s.LogoURL = &v
	return s
}

type ImageBlindPicWatermarkResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImageBlindPicWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImageBlindPicWatermarkResponse) String() string {
	return tea.Prettify(s)
}

func (s ImageBlindPicWatermarkResponse) GoString() string {
	return s.String()
}

func (s *ImageBlindPicWatermarkResponse) SetHeaders(v map[string]*string) *ImageBlindPicWatermarkResponse {
	s.Headers = v
	return s
}

func (s *ImageBlindPicWatermarkResponse) SetBody(v *ImageBlindPicWatermarkResponseBody) *ImageBlindPicWatermarkResponse {
	s.Body = v
	return s
}

type RemoveImageSubtitlesRequest struct {
	ImageURL *string  `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	BX       *float32 `json:"BX,omitempty" xml:"BX,omitempty"`
	BY       *float32 `json:"BY,omitempty" xml:"BY,omitempty"`
	BW       *float32 `json:"BW,omitempty" xml:"BW,omitempty"`
	BH       *float32 `json:"BH,omitempty" xml:"BH,omitempty"`
}

func (s RemoveImageSubtitlesRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageSubtitlesRequest) GoString() string {
	return s.String()
}

func (s *RemoveImageSubtitlesRequest) SetImageURL(v string) *RemoveImageSubtitlesRequest {
	s.ImageURL = &v
	return s
}

func (s *RemoveImageSubtitlesRequest) SetBX(v float32) *RemoveImageSubtitlesRequest {
	s.BX = &v
	return s
}

func (s *RemoveImageSubtitlesRequest) SetBY(v float32) *RemoveImageSubtitlesRequest {
	s.BY = &v
	return s
}

func (s *RemoveImageSubtitlesRequest) SetBW(v float32) *RemoveImageSubtitlesRequest {
	s.BW = &v
	return s
}

func (s *RemoveImageSubtitlesRequest) SetBH(v float32) *RemoveImageSubtitlesRequest {
	s.BH = &v
	return s
}

type RemoveImageSubtitlesAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	BX             *float32  `json:"BX,omitempty" xml:"BX,omitempty"`
	BY             *float32  `json:"BY,omitempty" xml:"BY,omitempty"`
	BW             *float32  `json:"BW,omitempty" xml:"BW,omitempty"`
	BH             *float32  `json:"BH,omitempty" xml:"BH,omitempty"`
}

func (s RemoveImageSubtitlesAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageSubtitlesAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RemoveImageSubtitlesAdvanceRequest) SetImageURLObject(v io.Reader) *RemoveImageSubtitlesAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RemoveImageSubtitlesAdvanceRequest) SetBX(v float32) *RemoveImageSubtitlesAdvanceRequest {
	s.BX = &v
	return s
}

func (s *RemoveImageSubtitlesAdvanceRequest) SetBY(v float32) *RemoveImageSubtitlesAdvanceRequest {
	s.BY = &v
	return s
}

func (s *RemoveImageSubtitlesAdvanceRequest) SetBW(v float32) *RemoveImageSubtitlesAdvanceRequest {
	s.BW = &v
	return s
}

func (s *RemoveImageSubtitlesAdvanceRequest) SetBH(v float32) *RemoveImageSubtitlesAdvanceRequest {
	s.BH = &v
	return s
}

type RemoveImageSubtitlesResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RemoveImageSubtitlesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RemoveImageSubtitlesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageSubtitlesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveImageSubtitlesResponseBody) SetRequestId(v string) *RemoveImageSubtitlesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveImageSubtitlesResponseBody) SetData(v *RemoveImageSubtitlesResponseBodyData) *RemoveImageSubtitlesResponseBody {
	s.Data = v
	return s
}

type RemoveImageSubtitlesResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RemoveImageSubtitlesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageSubtitlesResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemoveImageSubtitlesResponseBodyData) SetImageURL(v string) *RemoveImageSubtitlesResponseBodyData {
	s.ImageURL = &v
	return s
}

type RemoveImageSubtitlesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveImageSubtitlesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveImageSubtitlesResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageSubtitlesResponse) GoString() string {
	return s.String()
}

func (s *RemoveImageSubtitlesResponse) SetHeaders(v map[string]*string) *RemoveImageSubtitlesResponse {
	s.Headers = v
	return s
}

func (s *RemoveImageSubtitlesResponse) SetBody(v *RemoveImageSubtitlesResponseBody) *RemoveImageSubtitlesResponse {
	s.Body = v
	return s
}

type RecolorHDImageRequest struct {
	Url           *string                               `json:"Url,omitempty" xml:"Url,omitempty"`
	Mode          *string                               `json:"Mode,omitempty" xml:"Mode,omitempty"`
	RefUrl        *string                               `json:"RefUrl,omitempty" xml:"RefUrl,omitempty"`
	ColorCount    *int32                                `json:"ColorCount,omitempty" xml:"ColorCount,omitempty"`
	Degree        *string                               `json:"Degree,omitempty" xml:"Degree,omitempty"`
	Async         *bool                                 `json:"Async,omitempty" xml:"Async,omitempty"`
	ColorTemplate []*RecolorHDImageRequestColorTemplate `json:"ColorTemplate,omitempty" xml:"ColorTemplate,omitempty" type:"Repeated"`
}

func (s RecolorHDImageRequest) String() string {
	return tea.Prettify(s)
}

func (s RecolorHDImageRequest) GoString() string {
	return s.String()
}

func (s *RecolorHDImageRequest) SetUrl(v string) *RecolorHDImageRequest {
	s.Url = &v
	return s
}

func (s *RecolorHDImageRequest) SetMode(v string) *RecolorHDImageRequest {
	s.Mode = &v
	return s
}

func (s *RecolorHDImageRequest) SetRefUrl(v string) *RecolorHDImageRequest {
	s.RefUrl = &v
	return s
}

func (s *RecolorHDImageRequest) SetColorCount(v int32) *RecolorHDImageRequest {
	s.ColorCount = &v
	return s
}

func (s *RecolorHDImageRequest) SetDegree(v string) *RecolorHDImageRequest {
	s.Degree = &v
	return s
}

func (s *RecolorHDImageRequest) SetAsync(v bool) *RecolorHDImageRequest {
	s.Async = &v
	return s
}

func (s *RecolorHDImageRequest) SetColorTemplate(v []*RecolorHDImageRequestColorTemplate) *RecolorHDImageRequest {
	s.ColorTemplate = v
	return s
}

type RecolorHDImageRequestColorTemplate struct {
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
}

func (s RecolorHDImageRequestColorTemplate) String() string {
	return tea.Prettify(s)
}

func (s RecolorHDImageRequestColorTemplate) GoString() string {
	return s.String()
}

func (s *RecolorHDImageRequestColorTemplate) SetColor(v string) *RecolorHDImageRequestColorTemplate {
	s.Color = &v
	return s
}

type RecolorHDImageAdvanceRequest struct {
	UrlObject     io.Reader                                    `json:"UrlObject,omitempty" xml:"UrlObject,omitempty" require:"true"`
	Mode          *string                                      `json:"Mode,omitempty" xml:"Mode,omitempty"`
	RefUrl        *string                                      `json:"RefUrl,omitempty" xml:"RefUrl,omitempty"`
	ColorCount    *int32                                       `json:"ColorCount,omitempty" xml:"ColorCount,omitempty"`
	Degree        *string                                      `json:"Degree,omitempty" xml:"Degree,omitempty"`
	Async         *bool                                        `json:"Async,omitempty" xml:"Async,omitempty"`
	ColorTemplate []*RecolorHDImageAdvanceRequestColorTemplate `json:"ColorTemplate,omitempty" xml:"ColorTemplate,omitempty" type:"Repeated"`
}

func (s RecolorHDImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecolorHDImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecolorHDImageAdvanceRequest) SetUrlObject(v io.Reader) *RecolorHDImageAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *RecolorHDImageAdvanceRequest) SetMode(v string) *RecolorHDImageAdvanceRequest {
	s.Mode = &v
	return s
}

func (s *RecolorHDImageAdvanceRequest) SetRefUrl(v string) *RecolorHDImageAdvanceRequest {
	s.RefUrl = &v
	return s
}

func (s *RecolorHDImageAdvanceRequest) SetColorCount(v int32) *RecolorHDImageAdvanceRequest {
	s.ColorCount = &v
	return s
}

func (s *RecolorHDImageAdvanceRequest) SetDegree(v string) *RecolorHDImageAdvanceRequest {
	s.Degree = &v
	return s
}

func (s *RecolorHDImageAdvanceRequest) SetAsync(v bool) *RecolorHDImageAdvanceRequest {
	s.Async = &v
	return s
}

func (s *RecolorHDImageAdvanceRequest) SetColorTemplate(v []*RecolorHDImageAdvanceRequestColorTemplate) *RecolorHDImageAdvanceRequest {
	s.ColorTemplate = v
	return s
}

type RecolorHDImageAdvanceRequestColorTemplate struct {
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
}

func (s RecolorHDImageAdvanceRequestColorTemplate) String() string {
	return tea.Prettify(s)
}

func (s RecolorHDImageAdvanceRequestColorTemplate) GoString() string {
	return s.String()
}

func (s *RecolorHDImageAdvanceRequestColorTemplate) SetColor(v string) *RecolorHDImageAdvanceRequestColorTemplate {
	s.Color = &v
	return s
}

type RecolorHDImageResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecolorHDImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecolorHDImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecolorHDImageResponseBody) GoString() string {
	return s.String()
}

func (s *RecolorHDImageResponseBody) SetRequestId(v string) *RecolorHDImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecolorHDImageResponseBody) SetData(v *RecolorHDImageResponseBodyData) *RecolorHDImageResponseBody {
	s.Data = v
	return s
}

type RecolorHDImageResponseBodyData struct {
	ImageList []*string `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
}

func (s RecolorHDImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecolorHDImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecolorHDImageResponseBodyData) SetImageList(v []*string) *RecolorHDImageResponseBodyData {
	s.ImageList = v
	return s
}

type RecolorHDImageResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecolorHDImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecolorHDImageResponse) String() string {
	return tea.Prettify(s)
}

func (s RecolorHDImageResponse) GoString() string {
	return s.String()
}

func (s *RecolorHDImageResponse) SetHeaders(v map[string]*string) *RecolorHDImageResponse {
	s.Headers = v
	return s
}

func (s *RecolorHDImageResponse) SetBody(v *RecolorHDImageResponseBody) *RecolorHDImageResponse {
	s.Body = v
	return s
}

type ColorizeImageRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ColorizeImageRequest) String() string {
	return tea.Prettify(s)
}

func (s ColorizeImageRequest) GoString() string {
	return s.String()
}

func (s *ColorizeImageRequest) SetImageURL(v string) *ColorizeImageRequest {
	s.ImageURL = &v
	return s
}

type ColorizeImageAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s ColorizeImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ColorizeImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ColorizeImageAdvanceRequest) SetImageURLObject(v io.Reader) *ColorizeImageAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type ColorizeImageResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ColorizeImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ColorizeImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ColorizeImageResponseBody) GoString() string {
	return s.String()
}

func (s *ColorizeImageResponseBody) SetRequestId(v string) *ColorizeImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ColorizeImageResponseBody) SetData(v *ColorizeImageResponseBodyData) *ColorizeImageResponseBody {
	s.Data = v
	return s
}

type ColorizeImageResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ColorizeImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ColorizeImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ColorizeImageResponseBodyData) SetImageURL(v string) *ColorizeImageResponseBodyData {
	s.ImageURL = &v
	return s
}

type ColorizeImageResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ColorizeImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ColorizeImageResponse) String() string {
	return tea.Prettify(s)
}

func (s ColorizeImageResponse) GoString() string {
	return s.String()
}

func (s *ColorizeImageResponse) SetHeaders(v map[string]*string) *ColorizeImageResponse {
	s.Headers = v
	return s
}

func (s *ColorizeImageResponse) SetBody(v *ColorizeImageResponseBody) *ColorizeImageResponse {
	s.Body = v
	return s
}

type RecolorImageRequest struct {
	Url           *string                             `json:"Url,omitempty" xml:"Url,omitempty"`
	Mode          *string                             `json:"Mode,omitempty" xml:"Mode,omitempty"`
	RefUrl        *string                             `json:"RefUrl,omitempty" xml:"RefUrl,omitempty"`
	ColorCount    *int32                              `json:"ColorCount,omitempty" xml:"ColorCount,omitempty"`
	ColorTemplate []*RecolorImageRequestColorTemplate `json:"ColorTemplate,omitempty" xml:"ColorTemplate,omitempty" type:"Repeated"`
}

func (s RecolorImageRequest) String() string {
	return tea.Prettify(s)
}

func (s RecolorImageRequest) GoString() string {
	return s.String()
}

func (s *RecolorImageRequest) SetUrl(v string) *RecolorImageRequest {
	s.Url = &v
	return s
}

func (s *RecolorImageRequest) SetMode(v string) *RecolorImageRequest {
	s.Mode = &v
	return s
}

func (s *RecolorImageRequest) SetRefUrl(v string) *RecolorImageRequest {
	s.RefUrl = &v
	return s
}

func (s *RecolorImageRequest) SetColorCount(v int32) *RecolorImageRequest {
	s.ColorCount = &v
	return s
}

func (s *RecolorImageRequest) SetColorTemplate(v []*RecolorImageRequestColorTemplate) *RecolorImageRequest {
	s.ColorTemplate = v
	return s
}

type RecolorImageRequestColorTemplate struct {
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
}

func (s RecolorImageRequestColorTemplate) String() string {
	return tea.Prettify(s)
}

func (s RecolorImageRequestColorTemplate) GoString() string {
	return s.String()
}

func (s *RecolorImageRequestColorTemplate) SetColor(v string) *RecolorImageRequestColorTemplate {
	s.Color = &v
	return s
}

type RecolorImageResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecolorImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecolorImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecolorImageResponseBody) GoString() string {
	return s.String()
}

func (s *RecolorImageResponseBody) SetRequestId(v string) *RecolorImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecolorImageResponseBody) SetData(v *RecolorImageResponseBodyData) *RecolorImageResponseBody {
	s.Data = v
	return s
}

type RecolorImageResponseBodyData struct {
	ImageList []*string `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
}

func (s RecolorImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecolorImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecolorImageResponseBodyData) SetImageList(v []*string) *RecolorImageResponseBodyData {
	s.ImageList = v
	return s
}

type RecolorImageResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecolorImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecolorImageResponse) String() string {
	return tea.Prettify(s)
}

func (s RecolorImageResponse) GoString() string {
	return s.String()
}

func (s *RecolorImageResponse) SetHeaders(v map[string]*string) *RecolorImageResponse {
	s.Headers = v
	return s
}

func (s *RecolorImageResponse) SetBody(v *RecolorImageResponseBody) *RecolorImageResponse {
	s.Body = v
	return s
}

type AssessCompositionRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s AssessCompositionRequest) String() string {
	return tea.Prettify(s)
}

func (s AssessCompositionRequest) GoString() string {
	return s.String()
}

func (s *AssessCompositionRequest) SetImageURL(v string) *AssessCompositionRequest {
	s.ImageURL = &v
	return s
}

type AssessCompositionAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s AssessCompositionAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AssessCompositionAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AssessCompositionAdvanceRequest) SetImageURLObject(v io.Reader) *AssessCompositionAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type AssessCompositionResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *AssessCompositionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s AssessCompositionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssessCompositionResponseBody) GoString() string {
	return s.String()
}

func (s *AssessCompositionResponseBody) SetRequestId(v string) *AssessCompositionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssessCompositionResponseBody) SetData(v *AssessCompositionResponseBodyData) *AssessCompositionResponseBody {
	s.Data = v
	return s
}

type AssessCompositionResponseBodyData struct {
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s AssessCompositionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AssessCompositionResponseBodyData) GoString() string {
	return s.String()
}

func (s *AssessCompositionResponseBodyData) SetScore(v float32) *AssessCompositionResponseBodyData {
	s.Score = &v
	return s
}

type AssessCompositionResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssessCompositionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssessCompositionResponse) String() string {
	return tea.Prettify(s)
}

func (s AssessCompositionResponse) GoString() string {
	return s.String()
}

func (s *AssessCompositionResponse) SetHeaders(v map[string]*string) *AssessCompositionResponse {
	s.Headers = v
	return s
}

func (s *AssessCompositionResponse) SetBody(v *AssessCompositionResponseBody) *AssessCompositionResponse {
	s.Body = v
	return s
}

type AssessSharpnessRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s AssessSharpnessRequest) String() string {
	return tea.Prettify(s)
}

func (s AssessSharpnessRequest) GoString() string {
	return s.String()
}

func (s *AssessSharpnessRequest) SetImageURL(v string) *AssessSharpnessRequest {
	s.ImageURL = &v
	return s
}

type AssessSharpnessAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s AssessSharpnessAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AssessSharpnessAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AssessSharpnessAdvanceRequest) SetImageURLObject(v io.Reader) *AssessSharpnessAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type AssessSharpnessResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *AssessSharpnessResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s AssessSharpnessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssessSharpnessResponseBody) GoString() string {
	return s.String()
}

func (s *AssessSharpnessResponseBody) SetRequestId(v string) *AssessSharpnessResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssessSharpnessResponseBody) SetData(v *AssessSharpnessResponseBodyData) *AssessSharpnessResponseBody {
	s.Data = v
	return s
}

type AssessSharpnessResponseBodyData struct {
	Sharpness *float32 `json:"Sharpness,omitempty" xml:"Sharpness,omitempty"`
}

func (s AssessSharpnessResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AssessSharpnessResponseBodyData) GoString() string {
	return s.String()
}

func (s *AssessSharpnessResponseBodyData) SetSharpness(v float32) *AssessSharpnessResponseBodyData {
	s.Sharpness = &v
	return s
}

type AssessSharpnessResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssessSharpnessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssessSharpnessResponse) String() string {
	return tea.Prettify(s)
}

func (s AssessSharpnessResponse) GoString() string {
	return s.String()
}

func (s *AssessSharpnessResponse) SetHeaders(v map[string]*string) *AssessSharpnessResponse {
	s.Headers = v
	return s
}

func (s *AssessSharpnessResponse) SetBody(v *AssessSharpnessResponseBody) *AssessSharpnessResponse {
	s.Body = v
	return s
}

type ErasePersonRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	UserMask *string `json:"UserMask,omitempty" xml:"UserMask,omitempty"`
}

func (s ErasePersonRequest) String() string {
	return tea.Prettify(s)
}

func (s ErasePersonRequest) GoString() string {
	return s.String()
}

func (s *ErasePersonRequest) SetImageURL(v string) *ErasePersonRequest {
	s.ImageURL = &v
	return s
}

func (s *ErasePersonRequest) SetUserMask(v string) *ErasePersonRequest {
	s.UserMask = &v
	return s
}

type ErasePersonAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	UserMask       *string   `json:"UserMask,omitempty" xml:"UserMask,omitempty"`
}

func (s ErasePersonAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ErasePersonAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ErasePersonAdvanceRequest) SetImageURLObject(v io.Reader) *ErasePersonAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *ErasePersonAdvanceRequest) SetUserMask(v string) *ErasePersonAdvanceRequest {
	s.UserMask = &v
	return s
}

type ErasePersonResponseBody struct {
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ErasePersonResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ErasePersonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ErasePersonResponseBody) GoString() string {
	return s.String()
}

func (s *ErasePersonResponseBody) SetRequestId(v string) *ErasePersonResponseBody {
	s.RequestId = &v
	return s
}

func (s *ErasePersonResponseBody) SetData(v *ErasePersonResponseBodyData) *ErasePersonResponseBody {
	s.Data = v
	return s
}

type ErasePersonResponseBodyData struct {
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s ErasePersonResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ErasePersonResponseBodyData) GoString() string {
	return s.String()
}

func (s *ErasePersonResponseBodyData) SetImageUrl(v string) *ErasePersonResponseBodyData {
	s.ImageUrl = &v
	return s
}

type ErasePersonResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ErasePersonResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ErasePersonResponse) String() string {
	return tea.Prettify(s)
}

func (s ErasePersonResponse) GoString() string {
	return s.String()
}

func (s *ErasePersonResponse) SetHeaders(v map[string]*string) *ErasePersonResponse {
	s.Headers = v
	return s
}

func (s *ErasePersonResponse) SetBody(v *ErasePersonResponseBody) *ErasePersonResponse {
	s.Body = v
	return s
}

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

type ImitatePhotoStyleRequest struct {
	StyleUrl *string `json:"StyleUrl,omitempty" xml:"StyleUrl,omitempty"`
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ImitatePhotoStyleRequest) String() string {
	return tea.Prettify(s)
}

func (s ImitatePhotoStyleRequest) GoString() string {
	return s.String()
}

func (s *ImitatePhotoStyleRequest) SetStyleUrl(v string) *ImitatePhotoStyleRequest {
	s.StyleUrl = &v
	return s
}

func (s *ImitatePhotoStyleRequest) SetImageURL(v string) *ImitatePhotoStyleRequest {
	s.ImageURL = &v
	return s
}

type ImitatePhotoStyleAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	StyleUrl       *string   `json:"StyleUrl,omitempty" xml:"StyleUrl,omitempty"`
}

func (s ImitatePhotoStyleAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ImitatePhotoStyleAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ImitatePhotoStyleAdvanceRequest) SetImageURLObject(v io.Reader) *ImitatePhotoStyleAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *ImitatePhotoStyleAdvanceRequest) SetStyleUrl(v string) *ImitatePhotoStyleAdvanceRequest {
	s.StyleUrl = &v
	return s
}

type ImitatePhotoStyleResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ImitatePhotoStyleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ImitatePhotoStyleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImitatePhotoStyleResponseBody) GoString() string {
	return s.String()
}

func (s *ImitatePhotoStyleResponseBody) SetRequestId(v string) *ImitatePhotoStyleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImitatePhotoStyleResponseBody) SetData(v *ImitatePhotoStyleResponseBodyData) *ImitatePhotoStyleResponseBody {
	s.Data = v
	return s
}

type ImitatePhotoStyleResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ImitatePhotoStyleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ImitatePhotoStyleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImitatePhotoStyleResponseBodyData) SetImageURL(v string) *ImitatePhotoStyleResponseBodyData {
	s.ImageURL = &v
	return s
}

type ImitatePhotoStyleResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImitatePhotoStyleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImitatePhotoStyleResponse) String() string {
	return tea.Prettify(s)
}

func (s ImitatePhotoStyleResponse) GoString() string {
	return s.String()
}

func (s *ImitatePhotoStyleResponse) SetHeaders(v map[string]*string) *ImitatePhotoStyleResponse {
	s.Headers = v
	return s
}

func (s *ImitatePhotoStyleResponse) SetBody(v *ImitatePhotoStyleResponseBody) *ImitatePhotoStyleResponse {
	s.Body = v
	return s
}

type ChangeImageSizeRequest struct {
	Width  *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	Url    *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ChangeImageSizeRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeImageSizeRequest) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeRequest) SetWidth(v int32) *ChangeImageSizeRequest {
	s.Width = &v
	return s
}

func (s *ChangeImageSizeRequest) SetHeight(v int32) *ChangeImageSizeRequest {
	s.Height = &v
	return s
}

func (s *ChangeImageSizeRequest) SetUrl(v string) *ChangeImageSizeRequest {
	s.Url = &v
	return s
}

type ChangeImageSizeAdvanceRequest struct {
	UrlObject io.Reader `json:"UrlObject,omitempty" xml:"UrlObject,omitempty" require:"true"`
	Width     *int32    `json:"Width,omitempty" xml:"Width,omitempty"`
	Height    *int32    `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s ChangeImageSizeAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeImageSizeAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeAdvanceRequest) SetUrlObject(v io.Reader) *ChangeImageSizeAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *ChangeImageSizeAdvanceRequest) SetWidth(v int32) *ChangeImageSizeAdvanceRequest {
	s.Width = &v
	return s
}

func (s *ChangeImageSizeAdvanceRequest) SetHeight(v int32) *ChangeImageSizeAdvanceRequest {
	s.Height = &v
	return s
}

type ChangeImageSizeResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ChangeImageSizeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ChangeImageSizeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeImageSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeResponseBody) SetRequestId(v string) *ChangeImageSizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeImageSizeResponseBody) SetData(v *ChangeImageSizeResponseBodyData) *ChangeImageSizeResponseBody {
	s.Data = v
	return s
}

type ChangeImageSizeResponseBodyData struct {
	Url            *string                                        `json:"Url,omitempty" xml:"Url,omitempty"`
	RetainLocation *ChangeImageSizeResponseBodyDataRetainLocation `json:"RetainLocation,omitempty" xml:"RetainLocation,omitempty" type:"Struct"`
}

func (s ChangeImageSizeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ChangeImageSizeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeResponseBodyData) SetUrl(v string) *ChangeImageSizeResponseBodyData {
	s.Url = &v
	return s
}

func (s *ChangeImageSizeResponseBodyData) SetRetainLocation(v *ChangeImageSizeResponseBodyDataRetainLocation) *ChangeImageSizeResponseBodyData {
	s.RetainLocation = v
	return s
}

type ChangeImageSizeResponseBodyDataRetainLocation struct {
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y      *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
	X      *int32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s ChangeImageSizeResponseBodyDataRetainLocation) String() string {
	return tea.Prettify(s)
}

func (s ChangeImageSizeResponseBodyDataRetainLocation) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeResponseBodyDataRetainLocation) SetWidth(v int32) *ChangeImageSizeResponseBodyDataRetainLocation {
	s.Width = &v
	return s
}

func (s *ChangeImageSizeResponseBodyDataRetainLocation) SetHeight(v int32) *ChangeImageSizeResponseBodyDataRetainLocation {
	s.Height = &v
	return s
}

func (s *ChangeImageSizeResponseBodyDataRetainLocation) SetY(v int32) *ChangeImageSizeResponseBodyDataRetainLocation {
	s.Y = &v
	return s
}

func (s *ChangeImageSizeResponseBodyDataRetainLocation) SetX(v int32) *ChangeImageSizeResponseBodyDataRetainLocation {
	s.X = &v
	return s
}

type ChangeImageSizeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ChangeImageSizeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeImageSizeResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeImageSizeResponse) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeResponse) SetHeaders(v map[string]*string) *ChangeImageSizeResponse {
	s.Headers = v
	return s
}

func (s *ChangeImageSizeResponse) SetBody(v *ChangeImageSizeResponseBody) *ChangeImageSizeResponse {
	s.Body = v
	return s
}

type EnhanceImageColorRequest struct {
	ImageURL     *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	Mode         *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s EnhanceImageColorRequest) String() string {
	return tea.Prettify(s)
}

func (s EnhanceImageColorRequest) GoString() string {
	return s.String()
}

func (s *EnhanceImageColorRequest) SetImageURL(v string) *EnhanceImageColorRequest {
	s.ImageURL = &v
	return s
}

func (s *EnhanceImageColorRequest) SetOutputFormat(v string) *EnhanceImageColorRequest {
	s.OutputFormat = &v
	return s
}

func (s *EnhanceImageColorRequest) SetMode(v string) *EnhanceImageColorRequest {
	s.Mode = &v
	return s
}

type EnhanceImageColorAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	OutputFormat   *string   `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	Mode           *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s EnhanceImageColorAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s EnhanceImageColorAdvanceRequest) GoString() string {
	return s.String()
}

func (s *EnhanceImageColorAdvanceRequest) SetImageURLObject(v io.Reader) *EnhanceImageColorAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *EnhanceImageColorAdvanceRequest) SetOutputFormat(v string) *EnhanceImageColorAdvanceRequest {
	s.OutputFormat = &v
	return s
}

func (s *EnhanceImageColorAdvanceRequest) SetMode(v string) *EnhanceImageColorAdvanceRequest {
	s.Mode = &v
	return s
}

type EnhanceImageColorResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *EnhanceImageColorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s EnhanceImageColorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnhanceImageColorResponseBody) GoString() string {
	return s.String()
}

func (s *EnhanceImageColorResponseBody) SetRequestId(v string) *EnhanceImageColorResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnhanceImageColorResponseBody) SetData(v *EnhanceImageColorResponseBodyData) *EnhanceImageColorResponseBody {
	s.Data = v
	return s
}

type EnhanceImageColorResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s EnhanceImageColorResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnhanceImageColorResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnhanceImageColorResponseBodyData) SetImageURL(v string) *EnhanceImageColorResponseBodyData {
	s.ImageURL = &v
	return s
}

type EnhanceImageColorResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnhanceImageColorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnhanceImageColorResponse) String() string {
	return tea.Prettify(s)
}

func (s EnhanceImageColorResponse) GoString() string {
	return s.String()
}

func (s *EnhanceImageColorResponse) SetHeaders(v map[string]*string) *EnhanceImageColorResponse {
	s.Headers = v
	return s
}

func (s *EnhanceImageColorResponse) SetBody(v *EnhanceImageColorResponseBody) *EnhanceImageColorResponse {
	s.Body = v
	return s
}

type AssessExposureRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s AssessExposureRequest) String() string {
	return tea.Prettify(s)
}

func (s AssessExposureRequest) GoString() string {
	return s.String()
}

func (s *AssessExposureRequest) SetImageURL(v string) *AssessExposureRequest {
	s.ImageURL = &v
	return s
}

type AssessExposureAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s AssessExposureAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AssessExposureAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AssessExposureAdvanceRequest) SetImageURLObject(v io.Reader) *AssessExposureAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type AssessExposureResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *AssessExposureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s AssessExposureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssessExposureResponseBody) GoString() string {
	return s.String()
}

func (s *AssessExposureResponseBody) SetRequestId(v string) *AssessExposureResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssessExposureResponseBody) SetData(v *AssessExposureResponseBodyData) *AssessExposureResponseBody {
	s.Data = v
	return s
}

type AssessExposureResponseBodyData struct {
	Exposure *float32 `json:"Exposure,omitempty" xml:"Exposure,omitempty"`
}

func (s AssessExposureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AssessExposureResponseBodyData) GoString() string {
	return s.String()
}

func (s *AssessExposureResponseBodyData) SetExposure(v float32) *AssessExposureResponseBodyData {
	s.Exposure = &v
	return s
}

type AssessExposureResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssessExposureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssessExposureResponse) String() string {
	return tea.Prettify(s)
}

func (s AssessExposureResponse) GoString() string {
	return s.String()
}

func (s *AssessExposureResponse) SetHeaders(v map[string]*string) *AssessExposureResponse {
	s.Headers = v
	return s
}

func (s *AssessExposureResponse) SetBody(v *AssessExposureResponseBody) *AssessExposureResponse {
	s.Body = v
	return s
}

type MakeSuperResolutionImageRequest struct {
	Url           *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Mode          *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	UpscaleFactor *int64  `json:"UpscaleFactor,omitempty" xml:"UpscaleFactor,omitempty"`
}

func (s MakeSuperResolutionImageRequest) String() string {
	return tea.Prettify(s)
}

func (s MakeSuperResolutionImageRequest) GoString() string {
	return s.String()
}

func (s *MakeSuperResolutionImageRequest) SetUrl(v string) *MakeSuperResolutionImageRequest {
	s.Url = &v
	return s
}

func (s *MakeSuperResolutionImageRequest) SetMode(v string) *MakeSuperResolutionImageRequest {
	s.Mode = &v
	return s
}

func (s *MakeSuperResolutionImageRequest) SetUpscaleFactor(v int64) *MakeSuperResolutionImageRequest {
	s.UpscaleFactor = &v
	return s
}

type MakeSuperResolutionImageAdvanceRequest struct {
	UrlObject     io.Reader `json:"UrlObject,omitempty" xml:"UrlObject,omitempty" require:"true"`
	Mode          *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	UpscaleFactor *int64    `json:"UpscaleFactor,omitempty" xml:"UpscaleFactor,omitempty"`
}

func (s MakeSuperResolutionImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s MakeSuperResolutionImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *MakeSuperResolutionImageAdvanceRequest) SetUrlObject(v io.Reader) *MakeSuperResolutionImageAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *MakeSuperResolutionImageAdvanceRequest) SetMode(v string) *MakeSuperResolutionImageAdvanceRequest {
	s.Mode = &v
	return s
}

func (s *MakeSuperResolutionImageAdvanceRequest) SetUpscaleFactor(v int64) *MakeSuperResolutionImageAdvanceRequest {
	s.UpscaleFactor = &v
	return s
}

type MakeSuperResolutionImageResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *MakeSuperResolutionImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s MakeSuperResolutionImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MakeSuperResolutionImageResponseBody) GoString() string {
	return s.String()
}

func (s *MakeSuperResolutionImageResponseBody) SetRequestId(v string) *MakeSuperResolutionImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *MakeSuperResolutionImageResponseBody) SetData(v *MakeSuperResolutionImageResponseBodyData) *MakeSuperResolutionImageResponseBody {
	s.Data = v
	return s
}

type MakeSuperResolutionImageResponseBodyData struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s MakeSuperResolutionImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MakeSuperResolutionImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *MakeSuperResolutionImageResponseBodyData) SetUrl(v string) *MakeSuperResolutionImageResponseBodyData {
	s.Url = &v
	return s
}

type MakeSuperResolutionImageResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MakeSuperResolutionImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MakeSuperResolutionImageResponse) String() string {
	return tea.Prettify(s)
}

func (s MakeSuperResolutionImageResponse) GoString() string {
	return s.String()
}

func (s *MakeSuperResolutionImageResponse) SetHeaders(v map[string]*string) *MakeSuperResolutionImageResponse {
	s.Headers = v
	return s
}

func (s *MakeSuperResolutionImageResponse) SetBody(v *MakeSuperResolutionImageResponseBody) *MakeSuperResolutionImageResponse {
	s.Body = v
	return s
}

type IntelligentCompositionRequest struct {
	NumBoxes *int32  `json:"NumBoxes,omitempty" xml:"NumBoxes,omitempty"`
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s IntelligentCompositionRequest) String() string {
	return tea.Prettify(s)
}

func (s IntelligentCompositionRequest) GoString() string {
	return s.String()
}

func (s *IntelligentCompositionRequest) SetNumBoxes(v int32) *IntelligentCompositionRequest {
	s.NumBoxes = &v
	return s
}

func (s *IntelligentCompositionRequest) SetImageURL(v string) *IntelligentCompositionRequest {
	s.ImageURL = &v
	return s
}

type IntelligentCompositionAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	NumBoxes       *int32    `json:"NumBoxes,omitempty" xml:"NumBoxes,omitempty"`
}

func (s IntelligentCompositionAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s IntelligentCompositionAdvanceRequest) GoString() string {
	return s.String()
}

func (s *IntelligentCompositionAdvanceRequest) SetImageURLObject(v io.Reader) *IntelligentCompositionAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *IntelligentCompositionAdvanceRequest) SetNumBoxes(v int32) *IntelligentCompositionAdvanceRequest {
	s.NumBoxes = &v
	return s
}

type IntelligentCompositionResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *IntelligentCompositionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s IntelligentCompositionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IntelligentCompositionResponseBody) GoString() string {
	return s.String()
}

func (s *IntelligentCompositionResponseBody) SetRequestId(v string) *IntelligentCompositionResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntelligentCompositionResponseBody) SetData(v *IntelligentCompositionResponseBodyData) *IntelligentCompositionResponseBody {
	s.Data = v
	return s
}

type IntelligentCompositionResponseBodyData struct {
	Elements []*IntelligentCompositionResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s IntelligentCompositionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s IntelligentCompositionResponseBodyData) GoString() string {
	return s.String()
}

func (s *IntelligentCompositionResponseBodyData) SetElements(v []*IntelligentCompositionResponseBodyDataElements) *IntelligentCompositionResponseBodyData {
	s.Elements = v
	return s
}

type IntelligentCompositionResponseBodyDataElements struct {
	MinX  *int32   `json:"MinX,omitempty" xml:"MinX,omitempty"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	MaxY  *int32   `json:"MaxY,omitempty" xml:"MaxY,omitempty"`
	MaxX  *int32   `json:"MaxX,omitempty" xml:"MaxX,omitempty"`
	MinY  *int32   `json:"MinY,omitempty" xml:"MinY,omitempty"`
}

func (s IntelligentCompositionResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s IntelligentCompositionResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *IntelligentCompositionResponseBodyDataElements) SetMinX(v int32) *IntelligentCompositionResponseBodyDataElements {
	s.MinX = &v
	return s
}

func (s *IntelligentCompositionResponseBodyDataElements) SetScore(v float32) *IntelligentCompositionResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *IntelligentCompositionResponseBodyDataElements) SetMaxY(v int32) *IntelligentCompositionResponseBodyDataElements {
	s.MaxY = &v
	return s
}

func (s *IntelligentCompositionResponseBodyDataElements) SetMaxX(v int32) *IntelligentCompositionResponseBodyDataElements {
	s.MaxX = &v
	return s
}

func (s *IntelligentCompositionResponseBodyDataElements) SetMinY(v int32) *IntelligentCompositionResponseBodyDataElements {
	s.MinY = &v
	return s
}

type IntelligentCompositionResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *IntelligentCompositionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IntelligentCompositionResponse) String() string {
	return tea.Prettify(s)
}

func (s IntelligentCompositionResponse) GoString() string {
	return s.String()
}

func (s *IntelligentCompositionResponse) SetHeaders(v map[string]*string) *IntelligentCompositionResponse {
	s.Headers = v
	return s
}

func (s *IntelligentCompositionResponse) SetBody(v *IntelligentCompositionResponseBody) *IntelligentCompositionResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("imageenhan"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ExtendImageStyleWithOptions(request *ExtendImageStyleRequest, runtime *util.RuntimeOptions) (_result *ExtendImageStyleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExtendImageStyleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExtendImageStyle"), tea.String("2019-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExtendImageStyle(request *ExtendImageStyleRequest) (_result *ExtendImageStyleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExtendImageStyleResponse{}
	_body, _err := client.ExtendImageStyleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImageBlindCharacterWatermarkWithOptions(request *ImageBlindCharacterWatermarkRequest, runtime *util.RuntimeOptions) (_result *ImageBlindCharacterWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ImageBlindCharacterWatermarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ImageBlindCharacterWatermark"), tea.String("2019-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImageBlindCharacterWatermark(request *ImageBlindCharacterWatermarkRequest) (_result *ImageBlindCharacterWatermarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImageBlindCharacterWatermarkResponse{}
	_body, _err := client.ImageBlindCharacterWatermarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImageBlindCharacterWatermarkAdvance(request *ImageBlindCharacterWatermarkAdvanceRequest, runtime *util.RuntimeOptions) (_result *ImageBlindCharacterWatermarkResponse, _err error) {
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
		Product:  tea.String("imageenhan"),
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
	imageBlindCharacterWatermarkReq := &ImageBlindCharacterWatermarkRequest{}
	openapiutil.Convert(request, imageBlindCharacterWatermarkReq)
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
		Content:     request.OriginImageURLObject,
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
	imageBlindCharacterWatermarkReq.OriginImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	imageBlindCharacterWatermarkResp, _err := client.ImageBlindCharacterWatermarkWithOptions(imageBlindCharacterWatermarkReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = imageBlindCharacterWatermarkResp
	return _result, _err
}

func (client *Client) RemoveImageWatermarkWithOptions(request *RemoveImageWatermarkRequest, runtime *util.RuntimeOptions) (_result *RemoveImageWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveImageWatermarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveImageWatermark"), tea.String("2019-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveImageWatermark(request *RemoveImageWatermarkRequest) (_result *RemoveImageWatermarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveImageWatermarkResponse{}
	_body, _err := client.RemoveImageWatermarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveImageWatermarkAdvance(request *RemoveImageWatermarkAdvanceRequest, runtime *util.RuntimeOptions) (_result *RemoveImageWatermarkResponse, _err error) {
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
		Product:  tea.String("imageenhan"),
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
	removeImageWatermarkReq := &RemoveImageWatermarkRequest{}
	openapiutil.Convert(request, removeImageWatermarkReq)
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
		Content:     request.ImageURLObject,
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
	removeImageWatermarkReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	removeImageWatermarkResp, _err := client.RemoveImageWatermarkWithOptions(removeImageWatermarkReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = removeImageWatermarkResp
	return _result, _err
}

func (client *Client) GenerateDynamicImageWithOptions(request *GenerateDynamicImageRequest, runtime *util.RuntimeOptions) (_result *GenerateDynamicImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateDynamicImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateDynamicImage"), tea.String("2019-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateDynamicImage(request *GenerateDynamicImageRequest) (_result *GenerateDynamicImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateDynamicImageResponse{}
	_body, _err := client.GenerateDynamicImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateDynamicImageAdvance(request *GenerateDynamicImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *GenerateDynamicImageResponse, _err error) {
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
		Product:  tea.String("imageenhan"),
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
	generateDynamicImageReq := &GenerateDynamicImageRequest{}
	openapiutil.Convert(request, generateDynamicImageReq)
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
		Content:     request.UrlObject,
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
	generateDynamicImageReq.Url = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	generateDynamicImageResp, _err := client.GenerateDynamicImageWithOptions(generateDynamicImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = generateDynamicImageResp
	return _result, _err
}

func (client *Client) ImageBlindPicWatermarkWithOptions(request *ImageBlindPicWatermarkRequest, runtime *util.RuntimeOptions) (_result *ImageBlindPicWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ImageBlindPicWatermarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ImageBlindPicWatermark"), tea.String("2019-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImageBlindPicWatermark(request *ImageBlindPicWatermarkRequest) (_result *ImageBlindPicWatermarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImageBlindPicWatermarkResponse{}
	_body, _err := client.ImageBlindPicWatermarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImageBlindPicWatermarkAdvance(request *ImageBlindPicWatermarkAdvanceRequest, runtime *util.RuntimeOptions) (_result *ImageBlindPicWatermarkResponse, _err error) {
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
		Product:  tea.String("imageenhan"),
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
	imageBlindPicWatermarkReq := &ImageBlindPicWatermarkRequest{}
	openapiutil.Convert(request, imageBlindPicWatermarkReq)
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
		Content:     request.OriginImageURLObject,
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
	imageBlindPicWatermarkReq.OriginImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	imageBlindPicWatermarkResp, _err := client.ImageBlindPicWatermarkWithOptions(imageBlindPicWatermarkReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = imageBlindPicWatermarkResp
	return _result, _err
}

func (client *Client) RemoveImageSubtitlesWithOptions(request *RemoveImageSubtitlesRequest, runtime *util.RuntimeOptions) (_result *RemoveImageSubtitlesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveImageSubtitlesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveImageSubtitles"), tea.String("2019-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveImageSubtitles(request *RemoveImageSubtitlesRequest) (_result *RemoveImageSubtitlesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveImageSubtitlesResponse{}
	_body, _err := client.RemoveImageSubtitlesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveImageSubtitlesAdvance(request *RemoveImageSubtitlesAdvanceRequest, runtime *util.RuntimeOptions) (_result *RemoveImageSubtitlesResponse, _err error) {
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
		Product:  tea.String("imageenhan"),
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
	removeImageSubtitlesReq := &RemoveImageSubtitlesRequest{}
	openapiutil.Convert(request, removeImageSubtitlesReq)
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
		Content:     request.ImageURLObject,
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
	removeImageSubtitlesReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	removeImageSubtitlesResp, _err := client.RemoveImageSubtitlesWithOptions(removeImageSubtitlesReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = removeImageSubtitlesResp
	return _result, _err
}

func (client *Client) RecolorHDImageWithOptions(request *RecolorHDImageRequest, runtime *util.RuntimeOptions) (_result *RecolorHDImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecolorHDImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecolorHDImage"), tea.String("2019-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecolorHDImage(request *RecolorHDImageRequest) (_result *RecolorHDImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecolorHDImageResponse{}
	_body, _err := client.RecolorHDImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecolorHDImageAdvance(request *RecolorHDImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecolorHDImageResponse, _err error) {
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
		Product:  tea.String("imageenhan"),
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
	recolorHDImageReq := &RecolorHDImageRequest{}
	openapiutil.Convert(request, recolorHDImageReq)
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
		Content:     request.UrlObject,
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
	recolorHDImageReq.Url = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recolorHDImageResp, _err := client.RecolorHDImageWithOptions(recolorHDImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recolorHDImageResp
	return _result, _err
}

func (client *Client) ColorizeImageWithOptions(request *ColorizeImageRequest, runtime *util.RuntimeOptions) (_result *ColorizeImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ColorizeImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ColorizeImage"), tea.String("2019-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ColorizeImage(request *ColorizeImageRequest) (_result *ColorizeImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ColorizeImageResponse{}
	_body, _err := client.ColorizeImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ColorizeImageAdvance(request *ColorizeImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *ColorizeImageResponse, _err error) {
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
		Product:  tea.String("imageenhan"),
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
	colorizeImageReq := &ColorizeImageRequest{}
	openapiutil.Convert(request, colorizeImageReq)
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
		Content:     request.ImageURLObject,
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
	colorizeImageReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	colorizeImageResp, _err := client.ColorizeImageWithOptions(colorizeImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = colorizeImageResp
	return _result, _err
}

func (client *Client) RecolorImageWithOptions(request *RecolorImageRequest, runtime *util.RuntimeOptions) (_result *RecolorImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecolorImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecolorImage"), tea.String("2019-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecolorImage(request *RecolorImageRequest) (_result *RecolorImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecolorImageResponse{}
	_body, _err := client.RecolorImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssessCompositionWithOptions(request *AssessCompositionRequest, runtime *util.RuntimeOptions) (_result *AssessCompositionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssessCompositionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AssessComposition"), tea.String("2019-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssessComposition(request *AssessCompositionRequest) (_result *AssessCompositionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssessCompositionResponse{}
	_body, _err := client.AssessCompositionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssessCompositionAdvance(request *AssessCompositionAdvanceRequest, runtime *util.RuntimeOptions) (_result *AssessCompositionResponse, _err error) {
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
		Product:  tea.String("imageenhan"),
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
	assessCompositionReq := &AssessCompositionRequest{}
	openapiutil.Convert(request, assessCompositionReq)
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
		Content:     request.ImageURLObject,
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
	assessCompositionReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	assessCompositionResp, _err := client.AssessCompositionWithOptions(assessCompositionReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = assessCompositionResp
	return _result, _err
}

func (client *Client) AssessSharpnessWithOptions(request *AssessSharpnessRequest, runtime *util.RuntimeOptions) (_result *AssessSharpnessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssessSharpnessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AssessSharpness"), tea.String("2019-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssessSharpness(request *AssessSharpnessRequest) (_result *AssessSharpnessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssessSharpnessResponse{}
	_body, _err := client.AssessSharpnessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssessSharpnessAdvance(request *AssessSharpnessAdvanceRequest, runtime *util.RuntimeOptions) (_result *AssessSharpnessResponse, _err error) {
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
		Product:  tea.String("imageenhan"),
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
	assessSharpnessReq := &AssessSharpnessRequest{}
	openapiutil.Convert(request, assessSharpnessReq)
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
		Content:     request.ImageURLObject,
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
	assessSharpnessReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	assessSharpnessResp, _err := client.AssessSharpnessWithOptions(assessSharpnessReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = assessSharpnessResp
	return _result, _err
}

func (client *Client) ErasePersonWithOptions(request *ErasePersonRequest, runtime *util.RuntimeOptions) (_result *ErasePersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ErasePersonResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ErasePerson"), tea.String("2019-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ErasePerson(request *ErasePersonRequest) (_result *ErasePersonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ErasePersonResponse{}
	_body, _err := client.ErasePersonWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ErasePersonAdvance(request *ErasePersonAdvanceRequest, runtime *util.RuntimeOptions) (_result *ErasePersonResponse, _err error) {
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
		Product:  tea.String("imageenhan"),
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
	erasePersonReq := &ErasePersonRequest{}
	openapiutil.Convert(request, erasePersonReq)
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
		Content:     request.ImageURLObject,
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
	erasePersonReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	erasePersonResp, _err := client.ErasePersonWithOptions(erasePersonReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = erasePersonResp
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
	_body, _err := client.DoRPCRequest(tea.String("GetAsyncJobResult"), tea.String("2019-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ImitatePhotoStyleWithOptions(request *ImitatePhotoStyleRequest, runtime *util.RuntimeOptions) (_result *ImitatePhotoStyleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ImitatePhotoStyleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ImitatePhotoStyle"), tea.String("2019-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImitatePhotoStyle(request *ImitatePhotoStyleRequest) (_result *ImitatePhotoStyleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImitatePhotoStyleResponse{}
	_body, _err := client.ImitatePhotoStyleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImitatePhotoStyleAdvance(request *ImitatePhotoStyleAdvanceRequest, runtime *util.RuntimeOptions) (_result *ImitatePhotoStyleResponse, _err error) {
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
		Product:  tea.String("imageenhan"),
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
	imitatePhotoStyleReq := &ImitatePhotoStyleRequest{}
	openapiutil.Convert(request, imitatePhotoStyleReq)
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
		Content:     request.ImageURLObject,
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
	imitatePhotoStyleReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	imitatePhotoStyleResp, _err := client.ImitatePhotoStyleWithOptions(imitatePhotoStyleReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = imitatePhotoStyleResp
	return _result, _err
}

func (client *Client) ChangeImageSizeWithOptions(request *ChangeImageSizeRequest, runtime *util.RuntimeOptions) (_result *ChangeImageSizeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ChangeImageSizeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ChangeImageSize"), tea.String("2019-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeImageSize(request *ChangeImageSizeRequest) (_result *ChangeImageSizeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeImageSizeResponse{}
	_body, _err := client.ChangeImageSizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeImageSizeAdvance(request *ChangeImageSizeAdvanceRequest, runtime *util.RuntimeOptions) (_result *ChangeImageSizeResponse, _err error) {
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
		Product:  tea.String("imageenhan"),
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
	changeImageSizeReq := &ChangeImageSizeRequest{}
	openapiutil.Convert(request, changeImageSizeReq)
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
		Content:     request.UrlObject,
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
	changeImageSizeReq.Url = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	changeImageSizeResp, _err := client.ChangeImageSizeWithOptions(changeImageSizeReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = changeImageSizeResp
	return _result, _err
}

func (client *Client) EnhanceImageColorWithOptions(request *EnhanceImageColorRequest, runtime *util.RuntimeOptions) (_result *EnhanceImageColorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnhanceImageColorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnhanceImageColor"), tea.String("2019-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnhanceImageColor(request *EnhanceImageColorRequest) (_result *EnhanceImageColorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnhanceImageColorResponse{}
	_body, _err := client.EnhanceImageColorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnhanceImageColorAdvance(request *EnhanceImageColorAdvanceRequest, runtime *util.RuntimeOptions) (_result *EnhanceImageColorResponse, _err error) {
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
		Product:  tea.String("imageenhan"),
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
	enhanceImageColorReq := &EnhanceImageColorRequest{}
	openapiutil.Convert(request, enhanceImageColorReq)
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
		Content:     request.ImageURLObject,
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
	enhanceImageColorReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	enhanceImageColorResp, _err := client.EnhanceImageColorWithOptions(enhanceImageColorReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = enhanceImageColorResp
	return _result, _err
}

func (client *Client) AssessExposureWithOptions(request *AssessExposureRequest, runtime *util.RuntimeOptions) (_result *AssessExposureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssessExposureResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AssessExposure"), tea.String("2019-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssessExposure(request *AssessExposureRequest) (_result *AssessExposureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssessExposureResponse{}
	_body, _err := client.AssessExposureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssessExposureAdvance(request *AssessExposureAdvanceRequest, runtime *util.RuntimeOptions) (_result *AssessExposureResponse, _err error) {
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
		Product:  tea.String("imageenhan"),
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
	assessExposureReq := &AssessExposureRequest{}
	openapiutil.Convert(request, assessExposureReq)
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
		Content:     request.ImageURLObject,
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
	assessExposureReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	assessExposureResp, _err := client.AssessExposureWithOptions(assessExposureReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = assessExposureResp
	return _result, _err
}

func (client *Client) MakeSuperResolutionImageWithOptions(request *MakeSuperResolutionImageRequest, runtime *util.RuntimeOptions) (_result *MakeSuperResolutionImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MakeSuperResolutionImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MakeSuperResolutionImage"), tea.String("2019-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MakeSuperResolutionImage(request *MakeSuperResolutionImageRequest) (_result *MakeSuperResolutionImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MakeSuperResolutionImageResponse{}
	_body, _err := client.MakeSuperResolutionImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MakeSuperResolutionImageAdvance(request *MakeSuperResolutionImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *MakeSuperResolutionImageResponse, _err error) {
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
		Product:  tea.String("imageenhan"),
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
	makeSuperResolutionImageReq := &MakeSuperResolutionImageRequest{}
	openapiutil.Convert(request, makeSuperResolutionImageReq)
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
		Content:     request.UrlObject,
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
	makeSuperResolutionImageReq.Url = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	makeSuperResolutionImageResp, _err := client.MakeSuperResolutionImageWithOptions(makeSuperResolutionImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = makeSuperResolutionImageResp
	return _result, _err
}

func (client *Client) IntelligentCompositionWithOptions(request *IntelligentCompositionRequest, runtime *util.RuntimeOptions) (_result *IntelligentCompositionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &IntelligentCompositionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("IntelligentComposition"), tea.String("2019-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IntelligentComposition(request *IntelligentCompositionRequest) (_result *IntelligentCompositionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IntelligentCompositionResponse{}
	_body, _err := client.IntelligentCompositionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IntelligentCompositionAdvance(request *IntelligentCompositionAdvanceRequest, runtime *util.RuntimeOptions) (_result *IntelligentCompositionResponse, _err error) {
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
		Product:  tea.String("imageenhan"),
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
	intelligentCompositionReq := &IntelligentCompositionRequest{}
	openapiutil.Convert(request, intelligentCompositionReq)
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
		Content:     request.ImageURLObject,
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
	intelligentCompositionReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	intelligentCompositionResp, _err := client.IntelligentCompositionWithOptions(intelligentCompositionReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = intelligentCompositionResp
	return _result, _err
}
