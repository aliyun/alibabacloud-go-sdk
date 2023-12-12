// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryCopyrightRequest struct {
	CreateTimeEnd   *int64  `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	CreateTimeStart *int64  `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	JobId           *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Level           *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	PageNumber      *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryCopyrightRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCopyrightRequest) GoString() string {
	return s.String()
}

func (s *QueryCopyrightRequest) SetCreateTimeEnd(v int64) *QueryCopyrightRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *QueryCopyrightRequest) SetCreateTimeStart(v int64) *QueryCopyrightRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *QueryCopyrightRequest) SetJobId(v string) *QueryCopyrightRequest {
	s.JobId = &v
	return s
}

func (s *QueryCopyrightRequest) SetLevel(v int64) *QueryCopyrightRequest {
	s.Level = &v
	return s
}

func (s *QueryCopyrightRequest) SetPageNumber(v int64) *QueryCopyrightRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryCopyrightRequest) SetPageSize(v int64) *QueryCopyrightRequest {
	s.PageSize = &v
	return s
}

type QueryCopyrightResponseBody struct {
	Data       []*QueryCopyrightResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestID  *string                           `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	StatusCode *int64                            `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QueryCopyrightResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCopyrightResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCopyrightResponseBody) SetData(v []*QueryCopyrightResponseBodyData) *QueryCopyrightResponseBody {
	s.Data = v
	return s
}

func (s *QueryCopyrightResponseBody) SetRequestID(v string) *QueryCopyrightResponseBody {
	s.RequestID = &v
	return s
}

func (s *QueryCopyrightResponseBody) SetStatusCode(v int64) *QueryCopyrightResponseBody {
	s.StatusCode = &v
	return s
}

type QueryCopyrightResponseBodyData struct {
	Callback    *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Input       *string `json:"Input,omitempty" xml:"Input,omitempty"`
	JobId       *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Level       *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	MessageId   *int64  `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	Output      *string `json:"Output,omitempty" xml:"Output,omitempty"`
	Result      *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserData    *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	UserId      *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryCopyrightResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryCopyrightResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCopyrightResponseBodyData) SetCallback(v string) *QueryCopyrightResponseBodyData {
	s.Callback = &v
	return s
}

func (s *QueryCopyrightResponseBodyData) SetGmtCreate(v int64) *QueryCopyrightResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryCopyrightResponseBodyData) SetGmtModified(v int64) *QueryCopyrightResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *QueryCopyrightResponseBodyData) SetInput(v string) *QueryCopyrightResponseBodyData {
	s.Input = &v
	return s
}

func (s *QueryCopyrightResponseBodyData) SetJobId(v string) *QueryCopyrightResponseBodyData {
	s.JobId = &v
	return s
}

func (s *QueryCopyrightResponseBodyData) SetLevel(v int64) *QueryCopyrightResponseBodyData {
	s.Level = &v
	return s
}

func (s *QueryCopyrightResponseBodyData) SetMessage(v string) *QueryCopyrightResponseBodyData {
	s.Message = &v
	return s
}

func (s *QueryCopyrightResponseBodyData) SetMessageId(v int64) *QueryCopyrightResponseBodyData {
	s.MessageId = &v
	return s
}

func (s *QueryCopyrightResponseBodyData) SetOutput(v string) *QueryCopyrightResponseBodyData {
	s.Output = &v
	return s
}

func (s *QueryCopyrightResponseBodyData) SetResult(v string) *QueryCopyrightResponseBodyData {
	s.Result = &v
	return s
}

func (s *QueryCopyrightResponseBodyData) SetStatus(v string) *QueryCopyrightResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryCopyrightResponseBodyData) SetUserData(v string) *QueryCopyrightResponseBodyData {
	s.UserData = &v
	return s
}

func (s *QueryCopyrightResponseBodyData) SetUserId(v int64) *QueryCopyrightResponseBodyData {
	s.UserId = &v
	return s
}

type QueryCopyrightResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCopyrightResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCopyrightResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCopyrightResponse) GoString() string {
	return s.String()
}

func (s *QueryCopyrightResponse) SetHeaders(v map[string]*string) *QueryCopyrightResponse {
	s.Headers = v
	return s
}

func (s *QueryCopyrightResponse) SetStatusCode(v int32) *QueryCopyrightResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCopyrightResponse) SetBody(v *QueryCopyrightResponseBody) *QueryCopyrightResponse {
	s.Body = v
	return s
}

type QueryCopyrightExtractRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s QueryCopyrightExtractRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCopyrightExtractRequest) GoString() string {
	return s.String()
}

func (s *QueryCopyrightExtractRequest) SetJobId(v string) *QueryCopyrightExtractRequest {
	s.JobId = &v
	return s
}

type QueryCopyrightExtractResponseBody struct {
	Data       *QueryCopyrightExtractResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestID  *string                                `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	StatusCode *int64                                 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QueryCopyrightExtractResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCopyrightExtractResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCopyrightExtractResponseBody) SetData(v *QueryCopyrightExtractResponseBodyData) *QueryCopyrightExtractResponseBody {
	s.Data = v
	return s
}

