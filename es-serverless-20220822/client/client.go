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

type DataStreamMapping struct {
	CaseSensitive   *bool                `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	DocValues       *bool                `json:"docValues,omitempty" xml:"docValues,omitempty"`
	Index           *bool                `json:"index,omitempty" xml:"index,omitempty"`
	Key             *string              `json:"key,omitempty" xml:"key,omitempty"`
	Properties      []*DataStreamMapping `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
	TokenizeOnChars []*string            `json:"tokenizeOnChars,omitempty" xml:"tokenizeOnChars,omitempty" type:"Repeated"`
	Type            *string              `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DataStreamMapping) String() string {
	return tea.Prettify(s)
}

func (s DataStreamMapping) GoString() string {
	return s.String()
}

func (s *DataStreamMapping) SetCaseSensitive(v bool) *DataStreamMapping {
	s.CaseSensitive = &v
	return s
}

func (s *DataStreamMapping) SetDocValues(v bool) *DataStreamMapping {
	s.DocValues = &v
	return s
}

func (s *DataStreamMapping) SetIndex(v bool) *DataStreamMapping {
	s.Index = &v
	return s
}

func (s *DataStreamMapping) SetKey(v string) *DataStreamMapping {
	s.Key = &v
	return s
}

func (s *DataStreamMapping) SetProperties(v []*DataStreamMapping) *DataStreamMapping {
	s.Properties = v
	return s
}

func (s *DataStreamMapping) SetTokenizeOnChars(v []*string) *DataStreamMapping {
	s.TokenizeOnChars = v
	return s
}

func (s *DataStreamMapping) SetType(v string) *DataStreamMapping {
	s.Type = &v
	return s
}

type CreateAppRequest struct {
	// 应用名
	AppName    *string `json:"appName,omitempty" xml:"appName,omitempty"`
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// 应用备注
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
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

func (s *CreateAppRequest) SetChargeType(v string) *CreateAppRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateAppRequest) SetDescription(v string) *CreateAppRequest {
	s.Description = &v
	return s
}

type CreateAppResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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

type CreateAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateAppResponse) SetStatusCode(v int32) *CreateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppResponse) SetBody(v *CreateAppResponseBody) *CreateAppResponse {
	s.Body = v
	return s
}

type CreateDataStreamRequest struct {
	// 代表资源名称的资源属性字段
	DataStreamName *string `json:"dataStreamName,omitempty" xml:"dataStreamName,omitempty"`
	// 删除时间
	DeletePhase *string `json:"deletePhase,omitempty" xml:"deletePhase,omitempty"`
	Dynamic     *string `json:"dynamic,omitempty" xml:"dynamic,omitempty"`
	RegionId    *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// 数据流模板
	Template     *CreateDataStreamRequestTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	TimeStampKey *string                          `json:"timeStampKey,omitempty" xml:"timeStampKey,omitempty"`
	// 数据流类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateDataStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataStreamRequest) GoString() string {
	return s.String()
}

func (s *CreateDataStreamRequest) SetDataStreamName(v string) *CreateDataStreamRequest {
	s.DataStreamName = &v
	return s
}

func (s *CreateDataStreamRequest) SetDeletePhase(v string) *CreateDataStreamRequest {
	s.DeletePhase = &v
	return s
}

func (s *CreateDataStreamRequest) SetDynamic(v string) *CreateDataStreamRequest {
	s.Dynamic = &v
	return s
}

func (s *CreateDataStreamRequest) SetRegionId(v string) *CreateDataStreamRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDataStreamRequest) SetTemplate(v *CreateDataStreamRequestTemplate) *CreateDataStreamRequest {
	s.Template = v
	return s
}

func (s *CreateDataStreamRequest) SetTimeStampKey(v string) *CreateDataStreamRequest {
	s.TimeStampKey = &v
	return s
}

func (s *CreateDataStreamRequest) SetType(v string) *CreateDataStreamRequest {
	s.Type = &v
	return s
}

type CreateDataStreamRequestTemplate struct {
	// 索引字段设置
	Mappings []*DataStreamMapping `json:"mappings,omitempty" xml:"mappings,omitempty" type:"Repeated"`
}

func (s CreateDataStreamRequestTemplate) String() string {
	return tea.Prettify(s)
}

func (s CreateDataStreamRequestTemplate) GoString() string {
	return s.String()
}

func (s *CreateDataStreamRequestTemplate) SetMappings(v []*DataStreamMapping) *CreateDataStreamRequestTemplate {
	s.Mappings = v
	return s
}

type CreateDataStreamResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateDataStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataStreamResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataStreamResponseBody) SetRequestId(v string) *CreateDataStreamResponseBody {
	s.RequestId = &v
	return s
}

type CreateDataStreamResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDataStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDataStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataStreamResponse) GoString() string {
	return s.String()
}

func (s *CreateDataStreamResponse) SetHeaders(v map[string]*string) *CreateDataStreamResponse {
	s.Headers = v
	return s
}

func (s *CreateDataStreamResponse) SetStatusCode(v int32) *CreateDataStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataStreamResponse) SetBody(v *CreateDataStreamResponseBody) *CreateDataStreamResponse {
	s.Body = v
	return s
}

type DeleteAccessTokenResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteAccessTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessTokenResponseBody) SetRequestId(v string) *DeleteAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccessTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAccessTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessTokenResponse) SetHeaders(v map[string]*string) *DeleteAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessTokenResponse) SetStatusCode(v int32) *DeleteAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessTokenResponse) SetBody(v *DeleteAccessTokenResponseBody) *DeleteAccessTokenResponse {
	s.Body = v
	return s
}

type DeleteAppResponseBody struct {
	RequestId *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DeleteAppResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DeleteAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBody) SetRequestId(v string) *DeleteAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppResponseBody) SetResult(v *DeleteAppResponseBodyResult) *DeleteAppResponseBody {
	s.Result = v
	return s
}

type DeleteAppResponseBodyResult struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s DeleteAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBodyResult) SetInstanceId(v string) *DeleteAppResponseBodyResult {
	s.InstanceId = &v
	return s
}

type DeleteAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppResponse) SetHeaders(v map[string]*string) *DeleteAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppResponse) SetStatusCode(v int32) *DeleteAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppResponse) SetBody(v *DeleteAppResponseBody) *DeleteAppResponse {
	s.Body = v
	return s
}

type DeleteDataStreamResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteDataStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataStreamResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataStreamResponseBody) SetRequestId(v string) *DeleteDataStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataStreamResponseBody) SetResult(v bool) *DeleteDataStreamResponseBody {
	s.Result = &v
	return s
}

type DeleteDataStreamResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDataStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDataStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataStreamResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataStreamResponse) SetHeaders(v map[string]*string) *DeleteDataStreamResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataStreamResponse) SetStatusCode(v int32) *DeleteDataStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataStreamResponse) SetBody(v *DeleteDataStreamResponseBody) *DeleteDataStreamResponse {
	s.Body = v
	return s
}

type DescibeRegionsResponseBody struct {
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*DescibeRegionsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s DescibeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescibeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescibeRegionsResponseBody) SetRequestId(v string) *DescibeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescibeRegionsResponseBody) SetResult(v []*DescibeRegionsResponseBodyResult) *DescibeRegionsResponseBody {
	s.Result = v
	return s
}

type DescibeRegionsResponseBodyResult struct {
	Endpoint  *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	LocalName *string `json:"localName,omitempty" xml:"localName,omitempty"`
	RegionId  *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DescibeRegionsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescibeRegionsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescibeRegionsResponseBodyResult) SetEndpoint(v string) *DescibeRegionsResponseBodyResult {
	s.Endpoint = &v
	return s
}

func (s *DescibeRegionsResponseBodyResult) SetLocalName(v string) *DescibeRegionsResponseBodyResult {
	s.LocalName = &v
	return s
}

func (s *DescibeRegionsResponseBodyResult) SetRegionId(v string) *DescibeRegionsResponseBodyResult {
	s.RegionId = &v
	return s
}

type DescibeRegionsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescibeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescibeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescibeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescibeRegionsResponse) SetHeaders(v map[string]*string) *DescibeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescibeRegionsResponse) SetStatusCode(v int32) *DescibeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescibeRegionsResponse) SetBody(v *DescibeRegionsResponseBody) *DescibeRegionsResponse {
	s.Body = v
	return s
}

type DisableAccessTokenResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DisableAccessTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAccessTokenResponseBody) SetRequestId(v string) *DisableAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

type DisableAccessTokenResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableAccessTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *DisableAccessTokenResponse) SetHeaders(v map[string]*string) *DisableAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *DisableAccessTokenResponse) SetStatusCode(v int32) *DisableAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAccessTokenResponse) SetBody(v *DisableAccessTokenResponseBody) *DisableAccessTokenResponse {
	s.Body = v
	return s
}

type EnableAccessTokenResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s EnableAccessTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *EnableAccessTokenResponseBody) SetRequestId(v string) *EnableAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

type EnableAccessTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableAccessTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *EnableAccessTokenResponse) SetHeaders(v map[string]*string) *EnableAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *EnableAccessTokenResponse) SetStatusCode(v int32) *EnableAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableAccessTokenResponse) SetBody(v *EnableAccessTokenResponseBody) *EnableAccessTokenResponse {
	s.Body = v
	return s
}

type GenerateAcccessTokenResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GenerateAcccessTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateAcccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateAcccessTokenResponseBody) SetRequestId(v string) *GenerateAcccessTokenResponseBody {
	s.RequestId = &v
	return s
}

type GenerateAcccessTokenResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateAcccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateAcccessTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateAcccessTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateAcccessTokenResponse) SetHeaders(v map[string]*string) *GenerateAcccessTokenResponse {
	s.Headers = v
	return s
}

func (s *GenerateAcccessTokenResponse) SetStatusCode(v int32) *GenerateAcccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateAcccessTokenResponse) SetBody(v *GenerateAcccessTokenResponseBody) *GenerateAcccessTokenResponse {
	s.Body = v
	return s
}

type GetAppResponseBody struct {
	RequestId *string                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetAppResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppResponseBody) SetRequestId(v string) *GetAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppResponseBody) SetResult(v *GetAppResponseBodyResult) *GetAppResponseBody {
	s.Result = v
	return s
}

type GetAppResponseBodyResult struct {
	// 代表资源名称的资源属性字段
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// 代表创建时间的资源属性字段
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 应用备注
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 流量限流额
	FlowQuota *string `json:"flowQuota,omitempty" xml:"flowQuota,omitempty"`
	// OwnerID账号ID
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// 代表region的资源属性字段
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// 代表资源状态的资源属性字段
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 存储限流额
	StorageQuota *string `json:"storageQuota,omitempty" xml:"storageQuota,omitempty"`
}

func (s GetAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyResult) SetAppName(v string) *GetAppResponseBodyResult {
	s.AppName = &v
	return s
}

func (s *GetAppResponseBodyResult) SetCreateTime(v string) *GetAppResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetAppResponseBodyResult) SetDescription(v string) *GetAppResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetAppResponseBodyResult) SetFlowQuota(v string) *GetAppResponseBodyResult {
	s.FlowQuota = &v
	return s
}

func (s *GetAppResponseBodyResult) SetOwnerId(v string) *GetAppResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *GetAppResponseBodyResult) SetRegionId(v string) *GetAppResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *GetAppResponseBodyResult) SetStatus(v string) *GetAppResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetAppResponseBodyResult) SetStorageQuota(v string) *GetAppResponseBodyResult {
	s.StorageQuota = &v
	return s
}

type GetAppResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponse) GoString() string {
	return s.String()
}

func (s *GetAppResponse) SetHeaders(v map[string]*string) *GetAppResponse {
	s.Headers = v
	return s
}

func (s *GetAppResponse) SetStatusCode(v int32) *GetAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppResponse) SetBody(v *GetAppResponseBody) *GetAppResponse {
	s.Body = v
	return s
}

type GetDataStreamResponseBody struct {
	RequestId *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetDataStreamResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetDataStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataStreamResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataStreamResponseBody) SetRequestId(v string) *GetDataStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataStreamResponseBody) SetResult(v *GetDataStreamResponseBodyResult) *GetDataStreamResponseBody {
	s.Result = v
	return s
}

type GetDataStreamResponseBodyResult struct {
	// 关联的应用AppId
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// 代表创建时间的资源属性字段
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 代表资源一级ID的资源属性字段
	DataStreamId *string `json:"dataStreamId,omitempty" xml:"dataStreamId,omitempty"`
	// 代表资源名称的资源属性字段
	DataStreamName *string `json:"dataStreamName,omitempty" xml:"dataStreamName,omitempty"`
	// 删除时间
	DeletePhase *string `json:"deletePhase,omitempty" xml:"deletePhase,omitempty"`
	// 代表region的资源属性字段
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// 数据流模板
	Template     *GetDataStreamResponseBodyResultTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	TimeStampKey *string                                  `json:"timeStampKey,omitempty" xml:"timeStampKey,omitempty"`
	// 数据流类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetDataStreamResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetDataStreamResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDataStreamResponseBodyResult) SetAppName(v string) *GetDataStreamResponseBodyResult {
	s.AppName = &v
	return s
}

func (s *GetDataStreamResponseBodyResult) SetCreateTime(v string) *GetDataStreamResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetDataStreamResponseBodyResult) SetDataStreamId(v string) *GetDataStreamResponseBodyResult {
	s.DataStreamId = &v
	return s
}

func (s *GetDataStreamResponseBodyResult) SetDataStreamName(v string) *GetDataStreamResponseBodyResult {
	s.DataStreamName = &v
	return s
}

func (s *GetDataStreamResponseBodyResult) SetDeletePhase(v string) *GetDataStreamResponseBodyResult {
	s.DeletePhase = &v
	return s
}

func (s *GetDataStreamResponseBodyResult) SetRegionId(v string) *GetDataStreamResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *GetDataStreamResponseBodyResult) SetTemplate(v *GetDataStreamResponseBodyResultTemplate) *GetDataStreamResponseBodyResult {
	s.Template = v
	return s
}

func (s *GetDataStreamResponseBodyResult) SetTimeStampKey(v string) *GetDataStreamResponseBodyResult {
	s.TimeStampKey = &v
	return s
}

func (s *GetDataStreamResponseBodyResult) SetType(v string) *GetDataStreamResponseBodyResult {
	s.Type = &v
	return s
}

type GetDataStreamResponseBodyResultTemplate struct {
	// 索引字段设置
	Mappings []*DataStreamMapping `json:"mappings,omitempty" xml:"mappings,omitempty" type:"Repeated"`
}

func (s GetDataStreamResponseBodyResultTemplate) String() string {
	return tea.Prettify(s)
}

func (s GetDataStreamResponseBodyResultTemplate) GoString() string {
	return s.String()
}

func (s *GetDataStreamResponseBodyResultTemplate) SetMappings(v []*DataStreamMapping) *GetDataStreamResponseBodyResultTemplate {
	s.Mappings = v
	return s
}

type GetDataStreamResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDataStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDataStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataStreamResponse) GoString() string {
	return s.String()
}

func (s *GetDataStreamResponse) SetHeaders(v map[string]*string) *GetDataStreamResponse {
	s.Headers = v
	return s
}

func (s *GetDataStreamResponse) SetStatusCode(v int32) *GetDataStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataStreamResponse) SetBody(v *GetDataStreamResponseBody) *GetDataStreamResponse {
	s.Body = v
	return s
}

type GetRegionInfoResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRegionInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRegionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegionInfoResponseBody) SetRequestId(v string) *GetRegionInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetRegionInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRegionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRegionInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRegionInfoResponse) GoString() string {
	return s.String()
}

func (s *GetRegionInfoResponse) SetHeaders(v map[string]*string) *GetRegionInfoResponse {
	s.Headers = v
	return s
}

func (s *GetRegionInfoResponse) SetStatusCode(v int32) *GetRegionInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegionInfoResponse) SetBody(v *GetRegionInfoResponseBody) *GetRegionInfoResponse {
	s.Body = v
	return s
}

type ListAccessTokensRequest struct {
	TokenId *string `json:"tokenId,omitempty" xml:"tokenId,omitempty"`
}

func (s ListAccessTokensRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccessTokensRequest) GoString() string {
	return s.String()
}

func (s *ListAccessTokensRequest) SetTokenId(v string) *ListAccessTokensRequest {
	s.TokenId = &v
	return s
}

type ListAccessTokensResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAccessTokensResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccessTokensResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessTokensResponseBody) SetRequestId(v string) *ListAccessTokensResponseBody {
	s.RequestId = &v
	return s
}

type ListAccessTokensResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAccessTokensResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAccessTokensResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccessTokensResponse) GoString() string {
	return s.String()
}

func (s *ListAccessTokensResponse) SetHeaders(v map[string]*string) *ListAccessTokensResponse {
	s.Headers = v
	return s
}

func (s *ListAccessTokensResponse) SetStatusCode(v int32) *ListAccessTokensResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessTokensResponse) SetBody(v *ListAccessTokensResponseBody) *ListAccessTokensResponse {
	s.Body = v
	return s
}

type ListAppsRequest struct {
	AppName     *string `json:"appName,omitempty" xml:"appName,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	PageNumber  *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize    *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppsRequest) GoString() string {
	return s.String()
}

