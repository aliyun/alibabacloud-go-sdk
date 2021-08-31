// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryTraceMuRequest struct {
	// 创建时间起始
	CreateTimeEnd *int64 `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// 创建时间截止
	CreateTimeStart *int64 `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// 任务id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 水印强度
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// 水印信息id
	MessageId *int64 `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// 页偏移
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页数量
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// 返回数据结构
	Data []*QueryTraceMuResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// 返回信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求id
	RequestID *string `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	// 状态码
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
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
	// 创建时间
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 最后修改时间
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 任务id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 媒体id
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// 输出oss地址
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// 任务状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 溯源水印信息
	Trace *string `json:"Trace,omitempty" xml:"Trace,omitempty"`
	// 溯源水印信息id
	TraceId *int64 `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	// 用户自定义数据
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// uid
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTraceMuResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryTraceMuResponse) SetBody(v *QueryTraceMuResponseBody) *QueryTraceMuResponse {
	s.Body = v
	return s
}

type SubmitImageCopyrightRequest struct {
	// 需要加水印的图片oss地址(Input和url二选一)
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// 水印强度
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// 水印信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 水印图片输出oss地址
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// 外部url链接(Input和url二选一)
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SubmitImageCopyrightRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitImageCopyrightRequest) GoString() string {
	return s.String()
}

func (s *SubmitImageCopyrightRequest) SetInput(v string) *SubmitImageCopyrightRequest {
	s.Input = &v
	return s
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

func (s *SubmitImageCopyrightRequest) SetUrl(v string) *SubmitImageCopyrightRequest {
	s.Url = &v
	return s
}

type SubmitImageCopyrightResponseBody struct {
	// 返回数据
	Data *SubmitImageCopyrightResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 返回信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求id
	RequestID *string `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	// 状态码
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
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
	// 任务id
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitImageCopyrightResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SubmitImageCopyrightResponse) SetBody(v *SubmitImageCopyrightResponseBody) *SubmitImageCopyrightResponse {
	s.Body = v
	return s
}

type QueryImageCopyrightRequest struct {
	// 创建时间起始
	CreateTimeEnd *int64 `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// 创建时间截止
	CreateTimeStart *int64 `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// 任务ID
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 页偏移
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页数量
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryImageCopyrightRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryImageCopyrightRequest) GoString() string {
	return s.String()
}

func (s *QueryImageCopyrightRequest) SetCreateTimeEnd(v int64) *QueryImageCopyrightRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *QueryImageCopyrightRequest) SetCreateTimeStart(v int64) *QueryImageCopyrightRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *QueryImageCopyrightRequest) SetJobId(v string) *QueryImageCopyrightRequest {
	s.JobId = &v
	return s
}

func (s *QueryImageCopyrightRequest) SetPageNumber(v int64) *QueryImageCopyrightRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryImageCopyrightRequest) SetPageSize(v int64) *QueryImageCopyrightRequest {
	s.PageSize = &v
	return s
}

type QueryImageCopyrightResponseBody struct {
	// 返回数据
	Data []*QueryImageCopyrightResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// 返回信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求id
	RequestID *string `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	// 状态码
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QueryImageCopyrightResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryImageCopyrightResponseBody) GoString() string {
	return s.String()
}

func (s *QueryImageCopyrightResponseBody) SetData(v []*QueryImageCopyrightResponseBodyData) *QueryImageCopyrightResponseBody {
	s.Data = v
	return s
}

func (s *QueryImageCopyrightResponseBody) SetMessage(v string) *QueryImageCopyrightResponseBody {
	s.Message = &v
	return s
}

func (s *QueryImageCopyrightResponseBody) SetRequestID(v string) *QueryImageCopyrightResponseBody {
	s.RequestID = &v
	return s
}

func (s *QueryImageCopyrightResponseBody) SetStatusCode(v int64) *QueryImageCopyrightResponseBody {
	s.StatusCode = &v
	return s
}

type QueryImageCopyrightResponseBodyData struct {
	// 创建时间
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 最后修改时间
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 水印图片输入oss地址
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// 任务id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 水印强度
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// 水印信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 水印信息id
	MessageId *int64 `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// 加完水印后的输出oss地址
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// 任务状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 用户自定义数据
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// uid
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryImageCopyrightResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryImageCopyrightResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryImageCopyrightResponseBodyData) SetGmtCreate(v int64) *QueryImageCopyrightResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryImageCopyrightResponseBodyData) SetGmtModified(v int64) *QueryImageCopyrightResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *QueryImageCopyrightResponseBodyData) SetInput(v string) *QueryImageCopyrightResponseBodyData {
	s.Input = &v
	return s
}

func (s *QueryImageCopyrightResponseBodyData) SetJobId(v string) *QueryImageCopyrightResponseBodyData {
	s.JobId = &v
	return s
}

