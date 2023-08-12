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

type CreateAsyncTranslateRequest struct {
	ApiType        *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	FormatType     *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	Scene          *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	SourceText     *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s CreateAsyncTranslateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAsyncTranslateRequest) GoString() string {
	return s.String()
}

func (s *CreateAsyncTranslateRequest) SetApiType(v string) *CreateAsyncTranslateRequest {
	s.ApiType = &v
	return s
}

func (s *CreateAsyncTranslateRequest) SetFormatType(v string) *CreateAsyncTranslateRequest {
	s.FormatType = &v
	return s
}

func (s *CreateAsyncTranslateRequest) SetScene(v string) *CreateAsyncTranslateRequest {
	s.Scene = &v
	return s
}

func (s *CreateAsyncTranslateRequest) SetSourceLanguage(v string) *CreateAsyncTranslateRequest {
	s.SourceLanguage = &v
	return s
}

func (s *CreateAsyncTranslateRequest) SetSourceText(v string) *CreateAsyncTranslateRequest {
	s.SourceText = &v
	return s
}

func (s *CreateAsyncTranslateRequest) SetTargetLanguage(v string) *CreateAsyncTranslateRequest {
	s.TargetLanguage = &v
	return s
}

type CreateAsyncTranslateResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateAsyncTranslateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAsyncTranslateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAsyncTranslateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAsyncTranslateResponseBody) SetCode(v int32) *CreateAsyncTranslateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAsyncTranslateResponseBody) SetData(v *CreateAsyncTranslateResponseBodyData) *CreateAsyncTranslateResponseBody {
	s.Data = v
	return s
}

func (s *CreateAsyncTranslateResponseBody) SetMessage(v string) *CreateAsyncTranslateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAsyncTranslateResponseBody) SetRequestId(v string) *CreateAsyncTranslateResponseBody {
	s.RequestId = &v
	return s
}

type CreateAsyncTranslateResponseBodyData struct {
	JobId  *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateAsyncTranslateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateAsyncTranslateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAsyncTranslateResponseBodyData) SetJobId(v string) *CreateAsyncTranslateResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreateAsyncTranslateResponseBodyData) SetStatus(v string) *CreateAsyncTranslateResponseBodyData {
	s.Status = &v
	return s
}

type CreateAsyncTranslateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAsyncTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAsyncTranslateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAsyncTranslateResponse) GoString() string {
	return s.String()
}

func (s *CreateAsyncTranslateResponse) SetHeaders(v map[string]*string) *CreateAsyncTranslateResponse {
	s.Headers = v
	return s
}

func (s *CreateAsyncTranslateResponse) SetStatusCode(v int32) *CreateAsyncTranslateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAsyncTranslateResponse) SetBody(v *CreateAsyncTranslateResponseBody) *CreateAsyncTranslateResponse {
	s.Body = v
	return s
}

type CreateDocTranslateTaskRequest struct {
	CallbackUrl    *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	FileUrl        *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Scene          *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s CreateDocTranslateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDocTranslateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDocTranslateTaskRequest) SetCallbackUrl(v string) *CreateDocTranslateTaskRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateDocTranslateTaskRequest) SetClientToken(v string) *CreateDocTranslateTaskRequest {
	s.ClientToken = &v
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

func (s *CreateDocTranslateTaskRequest) SetSourceLanguage(v string) *CreateDocTranslateTaskRequest {
	s.SourceLanguage = &v
	return s
}

func (s *CreateDocTranslateTaskRequest) SetTargetLanguage(v string) *CreateDocTranslateTaskRequest {
	s.TargetLanguage = &v
	return s
}

type CreateDocTranslateTaskAdvanceRequest struct {
	CallbackUrl    *string   `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	ClientToken    *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	FileUrlObject  io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Scene          *string   `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SourceLanguage *string   `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	TargetLanguage *string   `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s CreateDocTranslateTaskAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDocTranslateTaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDocTranslateTaskAdvanceRequest) SetCallbackUrl(v string) *CreateDocTranslateTaskAdvanceRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateDocTranslateTaskAdvanceRequest) SetClientToken(v string) *CreateDocTranslateTaskAdvanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDocTranslateTaskAdvanceRequest) SetFileUrlObject(v io.Reader) *CreateDocTranslateTaskAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *CreateDocTranslateTaskAdvanceRequest) SetScene(v string) *CreateDocTranslateTaskAdvanceRequest {
	s.Scene = &v
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

type CreateDocTranslateTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateDocTranslateTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDocTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDocTranslateTaskResponseBody) SetRequestId(v string) *CreateDocTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDocTranslateTaskResponseBody) SetStatus(v string) *CreateDocTranslateTaskResponseBody {
	s.Status = &v
	return s
}

func (s *CreateDocTranslateTaskResponseBody) SetTaskId(v string) *CreateDocTranslateTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateDocTranslateTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDocTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDocTranslateTaskResponse) SetStatusCode(v int32) *CreateDocTranslateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDocTranslateTaskResponse) SetBody(v *CreateDocTranslateTaskResponseBody) *CreateDocTranslateTaskResponse {
	s.Body = v
	return s
}

type CreateImageTranslateTaskRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Extra          *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	UrlList        *string `json:"UrlList,omitempty" xml:"UrlList,omitempty"`
}

func (s CreateImageTranslateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageTranslateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateImageTranslateTaskRequest) SetClientToken(v string) *CreateImageTranslateTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateImageTranslateTaskRequest) SetExtra(v string) *CreateImageTranslateTaskRequest {
	s.Extra = &v
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

func (s *CreateImageTranslateTaskRequest) SetUrlList(v string) *CreateImageTranslateTaskRequest {
	s.UrlList = &v
	return s
}

type CreateImageTranslateTaskResponseBody struct {
	Code      *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateImageTranslateTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImageTranslateTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageTranslateTaskResponseBody) SetCode(v int32) *CreateImageTranslateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateImageTranslateTaskResponseBody) SetData(v *CreateImageTranslateTaskResponseBodyData) *CreateImageTranslateTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateImageTranslateTaskResponseBody) SetMessage(v string) *CreateImageTranslateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateImageTranslateTaskResponseBody) SetRequestId(v string) *CreateImageTranslateTaskResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateImageTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateImageTranslateTaskResponse) SetStatusCode(v int32) *CreateImageTranslateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageTranslateTaskResponse) SetBody(v *CreateImageTranslateTaskResponseBody) *CreateImageTranslateTaskResponse {
	s.Body = v
	return s
}

type GetAsyncTranslateRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAsyncTranslateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncTranslateRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncTranslateRequest) SetJobId(v string) *GetAsyncTranslateRequest {
	s.JobId = &v
	return s
}

type GetAsyncTranslateResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetAsyncTranslateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAsyncTranslateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncTranslateResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncTranslateResponseBody) SetCode(v int32) *GetAsyncTranslateResponseBody {
	s.Code = &v
	return s
}

func (s *GetAsyncTranslateResponseBody) SetData(v *GetAsyncTranslateResponseBodyData) *GetAsyncTranslateResponseBody {
	s.Data = v
	return s
}

