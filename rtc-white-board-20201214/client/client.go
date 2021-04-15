// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeAppsRequest struct {
	// 白板应用唯一标识符，默认查询所有应用ID
	AppID *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
	// 白板应用状态，默认查询所有状态。（取值：1:启用，2:停用）
	AppStatus *int64 `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	// 第几页，默认查询第1页。
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 每页显示个数，默认为10。
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppsRequest) SetAppID(v string) *DescribeAppsRequest {
	s.AppID = &v
	return s
}

func (s *DescribeAppsRequest) SetAppStatus(v int64) *DescribeAppsRequest {
	s.AppStatus = &v
	return s
}

func (s *DescribeAppsRequest) SetPageNum(v int64) *DescribeAppsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeAppsRequest) SetPageSize(v int64) *DescribeAppsRequest {
	s.PageSize = &v
	return s
}

type DescribeAppsResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求结果
	ResponseSuccess *bool `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// 返回结果体
	Result *DescribeAppsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBody) SetRequestId(v string) *DescribeAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppsResponseBody) SetResponseSuccess(v bool) *DescribeAppsResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *DescribeAppsResponseBody) SetErrorCode(v string) *DescribeAppsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeAppsResponseBody) SetErrorMsg(v string) *DescribeAppsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeAppsResponseBody) SetResult(v *DescribeAppsResponseBodyResult) *DescribeAppsResponseBody {
	s.Result = v
	return s
}

type DescribeAppsResponseBodyResult struct {
	TotalNum  *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	// App信息列表
	AppList []*DescribeAppsResponseBodyResultAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Repeated"`
}

func (s DescribeAppsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResult) SetTotalNum(v int64) *DescribeAppsResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetTotalPage(v int64) *DescribeAppsResponseBodyResult {
	s.TotalPage = &v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetAppList(v []*DescribeAppsResponseBodyResultAppList) *DescribeAppsResponseBodyResult {
	s.AppList = v
	return s
}

type DescribeAppsResponseBodyResultAppList struct {
	// 白板应用唯一标识符
	AppID *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
	// 白板应用名
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 白板应用状态（取值：1:启用，2:停用）
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 白板应用回调地址URL
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// 合法域名列表，使用英文逗号(,)分隔
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	// 白板应用创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s DescribeAppsResponseBodyResultAppList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResultAppList) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultAppList) SetAppID(v string) *DescribeAppsResponseBodyResultAppList {
	s.AppID = &v
	return s
}

func (s *DescribeAppsResponseBodyResultAppList) SetAppName(v string) *DescribeAppsResponseBodyResultAppList {
	s.AppName = &v
	return s
}

func (s *DescribeAppsResponseBodyResultAppList) SetStatus(v int64) *DescribeAppsResponseBodyResultAppList {
	s.Status = &v
	return s
}

func (s *DescribeAppsResponseBodyResultAppList) SetCallbackUrl(v string) *DescribeAppsResponseBodyResultAppList {
	s.CallbackUrl = &v
	return s
}

func (s *DescribeAppsResponseBodyResultAppList) SetDomainNames(v string) *DescribeAppsResponseBodyResultAppList {
	s.DomainNames = &v
	return s
}

func (s *DescribeAppsResponseBodyResultAppList) SetCreateTime(v string) *DescribeAppsResponseBodyResultAppList {
	s.CreateTime = &v
	return s
}

type DescribeAppsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponse) SetHeaders(v map[string]*string) *DescribeAppsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppsResponse) SetBody(v *DescribeAppsResponseBody) *DescribeAppsResponse {
	s.Body = v
	return s
}

type SetAppCallbackUrlRequest struct {
	// 白板应用唯一标识符
	AppID *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
	// 白板应用回调地址URL
	AppCallbackUrl *string `json:"AppCallbackUrl,omitempty" xml:"AppCallbackUrl,omitempty"`
	// 白板应用回调鉴权码，由8~20位大小写字母、数字或下划线组成
	AppCallbackAuthKey *string `json:"AppCallbackAuthKey,omitempty" xml:"AppCallbackAuthKey,omitempty"`
}

func (s SetAppCallbackUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAppCallbackUrlRequest) GoString() string {
	return s.String()
}

func (s *SetAppCallbackUrlRequest) SetAppID(v string) *SetAppCallbackUrlRequest {
	s.AppID = &v
	return s
}

