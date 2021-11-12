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

type BindInstance2VpcRequest struct {
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceVpcName *string `json:"InstanceVpcName,omitempty" xml:"InstanceVpcName,omitempty"`
	Network         *string `json:"Network,omitempty" xml:"Network,omitempty"`
	RegionNo        *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VirtualSwitchId *string `json:"VirtualSwitchId,omitempty" xml:"VirtualSwitchId,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s BindInstance2VpcRequest) String() string {
	return tea.Prettify(s)
}

func (s BindInstance2VpcRequest) GoString() string {
	return s.String()
}

func (s *BindInstance2VpcRequest) SetInstanceName(v string) *BindInstance2VpcRequest {
	s.InstanceName = &v
	return s
}

func (s *BindInstance2VpcRequest) SetInstanceVpcName(v string) *BindInstance2VpcRequest {
	s.InstanceVpcName = &v
	return s
}

func (s *BindInstance2VpcRequest) SetNetwork(v string) *BindInstance2VpcRequest {
	s.Network = &v
	return s
}

func (s *BindInstance2VpcRequest) SetRegionNo(v string) *BindInstance2VpcRequest {
	s.RegionNo = &v
	return s
}

func (s *BindInstance2VpcRequest) SetResourceOwnerId(v int64) *BindInstance2VpcRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindInstance2VpcRequest) SetVirtualSwitchId(v string) *BindInstance2VpcRequest {
	s.VirtualSwitchId = &v
	return s
}

func (s *BindInstance2VpcRequest) SetVpcId(v string) *BindInstance2VpcRequest {
	s.VpcId = &v
	return s
}

type BindInstance2VpcResponseBody struct {
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Endpoint  *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindInstance2VpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindInstance2VpcResponseBody) GoString() string {
	return s.String()
}

func (s *BindInstance2VpcResponseBody) SetDomain(v string) *BindInstance2VpcResponseBody {
	s.Domain = &v
	return s
}

func (s *BindInstance2VpcResponseBody) SetEndpoint(v string) *BindInstance2VpcResponseBody {
	s.Endpoint = &v
	return s
}

func (s *BindInstance2VpcResponseBody) SetRequestId(v string) *BindInstance2VpcResponseBody {
	s.RequestId = &v
	return s
}

type BindInstance2VpcResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindInstance2VpcResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindInstance2VpcResponse) String() string {
	return tea.Prettify(s)
}

func (s BindInstance2VpcResponse) GoString() string {
	return s.String()
}

func (s *BindInstance2VpcResponse) SetHeaders(v map[string]*string) *BindInstance2VpcResponse {
	s.Headers = v
	return s
}

func (s *BindInstance2VpcResponse) SetBody(v *BindInstance2VpcResponseBody) *BindInstance2VpcResponse {
	s.Body = v
	return s
}

type DeleteInstanceRequest struct {
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) SetInstanceName(v string) *DeleteInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *DeleteInstanceRequest) SetResourceOwnerId(v int64) *DeleteInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type DeleteTagsRequest struct {
	InstanceName    *string                     `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ResourceOwnerId *int64                      `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TagInfo         []*DeleteTagsRequestTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
}

func (s DeleteTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagsRequest) GoString() string {
	return s.String()
}

func (s *DeleteTagsRequest) SetInstanceName(v string) *DeleteTagsRequest {
	s.InstanceName = &v
	return s
}

func (s *DeleteTagsRequest) SetResourceOwnerId(v int64) *DeleteTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTagsRequest) SetTagInfo(v []*DeleteTagsRequestTagInfo) *DeleteTagsRequest {
	s.TagInfo = v
	return s
}

type DeleteTagsRequestTagInfo struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DeleteTagsRequestTagInfo) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagsRequestTagInfo) GoString() string {
	return s.String()
}

func (s *DeleteTagsRequestTagInfo) SetTagKey(v string) *DeleteTagsRequestTagInfo {
	s.TagKey = &v
	return s
}

func (s *DeleteTagsRequestTagInfo) SetTagValue(v string) *DeleteTagsRequestTagInfo {
	s.TagValue = &v
	return s
}

type DeleteTagsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTagsResponseBody) SetRequestId(v string) *DeleteTagsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTagsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagsResponse) GoString() string {
	return s.String()
}

func (s *DeleteTagsResponse) SetHeaders(v map[string]*string) *DeleteTagsResponse {
	s.Headers = v
	return s
}

func (s *DeleteTagsResponse) SetBody(v *DeleteTagsResponseBody) *DeleteTagsResponse {
	s.Body = v
	return s
}

type GetInstanceRequest struct {
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceRequest) SetInstanceName(v string) *GetInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceRequest) SetResourceOwnerId(v int64) *GetInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetInstanceResponseBody struct {
	InstanceInfo *GetInstanceResponseBodyInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Struct"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetInstanceInfo(v *GetInstanceResponseBodyInstanceInfo) *GetInstanceResponseBody {
	s.InstanceInfo = v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

type GetInstanceResponseBodyInstanceInfo struct {
	ClusterType   *string                                      `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	CreateTime    *string                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceName  *string                                      `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Network       *string                                      `json:"Network,omitempty" xml:"Network,omitempty"`
	Quota         *GetInstanceResponseBodyInstanceInfoQuota    `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
	ReadCapacity  *int32                                       `json:"ReadCapacity,omitempty" xml:"ReadCapacity,omitempty"`
	Status        *int32                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	TagInfos      *GetInstanceResponseBodyInstanceInfoTagInfos `json:"TagInfos,omitempty" xml:"TagInfos,omitempty" type:"Struct"`
	UserId        *string                                      `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WriteCapacity *int32                                       `json:"WriteCapacity,omitempty" xml:"WriteCapacity,omitempty"`
}

func (s GetInstanceResponseBodyInstanceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceInfo) SetClusterType(v string) *GetInstanceResponseBodyInstanceInfo {
	s.ClusterType = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceInfo) SetCreateTime(v string) *GetInstanceResponseBodyInstanceInfo {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceInfo) SetDescription(v string) *GetInstanceResponseBodyInstanceInfo {
	s.Description = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceInfo) SetInstanceName(v string) *GetInstanceResponseBodyInstanceInfo {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceInfo) SetNetwork(v string) *GetInstanceResponseBodyInstanceInfo {
	s.Network = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceInfo) SetQuota(v *GetInstanceResponseBodyInstanceInfoQuota) *GetInstanceResponseBodyInstanceInfo {
	s.Quota = v
	return s
}

func (s *GetInstanceResponseBodyInstanceInfo) SetReadCapacity(v int32) *GetInstanceResponseBodyInstanceInfo {
	s.ReadCapacity = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceInfo) SetStatus(v int32) *GetInstanceResponseBodyInstanceInfo {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceInfo) SetTagInfos(v *GetInstanceResponseBodyInstanceInfoTagInfos) *GetInstanceResponseBodyInstanceInfo {
	s.TagInfos = v
	return s
}

func (s *GetInstanceResponseBodyInstanceInfo) SetUserId(v string) *GetInstanceResponseBodyInstanceInfo {
	s.UserId = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceInfo) SetWriteCapacity(v int32) *GetInstanceResponseBodyInstanceInfo {
	s.WriteCapacity = &v
	return s
}

type GetInstanceResponseBodyInstanceInfoQuota struct {
	EntityQuota *int32 `json:"EntityQuota,omitempty" xml:"EntityQuota,omitempty"`
}

func (s GetInstanceResponseBodyInstanceInfoQuota) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceInfoQuota) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceInfoQuota) SetEntityQuota(v int32) *GetInstanceResponseBodyInstanceInfoQuota {
	s.EntityQuota = &v
	return s
}

type GetInstanceResponseBodyInstanceInfoTagInfos struct {
	TagInfo []*GetInstanceResponseBodyInstanceInfoTagInfosTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
}

func (s GetInstanceResponseBodyInstanceInfoTagInfos) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceInfoTagInfos) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceInfoTagInfos) SetTagInfo(v []*GetInstanceResponseBodyInstanceInfoTagInfosTagInfo) *GetInstanceResponseBodyInstanceInfoTagInfos {
	s.TagInfo = v
	return s
}

type GetInstanceResponseBodyInstanceInfoTagInfosTagInfo struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetInstanceResponseBodyInstanceInfoTagInfosTagInfo) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceInfoTagInfosTagInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceInfoTagInfosTagInfo) SetTagKey(v string) *GetInstanceResponseBodyInstanceInfoTagInfosTagInfo {
	s.TagKey = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceInfoTagInfosTagInfo) SetTagValue(v string) *GetInstanceResponseBodyInstanceInfoTagInfosTagInfo {
	s.TagValue = &v
	return s
}

type GetInstanceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) SetHeaders(v map[string]*string) *GetInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResponse) SetBody(v *GetInstanceResponseBody) *GetInstanceResponse {
	s.Body = v
	return s
}

type GetOtsServiceStatusRequest struct {
	// ownerId
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s GetOtsServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOtsServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetOtsServiceStatusRequest) SetOwnerId(v int64) *GetOtsServiceStatusRequest {
	s.OwnerId = &v
	return s
}

type GetOtsServiceStatusResponseBody struct {
	// data
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// dynamicCode
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamicMessage
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// errCode
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// httpStatusCode
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOtsServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOtsServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetOtsServiceStatusResponseBody) SetData(v map[string]interface{}) *GetOtsServiceStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetOtsServiceStatusResponseBody) SetDynamicCode(v string) *GetOtsServiceStatusResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetOtsServiceStatusResponseBody) SetDynamicMessage(v string) *GetOtsServiceStatusResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetOtsServiceStatusResponseBody) SetErrCode(v string) *GetOtsServiceStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetOtsServiceStatusResponseBody) SetHttpStatusCode(v int32) *GetOtsServiceStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetOtsServiceStatusResponseBody) SetMessage(v string) *GetOtsServiceStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetOtsServiceStatusResponseBody) SetRequestId(v string) *GetOtsServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOtsServiceStatusResponseBody) SetSuccess(v bool) *GetOtsServiceStatusResponseBody {
	s.Success = &v
	return s
}

type GetOtsServiceStatusResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOtsServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOtsServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOtsServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetOtsServiceStatusResponse) SetHeaders(v map[string]*string) *GetOtsServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetOtsServiceStatusResponse) SetBody(v *GetOtsServiceStatusResponseBody) *GetOtsServiceStatusResponse {
	s.Body = v
	return s
}

type InsertInstanceRequest struct {
	ClusterType     *string                         `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Description     *string                         `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceName    *string                         `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Network         *string                         `json:"Network,omitempty" xml:"Network,omitempty"`
	ResourceOwnerId *int64                          `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TagInfo         []*InsertInstanceRequestTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
}

func (s InsertInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertInstanceRequest) GoString() string {
	return s.String()
}

func (s *InsertInstanceRequest) SetClusterType(v string) *InsertInstanceRequest {
	s.ClusterType = &v
	return s
}

func (s *InsertInstanceRequest) SetDescription(v string) *InsertInstanceRequest {
	s.Description = &v
	return s
}

func (s *InsertInstanceRequest) SetInstanceName(v string) *InsertInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *InsertInstanceRequest) SetNetwork(v string) *InsertInstanceRequest {
	s.Network = &v
	return s
}

func (s *InsertInstanceRequest) SetResourceOwnerId(v int64) *InsertInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *InsertInstanceRequest) SetTagInfo(v []*InsertInstanceRequestTagInfo) *InsertInstanceRequest {
	s.TagInfo = v
	return s
}

type InsertInstanceRequestTagInfo struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s InsertInstanceRequestTagInfo) String() string {
	return tea.Prettify(s)
}

func (s InsertInstanceRequestTagInfo) GoString() string {
	return s.String()
}

func (s *InsertInstanceRequestTagInfo) SetTagKey(v string) *InsertInstanceRequestTagInfo {
	s.TagKey = &v
	return s
}

func (s *InsertInstanceRequestTagInfo) SetTagValue(v string) *InsertInstanceRequestTagInfo {
	s.TagValue = &v
	return s
}

type InsertInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InsertInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *InsertInstanceResponseBody) SetRequestId(v string) *InsertInstanceResponseBody {
	s.RequestId = &v
	return s
}

type InsertInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertInstanceResponse) GoString() string {
	return s.String()
}

func (s *InsertInstanceResponse) SetHeaders(v map[string]*string) *InsertInstanceResponse {
	s.Headers = v
	return s
}

func (s *InsertInstanceResponse) SetBody(v *InsertInstanceResponseBody) *InsertInstanceResponse {
	s.Body = v
	return s
}

type InsertTagsRequest struct {
	InstanceName    *string                     `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ResourceOwnerId *int64                      `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TagInfo         []*InsertTagsRequestTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
}

func (s InsertTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertTagsRequest) GoString() string {
	return s.String()
}

func (s *InsertTagsRequest) SetInstanceName(v string) *InsertTagsRequest {
	s.InstanceName = &v
	return s
}

func (s *InsertTagsRequest) SetResourceOwnerId(v int64) *InsertTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *InsertTagsRequest) SetTagInfo(v []*InsertTagsRequestTagInfo) *InsertTagsRequest {
	s.TagInfo = v
	return s
}

type InsertTagsRequestTagInfo struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s InsertTagsRequestTagInfo) String() string {
	return tea.Prettify(s)
}

func (s InsertTagsRequestTagInfo) GoString() string {
	return s.String()
}

func (s *InsertTagsRequestTagInfo) SetTagKey(v string) *InsertTagsRequestTagInfo {
	s.TagKey = &v
	return s
}

func (s *InsertTagsRequestTagInfo) SetTagValue(v string) *InsertTagsRequestTagInfo {
	s.TagValue = &v
	return s
}

type InsertTagsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InsertTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertTagsResponseBody) GoString() string {
	return s.String()
}

func (s *InsertTagsResponseBody) SetRequestId(v string) *InsertTagsResponseBody {
	s.RequestId = &v
	return s
}

type InsertTagsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertTagsResponse) GoString() string {
	return s.String()
}

func (s *InsertTagsResponse) SetHeaders(v map[string]*string) *InsertTagsResponse {
	s.Headers = v
	return s
}

func (s *InsertTagsResponse) SetBody(v *InsertTagsResponseBody) *InsertTagsResponse {
	s.Body = v
	return s
}

type ListClusterTypeRequest struct {
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListClusterTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTypeRequest) GoString() string {
	return s.String()
}

func (s *ListClusterTypeRequest) SetResourceOwnerId(v int64) *ListClusterTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListClusterTypeResponseBody struct {
	ClusterTypeInfos *ListClusterTypeResponseBodyClusterTypeInfos `json:"ClusterTypeInfos,omitempty" xml:"ClusterTypeInfos,omitempty" type:"Struct"`
	RequestId        *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClusterTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterTypeResponseBody) SetClusterTypeInfos(v *ListClusterTypeResponseBodyClusterTypeInfos) *ListClusterTypeResponseBody {
	s.ClusterTypeInfos = v
	return s
}

func (s *ListClusterTypeResponseBody) SetRequestId(v string) *ListClusterTypeResponseBody {
	s.RequestId = &v
	return s
}

type ListClusterTypeResponseBodyClusterTypeInfos struct {
	ClusterType []*string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty" type:"Repeated"`
}