func (s *GetAsyncTranslateResponseBody) SetMessage(v string) *GetAsyncTranslateResponseBody {
	s.Message = &v
	return s
}

func (s *GetAsyncTranslateResponseBody) SetRequestId(v string) *GetAsyncTranslateResponseBody {
	s.RequestId = &v
	return s
}

type GetAsyncTranslateResponseBodyData struct {
	DetectedLanguage *string `json:"DetectedLanguage,omitempty" xml:"DetectedLanguage,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TranslatedText   *string `json:"TranslatedText,omitempty" xml:"TranslatedText,omitempty"`
	WordCount        *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s GetAsyncTranslateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncTranslateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsyncTranslateResponseBodyData) SetDetectedLanguage(v string) *GetAsyncTranslateResponseBodyData {
	s.DetectedLanguage = &v
	return s
}

func (s *GetAsyncTranslateResponseBodyData) SetStatus(v string) *GetAsyncTranslateResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAsyncTranslateResponseBodyData) SetTranslatedText(v string) *GetAsyncTranslateResponseBodyData {
	s.TranslatedText = &v
	return s
}

func (s *GetAsyncTranslateResponseBodyData) SetWordCount(v string) *GetAsyncTranslateResponseBodyData {
	s.WordCount = &v
	return s
}

type GetAsyncTranslateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAsyncTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAsyncTranslateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncTranslateResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncTranslateResponse) SetHeaders(v map[string]*string) *GetAsyncTranslateResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncTranslateResponse) SetStatusCode(v int32) *GetAsyncTranslateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncTranslateResponse) SetBody(v *GetAsyncTranslateResponseBody) *GetAsyncTranslateResponse {
	s.Body = v
	return s
}

type GetBatchTranslateRequest struct {
	ApiType        *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	FormatType     *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	Scene          *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	SourceText     *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s GetBatchTranslateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBatchTranslateRequest) GoString() string {
	return s.String()
}

func (s *GetBatchTranslateRequest) SetApiType(v string) *GetBatchTranslateRequest {
	s.ApiType = &v
	return s
}

func (s *GetBatchTranslateRequest) SetFormatType(v string) *GetBatchTranslateRequest {
	s.FormatType = &v
	return s
}

func (s *GetBatchTranslateRequest) SetScene(v string) *GetBatchTranslateRequest {
	s.Scene = &v
	return s
}

func (s *GetBatchTranslateRequest) SetSourceLanguage(v string) *GetBatchTranslateRequest {
	s.SourceLanguage = &v
	return s
}

func (s *GetBatchTranslateRequest) SetSourceText(v string) *GetBatchTranslateRequest {
	s.SourceText = &v
	return s
}

func (s *GetBatchTranslateRequest) SetTargetLanguage(v string) *GetBatchTranslateRequest {
	s.TargetLanguage = &v
	return s
}

type GetBatchTranslateResponseBody struct {
	Code           *int32                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TranslatedList []map[string]interface{} `json:"TranslatedList,omitempty" xml:"TranslatedList,omitempty" type:"Repeated"`
}

func (s GetBatchTranslateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBatchTranslateResponseBody) GoString() string {
	return s.String()
}

func (s *GetBatchTranslateResponseBody) SetCode(v int32) *GetBatchTranslateResponseBody {
	s.Code = &v
	return s
}

func (s *GetBatchTranslateResponseBody) SetMessage(v string) *GetBatchTranslateResponseBody {
	s.Message = &v
	return s
}

func (s *GetBatchTranslateResponseBody) SetRequestId(v string) *GetBatchTranslateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBatchTranslateResponseBody) SetTranslatedList(v []map[string]interface{}) *GetBatchTranslateResponseBody {
	s.TranslatedList = v
	return s
}

type GetBatchTranslateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBatchTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBatchTranslateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBatchTranslateResponse) GoString() string {
	return s.String()
}

func (s *GetBatchTranslateResponse) SetHeaders(v map[string]*string) *GetBatchTranslateResponse {
	s.Headers = v
	return s
}

func (s *GetBatchTranslateResponse) SetStatusCode(v int32) *GetBatchTranslateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBatchTranslateResponse) SetBody(v *GetBatchTranslateResponseBody) *GetBatchTranslateResponse {
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
	DetectedLanguage *string `json:"DetectedLanguage,omitempty" xml:"DetectedLanguage,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDetectLanguageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDetectLanguageResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetectLanguageResponseBody) SetDetectedLanguage(v string) *GetDetectLanguageResponseBody {
	s.DetectedLanguage = &v
	return s
}

func (s *GetDetectLanguageResponseBody) SetRequestId(v string) *GetDetectLanguageResponseBody {
	s.RequestId = &v
	return s
}

type GetDetectLanguageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDetectLanguageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDetectLanguageResponse) SetStatusCode(v int32) *GetDetectLanguageResponse {
	s.StatusCode = &v
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
	PageCount             *int32  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId                *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TranslateErrorCode    *string `json:"TranslateErrorCode,omitempty" xml:"TranslateErrorCode,omitempty"`
	TranslateErrorMessage *string `json:"TranslateErrorMessage,omitempty" xml:"TranslateErrorMessage,omitempty"`
	TranslateFileUrl      *string `json:"TranslateFileUrl,omitempty" xml:"TranslateFileUrl,omitempty"`
}

func (s GetDocTranslateTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocTranslateTaskResponseBody) SetPageCount(v int32) *GetDocTranslateTaskResponseBody {
	s.PageCount = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetRequestId(v string) *GetDocTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetStatus(v string) *GetDocTranslateTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetTaskId(v string) *GetDocTranslateTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetTranslateErrorCode(v string) *GetDocTranslateTaskResponseBody {
	s.TranslateErrorCode = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetTranslateErrorMessage(v string) *GetDocTranslateTaskResponseBody {
	s.TranslateErrorMessage = &v
	return s
}

func (s *GetDocTranslateTaskResponseBody) SetTranslateFileUrl(v string) *GetDocTranslateTaskResponseBody {
	s.TranslateFileUrl = &v
	return s
}

type GetDocTranslateTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDocTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDocTranslateTaskResponse) SetStatusCode(v int32) *GetDocTranslateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocTranslateTaskResponse) SetBody(v *GetDocTranslateTaskResponseBody) *GetDocTranslateTaskResponse {
	s.Body = v
	return s
}

type GetImageDiagnoseRequest struct {
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	Url   *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetImageDiagnoseRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageDiagnoseRequest) GoString() string {
	return s.String()
}

func (s *GetImageDiagnoseRequest) SetExtra(v string) *GetImageDiagnoseRequest {
	s.Extra = &v
	return s
}

func (s *GetImageDiagnoseRequest) SetUrl(v string) *GetImageDiagnoseRequest {
	s.Url = &v
	return s
}

type GetImageDiagnoseResponseBody struct {
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetImageDiagnoseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetImageDiagnoseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageDiagnoseResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageDiagnoseResponseBody) SetCode(v int32) *GetImageDiagnoseResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageDiagnoseResponseBody) SetData(v *GetImageDiagnoseResponseBodyData) *GetImageDiagnoseResponseBody {
	s.Data = v
	return s
}