func (s *SetAppCallbackUrlRequest) SetAppCallbackUrl(v string) *SetAppCallbackUrlRequest {
	s.AppCallbackUrl = &v
	return s
}

func (s *SetAppCallbackUrlRequest) SetAppCallbackAuthKey(v string) *SetAppCallbackUrlRequest {
	s.AppCallbackAuthKey = &v
	return s
}

type SetAppCallbackUrlResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求结果
	ResponseSuccess *bool `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// 返回结果
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SetAppCallbackUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAppCallbackUrlResponseBody) GoString() string {
	return s.String()
}

func (s *SetAppCallbackUrlResponseBody) SetRequestId(v string) *SetAppCallbackUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAppCallbackUrlResponseBody) SetResponseSuccess(v bool) *SetAppCallbackUrlResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *SetAppCallbackUrlResponseBody) SetErrorCode(v string) *SetAppCallbackUrlResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SetAppCallbackUrlResponseBody) SetErrorMsg(v string) *SetAppCallbackUrlResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SetAppCallbackUrlResponseBody) SetResult(v bool) *SetAppCallbackUrlResponseBody {
	s.Result = &v
	return s
}

type SetAppCallbackUrlResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetAppCallbackUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetAppCallbackUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAppCallbackUrlResponse) GoString() string {
	return s.String()
}

func (s *SetAppCallbackUrlResponse) SetHeaders(v map[string]*string) *SetAppCallbackUrlResponse {
	s.Headers = v
	return s
}

func (s *SetAppCallbackUrlResponse) SetBody(v *SetAppCallbackUrlResponseBody) *SetAppCallbackUrlResponse {
	s.Body = v
	return s
}

type SetAppNameRequest struct {
	// 白板应用唯一标识符
	AppID *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
	// 白板应用名，由不超过32位的中文、英文、数字或下划线组成
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s SetAppNameRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAppNameRequest) GoString() string {
	return s.String()
}

func (s *SetAppNameRequest) SetAppID(v string) *SetAppNameRequest {
	s.AppID = &v
	return s
}

func (s *SetAppNameRequest) SetAppName(v string) *SetAppNameRequest {
	s.AppName = &v
	return s
}

type SetAppNameResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求结果
	ResponseSuccess *bool `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// 返回结果
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SetAppNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAppNameResponseBody) GoString() string {
	return s.String()
}

func (s *SetAppNameResponseBody) SetRequestId(v string) *SetAppNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAppNameResponseBody) SetResponseSuccess(v bool) *SetAppNameResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *SetAppNameResponseBody) SetErrorCode(v string) *SetAppNameResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SetAppNameResponseBody) SetErrorMsg(v string) *SetAppNameResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SetAppNameResponseBody) SetResult(v bool) *SetAppNameResponseBody {
	s.Result = &v
	return s
}

type SetAppNameResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetAppNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetAppNameResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAppNameResponse) GoString() string {
	return s.String()
}

func (s *SetAppNameResponse) SetHeaders(v map[string]*string) *SetAppNameResponse {
	s.Headers = v
	return s
}

func (s *SetAppNameResponse) SetBody(v *SetAppNameResponseBody) *SetAppNameResponse {
	s.Body = v
	return s
}

type DescribeWhiteBoardsRequest struct {
	// 白板应用唯一标识符
	AppID *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
	// 白板文档状态，默认查询所有状态。（取值：1:启用，2:停用）
	DocStatus *int64 `json:"DocStatus,omitempty" xml:"DocStatus,omitempty"`
	// 第几页，默认查询第1页
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 每页显示个数，默认为10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeWhiteBoardsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhiteBoardsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWhiteBoardsRequest) SetAppID(v string) *DescribeWhiteBoardsRequest {
	s.AppID = &v
	return s
}

func (s *DescribeWhiteBoardsRequest) SetDocStatus(v int64) *DescribeWhiteBoardsRequest {
	s.DocStatus = &v
	return s
}

func (s *DescribeWhiteBoardsRequest) SetPageNum(v int64) *DescribeWhiteBoardsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeWhiteBoardsRequest) SetPageSize(v int64) *DescribeWhiteBoardsRequest {
	s.PageSize = &v
	return s
}

type DescribeWhiteBoardsResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求结果
	ResponseSuccess *bool `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// 返回结果体
	Result *DescribeWhiteBoardsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeWhiteBoardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhiteBoardsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWhiteBoardsResponseBody) SetRequestId(v string) *DescribeWhiteBoardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWhiteBoardsResponseBody) SetResponseSuccess(v bool) *DescribeWhiteBoardsResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *DescribeWhiteBoardsResponseBody) SetErrorCode(v string) *DescribeWhiteBoardsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeWhiteBoardsResponseBody) SetErrorMsg(v string) *DescribeWhiteBoardsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeWhiteBoardsResponseBody) SetResult(v *DescribeWhiteBoardsResponseBodyResult) *DescribeWhiteBoardsResponseBody {
	s.Result = v
	return s
}

type DescribeWhiteBoardsResponseBodyResult struct {
	TotalNum  *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	// App信息列表
	DocList []*DescribeWhiteBoardsResponseBodyResultDocList `json:"DocList,omitempty" xml:"DocList,omitempty" type:"Repeated"`
}

func (s DescribeWhiteBoardsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhiteBoardsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeWhiteBoardsResponseBodyResult) SetTotalNum(v int64) *DescribeWhiteBoardsResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *DescribeWhiteBoardsResponseBodyResult) SetTotalPage(v int64) *DescribeWhiteBoardsResponseBodyResult {
	s.TotalPage = &v
	return s
}

func (s *DescribeWhiteBoardsResponseBodyResult) SetDocList(v []*DescribeWhiteBoardsResponseBodyResultDocList) *DescribeWhiteBoardsResponseBodyResult {
	s.DocList = v
	return s
}

type DescribeWhiteBoardsResponseBodyResultDocList struct {
	// 白板应用唯一标识符
	AppID *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
	// 文档唯一标识符
	DocKey *string `json:"DocKey,omitempty" xml:"DocKey,omitempty"`
	// 文档状态（取值：1:启用，2:停用）
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 创建白板的用户ID
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// 白板应用创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s DescribeWhiteBoardsResponseBodyResultDocList) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhiteBoardsResponseBodyResultDocList) GoString() string {
	return s.String()
}

func (s *DescribeWhiteBoardsResponseBodyResultDocList) SetAppID(v string) *DescribeWhiteBoardsResponseBodyResultDocList {
	s.AppID = &v
	return s
}

func (s *DescribeWhiteBoardsResponseBodyResultDocList) SetDocKey(v string) *DescribeWhiteBoardsResponseBodyResultDocList {
	s.DocKey = &v
	return s
}

func (s *DescribeWhiteBoardsResponseBodyResultDocList) SetStatus(v int64) *DescribeWhiteBoardsResponseBodyResultDocList {
	s.Status = &v
	return s
}

func (s *DescribeWhiteBoardsResponseBodyResultDocList) SetCreateUserId(v string) *DescribeWhiteBoardsResponseBodyResultDocList {
	s.CreateUserId = &v
	return s
}

func (s *DescribeWhiteBoardsResponseBodyResultDocList) SetCreateTime(v string) *DescribeWhiteBoardsResponseBodyResultDocList {
	s.CreateTime = &v
	return s
}

type DescribeWhiteBoardsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWhiteBoardsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWhiteBoardsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWhiteBoardsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWhiteBoardsResponse) SetHeaders(v map[string]*string) *DescribeWhiteBoardsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWhiteBoardsResponse) SetBody(v *DescribeWhiteBoardsResponseBody) *DescribeWhiteBoardsResponse {
	s.Body = v
	return s
}

type SetAppDomainNamesRequest struct {
	// 白板应用唯一标识符
	AppID *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
	// 所有会使用到白板应用的客户网站域名，多个使用英文逗号(,)分隔，最多传10个
	AppDomainNames *string `json:"AppDomainNames,omitempty" xml:"AppDomainNames,omitempty"`
}

func (s SetAppDomainNamesRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAppDomainNamesRequest) GoString() string {
	return s.String()
}

func (s *SetAppDomainNamesRequest) SetAppID(v string) *SetAppDomainNamesRequest {
	s.AppID = &v
	return s
}

func (s *SetAppDomainNamesRequest) SetAppDomainNames(v string) *SetAppDomainNamesRequest {
	s.AppDomainNames = &v
	return s
}

type SetAppDomainNamesResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求结果
	ResponseSuccess *bool `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// 返回结果
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SetAppDomainNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAppDomainNamesResponseBody) GoString() string {
	return s.String()
}

