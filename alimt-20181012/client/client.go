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

type CreateAsyncTranslateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// translate_standard
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// general
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Opinion: We have finally gotten some relief at the pump. But it may not last long
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
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
	// example:
	//
	// 200
	Code *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateAsyncTranslateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DC2DCCC9-C3DF-4F59-8D8E-78185729F16D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 98bbb007-71bb-448b-bab0-2695ce8f8599
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ready
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAsyncTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// http://callbackUrl
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// example:
	//
	// token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://fileUrl
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// general
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
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
	// example:
	//
	// http://callbackUrl
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// example:
	//
	// token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://fileUrl
	FileUrlObject io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// general
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
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
	// example:
	//
	// D3920BC3-A395-4CAD-9E84-7C39EB07507B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ready
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0586df512c8b4bb382d7d9a6a01b5854
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDocTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 1
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// {"have_ocr":"false","without_text":"false","have_psd":"true","ignore_entity":"false"}
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://xxx,http://yyy
	UrlList *string `json:"UrlList,omitempty" xml:"UrlList,omitempty"`
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
	// example:
	//
	// 200
	Code *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateImageTranslateTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A41F6E25-8520-4AF0-90EF-AF7E32840108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// A41F6E25-8520-4AF0-90EF-111111
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 77056ab7-7be1-4c2a-91a1-f20f63894048
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
	// example:
	//
	// 200
	Code *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAsyncTranslateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DC2DCCC9-C3DF-4F59-8D8E-78185729F16D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// zh
	DetectedLanguage *string `json:"DetectedLanguage,omitempty" xml:"DetectedLanguage,omitempty"`
	// example:
	//
	// ready
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// hello
	TranslatedText *string `json:"TranslatedText,omitempty" xml:"TranslatedText,omitempty"`
	// example:
	//
	// 2
	WordCount *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAsyncTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// translate_standard
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// general
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"11":"hello boy","12":"go home","13":"we can"}
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
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
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86D18195-D89C-4C8C-9DC4-5FCE789CE6D5
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBatchTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetBatchTranslateByVPCRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// translate_standard
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// general
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s GetBatchTranslateByVPCRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBatchTranslateByVPCRequest) GoString() string {
	return s.String()
}

func (s *GetBatchTranslateByVPCRequest) SetApiType(v string) *GetBatchTranslateByVPCRequest {
	s.ApiType = &v
	return s
}

func (s *GetBatchTranslateByVPCRequest) SetFormatType(v string) *GetBatchTranslateByVPCRequest {
	s.FormatType = &v
	return s
}

func (s *GetBatchTranslateByVPCRequest) SetScene(v string) *GetBatchTranslateByVPCRequest {
	s.Scene = &v
	return s
}

func (s *GetBatchTranslateByVPCRequest) SetSourceLanguage(v string) *GetBatchTranslateByVPCRequest {
	s.SourceLanguage = &v
	return s
}

func (s *GetBatchTranslateByVPCRequest) SetSourceText(v string) *GetBatchTranslateByVPCRequest {
	s.SourceText = &v
	return s
}

func (s *GetBatchTranslateByVPCRequest) SetTargetLanguage(v string) *GetBatchTranslateByVPCRequest {
	s.TargetLanguage = &v
	return s
}

type GetBatchTranslateByVPCResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Param checker failed! param:[sourceText]. reason: stringChecker not pass. Param length not less [0] and not greater[10000]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DC2DCCC9-C3DF-4F59-8D8E-78185729F16D
	RequestId      *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TranslatedList []map[string]interface{} `json:"TranslatedList,omitempty" xml:"TranslatedList,omitempty" type:"Repeated"`
}

func (s GetBatchTranslateByVPCResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBatchTranslateByVPCResponseBody) GoString() string {
	return s.String()
}

func (s *GetBatchTranslateByVPCResponseBody) SetCode(v int32) *GetBatchTranslateByVPCResponseBody {
	s.Code = &v
	return s
}