func (s ListClusterTypeResponseBodyClusterTypeInfos) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTypeResponseBodyClusterTypeInfos) GoString() string {
	return s.String()
}

func (s *ListClusterTypeResponseBodyClusterTypeInfos) SetClusterType(v []*string) *ListClusterTypeResponseBodyClusterTypeInfos {
	s.ClusterType = v
	return s
}

type ListClusterTypeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClusterTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTypeResponse) GoString() string {
	return s.String()
}

func (s *ListClusterTypeResponse) SetHeaders(v map[string]*string) *ListClusterTypeResponse {
	s.Headers = v
	return s
}

func (s *ListClusterTypeResponse) SetBody(v *ListClusterTypeResponseBody) *ListClusterTypeResponse {
	s.Body = v
	return s
}

type ListInstanceRequest struct {
	PageNum         *int64                        `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize        *int64                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerId *int64                        `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TagInfo         []*ListInstanceRequestTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
}

func (s ListInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRequest) SetPageNum(v int64) *ListInstanceRequest {
	s.PageNum = &v
	return s
}

func (s *ListInstanceRequest) SetPageSize(v int64) *ListInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceRequest) SetResourceOwnerId(v int64) *ListInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListInstanceRequest) SetTagInfo(v []*ListInstanceRequestTagInfo) *ListInstanceRequest {
	s.TagInfo = v
	return s
}

