// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DeleteUnbeianIpCheckTypeRequest struct {
	// This parameter is required.
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// This parameter is required.
	CheckType *int32 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// This parameter is required.
	Ip     *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s DeleteUnbeianIpCheckTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUnbeianIpCheckTypeRequest) GoString() string {
	return s.String()
}

func (s *DeleteUnbeianIpCheckTypeRequest) SetCaller(v string) *DeleteUnbeianIpCheckTypeRequest {
	s.Caller = &v
	return s
}

func (s *DeleteUnbeianIpCheckTypeRequest) SetCheckType(v int32) *DeleteUnbeianIpCheckTypeRequest {
	s.CheckType = &v
	return s
}

func (s *DeleteUnbeianIpCheckTypeRequest) SetIp(v string) *DeleteUnbeianIpCheckTypeRequest {
	s.Ip = &v
	return s
}

func (s *DeleteUnbeianIpCheckTypeRequest) SetRemark(v string) *DeleteUnbeianIpCheckTypeRequest {
	s.Remark = &v
	return s
}

type DeleteUnbeianIpCheckTypeResponseBody struct {
	ErrorCode                   *int32                                                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage                *string                                                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HuntressIpCheckTypeResultDO *DeleteUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO `json:"HuntressIpCheckTypeResultDO,omitempty" xml:"HuntressIpCheckTypeResultDO,omitempty" type:"Struct"`
	RequestId                   *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                     *bool                                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUnbeianIpCheckTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUnbeianIpCheckTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUnbeianIpCheckTypeResponseBody) SetErrorCode(v int32) *DeleteUnbeianIpCheckTypeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteUnbeianIpCheckTypeResponseBody) SetErrorMessage(v string) *DeleteUnbeianIpCheckTypeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteUnbeianIpCheckTypeResponseBody) SetHuntressIpCheckTypeResultDO(v *DeleteUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO) *DeleteUnbeianIpCheckTypeResponseBody {
	s.HuntressIpCheckTypeResultDO = v
	return s
}

func (s *DeleteUnbeianIpCheckTypeResponseBody) SetRequestId(v string) *DeleteUnbeianIpCheckTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUnbeianIpCheckTypeResponseBody) SetSuccess(v bool) *DeleteUnbeianIpCheckTypeResponseBody {
	s.Success = &v
	return s
}

type DeleteUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO struct {
	Msg     *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO) String() string {
	return tea.Prettify(s)
}

func (s DeleteUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO) GoString() string {
	return s.String()
}

func (s *DeleteUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO) SetMsg(v string) *DeleteUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO {
	s.Msg = &v
	return s
}

func (s *DeleteUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO) SetSuccess(v bool) *DeleteUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO {
	s.Success = &v
	return s
}

type DeleteUnbeianIpCheckTypeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUnbeianIpCheckTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUnbeianIpCheckTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUnbeianIpCheckTypeResponse) GoString() string {
	return s.String()
}

func (s *DeleteUnbeianIpCheckTypeResponse) SetHeaders(v map[string]*string) *DeleteUnbeianIpCheckTypeResponse {
	s.Headers = v
	return s
}

func (s *DeleteUnbeianIpCheckTypeResponse) SetStatusCode(v int32) *DeleteUnbeianIpCheckTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUnbeianIpCheckTypeResponse) SetBody(v *DeleteUnbeianIpCheckTypeResponseBody) *DeleteUnbeianIpCheckTypeResponse {
	s.Body = v
	return s
}

type GetMainDomainRequest struct {
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s GetMainDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMainDomainRequest) GoString() string {
	return s.String()
}

func (s *GetMainDomainRequest) SetDomain(v string) *GetMainDomainRequest {
	s.Domain = &v
	return s
}

type GetMainDomainResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMainDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMainDomainResponseBody) GoString() string {
	return s.String()
}

func (s *GetMainDomainResponseBody) SetCode(v string) *GetMainDomainResponseBody {
	s.Code = &v
	return s
}

func (s *GetMainDomainResponseBody) SetData(v string) *GetMainDomainResponseBody {
	s.Data = &v
	return s
}

func (s *GetMainDomainResponseBody) SetMessage(v string) *GetMainDomainResponseBody {
	s.Message = &v
	return s
}

func (s *GetMainDomainResponseBody) SetRequestId(v string) *GetMainDomainResponseBody {
	s.RequestId = &v
	return s
}

type GetMainDomainResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMainDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMainDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMainDomainResponse) GoString() string {
	return s.String()
}

func (s *GetMainDomainResponse) SetHeaders(v map[string]*string) *GetMainDomainResponse {
	s.Headers = v
	return s
}

func (s *GetMainDomainResponse) SetStatusCode(v int32) *GetMainDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMainDomainResponse) SetBody(v *GetMainDomainResponseBody) *GetMainDomainResponse {
	s.Body = v
	return s
}

type InsertUnbeianIpCheckTypeRequest struct {
	// This parameter is required.
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// This parameter is required.
	CheckType *int32 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// This parameter is required.
	Ip     *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s InsertUnbeianIpCheckTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertUnbeianIpCheckTypeRequest) GoString() string {
	return s.String()
}

func (s *InsertUnbeianIpCheckTypeRequest) SetCaller(v string) *InsertUnbeianIpCheckTypeRequest {
	s.Caller = &v
	return s
}

func (s *InsertUnbeianIpCheckTypeRequest) SetCheckType(v int32) *InsertUnbeianIpCheckTypeRequest {
	s.CheckType = &v
	return s
}

func (s *InsertUnbeianIpCheckTypeRequest) SetIp(v string) *InsertUnbeianIpCheckTypeRequest {
	s.Ip = &v
	return s
}

func (s *InsertUnbeianIpCheckTypeRequest) SetRemark(v string) *InsertUnbeianIpCheckTypeRequest {
	s.Remark = &v
	return s
}

