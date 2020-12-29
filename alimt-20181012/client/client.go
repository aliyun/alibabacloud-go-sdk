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

type CreateDocTranslateTaskRequest struct {
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	FileUrl        *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Scene          *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	CallbackUrl    *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateDocTranslateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDocTranslateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDocTranslateTaskRequest) SetSourceLanguage(v string) *CreateDocTranslateTaskRequest {
	s.SourceLanguage = &v
	return s
}

func (s *CreateDocTranslateTaskRequest) SetTargetLanguage(v string) *CreateDocTranslateTaskRequest {
	s.TargetLanguage = &v
	return s
}

func (s *CreateDocTranslateTaskRequest) SetFileUrl(v string) *CreateDocTranslateTaskRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateDocTranslateTaskRequest) SetScene(v string) *CreateDocTranslateTaskRequest {
	s.Scene = &v
	return s
}

func (s *CreateDocTranslateTaskRequest) SetCallbackUrl(v string) *CreateDocTranslateTaskRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateDocTranslateTaskRequest) SetClientToken(v string) *CreateDocTranslateTaskRequest {
	s.ClientToken = &v
	return s
}

type CreateDocTranslateTaskAdvanceRequest struct {
	FileUrlObject  io.Reader `json:"FileUrlObject,omitempty" xml:"FileUrlObject,omitempty" require:"true"`
	SourceLanguage *string   `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	TargetLanguage *string   `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	Scene          *string   `json:"Scene,omitempty" xml:"Scene,omitempty"`
	CallbackUrl    *string   `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	ClientToken    *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateDocTranslateTaskAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDocTranslateTaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDocTranslateTaskAdvanceRequest) SetFileUrlObject(v io.Reader) *CreateDocTranslateTaskAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *CreateDocTranslateTaskAdvanceRequest) SetSourceLanguage(v string) *CreateDocTranslateTaskAdvanceRequest {
	s.SourceLanguage = &v
	return s
}

func (s *CreateDocTranslateTaskAdvanceRequest) SetTargetLanguage(v string) *CreateDocTranslateTaskAdvanceRequest {
	s.TargetLanguage = &v
	return s
}

func (s *CreateDocTranslateTaskAdvanceRequest) SetScene(v string) *CreateDocTranslateTaskAdvanceRequest {
	s.Scene = &v
	return s
}

func (s *CreateDocTranslateTaskAdvanceRequest) SetCallbackUrl(v string) *CreateDocTranslateTaskAdvanceRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateDocTranslateTaskAdvanceRequest) SetClientToken(v string) *CreateDocTranslateTaskAdvanceRequest {
	s.ClientToken = &v
	return s
}

type CreateDocTranslateTaskResponseBody struct {
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDocTranslateTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDocTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDocTranslateTaskResponseBody) SetStatus(v string) *CreateDocTranslateTaskResponseBody {
	s.Status = &v
	return s
}

func (s *CreateDocTranslateTaskResponseBody) SetTaskId(v string) *CreateDocTranslateTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateDocTranslateTaskResponseBody) SetRequestId(v string) *CreateDocTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateDocTranslateTaskResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDocTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDocTranslateTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDocTranslateTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDocTranslateTaskResponse) SetHeaders(v map[string]*string) *CreateDocTranslateTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDocTranslateTaskResponse) SetBody(v *CreateDocTranslateTaskResponseBody) *CreateDocTranslateTaskResponse {
	s.Body = v
	return s
}

type CreateImageTranslateTaskRequest struct {
	UrlList        *string `json:"UrlList,omitempty" xml:"UrlList,omitempty"`
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	Extra          *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateImageTranslateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageTranslateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateImageTranslateTaskRequest) SetUrlList(v string) *CreateImageTranslateTaskRequest {
	s.UrlList = &v
	return s
}

func (s *CreateImageTranslateTaskRequest) SetSourceLanguage(v string) *CreateImageTranslateTaskRequest {
	s.SourceLanguage = &v
	return s
}

func (s *CreateImageTranslateTaskRequest) SetTargetLanguage(v string) *CreateImageTranslateTaskRequest {
	s.TargetLanguage = &v
	return s
}

func (s *CreateImageTranslateTaskRequest) SetExtra(v string) *CreateImageTranslateTaskRequest {
	s.Extra = &v
	return s
}

func (s *CreateImageTranslateTaskRequest) SetClientToken(v string) *CreateImageTranslateTaskRequest {
	s.ClientToken = &v
	return s
}

type CreateImageTranslateTaskResponseBody struct {
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CreateImageTranslateTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateImageTranslateTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageTranslateTaskResponseBody) SetMessage(v string) *CreateImageTranslateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateImageTranslateTaskResponseBody) SetRequestId(v string) *CreateImageTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageTranslateTaskResponseBody) SetData(v *CreateImageTranslateTaskResponseBodyData) *CreateImageTranslateTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateImageTranslateTaskResponseBody) SetCode(v int32) *CreateImageTranslateTaskResponseBody {
	s.Code = &v
	return s
}

type CreateImageTranslateTaskResponseBodyData struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateImageTranslateTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateImageTranslateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateImageTranslateTaskResponseBodyData) SetTaskId(v string) *CreateImageTranslateTaskResponseBodyData {
	s.TaskId = &v
	return s
}

type CreateImageTranslateTaskResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateImageTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateImageTranslateTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageTranslateTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateImageTranslateTaskResponse) SetHeaders(v map[string]*string) *CreateImageTranslateTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateImageTranslateTaskResponse) SetBody(v *CreateImageTranslateTaskResponseBody) *CreateImageTranslateTaskResponse {
	s.Body = v
	return s
}

type GetDetectLanguageRequest struct {
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
}

func (s GetDetectLanguageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDetectLanguageRequest) GoString() string {
	return s.String()
}

func (s *GetDetectLanguageRequest) SetSourceText(v string) *GetDetectLanguageRequest {
	s.SourceText = &v
	return s
}

type GetDetectLanguageResponseBody struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DetectedLanguage *string `json:"DetectedLanguage,omitempty" xml:"DetectedLanguage,omitempty"`
}

func (s GetDetectLanguageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDetectLanguageResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetectLanguageResponseBody) SetRequestId(v string) *GetDetectLanguageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDetectLanguageResponseBody) SetDetectedLanguage(v string) *GetDetectLanguageResponseBody {
	s.DetectedLanguage = &v
	return s
}

type GetDetectLanguageResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDetectLanguageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDetectLanguageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDetectLanguageResponse) GoString() string {
	return s.String()
}

func (s *GetDetectLanguageResponse) SetHeaders(v map[string]*string) *GetDetectLanguageResponse {
	s.Headers = v
	return s
}

func (s *GetDetectLanguageResponse) SetBody(v *GetDetectLanguageResponseBody) *GetDetectLanguageResponse {
	s.Body = v
	return s
}

type GetDocTranslateTaskRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetDocTranslateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocTranslateTaskRequest) GoString() string {
	return s.String()
}

func (s *GetDocTranslateTaskRequest) SetTaskId(v string) *GetDocTranslateTaskRequest {
	s.TaskId = &v
	return s
}

type GetDocTranslateTaskResponseBody struct {
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TranslateErrorMessage *string `json:"TranslateErrorMessage,omitempty" xml:"TranslateErrorMessage,omitempty"`
	TaskId                *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageCount             *int32  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	TranslateFileUrl      *string `json:"TranslateFileUrl,omitempty" xml:"TranslateFileUrl,omitempty"`
	TranslateErrorCode    *string `json:"TranslateErrorCode,omitempty" xml:"TranslateErrorCode,omitempty"`
}

func (s GetDocTranslateTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocTranslateTaskResponseBody) SetStatus(v string) *GetDocTranslateTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetTranslateErrorMessage(v string) *GetDocTranslateTaskResponseBody {
	s.TranslateErrorMessage = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetTaskId(v string) *GetDocTranslateTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetRequestId(v string) *GetDocTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetPageCount(v int32) *GetDocTranslateTaskResponseBody {
	s.PageCount = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetTranslateFileUrl(v string) *GetDocTranslateTaskResponseBody {
	s.TranslateFileUrl = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetTranslateErrorCode(v string) *GetDocTranslateTaskResponseBody {
	s.TranslateErrorCode = &v
	return s
}

type GetDocTranslateTaskResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDocTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDocTranslateTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocTranslateTaskResponse) GoString() string {
	return s.String()
}

func (s *GetDocTranslateTaskResponse) SetHeaders(v map[string]*string) *GetDocTranslateTaskResponse {
	s.Headers = v
	return s
}

func (s *GetDocTranslateTaskResponse) SetBody(v *GetDocTranslateTaskResponseBody) *GetDocTranslateTaskResponse {
	s.Body = v
	return s
}

type GetImageDiagnoseRequest struct {
	Url   *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s GetImageDiagnoseRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageDiagnoseRequest) GoString() string {
	return s.String()
}

func (s *GetImageDiagnoseRequest) SetUrl(v string) *GetImageDiagnoseRequest {
	s.Url = &v
	return s
}

func (s *GetImageDiagnoseRequest) SetExtra(v string) *GetImageDiagnoseRequest {
	s.Extra = &v
	return s
}

type GetImageDiagnoseResponseBody struct {
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetImageDiagnoseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetImageDiagnoseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageDiagnoseResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageDiagnoseResponseBody) SetMessage(v string) *GetImageDiagnoseResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageDiagnoseResponseBody) SetRequestId(v string) *GetImageDiagnoseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageDiagnoseResponseBody) SetData(v *GetImageDiagnoseResponseBodyData) *GetImageDiagnoseResponseBody {
	s.Data = v
	return s
}

func (s *GetImageDiagnoseResponseBody) SetCode(v int32) *GetImageDiagnoseResponseBody {
	s.Code = &v
	return s
}

type GetImageDiagnoseResponseBodyData struct {
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s GetImageDiagnoseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetImageDiagnoseResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetImageDiagnoseResponseBodyData) SetLanguage(v string) *GetImageDiagnoseResponseBodyData {
	s.Language = &v
	return s
}

type GetImageDiagnoseResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetImageDiagnoseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetImageDiagnoseResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImageDiagnoseResponse) GoString() string {
	return s.String()
}

func (s *GetImageDiagnoseResponse) SetHeaders(v map[string]*string) *GetImageDiagnoseResponse {
	s.Headers = v
	return s
}

func (s *GetImageDiagnoseResponse) SetBody(v *GetImageDiagnoseResponseBody) *GetImageDiagnoseResponse {
	s.Body = v
	return s
}

type GetImageTranslateRequest struct {
	Url            *string `json:"Url,omitempty" xml:"Url,omitempty"`
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	Extra          *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s GetImageTranslateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageTranslateRequest) GoString() string {
	return s.String()
}

func (s *GetImageTranslateRequest) SetUrl(v string) *GetImageTranslateRequest {
	s.Url = &v
	return s
}

func (s *GetImageTranslateRequest) SetSourceLanguage(v string) *GetImageTranslateRequest {
	s.SourceLanguage = &v
	return s
}

func (s *GetImageTranslateRequest) SetTargetLanguage(v string) *GetImageTranslateRequest {
	s.TargetLanguage = &v
	return s
}

func (s *GetImageTranslateRequest) SetExtra(v string) *GetImageTranslateRequest {
	s.Extra = &v
	return s
}

type GetImageTranslateResponseBody struct {
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetImageTranslateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetImageTranslateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageTranslateResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageTranslateResponseBody) SetMessage(v string) *GetImageTranslateResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageTranslateResponseBody) SetRequestId(v string) *GetImageTranslateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageTranslateResponseBody) SetData(v *GetImageTranslateResponseBodyData) *GetImageTranslateResponseBody {
	s.Data = v
	return s
}

func (s *GetImageTranslateResponseBody) SetCode(v int32) *GetImageTranslateResponseBody {
	s.Code = &v
	return s
}

type GetImageTranslateResponseBodyData struct {
	Url           *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Orc           *string `json:"Orc,omitempty" xml:"Orc,omitempty"`
	PictureEditor *string `json:"PictureEditor,omitempty" xml:"PictureEditor,omitempty"`
}

func (s GetImageTranslateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetImageTranslateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetImageTranslateResponseBodyData) SetUrl(v string) *GetImageTranslateResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetImageTranslateResponseBodyData) SetOrc(v string) *GetImageTranslateResponseBodyData {
	s.Orc = &v
	return s
}

func (s *GetImageTranslateResponseBodyData) SetPictureEditor(v string) *GetImageTranslateResponseBodyData {
	s.PictureEditor = &v
	return s
}

type GetImageTranslateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetImageTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetImageTranslateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImageTranslateResponse) GoString() string {
	return s.String()
}

func (s *GetImageTranslateResponse) SetHeaders(v map[string]*string) *GetImageTranslateResponse {
	s.Headers = v
	return s
}

func (s *GetImageTranslateResponse) SetBody(v *GetImageTranslateResponseBody) *GetImageTranslateResponse {
	s.Body = v
	return s
}

type GetImageTranslateTaskRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetImageTranslateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageTranslateTaskRequest) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskRequest) SetTaskId(v string) *GetImageTranslateTaskRequest {
	s.TaskId = &v
	return s
}

type GetImageTranslateTaskResponseBody struct {
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetImageTranslateTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetImageTranslateTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskResponseBody) SetMessage(v string) *GetImageTranslateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageTranslateTaskResponseBody) SetRequestId(v string) *GetImageTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageTranslateTaskResponseBody) SetData(v *GetImageTranslateTaskResponseBodyData) *GetImageTranslateTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetImageTranslateTaskResponseBody) SetCode(v int32) *GetImageTranslateTaskResponseBody {
	s.Code = &v
	return s
}

type GetImageTranslateTaskResponseBodyData struct {
	ImageData *string `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
}