type ListInstanceRequestTagInfo struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListInstanceRequestTagInfo) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceRequestTagInfo) GoString() string {
	return s.String()
}

func (s *ListInstanceRequestTagInfo) SetTagKey(v string) *ListInstanceRequestTagInfo {
	s.TagKey = &v
	return s
}

func (s *ListInstanceRequestTagInfo) SetTagValue(v string) *ListInstanceRequestTagInfo {
	s.TagValue = &v
	return s
}

type ListInstanceResponseBody struct {
	InstanceInfos *ListInstanceResponseBodyInstanceInfos `json:"InstanceInfos,omitempty" xml:"InstanceInfos,omitempty" type:"Struct"`
	PageNum       *int64                                 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize      *int64                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *int64                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBody) SetInstanceInfos(v *ListInstanceResponseBodyInstanceInfos) *ListInstanceResponseBody {
	s.InstanceInfos = v
	return s
}

func (s *ListInstanceResponseBody) SetPageNum(v int64) *ListInstanceResponseBody {
	s.PageNum = &v
	return s
}

func (s *ListInstanceResponseBody) SetPageSize(v int64) *ListInstanceResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInstanceResponseBody) SetRequestId(v string) *ListInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceResponseBody) SetTotalCount(v int64) *ListInstanceResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstanceResponseBodyInstanceInfos struct {
	InstanceInfo []*ListInstanceResponseBodyInstanceInfosInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Repeated"`
}