type InsertUnbeianIpCheckTypeResponseBody struct {
	ErrorCode                   *int32                                                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage                *string                                                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HuntressIpCheckTypeResultDO *InsertUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO `json:"HuntressIpCheckTypeResultDO,omitempty" xml:"HuntressIpCheckTypeResultDO,omitempty" type:"Struct"`
	RequestId                   *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                     *bool                                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertUnbeianIpCheckTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertUnbeianIpCheckTypeResponseBody) GoString() string {
	return s.String()
}

func (s *InsertUnbeianIpCheckTypeResponseBody) SetErrorCode(v int32) *InsertUnbeianIpCheckTypeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InsertUnbeianIpCheckTypeResponseBody) SetErrorMessage(v string) *InsertUnbeianIpCheckTypeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *InsertUnbeianIpCheckTypeResponseBody) SetHuntressIpCheckTypeResultDO(v *InsertUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO) *InsertUnbeianIpCheckTypeResponseBody {
	s.HuntressIpCheckTypeResultDO = v
	return s
}

func (s *InsertUnbeianIpCheckTypeResponseBody) SetRequestId(v string) *InsertUnbeianIpCheckTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertUnbeianIpCheckTypeResponseBody) SetSuccess(v bool) *InsertUnbeianIpCheckTypeResponseBody {
	s.Success = &v
	return s
}

type InsertUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO struct {
	Msg     *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO) String() string {
	return tea.Prettify(s)
}

func (s InsertUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO) GoString() string {
	return s.String()
}

func (s *InsertUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO) SetMsg(v string) *InsertUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO {
	s.Msg = &v
	return s
}

func (s *InsertUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO) SetSuccess(v bool) *InsertUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO {
	s.Success = &v
	return s
}

type InsertUnbeianIpCheckTypeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertUnbeianIpCheckTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertUnbeianIpCheckTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertUnbeianIpCheckTypeResponse) GoString() string {
	return s.String()
}

func (s *InsertUnbeianIpCheckTypeResponse) SetHeaders(v map[string]*string) *InsertUnbeianIpCheckTypeResponse {
	s.Headers = v
	return s
}

func (s *InsertUnbeianIpCheckTypeResponse) SetStatusCode(v int32) *InsertUnbeianIpCheckTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertUnbeianIpCheckTypeResponse) SetBody(v *InsertUnbeianIpCheckTypeResponseBody) *InsertUnbeianIpCheckTypeResponse {
	s.Body = v
	return s
}

type ListUnbeianIpCheckTypeRequest struct {
	// This parameter is required.
	Caller    *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	CheckType *int32  `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Limit     *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Page      *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	Remark    *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ListUnbeianIpCheckTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUnbeianIpCheckTypeRequest) GoString() string {
	return s.String()
}

func (s *ListUnbeianIpCheckTypeRequest) SetCaller(v string) *ListUnbeianIpCheckTypeRequest {
	s.Caller = &v
	return s
}

func (s *ListUnbeianIpCheckTypeRequest) SetCheckType(v int32) *ListUnbeianIpCheckTypeRequest {
	s.CheckType = &v
	return s
}

func (s *ListUnbeianIpCheckTypeRequest) SetIp(v string) *ListUnbeianIpCheckTypeRequest {
	s.Ip = &v
	return s
}

func (s *ListUnbeianIpCheckTypeRequest) SetLimit(v int32) *ListUnbeianIpCheckTypeRequest {
	s.Limit = &v
	return s
}

func (s *ListUnbeianIpCheckTypeRequest) SetPage(v int32) *ListUnbeianIpCheckTypeRequest {
	s.Page = &v
	return s
}

func (s *ListUnbeianIpCheckTypeRequest) SetRemark(v string) *ListUnbeianIpCheckTypeRequest {
	s.Remark = &v
	return s
}

type ListUnbeianIpCheckTypeResponseBody struct {
	ErrorCode                   *int32                                                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage                *string                                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HuntressIpCheckTypeResultDO *ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO `json:"HuntressIpCheckTypeResultDO,omitempty" xml:"HuntressIpCheckTypeResultDO,omitempty" type:"Struct"`
	RequestId                   *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                     *bool                                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUnbeianIpCheckTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUnbeianIpCheckTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ListUnbeianIpCheckTypeResponseBody) SetErrorCode(v int32) *ListUnbeianIpCheckTypeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListUnbeianIpCheckTypeResponseBody) SetErrorMessage(v string) *ListUnbeianIpCheckTypeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListUnbeianIpCheckTypeResponseBody) SetHuntressIpCheckTypeResultDO(v *ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO) *ListUnbeianIpCheckTypeResponseBody {
	s.HuntressIpCheckTypeResultDO = v
	return s
}

func (s *ListUnbeianIpCheckTypeResponseBody) SetRequestId(v string) *ListUnbeianIpCheckTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUnbeianIpCheckTypeResponseBody) SetSuccess(v bool) *ListUnbeianIpCheckTypeResponseBody {
	s.Success = &v
	return s
}

type ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO struct {
	List    []*ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDOList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Msg     *string                                                              `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Success *bool                                                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO) String() string {
	return tea.Prettify(s)
}

func (s ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO) GoString() string {
	return s.String()
}

func (s *ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO) SetList(v []*ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDOList) *ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO {
	s.List = v
	return s
}

func (s *ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO) SetMsg(v string) *ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO {
	s.Msg = &v
	return s
}

func (s *ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO) SetSuccess(v bool) *ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDO {
	s.Success = &v
	return s
}

type ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDOList struct {
	Aliuid    *int64  `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	Caller    *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	CheckType *int32  `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Remark    *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDOList) String() string {
	return tea.Prettify(s)
}

func (s ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDOList) GoString() string {
	return s.String()
}

func (s *ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDOList) SetAliuid(v int64) *ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDOList {
	s.Aliuid = &v
	return s
}

func (s *ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDOList) SetCaller(v string) *ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDOList {
	s.Caller = &v
	return s
}