func (s GetImageTranslateTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetImageTranslateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskResponseBodyData) SetImageData(v string) *GetImageTranslateTaskResponseBodyData {
	s.ImageData = &v
	return s
}

type GetImageTranslateTaskResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetImageTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetImageTranslateTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImageTranslateTaskResponse) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskResponse) SetHeaders(v map[string]*string) *GetImageTranslateTaskResponse {
	s.Headers = v
	return s
}

func (s *GetImageTranslateTaskResponse) SetBody(v *GetImageTranslateTaskResponseBody) *GetImageTranslateTaskResponse {
	s.Body = v
	return s
}

type GetTitleDiagnoseRequest struct {
	Title      *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Platform   *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Extra      *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s GetTitleDiagnoseRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTitleDiagnoseRequest) GoString() string {
	return s.String()
}

func (s *GetTitleDiagnoseRequest) SetTitle(v string) *GetTitleDiagnoseRequest {
	s.Title = &v
	return s
}

func (s *GetTitleDiagnoseRequest) SetLanguage(v string) *GetTitleDiagnoseRequest {
	s.Language = &v
	return s
}

func (s *GetTitleDiagnoseRequest) SetPlatform(v string) *GetTitleDiagnoseRequest {
	s.Platform = &v
	return s
}

func (s *GetTitleDiagnoseRequest) SetCategoryId(v string) *GetTitleDiagnoseRequest {
	s.CategoryId = &v
	return s
}

