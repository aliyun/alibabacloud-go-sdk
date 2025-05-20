// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DelMessageRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BizName        *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	ClientSource   *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	Cookies        *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	MsgId          *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	SrcUrl         *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	TenantCode     *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	UidType        *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s DelMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s DelMessageRequest) GoString() string {
	return s.String()
}

func (s *DelMessageRequest) SetAcceptLanguage(v string) *DelMessageRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DelMessageRequest) SetAppName(v string) *DelMessageRequest {
	s.AppName = &v
	return s
}

func (s *DelMessageRequest) SetBizName(v string) *DelMessageRequest {
	s.BizName = &v
	return s
}

func (s *DelMessageRequest) SetCallerProtocol(v string) *DelMessageRequest {
	s.CallerProtocol = &v
	return s
}

func (s *DelMessageRequest) SetClientSource(v string) *DelMessageRequest {
	s.ClientSource = &v
	return s
}

func (s *DelMessageRequest) SetCookies(v string) *DelMessageRequest {
	s.Cookies = &v
	return s
}

func (s *DelMessageRequest) SetMsgId(v string) *DelMessageRequest {
	s.MsgId = &v
	return s
}

func (s *DelMessageRequest) SetSrcUrl(v string) *DelMessageRequest {
	s.SrcUrl = &v
	return s
}

func (s *DelMessageRequest) SetTenantCode(v string) *DelMessageRequest {
	s.TenantCode = &v
	return s
}

func (s *DelMessageRequest) SetUidType(v string) *DelMessageRequest {
	s.UidType = &v
	return s
}

type DelMessageResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DelMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DelMessageResponseBody) GoString() string {
	return s.String()
}

func (s *DelMessageResponseBody) SetCode(v string) *DelMessageResponseBody {
	s.Code = &v
	return s
}

func (s *DelMessageResponseBody) SetData(v bool) *DelMessageResponseBody {
	s.Data = &v
	return s
}

func (s *DelMessageResponseBody) SetMessage(v string) *DelMessageResponseBody {
	s.Message = &v
	return s
}

func (s *DelMessageResponseBody) SetRequestId(v string) *DelMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DelMessageResponseBody) SetSuccess(v bool) *DelMessageResponseBody {
	s.Success = &v
	return s
}

type DelMessageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DelMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DelMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s DelMessageResponse) GoString() string {
	return s.String()
}

func (s *DelMessageResponse) SetHeaders(v map[string]*string) *DelMessageResponse {
	s.Headers = v
	return s
}

func (s *DelMessageResponse) SetStatusCode(v int32) *DelMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *DelMessageResponse) SetBody(v *DelMessageResponseBody) *DelMessageResponse {
	s.Body = v
	return s
}

type DeleteAllMessageRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BizName        *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	ClassId        *int64  `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	ClientSource   *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	Cookies        *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	SrcUrl         *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	TenantCode     *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	UidType        *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s DeleteAllMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAllMessageRequest) GoString() string {
	return s.String()
}

func (s *DeleteAllMessageRequest) SetAcceptLanguage(v string) *DeleteAllMessageRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteAllMessageRequest) SetAppName(v string) *DeleteAllMessageRequest {
	s.AppName = &v
	return s
}

func (s *DeleteAllMessageRequest) SetBizName(v string) *DeleteAllMessageRequest {
	s.BizName = &v
	return s
}

func (s *DeleteAllMessageRequest) SetCallerProtocol(v string) *DeleteAllMessageRequest {
	s.CallerProtocol = &v
	return s
}

func (s *DeleteAllMessageRequest) SetClassId(v int64) *DeleteAllMessageRequest {
	s.ClassId = &v
	return s
}

func (s *DeleteAllMessageRequest) SetClientSource(v string) *DeleteAllMessageRequest {
	s.ClientSource = &v
	return s
}

func (s *DeleteAllMessageRequest) SetCookies(v string) *DeleteAllMessageRequest {
	s.Cookies = &v
	return s
}

func (s *DeleteAllMessageRequest) SetSrcUrl(v string) *DeleteAllMessageRequest {
	s.SrcUrl = &v
	return s
}

func (s *DeleteAllMessageRequest) SetTenantCode(v string) *DeleteAllMessageRequest {
	s.TenantCode = &v
	return s
}

func (s *DeleteAllMessageRequest) SetUidType(v string) *DeleteAllMessageRequest {
	s.UidType = &v
	return s
}

type DeleteAllMessageResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAllMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAllMessageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAllMessageResponseBody) SetCode(v string) *DeleteAllMessageResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAllMessageResponseBody) SetData(v bool) *DeleteAllMessageResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAllMessageResponseBody) SetMessage(v string) *DeleteAllMessageResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAllMessageResponseBody) SetRequestId(v string) *DeleteAllMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAllMessageResponseBody) SetSuccess(v bool) *DeleteAllMessageResponseBody {
	s.Success = &v
	return s
}

type DeleteAllMessageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAllMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAllMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAllMessageResponse) GoString() string {
	return s.String()
}

func (s *DeleteAllMessageResponse) SetHeaders(v map[string]*string) *DeleteAllMessageResponse {
	s.Headers = v
	return s
}

func (s *DeleteAllMessageResponse) SetStatusCode(v int32) *DeleteAllMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAllMessageResponse) SetBody(v *DeleteAllMessageResponseBody) *DeleteAllMessageResponse {
	s.Body = v
	return s
}

type ReadAllMessageRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BizName        *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	ClassId        *int64  `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	ClientSource   *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	Cookies        *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	SrcUrl         *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	TenantCode     *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	UidType        *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadAllMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s ReadAllMessageRequest) GoString() string {
	return s.String()
}

func (s *ReadAllMessageRequest) SetAcceptLanguage(v string) *ReadAllMessageRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadAllMessageRequest) SetAppName(v string) *ReadAllMessageRequest {
	s.AppName = &v
	return s
}

func (s *ReadAllMessageRequest) SetBizName(v string) *ReadAllMessageRequest {
	s.BizName = &v
	return s
}

func (s *ReadAllMessageRequest) SetCallerProtocol(v string) *ReadAllMessageRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadAllMessageRequest) SetClassId(v int64) *ReadAllMessageRequest {
	s.ClassId = &v
	return s
}

func (s *ReadAllMessageRequest) SetClientSource(v string) *ReadAllMessageRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadAllMessageRequest) SetCookies(v string) *ReadAllMessageRequest {
	s.Cookies = &v
	return s
}