func (s ListInstanceResponseBodyInstanceInfos) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyInstanceInfos) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyInstanceInfos) SetInstanceInfo(v []*ListInstanceResponseBodyInstanceInfosInstanceInfo) *ListInstanceResponseBodyInstanceInfos {
	s.InstanceInfo = v
	return s
}

type ListInstanceResponseBodyInstanceInfosInstanceInfo struct {
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Timestamp    *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListInstanceResponseBodyInstanceInfosInstanceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyInstanceInfosInstanceInfo) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyInstanceInfosInstanceInfo) SetInstanceName(v string) *ListInstanceResponseBodyInstanceInfosInstanceInfo {
	s.InstanceName = &v
	return s
}

func (s *ListInstanceResponseBodyInstanceInfosInstanceInfo) SetTimestamp(v string) *ListInstanceResponseBodyInstanceInfosInstanceInfo {
	s.Timestamp = &v
	return s
}

type ListInstanceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceResponse) SetHeaders(v map[string]*string) *ListInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceResponse) SetBody(v *ListInstanceResponseBody) *ListInstanceResponse {
	s.Body = v
	return s
}

type ListTagsRequest struct {
	InstanceName    *string                   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PageNum         *int64                    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize        *int64                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerId *int64                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TagInfo         []*ListTagsRequestTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
}

func (s ListTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagsRequest) GoString() string {
	return s.String()
}

func (s *ListTagsRequest) SetInstanceName(v string) *ListTagsRequest {
	s.InstanceName = &v
	return s
}

func (s *ListTagsRequest) SetPageNum(v int64) *ListTagsRequest {
	s.PageNum = &v
	return s
}

func (s *ListTagsRequest) SetPageSize(v int64) *ListTagsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagsRequest) SetResourceOwnerId(v int64) *ListTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTagsRequest) SetTagInfo(v []*ListTagsRequestTagInfo) *ListTagsRequest {
	s.TagInfo = v
	return s
}

