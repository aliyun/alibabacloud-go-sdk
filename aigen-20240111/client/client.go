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

type GenerateCosplayImageRequest struct {
	FaceImageUrl     *string `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	Style            *int64  `json:"Style,omitempty" xml:"Style,omitempty"`
	TemplateImageUrl *string `json:"TemplateImageUrl,omitempty" xml:"TemplateImageUrl,omitempty"`
}

func (s GenerateCosplayImageRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateCosplayImageRequest) GoString() string {
	return s.String()
}

func (s *GenerateCosplayImageRequest) SetFaceImageUrl(v string) *GenerateCosplayImageRequest {
	s.FaceImageUrl = &v
	return s
}

func (s *GenerateCosplayImageRequest) SetStyle(v int64) *GenerateCosplayImageRequest {
	s.Style = &v
	return s
}

func (s *GenerateCosplayImageRequest) SetTemplateImageUrl(v string) *GenerateCosplayImageRequest {
	s.TemplateImageUrl = &v
	return s
}

type GenerateCosplayImageAdvanceRequest struct {
	FaceImageUrlObject     io.Reader `json:"FaceImageUrl,omitempty" xml:"FaceImageUrl,omitempty"`
	Style                  *int64    `json:"Style,omitempty" xml:"Style,omitempty"`
	TemplateImageUrlObject io.Reader `json:"TemplateImageUrl,omitempty" xml:"TemplateImageUrl,omitempty"`
}

func (s GenerateCosplayImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateCosplayImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateCosplayImageAdvanceRequest) SetFaceImageUrlObject(v io.Reader) *GenerateCosplayImageAdvanceRequest {
	s.FaceImageUrlObject = v
	return s
}

func (s *GenerateCosplayImageAdvanceRequest) SetStyle(v int64) *GenerateCosplayImageAdvanceRequest {
	s.Style = &v
	return s
}

func (s *GenerateCosplayImageAdvanceRequest) SetTemplateImageUrlObject(v io.Reader) *GenerateCosplayImageAdvanceRequest {
	s.TemplateImageUrlObject = v
	return s
}

type GenerateCosplayImageResponseBody struct {
	Data      *GenerateCosplayImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateCosplayImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateCosplayImageResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCosplayImageResponseBody) SetData(v *GenerateCosplayImageResponseBodyData) *GenerateCosplayImageResponseBody {
	s.Data = v
	return s
}

func (s *GenerateCosplayImageResponseBody) SetMessage(v string) *GenerateCosplayImageResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateCosplayImageResponseBody) SetRequestId(v string) *GenerateCosplayImageResponseBody {
	s.RequestId = &v
	return s
}

type GenerateCosplayImageResponseBodyData struct {
	ResultUrl *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
}

func (s GenerateCosplayImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateCosplayImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateCosplayImageResponseBodyData) SetResultUrl(v string) *GenerateCosplayImageResponseBodyData {
	s.ResultUrl = &v
	return s
}

type GenerateCosplayImageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateCosplayImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateCosplayImageResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateCosplayImageResponse) GoString() string {
	return s.String()
}

func (s *GenerateCosplayImageResponse) SetHeaders(v map[string]*string) *GenerateCosplayImageResponse {
	s.Headers = v
	return s
}

func (s *GenerateCosplayImageResponse) SetStatusCode(v int32) *GenerateCosplayImageResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateCosplayImageResponse) SetBody(v *GenerateCosplayImageResponseBody) *GenerateCosplayImageResponse {
	s.Body = v
	return s
}

type GenerateTextDeformationRequest struct {
	// 1
	Async            *bool   `json:"Async,omitempty" xml:"Async,omitempty"`
	FontName         *string `json:"FontName,omitempty" xml:"FontName,omitempty"`
	N                *int64  `json:"N,omitempty" xml:"N,omitempty"`
	OutputImageRatio *string `json:"OutputImageRatio,omitempty" xml:"OutputImageRatio,omitempty"`
	Prompt           *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	TextContent      *string `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	TtfUrl           *string `json:"TtfUrl,omitempty" xml:"TtfUrl,omitempty"`
}

func (s GenerateTextDeformationRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateTextDeformationRequest) GoString() string {
	return s.String()
}

func (s *GenerateTextDeformationRequest) SetAsync(v bool) *GenerateTextDeformationRequest {
	s.Async = &v
	return s
}

func (s *GenerateTextDeformationRequest) SetFontName(v string) *GenerateTextDeformationRequest {
	s.FontName = &v
	return s
}