func (s *GetBatchTranslateByVPCResponseBody) SetMessage(v string) *GetBatchTranslateByVPCResponseBody {
	s.Message = &v
	return s
}

func (s *GetBatchTranslateByVPCResponseBody) SetRequestId(v string) *GetBatchTranslateByVPCResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBatchTranslateByVPCResponseBody) SetTranslatedList(v []map[string]interface{}) *GetBatchTranslateByVPCResponseBody {
	s.TranslatedList = v
	return s
}

type GetBatchTranslateByVPCResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBatchTranslateByVPCResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBatchTranslateByVPCResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBatchTranslateByVPCResponse) GoString() string {
	return s.String()
}

func (s *GetBatchTranslateByVPCResponse) SetHeaders(v map[string]*string) *GetBatchTranslateByVPCResponse {
	s.Headers = v
	return s
}

func (s *GetBatchTranslateByVPCResponse) SetStatusCode(v int32) *GetBatchTranslateByVPCResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBatchTranslateByVPCResponse) SetBody(v *GetBatchTranslateByVPCResponseBody) *GetBatchTranslateByVPCResponse {
	s.Body = v
	return s
}

type GetDetectLanguageRequest struct {
	// This parameter is required.
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
	// example:
	//
	// zh
	DetectedLanguage      *string `json:"DetectedLanguage,omitempty" xml:"DetectedLanguage,omitempty"`
	LanguageProbabilities *string `json:"LanguageProbabilities,omitempty" xml:"LanguageProbabilities,omitempty"`
	// example:
	//
	// 0C5EC1EC-1A06-4D60-97E6-4D41350945E4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *GetDetectLanguageResponseBody) SetLanguageProbabilities(v string) *GetDetectLanguageResponseBody {
	s.LanguageProbabilities = &v
	return s
}

func (s *GetDetectLanguageResponseBody) SetRequestId(v string) *GetDetectLanguageResponseBody {
	s.RequestId = &v
	return s
}

type GetDetectLanguageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDetectLanguageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetDetectLanguageVpcRequest struct {
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
}

func (s GetDetectLanguageVpcRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDetectLanguageVpcRequest) GoString() string {
	return s.String()
}

func (s *GetDetectLanguageVpcRequest) SetSourceText(v string) *GetDetectLanguageVpcRequest {
	s.SourceText = &v
	return s
}