func (s *ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDOList) SetCheckType(v int32) *ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDOList {
	s.CheckType = &v
	return s
}

func (s *ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDOList) SetIp(v string) *ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDOList {
	s.Ip = &v
	return s
}

func (s *ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDOList) SetRemark(v string) *ListUnbeianIpCheckTypeResponseBodyHuntressIpCheckTypeResultDOList {
	s.Remark = &v
	return s
}

type ListUnbeianIpCheckTypeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUnbeianIpCheckTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUnbeianIpCheckTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUnbeianIpCheckTypeResponse) GoString() string {
	return s.String()
}

func (s *ListUnbeianIpCheckTypeResponse) SetHeaders(v map[string]*string) *ListUnbeianIpCheckTypeResponse {
	s.Headers = v
	return s
}

func (s *ListUnbeianIpCheckTypeResponse) SetStatusCode(v int32) *ListUnbeianIpCheckTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUnbeianIpCheckTypeResponse) SetBody(v *ListUnbeianIpCheckTypeResponseBody) *ListUnbeianIpCheckTypeResponse {
	s.Body = v
	return s
}

type ManageAccessorDomainRequest struct {
	// This parameter is required.
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// This parameter is required.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
}

func (s ManageAccessorDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s ManageAccessorDomainRequest) GoString() string {
	return s.String()
}

func (s *ManageAccessorDomainRequest) SetCaller(v string) *ManageAccessorDomainRequest {
	s.Caller = &v
	return s
}

func (s *ManageAccessorDomainRequest) SetDomain(v string) *ManageAccessorDomainRequest {
	s.Domain = &v
	return s
}

func (s *ManageAccessorDomainRequest) SetOperation(v string) *ManageAccessorDomainRequest {
	s.Operation = &v
	return s
}

type ManageAccessorDomainResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ManageAccessorDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ManageAccessorDomainResponseBody) GoString() string {
	return s.String()
}

func (s *ManageAccessorDomainResponseBody) SetCode(v int32) *ManageAccessorDomainResponseBody {
	s.Code = &v
	return s
}

func (s *ManageAccessorDomainResponseBody) SetMessage(v string) *ManageAccessorDomainResponseBody {
	s.Message = &v
	return s
}

func (s *ManageAccessorDomainResponseBody) SetRequestId(v string) *ManageAccessorDomainResponseBody {
	s.RequestId = &v
	return s
}

type ManageAccessorDomainResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManageAccessorDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManageAccessorDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s ManageAccessorDomainResponse) GoString() string {
	return s.String()
}

func (s *ManageAccessorDomainResponse) SetHeaders(v map[string]*string) *ManageAccessorDomainResponse {
	s.Headers = v
	return s
}

func (s *ManageAccessorDomainResponse) SetStatusCode(v int32) *ManageAccessorDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *ManageAccessorDomainResponse) SetBody(v *ManageAccessorDomainResponseBody) *ManageAccessorDomainResponse {
	s.Body = v
	return s
}

type ManageAccessorDomainWhiteListRequest struct {
	// This parameter is required.
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// This parameter is required.
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	EndTime *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	Remark    *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ManageAccessorDomainWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s ManageAccessorDomainWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ManageAccessorDomainWhiteListRequest) SetCaller(v string) *ManageAccessorDomainWhiteListRequest {
	s.Caller = &v
	return s
}

func (s *ManageAccessorDomainWhiteListRequest) SetDomains(v []*string) *ManageAccessorDomainWhiteListRequest {
	s.Domains = v
	return s
}

func (s *ManageAccessorDomainWhiteListRequest) SetEndTime(v string) *ManageAccessorDomainWhiteListRequest {
	s.EndTime = &v
	return s
}

func (s *ManageAccessorDomainWhiteListRequest) SetOperation(v string) *ManageAccessorDomainWhiteListRequest {
	s.Operation = &v
	return s
}

func (s *ManageAccessorDomainWhiteListRequest) SetRemark(v string) *ManageAccessorDomainWhiteListRequest {
	s.Remark = &v
	return s
}

func (s *ManageAccessorDomainWhiteListRequest) SetStartTime(v string) *ManageAccessorDomainWhiteListRequest {
	s.StartTime = &v
	return s
}

type ManageAccessorDomainWhiteListResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ManageAccessorDomainWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ManageAccessorDomainWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ManageAccessorDomainWhiteListResponseBody) SetCode(v int32) *ManageAccessorDomainWhiteListResponseBody {
	s.Code = &v
	return s
}

func (s *ManageAccessorDomainWhiteListResponseBody) SetMessage(v string) *ManageAccessorDomainWhiteListResponseBody {
	s.Message = &v
	return s
}

func (s *ManageAccessorDomainWhiteListResponseBody) SetRequestId(v string) *ManageAccessorDomainWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type ManageAccessorDomainWhiteListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManageAccessorDomainWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManageAccessorDomainWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s ManageAccessorDomainWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ManageAccessorDomainWhiteListResponse) SetHeaders(v map[string]*string) *ManageAccessorDomainWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ManageAccessorDomainWhiteListResponse) SetStatusCode(v int32) *ManageAccessorDomainWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *ManageAccessorDomainWhiteListResponse) SetBody(v *ManageAccessorDomainWhiteListResponseBody) *ManageAccessorDomainWhiteListResponse {
	s.Body = v
	return s
}

type ManageAccessorIpRequest struct {
	// This parameter is required.
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// This parameter is required.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// This parameter is required.
	IpVersion *int32 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// This parameter is required.
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	Remark    *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ManageAccessorIpRequest) String() string {
	return tea.Prettify(s)
}

func (s ManageAccessorIpRequest) GoString() string {
	return s.String()
}

func (s *ManageAccessorIpRequest) SetCaller(v string) *ManageAccessorIpRequest {
	s.Caller = &v
	return s
}