func (s *QueryCopyrightExtractResponseBody) SetMessage(v string) *QueryCopyrightExtractResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCopyrightExtractResponseBody) SetRequestID(v string) *QueryCopyrightExtractResponseBody {
	s.RequestID = &v
	return s
}

func (s *QueryCopyrightExtractResponseBody) SetStatusCode(v int64) *QueryCopyrightExtractResponseBody {
	s.StatusCode = &v
	return s
}

type QueryCopyrightExtractResponseBodyData struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s QueryCopyrightExtractResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryCopyrightExtractResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCopyrightExtractResponseBodyData) SetMessage(v string) *QueryCopyrightExtractResponseBodyData {
	s.Message = &v
	return s
}

type QueryCopyrightExtractResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCopyrightExtractResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCopyrightExtractResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCopyrightExtractResponse) GoString() string {
	return s.String()
}

func (s *QueryCopyrightExtractResponse) SetHeaders(v map[string]*string) *QueryCopyrightExtractResponse {
	s.Headers = v
	return s
}

func (s *QueryCopyrightExtractResponse) SetStatusCode(v int32) *QueryCopyrightExtractResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCopyrightExtractResponse) SetBody(v *QueryCopyrightExtractResponseBody) *QueryCopyrightExtractResponse {
	s.Body = v
	return s
}

type QueryTraceAbRequest struct {
	JobId   *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s QueryTraceAbRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTraceAbRequest) GoString() string {
	return s.String()
}

func (s *QueryTraceAbRequest) SetJobId(v string) *QueryTraceAbRequest {
	s.JobId = &v
	return s
}

func (s *QueryTraceAbRequest) SetMediaId(v string) *QueryTraceAbRequest {
	s.MediaId = &v
	return s
}

type QueryTraceAbResponseBody struct {
	Data       []*QueryTraceAbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestID  *string                         `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	StatusCode *int64                          `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QueryTraceAbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTraceAbResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTraceAbResponseBody) SetData(v []*QueryTraceAbResponseBodyData) *QueryTraceAbResponseBody {
	s.Data = v
	return s
}

func (s *QueryTraceAbResponseBody) SetMessage(v string) *QueryTraceAbResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTraceAbResponseBody) SetRequestID(v string) *QueryTraceAbResponseBody {
	s.RequestID = &v
	return s
}

func (s *QueryTraceAbResponseBody) SetStatusCode(v int64) *QueryTraceAbResponseBody {
	s.StatusCode = &v
	return s
}

type QueryTraceAbResponseBodyData struct {
	Callback    *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Input       *string `json:"Input,omitempty" xml:"Input,omitempty"`
	JobId       *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Level       *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	MediaId     *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Output      *string `json:"Output,omitempty" xml:"Output,omitempty"`
	Result      *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserData    *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	UserId      *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryTraceAbResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTraceAbResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTraceAbResponseBodyData) SetCallback(v string) *QueryTraceAbResponseBodyData {
	s.Callback = &v
	return s
}

func (s *QueryTraceAbResponseBodyData) SetGmtCreate(v int64) *QueryTraceAbResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryTraceAbResponseBodyData) SetGmtModified(v int64) *QueryTraceAbResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *QueryTraceAbResponseBodyData) SetInput(v string) *QueryTraceAbResponseBodyData {
	s.Input = &v
	return s
}

func (s *QueryTraceAbResponseBodyData) SetJobId(v string) *QueryTraceAbResponseBodyData {
	s.JobId = &v
	return s
}

func (s *QueryTraceAbResponseBodyData) SetLevel(v int64) *QueryTraceAbResponseBodyData {
	s.Level = &v
	return s
}

func (s *QueryTraceAbResponseBodyData) SetMediaId(v string) *QueryTraceAbResponseBodyData {
	s.MediaId = &v
	return s
}

func (s *QueryTraceAbResponseBodyData) SetOutput(v string) *QueryTraceAbResponseBodyData {
	s.Output = &v
	return s
}

func (s *QueryTraceAbResponseBodyData) SetResult(v string) *QueryTraceAbResponseBodyData {
	s.Result = &v
	return s
}

func (s *QueryTraceAbResponseBodyData) SetStatus(v string) *QueryTraceAbResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryTraceAbResponseBodyData) SetUserData(v string) *QueryTraceAbResponseBodyData {
	s.UserData = &v
	return s
}

func (s *QueryTraceAbResponseBodyData) SetUserId(v int64) *QueryTraceAbResponseBodyData {
	s.UserId = &v
	return s
}

type QueryTraceAbResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTraceAbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTraceAbResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTraceAbResponse) GoString() string {
	return s.String()
}

func (s *QueryTraceAbResponse) SetHeaders(v map[string]*string) *QueryTraceAbResponse {
	s.Headers = v
	return s
}

func (s *QueryTraceAbResponse) SetStatusCode(v int32) *QueryTraceAbResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTraceAbResponse) SetBody(v *QueryTraceAbResponseBody) *QueryTraceAbResponse {
	s.Body = v
	return s
}

type QueryTraceExtractRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s QueryTraceExtractRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTraceExtractRequest) GoString() string {
	return s.String()
}

func (s *QueryTraceExtractRequest) SetJobId(v string) *QueryTraceExtractRequest {
	s.JobId = &v
	return s
}

type QueryTraceExtractResponseBody struct {
	Data       *QueryTraceExtractResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestID  *string                            `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	StatusCode *int64                             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QueryTraceExtractResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTraceExtractResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTraceExtractResponseBody) SetData(v *QueryTraceExtractResponseBodyData) *QueryTraceExtractResponseBody {
	s.Data = v
	return s
}

func (s *QueryTraceExtractResponseBody) SetMessage(v string) *QueryTraceExtractResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTraceExtractResponseBody) SetRequestID(v string) *QueryTraceExtractResponseBody {
	s.RequestID = &v
	return s
}

func (s *QueryTraceExtractResponseBody) SetStatusCode(v int64) *QueryTraceExtractResponseBody {
	s.StatusCode = &v
	return s
}

type QueryTraceExtractResponseBodyData struct {
	Trace *string `json:"Trace,omitempty" xml:"Trace,omitempty"`
}

func (s QueryTraceExtractResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTraceExtractResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTraceExtractResponseBodyData) SetTrace(v string) *QueryTraceExtractResponseBodyData {
	s.Trace = &v
	return s
}

type QueryTraceExtractResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTraceExtractResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTraceExtractResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTraceExtractResponse) GoString() string {
	return s.String()
}

func (s *QueryTraceExtractResponse) SetHeaders(v map[string]*string) *QueryTraceExtractResponse {
	s.Headers = v
	return s
}

func (s *QueryTraceExtractResponse) SetStatusCode(v int32) *QueryTraceExtractResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTraceExtractResponse) SetBody(v *QueryTraceExtractResponseBody) *QueryTraceExtractResponse {
	s.Body = v
	return s
}

type QueryTraceMuRequest struct {
	CreateTimeEnd   *int64  `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	CreateTimeStart *int64  `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	JobId           *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Level           *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	MessageId       *int64  `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	PageNumber      *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryTraceMuRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTraceMuRequest) GoString() string {
	return s.String()
}

func (s *QueryTraceMuRequest) SetCreateTimeEnd(v int64) *QueryTraceMuRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *QueryTraceMuRequest) SetCreateTimeStart(v int64) *QueryTraceMuRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *QueryTraceMuRequest) SetJobId(v string) *QueryTraceMuRequest {
	s.JobId = &v
	return s
}

func (s *QueryTraceMuRequest) SetLevel(v int64) *QueryTraceMuRequest {
	s.Level = &v
	return s
}

func (s *QueryTraceMuRequest) SetMessageId(v int64) *QueryTraceMuRequest {
	s.MessageId = &v
	return s
}

func (s *QueryTraceMuRequest) SetPageNumber(v int64) *QueryTraceMuRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryTraceMuRequest) SetPageSize(v int64) *QueryTraceMuRequest {
	s.PageSize = &v
	return s
}

type QueryTraceMuResponseBody struct {
	Data       []*QueryTraceMuResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestID  *string                         `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	StatusCode *int64                          `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QueryTraceMuResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTraceMuResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTraceMuResponseBody) SetData(v []*QueryTraceMuResponseBodyData) *QueryTraceMuResponseBody {
	s.Data = v
	return s
}

func (s *QueryTraceMuResponseBody) SetMessage(v string) *QueryTraceMuResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTraceMuResponseBody) SetRequestID(v string) *QueryTraceMuResponseBody {
	s.RequestID = &v
	return s
}

func (s *QueryTraceMuResponseBody) SetStatusCode(v int64) *QueryTraceMuResponseBody {
	s.StatusCode = &v
	return s
}

type QueryTraceMuResponseBodyData struct {
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	JobId       *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId     *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Output      *string `json:"Output,omitempty" xml:"Output,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Trace       *string `json:"Trace,omitempty" xml:"Trace,omitempty"`
	TraceId     *int64  `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	UserData    *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	UserId      *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryTraceMuResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTraceMuResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTraceMuResponseBodyData) SetGmtCreate(v int64) *QueryTraceMuResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryTraceMuResponseBodyData) SetGmtModified(v int64) *QueryTraceMuResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *QueryTraceMuResponseBodyData) SetJobId(v string) *QueryTraceMuResponseBodyData {
	s.JobId = &v
	return s
}

func (s *QueryTraceMuResponseBodyData) SetMediaId(v string) *QueryTraceMuResponseBodyData {
	s.MediaId = &v
	return s
}

func (s *QueryTraceMuResponseBodyData) SetOutput(v string) *QueryTraceMuResponseBodyData {
	s.Output = &v
	return s
}

func (s *QueryTraceMuResponseBodyData) SetStatus(v string) *QueryTraceMuResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryTraceMuResponseBodyData) SetTrace(v string) *QueryTraceMuResponseBodyData {
	s.Trace = &v
	return s
}