func (s *SetAppDomainNamesResponseBody) SetRequestId(v string) *SetAppDomainNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAppDomainNamesResponseBody) SetResponseSuccess(v bool) *SetAppDomainNamesResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *SetAppDomainNamesResponseBody) SetErrorCode(v string) *SetAppDomainNamesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SetAppDomainNamesResponseBody) SetErrorMsg(v string) *SetAppDomainNamesResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SetAppDomainNamesResponseBody) SetResult(v bool) *SetAppDomainNamesResponseBody {
	s.Result = &v
	return s
}

type SetAppDomainNamesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetAppDomainNamesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetAppDomainNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAppDomainNamesResponse) GoString() string {
	return s.String()
}

func (s *SetAppDomainNamesResponse) SetHeaders(v map[string]*string) *SetAppDomainNamesResponse {
	s.Headers = v
	return s
}

func (s *SetAppDomainNamesResponse) SetBody(v *SetAppDomainNamesResponseBody) *SetAppDomainNamesResponse {
	s.Body = v
	return s
}

type OpenWhiteBoardRequest struct {
	// 白板应用唯一标识符
	AppID *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
	// 打开白板的用户ID（客户业务用户），由纯数字组成
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 白板文档唯一标识符
	DocKey *string `json:"DocKey,omitempty" xml:"DocKey,omitempty"`
}

func (s OpenWhiteBoardRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenWhiteBoardRequest) GoString() string {
	return s.String()
}

func (s *OpenWhiteBoardRequest) SetAppID(v string) *OpenWhiteBoardRequest {
	s.AppID = &v
	return s
}

func (s *OpenWhiteBoardRequest) SetUserId(v string) *OpenWhiteBoardRequest {
	s.UserId = &v
	return s
}

func (s *OpenWhiteBoardRequest) SetDocKey(v string) *OpenWhiteBoardRequest {
	s.DocKey = &v
	return s
}

type OpenWhiteBoardResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求结果
	ResponseSuccess *bool `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// 返回结果体
	Result *OpenWhiteBoardResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s OpenWhiteBoardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenWhiteBoardResponseBody) GoString() string {
	return s.String()
}

func (s *OpenWhiteBoardResponseBody) SetRequestId(v string) *OpenWhiteBoardResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenWhiteBoardResponseBody) SetResponseSuccess(v bool) *OpenWhiteBoardResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *OpenWhiteBoardResponseBody) SetErrorCode(v string) *OpenWhiteBoardResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *OpenWhiteBoardResponseBody) SetErrorMsg(v string) *OpenWhiteBoardResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *OpenWhiteBoardResponseBody) SetResult(v *OpenWhiteBoardResponseBodyResult) *OpenWhiteBoardResponseBody {
	s.Result = v
	return s
}

type OpenWhiteBoardResponseBodyResult struct {
	// 白板连接信息
	DocumentAccessInfo *OpenWhiteBoardResponseBodyResultDocumentAccessInfo `json:"DocumentAccessInfo,omitempty" xml:"DocumentAccessInfo,omitempty" type:"Struct"`
}

func (s OpenWhiteBoardResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s OpenWhiteBoardResponseBodyResult) GoString() string {
	return s.String()
}

func (s *OpenWhiteBoardResponseBodyResult) SetDocumentAccessInfo(v *OpenWhiteBoardResponseBodyResultDocumentAccessInfo) *OpenWhiteBoardResponseBodyResult {
	s.DocumentAccessInfo = v
	return s
}

type OpenWhiteBoardResponseBodyResultDocumentAccessInfo struct {
	// 连接签名
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// 白板长连接地址
	CollabHost *string `json:"CollabHost,omitempty" xml:"CollabHost,omitempty"`
	// 权限码，取值：0:无权限，1:只读，2:读写
	Permission *int64 `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// 用户信息
	UserInfo *OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
	// 新协议长连接服务域名
	WsDomain *string `json:"WsDomain,omitempty" xml:"WsDomain,omitempty"`
}

func (s OpenWhiteBoardResponseBodyResultDocumentAccessInfo) String() string {
	return tea.Prettify(s)
}

func (s OpenWhiteBoardResponseBodyResultDocumentAccessInfo) GoString() string {
	return s.String()
}

