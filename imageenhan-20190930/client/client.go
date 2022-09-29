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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	Data      *AssessCompositionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssessCompositionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssessCompositionResponseBody) GoString() string {
	return s.String()
}

func (s *AssessCompositionResponseBody) SetData(v *AssessCompositionResponseBodyData) *AssessCompositionResponseBody {
	s.Data = v
	return s
}

func (s *AssessCompositionResponseBody) SetRequestId(v string) *AssessCompositionResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AssessCompositionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AssessCompositionResponse) SetStatusCode(v int32) *AssessCompositionResponse {
	s.StatusCode = &v
	return s
}

func (s *AssessCompositionResponse) SetBody(v *AssessCompositionResponseBody) *AssessCompositionResponse {
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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	Data      *AssessExposureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssessExposureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssessExposureResponseBody) GoString() string {
	return s.String()
}

func (s *AssessExposureResponseBody) SetData(v *AssessExposureResponseBodyData) *AssessExposureResponseBody {
	s.Data = v
	return s
}

func (s *AssessExposureResponseBody) SetRequestId(v string) *AssessExposureResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AssessExposureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AssessExposureResponse) SetStatusCode(v int32) *AssessExposureResponse {
	s.StatusCode = &v
	return s
}

func (s *AssessExposureResponse) SetBody(v *AssessExposureResponseBody) *AssessExposureResponse {
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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	Data      *AssessSharpnessResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssessSharpnessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssessSharpnessResponseBody) GoString() string {
	return s.String()
}

func (s *AssessSharpnessResponseBody) SetData(v *AssessSharpnessResponseBodyData) *AssessSharpnessResponseBody {
	s.Data = v
	return s
}

func (s *AssessSharpnessResponseBody) SetRequestId(v string) *AssessSharpnessResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AssessSharpnessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AssessSharpnessResponse) SetStatusCode(v int32) *AssessSharpnessResponse {
	s.StatusCode = &v
	return s
}

func (s *AssessSharpnessResponse) SetBody(v *AssessSharpnessResponseBody) *AssessSharpnessResponse {
	s.Body = v
	return s
}

type ChangeImageSizeRequest struct {
	Height *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	Url    *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Width  *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ChangeImageSizeRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeImageSizeRequest) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeRequest) SetHeight(v int32) *ChangeImageSizeRequest {
	s.Height = &v
	return s
}

func (s *ChangeImageSizeRequest) SetUrl(v string) *ChangeImageSizeRequest {
	s.Url = &v
	return s
}

func (s *ChangeImageSizeRequest) SetWidth(v int32) *ChangeImageSizeRequest {
	s.Width = &v
	return s
}

type ChangeImageSizeAdvanceRequest struct {
	Height    *int32    `json:"Height,omitempty" xml:"Height,omitempty"`
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
	Width     *int32    `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ChangeImageSizeAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeImageSizeAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeAdvanceRequest) SetHeight(v int32) *ChangeImageSizeAdvanceRequest {
	s.Height = &v
	return s
}

func (s *ChangeImageSizeAdvanceRequest) SetUrlObject(v io.Reader) *ChangeImageSizeAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *ChangeImageSizeAdvanceRequest) SetWidth(v int32) *ChangeImageSizeAdvanceRequest {
	s.Width = &v
	return s
}

type ChangeImageSizeResponseBody struct {
	Data      *ChangeImageSizeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeImageSizeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeImageSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeResponseBody) SetData(v *ChangeImageSizeResponseBodyData) *ChangeImageSizeResponseBody {
	s.Data = v
	return s
}

func (s *ChangeImageSizeResponseBody) SetRequestId(v string) *ChangeImageSizeResponseBody {
	s.RequestId = &v
	return s
}

type ChangeImageSizeResponseBodyData struct {
	RetainLocation *ChangeImageSizeResponseBodyDataRetainLocation `json:"RetainLocation,omitempty" xml:"RetainLocation,omitempty" type:"Struct"`
	Url            *string                                        `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ChangeImageSizeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ChangeImageSizeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeResponseBodyData) SetRetainLocation(v *ChangeImageSizeResponseBodyDataRetainLocation) *ChangeImageSizeResponseBodyData {
	s.RetainLocation = v
	return s
}

func (s *ChangeImageSizeResponseBodyData) SetUrl(v string) *ChangeImageSizeResponseBodyData {
	s.Url = &v
	return s
}

type ChangeImageSizeResponseBodyDataRetainLocation struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *int32 `json:"X,omitempty" xml:"X,omitempty"`
	Y      *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ChangeImageSizeResponseBodyDataRetainLocation) String() string {
	return tea.Prettify(s)
}

func (s ChangeImageSizeResponseBodyDataRetainLocation) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeResponseBodyDataRetainLocation) SetHeight(v int32) *ChangeImageSizeResponseBodyDataRetainLocation {
	s.Height = &v
	return s
}

func (s *ChangeImageSizeResponseBodyDataRetainLocation) SetWidth(v int32) *ChangeImageSizeResponseBodyDataRetainLocation {
	s.Width = &v
	return s
}

func (s *ChangeImageSizeResponseBodyDataRetainLocation) SetX(v int32) *ChangeImageSizeResponseBodyDataRetainLocation {
	s.X = &v
	return s
}

func (s *ChangeImageSizeResponseBodyDataRetainLocation) SetY(v int32) *ChangeImageSizeResponseBodyDataRetainLocation {
	s.Y = &v
	return s
}

type ChangeImageSizeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ChangeImageSizeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ChangeImageSizeResponse) SetStatusCode(v int32) *ChangeImageSizeResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeImageSizeResponse) SetBody(v *ChangeImageSizeResponseBody) *ChangeImageSizeResponse {
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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	Data      *ColorizeImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ColorizeImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ColorizeImageResponseBody) GoString() string {
	return s.String()
}

func (s *ColorizeImageResponseBody) SetData(v *ColorizeImageResponseBodyData) *ColorizeImageResponseBody {
	s.Data = v
	return s
}

func (s *ColorizeImageResponseBody) SetRequestId(v string) *ColorizeImageResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ColorizeImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ColorizeImageResponse) SetStatusCode(v int32) *ColorizeImageResponse {
	s.StatusCode = &v
	return s
}

func (s *ColorizeImageResponse) SetBody(v *ColorizeImageResponseBody) *ColorizeImageResponse {
	s.Body = v
	return s
}