type ListTagsRequestTagInfo struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagsRequestTagInfo) String() string {
	return tea.Prettify(s)
}

func (s ListTagsRequestTagInfo) GoString() string {
	return s.String()
}

func (s *ListTagsRequestTagInfo) SetTagKey(v string) *ListTagsRequestTagInfo {
	s.TagKey = &v
	return s
}

func (s *ListTagsRequestTagInfo) SetTagValue(v string) *ListTagsRequestTagInfo {
	s.TagValue = &v
	return s
}

type ListTagsResponseBody struct {
	PageNum    *int64                        `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagInfos   *ListTagsResponseBodyTagInfos `json:"TagInfos,omitempty" xml:"TagInfos,omitempty" type:"Struct"`
	TotalCount *int64                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBody) SetPageNum(v int64) *ListTagsResponseBody {
	s.PageNum = &v
	return s
}

func (s *ListTagsResponseBody) SetPageSize(v int64) *ListTagsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagsResponseBody) SetRequestId(v string) *ListTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagsResponseBody) SetTagInfos(v *ListTagsResponseBodyTagInfos) *ListTagsResponseBody {
	s.TagInfos = v
	return s
}

func (s *ListTagsResponseBody) SetTotalCount(v int64) *ListTagsResponseBody {
	s.TotalCount = &v
	return s
}

type ListTagsResponseBodyTagInfos struct {
	TagInfo []*ListTagsResponseBodyTagInfosTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
}

func (s ListTagsResponseBodyTagInfos) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBodyTagInfos) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyTagInfos) SetTagInfo(v []*ListTagsResponseBodyTagInfosTagInfo) *ListTagsResponseBodyTagInfos {
	s.TagInfo = v
	return s
}

type ListTagsResponseBodyTagInfosTagInfo struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagsResponseBodyTagInfosTagInfo) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBodyTagInfosTagInfo) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyTagInfosTagInfo) SetTagKey(v string) *ListTagsResponseBodyTagInfosTagInfo {
	s.TagKey = &v
	return s
}

func (s *ListTagsResponseBodyTagInfosTagInfo) SetTagValue(v string) *ListTagsResponseBodyTagInfosTagInfo {
	s.TagValue = &v
	return s
}

type ListTagsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponse) GoString() string {
	return s.String()
}

func (s *ListTagsResponse) SetHeaders(v map[string]*string) *ListTagsResponse {
	s.Headers = v
	return s
}

func (s *ListTagsResponse) SetBody(v *ListTagsResponseBody) *ListTagsResponse {
	s.Body = v
	return s
}

type ListVpcInfoByInstanceRequest struct {
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PageNum         *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize        *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListVpcInfoByInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVpcInfoByInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByInstanceRequest) SetInstanceName(v string) *ListVpcInfoByInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *ListVpcInfoByInstanceRequest) SetPageNum(v int64) *ListVpcInfoByInstanceRequest {
	s.PageNum = &v
	return s
}

func (s *ListVpcInfoByInstanceRequest) SetPageSize(v int64) *ListVpcInfoByInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListVpcInfoByInstanceRequest) SetResourceOwnerId(v int64) *ListVpcInfoByInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListVpcInfoByInstanceResponseBody struct {
	InstanceName *string                                    `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PageNum      *int64                                     `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int64                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VpcInfos     *ListVpcInfoByInstanceResponseBodyVpcInfos `json:"VpcInfos,omitempty" xml:"VpcInfos,omitempty" type:"Struct"`
}

func (s ListVpcInfoByInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcInfoByInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByInstanceResponseBody) SetInstanceName(v string) *ListVpcInfoByInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *ListVpcInfoByInstanceResponseBody) SetPageNum(v int64) *ListVpcInfoByInstanceResponseBody {
	s.PageNum = &v
	return s
}

func (s *ListVpcInfoByInstanceResponseBody) SetPageSize(v int64) *ListVpcInfoByInstanceResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVpcInfoByInstanceResponseBody) SetRequestId(v string) *ListVpcInfoByInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcInfoByInstanceResponseBody) SetTotalCount(v int64) *ListVpcInfoByInstanceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpcInfoByInstanceResponseBody) SetVpcInfos(v *ListVpcInfoByInstanceResponseBodyVpcInfos) *ListVpcInfoByInstanceResponseBody {
	s.VpcInfos = v
	return s
}

type ListVpcInfoByInstanceResponseBodyVpcInfos struct {
	VpcInfo []*ListVpcInfoByInstanceResponseBodyVpcInfosVpcInfo `json:"VpcInfo,omitempty" xml:"VpcInfo,omitempty" type:"Repeated"`
}

func (s ListVpcInfoByInstanceResponseBodyVpcInfos) String() string {
	return tea.Prettify(s)
}

func (s ListVpcInfoByInstanceResponseBodyVpcInfos) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByInstanceResponseBodyVpcInfos) SetVpcInfo(v []*ListVpcInfoByInstanceResponseBodyVpcInfosVpcInfo) *ListVpcInfoByInstanceResponseBodyVpcInfos {
	s.VpcInfo = v
	return s
}

type ListVpcInfoByInstanceResponseBodyVpcInfosVpcInfo struct {
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Endpoint        *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	InstanceVpcName *string `json:"InstanceVpcName,omitempty" xml:"InstanceVpcName,omitempty"`
	RegionNo        *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListVpcInfoByInstanceResponseBodyVpcInfosVpcInfo) String() string {
	return tea.Prettify(s)
}

func (s ListVpcInfoByInstanceResponseBodyVpcInfosVpcInfo) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByInstanceResponseBodyVpcInfosVpcInfo) SetDomain(v string) *ListVpcInfoByInstanceResponseBodyVpcInfosVpcInfo {
	s.Domain = &v
	return s
}

func (s *ListVpcInfoByInstanceResponseBodyVpcInfosVpcInfo) SetEndpoint(v string) *ListVpcInfoByInstanceResponseBodyVpcInfosVpcInfo {
	s.Endpoint = &v
	return s
}

func (s *ListVpcInfoByInstanceResponseBodyVpcInfosVpcInfo) SetInstanceVpcName(v string) *ListVpcInfoByInstanceResponseBodyVpcInfosVpcInfo {
	s.InstanceVpcName = &v
	return s
}

func (s *ListVpcInfoByInstanceResponseBodyVpcInfosVpcInfo) SetRegionNo(v string) *ListVpcInfoByInstanceResponseBodyVpcInfosVpcInfo {
	s.RegionNo = &v
	return s
}

func (s *ListVpcInfoByInstanceResponseBodyVpcInfosVpcInfo) SetVpcId(v string) *ListVpcInfoByInstanceResponseBodyVpcInfosVpcInfo {
	s.VpcId = &v
	return s
}

type ListVpcInfoByInstanceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVpcInfoByInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVpcInfoByInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVpcInfoByInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByInstanceResponse) SetHeaders(v map[string]*string) *ListVpcInfoByInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListVpcInfoByInstanceResponse) SetBody(v *ListVpcInfoByInstanceResponseBody) *ListVpcInfoByInstanceResponse {
	s.Body = v
	return s
}

type ListVpcInfoByVpcRequest struct {
	PageNum         *int64                            `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize        *int64                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerId *int64                            `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TagInfo         []*ListVpcInfoByVpcRequestTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
	VpcId           *string                           `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListVpcInfoByVpcRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVpcInfoByVpcRequest) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByVpcRequest) SetPageNum(v int64) *ListVpcInfoByVpcRequest {
	s.PageNum = &v
	return s
}