func (s *GetTitleDiagnoseRequest) SetExtra(v string) *GetTitleDiagnoseRequest {
	s.Extra = &v
	return s
}

type GetTitleDiagnoseResponseBody struct {
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetTitleDiagnoseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetTitleDiagnoseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTitleDiagnoseResponseBody) GoString() string {
	return s.String()
}

func (s *GetTitleDiagnoseResponseBody) SetMessage(v string) *GetTitleDiagnoseResponseBody {
	s.Message = &v
	return s
}

func (s *GetTitleDiagnoseResponseBody) SetRequestId(v string) *GetTitleDiagnoseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTitleDiagnoseResponseBody) SetData(v *GetTitleDiagnoseResponseBodyData) *GetTitleDiagnoseResponseBody {
	s.Data = v
	return s
}

func (s *GetTitleDiagnoseResponseBody) SetCode(v int32) *GetTitleDiagnoseResponseBody {
	s.Code = &v
	return s
}

type GetTitleDiagnoseResponseBodyData struct {
	DuplicateWords          *string `json:"DuplicateWords,omitempty" xml:"DuplicateWords,omitempty"`
	ContainCoreClasses      *string `json:"ContainCoreClasses,omitempty" xml:"ContainCoreClasses,omitempty"`
	WordCount               *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
	LanguageQualityScore    *string `json:"LanguageQualityScore,omitempty" xml:"LanguageQualityScore,omitempty"`
	AllUppercaseWords       *string `json:"AllUppercaseWords,omitempty" xml:"AllUppercaseWords,omitempty"`
	OverLengthLimit         *string `json:"OverLengthLimit,omitempty" xml:"OverLengthLimit,omitempty"`
	DisableWords            *string `json:"DisableWords,omitempty" xml:"DisableWords,omitempty"`
	NoFirstUppercaseList    *string `json:"NoFirstUppercaseList,omitempty" xml:"NoFirstUppercaseList,omitempty"`
	TotalScore              *string `json:"TotalScore,omitempty" xml:"TotalScore,omitempty"`
	WordSpelledCorrectError *string `json:"WordSpelledCorrectError,omitempty" xml:"WordSpelledCorrectError,omitempty"`
}