type EnhanceImageColorRequest struct {
	ImageURL     *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Mode         *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
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

func (s *EnhanceImageColorRequest) SetMode(v string) *EnhanceImageColorRequest {
	s.Mode = &v
	return s
}

func (s *EnhanceImageColorRequest) SetOutputFormat(v string) *EnhanceImageColorRequest {
	s.OutputFormat = &v
	return s
}

type EnhanceImageColorAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Mode           *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	OutputFormat   *string   `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
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

func (s *EnhanceImageColorAdvanceRequest) SetMode(v string) *EnhanceImageColorAdvanceRequest {
	s.Mode = &v
	return s
}

func (s *EnhanceImageColorAdvanceRequest) SetOutputFormat(v string) *EnhanceImageColorAdvanceRequest {
	s.OutputFormat = &v
	return s
}

type EnhanceImageColorResponseBody struct {
	Data      *EnhanceImageColorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnhanceImageColorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnhanceImageColorResponseBody) GoString() string {
	return s.String()
}

func (s *EnhanceImageColorResponseBody) SetData(v *EnhanceImageColorResponseBodyData) *EnhanceImageColorResponseBody {
	s.Data = v
	return s
}

func (s *EnhanceImageColorResponseBody) SetRequestId(v string) *EnhanceImageColorResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnhanceImageColorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *EnhanceImageColorResponse) SetStatusCode(v int32) *EnhanceImageColorResponse {
	s.StatusCode = &v
	return s
}

func (s *EnhanceImageColorResponse) SetBody(v *EnhanceImageColorResponseBody) *EnhanceImageColorResponse {
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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	Data      *ErasePersonResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ErasePersonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ErasePersonResponseBody) GoString() string {
	return s.String()
}

func (s *ErasePersonResponseBody) SetData(v *ErasePersonResponseBodyData) *ErasePersonResponseBody {
	s.Data = v
	return s
}

func (s *ErasePersonResponseBody) SetRequestId(v string) *ErasePersonResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ErasePersonResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ErasePersonResponse) SetStatusCode(v int32) *ErasePersonResponse {
	s.StatusCode = &v
	return s
}

func (s *ErasePersonResponse) SetBody(v *ErasePersonResponseBody) *ErasePersonResponse {
	s.Body = v
	return s
}

type ExtendImageStyleRequest struct {
	MajorUrl *string `json:"MajorUrl,omitempty" xml:"MajorUrl,omitempty"`
	StyleUrl *string `json:"StyleUrl,omitempty" xml:"StyleUrl,omitempty"`
}

func (s ExtendImageStyleRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtendImageStyleRequest) GoString() string {
	return s.String()
}

func (s *ExtendImageStyleRequest) SetMajorUrl(v string) *ExtendImageStyleRequest {
	s.MajorUrl = &v
	return s
}

func (s *ExtendImageStyleRequest) SetStyleUrl(v string) *ExtendImageStyleRequest {
	s.StyleUrl = &v
	return s
}

type ExtendImageStyleAdvanceRequest struct {
	MajorUrlObject io.Reader `json:"MajorUrl,omitempty" xml:"MajorUrl,omitempty"`
	StyleUrlObject io.Reader `json:"StyleUrl,omitempty" xml:"StyleUrl,omitempty"`
}

func (s ExtendImageStyleAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtendImageStyleAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ExtendImageStyleAdvanceRequest) SetMajorUrlObject(v io.Reader) *ExtendImageStyleAdvanceRequest {
	s.MajorUrlObject = v
	return s
}

func (s *ExtendImageStyleAdvanceRequest) SetStyleUrlObject(v io.Reader) *ExtendImageStyleAdvanceRequest {
	s.StyleUrlObject = v
	return s
}

type ExtendImageStyleResponseBody struct {
	Data      *ExtendImageStyleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExtendImageStyleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExtendImageStyleResponseBody) GoString() string {
	return s.String()
}

func (s *ExtendImageStyleResponseBody) SetData(v *ExtendImageStyleResponseBodyData) *ExtendImageStyleResponseBody {
	s.Data = v
	return s
}

func (s *ExtendImageStyleResponseBody) SetRequestId(v string) *ExtendImageStyleResponseBody {
	s.RequestId = &v
	return s
}

type ExtendImageStyleResponseBodyData struct {
	MajorUrl *string `json:"MajorUrl,omitempty" xml:"MajorUrl,omitempty"`
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ExtendImageStyleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExtendImageStyleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExtendImageStyleResponseBodyData) SetMajorUrl(v string) *ExtendImageStyleResponseBodyData {
	s.MajorUrl = &v
	return s
}

func (s *ExtendImageStyleResponseBodyData) SetUrl(v string) *ExtendImageStyleResponseBodyData {
	s.Url = &v
	return s
}

type ExtendImageStyleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExtendImageStyleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ExtendImageStyleResponse) SetStatusCode(v int32) *ExtendImageStyleResponse {
	s.StatusCode = &v
	return s
}

func (s *ExtendImageStyleResponse) SetBody(v *ExtendImageStyleResponseBody) *ExtendImageStyleResponse {
	s.Body = v
	return s
}

type GenerateDynamicImageRequest struct {
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GenerateDynamicImageRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDynamicImageRequest) GoString() string {
	return s.String()
}

func (s *GenerateDynamicImageRequest) SetOperation(v string) *GenerateDynamicImageRequest {
	s.Operation = &v
	return s
}

func (s *GenerateDynamicImageRequest) SetUrl(v string) *GenerateDynamicImageRequest {
	s.Url = &v
	return s
}

type GenerateDynamicImageAdvanceRequest struct {
	Operation *string   `json:"Operation,omitempty" xml:"Operation,omitempty"`
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GenerateDynamicImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDynamicImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateDynamicImageAdvanceRequest) SetOperation(v string) *GenerateDynamicImageAdvanceRequest {
	s.Operation = &v
	return s
}

func (s *GenerateDynamicImageAdvanceRequest) SetUrlObject(v io.Reader) *GenerateDynamicImageAdvanceRequest {
	s.UrlObject = v
	return s
}

type GenerateDynamicImageResponseBody struct {
	Data      *GenerateDynamicImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateDynamicImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateDynamicImageResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDynamicImageResponseBody) SetData(v *GenerateDynamicImageResponseBodyData) *GenerateDynamicImageResponseBody {
	s.Data = v
	return s
}

func (s *GenerateDynamicImageResponseBody) SetRequestId(v string) *GenerateDynamicImageResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateDynamicImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GenerateDynamicImageResponse) SetStatusCode(v int32) *GenerateDynamicImageResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateDynamicImageResponse) SetBody(v *GenerateDynamicImageResponseBody) *GenerateDynamicImageResponse {
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

type ImageBlindCharacterWatermarkRequest struct {
	FunctionType      *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	OriginImageURL    *string `json:"OriginImageURL,omitempty" xml:"OriginImageURL,omitempty"`
	OutputFileType    *string `json:"OutputFileType,omitempty" xml:"OutputFileType,omitempty"`
	QualityFactor     *int32  `json:"QualityFactor,omitempty" xml:"QualityFactor,omitempty"`
	Text              *string `json:"Text,omitempty" xml:"Text,omitempty"`
	WatermarkImageURL *string `json:"WatermarkImageURL,omitempty" xml:"WatermarkImageURL,omitempty"`
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

func (s *ImageBlindCharacterWatermarkRequest) SetOriginImageURL(v string) *ImageBlindCharacterWatermarkRequest {
	s.OriginImageURL = &v
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

func (s *ImageBlindCharacterWatermarkRequest) SetText(v string) *ImageBlindCharacterWatermarkRequest {
	s.Text = &v
	return s
}

func (s *ImageBlindCharacterWatermarkRequest) SetWatermarkImageURL(v string) *ImageBlindCharacterWatermarkRequest {
	s.WatermarkImageURL = &v
	return s
}

type ImageBlindCharacterWatermarkAdvanceRequest struct {
	FunctionType            *string   `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	OriginImageURLObject    io.Reader `json:"OriginImageURL,omitempty" xml:"OriginImageURL,omitempty"`
	OutputFileType          *string   `json:"OutputFileType,omitempty" xml:"OutputFileType,omitempty"`
	QualityFactor           *int32    `json:"QualityFactor,omitempty" xml:"QualityFactor,omitempty"`
	Text                    *string   `json:"Text,omitempty" xml:"Text,omitempty"`
	WatermarkImageURLObject io.Reader `json:"WatermarkImageURL,omitempty" xml:"WatermarkImageURL,omitempty"`
}

func (s ImageBlindCharacterWatermarkAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ImageBlindCharacterWatermarkAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) SetFunctionType(v string) *ImageBlindCharacterWatermarkAdvanceRequest {
	s.FunctionType = &v
	return s
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) SetOriginImageURLObject(v io.Reader) *ImageBlindCharacterWatermarkAdvanceRequest {
	s.OriginImageURLObject = v
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

func (s *ImageBlindCharacterWatermarkAdvanceRequest) SetText(v string) *ImageBlindCharacterWatermarkAdvanceRequest {
	s.Text = &v
	return s
}

func (s *ImageBlindCharacterWatermarkAdvanceRequest) SetWatermarkImageURLObject(v io.Reader) *ImageBlindCharacterWatermarkAdvanceRequest {
	s.WatermarkImageURLObject = v
	return s
}

type ImageBlindCharacterWatermarkResponseBody struct {
	Data      *ImageBlindCharacterWatermarkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageBlindCharacterWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImageBlindCharacterWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *ImageBlindCharacterWatermarkResponseBody) SetData(v *ImageBlindCharacterWatermarkResponseBodyData) *ImageBlindCharacterWatermarkResponseBody {
	s.Data = v
	return s
}

func (s *ImageBlindCharacterWatermarkResponseBody) SetRequestId(v string) *ImageBlindCharacterWatermarkResponseBody {
	s.RequestId = &v
	return s
}

type ImageBlindCharacterWatermarkResponseBodyData struct {
	TextImageURL      *string `json:"TextImageURL,omitempty" xml:"TextImageURL,omitempty"`
	WatermarkImageURL *string `json:"WatermarkImageURL,omitempty" xml:"WatermarkImageURL,omitempty"`
}

func (s ImageBlindCharacterWatermarkResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ImageBlindCharacterWatermarkResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageBlindCharacterWatermarkResponseBodyData) SetTextImageURL(v string) *ImageBlindCharacterWatermarkResponseBodyData {
	s.TextImageURL = &v
	return s
}

func (s *ImageBlindCharacterWatermarkResponseBodyData) SetWatermarkImageURL(v string) *ImageBlindCharacterWatermarkResponseBodyData {
	s.WatermarkImageURL = &v
	return s
}

type ImageBlindCharacterWatermarkResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ImageBlindCharacterWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ImageBlindCharacterWatermarkResponse) SetStatusCode(v int32) *ImageBlindCharacterWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageBlindCharacterWatermarkResponse) SetBody(v *ImageBlindCharacterWatermarkResponseBody) *ImageBlindCharacterWatermarkResponse {
	s.Body = v
	return s
}