func (s *ManageAccessorIpRequest) SetIp(v string) *ManageAccessorIpRequest {
	s.Ip = &v
	return s
}

func (s *ManageAccessorIpRequest) SetIpVersion(v int32) *ManageAccessorIpRequest {
	s.IpVersion = &v
	return s
}

func (s *ManageAccessorIpRequest) SetOperation(v string) *ManageAccessorIpRequest {
	s.Operation = &v
	return s
}

func (s *ManageAccessorIpRequest) SetRemark(v string) *ManageAccessorIpRequest {
	s.Remark = &v
	return s
}

type ManageAccessorIpResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ManageAccessorIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ManageAccessorIpResponseBody) GoString() string {
	return s.String()
}

func (s *ManageAccessorIpResponseBody) SetCode(v int32) *ManageAccessorIpResponseBody {
	s.Code = &v
	return s
}

func (s *ManageAccessorIpResponseBody) SetMessage(v string) *ManageAccessorIpResponseBody {
	s.Message = &v
	return s
}

func (s *ManageAccessorIpResponseBody) SetRequestId(v string) *ManageAccessorIpResponseBody {
	s.RequestId = &v
	return s
}

type ManageAccessorIpResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManageAccessorIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManageAccessorIpResponse) String() string {
	return tea.Prettify(s)
}

func (s ManageAccessorIpResponse) GoString() string {
	return s.String()
}

func (s *ManageAccessorIpResponse) SetHeaders(v map[string]*string) *ManageAccessorIpResponse {
	s.Headers = v
	return s
}

func (s *ManageAccessorIpResponse) SetStatusCode(v int32) *ManageAccessorIpResponse {
	s.StatusCode = &v
	return s
}

func (s *ManageAccessorIpResponse) SetBody(v *ManageAccessorIpResponseBody) *ManageAccessorIpResponse {
	s.Body = v
	return s
}

type QueryAccessorDomainRequest struct {
	// This parameter is required.
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// This parameter is required.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s QueryAccessorDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorDomainRequest) GoString() string {
	return s.String()
}

func (s *QueryAccessorDomainRequest) SetCaller(v string) *QueryAccessorDomainRequest {
	s.Caller = &v
	return s
}

func (s *QueryAccessorDomainRequest) SetDomain(v string) *QueryAccessorDomainRequest {
	s.Domain = &v
	return s
}

type QueryAccessorDomainResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAccessorDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorDomainResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccessorDomainResponseBody) SetCode(v int32) *QueryAccessorDomainResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAccessorDomainResponseBody) SetData(v bool) *QueryAccessorDomainResponseBody {
	s.Data = &v
	return s
}

func (s *QueryAccessorDomainResponseBody) SetMessage(v string) *QueryAccessorDomainResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAccessorDomainResponseBody) SetRequestId(v string) *QueryAccessorDomainResponseBody {
	s.RequestId = &v
	return s
}

type QueryAccessorDomainResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAccessorDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAccessorDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorDomainResponse) GoString() string {
	return s.String()
}

func (s *QueryAccessorDomainResponse) SetHeaders(v map[string]*string) *QueryAccessorDomainResponse {
	s.Headers = v
	return s
}

func (s *QueryAccessorDomainResponse) SetStatusCode(v int32) *QueryAccessorDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAccessorDomainResponse) SetBody(v *QueryAccessorDomainResponseBody) *QueryAccessorDomainResponse {
	s.Body = v
	return s
}

type QueryAccessorDomainListRequest struct {
	// This parameter is required.
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// This parameter is required.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryAccessorDomainListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorDomainListRequest) GoString() string {
	return s.String()
}

func (s *QueryAccessorDomainListRequest) SetCaller(v string) *QueryAccessorDomainListRequest {
	s.Caller = &v
	return s
}

func (s *QueryAccessorDomainListRequest) SetDomain(v string) *QueryAccessorDomainListRequest {
	s.Domain = &v
	return s
}

func (s *QueryAccessorDomainListRequest) SetPageNo(v int32) *QueryAccessorDomainListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryAccessorDomainListRequest) SetPageSize(v int32) *QueryAccessorDomainListRequest {
	s.PageSize = &v
	return s
}

type QueryAccessorDomainListResponseBody struct {
	Code      *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryAccessorDomainListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAccessorDomainListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorDomainListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccessorDomainListResponseBody) SetCode(v int32) *QueryAccessorDomainListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAccessorDomainListResponseBody) SetData(v *QueryAccessorDomainListResponseBodyData) *QueryAccessorDomainListResponseBody {
	s.Data = v
	return s
}

func (s *QueryAccessorDomainListResponseBody) SetMessage(v string) *QueryAccessorDomainListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAccessorDomainListResponseBody) SetRequestId(v string) *QueryAccessorDomainListResponseBody {
	s.RequestId = &v
	return s
}

type QueryAccessorDomainListResponseBodyData struct {
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s QueryAccessorDomainListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorDomainListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAccessorDomainListResponseBodyData) SetDomains(v []*string) *QueryAccessorDomainListResponseBodyData {
	s.Domains = v
	return s
}

type QueryAccessorDomainListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAccessorDomainListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAccessorDomainListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorDomainListResponse) GoString() string {
	return s.String()
}

func (s *QueryAccessorDomainListResponse) SetHeaders(v map[string]*string) *QueryAccessorDomainListResponse {
	s.Headers = v
	return s
}

func (s *QueryAccessorDomainListResponse) SetStatusCode(v int32) *QueryAccessorDomainListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAccessorDomainListResponse) SetBody(v *QueryAccessorDomainListResponseBody) *QueryAccessorDomainListResponse {
	s.Body = v
	return s
}

type QueryAccessorDomainStatusRequest struct {
	// This parameter is required.
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// This parameter is required.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s QueryAccessorDomainStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorDomainStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryAccessorDomainStatusRequest) SetCaller(v string) *QueryAccessorDomainStatusRequest {
	s.Caller = &v
	return s
}