type GetDetectLanguageVpcResponseBody struct {
	DetectedLanguage      *string `json:"DetectedLanguage,omitempty" xml:"DetectedLanguage,omitempty"`
	LanguageProbabilities *string `json:"LanguageProbabilities,omitempty" xml:"LanguageProbabilities,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDetectLanguageVpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDetectLanguageVpcResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetectLanguageVpcResponseBody) SetDetectedLanguage(v string) *GetDetectLanguageVpcResponseBody {
	s.DetectedLanguage = &v
	return s
}

func (s *GetDetectLanguageVpcResponseBody) SetLanguageProbabilities(v string) *GetDetectLanguageVpcResponseBody {
	s.LanguageProbabilities = &v
	return s
}

func (s *GetDetectLanguageVpcResponseBody) SetRequestId(v string) *GetDetectLanguageVpcResponseBody {
	s.RequestId = &v
	return s
}

type GetDetectLanguageVpcResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDetectLanguageVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDetectLanguageVpcResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDetectLanguageVpcResponse) GoString() string {
	return s.String()
}

func (s *GetDetectLanguageVpcResponse) SetHeaders(v map[string]*string) *GetDetectLanguageVpcResponse {
	s.Headers = v
	return s
}

func (s *GetDetectLanguageVpcResponse) SetStatusCode(v int32) *GetDetectLanguageVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDetectLanguageVpcResponse) SetBody(v *GetDetectLanguageVpcResponseBody) *GetDetectLanguageVpcResponse {
	s.Body = v
	return s
}

type GetDocTranslateTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0586df512c8b4bb382d7d9a6a01b5854
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
	// example:
	//
	// 20
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// example:
	//
	// 12952E92-FDF3-4D3C-88E3-242D72BA7150
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// translated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0586df512c8b4bb382d7d9a6a01b5854
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// Error
	TranslateErrorCode *string `json:"TranslateErrorCode,omitempty" xml:"TranslateErrorCode,omitempty"`
	// example:
	//
	// Fail to get the page number of document.
	TranslateErrorMessage *string `json:"TranslateErrorMessage,omitempty" xml:"TranslateErrorMessage,omitempty"`
	// example:
	//
	// http://translateFileUrl
	TranslateFileUrl *string `json:"TranslateFileUrl,omitempty" xml:"TranslateFileUrl,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// {   "product_id": "1",   "platform": "ae" }
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://xxxxx.oss-cn-shenzhen.aliyuncs.com/jd/41209/xxxxx.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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
	// example:
	//
	// 200
	Code *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetImageDiagnoseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DC2DCCC9-C3DF-4F59-8D8E-78185729F16D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// zh
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageDiagnoseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// {"have_ocr": "false", "without_text":"true", "have_psd": "false", "ignore_entity": "false"}
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://xxxxxxx.oss-cn-shenzhen.aliyuncs.com/xxxxxx.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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
	// example:
	//
	// 200
	Code *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetImageTranslateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A41F6E25-8520-4AF0-90EF-AF7E32840108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// https://ae01.alicdn.com/kf/xxxxx.jpeg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// xxxxxx
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
	// example:
	//
	// 200
	Code *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetImageTranslateTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// xxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// [{xxxxxx}]
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 123
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// {   "product_id": "1",   "platform": "ae" }
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ae
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Amroe Japan Comic Theme Uzumaki Naruto Namikaze Minato 3D Visual Cartoon 7 Color USB Touch
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
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
	// example:
	//
	// 200
	Code *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTitleDiagnoseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DC2DCCC9-C3DF-4F59-8D8E-78185729F16D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// Boy
	AllUppercaseWords *string `json:"AllUppercaseWords,omitempty" xml:"AllUppercaseWords,omitempty"`
	// example:
	//
	// true
	ContainCoreClasses *string `json:"ContainCoreClasses,omitempty" xml:"ContainCoreClasses,omitempty"`
	// example:
	//
	// baba
	DisableWords *string `json:"DisableWords,omitempty" xml:"DisableWords,omitempty"`
	// example:
	//
	// hi
	DuplicateWords *string `json:"DuplicateWords,omitempty" xml:"DuplicateWords,omitempty"`
	// example:
	//
	// 2
	LanguageQualityScore *string `json:"LanguageQualityScore,omitempty" xml:"LanguageQualityScore,omitempty"`
	// example:
	//
	// no
	NoFirstUppercaseList *string `json:"NoFirstUppercaseList,omitempty" xml:"NoFirstUppercaseList,omitempty"`
	// example:
	//
	// 100
	OverLengthLimit *string `json:"OverLengthLimit,omitempty" xml:"OverLengthLimit,omitempty"`
	// example:
	//
	// 3
	TotalScore *string `json:"TotalScore,omitempty" xml:"TotalScore,omitempty"`
	// example:
	//
	// 10
	WordCount *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
	// example:
	//
	// ahk
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTitleDiagnoseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// {2:"None",10:"Plastic"}
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 127896018
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// {   "product_id": "1",   "platform": "ae" }
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// None,Plastic
	HotWords *string `json:"HotWords,omitempty" xml:"HotWords,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ae
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10pcs 80ml Kitchen Disposable
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
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
	// example:
	//
	// 200
	Code *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTitleGenerateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DC2DCCC9-C3DF-4F59-8D8E-78185729F16D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 10pcs 80ml Kitchen Disposable Plastic Sauce Cup Pot Chutney Container
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTitleGenerateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 111
	CatLevelThreeId *int64 `json:"CatLevelThreeId,omitempty" xml:"CatLevelThreeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 222
	CatLevelTwoId *int64 `json:"CatLevelTwoId,omitempty" xml:"CatLevelTwoId,omitempty"`
	// example:
	//
	// {"product_id":"1212"}
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// hello,apple
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ae
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
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
	// example:
	//
	// 200
	Code *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTitleIntelligenceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D70487B4-8891-4D24-BB64-8788E53671FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// Custom Hello Kitty PVC Cartoon Apple for Home Garden Complete Apple Bath Shower
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTitleIntelligenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// EEA28E6D-4828-5031-BD8C-8FF1B3216842
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
	// example:
	//
	// 200
	Code *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTranslateImageBatchResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DACD263C-C068-5116-83EC-117245AA35CF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// https://example.com/example.jpg
	FinalImageUrl *string `json:"FinalImageUrl,omitempty" xml:"FinalImageUrl,omitempty"`
	// example:
	//
	// https://example.com/example.jpg
	InPaintingUrl *string `json:"InPaintingUrl,omitempty" xml:"InPaintingUrl,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// https://example.com/example.jpg
	SourceImageUrl *string `json:"SourceImageUrl,omitempty" xml:"SourceImageUrl,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// {"TemplateJson": "Editor Template Json String	"}
	TemplateJson *string `json:"TemplateJson,omitempty" xml:"TemplateJson,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTranslateImageBatchResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-03-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-03-01 23:59:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
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
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"result":[{"time":"2021-10-21 00:00:00","total":100}]}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86D18195-D89C-4C8C-9DC4-5FCE789CE6D5
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTranslateReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// id
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// example:
	//
	// 123456
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// xxxx-xxxxx
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenAlimtServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// {\\"appName\\":\\"alynx\\"}
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// title
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
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
	// example:
	//
	// 200
	Code *int32                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *TranslateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86D18195-D89C-4C8C-9DC4-5FCE789CE6D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// Hello
	Translated *string `json:"Translated,omitempty" xml:"Translated,omitempty"`
	// example:
	//
	// 10
	WordCount *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// driving_license
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://imageurl
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// text
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// driving_license
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://imageurl
	ImageUrlObject io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// text
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
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
	Data *TranslateCertificateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 6640DE48-201C-4110-AE87-FB8FA34412B9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// name
	KeyTranslation *string `json:"KeyTranslation,omitempty" xml:"KeyTranslation,omitempty"`
	Value          *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// Solemn
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TranslateCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// context信息
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// social
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Hello
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
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
	// example:
	//
	// 200
	Code *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *TranslateECommerceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// An error occurred while translating.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CC93BB5C-EAB5-579B-AA44-F61528F48FF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	DetectedLanguage *string `json:"DetectedLanguage,omitempty" xml:"DetectedLanguage,omitempty"`
	Translated       *string `json:"Translated,omitempty" xml:"Translated,omitempty"`
	// example:
	//
	// 10
	WordCount *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s TranslateECommerceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TranslateECommerceResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateECommerceResponseBodyData) SetDetectedLanguage(v string) *TranslateECommerceResponseBodyData {
	s.DetectedLanguage = &v
	return s
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TranslateECommerceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// {\\"appName\\":\\"alynx\\"}
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// example:
	//
	// general
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
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
	// example:
	//
	// 200
	Code *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *TranslateGeneralResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86D18195-D89C-4C8C-9DC4-5FCE789CE6D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// Hello
	Translated *string `json:"Translated,omitempty" xml:"Translated,omitempty"`
	// example:
	//
	// 10
	WordCount *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TranslateGeneralResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type TranslateGeneralVpcRequest struct {
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// general
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s TranslateGeneralVpcRequest) String() string {
	return tea.Prettify(s)
}

func (s TranslateGeneralVpcRequest) GoString() string {
	return s.String()
}

func (s *TranslateGeneralVpcRequest) SetContext(v string) *TranslateGeneralVpcRequest {
	s.Context = &v
	return s
}

func (s *TranslateGeneralVpcRequest) SetFormatType(v string) *TranslateGeneralVpcRequest {
	s.FormatType = &v
	return s
}

func (s *TranslateGeneralVpcRequest) SetScene(v string) *TranslateGeneralVpcRequest {
	s.Scene = &v
	return s
}

func (s *TranslateGeneralVpcRequest) SetSourceLanguage(v string) *TranslateGeneralVpcRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TranslateGeneralVpcRequest) SetSourceText(v string) *TranslateGeneralVpcRequest {
	s.SourceText = &v
	return s
}

func (s *TranslateGeneralVpcRequest) SetTargetLanguage(v string) *TranslateGeneralVpcRequest {
	s.TargetLanguage = &v
	return s
}

type TranslateGeneralVpcResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 200
	Code *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *TranslateGeneralVpcResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// translate from source to target not support
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DC2DCCC9-C3DF-4F59-8D8E-78185729F16D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TranslateGeneralVpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TranslateGeneralVpcResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateGeneralVpcResponseBody) SetCode(v int32) *TranslateGeneralVpcResponseBody {
	s.Code = &v
	return s
}

func (s *TranslateGeneralVpcResponseBody) SetData(v *TranslateGeneralVpcResponseBodyData) *TranslateGeneralVpcResponseBody {
	s.Data = v
	return s
}

func (s *TranslateGeneralVpcResponseBody) SetMessage(v string) *TranslateGeneralVpcResponseBody {
	s.Message = &v
	return s
}

func (s *TranslateGeneralVpcResponseBody) SetRequestId(v string) *TranslateGeneralVpcResponseBody {
	s.RequestId = &v
	return s
}

type TranslateGeneralVpcResponseBodyData struct {
	// example:
	//
	// zh
	DetectedLanguage *string `json:"DetectedLanguage,omitempty" xml:"DetectedLanguage,omitempty"`
	// example:
	//
	// Hello
	Translated *string `json:"Translated,omitempty" xml:"Translated,omitempty"`
	// example:
	//
	// 10
	WordCount *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s TranslateGeneralVpcResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TranslateGeneralVpcResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateGeneralVpcResponseBodyData) SetDetectedLanguage(v string) *TranslateGeneralVpcResponseBodyData {
	s.DetectedLanguage = &v
	return s
}

func (s *TranslateGeneralVpcResponseBodyData) SetTranslated(v string) *TranslateGeneralVpcResponseBodyData {
	s.Translated = &v
	return s
}

func (s *TranslateGeneralVpcResponseBodyData) SetWordCount(v string) *TranslateGeneralVpcResponseBodyData {
	s.WordCount = &v
	return s
}

type TranslateGeneralVpcResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TranslateGeneralVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TranslateGeneralVpcResponse) String() string {
	return tea.Prettify(s)
}

func (s TranslateGeneralVpcResponse) GoString() string {
	return s.String()
}

func (s *TranslateGeneralVpcResponse) SetHeaders(v map[string]*string) *TranslateGeneralVpcResponse {
	s.Headers = v
	return s
}

func (s *TranslateGeneralVpcResponse) SetStatusCode(v int32) *TranslateGeneralVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *TranslateGeneralVpcResponse) SetBody(v *TranslateGeneralVpcResponseBody) *TranslateGeneralVpcResponse {
	s.Body = v
	return s
}

type TranslateImageRequest struct {
	// example:
	//
	// {"needEditorData": "false", "ignoreEntityRecognize": "true"}
	Ext *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// general
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// example:
	//
	// /9j/4...H/9k=
	ImageBase64 *string `json:"ImageBase64,omitempty" xml:"ImageBase64,omitempty"`
	// example:
	//
	// https://example.com/example.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// example:
	//
	// en
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
	// example:
	//
	// 200
	Code *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *TranslateImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Error Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D774D33D-F1CB-5A2C-A787-E0A2179239CE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// https://example.com/example.jpg
	FinalImageUrl *string `json:"FinalImageUrl,omitempty" xml:"FinalImageUrl,omitempty"`
	// example:
	//
	// https://example.com/example.jpg
	InPaintingUrl *string `json:"InPaintingUrl,omitempty" xml:"InPaintingUrl,omitempty"`
	// example:
	//
	// Editor Template Json String
	TemplateJson *string `json:"TemplateJson,omitempty" xml:"TemplateJson,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TranslateImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// my_awesome_task_1
	CustomTaskId *string `json:"CustomTaskId,omitempty" xml:"CustomTaskId,omitempty"`
	// example:
	//
	// {"needEditorData": "false", "ignoreEntityRecognize": "true"}
	Ext *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// general
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/1.jpg,https://example.com/2.jpg,https://example.com/3.jpg
	ImageUrls *string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
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
	// example:
	//
	// 200
	Code *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *TranslateImageBatchResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D774D33D-F1CB-5A2C-A787-E0A2179239CE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// EEA28E6D-4828-5031-BD8C-8FF1B3216842
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TranslateImageBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type TranslateSearchRequest struct {
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// text
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// query
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 今天天气不错
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s TranslateSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s TranslateSearchRequest) GoString() string {
	return s.String()
}

func (s *TranslateSearchRequest) SetFormatType(v string) *TranslateSearchRequest {
	s.FormatType = &v
	return s
}

func (s *TranslateSearchRequest) SetScene(v string) *TranslateSearchRequest {
	s.Scene = &v
	return s
}

func (s *TranslateSearchRequest) SetSourceLanguage(v string) *TranslateSearchRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TranslateSearchRequest) SetSourceText(v string) *TranslateSearchRequest {
	s.SourceText = &v
	return s
}