func (s *QueryTraceMuResponseBodyData) SetTraceId(v int64) *QueryTraceMuResponseBodyData {
	s.TraceId = &v
	return s
}

func (s *QueryTraceMuResponseBodyData) SetUserData(v string) *QueryTraceMuResponseBodyData {
	s.UserData = &v
	return s
}

func (s *QueryTraceMuResponseBodyData) SetUserId(v int64) *QueryTraceMuResponseBodyData {
	s.UserId = &v
	return s
}

type QueryTraceMuResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTraceMuResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTraceMuResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTraceMuResponse) GoString() string {
	return s.String()
}

func (s *QueryTraceMuResponse) SetHeaders(v map[string]*string) *QueryTraceMuResponse {
	s.Headers = v
	return s
}

func (s *QueryTraceMuResponse) SetStatusCode(v int32) *QueryTraceMuResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTraceMuResponse) SetBody(v *QueryTraceMuResponseBody) *QueryTraceMuResponse {
	s.Body = v
	return s
}

type SubmitCopyrightExtractRequest struct {
	CallBack *string `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	Input    *string `json:"Input,omitempty" xml:"Input,omitempty"`
	Params   *string `json:"Params,omitempty" xml:"Params,omitempty"`
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitCopyrightExtractRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitCopyrightExtractRequest) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightExtractRequest) SetCallBack(v string) *SubmitCopyrightExtractRequest {
	s.CallBack = &v
	return s
}

func (s *SubmitCopyrightExtractRequest) SetInput(v string) *SubmitCopyrightExtractRequest {
	s.Input = &v
	return s
}

func (s *SubmitCopyrightExtractRequest) SetParams(v string) *SubmitCopyrightExtractRequest {
	s.Params = &v
	return s
}

func (s *SubmitCopyrightExtractRequest) SetUrl(v string) *SubmitCopyrightExtractRequest {
	s.Url = &v
	return s
}

func (s *SubmitCopyrightExtractRequest) SetUserData(v string) *SubmitCopyrightExtractRequest {
	s.UserData = &v
	return s
}

type SubmitCopyrightExtractResponseBody struct {
	Data       *SubmitCopyrightExtractResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestID  *string                                 `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	StatusCode *int64                                  `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s SubmitCopyrightExtractResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitCopyrightExtractResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightExtractResponseBody) SetData(v *SubmitCopyrightExtractResponseBodyData) *SubmitCopyrightExtractResponseBody {
	s.Data = v
	return s
}

func (s *SubmitCopyrightExtractResponseBody) SetMessage(v string) *SubmitCopyrightExtractResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitCopyrightExtractResponseBody) SetRequestID(v string) *SubmitCopyrightExtractResponseBody {
	s.RequestID = &v
	return s
}

func (s *SubmitCopyrightExtractResponseBody) SetStatusCode(v int64) *SubmitCopyrightExtractResponseBody {
	s.StatusCode = &v
	return s
}

type SubmitCopyrightExtractResponseBodyData struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitCopyrightExtractResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitCopyrightExtractResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightExtractResponseBodyData) SetJobId(v string) *SubmitCopyrightExtractResponseBodyData {
	s.JobId = &v
	return s
}

type SubmitCopyrightExtractResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitCopyrightExtractResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitCopyrightExtractResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitCopyrightExtractResponse) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightExtractResponse) SetHeaders(v map[string]*string) *SubmitCopyrightExtractResponse {
	s.Headers = v
	return s
}

func (s *SubmitCopyrightExtractResponse) SetStatusCode(v int32) *SubmitCopyrightExtractResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitCopyrightExtractResponse) SetBody(v *SubmitCopyrightExtractResponseBody) *SubmitCopyrightExtractResponse {
	s.Body = v
	return s
}

type SubmitCopyrightJobRequest struct {
	CallBack       *string `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Input          *string `json:"Input,omitempty" xml:"Input,omitempty"`
	Level          *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Output         *string `json:"Output,omitempty" xml:"Output,omitempty"`
	Params         *string `json:"Params,omitempty" xml:"Params,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TotalTime      *int64  `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	Url            *string `json:"Url,omitempty" xml:"Url,omitempty"`
	UserData       *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	VisibleMessage *string `json:"VisibleMessage,omitempty" xml:"VisibleMessage,omitempty"`
}