func (s *QueryAccessorDomainStatusRequest) SetDomain(v string) *QueryAccessorDomainStatusRequest {
	s.Domain = &v
	return s
}

type QueryAccessorDomainStatusResponseBody struct {
	Code      *int32                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryAccessorDomainStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAccessorDomainStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorDomainStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccessorDomainStatusResponseBody) SetCode(v int32) *QueryAccessorDomainStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAccessorDomainStatusResponseBody) SetData(v *QueryAccessorDomainStatusResponseBodyData) *QueryAccessorDomainStatusResponseBody {
	s.Data = v
	return s
}

func (s *QueryAccessorDomainStatusResponseBody) SetMessage(v string) *QueryAccessorDomainStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAccessorDomainStatusResponseBody) SetRequestId(v string) *QueryAccessorDomainStatusResponseBody {
	s.RequestId = &v
	return s
}

type QueryAccessorDomainStatusResponseBodyData struct {
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Reason     *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	ReasonCode *int32  `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryAccessorDomainStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorDomainStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAccessorDomainStatusResponseBodyData) SetDomain(v string) *QueryAccessorDomainStatusResponseBodyData {
	s.Domain = &v
	return s
}

func (s *QueryAccessorDomainStatusResponseBodyData) SetReason(v string) *QueryAccessorDomainStatusResponseBodyData {
	s.Reason = &v
	return s
}

func (s *QueryAccessorDomainStatusResponseBodyData) SetReasonCode(v int32) *QueryAccessorDomainStatusResponseBodyData {
	s.ReasonCode = &v
	return s
}

func (s *QueryAccessorDomainStatusResponseBodyData) SetStatus(v string) *QueryAccessorDomainStatusResponseBodyData {
	s.Status = &v
	return s
}

type QueryAccessorDomainStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAccessorDomainStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAccessorDomainStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorDomainStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryAccessorDomainStatusResponse) SetHeaders(v map[string]*string) *QueryAccessorDomainStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryAccessorDomainStatusResponse) SetStatusCode(v int32) *QueryAccessorDomainStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAccessorDomainStatusResponse) SetBody(v *QueryAccessorDomainStatusResponseBody) *QueryAccessorDomainStatusResponse {
	s.Body = v
	return s
}

type QueryAccessorDomainWhiteListRequest struct {
	// This parameter is required.
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// This parameter is required.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s QueryAccessorDomainWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorDomainWhiteListRequest) GoString() string {
	return s.String()
}

func (s *QueryAccessorDomainWhiteListRequest) SetCaller(v string) *QueryAccessorDomainWhiteListRequest {
	s.Caller = &v
	return s
}

func (s *QueryAccessorDomainWhiteListRequest) SetDomain(v string) *QueryAccessorDomainWhiteListRequest {
	s.Domain = &v
	return s
}

type QueryAccessorDomainWhiteListResponseBody struct {
	Code      *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryAccessorDomainWhiteListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAccessorDomainWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorDomainWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccessorDomainWhiteListResponseBody) SetCode(v int32) *QueryAccessorDomainWhiteListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAccessorDomainWhiteListResponseBody) SetData(v *QueryAccessorDomainWhiteListResponseBodyData) *QueryAccessorDomainWhiteListResponseBody {
	s.Data = v
	return s
}

func (s *QueryAccessorDomainWhiteListResponseBody) SetMessage(v string) *QueryAccessorDomainWhiteListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAccessorDomainWhiteListResponseBody) SetRequestId(v string) *QueryAccessorDomainWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type QueryAccessorDomainWhiteListResponseBodyData struct {
	Items []*QueryAccessorDomainWhiteListResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	White *bool                                                `json:"White,omitempty" xml:"White,omitempty"`
}

func (s QueryAccessorDomainWhiteListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorDomainWhiteListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAccessorDomainWhiteListResponseBodyData) SetItems(v []*QueryAccessorDomainWhiteListResponseBodyDataItems) *QueryAccessorDomainWhiteListResponseBodyData {
	s.Items = v
	return s
}

func (s *QueryAccessorDomainWhiteListResponseBodyData) SetWhite(v bool) *QueryAccessorDomainWhiteListResponseBodyData {
	s.White = &v
	return s
}

type QueryAccessorDomainWhiteListResponseBodyDataItems struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Valid      *bool   `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s QueryAccessorDomainWhiteListResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorDomainWhiteListResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QueryAccessorDomainWhiteListResponseBodyDataItems) SetCreateTime(v string) *QueryAccessorDomainWhiteListResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *QueryAccessorDomainWhiteListResponseBodyDataItems) SetEndTime(v string) *QueryAccessorDomainWhiteListResponseBodyDataItems {
	s.EndTime = &v
	return s
}

func (s *QueryAccessorDomainWhiteListResponseBodyDataItems) SetStartTime(v string) *QueryAccessorDomainWhiteListResponseBodyDataItems {
	s.StartTime = &v
	return s
}

func (s *QueryAccessorDomainWhiteListResponseBodyDataItems) SetType(v string) *QueryAccessorDomainWhiteListResponseBodyDataItems {
	s.Type = &v
	return s
}

func (s *QueryAccessorDomainWhiteListResponseBodyDataItems) SetValid(v bool) *QueryAccessorDomainWhiteListResponseBodyDataItems {
	s.Valid = &v
	return s
}

type QueryAccessorDomainWhiteListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAccessorDomainWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAccessorDomainWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorDomainWhiteListResponse) GoString() string {
	return s.String()
}

func (s *QueryAccessorDomainWhiteListResponse) SetHeaders(v map[string]*string) *QueryAccessorDomainWhiteListResponse {
	s.Headers = v
	return s
}