func (s *GenerateTextDeformationRequest) SetN(v int64) *GenerateTextDeformationRequest {
	s.N = &v
	return s
}

func (s *GenerateTextDeformationRequest) SetOutputImageRatio(v string) *GenerateTextDeformationRequest {
	s.OutputImageRatio = &v
	return s
}

func (s *GenerateTextDeformationRequest) SetPrompt(v string) *GenerateTextDeformationRequest {
	s.Prompt = &v
	return s
}

func (s *GenerateTextDeformationRequest) SetTextContent(v string) *GenerateTextDeformationRequest {
	s.TextContent = &v
	return s
}

func (s *GenerateTextDeformationRequest) SetTtfUrl(v string) *GenerateTextDeformationRequest {
	s.TtfUrl = &v
	return s
}

type GenerateTextDeformationAdvanceRequest struct {
	// 1
	Async            *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
	FontName         *string   `json:"FontName,omitempty" xml:"FontName,omitempty"`
	N                *int64    `json:"N,omitempty" xml:"N,omitempty"`
	OutputImageRatio *string   `json:"OutputImageRatio,omitempty" xml:"OutputImageRatio,omitempty"`
	Prompt           *string   `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	TextContent      *string   `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	TtfUrlObject     io.Reader `json:"TtfUrl,omitempty" xml:"TtfUrl,omitempty"`
}

func (s GenerateTextDeformationAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateTextDeformationAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateTextDeformationAdvanceRequest) SetAsync(v bool) *GenerateTextDeformationAdvanceRequest {
	s.Async = &v
	return s
}

func (s *GenerateTextDeformationAdvanceRequest) SetFontName(v string) *GenerateTextDeformationAdvanceRequest {
	s.FontName = &v
	return s
}

func (s *GenerateTextDeformationAdvanceRequest) SetN(v int64) *GenerateTextDeformationAdvanceRequest {
	s.N = &v
	return s
}

func (s *GenerateTextDeformationAdvanceRequest) SetOutputImageRatio(v string) *GenerateTextDeformationAdvanceRequest {
	s.OutputImageRatio = &v
	return s
}

func (s *GenerateTextDeformationAdvanceRequest) SetPrompt(v string) *GenerateTextDeformationAdvanceRequest {
	s.Prompt = &v
	return s
}

func (s *GenerateTextDeformationAdvanceRequest) SetTextContent(v string) *GenerateTextDeformationAdvanceRequest {
	s.TextContent = &v
	return s
}

func (s *GenerateTextDeformationAdvanceRequest) SetTtfUrlObject(v io.Reader) *GenerateTextDeformationAdvanceRequest {
	s.TtfUrlObject = v
	return s
}

type GenerateTextDeformationResponseBody struct {
	Data      *GenerateTextDeformationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateTextDeformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateTextDeformationResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateTextDeformationResponseBody) SetData(v *GenerateTextDeformationResponseBodyData) *GenerateTextDeformationResponseBody {
	s.Data = v
	return s
}

func (s *GenerateTextDeformationResponseBody) SetMessage(v string) *GenerateTextDeformationResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateTextDeformationResponseBody) SetRequestId(v string) *GenerateTextDeformationResponseBody {
	s.RequestId = &v
	return s
}

type GenerateTextDeformationResponseBodyData struct {
	ResultUrl *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
}

func (s GenerateTextDeformationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateTextDeformationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateTextDeformationResponseBodyData) SetResultUrl(v string) *GenerateTextDeformationResponseBodyData {
	s.ResultUrl = &v
	return s
}

type GenerateTextDeformationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateTextDeformationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateTextDeformationResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateTextDeformationResponse) GoString() string {
	return s.String()
}

func (s *GenerateTextDeformationResponse) SetHeaders(v map[string]*string) *GenerateTextDeformationResponse {
	s.Headers = v
	return s
}

func (s *GenerateTextDeformationResponse) SetStatusCode(v int32) *GenerateTextDeformationResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateTextDeformationResponse) SetBody(v *GenerateTextDeformationResponseBody) *GenerateTextDeformationResponse {
	s.Body = v
	return s
}