func (s *ListVpcInfoByVpcRequest) SetPageSize(v int64) *ListVpcInfoByVpcRequest {
	s.PageSize = &v
	return s
}

func (s *ListVpcInfoByVpcRequest) SetResourceOwnerId(v int64) *ListVpcInfoByVpcRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListVpcInfoByVpcRequest) SetTagInfo(v []*ListVpcInfoByVpcRequestTagInfo) *ListVpcInfoByVpcRequest {
	s.TagInfo = v
	return s
}

func (s *ListVpcInfoByVpcRequest) SetVpcId(v string) *ListVpcInfoByVpcRequest {
	s.VpcId = &v
	return s
}

type ListVpcInfoByVpcRequestTagInfo struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListVpcInfoByVpcRequestTagInfo) String() string {
	return tea.Prettify(s)
}

func (s ListVpcInfoByVpcRequestTagInfo) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByVpcRequestTagInfo) SetTagKey(v string) *ListVpcInfoByVpcRequestTagInfo {
	s.TagKey = &v
	return s
}

func (s *ListVpcInfoByVpcRequestTagInfo) SetTagValue(v string) *ListVpcInfoByVpcRequestTagInfo {
	s.TagValue = &v
	return s
}

type ListVpcInfoByVpcResponseBody struct {
	PageNum    *int64                                `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VpcId      *string                               `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcInfos   *ListVpcInfoByVpcResponseBodyVpcInfos `json:"VpcInfos,omitempty" xml:"VpcInfos,omitempty" type:"Struct"`
}

func (s ListVpcInfoByVpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcInfoByVpcResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByVpcResponseBody) SetPageNum(v int64) *ListVpcInfoByVpcResponseBody {
	s.PageNum = &v
	return s
}

func (s *ListVpcInfoByVpcResponseBody) SetPageSize(v int64) *ListVpcInfoByVpcResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVpcInfoByVpcResponseBody) SetRequestId(v string) *ListVpcInfoByVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcInfoByVpcResponseBody) SetTotalCount(v int64) *ListVpcInfoByVpcResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpcInfoByVpcResponseBody) SetVpcId(v string) *ListVpcInfoByVpcResponseBody {
	s.VpcId = &v
	return s
}

func (s *ListVpcInfoByVpcResponseBody) SetVpcInfos(v *ListVpcInfoByVpcResponseBodyVpcInfos) *ListVpcInfoByVpcResponseBody {
	s.VpcInfos = v
	return s
}