func (s SubmitCopyrightJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitCopyrightJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightJobRequest) SetCallBack(v string) *SubmitCopyrightJobRequest {
	s.CallBack = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetDescription(v string) *SubmitCopyrightJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetInput(v string) *SubmitCopyrightJobRequest {
	s.Input = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetLevel(v int64) *SubmitCopyrightJobRequest {
	s.Level = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetMessage(v string) *SubmitCopyrightJobRequest {
	s.Message = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetOutput(v string) *SubmitCopyrightJobRequest {
	s.Output = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetParams(v string) *SubmitCopyrightJobRequest {
	s.Params = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetStartTime(v int64) *SubmitCopyrightJobRequest {
	s.StartTime = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetTotalTime(v int64) *SubmitCopyrightJobRequest {
	s.TotalTime = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetUrl(v string) *SubmitCopyrightJobRequest {
	s.Url = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetUserData(v string) *SubmitCopyrightJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetVisibleMessage(v string) *SubmitCopyrightJobRequest {
	s.VisibleMessage = &v
	return s
}

type SubmitCopyrightJobResponseBody struct {
	Data       *SubmitCopyrightJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestID  *string                             `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	StatusCode *int64                              `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s SubmitCopyrightJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitCopyrightJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightJobResponseBody) SetData(v *SubmitCopyrightJobResponseBodyData) *SubmitCopyrightJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitCopyrightJobResponseBody) SetMessage(v string) *SubmitCopyrightJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitCopyrightJobResponseBody) SetRequestID(v string) *SubmitCopyrightJobResponseBody {
	s.RequestID = &v
	return s
}

func (s *SubmitCopyrightJobResponseBody) SetStatusCode(v int64) *SubmitCopyrightJobResponseBody {
	s.StatusCode = &v
	return s
}

type SubmitCopyrightJobResponseBodyData struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitCopyrightJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitCopyrightJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightJobResponseBodyData) SetJobId(v string) *SubmitCopyrightJobResponseBodyData {
	s.JobId = &v
	return s
}

type SubmitCopyrightJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitCopyrightJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitCopyrightJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitCopyrightJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightJobResponse) SetHeaders(v map[string]*string) *SubmitCopyrightJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitCopyrightJobResponse) SetStatusCode(v int32) *SubmitCopyrightJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitCopyrightJobResponse) SetBody(v *SubmitCopyrightJobResponseBody) *SubmitCopyrightJobResponse {
	s.Body = v
	return s
}

type SubmitImageCopyrightRequest struct {
	Level   *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Output  *string `json:"Output,omitempty" xml:"Output,omitempty"`
	Params  *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s SubmitImageCopyrightRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitImageCopyrightRequest) GoString() string {
	return s.String()
}

func (s *SubmitImageCopyrightRequest) SetLevel(v int64) *SubmitImageCopyrightRequest {
	s.Level = &v
	return s
}

func (s *SubmitImageCopyrightRequest) SetMessage(v string) *SubmitImageCopyrightRequest {
	s.Message = &v
	return s
}

func (s *SubmitImageCopyrightRequest) SetOutput(v string) *SubmitImageCopyrightRequest {
	s.Output = &v
	return s
}

func (s *SubmitImageCopyrightRequest) SetParams(v string) *SubmitImageCopyrightRequest {
	s.Params = &v
	return s
}

type SubmitImageCopyrightResponseBody struct {
	Data       *SubmitImageCopyrightResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestID  *string                               `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	StatusCode *int64                                `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s SubmitImageCopyrightResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitImageCopyrightResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitImageCopyrightResponseBody) SetData(v *SubmitImageCopyrightResponseBodyData) *SubmitImageCopyrightResponseBody {
	s.Data = v
	return s
}

func (s *SubmitImageCopyrightResponseBody) SetMessage(v string) *SubmitImageCopyrightResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitImageCopyrightResponseBody) SetRequestID(v string) *SubmitImageCopyrightResponseBody {
	s.RequestID = &v
	return s
}

func (s *SubmitImageCopyrightResponseBody) SetStatusCode(v int64) *SubmitImageCopyrightResponseBody {
	s.StatusCode = &v
	return s
}

type SubmitImageCopyrightResponseBodyData struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitImageCopyrightResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitImageCopyrightResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitImageCopyrightResponseBodyData) SetJobId(v string) *SubmitImageCopyrightResponseBodyData {
	s.JobId = &v
	return s
}

type SubmitImageCopyrightResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitImageCopyrightResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitImageCopyrightResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitImageCopyrightResponse) GoString() string {
	return s.String()
}

func (s *SubmitImageCopyrightResponse) SetHeaders(v map[string]*string) *SubmitImageCopyrightResponse {
	s.Headers = v
	return s
}

func (s *SubmitImageCopyrightResponse) SetStatusCode(v int32) *SubmitImageCopyrightResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitImageCopyrightResponse) SetBody(v *SubmitImageCopyrightResponseBody) *SubmitImageCopyrightResponse {
	s.Body = v
	return s
}

type SubmitTraceAbRequest struct {
	CallBack       *string `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	CipherBase64ed *string `json:"CipherBase64ed,omitempty" xml:"CipherBase64ed,omitempty"`
	Input          *string `json:"Input,omitempty" xml:"Input,omitempty"`
	Level          *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	Output         *string `json:"Output,omitempty" xml:"Output,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TotalTime      *int64  `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	Url            *string `json:"Url,omitempty" xml:"Url,omitempty"`
	UserData       *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitTraceAbRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitTraceAbRequest) GoString() string {
	return s.String()
}

func (s *SubmitTraceAbRequest) SetCallBack(v string) *SubmitTraceAbRequest {
	s.CallBack = &v
	return s
}

func (s *SubmitTraceAbRequest) SetCipherBase64ed(v string) *SubmitTraceAbRequest {
	s.CipherBase64ed = &v
	return s
}