func (s *QueryImageCopyrightResponseBodyData) SetLevel(v int64) *QueryImageCopyrightResponseBodyData {
	s.Level = &v
	return s
}

func (s *QueryImageCopyrightResponseBodyData) SetMessage(v string) *QueryImageCopyrightResponseBodyData {
	s.Message = &v
	return s
}

func (s *QueryImageCopyrightResponseBodyData) SetMessageId(v int64) *QueryImageCopyrightResponseBodyData {
	s.MessageId = &v
	return s
}

func (s *QueryImageCopyrightResponseBodyData) SetOutput(v string) *QueryImageCopyrightResponseBodyData {
	s.Output = &v
	return s
}

func (s *QueryImageCopyrightResponseBodyData) SetStatus(v string) *QueryImageCopyrightResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryImageCopyrightResponseBodyData) SetUserData(v string) *QueryImageCopyrightResponseBodyData {
	s.UserData = &v
	return s
}

func (s *QueryImageCopyrightResponseBodyData) SetUserId(v int64) *QueryImageCopyrightResponseBodyData {
	s.UserId = &v
	return s
}

type QueryImageCopyrightResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryImageCopyrightResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryImageCopyrightResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryImageCopyrightResponse) GoString() string {
	return s.String()
}

func (s *QueryImageCopyrightResponse) SetHeaders(v map[string]*string) *QueryImageCopyrightResponse {
	s.Headers = v
	return s
}

func (s *QueryImageCopyrightResponse) SetBody(v *QueryImageCopyrightResponseBody) *QueryImageCopyrightResponse {
	s.Body = v
	return s
}

type QueryCopyrightRequest struct {
	// 创建时间截止
	CreateTimeEnd *int64 `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// 创建时间起始
	CreateTimeStart *int64 `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// 任务id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 水印强度
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// 翻页下标
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页数量
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	RequestID *string                           `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	Data      []*QueryCopyrightResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// 状态码
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QueryCopyrightResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCopyrightResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCopyrightResponseBody) SetRequestID(v string) *QueryCopyrightResponseBody {
	s.RequestID = &v
	return s
}

func (s *QueryCopyrightResponseBody) SetData(v []*QueryCopyrightResponseBodyData) *QueryCopyrightResponseBody {
	s.Data = v
	return s
}

func (s *QueryCopyrightResponseBody) SetStatusCode(v int64) *QueryCopyrightResponseBody {
	s.StatusCode = &v
	return s
}

type QueryCopyrightResponseBodyData struct {
	// 回调url
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// 创建时间
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 修改时间
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 水印视频输入
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// 任务id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 水印强度
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// 水印信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 水印信息id
	MessageId *int64 `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// 水印视频输出
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 用户数据
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// 用户ID
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCopyrightResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryCopyrightResponse) SetBody(v *QueryCopyrightResponseBody) *QueryCopyrightResponse {
	s.Body = v
	return s
}

type SubmitTracemuRequest struct {
	// ab流处理后的媒体id
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// m3u8文件输出oss地址
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// 溯源水印信息
	Trace *string `json:"Trace,omitempty" xml:"Trace,omitempty"`
}

func (s SubmitTracemuRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitTracemuRequest) GoString() string {
	return s.String()
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
	// 返回数据
	Data *SubmitTracemuResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 返回信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求Id
	RequestID *string `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	// 状态码
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
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
	// 任务id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitTracemuResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitTracemuResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitTracemuResponseBodyData) SetJobId(v string) *SubmitTracemuResponseBodyData {
	s.JobId = &v
	return s
}

type SubmitTracemuResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitTracemuResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SubmitTracemuResponse) SetBody(v *SubmitTracemuResponseBody) *SubmitTracemuResponse {
	s.Body = v
	return s
}

type QueryTraceAbRequest struct {
	// 任务id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 媒体id
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
	// 返回结构
	Data []*QueryTraceAbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// 返回信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求id
	RequestID *string `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	// 状态码
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
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
	// 任务结果回调
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// 创建时间
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 最后修改时间
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 输入oss地址
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// 任务id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 水印强度
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// 媒体id
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// 输出地址
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// 任务状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 用户数据
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// uid
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTraceAbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryTraceAbResponse) SetBody(v *QueryTraceAbResponseBody) *QueryTraceAbResponse {
	s.Body = v
	return s
}

type SubmitTraceAbRequest struct {
	// 任务结果回调
	CallBack *string `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	// 溯源水印ab流处理视频输入
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// 水印强度
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// 溯源水印ab流处理输出
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// 用户自定义数据，最大长度1024个字节
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// 外部url链接(Input和url二选一)
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

func (s *SubmitTraceAbRequest) SetUserData(v string) *SubmitTraceAbRequest {
	s.UserData = &v
	return s
}

func (s *SubmitTraceAbRequest) SetUrl(v string) *SubmitTraceAbRequest {
	s.Url = &v
	return s
}

type SubmitTraceAbResponseBody struct {
	// 返回数据
	Data *SubmitTraceAbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 返回信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求Id
	RequestID *string `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	// 状态码
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
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
	// 任务ID
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 媒体id
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitTraceAbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SubmitTraceAbResponse) SetBody(v *SubmitTraceAbResponseBody) *SubmitTraceAbResponse {
	s.Body = v
	return s
}