func (s *ReadAllMessageRequest) SetSrcUrl(v string) *ReadAllMessageRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadAllMessageRequest) SetTenantCode(v string) *ReadAllMessageRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadAllMessageRequest) SetUidType(v string) *ReadAllMessageRequest {
	s.UidType = &v
	return s
}

type ReadAllMessageResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadAllMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReadAllMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ReadAllMessageResponseBody) SetCode(v string) *ReadAllMessageResponseBody {
	s.Code = &v
	return s
}

func (s *ReadAllMessageResponseBody) SetData(v bool) *ReadAllMessageResponseBody {
	s.Data = &v
	return s
}

func (s *ReadAllMessageResponseBody) SetMessage(v string) *ReadAllMessageResponseBody {
	s.Message = &v
	return s
}

func (s *ReadAllMessageResponseBody) SetRequestId(v string) *ReadAllMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadAllMessageResponseBody) SetSuccess(v bool) *ReadAllMessageResponseBody {
	s.Success = &v
	return s
}

type ReadAllMessageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadAllMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadAllMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s ReadAllMessageResponse) GoString() string {
	return s.String()
}

func (s *ReadAllMessageResponse) SetHeaders(v map[string]*string) *ReadAllMessageResponse {
	s.Headers = v
	return s
}

func (s *ReadAllMessageResponse) SetStatusCode(v int32) *ReadAllMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadAllMessageResponse) SetBody(v *ReadAllMessageResponseBody) *ReadAllMessageResponse {
	s.Body = v
	return s
}

type ReadClassNameRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BizName        *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	ClientSource   *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	Cookies        *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	SrcUrl         *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	TenantCode     *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	UidType        *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadClassNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ReadClassNameRequest) GoString() string {
	return s.String()
}

func (s *ReadClassNameRequest) SetAcceptLanguage(v string) *ReadClassNameRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadClassNameRequest) SetAppName(v string) *ReadClassNameRequest {
	s.AppName = &v
	return s
}

func (s *ReadClassNameRequest) SetBizName(v string) *ReadClassNameRequest {
	s.BizName = &v
	return s
}

func (s *ReadClassNameRequest) SetCallerProtocol(v string) *ReadClassNameRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadClassNameRequest) SetClientSource(v string) *ReadClassNameRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadClassNameRequest) SetCookies(v string) *ReadClassNameRequest {
	s.Cookies = &v
	return s
}

func (s *ReadClassNameRequest) SetSrcUrl(v string) *ReadClassNameRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadClassNameRequest) SetTenantCode(v string) *ReadClassNameRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadClassNameRequest) SetUidType(v string) *ReadClassNameRequest {
	s.UidType = &v
	return s
}

type ReadClassNameResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ReadClassNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadClassNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReadClassNameResponseBody) GoString() string {
	return s.String()
}

func (s *ReadClassNameResponseBody) SetCode(v string) *ReadClassNameResponseBody {
	s.Code = &v
	return s
}

func (s *ReadClassNameResponseBody) SetData(v []*ReadClassNameResponseBodyData) *ReadClassNameResponseBody {
	s.Data = v
	return s
}

func (s *ReadClassNameResponseBody) SetMessage(v string) *ReadClassNameResponseBody {
	s.Message = &v
	return s
}

func (s *ReadClassNameResponseBody) SetRequestId(v string) *ReadClassNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadClassNameResponseBody) SetSuccess(v bool) *ReadClassNameResponseBody {
	s.Success = &v
	return s
}

type ReadClassNameResponseBodyData struct {
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ReadClassNameResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ReadClassNameResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadClassNameResponseBodyData) SetId(v int64) *ReadClassNameResponseBodyData {
	s.Id = &v
	return s
}

func (s *ReadClassNameResponseBodyData) SetName(v string) *ReadClassNameResponseBodyData {
	s.Name = &v
	return s
}

type ReadClassNameResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadClassNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadClassNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ReadClassNameResponse) GoString() string {
	return s.String()
}

func (s *ReadClassNameResponse) SetHeaders(v map[string]*string) *ReadClassNameResponse {
	s.Headers = v
	return s
}

func (s *ReadClassNameResponse) SetStatusCode(v int32) *ReadClassNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadClassNameResponse) SetBody(v *ReadClassNameResponseBody) *ReadClassNameResponse {
	s.Body = v
	return s
}

type ReadMessageRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BizName        *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	ClientSource   *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	Cookies        *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	MsgId          *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	SrcUrl         *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	TenantCode     *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	UidType        *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageRequest) GoString() string {
	return s.String()
}

func (s *ReadMessageRequest) SetAcceptLanguage(v string) *ReadMessageRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadMessageRequest) SetAppName(v string) *ReadMessageRequest {
	s.AppName = &v
	return s
}

func (s *ReadMessageRequest) SetBizName(v string) *ReadMessageRequest {
	s.BizName = &v
	return s
}

func (s *ReadMessageRequest) SetCallerProtocol(v string) *ReadMessageRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadMessageRequest) SetClientSource(v string) *ReadMessageRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadMessageRequest) SetCookies(v string) *ReadMessageRequest {
	s.Cookies = &v
	return s
}

func (s *ReadMessageRequest) SetMsgId(v string) *ReadMessageRequest {
	s.MsgId = &v
	return s
}

func (s *ReadMessageRequest) SetSrcUrl(v string) *ReadMessageRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadMessageRequest) SetTenantCode(v string) *ReadMessageRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadMessageRequest) SetUidType(v string) *ReadMessageRequest {
	s.UidType = &v
	return s
}

type ReadMessageResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ReadMessageResponseBody) SetCode(v string) *ReadMessageResponseBody {
	s.Code = &v
	return s
}

func (s *ReadMessageResponseBody) SetData(v bool) *ReadMessageResponseBody {
	s.Data = &v
	return s
}

func (s *ReadMessageResponseBody) SetMessage(v string) *ReadMessageResponseBody {
	s.Message = &v
	return s
}

func (s *ReadMessageResponseBody) SetRequestId(v string) *ReadMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadMessageResponseBody) SetSuccess(v bool) *ReadMessageResponseBody {
	s.Success = &v
	return s
}

type ReadMessageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageResponse) GoString() string {
	return s.String()
}

func (s *ReadMessageResponse) SetHeaders(v map[string]*string) *ReadMessageResponse {
	s.Headers = v
	return s
}

func (s *ReadMessageResponse) SetStatusCode(v int32) *ReadMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadMessageResponse) SetBody(v *ReadMessageResponseBody) *ReadMessageResponse {
	s.Body = v
	return s
}

type ReadMessageContentRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BizName        *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	ClassId        *int64  `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	ClientSource   *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	Cookies        *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	MsgId          *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	SrcUrl         *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantCode     *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	UidType        *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadMessageContentRequest) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageContentRequest) GoString() string {
	return s.String()
}

func (s *ReadMessageContentRequest) SetAcceptLanguage(v string) *ReadMessageContentRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadMessageContentRequest) SetAppName(v string) *ReadMessageContentRequest {
	s.AppName = &v
	return s
}

func (s *ReadMessageContentRequest) SetBizName(v string) *ReadMessageContentRequest {
	s.BizName = &v
	return s
}

func (s *ReadMessageContentRequest) SetCallerProtocol(v string) *ReadMessageContentRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadMessageContentRequest) SetClassId(v int64) *ReadMessageContentRequest {
	s.ClassId = &v
	return s
}

func (s *ReadMessageContentRequest) SetClientSource(v string) *ReadMessageContentRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadMessageContentRequest) SetCookies(v string) *ReadMessageContentRequest {
	s.Cookies = &v
	return s
}

func (s *ReadMessageContentRequest) SetMsgId(v string) *ReadMessageContentRequest {
	s.MsgId = &v
	return s
}

func (s *ReadMessageContentRequest) SetSrcUrl(v string) *ReadMessageContentRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadMessageContentRequest) SetStatus(v int32) *ReadMessageContentRequest {
	s.Status = &v
	return s
}

func (s *ReadMessageContentRequest) SetTenantCode(v string) *ReadMessageContentRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadMessageContentRequest) SetUidType(v string) *ReadMessageContentRequest {
	s.UidType = &v
	return s
}

type ReadMessageContentResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ReadMessageContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadMessageContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageContentResponseBody) GoString() string {
	return s.String()
}

func (s *ReadMessageContentResponseBody) SetCode(v string) *ReadMessageContentResponseBody {
	s.Code = &v
	return s
}

func (s *ReadMessageContentResponseBody) SetData(v *ReadMessageContentResponseBodyData) *ReadMessageContentResponseBody {
	s.Data = v
	return s
}

func (s *ReadMessageContentResponseBody) SetMessage(v string) *ReadMessageContentResponseBody {
	s.Message = &v
	return s
}

func (s *ReadMessageContentResponseBody) SetRequestId(v string) *ReadMessageContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadMessageContentResponseBody) SetSuccess(v bool) *ReadMessageContentResponseBody {
	s.Success = &v
	return s
}

type ReadMessageContentResponseBodyData struct {
	Datas *ReadMessageContentResponseBodyDataDatas `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Struct"`
}

func (s ReadMessageContentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageContentResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadMessageContentResponseBodyData) SetDatas(v *ReadMessageContentResponseBodyDataDatas) *ReadMessageContentResponseBodyData {
	s.Datas = v
	return s
}

type ReadMessageContentResponseBodyDataDatas struct {
	Item     []*ReadMessageContentResponseBodyDataDatasItem     `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	LastItem []*ReadMessageContentResponseBodyDataDatasLastItem `json:"LastItem,omitempty" xml:"LastItem,omitempty" type:"Repeated"`
	NextItem []*ReadMessageContentResponseBodyDataDatasNextItem `json:"NextItem,omitempty" xml:"NextItem,omitempty" type:"Repeated"`
}

func (s ReadMessageContentResponseBodyDataDatas) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageContentResponseBodyDataDatas) GoString() string {
	return s.String()
}

func (s *ReadMessageContentResponseBodyDataDatas) SetItem(v []*ReadMessageContentResponseBodyDataDatasItem) *ReadMessageContentResponseBodyDataDatas {
	s.Item = v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatas) SetLastItem(v []*ReadMessageContentResponseBodyDataDatasLastItem) *ReadMessageContentResponseBodyDataDatas {
	s.LastItem = v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatas) SetNextItem(v []*ReadMessageContentResponseBodyDataDatasNextItem) *ReadMessageContentResponseBodyDataDatas {
	s.NextItem = v
	return s
}

type ReadMessageContentResponseBodyDataDatasItem struct {
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	ClassId      *int64  `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Deleted      *int32  `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	GmtCreated   *int64  `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtUpdate    *int64  `json:"GmtUpdate,omitempty" xml:"GmtUpdate,omitempty"`
	MassId       *int64  `json:"MassId,omitempty" xml:"MassId,omitempty"`
	Memo         *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	MsgId        *int64  `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	Status       *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ReadMessageContentResponseBodyDataDatasItem) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageContentResponseBodyDataDatasItem) GoString() string {
	return s.String()
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetCategoryName(v string) *ReadMessageContentResponseBodyDataDatasItem {
	s.CategoryName = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetClassId(v int64) *ReadMessageContentResponseBodyDataDatasItem {
	s.ClassId = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetContent(v string) *ReadMessageContentResponseBodyDataDatasItem {
	s.Content = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetDeleted(v int32) *ReadMessageContentResponseBodyDataDatasItem {
	s.Deleted = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetGmtCreated(v int64) *ReadMessageContentResponseBodyDataDatasItem {
	s.GmtCreated = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetGmtUpdate(v int64) *ReadMessageContentResponseBodyDataDatasItem {
	s.GmtUpdate = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetMassId(v int64) *ReadMessageContentResponseBodyDataDatasItem {
	s.MassId = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetMemo(v string) *ReadMessageContentResponseBodyDataDatasItem {
	s.Memo = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetMsgId(v int64) *ReadMessageContentResponseBodyDataDatasItem {
	s.MsgId = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetStatus(v int64) *ReadMessageContentResponseBodyDataDatasItem {
	s.Status = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetTitle(v string) *ReadMessageContentResponseBodyDataDatasItem {
	s.Title = &v
	return s
}

type ReadMessageContentResponseBodyDataDatasLastItem struct {
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	ClassId      *int64  `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Deleted      *int32  `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	GmtCreated   *int64  `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtUpdate    *int64  `json:"GmtUpdate,omitempty" xml:"GmtUpdate,omitempty"`
	MassId       *int64  `json:"MassId,omitempty" xml:"MassId,omitempty"`
	Memo         *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	MsgId        *int64  `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	Status       *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ReadMessageContentResponseBodyDataDatasLastItem) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageContentResponseBodyDataDatasLastItem) GoString() string {
	return s.String()
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetCategoryName(v string) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.CategoryName = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetClassId(v int64) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.ClassId = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetContent(v string) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.Content = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetDeleted(v int32) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.Deleted = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetGmtCreated(v int64) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.GmtCreated = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetGmtUpdate(v int64) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.GmtUpdate = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetMassId(v int64) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.MassId = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetMemo(v string) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.Memo = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetMsgId(v int64) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.MsgId = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetStatus(v int64) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.Status = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetTitle(v string) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.Title = &v
	return s
}

type ReadMessageContentResponseBodyDataDatasNextItem struct {
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	ClassId      *int64  `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Deleted      *int32  `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	GmtCreated   *int64  `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtUpdate    *int64  `json:"GmtUpdate,omitempty" xml:"GmtUpdate,omitempty"`
	MassId       *int64  `json:"MassId,omitempty" xml:"MassId,omitempty"`
	Memo         *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	MsgId        *int64  `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	Status       *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ReadMessageContentResponseBodyDataDatasNextItem) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageContentResponseBodyDataDatasNextItem) GoString() string {
	return s.String()
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetCategoryName(v string) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.CategoryName = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetClassId(v int64) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.ClassId = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetContent(v string) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.Content = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetDeleted(v int32) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.Deleted = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetGmtCreated(v int64) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.GmtCreated = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetGmtUpdate(v int64) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.GmtUpdate = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetMassId(v int64) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.MassId = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetMemo(v string) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.Memo = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetMsgId(v int64) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.MsgId = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetStatus(v int64) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.Status = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetTitle(v string) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.Title = &v
	return s
}

type ReadMessageContentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadMessageContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadMessageContentResponse) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageContentResponse) GoString() string {
	return s.String()
}

func (s *ReadMessageContentResponse) SetHeaders(v map[string]*string) *ReadMessageContentResponse {
	s.Headers = v
	return s
}

func (s *ReadMessageContentResponse) SetStatusCode(v int32) *ReadMessageContentResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadMessageContentResponse) SetBody(v *ReadMessageContentResponseBody) *ReadMessageContentResponse {
	s.Body = v
	return s
}

type ReadMessageListRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BizName        *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	ClassId        *int64  `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	ClientSource   *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	Content        *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Cookies        *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	Loc            *string `json:"Loc,omitempty" xml:"Loc,omitempty"`
	MaxResults     *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Page           *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SrcUrl         *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantCode     *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	Title          *string `json:"Title,omitempty" xml:"Title,omitempty"`
	UidType        *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadMessageListRequest) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageListRequest) GoString() string {
	return s.String()
}