func (s *ListAppsRequest) SetAppName(v string) *ListAppsRequest {
	s.AppName = &v
	return s
}

func (s *ListAppsRequest) SetDescription(v string) *ListAppsRequest {
	s.Description = &v
	return s
}

func (s *ListAppsRequest) SetPageNumber(v int32) *ListAppsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppsRequest) SetPageSize(v int32) *ListAppsRequest {
	s.PageSize = &v
	return s
}

type ListAppsResponseBody struct {
	RequestId  *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result     []*ListAppsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int32                        `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBody) SetRequestId(v string) *ListAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppsResponseBody) SetResult(v []*ListAppsResponseBodyResult) *ListAppsResponseBody {
	s.Result = v
	return s
}

func (s *ListAppsResponseBody) SetTotalCount(v int32) *ListAppsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAppsResponseBodyResult struct {
	// 代表资源名称的资源属性字段
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// 代表创建时间的资源属性字段
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 应用备注
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 流量限流额
	FlowQuota *string `json:"flowQuota,omitempty" xml:"flowQuota,omitempty"`
	// OwnerID账号ID
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// 代表region的资源属性字段
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// 代表资源状态的资源属性字段
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 存储限流额
	StorageQuota *string `json:"storageQuota,omitempty" xml:"storageQuota,omitempty"`
}

func (s ListAppsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyResult) SetAppName(v string) *ListAppsResponseBodyResult {
	s.AppName = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetCreateTime(v string) *ListAppsResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetDescription(v string) *ListAppsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetFlowQuota(v string) *ListAppsResponseBodyResult {
	s.FlowQuota = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetOwnerId(v string) *ListAppsResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetRegionId(v string) *ListAppsResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetStatus(v string) *ListAppsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetStorageQuota(v string) *ListAppsResponseBodyResult {
	s.StorageQuota = &v
	return s
}

type ListAppsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponse) GoString() string {
	return s.String()
}

func (s *ListAppsResponse) SetHeaders(v map[string]*string) *ListAppsResponse {
	s.Headers = v
	return s
}

func (s *ListAppsResponse) SetStatusCode(v int32) *ListAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppsResponse) SetBody(v *ListAppsResponseBody) *ListAppsResponse {
	s.Body = v
	return s
}

type ListDataStreamsRequest struct {
	DataStreamName *string `json:"dataStreamName,omitempty" xml:"dataStreamName,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RegionId       *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListDataStreamsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataStreamsRequest) GoString() string {
	return s.String()
}

