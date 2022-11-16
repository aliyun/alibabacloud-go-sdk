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

type AddAccessControlListEntryRequest struct {
	AclEntrys            *string `json:"AclEntrys,omitempty" xml:"AclEntrys,omitempty"`
	AclId                *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddAccessControlListEntryRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAccessControlListEntryRequest) GoString() string {
	return s.String()
}

func (s *AddAccessControlListEntryRequest) SetAclEntrys(v string) *AddAccessControlListEntryRequest {
	s.AclEntrys = &v
	return s
}

func (s *AddAccessControlListEntryRequest) SetAclId(v string) *AddAccessControlListEntryRequest {
	s.AclId = &v
	return s
}

func (s *AddAccessControlListEntryRequest) SetOwnerAccount(v string) *AddAccessControlListEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddAccessControlListEntryRequest) SetOwnerId(v int64) *AddAccessControlListEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *AddAccessControlListEntryRequest) SetRegionId(v string) *AddAccessControlListEntryRequest {
	s.RegionId = &v
	return s
}

func (s *AddAccessControlListEntryRequest) SetResourceOwnerAccount(v string) *AddAccessControlListEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddAccessControlListEntryRequest) SetResourceOwnerId(v int64) *AddAccessControlListEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

type AddAccessControlListEntryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAccessControlListEntryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAccessControlListEntryResponseBody) GoString() string {
	return s.String()
}

func (s *AddAccessControlListEntryResponseBody) SetRequestId(v string) *AddAccessControlListEntryResponseBody {
	s.RequestId = &v
	return s
}

type AddAccessControlListEntryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddAccessControlListEntryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddAccessControlListEntryResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAccessControlListEntryResponse) GoString() string {
	return s.String()
}

func (s *AddAccessControlListEntryResponse) SetHeaders(v map[string]*string) *AddAccessControlListEntryResponse {
	s.Headers = v
	return s
}

func (s *AddAccessControlListEntryResponse) SetStatusCode(v int32) *AddAccessControlListEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAccessControlListEntryResponse) SetBody(v *AddAccessControlListEntryResponseBody) *AddAccessControlListEntryResponse {
	s.Body = v
	return s
}

type AddBackendServersRequest struct {
	BackendServers       *string `json:"BackendServers,omitempty" xml:"BackendServers,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddBackendServersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddBackendServersRequest) GoString() string {
	return s.String()
}

func (s *AddBackendServersRequest) SetBackendServers(v string) *AddBackendServersRequest {
	s.BackendServers = &v
	return s
}

func (s *AddBackendServersRequest) SetLoadBalancerId(v string) *AddBackendServersRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *AddBackendServersRequest) SetOwnerAccount(v string) *AddBackendServersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddBackendServersRequest) SetOwnerId(v int64) *AddBackendServersRequest {
	s.OwnerId = &v
	return s
}

func (s *AddBackendServersRequest) SetRegionId(v string) *AddBackendServersRequest {
	s.RegionId = &v
	return s
}

func (s *AddBackendServersRequest) SetResourceOwnerAccount(v string) *AddBackendServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddBackendServersRequest) SetResourceOwnerId(v int64) *AddBackendServersRequest {
	s.ResourceOwnerId = &v
	return s
}

type AddBackendServersResponseBody struct {
	BackendServers *AddBackendServersResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	LoadBalancerId *string                                      `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddBackendServersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddBackendServersResponseBody) GoString() string {
	return s.String()
}

func (s *AddBackendServersResponseBody) SetBackendServers(v *AddBackendServersResponseBodyBackendServers) *AddBackendServersResponseBody {
	s.BackendServers = v
	return s
}

func (s *AddBackendServersResponseBody) SetLoadBalancerId(v string) *AddBackendServersResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *AddBackendServersResponseBody) SetRequestId(v string) *AddBackendServersResponseBody {
	s.RequestId = &v
	return s
}

type AddBackendServersResponseBodyBackendServers struct {
	BackendServer []*AddBackendServersResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s AddBackendServersResponseBodyBackendServers) String() string {
	return tea.Prettify(s)
}

func (s AddBackendServersResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *AddBackendServersResponseBodyBackendServers) SetBackendServer(v []*AddBackendServersResponseBodyBackendServersBackendServer) *AddBackendServersResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

type AddBackendServersResponseBodyBackendServersBackendServer struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ServerId    *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight      *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s AddBackendServersResponseBodyBackendServersBackendServer) String() string {
	return tea.Prettify(s)
}

func (s AddBackendServersResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) SetDescription(v string) *AddBackendServersResponseBodyBackendServersBackendServer {
	s.Description = &v
	return s
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) SetServerId(v string) *AddBackendServersResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) SetType(v string) *AddBackendServersResponseBodyBackendServersBackendServer {
	s.Type = &v
	return s
}

func (s *AddBackendServersResponseBodyBackendServersBackendServer) SetWeight(v string) *AddBackendServersResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

type AddBackendServersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddBackendServersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddBackendServersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddBackendServersResponse) GoString() string {
	return s.String()
}

func (s *AddBackendServersResponse) SetHeaders(v map[string]*string) *AddBackendServersResponse {
	s.Headers = v
	return s
}

func (s *AddBackendServersResponse) SetStatusCode(v int32) *AddBackendServersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddBackendServersResponse) SetBody(v *AddBackendServersResponseBody) *AddBackendServersResponse {
	s.Body = v
	return s
}

type AddListenerWhiteListItemRequest struct {
	ListenerPort         *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerProtocol     *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SourceItems          *string `json:"SourceItems,omitempty" xml:"SourceItems,omitempty"`
}

func (s AddListenerWhiteListItemRequest) String() string {
	return tea.Prettify(s)
}

func (s AddListenerWhiteListItemRequest) GoString() string {
	return s.String()
}

func (s *AddListenerWhiteListItemRequest) SetListenerPort(v int32) *AddListenerWhiteListItemRequest {
	s.ListenerPort = &v
	return s
}

func (s *AddListenerWhiteListItemRequest) SetListenerProtocol(v string) *AddListenerWhiteListItemRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *AddListenerWhiteListItemRequest) SetLoadBalancerId(v string) *AddListenerWhiteListItemRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *AddListenerWhiteListItemRequest) SetOwnerAccount(v string) *AddListenerWhiteListItemRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddListenerWhiteListItemRequest) SetOwnerId(v int64) *AddListenerWhiteListItemRequest {
	s.OwnerId = &v
	return s
}

func (s *AddListenerWhiteListItemRequest) SetRegionId(v string) *AddListenerWhiteListItemRequest {
	s.RegionId = &v
	return s
}

func (s *AddListenerWhiteListItemRequest) SetResourceOwnerAccount(v string) *AddListenerWhiteListItemRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddListenerWhiteListItemRequest) SetResourceOwnerId(v int64) *AddListenerWhiteListItemRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddListenerWhiteListItemRequest) SetSourceItems(v string) *AddListenerWhiteListItemRequest {
	s.SourceItems = &v
	return s
}

type AddListenerWhiteListItemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddListenerWhiteListItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddListenerWhiteListItemResponseBody) GoString() string {
	return s.String()
}

func (s *AddListenerWhiteListItemResponseBody) SetRequestId(v string) *AddListenerWhiteListItemResponseBody {
	s.RequestId = &v
	return s
}

type AddListenerWhiteListItemResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddListenerWhiteListItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddListenerWhiteListItemResponse) String() string {
	return tea.Prettify(s)
}

func (s AddListenerWhiteListItemResponse) GoString() string {
	return s.String()
}

func (s *AddListenerWhiteListItemResponse) SetHeaders(v map[string]*string) *AddListenerWhiteListItemResponse {
	s.Headers = v
	return s
}

func (s *AddListenerWhiteListItemResponse) SetStatusCode(v int32) *AddListenerWhiteListItemResponse {
	s.StatusCode = &v
	return s
}

func (s *AddListenerWhiteListItemResponse) SetBody(v *AddListenerWhiteListItemResponseBody) *AddListenerWhiteListItemResponse {
	s.Body = v
	return s
}

type AddTagsRequest struct {
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s AddTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTagsRequest) GoString() string {
	return s.String()
}

func (s *AddTagsRequest) SetLoadBalancerId(v string) *AddTagsRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *AddTagsRequest) SetOwnerAccount(v string) *AddTagsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddTagsRequest) SetOwnerId(v int64) *AddTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *AddTagsRequest) SetRegionId(v string) *AddTagsRequest {
	s.RegionId = &v
	return s
}

func (s *AddTagsRequest) SetResourceOwnerAccount(v string) *AddTagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddTagsRequest) SetResourceOwnerId(v int64) *AddTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddTagsRequest) SetTags(v string) *AddTagsRequest {
	s.Tags = &v
	return s
}

type AddTagsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddTagsResponseBody) GoString() string {
	return s.String()
}

func (s *AddTagsResponseBody) SetRequestId(v string) *AddTagsResponseBody {
	s.RequestId = &v
	return s
}

type AddTagsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTagsResponse) GoString() string {
	return s.String()
}

func (s *AddTagsResponse) SetHeaders(v map[string]*string) *AddTagsResponse {
	s.Headers = v
	return s
}

func (s *AddTagsResponse) SetStatusCode(v int32) *AddTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTagsResponse) SetBody(v *AddTagsResponseBody) *AddTagsResponse {
	s.Body = v
	return s
}

type AddVServerGroupBackendServersRequest struct {
	BackendServers       *string `json:"BackendServers,omitempty" xml:"BackendServers,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VServerGroupId       *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s AddVServerGroupBackendServersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddVServerGroupBackendServersRequest) GoString() string {
	return s.String()
}

func (s *AddVServerGroupBackendServersRequest) SetBackendServers(v string) *AddVServerGroupBackendServersRequest {
	s.BackendServers = &v
	return s
}

func (s *AddVServerGroupBackendServersRequest) SetOwnerAccount(v string) *AddVServerGroupBackendServersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddVServerGroupBackendServersRequest) SetOwnerId(v int64) *AddVServerGroupBackendServersRequest {
	s.OwnerId = &v
	return s
}

func (s *AddVServerGroupBackendServersRequest) SetRegionId(v string) *AddVServerGroupBackendServersRequest {
	s.RegionId = &v
	return s
}

func (s *AddVServerGroupBackendServersRequest) SetResourceOwnerAccount(v string) *AddVServerGroupBackendServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddVServerGroupBackendServersRequest) SetResourceOwnerId(v int64) *AddVServerGroupBackendServersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddVServerGroupBackendServersRequest) SetVServerGroupId(v string) *AddVServerGroupBackendServersRequest {
	s.VServerGroupId = &v
	return s
}

type AddVServerGroupBackendServersResponseBody struct {
	BackendServers *AddVServerGroupBackendServersResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	RequestId      *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VServerGroupId *string                                                  `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s AddVServerGroupBackendServersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddVServerGroupBackendServersResponseBody) GoString() string {
	return s.String()
}

func (s *AddVServerGroupBackendServersResponseBody) SetBackendServers(v *AddVServerGroupBackendServersResponseBodyBackendServers) *AddVServerGroupBackendServersResponseBody {
	s.BackendServers = v
	return s
}

func (s *AddVServerGroupBackendServersResponseBody) SetRequestId(v string) *AddVServerGroupBackendServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddVServerGroupBackendServersResponseBody) SetVServerGroupId(v string) *AddVServerGroupBackendServersResponseBody {
	s.VServerGroupId = &v
	return s
}

type AddVServerGroupBackendServersResponseBodyBackendServers struct {
	BackendServer []*AddVServerGroupBackendServersResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s AddVServerGroupBackendServersResponseBodyBackendServers) String() string {
	return tea.Prettify(s)
}

func (s AddVServerGroupBackendServersResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *AddVServerGroupBackendServersResponseBodyBackendServers) SetBackendServer(v []*AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) *AddVServerGroupBackendServersResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

type AddVServerGroupBackendServersResponseBodyBackendServersBackendServer struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Port        *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ServerId    *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight      *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) String() string {
	return tea.Prettify(s)
}

func (s AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetDescription(v string) *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Description = &v
	return s
}

func (s *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetPort(v int32) *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Port = &v
	return s
}

func (s *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetServerId(v string) *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetType(v string) *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Type = &v
	return s
}

func (s *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetWeight(v int32) *AddVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

type AddVServerGroupBackendServersResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddVServerGroupBackendServersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddVServerGroupBackendServersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddVServerGroupBackendServersResponse) GoString() string {
	return s.String()
}

func (s *AddVServerGroupBackendServersResponse) SetHeaders(v map[string]*string) *AddVServerGroupBackendServersResponse {
	s.Headers = v
	return s
}

func (s *AddVServerGroupBackendServersResponse) SetStatusCode(v int32) *AddVServerGroupBackendServersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddVServerGroupBackendServersResponse) SetBody(v *AddVServerGroupBackendServersResponseBody) *AddVServerGroupBackendServersResponse {
	s.Body = v
	return s
}

type CreateAccessControlListRequest struct {
	AclName              *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	AddressIPVersion     *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateAccessControlListRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessControlListRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessControlListRequest) SetAclName(v string) *CreateAccessControlListRequest {
	s.AclName = &v
	return s
}

func (s *CreateAccessControlListRequest) SetAddressIPVersion(v string) *CreateAccessControlListRequest {
	s.AddressIPVersion = &v
	return s
}

func (s *CreateAccessControlListRequest) SetOwnerAccount(v string) *CreateAccessControlListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAccessControlListRequest) SetOwnerId(v int64) *CreateAccessControlListRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAccessControlListRequest) SetRegionId(v string) *CreateAccessControlListRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAccessControlListRequest) SetResourceGroupId(v string) *CreateAccessControlListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAccessControlListRequest) SetResourceOwnerAccount(v string) *CreateAccessControlListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAccessControlListRequest) SetResourceOwnerId(v int64) *CreateAccessControlListRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateAccessControlListResponseBody struct {
	AclId     *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccessControlListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessControlListResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessControlListResponseBody) SetAclId(v string) *CreateAccessControlListResponseBody {
	s.AclId = &v
	return s
}

func (s *CreateAccessControlListResponseBody) SetRequestId(v string) *CreateAccessControlListResponseBody {
	s.RequestId = &v
	return s
}

type CreateAccessControlListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAccessControlListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAccessControlListResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessControlListResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessControlListResponse) SetHeaders(v map[string]*string) *CreateAccessControlListResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessControlListResponse) SetStatusCode(v int32) *CreateAccessControlListResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessControlListResponse) SetBody(v *CreateAccessControlListResponseBody) *CreateAccessControlListResponse {
	s.Body = v
	return s
}

type CreateDomainExtensionRequest struct {
	Domain               *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ListenerPort         *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ServerCertificateId  *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
}

func (s CreateDomainExtensionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainExtensionRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainExtensionRequest) SetDomain(v string) *CreateDomainExtensionRequest {
	s.Domain = &v
	return s
}

func (s *CreateDomainExtensionRequest) SetListenerPort(v int32) *CreateDomainExtensionRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateDomainExtensionRequest) SetLoadBalancerId(v string) *CreateDomainExtensionRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateDomainExtensionRequest) SetOwnerAccount(v string) *CreateDomainExtensionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDomainExtensionRequest) SetOwnerId(v int64) *CreateDomainExtensionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDomainExtensionRequest) SetRegionId(v string) *CreateDomainExtensionRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDomainExtensionRequest) SetResourceOwnerAccount(v string) *CreateDomainExtensionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDomainExtensionRequest) SetResourceOwnerId(v int64) *CreateDomainExtensionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDomainExtensionRequest) SetServerCertificateId(v string) *CreateDomainExtensionRequest {
	s.ServerCertificateId = &v
	return s
}

type CreateDomainExtensionResponseBody struct {
	DomainExtensionId *string `json:"DomainExtensionId,omitempty" xml:"DomainExtensionId,omitempty"`
	ListenerPort      *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDomainExtensionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainExtensionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDomainExtensionResponseBody) SetDomainExtensionId(v string) *CreateDomainExtensionResponseBody {
	s.DomainExtensionId = &v
	return s
}

func (s *CreateDomainExtensionResponseBody) SetListenerPort(v int32) *CreateDomainExtensionResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *CreateDomainExtensionResponseBody) SetRequestId(v string) *CreateDomainExtensionResponseBody {
	s.RequestId = &v
	return s
}

type CreateDomainExtensionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDomainExtensionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDomainExtensionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainExtensionResponse) GoString() string {
	return s.String()
}

func (s *CreateDomainExtensionResponse) SetHeaders(v map[string]*string) *CreateDomainExtensionResponse {
	s.Headers = v
	return s
}

func (s *CreateDomainExtensionResponse) SetStatusCode(v int32) *CreateDomainExtensionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDomainExtensionResponse) SetBody(v *CreateDomainExtensionResponseBody) *CreateDomainExtensionResponse {
	s.Body = v
	return s
}

type CreateLoadBalancerRequest struct {
	Address                      *string `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressIPVersion             *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	AddressType                  *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	AutoPay                      *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	Bandwidth                    *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ClientToken                  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DeleteProtection             *string `json:"DeleteProtection,omitempty" xml:"DeleteProtection,omitempty"`
	Duration                     *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceChargeType           *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InternetChargeType           *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	LoadBalancerName             *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	LoadBalancerSpec             *string `json:"LoadBalancerSpec,omitempty" xml:"LoadBalancerSpec,omitempty"`
	MasterZoneId                 *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	ModificationProtectionReason *string `json:"ModificationProtectionReason,omitempty" xml:"ModificationProtectionReason,omitempty"`
	ModificationProtectionStatus *string `json:"ModificationProtectionStatus,omitempty" xml:"ModificationProtectionStatus,omitempty"`
	OwnerAccount                 *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PayType                      *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PricingCycle                 *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	RegionId                     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId              *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount         *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId              *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SlaveZoneId                  *string `json:"SlaveZoneId,omitempty" xml:"SlaveZoneId,omitempty"`
	VSwitchId                    *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                        *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateLoadBalancerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequest) SetAddress(v string) *CreateLoadBalancerRequest {
	s.Address = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetAddressIPVersion(v string) *CreateLoadBalancerRequest {
	s.AddressIPVersion = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetAddressType(v string) *CreateLoadBalancerRequest {
	s.AddressType = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetAutoPay(v bool) *CreateLoadBalancerRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetBandwidth(v int32) *CreateLoadBalancerRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetClientToken(v string) *CreateLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetDeleteProtection(v string) *CreateLoadBalancerRequest {
	s.DeleteProtection = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetDuration(v int32) *CreateLoadBalancerRequest {
	s.Duration = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetInstanceChargeType(v string) *CreateLoadBalancerRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetInternetChargeType(v string) *CreateLoadBalancerRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerName(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerSpec(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerSpec = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetMasterZoneId(v string) *CreateLoadBalancerRequest {
	s.MasterZoneId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetModificationProtectionReason(v string) *CreateLoadBalancerRequest {
	s.ModificationProtectionReason = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetModificationProtectionStatus(v string) *CreateLoadBalancerRequest {
	s.ModificationProtectionStatus = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetOwnerAccount(v string) *CreateLoadBalancerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetOwnerId(v int64) *CreateLoadBalancerRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetPayType(v string) *CreateLoadBalancerRequest {
	s.PayType = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetPricingCycle(v string) *CreateLoadBalancerRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetRegionId(v string) *CreateLoadBalancerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetResourceGroupId(v string) *CreateLoadBalancerRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetResourceOwnerAccount(v string) *CreateLoadBalancerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetResourceOwnerId(v int64) *CreateLoadBalancerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetSlaveZoneId(v string) *CreateLoadBalancerRequest {
	s.SlaveZoneId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetVSwitchId(v string) *CreateLoadBalancerRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetVpcId(v string) *CreateLoadBalancerRequest {
	s.VpcId = &v
	return s
}

type CreateLoadBalancerResponseBody struct {
	Address          *string `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	LoadBalancerId   *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	NetworkType      *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OrderId          *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId  *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId            *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateLoadBalancerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerResponseBody) SetAddress(v string) *CreateLoadBalancerResponseBody {
	s.Address = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetAddressIPVersion(v string) *CreateLoadBalancerResponseBody {
	s.AddressIPVersion = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetLoadBalancerId(v string) *CreateLoadBalancerResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetLoadBalancerName(v string) *CreateLoadBalancerResponseBody {
	s.LoadBalancerName = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetNetworkType(v string) *CreateLoadBalancerResponseBody {
	s.NetworkType = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetOrderId(v int64) *CreateLoadBalancerResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetRequestId(v string) *CreateLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetResourceGroupId(v string) *CreateLoadBalancerResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetVSwitchId(v string) *CreateLoadBalancerResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetVpcId(v string) *CreateLoadBalancerResponseBody {
	s.VpcId = &v
	return s
}

type CreateLoadBalancerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLoadBalancerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLoadBalancerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerResponse) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerResponse) SetHeaders(v map[string]*string) *CreateLoadBalancerResponse {
	s.Headers = v
	return s
}

func (s *CreateLoadBalancerResponse) SetStatusCode(v int32) *CreateLoadBalancerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLoadBalancerResponse) SetBody(v *CreateLoadBalancerResponseBody) *CreateLoadBalancerResponse {
	s.Body = v
	return s
}

type CreateLoadBalancerHTTPListenerRequest struct {
	AclId                  *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclStatus              *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	AclType                *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	BackendServerPort      *int32  `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	Bandwidth              *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Cookie                 *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	CookieTimeout          *int32  `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ForwardPort            *int32  `json:"ForwardPort,omitempty" xml:"ForwardPort,omitempty"`
	Gzip                   *string `json:"Gzip,omitempty" xml:"Gzip,omitempty"`
	HealthCheck            *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckConnectPort *int32  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckDomain      *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	HealthCheckHttpCode    *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckInterval    *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckMethod      *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	HealthCheckTimeout     *int32  `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	HealthCheckURI         *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	HealthyThreshold       *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	IdleTimeout            *int32  `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	ListenerForward        *string `json:"ListenerForward,omitempty" xml:"ListenerForward,omitempty"`
	ListenerPort           *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	LoadBalancerId         *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestTimeout         *int32  `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Scheduler              *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	StickySession          *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	StickySessionType      *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	UnhealthyThreshold     *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	VServerGroupId         *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	XForwardedFor          *string `json:"XForwardedFor,omitempty" xml:"XForwardedFor,omitempty"`
	XForwardedFor_SLBID    *string `json:"XForwardedFor_SLBID,omitempty" xml:"XForwardedFor_SLBID,omitempty"`
	XForwardedFor_SLBIP    *string `json:"XForwardedFor_SLBIP,omitempty" xml:"XForwardedFor_SLBIP,omitempty"`
	XForwardedFor_proto    *string `json:"XForwardedFor_proto,omitempty" xml:"XForwardedFor_proto,omitempty"`
}

func (s CreateLoadBalancerHTTPListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerHTTPListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetAclId(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.AclId = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetAclStatus(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.AclStatus = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetAclType(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.AclType = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetBackendServerPort(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.BackendServerPort = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetBandwidth(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetCookie(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.Cookie = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetCookieTimeout(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.CookieTimeout = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetDescription(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.Description = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetForwardPort(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.ForwardPort = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetGzip(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.Gzip = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthCheck(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthCheck = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckConnectPort(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckDomain(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthCheckDomain = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckHttpCode(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckInterval(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckMethod(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthCheckMethod = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckTimeout(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthCheckTimeout = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckURI(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthCheckURI = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetHealthyThreshold(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetIdleTimeout(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.IdleTimeout = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetListenerForward(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.ListenerForward = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetListenerPort(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetLoadBalancerId(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetOwnerAccount(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetOwnerId(v int64) *CreateLoadBalancerHTTPListenerRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetRegionId(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetRequestTimeout(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.RequestTimeout = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetResourceOwnerAccount(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetResourceOwnerId(v int64) *CreateLoadBalancerHTTPListenerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetScheduler(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.Scheduler = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetStickySession(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.StickySession = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetStickySessionType(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.StickySessionType = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetUnhealthyThreshold(v int32) *CreateLoadBalancerHTTPListenerRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetVServerGroupId(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.VServerGroupId = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetXForwardedFor(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.XForwardedFor = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetXForwardedFor_SLBID(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.XForwardedFor_SLBID = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetXForwardedFor_SLBIP(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.XForwardedFor_SLBIP = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerRequest) SetXForwardedFor_proto(v string) *CreateLoadBalancerHTTPListenerRequest {
	s.XForwardedFor_proto = &v
	return s
}

type CreateLoadBalancerHTTPListenerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLoadBalancerHTTPListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerHTTPListenerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerHTTPListenerResponseBody) SetRequestId(v string) *CreateLoadBalancerHTTPListenerResponseBody {
	s.RequestId = &v
	return s
}

type CreateLoadBalancerHTTPListenerResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLoadBalancerHTTPListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLoadBalancerHTTPListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerHTTPListenerResponse) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerHTTPListenerResponse) SetHeaders(v map[string]*string) *CreateLoadBalancerHTTPListenerResponse {
	s.Headers = v
	return s
}

func (s *CreateLoadBalancerHTTPListenerResponse) SetStatusCode(v int32) *CreateLoadBalancerHTTPListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerResponse) SetBody(v *CreateLoadBalancerHTTPListenerResponseBody) *CreateLoadBalancerHTTPListenerResponse {
	s.Body = v
	return s
}

type CreateLoadBalancerHTTPSListenerRequest struct {
	AclId                  *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclStatus              *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	AclType                *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	BackendServerPort      *int32  `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	Bandwidth              *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	CACertificateId        *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	Cookie                 *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	CookieTimeout          *int32  `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableHttp2            *string `json:"EnableHttp2,omitempty" xml:"EnableHttp2,omitempty"`
	Gzip                   *string `json:"Gzip,omitempty" xml:"Gzip,omitempty"`
	HealthCheck            *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckConnectPort *int32  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckDomain      *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	HealthCheckHttpCode    *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckInterval    *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckMethod      *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	HealthCheckTimeout     *int32  `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	HealthCheckURI         *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	HealthyThreshold       *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	IdleTimeout            *int32  `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	ListenerPort           *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	LoadBalancerId         *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestTimeout         *int32  `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Scheduler              *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	ServerCertificateId    *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
	StickySession          *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	StickySessionType      *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	TLSCipherPolicy        *string `json:"TLSCipherPolicy,omitempty" xml:"TLSCipherPolicy,omitempty"`
	UnhealthyThreshold     *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	VServerGroupId         *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	XForwardedFor          *string `json:"XForwardedFor,omitempty" xml:"XForwardedFor,omitempty"`
	XForwardedFor_SLBID    *string `json:"XForwardedFor_SLBID,omitempty" xml:"XForwardedFor_SLBID,omitempty"`
	XForwardedFor_SLBIP    *string `json:"XForwardedFor_SLBIP,omitempty" xml:"XForwardedFor_SLBIP,omitempty"`
	XForwardedFor_proto    *string `json:"XForwardedFor_proto,omitempty" xml:"XForwardedFor_proto,omitempty"`
}

func (s CreateLoadBalancerHTTPSListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerHTTPSListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetAclId(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.AclId = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetAclStatus(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.AclStatus = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetAclType(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.AclType = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetBackendServerPort(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.BackendServerPort = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetBandwidth(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetCACertificateId(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.CACertificateId = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetCookie(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.Cookie = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetCookieTimeout(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.CookieTimeout = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetDescription(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.Description = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetEnableHttp2(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.EnableHttp2 = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetGzip(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.Gzip = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheck(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.HealthCheck = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheckConnectPort(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheckDomain(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.HealthCheckDomain = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheckHttpCode(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheckInterval(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheckMethod(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.HealthCheckMethod = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheckTimeout(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.HealthCheckTimeout = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheckURI(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.HealthCheckURI = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetHealthyThreshold(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetIdleTimeout(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.IdleTimeout = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetListenerPort(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetLoadBalancerId(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetOwnerAccount(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetOwnerId(v int64) *CreateLoadBalancerHTTPSListenerRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetRegionId(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetRequestTimeout(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.RequestTimeout = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetResourceOwnerAccount(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetResourceOwnerId(v int64) *CreateLoadBalancerHTTPSListenerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetScheduler(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.Scheduler = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetServerCertificateId(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.ServerCertificateId = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetStickySession(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.StickySession = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetStickySessionType(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.StickySessionType = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetTLSCipherPolicy(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.TLSCipherPolicy = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetUnhealthyThreshold(v int32) *CreateLoadBalancerHTTPSListenerRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetVServerGroupId(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.VServerGroupId = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetXForwardedFor(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.XForwardedFor = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetXForwardedFor_SLBID(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.XForwardedFor_SLBID = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetXForwardedFor_SLBIP(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.XForwardedFor_SLBIP = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerRequest) SetXForwardedFor_proto(v string) *CreateLoadBalancerHTTPSListenerRequest {
	s.XForwardedFor_proto = &v
	return s
}

type CreateLoadBalancerHTTPSListenerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLoadBalancerHTTPSListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerHTTPSListenerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerHTTPSListenerResponseBody) SetRequestId(v string) *CreateLoadBalancerHTTPSListenerResponseBody {
	s.RequestId = &v
	return s
}

type CreateLoadBalancerHTTPSListenerResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLoadBalancerHTTPSListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLoadBalancerHTTPSListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerHTTPSListenerResponse) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerHTTPSListenerResponse) SetHeaders(v map[string]*string) *CreateLoadBalancerHTTPSListenerResponse {
	s.Headers = v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerResponse) SetStatusCode(v int32) *CreateLoadBalancerHTTPSListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerResponse) SetBody(v *CreateLoadBalancerHTTPSListenerResponseBody) *CreateLoadBalancerHTTPSListenerResponse {
	s.Body = v
	return s
}

type CreateLoadBalancerTCPListenerRequest struct {
	AclId                     *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclStatus                 *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	AclType                   *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	BackendServerPort         *int32  `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	Bandwidth                 *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ConnectionDrain           *string `json:"ConnectionDrain,omitempty" xml:"ConnectionDrain,omitempty"`
	ConnectionDrainTimeout    *int32  `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EstablishedTimeout        *int32  `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
	HealthCheckConnectPort    *int32  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckConnectTimeout *int32  `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	HealthCheckDomain         *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	HealthCheckHttpCode       *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckSwitch         *string `json:"HealthCheckSwitch,omitempty" xml:"HealthCheckSwitch,omitempty"`
	HealthCheckType           *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	HealthCheckURI            *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	HealthyThreshold          *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	ListenerPort              *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	LoadBalancerId            *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	MasterSlaveServerGroupId  *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PersistenceTimeout        *int32  `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	ProxyProtocolV2Enabled    *bool   `json:"ProxyProtocolV2Enabled,omitempty" xml:"ProxyProtocolV2Enabled,omitempty"`
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Scheduler                 *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	UnhealthyThreshold        *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	VServerGroupId            *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	HealthCheckInterval       *int32  `json:"healthCheckInterval,omitempty" xml:"healthCheckInterval,omitempty"`
}

func (s CreateLoadBalancerTCPListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerTCPListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerTCPListenerRequest) SetAclId(v string) *CreateLoadBalancerTCPListenerRequest {
	s.AclId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetAclStatus(v string) *CreateLoadBalancerTCPListenerRequest {
	s.AclStatus = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetAclType(v string) *CreateLoadBalancerTCPListenerRequest {
	s.AclType = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetBackendServerPort(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.BackendServerPort = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetBandwidth(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetConnectionDrain(v string) *CreateLoadBalancerTCPListenerRequest {
	s.ConnectionDrain = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetConnectionDrainTimeout(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetDescription(v string) *CreateLoadBalancerTCPListenerRequest {
	s.Description = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetEstablishedTimeout(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.EstablishedTimeout = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckConnectPort(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckConnectTimeout(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckDomain(v string) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckDomain = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckHttpCode(v string) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckSwitch(v string) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckSwitch = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckType(v string) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckType = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckURI(v string) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckURI = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthyThreshold(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetListenerPort(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetLoadBalancerId(v string) *CreateLoadBalancerTCPListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetMasterSlaveServerGroupId(v string) *CreateLoadBalancerTCPListenerRequest {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetOwnerAccount(v string) *CreateLoadBalancerTCPListenerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetOwnerId(v int64) *CreateLoadBalancerTCPListenerRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetPersistenceTimeout(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.PersistenceTimeout = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetProxyProtocolV2Enabled(v bool) *CreateLoadBalancerTCPListenerRequest {
	s.ProxyProtocolV2Enabled = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetRegionId(v string) *CreateLoadBalancerTCPListenerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetResourceOwnerAccount(v string) *CreateLoadBalancerTCPListenerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetResourceOwnerId(v int64) *CreateLoadBalancerTCPListenerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetScheduler(v string) *CreateLoadBalancerTCPListenerRequest {
	s.Scheduler = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetUnhealthyThreshold(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetVServerGroupId(v string) *CreateLoadBalancerTCPListenerRequest {
	s.VServerGroupId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckInterval(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckInterval = &v
	return s
}

type CreateLoadBalancerTCPListenerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLoadBalancerTCPListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerTCPListenerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerTCPListenerResponseBody) SetRequestId(v string) *CreateLoadBalancerTCPListenerResponseBody {
	s.RequestId = &v
	return s
}

type CreateLoadBalancerTCPListenerResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLoadBalancerTCPListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLoadBalancerTCPListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerTCPListenerResponse) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerTCPListenerResponse) SetHeaders(v map[string]*string) *CreateLoadBalancerTCPListenerResponse {
	s.Headers = v
	return s
}

func (s *CreateLoadBalancerTCPListenerResponse) SetStatusCode(v int32) *CreateLoadBalancerTCPListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerResponse) SetBody(v *CreateLoadBalancerTCPListenerResponseBody) *CreateLoadBalancerTCPListenerResponse {
	s.Body = v
	return s
}

type CreateLoadBalancerUDPListenerRequest struct {
	AclId                     *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclStatus                 *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	AclType                   *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	BackendServerPort         *int32  `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	Bandwidth                 *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HealthCheckConnectPort    *int32  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckConnectTimeout *int32  `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	HealthCheckSwitch         *string `json:"HealthCheckSwitch,omitempty" xml:"HealthCheckSwitch,omitempty"`
	HealthyThreshold          *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	ListenerPort              *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	LoadBalancerId            *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	MasterSlaveServerGroupId  *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProxyProtocolV2Enabled    *bool   `json:"ProxyProtocolV2Enabled,omitempty" xml:"ProxyProtocolV2Enabled,omitempty"`
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Scheduler                 *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	UnhealthyThreshold        *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	VServerGroupId            *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	HealthCheckExp            *string `json:"healthCheckExp,omitempty" xml:"healthCheckExp,omitempty"`
	HealthCheckInterval       *int32  `json:"healthCheckInterval,omitempty" xml:"healthCheckInterval,omitempty"`
	HealthCheckReq            *string `json:"healthCheckReq,omitempty" xml:"healthCheckReq,omitempty"`
}

func (s CreateLoadBalancerUDPListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerUDPListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerUDPListenerRequest) SetAclId(v string) *CreateLoadBalancerUDPListenerRequest {
	s.AclId = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetAclStatus(v string) *CreateLoadBalancerUDPListenerRequest {
	s.AclStatus = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetAclType(v string) *CreateLoadBalancerUDPListenerRequest {
	s.AclType = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetBackendServerPort(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.BackendServerPort = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetBandwidth(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetDescription(v string) *CreateLoadBalancerUDPListenerRequest {
	s.Description = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetHealthCheckConnectPort(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetHealthCheckConnectTimeout(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetHealthCheckSwitch(v string) *CreateLoadBalancerUDPListenerRequest {
	s.HealthCheckSwitch = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetHealthyThreshold(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetListenerPort(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetLoadBalancerId(v string) *CreateLoadBalancerUDPListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetMasterSlaveServerGroupId(v string) *CreateLoadBalancerUDPListenerRequest {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetOwnerAccount(v string) *CreateLoadBalancerUDPListenerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetOwnerId(v int64) *CreateLoadBalancerUDPListenerRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetProxyProtocolV2Enabled(v bool) *CreateLoadBalancerUDPListenerRequest {
	s.ProxyProtocolV2Enabled = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetRegionId(v string) *CreateLoadBalancerUDPListenerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetResourceOwnerAccount(v string) *CreateLoadBalancerUDPListenerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetResourceOwnerId(v int64) *CreateLoadBalancerUDPListenerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetScheduler(v string) *CreateLoadBalancerUDPListenerRequest {
	s.Scheduler = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetUnhealthyThreshold(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetVServerGroupId(v string) *CreateLoadBalancerUDPListenerRequest {
	s.VServerGroupId = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetHealthCheckExp(v string) *CreateLoadBalancerUDPListenerRequest {
	s.HealthCheckExp = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetHealthCheckInterval(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetHealthCheckReq(v string) *CreateLoadBalancerUDPListenerRequest {
	s.HealthCheckReq = &v
	return s
}

type CreateLoadBalancerUDPListenerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLoadBalancerUDPListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerUDPListenerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerUDPListenerResponseBody) SetRequestId(v string) *CreateLoadBalancerUDPListenerResponseBody {
	s.RequestId = &v
	return s
}

type CreateLoadBalancerUDPListenerResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLoadBalancerUDPListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLoadBalancerUDPListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerUDPListenerResponse) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerUDPListenerResponse) SetHeaders(v map[string]*string) *CreateLoadBalancerUDPListenerResponse {
	s.Headers = v
	return s
}

func (s *CreateLoadBalancerUDPListenerResponse) SetStatusCode(v int32) *CreateLoadBalancerUDPListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerResponse) SetBody(v *CreateLoadBalancerUDPListenerResponseBody) *CreateLoadBalancerUDPListenerResponse {
	s.Body = v
	return s
}

type CreateMasterSlaveServerGroupRequest struct {
	LoadBalancerId             *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	MasterSlaveBackendServers  *string `json:"MasterSlaveBackendServers,omitempty" xml:"MasterSlaveBackendServers,omitempty"`
	MasterSlaveServerGroupName *string `json:"MasterSlaveServerGroupName,omitempty" xml:"MasterSlaveServerGroupName,omitempty"`
	OwnerAccount               *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount       *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId            *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateMasterSlaveServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMasterSlaveServerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateMasterSlaveServerGroupRequest) SetLoadBalancerId(v string) *CreateMasterSlaveServerGroupRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateMasterSlaveServerGroupRequest) SetMasterSlaveBackendServers(v string) *CreateMasterSlaveServerGroupRequest {
	s.MasterSlaveBackendServers = &v
	return s
}

func (s *CreateMasterSlaveServerGroupRequest) SetMasterSlaveServerGroupName(v string) *CreateMasterSlaveServerGroupRequest {
	s.MasterSlaveServerGroupName = &v
	return s
}

func (s *CreateMasterSlaveServerGroupRequest) SetOwnerAccount(v string) *CreateMasterSlaveServerGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateMasterSlaveServerGroupRequest) SetOwnerId(v int64) *CreateMasterSlaveServerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMasterSlaveServerGroupRequest) SetRegionId(v string) *CreateMasterSlaveServerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMasterSlaveServerGroupRequest) SetResourceOwnerAccount(v string) *CreateMasterSlaveServerGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateMasterSlaveServerGroupRequest) SetResourceOwnerId(v int64) *CreateMasterSlaveServerGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateMasterSlaveServerGroupResponseBody struct {
	MasterSlaveBackendServers *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServers `json:"MasterSlaveBackendServers,omitempty" xml:"MasterSlaveBackendServers,omitempty" type:"Struct"`
	MasterSlaveServerGroupId  *string                                                            `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	RequestId                 *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMasterSlaveServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMasterSlaveServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMasterSlaveServerGroupResponseBody) SetMasterSlaveBackendServers(v *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServers) *CreateMasterSlaveServerGroupResponseBody {
	s.MasterSlaveBackendServers = v
	return s
}

func (s *CreateMasterSlaveServerGroupResponseBody) SetMasterSlaveServerGroupId(v string) *CreateMasterSlaveServerGroupResponseBody {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *CreateMasterSlaveServerGroupResponseBody) SetRequestId(v string) *CreateMasterSlaveServerGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServers struct {
	MasterSlaveBackendServer []*CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer `json:"MasterSlaveBackendServer,omitempty" xml:"MasterSlaveBackendServer,omitempty" type:"Repeated"`
}

func (s CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServers) String() string {
	return tea.Prettify(s)
}

func (s CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServers) GoString() string {
	return s.String()
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServers) SetMasterSlaveBackendServer(v []*CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServers {
	s.MasterSlaveBackendServer = v
	return s
}

type CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Port        *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ServerId    *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	ServerType  *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight      *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) String() string {
	return tea.Prettify(s)
}

func (s CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) GoString() string {
	return s.String()
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetDescription(v string) *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.Description = &v
	return s
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetPort(v int32) *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.Port = &v
	return s
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetServerId(v string) *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.ServerId = &v
	return s
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetServerType(v string) *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.ServerType = &v
	return s
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetType(v string) *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.Type = &v
	return s
}

func (s *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetWeight(v int32) *CreateMasterSlaveServerGroupResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.Weight = &v
	return s
}

type CreateMasterSlaveServerGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMasterSlaveServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMasterSlaveServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMasterSlaveServerGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateMasterSlaveServerGroupResponse) SetHeaders(v map[string]*string) *CreateMasterSlaveServerGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateMasterSlaveServerGroupResponse) SetStatusCode(v int32) *CreateMasterSlaveServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMasterSlaveServerGroupResponse) SetBody(v *CreateMasterSlaveServerGroupResponseBody) *CreateMasterSlaveServerGroupResponse {
	s.Body = v
	return s
}

type CreateRulesRequest struct {
	ListenerPort         *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerProtocol     *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RuleList             *string `json:"RuleList,omitempty" xml:"RuleList,omitempty"`
}

func (s CreateRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateRulesRequest) SetListenerPort(v int32) *CreateRulesRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateRulesRequest) SetListenerProtocol(v string) *CreateRulesRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *CreateRulesRequest) SetLoadBalancerId(v string) *CreateRulesRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateRulesRequest) SetOwnerAccount(v string) *CreateRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateRulesRequest) SetOwnerId(v int64) *CreateRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRulesRequest) SetRegionId(v string) *CreateRulesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRulesRequest) SetResourceOwnerAccount(v string) *CreateRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateRulesRequest) SetResourceOwnerId(v int64) *CreateRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateRulesRequest) SetRuleList(v string) *CreateRulesRequest {
	s.RuleList = &v
	return s
}

type CreateRulesResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     *CreateRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s CreateRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRulesResponseBody) SetRequestId(v string) *CreateRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRulesResponseBody) SetRules(v *CreateRulesResponseBodyRules) *CreateRulesResponseBody {
	s.Rules = v
	return s
}

type CreateRulesResponseBodyRules struct {
	Rule []*CreateRulesResponseBodyRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s CreateRulesResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *CreateRulesResponseBodyRules) SetRule(v []*CreateRulesResponseBodyRulesRule) *CreateRulesResponseBodyRules {
	s.Rule = v
	return s
}

type CreateRulesResponseBodyRulesRule struct {
	RuleId   *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s CreateRulesResponseBodyRulesRule) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesResponseBodyRulesRule) GoString() string {
	return s.String()
}

func (s *CreateRulesResponseBodyRulesRule) SetRuleId(v string) *CreateRulesResponseBodyRulesRule {
	s.RuleId = &v
	return s
}

func (s *CreateRulesResponseBodyRulesRule) SetRuleName(v string) *CreateRulesResponseBodyRulesRule {
	s.RuleName = &v
	return s
}

type CreateRulesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateRulesResponse) SetHeaders(v map[string]*string) *CreateRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateRulesResponse) SetStatusCode(v int32) *CreateRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRulesResponse) SetBody(v *CreateRulesResponseBody) *CreateRulesResponse {
	s.Body = v
	return s
}

type CreateTLSCipherPolicyRequest struct {
	Ciphers              []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	Name                 *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TLSVersions          []*string `json:"TLSVersions,omitempty" xml:"TLSVersions,omitempty" type:"Repeated"`
}

func (s CreateTLSCipherPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTLSCipherPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateTLSCipherPolicyRequest) SetCiphers(v []*string) *CreateTLSCipherPolicyRequest {
	s.Ciphers = v
	return s
}

func (s *CreateTLSCipherPolicyRequest) SetName(v string) *CreateTLSCipherPolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateTLSCipherPolicyRequest) SetOwnerAccount(v string) *CreateTLSCipherPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTLSCipherPolicyRequest) SetOwnerId(v int64) *CreateTLSCipherPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTLSCipherPolicyRequest) SetRegionId(v string) *CreateTLSCipherPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTLSCipherPolicyRequest) SetResourceOwnerAccount(v string) *CreateTLSCipherPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTLSCipherPolicyRequest) SetResourceOwnerId(v int64) *CreateTLSCipherPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTLSCipherPolicyRequest) SetTLSVersions(v []*string) *CreateTLSCipherPolicyRequest {
	s.TLSVersions = v
	return s
}

type CreateTLSCipherPolicyResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TLSCipherPolicyId *string `json:"TLSCipherPolicyId,omitempty" xml:"TLSCipherPolicyId,omitempty"`
}

func (s CreateTLSCipherPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTLSCipherPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTLSCipherPolicyResponseBody) SetRequestId(v string) *CreateTLSCipherPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTLSCipherPolicyResponseBody) SetTLSCipherPolicyId(v string) *CreateTLSCipherPolicyResponseBody {
	s.TLSCipherPolicyId = &v
	return s
}

type CreateTLSCipherPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTLSCipherPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTLSCipherPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTLSCipherPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateTLSCipherPolicyResponse) SetHeaders(v map[string]*string) *CreateTLSCipherPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateTLSCipherPolicyResponse) SetStatusCode(v int32) *CreateTLSCipherPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTLSCipherPolicyResponse) SetBody(v *CreateTLSCipherPolicyResponseBody) *CreateTLSCipherPolicyResponse {
	s.Body = v
	return s
}

type CreateVServerGroupRequest struct {
	BackendServers       *string `json:"BackendServers,omitempty" xml:"BackendServers,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VServerGroupName     *string `json:"VServerGroupName,omitempty" xml:"VServerGroupName,omitempty"`
}

func (s CreateVServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVServerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateVServerGroupRequest) SetBackendServers(v string) *CreateVServerGroupRequest {
	s.BackendServers = &v
	return s
}

func (s *CreateVServerGroupRequest) SetLoadBalancerId(v string) *CreateVServerGroupRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateVServerGroupRequest) SetOwnerAccount(v string) *CreateVServerGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVServerGroupRequest) SetOwnerId(v int64) *CreateVServerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVServerGroupRequest) SetRegionId(v string) *CreateVServerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVServerGroupRequest) SetResourceOwnerAccount(v string) *CreateVServerGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVServerGroupRequest) SetResourceOwnerId(v int64) *CreateVServerGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVServerGroupRequest) SetVServerGroupName(v string) *CreateVServerGroupRequest {
	s.VServerGroupName = &v
	return s
}

type CreateVServerGroupResponseBody struct {
	BackendServers *CreateVServerGroupResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VServerGroupId *string                                       `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s CreateVServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVServerGroupResponseBody) SetBackendServers(v *CreateVServerGroupResponseBodyBackendServers) *CreateVServerGroupResponseBody {
	s.BackendServers = v
	return s
}

func (s *CreateVServerGroupResponseBody) SetRequestId(v string) *CreateVServerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVServerGroupResponseBody) SetVServerGroupId(v string) *CreateVServerGroupResponseBody {
	s.VServerGroupId = &v
	return s
}

type CreateVServerGroupResponseBodyBackendServers struct {
	BackendServer []*CreateVServerGroupResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s CreateVServerGroupResponseBodyBackendServers) String() string {
	return tea.Prettify(s)
}

func (s CreateVServerGroupResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *CreateVServerGroupResponseBodyBackendServers) SetBackendServer(v []*CreateVServerGroupResponseBodyBackendServersBackendServer) *CreateVServerGroupResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

type CreateVServerGroupResponseBodyBackendServersBackendServer struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Port        *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ServerId    *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight      *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateVServerGroupResponseBodyBackendServersBackendServer) String() string {
	return tea.Prettify(s)
}

func (s CreateVServerGroupResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *CreateVServerGroupResponseBodyBackendServersBackendServer) SetDescription(v string) *CreateVServerGroupResponseBodyBackendServersBackendServer {
	s.Description = &v
	return s
}

func (s *CreateVServerGroupResponseBodyBackendServersBackendServer) SetPort(v int32) *CreateVServerGroupResponseBodyBackendServersBackendServer {
	s.Port = &v
	return s
}

func (s *CreateVServerGroupResponseBodyBackendServersBackendServer) SetServerId(v string) *CreateVServerGroupResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *CreateVServerGroupResponseBodyBackendServersBackendServer) SetType(v string) *CreateVServerGroupResponseBodyBackendServersBackendServer {
	s.Type = &v
	return s
}

func (s *CreateVServerGroupResponseBodyBackendServersBackendServer) SetWeight(v int32) *CreateVServerGroupResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

type CreateVServerGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVServerGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateVServerGroupResponse) SetHeaders(v map[string]*string) *CreateVServerGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateVServerGroupResponse) SetStatusCode(v int32) *CreateVServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVServerGroupResponse) SetBody(v *CreateVServerGroupResponseBody) *CreateVServerGroupResponse {
	s.Body = v
	return s
}

type DeleteAccessControlListRequest struct {
	AclId                *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteAccessControlListRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessControlListRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessControlListRequest) SetAclId(v string) *DeleteAccessControlListRequest {
	s.AclId = &v
	return s
}

func (s *DeleteAccessControlListRequest) SetOwnerAccount(v string) *DeleteAccessControlListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteAccessControlListRequest) SetOwnerId(v int64) *DeleteAccessControlListRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAccessControlListRequest) SetRegionId(v string) *DeleteAccessControlListRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAccessControlListRequest) SetResourceOwnerAccount(v string) *DeleteAccessControlListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAccessControlListRequest) SetResourceOwnerId(v int64) *DeleteAccessControlListRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteAccessControlListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessControlListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessControlListResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessControlListResponseBody) SetRequestId(v string) *DeleteAccessControlListResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccessControlListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAccessControlListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAccessControlListResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessControlListResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessControlListResponse) SetHeaders(v map[string]*string) *DeleteAccessControlListResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessControlListResponse) SetStatusCode(v int32) *DeleteAccessControlListResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessControlListResponse) SetBody(v *DeleteAccessControlListResponseBody) *DeleteAccessControlListResponse {
	s.Body = v
	return s
}

type DeleteCACertificateRequest struct {
	CACertificateId      *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteCACertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCACertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteCACertificateRequest) SetCACertificateId(v string) *DeleteCACertificateRequest {
	s.CACertificateId = &v
	return s
}

func (s *DeleteCACertificateRequest) SetOwnerAccount(v string) *DeleteCACertificateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCACertificateRequest) SetOwnerId(v int64) *DeleteCACertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCACertificateRequest) SetRegionId(v string) *DeleteCACertificateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCACertificateRequest) SetResourceOwnerAccount(v string) *DeleteCACertificateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCACertificateRequest) SetResourceOwnerId(v int64) *DeleteCACertificateRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteCACertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCACertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCACertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCACertificateResponseBody) SetRequestId(v string) *DeleteCACertificateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCACertificateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCACertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCACertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCACertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteCACertificateResponse) SetHeaders(v map[string]*string) *DeleteCACertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteCACertificateResponse) SetStatusCode(v int32) *DeleteCACertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCACertificateResponse) SetBody(v *DeleteCACertificateResponseBody) *DeleteCACertificateResponse {
	s.Body = v
	return s
}

type DeleteDomainExtensionRequest struct {
	DomainExtensionId    *string `json:"DomainExtensionId,omitempty" xml:"DomainExtensionId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDomainExtensionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainExtensionRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainExtensionRequest) SetDomainExtensionId(v string) *DeleteDomainExtensionRequest {
	s.DomainExtensionId = &v
	return s
}

func (s *DeleteDomainExtensionRequest) SetOwnerAccount(v string) *DeleteDomainExtensionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDomainExtensionRequest) SetOwnerId(v int64) *DeleteDomainExtensionRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDomainExtensionRequest) SetRegionId(v string) *DeleteDomainExtensionRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDomainExtensionRequest) SetResourceOwnerAccount(v string) *DeleteDomainExtensionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDomainExtensionRequest) SetResourceOwnerId(v int64) *DeleteDomainExtensionRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteDomainExtensionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainExtensionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainExtensionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainExtensionResponseBody) SetRequestId(v string) *DeleteDomainExtensionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDomainExtensionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDomainExtensionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDomainExtensionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainExtensionResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainExtensionResponse) SetHeaders(v map[string]*string) *DeleteDomainExtensionResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainExtensionResponse) SetStatusCode(v int32) *DeleteDomainExtensionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainExtensionResponse) SetBody(v *DeleteDomainExtensionResponseBody) *DeleteDomainExtensionResponse {
	s.Body = v
	return s
}

type DeleteLoadBalancerRequest struct {
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteLoadBalancerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerRequest) SetLoadBalancerId(v string) *DeleteLoadBalancerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetOwnerAccount(v string) *DeleteLoadBalancerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetOwnerId(v int64) *DeleteLoadBalancerRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetRegionId(v string) *DeleteLoadBalancerRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetResourceOwnerAccount(v string) *DeleteLoadBalancerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetResourceOwnerId(v int64) *DeleteLoadBalancerRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteLoadBalancerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLoadBalancerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerResponseBody) SetRequestId(v string) *DeleteLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLoadBalancerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLoadBalancerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLoadBalancerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerResponse) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerResponse) SetHeaders(v map[string]*string) *DeleteLoadBalancerResponse {
	s.Headers = v
	return s
}

func (s *DeleteLoadBalancerResponse) SetStatusCode(v int32) *DeleteLoadBalancerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLoadBalancerResponse) SetBody(v *DeleteLoadBalancerResponseBody) *DeleteLoadBalancerResponse {
	s.Body = v
	return s
}

type DeleteLoadBalancerListenerRequest struct {
	ListenerPort         *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerProtocol     *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteLoadBalancerListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerListenerRequest) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerListenerRequest) SetListenerPort(v int32) *DeleteLoadBalancerListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetListenerProtocol(v string) *DeleteLoadBalancerListenerRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetLoadBalancerId(v string) *DeleteLoadBalancerListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetOwnerAccount(v string) *DeleteLoadBalancerListenerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetOwnerId(v int64) *DeleteLoadBalancerListenerRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetRegionId(v string) *DeleteLoadBalancerListenerRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetResourceOwnerAccount(v string) *DeleteLoadBalancerListenerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetResourceOwnerId(v int64) *DeleteLoadBalancerListenerRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteLoadBalancerListenerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLoadBalancerListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerListenerResponseBody) SetRequestId(v string) *DeleteLoadBalancerListenerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLoadBalancerListenerResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLoadBalancerListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLoadBalancerListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerListenerResponse) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerListenerResponse) SetHeaders(v map[string]*string) *DeleteLoadBalancerListenerResponse {
	s.Headers = v
	return s
}

func (s *DeleteLoadBalancerListenerResponse) SetStatusCode(v int32) *DeleteLoadBalancerListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLoadBalancerListenerResponse) SetBody(v *DeleteLoadBalancerListenerResponseBody) *DeleteLoadBalancerListenerResponse {
	s.Body = v
	return s
}

type DeleteMasterSlaveServerGroupRequest struct {
	MasterSlaveServerGroupId *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount     *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteMasterSlaveServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMasterSlaveServerGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteMasterSlaveServerGroupRequest) SetMasterSlaveServerGroupId(v string) *DeleteMasterSlaveServerGroupRequest {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *DeleteMasterSlaveServerGroupRequest) SetOwnerAccount(v string) *DeleteMasterSlaveServerGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteMasterSlaveServerGroupRequest) SetOwnerId(v int64) *DeleteMasterSlaveServerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMasterSlaveServerGroupRequest) SetRegionId(v string) *DeleteMasterSlaveServerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMasterSlaveServerGroupRequest) SetResourceOwnerAccount(v string) *DeleteMasterSlaveServerGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteMasterSlaveServerGroupRequest) SetResourceOwnerId(v int64) *DeleteMasterSlaveServerGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteMasterSlaveServerGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMasterSlaveServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMasterSlaveServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMasterSlaveServerGroupResponseBody) SetRequestId(v string) *DeleteMasterSlaveServerGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMasterSlaveServerGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteMasterSlaveServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMasterSlaveServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMasterSlaveServerGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteMasterSlaveServerGroupResponse) SetHeaders(v map[string]*string) *DeleteMasterSlaveServerGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteMasterSlaveServerGroupResponse) SetStatusCode(v int32) *DeleteMasterSlaveServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMasterSlaveServerGroupResponse) SetBody(v *DeleteMasterSlaveServerGroupResponseBody) *DeleteMasterSlaveServerGroupResponse {
	s.Body = v
	return s
}

type DeleteRulesRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RuleIds              *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
}

func (s DeleteRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteRulesRequest) SetOwnerAccount(v string) *DeleteRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteRulesRequest) SetOwnerId(v int64) *DeleteRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRulesRequest) SetRegionId(v string) *DeleteRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRulesRequest) SetResourceOwnerAccount(v string) *DeleteRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteRulesRequest) SetResourceOwnerId(v int64) *DeleteRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteRulesRequest) SetRuleIds(v string) *DeleteRulesRequest {
	s.RuleIds = &v
	return s
}

type DeleteRulesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRulesResponseBody) SetRequestId(v string) *DeleteRulesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRulesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteRulesResponse) SetHeaders(v map[string]*string) *DeleteRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteRulesResponse) SetStatusCode(v int32) *DeleteRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRulesResponse) SetBody(v *DeleteRulesResponseBody) *DeleteRulesResponse {
	s.Body = v
	return s
}

type DeleteServerCertificateRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ServerCertificateId  *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
}

func (s DeleteServerCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteServerCertificateRequest) SetOwnerAccount(v string) *DeleteServerCertificateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteServerCertificateRequest) SetOwnerId(v int64) *DeleteServerCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteServerCertificateRequest) SetRegionId(v string) *DeleteServerCertificateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServerCertificateRequest) SetResourceOwnerAccount(v string) *DeleteServerCertificateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteServerCertificateRequest) SetResourceOwnerId(v int64) *DeleteServerCertificateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteServerCertificateRequest) SetServerCertificateId(v string) *DeleteServerCertificateRequest {
	s.ServerCertificateId = &v
	return s
}

type DeleteServerCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServerCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServerCertificateResponseBody) SetRequestId(v string) *DeleteServerCertificateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServerCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteServerCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServerCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteServerCertificateResponse) SetHeaders(v map[string]*string) *DeleteServerCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteServerCertificateResponse) SetStatusCode(v int32) *DeleteServerCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServerCertificateResponse) SetBody(v *DeleteServerCertificateResponseBody) *DeleteServerCertificateResponse {
	s.Body = v
	return s
}

type DeleteTLSCipherPolicyRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TLSCipherPolicyId    *string `json:"TLSCipherPolicyId,omitempty" xml:"TLSCipherPolicyId,omitempty"`
}

func (s DeleteTLSCipherPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTLSCipherPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteTLSCipherPolicyRequest) SetOwnerAccount(v string) *DeleteTLSCipherPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTLSCipherPolicyRequest) SetOwnerId(v int64) *DeleteTLSCipherPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTLSCipherPolicyRequest) SetRegionId(v string) *DeleteTLSCipherPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTLSCipherPolicyRequest) SetResourceOwnerAccount(v string) *DeleteTLSCipherPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTLSCipherPolicyRequest) SetResourceOwnerId(v int64) *DeleteTLSCipherPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTLSCipherPolicyRequest) SetTLSCipherPolicyId(v string) *DeleteTLSCipherPolicyRequest {
	s.TLSCipherPolicyId = &v
	return s
}

type DeleteTLSCipherPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTLSCipherPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTLSCipherPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTLSCipherPolicyResponseBody) SetRequestId(v string) *DeleteTLSCipherPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTLSCipherPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTLSCipherPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTLSCipherPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTLSCipherPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteTLSCipherPolicyResponse) SetHeaders(v map[string]*string) *DeleteTLSCipherPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteTLSCipherPolicyResponse) SetStatusCode(v int32) *DeleteTLSCipherPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTLSCipherPolicyResponse) SetBody(v *DeleteTLSCipherPolicyResponseBody) *DeleteTLSCipherPolicyResponse {
	s.Body = v
	return s
}

type DeleteVServerGroupRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VServerGroupId       *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DeleteVServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVServerGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteVServerGroupRequest) SetOwnerAccount(v string) *DeleteVServerGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVServerGroupRequest) SetOwnerId(v int64) *DeleteVServerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVServerGroupRequest) SetRegionId(v string) *DeleteVServerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVServerGroupRequest) SetResourceOwnerAccount(v string) *DeleteVServerGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVServerGroupRequest) SetResourceOwnerId(v int64) *DeleteVServerGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVServerGroupRequest) SetVServerGroupId(v string) *DeleteVServerGroupRequest {
	s.VServerGroupId = &v
	return s
}

type DeleteVServerGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVServerGroupResponseBody) SetRequestId(v string) *DeleteVServerGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVServerGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVServerGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteVServerGroupResponse) SetHeaders(v map[string]*string) *DeleteVServerGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteVServerGroupResponse) SetStatusCode(v int32) *DeleteVServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVServerGroupResponse) SetBody(v *DeleteVServerGroupResponseBody) *DeleteVServerGroupResponse {
	s.Body = v
	return s
}

type DescribeAccessControlListAttributeRequest struct {
	AclEntryComment      *string `json:"AclEntryComment,omitempty" xml:"AclEntryComment,omitempty"`
	AclId                *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAccessControlListAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessControlListAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListAttributeRequest) SetAclEntryComment(v string) *DescribeAccessControlListAttributeRequest {
	s.AclEntryComment = &v
	return s
}

func (s *DescribeAccessControlListAttributeRequest) SetAclId(v string) *DescribeAccessControlListAttributeRequest {
	s.AclId = &v
	return s
}

func (s *DescribeAccessControlListAttributeRequest) SetOwnerAccount(v string) *DescribeAccessControlListAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAccessControlListAttributeRequest) SetOwnerId(v int64) *DescribeAccessControlListAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccessControlListAttributeRequest) SetRegionId(v string) *DescribeAccessControlListAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccessControlListAttributeRequest) SetResourceOwnerAccount(v string) *DescribeAccessControlListAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAccessControlListAttributeRequest) SetResourceOwnerId(v int64) *DescribeAccessControlListAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeAccessControlListAttributeResponseBody struct {
	AclEntrys        *DescribeAccessControlListAttributeResponseBodyAclEntrys        `json:"AclEntrys,omitempty" xml:"AclEntrys,omitempty" type:"Struct"`
	AclId            *string                                                         `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclName          *string                                                         `json:"AclName,omitempty" xml:"AclName,omitempty"`
	AddressIPVersion *string                                                         `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	CreateTime       *string                                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RelatedListeners *DescribeAccessControlListAttributeResponseBodyRelatedListeners `json:"RelatedListeners,omitempty" xml:"RelatedListeners,omitempty" type:"Struct"`
	RequestId        *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId  *string                                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeAccessControlListAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessControlListAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListAttributeResponseBody) SetAclEntrys(v *DescribeAccessControlListAttributeResponseBodyAclEntrys) *DescribeAccessControlListAttributeResponseBody {
	s.AclEntrys = v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) SetAclId(v string) *DescribeAccessControlListAttributeResponseBody {
	s.AclId = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) SetAclName(v string) *DescribeAccessControlListAttributeResponseBody {
	s.AclName = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) SetAddressIPVersion(v string) *DescribeAccessControlListAttributeResponseBody {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) SetCreateTime(v string) *DescribeAccessControlListAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) SetRelatedListeners(v *DescribeAccessControlListAttributeResponseBodyRelatedListeners) *DescribeAccessControlListAttributeResponseBody {
	s.RelatedListeners = v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) SetRequestId(v string) *DescribeAccessControlListAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) SetResourceGroupId(v string) *DescribeAccessControlListAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

type DescribeAccessControlListAttributeResponseBodyAclEntrys struct {
	AclEntry []*DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry `json:"AclEntry,omitempty" xml:"AclEntry,omitempty" type:"Repeated"`
}

func (s DescribeAccessControlListAttributeResponseBodyAclEntrys) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessControlListAttributeResponseBodyAclEntrys) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListAttributeResponseBodyAclEntrys) SetAclEntry(v []*DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) *DescribeAccessControlListAttributeResponseBodyAclEntrys {
	s.AclEntry = v
	return s
}

type DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry struct {
	AclEntryComment *string `json:"AclEntryComment,omitempty" xml:"AclEntryComment,omitempty"`
	AclEntryIP      *string `json:"AclEntryIP,omitempty" xml:"AclEntryIP,omitempty"`
}

func (s DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) SetAclEntryComment(v string) *DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry {
	s.AclEntryComment = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) SetAclEntryIP(v string) *DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry {
	s.AclEntryIP = &v
	return s
}

type DescribeAccessControlListAttributeResponseBodyRelatedListeners struct {
	RelatedListener []*DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener `json:"RelatedListener,omitempty" xml:"RelatedListener,omitempty" type:"Repeated"`
}

func (s DescribeAccessControlListAttributeResponseBodyRelatedListeners) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessControlListAttributeResponseBodyRelatedListeners) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListAttributeResponseBodyRelatedListeners) SetRelatedListener(v []*DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener) *DescribeAccessControlListAttributeResponseBodyRelatedListeners {
	s.RelatedListener = v
	return s
}

type DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener struct {
	AclType        *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	ListenerPort   *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	Protocol       *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener) SetAclType(v string) *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener {
	s.AclType = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener) SetListenerPort(v int32) *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener {
	s.ListenerPort = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener) SetLoadBalancerId(v string) *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener) SetProtocol(v string) *DescribeAccessControlListAttributeResponseBodyRelatedListenersRelatedListener {
	s.Protocol = &v
	return s
}

type DescribeAccessControlListAttributeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAccessControlListAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccessControlListAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessControlListAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListAttributeResponse) SetHeaders(v map[string]*string) *DescribeAccessControlListAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessControlListAttributeResponse) SetStatusCode(v int32) *DescribeAccessControlListAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponse) SetBody(v *DescribeAccessControlListAttributeResponseBody) *DescribeAccessControlListAttributeResponse {
	s.Body = v
	return s
}

type DescribeAccessControlListsRequest struct {
	AclName              *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	AddressIPVersion     *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAccessControlListsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessControlListsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListsRequest) SetAclName(v string) *DescribeAccessControlListsRequest {
	s.AclName = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetAddressIPVersion(v string) *DescribeAccessControlListsRequest {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetOwnerAccount(v string) *DescribeAccessControlListsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetOwnerId(v int64) *DescribeAccessControlListsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetPageNumber(v int32) *DescribeAccessControlListsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetPageSize(v int32) *DescribeAccessControlListsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetRegionId(v string) *DescribeAccessControlListsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetResourceGroupId(v string) *DescribeAccessControlListsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetResourceOwnerAccount(v string) *DescribeAccessControlListsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetResourceOwnerId(v int64) *DescribeAccessControlListsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeAccessControlListsResponseBody struct {
	Acls       *DescribeAccessControlListsResponseBodyAcls `json:"Acls,omitempty" xml:"Acls,omitempty" type:"Struct"`
	Count      *int32                                      `json:"Count,omitempty" xml:"Count,omitempty"`
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccessControlListsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessControlListsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListsResponseBody) SetAcls(v *DescribeAccessControlListsResponseBodyAcls) *DescribeAccessControlListsResponseBody {
	s.Acls = v
	return s
}

func (s *DescribeAccessControlListsResponseBody) SetCount(v int32) *DescribeAccessControlListsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeAccessControlListsResponseBody) SetPageNumber(v int32) *DescribeAccessControlListsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessControlListsResponseBody) SetPageSize(v int32) *DescribeAccessControlListsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessControlListsResponseBody) SetRequestId(v string) *DescribeAccessControlListsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessControlListsResponseBody) SetTotalCount(v int32) *DescribeAccessControlListsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAccessControlListsResponseBodyAcls struct {
	Acl []*DescribeAccessControlListsResponseBodyAclsAcl `json:"Acl,omitempty" xml:"Acl,omitempty" type:"Repeated"`
}

func (s DescribeAccessControlListsResponseBodyAcls) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessControlListsResponseBodyAcls) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListsResponseBodyAcls) SetAcl(v []*DescribeAccessControlListsResponseBodyAclsAcl) *DescribeAccessControlListsResponseBodyAcls {
	s.Acl = v
	return s
}

type DescribeAccessControlListsResponseBodyAclsAcl struct {
	AclId            *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclName          *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ResourceGroupId  *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeAccessControlListsResponseBodyAclsAcl) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessControlListsResponseBodyAclsAcl) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) SetAclId(v string) *DescribeAccessControlListsResponseBodyAclsAcl {
	s.AclId = &v
	return s
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) SetAclName(v string) *DescribeAccessControlListsResponseBodyAclsAcl {
	s.AclName = &v
	return s
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) SetAddressIPVersion(v string) *DescribeAccessControlListsResponseBodyAclsAcl {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) SetCreateTime(v string) *DescribeAccessControlListsResponseBodyAclsAcl {
	s.CreateTime = &v
	return s
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) SetResourceGroupId(v string) *DescribeAccessControlListsResponseBodyAclsAcl {
	s.ResourceGroupId = &v
	return s
}

type DescribeAccessControlListsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAccessControlListsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccessControlListsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessControlListsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListsResponse) SetHeaders(v map[string]*string) *DescribeAccessControlListsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessControlListsResponse) SetStatusCode(v int32) *DescribeAccessControlListsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessControlListsResponse) SetBody(v *DescribeAccessControlListsResponseBody) *DescribeAccessControlListsResponse {
	s.Body = v
	return s
}

type DescribeAccessLogsDownloadAttributeRequest struct {
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LogType              *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s DescribeAccessLogsDownloadAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessLogsDownloadAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessLogsDownloadAttributeRequest) SetLoadBalancerId(v string) *DescribeAccessLogsDownloadAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeRequest) SetLogType(v string) *DescribeAccessLogsDownloadAttributeRequest {
	s.LogType = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeRequest) SetOwnerAccount(v string) *DescribeAccessLogsDownloadAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeRequest) SetOwnerId(v int64) *DescribeAccessLogsDownloadAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeRequest) SetPageNumber(v int32) *DescribeAccessLogsDownloadAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeRequest) SetPageSize(v int32) *DescribeAccessLogsDownloadAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeRequest) SetRegionId(v string) *DescribeAccessLogsDownloadAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeRequest) SetResourceOwnerAccount(v string) *DescribeAccessLogsDownloadAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeRequest) SetResourceOwnerId(v int64) *DescribeAccessLogsDownloadAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeRequest) SetTags(v string) *DescribeAccessLogsDownloadAttributeRequest {
	s.Tags = &v
	return s
}

type DescribeAccessLogsDownloadAttributeResponseBody struct {
	Count                  *int32                                                                 `json:"Count,omitempty" xml:"Count,omitempty"`
	LogsDownloadAttributes *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributes `json:"LogsDownloadAttributes,omitempty" xml:"LogsDownloadAttributes,omitempty" type:"Struct"`
	PageNumber             *int32                                                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize               *int32                                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId              *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount             *int32                                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccessLogsDownloadAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessLogsDownloadAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessLogsDownloadAttributeResponseBody) SetCount(v int32) *DescribeAccessLogsDownloadAttributeResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponseBody) SetLogsDownloadAttributes(v *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributes) *DescribeAccessLogsDownloadAttributeResponseBody {
	s.LogsDownloadAttributes = v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponseBody) SetPageNumber(v int32) *DescribeAccessLogsDownloadAttributeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponseBody) SetPageSize(v int32) *DescribeAccessLogsDownloadAttributeResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponseBody) SetRequestId(v string) *DescribeAccessLogsDownloadAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponseBody) SetTotalCount(v int32) *DescribeAccessLogsDownloadAttributeResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributes struct {
	LogsDownloadAttribute []*DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute `json:"LogsDownloadAttribute,omitempty" xml:"LogsDownloadAttribute,omitempty" type:"Repeated"`
}

func (s DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributes) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributes) GoString() string {
	return s.String()
}

func (s *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributes) SetLogsDownloadAttribute(v []*DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributes {
	s.LogsDownloadAttribute = v
	return s
}

type DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute struct {
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LogProject     *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	LogStore       *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	LogType        *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	Region         *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) GoString() string {
	return s.String()
}

func (s *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) SetLoadBalancerId(v string) *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) SetLogProject(v string) *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute {
	s.LogProject = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) SetLogStore(v string) *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute {
	s.LogStore = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) SetLogType(v string) *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute {
	s.LogType = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute) SetRegion(v string) *DescribeAccessLogsDownloadAttributeResponseBodyLogsDownloadAttributesLogsDownloadAttribute {
	s.Region = &v
	return s
}

type DescribeAccessLogsDownloadAttributeResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAccessLogsDownloadAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccessLogsDownloadAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessLogsDownloadAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessLogsDownloadAttributeResponse) SetHeaders(v map[string]*string) *DescribeAccessLogsDownloadAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponse) SetStatusCode(v int32) *DescribeAccessLogsDownloadAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponse) SetBody(v *DescribeAccessLogsDownloadAttributeResponseBody) *DescribeAccessLogsDownloadAttributeResponse {
	s.Body = v
	return s
}

type DescribeAvailableResourceRequest struct {
	AddressIPVersion     *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	AddressType          *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAvailableResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequest) SetAddressIPVersion(v string) *DescribeAvailableResourceRequest {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetAddressType(v string) *DescribeAvailableResourceRequest {
	s.AddressType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetRegionId(v string) *DescribeAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeAvailableResourceResponseBody struct {
	AvailableResources *DescribeAvailableResourceResponseBodyAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" type:"Struct"`
	RequestId          *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBody) SetAvailableResources(v *DescribeAvailableResourceResponseBodyAvailableResources) *DescribeAvailableResourceResponseBody {
	s.AvailableResources = v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetRequestId(v string) *DescribeAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableResources struct {
	AvailableResource []*DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource `json:"AvailableResource,omitempty" xml:"AvailableResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableResources) SetAvailableResource(v []*DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource) *DescribeAvailableResourceResponseBodyAvailableResources {
	s.AvailableResource = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource struct {
	MasterZoneId     *string                                                                                   `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	SlaveZoneId      *string                                                                                   `json:"SlaveZoneId,omitempty" xml:"SlaveZoneId,omitempty"`
	SupportResources *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResources `json:"SupportResources,omitempty" xml:"SupportResources,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource) SetMasterZoneId(v string) *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource {
	s.MasterZoneId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource) SetSlaveZoneId(v string) *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource {
	s.SlaveZoneId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource) SetSupportResources(v *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResources) *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResource {
	s.SupportResources = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResources struct {
	SupportResource []*DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource `json:"SupportResource,omitempty" xml:"SupportResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResources) SetSupportResource(v []*DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource) *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResources {
	s.SupportResource = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource struct {
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	AddressType      *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource) SetAddressIPVersion(v string) *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource) SetAddressType(v string) *DescribeAvailableResourceResponseBodyAvailableResourcesAvailableResourceSupportResourcesSupportResource {
	s.AddressType = &v
	return s
}

type DescribeAvailableResourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAvailableResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAvailableResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponse) SetHeaders(v map[string]*string) *DescribeAvailableResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableResourceResponse) SetStatusCode(v int32) *DescribeAvailableResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableResourceResponse) SetBody(v *DescribeAvailableResourceResponseBody) *DescribeAvailableResourceResponse {
	s.Body = v
	return s
}

type DescribeCACertificatesRequest struct {
	CACertificateId      *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeCACertificatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCACertificatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCACertificatesRequest) SetCACertificateId(v string) *DescribeCACertificatesRequest {
	s.CACertificateId = &v
	return s
}

func (s *DescribeCACertificatesRequest) SetOwnerAccount(v string) *DescribeCACertificatesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCACertificatesRequest) SetOwnerId(v int64) *DescribeCACertificatesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCACertificatesRequest) SetRegionId(v string) *DescribeCACertificatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCACertificatesRequest) SetResourceGroupId(v string) *DescribeCACertificatesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCACertificatesRequest) SetResourceOwnerAccount(v string) *DescribeCACertificatesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCACertificatesRequest) SetResourceOwnerId(v int64) *DescribeCACertificatesRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeCACertificatesResponseBody struct {
	CACertificates *DescribeCACertificatesResponseBodyCACertificates `json:"CACertificates,omitempty" xml:"CACertificates,omitempty" type:"Struct"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCACertificatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCACertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCACertificatesResponseBody) SetCACertificates(v *DescribeCACertificatesResponseBodyCACertificates) *DescribeCACertificatesResponseBody {
	s.CACertificates = v
	return s
}

func (s *DescribeCACertificatesResponseBody) SetRequestId(v string) *DescribeCACertificatesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCACertificatesResponseBodyCACertificates struct {
	CACertificate []*DescribeCACertificatesResponseBodyCACertificatesCACertificate `json:"CACertificate,omitempty" xml:"CACertificate,omitempty" type:"Repeated"`
}

func (s DescribeCACertificatesResponseBodyCACertificates) String() string {
	return tea.Prettify(s)
}

func (s DescribeCACertificatesResponseBodyCACertificates) GoString() string {
	return s.String()
}

func (s *DescribeCACertificatesResponseBodyCACertificates) SetCACertificate(v []*DescribeCACertificatesResponseBodyCACertificatesCACertificate) *DescribeCACertificatesResponseBodyCACertificates {
	s.CACertificate = v
	return s
}

type DescribeCACertificatesResponseBodyCACertificatesCACertificate struct {
	CACertificateId   *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	CACertificateName *string `json:"CACertificateName,omitempty" xml:"CACertificateName,omitempty"`
	CommonName        *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimeStamp   *int64  `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	ExpireTime        *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ExpireTimeStamp   *int64  `json:"ExpireTimeStamp,omitempty" xml:"ExpireTimeStamp,omitempty"`
	Fingerprint       *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId   *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeCACertificatesResponseBodyCACertificatesCACertificate) String() string {
	return tea.Prettify(s)
}

func (s DescribeCACertificatesResponseBodyCACertificatesCACertificate) GoString() string {
	return s.String()
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetCACertificateId(v string) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.CACertificateId = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetCACertificateName(v string) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.CACertificateName = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetCommonName(v string) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.CommonName = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetCreateTime(v string) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.CreateTime = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetCreateTimeStamp(v int64) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.CreateTimeStamp = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetExpireTime(v string) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.ExpireTime = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetExpireTimeStamp(v int64) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.ExpireTimeStamp = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetFingerprint(v string) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.Fingerprint = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetRegionId(v string) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.RegionId = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetResourceGroupId(v string) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.ResourceGroupId = &v
	return s
}

type DescribeCACertificatesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCACertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCACertificatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCACertificatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCACertificatesResponse) SetHeaders(v map[string]*string) *DescribeCACertificatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCACertificatesResponse) SetStatusCode(v int32) *DescribeCACertificatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCACertificatesResponse) SetBody(v *DescribeCACertificatesResponseBody) *DescribeCACertificatesResponse {
	s.Body = v
	return s
}

type DescribeDomainExtensionAttributeRequest struct {
	DomainExtensionId    *string `json:"DomainExtensionId,omitempty" xml:"DomainExtensionId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDomainExtensionAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainExtensionAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainExtensionAttributeRequest) SetDomainExtensionId(v string) *DescribeDomainExtensionAttributeRequest {
	s.DomainExtensionId = &v
	return s
}

func (s *DescribeDomainExtensionAttributeRequest) SetOwnerAccount(v string) *DescribeDomainExtensionAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDomainExtensionAttributeRequest) SetOwnerId(v int64) *DescribeDomainExtensionAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainExtensionAttributeRequest) SetRegionId(v string) *DescribeDomainExtensionAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDomainExtensionAttributeRequest) SetResourceOwnerAccount(v string) *DescribeDomainExtensionAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDomainExtensionAttributeRequest) SetResourceOwnerId(v int64) *DescribeDomainExtensionAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDomainExtensionAttributeResponseBody struct {
	Domain              *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainExtensionId   *string `json:"DomainExtensionId,omitempty" xml:"DomainExtensionId,omitempty"`
	ListenerPort        *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	LoadBalancerId      *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
}

func (s DescribeDomainExtensionAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainExtensionAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainExtensionAttributeResponseBody) SetDomain(v string) *DescribeDomainExtensionAttributeResponseBody {
	s.Domain = &v
	return s
}

func (s *DescribeDomainExtensionAttributeResponseBody) SetDomainExtensionId(v string) *DescribeDomainExtensionAttributeResponseBody {
	s.DomainExtensionId = &v
	return s
}

func (s *DescribeDomainExtensionAttributeResponseBody) SetListenerPort(v int32) *DescribeDomainExtensionAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *DescribeDomainExtensionAttributeResponseBody) SetLoadBalancerId(v string) *DescribeDomainExtensionAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeDomainExtensionAttributeResponseBody) SetRequestId(v string) *DescribeDomainExtensionAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainExtensionAttributeResponseBody) SetServerCertificateId(v string) *DescribeDomainExtensionAttributeResponseBody {
	s.ServerCertificateId = &v
	return s
}

type DescribeDomainExtensionAttributeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainExtensionAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainExtensionAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainExtensionAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainExtensionAttributeResponse) SetHeaders(v map[string]*string) *DescribeDomainExtensionAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainExtensionAttributeResponse) SetStatusCode(v int32) *DescribeDomainExtensionAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainExtensionAttributeResponse) SetBody(v *DescribeDomainExtensionAttributeResponseBody) *DescribeDomainExtensionAttributeResponse {
	s.Body = v
	return s
}

type DescribeDomainExtensionsRequest struct {
	DomainExtensionId    *string `json:"DomainExtensionId,omitempty" xml:"DomainExtensionId,omitempty"`
	ListenerPort         *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDomainExtensionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainExtensionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainExtensionsRequest) SetDomainExtensionId(v string) *DescribeDomainExtensionsRequest {
	s.DomainExtensionId = &v
	return s
}

func (s *DescribeDomainExtensionsRequest) SetListenerPort(v int32) *DescribeDomainExtensionsRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeDomainExtensionsRequest) SetLoadBalancerId(v string) *DescribeDomainExtensionsRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeDomainExtensionsRequest) SetOwnerAccount(v string) *DescribeDomainExtensionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDomainExtensionsRequest) SetOwnerId(v int64) *DescribeDomainExtensionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainExtensionsRequest) SetRegionId(v string) *DescribeDomainExtensionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDomainExtensionsRequest) SetResourceOwnerAccount(v string) *DescribeDomainExtensionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDomainExtensionsRequest) SetResourceOwnerId(v int64) *DescribeDomainExtensionsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDomainExtensionsResponseBody struct {
	DomainExtensions *DescribeDomainExtensionsResponseBodyDomainExtensions `json:"DomainExtensions,omitempty" xml:"DomainExtensions,omitempty" type:"Struct"`
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainExtensionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainExtensionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainExtensionsResponseBody) SetDomainExtensions(v *DescribeDomainExtensionsResponseBodyDomainExtensions) *DescribeDomainExtensionsResponseBody {
	s.DomainExtensions = v
	return s
}

func (s *DescribeDomainExtensionsResponseBody) SetRequestId(v string) *DescribeDomainExtensionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDomainExtensionsResponseBodyDomainExtensions struct {
	DomainExtension []*DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension `json:"DomainExtension,omitempty" xml:"DomainExtension,omitempty" type:"Repeated"`
}

func (s DescribeDomainExtensionsResponseBodyDomainExtensions) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainExtensionsResponseBodyDomainExtensions) GoString() string {
	return s.String()
}

func (s *DescribeDomainExtensionsResponseBodyDomainExtensions) SetDomainExtension(v []*DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension) *DescribeDomainExtensionsResponseBodyDomainExtensions {
	s.DomainExtension = v
	return s
}

type DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension struct {
	Domain              *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainExtensionId   *string `json:"DomainExtensionId,omitempty" xml:"DomainExtensionId,omitempty"`
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
}

func (s DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension) GoString() string {
	return s.String()
}

func (s *DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension) SetDomain(v string) *DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension {
	s.Domain = &v
	return s
}

func (s *DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension) SetDomainExtensionId(v string) *DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension {
	s.DomainExtensionId = &v
	return s
}

func (s *DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension) SetServerCertificateId(v string) *DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension {
	s.ServerCertificateId = &v
	return s
}

type DescribeDomainExtensionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainExtensionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainExtensionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainExtensionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainExtensionsResponse) SetHeaders(v map[string]*string) *DescribeDomainExtensionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainExtensionsResponse) SetStatusCode(v int32) *DescribeDomainExtensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainExtensionsResponse) SetBody(v *DescribeDomainExtensionsResponseBody) *DescribeDomainExtensionsResponse {
	s.Body = v
	return s
}

type DescribeHealthStatusRequest struct {
	ListenerPort         *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerProtocol     *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeHealthStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusRequest) SetListenerPort(v int32) *DescribeHealthStatusRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeHealthStatusRequest) SetListenerProtocol(v string) *DescribeHealthStatusRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *DescribeHealthStatusRequest) SetLoadBalancerId(v string) *DescribeHealthStatusRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeHealthStatusRequest) SetOwnerAccount(v string) *DescribeHealthStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeHealthStatusRequest) SetOwnerId(v int64) *DescribeHealthStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHealthStatusRequest) SetRegionId(v string) *DescribeHealthStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHealthStatusRequest) SetResourceOwnerAccount(v string) *DescribeHealthStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHealthStatusRequest) SetResourceOwnerId(v int64) *DescribeHealthStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeHealthStatusResponseBody struct {
	BackendServers *DescribeHealthStatusResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHealthStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBody) SetBackendServers(v *DescribeHealthStatusResponseBodyBackendServers) *DescribeHealthStatusResponseBody {
	s.BackendServers = v
	return s
}

func (s *DescribeHealthStatusResponseBody) SetRequestId(v string) *DescribeHealthStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeHealthStatusResponseBodyBackendServers struct {
	BackendServer []*DescribeHealthStatusResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s DescribeHealthStatusResponseBodyBackendServers) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyBackendServers) SetBackendServer(v []*DescribeHealthStatusResponseBodyBackendServersBackendServer) *DescribeHealthStatusResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

type DescribeHealthStatusResponseBodyBackendServersBackendServer struct {
	ListenerPort       *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	Port               *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol           *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	ServerHealthStatus *string `json:"ServerHealthStatus,omitempty" xml:"ServerHealthStatus,omitempty"`
	ServerId           *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	ServerIp           *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
}

func (s DescribeHealthStatusResponseBodyBackendServersBackendServer) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyBackendServersBackendServer) SetListenerPort(v int32) *DescribeHealthStatusResponseBodyBackendServersBackendServer {
	s.ListenerPort = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyBackendServersBackendServer) SetPort(v int32) *DescribeHealthStatusResponseBodyBackendServersBackendServer {
	s.Port = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyBackendServersBackendServer) SetProtocol(v string) *DescribeHealthStatusResponseBodyBackendServersBackendServer {
	s.Protocol = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyBackendServersBackendServer) SetServerHealthStatus(v string) *DescribeHealthStatusResponseBodyBackendServersBackendServer {
	s.ServerHealthStatus = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyBackendServersBackendServer) SetServerId(v string) *DescribeHealthStatusResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyBackendServersBackendServer) SetServerIp(v string) *DescribeHealthStatusResponseBodyBackendServersBackendServer {
	s.ServerIp = &v
	return s
}

type DescribeHealthStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeHealthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHealthStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponse) SetHeaders(v map[string]*string) *DescribeHealthStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeHealthStatusResponse) SetStatusCode(v int32) *DescribeHealthStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHealthStatusResponse) SetBody(v *DescribeHealthStatusResponseBody) *DescribeHealthStatusResponse {
	s.Body = v
	return s
}

type DescribeListenerAccessControlAttributeRequest struct {
	ListenerPort         *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerProtocol     *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeListenerAccessControlAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeListenerAccessControlAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeListenerAccessControlAttributeRequest) SetListenerPort(v int32) *DescribeListenerAccessControlAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeRequest) SetListenerProtocol(v string) *DescribeListenerAccessControlAttributeRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeRequest) SetLoadBalancerId(v string) *DescribeListenerAccessControlAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeRequest) SetOwnerAccount(v string) *DescribeListenerAccessControlAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeRequest) SetOwnerId(v int64) *DescribeListenerAccessControlAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeRequest) SetRegionId(v string) *DescribeListenerAccessControlAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeRequest) SetResourceOwnerAccount(v string) *DescribeListenerAccessControlAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeRequest) SetResourceOwnerId(v int64) *DescribeListenerAccessControlAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeListenerAccessControlAttributeResponseBody struct {
	AccessControlStatus *string `json:"AccessControlStatus,omitempty" xml:"AccessControlStatus,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceItems         *string `json:"SourceItems,omitempty" xml:"SourceItems,omitempty"`
}

func (s DescribeListenerAccessControlAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeListenerAccessControlAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeListenerAccessControlAttributeResponseBody) SetAccessControlStatus(v string) *DescribeListenerAccessControlAttributeResponseBody {
	s.AccessControlStatus = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeResponseBody) SetRequestId(v string) *DescribeListenerAccessControlAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeResponseBody) SetSourceItems(v string) *DescribeListenerAccessControlAttributeResponseBody {
	s.SourceItems = &v
	return s
}

type DescribeListenerAccessControlAttributeResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeListenerAccessControlAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeListenerAccessControlAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeListenerAccessControlAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeListenerAccessControlAttributeResponse) SetHeaders(v map[string]*string) *DescribeListenerAccessControlAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeListenerAccessControlAttributeResponse) SetStatusCode(v int32) *DescribeListenerAccessControlAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeResponse) SetBody(v *DescribeListenerAccessControlAttributeResponseBody) *DescribeListenerAccessControlAttributeResponse {
	s.Body = v
	return s
}

type DescribeLoadBalancerAttributeRequest struct {
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLoadBalancerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeRequest) SetLoadBalancerId(v string) *DescribeLoadBalancerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeRequest) SetOwnerAccount(v string) *DescribeLoadBalancerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerAttributeRequest) SetOwnerId(v int64) *DescribeLoadBalancerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeRequest) SetRegionId(v string) *DescribeLoadBalancerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeRequest) SetResourceOwnerAccount(v string) *DescribeLoadBalancerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerAttributeRequest) SetResourceOwnerId(v int64) *DescribeLoadBalancerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeLoadBalancerAttributeResponseBody struct {
	Address                      *string                                                            `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressIPVersion             *string                                                            `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	AddressType                  *string                                                            `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	AutoReleaseTime              *int64                                                             `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	BackendServers               *DescribeLoadBalancerAttributeResponseBodyBackendServers           `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	Bandwidth                    *int32                                                             `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	CreateTime                   *string                                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimeStamp              *int64                                                             `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	DeleteProtection             *string                                                            `json:"DeleteProtection,omitempty" xml:"DeleteProtection,omitempty"`
	EndTime                      *string                                                            `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimeStamp                 *int64                                                             `json:"EndTimeStamp,omitempty" xml:"EndTimeStamp,omitempty"`
	InstanceChargeType           *string                                                            `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InternetChargeType           *string                                                            `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	ListenerPorts                *DescribeLoadBalancerAttributeResponseBodyListenerPorts            `json:"ListenerPorts,omitempty" xml:"ListenerPorts,omitempty" type:"Struct"`
	ListenerPortsAndProtocal     *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal `json:"ListenerPortsAndProtocal,omitempty" xml:"ListenerPortsAndProtocal,omitempty" type:"Struct"`
	ListenerPortsAndProtocol     *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol `json:"ListenerPortsAndProtocol,omitempty" xml:"ListenerPortsAndProtocol,omitempty" type:"Struct"`
	LoadBalancerId               *string                                                            `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LoadBalancerName             *string                                                            `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	LoadBalancerSpec             *string                                                            `json:"LoadBalancerSpec,omitempty" xml:"LoadBalancerSpec,omitempty"`
	LoadBalancerStatus           *string                                                            `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	MasterZoneId                 *string                                                            `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	ModificationProtectionReason *string                                                            `json:"ModificationProtectionReason,omitempty" xml:"ModificationProtectionReason,omitempty"`
	ModificationProtectionStatus *string                                                            `json:"ModificationProtectionStatus,omitempty" xml:"ModificationProtectionStatus,omitempty"`
	NetworkType                  *string                                                            `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PayType                      *string                                                            `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId                     *string                                                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionIdAlias                *string                                                            `json:"RegionIdAlias,omitempty" xml:"RegionIdAlias,omitempty"`
	RenewalCycUnit               *string                                                            `json:"RenewalCycUnit,omitempty" xml:"RenewalCycUnit,omitempty"`
	RenewalDuration              *int32                                                             `json:"RenewalDuration,omitempty" xml:"RenewalDuration,omitempty"`
	RenewalStatus                *string                                                            `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	RequestId                    *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId              *string                                                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SlaveZoneId                  *string                                                            `json:"SlaveZoneId,omitempty" xml:"SlaveZoneId,omitempty"`
	VSwitchId                    *string                                                            `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                        *string                                                            `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeLoadBalancerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetAddress(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.Address = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetAddressIPVersion(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetAddressType(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.AddressType = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetAutoReleaseTime(v int64) *DescribeLoadBalancerAttributeResponseBody {
	s.AutoReleaseTime = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetBackendServers(v *DescribeLoadBalancerAttributeResponseBodyBackendServers) *DescribeLoadBalancerAttributeResponseBody {
	s.BackendServers = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetBandwidth(v int32) *DescribeLoadBalancerAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetCreateTime(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetCreateTimeStamp(v int64) *DescribeLoadBalancerAttributeResponseBody {
	s.CreateTimeStamp = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetDeleteProtection(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.DeleteProtection = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetEndTime(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetEndTimeStamp(v int64) *DescribeLoadBalancerAttributeResponseBody {
	s.EndTimeStamp = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetInstanceChargeType(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetInternetChargeType(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetListenerPorts(v *DescribeLoadBalancerAttributeResponseBodyListenerPorts) *DescribeLoadBalancerAttributeResponseBody {
	s.ListenerPorts = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetListenerPortsAndProtocal(v *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal) *DescribeLoadBalancerAttributeResponseBody {
	s.ListenerPortsAndProtocal = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetListenerPortsAndProtocol(v *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol) *DescribeLoadBalancerAttributeResponseBody {
	s.ListenerPortsAndProtocol = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetLoadBalancerId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetLoadBalancerName(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.LoadBalancerName = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetLoadBalancerSpec(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.LoadBalancerSpec = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetLoadBalancerStatus(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.LoadBalancerStatus = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetMasterZoneId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.MasterZoneId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetModificationProtectionReason(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.ModificationProtectionReason = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetModificationProtectionStatus(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.ModificationProtectionStatus = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetNetworkType(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.NetworkType = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetPayType(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetRegionId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetRegionIdAlias(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.RegionIdAlias = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetRenewalCycUnit(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.RenewalCycUnit = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetRenewalDuration(v int32) *DescribeLoadBalancerAttributeResponseBody {
	s.RenewalDuration = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetRenewalStatus(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.RenewalStatus = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetRequestId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetResourceGroupId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetSlaveZoneId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.SlaveZoneId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetVSwitchId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetVpcId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.VpcId = &v
	return s
}

type DescribeLoadBalancerAttributeResponseBodyBackendServers struct {
	BackendServer []*DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerAttributeResponseBodyBackendServers) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServers) SetBackendServer(v []*DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) *DescribeLoadBalancerAttributeResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

type DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ServerId    *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight      *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) SetDescription(v string) *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) SetServerId(v string) *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) SetType(v string) *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer {
	s.Type = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) SetWeight(v int32) *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

type DescribeLoadBalancerAttributeResponseBodyListenerPorts struct {
	ListenerPort []*int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPorts) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPorts) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPorts) SetListenerPort(v []*int32) *DescribeLoadBalancerAttributeResponseBodyListenerPorts {
	s.ListenerPort = v
	return s
}

type DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal struct {
	ListenerPortAndProtocal []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal `json:"ListenerPortAndProtocal,omitempty" xml:"ListenerPortAndProtocal,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal) SetListenerPortAndProtocal(v []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal {
	s.ListenerPortAndProtocal = v
	return s
}

type DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal struct {
	ListenerPort     *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerProtocal *string `json:"ListenerProtocal,omitempty" xml:"ListenerProtocal,omitempty"`
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) SetListenerPort(v int32) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) SetListenerProtocal(v string) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal {
	s.ListenerProtocal = &v
	return s
}

type DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol struct {
	ListenerPortAndProtocol []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol `json:"ListenerPortAndProtocol,omitempty" xml:"ListenerPortAndProtocol,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol) SetListenerPortAndProtocol(v []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol {
	s.ListenerPortAndProtocol = v
	return s
}

type DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol struct {
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ForwardPort      *int32  `json:"ForwardPort,omitempty" xml:"ForwardPort,omitempty"`
	ListenerForward  *string `json:"ListenerForward,omitempty" xml:"ListenerForward,omitempty"`
	ListenerPort     *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) SetDescription(v string) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) SetForwardPort(v int32) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol {
	s.ForwardPort = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) SetListenerForward(v string) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol {
	s.ListenerForward = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) SetListenerPort(v int32) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) SetListenerProtocol(v string) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol {
	s.ListenerProtocol = &v
	return s
}

type DescribeLoadBalancerAttributeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLoadBalancerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLoadBalancerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponse) SetHeaders(v map[string]*string) *DescribeLoadBalancerAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponse) SetStatusCode(v int32) *DescribeLoadBalancerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponse) SetBody(v *DescribeLoadBalancerAttributeResponseBody) *DescribeLoadBalancerAttributeResponse {
	s.Body = v
	return s
}

type DescribeLoadBalancerHTTPListenerAttributeRequest struct {
	ListenerPort         *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLoadBalancerHTTPListenerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerHTTPListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) SetListenerPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) SetLoadBalancerId(v string) *DescribeLoadBalancerHTTPListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) SetOwnerAccount(v string) *DescribeLoadBalancerHTTPListenerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) SetOwnerId(v int64) *DescribeLoadBalancerHTTPListenerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) SetRegionId(v string) *DescribeLoadBalancerHTTPListenerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) SetResourceOwnerAccount(v string) *DescribeLoadBalancerHTTPListenerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) SetResourceOwnerId(v int64) *DescribeLoadBalancerHTTPListenerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeLoadBalancerHTTPListenerAttributeResponseBody struct {
	AclId                  *string                                                     `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclStatus              *string                                                     `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	AclType                *string                                                     `json:"AclType,omitempty" xml:"AclType,omitempty"`
	BackendServerPort      *int32                                                      `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	Bandwidth              *int32                                                      `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Cookie                 *string                                                     `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	CookieTimeout          *int32                                                      `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	Description            *string                                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	ForwardPort            *int32                                                      `json:"ForwardPort,omitempty" xml:"ForwardPort,omitempty"`
	Gzip                   *string                                                     `json:"Gzip,omitempty" xml:"Gzip,omitempty"`
	HealthCheck            *string                                                     `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckConnectPort *int32                                                      `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckDomain      *string                                                     `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	HealthCheckHttpCode    *string                                                     `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckInterval    *int32                                                      `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckMethod      *string                                                     `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	HealthCheckTimeout     *int32                                                      `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	HealthCheckURI         *string                                                     `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	HealthyThreshold       *int32                                                      `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	IdleTimeout            *int32                                                      `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	ListenerForward        *string                                                     `json:"ListenerForward,omitempty" xml:"ListenerForward,omitempty"`
	ListenerPort           *int32                                                      `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	RequestId              *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestTimeout         *int32                                                      `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	Rules                  *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	Scheduler              *string                                                     `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	SecurityStatus         *string                                                     `json:"SecurityStatus,omitempty" xml:"SecurityStatus,omitempty"`
	Status                 *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	StickySession          *string                                                     `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	StickySessionType      *string                                                     `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	UnhealthyThreshold     *int32                                                      `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	VServerGroupId         *string                                                     `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	XForwardedFor          *string                                                     `json:"XForwardedFor,omitempty" xml:"XForwardedFor,omitempty"`
	XForwardedFor_SLBID    *string                                                     `json:"XForwardedFor_SLBID,omitempty" xml:"XForwardedFor_SLBID,omitempty"`
	XForwardedFor_SLBIP    *string                                                     `json:"XForwardedFor_SLBIP,omitempty" xml:"XForwardedFor_SLBIP,omitempty"`
	XForwardedFor_proto    *string                                                     `json:"XForwardedFor_proto,omitempty" xml:"XForwardedFor_proto,omitempty"`
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetAclId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.AclId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetAclStatus(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.AclStatus = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetAclType(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.AclType = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetBackendServerPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.BackendServerPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetBandwidth(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetCookie(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Cookie = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetCookieTimeout(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.CookieTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetDescription(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetForwardPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.ForwardPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetGzip(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Gzip = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthCheck(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthCheckDomain(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthCheckDomain = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthCheckHttpCode(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthCheckInterval(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthCheckMethod(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthCheckMethod = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthCheckTimeout(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthCheckTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthCheckURI(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthCheckURI = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetHealthyThreshold(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetIdleTimeout(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.IdleTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetListenerForward(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.ListenerForward = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetListenerPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetRequestId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetRequestTimeout(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.RequestTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetRules(v *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRules) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetScheduler(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Scheduler = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetSecurityStatus(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.SecurityStatus = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetStatus(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetStickySession(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.StickySession = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetStickySessionType(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.StickySessionType = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetVServerGroupId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetXForwardedFor(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.XForwardedFor = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetXForwardedFor_SLBID(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.XForwardedFor_SLBID = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetXForwardedFor_SLBIP(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.XForwardedFor_SLBIP = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBody) SetXForwardedFor_proto(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBody {
	s.XForwardedFor_proto = &v
	return s
}

type DescribeLoadBalancerHTTPListenerAttributeResponseBodyRules struct {
	Rule []*DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRules) SetRule(v []*DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRules {
	s.Rule = v
	return s
}

type DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule struct {
	Domain         *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	RuleId         *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName       *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Url            *string `json:"Url,omitempty" xml:"Url,omitempty"`
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) SetDomain(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule {
	s.Domain = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) SetRuleId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule {
	s.RuleId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) SetRuleName(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule {
	s.RuleName = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) SetUrl(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule {
	s.Url = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule) SetVServerGroupId(v string) *DescribeLoadBalancerHTTPListenerAttributeResponseBodyRulesRule {
	s.VServerGroupId = &v
	return s
}

type DescribeLoadBalancerHTTPListenerAttributeResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLoadBalancerHTTPListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerHTTPListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponse) SetHeaders(v map[string]*string) *DescribeLoadBalancerHTTPListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponse) SetStatusCode(v int32) *DescribeLoadBalancerHTTPListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeResponse) SetBody(v *DescribeLoadBalancerHTTPListenerAttributeResponseBody) *DescribeLoadBalancerHTTPListenerAttributeResponse {
	s.Body = v
	return s
}

type DescribeLoadBalancerHTTPSListenerAttributeRequest struct {
	ListenerPort         *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLoadBalancerHTTPSListenerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerHTTPSListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetListenerPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetLoadBalancerId(v string) *DescribeLoadBalancerHTTPSListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetOwnerAccount(v string) *DescribeLoadBalancerHTTPSListenerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetOwnerId(v int64) *DescribeLoadBalancerHTTPSListenerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetRegionId(v string) *DescribeLoadBalancerHTTPSListenerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetResourceOwnerAccount(v string) *DescribeLoadBalancerHTTPSListenerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetResourceOwnerId(v int64) *DescribeLoadBalancerHTTPSListenerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeLoadBalancerHTTPSListenerAttributeResponseBody struct {
	AclId                                *string                                                                 `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclStatus                            *string                                                                 `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	AclType                              *string                                                                 `json:"AclType,omitempty" xml:"AclType,omitempty"`
	BackendServerPort                    *int32                                                                  `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	Bandwidth                            *int32                                                                  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	CACertificateId                      *string                                                                 `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	Cookie                               *string                                                                 `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	CookieTimeout                        *int32                                                                  `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	Description                          *string                                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainExtensions                     *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensions `json:"DomainExtensions,omitempty" xml:"DomainExtensions,omitempty" type:"Struct"`
	EnableHttp2                          *string                                                                 `json:"EnableHttp2,omitempty" xml:"EnableHttp2,omitempty"`
	Gzip                                 *string                                                                 `json:"Gzip,omitempty" xml:"Gzip,omitempty"`
	HealthCheck                          *string                                                                 `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckConnectPort               *int32                                                                  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckDomain                    *string                                                                 `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	HealthCheckHttpCode                  *string                                                                 `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckInterval                  *int32                                                                  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckMethod                    *string                                                                 `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	HealthCheckTimeout                   *int32                                                                  `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	HealthCheckURI                       *string                                                                 `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	HealthyThreshold                     *int32                                                                  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	IdleTimeout                          *int32                                                                  `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	ListenerPort                         *int32                                                                  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	RequestId                            *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestTimeout                       *int32                                                                  `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	Rules                                *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRules            `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	Scheduler                            *string                                                                 `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	SecurityStatus                       *string                                                                 `json:"SecurityStatus,omitempty" xml:"SecurityStatus,omitempty"`
	ServerCertificateId                  *string                                                                 `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
	Status                               *string                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	StickySession                        *string                                                                 `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	StickySessionType                    *string                                                                 `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	TLSCipherPolicy                      *string                                                                 `json:"TLSCipherPolicy,omitempty" xml:"TLSCipherPolicy,omitempty"`
	UnhealthyThreshold                   *int32                                                                  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	VServerGroupId                       *string                                                                 `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	XForwardedFor                        *string                                                                 `json:"XForwardedFor,omitempty" xml:"XForwardedFor,omitempty"`
	XForwardedFor_ClientCertClientVerify *string                                                                 `json:"XForwardedFor_ClientCertClientVerify,omitempty" xml:"XForwardedFor_ClientCertClientVerify,omitempty"`
	XForwardedFor_ClientCertFingerprint  *string                                                                 `json:"XForwardedFor_ClientCertFingerprint,omitempty" xml:"XForwardedFor_ClientCertFingerprint,omitempty"`
	XForwardedFor_ClientCertIssuerDN     *string                                                                 `json:"XForwardedFor_ClientCertIssuerDN,omitempty" xml:"XForwardedFor_ClientCertIssuerDN,omitempty"`
	XForwardedFor_ClientCertSubjectDN    *string                                                                 `json:"XForwardedFor_ClientCertSubjectDN,omitempty" xml:"XForwardedFor_ClientCertSubjectDN,omitempty"`
	XForwardedFor_ClientSrcPort          *string                                                                 `json:"XForwardedFor_ClientSrcPort,omitempty" xml:"XForwardedFor_ClientSrcPort,omitempty"`
	XForwardedFor_SLBID                  *string                                                                 `json:"XForwardedFor_SLBID,omitempty" xml:"XForwardedFor_SLBID,omitempty"`
	XForwardedFor_SLBIP                  *string                                                                 `json:"XForwardedFor_SLBIP,omitempty" xml:"XForwardedFor_SLBIP,omitempty"`
	XForwardedFor_SLBPORT                *string                                                                 `json:"XForwardedFor_SLBPORT,omitempty" xml:"XForwardedFor_SLBPORT,omitempty"`
	XForwardedFor_proto                  *string                                                                 `json:"XForwardedFor_proto,omitempty" xml:"XForwardedFor_proto,omitempty"`
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetAclId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.AclId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetAclStatus(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.AclStatus = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetAclType(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.AclType = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetBackendServerPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.BackendServerPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetBandwidth(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetCACertificateId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.CACertificateId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetCookie(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Cookie = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetCookieTimeout(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.CookieTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetDescription(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetDomainExtensions(v *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensions) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.DomainExtensions = v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetEnableHttp2(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.EnableHttp2 = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetGzip(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Gzip = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheck(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckDomain(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckDomain = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckHttpCode(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckInterval(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckMethod(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckMethod = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckTimeout(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthCheckURI(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthCheckURI = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetHealthyThreshold(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetIdleTimeout(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.IdleTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetListenerPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetRequestId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetRequestTimeout(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.RequestTimeout = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetRules(v *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRules) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetScheduler(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Scheduler = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetSecurityStatus(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.SecurityStatus = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetServerCertificateId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.ServerCertificateId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetStatus(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetStickySession(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.StickySession = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetStickySessionType(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.StickySessionType = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetTLSCipherPolicy(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.TLSCipherPolicy = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetVServerGroupId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetXForwardedFor(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.XForwardedFor = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetXForwardedFor_ClientCertClientVerify(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.XForwardedFor_ClientCertClientVerify = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetXForwardedFor_ClientCertFingerprint(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.XForwardedFor_ClientCertFingerprint = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetXForwardedFor_ClientCertIssuerDN(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.XForwardedFor_ClientCertIssuerDN = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetXForwardedFor_ClientCertSubjectDN(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.XForwardedFor_ClientCertSubjectDN = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetXForwardedFor_ClientSrcPort(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.XForwardedFor_ClientSrcPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetXForwardedFor_SLBID(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.XForwardedFor_SLBID = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetXForwardedFor_SLBIP(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.XForwardedFor_SLBIP = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetXForwardedFor_SLBPORT(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.XForwardedFor_SLBPORT = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) SetXForwardedFor_proto(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBody {
	s.XForwardedFor_proto = &v
	return s
}

type DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensions struct {
	DomainExtension []*DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension `json:"DomainExtension,omitempty" xml:"DomainExtension,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensions) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensions) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensions) SetDomainExtension(v []*DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensions {
	s.DomainExtension = v
	return s
}

type DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension struct {
	Domain              *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainExtensionId   *string `json:"DomainExtensionId,omitempty" xml:"DomainExtensionId,omitempty"`
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension) SetDomain(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension {
	s.Domain = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension) SetDomainExtensionId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension {
	s.DomainExtensionId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension) SetServerCertificateId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyDomainExtensionsDomainExtension {
	s.ServerCertificateId = &v
	return s
}

type DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRules struct {
	Rule []*DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRules) SetRule(v []*DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRules {
	s.Rule = v
	return s
}

type DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule struct {
	Domain         *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	RuleId         *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName       *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Url            *string `json:"Url,omitempty" xml:"Url,omitempty"`
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) SetDomain(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule {
	s.Domain = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) SetRuleId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule {
	s.RuleId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) SetRuleName(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule {
	s.RuleName = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) SetUrl(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule {
	s.Url = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule) SetVServerGroupId(v string) *DescribeLoadBalancerHTTPSListenerAttributeResponseBodyRulesRule {
	s.VServerGroupId = &v
	return s
}

type DescribeLoadBalancerHTTPSListenerAttributeResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLoadBalancerHTTPSListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerHTTPSListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponse) SetHeaders(v map[string]*string) *DescribeLoadBalancerHTTPSListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponse) SetStatusCode(v int32) *DescribeLoadBalancerHTTPSListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeResponse) SetBody(v *DescribeLoadBalancerHTTPSListenerAttributeResponseBody) *DescribeLoadBalancerHTTPSListenerAttributeResponse {
	s.Body = v
	return s
}

type DescribeLoadBalancerListenersRequest struct {
	ListenerProtocol     *string   `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	LoadBalancerId       []*string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty" type:"Repeated"`
	MaxResults           *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken            *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLoadBalancerListenersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerListenersRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersRequest) SetListenerProtocol(v string) *DescribeLoadBalancerListenersRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetLoadBalancerId(v []*string) *DescribeLoadBalancerListenersRequest {
	s.LoadBalancerId = v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetMaxResults(v int32) *DescribeLoadBalancerListenersRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetNextToken(v string) *DescribeLoadBalancerListenersRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetOwnerAccount(v string) *DescribeLoadBalancerListenersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetOwnerId(v int64) *DescribeLoadBalancerListenersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetRegionId(v string) *DescribeLoadBalancerListenersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetResourceOwnerAccount(v string) *DescribeLoadBalancerListenersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetResourceOwnerId(v int64) *DescribeLoadBalancerListenersRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeLoadBalancerListenersResponseBody struct {
	Listeners  []*DescribeLoadBalancerListenersResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	MaxResults *int32                                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLoadBalancerListenersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerListenersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersResponseBody) SetListeners(v []*DescribeLoadBalancerListenersResponseBodyListeners) *DescribeLoadBalancerListenersResponseBody {
	s.Listeners = v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBody) SetMaxResults(v int32) *DescribeLoadBalancerListenersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBody) SetNextToken(v string) *DescribeLoadBalancerListenersResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBody) SetRequestId(v string) *DescribeLoadBalancerListenersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBody) SetTotalCount(v int32) *DescribeLoadBalancerListenersResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeLoadBalancerListenersResponseBodyListeners struct {
	AclId               *string                                                                `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclStatus           *string                                                                `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	AclType             *string                                                                `json:"AclType,omitempty" xml:"AclType,omitempty"`
	BackendServerPort   *int32                                                                 `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	Bandwidth           *int32                                                                 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Description         *string                                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	HTTPListenerConfig  *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig  `json:"HTTPListenerConfig,omitempty" xml:"HTTPListenerConfig,omitempty" type:"Struct"`
	HTTPSListenerConfig *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig `json:"HTTPSListenerConfig,omitempty" xml:"HTTPSListenerConfig,omitempty" type:"Struct"`
	ListenerPort        *int32                                                                 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerProtocol    *string                                                                `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	LoadBalancerId      *string                                                                `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	Scheduler           *string                                                                `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	Status              *string                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	TCPListenerConfig   *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig   `json:"TCPListenerConfig,omitempty" xml:"TCPListenerConfig,omitempty" type:"Struct"`
	UDPListenerConfig   *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig   `json:"UDPListenerConfig,omitempty" xml:"UDPListenerConfig,omitempty" type:"Struct"`
	VServerGroupId      *string                                                                `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DescribeLoadBalancerListenersResponseBodyListeners) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerListenersResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetAclId(v string) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.AclId = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetAclStatus(v string) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.AclStatus = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetAclType(v string) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.AclType = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetBackendServerPort(v int32) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.BackendServerPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetBandwidth(v int32) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetDescription(v string) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetHTTPListenerConfig(v *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.HTTPListenerConfig = v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetHTTPSListenerConfig(v *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.HTTPSListenerConfig = v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetListenerPort(v int32) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetListenerProtocol(v string) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.ListenerProtocol = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetLoadBalancerId(v string) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetScheduler(v string) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.Scheduler = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetStatus(v string) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.Status = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetTCPListenerConfig(v *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.TCPListenerConfig = v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetUDPListenerConfig(v *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.UDPListenerConfig = v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListeners) SetVServerGroupId(v string) *DescribeLoadBalancerListenersResponseBodyListeners {
	s.VServerGroupId = &v
	return s
}

type DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig struct {
	Cookie                      *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	CookieTimeout               *int32  `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	ForwardPort                 *int32  `json:"ForwardPort,omitempty" xml:"ForwardPort,omitempty"`
	Gzip                        *string `json:"Gzip,omitempty" xml:"Gzip,omitempty"`
	HealthCheck                 *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckConnectPort      *int32  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckDomain           *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	HealthCheckHttpCode         *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckHttpVersion      *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	HealthCheckInterval         *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckMethod           *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	HealthCheckTimeout          *int32  `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	HealthCheckType             *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	HealthCheckURI              *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	HealthyThreshold            *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	IdleTimeout                 *int32  `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	ListenerForward             *string `json:"ListenerForward,omitempty" xml:"ListenerForward,omitempty"`
	RequestTimeout              *int32  `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	StickySession               *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	StickySessionType           *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	UnhealthyThreshold          *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	XForwardedFor               *string `json:"XForwardedFor,omitempty" xml:"XForwardedFor,omitempty"`
	XForwardedFor_ClientSrcPort *string `json:"XForwardedFor_ClientSrcPort,omitempty" xml:"XForwardedFor_ClientSrcPort,omitempty"`
	XForwardedFor_SLBID         *string `json:"XForwardedFor_SLBID,omitempty" xml:"XForwardedFor_SLBID,omitempty"`
	XForwardedFor_SLBIP         *string `json:"XForwardedFor_SLBIP,omitempty" xml:"XForwardedFor_SLBIP,omitempty"`
	XForwardedFor_SLBPORT       *string `json:"XForwardedFor_SLBPORT,omitempty" xml:"XForwardedFor_SLBPORT,omitempty"`
	XForwardedFor_proto         *string `json:"XForwardedFor_proto,omitempty" xml:"XForwardedFor_proto,omitempty"`
}

func (s DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetCookie(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.Cookie = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetCookieTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.CookieTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetForwardPort(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.ForwardPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetGzip(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.Gzip = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthCheck(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthCheckDomain(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthCheckDomain = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthCheckHttpCode(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthCheckHttpVersion(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthCheckInterval(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthCheckMethod(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthCheckMethod = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthCheckTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthCheckTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthCheckType(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthCheckType = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthCheckURI(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthCheckURI = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetHealthyThreshold(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetIdleTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.IdleTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetListenerForward(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.ListenerForward = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetRequestTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.RequestTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetStickySession(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.StickySession = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetStickySessionType(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.StickySessionType = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetXForwardedFor(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.XForwardedFor = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetXForwardedFor_ClientSrcPort(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.XForwardedFor_ClientSrcPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetXForwardedFor_SLBID(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.XForwardedFor_SLBID = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetXForwardedFor_SLBIP(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.XForwardedFor_SLBIP = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetXForwardedFor_SLBPORT(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.XForwardedFor_SLBPORT = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig) SetXForwardedFor_proto(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPListenerConfig {
	s.XForwardedFor_proto = &v
	return s
}

type DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig struct {
	CACertificateId                      *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	Cookie                               *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	CookieTimeout                        *int32  `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	EnableHttp2                          *string `json:"EnableHttp2,omitempty" xml:"EnableHttp2,omitempty"`
	Gzip                                 *string `json:"Gzip,omitempty" xml:"Gzip,omitempty"`
	HealthCheck                          *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckConnectPort               *int32  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckDomain                    *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	HealthCheckHttpCode                  *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckHttpVersion               *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	HealthCheckInterval                  *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckMethod                    *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	HealthCheckTimeout                   *int32  `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	HealthCheckType                      *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	HealthCheckURI                       *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	HealthyThreshold                     *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	IdleTimeout                          *int32  `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	RequestTimeout                       *int32  `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	ServerCertificateId                  *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
	StickySession                        *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	StickySessionType                    *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	TLSCipherPolicy                      *string `json:"TLSCipherPolicy,omitempty" xml:"TLSCipherPolicy,omitempty"`
	UnhealthyThreshold                   *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	XForwardedFor                        *string `json:"XForwardedFor,omitempty" xml:"XForwardedFor,omitempty"`
	XForwardedFor_ClientCertClientVerify *string `json:"XForwardedFor_ClientCertClientVerify,omitempty" xml:"XForwardedFor_ClientCertClientVerify,omitempty"`
	XForwardedFor_ClientCertFingerprint  *string `json:"XForwardedFor_ClientCertFingerprint,omitempty" xml:"XForwardedFor_ClientCertFingerprint,omitempty"`
	XForwardedFor_ClientCertIssuerDN     *string `json:"XForwardedFor_ClientCertIssuerDN,omitempty" xml:"XForwardedFor_ClientCertIssuerDN,omitempty"`
	XForwardedFor_ClientCertSubjectDN    *string `json:"XForwardedFor_ClientCertSubjectDN,omitempty" xml:"XForwardedFor_ClientCertSubjectDN,omitempty"`
	XForwardedFor_ClientSrcPort          *string `json:"XForwardedFor_ClientSrcPort,omitempty" xml:"XForwardedFor_ClientSrcPort,omitempty"`
	XForwardedFor_SLBID                  *string `json:"XForwardedFor_SLBID,omitempty" xml:"XForwardedFor_SLBID,omitempty"`
	XForwardedFor_SLBIP                  *string `json:"XForwardedFor_SLBIP,omitempty" xml:"XForwardedFor_SLBIP,omitempty"`
	XForwardedFor_SLBPORT                *string `json:"XForwardedFor_SLBPORT,omitempty" xml:"XForwardedFor_SLBPORT,omitempty"`
	XForwardedFor_proto                  *string `json:"XForwardedFor_proto,omitempty" xml:"XForwardedFor_proto,omitempty"`
}

func (s DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetCACertificateId(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.CACertificateId = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetCookie(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.Cookie = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetCookieTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.CookieTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetEnableHttp2(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.EnableHttp2 = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetGzip(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.Gzip = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthCheck(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthCheckDomain(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthCheckDomain = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthCheckHttpCode(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthCheckHttpVersion(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthCheckInterval(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthCheckMethod(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthCheckMethod = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthCheckTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthCheckTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthCheckType(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthCheckType = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthCheckURI(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthCheckURI = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetHealthyThreshold(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetIdleTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.IdleTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetRequestTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.RequestTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetServerCertificateId(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.ServerCertificateId = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetStickySession(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.StickySession = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetStickySessionType(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.StickySessionType = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetTLSCipherPolicy(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.TLSCipherPolicy = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetXForwardedFor(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.XForwardedFor = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetXForwardedFor_ClientCertClientVerify(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.XForwardedFor_ClientCertClientVerify = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetXForwardedFor_ClientCertFingerprint(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.XForwardedFor_ClientCertFingerprint = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetXForwardedFor_ClientCertIssuerDN(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.XForwardedFor_ClientCertIssuerDN = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetXForwardedFor_ClientCertSubjectDN(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.XForwardedFor_ClientCertSubjectDN = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetXForwardedFor_ClientSrcPort(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.XForwardedFor_ClientSrcPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetXForwardedFor_SLBID(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.XForwardedFor_SLBID = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetXForwardedFor_SLBIP(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.XForwardedFor_SLBIP = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetXForwardedFor_SLBPORT(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.XForwardedFor_SLBPORT = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig) SetXForwardedFor_proto(v string) *DescribeLoadBalancerListenersResponseBodyListenersHTTPSListenerConfig {
	s.XForwardedFor_proto = &v
	return s
}

type DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig struct {
	ConnectionDrain           *string `json:"ConnectionDrain,omitempty" xml:"ConnectionDrain,omitempty"`
	ConnectionDrainTimeout    *int32  `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
	EstablishedTimeout        *int32  `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
	HealthCheck               *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckConnectPort    *int32  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckConnectTimeout *int32  `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	HealthCheckDomain         *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	HealthCheckHttpCode       *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckInterval       *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckMethod         *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	HealthCheckType           *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	HealthCheckURI            *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	HealthyThreshold          *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	MasterSlaveServerGroupId  *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	PersistenceTimeout        *int32  `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	ProxyProtocolV2Enabled    *string `json:"ProxyProtocolV2Enabled,omitempty" xml:"ProxyProtocolV2Enabled,omitempty"`
	UnhealthyThreshold        *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetConnectionDrain(v string) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.ConnectionDrain = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetConnectionDrainTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetEstablishedTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.EstablishedTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetHealthCheck(v string) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetHealthCheckConnectTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetHealthCheckDomain(v string) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.HealthCheckDomain = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetHealthCheckHttpCode(v string) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetHealthCheckInterval(v int32) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetHealthCheckMethod(v string) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.HealthCheckMethod = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetHealthCheckType(v string) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.HealthCheckType = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetHealthCheckURI(v string) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.HealthCheckURI = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetHealthyThreshold(v int32) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetMasterSlaveServerGroupId(v string) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetPersistenceTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.PersistenceTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetProxyProtocolV2Enabled(v string) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.ProxyProtocolV2Enabled = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerListenersResponseBodyListenersTCPListenerConfig {
	s.UnhealthyThreshold = &v
	return s
}

type DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig struct {
	ConnectionDrain           *string `json:"ConnectionDrain,omitempty" xml:"ConnectionDrain,omitempty"`
	ConnectionDrainTimeout    *int32  `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
	HealthCheck               *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckConnectPort    *int32  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckConnectTimeout *int32  `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	HealthCheckExp            *string `json:"HealthCheckExp,omitempty" xml:"HealthCheckExp,omitempty"`
	HealthCheckInterval       *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckReq            *string `json:"HealthCheckReq,omitempty" xml:"HealthCheckReq,omitempty"`
	HealthyThreshold          *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	MasterSlaveServerGroupId  *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	ProxyProtocolV2Enabled    *string `json:"ProxyProtocolV2Enabled,omitempty" xml:"ProxyProtocolV2Enabled,omitempty"`
	UnhealthyThreshold        *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetConnectionDrain(v string) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.ConnectionDrain = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetConnectionDrainTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetHealthCheck(v string) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetHealthCheckConnectTimeout(v int32) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetHealthCheckExp(v string) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.HealthCheckExp = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetHealthCheckInterval(v int32) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetHealthCheckReq(v string) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.HealthCheckReq = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetHealthyThreshold(v int32) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetMasterSlaveServerGroupId(v string) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetProxyProtocolV2Enabled(v string) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.ProxyProtocolV2Enabled = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerListenersResponseBodyListenersUDPListenerConfig {
	s.UnhealthyThreshold = &v
	return s
}

type DescribeLoadBalancerListenersResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLoadBalancerListenersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLoadBalancerListenersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerListenersResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersResponse) SetHeaders(v map[string]*string) *DescribeLoadBalancerListenersResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoadBalancerListenersResponse) SetStatusCode(v int32) *DescribeLoadBalancerListenersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoadBalancerListenersResponse) SetBody(v *DescribeLoadBalancerListenersResponseBody) *DescribeLoadBalancerListenersResponse {
	s.Body = v
	return s
}

type DescribeLoadBalancerTCPListenerAttributeRequest struct {
	ListenerPort         *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLoadBalancerTCPListenerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerTCPListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetListenerPort(v int32) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetLoadBalancerId(v string) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetOwnerAccount(v string) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetOwnerId(v int64) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetRegionId(v string) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetResourceOwnerAccount(v string) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetResourceOwnerId(v int64) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeLoadBalancerTCPListenerAttributeResponseBody struct {
	AclId                     *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclStatus                 *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	AclType                   *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	BackendServerPort         *int32  `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	Bandwidth                 *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ConnectionDrain           *string `json:"ConnectionDrain,omitempty" xml:"ConnectionDrain,omitempty"`
	ConnectionDrainTimeout    *int32  `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EstablishedTimeout        *int32  `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
	HealthCheck               *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckConnectPort    *int32  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckConnectTimeout *int32  `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	HealthCheckDomain         *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	HealthCheckHttpCode       *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckInterval       *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckMethod         *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	HealthCheckType           *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	HealthCheckURI            *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	HealthyThreshold          *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	ListenerPort              *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	MasterSlaveServerGroupId  *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	PersistenceTimeout        *int32  `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	ProxyProtocolV2Enabled    *bool   `json:"ProxyProtocolV2Enabled,omitempty" xml:"ProxyProtocolV2Enabled,omitempty"`
	RequestId                 *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Scheduler                 *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SynProxy                  *string `json:"SynProxy,omitempty" xml:"SynProxy,omitempty"`
	UnhealthyThreshold        *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	VServerGroupId            *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DescribeLoadBalancerTCPListenerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerTCPListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetAclId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.AclId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetAclStatus(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.AclStatus = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetAclType(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.AclType = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetBackendServerPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.BackendServerPort = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetBandwidth(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetConnectionDrain(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.ConnectionDrain = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetConnectionDrainTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetDescription(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetEstablishedTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.EstablishedTimeout = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheck(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckConnectTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckDomain(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckDomain = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckHttpCode(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckInterval(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckMethod(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckMethod = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckType(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckType = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckURI(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckURI = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthyThreshold(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetListenerPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetMasterSlaveServerGroupId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetPersistenceTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.PersistenceTimeout = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetProxyProtocolV2Enabled(v bool) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.ProxyProtocolV2Enabled = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetRequestId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetScheduler(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.Scheduler = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetStatus(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetSynProxy(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.SynProxy = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetVServerGroupId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.VServerGroupId = &v
	return s
}

type DescribeLoadBalancerTCPListenerAttributeResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLoadBalancerTCPListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLoadBalancerTCPListenerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerTCPListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponse) SetHeaders(v map[string]*string) *DescribeLoadBalancerTCPListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponse) SetStatusCode(v int32) *DescribeLoadBalancerTCPListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponse) SetBody(v *DescribeLoadBalancerTCPListenerAttributeResponseBody) *DescribeLoadBalancerTCPListenerAttributeResponse {
	s.Body = v
	return s
}

type DescribeLoadBalancerUDPListenerAttributeRequest struct {
	ListenerPort         *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLoadBalancerUDPListenerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerUDPListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) SetListenerPort(v int32) *DescribeLoadBalancerUDPListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) SetLoadBalancerId(v string) *DescribeLoadBalancerUDPListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) SetOwnerAccount(v string) *DescribeLoadBalancerUDPListenerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) SetOwnerId(v int64) *DescribeLoadBalancerUDPListenerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) SetRegionId(v string) *DescribeLoadBalancerUDPListenerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) SetResourceOwnerAccount(v string) *DescribeLoadBalancerUDPListenerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) SetResourceOwnerId(v int64) *DescribeLoadBalancerUDPListenerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeLoadBalancerUDPListenerAttributeResponseBody struct {
	AclId                     *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclStatus                 *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	AclType                   *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	BackendServerPort         *int32  `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	Bandwidth                 *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HealthCheck               *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckConnectPort    *int32  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckConnectTimeout *int32  `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	HealthCheckExp            *string `json:"HealthCheckExp,omitempty" xml:"HealthCheckExp,omitempty"`
	HealthCheckInterval       *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckReq            *string `json:"HealthCheckReq,omitempty" xml:"HealthCheckReq,omitempty"`
	HealthyThreshold          *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	ListenerPort              *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	MasterSlaveServerGroupId  *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	ProxyProtocolV2Enabled    *bool   `json:"ProxyProtocolV2Enabled,omitempty" xml:"ProxyProtocolV2Enabled,omitempty"`
	RequestId                 *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Scheduler                 *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UnhealthyThreshold        *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	VServerGroupId            *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DescribeLoadBalancerUDPListenerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerUDPListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetAclId(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.AclId = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetAclStatus(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.AclStatus = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetAclType(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.AclType = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetBackendServerPort(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.BackendServerPort = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetBandwidth(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetDescription(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthCheck(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthCheckConnectTimeout(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthCheckExp(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthCheckExp = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthCheckInterval(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthCheckReq(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthCheckReq = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthyThreshold(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetListenerPort(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetMasterSlaveServerGroupId(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetProxyProtocolV2Enabled(v bool) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.ProxyProtocolV2Enabled = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetRequestId(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetScheduler(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.Scheduler = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetStatus(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetVServerGroupId(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.VServerGroupId = &v
	return s
}

type DescribeLoadBalancerUDPListenerAttributeResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLoadBalancerUDPListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLoadBalancerUDPListenerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancerUDPListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponse) SetHeaders(v map[string]*string) *DescribeLoadBalancerUDPListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponse) SetStatusCode(v int32) *DescribeLoadBalancerUDPListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponse) SetBody(v *DescribeLoadBalancerUDPListenerAttributeResponseBody) *DescribeLoadBalancerUDPListenerAttributeResponse {
	s.Body = v
	return s
}

type DescribeLoadBalancersRequest struct {
	Address               *string `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressIPVersion      *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	AddressType           *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	InternetChargeType    *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	LoadBalancerId        *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LoadBalancerName      *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	LoadBalancerStatus    *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	MasterZoneId          *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	NetworkType           *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber            *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PayType               *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId       *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ServerId              *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	ServerIntranetAddress *string `json:"ServerIntranetAddress,omitempty" xml:"ServerIntranetAddress,omitempty"`
	SlaveZoneId           *string `json:"SlaveZoneId,omitempty" xml:"SlaveZoneId,omitempty"`
	Tags                  *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	VSwitchId             *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                 *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeLoadBalancersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancersRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancersRequest) SetAddress(v string) *DescribeLoadBalancersRequest {
	s.Address = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetAddressIPVersion(v string) *DescribeLoadBalancersRequest {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetAddressType(v string) *DescribeLoadBalancersRequest {
	s.AddressType = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetInternetChargeType(v string) *DescribeLoadBalancersRequest {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetLoadBalancerId(v string) *DescribeLoadBalancersRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetLoadBalancerName(v string) *DescribeLoadBalancersRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetLoadBalancerStatus(v string) *DescribeLoadBalancersRequest {
	s.LoadBalancerStatus = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetMasterZoneId(v string) *DescribeLoadBalancersRequest {
	s.MasterZoneId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetNetworkType(v string) *DescribeLoadBalancersRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetOwnerAccount(v string) *DescribeLoadBalancersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetOwnerId(v int64) *DescribeLoadBalancersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetPageNumber(v int32) *DescribeLoadBalancersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetPageSize(v int32) *DescribeLoadBalancersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetPayType(v string) *DescribeLoadBalancersRequest {
	s.PayType = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetRegionId(v string) *DescribeLoadBalancersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetResourceGroupId(v string) *DescribeLoadBalancersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetResourceOwnerAccount(v string) *DescribeLoadBalancersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetResourceOwnerId(v int64) *DescribeLoadBalancersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetServerId(v string) *DescribeLoadBalancersRequest {
	s.ServerId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetServerIntranetAddress(v string) *DescribeLoadBalancersRequest {
	s.ServerIntranetAddress = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetSlaveZoneId(v string) *DescribeLoadBalancersRequest {
	s.SlaveZoneId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetTags(v string) *DescribeLoadBalancersRequest {
	s.Tags = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetVSwitchId(v string) *DescribeLoadBalancersRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetVpcId(v string) *DescribeLoadBalancersRequest {
	s.VpcId = &v
	return s
}

type DescribeLoadBalancersResponseBody struct {
	LoadBalancers *DescribeLoadBalancersResponseBodyLoadBalancers `json:"LoadBalancers,omitempty" xml:"LoadBalancers,omitempty" type:"Struct"`
	PageNumber    *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLoadBalancersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancersResponseBody) SetLoadBalancers(v *DescribeLoadBalancersResponseBodyLoadBalancers) *DescribeLoadBalancersResponseBody {
	s.LoadBalancers = v
	return s
}

func (s *DescribeLoadBalancersResponseBody) SetPageNumber(v int32) *DescribeLoadBalancersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLoadBalancersResponseBody) SetPageSize(v int32) *DescribeLoadBalancersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLoadBalancersResponseBody) SetRequestId(v string) *DescribeLoadBalancersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBody) SetTotalCount(v int32) *DescribeLoadBalancersResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeLoadBalancersResponseBodyLoadBalancers struct {
	LoadBalancer []*DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer `json:"LoadBalancer,omitempty" xml:"LoadBalancer,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancersResponseBodyLoadBalancers) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancersResponseBodyLoadBalancers) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancers) SetLoadBalancer(v []*DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) *DescribeLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancer = v
	return s
}

type DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer struct {
	Address                      *string                                                         `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressIPVersion             *string                                                         `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	AddressType                  *string                                                         `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	Bandwidth                    *int32                                                          `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	CreateTime                   *string                                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimeStamp              *int64                                                          `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	DeleteProtection             *string                                                         `json:"DeleteProtection,omitempty" xml:"DeleteProtection,omitempty"`
	InstanceChargeType           *string                                                         `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InternetChargeType           *string                                                         `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetChargeTypeAlias      *string                                                         `json:"InternetChargeTypeAlias,omitempty" xml:"InternetChargeTypeAlias,omitempty"`
	LoadBalancerId               *string                                                         `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LoadBalancerName             *string                                                         `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	LoadBalancerSpec             *string                                                         `json:"LoadBalancerSpec,omitempty" xml:"LoadBalancerSpec,omitempty"`
	LoadBalancerStatus           *string                                                         `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	MasterZoneId                 *string                                                         `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	ModificationProtectionReason *string                                                         `json:"ModificationProtectionReason,omitempty" xml:"ModificationProtectionReason,omitempty"`
	ModificationProtectionStatus *string                                                         `json:"ModificationProtectionStatus,omitempty" xml:"ModificationProtectionStatus,omitempty"`
	NetworkType                  *string                                                         `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PayType                      *string                                                         `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId                     *string                                                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionIdAlias                *string                                                         `json:"RegionIdAlias,omitempty" xml:"RegionIdAlias,omitempty"`
	ResourceGroupId              *string                                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SlaveZoneId                  *string                                                         `json:"SlaveZoneId,omitempty" xml:"SlaveZoneId,omitempty"`
	Tags                         *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VSwitchId                    *string                                                         `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                        *string                                                         `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetAddress(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.Address = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetAddressIPVersion(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetAddressType(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.AddressType = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetBandwidth(v int32) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetCreateTime(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.CreateTime = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetCreateTimeStamp(v int64) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.CreateTimeStamp = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetDeleteProtection(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.DeleteProtection = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetInstanceChargeType(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetInternetChargeType(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetInternetChargeTypeAlias(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.InternetChargeTypeAlias = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetLoadBalancerId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetLoadBalancerName(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.LoadBalancerName = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetLoadBalancerSpec(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.LoadBalancerSpec = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetLoadBalancerStatus(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.LoadBalancerStatus = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetMasterZoneId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.MasterZoneId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetModificationProtectionReason(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.ModificationProtectionReason = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetModificationProtectionStatus(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.ModificationProtectionStatus = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetNetworkType(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.NetworkType = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetPayType(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.PayType = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetRegionId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetRegionIdAlias(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.RegionIdAlias = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetResourceGroupId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetSlaveZoneId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.SlaveZoneId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetTags(v *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTags) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.Tags = v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetVSwitchId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.VSwitchId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetVpcId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.VpcId = &v
	return s
}

type DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTags struct {
	Tag []*DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTags) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTags) SetTag(v []*DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTags {
	s.Tag = v
	return s
}

type DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag) SetTagKey(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag) SetTagValue(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancerTagsTag {
	s.TagValue = &v
	return s
}

type DescribeLoadBalancersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLoadBalancersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLoadBalancersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadBalancersResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancersResponse) SetHeaders(v map[string]*string) *DescribeLoadBalancersResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoadBalancersResponse) SetStatusCode(v int32) *DescribeLoadBalancersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoadBalancersResponse) SetBody(v *DescribeLoadBalancersResponseBody) *DescribeLoadBalancersResponse {
	s.Body = v
	return s
}

type DescribeMasterSlaveServerGroupAttributeRequest struct {
	MasterSlaveServerGroupId *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount     *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeMasterSlaveServerGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupAttributeRequest) SetMasterSlaveServerGroupId(v string) *DescribeMasterSlaveServerGroupAttributeRequest {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeRequest) SetOwnerAccount(v string) *DescribeMasterSlaveServerGroupAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeRequest) SetOwnerId(v int64) *DescribeMasterSlaveServerGroupAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeRequest) SetRegionId(v string) *DescribeMasterSlaveServerGroupAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeRequest) SetResourceOwnerAccount(v string) *DescribeMasterSlaveServerGroupAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeRequest) SetResourceOwnerId(v int64) *DescribeMasterSlaveServerGroupAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeMasterSlaveServerGroupAttributeResponseBody struct {
	CreateTime                 *string                                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LoadBalancerId             *string                                                                       `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	MasterSlaveBackendServers  *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServers `json:"MasterSlaveBackendServers,omitempty" xml:"MasterSlaveBackendServers,omitempty" type:"Struct"`
	MasterSlaveServerGroupId   *string                                                                       `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	MasterSlaveServerGroupName *string                                                                       `json:"MasterSlaveServerGroupName,omitempty" xml:"MasterSlaveServerGroupName,omitempty"`
	RequestId                  *string                                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMasterSlaveServerGroupAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) SetCreateTime(v string) *DescribeMasterSlaveServerGroupAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) SetLoadBalancerId(v string) *DescribeMasterSlaveServerGroupAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) SetMasterSlaveBackendServers(v *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServers) *DescribeMasterSlaveServerGroupAttributeResponseBody {
	s.MasterSlaveBackendServers = v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) SetMasterSlaveServerGroupId(v string) *DescribeMasterSlaveServerGroupAttributeResponseBody {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) SetMasterSlaveServerGroupName(v string) *DescribeMasterSlaveServerGroupAttributeResponseBody {
	s.MasterSlaveServerGroupName = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBody) SetRequestId(v string) *DescribeMasterSlaveServerGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServers struct {
	MasterSlaveBackendServer []*DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer `json:"MasterSlaveBackendServer,omitempty" xml:"MasterSlaveBackendServer,omitempty" type:"Repeated"`
}

func (s DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServers) String() string {
	return tea.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServers) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServers) SetMasterSlaveBackendServer(v []*DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServers {
	s.MasterSlaveBackendServer = v
	return s
}

type DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Port        *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ServerId    *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	ServerType  *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight      *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) String() string {
	return tea.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetDescription(v string) *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.Description = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetPort(v int32) *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.Port = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetServerId(v string) *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.ServerId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetServerType(v string) *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.ServerType = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetType(v string) *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.Type = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer) SetWeight(v int32) *DescribeMasterSlaveServerGroupAttributeResponseBodyMasterSlaveBackendServersMasterSlaveBackendServer {
	s.Weight = &v
	return s
}

type DescribeMasterSlaveServerGroupAttributeResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMasterSlaveServerGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMasterSlaveServerGroupAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupAttributeResponse) SetHeaders(v map[string]*string) *DescribeMasterSlaveServerGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponse) SetStatusCode(v int32) *DescribeMasterSlaveServerGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponse) SetBody(v *DescribeMasterSlaveServerGroupAttributeResponseBody) *DescribeMasterSlaveServerGroupAttributeResponse {
	s.Body = v
	return s
}

type DescribeMasterSlaveServerGroupsRequest struct {
	IncludeListener      *bool   `json:"IncludeListener,omitempty" xml:"IncludeListener,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeMasterSlaveServerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupsRequest) SetIncludeListener(v bool) *DescribeMasterSlaveServerGroupsRequest {
	s.IncludeListener = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsRequest) SetLoadBalancerId(v string) *DescribeMasterSlaveServerGroupsRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsRequest) SetOwnerAccount(v string) *DescribeMasterSlaveServerGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsRequest) SetOwnerId(v int64) *DescribeMasterSlaveServerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsRequest) SetRegionId(v string) *DescribeMasterSlaveServerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsRequest) SetResourceOwnerAccount(v string) *DescribeMasterSlaveServerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsRequest) SetResourceOwnerId(v int64) *DescribeMasterSlaveServerGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeMasterSlaveServerGroupsResponseBody struct {
	MasterSlaveServerGroups *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroups `json:"MasterSlaveServerGroups,omitempty" xml:"MasterSlaveServerGroups,omitempty" type:"Struct"`
	RequestId               *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMasterSlaveServerGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupsResponseBody) SetMasterSlaveServerGroups(v *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroups) *DescribeMasterSlaveServerGroupsResponseBody {
	s.MasterSlaveServerGroups = v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponseBody) SetRequestId(v string) *DescribeMasterSlaveServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroups struct {
	MasterSlaveServerGroup []*DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup `json:"MasterSlaveServerGroup,omitempty" xml:"MasterSlaveServerGroup,omitempty" type:"Repeated"`
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroups) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroups) SetMasterSlaveServerGroup(v []*DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroups {
	s.MasterSlaveServerGroup = v
	return s
}

type DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup struct {
	AssociatedObjects          *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjects `json:"AssociatedObjects,omitempty" xml:"AssociatedObjects,omitempty" type:"Struct"`
	CreateTime                 *string                                                                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MasterSlaveServerGroupId   *string                                                                                                    `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	MasterSlaveServerGroupName *string                                                                                                    `json:"MasterSlaveServerGroupName,omitempty" xml:"MasterSlaveServerGroupName,omitempty"`
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) SetAssociatedObjects(v *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjects) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup {
	s.AssociatedObjects = v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) SetCreateTime(v string) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup {
	s.CreateTime = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) SetMasterSlaveServerGroupId(v string) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup) SetMasterSlaveServerGroupName(v string) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroup {
	s.MasterSlaveServerGroupName = &v
	return s
}

type DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjects struct {
	Listeners *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Struct"`
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjects) String() string {
	return tea.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjects) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjects) SetListeners(v *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListeners) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjects {
	s.Listeners = v
	return s
}

type DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListeners struct {
	Listener []*DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener `json:"Listener,omitempty" xml:"Listener,omitempty" type:"Repeated"`
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListeners) String() string {
	return tea.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListeners) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListeners) SetListener(v []*DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListeners {
	s.Listener = v
	return s
}

type DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener struct {
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener) String() string {
	return tea.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener) SetPort(v int32) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener {
	s.Port = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener) SetProtocol(v string) *DescribeMasterSlaveServerGroupsResponseBodyMasterSlaveServerGroupsMasterSlaveServerGroupAssociatedObjectsListenersListener {
	s.Protocol = &v
	return s
}

type DescribeMasterSlaveServerGroupsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMasterSlaveServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMasterSlaveServerGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupsResponse) SetHeaders(v map[string]*string) *DescribeMasterSlaveServerGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponse) SetStatusCode(v int32) *DescribeMasterSlaveServerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponse) SetBody(v *DescribeMasterSlaveServerGroupsResponseBody) *DescribeMasterSlaveServerGroupsResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	AcceptLanguage       *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerAccount(v string) *DescribeRegionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerId(v int64) *DescribeRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerAccount(v string) *DescribeRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerId(v int64) *DescribeRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeRuleAttributeRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RuleId               *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeRuleAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleAttributeRequest) SetOwnerAccount(v string) *DescribeRuleAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRuleAttributeRequest) SetOwnerId(v int64) *DescribeRuleAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRuleAttributeRequest) SetRegionId(v string) *DescribeRuleAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRuleAttributeRequest) SetResourceOwnerAccount(v string) *DescribeRuleAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRuleAttributeRequest) SetResourceOwnerId(v int64) *DescribeRuleAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRuleAttributeRequest) SetRuleId(v string) *DescribeRuleAttributeRequest {
	s.RuleId = &v
	return s
}

type DescribeRuleAttributeResponseBody struct {
	Cookie                 *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	CookieTimeout          *int32  `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	Domain                 *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	HealthCheck            *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckConnectPort *int32  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckDomain      *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	HealthCheckHttpCode    *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckInterval    *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckTimeout     *int32  `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	HealthCheckURI         *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	HealthyThreshold       *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	ListenerPort           *string `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerSync           *string `json:"ListenerSync,omitempty" xml:"ListenerSync,omitempty"`
	LoadBalancerId         *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleId                 *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName               *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Scheduler              *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	StickySession          *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	StickySessionType      *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	UnhealthyThreshold     *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	Url                    *string `json:"Url,omitempty" xml:"Url,omitempty"`
	VServerGroupId         *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DescribeRuleAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleAttributeResponseBody) SetCookie(v string) *DescribeRuleAttributeResponseBody {
	s.Cookie = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetCookieTimeout(v int32) *DescribeRuleAttributeResponseBody {
	s.CookieTimeout = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetDomain(v string) *DescribeRuleAttributeResponseBody {
	s.Domain = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetHealthCheck(v string) *DescribeRuleAttributeResponseBody {
	s.HealthCheck = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetHealthCheckConnectPort(v int32) *DescribeRuleAttributeResponseBody {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetHealthCheckDomain(v string) *DescribeRuleAttributeResponseBody {
	s.HealthCheckDomain = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetHealthCheckHttpCode(v string) *DescribeRuleAttributeResponseBody {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetHealthCheckInterval(v int32) *DescribeRuleAttributeResponseBody {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetHealthCheckTimeout(v int32) *DescribeRuleAttributeResponseBody {
	s.HealthCheckTimeout = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetHealthCheckURI(v string) *DescribeRuleAttributeResponseBody {
	s.HealthCheckURI = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetHealthyThreshold(v int32) *DescribeRuleAttributeResponseBody {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetListenerPort(v string) *DescribeRuleAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetListenerSync(v string) *DescribeRuleAttributeResponseBody {
	s.ListenerSync = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetLoadBalancerId(v string) *DescribeRuleAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetRequestId(v string) *DescribeRuleAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetRuleId(v string) *DescribeRuleAttributeResponseBody {
	s.RuleId = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetRuleName(v string) *DescribeRuleAttributeResponseBody {
	s.RuleName = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetScheduler(v string) *DescribeRuleAttributeResponseBody {
	s.Scheduler = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetStickySession(v string) *DescribeRuleAttributeResponseBody {
	s.StickySession = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetStickySessionType(v string) *DescribeRuleAttributeResponseBody {
	s.StickySessionType = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetUnhealthyThreshold(v int32) *DescribeRuleAttributeResponseBody {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetUrl(v string) *DescribeRuleAttributeResponseBody {
	s.Url = &v
	return s
}

func (s *DescribeRuleAttributeResponseBody) SetVServerGroupId(v string) *DescribeRuleAttributeResponseBody {
	s.VServerGroupId = &v
	return s
}

type DescribeRuleAttributeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRuleAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleAttributeResponse) SetHeaders(v map[string]*string) *DescribeRuleAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleAttributeResponse) SetStatusCode(v int32) *DescribeRuleAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleAttributeResponse) SetBody(v *DescribeRuleAttributeResponseBody) *DescribeRuleAttributeResponse {
	s.Body = v
	return s
}

type DescribeRulesRequest struct {
	ListenerPort         *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerProtocol     *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRulesRequest) SetListenerPort(v int32) *DescribeRulesRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeRulesRequest) SetListenerProtocol(v string) *DescribeRulesRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *DescribeRulesRequest) SetLoadBalancerId(v string) *DescribeRulesRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeRulesRequest) SetOwnerAccount(v string) *DescribeRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRulesRequest) SetOwnerId(v int64) *DescribeRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRulesRequest) SetRegionId(v string) *DescribeRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRulesRequest) SetResourceOwnerAccount(v string) *DescribeRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRulesRequest) SetResourceOwnerId(v int64) *DescribeRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeRulesResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     *DescribeRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s DescribeRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponseBody) SetRequestId(v string) *DescribeRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRulesResponseBody) SetRules(v *DescribeRulesResponseBodyRules) *DescribeRulesResponseBody {
	s.Rules = v
	return s
}

type DescribeRulesResponseBodyRules struct {
	Rule []*DescribeRulesResponseBodyRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeRulesResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponseBodyRules) SetRule(v []*DescribeRulesResponseBodyRulesRule) *DescribeRulesResponseBodyRules {
	s.Rule = v
	return s
}

type DescribeRulesResponseBodyRulesRule struct {
	Cookie                 *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	CookieTimeout          *int32  `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	Domain                 *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	HealthCheck            *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckConnectPort *int32  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckDomain      *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	HealthCheckHttpCode    *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckInterval    *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckTimeout     *int32  `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	HealthCheckURI         *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	HealthyThreshold       *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	ListenerSync           *string `json:"ListenerSync,omitempty" xml:"ListenerSync,omitempty"`
	RuleId                 *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName               *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Scheduler              *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	StickySession          *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	StickySessionType      *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	UnhealthyThreshold     *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	Url                    *string `json:"Url,omitempty" xml:"Url,omitempty"`
	VServerGroupId         *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DescribeRulesResponseBodyRulesRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesResponseBodyRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponseBodyRulesRule) SetCookie(v string) *DescribeRulesResponseBodyRulesRule {
	s.Cookie = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetCookieTimeout(v int32) *DescribeRulesResponseBodyRulesRule {
	s.CookieTimeout = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetDomain(v string) *DescribeRulesResponseBodyRulesRule {
	s.Domain = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetHealthCheck(v string) *DescribeRulesResponseBodyRulesRule {
	s.HealthCheck = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetHealthCheckConnectPort(v int32) *DescribeRulesResponseBodyRulesRule {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetHealthCheckDomain(v string) *DescribeRulesResponseBodyRulesRule {
	s.HealthCheckDomain = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetHealthCheckHttpCode(v string) *DescribeRulesResponseBodyRulesRule {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetHealthCheckInterval(v int32) *DescribeRulesResponseBodyRulesRule {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetHealthCheckTimeout(v int32) *DescribeRulesResponseBodyRulesRule {
	s.HealthCheckTimeout = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetHealthCheckURI(v string) *DescribeRulesResponseBodyRulesRule {
	s.HealthCheckURI = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetHealthyThreshold(v int32) *DescribeRulesResponseBodyRulesRule {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetListenerSync(v string) *DescribeRulesResponseBodyRulesRule {
	s.ListenerSync = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetRuleId(v string) *DescribeRulesResponseBodyRulesRule {
	s.RuleId = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetRuleName(v string) *DescribeRulesResponseBodyRulesRule {
	s.RuleName = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetScheduler(v string) *DescribeRulesResponseBodyRulesRule {
	s.Scheduler = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetStickySession(v string) *DescribeRulesResponseBodyRulesRule {
	s.StickySession = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetStickySessionType(v string) *DescribeRulesResponseBodyRulesRule {
	s.StickySessionType = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetUnhealthyThreshold(v int32) *DescribeRulesResponseBodyRulesRule {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetUrl(v string) *DescribeRulesResponseBodyRulesRule {
	s.Url = &v
	return s
}

func (s *DescribeRulesResponseBodyRulesRule) SetVServerGroupId(v string) *DescribeRulesResponseBodyRulesRule {
	s.VServerGroupId = &v
	return s
}

type DescribeRulesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponse) SetHeaders(v map[string]*string) *DescribeRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRulesResponse) SetStatusCode(v int32) *DescribeRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRulesResponse) SetBody(v *DescribeRulesResponseBody) *DescribeRulesResponse {
	s.Body = v
	return s
}

type DescribeServerCertificatesRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ServerCertificateId  *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
}

func (s DescribeServerCertificatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServerCertificatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeServerCertificatesRequest) SetOwnerAccount(v string) *DescribeServerCertificatesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeServerCertificatesRequest) SetOwnerId(v int64) *DescribeServerCertificatesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeServerCertificatesRequest) SetRegionId(v string) *DescribeServerCertificatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeServerCertificatesRequest) SetResourceGroupId(v string) *DescribeServerCertificatesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeServerCertificatesRequest) SetResourceOwnerAccount(v string) *DescribeServerCertificatesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeServerCertificatesRequest) SetResourceOwnerId(v int64) *DescribeServerCertificatesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeServerCertificatesRequest) SetServerCertificateId(v string) *DescribeServerCertificatesRequest {
	s.ServerCertificateId = &v
	return s
}

type DescribeServerCertificatesResponseBody struct {
	RequestId          *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServerCertificates *DescribeServerCertificatesResponseBodyServerCertificates `json:"ServerCertificates,omitempty" xml:"ServerCertificates,omitempty" type:"Struct"`
}

func (s DescribeServerCertificatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServerCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServerCertificatesResponseBody) SetRequestId(v string) *DescribeServerCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServerCertificatesResponseBody) SetServerCertificates(v *DescribeServerCertificatesResponseBodyServerCertificates) *DescribeServerCertificatesResponseBody {
	s.ServerCertificates = v
	return s
}

type DescribeServerCertificatesResponseBodyServerCertificates struct {
	ServerCertificate []*DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate `json:"ServerCertificate,omitempty" xml:"ServerCertificate,omitempty" type:"Repeated"`
}

func (s DescribeServerCertificatesResponseBodyServerCertificates) String() string {
	return tea.Prettify(s)
}

func (s DescribeServerCertificatesResponseBodyServerCertificates) GoString() string {
	return s.String()
}

func (s *DescribeServerCertificatesResponseBodyServerCertificates) SetServerCertificate(v []*DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) *DescribeServerCertificatesResponseBodyServerCertificates {
	s.ServerCertificate = v
	return s
}

type DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate struct {
	AliCloudCertificateId   *string                                                                                           `json:"AliCloudCertificateId,omitempty" xml:"AliCloudCertificateId,omitempty"`
	AliCloudCertificateName *string                                                                                           `json:"AliCloudCertificateName,omitempty" xml:"AliCloudCertificateName,omitempty"`
	CommonName              *string                                                                                           `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	CreateTime              *string                                                                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimeStamp         *int64                                                                                            `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	ExpireTime              *string                                                                                           `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ExpireTimeStamp         *int64                                                                                            `json:"ExpireTimeStamp,omitempty" xml:"ExpireTimeStamp,omitempty"`
	Fingerprint             *string                                                                                           `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	IsAliCloudCertificate   *int32                                                                                            `json:"IsAliCloudCertificate,omitempty" xml:"IsAliCloudCertificate,omitempty"`
	RegionId                *string                                                                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId         *string                                                                                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServerCertificateId     *string                                                                                           `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
	ServerCertificateName   *string                                                                                           `json:"ServerCertificateName,omitempty" xml:"ServerCertificateName,omitempty"`
	SubjectAlternativeNames *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateSubjectAlternativeNames `json:"SubjectAlternativeNames,omitempty" xml:"SubjectAlternativeNames,omitempty" type:"Struct"`
}

func (s DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) String() string {
	return tea.Prettify(s)
}

func (s DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) GoString() string {
	return s.String()
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetAliCloudCertificateId(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.AliCloudCertificateId = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetAliCloudCertificateName(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.AliCloudCertificateName = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetCommonName(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.CommonName = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetCreateTime(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.CreateTime = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetCreateTimeStamp(v int64) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.CreateTimeStamp = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetExpireTime(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.ExpireTime = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetExpireTimeStamp(v int64) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.ExpireTimeStamp = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetFingerprint(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.Fingerprint = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetIsAliCloudCertificate(v int32) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.IsAliCloudCertificate = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetRegionId(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.RegionId = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetResourceGroupId(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetServerCertificateId(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.ServerCertificateId = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetServerCertificateName(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.ServerCertificateName = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetSubjectAlternativeNames(v *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateSubjectAlternativeNames) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.SubjectAlternativeNames = v
	return s
}

type DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateSubjectAlternativeNames struct {
	SubjectAlternativeName []*string `json:"SubjectAlternativeName,omitempty" xml:"SubjectAlternativeName,omitempty" type:"Repeated"`
}

func (s DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateSubjectAlternativeNames) String() string {
	return tea.Prettify(s)
}

func (s DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateSubjectAlternativeNames) GoString() string {
	return s.String()
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateSubjectAlternativeNames) SetSubjectAlternativeName(v []*string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateSubjectAlternativeNames {
	s.SubjectAlternativeName = v
	return s
}

type DescribeServerCertificatesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeServerCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServerCertificatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServerCertificatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeServerCertificatesResponse) SetHeaders(v map[string]*string) *DescribeServerCertificatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeServerCertificatesResponse) SetStatusCode(v int32) *DescribeServerCertificatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServerCertificatesResponse) SetBody(v *DescribeServerCertificatesResponseBody) *DescribeServerCertificatesResponse {
	s.Body = v
	return s
}

type DescribeTagsRequest struct {
	DistinctKey          *bool   `json:"DistinctKey,omitempty" xml:"DistinctKey,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s DescribeTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequest) SetDistinctKey(v bool) *DescribeTagsRequest {
	s.DistinctKey = &v
	return s
}

func (s *DescribeTagsRequest) SetLoadBalancerId(v string) *DescribeTagsRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeTagsRequest) SetOwnerAccount(v string) *DescribeTagsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTagsRequest) SetOwnerId(v int64) *DescribeTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTagsRequest) SetPageNumber(v int32) *DescribeTagsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagsRequest) SetPageSize(v int32) *DescribeTagsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTagsRequest) SetRegionId(v string) *DescribeTagsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceOwnerAccount(v string) *DescribeTagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceOwnerId(v int64) *DescribeTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTagsRequest) SetTags(v string) *DescribeTagsRequest {
	s.Tags = &v
	return s
}

type DescribeTagsResponseBody struct {
	PageNumber *int32                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagSets    *DescribeTagsResponseBodyTagSets `json:"TagSets,omitempty" xml:"TagSets,omitempty" type:"Struct"`
	TotalCount *int32                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBody) SetPageNumber(v int32) *DescribeTagsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagsResponseBody) SetPageSize(v int32) *DescribeTagsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTagsResponseBody) SetRequestId(v string) *DescribeTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsResponseBody) SetTagSets(v *DescribeTagsResponseBodyTagSets) *DescribeTagsResponseBody {
	s.TagSets = v
	return s
}

func (s *DescribeTagsResponseBody) SetTotalCount(v int32) *DescribeTagsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTagsResponseBodyTagSets struct {
	TagSet []*DescribeTagsResponseBodyTagSetsTagSet `json:"TagSet,omitempty" xml:"TagSet,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBodyTagSets) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBodyTagSets) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTagSets) SetTagSet(v []*DescribeTagsResponseBodyTagSetsTagSet) *DescribeTagsResponseBodyTagSets {
	s.TagSet = v
	return s
}

type DescribeTagsResponseBodyTagSetsTagSet struct {
	InstanceCount *int32  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	TagKey        *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue      *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeTagsResponseBodyTagSetsTagSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBodyTagSetsTagSet) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTagSetsTagSet) SetInstanceCount(v int32) *DescribeTagsResponseBodyTagSetsTagSet {
	s.InstanceCount = &v
	return s
}

func (s *DescribeTagsResponseBodyTagSetsTagSet) SetTagKey(v string) *DescribeTagsResponseBodyTagSetsTagSet {
	s.TagKey = &v
	return s
}

func (s *DescribeTagsResponseBodyTagSetsTagSet) SetTagValue(v string) *DescribeTagsResponseBodyTagSetsTagSet {
	s.TagValue = &v
	return s
}

type DescribeTagsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponse) SetHeaders(v map[string]*string) *DescribeTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagsResponse) SetStatusCode(v int32) *DescribeTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagsResponse) SetBody(v *DescribeTagsResponseBody) *DescribeTagsResponse {
	s.Body = v
	return s
}

type DescribeVServerGroupAttributeRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VServerGroupId       *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DescribeVServerGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVServerGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupAttributeRequest) SetOwnerAccount(v string) *DescribeVServerGroupAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVServerGroupAttributeRequest) SetOwnerId(v int64) *DescribeVServerGroupAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVServerGroupAttributeRequest) SetRegionId(v string) *DescribeVServerGroupAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVServerGroupAttributeRequest) SetResourceOwnerAccount(v string) *DescribeVServerGroupAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVServerGroupAttributeRequest) SetResourceOwnerId(v int64) *DescribeVServerGroupAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVServerGroupAttributeRequest) SetVServerGroupId(v string) *DescribeVServerGroupAttributeRequest {
	s.VServerGroupId = &v
	return s
}

type DescribeVServerGroupAttributeResponseBody struct {
	BackendServers   *DescribeVServerGroupAttributeResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	CreateTime       *string                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LoadBalancerId   *string                                                  `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	RequestId        *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VServerGroupId   *string                                                  `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	VServerGroupName *string                                                  `json:"VServerGroupName,omitempty" xml:"VServerGroupName,omitempty"`
}

func (s DescribeVServerGroupAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVServerGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupAttributeResponseBody) SetBackendServers(v *DescribeVServerGroupAttributeResponseBodyBackendServers) *DescribeVServerGroupAttributeResponseBody {
	s.BackendServers = v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBody) SetCreateTime(v string) *DescribeVServerGroupAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBody) SetLoadBalancerId(v string) *DescribeVServerGroupAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBody) SetRequestId(v string) *DescribeVServerGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBody) SetVServerGroupId(v string) *DescribeVServerGroupAttributeResponseBody {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBody) SetVServerGroupName(v string) *DescribeVServerGroupAttributeResponseBody {
	s.VServerGroupName = &v
	return s
}

type DescribeVServerGroupAttributeResponseBodyBackendServers struct {
	BackendServer []*DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s DescribeVServerGroupAttributeResponseBodyBackendServers) String() string {
	return tea.Prettify(s)
}

func (s DescribeVServerGroupAttributeResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServers) SetBackendServer(v []*DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) *DescribeVServerGroupAttributeResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

type DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Port        *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ServerId    *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	ServerIp    *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight      *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) String() string {
	return tea.Prettify(s)
}

func (s DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) SetDescription(v string) *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.Description = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) SetPort(v int32) *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.Port = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) SetServerId(v string) *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) SetServerIp(v string) *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.ServerIp = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) SetType(v string) *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.Type = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer) SetWeight(v int32) *DescribeVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

type DescribeVServerGroupAttributeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVServerGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVServerGroupAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVServerGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupAttributeResponse) SetHeaders(v map[string]*string) *DescribeVServerGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeVServerGroupAttributeResponse) SetStatusCode(v int32) *DescribeVServerGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponse) SetBody(v *DescribeVServerGroupAttributeResponseBody) *DescribeVServerGroupAttributeResponse {
	s.Body = v
	return s
}

type DescribeVServerGroupsRequest struct {
	IncludeListener      *bool   `json:"IncludeListener,omitempty" xml:"IncludeListener,omitempty"`
	IncludeRule          *bool   `json:"IncludeRule,omitempty" xml:"IncludeRule,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeVServerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsRequest) SetIncludeListener(v bool) *DescribeVServerGroupsRequest {
	s.IncludeListener = &v
	return s
}

func (s *DescribeVServerGroupsRequest) SetIncludeRule(v bool) *DescribeVServerGroupsRequest {
	s.IncludeRule = &v
	return s
}

func (s *DescribeVServerGroupsRequest) SetLoadBalancerId(v string) *DescribeVServerGroupsRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeVServerGroupsRequest) SetOwnerAccount(v string) *DescribeVServerGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVServerGroupsRequest) SetOwnerId(v int64) *DescribeVServerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVServerGroupsRequest) SetRegionId(v string) *DescribeVServerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVServerGroupsRequest) SetResourceOwnerAccount(v string) *DescribeVServerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVServerGroupsRequest) SetResourceOwnerId(v int64) *DescribeVServerGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeVServerGroupsResponseBody struct {
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VServerGroups *DescribeVServerGroupsResponseBodyVServerGroups `json:"VServerGroups,omitempty" xml:"VServerGroups,omitempty" type:"Struct"`
}

func (s DescribeVServerGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsResponseBody) SetRequestId(v string) *DescribeVServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVServerGroupsResponseBody) SetVServerGroups(v *DescribeVServerGroupsResponseBodyVServerGroups) *DescribeVServerGroupsResponseBody {
	s.VServerGroups = v
	return s
}

type DescribeVServerGroupsResponseBodyVServerGroups struct {
	VServerGroup []*DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup `json:"VServerGroup,omitempty" xml:"VServerGroup,omitempty" type:"Repeated"`
}

func (s DescribeVServerGroupsResponseBodyVServerGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeVServerGroupsResponseBodyVServerGroups) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsResponseBodyVServerGroups) SetVServerGroup(v []*DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) *DescribeVServerGroupsResponseBodyVServerGroups {
	s.VServerGroup = v
	return s
}

type DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup struct {
	AssociatedObjects *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects `json:"AssociatedObjects,omitempty" xml:"AssociatedObjects,omitempty" type:"Struct"`
	CreateTime        *string                                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ServerCount       *int64                                                                       `json:"ServerCount,omitempty" xml:"ServerCount,omitempty"`
	VServerGroupId    *string                                                                      `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	VServerGroupName  *string                                                                      `json:"VServerGroupName,omitempty" xml:"VServerGroupName,omitempty"`
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) SetAssociatedObjects(v *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup {
	s.AssociatedObjects = v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) SetCreateTime(v string) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup {
	s.CreateTime = &v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) SetServerCount(v int64) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup {
	s.ServerCount = &v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) SetVServerGroupId(v string) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) SetVServerGroupName(v string) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup {
	s.VServerGroupName = &v
	return s
}

type DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects struct {
	Listeners *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Struct"`
	Rules     *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRules     `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects) String() string {
	return tea.Prettify(s)
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects) SetListeners(v *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListeners) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects {
	s.Listeners = v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects) SetRules(v *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRules) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects {
	s.Rules = v
	return s
}

type DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListeners struct {
	Listener []*DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener `json:"Listener,omitempty" xml:"Listener,omitempty" type:"Repeated"`
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListeners) String() string {
	return tea.Prettify(s)
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListeners) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListeners) SetListener(v []*DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListeners {
	s.Listener = v
	return s
}

type DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener struct {
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener) String() string {
	return tea.Prettify(s)
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener) SetPort(v int32) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener {
	s.Port = &v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener) SetProtocol(v string) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener {
	s.Protocol = &v
	return s
}

type DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRules struct {
	Rule []*DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRules) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRules) SetRule(v []*DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRules {
	s.Rule = v
	return s
}

type DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule struct {
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	RuleId   *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule) SetDomain(v string) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule {
	s.Domain = &v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule) SetRuleId(v string) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule {
	s.RuleId = &v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule) SetRuleName(v string) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule {
	s.RuleName = &v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule) SetUrl(v string) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule {
	s.Url = &v
	return s
}

type DescribeVServerGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVServerGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVServerGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsResponse) SetHeaders(v map[string]*string) *DescribeVServerGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVServerGroupsResponse) SetStatusCode(v int32) *DescribeVServerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVServerGroupsResponse) SetBody(v *DescribeVServerGroupsResponseBody) *DescribeVServerGroupsResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetOwnerAccount(v string) *DescribeZonesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeZonesRequest) SetOwnerId(v int64) *DescribeZonesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeZonesRequest) SetResourceOwnerAccount(v string) *DescribeZonesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeZonesRequest) SetResourceOwnerId(v int64) *DescribeZonesRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeZonesResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Zones     *DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v *DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

type DescribeZonesResponseBodyZones struct {
	Zone []*DescribeZonesResponseBodyZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) SetZone(v []*DescribeZonesResponseBodyZonesZone) *DescribeZonesResponseBodyZones {
	s.Zone = v
	return s
}

type DescribeZonesResponseBodyZonesZone struct {
	LocalName  *string                                       `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	SlaveZones *DescribeZonesResponseBodyZonesZoneSlaveZones `json:"SlaveZones,omitempty" xml:"SlaveZones,omitempty" type:"Struct"`
	ZoneId     *string                                       `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZone) SetLocalName(v string) *DescribeZonesResponseBodyZonesZone {
	s.LocalName = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetSlaveZones(v *DescribeZonesResponseBodyZonesZoneSlaveZones) *DescribeZonesResponseBodyZonesZone {
	s.SlaveZones = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneId(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneId = &v
	return s
}

type DescribeZonesResponseBodyZonesZoneSlaveZones struct {
	SlaveZone []*DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone `json:"SlaveZone,omitempty" xml:"SlaveZone,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneSlaveZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneSlaveZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneSlaveZones) SetSlaveZone(v []*DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone) *DescribeZonesResponseBodyZonesZoneSlaveZones {
	s.SlaveZone = v
	return s
}

type DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone) SetLocalName(v string) *DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone {
	s.LocalName = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone) SetZoneId(v string) *DescribeZonesResponseBodyZonesZoneSlaveZonesSlaveZone {
	s.ZoneId = &v
	return s
}

type DescribeZonesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetHeaders(v map[string]*string) *DescribeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesResponse) SetStatusCode(v int32) *DescribeZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

type ListTLSCipherPoliciesRequest struct {
	IncludeListener      *bool   `json:"IncludeListener,omitempty" xml:"IncludeListener,omitempty"`
	MaxItems             *int32  `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TLSCipherPolicyId    *string `json:"TLSCipherPolicyId,omitempty" xml:"TLSCipherPolicyId,omitempty"`
}

func (s ListTLSCipherPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTLSCipherPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListTLSCipherPoliciesRequest) SetIncludeListener(v bool) *ListTLSCipherPoliciesRequest {
	s.IncludeListener = &v
	return s
}

func (s *ListTLSCipherPoliciesRequest) SetMaxItems(v int32) *ListTLSCipherPoliciesRequest {
	s.MaxItems = &v
	return s
}

func (s *ListTLSCipherPoliciesRequest) SetName(v string) *ListTLSCipherPoliciesRequest {
	s.Name = &v
	return s
}

func (s *ListTLSCipherPoliciesRequest) SetNextToken(v string) *ListTLSCipherPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTLSCipherPoliciesRequest) SetOwnerAccount(v string) *ListTLSCipherPoliciesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTLSCipherPoliciesRequest) SetOwnerId(v int64) *ListTLSCipherPoliciesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTLSCipherPoliciesRequest) SetRegionId(v string) *ListTLSCipherPoliciesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTLSCipherPoliciesRequest) SetResourceOwnerAccount(v string) *ListTLSCipherPoliciesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTLSCipherPoliciesRequest) SetResourceOwnerId(v int64) *ListTLSCipherPoliciesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTLSCipherPoliciesRequest) SetTLSCipherPolicyId(v string) *ListTLSCipherPoliciesRequest {
	s.TLSCipherPolicyId = &v
	return s
}

type ListTLSCipherPoliciesResponseBody struct {
	IsTruncated       *bool                                                 `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	NextToken         *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TLSCipherPolicies []*ListTLSCipherPoliciesResponseBodyTLSCipherPolicies `json:"TLSCipherPolicies,omitempty" xml:"TLSCipherPolicies,omitempty" type:"Repeated"`
	TotalCount        *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTLSCipherPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTLSCipherPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTLSCipherPoliciesResponseBody) SetIsTruncated(v bool) *ListTLSCipherPoliciesResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListTLSCipherPoliciesResponseBody) SetNextToken(v string) *ListTLSCipherPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTLSCipherPoliciesResponseBody) SetRequestId(v string) *ListTLSCipherPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTLSCipherPoliciesResponseBody) SetTLSCipherPolicies(v []*ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) *ListTLSCipherPoliciesResponseBody {
	s.TLSCipherPolicies = v
	return s
}

func (s *ListTLSCipherPoliciesResponseBody) SetTotalCount(v int32) *ListTLSCipherPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

type ListTLSCipherPoliciesResponseBodyTLSCipherPolicies struct {
	Ciphers         []*string                                                            `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	CreateTime      *int64                                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceId      *string                                                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name            *string                                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	RelateListeners []*ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners `json:"RelateListeners,omitempty" xml:"RelateListeners,omitempty" type:"Repeated"`
	Status          *string                                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	TLSVersions     []*string                                                            `json:"TLSVersions,omitempty" xml:"TLSVersions,omitempty" type:"Repeated"`
}

func (s ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) GoString() string {
	return s.String()
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) SetCiphers(v []*string) *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies {
	s.Ciphers = v
	return s
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) SetCreateTime(v int64) *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies {
	s.CreateTime = &v
	return s
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) SetInstanceId(v string) *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies {
	s.InstanceId = &v
	return s
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) SetName(v string) *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies {
	s.Name = &v
	return s
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) SetRelateListeners(v []*ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners) *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies {
	s.RelateListeners = v
	return s
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) SetStatus(v string) *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies {
	s.Status = &v
	return s
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) SetTLSVersions(v []*string) *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies {
	s.TLSVersions = v
	return s
}

type ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners struct {
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	Port           *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol       *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners) String() string {
	return tea.Prettify(s)
}

func (s ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners) GoString() string {
	return s.String()
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners) SetLoadBalancerId(v string) *ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners {
	s.LoadBalancerId = &v
	return s
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners) SetPort(v int32) *ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners {
	s.Port = &v
	return s
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners) SetProtocol(v string) *ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners {
	s.Protocol = &v
	return s
}

type ListTLSCipherPoliciesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTLSCipherPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTLSCipherPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTLSCipherPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListTLSCipherPoliciesResponse) SetHeaders(v map[string]*string) *ListTLSCipherPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListTLSCipherPoliciesResponse) SetStatusCode(v int32) *ListTLSCipherPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTLSCipherPoliciesResponse) SetBody(v *ListTLSCipherPoliciesResponseBody) *ListTLSCipherPoliciesResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	NextToken            *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string                       `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                        `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag                  []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerAccount(v string) *ListTagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerId(v int64) *ListTagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyLoadBalancerInstanceChargeTypeRequest struct {
	Bandwidth            *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	InstanceChargeType   *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InternetChargeType   *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LoadBalancerSpec     *string `json:"LoadBalancerSpec,omitempty" xml:"LoadBalancerSpec,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyLoadBalancerInstanceChargeTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLoadBalancerInstanceChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) SetBandwidth(v int32) *ModifyLoadBalancerInstanceChargeTypeRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) SetInstanceChargeType(v string) *ModifyLoadBalancerInstanceChargeTypeRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) SetInternetChargeType(v string) *ModifyLoadBalancerInstanceChargeTypeRequest {
	s.InternetChargeType = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) SetLoadBalancerId(v string) *ModifyLoadBalancerInstanceChargeTypeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) SetLoadBalancerSpec(v string) *ModifyLoadBalancerInstanceChargeTypeRequest {
	s.LoadBalancerSpec = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) SetOwnerAccount(v string) *ModifyLoadBalancerInstanceChargeTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) SetOwnerId(v int64) *ModifyLoadBalancerInstanceChargeTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) SetRegionId(v string) *ModifyLoadBalancerInstanceChargeTypeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) SetResourceOwnerAccount(v string) *ModifyLoadBalancerInstanceChargeTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeRequest) SetResourceOwnerId(v int64) *ModifyLoadBalancerInstanceChargeTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyLoadBalancerInstanceChargeTypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLoadBalancerInstanceChargeTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLoadBalancerInstanceChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerInstanceChargeTypeResponseBody) SetRequestId(v string) *ModifyLoadBalancerInstanceChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyLoadBalancerInstanceChargeTypeResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyLoadBalancerInstanceChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyLoadBalancerInstanceChargeTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLoadBalancerInstanceChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerInstanceChargeTypeResponse) SetHeaders(v map[string]*string) *ModifyLoadBalancerInstanceChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeResponse) SetStatusCode(v int32) *ModifyLoadBalancerInstanceChargeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeResponse) SetBody(v *ModifyLoadBalancerInstanceChargeTypeResponseBody) *ModifyLoadBalancerInstanceChargeTypeResponse {
	s.Body = v
	return s
}

type ModifyLoadBalancerInstanceSpecRequest struct {
	AutoPay              *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LoadBalancerSpec     *string `json:"LoadBalancerSpec,omitempty" xml:"LoadBalancerSpec,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyLoadBalancerInstanceSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLoadBalancerInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerInstanceSpecRequest) SetAutoPay(v bool) *ModifyLoadBalancerInstanceSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecRequest) SetLoadBalancerId(v string) *ModifyLoadBalancerInstanceSpecRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecRequest) SetLoadBalancerSpec(v string) *ModifyLoadBalancerInstanceSpecRequest {
	s.LoadBalancerSpec = &v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecRequest) SetOwnerAccount(v string) *ModifyLoadBalancerInstanceSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecRequest) SetOwnerId(v int64) *ModifyLoadBalancerInstanceSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecRequest) SetRegionId(v string) *ModifyLoadBalancerInstanceSpecRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecRequest) SetResourceOwnerAccount(v string) *ModifyLoadBalancerInstanceSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecRequest) SetResourceOwnerId(v int64) *ModifyLoadBalancerInstanceSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyLoadBalancerInstanceSpecResponseBody struct {
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLoadBalancerInstanceSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLoadBalancerInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerInstanceSpecResponseBody) SetOrderId(v int64) *ModifyLoadBalancerInstanceSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecResponseBody) SetRequestId(v string) *ModifyLoadBalancerInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

type ModifyLoadBalancerInstanceSpecResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyLoadBalancerInstanceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyLoadBalancerInstanceSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLoadBalancerInstanceSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerInstanceSpecResponse) SetHeaders(v map[string]*string) *ModifyLoadBalancerInstanceSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecResponse) SetStatusCode(v int32) *ModifyLoadBalancerInstanceSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecResponse) SetBody(v *ModifyLoadBalancerInstanceSpecResponseBody) *ModifyLoadBalancerInstanceSpecResponse {
	s.Body = v
	return s
}

type ModifyLoadBalancerInternetSpecRequest struct {
	AutoPay              *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	Bandwidth            *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	InternetChargeType   *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyLoadBalancerInternetSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLoadBalancerInternetSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerInternetSpecRequest) SetAutoPay(v bool) *ModifyLoadBalancerInternetSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecRequest) SetBandwidth(v int32) *ModifyLoadBalancerInternetSpecRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecRequest) SetInternetChargeType(v string) *ModifyLoadBalancerInternetSpecRequest {
	s.InternetChargeType = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecRequest) SetLoadBalancerId(v string) *ModifyLoadBalancerInternetSpecRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecRequest) SetOwnerAccount(v string) *ModifyLoadBalancerInternetSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecRequest) SetOwnerId(v int64) *ModifyLoadBalancerInternetSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecRequest) SetRegionId(v string) *ModifyLoadBalancerInternetSpecRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecRequest) SetResourceOwnerAccount(v string) *ModifyLoadBalancerInternetSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecRequest) SetResourceOwnerId(v int64) *ModifyLoadBalancerInternetSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyLoadBalancerInternetSpecResponseBody struct {
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLoadBalancerInternetSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLoadBalancerInternetSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerInternetSpecResponseBody) SetOrderId(v int64) *ModifyLoadBalancerInternetSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecResponseBody) SetRequestId(v string) *ModifyLoadBalancerInternetSpecResponseBody {
	s.RequestId = &v
	return s
}

type ModifyLoadBalancerInternetSpecResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyLoadBalancerInternetSpecResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyLoadBalancerInternetSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLoadBalancerInternetSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerInternetSpecResponse) SetHeaders(v map[string]*string) *ModifyLoadBalancerInternetSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyLoadBalancerInternetSpecResponse) SetStatusCode(v int32) *ModifyLoadBalancerInternetSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecResponse) SetBody(v *ModifyLoadBalancerInternetSpecResponseBody) *ModifyLoadBalancerInternetSpecResponse {
	s.Body = v
	return s
}

type ModifyLoadBalancerPayTypeRequest struct {
	AutoPay              *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	Duration             *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PricingCycle         *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyLoadBalancerPayTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLoadBalancerPayTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerPayTypeRequest) SetAutoPay(v bool) *ModifyLoadBalancerPayTypeRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeRequest) SetDuration(v int32) *ModifyLoadBalancerPayTypeRequest {
	s.Duration = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeRequest) SetLoadBalancerId(v string) *ModifyLoadBalancerPayTypeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeRequest) SetOwnerAccount(v string) *ModifyLoadBalancerPayTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeRequest) SetOwnerId(v int64) *ModifyLoadBalancerPayTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeRequest) SetPayType(v string) *ModifyLoadBalancerPayTypeRequest {
	s.PayType = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeRequest) SetPricingCycle(v string) *ModifyLoadBalancerPayTypeRequest {
	s.PricingCycle = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeRequest) SetRegionId(v string) *ModifyLoadBalancerPayTypeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeRequest) SetResourceOwnerAccount(v string) *ModifyLoadBalancerPayTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeRequest) SetResourceOwnerId(v int64) *ModifyLoadBalancerPayTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyLoadBalancerPayTypeResponseBody struct {
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLoadBalancerPayTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLoadBalancerPayTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerPayTypeResponseBody) SetOrderId(v int64) *ModifyLoadBalancerPayTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeResponseBody) SetRequestId(v string) *ModifyLoadBalancerPayTypeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyLoadBalancerPayTypeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyLoadBalancerPayTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyLoadBalancerPayTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLoadBalancerPayTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerPayTypeResponse) SetHeaders(v map[string]*string) *ModifyLoadBalancerPayTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyLoadBalancerPayTypeResponse) SetStatusCode(v int32) *ModifyLoadBalancerPayTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeResponse) SetBody(v *ModifyLoadBalancerPayTypeResponseBody) *ModifyLoadBalancerPayTypeResponse {
	s.Body = v
	return s
}

type ModifyVServerGroupBackendServersRequest struct {
	NewBackendServers    *string `json:"NewBackendServers,omitempty" xml:"NewBackendServers,omitempty"`
	OldBackendServers    *string `json:"OldBackendServers,omitempty" xml:"OldBackendServers,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VServerGroupId       *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s ModifyVServerGroupBackendServersRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVServerGroupBackendServersRequest) GoString() string {
	return s.String()
}

func (s *ModifyVServerGroupBackendServersRequest) SetNewBackendServers(v string) *ModifyVServerGroupBackendServersRequest {
	s.NewBackendServers = &v
	return s
}

func (s *ModifyVServerGroupBackendServersRequest) SetOldBackendServers(v string) *ModifyVServerGroupBackendServersRequest {
	s.OldBackendServers = &v
	return s
}

func (s *ModifyVServerGroupBackendServersRequest) SetOwnerAccount(v string) *ModifyVServerGroupBackendServersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyVServerGroupBackendServersRequest) SetOwnerId(v int64) *ModifyVServerGroupBackendServersRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyVServerGroupBackendServersRequest) SetRegionId(v string) *ModifyVServerGroupBackendServersRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVServerGroupBackendServersRequest) SetResourceOwnerAccount(v string) *ModifyVServerGroupBackendServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyVServerGroupBackendServersRequest) SetResourceOwnerId(v int64) *ModifyVServerGroupBackendServersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyVServerGroupBackendServersRequest) SetVServerGroupId(v string) *ModifyVServerGroupBackendServersRequest {
	s.VServerGroupId = &v
	return s
}

type ModifyVServerGroupBackendServersResponseBody struct {
	BackendServers *ModifyVServerGroupBackendServersResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	RequestId      *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VServerGroupId *string                                                     `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s ModifyVServerGroupBackendServersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVServerGroupBackendServersResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVServerGroupBackendServersResponseBody) SetBackendServers(v *ModifyVServerGroupBackendServersResponseBodyBackendServers) *ModifyVServerGroupBackendServersResponseBody {
	s.BackendServers = v
	return s
}

func (s *ModifyVServerGroupBackendServersResponseBody) SetRequestId(v string) *ModifyVServerGroupBackendServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVServerGroupBackendServersResponseBody) SetVServerGroupId(v string) *ModifyVServerGroupBackendServersResponseBody {
	s.VServerGroupId = &v
	return s
}

type ModifyVServerGroupBackendServersResponseBodyBackendServers struct {
	BackendServer []*ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s ModifyVServerGroupBackendServersResponseBodyBackendServers) String() string {
	return tea.Prettify(s)
}

func (s ModifyVServerGroupBackendServersResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *ModifyVServerGroupBackendServersResponseBodyBackendServers) SetBackendServer(v []*ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) *ModifyVServerGroupBackendServersResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

type ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Port        *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ServerId    *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight      *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) String() string {
	return tea.Prettify(s)
}

func (s ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetDescription(v string) *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Description = &v
	return s
}

func (s *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetPort(v int32) *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Port = &v
	return s
}

func (s *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetServerId(v string) *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetType(v string) *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Type = &v
	return s
}

func (s *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetWeight(v int32) *ModifyVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

type ModifyVServerGroupBackendServersResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyVServerGroupBackendServersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyVServerGroupBackendServersResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVServerGroupBackendServersResponse) GoString() string {
	return s.String()
}

func (s *ModifyVServerGroupBackendServersResponse) SetHeaders(v map[string]*string) *ModifyVServerGroupBackendServersResponse {
	s.Headers = v
	return s
}

func (s *ModifyVServerGroupBackendServersResponse) SetStatusCode(v int32) *ModifyVServerGroupBackendServersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVServerGroupBackendServersResponse) SetBody(v *ModifyVServerGroupBackendServersResponseBody) *ModifyVServerGroupBackendServersResponse {
	s.Body = v
	return s
}

type RemoveAccessControlListEntryRequest struct {
	AclEntrys            *string `json:"AclEntrys,omitempty" xml:"AclEntrys,omitempty"`
	AclId                *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RemoveAccessControlListEntryRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveAccessControlListEntryRequest) GoString() string {
	return s.String()
}

func (s *RemoveAccessControlListEntryRequest) SetAclEntrys(v string) *RemoveAccessControlListEntryRequest {
	s.AclEntrys = &v
	return s
}

func (s *RemoveAccessControlListEntryRequest) SetAclId(v string) *RemoveAccessControlListEntryRequest {
	s.AclId = &v
	return s
}

func (s *RemoveAccessControlListEntryRequest) SetOwnerAccount(v string) *RemoveAccessControlListEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveAccessControlListEntryRequest) SetOwnerId(v int64) *RemoveAccessControlListEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveAccessControlListEntryRequest) SetRegionId(v string) *RemoveAccessControlListEntryRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveAccessControlListEntryRequest) SetResourceOwnerAccount(v string) *RemoveAccessControlListEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveAccessControlListEntryRequest) SetResourceOwnerId(v int64) *RemoveAccessControlListEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

type RemoveAccessControlListEntryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveAccessControlListEntryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveAccessControlListEntryResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAccessControlListEntryResponseBody) SetRequestId(v string) *RemoveAccessControlListEntryResponseBody {
	s.RequestId = &v
	return s
}

type RemoveAccessControlListEntryResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveAccessControlListEntryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveAccessControlListEntryResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveAccessControlListEntryResponse) GoString() string {
	return s.String()
}

func (s *RemoveAccessControlListEntryResponse) SetHeaders(v map[string]*string) *RemoveAccessControlListEntryResponse {
	s.Headers = v
	return s
}

func (s *RemoveAccessControlListEntryResponse) SetStatusCode(v int32) *RemoveAccessControlListEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveAccessControlListEntryResponse) SetBody(v *RemoveAccessControlListEntryResponseBody) *RemoveAccessControlListEntryResponse {
	s.Body = v
	return s
}

type RemoveBackendServersRequest struct {
	BackendServers       *string `json:"BackendServers,omitempty" xml:"BackendServers,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RemoveBackendServersRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveBackendServersRequest) GoString() string {
	return s.String()
}

func (s *RemoveBackendServersRequest) SetBackendServers(v string) *RemoveBackendServersRequest {
	s.BackendServers = &v
	return s
}

func (s *RemoveBackendServersRequest) SetLoadBalancerId(v string) *RemoveBackendServersRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *RemoveBackendServersRequest) SetOwnerAccount(v string) *RemoveBackendServersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveBackendServersRequest) SetOwnerId(v int64) *RemoveBackendServersRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveBackendServersRequest) SetRegionId(v string) *RemoveBackendServersRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveBackendServersRequest) SetResourceOwnerAccount(v string) *RemoveBackendServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveBackendServersRequest) SetResourceOwnerId(v int64) *RemoveBackendServersRequest {
	s.ResourceOwnerId = &v
	return s
}

type RemoveBackendServersResponseBody struct {
	BackendServers *RemoveBackendServersResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	LoadBalancerId *string                                         `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveBackendServersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveBackendServersResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveBackendServersResponseBody) SetBackendServers(v *RemoveBackendServersResponseBodyBackendServers) *RemoveBackendServersResponseBody {
	s.BackendServers = v
	return s
}

func (s *RemoveBackendServersResponseBody) SetLoadBalancerId(v string) *RemoveBackendServersResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *RemoveBackendServersResponseBody) SetRequestId(v string) *RemoveBackendServersResponseBody {
	s.RequestId = &v
	return s
}

type RemoveBackendServersResponseBodyBackendServers struct {
	BackendServer []*RemoveBackendServersResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s RemoveBackendServersResponseBodyBackendServers) String() string {
	return tea.Prettify(s)
}

func (s RemoveBackendServersResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *RemoveBackendServersResponseBodyBackendServers) SetBackendServer(v []*RemoveBackendServersResponseBodyBackendServersBackendServer) *RemoveBackendServersResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

type RemoveBackendServersResponseBodyBackendServersBackendServer struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ServerId    *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight      *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s RemoveBackendServersResponseBodyBackendServersBackendServer) String() string {
	return tea.Prettify(s)
}

func (s RemoveBackendServersResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *RemoveBackendServersResponseBodyBackendServersBackendServer) SetDescription(v string) *RemoveBackendServersResponseBodyBackendServersBackendServer {
	s.Description = &v
	return s
}

func (s *RemoveBackendServersResponseBodyBackendServersBackendServer) SetServerId(v string) *RemoveBackendServersResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *RemoveBackendServersResponseBodyBackendServersBackendServer) SetType(v string) *RemoveBackendServersResponseBodyBackendServersBackendServer {
	s.Type = &v
	return s
}

func (s *RemoveBackendServersResponseBodyBackendServersBackendServer) SetWeight(v int32) *RemoveBackendServersResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

type RemoveBackendServersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveBackendServersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveBackendServersResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveBackendServersResponse) GoString() string {
	return s.String()
}

func (s *RemoveBackendServersResponse) SetHeaders(v map[string]*string) *RemoveBackendServersResponse {
	s.Headers = v
	return s
}

func (s *RemoveBackendServersResponse) SetStatusCode(v int32) *RemoveBackendServersResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveBackendServersResponse) SetBody(v *RemoveBackendServersResponseBody) *RemoveBackendServersResponse {
	s.Body = v
	return s
}

type RemoveListenerWhiteListItemRequest struct {
	ListenerPort         *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerProtocol     *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SourceItems          *string `json:"SourceItems,omitempty" xml:"SourceItems,omitempty"`
}

func (s RemoveListenerWhiteListItemRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveListenerWhiteListItemRequest) GoString() string {
	return s.String()
}

func (s *RemoveListenerWhiteListItemRequest) SetListenerPort(v int32) *RemoveListenerWhiteListItemRequest {
	s.ListenerPort = &v
	return s
}

func (s *RemoveListenerWhiteListItemRequest) SetListenerProtocol(v string) *RemoveListenerWhiteListItemRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *RemoveListenerWhiteListItemRequest) SetLoadBalancerId(v string) *RemoveListenerWhiteListItemRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *RemoveListenerWhiteListItemRequest) SetOwnerAccount(v string) *RemoveListenerWhiteListItemRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveListenerWhiteListItemRequest) SetOwnerId(v int64) *RemoveListenerWhiteListItemRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveListenerWhiteListItemRequest) SetRegionId(v string) *RemoveListenerWhiteListItemRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveListenerWhiteListItemRequest) SetResourceOwnerAccount(v string) *RemoveListenerWhiteListItemRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveListenerWhiteListItemRequest) SetResourceOwnerId(v int64) *RemoveListenerWhiteListItemRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveListenerWhiteListItemRequest) SetSourceItems(v string) *RemoveListenerWhiteListItemRequest {
	s.SourceItems = &v
	return s
}

type RemoveListenerWhiteListItemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveListenerWhiteListItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveListenerWhiteListItemResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveListenerWhiteListItemResponseBody) SetRequestId(v string) *RemoveListenerWhiteListItemResponseBody {
	s.RequestId = &v
	return s
}

type RemoveListenerWhiteListItemResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveListenerWhiteListItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveListenerWhiteListItemResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveListenerWhiteListItemResponse) GoString() string {
	return s.String()
}

func (s *RemoveListenerWhiteListItemResponse) SetHeaders(v map[string]*string) *RemoveListenerWhiteListItemResponse {
	s.Headers = v
	return s
}

func (s *RemoveListenerWhiteListItemResponse) SetStatusCode(v int32) *RemoveListenerWhiteListItemResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveListenerWhiteListItemResponse) SetBody(v *RemoveListenerWhiteListItemResponseBody) *RemoveListenerWhiteListItemResponse {
	s.Body = v
	return s
}

type RemoveTagsRequest struct {
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s RemoveTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagsRequest) GoString() string {
	return s.String()
}

func (s *RemoveTagsRequest) SetLoadBalancerId(v string) *RemoveTagsRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *RemoveTagsRequest) SetOwnerAccount(v string) *RemoveTagsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveTagsRequest) SetOwnerId(v int64) *RemoveTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveTagsRequest) SetRegionId(v string) *RemoveTagsRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveTagsRequest) SetResourceOwnerAccount(v string) *RemoveTagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveTagsRequest) SetResourceOwnerId(v int64) *RemoveTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveTagsRequest) SetTags(v string) *RemoveTagsRequest {
	s.Tags = &v
	return s
}

type RemoveTagsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTagsResponseBody) SetRequestId(v string) *RemoveTagsResponseBody {
	s.RequestId = &v
	return s
}

type RemoveTagsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagsResponse) GoString() string {
	return s.String()
}

func (s *RemoveTagsResponse) SetHeaders(v map[string]*string) *RemoveTagsResponse {
	s.Headers = v
	return s
}

func (s *RemoveTagsResponse) SetStatusCode(v int32) *RemoveTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveTagsResponse) SetBody(v *RemoveTagsResponseBody) *RemoveTagsResponse {
	s.Body = v
	return s
}

type RemoveVServerGroupBackendServersRequest struct {
	BackendServers       *string `json:"BackendServers,omitempty" xml:"BackendServers,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VServerGroupId       *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s RemoveVServerGroupBackendServersRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveVServerGroupBackendServersRequest) GoString() string {
	return s.String()
}

func (s *RemoveVServerGroupBackendServersRequest) SetBackendServers(v string) *RemoveVServerGroupBackendServersRequest {
	s.BackendServers = &v
	return s
}

func (s *RemoveVServerGroupBackendServersRequest) SetOwnerAccount(v string) *RemoveVServerGroupBackendServersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveVServerGroupBackendServersRequest) SetOwnerId(v int64) *RemoveVServerGroupBackendServersRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveVServerGroupBackendServersRequest) SetRegionId(v string) *RemoveVServerGroupBackendServersRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveVServerGroupBackendServersRequest) SetResourceOwnerAccount(v string) *RemoveVServerGroupBackendServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveVServerGroupBackendServersRequest) SetResourceOwnerId(v int64) *RemoveVServerGroupBackendServersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveVServerGroupBackendServersRequest) SetVServerGroupId(v string) *RemoveVServerGroupBackendServersRequest {
	s.VServerGroupId = &v
	return s
}

type RemoveVServerGroupBackendServersResponseBody struct {
	BackendServers *RemoveVServerGroupBackendServersResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	RequestId      *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VServerGroupId *string                                                     `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s RemoveVServerGroupBackendServersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveVServerGroupBackendServersResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveVServerGroupBackendServersResponseBody) SetBackendServers(v *RemoveVServerGroupBackendServersResponseBodyBackendServers) *RemoveVServerGroupBackendServersResponseBody {
	s.BackendServers = v
	return s
}

func (s *RemoveVServerGroupBackendServersResponseBody) SetRequestId(v string) *RemoveVServerGroupBackendServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveVServerGroupBackendServersResponseBody) SetVServerGroupId(v string) *RemoveVServerGroupBackendServersResponseBody {
	s.VServerGroupId = &v
	return s
}

type RemoveVServerGroupBackendServersResponseBodyBackendServers struct {
	BackendServer []*RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s RemoveVServerGroupBackendServersResponseBodyBackendServers) String() string {
	return tea.Prettify(s)
}

func (s RemoveVServerGroupBackendServersResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *RemoveVServerGroupBackendServersResponseBodyBackendServers) SetBackendServer(v []*RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer) *RemoveVServerGroupBackendServersResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

type RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer struct {
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight   *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer) String() string {
	return tea.Prettify(s)
}

func (s RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetPort(v int32) *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Port = &v
	return s
}

func (s *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetServerId(v string) *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetType(v string) *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Type = &v
	return s
}

func (s *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer) SetWeight(v int32) *RemoveVServerGroupBackendServersResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

type RemoveVServerGroupBackendServersResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveVServerGroupBackendServersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveVServerGroupBackendServersResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveVServerGroupBackendServersResponse) GoString() string {
	return s.String()
}

func (s *RemoveVServerGroupBackendServersResponse) SetHeaders(v map[string]*string) *RemoveVServerGroupBackendServersResponse {
	s.Headers = v
	return s
}

func (s *RemoveVServerGroupBackendServersResponse) SetStatusCode(v int32) *RemoveVServerGroupBackendServersResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveVServerGroupBackendServersResponse) SetBody(v *RemoveVServerGroupBackendServersResponseBody) *RemoveVServerGroupBackendServersResponse {
	s.Body = v
	return s
}

type SetAccessControlListAttributeRequest struct {
	AclId                *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclName              *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetAccessControlListAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAccessControlListAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetAccessControlListAttributeRequest) SetAclId(v string) *SetAccessControlListAttributeRequest {
	s.AclId = &v
	return s
}

func (s *SetAccessControlListAttributeRequest) SetAclName(v string) *SetAccessControlListAttributeRequest {
	s.AclName = &v
	return s
}

func (s *SetAccessControlListAttributeRequest) SetOwnerAccount(v string) *SetAccessControlListAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetAccessControlListAttributeRequest) SetOwnerId(v int64) *SetAccessControlListAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *SetAccessControlListAttributeRequest) SetRegionId(v string) *SetAccessControlListAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *SetAccessControlListAttributeRequest) SetResourceOwnerAccount(v string) *SetAccessControlListAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetAccessControlListAttributeRequest) SetResourceOwnerId(v int64) *SetAccessControlListAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type SetAccessControlListAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAccessControlListAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAccessControlListAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *SetAccessControlListAttributeResponseBody) SetRequestId(v string) *SetAccessControlListAttributeResponseBody {
	s.RequestId = &v
	return s
}

type SetAccessControlListAttributeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetAccessControlListAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetAccessControlListAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAccessControlListAttributeResponse) GoString() string {
	return s.String()
}

func (s *SetAccessControlListAttributeResponse) SetHeaders(v map[string]*string) *SetAccessControlListAttributeResponse {
	s.Headers = v
	return s
}

func (s *SetAccessControlListAttributeResponse) SetStatusCode(v int32) *SetAccessControlListAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAccessControlListAttributeResponse) SetBody(v *SetAccessControlListAttributeResponseBody) *SetAccessControlListAttributeResponse {
	s.Body = v
	return s
}

type SetBackendServersRequest struct {
	BackendServers       *string `json:"BackendServers,omitempty" xml:"BackendServers,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetBackendServersRequest) String() string {
	return tea.Prettify(s)
}

func (s SetBackendServersRequest) GoString() string {
	return s.String()
}

func (s *SetBackendServersRequest) SetBackendServers(v string) *SetBackendServersRequest {
	s.BackendServers = &v
	return s
}

func (s *SetBackendServersRequest) SetLoadBalancerId(v string) *SetBackendServersRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetBackendServersRequest) SetOwnerAccount(v string) *SetBackendServersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetBackendServersRequest) SetOwnerId(v int64) *SetBackendServersRequest {
	s.OwnerId = &v
	return s
}

func (s *SetBackendServersRequest) SetRegionId(v string) *SetBackendServersRequest {
	s.RegionId = &v
	return s
}

func (s *SetBackendServersRequest) SetResourceOwnerAccount(v string) *SetBackendServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetBackendServersRequest) SetResourceOwnerId(v int64) *SetBackendServersRequest {
	s.ResourceOwnerId = &v
	return s
}

type SetBackendServersResponseBody struct {
	BackendServers *SetBackendServersResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	LoadBalancerId *string                                      `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetBackendServersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetBackendServersResponseBody) GoString() string {
	return s.String()
}

func (s *SetBackendServersResponseBody) SetBackendServers(v *SetBackendServersResponseBodyBackendServers) *SetBackendServersResponseBody {
	s.BackendServers = v
	return s
}

func (s *SetBackendServersResponseBody) SetLoadBalancerId(v string) *SetBackendServersResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *SetBackendServersResponseBody) SetRequestId(v string) *SetBackendServersResponseBody {
	s.RequestId = &v
	return s
}

type SetBackendServersResponseBodyBackendServers struct {
	BackendServer []*SetBackendServersResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s SetBackendServersResponseBodyBackendServers) String() string {
	return tea.Prettify(s)
}

func (s SetBackendServersResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *SetBackendServersResponseBodyBackendServers) SetBackendServer(v []*SetBackendServersResponseBodyBackendServersBackendServer) *SetBackendServersResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

type SetBackendServersResponseBodyBackendServersBackendServer struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ServerId    *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight      *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s SetBackendServersResponseBodyBackendServersBackendServer) String() string {
	return tea.Prettify(s)
}

func (s SetBackendServersResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) SetDescription(v string) *SetBackendServersResponseBodyBackendServersBackendServer {
	s.Description = &v
	return s
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) SetServerId(v string) *SetBackendServersResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) SetType(v string) *SetBackendServersResponseBodyBackendServersBackendServer {
	s.Type = &v
	return s
}

func (s *SetBackendServersResponseBodyBackendServersBackendServer) SetWeight(v string) *SetBackendServersResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

type SetBackendServersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetBackendServersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetBackendServersResponse) String() string {
	return tea.Prettify(s)
}

func (s SetBackendServersResponse) GoString() string {
	return s.String()
}

func (s *SetBackendServersResponse) SetHeaders(v map[string]*string) *SetBackendServersResponse {
	s.Headers = v
	return s
}

func (s *SetBackendServersResponse) SetStatusCode(v int32) *SetBackendServersResponse {
	s.StatusCode = &v
	return s
}

func (s *SetBackendServersResponse) SetBody(v *SetBackendServersResponseBody) *SetBackendServersResponse {
	s.Body = v
	return s
}

type SetCACertificateNameRequest struct {
	CACertificateId      *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	CACertificateName    *string `json:"CACertificateName,omitempty" xml:"CACertificateName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetCACertificateNameRequest) String() string {
	return tea.Prettify(s)
}

func (s SetCACertificateNameRequest) GoString() string {
	return s.String()
}

func (s *SetCACertificateNameRequest) SetCACertificateId(v string) *SetCACertificateNameRequest {
	s.CACertificateId = &v
	return s
}

func (s *SetCACertificateNameRequest) SetCACertificateName(v string) *SetCACertificateNameRequest {
	s.CACertificateName = &v
	return s
}

func (s *SetCACertificateNameRequest) SetOwnerAccount(v string) *SetCACertificateNameRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetCACertificateNameRequest) SetOwnerId(v int64) *SetCACertificateNameRequest {
	s.OwnerId = &v
	return s
}

func (s *SetCACertificateNameRequest) SetRegionId(v string) *SetCACertificateNameRequest {
	s.RegionId = &v
	return s
}

func (s *SetCACertificateNameRequest) SetResourceOwnerAccount(v string) *SetCACertificateNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetCACertificateNameRequest) SetResourceOwnerId(v int64) *SetCACertificateNameRequest {
	s.ResourceOwnerId = &v
	return s
}

type SetCACertificateNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCACertificateNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetCACertificateNameResponseBody) GoString() string {
	return s.String()
}

func (s *SetCACertificateNameResponseBody) SetRequestId(v string) *SetCACertificateNameResponseBody {
	s.RequestId = &v
	return s
}

type SetCACertificateNameResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetCACertificateNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetCACertificateNameResponse) String() string {
	return tea.Prettify(s)
}

func (s SetCACertificateNameResponse) GoString() string {
	return s.String()
}

func (s *SetCACertificateNameResponse) SetHeaders(v map[string]*string) *SetCACertificateNameResponse {
	s.Headers = v
	return s
}

func (s *SetCACertificateNameResponse) SetStatusCode(v int32) *SetCACertificateNameResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCACertificateNameResponse) SetBody(v *SetCACertificateNameResponseBody) *SetCACertificateNameResponse {
	s.Body = v
	return s
}

type SetDomainExtensionAttributeRequest struct {
	DomainExtensionId    *string `json:"DomainExtensionId,omitempty" xml:"DomainExtensionId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ServerCertificateId  *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
}

func (s SetDomainExtensionAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDomainExtensionAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetDomainExtensionAttributeRequest) SetDomainExtensionId(v string) *SetDomainExtensionAttributeRequest {
	s.DomainExtensionId = &v
	return s
}

func (s *SetDomainExtensionAttributeRequest) SetOwnerAccount(v string) *SetDomainExtensionAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetDomainExtensionAttributeRequest) SetOwnerId(v int64) *SetDomainExtensionAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDomainExtensionAttributeRequest) SetRegionId(v string) *SetDomainExtensionAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *SetDomainExtensionAttributeRequest) SetResourceOwnerAccount(v string) *SetDomainExtensionAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetDomainExtensionAttributeRequest) SetResourceOwnerId(v int64) *SetDomainExtensionAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetDomainExtensionAttributeRequest) SetServerCertificateId(v string) *SetDomainExtensionAttributeRequest {
	s.ServerCertificateId = &v
	return s
}

type SetDomainExtensionAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDomainExtensionAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDomainExtensionAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *SetDomainExtensionAttributeResponseBody) SetRequestId(v string) *SetDomainExtensionAttributeResponseBody {
	s.RequestId = &v
	return s
}

type SetDomainExtensionAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDomainExtensionAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDomainExtensionAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDomainExtensionAttributeResponse) GoString() string {
	return s.String()
}

func (s *SetDomainExtensionAttributeResponse) SetHeaders(v map[string]*string) *SetDomainExtensionAttributeResponse {
	s.Headers = v
	return s
}

func (s *SetDomainExtensionAttributeResponse) SetStatusCode(v int32) *SetDomainExtensionAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDomainExtensionAttributeResponse) SetBody(v *SetDomainExtensionAttributeResponseBody) *SetDomainExtensionAttributeResponse {
	s.Body = v
	return s
}

type SetListenerAccessControlStatusRequest struct {
	AccessControlStatus  *string `json:"AccessControlStatus,omitempty" xml:"AccessControlStatus,omitempty"`
	ListenerPort         *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerProtocol     *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetListenerAccessControlStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetListenerAccessControlStatusRequest) GoString() string {
	return s.String()
}

func (s *SetListenerAccessControlStatusRequest) SetAccessControlStatus(v string) *SetListenerAccessControlStatusRequest {
	s.AccessControlStatus = &v
	return s
}

func (s *SetListenerAccessControlStatusRequest) SetListenerPort(v int32) *SetListenerAccessControlStatusRequest {
	s.ListenerPort = &v
	return s
}

func (s *SetListenerAccessControlStatusRequest) SetListenerProtocol(v string) *SetListenerAccessControlStatusRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *SetListenerAccessControlStatusRequest) SetLoadBalancerId(v string) *SetListenerAccessControlStatusRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetListenerAccessControlStatusRequest) SetOwnerAccount(v string) *SetListenerAccessControlStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetListenerAccessControlStatusRequest) SetOwnerId(v int64) *SetListenerAccessControlStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *SetListenerAccessControlStatusRequest) SetRegionId(v string) *SetListenerAccessControlStatusRequest {
	s.RegionId = &v
	return s
}

func (s *SetListenerAccessControlStatusRequest) SetResourceOwnerAccount(v string) *SetListenerAccessControlStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetListenerAccessControlStatusRequest) SetResourceOwnerId(v int64) *SetListenerAccessControlStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

type SetListenerAccessControlStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetListenerAccessControlStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetListenerAccessControlStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetListenerAccessControlStatusResponseBody) SetRequestId(v string) *SetListenerAccessControlStatusResponseBody {
	s.RequestId = &v
	return s
}

type SetListenerAccessControlStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetListenerAccessControlStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetListenerAccessControlStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetListenerAccessControlStatusResponse) GoString() string {
	return s.String()
}

func (s *SetListenerAccessControlStatusResponse) SetHeaders(v map[string]*string) *SetListenerAccessControlStatusResponse {
	s.Headers = v
	return s
}

func (s *SetListenerAccessControlStatusResponse) SetStatusCode(v int32) *SetListenerAccessControlStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetListenerAccessControlStatusResponse) SetBody(v *SetListenerAccessControlStatusResponseBody) *SetListenerAccessControlStatusResponse {
	s.Body = v
	return s
}

type SetLoadBalancerDeleteProtectionRequest struct {
	DeleteProtection     *string `json:"DeleteProtection,omitempty" xml:"DeleteProtection,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetLoadBalancerDeleteProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerDeleteProtectionRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerDeleteProtectionRequest) SetDeleteProtection(v string) *SetLoadBalancerDeleteProtectionRequest {
	s.DeleteProtection = &v
	return s
}

func (s *SetLoadBalancerDeleteProtectionRequest) SetLoadBalancerId(v string) *SetLoadBalancerDeleteProtectionRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerDeleteProtectionRequest) SetOwnerAccount(v string) *SetLoadBalancerDeleteProtectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerDeleteProtectionRequest) SetOwnerId(v int64) *SetLoadBalancerDeleteProtectionRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerDeleteProtectionRequest) SetRegionId(v string) *SetLoadBalancerDeleteProtectionRequest {
	s.RegionId = &v
	return s
}

func (s *SetLoadBalancerDeleteProtectionRequest) SetResourceOwnerAccount(v string) *SetLoadBalancerDeleteProtectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetLoadBalancerDeleteProtectionRequest) SetResourceOwnerId(v int64) *SetLoadBalancerDeleteProtectionRequest {
	s.ResourceOwnerId = &v
	return s
}

type SetLoadBalancerDeleteProtectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLoadBalancerDeleteProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerDeleteProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerDeleteProtectionResponseBody) SetRequestId(v string) *SetLoadBalancerDeleteProtectionResponseBody {
	s.RequestId = &v
	return s
}

type SetLoadBalancerDeleteProtectionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetLoadBalancerDeleteProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetLoadBalancerDeleteProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerDeleteProtectionResponse) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerDeleteProtectionResponse) SetHeaders(v map[string]*string) *SetLoadBalancerDeleteProtectionResponse {
	s.Headers = v
	return s
}

func (s *SetLoadBalancerDeleteProtectionResponse) SetStatusCode(v int32) *SetLoadBalancerDeleteProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLoadBalancerDeleteProtectionResponse) SetBody(v *SetLoadBalancerDeleteProtectionResponseBody) *SetLoadBalancerDeleteProtectionResponse {
	s.Body = v
	return s
}

type SetLoadBalancerHTTPListenerAttributeRequest struct {
	AclId                  *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclStatus              *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	AclType                *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	Bandwidth              *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Cookie                 *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	CookieTimeout          *int32  `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Gzip                   *string `json:"Gzip,omitempty" xml:"Gzip,omitempty"`
	HealthCheck            *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckConnectPort *int32  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckDomain      *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	HealthCheckHttpCode    *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckInterval    *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckMethod      *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	HealthCheckTimeout     *int32  `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	HealthCheckURI         *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	HealthyThreshold       *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	IdleTimeout            *int32  `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	ListenerPort           *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	LoadBalancerId         *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestTimeout         *int32  `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Scheduler              *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	StickySession          *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	StickySessionType      *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	UnhealthyThreshold     *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	VServerGroup           *string `json:"VServerGroup,omitempty" xml:"VServerGroup,omitempty"`
	VServerGroupId         *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	XForwardedFor          *string `json:"XForwardedFor,omitempty" xml:"XForwardedFor,omitempty"`
	XForwardedFor_SLBID    *string `json:"XForwardedFor_SLBID,omitempty" xml:"XForwardedFor_SLBID,omitempty"`
	XForwardedFor_SLBIP    *string `json:"XForwardedFor_SLBIP,omitempty" xml:"XForwardedFor_SLBIP,omitempty"`
	XForwardedFor_proto    *string `json:"XForwardedFor_proto,omitempty" xml:"XForwardedFor_proto,omitempty"`
}

func (s SetLoadBalancerHTTPListenerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerHTTPListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetAclId(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.AclId = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetAclStatus(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.AclStatus = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetAclType(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.AclType = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetBandwidth(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.Bandwidth = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetCookie(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.Cookie = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetCookieTimeout(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.CookieTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetDescription(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.Description = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetGzip(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.Gzip = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheck(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthCheck = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckConnectPort(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckDomain(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthCheckDomain = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckHttpCode(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckInterval(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckMethod(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthCheckMethod = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckTimeout(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthCheckTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckURI(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthCheckURI = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthyThreshold(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetIdleTimeout(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.IdleTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetListenerPort(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetLoadBalancerId(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetOwnerAccount(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetOwnerId(v int64) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetRegionId(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetRequestTimeout(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.RequestTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetResourceOwnerAccount(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetResourceOwnerId(v int64) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetScheduler(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.Scheduler = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetStickySession(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.StickySession = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetStickySessionType(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.StickySessionType = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetUnhealthyThreshold(v int32) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetVServerGroup(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.VServerGroup = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetVServerGroupId(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.VServerGroupId = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetXForwardedFor(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.XForwardedFor = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetXForwardedFor_SLBID(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.XForwardedFor_SLBID = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetXForwardedFor_SLBIP(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.XForwardedFor_SLBIP = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeRequest) SetXForwardedFor_proto(v string) *SetLoadBalancerHTTPListenerAttributeRequest {
	s.XForwardedFor_proto = &v
	return s
}

type SetLoadBalancerHTTPListenerAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLoadBalancerHTTPListenerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerHTTPListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerHTTPListenerAttributeResponseBody) SetRequestId(v string) *SetLoadBalancerHTTPListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

type SetLoadBalancerHTTPListenerAttributeResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetLoadBalancerHTTPListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetLoadBalancerHTTPListenerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerHTTPListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerHTTPListenerAttributeResponse) SetHeaders(v map[string]*string) *SetLoadBalancerHTTPListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeResponse) SetStatusCode(v int32) *SetLoadBalancerHTTPListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeResponse) SetBody(v *SetLoadBalancerHTTPListenerAttributeResponseBody) *SetLoadBalancerHTTPListenerAttributeResponse {
	s.Body = v
	return s
}

type SetLoadBalancerHTTPSListenerAttributeRequest struct {
	AclId                  *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclStatus              *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	AclType                *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	Bandwidth              *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	CACertificateId        *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	Cookie                 *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	CookieTimeout          *int32  `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableHttp2            *string `json:"EnableHttp2,omitempty" xml:"EnableHttp2,omitempty"`
	Gzip                   *string `json:"Gzip,omitempty" xml:"Gzip,omitempty"`
	HealthCheck            *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckConnectPort *int32  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckDomain      *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	HealthCheckHttpCode    *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckInterval    *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckMethod      *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	HealthCheckTimeout     *int32  `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	HealthCheckURI         *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	HealthyThreshold       *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	IdleTimeout            *int32  `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	ListenerPort           *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	LoadBalancerId         *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestTimeout         *int32  `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Scheduler              *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	ServerCertificateId    *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
	StickySession          *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	StickySessionType      *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	TLSCipherPolicy        *string `json:"TLSCipherPolicy,omitempty" xml:"TLSCipherPolicy,omitempty"`
	UnhealthyThreshold     *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	VServerGroup           *string `json:"VServerGroup,omitempty" xml:"VServerGroup,omitempty"`
	VServerGroupId         *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	XForwardedFor          *string `json:"XForwardedFor,omitempty" xml:"XForwardedFor,omitempty"`
	XForwardedFor_SLBID    *string `json:"XForwardedFor_SLBID,omitempty" xml:"XForwardedFor_SLBID,omitempty"`
	XForwardedFor_SLBIP    *string `json:"XForwardedFor_SLBIP,omitempty" xml:"XForwardedFor_SLBIP,omitempty"`
	XForwardedFor_proto    *string `json:"XForwardedFor_proto,omitempty" xml:"XForwardedFor_proto,omitempty"`
}

func (s SetLoadBalancerHTTPSListenerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerHTTPSListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetAclId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.AclId = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetAclStatus(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.AclStatus = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetAclType(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.AclType = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetBandwidth(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.Bandwidth = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetCACertificateId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.CACertificateId = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetCookie(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.Cookie = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetCookieTimeout(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.CookieTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetDescription(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.Description = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetEnableHttp2(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.EnableHttp2 = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetGzip(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.Gzip = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheck(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheck = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckConnectPort(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckDomain(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckDomain = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckHttpCode(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckInterval(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckMethod(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckMethod = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckTimeout(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckURI(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthCheckURI = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthyThreshold(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetIdleTimeout(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.IdleTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetListenerPort(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetLoadBalancerId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetOwnerAccount(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetOwnerId(v int64) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetRegionId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetRequestTimeout(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.RequestTimeout = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetResourceOwnerAccount(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetResourceOwnerId(v int64) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetScheduler(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.Scheduler = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetServerCertificateId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.ServerCertificateId = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetStickySession(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.StickySession = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetStickySessionType(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.StickySessionType = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetTLSCipherPolicy(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.TLSCipherPolicy = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetUnhealthyThreshold(v int32) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetVServerGroup(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.VServerGroup = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetVServerGroupId(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.VServerGroupId = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetXForwardedFor(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.XForwardedFor = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetXForwardedFor_SLBID(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.XForwardedFor_SLBID = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetXForwardedFor_SLBIP(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.XForwardedFor_SLBIP = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeRequest) SetXForwardedFor_proto(v string) *SetLoadBalancerHTTPSListenerAttributeRequest {
	s.XForwardedFor_proto = &v
	return s
}

type SetLoadBalancerHTTPSListenerAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLoadBalancerHTTPSListenerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerHTTPSListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerHTTPSListenerAttributeResponseBody) SetRequestId(v string) *SetLoadBalancerHTTPSListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

type SetLoadBalancerHTTPSListenerAttributeResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetLoadBalancerHTTPSListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetLoadBalancerHTTPSListenerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerHTTPSListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerHTTPSListenerAttributeResponse) SetHeaders(v map[string]*string) *SetLoadBalancerHTTPSListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeResponse) SetStatusCode(v int32) *SetLoadBalancerHTTPSListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeResponse) SetBody(v *SetLoadBalancerHTTPSListenerAttributeResponseBody) *SetLoadBalancerHTTPSListenerAttributeResponse {
	s.Body = v
	return s
}

type SetLoadBalancerModificationProtectionRequest struct {
	LoadBalancerId               *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	ModificationProtectionReason *string `json:"ModificationProtectionReason,omitempty" xml:"ModificationProtectionReason,omitempty"`
	ModificationProtectionStatus *string `json:"ModificationProtectionStatus,omitempty" xml:"ModificationProtectionStatus,omitempty"`
	OwnerAccount                 *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount         *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId              *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetLoadBalancerModificationProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerModificationProtectionRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerModificationProtectionRequest) SetLoadBalancerId(v string) *SetLoadBalancerModificationProtectionRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerModificationProtectionRequest) SetModificationProtectionReason(v string) *SetLoadBalancerModificationProtectionRequest {
	s.ModificationProtectionReason = &v
	return s
}

func (s *SetLoadBalancerModificationProtectionRequest) SetModificationProtectionStatus(v string) *SetLoadBalancerModificationProtectionRequest {
	s.ModificationProtectionStatus = &v
	return s
}

func (s *SetLoadBalancerModificationProtectionRequest) SetOwnerAccount(v string) *SetLoadBalancerModificationProtectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerModificationProtectionRequest) SetOwnerId(v int64) *SetLoadBalancerModificationProtectionRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerModificationProtectionRequest) SetRegionId(v string) *SetLoadBalancerModificationProtectionRequest {
	s.RegionId = &v
	return s
}

func (s *SetLoadBalancerModificationProtectionRequest) SetResourceOwnerAccount(v string) *SetLoadBalancerModificationProtectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetLoadBalancerModificationProtectionRequest) SetResourceOwnerId(v int64) *SetLoadBalancerModificationProtectionRequest {
	s.ResourceOwnerId = &v
	return s
}

type SetLoadBalancerModificationProtectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLoadBalancerModificationProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerModificationProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerModificationProtectionResponseBody) SetRequestId(v string) *SetLoadBalancerModificationProtectionResponseBody {
	s.RequestId = &v
	return s
}

type SetLoadBalancerModificationProtectionResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetLoadBalancerModificationProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetLoadBalancerModificationProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerModificationProtectionResponse) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerModificationProtectionResponse) SetHeaders(v map[string]*string) *SetLoadBalancerModificationProtectionResponse {
	s.Headers = v
	return s
}

func (s *SetLoadBalancerModificationProtectionResponse) SetStatusCode(v int32) *SetLoadBalancerModificationProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLoadBalancerModificationProtectionResponse) SetBody(v *SetLoadBalancerModificationProtectionResponseBody) *SetLoadBalancerModificationProtectionResponse {
	s.Body = v
	return s
}

type SetLoadBalancerNameRequest struct {
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LoadBalancerName     *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetLoadBalancerNameRequest) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerNameRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerNameRequest) SetLoadBalancerId(v string) *SetLoadBalancerNameRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerNameRequest) SetLoadBalancerName(v string) *SetLoadBalancerNameRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *SetLoadBalancerNameRequest) SetOwnerAccount(v string) *SetLoadBalancerNameRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerNameRequest) SetOwnerId(v int64) *SetLoadBalancerNameRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerNameRequest) SetRegionId(v string) *SetLoadBalancerNameRequest {
	s.RegionId = &v
	return s
}

func (s *SetLoadBalancerNameRequest) SetResourceOwnerAccount(v string) *SetLoadBalancerNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetLoadBalancerNameRequest) SetResourceOwnerId(v int64) *SetLoadBalancerNameRequest {
	s.ResourceOwnerId = &v
	return s
}

type SetLoadBalancerNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLoadBalancerNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerNameResponseBody) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerNameResponseBody) SetRequestId(v string) *SetLoadBalancerNameResponseBody {
	s.RequestId = &v
	return s
}

type SetLoadBalancerNameResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetLoadBalancerNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetLoadBalancerNameResponse) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerNameResponse) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerNameResponse) SetHeaders(v map[string]*string) *SetLoadBalancerNameResponse {
	s.Headers = v
	return s
}

func (s *SetLoadBalancerNameResponse) SetStatusCode(v int32) *SetLoadBalancerNameResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLoadBalancerNameResponse) SetBody(v *SetLoadBalancerNameResponseBody) *SetLoadBalancerNameResponse {
	s.Body = v
	return s
}

type SetLoadBalancerStatusRequest struct {
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LoadBalancerStatus   *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetLoadBalancerStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerStatusRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerStatusRequest) SetLoadBalancerId(v string) *SetLoadBalancerStatusRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) SetLoadBalancerStatus(v string) *SetLoadBalancerStatusRequest {
	s.LoadBalancerStatus = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) SetOwnerAccount(v string) *SetLoadBalancerStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) SetOwnerId(v int64) *SetLoadBalancerStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) SetRegionId(v string) *SetLoadBalancerStatusRequest {
	s.RegionId = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) SetResourceOwnerAccount(v string) *SetLoadBalancerStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetLoadBalancerStatusRequest) SetResourceOwnerId(v int64) *SetLoadBalancerStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

type SetLoadBalancerStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLoadBalancerStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerStatusResponseBody) SetRequestId(v string) *SetLoadBalancerStatusResponseBody {
	s.RequestId = &v
	return s
}

type SetLoadBalancerStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetLoadBalancerStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetLoadBalancerStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerStatusResponse) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerStatusResponse) SetHeaders(v map[string]*string) *SetLoadBalancerStatusResponse {
	s.Headers = v
	return s
}

func (s *SetLoadBalancerStatusResponse) SetStatusCode(v int32) *SetLoadBalancerStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLoadBalancerStatusResponse) SetBody(v *SetLoadBalancerStatusResponseBody) *SetLoadBalancerStatusResponse {
	s.Body = v
	return s
}

type SetLoadBalancerTCPListenerAttributeRequest struct {
	AclId                     *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclStatus                 *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	AclType                   *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	Bandwidth                 *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ConnectionDrain           *string `json:"ConnectionDrain,omitempty" xml:"ConnectionDrain,omitempty"`
	ConnectionDrainTimeout    *int32  `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EstablishedTimeout        *int32  `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
	HealthCheckConnectPort    *int32  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckConnectTimeout *int32  `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	HealthCheckDomain         *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	HealthCheckHttpCode       *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckInterval       *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckSwitch         *string `json:"HealthCheckSwitch,omitempty" xml:"HealthCheckSwitch,omitempty"`
	HealthCheckType           *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	HealthCheckURI            *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	HealthyThreshold          *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	ListenerPort              *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	LoadBalancerId            *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	MasterSlaveServerGroup    *string `json:"MasterSlaveServerGroup,omitempty" xml:"MasterSlaveServerGroup,omitempty"`
	MasterSlaveServerGroupId  *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PersistenceTimeout        *int32  `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	ProxyProtocolV2Enabled    *bool   `json:"ProxyProtocolV2Enabled,omitempty" xml:"ProxyProtocolV2Enabled,omitempty"`
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Scheduler                 *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	SynProxy                  *string `json:"SynProxy,omitempty" xml:"SynProxy,omitempty"`
	UnhealthyThreshold        *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	VServerGroup              *string `json:"VServerGroup,omitempty" xml:"VServerGroup,omitempty"`
	VServerGroupId            *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s SetLoadBalancerTCPListenerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerTCPListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetAclId(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.AclId = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetAclStatus(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.AclStatus = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetAclType(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.AclType = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetBandwidth(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.Bandwidth = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetConnectionDrain(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.ConnectionDrain = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetConnectionDrainTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetDescription(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.Description = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetEstablishedTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.EstablishedTimeout = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckConnectPort(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckConnectTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckDomain(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckDomain = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckHttpCode(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckInterval(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckSwitch(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckSwitch = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckType(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckType = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckURI(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckURI = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthyThreshold(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetListenerPort(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetLoadBalancerId(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetMasterSlaveServerGroup(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.MasterSlaveServerGroup = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetMasterSlaveServerGroupId(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetOwnerAccount(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetOwnerId(v int64) *SetLoadBalancerTCPListenerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetPersistenceTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.PersistenceTimeout = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetProxyProtocolV2Enabled(v bool) *SetLoadBalancerTCPListenerAttributeRequest {
	s.ProxyProtocolV2Enabled = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetRegionId(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetResourceOwnerAccount(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetResourceOwnerId(v int64) *SetLoadBalancerTCPListenerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetScheduler(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.Scheduler = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetSynProxy(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.SynProxy = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetUnhealthyThreshold(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetVServerGroup(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.VServerGroup = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetVServerGroupId(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.VServerGroupId = &v
	return s
}

type SetLoadBalancerTCPListenerAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLoadBalancerTCPListenerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerTCPListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerTCPListenerAttributeResponseBody) SetRequestId(v string) *SetLoadBalancerTCPListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

type SetLoadBalancerTCPListenerAttributeResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetLoadBalancerTCPListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetLoadBalancerTCPListenerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerTCPListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerTCPListenerAttributeResponse) SetHeaders(v map[string]*string) *SetLoadBalancerTCPListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeResponse) SetStatusCode(v int32) *SetLoadBalancerTCPListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeResponse) SetBody(v *SetLoadBalancerTCPListenerAttributeResponseBody) *SetLoadBalancerTCPListenerAttributeResponse {
	s.Body = v
	return s
}

type SetLoadBalancerUDPListenerAttributeRequest struct {
	AclId                     *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclStatus                 *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	AclType                   *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	Bandwidth                 *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HealthCheckConnectPort    *int32  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckConnectTimeout *int32  `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	HealthCheckInterval       *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckSwitch         *string `json:"HealthCheckSwitch,omitempty" xml:"HealthCheckSwitch,omitempty"`
	HealthyThreshold          *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	ListenerPort              *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	LoadBalancerId            *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	MasterSlaveServerGroup    *string `json:"MasterSlaveServerGroup,omitempty" xml:"MasterSlaveServerGroup,omitempty"`
	MasterSlaveServerGroupId  *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProxyProtocolV2Enabled    *bool   `json:"ProxyProtocolV2Enabled,omitempty" xml:"ProxyProtocolV2Enabled,omitempty"`
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Scheduler                 *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	UnhealthyThreshold        *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	VServerGroup              *string `json:"VServerGroup,omitempty" xml:"VServerGroup,omitempty"`
	VServerGroupId            *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	HealthCheckExp            *string `json:"healthCheckExp,omitempty" xml:"healthCheckExp,omitempty"`
	HealthCheckReq            *string `json:"healthCheckReq,omitempty" xml:"healthCheckReq,omitempty"`
}

func (s SetLoadBalancerUDPListenerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerUDPListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetAclId(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.AclId = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetAclStatus(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.AclStatus = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetAclType(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.AclType = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetBandwidth(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.Bandwidth = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetDescription(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.Description = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetHealthCheckConnectPort(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetHealthCheckConnectTimeout(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetHealthCheckInterval(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetHealthCheckSwitch(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.HealthCheckSwitch = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetHealthyThreshold(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetListenerPort(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetLoadBalancerId(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetMasterSlaveServerGroup(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.MasterSlaveServerGroup = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetMasterSlaveServerGroupId(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetOwnerAccount(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetOwnerId(v int64) *SetLoadBalancerUDPListenerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetProxyProtocolV2Enabled(v bool) *SetLoadBalancerUDPListenerAttributeRequest {
	s.ProxyProtocolV2Enabled = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetRegionId(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetResourceOwnerAccount(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetResourceOwnerId(v int64) *SetLoadBalancerUDPListenerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetScheduler(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.Scheduler = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetUnhealthyThreshold(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetVServerGroup(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.VServerGroup = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetVServerGroupId(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.VServerGroupId = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetHealthCheckExp(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.HealthCheckExp = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetHealthCheckReq(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.HealthCheckReq = &v
	return s
}

type SetLoadBalancerUDPListenerAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLoadBalancerUDPListenerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerUDPListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerUDPListenerAttributeResponseBody) SetRequestId(v string) *SetLoadBalancerUDPListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

type SetLoadBalancerUDPListenerAttributeResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetLoadBalancerUDPListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetLoadBalancerUDPListenerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s SetLoadBalancerUDPListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerUDPListenerAttributeResponse) SetHeaders(v map[string]*string) *SetLoadBalancerUDPListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeResponse) SetStatusCode(v int32) *SetLoadBalancerUDPListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeResponse) SetBody(v *SetLoadBalancerUDPListenerAttributeResponseBody) *SetLoadBalancerUDPListenerAttributeResponse {
	s.Body = v
	return s
}

type SetRuleRequest struct {
	Cookie                 *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	CookieTimeout          *int32  `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	HealthCheck            *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckConnectPort *int32  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckDomain      *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	HealthCheckHttpCode    *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckInterval    *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckTimeout     *int32  `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	HealthCheckURI         *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	HealthyThreshold       *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	ListenerSync           *string `json:"ListenerSync,omitempty" xml:"ListenerSync,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RuleId                 *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName               *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Scheduler              *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	StickySession          *string `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	StickySessionType      *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	UnhealthyThreshold     *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	VServerGroupId         *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s SetRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s SetRuleRequest) GoString() string {
	return s.String()
}

func (s *SetRuleRequest) SetCookie(v string) *SetRuleRequest {
	s.Cookie = &v
	return s
}

func (s *SetRuleRequest) SetCookieTimeout(v int32) *SetRuleRequest {
	s.CookieTimeout = &v
	return s
}

func (s *SetRuleRequest) SetHealthCheck(v string) *SetRuleRequest {
	s.HealthCheck = &v
	return s
}

func (s *SetRuleRequest) SetHealthCheckConnectPort(v int32) *SetRuleRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *SetRuleRequest) SetHealthCheckDomain(v string) *SetRuleRequest {
	s.HealthCheckDomain = &v
	return s
}

func (s *SetRuleRequest) SetHealthCheckHttpCode(v string) *SetRuleRequest {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *SetRuleRequest) SetHealthCheckInterval(v int32) *SetRuleRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *SetRuleRequest) SetHealthCheckTimeout(v int32) *SetRuleRequest {
	s.HealthCheckTimeout = &v
	return s
}

func (s *SetRuleRequest) SetHealthCheckURI(v string) *SetRuleRequest {
	s.HealthCheckURI = &v
	return s
}

func (s *SetRuleRequest) SetHealthyThreshold(v int32) *SetRuleRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *SetRuleRequest) SetListenerSync(v string) *SetRuleRequest {
	s.ListenerSync = &v
	return s
}

func (s *SetRuleRequest) SetOwnerAccount(v string) *SetRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetRuleRequest) SetOwnerId(v int64) *SetRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *SetRuleRequest) SetRegionId(v string) *SetRuleRequest {
	s.RegionId = &v
	return s
}

func (s *SetRuleRequest) SetResourceOwnerAccount(v string) *SetRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetRuleRequest) SetResourceOwnerId(v int64) *SetRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetRuleRequest) SetRuleId(v string) *SetRuleRequest {
	s.RuleId = &v
	return s
}

func (s *SetRuleRequest) SetRuleName(v string) *SetRuleRequest {
	s.RuleName = &v
	return s
}

func (s *SetRuleRequest) SetScheduler(v string) *SetRuleRequest {
	s.Scheduler = &v
	return s
}

func (s *SetRuleRequest) SetStickySession(v string) *SetRuleRequest {
	s.StickySession = &v
	return s
}

func (s *SetRuleRequest) SetStickySessionType(v string) *SetRuleRequest {
	s.StickySessionType = &v
	return s
}

func (s *SetRuleRequest) SetUnhealthyThreshold(v int32) *SetRuleRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *SetRuleRequest) SetVServerGroupId(v string) *SetRuleRequest {
	s.VServerGroupId = &v
	return s
}

type SetRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetRuleResponseBody) GoString() string {
	return s.String()
}

func (s *SetRuleResponseBody) SetRequestId(v string) *SetRuleResponseBody {
	s.RequestId = &v
	return s
}

type SetRuleResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s SetRuleResponse) GoString() string {
	return s.String()
}

func (s *SetRuleResponse) SetHeaders(v map[string]*string) *SetRuleResponse {
	s.Headers = v
	return s
}

func (s *SetRuleResponse) SetStatusCode(v int32) *SetRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *SetRuleResponse) SetBody(v *SetRuleResponseBody) *SetRuleResponse {
	s.Body = v
	return s
}

type SetServerCertificateNameRequest struct {
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ServerCertificateId   *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
	ServerCertificateName *string `json:"ServerCertificateName,omitempty" xml:"ServerCertificateName,omitempty"`
}

func (s SetServerCertificateNameRequest) String() string {
	return tea.Prettify(s)
}

func (s SetServerCertificateNameRequest) GoString() string {
	return s.String()
}

func (s *SetServerCertificateNameRequest) SetOwnerAccount(v string) *SetServerCertificateNameRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetServerCertificateNameRequest) SetOwnerId(v int64) *SetServerCertificateNameRequest {
	s.OwnerId = &v
	return s
}

func (s *SetServerCertificateNameRequest) SetRegionId(v string) *SetServerCertificateNameRequest {
	s.RegionId = &v
	return s
}

func (s *SetServerCertificateNameRequest) SetResourceOwnerAccount(v string) *SetServerCertificateNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetServerCertificateNameRequest) SetResourceOwnerId(v int64) *SetServerCertificateNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetServerCertificateNameRequest) SetServerCertificateId(v string) *SetServerCertificateNameRequest {
	s.ServerCertificateId = &v
	return s
}

func (s *SetServerCertificateNameRequest) SetServerCertificateName(v string) *SetServerCertificateNameRequest {
	s.ServerCertificateName = &v
	return s
}

type SetServerCertificateNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetServerCertificateNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetServerCertificateNameResponseBody) GoString() string {
	return s.String()
}

func (s *SetServerCertificateNameResponseBody) SetRequestId(v string) *SetServerCertificateNameResponseBody {
	s.RequestId = &v
	return s
}

type SetServerCertificateNameResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetServerCertificateNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetServerCertificateNameResponse) String() string {
	return tea.Prettify(s)
}

func (s SetServerCertificateNameResponse) GoString() string {
	return s.String()
}

func (s *SetServerCertificateNameResponse) SetHeaders(v map[string]*string) *SetServerCertificateNameResponse {
	s.Headers = v
	return s
}

func (s *SetServerCertificateNameResponse) SetStatusCode(v int32) *SetServerCertificateNameResponse {
	s.StatusCode = &v
	return s
}

func (s *SetServerCertificateNameResponse) SetBody(v *SetServerCertificateNameResponseBody) *SetServerCertificateNameResponse {
	s.Body = v
	return s
}

type SetTLSCipherPolicyAttributeRequest struct {
	Ciphers              []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	Name                 *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TLSCipherPolicyId    *string   `json:"TLSCipherPolicyId,omitempty" xml:"TLSCipherPolicyId,omitempty"`
	TLSVersions          []*string `json:"TLSVersions,omitempty" xml:"TLSVersions,omitempty" type:"Repeated"`
}

func (s SetTLSCipherPolicyAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s SetTLSCipherPolicyAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetTLSCipherPolicyAttributeRequest) SetCiphers(v []*string) *SetTLSCipherPolicyAttributeRequest {
	s.Ciphers = v
	return s
}

func (s *SetTLSCipherPolicyAttributeRequest) SetName(v string) *SetTLSCipherPolicyAttributeRequest {
	s.Name = &v
	return s
}

func (s *SetTLSCipherPolicyAttributeRequest) SetOwnerAccount(v string) *SetTLSCipherPolicyAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetTLSCipherPolicyAttributeRequest) SetOwnerId(v int64) *SetTLSCipherPolicyAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *SetTLSCipherPolicyAttributeRequest) SetRegionId(v string) *SetTLSCipherPolicyAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *SetTLSCipherPolicyAttributeRequest) SetResourceOwnerAccount(v string) *SetTLSCipherPolicyAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetTLSCipherPolicyAttributeRequest) SetResourceOwnerId(v int64) *SetTLSCipherPolicyAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetTLSCipherPolicyAttributeRequest) SetTLSCipherPolicyId(v string) *SetTLSCipherPolicyAttributeRequest {
	s.TLSCipherPolicyId = &v
	return s
}

func (s *SetTLSCipherPolicyAttributeRequest) SetTLSVersions(v []*string) *SetTLSCipherPolicyAttributeRequest {
	s.TLSVersions = v
	return s
}

type SetTLSCipherPolicyAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SetTLSCipherPolicyAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetTLSCipherPolicyAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *SetTLSCipherPolicyAttributeResponseBody) SetRequestId(v string) *SetTLSCipherPolicyAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetTLSCipherPolicyAttributeResponseBody) SetTaskId(v string) *SetTLSCipherPolicyAttributeResponseBody {
	s.TaskId = &v
	return s
}

type SetTLSCipherPolicyAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetTLSCipherPolicyAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetTLSCipherPolicyAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s SetTLSCipherPolicyAttributeResponse) GoString() string {
	return s.String()
}

func (s *SetTLSCipherPolicyAttributeResponse) SetHeaders(v map[string]*string) *SetTLSCipherPolicyAttributeResponse {
	s.Headers = v
	return s
}

func (s *SetTLSCipherPolicyAttributeResponse) SetStatusCode(v int32) *SetTLSCipherPolicyAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetTLSCipherPolicyAttributeResponse) SetBody(v *SetTLSCipherPolicyAttributeResponseBody) *SetTLSCipherPolicyAttributeResponse {
	s.Body = v
	return s
}

type SetVServerGroupAttributeRequest struct {
	BackendServers       *string `json:"BackendServers,omitempty" xml:"BackendServers,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VServerGroupId       *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	VServerGroupName     *string `json:"VServerGroupName,omitempty" xml:"VServerGroupName,omitempty"`
}

func (s SetVServerGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s SetVServerGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetVServerGroupAttributeRequest) SetBackendServers(v string) *SetVServerGroupAttributeRequest {
	s.BackendServers = &v
	return s
}

func (s *SetVServerGroupAttributeRequest) SetOwnerAccount(v string) *SetVServerGroupAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetVServerGroupAttributeRequest) SetOwnerId(v int64) *SetVServerGroupAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *SetVServerGroupAttributeRequest) SetRegionId(v string) *SetVServerGroupAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *SetVServerGroupAttributeRequest) SetResourceOwnerAccount(v string) *SetVServerGroupAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetVServerGroupAttributeRequest) SetResourceOwnerId(v int64) *SetVServerGroupAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetVServerGroupAttributeRequest) SetVServerGroupId(v string) *SetVServerGroupAttributeRequest {
	s.VServerGroupId = &v
	return s
}

func (s *SetVServerGroupAttributeRequest) SetVServerGroupName(v string) *SetVServerGroupAttributeRequest {
	s.VServerGroupName = &v
	return s
}

type SetVServerGroupAttributeResponseBody struct {
	BackendServers   *SetVServerGroupAttributeResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VServerGroupId   *string                                             `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	VServerGroupName *string                                             `json:"VServerGroupName,omitempty" xml:"VServerGroupName,omitempty"`
}

func (s SetVServerGroupAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetVServerGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *SetVServerGroupAttributeResponseBody) SetBackendServers(v *SetVServerGroupAttributeResponseBodyBackendServers) *SetVServerGroupAttributeResponseBody {
	s.BackendServers = v
	return s
}

func (s *SetVServerGroupAttributeResponseBody) SetRequestId(v string) *SetVServerGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetVServerGroupAttributeResponseBody) SetVServerGroupId(v string) *SetVServerGroupAttributeResponseBody {
	s.VServerGroupId = &v
	return s
}

func (s *SetVServerGroupAttributeResponseBody) SetVServerGroupName(v string) *SetVServerGroupAttributeResponseBody {
	s.VServerGroupName = &v
	return s
}

type SetVServerGroupAttributeResponseBodyBackendServers struct {
	BackendServer []*SetVServerGroupAttributeResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s SetVServerGroupAttributeResponseBodyBackendServers) String() string {
	return tea.Prettify(s)
}

func (s SetVServerGroupAttributeResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *SetVServerGroupAttributeResponseBodyBackendServers) SetBackendServer(v []*SetVServerGroupAttributeResponseBodyBackendServersBackendServer) *SetVServerGroupAttributeResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

type SetVServerGroupAttributeResponseBodyBackendServersBackendServer struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Port        *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ServerId    *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight      *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s SetVServerGroupAttributeResponseBodyBackendServersBackendServer) String() string {
	return tea.Prettify(s)
}

func (s SetVServerGroupAttributeResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *SetVServerGroupAttributeResponseBodyBackendServersBackendServer) SetDescription(v string) *SetVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.Description = &v
	return s
}

func (s *SetVServerGroupAttributeResponseBodyBackendServersBackendServer) SetPort(v int32) *SetVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.Port = &v
	return s
}

func (s *SetVServerGroupAttributeResponseBodyBackendServersBackendServer) SetServerId(v string) *SetVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *SetVServerGroupAttributeResponseBodyBackendServersBackendServer) SetType(v string) *SetVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.Type = &v
	return s
}

func (s *SetVServerGroupAttributeResponseBodyBackendServersBackendServer) SetWeight(v int32) *SetVServerGroupAttributeResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

type SetVServerGroupAttributeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetVServerGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetVServerGroupAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s SetVServerGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *SetVServerGroupAttributeResponse) SetHeaders(v map[string]*string) *SetVServerGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *SetVServerGroupAttributeResponse) SetStatusCode(v int32) *SetVServerGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetVServerGroupAttributeResponse) SetBody(v *SetVServerGroupAttributeResponseBody) *SetVServerGroupAttributeResponse {
	s.Body = v
	return s
}

type StartLoadBalancerListenerRequest struct {
	ListenerPort         *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerProtocol     *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s StartLoadBalancerListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s StartLoadBalancerListenerRequest) GoString() string {
	return s.String()
}

func (s *StartLoadBalancerListenerRequest) SetListenerPort(v int32) *StartLoadBalancerListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *StartLoadBalancerListenerRequest) SetListenerProtocol(v string) *StartLoadBalancerListenerRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *StartLoadBalancerListenerRequest) SetLoadBalancerId(v string) *StartLoadBalancerListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *StartLoadBalancerListenerRequest) SetOwnerAccount(v string) *StartLoadBalancerListenerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StartLoadBalancerListenerRequest) SetOwnerId(v int64) *StartLoadBalancerListenerRequest {
	s.OwnerId = &v
	return s
}

func (s *StartLoadBalancerListenerRequest) SetRegionId(v string) *StartLoadBalancerListenerRequest {
	s.RegionId = &v
	return s
}

func (s *StartLoadBalancerListenerRequest) SetResourceOwnerAccount(v string) *StartLoadBalancerListenerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartLoadBalancerListenerRequest) SetResourceOwnerId(v int64) *StartLoadBalancerListenerRequest {
	s.ResourceOwnerId = &v
	return s
}

type StartLoadBalancerListenerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartLoadBalancerListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartLoadBalancerListenerResponseBody) GoString() string {
	return s.String()
}

func (s *StartLoadBalancerListenerResponseBody) SetRequestId(v string) *StartLoadBalancerListenerResponseBody {
	s.RequestId = &v
	return s
}

type StartLoadBalancerListenerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartLoadBalancerListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartLoadBalancerListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s StartLoadBalancerListenerResponse) GoString() string {
	return s.String()
}

func (s *StartLoadBalancerListenerResponse) SetHeaders(v map[string]*string) *StartLoadBalancerListenerResponse {
	s.Headers = v
	return s
}

func (s *StartLoadBalancerListenerResponse) SetStatusCode(v int32) *StartLoadBalancerListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *StartLoadBalancerListenerResponse) SetBody(v *StartLoadBalancerListenerResponseBody) *StartLoadBalancerListenerResponse {
	s.Body = v
	return s
}

type StopLoadBalancerListenerRequest struct {
	ListenerPort         *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerProtocol     *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s StopLoadBalancerListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s StopLoadBalancerListenerRequest) GoString() string {
	return s.String()
}

func (s *StopLoadBalancerListenerRequest) SetListenerPort(v int32) *StopLoadBalancerListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *StopLoadBalancerListenerRequest) SetListenerProtocol(v string) *StopLoadBalancerListenerRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *StopLoadBalancerListenerRequest) SetLoadBalancerId(v string) *StopLoadBalancerListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *StopLoadBalancerListenerRequest) SetOwnerAccount(v string) *StopLoadBalancerListenerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StopLoadBalancerListenerRequest) SetOwnerId(v int64) *StopLoadBalancerListenerRequest {
	s.OwnerId = &v
	return s
}

func (s *StopLoadBalancerListenerRequest) SetRegionId(v string) *StopLoadBalancerListenerRequest {
	s.RegionId = &v
	return s
}

func (s *StopLoadBalancerListenerRequest) SetResourceOwnerAccount(v string) *StopLoadBalancerListenerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StopLoadBalancerListenerRequest) SetResourceOwnerId(v int64) *StopLoadBalancerListenerRequest {
	s.ResourceOwnerId = &v
	return s
}

type StopLoadBalancerListenerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopLoadBalancerListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopLoadBalancerListenerResponseBody) GoString() string {
	return s.String()
}

func (s *StopLoadBalancerListenerResponseBody) SetRequestId(v string) *StopLoadBalancerListenerResponseBody {
	s.RequestId = &v
	return s
}

type StopLoadBalancerListenerResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopLoadBalancerListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopLoadBalancerListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s StopLoadBalancerListenerResponse) GoString() string {
	return s.String()
}

func (s *StopLoadBalancerListenerResponse) SetHeaders(v map[string]*string) *StopLoadBalancerListenerResponse {
	s.Headers = v
	return s
}

func (s *StopLoadBalancerListenerResponse) SetStatusCode(v int32) *StopLoadBalancerListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *StopLoadBalancerListenerResponse) SetBody(v *StopLoadBalancerListenerResponseBody) *StopLoadBalancerListenerResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerAccount         *string                   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag                  []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetOwnerAccount(v string) *TagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerId(v int64) *TagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	All                  *bool     `json:"All,omitempty" xml:"All,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey               []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerAccount(v string) *UntagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerAccount(v string) *UntagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerId(v int64) *UntagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UploadCACertificateRequest struct {
	CACertificate        *string `json:"CACertificate,omitempty" xml:"CACertificate,omitempty"`
	CACertificateName    *string `json:"CACertificateName,omitempty" xml:"CACertificateName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UploadCACertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadCACertificateRequest) GoString() string {
	return s.String()
}

func (s *UploadCACertificateRequest) SetCACertificate(v string) *UploadCACertificateRequest {
	s.CACertificate = &v
	return s
}

func (s *UploadCACertificateRequest) SetCACertificateName(v string) *UploadCACertificateRequest {
	s.CACertificateName = &v
	return s
}

func (s *UploadCACertificateRequest) SetOwnerAccount(v string) *UploadCACertificateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UploadCACertificateRequest) SetOwnerId(v int64) *UploadCACertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *UploadCACertificateRequest) SetRegionId(v string) *UploadCACertificateRequest {
	s.RegionId = &v
	return s
}

func (s *UploadCACertificateRequest) SetResourceGroupId(v string) *UploadCACertificateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UploadCACertificateRequest) SetResourceOwnerAccount(v string) *UploadCACertificateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UploadCACertificateRequest) SetResourceOwnerId(v int64) *UploadCACertificateRequest {
	s.ResourceOwnerId = &v
	return s
}

type UploadCACertificateResponseBody struct {
	CACertificateId   *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	CACertificateName *string `json:"CACertificateName,omitempty" xml:"CACertificateName,omitempty"`
	CommonName        *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimeStamp   *int64  `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	ExpireTime        *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ExpireTimeStamp   *int64  `json:"ExpireTimeStamp,omitempty" xml:"ExpireTimeStamp,omitempty"`
	Fingerprint       *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId   *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s UploadCACertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadCACertificateResponseBody) GoString() string {
	return s.String()
}

func (s *UploadCACertificateResponseBody) SetCACertificateId(v string) *UploadCACertificateResponseBody {
	s.CACertificateId = &v
	return s
}

func (s *UploadCACertificateResponseBody) SetCACertificateName(v string) *UploadCACertificateResponseBody {
	s.CACertificateName = &v
	return s
}

func (s *UploadCACertificateResponseBody) SetCommonName(v string) *UploadCACertificateResponseBody {
	s.CommonName = &v
	return s
}

func (s *UploadCACertificateResponseBody) SetCreateTime(v string) *UploadCACertificateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *UploadCACertificateResponseBody) SetCreateTimeStamp(v int64) *UploadCACertificateResponseBody {
	s.CreateTimeStamp = &v
	return s
}

func (s *UploadCACertificateResponseBody) SetExpireTime(v string) *UploadCACertificateResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *UploadCACertificateResponseBody) SetExpireTimeStamp(v int64) *UploadCACertificateResponseBody {
	s.ExpireTimeStamp = &v
	return s
}

func (s *UploadCACertificateResponseBody) SetFingerprint(v string) *UploadCACertificateResponseBody {
	s.Fingerprint = &v
	return s
}

func (s *UploadCACertificateResponseBody) SetRequestId(v string) *UploadCACertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadCACertificateResponseBody) SetResourceGroupId(v string) *UploadCACertificateResponseBody {
	s.ResourceGroupId = &v
	return s
}

type UploadCACertificateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadCACertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadCACertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadCACertificateResponse) GoString() string {
	return s.String()
}

func (s *UploadCACertificateResponse) SetHeaders(v map[string]*string) *UploadCACertificateResponse {
	s.Headers = v
	return s
}

func (s *UploadCACertificateResponse) SetStatusCode(v int32) *UploadCACertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadCACertificateResponse) SetBody(v *UploadCACertificateResponseBody) *UploadCACertificateResponse {
	s.Body = v
	return s
}

type UploadServerCertificateRequest struct {
	AliCloudCertificateId       *string `json:"AliCloudCertificateId,omitempty" xml:"AliCloudCertificateId,omitempty"`
	AliCloudCertificateName     *string `json:"AliCloudCertificateName,omitempty" xml:"AliCloudCertificateName,omitempty"`
	AliCloudCertificateRegionId *string `json:"AliCloudCertificateRegionId,omitempty" xml:"AliCloudCertificateRegionId,omitempty"`
	OwnerAccount                *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PrivateKey                  *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	RegionId                    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId             *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount        *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId             *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ServerCertificate           *string `json:"ServerCertificate,omitempty" xml:"ServerCertificate,omitempty"`
	ServerCertificateName       *string `json:"ServerCertificateName,omitempty" xml:"ServerCertificateName,omitempty"`
}

func (s UploadServerCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadServerCertificateRequest) GoString() string {
	return s.String()
}

func (s *UploadServerCertificateRequest) SetAliCloudCertificateId(v string) *UploadServerCertificateRequest {
	s.AliCloudCertificateId = &v
	return s
}

func (s *UploadServerCertificateRequest) SetAliCloudCertificateName(v string) *UploadServerCertificateRequest {
	s.AliCloudCertificateName = &v
	return s
}

func (s *UploadServerCertificateRequest) SetAliCloudCertificateRegionId(v string) *UploadServerCertificateRequest {
	s.AliCloudCertificateRegionId = &v
	return s
}

func (s *UploadServerCertificateRequest) SetOwnerAccount(v string) *UploadServerCertificateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UploadServerCertificateRequest) SetOwnerId(v int64) *UploadServerCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *UploadServerCertificateRequest) SetPrivateKey(v string) *UploadServerCertificateRequest {
	s.PrivateKey = &v
	return s
}

func (s *UploadServerCertificateRequest) SetRegionId(v string) *UploadServerCertificateRequest {
	s.RegionId = &v
	return s
}

func (s *UploadServerCertificateRequest) SetResourceGroupId(v string) *UploadServerCertificateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UploadServerCertificateRequest) SetResourceOwnerAccount(v string) *UploadServerCertificateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UploadServerCertificateRequest) SetResourceOwnerId(v int64) *UploadServerCertificateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UploadServerCertificateRequest) SetServerCertificate(v string) *UploadServerCertificateRequest {
	s.ServerCertificate = &v
	return s
}

func (s *UploadServerCertificateRequest) SetServerCertificateName(v string) *UploadServerCertificateRequest {
	s.ServerCertificateName = &v
	return s
}

type UploadServerCertificateResponseBody struct {
	AliCloudCertificateId   *string                                                     `json:"AliCloudCertificateId,omitempty" xml:"AliCloudCertificateId,omitempty"`
	AliCloudCertificateName *string                                                     `json:"AliCloudCertificateName,omitempty" xml:"AliCloudCertificateName,omitempty"`
	CommonName              *string                                                     `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	CreateTime              *string                                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimeStamp         *int64                                                      `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	ExpireTime              *string                                                     `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ExpireTimeStamp         *int64                                                      `json:"ExpireTimeStamp,omitempty" xml:"ExpireTimeStamp,omitempty"`
	Fingerprint             *string                                                     `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	IsAliCloudCertificate   *int32                                                      `json:"IsAliCloudCertificate,omitempty" xml:"IsAliCloudCertificate,omitempty"`
	RegionId                *string                                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId               *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId         *string                                                     `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServerCertificateId     *string                                                     `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
	ServerCertificateName   *string                                                     `json:"ServerCertificateName,omitempty" xml:"ServerCertificateName,omitempty"`
	SubjectAlternativeNames *UploadServerCertificateResponseBodySubjectAlternativeNames `json:"SubjectAlternativeNames,omitempty" xml:"SubjectAlternativeNames,omitempty" type:"Struct"`
}

func (s UploadServerCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadServerCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *UploadServerCertificateResponseBody) SetAliCloudCertificateId(v string) *UploadServerCertificateResponseBody {
	s.AliCloudCertificateId = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetAliCloudCertificateName(v string) *UploadServerCertificateResponseBody {
	s.AliCloudCertificateName = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetCommonName(v string) *UploadServerCertificateResponseBody {
	s.CommonName = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetCreateTime(v string) *UploadServerCertificateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetCreateTimeStamp(v int64) *UploadServerCertificateResponseBody {
	s.CreateTimeStamp = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetExpireTime(v string) *UploadServerCertificateResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetExpireTimeStamp(v int64) *UploadServerCertificateResponseBody {
	s.ExpireTimeStamp = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetFingerprint(v string) *UploadServerCertificateResponseBody {
	s.Fingerprint = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetIsAliCloudCertificate(v int32) *UploadServerCertificateResponseBody {
	s.IsAliCloudCertificate = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetRegionId(v string) *UploadServerCertificateResponseBody {
	s.RegionId = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetRequestId(v string) *UploadServerCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetResourceGroupId(v string) *UploadServerCertificateResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetServerCertificateId(v string) *UploadServerCertificateResponseBody {
	s.ServerCertificateId = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetServerCertificateName(v string) *UploadServerCertificateResponseBody {
	s.ServerCertificateName = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetSubjectAlternativeNames(v *UploadServerCertificateResponseBodySubjectAlternativeNames) *UploadServerCertificateResponseBody {
	s.SubjectAlternativeNames = v
	return s
}

type UploadServerCertificateResponseBodySubjectAlternativeNames struct {
	SubjectAlternativeName []*string `json:"SubjectAlternativeName,omitempty" xml:"SubjectAlternativeName,omitempty" type:"Repeated"`
}

func (s UploadServerCertificateResponseBodySubjectAlternativeNames) String() string {
	return tea.Prettify(s)
}

func (s UploadServerCertificateResponseBodySubjectAlternativeNames) GoString() string {
	return s.String()
}

func (s *UploadServerCertificateResponseBodySubjectAlternativeNames) SetSubjectAlternativeName(v []*string) *UploadServerCertificateResponseBodySubjectAlternativeNames {
	s.SubjectAlternativeName = v
	return s
}

type UploadServerCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadServerCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadServerCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadServerCertificateResponse) GoString() string {
	return s.String()
}

func (s *UploadServerCertificateResponse) SetHeaders(v map[string]*string) *UploadServerCertificateResponse {
	s.Headers = v
	return s
}

func (s *UploadServerCertificateResponse) SetStatusCode(v int32) *UploadServerCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadServerCertificateResponse) SetBody(v *UploadServerCertificateResponseBody) *UploadServerCertificateResponse {
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
		"cn-qingdao":                  tea.String("slb.aliyuncs.com"),
		"cn-beijing":                  tea.String("slb.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("slb.aliyuncs.com"),
		"cn-shanghai":                 tea.String("slb.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("slb.aliyuncs.com"),
		"cn-hongkong":                 tea.String("slb.aliyuncs.com"),
		"ap-southeast-1":              tea.String("slb.aliyuncs.com"),
		"us-east-1":                   tea.String("slb.aliyuncs.com"),
		"us-west-1":                   tea.String("slb.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("slb.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("slb.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("slb.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("slb.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("slb.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("slb.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("slb.aliyuncs.com"),
		"cn-edge-1":                   tea.String("slb.aliyuncs.com"),
		"cn-fujian":                   tea.String("slb.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("slb.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("slb.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("slb.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("slb.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("slb.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("slb.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("slb.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("slb.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("slb.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("slb-api.cn-qingdao-nebula.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("slb.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("slb.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("slb.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("slb.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("slb.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("slb.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("slb.aliyuncs.com"),
		"cn-wuhan":                    tea.String("slb.aliyuncs.com"),
		"cn-yushanfang":               tea.String("slb.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("slb.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("slb.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("slb.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("slb.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("slb.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("slb.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("slb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddAccessControlListEntryWithOptions(request *AddAccessControlListEntryRequest, runtime *util.RuntimeOptions) (_result *AddAccessControlListEntryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclEntrys)) {
		query["AclEntrys"] = request.AclEntrys
	}

	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddAccessControlListEntry"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddAccessControlListEntryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddAccessControlListEntry(request *AddAccessControlListEntryRequest) (_result *AddAccessControlListEntryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddAccessControlListEntryResponse{}
	_body, _err := client.AddAccessControlListEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddBackendServersWithOptions(request *AddBackendServersRequest, runtime *util.RuntimeOptions) (_result *AddBackendServersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackendServers)) {
		query["BackendServers"] = request.BackendServers
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddBackendServers"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddBackendServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddBackendServers(request *AddBackendServersRequest) (_result *AddBackendServersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddBackendServersResponse{}
	_body, _err := client.AddBackendServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddListenerWhiteListItemWithOptions(request *AddListenerWhiteListItemRequest, runtime *util.RuntimeOptions) (_result *AddListenerWhiteListItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerProtocol)) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceItems)) {
		query["SourceItems"] = request.SourceItems
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddListenerWhiteListItem"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddListenerWhiteListItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddListenerWhiteListItem(request *AddListenerWhiteListItemRequest) (_result *AddListenerWhiteListItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddListenerWhiteListItemResponse{}
	_body, _err := client.AddListenerWhiteListItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddTagsWithOptions(request *AddTagsRequest, runtime *util.RuntimeOptions) (_result *AddTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddTags"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddTags(request *AddTagsRequest) (_result *AddTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddTagsResponse{}
	_body, _err := client.AddTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddVServerGroupBackendServersWithOptions(request *AddVServerGroupBackendServersRequest, runtime *util.RuntimeOptions) (_result *AddVServerGroupBackendServersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackendServers)) {
		query["BackendServers"] = request.BackendServers
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroupId)) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddVServerGroupBackendServers"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddVServerGroupBackendServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddVServerGroupBackendServers(request *AddVServerGroupBackendServersRequest) (_result *AddVServerGroupBackendServersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddVServerGroupBackendServersResponse{}
	_body, _err := client.AddVServerGroupBackendServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAccessControlListWithOptions(request *CreateAccessControlListRequest, runtime *util.RuntimeOptions) (_result *CreateAccessControlListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclName)) {
		query["AclName"] = request.AclName
	}

	if !tea.BoolValue(util.IsUnset(request.AddressIPVersion)) {
		query["AddressIPVersion"] = request.AddressIPVersion
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccessControlList"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccessControlListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccessControlList(request *CreateAccessControlListRequest) (_result *CreateAccessControlListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccessControlListResponse{}
	_body, _err := client.CreateAccessControlListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDomainExtensionWithOptions(request *CreateDomainExtensionRequest, runtime *util.RuntimeOptions) (_result *CreateDomainExtensionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerCertificateId)) {
		query["ServerCertificateId"] = request.ServerCertificateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDomainExtension"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDomainExtensionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDomainExtension(request *CreateDomainExtensionRequest) (_result *CreateDomainExtensionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDomainExtensionResponse{}
	_body, _err := client.CreateDomainExtensionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLoadBalancerWithOptions(request *CreateLoadBalancerRequest, runtime *util.RuntimeOptions) (_result *CreateLoadBalancerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.AddressIPVersion)) {
		query["AddressIPVersion"] = request.AddressIPVersion
	}

	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		query["AddressType"] = request.AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DeleteProtection)) {
		query["DeleteProtection"] = request.DeleteProtection
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceChargeType)) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.InternetChargeType)) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerName)) {
		query["LoadBalancerName"] = request.LoadBalancerName
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerSpec)) {
		query["LoadBalancerSpec"] = request.LoadBalancerSpec
	}

	if !tea.BoolValue(util.IsUnset(request.MasterZoneId)) {
		query["MasterZoneId"] = request.MasterZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.ModificationProtectionReason)) {
		query["ModificationProtectionReason"] = request.ModificationProtectionReason
	}

	if !tea.BoolValue(util.IsUnset(request.ModificationProtectionStatus)) {
		query["ModificationProtectionStatus"] = request.ModificationProtectionStatus
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SlaveZoneId)) {
		query["SlaveZoneId"] = request.SlaveZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLoadBalancer"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLoadBalancerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (_result *CreateLoadBalancerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLoadBalancerResponse{}
	_body, _err := client.CreateLoadBalancerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLoadBalancerHTTPListenerWithOptions(request *CreateLoadBalancerHTTPListenerRequest, runtime *util.RuntimeOptions) (_result *CreateLoadBalancerHTTPListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.AclStatus)) {
		query["AclStatus"] = request.AclStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AclType)) {
		query["AclType"] = request.AclType
	}

	if !tea.BoolValue(util.IsUnset(request.BackendServerPort)) {
		query["BackendServerPort"] = request.BackendServerPort
	}

	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.Cookie)) {
		query["Cookie"] = request.Cookie
	}

	if !tea.BoolValue(util.IsUnset(request.CookieTimeout)) {
		query["CookieTimeout"] = request.CookieTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ForwardPort)) {
		query["ForwardPort"] = request.ForwardPort
	}

	if !tea.BoolValue(util.IsUnset(request.Gzip)) {
		query["Gzip"] = request.Gzip
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheck)) {
		query["HealthCheck"] = request.HealthCheck
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckConnectPort)) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckDomain)) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckHttpCode)) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckInterval)) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckMethod)) {
		query["HealthCheckMethod"] = request.HealthCheckMethod
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckTimeout)) {
		query["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckURI)) {
		query["HealthCheckURI"] = request.HealthCheckURI
	}

	if !tea.BoolValue(util.IsUnset(request.HealthyThreshold)) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.IdleTimeout)) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerForward)) {
		query["ListenerForward"] = request.ListenerForward
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestTimeout)) {
		query["RequestTimeout"] = request.RequestTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Scheduler)) {
		query["Scheduler"] = request.Scheduler
	}

	if !tea.BoolValue(util.IsUnset(request.StickySession)) {
		query["StickySession"] = request.StickySession
	}

	if !tea.BoolValue(util.IsUnset(request.StickySessionType)) {
		query["StickySessionType"] = request.StickySessionType
	}

	if !tea.BoolValue(util.IsUnset(request.UnhealthyThreshold)) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroupId)) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedFor)) {
		query["XForwardedFor"] = request.XForwardedFor
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedFor_SLBID)) {
		query["XForwardedFor_SLBID"] = request.XForwardedFor_SLBID
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedFor_SLBIP)) {
		query["XForwardedFor_SLBIP"] = request.XForwardedFor_SLBIP
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedFor_proto)) {
		query["XForwardedFor_proto"] = request.XForwardedFor_proto
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLoadBalancerHTTPListener"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLoadBalancerHTTPListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLoadBalancerHTTPListener(request *CreateLoadBalancerHTTPListenerRequest) (_result *CreateLoadBalancerHTTPListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLoadBalancerHTTPListenerResponse{}
	_body, _err := client.CreateLoadBalancerHTTPListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLoadBalancerHTTPSListenerWithOptions(request *CreateLoadBalancerHTTPSListenerRequest, runtime *util.RuntimeOptions) (_result *CreateLoadBalancerHTTPSListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.AclStatus)) {
		query["AclStatus"] = request.AclStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AclType)) {
		query["AclType"] = request.AclType
	}

	if !tea.BoolValue(util.IsUnset(request.BackendServerPort)) {
		query["BackendServerPort"] = request.BackendServerPort
	}

	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.CACertificateId)) {
		query["CACertificateId"] = request.CACertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.Cookie)) {
		query["Cookie"] = request.Cookie
	}

	if !tea.BoolValue(util.IsUnset(request.CookieTimeout)) {
		query["CookieTimeout"] = request.CookieTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EnableHttp2)) {
		query["EnableHttp2"] = request.EnableHttp2
	}

	if !tea.BoolValue(util.IsUnset(request.Gzip)) {
		query["Gzip"] = request.Gzip
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheck)) {
		query["HealthCheck"] = request.HealthCheck
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckConnectPort)) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckDomain)) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckHttpCode)) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckInterval)) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckMethod)) {
		query["HealthCheckMethod"] = request.HealthCheckMethod
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckTimeout)) {
		query["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckURI)) {
		query["HealthCheckURI"] = request.HealthCheckURI
	}

	if !tea.BoolValue(util.IsUnset(request.HealthyThreshold)) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.IdleTimeout)) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestTimeout)) {
		query["RequestTimeout"] = request.RequestTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Scheduler)) {
		query["Scheduler"] = request.Scheduler
	}

	if !tea.BoolValue(util.IsUnset(request.ServerCertificateId)) {
		query["ServerCertificateId"] = request.ServerCertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.StickySession)) {
		query["StickySession"] = request.StickySession
	}

	if !tea.BoolValue(util.IsUnset(request.StickySessionType)) {
		query["StickySessionType"] = request.StickySessionType
	}

	if !tea.BoolValue(util.IsUnset(request.TLSCipherPolicy)) {
		query["TLSCipherPolicy"] = request.TLSCipherPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.UnhealthyThreshold)) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroupId)) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedFor)) {
		query["XForwardedFor"] = request.XForwardedFor
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedFor_SLBID)) {
		query["XForwardedFor_SLBID"] = request.XForwardedFor_SLBID
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedFor_SLBIP)) {
		query["XForwardedFor_SLBIP"] = request.XForwardedFor_SLBIP
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedFor_proto)) {
		query["XForwardedFor_proto"] = request.XForwardedFor_proto
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLoadBalancerHTTPSListener"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLoadBalancerHTTPSListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLoadBalancerHTTPSListener(request *CreateLoadBalancerHTTPSListenerRequest) (_result *CreateLoadBalancerHTTPSListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLoadBalancerHTTPSListenerResponse{}
	_body, _err := client.CreateLoadBalancerHTTPSListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLoadBalancerTCPListenerWithOptions(request *CreateLoadBalancerTCPListenerRequest, runtime *util.RuntimeOptions) (_result *CreateLoadBalancerTCPListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.AclStatus)) {
		query["AclStatus"] = request.AclStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AclType)) {
		query["AclType"] = request.AclType
	}

	if !tea.BoolValue(util.IsUnset(request.BackendServerPort)) {
		query["BackendServerPort"] = request.BackendServerPort
	}

	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionDrain)) {
		query["ConnectionDrain"] = request.ConnectionDrain
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionDrainTimeout)) {
		query["ConnectionDrainTimeout"] = request.ConnectionDrainTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EstablishedTimeout)) {
		query["EstablishedTimeout"] = request.EstablishedTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckConnectPort)) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckConnectTimeout)) {
		query["HealthCheckConnectTimeout"] = request.HealthCheckConnectTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckDomain)) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckHttpCode)) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckSwitch)) {
		query["HealthCheckSwitch"] = request.HealthCheckSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckType)) {
		query["HealthCheckType"] = request.HealthCheckType
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckURI)) {
		query["HealthCheckURI"] = request.HealthCheckURI
	}

	if !tea.BoolValue(util.IsUnset(request.HealthyThreshold)) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.MasterSlaveServerGroupId)) {
		query["MasterSlaveServerGroupId"] = request.MasterSlaveServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PersistenceTimeout)) {
		query["PersistenceTimeout"] = request.PersistenceTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyProtocolV2Enabled)) {
		query["ProxyProtocolV2Enabled"] = request.ProxyProtocolV2Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Scheduler)) {
		query["Scheduler"] = request.Scheduler
	}

	if !tea.BoolValue(util.IsUnset(request.UnhealthyThreshold)) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroupId)) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckInterval)) {
		query["healthCheckInterval"] = request.HealthCheckInterval
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLoadBalancerTCPListener"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLoadBalancerTCPListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLoadBalancerTCPListener(request *CreateLoadBalancerTCPListenerRequest) (_result *CreateLoadBalancerTCPListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLoadBalancerTCPListenerResponse{}
	_body, _err := client.CreateLoadBalancerTCPListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLoadBalancerUDPListenerWithOptions(request *CreateLoadBalancerUDPListenerRequest, runtime *util.RuntimeOptions) (_result *CreateLoadBalancerUDPListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.AclStatus)) {
		query["AclStatus"] = request.AclStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AclType)) {
		query["AclType"] = request.AclType
	}

	if !tea.BoolValue(util.IsUnset(request.BackendServerPort)) {
		query["BackendServerPort"] = request.BackendServerPort
	}

	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckConnectPort)) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckConnectTimeout)) {
		query["HealthCheckConnectTimeout"] = request.HealthCheckConnectTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckSwitch)) {
		query["HealthCheckSwitch"] = request.HealthCheckSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.HealthyThreshold)) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.MasterSlaveServerGroupId)) {
		query["MasterSlaveServerGroupId"] = request.MasterSlaveServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyProtocolV2Enabled)) {
		query["ProxyProtocolV2Enabled"] = request.ProxyProtocolV2Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Scheduler)) {
		query["Scheduler"] = request.Scheduler
	}

	if !tea.BoolValue(util.IsUnset(request.UnhealthyThreshold)) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroupId)) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckExp)) {
		query["healthCheckExp"] = request.HealthCheckExp
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckInterval)) {
		query["healthCheckInterval"] = request.HealthCheckInterval
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckReq)) {
		query["healthCheckReq"] = request.HealthCheckReq
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLoadBalancerUDPListener"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLoadBalancerUDPListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLoadBalancerUDPListener(request *CreateLoadBalancerUDPListenerRequest) (_result *CreateLoadBalancerUDPListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLoadBalancerUDPListenerResponse{}
	_body, _err := client.CreateLoadBalancerUDPListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMasterSlaveServerGroupWithOptions(request *CreateMasterSlaveServerGroupRequest, runtime *util.RuntimeOptions) (_result *CreateMasterSlaveServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.MasterSlaveBackendServers)) {
		query["MasterSlaveBackendServers"] = request.MasterSlaveBackendServers
	}

	if !tea.BoolValue(util.IsUnset(request.MasterSlaveServerGroupName)) {
		query["MasterSlaveServerGroupName"] = request.MasterSlaveServerGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMasterSlaveServerGroup"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMasterSlaveServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMasterSlaveServerGroup(request *CreateMasterSlaveServerGroupRequest) (_result *CreateMasterSlaveServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMasterSlaveServerGroupResponse{}
	_body, _err := client.CreateMasterSlaveServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRulesWithOptions(request *CreateRulesRequest, runtime *util.RuntimeOptions) (_result *CreateRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerProtocol)) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleList)) {
		query["RuleList"] = request.RuleList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRules"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRules(request *CreateRulesRequest) (_result *CreateRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRulesResponse{}
	_body, _err := client.CreateRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTLSCipherPolicyWithOptions(request *CreateTLSCipherPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateTLSCipherPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ciphers)) {
		query["Ciphers"] = request.Ciphers
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TLSVersions)) {
		query["TLSVersions"] = request.TLSVersions
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTLSCipherPolicy"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTLSCipherPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTLSCipherPolicy(request *CreateTLSCipherPolicyRequest) (_result *CreateTLSCipherPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTLSCipherPolicyResponse{}
	_body, _err := client.CreateTLSCipherPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVServerGroupWithOptions(request *CreateVServerGroupRequest, runtime *util.RuntimeOptions) (_result *CreateVServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackendServers)) {
		query["BackendServers"] = request.BackendServers
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroupName)) {
		query["VServerGroupName"] = request.VServerGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVServerGroup"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVServerGroup(request *CreateVServerGroupRequest) (_result *CreateVServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVServerGroupResponse{}
	_body, _err := client.CreateVServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAccessControlListWithOptions(request *DeleteAccessControlListRequest, runtime *util.RuntimeOptions) (_result *DeleteAccessControlListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccessControlList"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccessControlListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAccessControlList(request *DeleteAccessControlListRequest) (_result *DeleteAccessControlListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccessControlListResponse{}
	_body, _err := client.DeleteAccessControlListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCACertificateWithOptions(request *DeleteCACertificateRequest, runtime *util.RuntimeOptions) (_result *DeleteCACertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CACertificateId)) {
		query["CACertificateId"] = request.CACertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCACertificate"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCACertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCACertificate(request *DeleteCACertificateRequest) (_result *DeleteCACertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCACertificateResponse{}
	_body, _err := client.DeleteCACertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDomainExtensionWithOptions(request *DeleteDomainExtensionRequest, runtime *util.RuntimeOptions) (_result *DeleteDomainExtensionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainExtensionId)) {
		query["DomainExtensionId"] = request.DomainExtensionId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomainExtension"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDomainExtensionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDomainExtension(request *DeleteDomainExtensionRequest) (_result *DeleteDomainExtensionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDomainExtensionResponse{}
	_body, _err := client.DeleteDomainExtensionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLoadBalancerWithOptions(request *DeleteLoadBalancerRequest, runtime *util.RuntimeOptions) (_result *DeleteLoadBalancerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLoadBalancer"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLoadBalancerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLoadBalancer(request *DeleteLoadBalancerRequest) (_result *DeleteLoadBalancerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLoadBalancerResponse{}
	_body, _err := client.DeleteLoadBalancerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLoadBalancerListenerWithOptions(request *DeleteLoadBalancerListenerRequest, runtime *util.RuntimeOptions) (_result *DeleteLoadBalancerListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerProtocol)) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLoadBalancerListener"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLoadBalancerListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLoadBalancerListener(request *DeleteLoadBalancerListenerRequest) (_result *DeleteLoadBalancerListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLoadBalancerListenerResponse{}
	_body, _err := client.DeleteLoadBalancerListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMasterSlaveServerGroupWithOptions(request *DeleteMasterSlaveServerGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteMasterSlaveServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MasterSlaveServerGroupId)) {
		query["MasterSlaveServerGroupId"] = request.MasterSlaveServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMasterSlaveServerGroup"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMasterSlaveServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMasterSlaveServerGroup(request *DeleteMasterSlaveServerGroupRequest) (_result *DeleteMasterSlaveServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMasterSlaveServerGroupResponse{}
	_body, _err := client.DeleteMasterSlaveServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRulesWithOptions(request *DeleteRulesRequest, runtime *util.RuntimeOptions) (_result *DeleteRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleIds)) {
		query["RuleIds"] = request.RuleIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRules"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRules(request *DeleteRulesRequest) (_result *DeleteRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRulesResponse{}
	_body, _err := client.DeleteRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServerCertificateWithOptions(request *DeleteServerCertificateRequest, runtime *util.RuntimeOptions) (_result *DeleteServerCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerCertificateId)) {
		query["ServerCertificateId"] = request.ServerCertificateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServerCertificate"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServerCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteServerCertificate(request *DeleteServerCertificateRequest) (_result *DeleteServerCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServerCertificateResponse{}
	_body, _err := client.DeleteServerCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTLSCipherPolicyWithOptions(request *DeleteTLSCipherPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteTLSCipherPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TLSCipherPolicyId)) {
		query["TLSCipherPolicyId"] = request.TLSCipherPolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTLSCipherPolicy"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTLSCipherPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTLSCipherPolicy(request *DeleteTLSCipherPolicyRequest) (_result *DeleteTLSCipherPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTLSCipherPolicyResponse{}
	_body, _err := client.DeleteTLSCipherPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVServerGroupWithOptions(request *DeleteVServerGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteVServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroupId)) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVServerGroup"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVServerGroup(request *DeleteVServerGroupRequest) (_result *DeleteVServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVServerGroupResponse{}
	_body, _err := client.DeleteVServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccessControlListAttributeWithOptions(request *DescribeAccessControlListAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeAccessControlListAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclEntryComment)) {
		query["AclEntryComment"] = request.AclEntryComment
	}

	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccessControlListAttribute"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccessControlListAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccessControlListAttribute(request *DescribeAccessControlListAttributeRequest) (_result *DescribeAccessControlListAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccessControlListAttributeResponse{}
	_body, _err := client.DescribeAccessControlListAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccessControlListsWithOptions(request *DescribeAccessControlListsRequest, runtime *util.RuntimeOptions) (_result *DescribeAccessControlListsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclName)) {
		query["AclName"] = request.AclName
	}

	if !tea.BoolValue(util.IsUnset(request.AddressIPVersion)) {
		query["AddressIPVersion"] = request.AddressIPVersion
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccessControlLists"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccessControlListsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccessControlLists(request *DescribeAccessControlListsRequest) (_result *DescribeAccessControlListsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccessControlListsResponse{}
	_body, _err := client.DescribeAccessControlListsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccessLogsDownloadAttributeWithOptions(request *DescribeAccessLogsDownloadAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeAccessLogsDownloadAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.LogType)) {
		query["LogType"] = request.LogType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccessLogsDownloadAttribute"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccessLogsDownloadAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccessLogsDownloadAttribute(request *DescribeAccessLogsDownloadAttributeRequest) (_result *DescribeAccessLogsDownloadAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccessLogsDownloadAttributeResponse{}
	_body, _err := client.DescribeAccessLogsDownloadAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAvailableResourceWithOptions(request *DescribeAvailableResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressIPVersion)) {
		query["AddressIPVersion"] = request.AddressIPVersion
	}

	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		query["AddressType"] = request.AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAvailableResource"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableResource(request *DescribeAvailableResourceRequest) (_result *DescribeAvailableResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.DescribeAvailableResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCACertificatesWithOptions(request *DescribeCACertificatesRequest, runtime *util.RuntimeOptions) (_result *DescribeCACertificatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CACertificateId)) {
		query["CACertificateId"] = request.CACertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCACertificates"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCACertificatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCACertificates(request *DescribeCACertificatesRequest) (_result *DescribeCACertificatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCACertificatesResponse{}
	_body, _err := client.DescribeCACertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainExtensionAttributeWithOptions(request *DescribeDomainExtensionAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainExtensionAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainExtensionId)) {
		query["DomainExtensionId"] = request.DomainExtensionId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainExtensionAttribute"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainExtensionAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainExtensionAttribute(request *DescribeDomainExtensionAttributeRequest) (_result *DescribeDomainExtensionAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainExtensionAttributeResponse{}
	_body, _err := client.DescribeDomainExtensionAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainExtensionsWithOptions(request *DescribeDomainExtensionsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainExtensionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainExtensionId)) {
		query["DomainExtensionId"] = request.DomainExtensionId
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainExtensions"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainExtensionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainExtensions(request *DescribeDomainExtensionsRequest) (_result *DescribeDomainExtensionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainExtensionsResponse{}
	_body, _err := client.DescribeDomainExtensionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHealthStatusWithOptions(request *DescribeHealthStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeHealthStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerProtocol)) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHealthStatus"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHealthStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHealthStatus(request *DescribeHealthStatusRequest) (_result *DescribeHealthStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHealthStatusResponse{}
	_body, _err := client.DescribeHealthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeListenerAccessControlAttributeWithOptions(request *DescribeListenerAccessControlAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeListenerAccessControlAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerProtocol)) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeListenerAccessControlAttribute"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeListenerAccessControlAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeListenerAccessControlAttribute(request *DescribeListenerAccessControlAttributeRequest) (_result *DescribeListenerAccessControlAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeListenerAccessControlAttributeResponse{}
	_body, _err := client.DescribeListenerAccessControlAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLoadBalancerAttributeWithOptions(request *DescribeLoadBalancerAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeLoadBalancerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLoadBalancerAttribute"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLoadBalancerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLoadBalancerAttribute(request *DescribeLoadBalancerAttributeRequest) (_result *DescribeLoadBalancerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLoadBalancerAttributeResponse{}
	_body, _err := client.DescribeLoadBalancerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLoadBalancerHTTPListenerAttributeWithOptions(request *DescribeLoadBalancerHTTPListenerAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeLoadBalancerHTTPListenerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLoadBalancerHTTPListenerAttribute"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLoadBalancerHTTPListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLoadBalancerHTTPListenerAttribute(request *DescribeLoadBalancerHTTPListenerAttributeRequest) (_result *DescribeLoadBalancerHTTPListenerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLoadBalancerHTTPListenerAttributeResponse{}
	_body, _err := client.DescribeLoadBalancerHTTPListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLoadBalancerHTTPSListenerAttributeWithOptions(request *DescribeLoadBalancerHTTPSListenerAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeLoadBalancerHTTPSListenerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLoadBalancerHTTPSListenerAttribute"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLoadBalancerHTTPSListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLoadBalancerHTTPSListenerAttribute(request *DescribeLoadBalancerHTTPSListenerAttributeRequest) (_result *DescribeLoadBalancerHTTPSListenerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLoadBalancerHTTPSListenerAttributeResponse{}
	_body, _err := client.DescribeLoadBalancerHTTPSListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLoadBalancerListenersWithOptions(request *DescribeLoadBalancerListenersRequest, runtime *util.RuntimeOptions) (_result *DescribeLoadBalancerListenersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerProtocol)) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLoadBalancerListeners"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLoadBalancerListenersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLoadBalancerListeners(request *DescribeLoadBalancerListenersRequest) (_result *DescribeLoadBalancerListenersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLoadBalancerListenersResponse{}
	_body, _err := client.DescribeLoadBalancerListenersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLoadBalancerTCPListenerAttributeWithOptions(request *DescribeLoadBalancerTCPListenerAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeLoadBalancerTCPListenerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLoadBalancerTCPListenerAttribute"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLoadBalancerTCPListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLoadBalancerTCPListenerAttribute(request *DescribeLoadBalancerTCPListenerAttributeRequest) (_result *DescribeLoadBalancerTCPListenerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLoadBalancerTCPListenerAttributeResponse{}
	_body, _err := client.DescribeLoadBalancerTCPListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLoadBalancerUDPListenerAttributeWithOptions(request *DescribeLoadBalancerUDPListenerAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeLoadBalancerUDPListenerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLoadBalancerUDPListenerAttribute"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLoadBalancerUDPListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLoadBalancerUDPListenerAttribute(request *DescribeLoadBalancerUDPListenerAttributeRequest) (_result *DescribeLoadBalancerUDPListenerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLoadBalancerUDPListenerAttributeResponse{}
	_body, _err := client.DescribeLoadBalancerUDPListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLoadBalancersWithOptions(request *DescribeLoadBalancersRequest, runtime *util.RuntimeOptions) (_result *DescribeLoadBalancersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.AddressIPVersion)) {
		query["AddressIPVersion"] = request.AddressIPVersion
	}

	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		query["AddressType"] = request.AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.InternetChargeType)) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerName)) {
		query["LoadBalancerName"] = request.LoadBalancerName
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerStatus)) {
		query["LoadBalancerStatus"] = request.LoadBalancerStatus
	}

	if !tea.BoolValue(util.IsUnset(request.MasterZoneId)) {
		query["MasterZoneId"] = request.MasterZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		query["NetworkType"] = request.NetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerId)) {
		query["ServerId"] = request.ServerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerIntranetAddress)) {
		query["ServerIntranetAddress"] = request.ServerIntranetAddress
	}

	if !tea.BoolValue(util.IsUnset(request.SlaveZoneId)) {
		query["SlaveZoneId"] = request.SlaveZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLoadBalancers"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLoadBalancersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLoadBalancers(request *DescribeLoadBalancersRequest) (_result *DescribeLoadBalancersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLoadBalancersResponse{}
	_body, _err := client.DescribeLoadBalancersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMasterSlaveServerGroupAttributeWithOptions(request *DescribeMasterSlaveServerGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeMasterSlaveServerGroupAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MasterSlaveServerGroupId)) {
		query["MasterSlaveServerGroupId"] = request.MasterSlaveServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMasterSlaveServerGroupAttribute"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMasterSlaveServerGroupAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMasterSlaveServerGroupAttribute(request *DescribeMasterSlaveServerGroupAttributeRequest) (_result *DescribeMasterSlaveServerGroupAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMasterSlaveServerGroupAttributeResponse{}
	_body, _err := client.DescribeMasterSlaveServerGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMasterSlaveServerGroupsWithOptions(request *DescribeMasterSlaveServerGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeMasterSlaveServerGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncludeListener)) {
		query["IncludeListener"] = request.IncludeListener
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMasterSlaveServerGroups"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMasterSlaveServerGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMasterSlaveServerGroups(request *DescribeMasterSlaveServerGroupsRequest) (_result *DescribeMasterSlaveServerGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMasterSlaveServerGroupsResponse{}
	_body, _err := client.DescribeMasterSlaveServerGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRuleAttributeWithOptions(request *DescribeRuleAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeRuleAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRuleAttribute"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRuleAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRuleAttribute(request *DescribeRuleAttributeRequest) (_result *DescribeRuleAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRuleAttributeResponse{}
	_body, _err := client.DescribeRuleAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRulesWithOptions(request *DescribeRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerProtocol)) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRules"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRules(request *DescribeRulesRequest) (_result *DescribeRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRulesResponse{}
	_body, _err := client.DescribeRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServerCertificatesWithOptions(request *DescribeServerCertificatesRequest, runtime *util.RuntimeOptions) (_result *DescribeServerCertificatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerCertificateId)) {
		query["ServerCertificateId"] = request.ServerCertificateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServerCertificates"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServerCertificatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServerCertificates(request *DescribeServerCertificatesRequest) (_result *DescribeServerCertificatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServerCertificatesResponse{}
	_body, _err := client.DescribeServerCertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTagsWithOptions(request *DescribeTagsRequest, runtime *util.RuntimeOptions) (_result *DescribeTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DistinctKey)) {
		query["DistinctKey"] = request.DistinctKey
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTags"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTags(request *DescribeTagsRequest) (_result *DescribeTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTagsResponse{}
	_body, _err := client.DescribeTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVServerGroupAttributeWithOptions(request *DescribeVServerGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeVServerGroupAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroupId)) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVServerGroupAttribute"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVServerGroupAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVServerGroupAttribute(request *DescribeVServerGroupAttributeRequest) (_result *DescribeVServerGroupAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVServerGroupAttributeResponse{}
	_body, _err := client.DescribeVServerGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVServerGroupsWithOptions(request *DescribeVServerGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeVServerGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncludeListener)) {
		query["IncludeListener"] = request.IncludeListener
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeRule)) {
		query["IncludeRule"] = request.IncludeRule
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVServerGroups"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVServerGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVServerGroups(request *DescribeVServerGroupsRequest) (_result *DescribeVServerGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVServerGroupsResponse{}
	_body, _err := client.DescribeVServerGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeZones"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTLSCipherPoliciesWithOptions(request *ListTLSCipherPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListTLSCipherPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncludeListener)) {
		query["IncludeListener"] = request.IncludeListener
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItems)) {
		query["MaxItems"] = request.MaxItems
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TLSCipherPolicyId)) {
		query["TLSCipherPolicyId"] = request.TLSCipherPolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTLSCipherPolicies"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTLSCipherPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTLSCipherPolicies(request *ListTLSCipherPoliciesRequest) (_result *ListTLSCipherPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTLSCipherPoliciesResponse{}
	_body, _err := client.ListTLSCipherPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyLoadBalancerInstanceChargeTypeWithOptions(request *ModifyLoadBalancerInstanceChargeTypeRequest, runtime *util.RuntimeOptions) (_result *ModifyLoadBalancerInstanceChargeTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceChargeType)) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.InternetChargeType)) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerSpec)) {
		query["LoadBalancerSpec"] = request.LoadBalancerSpec
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyLoadBalancerInstanceChargeType"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyLoadBalancerInstanceChargeTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyLoadBalancerInstanceChargeType(request *ModifyLoadBalancerInstanceChargeTypeRequest) (_result *ModifyLoadBalancerInstanceChargeTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLoadBalancerInstanceChargeTypeResponse{}
	_body, _err := client.ModifyLoadBalancerInstanceChargeTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyLoadBalancerInstanceSpecWithOptions(request *ModifyLoadBalancerInstanceSpecRequest, runtime *util.RuntimeOptions) (_result *ModifyLoadBalancerInstanceSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerSpec)) {
		query["LoadBalancerSpec"] = request.LoadBalancerSpec
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyLoadBalancerInstanceSpec"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyLoadBalancerInstanceSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyLoadBalancerInstanceSpec(request *ModifyLoadBalancerInstanceSpecRequest) (_result *ModifyLoadBalancerInstanceSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLoadBalancerInstanceSpecResponse{}
	_body, _err := client.ModifyLoadBalancerInstanceSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyLoadBalancerInternetSpecWithOptions(request *ModifyLoadBalancerInternetSpecRequest, runtime *util.RuntimeOptions) (_result *ModifyLoadBalancerInternetSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.InternetChargeType)) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyLoadBalancerInternetSpec"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyLoadBalancerInternetSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyLoadBalancerInternetSpec(request *ModifyLoadBalancerInternetSpecRequest) (_result *ModifyLoadBalancerInternetSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLoadBalancerInternetSpecResponse{}
	_body, _err := client.ModifyLoadBalancerInternetSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyLoadBalancerPayTypeWithOptions(request *ModifyLoadBalancerPayTypeRequest, runtime *util.RuntimeOptions) (_result *ModifyLoadBalancerPayTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyLoadBalancerPayType"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyLoadBalancerPayTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyLoadBalancerPayType(request *ModifyLoadBalancerPayTypeRequest) (_result *ModifyLoadBalancerPayTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLoadBalancerPayTypeResponse{}
	_body, _err := client.ModifyLoadBalancerPayTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyVServerGroupBackendServersWithOptions(request *ModifyVServerGroupBackendServersRequest, runtime *util.RuntimeOptions) (_result *ModifyVServerGroupBackendServersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewBackendServers)) {
		query["NewBackendServers"] = request.NewBackendServers
	}

	if !tea.BoolValue(util.IsUnset(request.OldBackendServers)) {
		query["OldBackendServers"] = request.OldBackendServers
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroupId)) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyVServerGroupBackendServers"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyVServerGroupBackendServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyVServerGroupBackendServers(request *ModifyVServerGroupBackendServersRequest) (_result *ModifyVServerGroupBackendServersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVServerGroupBackendServersResponse{}
	_body, _err := client.ModifyVServerGroupBackendServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveAccessControlListEntryWithOptions(request *RemoveAccessControlListEntryRequest, runtime *util.RuntimeOptions) (_result *RemoveAccessControlListEntryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclEntrys)) {
		query["AclEntrys"] = request.AclEntrys
	}

	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveAccessControlListEntry"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveAccessControlListEntryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveAccessControlListEntry(request *RemoveAccessControlListEntryRequest) (_result *RemoveAccessControlListEntryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveAccessControlListEntryResponse{}
	_body, _err := client.RemoveAccessControlListEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveBackendServersWithOptions(request *RemoveBackendServersRequest, runtime *util.RuntimeOptions) (_result *RemoveBackendServersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackendServers)) {
		query["BackendServers"] = request.BackendServers
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveBackendServers"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveBackendServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveBackendServers(request *RemoveBackendServersRequest) (_result *RemoveBackendServersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveBackendServersResponse{}
	_body, _err := client.RemoveBackendServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveListenerWhiteListItemWithOptions(request *RemoveListenerWhiteListItemRequest, runtime *util.RuntimeOptions) (_result *RemoveListenerWhiteListItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerProtocol)) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceItems)) {
		query["SourceItems"] = request.SourceItems
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveListenerWhiteListItem"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveListenerWhiteListItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveListenerWhiteListItem(request *RemoveListenerWhiteListItemRequest) (_result *RemoveListenerWhiteListItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveListenerWhiteListItemResponse{}
	_body, _err := client.RemoveListenerWhiteListItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveTagsWithOptions(request *RemoveTagsRequest, runtime *util.RuntimeOptions) (_result *RemoveTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveTags"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveTags(request *RemoveTagsRequest) (_result *RemoveTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveTagsResponse{}
	_body, _err := client.RemoveTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveVServerGroupBackendServersWithOptions(request *RemoveVServerGroupBackendServersRequest, runtime *util.RuntimeOptions) (_result *RemoveVServerGroupBackendServersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackendServers)) {
		query["BackendServers"] = request.BackendServers
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroupId)) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveVServerGroupBackendServers"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveVServerGroupBackendServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveVServerGroupBackendServers(request *RemoveVServerGroupBackendServersRequest) (_result *RemoveVServerGroupBackendServersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveVServerGroupBackendServersResponse{}
	_body, _err := client.RemoveVServerGroupBackendServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAccessControlListAttributeWithOptions(request *SetAccessControlListAttributeRequest, runtime *util.RuntimeOptions) (_result *SetAccessControlListAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.AclName)) {
		query["AclName"] = request.AclName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetAccessControlListAttribute"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetAccessControlListAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAccessControlListAttribute(request *SetAccessControlListAttributeRequest) (_result *SetAccessControlListAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAccessControlListAttributeResponse{}
	_body, _err := client.SetAccessControlListAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetBackendServersWithOptions(request *SetBackendServersRequest, runtime *util.RuntimeOptions) (_result *SetBackendServersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackendServers)) {
		query["BackendServers"] = request.BackendServers
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetBackendServers"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetBackendServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetBackendServers(request *SetBackendServersRequest) (_result *SetBackendServersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetBackendServersResponse{}
	_body, _err := client.SetBackendServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetCACertificateNameWithOptions(request *SetCACertificateNameRequest, runtime *util.RuntimeOptions) (_result *SetCACertificateNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CACertificateId)) {
		query["CACertificateId"] = request.CACertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.CACertificateName)) {
		query["CACertificateName"] = request.CACertificateName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetCACertificateName"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetCACertificateNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetCACertificateName(request *SetCACertificateNameRequest) (_result *SetCACertificateNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetCACertificateNameResponse{}
	_body, _err := client.SetCACertificateNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDomainExtensionAttributeWithOptions(request *SetDomainExtensionAttributeRequest, runtime *util.RuntimeOptions) (_result *SetDomainExtensionAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainExtensionId)) {
		query["DomainExtensionId"] = request.DomainExtensionId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerCertificateId)) {
		query["ServerCertificateId"] = request.ServerCertificateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDomainExtensionAttribute"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDomainExtensionAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDomainExtensionAttribute(request *SetDomainExtensionAttributeRequest) (_result *SetDomainExtensionAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDomainExtensionAttributeResponse{}
	_body, _err := client.SetDomainExtensionAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetListenerAccessControlStatusWithOptions(request *SetListenerAccessControlStatusRequest, runtime *util.RuntimeOptions) (_result *SetListenerAccessControlStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessControlStatus)) {
		query["AccessControlStatus"] = request.AccessControlStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerProtocol)) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetListenerAccessControlStatus"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetListenerAccessControlStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetListenerAccessControlStatus(request *SetListenerAccessControlStatusRequest) (_result *SetListenerAccessControlStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetListenerAccessControlStatusResponse{}
	_body, _err := client.SetListenerAccessControlStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetLoadBalancerDeleteProtectionWithOptions(request *SetLoadBalancerDeleteProtectionRequest, runtime *util.RuntimeOptions) (_result *SetLoadBalancerDeleteProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeleteProtection)) {
		query["DeleteProtection"] = request.DeleteProtection
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetLoadBalancerDeleteProtection"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetLoadBalancerDeleteProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetLoadBalancerDeleteProtection(request *SetLoadBalancerDeleteProtectionRequest) (_result *SetLoadBalancerDeleteProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetLoadBalancerDeleteProtectionResponse{}
	_body, _err := client.SetLoadBalancerDeleteProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetLoadBalancerHTTPListenerAttributeWithOptions(request *SetLoadBalancerHTTPListenerAttributeRequest, runtime *util.RuntimeOptions) (_result *SetLoadBalancerHTTPListenerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.AclStatus)) {
		query["AclStatus"] = request.AclStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AclType)) {
		query["AclType"] = request.AclType
	}

	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.Cookie)) {
		query["Cookie"] = request.Cookie
	}

	if !tea.BoolValue(util.IsUnset(request.CookieTimeout)) {
		query["CookieTimeout"] = request.CookieTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Gzip)) {
		query["Gzip"] = request.Gzip
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheck)) {
		query["HealthCheck"] = request.HealthCheck
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckConnectPort)) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckDomain)) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckHttpCode)) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckInterval)) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckMethod)) {
		query["HealthCheckMethod"] = request.HealthCheckMethod
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckTimeout)) {
		query["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckURI)) {
		query["HealthCheckURI"] = request.HealthCheckURI
	}

	if !tea.BoolValue(util.IsUnset(request.HealthyThreshold)) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.IdleTimeout)) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestTimeout)) {
		query["RequestTimeout"] = request.RequestTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Scheduler)) {
		query["Scheduler"] = request.Scheduler
	}

	if !tea.BoolValue(util.IsUnset(request.StickySession)) {
		query["StickySession"] = request.StickySession
	}

	if !tea.BoolValue(util.IsUnset(request.StickySessionType)) {
		query["StickySessionType"] = request.StickySessionType
	}

	if !tea.BoolValue(util.IsUnset(request.UnhealthyThreshold)) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroup)) {
		query["VServerGroup"] = request.VServerGroup
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroupId)) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedFor)) {
		query["XForwardedFor"] = request.XForwardedFor
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedFor_SLBID)) {
		query["XForwardedFor_SLBID"] = request.XForwardedFor_SLBID
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedFor_SLBIP)) {
		query["XForwardedFor_SLBIP"] = request.XForwardedFor_SLBIP
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedFor_proto)) {
		query["XForwardedFor_proto"] = request.XForwardedFor_proto
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetLoadBalancerHTTPListenerAttribute"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetLoadBalancerHTTPListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetLoadBalancerHTTPListenerAttribute(request *SetLoadBalancerHTTPListenerAttributeRequest) (_result *SetLoadBalancerHTTPListenerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetLoadBalancerHTTPListenerAttributeResponse{}
	_body, _err := client.SetLoadBalancerHTTPListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetLoadBalancerHTTPSListenerAttributeWithOptions(request *SetLoadBalancerHTTPSListenerAttributeRequest, runtime *util.RuntimeOptions) (_result *SetLoadBalancerHTTPSListenerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.AclStatus)) {
		query["AclStatus"] = request.AclStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AclType)) {
		query["AclType"] = request.AclType
	}

	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.CACertificateId)) {
		query["CACertificateId"] = request.CACertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.Cookie)) {
		query["Cookie"] = request.Cookie
	}

	if !tea.BoolValue(util.IsUnset(request.CookieTimeout)) {
		query["CookieTimeout"] = request.CookieTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EnableHttp2)) {
		query["EnableHttp2"] = request.EnableHttp2
	}

	if !tea.BoolValue(util.IsUnset(request.Gzip)) {
		query["Gzip"] = request.Gzip
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheck)) {
		query["HealthCheck"] = request.HealthCheck
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckConnectPort)) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckDomain)) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckHttpCode)) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckInterval)) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckMethod)) {
		query["HealthCheckMethod"] = request.HealthCheckMethod
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckTimeout)) {
		query["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckURI)) {
		query["HealthCheckURI"] = request.HealthCheckURI
	}

	if !tea.BoolValue(util.IsUnset(request.HealthyThreshold)) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.IdleTimeout)) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestTimeout)) {
		query["RequestTimeout"] = request.RequestTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Scheduler)) {
		query["Scheduler"] = request.Scheduler
	}

	if !tea.BoolValue(util.IsUnset(request.ServerCertificateId)) {
		query["ServerCertificateId"] = request.ServerCertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.StickySession)) {
		query["StickySession"] = request.StickySession
	}

	if !tea.BoolValue(util.IsUnset(request.StickySessionType)) {
		query["StickySessionType"] = request.StickySessionType
	}

	if !tea.BoolValue(util.IsUnset(request.TLSCipherPolicy)) {
		query["TLSCipherPolicy"] = request.TLSCipherPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.UnhealthyThreshold)) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroup)) {
		query["VServerGroup"] = request.VServerGroup
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroupId)) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedFor)) {
		query["XForwardedFor"] = request.XForwardedFor
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedFor_SLBID)) {
		query["XForwardedFor_SLBID"] = request.XForwardedFor_SLBID
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedFor_SLBIP)) {
		query["XForwardedFor_SLBIP"] = request.XForwardedFor_SLBIP
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedFor_proto)) {
		query["XForwardedFor_proto"] = request.XForwardedFor_proto
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetLoadBalancerHTTPSListenerAttribute"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetLoadBalancerHTTPSListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetLoadBalancerHTTPSListenerAttribute(request *SetLoadBalancerHTTPSListenerAttributeRequest) (_result *SetLoadBalancerHTTPSListenerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetLoadBalancerHTTPSListenerAttributeResponse{}
	_body, _err := client.SetLoadBalancerHTTPSListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetLoadBalancerModificationProtectionWithOptions(request *SetLoadBalancerModificationProtectionRequest, runtime *util.RuntimeOptions) (_result *SetLoadBalancerModificationProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.ModificationProtectionReason)) {
		query["ModificationProtectionReason"] = request.ModificationProtectionReason
	}

	if !tea.BoolValue(util.IsUnset(request.ModificationProtectionStatus)) {
		query["ModificationProtectionStatus"] = request.ModificationProtectionStatus
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetLoadBalancerModificationProtection"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetLoadBalancerModificationProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetLoadBalancerModificationProtection(request *SetLoadBalancerModificationProtectionRequest) (_result *SetLoadBalancerModificationProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetLoadBalancerModificationProtectionResponse{}
	_body, _err := client.SetLoadBalancerModificationProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetLoadBalancerNameWithOptions(request *SetLoadBalancerNameRequest, runtime *util.RuntimeOptions) (_result *SetLoadBalancerNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerName)) {
		query["LoadBalancerName"] = request.LoadBalancerName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetLoadBalancerName"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetLoadBalancerNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetLoadBalancerName(request *SetLoadBalancerNameRequest) (_result *SetLoadBalancerNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetLoadBalancerNameResponse{}
	_body, _err := client.SetLoadBalancerNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetLoadBalancerStatusWithOptions(request *SetLoadBalancerStatusRequest, runtime *util.RuntimeOptions) (_result *SetLoadBalancerStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerStatus)) {
		query["LoadBalancerStatus"] = request.LoadBalancerStatus
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetLoadBalancerStatus"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetLoadBalancerStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetLoadBalancerStatus(request *SetLoadBalancerStatusRequest) (_result *SetLoadBalancerStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetLoadBalancerStatusResponse{}
	_body, _err := client.SetLoadBalancerStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetLoadBalancerTCPListenerAttributeWithOptions(request *SetLoadBalancerTCPListenerAttributeRequest, runtime *util.RuntimeOptions) (_result *SetLoadBalancerTCPListenerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.AclStatus)) {
		query["AclStatus"] = request.AclStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AclType)) {
		query["AclType"] = request.AclType
	}

	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionDrain)) {
		query["ConnectionDrain"] = request.ConnectionDrain
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionDrainTimeout)) {
		query["ConnectionDrainTimeout"] = request.ConnectionDrainTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EstablishedTimeout)) {
		query["EstablishedTimeout"] = request.EstablishedTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckConnectPort)) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckConnectTimeout)) {
		query["HealthCheckConnectTimeout"] = request.HealthCheckConnectTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckDomain)) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckHttpCode)) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckInterval)) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckSwitch)) {
		query["HealthCheckSwitch"] = request.HealthCheckSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckType)) {
		query["HealthCheckType"] = request.HealthCheckType
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckURI)) {
		query["HealthCheckURI"] = request.HealthCheckURI
	}

	if !tea.BoolValue(util.IsUnset(request.HealthyThreshold)) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.MasterSlaveServerGroup)) {
		query["MasterSlaveServerGroup"] = request.MasterSlaveServerGroup
	}

	if !tea.BoolValue(util.IsUnset(request.MasterSlaveServerGroupId)) {
		query["MasterSlaveServerGroupId"] = request.MasterSlaveServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PersistenceTimeout)) {
		query["PersistenceTimeout"] = request.PersistenceTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyProtocolV2Enabled)) {
		query["ProxyProtocolV2Enabled"] = request.ProxyProtocolV2Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Scheduler)) {
		query["Scheduler"] = request.Scheduler
	}

	if !tea.BoolValue(util.IsUnset(request.SynProxy)) {
		query["SynProxy"] = request.SynProxy
	}

	if !tea.BoolValue(util.IsUnset(request.UnhealthyThreshold)) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroup)) {
		query["VServerGroup"] = request.VServerGroup
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroupId)) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetLoadBalancerTCPListenerAttribute"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetLoadBalancerTCPListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetLoadBalancerTCPListenerAttribute(request *SetLoadBalancerTCPListenerAttributeRequest) (_result *SetLoadBalancerTCPListenerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetLoadBalancerTCPListenerAttributeResponse{}
	_body, _err := client.SetLoadBalancerTCPListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetLoadBalancerUDPListenerAttributeWithOptions(request *SetLoadBalancerUDPListenerAttributeRequest, runtime *util.RuntimeOptions) (_result *SetLoadBalancerUDPListenerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.AclStatus)) {
		query["AclStatus"] = request.AclStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AclType)) {
		query["AclType"] = request.AclType
	}

	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckConnectPort)) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckConnectTimeout)) {
		query["HealthCheckConnectTimeout"] = request.HealthCheckConnectTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckInterval)) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckSwitch)) {
		query["HealthCheckSwitch"] = request.HealthCheckSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.HealthyThreshold)) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.MasterSlaveServerGroup)) {
		query["MasterSlaveServerGroup"] = request.MasterSlaveServerGroup
	}

	if !tea.BoolValue(util.IsUnset(request.MasterSlaveServerGroupId)) {
		query["MasterSlaveServerGroupId"] = request.MasterSlaveServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyProtocolV2Enabled)) {
		query["ProxyProtocolV2Enabled"] = request.ProxyProtocolV2Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Scheduler)) {
		query["Scheduler"] = request.Scheduler
	}

	if !tea.BoolValue(util.IsUnset(request.UnhealthyThreshold)) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroup)) {
		query["VServerGroup"] = request.VServerGroup
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroupId)) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckExp)) {
		query["healthCheckExp"] = request.HealthCheckExp
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckReq)) {
		query["healthCheckReq"] = request.HealthCheckReq
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetLoadBalancerUDPListenerAttribute"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetLoadBalancerUDPListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetLoadBalancerUDPListenerAttribute(request *SetLoadBalancerUDPListenerAttributeRequest) (_result *SetLoadBalancerUDPListenerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetLoadBalancerUDPListenerAttributeResponse{}
	_body, _err := client.SetLoadBalancerUDPListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetRuleWithOptions(request *SetRuleRequest, runtime *util.RuntimeOptions) (_result *SetRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cookie)) {
		query["Cookie"] = request.Cookie
	}

	if !tea.BoolValue(util.IsUnset(request.CookieTimeout)) {
		query["CookieTimeout"] = request.CookieTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheck)) {
		query["HealthCheck"] = request.HealthCheck
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckConnectPort)) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckDomain)) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckHttpCode)) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckInterval)) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckTimeout)) {
		query["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckURI)) {
		query["HealthCheckURI"] = request.HealthCheckURI
	}

	if !tea.BoolValue(util.IsUnset(request.HealthyThreshold)) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerSync)) {
		query["ListenerSync"] = request.ListenerSync
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.Scheduler)) {
		query["Scheduler"] = request.Scheduler
	}

	if !tea.BoolValue(util.IsUnset(request.StickySession)) {
		query["StickySession"] = request.StickySession
	}

	if !tea.BoolValue(util.IsUnset(request.StickySessionType)) {
		query["StickySessionType"] = request.StickySessionType
	}

	if !tea.BoolValue(util.IsUnset(request.UnhealthyThreshold)) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroupId)) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetRule"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetRule(request *SetRuleRequest) (_result *SetRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetRuleResponse{}
	_body, _err := client.SetRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetServerCertificateNameWithOptions(request *SetServerCertificateNameRequest, runtime *util.RuntimeOptions) (_result *SetServerCertificateNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerCertificateId)) {
		query["ServerCertificateId"] = request.ServerCertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerCertificateName)) {
		query["ServerCertificateName"] = request.ServerCertificateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetServerCertificateName"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetServerCertificateNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetServerCertificateName(request *SetServerCertificateNameRequest) (_result *SetServerCertificateNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetServerCertificateNameResponse{}
	_body, _err := client.SetServerCertificateNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetTLSCipherPolicyAttributeWithOptions(request *SetTLSCipherPolicyAttributeRequest, runtime *util.RuntimeOptions) (_result *SetTLSCipherPolicyAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ciphers)) {
		query["Ciphers"] = request.Ciphers
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TLSCipherPolicyId)) {
		query["TLSCipherPolicyId"] = request.TLSCipherPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.TLSVersions)) {
		query["TLSVersions"] = request.TLSVersions
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetTLSCipherPolicyAttribute"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetTLSCipherPolicyAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetTLSCipherPolicyAttribute(request *SetTLSCipherPolicyAttributeRequest) (_result *SetTLSCipherPolicyAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetTLSCipherPolicyAttributeResponse{}
	_body, _err := client.SetTLSCipherPolicyAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetVServerGroupAttributeWithOptions(request *SetVServerGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *SetVServerGroupAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackendServers)) {
		query["BackendServers"] = request.BackendServers
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroupId)) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroupName)) {
		query["VServerGroupName"] = request.VServerGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetVServerGroupAttribute"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetVServerGroupAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetVServerGroupAttribute(request *SetVServerGroupAttributeRequest) (_result *SetVServerGroupAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetVServerGroupAttributeResponse{}
	_body, _err := client.SetVServerGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartLoadBalancerListenerWithOptions(request *StartLoadBalancerListenerRequest, runtime *util.RuntimeOptions) (_result *StartLoadBalancerListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerProtocol)) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartLoadBalancerListener"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartLoadBalancerListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartLoadBalancerListener(request *StartLoadBalancerListenerRequest) (_result *StartLoadBalancerListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartLoadBalancerListenerResponse{}
	_body, _err := client.StartLoadBalancerListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopLoadBalancerListenerWithOptions(request *StopLoadBalancerListenerRequest, runtime *util.RuntimeOptions) (_result *StopLoadBalancerListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerProtocol)) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopLoadBalancerListener"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopLoadBalancerListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopLoadBalancerListener(request *StopLoadBalancerListenerRequest) (_result *StopLoadBalancerListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopLoadBalancerListenerResponse{}
	_body, _err := client.StopLoadBalancerListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadCACertificateWithOptions(request *UploadCACertificateRequest, runtime *util.RuntimeOptions) (_result *UploadCACertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CACertificate)) {
		query["CACertificate"] = request.CACertificate
	}

	if !tea.BoolValue(util.IsUnset(request.CACertificateName)) {
		query["CACertificateName"] = request.CACertificateName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadCACertificate"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadCACertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadCACertificate(request *UploadCACertificateRequest) (_result *UploadCACertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadCACertificateResponse{}
	_body, _err := client.UploadCACertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadServerCertificateWithOptions(request *UploadServerCertificateRequest, runtime *util.RuntimeOptions) (_result *UploadServerCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliCloudCertificateId)) {
		query["AliCloudCertificateId"] = request.AliCloudCertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.AliCloudCertificateName)) {
		query["AliCloudCertificateName"] = request.AliCloudCertificateName
	}

	if !tea.BoolValue(util.IsUnset(request.AliCloudCertificateRegionId)) {
		query["AliCloudCertificateRegionId"] = request.AliCloudCertificateRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateKey)) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerCertificate)) {
		query["ServerCertificate"] = request.ServerCertificate
	}

	if !tea.BoolValue(util.IsUnset(request.ServerCertificateName)) {
		query["ServerCertificateName"] = request.ServerCertificateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadServerCertificate"),
		Version:     tea.String("2014-05-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadServerCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadServerCertificate(request *UploadServerCertificateRequest) (_result *UploadServerCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadServerCertificateResponse{}
	_body, _err := client.UploadServerCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