func (s *ReadMessageListRequest) SetAcceptLanguage(v string) *ReadMessageListRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadMessageListRequest) SetAppName(v string) *ReadMessageListRequest {
	s.AppName = &v
	return s
}

func (s *ReadMessageListRequest) SetBizName(v string) *ReadMessageListRequest {
	s.BizName = &v
	return s
}

func (s *ReadMessageListRequest) SetCallerProtocol(v string) *ReadMessageListRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadMessageListRequest) SetClassId(v int64) *ReadMessageListRequest {
	s.ClassId = &v
	return s
}

func (s *ReadMessageListRequest) SetClientSource(v string) *ReadMessageListRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadMessageListRequest) SetContent(v string) *ReadMessageListRequest {
	s.Content = &v
	return s
}

func (s *ReadMessageListRequest) SetCookies(v string) *ReadMessageListRequest {
	s.Cookies = &v
	return s
}

func (s *ReadMessageListRequest) SetLoc(v string) *ReadMessageListRequest {
	s.Loc = &v
	return s
}

func (s *ReadMessageListRequest) SetMaxResults(v int32) *ReadMessageListRequest {
	s.MaxResults = &v
	return s
}

func (s *ReadMessageListRequest) SetNextToken(v string) *ReadMessageListRequest {
	s.NextToken = &v
	return s
}

func (s *ReadMessageListRequest) SetPage(v int32) *ReadMessageListRequest {
	s.Page = &v
	return s
}

func (s *ReadMessageListRequest) SetPageSize(v int32) *ReadMessageListRequest {
	s.PageSize = &v
	return s
}

func (s *ReadMessageListRequest) SetSrcUrl(v string) *ReadMessageListRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadMessageListRequest) SetStatus(v int32) *ReadMessageListRequest {
	s.Status = &v
	return s
}

func (s *ReadMessageListRequest) SetTenantCode(v string) *ReadMessageListRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadMessageListRequest) SetTitle(v string) *ReadMessageListRequest {
	s.Title = &v
	return s
}

func (s *ReadMessageListRequest) SetUidType(v string) *ReadMessageListRequest {
	s.UidType = &v
	return s
}

type ReadMessageListResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ReadMessageListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadMessageListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageListResponseBody) GoString() string {
	return s.String()
}

func (s *ReadMessageListResponseBody) SetCode(v string) *ReadMessageListResponseBody {
	s.Code = &v
	return s
}

func (s *ReadMessageListResponseBody) SetData(v *ReadMessageListResponseBodyData) *ReadMessageListResponseBody {
	s.Data = v
	return s
}

func (s *ReadMessageListResponseBody) SetMessage(v string) *ReadMessageListResponseBody {
	s.Message = &v
	return s
}

func (s *ReadMessageListResponseBody) SetRequestId(v string) *ReadMessageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadMessageListResponseBody) SetSuccess(v bool) *ReadMessageListResponseBody {
	s.Success = &v
	return s
}

type ReadMessageListResponseBodyData struct {
	Count      *int64                                 `json:"Count,omitempty" xml:"Count,omitempty"`
	MaxResults *int64                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Page       *int32                                 `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Rows       []*ReadMessageListResponseBodyDataRows `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
}

func (s ReadMessageListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadMessageListResponseBodyData) SetCount(v int64) *ReadMessageListResponseBodyData {
	s.Count = &v
	return s
}

func (s *ReadMessageListResponseBodyData) SetMaxResults(v int64) *ReadMessageListResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ReadMessageListResponseBodyData) SetNextToken(v string) *ReadMessageListResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ReadMessageListResponseBodyData) SetPage(v int32) *ReadMessageListResponseBodyData {
	s.Page = &v
	return s
}

func (s *ReadMessageListResponseBodyData) SetPageSize(v int32) *ReadMessageListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ReadMessageListResponseBodyData) SetRows(v []*ReadMessageListResponseBodyDataRows) *ReadMessageListResponseBodyData {
	s.Rows = v
	return s
}