type GenerateTextTextureRequest struct {
	AlphaChannel     *bool   `json:"AlphaChannel,omitempty" xml:"AlphaChannel,omitempty"`
	FontName         *string `json:"FontName,omitempty" xml:"FontName,omitempty"`
	ImageShortSize   *int64  `json:"ImageShortSize,omitempty" xml:"ImageShortSize,omitempty"`
	ImageUrl         *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	N                *int64  `json:"N,omitempty" xml:"N,omitempty"`
	OutputImageRatio *string `json:"OutputImageRatio,omitempty" xml:"OutputImageRatio,omitempty"`
	Prompt           *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	TextContent      *string `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	TextureStyle     *string `json:"TextureStyle,omitempty" xml:"TextureStyle,omitempty"`
	TtfUrl           *string `json:"TtfUrl,omitempty" xml:"TtfUrl,omitempty"`
}

func (s GenerateTextTextureRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateTextTextureRequest) GoString() string {
	return s.String()
}

func (s *GenerateTextTextureRequest) SetAlphaChannel(v bool) *GenerateTextTextureRequest {
	s.AlphaChannel = &v
	return s
}

func (s *GenerateTextTextureRequest) SetFontName(v string) *GenerateTextTextureRequest {
	s.FontName = &v
	return s
}

func (s *GenerateTextTextureRequest) SetImageShortSize(v int64) *GenerateTextTextureRequest {
	s.ImageShortSize = &v
	return s
}

func (s *GenerateTextTextureRequest) SetImageUrl(v string) *GenerateTextTextureRequest {
	s.ImageUrl = &v
	return s
}

func (s *GenerateTextTextureRequest) SetN(v int64) *GenerateTextTextureRequest {
	s.N = &v
	return s
}

func (s *GenerateTextTextureRequest) SetOutputImageRatio(v string) *GenerateTextTextureRequest {
	s.OutputImageRatio = &v
	return s
}

func (s *GenerateTextTextureRequest) SetPrompt(v string) *GenerateTextTextureRequest {
	s.Prompt = &v
	return s
}

func (s *GenerateTextTextureRequest) SetTextContent(v string) *GenerateTextTextureRequest {
	s.TextContent = &v
	return s
}

func (s *GenerateTextTextureRequest) SetTextureStyle(v string) *GenerateTextTextureRequest {
	s.TextureStyle = &v
	return s
}

func (s *GenerateTextTextureRequest) SetTtfUrl(v string) *GenerateTextTextureRequest {
	s.TtfUrl = &v
	return s
}

type GenerateTextTextureAdvanceRequest struct {
	AlphaChannel     *bool     `json:"AlphaChannel,omitempty" xml:"AlphaChannel,omitempty"`
	FontName         *string   `json:"FontName,omitempty" xml:"FontName,omitempty"`
	ImageShortSize   *int64    `json:"ImageShortSize,omitempty" xml:"ImageShortSize,omitempty"`
	ImageUrlObject   io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	N                *int64    `json:"N,omitempty" xml:"N,omitempty"`
	OutputImageRatio *string   `json:"OutputImageRatio,omitempty" xml:"OutputImageRatio,omitempty"`
	Prompt           *string   `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	TextContent      *string   `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	TextureStyle     *string   `json:"TextureStyle,omitempty" xml:"TextureStyle,omitempty"`
	TtfUrlObject     io.Reader `json:"TtfUrl,omitempty" xml:"TtfUrl,omitempty"`
}

func (s GenerateTextTextureAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateTextTextureAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateTextTextureAdvanceRequest) SetAlphaChannel(v bool) *GenerateTextTextureAdvanceRequest {
	s.AlphaChannel = &v
	return s
}

func (s *GenerateTextTextureAdvanceRequest) SetFontName(v string) *GenerateTextTextureAdvanceRequest {
	s.FontName = &v
	return s
}

func (s *GenerateTextTextureAdvanceRequest) SetImageShortSize(v int64) *GenerateTextTextureAdvanceRequest {
	s.ImageShortSize = &v
	return s
}

func (s *GenerateTextTextureAdvanceRequest) SetImageUrlObject(v io.Reader) *GenerateTextTextureAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *GenerateTextTextureAdvanceRequest) SetN(v int64) *GenerateTextTextureAdvanceRequest {
	s.N = &v
	return s
}

func (s *GenerateTextTextureAdvanceRequest) SetOutputImageRatio(v string) *GenerateTextTextureAdvanceRequest {
	s.OutputImageRatio = &v
	return s
}

func (s *GenerateTextTextureAdvanceRequest) SetPrompt(v string) *GenerateTextTextureAdvanceRequest {
	s.Prompt = &v
	return s
}

func (s *GenerateTextTextureAdvanceRequest) SetTextContent(v string) *GenerateTextTextureAdvanceRequest {
	s.TextContent = &v
	return s
}

func (s *GenerateTextTextureAdvanceRequest) SetTextureStyle(v string) *GenerateTextTextureAdvanceRequest {
	s.TextureStyle = &v
	return s
}

func (s *GenerateTextTextureAdvanceRequest) SetTtfUrlObject(v io.Reader) *GenerateTextTextureAdvanceRequest {
	s.TtfUrlObject = v
	return s
}

type GenerateTextTextureResponseBody struct {
	Data      *GenerateTextTextureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateTextTextureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateTextTextureResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateTextTextureResponseBody) SetData(v *GenerateTextTextureResponseBodyData) *GenerateTextTextureResponseBody {
	s.Data = v
	return s
}

func (s *GenerateTextTextureResponseBody) SetMessage(v string) *GenerateTextTextureResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateTextTextureResponseBody) SetRequestId(v string) *GenerateTextTextureResponseBody {
	s.RequestId = &v
	return s
}

type GenerateTextTextureResponseBodyData struct {
	ResultUrl *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
}

func (s GenerateTextTextureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateTextTextureResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateTextTextureResponseBodyData) SetResultUrl(v string) *GenerateTextTextureResponseBodyData {
	s.ResultUrl = &v
	return s
}

type GenerateTextTextureResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateTextTextureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateTextTextureResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateTextTextureResponse) GoString() string {
	return s.String()
}

func (s *GenerateTextTextureResponse) SetHeaders(v map[string]*string) *GenerateTextTextureResponse {
	s.Headers = v
	return s
}

func (s *GenerateTextTextureResponse) SetStatusCode(v int32) *GenerateTextTextureResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateTextTextureResponse) SetBody(v *GenerateTextTextureResponseBody) *GenerateTextTextureResponse {
	s.Body = v
	return s
}

type InteractiveFullSegmentationRequest struct {
	ImageUrl     *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	ReturnFormat *string `json:"ReturnFormat,omitempty" xml:"ReturnFormat,omitempty"`
}

func (s InteractiveFullSegmentationRequest) String() string {
	return tea.Prettify(s)
}

func (s InteractiveFullSegmentationRequest) GoString() string {
	return s.String()
}

func (s *InteractiveFullSegmentationRequest) SetImageUrl(v string) *InteractiveFullSegmentationRequest {
	s.ImageUrl = &v
	return s
}

func (s *InteractiveFullSegmentationRequest) SetReturnFormat(v string) *InteractiveFullSegmentationRequest {
	s.ReturnFormat = &v
	return s
}

type InteractiveFullSegmentationAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	ReturnFormat   *string   `json:"ReturnFormat,omitempty" xml:"ReturnFormat,omitempty"`
}

func (s InteractiveFullSegmentationAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s InteractiveFullSegmentationAdvanceRequest) GoString() string {
	return s.String()
}

func (s *InteractiveFullSegmentationAdvanceRequest) SetImageUrlObject(v io.Reader) *InteractiveFullSegmentationAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *InteractiveFullSegmentationAdvanceRequest) SetReturnFormat(v string) *InteractiveFullSegmentationAdvanceRequest {
	s.ReturnFormat = &v
	return s
}

type InteractiveFullSegmentationResponseBody struct {
	Data      *InteractiveFullSegmentationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InteractiveFullSegmentationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InteractiveFullSegmentationResponseBody) GoString() string {
	return s.String()
}

func (s *InteractiveFullSegmentationResponseBody) SetData(v *InteractiveFullSegmentationResponseBodyData) *InteractiveFullSegmentationResponseBody {
	s.Data = v
	return s
}

func (s *InteractiveFullSegmentationResponseBody) SetMessage(v string) *InteractiveFullSegmentationResponseBody {
	s.Message = &v
	return s
}

func (s *InteractiveFullSegmentationResponseBody) SetRequestId(v string) *InteractiveFullSegmentationResponseBody {
	s.RequestId = &v
	return s
}

type InteractiveFullSegmentationResponseBodyData struct {
	ResultUrl *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
}

func (s InteractiveFullSegmentationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InteractiveFullSegmentationResponseBodyData) GoString() string {
	return s.String()
}

func (s *InteractiveFullSegmentationResponseBodyData) SetResultUrl(v string) *InteractiveFullSegmentationResponseBodyData {
	s.ResultUrl = &v
	return s
}

type InteractiveFullSegmentationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InteractiveFullSegmentationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InteractiveFullSegmentationResponse) String() string {
	return tea.Prettify(s)
}

func (s InteractiveFullSegmentationResponse) GoString() string {
	return s.String()
}

func (s *InteractiveFullSegmentationResponse) SetHeaders(v map[string]*string) *InteractiveFullSegmentationResponse {
	s.Headers = v
	return s
}

func (s *InteractiveFullSegmentationResponse) SetStatusCode(v int32) *InteractiveFullSegmentationResponse {
	s.StatusCode = &v
	return s
}

func (s *InteractiveFullSegmentationResponse) SetBody(v *InteractiveFullSegmentationResponseBody) *InteractiveFullSegmentationResponse {
	s.Body = v
	return s
}

type InteractiveScribbleSegmentationRequest struct {
	EdgeFeathering      *string `json:"EdgeFeathering,omitempty" xml:"EdgeFeathering,omitempty"`
	ImageUrl            *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	IntegratedMaskUrl   *string `json:"IntegratedMaskUrl,omitempty" xml:"IntegratedMaskUrl,omitempty"`
	MaskImageUrl        *string `json:"MaskImageUrl,omitempty" xml:"MaskImageUrl,omitempty"`
	NegScribbleImageUrl *string `json:"NegScribbleImageUrl,omitempty" xml:"NegScribbleImageUrl,omitempty"`
	PosScribbleImageUrl *string `json:"PosScribbleImageUrl,omitempty" xml:"PosScribbleImageUrl,omitempty"`
	PostprocessOption   *string `json:"PostprocessOption,omitempty" xml:"PostprocessOption,omitempty"`
	ReturnForm          *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
	ReturnFormat        *string `json:"ReturnFormat,omitempty" xml:"ReturnFormat,omitempty"`
}

func (s InteractiveScribbleSegmentationRequest) String() string {
	return tea.Prettify(s)
}

func (s InteractiveScribbleSegmentationRequest) GoString() string {
	return s.String()
}

func (s *InteractiveScribbleSegmentationRequest) SetEdgeFeathering(v string) *InteractiveScribbleSegmentationRequest {
	s.EdgeFeathering = &v
	return s
}

func (s *InteractiveScribbleSegmentationRequest) SetImageUrl(v string) *InteractiveScribbleSegmentationRequest {
	s.ImageUrl = &v
	return s
}

func (s *InteractiveScribbleSegmentationRequest) SetIntegratedMaskUrl(v string) *InteractiveScribbleSegmentationRequest {
	s.IntegratedMaskUrl = &v
	return s
}

func (s *InteractiveScribbleSegmentationRequest) SetMaskImageUrl(v string) *InteractiveScribbleSegmentationRequest {
	s.MaskImageUrl = &v
	return s
}

func (s *InteractiveScribbleSegmentationRequest) SetNegScribbleImageUrl(v string) *InteractiveScribbleSegmentationRequest {
	s.NegScribbleImageUrl = &v
	return s
}

func (s *InteractiveScribbleSegmentationRequest) SetPosScribbleImageUrl(v string) *InteractiveScribbleSegmentationRequest {
	s.PosScribbleImageUrl = &v
	return s
}

func (s *InteractiveScribbleSegmentationRequest) SetPostprocessOption(v string) *InteractiveScribbleSegmentationRequest {
	s.PostprocessOption = &v
	return s
}

func (s *InteractiveScribbleSegmentationRequest) SetReturnForm(v string) *InteractiveScribbleSegmentationRequest {
	s.ReturnForm = &v
	return s
}

func (s *InteractiveScribbleSegmentationRequest) SetReturnFormat(v string) *InteractiveScribbleSegmentationRequest {
	s.ReturnFormat = &v
	return s
}

type InteractiveScribbleSegmentationAdvanceRequest struct {
	EdgeFeathering            *string   `json:"EdgeFeathering,omitempty" xml:"EdgeFeathering,omitempty"`
	ImageUrlObject            io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	IntegratedMaskUrlObject   io.Reader `json:"IntegratedMaskUrl,omitempty" xml:"IntegratedMaskUrl,omitempty"`
	MaskImageUrlObject        io.Reader `json:"MaskImageUrl,omitempty" xml:"MaskImageUrl,omitempty"`
	NegScribbleImageUrlObject io.Reader `json:"NegScribbleImageUrl,omitempty" xml:"NegScribbleImageUrl,omitempty"`
	PosScribbleImageUrlObject io.Reader `json:"PosScribbleImageUrl,omitempty" xml:"PosScribbleImageUrl,omitempty"`
	PostprocessOption         *string   `json:"PostprocessOption,omitempty" xml:"PostprocessOption,omitempty"`
	ReturnForm                *string   `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
	ReturnFormat              *string   `json:"ReturnFormat,omitempty" xml:"ReturnFormat,omitempty"`
}

func (s InteractiveScribbleSegmentationAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s InteractiveScribbleSegmentationAdvanceRequest) GoString() string {
	return s.String()
}

func (s *InteractiveScribbleSegmentationAdvanceRequest) SetEdgeFeathering(v string) *InteractiveScribbleSegmentationAdvanceRequest {
	s.EdgeFeathering = &v
	return s
}

func (s *InteractiveScribbleSegmentationAdvanceRequest) SetImageUrlObject(v io.Reader) *InteractiveScribbleSegmentationAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *InteractiveScribbleSegmentationAdvanceRequest) SetIntegratedMaskUrlObject(v io.Reader) *InteractiveScribbleSegmentationAdvanceRequest {
	s.IntegratedMaskUrlObject = v
	return s
}

func (s *InteractiveScribbleSegmentationAdvanceRequest) SetMaskImageUrlObject(v io.Reader) *InteractiveScribbleSegmentationAdvanceRequest {
	s.MaskImageUrlObject = v
	return s
}

func (s *InteractiveScribbleSegmentationAdvanceRequest) SetNegScribbleImageUrlObject(v io.Reader) *InteractiveScribbleSegmentationAdvanceRequest {
	s.NegScribbleImageUrlObject = v
	return s
}

func (s *InteractiveScribbleSegmentationAdvanceRequest) SetPosScribbleImageUrlObject(v io.Reader) *InteractiveScribbleSegmentationAdvanceRequest {
	s.PosScribbleImageUrlObject = v
	return s
}