type SubmitCopyrightJobRequest struct {
	// 任务结果回调url
	CallBack *string `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	// 水印信息描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 输入的视频，oss三元组
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// 水印强度，取值1，2，3
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// 水印信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 水印起始时间(单位是秒)，不填写默认为0
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 水印结束时间(单位是秒)，不填默认为60000
	TotalTime *int64 `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	// 输出的视频，oss三元组
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// 用户自定义数据
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// 外部url链接(Input和url二选一)
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

func (s *SubmitCopyrightJobRequest) SetStartTime(v int64) *SubmitCopyrightJobRequest {
	s.StartTime = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetTotalTime(v int64) *SubmitCopyrightJobRequest {
	s.TotalTime = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetOutput(v string) *SubmitCopyrightJobRequest {
	s.Output = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetUserData(v string) *SubmitCopyrightJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetUrl(v string) *SubmitCopyrightJobRequest {
	s.Url = &v
	return s
}

type SubmitCopyrightJobResponseBody struct {
	// 请求Id
	RequestID *string `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	// 返回信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回数据
	Data *SubmitCopyrightJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 状态码
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s SubmitCopyrightJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitCopyrightJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightJobResponseBody) SetRequestID(v string) *SubmitCopyrightJobResponseBody {
	s.RequestID = &v
	return s
}

func (s *SubmitCopyrightJobResponseBody) SetMessage(v string) *SubmitCopyrightJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitCopyrightJobResponseBody) SetData(v *SubmitCopyrightJobResponseBodyData) *SubmitCopyrightJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitCopyrightJobResponseBody) SetStatusCode(v int64) *SubmitCopyrightJobResponseBody {
	s.StatusCode = &v
	return s
}

type SubmitCopyrightJobResponseBodyData struct {
	// 任务id
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitCopyrightJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SubmitCopyrightJobResponse) SetBody(v *SubmitCopyrightJobResponseBody) *SubmitCopyrightJobResponse {
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
	_result = &QueryTraceMuResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryTraceMu"), tea.String("2021-07-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/queryTraceM3u8"), tea.String("json"), req, runtime)
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

func (client *Client) SubmitImageCopyrightWithOptions(request *SubmitImageCopyrightRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitImageCopyrightResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &SubmitImageCopyrightResponse{}
	_body, _err := client.DoROARequest(tea.String("SubmitImageCopyright"), tea.String("2021-07-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/submitImageCopyright"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryImageCopyright(request *QueryImageCopyrightRequest) (_result *QueryImageCopyrightResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryImageCopyrightResponse{}
	_body, _err := client.QueryImageCopyrightWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryImageCopyrightWithOptions(request *QueryImageCopyrightRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryImageCopyrightResponse, _err error) {
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
	_result = &QueryImageCopyrightResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryImageCopyright"), tea.String("2021-07-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/queryImageCopyright"), tea.String("json"), req, runtime)
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
	_result = &QueryCopyrightResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCopyright"), tea.String("2021-07-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/queryCopyrightJob"), tea.String("json"), req, runtime)
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

func (client *Client) SubmitTracemuWithOptions(request *SubmitTracemuRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitTracemuResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
	_result = &SubmitTracemuResponse{}
	_body, _err := client.DoROARequest(tea.String("SubmitTracemu"), tea.String("2021-07-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/submitTraceM3u8"), tea.String("json"), req, runtime)
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
	_result = &QueryTraceAbResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryTraceAb"), tea.String("2021-07-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/queryTraceAb"), tea.String("json"), req, runtime)
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

func (client *Client) SubmitTraceAbWithOptions(request *SubmitTraceAbRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitTraceAbResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.Output)) {
		body["Output"] = request.Output
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		body["UserData"] = request.UserData
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &SubmitTraceAbResponse{}
	_body, _err := client.DoROARequest(tea.String("SubmitTraceAb"), tea.String("2021-07-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/submitTraceAb"), tea.String("json"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TotalTime)) {
		body["TotalTime"] = request.TotalTime
	}

	if !tea.BoolValue(util.IsUnset(request.Output)) {
		body["Output"] = request.Output
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		body["UserData"] = request.UserData
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &SubmitCopyrightJobResponse{}
	_body, _err := client.DoROARequest(tea.String("SubmitCopyrightJob"), tea.String("2021-07-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/submitCopyrightJob"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