func (s *TranslateSearchRequest) SetTargetLanguage(v string) *TranslateSearchRequest {
	s.TargetLanguage = &v
	return s
}

type TranslateSearchResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *TranslateSearchResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TranslateSearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TranslateSearchResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateSearchResponseBody) SetCode(v string) *TranslateSearchResponseBody {
	s.Code = &v
	return s
}

func (s *TranslateSearchResponseBody) SetData(v *TranslateSearchResponseBodyData) *TranslateSearchResponseBody {
	s.Data = v
	return s
}

func (s *TranslateSearchResponseBody) SetMessage(v string) *TranslateSearchResponseBody {
	s.Message = &v
	return s
}

func (s *TranslateSearchResponseBody) SetRequestId(v string) *TranslateSearchResponseBody {
	s.RequestId = &v
	return s
}

type TranslateSearchResponseBodyData struct {
	Translated *string `json:"Translated,omitempty" xml:"Translated,omitempty"`
}

func (s TranslateSearchResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TranslateSearchResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateSearchResponseBodyData) SetTranslated(v string) *TranslateSearchResponseBodyData {
	s.Translated = &v
	return s
}

type TranslateSearchResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TranslateSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TranslateSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s TranslateSearchResponse) GoString() string {
	return s.String()
}