func (s *SubmitTraceAbRequest) SetInput(v string) *SubmitTraceAbRequest {
	s.Input = &v
	return s
}

func (s *SubmitTraceAbRequest) SetLevel(v int64) *SubmitTraceAbRequest {
	s.Level = &v
	return s
}

func (s *SubmitTraceAbRequest) SetOutput(v string) *SubmitTraceAbRequest {
	s.Output = &v
	return s
}

func (s *SubmitTraceAbRequest) SetStartTime(v int64) *SubmitTraceAbRequest {
	s.StartTime = &v
	return s
}

func (s *SubmitTraceAbRequest) SetTotalTime(v int64) *SubmitTraceAbRequest {
	s.TotalTime = &v
	return s
}

func (s *SubmitTraceAbRequest) SetUrl(v string) *SubmitTraceAbRequest {
	s.Url = &v
	return s
}

func (s *SubmitTraceAbRequest) SetUserData(v string) *SubmitTraceAbRequest {
	s.UserData = &v
	return s
}

type SubmitTraceAbResponseBody struct {
	Data       *SubmitTraceAbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestID  *string                        `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	StatusCode *int64                         `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s SubmitTraceAbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitTraceAbResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTraceAbResponseBody) SetData(v *SubmitTraceAbResponseBodyData) *SubmitTraceAbResponseBody {
	s.Data = v
	return s
}

func (s *SubmitTraceAbResponseBody) SetMessage(v string) *SubmitTraceAbResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitTraceAbResponseBody) SetRequestID(v string) *SubmitTraceAbResponseBody {
	s.RequestID = &v
	return s
}

func (s *SubmitTraceAbResponseBody) SetStatusCode(v int64) *SubmitTraceAbResponseBody {
	s.StatusCode = &v
	return s
}

type SubmitTraceAbResponseBodyData struct {
	JobId   *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s SubmitTraceAbResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitTraceAbResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitTraceAbResponseBodyData) SetJobId(v string) *SubmitTraceAbResponseBodyData {
	s.JobId = &v
	return s
}

func (s *SubmitTraceAbResponseBodyData) SetMediaId(v string) *SubmitTraceAbResponseBodyData {
	s.MediaId = &v
	return s
}

type SubmitTraceAbResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitTraceAbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitTraceAbResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitTraceAbResponse) GoString() string {
	return s.String()
}

func (s *SubmitTraceAbResponse) SetHeaders(v map[string]*string) *SubmitTraceAbResponse {
	s.Headers = v
	return s
}

func (s *SubmitTraceAbResponse) SetStatusCode(v int32) *SubmitTraceAbResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitTraceAbResponse) SetBody(v *SubmitTraceAbResponseBody) *SubmitTraceAbResponse {
	s.Body = v
	return s
}