type ReadMessageListResponseBodyDataRows struct {
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	Class        *string `json:"Class,omitempty" xml:"Class,omitempty"`
	ClassId      *int64  `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Deleted      *int32  `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	GmtCreated   *int64  `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtUpdate    *int64  `json:"GmtUpdate,omitempty" xml:"GmtUpdate,omitempty"`
	MassId       *int64  `json:"MassId,omitempty" xml:"MassId,omitempty"`
	Memo         *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	MsgId        *int64  `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	Status       *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ReadMessageListResponseBodyDataRows) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageListResponseBodyDataRows) GoString() string {
	return s.String()
}

func (s *ReadMessageListResponseBodyDataRows) SetCategoryName(v string) *ReadMessageListResponseBodyDataRows {
	s.CategoryName = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetClass(v string) *ReadMessageListResponseBodyDataRows {
	s.Class = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetClassId(v int64) *ReadMessageListResponseBodyDataRows {
	s.ClassId = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetContent(v string) *ReadMessageListResponseBodyDataRows {
	s.Content = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetDeleted(v int32) *ReadMessageListResponseBodyDataRows {
	s.Deleted = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetGmtCreated(v int64) *ReadMessageListResponseBodyDataRows {
	s.GmtCreated = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetGmtUpdate(v int64) *ReadMessageListResponseBodyDataRows {
	s.GmtUpdate = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetMassId(v int64) *ReadMessageListResponseBodyDataRows {
	s.MassId = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetMemo(v string) *ReadMessageListResponseBodyDataRows {
	s.Memo = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetMsgId(v int64) *ReadMessageListResponseBodyDataRows {
	s.MsgId = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetStatus(v int64) *ReadMessageListResponseBodyDataRows {
	s.Status = &v
	return s
}

func (s *ReadMessageListResponseBodyDataRows) SetTitle(v string) *ReadMessageListResponseBodyDataRows {
	s.Title = &v
	return s
}

type ReadMessageListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadMessageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadMessageListResponse) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageListResponse) GoString() string {
	return s.String()
}

func (s *ReadMessageListResponse) SetHeaders(v map[string]*string) *ReadMessageListResponse {
	s.Headers = v
	return s
}

func (s *ReadMessageListResponse) SetStatusCode(v int32) *ReadMessageListResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadMessageListResponse) SetBody(v *ReadMessageListResponseBody) *ReadMessageListResponse {
	s.Body = v
	return s
}

type ReadMessageNewTotalRequest struct {
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BizName        *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	ClientSource   *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	Cookies        *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	SrcUrl         *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	TenantCode     *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	UidType        *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadMessageNewTotalRequest) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageNewTotalRequest) GoString() string {
	return s.String()
}

func (s *ReadMessageNewTotalRequest) SetAcceptLanguage(v string) *ReadMessageNewTotalRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadMessageNewTotalRequest) SetAppName(v string) *ReadMessageNewTotalRequest {
	s.AppName = &v
	return s
}

func (s *ReadMessageNewTotalRequest) SetBizName(v string) *ReadMessageNewTotalRequest {
	s.BizName = &v
	return s
}

func (s *ReadMessageNewTotalRequest) SetCallerProtocol(v string) *ReadMessageNewTotalRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadMessageNewTotalRequest) SetClientSource(v string) *ReadMessageNewTotalRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadMessageNewTotalRequest) SetCookies(v string) *ReadMessageNewTotalRequest {
	s.Cookies = &v
	return s
}

func (s *ReadMessageNewTotalRequest) SetSrcUrl(v string) *ReadMessageNewTotalRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadMessageNewTotalRequest) SetTenantCode(v string) *ReadMessageNewTotalRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadMessageNewTotalRequest) SetUidType(v string) *ReadMessageNewTotalRequest {
	s.UidType = &v
	return s
}

type ReadMessageNewTotalResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadMessageNewTotalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageNewTotalResponseBody) GoString() string {
	return s.String()
}

func (s *ReadMessageNewTotalResponseBody) SetCode(v string) *ReadMessageNewTotalResponseBody {
	s.Code = &v
	return s
}

func (s *ReadMessageNewTotalResponseBody) SetData(v int64) *ReadMessageNewTotalResponseBody {
	s.Data = &v
	return s
}

func (s *ReadMessageNewTotalResponseBody) SetMessage(v string) *ReadMessageNewTotalResponseBody {
	s.Message = &v
	return s
}

func (s *ReadMessageNewTotalResponseBody) SetRequestId(v string) *ReadMessageNewTotalResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadMessageNewTotalResponseBody) SetSuccess(v bool) *ReadMessageNewTotalResponseBody {
	s.Success = &v
	return s
}

type ReadMessageNewTotalResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadMessageNewTotalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadMessageNewTotalResponse) String() string {
	return tea.Prettify(s)
}

func (s ReadMessageNewTotalResponse) GoString() string {
	return s.String()
}

func (s *ReadMessageNewTotalResponse) SetHeaders(v map[string]*string) *ReadMessageNewTotalResponse {
	s.Headers = v
	return s
}

func (s *ReadMessageNewTotalResponse) SetStatusCode(v int32) *ReadMessageNewTotalResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadMessageNewTotalResponse) SetBody(v *ReadMessageNewTotalResponseBody) *ReadMessageNewTotalResponse {
	s.Body = v
	return s
}

type ReadNumGroupByClassRequest struct {
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BizName        *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	ClientSource   *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	Cookies        *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	SrcUrl         *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	TenantCode     *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	UidType        *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadNumGroupByClassRequest) String() string {
	return tea.Prettify(s)
}

func (s ReadNumGroupByClassRequest) GoString() string {
	return s.String()
}

func (s *ReadNumGroupByClassRequest) SetAcceptLanguage(v string) *ReadNumGroupByClassRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadNumGroupByClassRequest) SetAppName(v string) *ReadNumGroupByClassRequest {
	s.AppName = &v
	return s
}

func (s *ReadNumGroupByClassRequest) SetBizName(v string) *ReadNumGroupByClassRequest {
	s.BizName = &v
	return s
}

func (s *ReadNumGroupByClassRequest) SetCallerProtocol(v string) *ReadNumGroupByClassRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadNumGroupByClassRequest) SetClientSource(v string) *ReadNumGroupByClassRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadNumGroupByClassRequest) SetCookies(v string) *ReadNumGroupByClassRequest {
	s.Cookies = &v
	return s
}

func (s *ReadNumGroupByClassRequest) SetSrcUrl(v string) *ReadNumGroupByClassRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadNumGroupByClassRequest) SetTenantCode(v string) *ReadNumGroupByClassRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadNumGroupByClassRequest) SetUidType(v string) *ReadNumGroupByClassRequest {
	s.UidType = &v
	return s
}