func (s *GetImageDiagnoseResponseBody) SetMessage(v string) *GetImageDiagnoseResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageDiagnoseResponseBody) SetRequestId(v string) *GetImageDiagnoseResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetImageDiagnoseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetImageDiagnoseResponse) SetStatusCode(v int32) *GetImageDiagnoseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageDiagnoseResponse) SetBody(v *GetImageDiagnoseResponseBody) *GetImageDiagnoseResponse {
	s.Body = v
	return s
}

type GetImageTranslateRequest struct {
	Extra          *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	Url            *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetImageTranslateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageTranslateRequest) GoString() string {
	return s.String()
}

func (s *GetImageTranslateRequest) SetExtra(v string) *GetImageTranslateRequest {
	s.Extra = &v
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

func (s *GetImageTranslateRequest) SetUrl(v string) *GetImageTranslateRequest {
	s.Url = &v
	return s
}

type GetImageTranslateResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetImageTranslateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetImageTranslateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageTranslateResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageTranslateResponseBody) SetCode(v int32) *GetImageTranslateResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageTranslateResponseBody) SetData(v *GetImageTranslateResponseBodyData) *GetImageTranslateResponseBody {
	s.Data = v
	return s
}

func (s *GetImageTranslateResponseBody) SetMessage(v string) *GetImageTranslateResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageTranslateResponseBody) SetRequestId(v string) *GetImageTranslateResponseBody {
	s.RequestId = &v
	return s
}

type GetImageTranslateResponseBodyData struct {
	Orc           *string `json:"Orc,omitempty" xml:"Orc,omitempty"`
	PictureEditor *string `json:"PictureEditor,omitempty" xml:"PictureEditor,omitempty"`
	Url           *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetImageTranslateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetImageTranslateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetImageTranslateResponseBodyData) SetOrc(v string) *GetImageTranslateResponseBodyData {
	s.Orc = &v
	return s
}

func (s *GetImageTranslateResponseBodyData) SetPictureEditor(v string) *GetImageTranslateResponseBodyData {
	s.PictureEditor = &v
	return s
}

func (s *GetImageTranslateResponseBodyData) SetUrl(v string) *GetImageTranslateResponseBodyData {
	s.Url = &v
	return s
}

type GetImageTranslateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetImageTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetImageTranslateResponse) SetStatusCode(v int32) *GetImageTranslateResponse {
	s.StatusCode = &v
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
	Code      *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetImageTranslateTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetImageTranslateTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskResponseBody) SetCode(v int32) *GetImageTranslateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageTranslateTaskResponseBody) SetData(v *GetImageTranslateTaskResponseBodyData) *GetImageTranslateTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetImageTranslateTaskResponseBody) SetMessage(v string) *GetImageTranslateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageTranslateTaskResponseBody) SetRequestId(v string) *GetImageTranslateTaskResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetImageTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetImageTranslateTaskResponse) SetStatusCode(v int32) *GetImageTranslateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageTranslateTaskResponse) SetBody(v *GetImageTranslateTaskResponseBody) *GetImageTranslateTaskResponse {
	s.Body = v
	return s
}

type GetTitleDiagnoseRequest struct {
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Extra      *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Platform   *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Title      *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetTitleDiagnoseRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTitleDiagnoseRequest) GoString() string {
	return s.String()
}

func (s *GetTitleDiagnoseRequest) SetCategoryId(v string) *GetTitleDiagnoseRequest {
	s.CategoryId = &v
	return s
}