func (s *OpenWhiteBoardResponseBodyResultDocumentAccessInfo) SetAccessToken(v string) *OpenWhiteBoardResponseBodyResultDocumentAccessInfo {
	s.AccessToken = &v
	return s
}

func (s *OpenWhiteBoardResponseBodyResultDocumentAccessInfo) SetCollabHost(v string) *OpenWhiteBoardResponseBodyResultDocumentAccessInfo {
	s.CollabHost = &v
	return s
}

func (s *OpenWhiteBoardResponseBodyResultDocumentAccessInfo) SetPermission(v int64) *OpenWhiteBoardResponseBodyResultDocumentAccessInfo {
	s.Permission = &v
	return s
}

func (s *OpenWhiteBoardResponseBodyResultDocumentAccessInfo) SetUserInfo(v *OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo) *OpenWhiteBoardResponseBodyResultDocumentAccessInfo {
	s.UserInfo = v
	return s
}

func (s *OpenWhiteBoardResponseBodyResultDocumentAccessInfo) SetWsDomain(v string) *OpenWhiteBoardResponseBodyResultDocumentAccessInfo {
	s.WsDomain = &v
	return s
}

type OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo struct {
	// 用户头像的URL
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	// 用户昵称
	Nick *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	// 用户的拼音昵称
	NickPinyin *string `json:"NickPinyin,omitempty" xml:"NickPinyin,omitempty"`
	// 打开白板的用户ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo) String() string {
	return tea.Prettify(s)
}

func (s OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo) GoString() string {
	return s.String()
}

func (s *OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo) SetAvatarUrl(v string) *OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo {
	s.AvatarUrl = &v
	return s
}

func (s *OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo) SetNick(v string) *OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo {
	s.Nick = &v
	return s
}

func (s *OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo) SetNickPinyin(v string) *OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo {
	s.NickPinyin = &v
	return s
}

func (s *OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo) SetUserId(v string) *OpenWhiteBoardResponseBodyResultDocumentAccessInfoUserInfo {
	s.UserId = &v
	return s
}

type OpenWhiteBoardResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenWhiteBoardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenWhiteBoardResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenWhiteBoardResponse) GoString() string {
	return s.String()
}

func (s *OpenWhiteBoardResponse) SetHeaders(v map[string]*string) *OpenWhiteBoardResponse {
	s.Headers = v
	return s
}

func (s *OpenWhiteBoardResponse) SetBody(v *OpenWhiteBoardResponseBody) *OpenWhiteBoardResponse {
	s.Body = v
	return s
}

type RefreshUsersPermissionsRequest struct {
	// 需要刷新权限的用户ID，多个用英文逗号（,）分隔，最多30个，每个ID由纯数字组成
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
	// 白板文档唯一标识符
	DocKey *string `json:"DocKey,omitempty" xml:"DocKey,omitempty"`
	// 白板应用唯一标识符
	AppID *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
}

func (s RefreshUsersPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshUsersPermissionsRequest) GoString() string {
	return s.String()
}

func (s *RefreshUsersPermissionsRequest) SetUserIds(v string) *RefreshUsersPermissionsRequest {
	s.UserIds = &v
	return s
}

func (s *RefreshUsersPermissionsRequest) SetDocKey(v string) *RefreshUsersPermissionsRequest {
	s.DocKey = &v
	return s
}

func (s *RefreshUsersPermissionsRequest) SetAppID(v string) *RefreshUsersPermissionsRequest {
	s.AppID = &v
	return s
}

type RefreshUsersPermissionsResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求结果
	ResponseSuccess *bool `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// 返回结果
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RefreshUsersPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshUsersPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshUsersPermissionsResponseBody) SetRequestId(v string) *RefreshUsersPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshUsersPermissionsResponseBody) SetResponseSuccess(v bool) *RefreshUsersPermissionsResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *RefreshUsersPermissionsResponseBody) SetErrorCode(v string) *RefreshUsersPermissionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RefreshUsersPermissionsResponseBody) SetErrorMsg(v string) *RefreshUsersPermissionsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RefreshUsersPermissionsResponseBody) SetResult(v bool) *RefreshUsersPermissionsResponseBody {
	s.Result = &v
	return s
}

type RefreshUsersPermissionsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefreshUsersPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshUsersPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshUsersPermissionsResponse) GoString() string {
	return s.String()
}

func (s *RefreshUsersPermissionsResponse) SetHeaders(v map[string]*string) *RefreshUsersPermissionsResponse {
	s.Headers = v
	return s
}

func (s *RefreshUsersPermissionsResponse) SetBody(v *RefreshUsersPermissionsResponseBody) *RefreshUsersPermissionsResponse {
	s.Body = v
	return s
}

type CreateAppRequest struct {
	// 白板应用名，由不超过32位的中文、英文、数字或下划线组成
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s CreateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) SetAppName(v string) *CreateAppRequest {
	s.AppName = &v
	return s
}

type CreateAppResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求结果
	ResponseSuccess *bool `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// 返回结果体
	Result *CreateAppResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResponseBody) SetResponseSuccess(v bool) *CreateAppResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *CreateAppResponseBody) SetErrorCode(v string) *CreateAppResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateAppResponseBody) SetErrorMsg(v string) *CreateAppResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateAppResponseBody) SetResult(v *CreateAppResponseBodyResult) *CreateAppResponseBody {
	s.Result = v
	return s
}

type CreateAppResponseBodyResult struct {
	// 白板应用唯一标识符
	AppID *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
}

func (s CreateAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResult) SetAppID(v string) *CreateAppResponseBodyResult {
	s.AppID = &v
	return s
}

type CreateAppResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponse) GoString() string {
	return s.String()
}

func (s *CreateAppResponse) SetHeaders(v map[string]*string) *CreateAppResponse {
	s.Headers = v
	return s
}

func (s *CreateAppResponse) SetBody(v *CreateAppResponseBody) *CreateAppResponse {
	s.Body = v
	return s
}

type CreateWhiteBoardRequest struct {
	// 创建白板的用户ID（客户业务用户），由纯数字组成。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 白板应用唯一标识符
	AppID *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
}

func (s CreateWhiteBoardRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWhiteBoardRequest) GoString() string {
	return s.String()
}

func (s *CreateWhiteBoardRequest) SetUserId(v string) *CreateWhiteBoardRequest {
	s.UserId = &v
	return s
}

func (s *CreateWhiteBoardRequest) SetAppID(v string) *CreateWhiteBoardRequest {
	s.AppID = &v
	return s
}

type CreateWhiteBoardResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求结果
	ResponseSuccess *bool `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// 返回结果体
	Result *CreateWhiteBoardResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateWhiteBoardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWhiteBoardResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWhiteBoardResponseBody) SetRequestId(v string) *CreateWhiteBoardResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWhiteBoardResponseBody) SetResponseSuccess(v bool) *CreateWhiteBoardResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *CreateWhiteBoardResponseBody) SetErrorCode(v string) *CreateWhiteBoardResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateWhiteBoardResponseBody) SetErrorMsg(v string) *CreateWhiteBoardResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateWhiteBoardResponseBody) SetResult(v *CreateWhiteBoardResponseBodyResult) *CreateWhiteBoardResponseBody {
	s.Result = v
	return s
}

type CreateWhiteBoardResponseBodyResult struct {
	// 文档唯一标识符，由大小写字母和数字组成
	DocKey *string `json:"DocKey,omitempty" xml:"DocKey,omitempty"`
}

func (s CreateWhiteBoardResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateWhiteBoardResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateWhiteBoardResponseBodyResult) SetDocKey(v string) *CreateWhiteBoardResponseBodyResult {
	s.DocKey = &v
	return s
}

type CreateWhiteBoardResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateWhiteBoardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWhiteBoardResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWhiteBoardResponse) GoString() string {
	return s.String()
}

func (s *CreateWhiteBoardResponse) SetHeaders(v map[string]*string) *CreateWhiteBoardResponse {
	s.Headers = v
	return s
}

func (s *CreateWhiteBoardResponse) SetBody(v *CreateWhiteBoardResponseBody) *CreateWhiteBoardResponse {
	s.Body = v
	return s
}

type SetAppStatusRequest struct {
	// 白板应用唯一标识符
	AppID *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
	// 白板应用状态（取值：1:启用，2:停用）
	AppStatus *int64 `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
}

func (s SetAppStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAppStatusRequest) GoString() string {
	return s.String()
}

func (s *SetAppStatusRequest) SetAppID(v string) *SetAppStatusRequest {
	s.AppID = &v
	return s
}