type ImageBlindPicWatermarkRequest struct {
	FunctionType      *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	LogoURL           *string `json:"LogoURL,omitempty" xml:"LogoURL,omitempty"`
	OriginImageURL    *string `json:"OriginImageURL,omitempty" xml:"OriginImageURL,omitempty"`
	OutputFileType    *string `json:"OutputFileType,omitempty" xml:"OutputFileType,omitempty"`
	QualityFactor     *int32  `json:"QualityFactor,omitempty" xml:"QualityFactor,omitempty"`
	WatermarkImageURL *string `json:"WatermarkImageURL,omitempty" xml:"WatermarkImageURL,omitempty"`
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

func (s *ImageBlindPicWatermarkRequest) SetOriginImageURL(v string) *ImageBlindPicWatermarkRequest {
	s.OriginImageURL = &v
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

func (s *ImageBlindPicWatermarkRequest) SetWatermarkImageURL(v string) *ImageBlindPicWatermarkRequest {
	s.WatermarkImageURL = &v
	return s
}

type ImageBlindPicWatermarkAdvanceRequest struct {
	FunctionType            *string   `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	LogoURLObject           io.Reader `json:"LogoURL,omitempty" xml:"LogoURL,omitempty"`
	OriginImageURLObject    io.Reader `json:"OriginImageURL,omitempty" xml:"OriginImageURL,omitempty"`
	OutputFileType          *string   `json:"OutputFileType,omitempty" xml:"OutputFileType,omitempty"`
	QualityFactor           *int32    `json:"QualityFactor,omitempty" xml:"QualityFactor,omitempty"`
	WatermarkImageURLObject io.Reader `json:"WatermarkImageURL,omitempty" xml:"WatermarkImageURL,omitempty"`
}

func (s ImageBlindPicWatermarkAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ImageBlindPicWatermarkAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ImageBlindPicWatermarkAdvanceRequest) SetFunctionType(v string) *ImageBlindPicWatermarkAdvanceRequest {
	s.FunctionType = &v
	return s
}

func (s *ImageBlindPicWatermarkAdvanceRequest) SetLogoURLObject(v io.Reader) *ImageBlindPicWatermarkAdvanceRequest {
	s.LogoURLObject = v
	return s
}

func (s *ImageBlindPicWatermarkAdvanceRequest) SetOriginImageURLObject(v io.Reader) *ImageBlindPicWatermarkAdvanceRequest {
	s.OriginImageURLObject = v
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

func (s *ImageBlindPicWatermarkAdvanceRequest) SetWatermarkImageURLObject(v io.Reader) *ImageBlindPicWatermarkAdvanceRequest {
	s.WatermarkImageURLObject = v
	return s
}

type ImageBlindPicWatermarkResponseBody struct {
	Data      *ImageBlindPicWatermarkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageBlindPicWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImageBlindPicWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *ImageBlindPicWatermarkResponseBody) SetData(v *ImageBlindPicWatermarkResponseBodyData) *ImageBlindPicWatermarkResponseBody {
	s.Data = v
	return s
}

func (s *ImageBlindPicWatermarkResponseBody) SetRequestId(v string) *ImageBlindPicWatermarkResponseBody {
	s.RequestId = &v
	return s
}

type ImageBlindPicWatermarkResponseBodyData struct {
	LogoURL           *string `json:"LogoURL,omitempty" xml:"LogoURL,omitempty"`
	WatermarkImageURL *string `json:"WatermarkImageURL,omitempty" xml:"WatermarkImageURL,omitempty"`
}

func (s ImageBlindPicWatermarkResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ImageBlindPicWatermarkResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageBlindPicWatermarkResponseBodyData) SetLogoURL(v string) *ImageBlindPicWatermarkResponseBodyData {
	s.LogoURL = &v
	return s
}

func (s *ImageBlindPicWatermarkResponseBodyData) SetWatermarkImageURL(v string) *ImageBlindPicWatermarkResponseBodyData {
	s.WatermarkImageURL = &v
	return s
}

type ImageBlindPicWatermarkResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ImageBlindPicWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ImageBlindPicWatermarkResponse) SetStatusCode(v int32) *ImageBlindPicWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageBlindPicWatermarkResponse) SetBody(v *ImageBlindPicWatermarkResponseBody) *ImageBlindPicWatermarkResponse {
	s.Body = v
	return s
}

type ImitatePhotoStyleRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	StyleUrl *string `json:"StyleUrl,omitempty" xml:"StyleUrl,omitempty"`
}

func (s ImitatePhotoStyleRequest) String() string {
	return tea.Prettify(s)
}

func (s ImitatePhotoStyleRequest) GoString() string {
	return s.String()
}

func (s *ImitatePhotoStyleRequest) SetImageURL(v string) *ImitatePhotoStyleRequest {
	s.ImageURL = &v
	return s
}

func (s *ImitatePhotoStyleRequest) SetStyleUrl(v string) *ImitatePhotoStyleRequest {
	s.StyleUrl = &v
	return s
}

type ImitatePhotoStyleAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	Data      *ImitatePhotoStyleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImitatePhotoStyleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImitatePhotoStyleResponseBody) GoString() string {
	return s.String()
}

func (s *ImitatePhotoStyleResponseBody) SetData(v *ImitatePhotoStyleResponseBodyData) *ImitatePhotoStyleResponseBody {
	s.Data = v
	return s
}

func (s *ImitatePhotoStyleResponseBody) SetRequestId(v string) *ImitatePhotoStyleResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ImitatePhotoStyleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ImitatePhotoStyleResponse) SetStatusCode(v int32) *ImitatePhotoStyleResponse {
	s.StatusCode = &v
	return s
}

func (s *ImitatePhotoStyleResponse) SetBody(v *ImitatePhotoStyleResponseBody) *ImitatePhotoStyleResponse {
	s.Body = v
	return s
}

type IntelligentCompositionRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	NumBoxes *int32  `json:"NumBoxes,omitempty" xml:"NumBoxes,omitempty"`
}

func (s IntelligentCompositionRequest) String() string {
	return tea.Prettify(s)
}

func (s IntelligentCompositionRequest) GoString() string {
	return s.String()
}

func (s *IntelligentCompositionRequest) SetImageURL(v string) *IntelligentCompositionRequest {
	s.ImageURL = &v
	return s
}

func (s *IntelligentCompositionRequest) SetNumBoxes(v int32) *IntelligentCompositionRequest {
	s.NumBoxes = &v
	return s
}

type IntelligentCompositionAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	Data      *IntelligentCompositionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IntelligentCompositionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IntelligentCompositionResponseBody) GoString() string {
	return s.String()
}

func (s *IntelligentCompositionResponseBody) SetData(v *IntelligentCompositionResponseBodyData) *IntelligentCompositionResponseBody {
	s.Data = v
	return s
}

func (s *IntelligentCompositionResponseBody) SetRequestId(v string) *IntelligentCompositionResponseBody {
	s.RequestId = &v
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
	MaxX  *int32   `json:"MaxX,omitempty" xml:"MaxX,omitempty"`
	MaxY  *int32   `json:"MaxY,omitempty" xml:"MaxY,omitempty"`
	MinX  *int32   `json:"MinX,omitempty" xml:"MinX,omitempty"`
	MinY  *int32   `json:"MinY,omitempty" xml:"MinY,omitempty"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s IntelligentCompositionResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s IntelligentCompositionResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *IntelligentCompositionResponseBodyDataElements) SetMaxX(v int32) *IntelligentCompositionResponseBodyDataElements {
	s.MaxX = &v
	return s
}

func (s *IntelligentCompositionResponseBodyDataElements) SetMaxY(v int32) *IntelligentCompositionResponseBodyDataElements {
	s.MaxY = &v
	return s
}

func (s *IntelligentCompositionResponseBodyDataElements) SetMinX(v int32) *IntelligentCompositionResponseBodyDataElements {
	s.MinX = &v
	return s
}

func (s *IntelligentCompositionResponseBodyDataElements) SetMinY(v int32) *IntelligentCompositionResponseBodyDataElements {
	s.MinY = &v
	return s
}

func (s *IntelligentCompositionResponseBodyDataElements) SetScore(v float32) *IntelligentCompositionResponseBodyDataElements {
	s.Score = &v
	return s
}

type IntelligentCompositionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IntelligentCompositionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *IntelligentCompositionResponse) SetStatusCode(v int32) *IntelligentCompositionResponse {
	s.StatusCode = &v
	return s
}

func (s *IntelligentCompositionResponse) SetBody(v *IntelligentCompositionResponseBody) *IntelligentCompositionResponse {
	s.Body = v
	return s
}

type MakeSuperResolutionImageRequest struct {
	Mode          *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	OutputFormat  *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	OutputQuality *int64  `json:"OutputQuality,omitempty" xml:"OutputQuality,omitempty"`
	UpscaleFactor *int64  `json:"UpscaleFactor,omitempty" xml:"UpscaleFactor,omitempty"`
	Url           *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s MakeSuperResolutionImageRequest) String() string {
	return tea.Prettify(s)
}

func (s MakeSuperResolutionImageRequest) GoString() string {
	return s.String()
}

func (s *MakeSuperResolutionImageRequest) SetMode(v string) *MakeSuperResolutionImageRequest {
	s.Mode = &v
	return s
}

func (s *MakeSuperResolutionImageRequest) SetOutputFormat(v string) *MakeSuperResolutionImageRequest {
	s.OutputFormat = &v
	return s
}

func (s *MakeSuperResolutionImageRequest) SetOutputQuality(v int64) *MakeSuperResolutionImageRequest {
	s.OutputQuality = &v
	return s
}

func (s *MakeSuperResolutionImageRequest) SetUpscaleFactor(v int64) *MakeSuperResolutionImageRequest {
	s.UpscaleFactor = &v
	return s
}

func (s *MakeSuperResolutionImageRequest) SetUrl(v string) *MakeSuperResolutionImageRequest {
	s.Url = &v
	return s
}

type MakeSuperResolutionImageAdvanceRequest struct {
	Mode          *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	OutputFormat  *string   `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	OutputQuality *int64    `json:"OutputQuality,omitempty" xml:"OutputQuality,omitempty"`
	UpscaleFactor *int64    `json:"UpscaleFactor,omitempty" xml:"UpscaleFactor,omitempty"`
	UrlObject     io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s MakeSuperResolutionImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s MakeSuperResolutionImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *MakeSuperResolutionImageAdvanceRequest) SetMode(v string) *MakeSuperResolutionImageAdvanceRequest {
	s.Mode = &v
	return s
}

func (s *MakeSuperResolutionImageAdvanceRequest) SetOutputFormat(v string) *MakeSuperResolutionImageAdvanceRequest {
	s.OutputFormat = &v
	return s
}

func (s *MakeSuperResolutionImageAdvanceRequest) SetOutputQuality(v int64) *MakeSuperResolutionImageAdvanceRequest {
	s.OutputQuality = &v
	return s
}

func (s *MakeSuperResolutionImageAdvanceRequest) SetUpscaleFactor(v int64) *MakeSuperResolutionImageAdvanceRequest {
	s.UpscaleFactor = &v
	return s
}

func (s *MakeSuperResolutionImageAdvanceRequest) SetUrlObject(v io.Reader) *MakeSuperResolutionImageAdvanceRequest {
	s.UrlObject = v
	return s
}

type MakeSuperResolutionImageResponseBody struct {
	Data      *MakeSuperResolutionImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MakeSuperResolutionImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MakeSuperResolutionImageResponseBody) GoString() string {
	return s.String()
}

func (s *MakeSuperResolutionImageResponseBody) SetData(v *MakeSuperResolutionImageResponseBodyData) *MakeSuperResolutionImageResponseBody {
	s.Data = v
	return s
}

func (s *MakeSuperResolutionImageResponseBody) SetRequestId(v string) *MakeSuperResolutionImageResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MakeSuperResolutionImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *MakeSuperResolutionImageResponse) SetStatusCode(v int32) *MakeSuperResolutionImageResponse {
	s.StatusCode = &v
	return s
}

func (s *MakeSuperResolutionImageResponse) SetBody(v *MakeSuperResolutionImageResponseBody) *MakeSuperResolutionImageResponse {
	s.Body = v
	return s
}

type RecolorHDImageRequest struct {
	ColorCount    *int32                                `json:"ColorCount,omitempty" xml:"ColorCount,omitempty"`
	ColorTemplate []*RecolorHDImageRequestColorTemplate `json:"ColorTemplate,omitempty" xml:"ColorTemplate,omitempty" type:"Repeated"`
	Degree        *string                               `json:"Degree,omitempty" xml:"Degree,omitempty"`
	Mode          *string                               `json:"Mode,omitempty" xml:"Mode,omitempty"`
	RefUrl        *string                               `json:"RefUrl,omitempty" xml:"RefUrl,omitempty"`
	Url           *string                               `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecolorHDImageRequest) String() string {
	return tea.Prettify(s)
}

func (s RecolorHDImageRequest) GoString() string {
	return s.String()
}

func (s *RecolorHDImageRequest) SetColorCount(v int32) *RecolorHDImageRequest {
	s.ColorCount = &v
	return s
}

func (s *RecolorHDImageRequest) SetColorTemplate(v []*RecolorHDImageRequestColorTemplate) *RecolorHDImageRequest {
	s.ColorTemplate = v
	return s
}

func (s *RecolorHDImageRequest) SetDegree(v string) *RecolorHDImageRequest {
	s.Degree = &v
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

func (s *RecolorHDImageRequest) SetUrl(v string) *RecolorHDImageRequest {
	s.Url = &v
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
	ColorCount    *int32                                       `json:"ColorCount,omitempty" xml:"ColorCount,omitempty"`
	ColorTemplate []*RecolorHDImageAdvanceRequestColorTemplate `json:"ColorTemplate,omitempty" xml:"ColorTemplate,omitempty" type:"Repeated"`
	Degree        *string                                      `json:"Degree,omitempty" xml:"Degree,omitempty"`
	Mode          *string                                      `json:"Mode,omitempty" xml:"Mode,omitempty"`
	RefUrlObject  io.Reader                                    `json:"RefUrl,omitempty" xml:"RefUrl,omitempty"`
	UrlObject     io.Reader                                    `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecolorHDImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecolorHDImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecolorHDImageAdvanceRequest) SetColorCount(v int32) *RecolorHDImageAdvanceRequest {
	s.ColorCount = &v
	return s
}

func (s *RecolorHDImageAdvanceRequest) SetColorTemplate(v []*RecolorHDImageAdvanceRequestColorTemplate) *RecolorHDImageAdvanceRequest {
	s.ColorTemplate = v
	return s
}

func (s *RecolorHDImageAdvanceRequest) SetDegree(v string) *RecolorHDImageAdvanceRequest {
	s.Degree = &v
	return s
}

func (s *RecolorHDImageAdvanceRequest) SetMode(v string) *RecolorHDImageAdvanceRequest {
	s.Mode = &v
	return s
}

func (s *RecolorHDImageAdvanceRequest) SetRefUrlObject(v io.Reader) *RecolorHDImageAdvanceRequest {
	s.RefUrlObject = v
	return s
}

func (s *RecolorHDImageAdvanceRequest) SetUrlObject(v io.Reader) *RecolorHDImageAdvanceRequest {
	s.UrlObject = v
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
	Data      *RecolorHDImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecolorHDImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecolorHDImageResponseBody) GoString() string {
	return s.String()
}

func (s *RecolorHDImageResponseBody) SetData(v *RecolorHDImageResponseBodyData) *RecolorHDImageResponseBody {
	s.Data = v
	return s
}

func (s *RecolorHDImageResponseBody) SetRequestId(v string) *RecolorHDImageResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecolorHDImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecolorHDImageResponse) SetStatusCode(v int32) *RecolorHDImageResponse {
	s.StatusCode = &v
	return s
}

func (s *RecolorHDImageResponse) SetBody(v *RecolorHDImageResponseBody) *RecolorHDImageResponse {
	s.Body = v
	return s
}

type RecolorImageRequest struct {
	ColorCount    *int32                              `json:"ColorCount,omitempty" xml:"ColorCount,omitempty"`
	ColorTemplate []*RecolorImageRequestColorTemplate `json:"ColorTemplate,omitempty" xml:"ColorTemplate,omitempty" type:"Repeated"`
	Mode          *string                             `json:"Mode,omitempty" xml:"Mode,omitempty"`
	RefUrl        *string                             `json:"RefUrl,omitempty" xml:"RefUrl,omitempty"`
	Url           *string                             `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecolorImageRequest) String() string {
	return tea.Prettify(s)
}

func (s RecolorImageRequest) GoString() string {
	return s.String()
}

func (s *RecolorImageRequest) SetColorCount(v int32) *RecolorImageRequest {
	s.ColorCount = &v
	return s
}

func (s *RecolorImageRequest) SetColorTemplate(v []*RecolorImageRequestColorTemplate) *RecolorImageRequest {
	s.ColorTemplate = v
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

func (s *RecolorImageRequest) SetUrl(v string) *RecolorImageRequest {
	s.Url = &v
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

type RecolorImageAdvanceRequest struct {
	ColorCount    *int32                                     `json:"ColorCount,omitempty" xml:"ColorCount,omitempty"`
	ColorTemplate []*RecolorImageAdvanceRequestColorTemplate `json:"ColorTemplate,omitempty" xml:"ColorTemplate,omitempty" type:"Repeated"`
	Mode          *string                                    `json:"Mode,omitempty" xml:"Mode,omitempty"`
	RefUrlObject  io.Reader                                  `json:"RefUrl,omitempty" xml:"RefUrl,omitempty"`
	UrlObject     io.Reader                                  `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecolorImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecolorImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecolorImageAdvanceRequest) SetColorCount(v int32) *RecolorImageAdvanceRequest {
	s.ColorCount = &v
	return s
}

func (s *RecolorImageAdvanceRequest) SetColorTemplate(v []*RecolorImageAdvanceRequestColorTemplate) *RecolorImageAdvanceRequest {
	s.ColorTemplate = v
	return s
}

func (s *RecolorImageAdvanceRequest) SetMode(v string) *RecolorImageAdvanceRequest {
	s.Mode = &v
	return s
}

func (s *RecolorImageAdvanceRequest) SetRefUrlObject(v io.Reader) *RecolorImageAdvanceRequest {
	s.RefUrlObject = v
	return s
}

func (s *RecolorImageAdvanceRequest) SetUrlObject(v io.Reader) *RecolorImageAdvanceRequest {
	s.UrlObject = v
	return s
}

type RecolorImageAdvanceRequestColorTemplate struct {
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
}

func (s RecolorImageAdvanceRequestColorTemplate) String() string {
	return tea.Prettify(s)
}

func (s RecolorImageAdvanceRequestColorTemplate) GoString() string {
	return s.String()
}

func (s *RecolorImageAdvanceRequestColorTemplate) SetColor(v string) *RecolorImageAdvanceRequestColorTemplate {
	s.Color = &v
	return s
}

type RecolorImageResponseBody struct {
	Data      *RecolorImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecolorImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecolorImageResponseBody) GoString() string {
	return s.String()
}

func (s *RecolorImageResponseBody) SetData(v *RecolorImageResponseBodyData) *RecolorImageResponseBody {
	s.Data = v
	return s
}

func (s *RecolorImageResponseBody) SetRequestId(v string) *RecolorImageResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecolorImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecolorImageResponse) SetStatusCode(v int32) *RecolorImageResponse {
	s.StatusCode = &v
	return s
}

func (s *RecolorImageResponse) SetBody(v *RecolorImageResponseBody) *RecolorImageResponse {
	s.Body = v
	return s
}

type RemoveImageSubtitlesRequest struct {
	BH       *float32 `json:"BH,omitempty" xml:"BH,omitempty"`
	BW       *float32 `json:"BW,omitempty" xml:"BW,omitempty"`
	BX       *float32 `json:"BX,omitempty" xml:"BX,omitempty"`
	BY       *float32 `json:"BY,omitempty" xml:"BY,omitempty"`
	ImageURL *string  `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RemoveImageSubtitlesRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageSubtitlesRequest) GoString() string {
	return s.String()
}

func (s *RemoveImageSubtitlesRequest) SetBH(v float32) *RemoveImageSubtitlesRequest {
	s.BH = &v
	return s
}

func (s *RemoveImageSubtitlesRequest) SetBW(v float32) *RemoveImageSubtitlesRequest {
	s.BW = &v
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

func (s *RemoveImageSubtitlesRequest) SetImageURL(v string) *RemoveImageSubtitlesRequest {
	s.ImageURL = &v
	return s
}

type RemoveImageSubtitlesAdvanceRequest struct {
	BH             *float32  `json:"BH,omitempty" xml:"BH,omitempty"`
	BW             *float32  `json:"BW,omitempty" xml:"BW,omitempty"`
	BX             *float32  `json:"BX,omitempty" xml:"BX,omitempty"`
	BY             *float32  `json:"BY,omitempty" xml:"BY,omitempty"`
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RemoveImageSubtitlesAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageSubtitlesAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RemoveImageSubtitlesAdvanceRequest) SetBH(v float32) *RemoveImageSubtitlesAdvanceRequest {
	s.BH = &v
	return s
}

func (s *RemoveImageSubtitlesAdvanceRequest) SetBW(v float32) *RemoveImageSubtitlesAdvanceRequest {
	s.BW = &v
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

func (s *RemoveImageSubtitlesAdvanceRequest) SetImageURLObject(v io.Reader) *RemoveImageSubtitlesAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RemoveImageSubtitlesResponseBody struct {
	Data      *RemoveImageSubtitlesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveImageSubtitlesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageSubtitlesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveImageSubtitlesResponseBody) SetData(v *RemoveImageSubtitlesResponseBodyData) *RemoveImageSubtitlesResponseBody {
	s.Data = v
	return s
}

func (s *RemoveImageSubtitlesResponseBody) SetRequestId(v string) *RemoveImageSubtitlesResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveImageSubtitlesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveImageSubtitlesResponse) SetStatusCode(v int32) *RemoveImageSubtitlesResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveImageSubtitlesResponse) SetBody(v *RemoveImageSubtitlesResponseBody) *RemoveImageSubtitlesResponse {
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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	Data      *RemoveImageWatermarkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveImageWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveImageWatermarkResponseBody) SetData(v *RemoveImageWatermarkResponseBodyData) *RemoveImageWatermarkResponseBody {
	s.Data = v
	return s
}

func (s *RemoveImageWatermarkResponseBody) SetRequestId(v string) *RemoveImageWatermarkResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveImageWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveImageWatermarkResponse) SetStatusCode(v int32) *RemoveImageWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveImageWatermarkResponse) SetBody(v *RemoveImageWatermarkResponseBody) *RemoveImageWatermarkResponse {
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

func (client *Client) AssessCompositionWithOptions(request *AssessCompositionRequest, runtime *util.RuntimeOptions) (_result *AssessCompositionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AssessComposition"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssessCompositionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		assessCompositionReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	assessCompositionResp, _err := client.AssessCompositionWithOptions(assessCompositionReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = assessCompositionResp
	return _result, _err
}

func (client *Client) AssessExposureWithOptions(request *AssessExposureRequest, runtime *util.RuntimeOptions) (_result *AssessExposureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AssessExposure"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssessExposureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		assessExposureReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	assessExposureResp, _err := client.AssessExposureWithOptions(assessExposureReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = assessExposureResp
	return _result, _err
}

func (client *Client) AssessSharpnessWithOptions(request *AssessSharpnessRequest, runtime *util.RuntimeOptions) (_result *AssessSharpnessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AssessSharpness"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssessSharpnessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		assessSharpnessReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	assessSharpnessResp, _err := client.AssessSharpnessWithOptions(assessSharpnessReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = assessSharpnessResp
	return _result, _err
}

func (client *Client) ChangeImageSizeWithOptions(request *ChangeImageSizeRequest, runtime *util.RuntimeOptions) (_result *ChangeImageSizeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Height)) {
		body["Height"] = request.Height
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	if !tea.BoolValue(util.IsUnset(request.Width)) {
		body["Width"] = request.Width
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeImageSize"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeImageSizeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.UrlObject)) {
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
			Content:     request.UrlObject,
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
		changeImageSizeReq.Url = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	changeImageSizeResp, _err := client.ChangeImageSizeWithOptions(changeImageSizeReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = changeImageSizeResp
	return _result, _err
}

func (client *Client) ColorizeImageWithOptions(request *ColorizeImageRequest, runtime *util.RuntimeOptions) (_result *ColorizeImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ColorizeImage"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ColorizeImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		colorizeImageReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	colorizeImageResp, _err := client.ColorizeImageWithOptions(colorizeImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = colorizeImageResp
	return _result, _err
}

func (client *Client) EnhanceImageColorWithOptions(request *EnhanceImageColorRequest, runtime *util.RuntimeOptions) (_result *EnhanceImageColorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.OutputFormat)) {
		body["OutputFormat"] = request.OutputFormat
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnhanceImageColor"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnhanceImageColorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		enhanceImageColorReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	enhanceImageColorResp, _err := client.EnhanceImageColorWithOptions(enhanceImageColorReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = enhanceImageColorResp
	return _result, _err
}

func (client *Client) ErasePersonWithOptions(request *ErasePersonRequest, runtime *util.RuntimeOptions) (_result *ErasePersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.UserMask)) {
		body["UserMask"] = request.UserMask
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ErasePerson"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ErasePersonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		erasePersonReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	erasePersonResp, _err := client.ErasePersonWithOptions(erasePersonReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = erasePersonResp
	return _result, _err
}

func (client *Client) ExtendImageStyleWithOptions(request *ExtendImageStyleRequest, runtime *util.RuntimeOptions) (_result *ExtendImageStyleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MajorUrl)) {
		body["MajorUrl"] = request.MajorUrl
	}

	if !tea.BoolValue(util.IsUnset(request.StyleUrl)) {
		body["StyleUrl"] = request.StyleUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExtendImageStyle"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExtendImageStyleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ExtendImageStyleAdvance(request *ExtendImageStyleAdvanceRequest, runtime *util.RuntimeOptions) (_result *ExtendImageStyleResponse, _err error) {
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
	extendImageStyleReq := &ExtendImageStyleRequest{}
	openapiutil.Convert(request, extendImageStyleReq)
	if !tea.BoolValue(util.IsUnset(request.MajorUrlObject)) {
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
			Content:     request.MajorUrlObject,
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
		extendImageStyleReq.MajorUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	if !tea.BoolValue(util.IsUnset(request.StyleUrlObject)) {
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
			Content:     request.StyleUrlObject,
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
		extendImageStyleReq.StyleUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	extendImageStyleResp, _err := client.ExtendImageStyleWithOptions(extendImageStyleReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = extendImageStyleResp
	return _result, _err
}

func (client *Client) GenerateDynamicImageWithOptions(request *GenerateDynamicImageRequest, runtime *util.RuntimeOptions) (_result *GenerateDynamicImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		body["Operation"] = request.Operation
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateDynamicImage"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateDynamicImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.UrlObject)) {
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
			Content:     request.UrlObject,
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
		generateDynamicImageReq.Url = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	generateDynamicImageResp, _err := client.GenerateDynamicImageWithOptions(generateDynamicImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = generateDynamicImageResp
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
		Version:     tea.String("2019-09-30"),
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

func (client *Client) ImageBlindCharacterWatermarkWithOptions(request *ImageBlindCharacterWatermarkRequest, runtime *util.RuntimeOptions) (_result *ImageBlindCharacterWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FunctionType)) {
		body["FunctionType"] = request.FunctionType
	}

	if !tea.BoolValue(util.IsUnset(request.OriginImageURL)) {
		body["OriginImageURL"] = request.OriginImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.OutputFileType)) {
		body["OutputFileType"] = request.OutputFileType
	}

	if !tea.BoolValue(util.IsUnset(request.QualityFactor)) {
		body["QualityFactor"] = request.QualityFactor
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkImageURL)) {
		body["WatermarkImageURL"] = request.WatermarkImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ImageBlindCharacterWatermark"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImageBlindCharacterWatermarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.OriginImageURLObject)) {
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
			Content:     request.OriginImageURLObject,
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
		imageBlindCharacterWatermarkReq.OriginImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkImageURLObject)) {
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
			Content:     request.WatermarkImageURLObject,
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
		imageBlindCharacterWatermarkReq.WatermarkImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	imageBlindCharacterWatermarkResp, _err := client.ImageBlindCharacterWatermarkWithOptions(imageBlindCharacterWatermarkReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = imageBlindCharacterWatermarkResp
	return _result, _err
}

func (client *Client) ImageBlindPicWatermarkWithOptions(request *ImageBlindPicWatermarkRequest, runtime *util.RuntimeOptions) (_result *ImageBlindPicWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FunctionType)) {
		body["FunctionType"] = request.FunctionType
	}

	if !tea.BoolValue(util.IsUnset(request.LogoURL)) {
		body["LogoURL"] = request.LogoURL
	}

	if !tea.BoolValue(util.IsUnset(request.OriginImageURL)) {
		body["OriginImageURL"] = request.OriginImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.OutputFileType)) {
		body["OutputFileType"] = request.OutputFileType
	}

	if !tea.BoolValue(util.IsUnset(request.QualityFactor)) {
		body["QualityFactor"] = request.QualityFactor
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkImageURL)) {
		body["WatermarkImageURL"] = request.WatermarkImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ImageBlindPicWatermark"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImageBlindPicWatermarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.LogoURLObject)) {
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
			Content:     request.LogoURLObject,
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
		imageBlindPicWatermarkReq.LogoURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	if !tea.BoolValue(util.IsUnset(request.OriginImageURLObject)) {
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
			Content:     request.OriginImageURLObject,
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
		imageBlindPicWatermarkReq.OriginImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkImageURLObject)) {
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
			Content:     request.WatermarkImageURLObject,
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
		imageBlindPicWatermarkReq.WatermarkImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	imageBlindPicWatermarkResp, _err := client.ImageBlindPicWatermarkWithOptions(imageBlindPicWatermarkReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = imageBlindPicWatermarkResp
	return _result, _err
}

func (client *Client) ImitatePhotoStyleWithOptions(request *ImitatePhotoStyleRequest, runtime *util.RuntimeOptions) (_result *ImitatePhotoStyleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.StyleUrl)) {
		body["StyleUrl"] = request.StyleUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ImitatePhotoStyle"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImitatePhotoStyleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		imitatePhotoStyleReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	imitatePhotoStyleResp, _err := client.ImitatePhotoStyleWithOptions(imitatePhotoStyleReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = imitatePhotoStyleResp
	return _result, _err
}