func (s *QueryAccessorDomainWhiteListResponse) SetStatusCode(v int32) *QueryAccessorDomainWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAccessorDomainWhiteListResponse) SetBody(v *QueryAccessorDomainWhiteListResponseBody) *QueryAccessorDomainWhiteListResponse {
	s.Body = v
	return s
}

type QueryAccessorDomainsStatusRequest struct {
	// This parameter is required.
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// This parameter is required.
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s QueryAccessorDomainsStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorDomainsStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryAccessorDomainsStatusRequest) SetCaller(v string) *QueryAccessorDomainsStatusRequest {
	s.Caller = &v
	return s
}

func (s *QueryAccessorDomainsStatusRequest) SetDomains(v []*string) *QueryAccessorDomainsStatusRequest {
	s.Domains = v
	return s
}

type QueryAccessorDomainsStatusResponseBody struct {
	Code      *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*QueryAccessorDomainsStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAccessorDomainsStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorDomainsStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccessorDomainsStatusResponseBody) SetCode(v int32) *QueryAccessorDomainsStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAccessorDomainsStatusResponseBody) SetData(v []*QueryAccessorDomainsStatusResponseBodyData) *QueryAccessorDomainsStatusResponseBody {
	s.Data = v
	return s
}

func (s *QueryAccessorDomainsStatusResponseBody) SetMessage(v string) *QueryAccessorDomainsStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAccessorDomainsStatusResponseBody) SetRequestId(v string) *QueryAccessorDomainsStatusResponseBody {
	s.RequestId = &v
	return s
}