func (s *GetTitleDiagnoseRequest) SetExtra(v string) *GetTitleDiagnoseRequest {
	s.Extra = &v
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

func (s *GetTitleDiagnoseRequest) SetTitle(v string) *GetTitleDiagnoseRequest {
	s.Title = &v
	return s
}

type GetTitleDiagnoseResponseBody struct {
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetTitleDiagnoseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTitleDiagnoseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTitleDiagnoseResponseBody) GoString() string {
	return s.String()
}

func (s *GetTitleDiagnoseResponseBody) SetCode(v int32) *GetTitleDiagnoseResponseBody {
	s.Code = &v
	return s
}

func (s *GetTitleDiagnoseResponseBody) SetData(v *GetTitleDiagnoseResponseBodyData) *GetTitleDiagnoseResponseBody {
	s.Data = v
	return s
}

func (s *GetTitleDiagnoseResponseBody) SetMessage(v string) *GetTitleDiagnoseResponseBody {
	s.Message = &v
	return s
}

func (s *GetTitleDiagnoseResponseBody) SetRequestId(v string) *GetTitleDiagnoseResponseBody {
	s.RequestId = &v
	return s
}

type GetTitleDiagnoseResponseBodyData struct {
	AllUppercaseWords       *string `json:"AllUppercaseWords,omitempty" xml:"AllUppercaseWords,omitempty"`
	ContainCoreClasses      *string `json:"ContainCoreClasses,omitempty" xml:"ContainCoreClasses,omitempty"`
	DisableWords            *string `json:"DisableWords,omitempty" xml:"DisableWords,omitempty"`
	DuplicateWords          *string `json:"DuplicateWords,omitempty" xml:"DuplicateWords,omitempty"`
	LanguageQualityScore    *string `json:"LanguageQualityScore,omitempty" xml:"LanguageQualityScore,omitempty"`
	NoFirstUppercaseList    *string `json:"NoFirstUppercaseList,omitempty" xml:"NoFirstUppercaseList,omitempty"`
	OverLengthLimit         *string `json:"OverLengthLimit,omitempty" xml:"OverLengthLimit,omitempty"`
	TotalScore              *string `json:"TotalScore,omitempty" xml:"TotalScore,omitempty"`
	WordCount               *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
	WordSpelledCorrectError *string `json:"WordSpelledCorrectError,omitempty" xml:"WordSpelledCorrectError,omitempty"`
}

func (s GetTitleDiagnoseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTitleDiagnoseResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTitleDiagnoseResponseBodyData) SetAllUppercaseWords(v string) *GetTitleDiagnoseResponseBodyData {
	s.AllUppercaseWords = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetContainCoreClasses(v string) *GetTitleDiagnoseResponseBodyData {
	s.ContainCoreClasses = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetDisableWords(v string) *GetTitleDiagnoseResponseBodyData {
	s.DisableWords = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetDuplicateWords(v string) *GetTitleDiagnoseResponseBodyData {
	s.DuplicateWords = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetLanguageQualityScore(v string) *GetTitleDiagnoseResponseBodyData {
	s.LanguageQualityScore = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetNoFirstUppercaseList(v string) *GetTitleDiagnoseResponseBodyData {
	s.NoFirstUppercaseList = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetOverLengthLimit(v string) *GetTitleDiagnoseResponseBodyData {
	s.OverLengthLimit = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetTotalScore(v string) *GetTitleDiagnoseResponseBodyData {
	s.TotalScore = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetWordCount(v string) *GetTitleDiagnoseResponseBodyData {
	s.WordCount = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetWordSpelledCorrectError(v string) *GetTitleDiagnoseResponseBodyData {
	s.WordSpelledCorrectError = &v
	return s
}

type GetTitleDiagnoseResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTitleDiagnoseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTitleDiagnoseResponse) SetStatusCode(v int32) *GetTitleDiagnoseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTitleDiagnoseResponse) SetBody(v *GetTitleDiagnoseResponseBody) *GetTitleDiagnoseResponse {
	s.Body = v
	return s
}

type GetTitleGenerateRequest struct {
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Extra      *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	HotWords   *string `json:"HotWords,omitempty" xml:"HotWords,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Platform   *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Title      *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetTitleGenerateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTitleGenerateRequest) GoString() string {
	return s.String()
}

func (s *GetTitleGenerateRequest) SetAttributes(v string) *GetTitleGenerateRequest {
	s.Attributes = &v
	return s
}

func (s *GetTitleGenerateRequest) SetCategoryId(v string) *GetTitleGenerateRequest {
	s.CategoryId = &v
	return s
}

func (s *GetTitleGenerateRequest) SetExtra(v string) *GetTitleGenerateRequest {
	s.Extra = &v
	return s
}

func (s *GetTitleGenerateRequest) SetHotWords(v string) *GetTitleGenerateRequest {
	s.HotWords = &v
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

func (s *GetTitleGenerateRequest) SetTitle(v string) *GetTitleGenerateRequest {
	s.Title = &v
	return s
}

type GetTitleGenerateResponseBody struct {
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetTitleGenerateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTitleGenerateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTitleGenerateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTitleGenerateResponseBody) SetCode(v int32) *GetTitleGenerateResponseBody {
	s.Code = &v
	return s
}

func (s *GetTitleGenerateResponseBody) SetData(v *GetTitleGenerateResponseBodyData) *GetTitleGenerateResponseBody {
	s.Data = v
	return s
}

func (s *GetTitleGenerateResponseBody) SetMessage(v string) *GetTitleGenerateResponseBody {
	s.Message = &v
	return s
}

func (s *GetTitleGenerateResponseBody) SetRequestId(v string) *GetTitleGenerateResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTitleGenerateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTitleGenerateResponse) SetStatusCode(v int32) *GetTitleGenerateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTitleGenerateResponse) SetBody(v *GetTitleGenerateResponseBody) *GetTitleGenerateResponse {
	s.Body = v
	return s
}

type GetTitleIntelligenceRequest struct {
	CatLevelThreeId *int64  `json:"CatLevelThreeId,omitempty" xml:"CatLevelThreeId,omitempty"`
	CatLevelTwoId   *int64  `json:"CatLevelTwoId,omitempty" xml:"CatLevelTwoId,omitempty"`
	Extra           *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	Keywords        *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	Platform        *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s GetTitleIntelligenceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTitleIntelligenceRequest) GoString() string {
	return s.String()
}

func (s *GetTitleIntelligenceRequest) SetCatLevelThreeId(v int64) *GetTitleIntelligenceRequest {
	s.CatLevelThreeId = &v
	return s
}

func (s *GetTitleIntelligenceRequest) SetCatLevelTwoId(v int64) *GetTitleIntelligenceRequest {
	s.CatLevelTwoId = &v
	return s
}

func (s *GetTitleIntelligenceRequest) SetExtra(v string) *GetTitleIntelligenceRequest {
	s.Extra = &v
	return s
}

func (s *GetTitleIntelligenceRequest) SetKeywords(v string) *GetTitleIntelligenceRequest {
	s.Keywords = &v
	return s
}

func (s *GetTitleIntelligenceRequest) SetPlatform(v string) *GetTitleIntelligenceRequest {
	s.Platform = &v
	return s
}

type GetTitleIntelligenceResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetTitleIntelligenceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTitleIntelligenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTitleIntelligenceResponseBody) GoString() string {
	return s.String()
}

func (s *GetTitleIntelligenceResponseBody) SetCode(v int32) *GetTitleIntelligenceResponseBody {
	s.Code = &v
	return s
}

func (s *GetTitleIntelligenceResponseBody) SetData(v *GetTitleIntelligenceResponseBodyData) *GetTitleIntelligenceResponseBody {
	s.Data = v
	return s
}

func (s *GetTitleIntelligenceResponseBody) SetMessage(v string) *GetTitleIntelligenceResponseBody {
	s.Message = &v
	return s
}

func (s *GetTitleIntelligenceResponseBody) SetRequestId(v string) *GetTitleIntelligenceResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTitleIntelligenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTitleIntelligenceResponse) SetStatusCode(v int32) *GetTitleIntelligenceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTitleIntelligenceResponse) SetBody(v *GetTitleIntelligenceResponseBody) *GetTitleIntelligenceResponse {
	s.Body = v
	return s
}

type GetTranslateImageBatchResultRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTranslateImageBatchResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTranslateImageBatchResultRequest) GoString() string {
	return s.String()
}

func (s *GetTranslateImageBatchResultRequest) SetTaskId(v string) *GetTranslateImageBatchResultRequest {
	s.TaskId = &v
	return s
}

type GetTranslateImageBatchResultResponseBody struct {
	Code      *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetTranslateImageBatchResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTranslateImageBatchResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTranslateImageBatchResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranslateImageBatchResultResponseBody) SetCode(v int32) *GetTranslateImageBatchResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetTranslateImageBatchResultResponseBody) SetData(v *GetTranslateImageBatchResultResponseBodyData) *GetTranslateImageBatchResultResponseBody {
	s.Data = v
	return s
}

func (s *GetTranslateImageBatchResultResponseBody) SetMessage(v string) *GetTranslateImageBatchResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetTranslateImageBatchResultResponseBody) SetRequestId(v string) *GetTranslateImageBatchResultResponseBody {
	s.RequestId = &v
	return s
}

type GetTranslateImageBatchResultResponseBodyData struct {
	Result []*GetTranslateImageBatchResultResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Status *string                                               `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTranslateImageBatchResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTranslateImageBatchResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTranslateImageBatchResultResponseBodyData) SetResult(v []*GetTranslateImageBatchResultResponseBodyDataResult) *GetTranslateImageBatchResultResponseBodyData {
	s.Result = v
	return s
}

func (s *GetTranslateImageBatchResultResponseBodyData) SetStatus(v string) *GetTranslateImageBatchResultResponseBodyData {
	s.Status = &v
	return s
}

type GetTranslateImageBatchResultResponseBodyDataResult struct {
	Code           *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	FinalImageUrl  *string `json:"FinalImageUrl,omitempty" xml:"FinalImageUrl,omitempty"`
	InPaintingUrl  *string `json:"InPaintingUrl,omitempty" xml:"InPaintingUrl,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	SourceImageUrl *string `json:"SourceImageUrl,omitempty" xml:"SourceImageUrl,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TemplateJson   *string `json:"TemplateJson,omitempty" xml:"TemplateJson,omitempty"`
}

func (s GetTranslateImageBatchResultResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s GetTranslateImageBatchResultResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) SetCode(v int32) *GetTranslateImageBatchResultResponseBodyDataResult {
	s.Code = &v
	return s
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) SetFinalImageUrl(v string) *GetTranslateImageBatchResultResponseBodyDataResult {
	s.FinalImageUrl = &v
	return s
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) SetInPaintingUrl(v string) *GetTranslateImageBatchResultResponseBodyDataResult {
	s.InPaintingUrl = &v
	return s
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) SetMessage(v string) *GetTranslateImageBatchResultResponseBodyDataResult {
	s.Message = &v
	return s
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) SetSourceImageUrl(v string) *GetTranslateImageBatchResultResponseBodyDataResult {
	s.SourceImageUrl = &v
	return s
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) SetSuccess(v bool) *GetTranslateImageBatchResultResponseBodyDataResult {
	s.Success = &v
	return s
}

func (s *GetTranslateImageBatchResultResponseBodyDataResult) SetTemplateJson(v string) *GetTranslateImageBatchResultResponseBodyDataResult {
	s.TemplateJson = &v
	return s
}

type GetTranslateImageBatchResultResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTranslateImageBatchResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTranslateImageBatchResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTranslateImageBatchResultResponse) GoString() string {
	return s.String()
}

func (s *GetTranslateImageBatchResultResponse) SetHeaders(v map[string]*string) *GetTranslateImageBatchResultResponse {
	s.Headers = v
	return s
}

func (s *GetTranslateImageBatchResultResponse) SetStatusCode(v int32) *GetTranslateImageBatchResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTranslateImageBatchResultResponse) SetBody(v *GetTranslateImageBatchResultResponseBody) *GetTranslateImageBatchResultResponse {
	s.Body = v
	return s
}

type GetTranslateReportRequest struct {
	ApiName   *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Group     *string `json:"Group,omitempty" xml:"Group,omitempty"`
}

func (s GetTranslateReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTranslateReportRequest) GoString() string {
	return s.String()
}

func (s *GetTranslateReportRequest) SetApiName(v string) *GetTranslateReportRequest {
	s.ApiName = &v
	return s
}

func (s *GetTranslateReportRequest) SetBeginTime(v string) *GetTranslateReportRequest {
	s.BeginTime = &v
	return s
}

func (s *GetTranslateReportRequest) SetEndTime(v string) *GetTranslateReportRequest {
	s.EndTime = &v
	return s
}

func (s *GetTranslateReportRequest) SetGroup(v string) *GetTranslateReportRequest {
	s.Group = &v
	return s
}

type GetTranslateReportResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTranslateReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTranslateReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranslateReportResponseBody) SetCode(v int32) *GetTranslateReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetTranslateReportResponseBody) SetData(v string) *GetTranslateReportResponseBody {
	s.Data = &v
	return s
}

func (s *GetTranslateReportResponseBody) SetMessage(v string) *GetTranslateReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetTranslateReportResponseBody) SetRequestId(v string) *GetTranslateReportResponseBody {
	s.RequestId = &v
	return s
}

type GetTranslateReportResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTranslateReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTranslateReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTranslateReportResponse) GoString() string {
	return s.String()
}

func (s *GetTranslateReportResponse) SetHeaders(v map[string]*string) *GetTranslateReportResponse {
	s.Headers = v
	return s
}

func (s *GetTranslateReportResponse) SetStatusCode(v int32) *GetTranslateReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTranslateReportResponse) SetBody(v *GetTranslateReportResponseBody) *GetTranslateReportResponse {
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
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenAlimtServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenAlimtServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenAlimtServiceResponseBody) SetOrderId(v string) *OpenAlimtServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenAlimtServiceResponseBody) SetRequestId(v string) *OpenAlimtServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenAlimtServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenAlimtServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OpenAlimtServiceResponse) SetStatusCode(v int32) *OpenAlimtServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenAlimtServiceResponse) SetBody(v *OpenAlimtServiceResponseBody) *OpenAlimtServiceResponse {
	s.Body = v
	return s
}