func (client *Client) IntelligentCompositionWithOptions(request *IntelligentCompositionRequest, runtime *util.RuntimeOptions) (_result *IntelligentCompositionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.NumBoxes)) {
		body["NumBoxes"] = request.NumBoxes
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("IntelligentComposition"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &IntelligentCompositionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		intelligentCompositionReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	intelligentCompositionResp, _err := client.IntelligentCompositionWithOptions(intelligentCompositionReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = intelligentCompositionResp
	return _result, _err
}

func (client *Client) MakeSuperResolutionImageWithOptions(request *MakeSuperResolutionImageRequest, runtime *util.RuntimeOptions) (_result *MakeSuperResolutionImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.OutputFormat)) {
		body["OutputFormat"] = request.OutputFormat
	}

	if !tea.BoolValue(util.IsUnset(request.OutputQuality)) {
		body["OutputQuality"] = request.OutputQuality
	}

	if !tea.BoolValue(util.IsUnset(request.UpscaleFactor)) {
		body["UpscaleFactor"] = request.UpscaleFactor
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MakeSuperResolutionImage"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MakeSuperResolutionImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.UrlObject)) {
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
			Content:     request.UrlObject,
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
		makeSuperResolutionImageReq.Url = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	makeSuperResolutionImageResp, _err := client.MakeSuperResolutionImageWithOptions(makeSuperResolutionImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = makeSuperResolutionImageResp
	return _result, _err
}

func (client *Client) RecolorHDImageWithOptions(request *RecolorHDImageRequest, runtime *util.RuntimeOptions) (_result *RecolorHDImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ColorCount)) {
		body["ColorCount"] = request.ColorCount
	}

	if !tea.BoolValue(util.IsUnset(request.ColorTemplate)) {
		body["ColorTemplate"] = request.ColorTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.Degree)) {
		body["Degree"] = request.Degree
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.RefUrl)) {
		body["RefUrl"] = request.RefUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecolorHDImage"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecolorHDImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.RefUrlObject)) {
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
			Content:     request.RefUrlObject,
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
		recolorHDImageReq.RefUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	if !tea.BoolValue(util.IsUnset(request.UrlObject)) {
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
			Content:     request.UrlObject,
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
		recolorHDImageReq.Url = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recolorHDImageResp, _err := client.RecolorHDImageWithOptions(recolorHDImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recolorHDImageResp
	return _result, _err
}

func (client *Client) RecolorImageWithOptions(request *RecolorImageRequest, runtime *util.RuntimeOptions) (_result *RecolorImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ColorCount)) {
		body["ColorCount"] = request.ColorCount
	}

	if !tea.BoolValue(util.IsUnset(request.ColorTemplate)) {
		body["ColorTemplate"] = request.ColorTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.RefUrl)) {
		body["RefUrl"] = request.RefUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecolorImage"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecolorImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) RecolorImageAdvance(request *RecolorImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecolorImageResponse, _err error) {
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
	recolorImageReq := &RecolorImageRequest{}
	openapiutil.Convert(request, recolorImageReq)
	if !tea.BoolValue(util.IsUnset(request.RefUrlObject)) {
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
			Content:     request.RefUrlObject,
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
		recolorImageReq.RefUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	if !tea.BoolValue(util.IsUnset(request.UrlObject)) {
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
			Content:     request.UrlObject,
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
		recolorImageReq.Url = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recolorImageResp, _err := client.RecolorImageWithOptions(recolorImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recolorImageResp
	return _result, _err
}

func (client *Client) RemoveImageSubtitlesWithOptions(request *RemoveImageSubtitlesRequest, runtime *util.RuntimeOptions) (_result *RemoveImageSubtitlesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveImageSubtitles"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveImageSubtitlesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		removeImageSubtitlesReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	removeImageSubtitlesResp, _err := client.RemoveImageSubtitlesWithOptions(removeImageSubtitlesReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = removeImageSubtitlesResp
	return _result, _err
}

func (client *Client) RemoveImageWatermarkWithOptions(request *RemoveImageWatermarkRequest, runtime *util.RuntimeOptions) (_result *RemoveImageWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveImageWatermark"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveImageWatermarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
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
			Content:     request.ImageURLObject,
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
		removeImageWatermarkReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	removeImageWatermarkResp, _err := client.RemoveImageWatermarkWithOptions(removeImageWatermarkReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = removeImageWatermarkResp
	return _result, _err
}