func (s *ListDataStreamsRequest) SetDataStreamName(v string) *ListDataStreamsRequest {
	s.DataStreamName = &v
	return s
}

func (s *ListDataStreamsRequest) SetPageNumber(v int32) *ListDataStreamsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataStreamsRequest) SetPageSize(v int32) *ListDataStreamsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataStreamsRequest) SetRegionId(v string) *ListDataStreamsRequest {
	s.RegionId = &v
	return s
}

type ListDataStreamsResponseBody struct {
	RequestId  *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result     []*ListDataStreamsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	TotalCount *int32                               `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDataStreamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataStreamsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataStreamsResponseBody) SetRequestId(v string) *ListDataStreamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataStreamsResponseBody) SetResult(v []*ListDataStreamsResponseBodyResult) *ListDataStreamsResponseBody {
	s.Result = v
	return s
}

func (s *ListDataStreamsResponseBody) SetTotalCount(v int32) *ListDataStreamsResponseBody {
	s.TotalCount = &v
	return s
}

type ListDataStreamsResponseBodyResult struct {
	// 关联的应用AppId
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// 代表创建时间的资源属性字段
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 代表资源一级ID的资源属性字段
	DataStreamId *string `json:"dataStreamId,omitempty" xml:"dataStreamId,omitempty"`
	// 代表资源名称的资源属性字段
	DataStreamName *string `json:"dataStreamName,omitempty" xml:"dataStreamName,omitempty"`
	// 删除时间
	DeletePhase *string `json:"deletePhase,omitempty" xml:"deletePhase,omitempty"`
	// 代表region的资源属性字段
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// 数据流模板
	Template     *ListDataStreamsResponseBodyResultTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	TimeStampKey *string                                    `json:"timeStampKey,omitempty" xml:"timeStampKey,omitempty"`
	// 数据流类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDataStreamsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDataStreamsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataStreamsResponseBodyResult) SetAppName(v string) *ListDataStreamsResponseBodyResult {
	s.AppName = &v
	return s
}

func (s *ListDataStreamsResponseBodyResult) SetCreateTime(v string) *ListDataStreamsResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListDataStreamsResponseBodyResult) SetDataStreamId(v string) *ListDataStreamsResponseBodyResult {
	s.DataStreamId = &v
	return s
}

func (s *ListDataStreamsResponseBodyResult) SetDataStreamName(v string) *ListDataStreamsResponseBodyResult {
	s.DataStreamName = &v
	return s
}

func (s *ListDataStreamsResponseBodyResult) SetDeletePhase(v string) *ListDataStreamsResponseBodyResult {
	s.DeletePhase = &v
	return s
}

func (s *ListDataStreamsResponseBodyResult) SetRegionId(v string) *ListDataStreamsResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *ListDataStreamsResponseBodyResult) SetTemplate(v *ListDataStreamsResponseBodyResultTemplate) *ListDataStreamsResponseBodyResult {
	s.Template = v
	return s
}

func (s *ListDataStreamsResponseBodyResult) SetTimeStampKey(v string) *ListDataStreamsResponseBodyResult {
	s.TimeStampKey = &v
	return s
}

func (s *ListDataStreamsResponseBodyResult) SetType(v string) *ListDataStreamsResponseBodyResult {
	s.Type = &v
	return s
}

type ListDataStreamsResponseBodyResultTemplate struct {
	// 索引字段设置
	Mappings []*DataStreamMapping `json:"mappings,omitempty" xml:"mappings,omitempty" type:"Repeated"`
}

func (s ListDataStreamsResponseBodyResultTemplate) String() string {
	return tea.Prettify(s)
}

func (s ListDataStreamsResponseBodyResultTemplate) GoString() string {
	return s.String()
}

func (s *ListDataStreamsResponseBodyResultTemplate) SetMappings(v []*DataStreamMapping) *ListDataStreamsResponseBodyResultTemplate {
	s.Mappings = v
	return s
}

type ListDataStreamsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDataStreamsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataStreamsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataStreamsResponse) GoString() string {
	return s.String()
}

func (s *ListDataStreamsResponse) SetHeaders(v map[string]*string) *ListDataStreamsResponse {
	s.Headers = v
	return s
}

func (s *ListDataStreamsResponse) SetStatusCode(v int32) *ListDataStreamsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataStreamsResponse) SetBody(v *ListDataStreamsResponseBody) *ListDataStreamsResponse {
	s.Body = v
	return s
}

type UpdateAppRequest struct {
	// 应用备注
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppRequest) SetDescription(v string) *UpdateAppRequest {
	s.Description = &v
	return s
}

type UpdateAppResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppResponseBody) SetRequestId(v string) *UpdateAppResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppResponse) SetHeaders(v map[string]*string) *UpdateAppResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppResponse) SetStatusCode(v int32) *UpdateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppResponse) SetBody(v *UpdateAppResponseBody) *UpdateAppResponse {
	s.Body = v
	return s
}

type UpdateDataStreamRequest struct {
	// 删除时间
	DeletePhase  *string                          `json:"deletePhase,omitempty" xml:"deletePhase,omitempty"`
	Dynamic      *string                          `json:"dynamic,omitempty" xml:"dynamic,omitempty"`
	Template     *UpdateDataStreamRequestTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	TimeStampKey *string                          `json:"timeStampKey,omitempty" xml:"timeStampKey,omitempty"`
}

func (s UpdateDataStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataStreamRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataStreamRequest) SetDeletePhase(v string) *UpdateDataStreamRequest {
	s.DeletePhase = &v
	return s
}

func (s *UpdateDataStreamRequest) SetDynamic(v string) *UpdateDataStreamRequest {
	s.Dynamic = &v
	return s
}

func (s *UpdateDataStreamRequest) SetTemplate(v *UpdateDataStreamRequestTemplate) *UpdateDataStreamRequest {
	s.Template = v
	return s
}

func (s *UpdateDataStreamRequest) SetTimeStampKey(v string) *UpdateDataStreamRequest {
	s.TimeStampKey = &v
	return s
}

type UpdateDataStreamRequestTemplate struct {
	Mappings []*DataStreamMapping `json:"mappings,omitempty" xml:"mappings,omitempty" type:"Repeated"`
}

func (s UpdateDataStreamRequestTemplate) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataStreamRequestTemplate) GoString() string {
	return s.String()
}

func (s *UpdateDataStreamRequestTemplate) SetMappings(v []*DataStreamMapping) *UpdateDataStreamRequestTemplate {
	s.Mappings = v
	return s
}

type UpdateDataStreamResponseBody struct {
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdateDataStreamResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateDataStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataStreamResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataStreamResponseBody) SetRequestId(v string) *UpdateDataStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataStreamResponseBody) SetResult(v *UpdateDataStreamResponseBodyResult) *UpdateDataStreamResponseBody {
	s.Result = v
	return s
}

type UpdateDataStreamResponseBodyResult struct {
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
}

func (s UpdateDataStreamResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataStreamResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateDataStreamResponseBodyResult) SetAppName(v string) *UpdateDataStreamResponseBodyResult {
	s.AppName = &v
	return s
}

type UpdateDataStreamResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDataStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDataStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataStreamResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataStreamResponse) SetHeaders(v map[string]*string) *UpdateDataStreamResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataStreamResponse) SetStatusCode(v int32) *UpdateDataStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataStreamResponse) SetBody(v *UpdateDataStreamResponseBody) *UpdateDataStreamResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("es-serverless"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateAppWithOptions(request *CreateAppRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["appName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["chargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApp"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/xops/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDataStreamWithOptions(appName *string, request *CreateDataStreamRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDataStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataStreamName)) {
		body["dataStreamName"] = request.DataStreamName
	}

	if !tea.BoolValue(util.IsUnset(request.DeletePhase)) {
		body["deletePhase"] = request.DeletePhase
	}

	if !tea.BoolValue(util.IsUnset(request.Dynamic)) {
		body["dynamic"] = request.Dynamic
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Template)) {
		body["template"] = request.Template
	}

	if !tea.BoolValue(util.IsUnset(request.TimeStampKey)) {
		body["timeStampKey"] = request.TimeStampKey
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDataStream"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/xops/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/data-streams"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDataStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDataStream(appName *string, request *CreateDataStreamRequest) (_result *CreateDataStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDataStreamResponse{}
	_body, _err := client.CreateDataStreamWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAccessTokenWithOptions(tokenId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAccessTokenResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccessToken"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/xops/tokens/" + tea.StringValue(openapiutil.GetEncodeParam(tokenId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccessTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAccessToken(tokenId *string) (_result *DeleteAccessTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAccessTokenResponse{}
	_body, _err := client.DeleteAccessTokenWithOptions(tokenId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppWithOptions(appName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApp"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/xops/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApp(appName *string) (_result *DeleteAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAppResponse{}
	_body, _err := client.DeleteAppWithOptions(appName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDataStreamWithOptions(appName *string, dataStreamName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDataStreamResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataStream"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/xops/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/data-streams/" + tea.StringValue(openapiutil.GetEncodeParam(dataStreamName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDataStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDataStream(appName *string, dataStreamName *string) (_result *DeleteDataStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDataStreamResponse{}
	_body, _err := client.DeleteDataStreamWithOptions(appName, dataStreamName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescibeRegionsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescibeRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescibeRegions"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/xops/regions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescibeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescibeRegions() (_result *DescibeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescibeRegionsResponse{}
	_body, _err := client.DescibeRegionsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableAccessTokenWithOptions(tokenId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DisableAccessTokenResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DisableAccessToken"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/xops/tokens/" + tea.StringValue(openapiutil.GetEncodeParam(tokenId)) + "/actions/disable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableAccessTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableAccessToken(tokenId *string) (_result *DisableAccessTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableAccessTokenResponse{}
	_body, _err := client.DisableAccessTokenWithOptions(tokenId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableAccessTokenWithOptions(tokenId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EnableAccessTokenResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("EnableAccessToken"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/xops/tokens/" + tea.StringValue(openapiutil.GetEncodeParam(tokenId)) + "/actions/enable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableAccessTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableAccessToken(tokenId *string) (_result *EnableAccessTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableAccessTokenResponse{}
	_body, _err := client.EnableAccessTokenWithOptions(tokenId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateAcccessTokenWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenerateAcccessTokenResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateAcccessToken"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/xops/tokens"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateAcccessTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateAcccessToken() (_result *GenerateAcccessTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateAcccessTokenResponse{}
	_body, _err := client.GenerateAcccessTokenWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppWithOptions(appName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAppResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetApp"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/xops/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApp(appName *string) (_result *GetAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAppResponse{}
	_body, _err := client.GetAppWithOptions(appName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDataStreamWithOptions(appName *string, dataStreamName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDataStreamResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetDataStream"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/xops/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/data-streams/" + tea.StringValue(openapiutil.GetEncodeParam(dataStreamName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDataStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataStream(appName *string, dataStreamName *string) (_result *GetDataStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDataStreamResponse{}
	_body, _err := client.GetDataStreamWithOptions(appName, dataStreamName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRegionInfoWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRegionInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetRegionInfo"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/xops/regions/info"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRegionInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRegionInfo() (_result *GetRegionInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRegionInfoResponse{}
	_body, _err := client.GetRegionInfoWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAccessTokensWithOptions(request *ListAccessTokensRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAccessTokensResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TokenId)) {
		query["tokenId"] = request.TokenId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccessTokens"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/xops/tokens"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccessTokensResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAccessTokens(request *ListAccessTokensRequest) (_result *ListAccessTokensResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAccessTokensResponse{}
	_body, _err := client.ListAccessTokensWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppsWithOptions(request *ListAppsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["appName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApps"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/xops/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApps(request *ListAppsRequest) (_result *ListAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAppsResponse{}
	_body, _err := client.ListAppsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataStreamsWithOptions(appName *string, request *ListDataStreamsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataStreamsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataStreamName)) {
		query["dataStreamName"] = request.DataStreamName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataStreams"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/xops/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/data-streams"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataStreamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataStreams(appName *string, request *ListDataStreamsRequest) (_result *ListDataStreamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataStreamsResponse{}
	_body, _err := client.ListDataStreamsWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppWithOptions(appName *string, request *UpdateAppRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApp"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/xops/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApp(appName *string, request *UpdateAppRequest) (_result *UpdateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAppResponse{}
	_body, _err := client.UpdateAppWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDataStreamWithOptions(dataStreamName *string, appName *string, request *UpdateDataStreamRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateDataStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeletePhase)) {
		body["deletePhase"] = request.DeletePhase
	}

	if !tea.BoolValue(util.IsUnset(request.Dynamic)) {
		body["dynamic"] = request.Dynamic
	}

	if !tea.BoolValue(util.IsUnset(request.Template)) {
		body["template"] = request.Template
	}

	if !tea.BoolValue(util.IsUnset(request.TimeStampKey)) {
		body["timeStampKey"] = request.TimeStampKey
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDataStream"),
		Version:     tea.String("2022-08-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/xops/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/data-streams/" + tea.StringValue(openapiutil.GetEncodeParam(dataStreamName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDataStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDataStream(dataStreamName *string, appName *string, request *UpdateDataStreamRequest) (_result *UpdateDataStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDataStreamResponse{}
	_body, _err := client.UpdateDataStreamWithOptions(dataStreamName, appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