type TranslateRequest struct {
	Context        *string `json:"Context,omitempty" xml:"Context,omitempty"`
	FormatType     *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	Scene          *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	SourceText     *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s TranslateRequest) String() string {
	return tea.Prettify(s)
}

func (s TranslateRequest) GoString() string {
	return s.String()
}

func (s *TranslateRequest) SetContext(v string) *TranslateRequest {
	s.Context = &v
	return s
}

func (s *TranslateRequest) SetFormatType(v string) *TranslateRequest {
	s.FormatType = &v
	return s
}

func (s *TranslateRequest) SetScene(v string) *TranslateRequest {
	s.Scene = &v
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

func (s *TranslateRequest) SetTargetLanguage(v string) *TranslateRequest {
	s.TargetLanguage = &v
	return s
}

type TranslateResponseBody struct {
	Code      *int32                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *TranslateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TranslateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TranslateResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateResponseBody) SetCode(v int32) *TranslateResponseBody {
	s.Code = &v
	return s
}

func (s *TranslateResponseBody) SetData(v *TranslateResponseBodyData) *TranslateResponseBody {
	s.Data = v
	return s
}

func (s *TranslateResponseBody) SetMessage(v string) *TranslateResponseBody {
	s.Message = &v
	return s
}

func (s *TranslateResponseBody) SetRequestId(v string) *TranslateResponseBody {
	s.RequestId = &v
	return s
}

type TranslateResponseBodyData struct {
	DetectedLanguage *string `json:"DetectedLanguage,omitempty" xml:"DetectedLanguage,omitempty"`
	Translated       *string `json:"Translated,omitempty" xml:"Translated,omitempty"`
	WordCount        *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s TranslateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TranslateResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateResponseBodyData) SetDetectedLanguage(v string) *TranslateResponseBodyData {
	s.DetectedLanguage = &v
	return s
}

func (s *TranslateResponseBodyData) SetTranslated(v string) *TranslateResponseBodyData {
	s.Translated = &v
	return s
}

func (s *TranslateResponseBodyData) SetWordCount(v string) *TranslateResponseBodyData {
	s.WordCount = &v
	return s
}

type TranslateResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TranslateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TranslateResponse) SetStatusCode(v int32) *TranslateResponse {
	s.StatusCode = &v
	return s
}

func (s *TranslateResponse) SetBody(v *TranslateResponseBody) *TranslateResponse {
	s.Body = v
	return s
}

type TranslateCertificateRequest struct {
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	ImageUrl        *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	ResultType      *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	SourceLanguage  *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	TargetLanguage  *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s TranslateCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s TranslateCertificateRequest) GoString() string {
	return s.String()
}

func (s *TranslateCertificateRequest) SetCertificateType(v string) *TranslateCertificateRequest {
	s.CertificateType = &v
	return s
}

func (s *TranslateCertificateRequest) SetImageUrl(v string) *TranslateCertificateRequest {
	s.ImageUrl = &v
	return s
}

func (s *TranslateCertificateRequest) SetResultType(v string) *TranslateCertificateRequest {
	s.ResultType = &v
	return s
}

func (s *TranslateCertificateRequest) SetSourceLanguage(v string) *TranslateCertificateRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TranslateCertificateRequest) SetTargetLanguage(v string) *TranslateCertificateRequest {
	s.TargetLanguage = &v
	return s
}