type ListVpcInfoByVpcResponseBodyVpcInfos struct {
	VpcInfo []*ListVpcInfoByVpcResponseBodyVpcInfosVpcInfo `json:"VpcInfo,omitempty" xml:"VpcInfo,omitempty" type:"Repeated"`
}

func (s ListVpcInfoByVpcResponseBodyVpcInfos) String() string {
	return tea.Prettify(s)
}

func (s ListVpcInfoByVpcResponseBodyVpcInfos) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByVpcResponseBodyVpcInfos) SetVpcInfo(v []*ListVpcInfoByVpcResponseBodyVpcInfosVpcInfo) *ListVpcInfoByVpcResponseBodyVpcInfos {
	s.VpcInfo = v
	return s
}

type ListVpcInfoByVpcResponseBodyVpcInfosVpcInfo struct {
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Endpoint        *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceVpcName *string `json:"InstanceVpcName,omitempty" xml:"InstanceVpcName,omitempty"`
	RegionNo        *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s ListVpcInfoByVpcResponseBodyVpcInfosVpcInfo) String() string {
	return tea.Prettify(s)
}

func (s ListVpcInfoByVpcResponseBodyVpcInfosVpcInfo) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByVpcResponseBodyVpcInfosVpcInfo) SetDomain(v string) *ListVpcInfoByVpcResponseBodyVpcInfosVpcInfo {
	s.Domain = &v
	return s
}

func (s *ListVpcInfoByVpcResponseBodyVpcInfosVpcInfo) SetEndpoint(v string) *ListVpcInfoByVpcResponseBodyVpcInfosVpcInfo {
	s.Endpoint = &v
	return s
}

func (s *ListVpcInfoByVpcResponseBodyVpcInfosVpcInfo) SetInstanceName(v string) *ListVpcInfoByVpcResponseBodyVpcInfosVpcInfo {
	s.InstanceName = &v
	return s
}

func (s *ListVpcInfoByVpcResponseBodyVpcInfosVpcInfo) SetInstanceVpcName(v string) *ListVpcInfoByVpcResponseBodyVpcInfosVpcInfo {
	s.InstanceVpcName = &v
	return s
}

func (s *ListVpcInfoByVpcResponseBodyVpcInfosVpcInfo) SetRegionNo(v string) *ListVpcInfoByVpcResponseBodyVpcInfosVpcInfo {
	s.RegionNo = &v
	return s
}

type ListVpcInfoByVpcResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVpcInfoByVpcResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVpcInfoByVpcResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVpcInfoByVpcResponse) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByVpcResponse) SetHeaders(v map[string]*string) *ListVpcInfoByVpcResponse {
	s.Headers = v
	return s
}

func (s *ListVpcInfoByVpcResponse) SetBody(v *ListVpcInfoByVpcResponseBody) *ListVpcInfoByVpcResponse {
	s.Body = v
	return s
}

type OpenOtsServiceResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenOtsServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenOtsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenOtsServiceResponseBody) SetOrderId(v string) *OpenOtsServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenOtsServiceResponseBody) SetRequestId(v string) *OpenOtsServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenOtsServiceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenOtsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenOtsServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenOtsServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenOtsServiceResponse) SetHeaders(v map[string]*string) *OpenOtsServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenOtsServiceResponse) SetBody(v *OpenOtsServiceResponseBody) *OpenOtsServiceResponse {
	s.Body = v
	return s
}

type UnbindInstance2VpcRequest struct {
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceVpcName *string `json:"InstanceVpcName,omitempty" xml:"InstanceVpcName,omitempty"`
	RegionNo        *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UnbindInstance2VpcRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindInstance2VpcRequest) GoString() string {
	return s.String()
}

func (s *UnbindInstance2VpcRequest) SetInstanceName(v string) *UnbindInstance2VpcRequest {
	s.InstanceName = &v
	return s
}

func (s *UnbindInstance2VpcRequest) SetInstanceVpcName(v string) *UnbindInstance2VpcRequest {
	s.InstanceVpcName = &v
	return s
}

func (s *UnbindInstance2VpcRequest) SetRegionNo(v string) *UnbindInstance2VpcRequest {
	s.RegionNo = &v
	return s
}

func (s *UnbindInstance2VpcRequest) SetResourceOwnerId(v int64) *UnbindInstance2VpcRequest {
	s.ResourceOwnerId = &v
	return s
}

type UnbindInstance2VpcResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindInstance2VpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindInstance2VpcResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindInstance2VpcResponseBody) SetRequestId(v string) *UnbindInstance2VpcResponseBody {
	s.RequestId = &v
	return s
}

type UnbindInstance2VpcResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindInstance2VpcResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindInstance2VpcResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindInstance2VpcResponse) GoString() string {
	return s.String()
}

func (s *UnbindInstance2VpcResponse) SetHeaders(v map[string]*string) *UnbindInstance2VpcResponse {
	s.Headers = v
	return s
}

func (s *UnbindInstance2VpcResponse) SetBody(v *UnbindInstance2VpcResponseBody) *UnbindInstance2VpcResponse {
	s.Body = v
	return s
}