func (s *SetAppStatusRequest) SetAppStatus(v int64) *SetAppStatusRequest {
	s.AppStatus = &v
	return s
}

type SetAppStatusResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求结果
	ResponseSuccess *bool `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// 返回结果
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SetAppStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAppStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetAppStatusResponseBody) SetRequestId(v string) *SetAppStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAppStatusResponseBody) SetResponseSuccess(v bool) *SetAppStatusResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *SetAppStatusResponseBody) SetErrorCode(v string) *SetAppStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SetAppStatusResponseBody) SetErrorMsg(v string) *SetAppStatusResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SetAppStatusResponseBody) SetResult(v bool) *SetAppStatusResponseBody {
	s.Result = &v
	return s
}

type SetAppStatusResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetAppStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetAppStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAppStatusResponse) GoString() string {
	return s.String()
}

func (s *SetAppStatusResponse) SetHeaders(v map[string]*string) *SetAppStatusResponse {
	s.Headers = v
	return s
}

func (s *SetAppStatusResponse) SetBody(v *SetAppStatusResponseBody) *SetAppStatusResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("rtc-white-board"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) DescribeAppsWithOptions(request *DescribeAppsRequest, runtime *util.RuntimeOptions) (_result *DescribeAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAppsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeApps"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApps(request *DescribeAppsRequest) (_result *DescribeAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppsResponse{}
	_body, _err := client.DescribeAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAppCallbackUrlWithOptions(request *SetAppCallbackUrlRequest, runtime *util.RuntimeOptions) (_result *SetAppCallbackUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetAppCallbackUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetAppCallbackUrl"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAppCallbackUrl(request *SetAppCallbackUrlRequest) (_result *SetAppCallbackUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAppCallbackUrlResponse{}
	_body, _err := client.SetAppCallbackUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAppNameWithOptions(request *SetAppNameRequest, runtime *util.RuntimeOptions) (_result *SetAppNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetAppNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetAppName"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAppName(request *SetAppNameRequest) (_result *SetAppNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAppNameResponse{}
	_body, _err := client.SetAppNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWhiteBoardsWithOptions(request *DescribeWhiteBoardsRequest, runtime *util.RuntimeOptions) (_result *DescribeWhiteBoardsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWhiteBoardsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWhiteBoards"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWhiteBoards(request *DescribeWhiteBoardsRequest) (_result *DescribeWhiteBoardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWhiteBoardsResponse{}
	_body, _err := client.DescribeWhiteBoardsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAppDomainNamesWithOptions(request *SetAppDomainNamesRequest, runtime *util.RuntimeOptions) (_result *SetAppDomainNamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetAppDomainNamesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetAppDomainNames"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAppDomainNames(request *SetAppDomainNamesRequest) (_result *SetAppDomainNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAppDomainNamesResponse{}
	_body, _err := client.SetAppDomainNamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenWhiteBoardWithOptions(request *OpenWhiteBoardRequest, runtime *util.RuntimeOptions) (_result *OpenWhiteBoardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OpenWhiteBoardResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpenWhiteBoard"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenWhiteBoard(request *OpenWhiteBoardRequest) (_result *OpenWhiteBoardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenWhiteBoardResponse{}
	_body, _err := client.OpenWhiteBoardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshUsersPermissionsWithOptions(request *RefreshUsersPermissionsRequest, runtime *util.RuntimeOptions) (_result *RefreshUsersPermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RefreshUsersPermissionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RefreshUsersPermissions"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshUsersPermissions(request *RefreshUsersPermissionsRequest) (_result *RefreshUsersPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshUsersPermissionsResponse{}
	_body, _err := client.RefreshUsersPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppWithOptions(request *CreateAppRequest, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateApp"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWhiteBoardWithOptions(request *CreateWhiteBoardRequest, runtime *util.RuntimeOptions) (_result *CreateWhiteBoardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateWhiteBoardResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateWhiteBoard"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWhiteBoard(request *CreateWhiteBoardRequest) (_result *CreateWhiteBoardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWhiteBoardResponse{}
	_body, _err := client.CreateWhiteBoardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAppStatusWithOptions(request *SetAppStatusRequest, runtime *util.RuntimeOptions) (_result *SetAppStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetAppStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetAppStatus"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAppStatus(request *SetAppStatusRequest) (_result *SetAppStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAppStatusResponse{}
	_body, _err := client.SetAppStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