func (s *InteractiveScribbleSegmentationAdvanceRequest) SetPostprocessOption(v string) *InteractiveScribbleSegmentationAdvanceRequest {
	s.PostprocessOption = &v
	return s
}

func (s *InteractiveScribbleSegmentationAdvanceRequest) SetReturnForm(v string) *InteractiveScribbleSegmentationAdvanceRequest {
	s.ReturnForm = &v
	return s
}

func (s *InteractiveScribbleSegmentationAdvanceRequest) SetReturnFormat(v string) *InteractiveScribbleSegmentationAdvanceRequest {
	s.ReturnFormat = &v
	return s
}

type InteractiveScribbleSegmentationResponseBody struct {
	Data      *InteractiveScribbleSegmentationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InteractiveScribbleSegmentationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InteractiveScribbleSegmentationResponseBody) GoString() string {
	return s.String()
}

func (s *InteractiveScribbleSegmentationResponseBody) SetData(v *InteractiveScribbleSegmentationResponseBodyData) *InteractiveScribbleSegmentationResponseBody {
	s.Data = v
	return s
}

func (s *InteractiveScribbleSegmentationResponseBody) SetRequestId(v string) *InteractiveScribbleSegmentationResponseBody {
	s.RequestId = &v
	return s
}

type InteractiveScribbleSegmentationResponseBodyData struct {
	ResultUrl *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
}

func (s InteractiveScribbleSegmentationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InteractiveScribbleSegmentationResponseBodyData) GoString() string {
	return s.String()
}

func (s *InteractiveScribbleSegmentationResponseBodyData) SetResultUrl(v string) *InteractiveScribbleSegmentationResponseBodyData {
	s.ResultUrl = &v
	return s
}

type InteractiveScribbleSegmentationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InteractiveScribbleSegmentationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InteractiveScribbleSegmentationResponse) String() string {
	return tea.Prettify(s)
}

func (s InteractiveScribbleSegmentationResponse) GoString() string {
	return s.String()
}

func (s *InteractiveScribbleSegmentationResponse) SetHeaders(v map[string]*string) *InteractiveScribbleSegmentationResponse {
	s.Headers = v
	return s
}

func (s *InteractiveScribbleSegmentationResponse) SetStatusCode(v int32) *InteractiveScribbleSegmentationResponse {
	s.StatusCode = &v
	return s
}

func (s *InteractiveScribbleSegmentationResponse) SetBody(v *InteractiveScribbleSegmentationResponseBody) *InteractiveScribbleSegmentationResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("aigen"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GenerateCosplayImageWithOptions(request *GenerateCosplayImageRequest, runtime *util.RuntimeOptions) (_result *GenerateCosplayImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FaceImageUrl)) {
		body["FaceImageUrl"] = request.FaceImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Style)) {
		body["Style"] = request.Style
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateImageUrl)) {
		body["TemplateImageUrl"] = request.TemplateImageUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateCosplayImage"),
		Version:     tea.String("2024-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateCosplayImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateCosplayImage(request *GenerateCosplayImageRequest) (_result *GenerateCosplayImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateCosplayImageResponse{}
	_body, _err := client.GenerateCosplayImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateCosplayImageAdvance(request *GenerateCosplayImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *GenerateCosplayImageResponse, _err error) {
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
		Product:  tea.String("aigen"),
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
	generateCosplayImageReq := &GenerateCosplayImageRequest{}
	openapiutil.Convert(request, generateCosplayImageReq)
	if !tea.BoolValue(util.IsUnset(request.FaceImageUrlObject)) {
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
			Content:     request.FaceImageUrlObject,
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
		generateCosplayImageReq.FaceImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateImageUrlObject)) {
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
			Content:     request.TemplateImageUrlObject,
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
		generateCosplayImageReq.TemplateImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	generateCosplayImageResp, _err := client.GenerateCosplayImageWithOptions(generateCosplayImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = generateCosplayImageResp
	return _result, _err
}

func (client *Client) GenerateTextDeformationWithOptions(request *GenerateTextDeformationRequest, runtime *util.RuntimeOptions) (_result *GenerateTextDeformationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Async)) {
		body["Async"] = request.Async
	}

	if !tea.BoolValue(util.IsUnset(request.FontName)) {
		body["FontName"] = request.FontName
	}

	if !tea.BoolValue(util.IsUnset(request.N)) {
		body["N"] = request.N
	}

	if !tea.BoolValue(util.IsUnset(request.OutputImageRatio)) {
		body["OutputImageRatio"] = request.OutputImageRatio
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["Prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.TextContent)) {
		body["TextContent"] = request.TextContent
	}

	if !tea.BoolValue(util.IsUnset(request.TtfUrl)) {
		body["TtfUrl"] = request.TtfUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateTextDeformation"),
		Version:     tea.String("2024-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateTextDeformationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateTextDeformation(request *GenerateTextDeformationRequest) (_result *GenerateTextDeformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateTextDeformationResponse{}
	_body, _err := client.GenerateTextDeformationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateTextDeformationAdvance(request *GenerateTextDeformationAdvanceRequest, runtime *util.RuntimeOptions) (_result *GenerateTextDeformationResponse, _err error) {
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
		Product:  tea.String("aigen"),
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
	generateTextDeformationReq := &GenerateTextDeformationRequest{}
	openapiutil.Convert(request, generateTextDeformationReq)
	if !tea.BoolValue(util.IsUnset(request.TtfUrlObject)) {
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
			Content:     request.TtfUrlObject,
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
		generateTextDeformationReq.TtfUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	generateTextDeformationResp, _err := client.GenerateTextDeformationWithOptions(generateTextDeformationReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = generateTextDeformationResp
	return _result, _err
}

func (client *Client) GenerateTextTextureWithOptions(request *GenerateTextTextureRequest, runtime *util.RuntimeOptions) (_result *GenerateTextTextureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TextureStyle)) {
		query["TextureStyle"] = request.TextureStyle
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlphaChannel)) {
		body["AlphaChannel"] = request.AlphaChannel
	}

	if !tea.BoolValue(util.IsUnset(request.FontName)) {
		body["FontName"] = request.FontName
	}

	if !tea.BoolValue(util.IsUnset(request.ImageShortSize)) {
		body["ImageShortSize"] = request.ImageShortSize
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.N)) {
		body["N"] = request.N
	}

	if !tea.BoolValue(util.IsUnset(request.OutputImageRatio)) {
		body["OutputImageRatio"] = request.OutputImageRatio
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["Prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.TextContent)) {
		body["TextContent"] = request.TextContent
	}

	if !tea.BoolValue(util.IsUnset(request.TtfUrl)) {
		body["TtfUrl"] = request.TtfUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateTextTexture"),
		Version:     tea.String("2024-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateTextTextureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateTextTexture(request *GenerateTextTextureRequest) (_result *GenerateTextTextureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateTextTextureResponse{}
	_body, _err := client.GenerateTextTextureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateTextTextureAdvance(request *GenerateTextTextureAdvanceRequest, runtime *util.RuntimeOptions) (_result *GenerateTextTextureResponse, _err error) {
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
		Product:  tea.String("aigen"),
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
	generateTextTextureReq := &GenerateTextTextureRequest{}
	openapiutil.Convert(request, generateTextTextureReq)
	if !tea.BoolValue(util.IsUnset(request.ImageUrlObject)) {
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
			Content:     request.ImageUrlObject,
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
		generateTextTextureReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	if !tea.BoolValue(util.IsUnset(request.TtfUrlObject)) {
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
			Content:     request.TtfUrlObject,
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
		generateTextTextureReq.TtfUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	generateTextTextureResp, _err := client.GenerateTextTextureWithOptions(generateTextTextureReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = generateTextTextureResp
	return _result, _err
}

func (client *Client) InteractiveFullSegmentationWithOptions(request *InteractiveFullSegmentationRequest, runtime *util.RuntimeOptions) (_result *InteractiveFullSegmentationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnFormat)) {
		body["ReturnFormat"] = request.ReturnFormat
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InteractiveFullSegmentation"),
		Version:     tea.String("2024-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InteractiveFullSegmentationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InteractiveFullSegmentation(request *InteractiveFullSegmentationRequest) (_result *InteractiveFullSegmentationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InteractiveFullSegmentationResponse{}
	_body, _err := client.InteractiveFullSegmentationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InteractiveFullSegmentationAdvance(request *InteractiveFullSegmentationAdvanceRequest, runtime *util.RuntimeOptions) (_result *InteractiveFullSegmentationResponse, _err error) {
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
		Product:  tea.String("aigen"),
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
	interactiveFullSegmentationReq := &InteractiveFullSegmentationRequest{}
	openapiutil.Convert(request, interactiveFullSegmentationReq)
	if !tea.BoolValue(util.IsUnset(request.ImageUrlObject)) {
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
			Content:     request.ImageUrlObject,
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
		interactiveFullSegmentationReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	interactiveFullSegmentationResp, _err := client.InteractiveFullSegmentationWithOptions(interactiveFullSegmentationReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = interactiveFullSegmentationResp
	return _result, _err
}

func (client *Client) InteractiveScribbleSegmentationWithOptions(request *InteractiveScribbleSegmentationRequest, runtime *util.RuntimeOptions) (_result *InteractiveScribbleSegmentationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EdgeFeathering)) {
		body["EdgeFeathering"] = request.EdgeFeathering
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IntegratedMaskUrl)) {
		body["IntegratedMaskUrl"] = request.IntegratedMaskUrl
	}

	if !tea.BoolValue(util.IsUnset(request.MaskImageUrl)) {
		body["MaskImageUrl"] = request.MaskImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.NegScribbleImageUrl)) {
		body["NegScribbleImageUrl"] = request.NegScribbleImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.PosScribbleImageUrl)) {
		body["PosScribbleImageUrl"] = request.PosScribbleImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.PostprocessOption)) {
		body["PostprocessOption"] = request.PostprocessOption
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnForm)) {
		body["ReturnForm"] = request.ReturnForm
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnFormat)) {
		body["ReturnFormat"] = request.ReturnFormat
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InteractiveScribbleSegmentation"),
		Version:     tea.String("2024-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InteractiveScribbleSegmentationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InteractiveScribbleSegmentation(request *InteractiveScribbleSegmentationRequest) (_result *InteractiveScribbleSegmentationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InteractiveScribbleSegmentationResponse{}
	_body, _err := client.InteractiveScribbleSegmentationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InteractiveScribbleSegmentationAdvance(request *InteractiveScribbleSegmentationAdvanceRequest, runtime *util.RuntimeOptions) (_result *InteractiveScribbleSegmentationResponse, _err error) {
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
		Product:  tea.String("aigen"),
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
	interactiveScribbleSegmentationReq := &InteractiveScribbleSegmentationRequest{}
	openapiutil.Convert(request, interactiveScribbleSegmentationReq)
	if !tea.BoolValue(util.IsUnset(request.ImageUrlObject)) {
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
			Content:     request.ImageUrlObject,
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
		interactiveScribbleSegmentationReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	if !tea.BoolValue(util.IsUnset(request.IntegratedMaskUrlObject)) {
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
			Content:     request.IntegratedMaskUrlObject,
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
		interactiveScribbleSegmentationReq.IntegratedMaskUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	if !tea.BoolValue(util.IsUnset(request.MaskImageUrlObject)) {
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
			Content:     request.MaskImageUrlObject,
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
		interactiveScribbleSegmentationReq.MaskImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	if !tea.BoolValue(util.IsUnset(request.NegScribbleImageUrlObject)) {
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
			Content:     request.NegScribbleImageUrlObject,
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
		interactiveScribbleSegmentationReq.NegScribbleImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	if !tea.BoolValue(util.IsUnset(request.PosScribbleImageUrlObject)) {
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
			Content:     request.PosScribbleImageUrlObject,
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
		interactiveScribbleSegmentationReq.PosScribbleImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	interactiveScribbleSegmentationResp, _err := client.InteractiveScribbleSegmentationWithOptions(interactiveScribbleSegmentationReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = interactiveScribbleSegmentationResp
	return _result, _err
}