type QueryAccessorDomainsStatusResponseBodyData struct {
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Reason     *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	ReasonCode *int32  `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryAccessorDomainsStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorDomainsStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAccessorDomainsStatusResponseBodyData) SetDomain(v string) *QueryAccessorDomainsStatusResponseBodyData {
	s.Domain = &v
	return s
}

func (s *QueryAccessorDomainsStatusResponseBodyData) SetReason(v string) *QueryAccessorDomainsStatusResponseBodyData {
	s.Reason = &v
	return s
}

func (s *QueryAccessorDomainsStatusResponseBodyData) SetReasonCode(v int32) *QueryAccessorDomainsStatusResponseBodyData {
	s.ReasonCode = &v
	return s
}

func (s *QueryAccessorDomainsStatusResponseBodyData) SetStatus(v string) *QueryAccessorDomainsStatusResponseBodyData {
	s.Status = &v
	return s
}

type QueryAccessorDomainsStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAccessorDomainsStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAccessorDomainsStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorDomainsStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryAccessorDomainsStatusResponse) SetHeaders(v map[string]*string) *QueryAccessorDomainsStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryAccessorDomainsStatusResponse) SetStatusCode(v int32) *QueryAccessorDomainsStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAccessorDomainsStatusResponse) SetBody(v *QueryAccessorDomainsStatusResponseBody) *QueryAccessorDomainsStatusResponse {
	s.Body = v
	return s
}

type QueryAccessorIpRequest struct {
	// This parameter is required.
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// This parameter is required.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s QueryAccessorIpRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorIpRequest) GoString() string {
	return s.String()
}

func (s *QueryAccessorIpRequest) SetCaller(v string) *QueryAccessorIpRequest {
	s.Caller = &v
	return s
}

func (s *QueryAccessorIpRequest) SetIp(v string) *QueryAccessorIpRequest {
	s.Ip = &v
	return s
}

type QueryAccessorIpResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAccessorIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorIpResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccessorIpResponseBody) SetCode(v string) *QueryAccessorIpResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAccessorIpResponseBody) SetData(v bool) *QueryAccessorIpResponseBody {
	s.Data = &v
	return s
}

func (s *QueryAccessorIpResponseBody) SetMessage(v string) *QueryAccessorIpResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAccessorIpResponseBody) SetRequestId(v string) *QueryAccessorIpResponseBody {
	s.RequestId = &v
	return s
}

type QueryAccessorIpResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAccessorIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAccessorIpResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAccessorIpResponse) GoString() string {
	return s.String()
}

func (s *QueryAccessorIpResponse) SetHeaders(v map[string]*string) *QueryAccessorIpResponse {
	s.Headers = v
	return s
}

func (s *QueryAccessorIpResponse) SetStatusCode(v int32) *QueryAccessorIpResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAccessorIpResponse) SetBody(v *QueryAccessorIpResponseBody) *QueryAccessorIpResponse {
	s.Body = v
	return s
}

type SubmitAccessorFullDomainsOssListRequest struct {
	// This parameter is required.
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// This parameter is required.
	OssList []*string `json:"OssList,omitempty" xml:"OssList,omitempty" type:"Repeated"`
}

func (s SubmitAccessorFullDomainsOssListRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAccessorFullDomainsOssListRequest) GoString() string {
	return s.String()
}

func (s *SubmitAccessorFullDomainsOssListRequest) SetCaller(v string) *SubmitAccessorFullDomainsOssListRequest {
	s.Caller = &v
	return s
}

func (s *SubmitAccessorFullDomainsOssListRequest) SetOssList(v []*string) *SubmitAccessorFullDomainsOssListRequest {
	s.OssList = v
	return s
}

type SubmitAccessorFullDomainsOssListResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAccessorFullDomainsOssListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitAccessorFullDomainsOssListResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAccessorFullDomainsOssListResponseBody) SetCode(v int32) *SubmitAccessorFullDomainsOssListResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitAccessorFullDomainsOssListResponseBody) SetMessage(v string) *SubmitAccessorFullDomainsOssListResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitAccessorFullDomainsOssListResponseBody) SetRequestId(v string) *SubmitAccessorFullDomainsOssListResponseBody {
	s.RequestId = &v
	return s
}

type SubmitAccessorFullDomainsOssListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAccessorFullDomainsOssListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAccessorFullDomainsOssListResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitAccessorFullDomainsOssListResponse) GoString() string {
	return s.String()
}

func (s *SubmitAccessorFullDomainsOssListResponse) SetHeaders(v map[string]*string) *SubmitAccessorFullDomainsOssListResponse {
	s.Headers = v
	return s
}

func (s *SubmitAccessorFullDomainsOssListResponse) SetStatusCode(v int32) *SubmitAccessorFullDomainsOssListResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAccessorFullDomainsOssListResponse) SetBody(v *SubmitAccessorFullDomainsOssListResponseBody) *SubmitAccessorFullDomainsOssListResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("beian"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - DeleteUnbeianIpCheckTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUnbeianIpCheckTypeResponse
func (client *Client) DeleteUnbeianIpCheckTypeWithOptions(request *DeleteUnbeianIpCheckTypeRequest, runtime *util.RuntimeOptions) (_result *DeleteUnbeianIpCheckTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.CheckType)) {
		query["CheckType"] = request.CheckType
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUnbeianIpCheckType"),
		Version:     tea.String("2016-08-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUnbeianIpCheckTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteUnbeianIpCheckTypeRequest
//
// @return DeleteUnbeianIpCheckTypeResponse
func (client *Client) DeleteUnbeianIpCheckType(request *DeleteUnbeianIpCheckTypeRequest) (_result *DeleteUnbeianIpCheckTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUnbeianIpCheckTypeResponse{}
	_body, _err := client.DeleteUnbeianIpCheckTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取主域名接口
//
// @param request - GetMainDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMainDomainResponse
func (client *Client) GetMainDomainWithOptions(request *GetMainDomainRequest, runtime *util.RuntimeOptions) (_result *GetMainDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMainDomain"),
		Version:     tea.String("2016-08-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMainDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取主域名接口
//
// @param request - GetMainDomainRequest
//
// @return GetMainDomainResponse
func (client *Client) GetMainDomain(request *GetMainDomainRequest) (_result *GetMainDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMainDomainResponse{}
	_body, _err := client.GetMainDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - InsertUnbeianIpCheckTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertUnbeianIpCheckTypeResponse
func (client *Client) InsertUnbeianIpCheckTypeWithOptions(request *InsertUnbeianIpCheckTypeRequest, runtime *util.RuntimeOptions) (_result *InsertUnbeianIpCheckTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.CheckType)) {
		query["CheckType"] = request.CheckType
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InsertUnbeianIpCheckType"),
		Version:     tea.String("2016-08-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InsertUnbeianIpCheckTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - InsertUnbeianIpCheckTypeRequest
//
// @return InsertUnbeianIpCheckTypeResponse
func (client *Client) InsertUnbeianIpCheckType(request *InsertUnbeianIpCheckTypeRequest) (_result *InsertUnbeianIpCheckTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertUnbeianIpCheckTypeResponse{}
	_body, _err := client.InsertUnbeianIpCheckTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListUnbeianIpCheckTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUnbeianIpCheckTypeResponse
func (client *Client) ListUnbeianIpCheckTypeWithOptions(request *ListUnbeianIpCheckTypeRequest, runtime *util.RuntimeOptions) (_result *ListUnbeianIpCheckTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.CheckType)) {
		query["CheckType"] = request.CheckType
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUnbeianIpCheckType"),
		Version:     tea.String("2016-08-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUnbeianIpCheckTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListUnbeianIpCheckTypeRequest
//
// @return ListUnbeianIpCheckTypeResponse
func (client *Client) ListUnbeianIpCheckType(request *ListUnbeianIpCheckTypeRequest) (_result *ListUnbeianIpCheckTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUnbeianIpCheckTypeResponse{}
	_body, _err := client.ListUnbeianIpCheckTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 接入方服务域名上报接口
//
// @param request - ManageAccessorDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManageAccessorDomainResponse
func (client *Client) ManageAccessorDomainWithOptions(request *ManageAccessorDomainRequest, runtime *util.RuntimeOptions) (_result *ManageAccessorDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		query["Operation"] = request.Operation
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ManageAccessorDomain"),
		Version:     tea.String("2016-08-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ManageAccessorDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 接入方服务域名上报接口
//
// @param request - ManageAccessorDomainRequest
//
// @return ManageAccessorDomainResponse
func (client *Client) ManageAccessorDomain(request *ManageAccessorDomainRequest) (_result *ManageAccessorDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ManageAccessorDomainResponse{}
	_body, _err := client.ManageAccessorDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 接入方域名白名单上报接口
//
// @param request - ManageAccessorDomainWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManageAccessorDomainWhiteListResponse
func (client *Client) ManageAccessorDomainWhiteListWithOptions(request *ManageAccessorDomainWhiteListRequest, runtime *util.RuntimeOptions) (_result *ManageAccessorDomainWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.Domains)) {
		query["Domains"] = request.Domains
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		query["Operation"] = request.Operation
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ManageAccessorDomainWhiteList"),
		Version:     tea.String("2016-08-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ManageAccessorDomainWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 接入方域名白名单上报接口
//
// @param request - ManageAccessorDomainWhiteListRequest
//
// @return ManageAccessorDomainWhiteListResponse
func (client *Client) ManageAccessorDomainWhiteList(request *ManageAccessorDomainWhiteListRequest) (_result *ManageAccessorDomainWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ManageAccessorDomainWhiteListResponse{}
	_body, _err := client.ManageAccessorDomainWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 接入方管控IP上报接口
//
// @param request - ManageAccessorIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManageAccessorIpResponse
func (client *Client) ManageAccessorIpWithOptions(request *ManageAccessorIpRequest, runtime *util.RuntimeOptions) (_result *ManageAccessorIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		query["Operation"] = request.Operation
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ManageAccessorIp"),
		Version:     tea.String("2016-08-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ManageAccessorIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 接入方管控IP上报接口
//
// @param request - ManageAccessorIpRequest
//
// @return ManageAccessorIpResponse
func (client *Client) ManageAccessorIp(request *ManageAccessorIpRequest) (_result *ManageAccessorIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ManageAccessorIpResponse{}
	_body, _err := client.ManageAccessorIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 接入方服务域名是否上报查询接口
//
// @param request - QueryAccessorDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccessorDomainResponse
func (client *Client) QueryAccessorDomainWithOptions(request *QueryAccessorDomainRequest, runtime *util.RuntimeOptions) (_result *QueryAccessorDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAccessorDomain"),
		Version:     tea.String("2016-08-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAccessorDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 接入方服务域名是否上报查询接口
//
// @param request - QueryAccessorDomainRequest
//
// @return QueryAccessorDomainResponse
func (client *Client) QueryAccessorDomain(request *QueryAccessorDomainRequest) (_result *QueryAccessorDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAccessorDomainResponse{}
	_body, _err := client.QueryAccessorDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 接入方查询服务域名列表接口
//
// @param request - QueryAccessorDomainListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccessorDomainListResponse
func (client *Client) QueryAccessorDomainListWithOptions(request *QueryAccessorDomainListRequest, runtime *util.RuntimeOptions) (_result *QueryAccessorDomainListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAccessorDomainList"),
		Version:     tea.String("2016-08-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAccessorDomainListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 接入方查询服务域名列表接口
//
// @param request - QueryAccessorDomainListRequest
//
// @return QueryAccessorDomainListResponse
func (client *Client) QueryAccessorDomainList(request *QueryAccessorDomainListRequest) (_result *QueryAccessorDomainListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAccessorDomainListResponse{}
	_body, _err := client.QueryAccessorDomainListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 接入方域名状态查询接口
//
// @param request - QueryAccessorDomainStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccessorDomainStatusResponse
func (client *Client) QueryAccessorDomainStatusWithOptions(request *QueryAccessorDomainStatusRequest, runtime *util.RuntimeOptions) (_result *QueryAccessorDomainStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAccessorDomainStatus"),
		Version:     tea.String("2016-08-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAccessorDomainStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 接入方域名状态查询接口
//
// @param request - QueryAccessorDomainStatusRequest
//
// @return QueryAccessorDomainStatusResponse
func (client *Client) QueryAccessorDomainStatus(request *QueryAccessorDomainStatusRequest) (_result *QueryAccessorDomainStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAccessorDomainStatusResponse{}
	_body, _err := client.QueryAccessorDomainStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 接入方域名白名单上报查询接口
//
// @param request - QueryAccessorDomainWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccessorDomainWhiteListResponse
func (client *Client) QueryAccessorDomainWhiteListWithOptions(request *QueryAccessorDomainWhiteListRequest, runtime *util.RuntimeOptions) (_result *QueryAccessorDomainWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAccessorDomainWhiteList"),
		Version:     tea.String("2016-08-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAccessorDomainWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 接入方域名白名单上报查询接口
//
// @param request - QueryAccessorDomainWhiteListRequest
//
// @return QueryAccessorDomainWhiteListResponse
func (client *Client) QueryAccessorDomainWhiteList(request *QueryAccessorDomainWhiteListRequest) (_result *QueryAccessorDomainWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAccessorDomainWhiteListResponse{}
	_body, _err := client.QueryAccessorDomainWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 接入方域名状态批量查询接口
//
// @param request - QueryAccessorDomainsStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccessorDomainsStatusResponse
func (client *Client) QueryAccessorDomainsStatusWithOptions(request *QueryAccessorDomainsStatusRequest, runtime *util.RuntimeOptions) (_result *QueryAccessorDomainsStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.Domains)) {
		query["Domains"] = request.Domains
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAccessorDomainsStatus"),
		Version:     tea.String("2016-08-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAccessorDomainsStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 接入方域名状态批量查询接口
//
// @param request - QueryAccessorDomainsStatusRequest
//
// @return QueryAccessorDomainsStatusResponse
func (client *Client) QueryAccessorDomainsStatus(request *QueryAccessorDomainsStatusRequest) (_result *QueryAccessorDomainsStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAccessorDomainsStatusResponse{}
	_body, _err := client.QueryAccessorDomainsStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 接入方服务域名上报接口
//
// @param request - QueryAccessorIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccessorIpResponse
func (client *Client) QueryAccessorIpWithOptions(request *QueryAccessorIpRequest, runtime *util.RuntimeOptions) (_result *QueryAccessorIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAccessorIp"),
		Version:     tea.String("2016-08-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAccessorIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 接入方服务域名上报接口
//
// @param request - QueryAccessorIpRequest
//
// @return QueryAccessorIpResponse
func (client *Client) QueryAccessorIp(request *QueryAccessorIpRequest) (_result *QueryAccessorIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAccessorIpResponse{}
	_body, _err := client.QueryAccessorIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 接入方服务域名全量上报接口
//
// @param request - SubmitAccessorFullDomainsOssListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAccessorFullDomainsOssListResponse
func (client *Client) SubmitAccessorFullDomainsOssListWithOptions(request *SubmitAccessorFullDomainsOssListRequest, runtime *util.RuntimeOptions) (_result *SubmitAccessorFullDomainsOssListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.OssList)) {
		query["OssList"] = request.OssList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitAccessorFullDomainsOssList"),
		Version:     tea.String("2016-08-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitAccessorFullDomainsOssListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 接入方服务域名全量上报接口
//
// @param request - SubmitAccessorFullDomainsOssListRequest
//
// @return SubmitAccessorFullDomainsOssListResponse
func (client *Client) SubmitAccessorFullDomainsOssList(request *SubmitAccessorFullDomainsOssListRequest) (_result *SubmitAccessorFullDomainsOssListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitAccessorFullDomainsOssListResponse{}
	_body, _err := client.SubmitAccessorFullDomainsOssListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