type TranslateCertificateAdvanceRequest struct {
	CertificateType *string   `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	ImageUrlObject  io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	ResultType      *string   `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	SourceLanguage  *string   `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	TargetLanguage  *string   `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s TranslateCertificateAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s TranslateCertificateAdvanceRequest) GoString() string {
	return s.String()
}

func (s *TranslateCertificateAdvanceRequest) SetCertificateType(v string) *TranslateCertificateAdvanceRequest {
	s.CertificateType = &v
	return s
}

func (s *TranslateCertificateAdvanceRequest) SetImageUrlObject(v io.Reader) *TranslateCertificateAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *TranslateCertificateAdvanceRequest) SetResultType(v string) *TranslateCertificateAdvanceRequest {
	s.ResultType = &v
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

type TranslateCertificateResponseBody struct {
	Data      *TranslateCertificateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TranslateCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TranslateCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateCertificateResponseBody) SetData(v *TranslateCertificateResponseBodyData) *TranslateCertificateResponseBody {
	s.Data = v
	return s
}

func (s *TranslateCertificateResponseBody) SetRequestId(v string) *TranslateCertificateResponseBody {
	s.RequestId = &v
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
	KeyTranslation   *string `json:"KeyTranslation,omitempty" xml:"KeyTranslation,omitempty"`
	Value            *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueTranslation *string `json:"ValueTranslation,omitempty" xml:"ValueTranslation,omitempty"`
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

func (s *TranslateCertificateResponseBodyDataTranslatedValues) SetKeyTranslation(v string) *TranslateCertificateResponseBodyDataTranslatedValues {
	s.KeyTranslation = &v
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

type TranslateCertificateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TranslateCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TranslateCertificateResponse) SetStatusCode(v int32) *TranslateCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *TranslateCertificateResponse) SetBody(v *TranslateCertificateResponseBody) *TranslateCertificateResponse {
	s.Body = v
	return s
}

type TranslateECommerceRequest struct {
	Context        *string `json:"Context,omitempty" xml:"Context,omitempty"`
	FormatType     *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	Scene          *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	SourceText     *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s TranslateECommerceRequest) String() string {
	return tea.Prettify(s)
}

func (s TranslateECommerceRequest) GoString() string {
	return s.String()
}

func (s *TranslateECommerceRequest) SetContext(v string) *TranslateECommerceRequest {
	s.Context = &v
	return s
}

func (s *TranslateECommerceRequest) SetFormatType(v string) *TranslateECommerceRequest {
	s.FormatType = &v
	return s
}

func (s *TranslateECommerceRequest) SetScene(v string) *TranslateECommerceRequest {
	s.Scene = &v
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

func (s *TranslateECommerceRequest) SetTargetLanguage(v string) *TranslateECommerceRequest {
	s.TargetLanguage = &v
	return s
}

type TranslateECommerceResponseBody struct {
	Code      *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *TranslateECommerceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TranslateECommerceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TranslateECommerceResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateECommerceResponseBody) SetCode(v int32) *TranslateECommerceResponseBody {
	s.Code = &v
	return s
}

func (s *TranslateECommerceResponseBody) SetData(v *TranslateECommerceResponseBodyData) *TranslateECommerceResponseBody {
	s.Data = v
	return s
}

func (s *TranslateECommerceResponseBody) SetMessage(v string) *TranslateECommerceResponseBody {
	s.Message = &v
	return s
}

func (s *TranslateECommerceResponseBody) SetRequestId(v string) *TranslateECommerceResponseBody {
	s.RequestId = &v
	return s
}

type TranslateECommerceResponseBodyData struct {
	Translated *string `json:"Translated,omitempty" xml:"Translated,omitempty"`
	WordCount  *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s TranslateECommerceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TranslateECommerceResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateECommerceResponseBodyData) SetTranslated(v string) *TranslateECommerceResponseBodyData {
	s.Translated = &v
	return s
}

func (s *TranslateECommerceResponseBodyData) SetWordCount(v string) *TranslateECommerceResponseBodyData {
	s.WordCount = &v
	return s
}

type TranslateECommerceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TranslateECommerceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TranslateECommerceResponse) SetStatusCode(v int32) *TranslateECommerceResponse {
	s.StatusCode = &v
	return s
}

func (s *TranslateECommerceResponse) SetBody(v *TranslateECommerceResponseBody) *TranslateECommerceResponse {
	s.Body = v
	return s
}

type TranslateGeneralRequest struct {
	Context        *string `json:"Context,omitempty" xml:"Context,omitempty"`
	FormatType     *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	Scene          *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	SourceText     *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s TranslateGeneralRequest) String() string {
	return tea.Prettify(s)
}

func (s TranslateGeneralRequest) GoString() string {
	return s.String()
}

func (s *TranslateGeneralRequest) SetContext(v string) *TranslateGeneralRequest {
	s.Context = &v
	return s
}

func (s *TranslateGeneralRequest) SetFormatType(v string) *TranslateGeneralRequest {
	s.FormatType = &v
	return s
}

func (s *TranslateGeneralRequest) SetScene(v string) *TranslateGeneralRequest {
	s.Scene = &v
	return s
}

func (s *TranslateGeneralRequest) SetSourceLanguage(v string) *TranslateGeneralRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TranslateGeneralRequest) SetSourceText(v string) *TranslateGeneralRequest {
	s.SourceText = &v
	return s
}

func (s *TranslateGeneralRequest) SetTargetLanguage(v string) *TranslateGeneralRequest {
	s.TargetLanguage = &v
	return s
}

type TranslateGeneralResponseBody struct {
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *TranslateGeneralResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TranslateGeneralResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TranslateGeneralResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateGeneralResponseBody) SetCode(v int32) *TranslateGeneralResponseBody {
	s.Code = &v
	return s
}

func (s *TranslateGeneralResponseBody) SetData(v *TranslateGeneralResponseBodyData) *TranslateGeneralResponseBody {
	s.Data = v
	return s
}

func (s *TranslateGeneralResponseBody) SetMessage(v string) *TranslateGeneralResponseBody {
	s.Message = &v
	return s
}

func (s *TranslateGeneralResponseBody) SetRequestId(v string) *TranslateGeneralResponseBody {
	s.RequestId = &v
	return s
}

type TranslateGeneralResponseBodyData struct {
	DetectedLanguage *string `json:"DetectedLanguage,omitempty" xml:"DetectedLanguage,omitempty"`
	Translated       *string `json:"Translated,omitempty" xml:"Translated,omitempty"`
	WordCount        *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s TranslateGeneralResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TranslateGeneralResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateGeneralResponseBodyData) SetDetectedLanguage(v string) *TranslateGeneralResponseBodyData {
	s.DetectedLanguage = &v
	return s
}

func (s *TranslateGeneralResponseBodyData) SetTranslated(v string) *TranslateGeneralResponseBodyData {
	s.Translated = &v
	return s
}

func (s *TranslateGeneralResponseBodyData) SetWordCount(v string) *TranslateGeneralResponseBodyData {
	s.WordCount = &v
	return s
}

type TranslateGeneralResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TranslateGeneralResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TranslateGeneralResponse) SetStatusCode(v int32) *TranslateGeneralResponse {
	s.StatusCode = &v
	return s
}

func (s *TranslateGeneralResponse) SetBody(v *TranslateGeneralResponseBody) *TranslateGeneralResponse {
	s.Body = v
	return s
}

type TranslateImageRequest struct {
	Ext            *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Field          *string `json:"Field,omitempty" xml:"Field,omitempty"`
	ImageBase64    *string `json:"ImageBase64,omitempty" xml:"ImageBase64,omitempty"`
	ImageUrl       *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s TranslateImageRequest) String() string {
	return tea.Prettify(s)
}

func (s TranslateImageRequest) GoString() string {
	return s.String()
}

func (s *TranslateImageRequest) SetExt(v string) *TranslateImageRequest {
	s.Ext = &v
	return s
}

func (s *TranslateImageRequest) SetField(v string) *TranslateImageRequest {
	s.Field = &v
	return s
}

func (s *TranslateImageRequest) SetImageBase64(v string) *TranslateImageRequest {
	s.ImageBase64 = &v
	return s
}

func (s *TranslateImageRequest) SetImageUrl(v string) *TranslateImageRequest {
	s.ImageUrl = &v
	return s
}

func (s *TranslateImageRequest) SetSourceLanguage(v string) *TranslateImageRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TranslateImageRequest) SetTargetLanguage(v string) *TranslateImageRequest {
	s.TargetLanguage = &v
	return s
}

type TranslateImageResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *TranslateImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TranslateImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TranslateImageResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateImageResponseBody) SetCode(v int32) *TranslateImageResponseBody {
	s.Code = &v
	return s
}

func (s *TranslateImageResponseBody) SetData(v *TranslateImageResponseBodyData) *TranslateImageResponseBody {
	s.Data = v
	return s
}

func (s *TranslateImageResponseBody) SetMessage(v string) *TranslateImageResponseBody {
	s.Message = &v
	return s
}

func (s *TranslateImageResponseBody) SetRequestId(v string) *TranslateImageResponseBody {
	s.RequestId = &v
	return s
}

type TranslateImageResponseBodyData struct {
	FinalImageUrl *string `json:"FinalImageUrl,omitempty" xml:"FinalImageUrl,omitempty"`
	InPaintingUrl *string `json:"InPaintingUrl,omitempty" xml:"InPaintingUrl,omitempty"`
	TemplateJson  *string `json:"TemplateJson,omitempty" xml:"TemplateJson,omitempty"`
}

func (s TranslateImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TranslateImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateImageResponseBodyData) SetFinalImageUrl(v string) *TranslateImageResponseBodyData {
	s.FinalImageUrl = &v
	return s
}

func (s *TranslateImageResponseBodyData) SetInPaintingUrl(v string) *TranslateImageResponseBodyData {
	s.InPaintingUrl = &v
	return s
}

func (s *TranslateImageResponseBodyData) SetTemplateJson(v string) *TranslateImageResponseBodyData {
	s.TemplateJson = &v
	return s
}

type TranslateImageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TranslateImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TranslateImageResponse) String() string {
	return tea.Prettify(s)
}

func (s TranslateImageResponse) GoString() string {
	return s.String()
}

func (s *TranslateImageResponse) SetHeaders(v map[string]*string) *TranslateImageResponse {
	s.Headers = v
	return s
}

func (s *TranslateImageResponse) SetStatusCode(v int32) *TranslateImageResponse {
	s.StatusCode = &v
	return s
}

func (s *TranslateImageResponse) SetBody(v *TranslateImageResponseBody) *TranslateImageResponse {
	s.Body = v
	return s
}

type TranslateImageBatchRequest struct {
	CustomTaskId   *string `json:"CustomTaskId,omitempty" xml:"CustomTaskId,omitempty"`
	Ext            *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Field          *string `json:"Field,omitempty" xml:"Field,omitempty"`
	ImageUrls      *string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty"`
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s TranslateImageBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s TranslateImageBatchRequest) GoString() string {
	return s.String()
}

func (s *TranslateImageBatchRequest) SetCustomTaskId(v string) *TranslateImageBatchRequest {
	s.CustomTaskId = &v
	return s
}

func (s *TranslateImageBatchRequest) SetExt(v string) *TranslateImageBatchRequest {
	s.Ext = &v
	return s
}

func (s *TranslateImageBatchRequest) SetField(v string) *TranslateImageBatchRequest {
	s.Field = &v
	return s
}

func (s *TranslateImageBatchRequest) SetImageUrls(v string) *TranslateImageBatchRequest {
	s.ImageUrls = &v
	return s
}

func (s *TranslateImageBatchRequest) SetSourceLanguage(v string) *TranslateImageBatchRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TranslateImageBatchRequest) SetTargetLanguage(v string) *TranslateImageBatchRequest {
	s.TargetLanguage = &v
	return s
}

type TranslateImageBatchResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *TranslateImageBatchResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TranslateImageBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TranslateImageBatchResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateImageBatchResponseBody) SetCode(v int32) *TranslateImageBatchResponseBody {
	s.Code = &v
	return s
}

func (s *TranslateImageBatchResponseBody) SetData(v *TranslateImageBatchResponseBodyData) *TranslateImageBatchResponseBody {
	s.Data = v
	return s
}

func (s *TranslateImageBatchResponseBody) SetMessage(v string) *TranslateImageBatchResponseBody {
	s.Message = &v
	return s
}

func (s *TranslateImageBatchResponseBody) SetRequestId(v string) *TranslateImageBatchResponseBody {
	s.RequestId = &v
	return s
}

type TranslateImageBatchResponseBodyData struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TranslateImageBatchResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TranslateImageBatchResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateImageBatchResponseBodyData) SetTaskId(v string) *TranslateImageBatchResponseBodyData {
	s.TaskId = &v
	return s
}

type TranslateImageBatchResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TranslateImageBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TranslateImageBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s TranslateImageBatchResponse) GoString() string {
	return s.String()
}

func (s *TranslateImageBatchResponse) SetHeaders(v map[string]*string) *TranslateImageBatchResponse {
	s.Headers = v
	return s
}

func (s *TranslateImageBatchResponse) SetStatusCode(v int32) *TranslateImageBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *TranslateImageBatchResponse) SetBody(v *TranslateImageBatchResponseBody) *TranslateImageBatchResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
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