type UpdateInstanceRequest struct {
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Network         *string `json:"Network,omitempty" xml:"Network,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) SetInstanceName(v string) *UpdateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceRequest) SetNetwork(v string) *UpdateInstanceRequest {
	s.Network = &v
	return s
}

func (s *UpdateInstanceRequest) SetResourceOwnerId(v int64) *UpdateInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

type UpdateInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponse) SetHeaders(v map[string]*string) *UpdateInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceResponse) SetBody(v *UpdateInstanceResponseBody) *UpdateInstanceResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("ots"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) BindInstance2VpcWithOptions(request *BindInstance2VpcRequest, runtime *util.RuntimeOptions) (_result *BindInstance2VpcResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BindInstance2VpcResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BindInstance2Vpc"), tea.String("2016-06-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindInstance2Vpc(request *BindInstance2VpcRequest) (_result *BindInstance2VpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindInstance2VpcResponse{}
	_body, _err := client.BindInstance2VpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteInstance"), tea.String("2016-06-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTagsWithOptions(request *DeleteTagsRequest, runtime *util.RuntimeOptions) (_result *DeleteTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteTags"), tea.String("2016-06-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTags(request *DeleteTagsRequest) (_result *DeleteTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTagsResponse{}
	_body, _err := client.DeleteTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceWithOptions(request *GetInstanceRequest, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetInstance"), tea.String("2016-06-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstance(request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOtsServiceStatusWithOptions(request *GetOtsServiceStatusRequest, runtime *util.RuntimeOptions) (_result *GetOtsServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetOtsServiceStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOtsServiceStatus"), tea.String("2016-06-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOtsServiceStatus(request *GetOtsServiceStatusRequest) (_result *GetOtsServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOtsServiceStatusResponse{}
	_body, _err := client.GetOtsServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertInstanceWithOptions(request *InsertInstanceRequest, runtime *util.RuntimeOptions) (_result *InsertInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InsertInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InsertInstance"), tea.String("2016-06-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertInstance(request *InsertInstanceRequest) (_result *InsertInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertInstanceResponse{}
	_body, _err := client.InsertInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertTagsWithOptions(request *InsertTagsRequest, runtime *util.RuntimeOptions) (_result *InsertTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InsertTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InsertTags"), tea.String("2016-06-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertTags(request *InsertTagsRequest) (_result *InsertTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertTagsResponse{}
	_body, _err := client.InsertTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterTypeWithOptions(request *ListClusterTypeRequest, runtime *util.RuntimeOptions) (_result *ListClusterTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListClusterTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListClusterType"), tea.String("2016-06-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterType(request *ListClusterTypeRequest) (_result *ListClusterTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterTypeResponse{}
	_body, _err := client.ListClusterTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstanceWithOptions(request *ListInstanceRequest, runtime *util.RuntimeOptions) (_result *ListInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInstance"), tea.String("2016-06-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstance(request *ListInstanceRequest) (_result *ListInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstanceResponse{}
	_body, _err := client.ListInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagsWithOptions(request *ListTagsRequest, runtime *util.RuntimeOptions) (_result *ListTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTags"), tea.String("2016-06-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTags(request *ListTagsRequest) (_result *ListTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagsResponse{}
	_body, _err := client.ListTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVpcInfoByInstanceWithOptions(request *ListVpcInfoByInstanceRequest, runtime *util.RuntimeOptions) (_result *ListVpcInfoByInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListVpcInfoByInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVpcInfoByInstance"), tea.String("2016-06-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVpcInfoByInstance(request *ListVpcInfoByInstanceRequest) (_result *ListVpcInfoByInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVpcInfoByInstanceResponse{}
	_body, _err := client.ListVpcInfoByInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVpcInfoByVpcWithOptions(request *ListVpcInfoByVpcRequest, runtime *util.RuntimeOptions) (_result *ListVpcInfoByVpcResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListVpcInfoByVpcResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVpcInfoByVpc"), tea.String("2016-06-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVpcInfoByVpc(request *ListVpcInfoByVpcRequest) (_result *ListVpcInfoByVpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVpcInfoByVpcResponse{}
	_body, _err := client.ListVpcInfoByVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenOtsServiceWithOptions(runtime *util.RuntimeOptions) (_result *OpenOtsServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &OpenOtsServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpenOtsService"), tea.String("2016-06-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenOtsService() (_result *OpenOtsServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenOtsServiceResponse{}
	_body, _err := client.OpenOtsServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindInstance2VpcWithOptions(request *UnbindInstance2VpcRequest, runtime *util.RuntimeOptions) (_result *UnbindInstance2VpcResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnbindInstance2VpcResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnbindInstance2Vpc"), tea.String("2016-06-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindInstance2Vpc(request *UnbindInstance2VpcRequest) (_result *UnbindInstance2VpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindInstance2VpcResponse{}
	_body, _err := client.UnbindInstance2VpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstanceWithOptions(request *UpdateInstanceRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateInstance"), tea.String("2016-06-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstance(request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