type ReadNumGroupByClassResponseBody struct {
	// example:
	//
	// SUCCESS
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ReadNumGroupByClassResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadNumGroupByClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReadNumGroupByClassResponseBody) GoString() string {
	return s.String()
}

func (s *ReadNumGroupByClassResponseBody) SetCode(v string) *ReadNumGroupByClassResponseBody {
	s.Code = &v
	return s
}

func (s *ReadNumGroupByClassResponseBody) SetData(v []*ReadNumGroupByClassResponseBodyData) *ReadNumGroupByClassResponseBody {
	s.Data = v
	return s
}

func (s *ReadNumGroupByClassResponseBody) SetMessage(v string) *ReadNumGroupByClassResponseBody {
	s.Message = &v
	return s
}

func (s *ReadNumGroupByClassResponseBody) SetRequestId(v string) *ReadNumGroupByClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadNumGroupByClassResponseBody) SetSuccess(v bool) *ReadNumGroupByClassResponseBody {
	s.Success = &v
	return s
}

type ReadNumGroupByClassResponseBodyData struct {
	ClassId  *int64 `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	MsgCount *int64 `json:"MsgCount,omitempty" xml:"MsgCount,omitempty"`
}

func (s ReadNumGroupByClassResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ReadNumGroupByClassResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadNumGroupByClassResponseBodyData) SetClassId(v int64) *ReadNumGroupByClassResponseBodyData {
	s.ClassId = &v
	return s
}

func (s *ReadNumGroupByClassResponseBodyData) SetMsgCount(v int64) *ReadNumGroupByClassResponseBodyData {
	s.MsgCount = &v
	return s
}

type ReadNumGroupByClassResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadNumGroupByClassResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadNumGroupByClassResponse) String() string {
	return tea.Prettify(s)
}

func (s ReadNumGroupByClassResponse) GoString() string {
	return s.String()
}

func (s *ReadNumGroupByClassResponse) SetHeaders(v map[string]*string) *ReadNumGroupByClassResponse {
	s.Headers = v
	return s
}

func (s *ReadNumGroupByClassResponse) SetStatusCode(v int32) *ReadNumGroupByClassResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadNumGroupByClassResponse) SetBody(v *ReadNumGroupByClassResponseBody) *ReadNumGroupByClassResponse {
	s.Body = v
	return s
}

type ReadNumGroupTotalRequest struct {
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BizName        *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	ClientSource   *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	Cookies        *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	SrcUrl         *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	TenantCode     *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	UidType        *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadNumGroupTotalRequest) String() string {
	return tea.Prettify(s)
}

func (s ReadNumGroupTotalRequest) GoString() string {
	return s.String()
}

func (s *ReadNumGroupTotalRequest) SetAcceptLanguage(v string) *ReadNumGroupTotalRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadNumGroupTotalRequest) SetAppName(v string) *ReadNumGroupTotalRequest {
	s.AppName = &v
	return s
}

func (s *ReadNumGroupTotalRequest) SetBizName(v string) *ReadNumGroupTotalRequest {
	s.BizName = &v
	return s
}

func (s *ReadNumGroupTotalRequest) SetCallerProtocol(v string) *ReadNumGroupTotalRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadNumGroupTotalRequest) SetClientSource(v string) *ReadNumGroupTotalRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadNumGroupTotalRequest) SetCookies(v string) *ReadNumGroupTotalRequest {
	s.Cookies = &v
	return s
}

func (s *ReadNumGroupTotalRequest) SetSrcUrl(v string) *ReadNumGroupTotalRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadNumGroupTotalRequest) SetTenantCode(v string) *ReadNumGroupTotalRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadNumGroupTotalRequest) SetUidType(v string) *ReadNumGroupTotalRequest {
	s.UidType = &v
	return s
}

type ReadNumGroupTotalResponseBody struct {
	// example:
	//
	// SUCCESS
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ReadNumGroupTotalResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadNumGroupTotalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReadNumGroupTotalResponseBody) GoString() string {
	return s.String()
}

func (s *ReadNumGroupTotalResponseBody) SetCode(v string) *ReadNumGroupTotalResponseBody {
	s.Code = &v
	return s
}

func (s *ReadNumGroupTotalResponseBody) SetData(v []*ReadNumGroupTotalResponseBodyData) *ReadNumGroupTotalResponseBody {
	s.Data = v
	return s
}

func (s *ReadNumGroupTotalResponseBody) SetMessage(v string) *ReadNumGroupTotalResponseBody {
	s.Message = &v
	return s
}

func (s *ReadNumGroupTotalResponseBody) SetRequestId(v string) *ReadNumGroupTotalResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadNumGroupTotalResponseBody) SetSuccess(v bool) *ReadNumGroupTotalResponseBody {
	s.Success = &v
	return s
}

type ReadNumGroupTotalResponseBodyData struct {
	Id          *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	ReadCount   *int64 `json:"ReadCount,omitempty" xml:"ReadCount,omitempty"`
	TotalCount  *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UnReadCount *int64 `json:"UnReadCount,omitempty" xml:"UnReadCount,omitempty"`
}

func (s ReadNumGroupTotalResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ReadNumGroupTotalResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadNumGroupTotalResponseBodyData) SetId(v int64) *ReadNumGroupTotalResponseBodyData {
	s.Id = &v
	return s
}

func (s *ReadNumGroupTotalResponseBodyData) SetReadCount(v int64) *ReadNumGroupTotalResponseBodyData {
	s.ReadCount = &v
	return s
}

func (s *ReadNumGroupTotalResponseBodyData) SetTotalCount(v int64) *ReadNumGroupTotalResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ReadNumGroupTotalResponseBodyData) SetUnReadCount(v int64) *ReadNumGroupTotalResponseBodyData {
	s.UnReadCount = &v
	return s
}

type ReadNumGroupTotalResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadNumGroupTotalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadNumGroupTotalResponse) String() string {
	return tea.Prettify(s)
}

func (s ReadNumGroupTotalResponse) GoString() string {
	return s.String()
}

func (s *ReadNumGroupTotalResponse) SetHeaders(v map[string]*string) *ReadNumGroupTotalResponse {
	s.Headers = v
	return s
}

func (s *ReadNumGroupTotalResponse) SetStatusCode(v int32) *ReadNumGroupTotalResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadNumGroupTotalResponse) SetBody(v *ReadNumGroupTotalResponseBody) *ReadNumGroupTotalResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("notifications"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 方法描述：删除消息
//
// @param request - DelMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DelMessageResponse
func (client *Client) DelMessageWithOptions(request *DelMessageRequest, runtime *util.RuntimeOptions) (_result *DelMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.BizName)) {
		body["BizName"] = request.BizName
	}

	if !tea.BoolValue(util.IsUnset(request.CallerProtocol)) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSource)) {
		body["ClientSource"] = request.ClientSource
	}

	if !tea.BoolValue(util.IsUnset(request.Cookies)) {
		body["Cookies"] = request.Cookies
	}

	if !tea.BoolValue(util.IsUnset(request.MsgId)) {
		body["MsgId"] = request.MsgId
	}

	if !tea.BoolValue(util.IsUnset(request.SrcUrl)) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !tea.BoolValue(util.IsUnset(request.TenantCode)) {
		body["TenantCode"] = request.TenantCode
	}

	if !tea.BoolValue(util.IsUnset(request.UidType)) {
		body["UidType"] = request.UidType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DelMessage"),
		Version:     tea.String("2024-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DelMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：删除消息
//
// @param request - DelMessageRequest
//
// @return DelMessageResponse
func (client *Client) DelMessage(request *DelMessageRequest) (_result *DelMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DelMessageResponse{}
	_body, _err := client.DelMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 方法描述：站内信全部删除（逻辑删除）
//
// @param request - DeleteAllMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAllMessageResponse
func (client *Client) DeleteAllMessageWithOptions(request *DeleteAllMessageRequest, runtime *util.RuntimeOptions) (_result *DeleteAllMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.BizName)) {
		body["BizName"] = request.BizName
	}

	if !tea.BoolValue(util.IsUnset(request.CallerProtocol)) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		body["ClassId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSource)) {
		body["ClientSource"] = request.ClientSource
	}

	if !tea.BoolValue(util.IsUnset(request.Cookies)) {
		body["Cookies"] = request.Cookies
	}

	if !tea.BoolValue(util.IsUnset(request.SrcUrl)) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !tea.BoolValue(util.IsUnset(request.TenantCode)) {
		body["TenantCode"] = request.TenantCode
	}

	if !tea.BoolValue(util.IsUnset(request.UidType)) {
		body["UidType"] = request.UidType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAllMessage"),
		Version:     tea.String("2024-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAllMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：站内信全部删除（逻辑删除）
//
// @param request - DeleteAllMessageRequest
//
// @return DeleteAllMessageResponse
func (client *Client) DeleteAllMessage(request *DeleteAllMessageRequest) (_result *DeleteAllMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAllMessageResponse{}
	_body, _err := client.DeleteAllMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 方法描述：分类全部标记为已读，不填则全部标记
//
// @param request - ReadAllMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadAllMessageResponse
func (client *Client) ReadAllMessageWithOptions(request *ReadAllMessageRequest, runtime *util.RuntimeOptions) (_result *ReadAllMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.BizName)) {
		body["BizName"] = request.BizName
	}

	if !tea.BoolValue(util.IsUnset(request.CallerProtocol)) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		body["ClassId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSource)) {
		body["ClientSource"] = request.ClientSource
	}

	if !tea.BoolValue(util.IsUnset(request.Cookies)) {
		body["Cookies"] = request.Cookies
	}

	if !tea.BoolValue(util.IsUnset(request.SrcUrl)) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !tea.BoolValue(util.IsUnset(request.TenantCode)) {
		body["TenantCode"] = request.TenantCode
	}

	if !tea.BoolValue(util.IsUnset(request.UidType)) {
		body["UidType"] = request.UidType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReadAllMessage"),
		Version:     tea.String("2024-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReadAllMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：分类全部标记为已读，不填则全部标记
//
// @param request - ReadAllMessageRequest
//
// @return ReadAllMessageResponse
func (client *Client) ReadAllMessage(request *ReadAllMessageRequest) (_result *ReadAllMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReadAllMessageResponse{}
	_body, _err := client.ReadAllMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 方法描述：获取各分类已读消息数
//
// @param request - ReadClassNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadClassNameResponse
func (client *Client) ReadClassNameWithOptions(request *ReadClassNameRequest, runtime *util.RuntimeOptions) (_result *ReadClassNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.BizName)) {
		body["BizName"] = request.BizName
	}

	if !tea.BoolValue(util.IsUnset(request.CallerProtocol)) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSource)) {
		body["ClientSource"] = request.ClientSource
	}

	if !tea.BoolValue(util.IsUnset(request.Cookies)) {
		body["Cookies"] = request.Cookies
	}

	if !tea.BoolValue(util.IsUnset(request.SrcUrl)) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !tea.BoolValue(util.IsUnset(request.TenantCode)) {
		body["TenantCode"] = request.TenantCode
	}

	if !tea.BoolValue(util.IsUnset(request.UidType)) {
		body["UidType"] = request.UidType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReadClassName"),
		Version:     tea.String("2024-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReadClassNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：获取各分类已读消息数
//
// @param request - ReadClassNameRequest
//
// @return ReadClassNameResponse
func (client *Client) ReadClassName(request *ReadClassNameRequest) (_result *ReadClassNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReadClassNameResponse{}
	_body, _err := client.ReadClassNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 方法描述：消息标记为已读
//
// @param request - ReadMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMessageResponse
func (client *Client) ReadMessageWithOptions(request *ReadMessageRequest, runtime *util.RuntimeOptions) (_result *ReadMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.BizName)) {
		body["BizName"] = request.BizName
	}

	if !tea.BoolValue(util.IsUnset(request.CallerProtocol)) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSource)) {
		body["ClientSource"] = request.ClientSource
	}

	if !tea.BoolValue(util.IsUnset(request.Cookies)) {
		body["Cookies"] = request.Cookies
	}

	if !tea.BoolValue(util.IsUnset(request.MsgId)) {
		body["MsgId"] = request.MsgId
	}

	if !tea.BoolValue(util.IsUnset(request.SrcUrl)) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !tea.BoolValue(util.IsUnset(request.TenantCode)) {
		body["TenantCode"] = request.TenantCode
	}

	if !tea.BoolValue(util.IsUnset(request.UidType)) {
		body["UidType"] = request.UidType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReadMessage"),
		Version:     tea.String("2024-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReadMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：消息标记为已读
//
// @param request - ReadMessageRequest
//
// @return ReadMessageResponse
func (client *Client) ReadMessage(request *ReadMessageRequest) (_result *ReadMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReadMessageResponse{}
	_body, _err := client.ReadMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 方法描述：获取消息正文
//
// @param request - ReadMessageContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMessageContentResponse
func (client *Client) ReadMessageContentWithOptions(request *ReadMessageContentRequest, runtime *util.RuntimeOptions) (_result *ReadMessageContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.BizName)) {
		body["BizName"] = request.BizName
	}

	if !tea.BoolValue(util.IsUnset(request.CallerProtocol)) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		body["ClassId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSource)) {
		body["ClientSource"] = request.ClientSource
	}

	if !tea.BoolValue(util.IsUnset(request.Cookies)) {
		body["Cookies"] = request.Cookies
	}

	if !tea.BoolValue(util.IsUnset(request.MsgId)) {
		body["MsgId"] = request.MsgId
	}

	if !tea.BoolValue(util.IsUnset(request.SrcUrl)) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TenantCode)) {
		body["TenantCode"] = request.TenantCode
	}

	if !tea.BoolValue(util.IsUnset(request.UidType)) {
		body["UidType"] = request.UidType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReadMessageContent"),
		Version:     tea.String("2024-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReadMessageContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：获取消息正文
//
// @param request - ReadMessageContentRequest
//
// @return ReadMessageContentResponse
func (client *Client) ReadMessageContent(request *ReadMessageContentRequest) (_result *ReadMessageContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReadMessageContentResponse{}
	_body, _err := client.ReadMessageContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 方法描述：获取消息列表
//
// @param request - ReadMessageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMessageListResponse
func (client *Client) ReadMessageListWithOptions(request *ReadMessageListRequest, runtime *util.RuntimeOptions) (_result *ReadMessageListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.BizName)) {
		body["BizName"] = request.BizName
	}

	if !tea.BoolValue(util.IsUnset(request.CallerProtocol)) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		body["ClassId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSource)) {
		body["ClientSource"] = request.ClientSource
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Cookies)) {
		body["Cookies"] = request.Cookies
	}

	if !tea.BoolValue(util.IsUnset(request.Loc)) {
		body["Loc"] = request.Loc
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		body["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SrcUrl)) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TenantCode)) {
		body["TenantCode"] = request.TenantCode
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UidType)) {
		body["UidType"] = request.UidType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReadMessageList"),
		Version:     tea.String("2024-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReadMessageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：获取消息列表
//
// @param request - ReadMessageListRequest
//
// @return ReadMessageListResponse
func (client *Client) ReadMessageList(request *ReadMessageListRequest) (_result *ReadMessageListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReadMessageListResponse{}
	_body, _err := client.ReadMessageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 方法描述：获取未读消息总数
//
// @param request - ReadMessageNewTotalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMessageNewTotalResponse
func (client *Client) ReadMessageNewTotalWithOptions(request *ReadMessageNewTotalRequest, runtime *util.RuntimeOptions) (_result *ReadMessageNewTotalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.BizName)) {
		body["BizName"] = request.BizName
	}

	if !tea.BoolValue(util.IsUnset(request.CallerProtocol)) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSource)) {
		body["ClientSource"] = request.ClientSource
	}

	if !tea.BoolValue(util.IsUnset(request.Cookies)) {
		body["Cookies"] = request.Cookies
	}

	if !tea.BoolValue(util.IsUnset(request.SrcUrl)) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !tea.BoolValue(util.IsUnset(request.TenantCode)) {
		body["TenantCode"] = request.TenantCode
	}

	if !tea.BoolValue(util.IsUnset(request.UidType)) {
		body["UidType"] = request.UidType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReadMessageNewTotal"),
		Version:     tea.String("2024-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReadMessageNewTotalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：获取未读消息总数
//
// @param request - ReadMessageNewTotalRequest
//
// @return ReadMessageNewTotalResponse
func (client *Client) ReadMessageNewTotal(request *ReadMessageNewTotalRequest) (_result *ReadMessageNewTotalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReadMessageNewTotalResponse{}
	_body, _err := client.ReadMessageNewTotalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 方法描述：获取各分类已读消息数
//
// @param request - ReadNumGroupByClassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadNumGroupByClassResponse
func (client *Client) ReadNumGroupByClassWithOptions(request *ReadNumGroupByClassRequest, runtime *util.RuntimeOptions) (_result *ReadNumGroupByClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.BizName)) {
		body["BizName"] = request.BizName
	}

	if !tea.BoolValue(util.IsUnset(request.CallerProtocol)) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSource)) {
		body["ClientSource"] = request.ClientSource
	}

	if !tea.BoolValue(util.IsUnset(request.Cookies)) {
		body["Cookies"] = request.Cookies
	}

	if !tea.BoolValue(util.IsUnset(request.SrcUrl)) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !tea.BoolValue(util.IsUnset(request.TenantCode)) {
		body["TenantCode"] = request.TenantCode
	}

	if !tea.BoolValue(util.IsUnset(request.UidType)) {
		body["UidType"] = request.UidType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReadNumGroupByClass"),
		Version:     tea.String("2024-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReadNumGroupByClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：获取各分类已读消息数
//
// @param request - ReadNumGroupByClassRequest
//
// @return ReadNumGroupByClassResponse
func (client *Client) ReadNumGroupByClass(request *ReadNumGroupByClassRequest) (_result *ReadNumGroupByClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReadNumGroupByClassResponse{}
	_body, _err := client.ReadNumGroupByClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 方法描述：获取所有分类下的信息
//
// @param request - ReadNumGroupTotalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadNumGroupTotalResponse
func (client *Client) ReadNumGroupTotalWithOptions(request *ReadNumGroupTotalRequest, runtime *util.RuntimeOptions) (_result *ReadNumGroupTotalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.BizName)) {
		body["BizName"] = request.BizName
	}

	if !tea.BoolValue(util.IsUnset(request.CallerProtocol)) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSource)) {
		body["ClientSource"] = request.ClientSource
	}

	if !tea.BoolValue(util.IsUnset(request.Cookies)) {
		body["Cookies"] = request.Cookies
	}

	if !tea.BoolValue(util.IsUnset(request.SrcUrl)) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !tea.BoolValue(util.IsUnset(request.TenantCode)) {
		body["TenantCode"] = request.TenantCode
	}

	if !tea.BoolValue(util.IsUnset(request.UidType)) {
		body["UidType"] = request.UidType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReadNumGroupTotal"),
		Version:     tea.String("2024-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReadNumGroupTotalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：获取所有分类下的信息
//
// @param request - ReadNumGroupTotalRequest
//
// @return ReadNumGroupTotalResponse
func (client *Client) ReadNumGroupTotal(request *ReadNumGroupTotalRequest) (_result *ReadNumGroupTotalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReadNumGroupTotalResponse{}
	_body, _err := client.ReadNumGroupTotalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