func (client *Client) CreateAsyncTranslateWithOptions(request *CreateAsyncTranslateRequest, runtime *util.RuntimeOptions) (_result *CreateAsyncTranslateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		body["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.FormatType)) {
		body["FormatType"] = request.FormatType
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.SourceLanguage)) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.SourceText)) {
		body["SourceText"] = request.SourceText
	}

	if !tea.BoolValue(util.IsUnset(request.TargetLanguage)) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAsyncTranslate"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAsyncTranslateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAsyncTranslate(request *CreateAsyncTranslateRequest) (_result *CreateAsyncTranslateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAsyncTranslateResponse{}
	_body, _err := client.CreateAsyncTranslateWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		body["CallbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		body["FileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.SourceLanguage)) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.TargetLanguage)) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDocTranslateTask"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDocTranslateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.FileUrlObject)) {
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
			Content:     request.FileUrlObject,
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
		createDocTranslateTaskReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		body["Extra"] = request.Extra
	}

	if !tea.BoolValue(util.IsUnset(request.SourceLanguage)) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.TargetLanguage)) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.UrlList)) {
		body["UrlList"] = request.UrlList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateImageTranslateTask"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateImageTranslateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetAsyncTranslateWithOptions(request *GetAsyncTranslateRequest, runtime *util.RuntimeOptions) (_result *GetAsyncTranslateResponse, _err error) {
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
		Action:      tea.String("GetAsyncTranslate"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncTranslateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncTranslate(request *GetAsyncTranslateRequest) (_result *GetAsyncTranslateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncTranslateResponse{}
	_body, _err := client.GetAsyncTranslateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBatchTranslateWithOptions(request *GetBatchTranslateRequest, runtime *util.RuntimeOptions) (_result *GetBatchTranslateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		body["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.FormatType)) {
		body["FormatType"] = request.FormatType
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.SourceLanguage)) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.SourceText)) {
		body["SourceText"] = request.SourceText
	}

	if !tea.BoolValue(util.IsUnset(request.TargetLanguage)) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBatchTranslate"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBatchTranslateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBatchTranslate(request *GetBatchTranslateRequest) (_result *GetBatchTranslateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBatchTranslateResponse{}
	_body, _err := client.GetBatchTranslateWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceText)) {
		body["SourceText"] = request.SourceText
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDetectLanguage"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDetectLanguageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocTranslateTask"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocTranslateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		body["Extra"] = request.Extra
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetImageDiagnose"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetImageDiagnoseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		body["Extra"] = request.Extra
	}

	if !tea.BoolValue(util.IsUnset(request.SourceLanguage)) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.TargetLanguage)) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetImageTranslate"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetImageTranslateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetImageTranslateTask"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetImageTranslateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		body["Extra"] = request.Extra
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		body["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTitleDiagnose"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTitleDiagnoseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Attributes)) {
		body["Attributes"] = request.Attributes
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		body["Extra"] = request.Extra
	}

	if !tea.BoolValue(util.IsUnset(request.HotWords)) {
		body["HotWords"] = request.HotWords
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		body["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTitleGenerate"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTitleGenerateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatLevelThreeId)) {
		body["CatLevelThreeId"] = request.CatLevelThreeId
	}

	if !tea.BoolValue(util.IsUnset(request.CatLevelTwoId)) {
		body["CatLevelTwoId"] = request.CatLevelTwoId
	}

	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		body["Extra"] = request.Extra
	}

	if !tea.BoolValue(util.IsUnset(request.Keywords)) {
		body["Keywords"] = request.Keywords
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		body["Platform"] = request.Platform
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTitleIntelligence"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTitleIntelligenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetTranslateImageBatchResultWithOptions(request *GetTranslateImageBatchResultRequest, runtime *util.RuntimeOptions) (_result *GetTranslateImageBatchResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTranslateImageBatchResult"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTranslateImageBatchResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTranslateImageBatchResult(request *GetTranslateImageBatchResultRequest) (_result *GetTranslateImageBatchResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTranslateImageBatchResultResponse{}
	_body, _err := client.GetTranslateImageBatchResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTranslateReportWithOptions(request *GetTranslateReportRequest, runtime *util.RuntimeOptions) (_result *GetTranslateReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiName)) {
		query["ApiName"] = request.ApiName
	}

	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Group)) {
		query["Group"] = request.Group
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTranslateReport"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTranslateReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTranslateReport(request *GetTranslateReportRequest) (_result *GetTranslateReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTranslateReportResponse{}
	_body, _err := client.GetTranslateReportWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenAlimtService"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenAlimtServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Context)) {
		query["Context"] = request.Context
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FormatType)) {
		body["FormatType"] = request.FormatType
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.SourceLanguage)) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.SourceText)) {
		body["SourceText"] = request.SourceText
	}

	if !tea.BoolValue(util.IsUnset(request.TargetLanguage)) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Translate"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TranslateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertificateType)) {
		body["CertificateType"] = request.CertificateType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ResultType)) {
		body["ResultType"] = request.ResultType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceLanguage)) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.TargetLanguage)) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TranslateCertificate"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TranslateCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
		translateCertificateReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	translateCertificateResp, _err := client.TranslateCertificateWithOptions(translateCertificateReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = translateCertificateResp
	return _result, _err
}

/**
 * @deprecated : TranslateECommerce is deprecated, please use alimt::2018-10-12::Translate instead.
 *
 * @param request TranslateECommerceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return TranslateECommerceResponse
 */
// Deprecated
func (client *Client) TranslateECommerceWithOptions(request *TranslateECommerceRequest, runtime *util.RuntimeOptions) (_result *TranslateECommerceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Context)) {
		query["Context"] = request.Context
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FormatType)) {
		body["FormatType"] = request.FormatType
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.SourceLanguage)) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.SourceText)) {
		body["SourceText"] = request.SourceText
	}

	if !tea.BoolValue(util.IsUnset(request.TargetLanguage)) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TranslateECommerce"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TranslateECommerceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated : TranslateECommerce is deprecated, please use alimt::2018-10-12::Translate instead.
 *
 * @param request TranslateECommerceRequest
 * @return TranslateECommerceResponse
 */
// Deprecated
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Context)) {
		query["Context"] = request.Context
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FormatType)) {
		body["FormatType"] = request.FormatType
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.SourceLanguage)) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.SourceText)) {
		body["SourceText"] = request.SourceText
	}

	if !tea.BoolValue(util.IsUnset(request.TargetLanguage)) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TranslateGeneral"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TranslateGeneralResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) TranslateImageWithOptions(request *TranslateImageRequest, runtime *util.RuntimeOptions) (_result *TranslateImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["Ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.Field)) {
		body["Field"] = request.Field
	}

	if !tea.BoolValue(util.IsUnset(request.ImageBase64)) {
		body["ImageBase64"] = request.ImageBase64
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SourceLanguage)) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.TargetLanguage)) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TranslateImage"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TranslateImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TranslateImage(request *TranslateImageRequest) (_result *TranslateImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TranslateImageResponse{}
	_body, _err := client.TranslateImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TranslateImageBatchWithOptions(request *TranslateImageBatchRequest, runtime *util.RuntimeOptions) (_result *TranslateImageBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomTaskId)) {
		body["CustomTaskId"] = request.CustomTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["Ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.Field)) {
		body["Field"] = request.Field
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrls)) {
		body["ImageUrls"] = request.ImageUrls
	}

	if !tea.BoolValue(util.IsUnset(request.SourceLanguage)) {
		body["SourceLanguage"] = request.SourceLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.TargetLanguage)) {
		body["TargetLanguage"] = request.TargetLanguage
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TranslateImageBatch"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TranslateImageBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TranslateImageBatch(request *TranslateImageBatchRequest) (_result *TranslateImageBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TranslateImageBatchResponse{}
	_body, _err := client.TranslateImageBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