func (s GetTitleDiagnoseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTitleDiagnoseResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTitleDiagnoseResponseBodyData) SetDuplicateWords(v string) *GetTitleDiagnoseResponseBodyData {
	s.DuplicateWords = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetContainCoreClasses(v string) *GetTitleDiagnoseResponseBodyData {
	s.ContainCoreClasses = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetWordCount(v string) *GetTitleDiagnoseResponseBodyData {
	s.WordCount = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetLanguageQualityScore(v string) *GetTitleDiagnoseResponseBodyData {
	s.LanguageQualityScore = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetAllUppercaseWords(v string) *GetTitleDiagnoseResponseBodyData {
	s.AllUppercaseWords = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetOverLengthLimit(v string) *GetTitleDiagnoseResponseBodyData {
	s.OverLengthLimit = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetDisableWords(v string) *GetTitleDiagnoseResponseBodyData {
	s.DisableWords = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetNoFirstUppercaseList(v string) *GetTitleDiagnoseResponseBodyData {
	s.NoFirstUppercaseList = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetTotalScore(v string) *GetTitleDiagnoseResponseBodyData {
	s.TotalScore = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetWordSpelledCorrectError(v string) *GetTitleDiagnoseResponseBodyData {
	s.WordSpelledCorrectError = &v
	return s
}

type GetTitleDiagnoseResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTitleDiagnoseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTitleDiagnoseResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTitleDiagnoseResponse) GoString() string {
	return s.String()
}

func (s *GetTitleDiagnoseResponse) SetHeaders(v map[string]*string) *GetTitleDiagnoseResponse {
	s.Headers = v
	return s
}

func (s *GetTitleDiagnoseResponse) SetBody(v *GetTitleDiagnoseResponseBody) *GetTitleDiagnoseResponse {
	s.Body = v
	return s
}

type GetTitleGenerateRequest struct {
	Title      *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Platform   *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	HotWords   *string `json:"HotWords,omitempty" xml:"HotWords,omitempty"`
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	Extra      *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s GetTitleGenerateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTitleGenerateRequest) GoString() string {
	return s.String()
}

func (s *GetTitleGenerateRequest) SetTitle(v string) *GetTitleGenerateRequest {
	s.Title = &v
	return s
}

func (s *GetTitleGenerateRequest) SetLanguage(v string) *GetTitleGenerateRequest {
	s.Language = &v
	return s
}

func (s *GetTitleGenerateRequest) SetPlatform(v string) *GetTitleGenerateRequest {
	s.Platform = &v
	return s
}

func (s *GetTitleGenerateRequest) SetCategoryId(v string) *GetTitleGenerateRequest {
	s.CategoryId = &v
	return s
}

func (s *GetTitleGenerateRequest) SetHotWords(v string) *GetTitleGenerateRequest {
	s.HotWords = &v
	return s
}

func (s *GetTitleGenerateRequest) SetAttributes(v string) *GetTitleGenerateRequest {
	s.Attributes = &v
	return s
}

func (s *GetTitleGenerateRequest) SetExtra(v string) *GetTitleGenerateRequest {
	s.Extra = &v
	return s
}

type GetTitleGenerateResponseBody struct {
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetTitleGenerateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetTitleGenerateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTitleGenerateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTitleGenerateResponseBody) SetMessage(v string) *GetTitleGenerateResponseBody {
	s.Message = &v
	return s
}

func (s *GetTitleGenerateResponseBody) SetRequestId(v string) *GetTitleGenerateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTitleGenerateResponseBody) SetData(v *GetTitleGenerateResponseBodyData) *GetTitleGenerateResponseBody {
	s.Data = v
	return s
}

func (s *GetTitleGenerateResponseBody) SetCode(v int32) *GetTitleGenerateResponseBody {
	s.Code = &v
	return s
}

type GetTitleGenerateResponseBodyData struct {
	Titles *string `json:"Titles,omitempty" xml:"Titles,omitempty"`
}

func (s GetTitleGenerateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTitleGenerateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTitleGenerateResponseBodyData) SetTitles(v string) *GetTitleGenerateResponseBodyData {
	s.Titles = &v
	return s
}

type GetTitleGenerateResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTitleGenerateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTitleGenerateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTitleGenerateResponse) GoString() string {
	return s.String()
}

func (s *GetTitleGenerateResponse) SetHeaders(v map[string]*string) *GetTitleGenerateResponse {
	s.Headers = v
	return s
}

func (s *GetTitleGenerateResponse) SetBody(v *GetTitleGenerateResponseBody) *GetTitleGenerateResponse {
	s.Body = v
	return s
}

type GetTitleIntelligenceRequest struct {
	Platform        *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Extra           *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	CatLevelThreeId *int64  `json:"CatLevelThreeId,omitempty" xml:"CatLevelThreeId,omitempty"`
	CatLevelTwoId   *int64  `json:"CatLevelTwoId,omitempty" xml:"CatLevelTwoId,omitempty"`
	Keywords        *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
}

func (s GetTitleIntelligenceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTitleIntelligenceRequest) GoString() string {
	return s.String()
}

func (s *GetTitleIntelligenceRequest) SetPlatform(v string) *GetTitleIntelligenceRequest {
	s.Platform = &v
	return s
}

func (s *GetTitleIntelligenceRequest) SetExtra(v string) *GetTitleIntelligenceRequest {
	s.Extra = &v
	return s
}

func (s *GetTitleIntelligenceRequest) SetCatLevelThreeId(v int64) *GetTitleIntelligenceRequest {
	s.CatLevelThreeId = &v
	return s
}

func (s *GetTitleIntelligenceRequest) SetCatLevelTwoId(v int64) *GetTitleIntelligenceRequest {
	s.CatLevelTwoId = &v
	return s
}

func (s *GetTitleIntelligenceRequest) SetKeywords(v string) *GetTitleIntelligenceRequest {
	s.Keywords = &v
	return s
}

type GetTitleIntelligenceResponseBody struct {
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetTitleIntelligenceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetTitleIntelligenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTitleIntelligenceResponseBody) GoString() string {
	return s.String()
}

func (s *GetTitleIntelligenceResponseBody) SetMessage(v string) *GetTitleIntelligenceResponseBody {
	s.Message = &v
	return s
}

func (s *GetTitleIntelligenceResponseBody) SetRequestId(v string) *GetTitleIntelligenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTitleIntelligenceResponseBody) SetData(v *GetTitleIntelligenceResponseBodyData) *GetTitleIntelligenceResponseBody {
	s.Data = v
	return s
}

func (s *GetTitleIntelligenceResponseBody) SetCode(v int32) *GetTitleIntelligenceResponseBody {
	s.Code = &v
	return s
}

type GetTitleIntelligenceResponseBodyData struct {
	Titles *string `json:"Titles,omitempty" xml:"Titles,omitempty"`
}

func (s GetTitleIntelligenceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTitleIntelligenceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTitleIntelligenceResponseBodyData) SetTitles(v string) *GetTitleIntelligenceResponseBodyData {
	s.Titles = &v
	return s
}

type GetTitleIntelligenceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTitleIntelligenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTitleIntelligenceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTitleIntelligenceResponse) GoString() string {
	return s.String()
}

func (s *GetTitleIntelligenceResponse) SetHeaders(v map[string]*string) *GetTitleIntelligenceResponse {
	s.Headers = v
	return s
}

func (s *GetTitleIntelligenceResponse) SetBody(v *GetTitleIntelligenceResponseBody) *GetTitleIntelligenceResponse {
	s.Body = v
	return s
}

type OpenAlimtServiceRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OpenAlimtServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenAlimtServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenAlimtServiceRequest) SetOwnerId(v int64) *OpenAlimtServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenAlimtServiceRequest) SetType(v string) *OpenAlimtServiceRequest {
	s.Type = &v
	return s
}

type OpenAlimtServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s OpenAlimtServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenAlimtServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenAlimtServiceResponseBody) SetRequestId(v string) *OpenAlimtServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenAlimtServiceResponseBody) SetOrderId(v string) *OpenAlimtServiceResponseBody {
	s.OrderId = &v
	return s
}

type OpenAlimtServiceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenAlimtServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenAlimtServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenAlimtServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenAlimtServiceResponse) SetHeaders(v map[string]*string) *OpenAlimtServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenAlimtServiceResponse) SetBody(v *OpenAlimtServiceResponseBody) *OpenAlimtServiceResponse {
	s.Body = v
	return s
}

type TranslateRequest struct {
	FormatType     *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	SourceText     *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	Scene          *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s TranslateRequest) String() string {
	return tea.Prettify(s)
}

func (s TranslateRequest) GoString() string {
	return s.String()
}

func (s *TranslateRequest) SetFormatType(v string) *TranslateRequest {
	s.FormatType = &v
	return s
}

func (s *TranslateRequest) SetTargetLanguage(v string) *TranslateRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TranslateRequest) SetSourceLanguage(v string) *TranslateRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TranslateRequest) SetSourceText(v string) *TranslateRequest {
	s.SourceText = &v
	return s
}

func (s *TranslateRequest) SetScene(v string) *TranslateRequest {
	s.Scene = &v
	return s
}

type TranslateResponseBody struct {
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *TranslateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                     `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s TranslateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TranslateResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateResponseBody) SetMessage(v string) *TranslateResponseBody {
	s.Message = &v
	return s
}

func (s *TranslateResponseBody) SetRequestId(v string) *TranslateResponseBody {
	s.RequestId = &v
	return s
}

func (s *TranslateResponseBody) SetData(v *TranslateResponseBodyData) *TranslateResponseBody {
	s.Data = v
	return s
}

func (s *TranslateResponseBody) SetCode(v int32) *TranslateResponseBody {
	s.Code = &v
	return s
}

type TranslateResponseBodyData struct {
	WordCount  *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
	Translated *string `json:"Translated,omitempty" xml:"Translated,omitempty"`
}

func (s TranslateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TranslateResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateResponseBodyData) SetWordCount(v string) *TranslateResponseBodyData {
	s.WordCount = &v
	return s
}

func (s *TranslateResponseBodyData) SetTranslated(v string) *TranslateResponseBodyData {
	s.Translated = &v
	return s
}

type TranslateResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TranslateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TranslateResponse) String() string {
	return tea.Prettify(s)
}

func (s TranslateResponse) GoString() string {
	return s.String()
}

func (s *TranslateResponse) SetHeaders(v map[string]*string) *TranslateResponse {
	s.Headers = v
	return s
}

func (s *TranslateResponse) SetBody(v *TranslateResponseBody) *TranslateResponse {
	s.Body = v
	return s
}

type TranslateCertificateRequest struct {
	SourceLanguage  *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	TargetLanguage  *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	ImageUrl        *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	ResultType      *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
}

func (s TranslateCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s TranslateCertificateRequest) GoString() string {
	return s.String()
}

func (s *TranslateCertificateRequest) SetSourceLanguage(v string) *TranslateCertificateRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TranslateCertificateRequest) SetTargetLanguage(v string) *TranslateCertificateRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TranslateCertificateRequest) SetImageUrl(v string) *TranslateCertificateRequest {
	s.ImageUrl = &v
	return s
}

func (s *TranslateCertificateRequest) SetCertificateType(v string) *TranslateCertificateRequest {
	s.CertificateType = &v
	return s
}

func (s *TranslateCertificateRequest) SetResultType(v string) *TranslateCertificateRequest {
	s.ResultType = &v
	return s
}

type TranslateCertificateAdvanceRequest struct {
	ImageUrlObject  io.Reader `json:"ImageUrlObject,omitempty" xml:"ImageUrlObject,omitempty" require:"true"`
	SourceLanguage  *string   `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	TargetLanguage  *string   `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	CertificateType *string   `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	ResultType      *string   `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
}

func (s TranslateCertificateAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s TranslateCertificateAdvanceRequest) GoString() string {
	return s.String()
}

func (s *TranslateCertificateAdvanceRequest) SetImageUrlObject(v io.Reader) *TranslateCertificateAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *TranslateCertificateAdvanceRequest) SetSourceLanguage(v string) *TranslateCertificateAdvanceRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TranslateCertificateAdvanceRequest) SetTargetLanguage(v string) *TranslateCertificateAdvanceRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TranslateCertificateAdvanceRequest) SetCertificateType(v string) *TranslateCertificateAdvanceRequest {
	s.CertificateType = &v
	return s
}

func (s *TranslateCertificateAdvanceRequest) SetResultType(v string) *TranslateCertificateAdvanceRequest {
	s.ResultType = &v
	return s
}

type TranslateCertificateResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *TranslateCertificateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s TranslateCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TranslateCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateCertificateResponseBody) SetRequestId(v string) *TranslateCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *TranslateCertificateResponseBody) SetData(v *TranslateCertificateResponseBodyData) *TranslateCertificateResponseBody {
	s.Data = v
	return s
}

type TranslateCertificateResponseBodyData struct {
	TranslatedValues []*TranslateCertificateResponseBodyDataTranslatedValues `json:"TranslatedValues,omitempty" xml:"TranslatedValues,omitempty" type:"Repeated"`
}

func (s TranslateCertificateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TranslateCertificateResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateCertificateResponseBodyData) SetTranslatedValues(v []*TranslateCertificateResponseBodyDataTranslatedValues) *TranslateCertificateResponseBodyData {
	s.TranslatedValues = v
	return s
}

type TranslateCertificateResponseBodyDataTranslatedValues struct {
	Key              *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value            *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueTranslation *string `json:"ValueTranslation,omitempty" xml:"ValueTranslation,omitempty"`
	KeyTranslation   *string `json:"KeyTranslation,omitempty" xml:"KeyTranslation,omitempty"`
}

func (s TranslateCertificateResponseBodyDataTranslatedValues) String() string {
	return tea.Prettify(s)
}

func (s TranslateCertificateResponseBodyDataTranslatedValues) GoString() string {
	return s.String()
}

func (s *TranslateCertificateResponseBodyDataTranslatedValues) SetKey(v string) *TranslateCertificateResponseBodyDataTranslatedValues {
	s.Key = &v
	return s
}

func (s *TranslateCertificateResponseBodyDataTranslatedValues) SetValue(v string) *TranslateCertificateResponseBodyDataTranslatedValues {
	s.Value = &v
	return s
}

func (s *TranslateCertificateResponseBodyDataTranslatedValues) SetValueTranslation(v string) *TranslateCertificateResponseBodyDataTranslatedValues {
	s.ValueTranslation = &v
	return s
}

func (s *TranslateCertificateResponseBodyDataTranslatedValues) SetKeyTranslation(v string) *TranslateCertificateResponseBodyDataTranslatedValues {
	s.KeyTranslation = &v
	return s
}

type TranslateCertificateResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TranslateCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TranslateCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s TranslateCertificateResponse) GoString() string {
	return s.String()
}

func (s *TranslateCertificateResponse) SetHeaders(v map[string]*string) *TranslateCertificateResponse {
	s.Headers = v
	return s
}

func (s *TranslateCertificateResponse) SetBody(v *TranslateCertificateResponseBody) *TranslateCertificateResponse {
	s.Body = v
	return s
}

type TranslateECommerceRequest struct {
	FormatType     *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	SourceText     *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	Scene          *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s TranslateECommerceRequest) String() string {
	return tea.Prettify(s)
}

func (s TranslateECommerceRequest) GoString() string {
	return s.String()
}

func (s *TranslateECommerceRequest) SetFormatType(v string) *TranslateECommerceRequest {
	s.FormatType = &v
	return s
}

func (s *TranslateECommerceRequest) SetTargetLanguage(v string) *TranslateECommerceRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TranslateECommerceRequest) SetSourceLanguage(v string) *TranslateECommerceRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TranslateECommerceRequest) SetSourceText(v string) *TranslateECommerceRequest {
	s.SourceText = &v
	return s
}

func (s *TranslateECommerceRequest) SetScene(v string) *TranslateECommerceRequest {
	s.Scene = &v
	return s
}

type TranslateECommerceResponseBody struct {
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *TranslateECommerceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s TranslateECommerceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TranslateECommerceResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateECommerceResponseBody) SetMessage(v string) *TranslateECommerceResponseBody {
	s.Message = &v
	return s
}

func (s *TranslateECommerceResponseBody) SetRequestId(v string) *TranslateECommerceResponseBody {
	s.RequestId = &v
	return s
}

func (s *TranslateECommerceResponseBody) SetData(v *TranslateECommerceResponseBodyData) *TranslateECommerceResponseBody {
	s.Data = v
	return s
}

func (s *TranslateECommerceResponseBody) SetCode(v int32) *TranslateECommerceResponseBody {
	s.Code = &v
	return s
}

type TranslateECommerceResponseBodyData struct {
	WordCount  *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
	Translated *string `json:"Translated,omitempty" xml:"Translated,omitempty"`
}

func (s TranslateECommerceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TranslateECommerceResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateECommerceResponseBodyData) SetWordCount(v string) *TranslateECommerceResponseBodyData {
	s.WordCount = &v
	return s
}

func (s *TranslateECommerceResponseBodyData) SetTranslated(v string) *TranslateECommerceResponseBodyData {
	s.Translated = &v
	return s
}

type TranslateECommerceResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TranslateECommerceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TranslateECommerceResponse) String() string {
	return tea.Prettify(s)
}

func (s TranslateECommerceResponse) GoString() string {
	return s.String()
}

func (s *TranslateECommerceResponse) SetHeaders(v map[string]*string) *TranslateECommerceResponse {
	s.Headers = v
	return s
}

func (s *TranslateECommerceResponse) SetBody(v *TranslateECommerceResponseBody) *TranslateECommerceResponse {
	s.Body = v
	return s
}

type TranslateGeneralRequest struct {
	FormatType     *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	SourceText     *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	Scene          *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s TranslateGeneralRequest) String() string {
	return tea.Prettify(s)
}

func (s TranslateGeneralRequest) GoString() string {
	return s.String()
}

func (s *TranslateGeneralRequest) SetFormatType(v string) *TranslateGeneralRequest {
	s.FormatType = &v
	return s
}

func (s *TranslateGeneralRequest) SetSourceLanguage(v string) *TranslateGeneralRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TranslateGeneralRequest) SetTargetLanguage(v string) *TranslateGeneralRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TranslateGeneralRequest) SetSourceText(v string) *TranslateGeneralRequest {
	s.SourceText = &v
	return s
}

func (s *TranslateGeneralRequest) SetScene(v string) *TranslateGeneralRequest {
	s.Scene = &v
	return s
}

type TranslateGeneralResponseBody struct {
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *TranslateGeneralResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s TranslateGeneralResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TranslateGeneralResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateGeneralResponseBody) SetMessage(v string) *TranslateGeneralResponseBody {
	s.Message = &v
	return s
}

func (s *TranslateGeneralResponseBody) SetRequestId(v string) *TranslateGeneralResponseBody {
	s.RequestId = &v
	return s
}

func (s *TranslateGeneralResponseBody) SetData(v *TranslateGeneralResponseBodyData) *TranslateGeneralResponseBody {
	s.Data = v
	return s
}

func (s *TranslateGeneralResponseBody) SetCode(v int32) *TranslateGeneralResponseBody {
	s.Code = &v
	return s
}

type TranslateGeneralResponseBodyData struct {
	WordCount  *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
	Translated *string `json:"Translated,omitempty" xml:"Translated,omitempty"`
}

func (s TranslateGeneralResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TranslateGeneralResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateGeneralResponseBodyData) SetWordCount(v string) *TranslateGeneralResponseBodyData {
	s.WordCount = &v
	return s
}

func (s *TranslateGeneralResponseBodyData) SetTranslated(v string) *TranslateGeneralResponseBodyData {
	s.Translated = &v
	return s
}

type TranslateGeneralResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TranslateGeneralResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TranslateGeneralResponse) String() string {
	return tea.Prettify(s)
}

func (s TranslateGeneralResponse) GoString() string {
	return s.String()
}

func (s *TranslateGeneralResponse) SetHeaders(v map[string]*string) *TranslateGeneralResponse {
	s.Headers = v
	return s
}

func (s *TranslateGeneralResponse) SetBody(v *TranslateGeneralResponseBody) *TranslateGeneralResponse {
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
	client.EndpointMap = map[string]*string{
		"cn-hangzhou":                 tea.String("mt.cn-hangzhou.aliyuncs.com"),
		"ap-northeast-1":              tea.String("mt.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("mt.aliyuncs.com"),
		"ap-south-1":                  tea.String("mt.aliyuncs.com"),
		"ap-southeast-1":              tea.String("mt.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":              tea.String("mt.aliyuncs.com"),
		"ap-southeast-3":              tea.String("mt.aliyuncs.com"),
		"ap-southeast-5":              tea.String("mt.aliyuncs.com"),
		"cn-beijing":                  tea.String("mt.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("mt.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("mt.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("mt.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("mt.aliyuncs.com"),
		"cn-chengdu":                  tea.String("mt.aliyuncs.com"),
		"cn-edge-1":                   tea.String("mt.aliyuncs.com"),
		"cn-fujian":                   tea.String("mt.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("mt.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("mt.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("mt.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("mt.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("mt.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("mt.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("mt.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("mt.aliyuncs.com"),
		"cn-hongkong":                 tea.String("mt.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("mt.aliyuncs.com"),
		"cn-huhehaote":                tea.String("mt.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("mt.aliyuncs.com"),
		"cn-qingdao":                  tea.String("mt.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("mt.aliyuncs.com"),
		"cn-shanghai":                 tea.String("mt.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("mt.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("mt.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("mt.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("mt.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("mt.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("mt.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("mt.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("mt.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("mt.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("mt.aliyuncs.com"),
		"cn-wuhan":                    tea.String("mt.aliyuncs.com"),
		"cn-yushanfang":               tea.String("mt.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("mt.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("mt.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("mt.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("mt.aliyuncs.com"),
		"eu-central-1":                tea.String("mt.aliyuncs.com"),
		"eu-west-1":                   tea.String("mt.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("mt.aliyuncs.com"),
		"me-east-1":                   tea.String("mt.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("mt.aliyuncs.com"),
		"us-east-1":                   tea.String("mt.aliyuncs.com"),
		"us-west-1":                   tea.String("mt.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("alimt"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateDocTranslateTaskWithOptions(request *CreateDocTranslateTaskRequest, runtime *util.RuntimeOptions) (_result *CreateDocTranslateTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDocTranslateTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDocTranslateTask"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDocTranslateTask(request *CreateDocTranslateTaskRequest) (_result *CreateDocTranslateTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDocTranslateTaskResponse{}
	_body, _err := client.CreateDocTranslateTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDocTranslateTaskAdvance(request *CreateDocTranslateTaskAdvanceRequest, runtime *util.RuntimeOptions) (_result *CreateDocTranslateTaskResponse, _err error) {
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
		Product:  tea.String("alimt"),
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
	createDocTranslateTaskReq := &CreateDocTranslateTaskRequest{}
	openapiutil.Convert(request, createDocTranslateTaskReq)
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
		Content:     request.FileUrlObject,
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
	createDocTranslateTaskReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	createDocTranslateTaskResp, _err := client.CreateDocTranslateTaskWithOptions(createDocTranslateTaskReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = createDocTranslateTaskResp
	return _result, _err
}

func (client *Client) CreateImageTranslateTaskWithOptions(request *CreateImageTranslateTaskRequest, runtime *util.RuntimeOptions) (_result *CreateImageTranslateTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateImageTranslateTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateImageTranslateTask"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImageTranslateTask(request *CreateImageTranslateTaskRequest) (_result *CreateImageTranslateTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageTranslateTaskResponse{}
	_body, _err := client.CreateImageTranslateTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDetectLanguageWithOptions(request *GetDetectLanguageRequest, runtime *util.RuntimeOptions) (_result *GetDetectLanguageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDetectLanguageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDetectLanguage"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDetectLanguage(request *GetDetectLanguageRequest) (_result *GetDetectLanguageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDetectLanguageResponse{}
	_body, _err := client.GetDetectLanguageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDocTranslateTaskWithOptions(request *GetDocTranslateTaskRequest, runtime *util.RuntimeOptions) (_result *GetDocTranslateTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetDocTranslateTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDocTranslateTask"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDocTranslateTask(request *GetDocTranslateTaskRequest) (_result *GetDocTranslateTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDocTranslateTaskResponse{}
	_body, _err := client.GetDocTranslateTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetImageDiagnoseWithOptions(request *GetImageDiagnoseRequest, runtime *util.RuntimeOptions) (_result *GetImageDiagnoseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetImageDiagnoseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetImageDiagnose"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetImageDiagnose(request *GetImageDiagnoseRequest) (_result *GetImageDiagnoseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetImageDiagnoseResponse{}
	_body, _err := client.GetImageDiagnoseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetImageTranslateWithOptions(request *GetImageTranslateRequest, runtime *util.RuntimeOptions) (_result *GetImageTranslateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetImageTranslateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetImageTranslate"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetImageTranslate(request *GetImageTranslateRequest) (_result *GetImageTranslateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetImageTranslateResponse{}
	_body, _err := client.GetImageTranslateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetImageTranslateTaskWithOptions(request *GetImageTranslateTaskRequest, runtime *util.RuntimeOptions) (_result *GetImageTranslateTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetImageTranslateTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetImageTranslateTask"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetImageTranslateTask(request *GetImageTranslateTaskRequest) (_result *GetImageTranslateTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetImageTranslateTaskResponse{}
	_body, _err := client.GetImageTranslateTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTitleDiagnoseWithOptions(request *GetTitleDiagnoseRequest, runtime *util.RuntimeOptions) (_result *GetTitleDiagnoseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTitleDiagnoseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTitleDiagnose"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTitleDiagnose(request *GetTitleDiagnoseRequest) (_result *GetTitleDiagnoseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTitleDiagnoseResponse{}
	_body, _err := client.GetTitleDiagnoseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTitleGenerateWithOptions(request *GetTitleGenerateRequest, runtime *util.RuntimeOptions) (_result *GetTitleGenerateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTitleGenerateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTitleGenerate"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTitleGenerate(request *GetTitleGenerateRequest) (_result *GetTitleGenerateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTitleGenerateResponse{}
	_body, _err := client.GetTitleGenerateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTitleIntelligenceWithOptions(request *GetTitleIntelligenceRequest, runtime *util.RuntimeOptions) (_result *GetTitleIntelligenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTitleIntelligenceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTitleIntelligence"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTitleIntelligence(request *GetTitleIntelligenceRequest) (_result *GetTitleIntelligenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTitleIntelligenceResponse{}
	_body, _err := client.GetTitleIntelligenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenAlimtServiceWithOptions(request *OpenAlimtServiceRequest, runtime *util.RuntimeOptions) (_result *OpenAlimtServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OpenAlimtServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpenAlimtService"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenAlimtService(request *OpenAlimtServiceRequest) (_result *OpenAlimtServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenAlimtServiceResponse{}
	_body, _err := client.OpenAlimtServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TranslateWithOptions(request *TranslateRequest, runtime *util.RuntimeOptions) (_result *TranslateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TranslateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("Translate"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Translate(request *TranslateRequest) (_result *TranslateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TranslateResponse{}
	_body, _err := client.TranslateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TranslateCertificateWithOptions(request *TranslateCertificateRequest, runtime *util.RuntimeOptions) (_result *TranslateCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TranslateCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TranslateCertificate"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TranslateCertificate(request *TranslateCertificateRequest) (_result *TranslateCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TranslateCertificateResponse{}
	_body, _err := client.TranslateCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TranslateCertificateAdvance(request *TranslateCertificateAdvanceRequest, runtime *util.RuntimeOptions) (_result *TranslateCertificateResponse, _err error) {
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
		Product:  tea.String("alimt"),
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
	translateCertificateReq := &TranslateCertificateRequest{}
	openapiutil.Convert(request, translateCertificateReq)
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
		Content:     request.ImageUrlObject,
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
	translateCertificateReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	translateCertificateResp, _err := client.TranslateCertificateWithOptions(translateCertificateReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = translateCertificateResp
	return _result, _err
}

func (client *Client) TranslateECommerceWithOptions(request *TranslateECommerceRequest, runtime *util.RuntimeOptions) (_result *TranslateECommerceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TranslateECommerceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TranslateECommerce"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TranslateECommerce(request *TranslateECommerceRequest) (_result *TranslateECommerceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TranslateECommerceResponse{}
	_body, _err := client.TranslateECommerceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TranslateGeneralWithOptions(request *TranslateGeneralRequest, runtime *util.RuntimeOptions) (_result *TranslateGeneralResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TranslateGeneralResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TranslateGeneral"), tea.String("2018-10-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TranslateGeneral(request *TranslateGeneralRequest) (_result *TranslateGeneralResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TranslateGeneralResponse{}
	_body, _err := client.TranslateGeneralWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