func (s *TranslateSearchResponse) SetHeaders(v map[string]*string) *TranslateSearchResponse {
	s.Headers = v
	return s
}

func (s *TranslateSearchResponse) SetStatusCode(v int32) *TranslateSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *TranslateSearchResponse) SetBody(v *TranslateSearchResponseBody) *TranslateSearchResponse {
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

// Summary:
//
// 大文本异步翻译，支持5000-50000字翻译
//
// @param request - CreateAsyncTranslateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAsyncTranslateResponse
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

// Summary:
//
// 大文本异步翻译，支持5000-50000字翻译
//
// @param request - CreateAsyncTranslateRequest
//
// @return CreateAsyncTranslateResponse
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

// @param request - CreateDocTranslateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDocTranslateTaskResponse
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

// @param request - CreateDocTranslateTaskRequest
//
// @return CreateDocTranslateTaskResponse
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
		Product:  tea.String("alimt"),
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

// @param request - CreateImageTranslateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageTranslateTaskResponse
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

// @param request - CreateImageTranslateTaskRequest
//
// @return CreateImageTranslateTaskResponse
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

// Summary:
//
// 大文本异步翻译，支持5000-50000字翻译
//
// @param request - GetAsyncTranslateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsyncTranslateResponse
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

// Summary:
//
// 大文本异步翻译，支持5000-50000字翻译
//
// @param request - GetAsyncTranslateRequest
//
// @return GetAsyncTranslateResponse
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

// @param request - GetBatchTranslateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTranslateResponse
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

// @param request - GetBatchTranslateRequest
//
// @return GetBatchTranslateResponse
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

// Summary:
//
// # GetBatchTranslateByVPC
//
// @param request - GetBatchTranslateByVPCRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTranslateByVPCResponse
func (client *Client) GetBatchTranslateByVPCWithOptions(request *GetBatchTranslateByVPCRequest, runtime *util.RuntimeOptions) (_result *GetBatchTranslateByVPCResponse, _err error) {
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
		Action:      tea.String("GetBatchTranslateByVPC"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBatchTranslateByVPCResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetBatchTranslateByVPC
//
// @param request - GetBatchTranslateByVPCRequest
//
// @return GetBatchTranslateByVPCResponse
func (client *Client) GetBatchTranslateByVPC(request *GetBatchTranslateByVPCRequest) (_result *GetBatchTranslateByVPCResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBatchTranslateByVPCResponse{}
	_body, _err := client.GetBatchTranslateByVPCWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 语种识别
//
// @param request - GetDetectLanguageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDetectLanguageResponse
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

// Summary:
//
// 语种识别
//
// @param request - GetDetectLanguageRequest
//
// @return GetDetectLanguageResponse
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

// Summary:
//
// 语种识别
//
// @param request - GetDetectLanguageVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDetectLanguageVpcResponse
func (client *Client) GetDetectLanguageVpcWithOptions(request *GetDetectLanguageVpcRequest, runtime *util.RuntimeOptions) (_result *GetDetectLanguageVpcResponse, _err error) {
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
		Action:      tea.String("GetDetectLanguageVpc"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDetectLanguageVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 语种识别
//
// @param request - GetDetectLanguageVpcRequest
//
// @return GetDetectLanguageVpcResponse
func (client *Client) GetDetectLanguageVpc(request *GetDetectLanguageVpcRequest) (_result *GetDetectLanguageVpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDetectLanguageVpcResponse{}
	_body, _err := client.GetDetectLanguageVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档翻译任务
//
// @param request - GetDocTranslateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocTranslateTaskResponse
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

// Summary:
//
// 获取文档翻译任务
//
// @param request - GetDocTranslateTaskRequest
//
// @return GetDocTranslateTaskResponse
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

// @param request - GetImageDiagnoseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageDiagnoseResponse
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

// @param request - GetImageDiagnoseRequest
//
// @return GetImageDiagnoseResponse
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

// @param request - GetImageTranslateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageTranslateResponse
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

// @param request - GetImageTranslateRequest
//
// @return GetImageTranslateResponse
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

// @param request - GetImageTranslateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageTranslateTaskResponse
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

// @param request - GetImageTranslateTaskRequest
//
// @return GetImageTranslateTaskResponse
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

// Summary:
//
// # GetTitleDiagnose
//
// @param request - GetTitleDiagnoseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTitleDiagnoseResponse
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

// Summary:
//
// # GetTitleDiagnose
//
// @param request - GetTitleDiagnoseRequest
//
// @return GetTitleDiagnoseResponse
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

// Summary:
//
// # GetTitleGenerate
//
// @param request - GetTitleGenerateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTitleGenerateResponse
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

// Summary:
//
// # GetTitleGenerate
//
// @param request - GetTitleGenerateRequest
//
// @return GetTitleGenerateResponse
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

// Summary:
//
// # GetTitleIntelligence
//
// @param request - GetTitleIntelligenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTitleIntelligenceResponse
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

// Summary:
//
// # GetTitleIntelligence
//
// @param request - GetTitleIntelligenceRequest
//
// @return GetTitleIntelligenceResponse
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

// Summary:
//
// 获取图片批量翻译结果
//
// @param request - GetTranslateImageBatchResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranslateImageBatchResultResponse
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

// Summary:
//
// 获取图片批量翻译结果
//
// @param request - GetTranslateImageBatchResultRequest
//
// @return GetTranslateImageBatchResultResponse
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

// Summary:
//
// # GetTranslateReport
//
// @param request - GetTranslateReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranslateReportResponse
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

// Summary:
//
// # GetTranslateReport
//
// @param request - GetTranslateReportRequest
//
// @return GetTranslateReportResponse
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

// Summary:
//
// 开通服务
//
// @param request - OpenAlimtServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenAlimtServiceResponse
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

// Summary:
//
// 开通服务
//
// @param request - OpenAlimtServiceRequest
//
// @return OpenAlimtServiceResponse
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

// Summary:
//
// 专业文本翻译
//
// @param request - TranslateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateResponse
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

// Summary:
//
// 专业文本翻译
//
// @param request - TranslateRequest
//
// @return TranslateResponse
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

// Summary:
//
// # TranslateCertificate
//
// @param request - TranslateCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateCertificateResponse
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

// Summary:
//
// # TranslateCertificate
//
// @param request - TranslateCertificateRequest
//
// @return TranslateCertificateResponse
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
		Product:  tea.String("alimt"),
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

// Deprecated: OpenAPI TranslateECommerce is deprecated, please use alimt::2018-10-12::Translate instead.
//
// Summary:
//
// # TranslateECommerce
//
// @param request - TranslateECommerceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateECommerceResponse
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

// Deprecated: OpenAPI TranslateECommerce is deprecated, please use alimt::2018-10-12::Translate instead.
//
// Summary:
//
// # TranslateECommerce
//
// @param request - TranslateECommerceRequest
//
// @return TranslateECommerceResponse
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

// Summary:
//
// 文本通用翻译
//
// @param request - TranslateGeneralRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateGeneralResponse
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

// Summary:
//
// 文本通用翻译
//
// @param request - TranslateGeneralRequest
//
// @return TranslateGeneralResponse
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

// Summary:
//
// # TranslateGeneralVpc
//
// @param request - TranslateGeneralVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateGeneralVpcResponse
func (client *Client) TranslateGeneralVpcWithOptions(request *TranslateGeneralVpcRequest, runtime *util.RuntimeOptions) (_result *TranslateGeneralVpcResponse, _err error) {
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
		Action:      tea.String("TranslateGeneralVpc"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TranslateGeneralVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # TranslateGeneralVpc
//
// @param request - TranslateGeneralVpcRequest
//
// @return TranslateGeneralVpcResponse
func (client *Client) TranslateGeneralVpc(request *TranslateGeneralVpcRequest) (_result *TranslateGeneralVpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TranslateGeneralVpcResponse{}
	_body, _err := client.TranslateGeneralVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 公有云图片翻译产品API
//
// @param request - TranslateImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateImageResponse
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

// Summary:
//
// 公有云图片翻译产品API
//
// @param request - TranslateImageRequest
//
// @return TranslateImageResponse
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

// Summary:
//
// 批量图片翻译接口
//
// @param request - TranslateImageBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateImageBatchResponse
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

// Summary:
//
// 批量图片翻译接口
//
// @param request - TranslateImageBatchRequest
//
// @return TranslateImageBatchResponse
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

// Summary:
//
// 搜索翻译
//
// @param request - TranslateSearchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TranslateSearchResponse
func (client *Client) TranslateSearchWithOptions(request *TranslateSearchRequest, runtime *util.RuntimeOptions) (_result *TranslateSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
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
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TranslateSearch"),
		Version:     tea.String("2018-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TranslateSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索翻译
//
// @param request - TranslateSearchRequest
//
// @return TranslateSearchResponse
func (client *Client) TranslateSearch(request *TranslateSearchRequest) (_result *TranslateSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TranslateSearchResponse{}
	_body, _err := client.TranslateSearchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