type SubmitTraceExtractRequest struct {
	CallBack *string `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	Input    *string `json:"Input,omitempty" xml:"Input,omitempty"`
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitTraceExtractRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitTraceExtractRequest) GoString() string {
	return s.String()
}

func (s *SubmitTraceExtractRequest) SetCallBack(v string) *SubmitTraceExtractRequest {
	s.CallBack = &v
	return s
}

func (s *SubmitTraceExtractRequest) SetInput(v string) *SubmitTraceExtractRequest {
	s.Input = &v
	return s
}

func (s *SubmitTraceExtractRequest) SetUrl(v string) *SubmitTraceExtractRequest {
	s.Url = &v
	return s
}

func (s *SubmitTraceExtractRequest) SetUserData(v string) *SubmitTraceExtractRequest {
	s.UserData = &v
	return s
}

type SubmitTraceExtractResponseBody struct {
	Data       *SubmitTraceExtractResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestID  *string                             `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	StatusCode *int64                              `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s SubmitTraceExtractResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitTraceExtractResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTraceExtractResponseBody) SetData(v *SubmitTraceExtractResponseBodyData) *SubmitTraceExtractResponseBody {
	s.Data = v
	return s
}

func (s *SubmitTraceExtractResponseBody) SetMessage(v string) *SubmitTraceExtractResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitTraceExtractResponseBody) SetRequestID(v string) *SubmitTraceExtractResponseBody {
	s.RequestID = &v
	return s
}

func (s *SubmitTraceExtractResponseBody) SetStatusCode(v int64) *SubmitTraceExtractResponseBody {
	s.StatusCode = &v
	return s
}

type SubmitTraceExtractResponseBodyData struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitTraceExtractResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitTraceExtractResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitTraceExtractResponseBodyData) SetJobId(v string) *SubmitTraceExtractResponseBodyData {
	s.JobId = &v
	return s
}

type SubmitTraceExtractResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitTraceExtractResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitTraceExtractResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitTraceExtractResponse) GoString() string {
	return s.String()
}

func (s *SubmitTraceExtractResponse) SetHeaders(v map[string]*string) *SubmitTraceExtractResponse {
	s.Headers = v
	return s
}

func (s *SubmitTraceExtractResponse) SetStatusCode(v int32) *SubmitTraceExtractResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitTraceExtractResponse) SetBody(v *SubmitTraceExtractResponseBody) *SubmitTraceExtractResponse {
	s.Body = v
	return s
}

type SubmitTracemuRequest struct {
	KeyUri  *string `json:"KeyUri,omitempty" xml:"KeyUri,omitempty"`
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Output  *string `json:"Output,omitempty" xml:"Output,omitempty"`
	Trace   *string `json:"Trace,omitempty" xml:"Trace,omitempty"`
}

func (s SubmitTracemuRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitTracemuRequest) GoString() string {
	return s.String()
}

func (s *SubmitTracemuRequest) SetKeyUri(v string) *SubmitTracemuRequest {
	s.KeyUri = &v
	return s
}

func (s *SubmitTracemuRequest) SetMediaId(v string) *SubmitTracemuRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitTracemuRequest) SetOutput(v string) *SubmitTracemuRequest {
	s.Output = &v
	return s
}

func (s *SubmitTracemuRequest) SetTrace(v string) *SubmitTracemuRequest {
	s.Trace = &v
	return s
}

type SubmitTracemuResponseBody struct {
	Data       *SubmitTracemuResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestID  *string                        `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	StatusCode *int64                         `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s SubmitTracemuResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitTracemuResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTracemuResponseBody) SetData(v *SubmitTracemuResponseBodyData) *SubmitTracemuResponseBody {
	s.Data = v
	return s
}

func (s *SubmitTracemuResponseBody) SetMessage(v string) *SubmitTracemuResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitTracemuResponseBody) SetRequestID(v string) *SubmitTracemuResponseBody {
	s.RequestID = &v
	return s
}

func (s *SubmitTracemuResponseBody) SetStatusCode(v int64) *SubmitTracemuResponseBody {
	s.StatusCode = &v
	return s
}

type SubmitTracemuResponseBodyData struct {
	Code  *string `json:"Code,omitempty" xml:"Code,omitempty"`
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitTracemuResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitTracemuResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitTracemuResponseBodyData) SetCode(v string) *SubmitTracemuResponseBodyData {
	s.Code = &v
	return s
}

func (s *SubmitTracemuResponseBodyData) SetJobId(v string) *SubmitTracemuResponseBodyData {
	s.JobId = &v
	return s
}

type SubmitTracemuResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitTracemuResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitTracemuResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitTracemuResponse) GoString() string {
	return s.String()
}

func (s *SubmitTracemuResponse) SetHeaders(v map[string]*string) *SubmitTracemuResponse {
	s.Headers = v
	return s
}

func (s *SubmitTracemuResponse) SetStatusCode(v int32) *SubmitTracemuResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitTracemuResponse) SetBody(v *SubmitTracemuResponseBody) *SubmitTracemuResponse {
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
		"ap-northeast-2-pop":          tea.String("mts.aliyuncs.com"),
		"ap-southeast-2":              tea.String("mts.aliyuncs.com"),
		"ap-southeast-3":              tea.String("mts.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("mts.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("mts.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("mts.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("mts.aliyuncs.com"),
		"cn-chengdu":                  tea.String("mts.aliyuncs.com"),
		"cn-edge-1":                   tea.String("mts.aliyuncs.com"),
		"cn-fujian":                   tea.String("mts.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("mts.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("mts.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("mts.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("mts.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("mts.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("mts.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("mts.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("mts.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("mts.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("mts.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("mts.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("mts.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("mts.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("mts.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("mts.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("mts.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("mts.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("mts.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("mts.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("mts.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("mts.aliyuncs.com"),
		"cn-wuhan":                    tea.String("mts.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("mts.aliyuncs.com"),
		"cn-yushanfang":               tea.String("mts.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("mts.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("mts.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("mts.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("mts.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("mts.aliyuncs.com"),
		"me-east-1":                   tea.String("mts.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("mts.aliyuncs.com"),
		"us-east-1":                   tea.String("mts.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("mts"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) QueryCopyrightWithOptions(request *QueryCopyrightRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryCopyrightResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateTimeEnd)) {
		body["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeStart)) {
		body["CreateTimeStart"] = request.CreateTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCopyright"),
		Version:     tea.String("2021-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/queryCopyrightJob"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCopyrightResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCopyright(request *QueryCopyrightRequest) (_result *QueryCopyrightResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryCopyrightResponse{}
	_body, _err := client.QueryCopyrightWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCopyrightExtractWithOptions(request *QueryCopyrightExtractRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryCopyrightExtractResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCopyrightExtract"),
		Version:     tea.String("2021-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/queryCopyrightExtract"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCopyrightExtractResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCopyrightExtract(request *QueryCopyrightExtractRequest) (_result *QueryCopyrightExtractResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryCopyrightExtractResponse{}
	_body, _err := client.QueryCopyrightExtractWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTraceAbWithOptions(request *QueryTraceAbRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryTraceAbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		body["MediaId"] = request.MediaId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTraceAb"),
		Version:     tea.String("2021-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/queryTraceAb"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTraceAbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTraceAb(request *QueryTraceAbRequest) (_result *QueryTraceAbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryTraceAbResponse{}
	_body, _err := client.QueryTraceAbWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTraceExtractWithOptions(request *QueryTraceExtractRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryTraceExtractResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTraceExtract"),
		Version:     tea.String("2021-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/queryTraceExtract"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTraceExtractResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTraceExtract(request *QueryTraceExtractRequest) (_result *QueryTraceExtractResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryTraceExtractResponse{}
	_body, _err := client.QueryTraceExtractWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTraceMuWithOptions(request *QueryTraceMuRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryTraceMuResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateTimeEnd)) {
		body["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeStart)) {
		body["CreateTimeStart"] = request.CreateTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		body["MessageId"] = request.MessageId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTraceMu"),
		Version:     tea.String("2021-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/queryTraceM3u8"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTraceMuResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTraceMu(request *QueryTraceMuRequest) (_result *QueryTraceMuResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryTraceMuResponse{}
	_body, _err := client.QueryTraceMuWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitCopyrightExtractWithOptions(request *SubmitCopyrightExtractRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitCopyrightExtractResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallBack)) {
		body["CallBack"] = request.CallBack
	}

	if !tea.BoolValue(util.IsUnset(request.Input)) {
		body["Input"] = request.Input
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		body["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitCopyrightExtract"),
		Version:     tea.String("2021-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/submitCopyrightExtract"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitCopyrightExtractResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitCopyrightExtract(request *SubmitCopyrightExtractRequest) (_result *SubmitCopyrightExtractResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitCopyrightExtractResponse{}
	_body, _err := client.SubmitCopyrightExtractWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitCopyrightJobWithOptions(request *SubmitCopyrightJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitCopyrightJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallBack)) {
		body["CallBack"] = request.CallBack
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Input)) {
		body["Input"] = request.Input
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		body["Message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.Output)) {
		body["Output"] = request.Output
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TotalTime)) {
		body["TotalTime"] = request.TotalTime
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		body["UserData"] = request.UserData
	}

	if !tea.BoolValue(util.IsUnset(request.VisibleMessage)) {
		body["VisibleMessage"] = request.VisibleMessage
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitCopyrightJob"),
		Version:     tea.String("2021-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/submitCopyrightJob"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitCopyrightJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitCopyrightJob(request *SubmitCopyrightJobRequest) (_result *SubmitCopyrightJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitCopyrightJobResponse{}
	_body, _err := client.SubmitCopyrightJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitImageCopyrightWithOptions(request *SubmitImageCopyrightRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitImageCopyrightResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		body["Message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.Output)) {
		body["Output"] = request.Output
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["Params"] = request.Params
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitImageCopyright"),
		Version:     tea.String("2021-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/submitImageCopyright"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitImageCopyrightResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitImageCopyright(request *SubmitImageCopyrightRequest) (_result *SubmitImageCopyrightResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitImageCopyrightResponse{}
	_body, _err := client.SubmitImageCopyrightWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitTraceAbWithOptions(request *SubmitTraceAbRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitTraceAbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallBack)) {
		body["CallBack"] = request.CallBack
	}

	if !tea.BoolValue(util.IsUnset(request.CipherBase64ed)) {
		body["CipherBase64ed"] = request.CipherBase64ed
	}

	if !tea.BoolValue(util.IsUnset(request.Input)) {
		body["Input"] = request.Input
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.Output)) {
		body["Output"] = request.Output
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TotalTime)) {
		body["TotalTime"] = request.TotalTime
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		body["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitTraceAb"),
		Version:     tea.String("2021-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/submitTraceAb"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitTraceAbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitTraceAb(request *SubmitTraceAbRequest) (_result *SubmitTraceAbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitTraceAbResponse{}
	_body, _err := client.SubmitTraceAbWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitTraceExtractWithOptions(request *SubmitTraceExtractRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitTraceExtractResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallBack)) {
		body["CallBack"] = request.CallBack
	}

	if !tea.BoolValue(util.IsUnset(request.Input)) {
		body["Input"] = request.Input
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		body["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitTraceExtract"),
		Version:     tea.String("2021-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/submitTraceExtract"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitTraceExtractResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitTraceExtract(request *SubmitTraceExtractRequest) (_result *SubmitTraceExtractResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitTraceExtractResponse{}
	_body, _err := client.SubmitTraceExtractWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitTracemuWithOptions(request *SubmitTracemuRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitTracemuResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyUri)) {
		body["KeyUri"] = request.KeyUri
	}

	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		body["MediaId"] = request.MediaId
	}

	if !tea.BoolValue(util.IsUnset(request.Output)) {
		body["Output"] = request.Output
	}

	if !tea.BoolValue(util.IsUnset(request.Trace)) {
		body["Trace"] = request.Trace
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitTracemu"),
		Version:     tea.String("2021-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/submitTraceM3u8"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitTracemuResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitTracemu(request *SubmitTracemuRequest) (_result *SubmitTracemuResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitTracemuResponse{}
	_body, _err := client.SubmitTracemuWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
