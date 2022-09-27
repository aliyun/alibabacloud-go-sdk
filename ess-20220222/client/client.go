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

type AttachAlbServerGroupsRequest struct {
	AlbServerGroups      []*AttachAlbServerGroupsRequestAlbServerGroups `json:"AlbServerGroups,omitempty" xml:"AlbServerGroups,omitempty" type:"Repeated"`
	ClientToken          *string                                        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ForceAttach          *bool                                          `json:"ForceAttach,omitempty" xml:"ForceAttach,omitempty"`
	OwnerId              *int64                                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string                                        `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string                                        `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s AttachAlbServerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachAlbServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *AttachAlbServerGroupsRequest) SetAlbServerGroups(v []*AttachAlbServerGroupsRequestAlbServerGroups) *AttachAlbServerGroupsRequest {
	s.AlbServerGroups = v
	return s
}

func (s *AttachAlbServerGroupsRequest) SetClientToken(v string) *AttachAlbServerGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachAlbServerGroupsRequest) SetForceAttach(v bool) *AttachAlbServerGroupsRequest {
	s.ForceAttach = &v
	return s
}

func (s *AttachAlbServerGroupsRequest) SetOwnerId(v int64) *AttachAlbServerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachAlbServerGroupsRequest) SetRegionId(v string) *AttachAlbServerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *AttachAlbServerGroupsRequest) SetResourceOwnerAccount(v string) *AttachAlbServerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachAlbServerGroupsRequest) SetScalingGroupId(v string) *AttachAlbServerGroupsRequest {
	s.ScalingGroupId = &v
	return s
}

type AttachAlbServerGroupsRequestAlbServerGroups struct {
	AlbServerGroupId *string `json:"AlbServerGroupId,omitempty" xml:"AlbServerGroupId,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Weight           *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s AttachAlbServerGroupsRequestAlbServerGroups) String() string {
	return tea.Prettify(s)
}

func (s AttachAlbServerGroupsRequestAlbServerGroups) GoString() string {
	return s.String()
}

func (s *AttachAlbServerGroupsRequestAlbServerGroups) SetAlbServerGroupId(v string) *AttachAlbServerGroupsRequestAlbServerGroups {
	s.AlbServerGroupId = &v
	return s
}

func (s *AttachAlbServerGroupsRequestAlbServerGroups) SetPort(v int32) *AttachAlbServerGroupsRequestAlbServerGroups {
	s.Port = &v
	return s
}

func (s *AttachAlbServerGroupsRequestAlbServerGroups) SetWeight(v int32) *AttachAlbServerGroupsRequestAlbServerGroups {
	s.Weight = &v
	return s
}

type AttachAlbServerGroupsResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s AttachAlbServerGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachAlbServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *AttachAlbServerGroupsResponseBody) SetRequestId(v string) *AttachAlbServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachAlbServerGroupsResponseBody) SetScalingActivityId(v string) *AttachAlbServerGroupsResponseBody {
	s.ScalingActivityId = &v
	return s
}

type AttachAlbServerGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachAlbServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachAlbServerGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachAlbServerGroupsResponse) GoString() string {
	return s.String()
}

func (s *AttachAlbServerGroupsResponse) SetHeaders(v map[string]*string) *AttachAlbServerGroupsResponse {
	s.Headers = v
	return s
}

func (s *AttachAlbServerGroupsResponse) SetStatusCode(v int32) *AttachAlbServerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachAlbServerGroupsResponse) SetBody(v *AttachAlbServerGroupsResponseBody) *AttachAlbServerGroupsResponse {
	s.Body = v
	return s
}

type AttachDBInstancesRequest struct {
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstances          []*string `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	ForceAttach          *bool     `json:"ForceAttach,omitempty" xml:"ForceAttach,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s AttachDBInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *AttachDBInstancesRequest) SetClientToken(v string) *AttachDBInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachDBInstancesRequest) SetDBInstances(v []*string) *AttachDBInstancesRequest {
	s.DBInstances = v
	return s
}

func (s *AttachDBInstancesRequest) SetForceAttach(v bool) *AttachDBInstancesRequest {
	s.ForceAttach = &v
	return s
}

func (s *AttachDBInstancesRequest) SetOwnerId(v int64) *AttachDBInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachDBInstancesRequest) SetRegionId(v string) *AttachDBInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *AttachDBInstancesRequest) SetResourceOwnerAccount(v string) *AttachDBInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachDBInstancesRequest) SetScalingGroupId(v string) *AttachDBInstancesRequest {
	s.ScalingGroupId = &v
	return s
}

type AttachDBInstancesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachDBInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachDBInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *AttachDBInstancesResponseBody) SetRequestId(v string) *AttachDBInstancesResponseBody {
	s.RequestId = &v
	return s
}

type AttachDBInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachDBInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachDBInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachDBInstancesResponse) GoString() string {
	return s.String()
}

func (s *AttachDBInstancesResponse) SetHeaders(v map[string]*string) *AttachDBInstancesResponse {
	s.Headers = v
	return s
}

func (s *AttachDBInstancesResponse) SetStatusCode(v int32) *AttachDBInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachDBInstancesResponse) SetBody(v *AttachDBInstancesResponseBody) *AttachDBInstancesResponse {
	s.Body = v
	return s
}

type AttachInstancesRequest struct {
	Entrusted            *bool     `json:"Entrusted,omitempty" xml:"Entrusted,omitempty"`
	InstanceIds          []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	LifecycleHook        *bool     `json:"LifecycleHook,omitempty" xml:"LifecycleHook,omitempty"`
	LoadBalancerWeights  []*int32  `json:"LoadBalancerWeights,omitempty" xml:"LoadBalancerWeights,omitempty" type:"Repeated"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s AttachInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesRequest) GoString() string {
	return s.String()
}

func (s *AttachInstancesRequest) SetEntrusted(v bool) *AttachInstancesRequest {
	s.Entrusted = &v
	return s
}

func (s *AttachInstancesRequest) SetInstanceIds(v []*string) *AttachInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *AttachInstancesRequest) SetLifecycleHook(v bool) *AttachInstancesRequest {
	s.LifecycleHook = &v
	return s
}

func (s *AttachInstancesRequest) SetLoadBalancerWeights(v []*int32) *AttachInstancesRequest {
	s.LoadBalancerWeights = v
	return s
}

func (s *AttachInstancesRequest) SetOwnerAccount(v string) *AttachInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AttachInstancesRequest) SetOwnerId(v int64) *AttachInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachInstancesRequest) SetRegionId(v string) *AttachInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *AttachInstancesRequest) SetResourceOwnerAccount(v string) *AttachInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachInstancesRequest) SetResourceOwnerId(v int64) *AttachInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AttachInstancesRequest) SetScalingGroupId(v string) *AttachInstancesRequest {
	s.ScalingGroupId = &v
	return s
}

type AttachInstancesResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s AttachInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *AttachInstancesResponseBody) SetRequestId(v string) *AttachInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachInstancesResponseBody) SetScalingActivityId(v string) *AttachInstancesResponseBody {
	s.ScalingActivityId = &v
	return s
}

type AttachInstancesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesResponse) GoString() string {
	return s.String()
}

func (s *AttachInstancesResponse) SetHeaders(v map[string]*string) *AttachInstancesResponse {
	s.Headers = v
	return s
}

func (s *AttachInstancesResponse) SetStatusCode(v int32) *AttachInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachInstancesResponse) SetBody(v *AttachInstancesResponseBody) *AttachInstancesResponse {
	s.Body = v
	return s
}

type AttachLoadBalancersRequest struct {
	Async                *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ForceAttach          *bool     `json:"ForceAttach,omitempty" xml:"ForceAttach,omitempty"`
	LoadBalancers        []*string `json:"LoadBalancers,omitempty" xml:"LoadBalancers,omitempty" type:"Repeated"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s AttachLoadBalancersRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachLoadBalancersRequest) GoString() string {
	return s.String()
}

func (s *AttachLoadBalancersRequest) SetAsync(v bool) *AttachLoadBalancersRequest {
	s.Async = &v
	return s
}

func (s *AttachLoadBalancersRequest) SetClientToken(v string) *AttachLoadBalancersRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachLoadBalancersRequest) SetForceAttach(v bool) *AttachLoadBalancersRequest {
	s.ForceAttach = &v
	return s
}

func (s *AttachLoadBalancersRequest) SetLoadBalancers(v []*string) *AttachLoadBalancersRequest {
	s.LoadBalancers = v
	return s
}

func (s *AttachLoadBalancersRequest) SetOwnerId(v int64) *AttachLoadBalancersRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachLoadBalancersRequest) SetResourceOwnerAccount(v string) *AttachLoadBalancersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachLoadBalancersRequest) SetScalingGroupId(v string) *AttachLoadBalancersRequest {
	s.ScalingGroupId = &v
	return s
}

type AttachLoadBalancersResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s AttachLoadBalancersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachLoadBalancersResponseBody) GoString() string {
	return s.String()
}

func (s *AttachLoadBalancersResponseBody) SetRequestId(v string) *AttachLoadBalancersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachLoadBalancersResponseBody) SetScalingActivityId(v string) *AttachLoadBalancersResponseBody {
	s.ScalingActivityId = &v
	return s
}

type AttachLoadBalancersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachLoadBalancersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachLoadBalancersResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachLoadBalancersResponse) GoString() string {
	return s.String()
}

func (s *AttachLoadBalancersResponse) SetHeaders(v map[string]*string) *AttachLoadBalancersResponse {
	s.Headers = v
	return s
}

func (s *AttachLoadBalancersResponse) SetStatusCode(v int32) *AttachLoadBalancersResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachLoadBalancersResponse) SetBody(v *AttachLoadBalancersResponseBody) *AttachLoadBalancersResponse {
	s.Body = v
	return s
}

type AttachVServerGroupsRequest struct {
	ClientToken          *string                                    `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ForceAttach          *bool                                      `json:"ForceAttach,omitempty" xml:"ForceAttach,omitempty"`
	OwnerId              *int64                                     `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string                                    `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string                                    `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	VServerGroups        []*AttachVServerGroupsRequestVServerGroups `json:"VServerGroups,omitempty" xml:"VServerGroups,omitempty" type:"Repeated"`
}

func (s AttachVServerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachVServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *AttachVServerGroupsRequest) SetClientToken(v string) *AttachVServerGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachVServerGroupsRequest) SetForceAttach(v bool) *AttachVServerGroupsRequest {
	s.ForceAttach = &v
	return s
}

func (s *AttachVServerGroupsRequest) SetOwnerId(v int64) *AttachVServerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachVServerGroupsRequest) SetRegionId(v string) *AttachVServerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *AttachVServerGroupsRequest) SetResourceOwnerAccount(v string) *AttachVServerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachVServerGroupsRequest) SetScalingGroupId(v string) *AttachVServerGroupsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *AttachVServerGroupsRequest) SetVServerGroups(v []*AttachVServerGroupsRequestVServerGroups) *AttachVServerGroupsRequest {
	s.VServerGroups = v
	return s
}

type AttachVServerGroupsRequestVServerGroups struct {
	LoadBalancerId         *string                                                          `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	VServerGroupAttributes []*AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes `json:"VServerGroupAttributes,omitempty" xml:"VServerGroupAttributes,omitempty" type:"Repeated"`
}

func (s AttachVServerGroupsRequestVServerGroups) String() string {
	return tea.Prettify(s)
}

func (s AttachVServerGroupsRequestVServerGroups) GoString() string {
	return s.String()
}

func (s *AttachVServerGroupsRequestVServerGroups) SetLoadBalancerId(v string) *AttachVServerGroupsRequestVServerGroups {
	s.LoadBalancerId = &v
	return s
}

func (s *AttachVServerGroupsRequestVServerGroups) SetVServerGroupAttributes(v []*AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes) *AttachVServerGroupsRequestVServerGroups {
	s.VServerGroupAttributes = v
	return s
}

type AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes struct {
	Port           *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	Weight         *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes) String() string {
	return tea.Prettify(s)
}

func (s AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes) GoString() string {
	return s.String()
}

func (s *AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes) SetPort(v int32) *AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes {
	s.Port = &v
	return s
}

func (s *AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes) SetVServerGroupId(v string) *AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes {
	s.VServerGroupId = &v
	return s
}

func (s *AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes) SetWeight(v int32) *AttachVServerGroupsRequestVServerGroupsVServerGroupAttributes {
	s.Weight = &v
	return s
}

type AttachVServerGroupsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachVServerGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachVServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *AttachVServerGroupsResponseBody) SetRequestId(v string) *AttachVServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

type AttachVServerGroupsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachVServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachVServerGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachVServerGroupsResponse) GoString() string {
	return s.String()
}

func (s *AttachVServerGroupsResponse) SetHeaders(v map[string]*string) *AttachVServerGroupsResponse {
	s.Headers = v
	return s
}

func (s *AttachVServerGroupsResponse) SetStatusCode(v int32) *AttachVServerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachVServerGroupsResponse) SetBody(v *AttachVServerGroupsResponseBody) *AttachVServerGroupsResponse {
	s.Body = v
	return s
}

type CompleteLifecycleActionRequest struct {
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	LifecycleActionResult *string `json:"LifecycleActionResult,omitempty" xml:"LifecycleActionResult,omitempty"`
	LifecycleActionToken  *string `json:"LifecycleActionToken,omitempty" xml:"LifecycleActionToken,omitempty"`
	LifecycleHookId       *string `json:"LifecycleHookId,omitempty" xml:"LifecycleHookId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s CompleteLifecycleActionRequest) String() string {
	return tea.Prettify(s)
}

func (s CompleteLifecycleActionRequest) GoString() string {
	return s.String()
}

func (s *CompleteLifecycleActionRequest) SetClientToken(v string) *CompleteLifecycleActionRequest {
	s.ClientToken = &v
	return s
}

func (s *CompleteLifecycleActionRequest) SetLifecycleActionResult(v string) *CompleteLifecycleActionRequest {
	s.LifecycleActionResult = &v
	return s
}

func (s *CompleteLifecycleActionRequest) SetLifecycleActionToken(v string) *CompleteLifecycleActionRequest {
	s.LifecycleActionToken = &v
	return s
}

func (s *CompleteLifecycleActionRequest) SetLifecycleHookId(v string) *CompleteLifecycleActionRequest {
	s.LifecycleHookId = &v
	return s
}

func (s *CompleteLifecycleActionRequest) SetOwnerAccount(v string) *CompleteLifecycleActionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CompleteLifecycleActionRequest) SetOwnerId(v int64) *CompleteLifecycleActionRequest {
	s.OwnerId = &v
	return s
}

func (s *CompleteLifecycleActionRequest) SetRegionId(v string) *CompleteLifecycleActionRequest {
	s.RegionId = &v
	return s
}

func (s *CompleteLifecycleActionRequest) SetResourceOwnerAccount(v string) *CompleteLifecycleActionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type CompleteLifecycleActionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CompleteLifecycleActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CompleteLifecycleActionResponseBody) GoString() string {
	return s.String()
}

func (s *CompleteLifecycleActionResponseBody) SetRequestId(v string) *CompleteLifecycleActionResponseBody {
	s.RequestId = &v
	return s
}

type CompleteLifecycleActionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CompleteLifecycleActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CompleteLifecycleActionResponse) String() string {
	return tea.Prettify(s)
}

func (s CompleteLifecycleActionResponse) GoString() string {
	return s.String()
}

func (s *CompleteLifecycleActionResponse) SetHeaders(v map[string]*string) *CompleteLifecycleActionResponse {
	s.Headers = v
	return s
}

func (s *CompleteLifecycleActionResponse) SetStatusCode(v int32) *CompleteLifecycleActionResponse {
	s.StatusCode = &v
	return s
}

func (s *CompleteLifecycleActionResponse) SetBody(v *CompleteLifecycleActionResponseBody) *CompleteLifecycleActionResponse {
	s.Body = v
	return s
}

type CreateAlarmRequest struct {
	AlarmActions             []*string                        `json:"AlarmActions,omitempty" xml:"AlarmActions,omitempty" type:"Repeated"`
	ComparisonOperator       *string                          `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Description              *string                          `json:"Description,omitempty" xml:"Description,omitempty"`
	Dimensions               []*CreateAlarmRequestDimensions  `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	Effective                *string                          `json:"Effective,omitempty" xml:"Effective,omitempty"`
	EvaluationCount          *int32                           `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	Expressions              []*CreateAlarmRequestExpressions `json:"Expressions,omitempty" xml:"Expressions,omitempty" type:"Repeated"`
	ExpressionsLogicOperator *string                          `json:"ExpressionsLogicOperator,omitempty" xml:"ExpressionsLogicOperator,omitempty"`
	GroupId                  *int32                           `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MetricName               *string                          `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	MetricType               *string                          `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Name                     *string                          `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId                  *int64                           `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Period                   *int32                           `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId                 *string                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount     *string                          `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId           *string                          `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	Statistics               *string                          `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold                *float32                         `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmRequest) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequest) SetAlarmActions(v []*string) *CreateAlarmRequest {
	s.AlarmActions = v
	return s
}

func (s *CreateAlarmRequest) SetComparisonOperator(v string) *CreateAlarmRequest {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateAlarmRequest) SetDescription(v string) *CreateAlarmRequest {
	s.Description = &v
	return s
}

func (s *CreateAlarmRequest) SetDimensions(v []*CreateAlarmRequestDimensions) *CreateAlarmRequest {
	s.Dimensions = v
	return s
}

func (s *CreateAlarmRequest) SetEffective(v string) *CreateAlarmRequest {
	s.Effective = &v
	return s
}

func (s *CreateAlarmRequest) SetEvaluationCount(v int32) *CreateAlarmRequest {
	s.EvaluationCount = &v
	return s
}

func (s *CreateAlarmRequest) SetExpressions(v []*CreateAlarmRequestExpressions) *CreateAlarmRequest {
	s.Expressions = v
	return s
}

func (s *CreateAlarmRequest) SetExpressionsLogicOperator(v string) *CreateAlarmRequest {
	s.ExpressionsLogicOperator = &v
	return s
}

func (s *CreateAlarmRequest) SetGroupId(v int32) *CreateAlarmRequest {
	s.GroupId = &v
	return s
}

func (s *CreateAlarmRequest) SetMetricName(v string) *CreateAlarmRequest {
	s.MetricName = &v
	return s
}

func (s *CreateAlarmRequest) SetMetricType(v string) *CreateAlarmRequest {
	s.MetricType = &v
	return s
}

func (s *CreateAlarmRequest) SetName(v string) *CreateAlarmRequest {
	s.Name = &v
	return s
}

func (s *CreateAlarmRequest) SetOwnerId(v int64) *CreateAlarmRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAlarmRequest) SetPeriod(v int32) *CreateAlarmRequest {
	s.Period = &v
	return s
}

func (s *CreateAlarmRequest) SetRegionId(v string) *CreateAlarmRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAlarmRequest) SetResourceOwnerAccount(v string) *CreateAlarmRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAlarmRequest) SetScalingGroupId(v string) *CreateAlarmRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *CreateAlarmRequest) SetStatistics(v string) *CreateAlarmRequest {
	s.Statistics = &v
	return s
}

func (s *CreateAlarmRequest) SetThreshold(v float32) *CreateAlarmRequest {
	s.Threshold = &v
	return s
}

type CreateAlarmRequestDimensions struct {
	DimensionKey   *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
}

func (s CreateAlarmRequestDimensions) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmRequestDimensions) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestDimensions) SetDimensionKey(v string) *CreateAlarmRequestDimensions {
	s.DimensionKey = &v
	return s
}

func (s *CreateAlarmRequestDimensions) SetDimensionValue(v string) *CreateAlarmRequestDimensions {
	s.DimensionValue = &v
	return s
}

type CreateAlarmRequestExpressions struct {
	ComparisonOperator *string  `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	MetricName         *string  `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Period             *int32   `json:"Period,omitempty" xml:"Period,omitempty"`
	Statistics         *string  `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateAlarmRequestExpressions) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmRequestExpressions) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestExpressions) SetComparisonOperator(v string) *CreateAlarmRequestExpressions {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateAlarmRequestExpressions) SetMetricName(v string) *CreateAlarmRequestExpressions {
	s.MetricName = &v
	return s
}

func (s *CreateAlarmRequestExpressions) SetPeriod(v int32) *CreateAlarmRequestExpressions {
	s.Period = &v
	return s
}

func (s *CreateAlarmRequestExpressions) SetStatistics(v string) *CreateAlarmRequestExpressions {
	s.Statistics = &v
	return s
}

func (s *CreateAlarmRequestExpressions) SetThreshold(v float32) *CreateAlarmRequestExpressions {
	s.Threshold = &v
	return s
}

type CreateAlarmResponseBody struct {
	AlarmTaskId *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlarmResponseBody) SetAlarmTaskId(v string) *CreateAlarmResponseBody {
	s.AlarmTaskId = &v
	return s
}

func (s *CreateAlarmResponseBody) SetRequestId(v string) *CreateAlarmResponseBody {
	s.RequestId = &v
	return s
}

type CreateAlarmResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmResponse) GoString() string {
	return s.String()
}

func (s *CreateAlarmResponse) SetHeaders(v map[string]*string) *CreateAlarmResponse {
	s.Headers = v
	return s
}

func (s *CreateAlarmResponse) SetStatusCode(v int32) *CreateAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlarmResponse) SetBody(v *CreateAlarmResponseBody) *CreateAlarmResponse {
	s.Body = v
	return s
}

type CreateEciScalingConfigurationRequest struct {
	AcrRegistryInfos              []*CreateEciScalingConfigurationRequestAcrRegistryInfos         `json:"AcrRegistryInfos,omitempty" xml:"AcrRegistryInfos,omitempty" type:"Repeated"`
	ActiveDeadlineSeconds         *int64                                                          `json:"ActiveDeadlineSeconds,omitempty" xml:"ActiveDeadlineSeconds,omitempty"`
	AutoCreateEip                 *bool                                                           `json:"AutoCreateEip,omitempty" xml:"AutoCreateEip,omitempty"`
	AutoMatchImageCache           *bool                                                           `json:"AutoMatchImageCache,omitempty" xml:"AutoMatchImageCache,omitempty"`
	ContainerGroupName            *string                                                         `json:"ContainerGroupName,omitempty" xml:"ContainerGroupName,omitempty"`
	Containers                    []*CreateEciScalingConfigurationRequestContainers               `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	CostOptimization              *bool                                                           `json:"CostOptimization,omitempty" xml:"CostOptimization,omitempty"`
	Cpu                           *float32                                                        `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CpuOptionsCore                *int32                                                          `json:"CpuOptionsCore,omitempty" xml:"CpuOptionsCore,omitempty"`
	CpuOptionsThreadsPerCore      *int32                                                          `json:"CpuOptionsThreadsPerCore,omitempty" xml:"CpuOptionsThreadsPerCore,omitempty"`
	Description                   *string                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	DnsConfigNameServers          []*string                                                       `json:"DnsConfigNameServers,omitempty" xml:"DnsConfigNameServers,omitempty" type:"Repeated"`
	DnsConfigOptions              []*CreateEciScalingConfigurationRequestDnsConfigOptions         `json:"DnsConfigOptions,omitempty" xml:"DnsConfigOptions,omitempty" type:"Repeated"`
	DnsConfigSearchs              []*string                                                       `json:"DnsConfigSearchs,omitempty" xml:"DnsConfigSearchs,omitempty" type:"Repeated"`
	DnsPolicy                     *string                                                         `json:"DnsPolicy,omitempty" xml:"DnsPolicy,omitempty"`
	EgressBandwidth               *int64                                                          `json:"EgressBandwidth,omitempty" xml:"EgressBandwidth,omitempty"`
	EipBandwidth                  *int32                                                          `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	EnableSls                     *bool                                                           `json:"EnableSls,omitempty" xml:"EnableSls,omitempty"`
	EphemeralStorage              *int32                                                          `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
	HostAliases                   []*CreateEciScalingConfigurationRequestHostAliases              `json:"HostAliases,omitempty" xml:"HostAliases,omitempty" type:"Repeated"`
	HostName                      *string                                                         `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ImageRegistryCredentials      []*CreateEciScalingConfigurationRequestImageRegistryCredentials `json:"ImageRegistryCredentials,omitempty" xml:"ImageRegistryCredentials,omitempty" type:"Repeated"`
	ImageSnapshotId               *string                                                         `json:"ImageSnapshotId,omitempty" xml:"ImageSnapshotId,omitempty"`
	IngressBandwidth              *int64                                                          `json:"IngressBandwidth,omitempty" xml:"IngressBandwidth,omitempty"`
	InitContainers                []*CreateEciScalingConfigurationRequestInitContainers           `json:"InitContainers,omitempty" xml:"InitContainers,omitempty" type:"Repeated"`
	InstanceFamilyLevel           *string                                                         `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	Ipv6AddressCount              *int32                                                          `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	LoadBalancerWeight            *int32                                                          `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	Memory                        *float32                                                        `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NtpServers                    []*string                                                       `json:"NtpServers,omitempty" xml:"NtpServers,omitempty" type:"Repeated"`
	OwnerId                       *int64                                                          `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RamRoleName                   *string                                                         `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	ResourceGroupId               *string                                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount          *string                                                         `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	RestartPolicy                 *string                                                         `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	ScalingConfigurationName      *string                                                         `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	ScalingGroupId                *string                                                         `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	SecurityContextSysctls        []*CreateEciScalingConfigurationRequestSecurityContextSysctls   `json:"SecurityContextSysctls,omitempty" xml:"SecurityContextSysctls,omitempty" type:"Repeated"`
	SecurityGroupId               *string                                                         `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SpotPriceLimit                *float32                                                        `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	SpotStrategy                  *string                                                         `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	Tags                          []*CreateEciScalingConfigurationRequestTags                     `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TerminationGracePeriodSeconds *int64                                                          `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	Volumes                       []*CreateEciScalingConfigurationRequestVolumes                  `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Repeated"`
}

func (s CreateEciScalingConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequest) SetAcrRegistryInfos(v []*CreateEciScalingConfigurationRequestAcrRegistryInfos) *CreateEciScalingConfigurationRequest {
	s.AcrRegistryInfos = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetActiveDeadlineSeconds(v int64) *CreateEciScalingConfigurationRequest {
	s.ActiveDeadlineSeconds = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetAutoCreateEip(v bool) *CreateEciScalingConfigurationRequest {
	s.AutoCreateEip = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetAutoMatchImageCache(v bool) *CreateEciScalingConfigurationRequest {
	s.AutoMatchImageCache = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetContainerGroupName(v string) *CreateEciScalingConfigurationRequest {
	s.ContainerGroupName = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetContainers(v []*CreateEciScalingConfigurationRequestContainers) *CreateEciScalingConfigurationRequest {
	s.Containers = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetCostOptimization(v bool) *CreateEciScalingConfigurationRequest {
	s.CostOptimization = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetCpu(v float32) *CreateEciScalingConfigurationRequest {
	s.Cpu = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetCpuOptionsCore(v int32) *CreateEciScalingConfigurationRequest {
	s.CpuOptionsCore = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetCpuOptionsThreadsPerCore(v int32) *CreateEciScalingConfigurationRequest {
	s.CpuOptionsThreadsPerCore = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetDescription(v string) *CreateEciScalingConfigurationRequest {
	s.Description = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetDnsConfigNameServers(v []*string) *CreateEciScalingConfigurationRequest {
	s.DnsConfigNameServers = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetDnsConfigOptions(v []*CreateEciScalingConfigurationRequestDnsConfigOptions) *CreateEciScalingConfigurationRequest {
	s.DnsConfigOptions = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetDnsConfigSearchs(v []*string) *CreateEciScalingConfigurationRequest {
	s.DnsConfigSearchs = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetDnsPolicy(v string) *CreateEciScalingConfigurationRequest {
	s.DnsPolicy = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetEgressBandwidth(v int64) *CreateEciScalingConfigurationRequest {
	s.EgressBandwidth = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetEipBandwidth(v int32) *CreateEciScalingConfigurationRequest {
	s.EipBandwidth = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetEnableSls(v bool) *CreateEciScalingConfigurationRequest {
	s.EnableSls = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetEphemeralStorage(v int32) *CreateEciScalingConfigurationRequest {
	s.EphemeralStorage = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetHostAliases(v []*CreateEciScalingConfigurationRequestHostAliases) *CreateEciScalingConfigurationRequest {
	s.HostAliases = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetHostName(v string) *CreateEciScalingConfigurationRequest {
	s.HostName = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetImageRegistryCredentials(v []*CreateEciScalingConfigurationRequestImageRegistryCredentials) *CreateEciScalingConfigurationRequest {
	s.ImageRegistryCredentials = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetImageSnapshotId(v string) *CreateEciScalingConfigurationRequest {
	s.ImageSnapshotId = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetIngressBandwidth(v int64) *CreateEciScalingConfigurationRequest {
	s.IngressBandwidth = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetInitContainers(v []*CreateEciScalingConfigurationRequestInitContainers) *CreateEciScalingConfigurationRequest {
	s.InitContainers = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetInstanceFamilyLevel(v string) *CreateEciScalingConfigurationRequest {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetIpv6AddressCount(v int32) *CreateEciScalingConfigurationRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetLoadBalancerWeight(v int32) *CreateEciScalingConfigurationRequest {
	s.LoadBalancerWeight = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetMemory(v float32) *CreateEciScalingConfigurationRequest {
	s.Memory = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetNtpServers(v []*string) *CreateEciScalingConfigurationRequest {
	s.NtpServers = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetOwnerId(v int64) *CreateEciScalingConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetRamRoleName(v string) *CreateEciScalingConfigurationRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetResourceGroupId(v string) *CreateEciScalingConfigurationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetResourceOwnerAccount(v string) *CreateEciScalingConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetRestartPolicy(v string) *CreateEciScalingConfigurationRequest {
	s.RestartPolicy = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetScalingConfigurationName(v string) *CreateEciScalingConfigurationRequest {
	s.ScalingConfigurationName = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetScalingGroupId(v string) *CreateEciScalingConfigurationRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetSecurityContextSysctls(v []*CreateEciScalingConfigurationRequestSecurityContextSysctls) *CreateEciScalingConfigurationRequest {
	s.SecurityContextSysctls = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetSecurityGroupId(v string) *CreateEciScalingConfigurationRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetSpotPriceLimit(v float32) *CreateEciScalingConfigurationRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetSpotStrategy(v string) *CreateEciScalingConfigurationRequest {
	s.SpotStrategy = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetTags(v []*CreateEciScalingConfigurationRequestTags) *CreateEciScalingConfigurationRequest {
	s.Tags = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetTerminationGracePeriodSeconds(v int64) *CreateEciScalingConfigurationRequest {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetVolumes(v []*CreateEciScalingConfigurationRequestVolumes) *CreateEciScalingConfigurationRequest {
	s.Volumes = v
	return s
}

type CreateEciScalingConfigurationRequestAcrRegistryInfos struct {
	Domains      []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateEciScalingConfigurationRequestAcrRegistryInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestAcrRegistryInfos) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestAcrRegistryInfos) SetDomains(v []*string) *CreateEciScalingConfigurationRequestAcrRegistryInfos {
	s.Domains = v
	return s
}

func (s *CreateEciScalingConfigurationRequestAcrRegistryInfos) SetInstanceId(v string) *CreateEciScalingConfigurationRequestAcrRegistryInfos {
	s.InstanceId = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestAcrRegistryInfos) SetInstanceName(v string) *CreateEciScalingConfigurationRequestAcrRegistryInfos {
	s.InstanceName = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestAcrRegistryInfos) SetRegionId(v string) *CreateEciScalingConfigurationRequestAcrRegistryInfos {
	s.RegionId = &v
	return s
}

type CreateEciScalingConfigurationRequestContainers struct {
	LivenessProbe   *CreateEciScalingConfigurationRequestContainersLivenessProbe     `json:"LivenessProbe,omitempty" xml:"LivenessProbe,omitempty" require:"true" type:"Struct"`
	ReadinessProbe  *CreateEciScalingConfigurationRequestContainersReadinessProbe    `json:"ReadinessProbe,omitempty" xml:"ReadinessProbe,omitempty" require:"true" type:"Struct"`
	SecurityContext *CreateEciScalingConfigurationRequestContainersSecurityContext   `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" require:"true" type:"Struct"`
	Args            []*string                                                        `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	Commands        []*string                                                        `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	Cpu             *float32                                                         `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	EnvironmentVars []*CreateEciScalingConfigurationRequestContainersEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	Gpu             *int32                                                           `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Image           *string                                                          `json:"Image,omitempty" xml:"Image,omitempty"`
	ImagePullPolicy *string                                                          `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	Memory          *float32                                                         `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Name            *string                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	Ports           []*CreateEciScalingConfigurationRequestContainersPorts           `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	Stdin           *bool                                                            `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	StdinOnce       *bool                                                            `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	Tty             *bool                                                            `json:"Tty,omitempty" xml:"Tty,omitempty"`
	VolumeMounts    []*CreateEciScalingConfigurationRequestContainersVolumeMounts    `json:"VolumeMounts,omitempty" xml:"VolumeMounts,omitempty" type:"Repeated"`
	WorkingDir      *string                                                          `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainers) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainers) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainers) SetLivenessProbe(v *CreateEciScalingConfigurationRequestContainersLivenessProbe) *CreateEciScalingConfigurationRequestContainers {
	s.LivenessProbe = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetReadinessProbe(v *CreateEciScalingConfigurationRequestContainersReadinessProbe) *CreateEciScalingConfigurationRequestContainers {
	s.ReadinessProbe = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetSecurityContext(v *CreateEciScalingConfigurationRequestContainersSecurityContext) *CreateEciScalingConfigurationRequestContainers {
	s.SecurityContext = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetArgs(v []*string) *CreateEciScalingConfigurationRequestContainers {
	s.Args = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetCommands(v []*string) *CreateEciScalingConfigurationRequestContainers {
	s.Commands = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetCpu(v float32) *CreateEciScalingConfigurationRequestContainers {
	s.Cpu = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetEnvironmentVars(v []*CreateEciScalingConfigurationRequestContainersEnvironmentVars) *CreateEciScalingConfigurationRequestContainers {
	s.EnvironmentVars = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetGpu(v int32) *CreateEciScalingConfigurationRequestContainers {
	s.Gpu = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetImage(v string) *CreateEciScalingConfigurationRequestContainers {
	s.Image = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetImagePullPolicy(v string) *CreateEciScalingConfigurationRequestContainers {
	s.ImagePullPolicy = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetMemory(v float32) *CreateEciScalingConfigurationRequestContainers {
	s.Memory = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetName(v string) *CreateEciScalingConfigurationRequestContainers {
	s.Name = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetPorts(v []*CreateEciScalingConfigurationRequestContainersPorts) *CreateEciScalingConfigurationRequestContainers {
	s.Ports = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetStdin(v bool) *CreateEciScalingConfigurationRequestContainers {
	s.Stdin = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetStdinOnce(v bool) *CreateEciScalingConfigurationRequestContainers {
	s.StdinOnce = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetTty(v bool) *CreateEciScalingConfigurationRequestContainers {
	s.Tty = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetVolumeMounts(v []*CreateEciScalingConfigurationRequestContainersVolumeMounts) *CreateEciScalingConfigurationRequestContainers {
	s.VolumeMounts = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetWorkingDir(v string) *CreateEciScalingConfigurationRequestContainers {
	s.WorkingDir = &v
	return s
}

type CreateEciScalingConfigurationRequestContainersLivenessProbe struct {
	Exec                *CreateEciScalingConfigurationRequestContainersLivenessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" require:"true" type:"Struct"`
	FailureThreshold    *int32                                                                `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" require:"true" type:"Struct"`
	InitialDelaySeconds *int32                                                                `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                                `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                                `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *CreateEciScalingConfigurationRequestContainersLivenessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" require:"true" type:"Struct"`
	TimeoutSeconds      *int32                                                                `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainersLivenessProbe) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersLivenessProbe) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) SetExec(v *CreateEciScalingConfigurationRequestContainersLivenessProbeExec) *CreateEciScalingConfigurationRequestContainersLivenessProbe {
	s.Exec = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) SetFailureThreshold(v int32) *CreateEciScalingConfigurationRequestContainersLivenessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) SetHttpGet(v *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet) *CreateEciScalingConfigurationRequestContainersLivenessProbe {
	s.HttpGet = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) SetInitialDelaySeconds(v int32) *CreateEciScalingConfigurationRequestContainersLivenessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) SetPeriodSeconds(v int32) *CreateEciScalingConfigurationRequestContainersLivenessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) SetSuccessThreshold(v int32) *CreateEciScalingConfigurationRequestContainersLivenessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) SetTcpSocket(v *CreateEciScalingConfigurationRequestContainersLivenessProbeTcpSocket) *CreateEciScalingConfigurationRequestContainersLivenessProbe {
	s.TcpSocket = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) SetTimeoutSeconds(v int32) *CreateEciScalingConfigurationRequestContainersLivenessProbe {
	s.TimeoutSeconds = &v
	return s
}

type CreateEciScalingConfigurationRequestContainersLivenessProbeExec struct {
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
}

func (s CreateEciScalingConfigurationRequestContainersLivenessProbeExec) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersLivenessProbeExec) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbeExec) SetCommands(v []*string) *CreateEciScalingConfigurationRequestContainersLivenessProbeExec {
	s.Commands = v
	return s
}

type CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet struct {
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet) SetPath(v string) *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet) SetPort(v int32) *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet) SetScheme(v string) *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet {
	s.Scheme = &v
	return s
}

type CreateEciScalingConfigurationRequestContainersLivenessProbeTcpSocket struct {
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainersLivenessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersLivenessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbeTcpSocket) SetPort(v int32) *CreateEciScalingConfigurationRequestContainersLivenessProbeTcpSocket {
	s.Port = &v
	return s
}

type CreateEciScalingConfigurationRequestContainersReadinessProbe struct {
	Exec                *CreateEciScalingConfigurationRequestContainersReadinessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" require:"true" type:"Struct"`
	FailureThreshold    *int32                                                                 `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" require:"true" type:"Struct"`
	InitialDelaySeconds *int32                                                                 `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                                 `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                                 `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *CreateEciScalingConfigurationRequestContainersReadinessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" require:"true" type:"Struct"`
	TimeoutSeconds      *int32                                                                 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainersReadinessProbe) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersReadinessProbe) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) SetExec(v *CreateEciScalingConfigurationRequestContainersReadinessProbeExec) *CreateEciScalingConfigurationRequestContainersReadinessProbe {
	s.Exec = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) SetFailureThreshold(v int32) *CreateEciScalingConfigurationRequestContainersReadinessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) SetHttpGet(v *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet) *CreateEciScalingConfigurationRequestContainersReadinessProbe {
	s.HttpGet = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) SetInitialDelaySeconds(v int32) *CreateEciScalingConfigurationRequestContainersReadinessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) SetPeriodSeconds(v int32) *CreateEciScalingConfigurationRequestContainersReadinessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) SetSuccessThreshold(v int32) *CreateEciScalingConfigurationRequestContainersReadinessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) SetTcpSocket(v *CreateEciScalingConfigurationRequestContainersReadinessProbeTcpSocket) *CreateEciScalingConfigurationRequestContainersReadinessProbe {
	s.TcpSocket = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) SetTimeoutSeconds(v int32) *CreateEciScalingConfigurationRequestContainersReadinessProbe {
	s.TimeoutSeconds = &v
	return s
}

type CreateEciScalingConfigurationRequestContainersReadinessProbeExec struct {
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
}

func (s CreateEciScalingConfigurationRequestContainersReadinessProbeExec) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersReadinessProbeExec) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbeExec) SetCommands(v []*string) *CreateEciScalingConfigurationRequestContainersReadinessProbeExec {
	s.Commands = v
	return s
}

type CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet struct {
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet) SetPath(v string) *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet) SetPort(v int32) *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet) SetScheme(v string) *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet {
	s.Scheme = &v
	return s
}

type CreateEciScalingConfigurationRequestContainersReadinessProbeTcpSocket struct {
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainersReadinessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersReadinessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbeTcpSocket) SetPort(v int32) *CreateEciScalingConfigurationRequestContainersReadinessProbeTcpSocket {
	s.Port = &v
	return s
}

type CreateEciScalingConfigurationRequestContainersSecurityContext struct {
	Capability             *CreateEciScalingConfigurationRequestContainersSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" require:"true" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                                    `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                                   `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainersSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersSecurityContext) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersSecurityContext) SetCapability(v *CreateEciScalingConfigurationRequestContainersSecurityContextCapability) *CreateEciScalingConfigurationRequestContainersSecurityContext {
	s.Capability = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersSecurityContext) SetReadOnlyRootFilesystem(v bool) *CreateEciScalingConfigurationRequestContainersSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersSecurityContext) SetRunAsUser(v int64) *CreateEciScalingConfigurationRequestContainersSecurityContext {
	s.RunAsUser = &v
	return s
}

type CreateEciScalingConfigurationRequestContainersSecurityContextCapability struct {
	Add []*string `json:"Add,omitempty" xml:"Add,omitempty" type:"Repeated"`
}

func (s CreateEciScalingConfigurationRequestContainersSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersSecurityContextCapability) SetAdd(v []*string) *CreateEciScalingConfigurationRequestContainersSecurityContextCapability {
	s.Add = v
	return s
}

type CreateEciScalingConfigurationRequestContainersEnvironmentVars struct {
	FieldRefFieldPath *string `json:"FieldRefFieldPath,omitempty" xml:"FieldRefFieldPath,omitempty"`
	Key               *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value             *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainersEnvironmentVars) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersEnvironmentVars) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersEnvironmentVars) SetFieldRefFieldPath(v string) *CreateEciScalingConfigurationRequestContainersEnvironmentVars {
	s.FieldRefFieldPath = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersEnvironmentVars) SetKey(v string) *CreateEciScalingConfigurationRequestContainersEnvironmentVars {
	s.Key = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersEnvironmentVars) SetValue(v string) *CreateEciScalingConfigurationRequestContainersEnvironmentVars {
	s.Value = &v
	return s
}

type CreateEciScalingConfigurationRequestContainersPorts struct {
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainersPorts) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersPorts) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersPorts) SetPort(v int32) *CreateEciScalingConfigurationRequestContainersPorts {
	s.Port = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersPorts) SetProtocol(v string) *CreateEciScalingConfigurationRequestContainersPorts {
	s.Protocol = &v
	return s
}

type CreateEciScalingConfigurationRequestContainersVolumeMounts struct {
	MountPath        *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ReadOnly         *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	SubPath          *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainersVolumeMounts) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersVolumeMounts) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersVolumeMounts) SetMountPath(v string) *CreateEciScalingConfigurationRequestContainersVolumeMounts {
	s.MountPath = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersVolumeMounts) SetMountPropagation(v string) *CreateEciScalingConfigurationRequestContainersVolumeMounts {
	s.MountPropagation = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersVolumeMounts) SetName(v string) *CreateEciScalingConfigurationRequestContainersVolumeMounts {
	s.Name = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersVolumeMounts) SetReadOnly(v bool) *CreateEciScalingConfigurationRequestContainersVolumeMounts {
	s.ReadOnly = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersVolumeMounts) SetSubPath(v string) *CreateEciScalingConfigurationRequestContainersVolumeMounts {
	s.SubPath = &v
	return s
}

type CreateEciScalingConfigurationRequestDnsConfigOptions struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEciScalingConfigurationRequestDnsConfigOptions) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestDnsConfigOptions) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestDnsConfigOptions) SetName(v string) *CreateEciScalingConfigurationRequestDnsConfigOptions {
	s.Name = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestDnsConfigOptions) SetValue(v string) *CreateEciScalingConfigurationRequestDnsConfigOptions {
	s.Value = &v
	return s
}

type CreateEciScalingConfigurationRequestHostAliases struct {
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	Ip        *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s CreateEciScalingConfigurationRequestHostAliases) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestHostAliases) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestHostAliases) SetHostnames(v []*string) *CreateEciScalingConfigurationRequestHostAliases {
	s.Hostnames = v
	return s
}

func (s *CreateEciScalingConfigurationRequestHostAliases) SetIp(v string) *CreateEciScalingConfigurationRequestHostAliases {
	s.Ip = &v
	return s
}

type CreateEciScalingConfigurationRequestImageRegistryCredentials struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateEciScalingConfigurationRequestImageRegistryCredentials) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestImageRegistryCredentials) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestImageRegistryCredentials) SetPassword(v string) *CreateEciScalingConfigurationRequestImageRegistryCredentials {
	s.Password = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestImageRegistryCredentials) SetServer(v string) *CreateEciScalingConfigurationRequestImageRegistryCredentials {
	s.Server = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestImageRegistryCredentials) SetUserName(v string) *CreateEciScalingConfigurationRequestImageRegistryCredentials {
	s.UserName = &v
	return s
}

type CreateEciScalingConfigurationRequestInitContainers struct {
	SecurityContext              *CreateEciScalingConfigurationRequestInitContainersSecurityContext                `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" require:"true" type:"Struct"`
	Args                         []*string                                                                         `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	Commands                     []*string                                                                         `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	Cpu                          *float32                                                                          `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu                          *int32                                                                            `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Image                        *string                                                                           `json:"Image,omitempty" xml:"Image,omitempty"`
	ImagePullPolicy              *string                                                                           `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	InitContainerEnvironmentVars []*CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars `json:"InitContainerEnvironmentVars,omitempty" xml:"InitContainerEnvironmentVars,omitempty" type:"Repeated"`
	InitContainerPorts           []*CreateEciScalingConfigurationRequestInitContainersInitContainerPorts           `json:"InitContainerPorts,omitempty" xml:"InitContainerPorts,omitempty" type:"Repeated"`
	InitContainerVolumeMounts    []*CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts    `json:"InitContainerVolumeMounts,omitempty" xml:"InitContainerVolumeMounts,omitempty" type:"Repeated"`
	Memory                       *float32                                                                          `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Name                         *string                                                                           `json:"Name,omitempty" xml:"Name,omitempty"`
	WorkingDir                   *string                                                                           `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s CreateEciScalingConfigurationRequestInitContainers) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestInitContainers) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetSecurityContext(v *CreateEciScalingConfigurationRequestInitContainersSecurityContext) *CreateEciScalingConfigurationRequestInitContainers {
	s.SecurityContext = v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetArgs(v []*string) *CreateEciScalingConfigurationRequestInitContainers {
	s.Args = v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetCommands(v []*string) *CreateEciScalingConfigurationRequestInitContainers {
	s.Commands = v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetCpu(v float32) *CreateEciScalingConfigurationRequestInitContainers {
	s.Cpu = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetGpu(v int32) *CreateEciScalingConfigurationRequestInitContainers {
	s.Gpu = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetImage(v string) *CreateEciScalingConfigurationRequestInitContainers {
	s.Image = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetImagePullPolicy(v string) *CreateEciScalingConfigurationRequestInitContainers {
	s.ImagePullPolicy = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetInitContainerEnvironmentVars(v []*CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) *CreateEciScalingConfigurationRequestInitContainers {
	s.InitContainerEnvironmentVars = v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetInitContainerPorts(v []*CreateEciScalingConfigurationRequestInitContainersInitContainerPorts) *CreateEciScalingConfigurationRequestInitContainers {
	s.InitContainerPorts = v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetInitContainerVolumeMounts(v []*CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) *CreateEciScalingConfigurationRequestInitContainers {
	s.InitContainerVolumeMounts = v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetMemory(v float32) *CreateEciScalingConfigurationRequestInitContainers {
	s.Memory = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetName(v string) *CreateEciScalingConfigurationRequestInitContainers {
	s.Name = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetWorkingDir(v string) *CreateEciScalingConfigurationRequestInitContainers {
	s.WorkingDir = &v
	return s
}

type CreateEciScalingConfigurationRequestInitContainersSecurityContext struct {
	Capability             *CreateEciScalingConfigurationRequestInitContainersSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" require:"true" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                                        `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                                       `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s CreateEciScalingConfigurationRequestInitContainersSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestInitContainersSecurityContext) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestInitContainersSecurityContext) SetCapability(v *CreateEciScalingConfigurationRequestInitContainersSecurityContextCapability) *CreateEciScalingConfigurationRequestInitContainersSecurityContext {
	s.Capability = v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersSecurityContext) SetReadOnlyRootFilesystem(v bool) *CreateEciScalingConfigurationRequestInitContainersSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersSecurityContext) SetRunAsUser(v int64) *CreateEciScalingConfigurationRequestInitContainersSecurityContext {
	s.RunAsUser = &v
	return s
}

type CreateEciScalingConfigurationRequestInitContainersSecurityContextCapability struct {
	Adds []*string `json:"Adds,omitempty" xml:"Adds,omitempty" type:"Repeated"`
}

func (s CreateEciScalingConfigurationRequestInitContainersSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestInitContainersSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestInitContainersSecurityContextCapability) SetAdds(v []*string) *CreateEciScalingConfigurationRequestInitContainersSecurityContextCapability {
	s.Adds = v
	return s
}

type CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars struct {
	FieldRefFieldPath *string `json:"FieldRefFieldPath,omitempty" xml:"FieldRefFieldPath,omitempty"`
	Key               *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value             *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) SetFieldRefFieldPath(v string) *CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars {
	s.FieldRefFieldPath = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) SetKey(v string) *CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars {
	s.Key = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) SetValue(v string) *CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars {
	s.Value = &v
	return s
}

type CreateEciScalingConfigurationRequestInitContainersInitContainerPorts struct {
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s CreateEciScalingConfigurationRequestInitContainersInitContainerPorts) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestInitContainersInitContainerPorts) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerPorts) SetPort(v int32) *CreateEciScalingConfigurationRequestInitContainersInitContainerPorts {
	s.Port = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerPorts) SetProtocol(v string) *CreateEciScalingConfigurationRequestInitContainersInitContainerPorts {
	s.Protocol = &v
	return s
}

type CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts struct {
	MountPath        *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ReadOnly         *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	SubPath          *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) SetMountPath(v string) *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	s.MountPath = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) SetMountPropagation(v string) *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	s.MountPropagation = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) SetName(v string) *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	s.Name = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) SetReadOnly(v bool) *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	s.ReadOnly = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) SetSubPath(v string) *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	s.SubPath = &v
	return s
}

type CreateEciScalingConfigurationRequestSecurityContextSysctls struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEciScalingConfigurationRequestSecurityContextSysctls) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestSecurityContextSysctls) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestSecurityContextSysctls) SetName(v string) *CreateEciScalingConfigurationRequestSecurityContextSysctls {
	s.Name = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestSecurityContextSysctls) SetValue(v string) *CreateEciScalingConfigurationRequestSecurityContextSysctls {
	s.Value = &v
	return s
}

type CreateEciScalingConfigurationRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEciScalingConfigurationRequestTags) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestTags) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestTags) SetKey(v string) *CreateEciScalingConfigurationRequestTags {
	s.Key = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestTags) SetValue(v string) *CreateEciScalingConfigurationRequestTags {
	s.Value = &v
	return s
}

type CreateEciScalingConfigurationRequestVolumes struct {
	DiskVolume                        *CreateEciScalingConfigurationRequestVolumesDiskVolume                          `json:"DiskVolume,omitempty" xml:"DiskVolume,omitempty" require:"true" type:"Struct"`
	EmptyDirVolume                    *CreateEciScalingConfigurationRequestVolumesEmptyDirVolume                      `json:"EmptyDirVolume,omitempty" xml:"EmptyDirVolume,omitempty" require:"true" type:"Struct"`
	FlexVolume                        *CreateEciScalingConfigurationRequestVolumesFlexVolume                          `json:"FlexVolume,omitempty" xml:"FlexVolume,omitempty" require:"true" type:"Struct"`
	HostPathVolume                    *CreateEciScalingConfigurationRequestVolumesHostPathVolume                      `json:"HostPathVolume,omitempty" xml:"HostPathVolume,omitempty" require:"true" type:"Struct"`
	NFSVolume                         *CreateEciScalingConfigurationRequestVolumesNFSVolume                           `json:"NFSVolume,omitempty" xml:"NFSVolume,omitempty" require:"true" type:"Struct"`
	ConfigFileVolumeConfigFileToPaths []*CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths `json:"ConfigFileVolumeConfigFileToPaths,omitempty" xml:"ConfigFileVolumeConfigFileToPaths,omitempty" type:"Repeated"`
	ConfigFileVolumeDefaultMode       *int32                                                                          `json:"ConfigFileVolumeDefaultMode,omitempty" xml:"ConfigFileVolumeDefaultMode,omitempty"`
	Name                              *string                                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	Type                              *string                                                                         `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateEciScalingConfigurationRequestVolumes) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestVolumes) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestVolumes) SetDiskVolume(v *CreateEciScalingConfigurationRequestVolumesDiskVolume) *CreateEciScalingConfigurationRequestVolumes {
	s.DiskVolume = v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumes) SetEmptyDirVolume(v *CreateEciScalingConfigurationRequestVolumesEmptyDirVolume) *CreateEciScalingConfigurationRequestVolumes {
	s.EmptyDirVolume = v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumes) SetFlexVolume(v *CreateEciScalingConfigurationRequestVolumesFlexVolume) *CreateEciScalingConfigurationRequestVolumes {
	s.FlexVolume = v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumes) SetHostPathVolume(v *CreateEciScalingConfigurationRequestVolumesHostPathVolume) *CreateEciScalingConfigurationRequestVolumes {
	s.HostPathVolume = v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumes) SetNFSVolume(v *CreateEciScalingConfigurationRequestVolumesNFSVolume) *CreateEciScalingConfigurationRequestVolumes {
	s.NFSVolume = v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumes) SetConfigFileVolumeConfigFileToPaths(v []*CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths) *CreateEciScalingConfigurationRequestVolumes {
	s.ConfigFileVolumeConfigFileToPaths = v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumes) SetConfigFileVolumeDefaultMode(v int32) *CreateEciScalingConfigurationRequestVolumes {
	s.ConfigFileVolumeDefaultMode = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumes) SetName(v string) *CreateEciScalingConfigurationRequestVolumes {
	s.Name = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumes) SetType(v string) *CreateEciScalingConfigurationRequestVolumes {
	s.Type = &v
	return s
}

type CreateEciScalingConfigurationRequestVolumesDiskVolume struct {
	DiskId   *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	DiskSize *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	FsType   *string `json:"FsType,omitempty" xml:"FsType,omitempty"`
}

func (s CreateEciScalingConfigurationRequestVolumesDiskVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestVolumesDiskVolume) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestVolumesDiskVolume) SetDiskId(v string) *CreateEciScalingConfigurationRequestVolumesDiskVolume {
	s.DiskId = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesDiskVolume) SetDiskSize(v int32) *CreateEciScalingConfigurationRequestVolumesDiskVolume {
	s.DiskSize = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesDiskVolume) SetFsType(v string) *CreateEciScalingConfigurationRequestVolumesDiskVolume {
	s.FsType = &v
	return s
}

type CreateEciScalingConfigurationRequestVolumesEmptyDirVolume struct {
	Medium *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
}

func (s CreateEciScalingConfigurationRequestVolumesEmptyDirVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestVolumesEmptyDirVolume) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestVolumesEmptyDirVolume) SetMedium(v string) *CreateEciScalingConfigurationRequestVolumesEmptyDirVolume {
	s.Medium = &v
	return s
}

type CreateEciScalingConfigurationRequestVolumesFlexVolume struct {
	Driver  *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
	FsType  *string `json:"FsType,omitempty" xml:"FsType,omitempty"`
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
}

func (s CreateEciScalingConfigurationRequestVolumesFlexVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestVolumesFlexVolume) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestVolumesFlexVolume) SetDriver(v string) *CreateEciScalingConfigurationRequestVolumesFlexVolume {
	s.Driver = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesFlexVolume) SetFsType(v string) *CreateEciScalingConfigurationRequestVolumesFlexVolume {
	s.FsType = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesFlexVolume) SetOptions(v string) *CreateEciScalingConfigurationRequestVolumesFlexVolume {
	s.Options = &v
	return s
}

type CreateEciScalingConfigurationRequestVolumesHostPathVolume struct {
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateEciScalingConfigurationRequestVolumesHostPathVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestVolumesHostPathVolume) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestVolumesHostPathVolume) SetPath(v string) *CreateEciScalingConfigurationRequestVolumesHostPathVolume {
	s.Path = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesHostPathVolume) SetType(v string) *CreateEciScalingConfigurationRequestVolumesHostPathVolume {
	s.Type = &v
	return s
}

type CreateEciScalingConfigurationRequestVolumesNFSVolume struct {
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty"`
	ReadOnly *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
}

func (s CreateEciScalingConfigurationRequestVolumesNFSVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestVolumesNFSVolume) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestVolumesNFSVolume) SetPath(v string) *CreateEciScalingConfigurationRequestVolumesNFSVolume {
	s.Path = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesNFSVolume) SetReadOnly(v bool) *CreateEciScalingConfigurationRequestVolumesNFSVolume {
	s.ReadOnly = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesNFSVolume) SetServer(v string) *CreateEciScalingConfigurationRequestVolumesNFSVolume {
	s.Server = &v
	return s
}

type CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Mode    *int32  `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Path    *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths) SetContent(v string) *CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths {
	s.Content = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths) SetMode(v int32) *CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths {
	s.Mode = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths) SetPath(v string) *CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths {
	s.Path = &v
	return s
}

type CreateEciScalingConfigurationResponseBody struct {
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
}

func (s CreateEciScalingConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationResponseBody) SetRequestId(v string) *CreateEciScalingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEciScalingConfigurationResponseBody) SetScalingConfigurationId(v string) *CreateEciScalingConfigurationResponseBody {
	s.ScalingConfigurationId = &v
	return s
}

type CreateEciScalingConfigurationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateEciScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEciScalingConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEciScalingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationResponse) SetHeaders(v map[string]*string) *CreateEciScalingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *CreateEciScalingConfigurationResponse) SetStatusCode(v int32) *CreateEciScalingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEciScalingConfigurationResponse) SetBody(v *CreateEciScalingConfigurationResponseBody) *CreateEciScalingConfigurationResponse {
	s.Body = v
	return s
}

type CreateLifecycleHookRequest struct {
	DefaultResult        *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	HeartbeatTimeout     *int32  `json:"HeartbeatTimeout,omitempty" xml:"HeartbeatTimeout,omitempty"`
	LifecycleHookName    *string `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	LifecycleTransition  *string `json:"LifecycleTransition,omitempty" xml:"LifecycleTransition,omitempty"`
	NotificationArn      *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	NotificationMetadata *string `json:"NotificationMetadata,omitempty" xml:"NotificationMetadata,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s CreateLifecycleHookRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLifecycleHookRequest) GoString() string {
	return s.String()
}

func (s *CreateLifecycleHookRequest) SetDefaultResult(v string) *CreateLifecycleHookRequest {
	s.DefaultResult = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetHeartbeatTimeout(v int32) *CreateLifecycleHookRequest {
	s.HeartbeatTimeout = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetLifecycleHookName(v string) *CreateLifecycleHookRequest {
	s.LifecycleHookName = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetLifecycleTransition(v string) *CreateLifecycleHookRequest {
	s.LifecycleTransition = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetNotificationArn(v string) *CreateLifecycleHookRequest {
	s.NotificationArn = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetNotificationMetadata(v string) *CreateLifecycleHookRequest {
	s.NotificationMetadata = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetOwnerAccount(v string) *CreateLifecycleHookRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetOwnerId(v int64) *CreateLifecycleHookRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetResourceOwnerAccount(v string) *CreateLifecycleHookRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetScalingGroupId(v string) *CreateLifecycleHookRequest {
	s.ScalingGroupId = &v
	return s
}

type CreateLifecycleHookResponseBody struct {
	LifecycleHookId *string `json:"LifecycleHookId,omitempty" xml:"LifecycleHookId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLifecycleHookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLifecycleHookResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLifecycleHookResponseBody) SetLifecycleHookId(v string) *CreateLifecycleHookResponseBody {
	s.LifecycleHookId = &v
	return s
}

func (s *CreateLifecycleHookResponseBody) SetRequestId(v string) *CreateLifecycleHookResponseBody {
	s.RequestId = &v
	return s
}

type CreateLifecycleHookResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLifecycleHookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLifecycleHookResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLifecycleHookResponse) GoString() string {
	return s.String()
}

func (s *CreateLifecycleHookResponse) SetHeaders(v map[string]*string) *CreateLifecycleHookResponse {
	s.Headers = v
	return s
}

func (s *CreateLifecycleHookResponse) SetStatusCode(v int32) *CreateLifecycleHookResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLifecycleHookResponse) SetBody(v *CreateLifecycleHookResponseBody) *CreateLifecycleHookResponse {
	s.Body = v
	return s
}

type CreateNotificationConfigurationRequest struct {
	NotificationArn      *string   `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	NotificationTypes    []*string `json:"NotificationTypes,omitempty" xml:"NotificationTypes,omitempty" type:"Repeated"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s CreateNotificationConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNotificationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateNotificationConfigurationRequest) SetNotificationArn(v string) *CreateNotificationConfigurationRequest {
	s.NotificationArn = &v
	return s
}

func (s *CreateNotificationConfigurationRequest) SetNotificationTypes(v []*string) *CreateNotificationConfigurationRequest {
	s.NotificationTypes = v
	return s
}

func (s *CreateNotificationConfigurationRequest) SetOwnerId(v int64) *CreateNotificationConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateNotificationConfigurationRequest) SetRegionId(v string) *CreateNotificationConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNotificationConfigurationRequest) SetResourceOwnerAccount(v string) *CreateNotificationConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateNotificationConfigurationRequest) SetScalingGroupId(v string) *CreateNotificationConfigurationRequest {
	s.ScalingGroupId = &v
	return s
}

type CreateNotificationConfigurationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNotificationConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNotificationConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNotificationConfigurationResponseBody) SetRequestId(v string) *CreateNotificationConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type CreateNotificationConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateNotificationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateNotificationConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNotificationConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CreateNotificationConfigurationResponse) SetHeaders(v map[string]*string) *CreateNotificationConfigurationResponse {
	s.Headers = v
	return s
}

func (s *CreateNotificationConfigurationResponse) SetStatusCode(v int32) *CreateNotificationConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNotificationConfigurationResponse) SetBody(v *CreateNotificationConfigurationResponseBody) *CreateNotificationConfigurationResponse {
	s.Body = v
	return s
}

type CreateScalingConfigurationRequest struct {
	PrivatePoolOptions          *CreateScalingConfigurationRequestPrivatePoolOptions      `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	SystemDisk                  *CreateScalingConfigurationRequestSystemDisk              `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	Affinity                    *string                                                   `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	ClientToken                 *string                                                   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Cpu                         *int32                                                    `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreditSpecification         *string                                                   `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	DataDisks                   []*CreateScalingConfigurationRequestDataDisks             `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	DedicatedHostId             *string                                                   `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	DeploymentSetId             *string                                                   `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	HostName                    *string                                                   `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HpcClusterId                *string                                                   `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	ImageFamily                 *string                                                   `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	ImageId                     *string                                                   `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName                   *string                                                   `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	InstanceDescription         *string                                                   `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	InstanceName                *string                                                   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstancePatternInfos        []*CreateScalingConfigurationRequestInstancePatternInfos  `json:"InstancePatternInfos,omitempty" xml:"InstancePatternInfos,omitempty" type:"Repeated"`
	InstanceType                *string                                                   `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceTypeOverrides       []*CreateScalingConfigurationRequestInstanceTypeOverrides `json:"InstanceTypeOverrides,omitempty" xml:"InstanceTypeOverrides,omitempty" type:"Repeated"`
	InstanceTypes               []*string                                                 `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	InternetChargeType          *string                                                   `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandwidthIn      *int32                                                    `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	InternetMaxBandwidthOut     *int32                                                    `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	IoOptimized                 *string                                                   `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	Ipv6AddressCount            *int32                                                    `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	KeyPairName                 *string                                                   `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	LoadBalancerWeight          *int32                                                    `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	Memory                      *int32                                                    `json:"Memory,omitempty" xml:"Memory,omitempty"`
	OwnerAccount                *string                                                   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                     *int64                                                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Password                    *string                                                   `json:"Password,omitempty" xml:"Password,omitempty"`
	PasswordInherit             *bool                                                     `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	RamRoleName                 *string                                                   `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	ResourceGroupId             *string                                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount        *string                                                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingConfigurationName    *string                                                   `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	ScalingGroupId              *string                                                   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	SchedulerOptions            map[string]interface{}                                    `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty"`
	SecurityEnhancementStrategy *string                                                   `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	SecurityGroupId             *string                                                   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupIds            []*string                                                 `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	SpotDuration                *int32                                                    `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	SpotInterruptionBehavior    *string                                                   `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	SpotPriceLimits             []*CreateScalingConfigurationRequestSpotPriceLimits       `json:"SpotPriceLimits,omitempty" xml:"SpotPriceLimits,omitempty" type:"Repeated"`
	SpotStrategy                *string                                                   `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDiskCategories        []*string                                                 `json:"SystemDiskCategories,omitempty" xml:"SystemDiskCategories,omitempty" type:"Repeated"`
	Tags                        *string                                                   `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Tenancy                     *string                                                   `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	UserData                    *string                                                   `json:"UserData,omitempty" xml:"UserData,omitempty"`
	ZoneId                      *string                                                   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateScalingConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequest) SetPrivatePoolOptions(v *CreateScalingConfigurationRequestPrivatePoolOptions) *CreateScalingConfigurationRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSystemDisk(v *CreateScalingConfigurationRequestSystemDisk) *CreateScalingConfigurationRequest {
	s.SystemDisk = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetAffinity(v string) *CreateScalingConfigurationRequest {
	s.Affinity = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetClientToken(v string) *CreateScalingConfigurationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetCpu(v int32) *CreateScalingConfigurationRequest {
	s.Cpu = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetCreditSpecification(v string) *CreateScalingConfigurationRequest {
	s.CreditSpecification = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetDataDisks(v []*CreateScalingConfigurationRequestDataDisks) *CreateScalingConfigurationRequest {
	s.DataDisks = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetDedicatedHostId(v string) *CreateScalingConfigurationRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetDeploymentSetId(v string) *CreateScalingConfigurationRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetHostName(v string) *CreateScalingConfigurationRequest {
	s.HostName = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetHpcClusterId(v string) *CreateScalingConfigurationRequest {
	s.HpcClusterId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetImageFamily(v string) *CreateScalingConfigurationRequest {
	s.ImageFamily = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetImageId(v string) *CreateScalingConfigurationRequest {
	s.ImageId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetImageName(v string) *CreateScalingConfigurationRequest {
	s.ImageName = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInstanceDescription(v string) *CreateScalingConfigurationRequest {
	s.InstanceDescription = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInstanceName(v string) *CreateScalingConfigurationRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInstancePatternInfos(v []*CreateScalingConfigurationRequestInstancePatternInfos) *CreateScalingConfigurationRequest {
	s.InstancePatternInfos = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInstanceType(v string) *CreateScalingConfigurationRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInstanceTypeOverrides(v []*CreateScalingConfigurationRequestInstanceTypeOverrides) *CreateScalingConfigurationRequest {
	s.InstanceTypeOverrides = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInstanceTypes(v []*string) *CreateScalingConfigurationRequest {
	s.InstanceTypes = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInternetChargeType(v string) *CreateScalingConfigurationRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInternetMaxBandwidthIn(v int32) *CreateScalingConfigurationRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInternetMaxBandwidthOut(v int32) *CreateScalingConfigurationRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetIoOptimized(v string) *CreateScalingConfigurationRequest {
	s.IoOptimized = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetIpv6AddressCount(v int32) *CreateScalingConfigurationRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetKeyPairName(v string) *CreateScalingConfigurationRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetLoadBalancerWeight(v int32) *CreateScalingConfigurationRequest {
	s.LoadBalancerWeight = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetMemory(v int32) *CreateScalingConfigurationRequest {
	s.Memory = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetOwnerAccount(v string) *CreateScalingConfigurationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetOwnerId(v int64) *CreateScalingConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetPassword(v string) *CreateScalingConfigurationRequest {
	s.Password = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetPasswordInherit(v bool) *CreateScalingConfigurationRequest {
	s.PasswordInherit = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetRamRoleName(v string) *CreateScalingConfigurationRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetResourceGroupId(v string) *CreateScalingConfigurationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetResourceOwnerAccount(v string) *CreateScalingConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetScalingConfigurationName(v string) *CreateScalingConfigurationRequest {
	s.ScalingConfigurationName = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetScalingGroupId(v string) *CreateScalingConfigurationRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSchedulerOptions(v map[string]interface{}) *CreateScalingConfigurationRequest {
	s.SchedulerOptions = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSecurityEnhancementStrategy(v string) *CreateScalingConfigurationRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSecurityGroupId(v string) *CreateScalingConfigurationRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSecurityGroupIds(v []*string) *CreateScalingConfigurationRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSpotDuration(v int32) *CreateScalingConfigurationRequest {
	s.SpotDuration = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSpotInterruptionBehavior(v string) *CreateScalingConfigurationRequest {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSpotPriceLimits(v []*CreateScalingConfigurationRequestSpotPriceLimits) *CreateScalingConfigurationRequest {
	s.SpotPriceLimits = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSpotStrategy(v string) *CreateScalingConfigurationRequest {
	s.SpotStrategy = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSystemDiskCategories(v []*string) *CreateScalingConfigurationRequest {
	s.SystemDiskCategories = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetTags(v string) *CreateScalingConfigurationRequest {
	s.Tags = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetTenancy(v string) *CreateScalingConfigurationRequest {
	s.Tenancy = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetUserData(v string) *CreateScalingConfigurationRequest {
	s.UserData = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetZoneId(v string) *CreateScalingConfigurationRequest {
	s.ZoneId = &v
	return s
}

type CreateScalingConfigurationRequestPrivatePoolOptions struct {
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
}

func (s CreateScalingConfigurationRequestPrivatePoolOptions) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestPrivatePoolOptions) SetId(v string) *CreateScalingConfigurationRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *CreateScalingConfigurationRequestPrivatePoolOptions) SetMatchCriteria(v string) *CreateScalingConfigurationRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

type CreateScalingConfigurationRequestSystemDisk struct {
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	BurstingEnabled      *bool   `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DiskName             *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	EncryptAlgorithm     *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	Encrypted            *bool   `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	KMSKeyId             *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	PerformanceLevel     *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	ProvisionedIops      *int64  `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	Size                 *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateScalingConfigurationRequestSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetAutoSnapshotPolicyId(v string) *CreateScalingConfigurationRequestSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetBurstingEnabled(v bool) *CreateScalingConfigurationRequestSystemDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetCategory(v string) *CreateScalingConfigurationRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetDescription(v string) *CreateScalingConfigurationRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetDiskName(v string) *CreateScalingConfigurationRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetEncryptAlgorithm(v string) *CreateScalingConfigurationRequestSystemDisk {
	s.EncryptAlgorithm = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetEncrypted(v bool) *CreateScalingConfigurationRequestSystemDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetKMSKeyId(v string) *CreateScalingConfigurationRequestSystemDisk {
	s.KMSKeyId = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetPerformanceLevel(v string) *CreateScalingConfigurationRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetProvisionedIops(v int64) *CreateScalingConfigurationRequestSystemDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetSize(v int32) *CreateScalingConfigurationRequestSystemDisk {
	s.Size = &v
	return s
}

type CreateScalingConfigurationRequestDataDisks struct {
	AutoSnapshotPolicyId *string   `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	BurstingEnabled      *bool     `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	Categories           []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	Category             *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteWithInstance   *bool     `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	Description          *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Device               *string   `json:"Device,omitempty" xml:"Device,omitempty"`
	DiskName             *string   `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	Encrypted            *string   `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	KMSKeyId             *string   `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	PerformanceLevel     *string   `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	ProvisionedIops      *int64    `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	Size                 *int32    `json:"Size,omitempty" xml:"Size,omitempty"`
	SnapshotId           *string   `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateScalingConfigurationRequestDataDisks) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationRequestDataDisks) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestDataDisks) SetAutoSnapshotPolicyId(v string) *CreateScalingConfigurationRequestDataDisks {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetBurstingEnabled(v bool) *CreateScalingConfigurationRequestDataDisks {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetCategories(v []*string) *CreateScalingConfigurationRequestDataDisks {
	s.Categories = v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetCategory(v string) *CreateScalingConfigurationRequestDataDisks {
	s.Category = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetDeleteWithInstance(v bool) *CreateScalingConfigurationRequestDataDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetDescription(v string) *CreateScalingConfigurationRequestDataDisks {
	s.Description = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetDevice(v string) *CreateScalingConfigurationRequestDataDisks {
	s.Device = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetDiskName(v string) *CreateScalingConfigurationRequestDataDisks {
	s.DiskName = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetEncrypted(v string) *CreateScalingConfigurationRequestDataDisks {
	s.Encrypted = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetKMSKeyId(v string) *CreateScalingConfigurationRequestDataDisks {
	s.KMSKeyId = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetPerformanceLevel(v string) *CreateScalingConfigurationRequestDataDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetProvisionedIops(v int64) *CreateScalingConfigurationRequestDataDisks {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetSize(v int32) *CreateScalingConfigurationRequestDataDisks {
	s.Size = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetSnapshotId(v string) *CreateScalingConfigurationRequestDataDisks {
	s.SnapshotId = &v
	return s
}

type CreateScalingConfigurationRequestInstancePatternInfos struct {
	Architectures         []*string `json:"Architectures,omitempty" xml:"Architectures,omitempty" type:"Repeated"`
	BurstablePerformance  *string   `json:"BurstablePerformance,omitempty" xml:"BurstablePerformance,omitempty"`
	Cores                 *int32    `json:"Cores,omitempty" xml:"Cores,omitempty"`
	ExcludedInstanceTypes []*string `json:"ExcludedInstanceTypes,omitempty" xml:"ExcludedInstanceTypes,omitempty" type:"Repeated"`
	InstanceFamilyLevel   *string   `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	MaxPrice              *float32  `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	Memory                *float32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s CreateScalingConfigurationRequestInstancePatternInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationRequestInstancePatternInfos) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetArchitectures(v []*string) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.Architectures = v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetBurstablePerformance(v string) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.BurstablePerformance = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetCores(v int32) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.Cores = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetExcludedInstanceTypes(v []*string) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.ExcludedInstanceTypes = v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetInstanceFamilyLevel(v string) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetMaxPrice(v float32) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.MaxPrice = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetMemory(v float32) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.Memory = &v
	return s
}

type CreateScalingConfigurationRequestInstanceTypeOverrides struct {
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	WeightedCapacity *int32  `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s CreateScalingConfigurationRequestInstanceTypeOverrides) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationRequestInstanceTypeOverrides) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestInstanceTypeOverrides) SetInstanceType(v string) *CreateScalingConfigurationRequestInstanceTypeOverrides {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstanceTypeOverrides) SetWeightedCapacity(v int32) *CreateScalingConfigurationRequestInstanceTypeOverrides {
	s.WeightedCapacity = &v
	return s
}

type CreateScalingConfigurationRequestSpotPriceLimits struct {
	InstanceType *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	PriceLimit   *float32 `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
}

func (s CreateScalingConfigurationRequestSpotPriceLimits) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationRequestSpotPriceLimits) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestSpotPriceLimits) SetInstanceType(v string) *CreateScalingConfigurationRequestSpotPriceLimits {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationRequestSpotPriceLimits) SetPriceLimit(v float32) *CreateScalingConfigurationRequestSpotPriceLimits {
	s.PriceLimit = &v
	return s
}

type CreateScalingConfigurationShrinkRequest struct {
	PrivatePoolOptions          *CreateScalingConfigurationShrinkRequestPrivatePoolOptions      `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	SystemDisk                  *CreateScalingConfigurationShrinkRequestSystemDisk              `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	Affinity                    *string                                                         `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	ClientToken                 *string                                                         `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Cpu                         *int32                                                          `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreditSpecification         *string                                                         `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	DataDisks                   []*CreateScalingConfigurationShrinkRequestDataDisks             `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	DedicatedHostId             *string                                                         `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	DeploymentSetId             *string                                                         `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	HostName                    *string                                                         `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HpcClusterId                *string                                                         `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	ImageFamily                 *string                                                         `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	ImageId                     *string                                                         `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName                   *string                                                         `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	InstanceDescription         *string                                                         `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	InstanceName                *string                                                         `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstancePatternInfos        []*CreateScalingConfigurationShrinkRequestInstancePatternInfos  `json:"InstancePatternInfos,omitempty" xml:"InstancePatternInfos,omitempty" type:"Repeated"`
	InstanceType                *string                                                         `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceTypeOverrides       []*CreateScalingConfigurationShrinkRequestInstanceTypeOverrides `json:"InstanceTypeOverrides,omitempty" xml:"InstanceTypeOverrides,omitempty" type:"Repeated"`
	InstanceTypes               []*string                                                       `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	InternetChargeType          *string                                                         `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandwidthIn      *int32                                                          `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	InternetMaxBandwidthOut     *int32                                                          `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	IoOptimized                 *string                                                         `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	Ipv6AddressCount            *int32                                                          `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	KeyPairName                 *string                                                         `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	LoadBalancerWeight          *int32                                                          `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	Memory                      *int32                                                          `json:"Memory,omitempty" xml:"Memory,omitempty"`
	OwnerAccount                *string                                                         `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                     *int64                                                          `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Password                    *string                                                         `json:"Password,omitempty" xml:"Password,omitempty"`
	PasswordInherit             *bool                                                           `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	RamRoleName                 *string                                                         `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	ResourceGroupId             *string                                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount        *string                                                         `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingConfigurationName    *string                                                         `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	ScalingGroupId              *string                                                         `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	SchedulerOptionsShrink      *string                                                         `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty"`
	SecurityEnhancementStrategy *string                                                         `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	SecurityGroupId             *string                                                         `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupIds            []*string                                                       `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	SpotDuration                *int32                                                          `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	SpotInterruptionBehavior    *string                                                         `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	SpotPriceLimits             []*CreateScalingConfigurationShrinkRequestSpotPriceLimits       `json:"SpotPriceLimits,omitempty" xml:"SpotPriceLimits,omitempty" type:"Repeated"`
	SpotStrategy                *string                                                         `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDiskCategories        []*string                                                       `json:"SystemDiskCategories,omitempty" xml:"SystemDiskCategories,omitempty" type:"Repeated"`
	Tags                        *string                                                         `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Tenancy                     *string                                                         `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	UserData                    *string                                                         `json:"UserData,omitempty" xml:"UserData,omitempty"`
	ZoneId                      *string                                                         `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateScalingConfigurationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequest) SetPrivatePoolOptions(v *CreateScalingConfigurationShrinkRequestPrivatePoolOptions) *CreateScalingConfigurationShrinkRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSystemDisk(v *CreateScalingConfigurationShrinkRequestSystemDisk) *CreateScalingConfigurationShrinkRequest {
	s.SystemDisk = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetAffinity(v string) *CreateScalingConfigurationShrinkRequest {
	s.Affinity = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetClientToken(v string) *CreateScalingConfigurationShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetCpu(v int32) *CreateScalingConfigurationShrinkRequest {
	s.Cpu = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetCreditSpecification(v string) *CreateScalingConfigurationShrinkRequest {
	s.CreditSpecification = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetDataDisks(v []*CreateScalingConfigurationShrinkRequestDataDisks) *CreateScalingConfigurationShrinkRequest {
	s.DataDisks = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetDedicatedHostId(v string) *CreateScalingConfigurationShrinkRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetDeploymentSetId(v string) *CreateScalingConfigurationShrinkRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetHostName(v string) *CreateScalingConfigurationShrinkRequest {
	s.HostName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetHpcClusterId(v string) *CreateScalingConfigurationShrinkRequest {
	s.HpcClusterId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetImageFamily(v string) *CreateScalingConfigurationShrinkRequest {
	s.ImageFamily = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetImageId(v string) *CreateScalingConfigurationShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetImageName(v string) *CreateScalingConfigurationShrinkRequest {
	s.ImageName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInstanceDescription(v string) *CreateScalingConfigurationShrinkRequest {
	s.InstanceDescription = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInstanceName(v string) *CreateScalingConfigurationShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInstancePatternInfos(v []*CreateScalingConfigurationShrinkRequestInstancePatternInfos) *CreateScalingConfigurationShrinkRequest {
	s.InstancePatternInfos = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInstanceType(v string) *CreateScalingConfigurationShrinkRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInstanceTypeOverrides(v []*CreateScalingConfigurationShrinkRequestInstanceTypeOverrides) *CreateScalingConfigurationShrinkRequest {
	s.InstanceTypeOverrides = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInstanceTypes(v []*string) *CreateScalingConfigurationShrinkRequest {
	s.InstanceTypes = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInternetChargeType(v string) *CreateScalingConfigurationShrinkRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInternetMaxBandwidthIn(v int32) *CreateScalingConfigurationShrinkRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInternetMaxBandwidthOut(v int32) *CreateScalingConfigurationShrinkRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetIoOptimized(v string) *CreateScalingConfigurationShrinkRequest {
	s.IoOptimized = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetIpv6AddressCount(v int32) *CreateScalingConfigurationShrinkRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetKeyPairName(v string) *CreateScalingConfigurationShrinkRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetLoadBalancerWeight(v int32) *CreateScalingConfigurationShrinkRequest {
	s.LoadBalancerWeight = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetMemory(v int32) *CreateScalingConfigurationShrinkRequest {
	s.Memory = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetOwnerAccount(v string) *CreateScalingConfigurationShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetOwnerId(v int64) *CreateScalingConfigurationShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetPassword(v string) *CreateScalingConfigurationShrinkRequest {
	s.Password = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetPasswordInherit(v bool) *CreateScalingConfigurationShrinkRequest {
	s.PasswordInherit = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetRamRoleName(v string) *CreateScalingConfigurationShrinkRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetResourceGroupId(v string) *CreateScalingConfigurationShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetResourceOwnerAccount(v string) *CreateScalingConfigurationShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetScalingConfigurationName(v string) *CreateScalingConfigurationShrinkRequest {
	s.ScalingConfigurationName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetScalingGroupId(v string) *CreateScalingConfigurationShrinkRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSchedulerOptionsShrink(v string) *CreateScalingConfigurationShrinkRequest {
	s.SchedulerOptionsShrink = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSecurityEnhancementStrategy(v string) *CreateScalingConfigurationShrinkRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSecurityGroupId(v string) *CreateScalingConfigurationShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSecurityGroupIds(v []*string) *CreateScalingConfigurationShrinkRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSpotDuration(v int32) *CreateScalingConfigurationShrinkRequest {
	s.SpotDuration = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSpotInterruptionBehavior(v string) *CreateScalingConfigurationShrinkRequest {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSpotPriceLimits(v []*CreateScalingConfigurationShrinkRequestSpotPriceLimits) *CreateScalingConfigurationShrinkRequest {
	s.SpotPriceLimits = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSpotStrategy(v string) *CreateScalingConfigurationShrinkRequest {
	s.SpotStrategy = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSystemDiskCategories(v []*string) *CreateScalingConfigurationShrinkRequest {
	s.SystemDiskCategories = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetTags(v string) *CreateScalingConfigurationShrinkRequest {
	s.Tags = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetTenancy(v string) *CreateScalingConfigurationShrinkRequest {
	s.Tenancy = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetUserData(v string) *CreateScalingConfigurationShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetZoneId(v string) *CreateScalingConfigurationShrinkRequest {
	s.ZoneId = &v
	return s
}

type CreateScalingConfigurationShrinkRequestPrivatePoolOptions struct {
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
}

func (s CreateScalingConfigurationShrinkRequestPrivatePoolOptions) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestPrivatePoolOptions) SetId(v string) *CreateScalingConfigurationShrinkRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestPrivatePoolOptions) SetMatchCriteria(v string) *CreateScalingConfigurationShrinkRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

type CreateScalingConfigurationShrinkRequestSystemDisk struct {
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	BurstingEnabled      *bool   `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DiskName             *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	EncryptAlgorithm     *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	Encrypted            *bool   `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	KMSKeyId             *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	PerformanceLevel     *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	ProvisionedIops      *int64  `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	Size                 *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateScalingConfigurationShrinkRequestSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetAutoSnapshotPolicyId(v string) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetBurstingEnabled(v bool) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetCategory(v string) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetDescription(v string) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetDiskName(v string) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetEncryptAlgorithm(v string) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.EncryptAlgorithm = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetEncrypted(v bool) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetKMSKeyId(v string) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.KMSKeyId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetPerformanceLevel(v string) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetProvisionedIops(v int64) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetSize(v int32) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.Size = &v
	return s
}

type CreateScalingConfigurationShrinkRequestDataDisks struct {
	AutoSnapshotPolicyId *string   `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	BurstingEnabled      *bool     `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	Categories           []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	Category             *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteWithInstance   *bool     `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	Description          *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Device               *string   `json:"Device,omitempty" xml:"Device,omitempty"`
	DiskName             *string   `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	Encrypted            *string   `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	KMSKeyId             *string   `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	PerformanceLevel     *string   `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	ProvisionedIops      *int64    `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	Size                 *int32    `json:"Size,omitempty" xml:"Size,omitempty"`
	SnapshotId           *string   `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateScalingConfigurationShrinkRequestDataDisks) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestDataDisks) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetAutoSnapshotPolicyId(v string) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetBurstingEnabled(v bool) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetCategories(v []*string) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.Categories = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetCategory(v string) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.Category = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetDeleteWithInstance(v bool) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetDescription(v string) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.Description = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetDevice(v string) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.Device = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetDiskName(v string) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.DiskName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetEncrypted(v string) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.Encrypted = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetKMSKeyId(v string) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.KMSKeyId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetPerformanceLevel(v string) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetProvisionedIops(v int64) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetSize(v int32) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.Size = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetSnapshotId(v string) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.SnapshotId = &v
	return s
}

type CreateScalingConfigurationShrinkRequestInstancePatternInfos struct {
	Architectures         []*string `json:"Architectures,omitempty" xml:"Architectures,omitempty" type:"Repeated"`
	BurstablePerformance  *string   `json:"BurstablePerformance,omitempty" xml:"BurstablePerformance,omitempty"`
	Cores                 *int32    `json:"Cores,omitempty" xml:"Cores,omitempty"`
	ExcludedInstanceTypes []*string `json:"ExcludedInstanceTypes,omitempty" xml:"ExcludedInstanceTypes,omitempty" type:"Repeated"`
	InstanceFamilyLevel   *string   `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	MaxPrice              *float32  `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	Memory                *float32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s CreateScalingConfigurationShrinkRequestInstancePatternInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestInstancePatternInfos) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetArchitectures(v []*string) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.Architectures = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetBurstablePerformance(v string) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.BurstablePerformance = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetCores(v int32) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.Cores = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetExcludedInstanceTypes(v []*string) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.ExcludedInstanceTypes = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetInstanceFamilyLevel(v string) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetMaxPrice(v float32) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MaxPrice = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetMemory(v float32) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.Memory = &v
	return s
}

type CreateScalingConfigurationShrinkRequestInstanceTypeOverrides struct {
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	WeightedCapacity *int32  `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s CreateScalingConfigurationShrinkRequestInstanceTypeOverrides) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestInstanceTypeOverrides) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestInstanceTypeOverrides) SetInstanceType(v string) *CreateScalingConfigurationShrinkRequestInstanceTypeOverrides {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstanceTypeOverrides) SetWeightedCapacity(v int32) *CreateScalingConfigurationShrinkRequestInstanceTypeOverrides {
	s.WeightedCapacity = &v
	return s
}

type CreateScalingConfigurationShrinkRequestSpotPriceLimits struct {
	InstanceType *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	PriceLimit   *float32 `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
}

func (s CreateScalingConfigurationShrinkRequestSpotPriceLimits) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestSpotPriceLimits) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestSpotPriceLimits) SetInstanceType(v string) *CreateScalingConfigurationShrinkRequestSpotPriceLimits {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSpotPriceLimits) SetPriceLimit(v float32) *CreateScalingConfigurationShrinkRequestSpotPriceLimits {
	s.PriceLimit = &v
	return s
}

type CreateScalingConfigurationResponseBody struct {
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
}

func (s CreateScalingConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationResponseBody) SetRequestId(v string) *CreateScalingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScalingConfigurationResponseBody) SetScalingConfigurationId(v string) *CreateScalingConfigurationResponseBody {
	s.ScalingConfigurationId = &v
	return s
}

type CreateScalingConfigurationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateScalingConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationResponse) SetHeaders(v map[string]*string) *CreateScalingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *CreateScalingConfigurationResponse) SetStatusCode(v int32) *CreateScalingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScalingConfigurationResponse) SetBody(v *CreateScalingConfigurationResponseBody) *CreateScalingConfigurationResponse {
	s.Body = v
	return s
}

type CreateScalingGroupRequest struct {
	AlbServerGroups                     []*CreateScalingGroupRequestAlbServerGroups         `json:"AlbServerGroups,omitempty" xml:"AlbServerGroups,omitempty" type:"Repeated"`
	AllocationStrategy                  *string                                             `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	AzBalance                           *bool                                               `json:"AzBalance,omitempty" xml:"AzBalance,omitempty"`
	ClientToken                         *string                                             `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CompensateWithOnDemand              *bool                                               `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	ContainerGroupId                    *string                                             `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	DBInstanceIds                       *string                                             `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	DefaultCooldown                     *int32                                              `json:"DefaultCooldown,omitempty" xml:"DefaultCooldown,omitempty"`
	DesiredCapacity                     *int32                                              `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	GroupDeletionProtection             *bool                                               `json:"GroupDeletionProtection,omitempty" xml:"GroupDeletionProtection,omitempty"`
	GroupType                           *string                                             `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	HealthCheckType                     *string                                             `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	InstanceId                          *string                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LaunchTemplateId                    *string                                             `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	LaunchTemplateOverrides             []*CreateScalingGroupRequestLaunchTemplateOverrides `json:"LaunchTemplateOverrides,omitempty" xml:"LaunchTemplateOverrides,omitempty" type:"Repeated"`
	LaunchTemplateVersion               *string                                             `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	LifecycleHooks                      []*CreateScalingGroupRequestLifecycleHooks          `json:"LifecycleHooks,omitempty" xml:"LifecycleHooks,omitempty" type:"Repeated"`
	LoadBalancerIds                     *string                                             `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty"`
	MaxInstanceLifetime                 *int32                                              `json:"MaxInstanceLifetime,omitempty" xml:"MaxInstanceLifetime,omitempty"`
	MaxSize                             *int32                                              `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	MinSize                             *int32                                              `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	MultiAZPolicy                       *string                                             `json:"MultiAZPolicy,omitempty" xml:"MultiAZPolicy,omitempty"`
	OnDemandBaseCapacity                *int32                                              `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	OnDemandPercentageAboveBaseCapacity *int32                                              `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	OwnerAccount                        *string                                             `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                             *int64                                              `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                            *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RemovalPolicies                     []*string                                           `json:"RemovalPolicies,omitempty" xml:"RemovalPolicies,omitempty" type:"Repeated"`
	ResourceOwnerAccount                *string                                             `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupName                    *string                                             `json:"ScalingGroupName,omitempty" xml:"ScalingGroupName,omitempty"`
	ScalingPolicy                       *string                                             `json:"ScalingPolicy,omitempty" xml:"ScalingPolicy,omitempty"`
	SpotAllocationStrategy              *string                                             `json:"SpotAllocationStrategy,omitempty" xml:"SpotAllocationStrategy,omitempty"`
	SpotInstancePools                   *int32                                              `json:"SpotInstancePools,omitempty" xml:"SpotInstancePools,omitempty"`
	SpotInstanceRemedy                  *bool                                               `json:"SpotInstanceRemedy,omitempty" xml:"SpotInstanceRemedy,omitempty"`
	SyncAlarmRuleToCms                  *bool                                               `json:"SyncAlarmRuleToCms,omitempty" xml:"SyncAlarmRuleToCms,omitempty"`
	Tags                                []*CreateScalingGroupRequestTags                    `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	VServerGroups                       []*CreateScalingGroupRequestVServerGroups           `json:"VServerGroups,omitempty" xml:"VServerGroups,omitempty" type:"Repeated"`
	VSwitchId                           *string                                             `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VSwitchIds                          []*string                                           `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s CreateScalingGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequest) SetAlbServerGroups(v []*CreateScalingGroupRequestAlbServerGroups) *CreateScalingGroupRequest {
	s.AlbServerGroups = v
	return s
}

func (s *CreateScalingGroupRequest) SetAllocationStrategy(v string) *CreateScalingGroupRequest {
	s.AllocationStrategy = &v
	return s
}

func (s *CreateScalingGroupRequest) SetAzBalance(v bool) *CreateScalingGroupRequest {
	s.AzBalance = &v
	return s
}

func (s *CreateScalingGroupRequest) SetClientToken(v string) *CreateScalingGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateScalingGroupRequest) SetCompensateWithOnDemand(v bool) *CreateScalingGroupRequest {
	s.CompensateWithOnDemand = &v
	return s
}

func (s *CreateScalingGroupRequest) SetContainerGroupId(v string) *CreateScalingGroupRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetDBInstanceIds(v string) *CreateScalingGroupRequest {
	s.DBInstanceIds = &v
	return s
}

func (s *CreateScalingGroupRequest) SetDefaultCooldown(v int32) *CreateScalingGroupRequest {
	s.DefaultCooldown = &v
	return s
}

func (s *CreateScalingGroupRequest) SetDesiredCapacity(v int32) *CreateScalingGroupRequest {
	s.DesiredCapacity = &v
	return s
}

func (s *CreateScalingGroupRequest) SetGroupDeletionProtection(v bool) *CreateScalingGroupRequest {
	s.GroupDeletionProtection = &v
	return s
}

func (s *CreateScalingGroupRequest) SetGroupType(v string) *CreateScalingGroupRequest {
	s.GroupType = &v
	return s
}

func (s *CreateScalingGroupRequest) SetHealthCheckType(v string) *CreateScalingGroupRequest {
	s.HealthCheckType = &v
	return s
}

func (s *CreateScalingGroupRequest) SetInstanceId(v string) *CreateScalingGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetLaunchTemplateId(v string) *CreateScalingGroupRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetLaunchTemplateOverrides(v []*CreateScalingGroupRequestLaunchTemplateOverrides) *CreateScalingGroupRequest {
	s.LaunchTemplateOverrides = v
	return s
}

func (s *CreateScalingGroupRequest) SetLaunchTemplateVersion(v string) *CreateScalingGroupRequest {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *CreateScalingGroupRequest) SetLifecycleHooks(v []*CreateScalingGroupRequestLifecycleHooks) *CreateScalingGroupRequest {
	s.LifecycleHooks = v
	return s
}

func (s *CreateScalingGroupRequest) SetLoadBalancerIds(v string) *CreateScalingGroupRequest {
	s.LoadBalancerIds = &v
	return s
}

func (s *CreateScalingGroupRequest) SetMaxInstanceLifetime(v int32) *CreateScalingGroupRequest {
	s.MaxInstanceLifetime = &v
	return s
}

func (s *CreateScalingGroupRequest) SetMaxSize(v int32) *CreateScalingGroupRequest {
	s.MaxSize = &v
	return s
}

func (s *CreateScalingGroupRequest) SetMinSize(v int32) *CreateScalingGroupRequest {
	s.MinSize = &v
	return s
}

func (s *CreateScalingGroupRequest) SetMultiAZPolicy(v string) *CreateScalingGroupRequest {
	s.MultiAZPolicy = &v
	return s
}

func (s *CreateScalingGroupRequest) SetOnDemandBaseCapacity(v int32) *CreateScalingGroupRequest {
	s.OnDemandBaseCapacity = &v
	return s
}

func (s *CreateScalingGroupRequest) SetOnDemandPercentageAboveBaseCapacity(v int32) *CreateScalingGroupRequest {
	s.OnDemandPercentageAboveBaseCapacity = &v
	return s
}

func (s *CreateScalingGroupRequest) SetOwnerAccount(v string) *CreateScalingGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateScalingGroupRequest) SetOwnerId(v int64) *CreateScalingGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetRegionId(v string) *CreateScalingGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetRemovalPolicies(v []*string) *CreateScalingGroupRequest {
	s.RemovalPolicies = v
	return s
}

func (s *CreateScalingGroupRequest) SetResourceOwnerAccount(v string) *CreateScalingGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateScalingGroupRequest) SetScalingGroupName(v string) *CreateScalingGroupRequest {
	s.ScalingGroupName = &v
	return s
}

func (s *CreateScalingGroupRequest) SetScalingPolicy(v string) *CreateScalingGroupRequest {
	s.ScalingPolicy = &v
	return s
}

func (s *CreateScalingGroupRequest) SetSpotAllocationStrategy(v string) *CreateScalingGroupRequest {
	s.SpotAllocationStrategy = &v
	return s
}

func (s *CreateScalingGroupRequest) SetSpotInstancePools(v int32) *CreateScalingGroupRequest {
	s.SpotInstancePools = &v
	return s
}

func (s *CreateScalingGroupRequest) SetSpotInstanceRemedy(v bool) *CreateScalingGroupRequest {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *CreateScalingGroupRequest) SetSyncAlarmRuleToCms(v bool) *CreateScalingGroupRequest {
	s.SyncAlarmRuleToCms = &v
	return s
}

func (s *CreateScalingGroupRequest) SetTags(v []*CreateScalingGroupRequestTags) *CreateScalingGroupRequest {
	s.Tags = v
	return s
}

func (s *CreateScalingGroupRequest) SetVServerGroups(v []*CreateScalingGroupRequestVServerGroups) *CreateScalingGroupRequest {
	s.VServerGroups = v
	return s
}

func (s *CreateScalingGroupRequest) SetVSwitchId(v string) *CreateScalingGroupRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetVSwitchIds(v []*string) *CreateScalingGroupRequest {
	s.VSwitchIds = v
	return s
}

type CreateScalingGroupRequestAlbServerGroups struct {
	AlbServerGroupId *string `json:"AlbServerGroupId,omitempty" xml:"AlbServerGroupId,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Weight           *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateScalingGroupRequestAlbServerGroups) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupRequestAlbServerGroups) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestAlbServerGroups) SetAlbServerGroupId(v string) *CreateScalingGroupRequestAlbServerGroups {
	s.AlbServerGroupId = &v
	return s
}

func (s *CreateScalingGroupRequestAlbServerGroups) SetPort(v int32) *CreateScalingGroupRequestAlbServerGroups {
	s.Port = &v
	return s
}

func (s *CreateScalingGroupRequestAlbServerGroups) SetWeight(v int32) *CreateScalingGroupRequestAlbServerGroups {
	s.Weight = &v
	return s
}

type CreateScalingGroupRequestLaunchTemplateOverrides struct {
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	WeightedCapacity *int32  `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s CreateScalingGroupRequestLaunchTemplateOverrides) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupRequestLaunchTemplateOverrides) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestLaunchTemplateOverrides) SetInstanceType(v string) *CreateScalingGroupRequestLaunchTemplateOverrides {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingGroupRequestLaunchTemplateOverrides) SetWeightedCapacity(v int32) *CreateScalingGroupRequestLaunchTemplateOverrides {
	s.WeightedCapacity = &v
	return s
}

type CreateScalingGroupRequestLifecycleHooks struct {
	DefaultResult        *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	HeartbeatTimeout     *int32  `json:"HeartbeatTimeout,omitempty" xml:"HeartbeatTimeout,omitempty"`
	LifecycleHookName    *string `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	LifecycleTransition  *string `json:"LifecycleTransition,omitempty" xml:"LifecycleTransition,omitempty"`
	NotificationArn      *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	NotificationMetadata *string `json:"NotificationMetadata,omitempty" xml:"NotificationMetadata,omitempty"`
}

func (s CreateScalingGroupRequestLifecycleHooks) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupRequestLifecycleHooks) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestLifecycleHooks) SetDefaultResult(v string) *CreateScalingGroupRequestLifecycleHooks {
	s.DefaultResult = &v
	return s
}

func (s *CreateScalingGroupRequestLifecycleHooks) SetHeartbeatTimeout(v int32) *CreateScalingGroupRequestLifecycleHooks {
	s.HeartbeatTimeout = &v
	return s
}

func (s *CreateScalingGroupRequestLifecycleHooks) SetLifecycleHookName(v string) *CreateScalingGroupRequestLifecycleHooks {
	s.LifecycleHookName = &v
	return s
}

func (s *CreateScalingGroupRequestLifecycleHooks) SetLifecycleTransition(v string) *CreateScalingGroupRequestLifecycleHooks {
	s.LifecycleTransition = &v
	return s
}

func (s *CreateScalingGroupRequestLifecycleHooks) SetNotificationArn(v string) *CreateScalingGroupRequestLifecycleHooks {
	s.NotificationArn = &v
	return s
}

func (s *CreateScalingGroupRequestLifecycleHooks) SetNotificationMetadata(v string) *CreateScalingGroupRequestLifecycleHooks {
	s.NotificationMetadata = &v
	return s
}

type CreateScalingGroupRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateScalingGroupRequestTags) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupRequestTags) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestTags) SetKey(v string) *CreateScalingGroupRequestTags {
	s.Key = &v
	return s
}

func (s *CreateScalingGroupRequestTags) SetValue(v string) *CreateScalingGroupRequestTags {
	s.Value = &v
	return s
}

type CreateScalingGroupRequestVServerGroups struct {
	LoadBalancerId         *string                                                         `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	VServerGroupAttributes []*CreateScalingGroupRequestVServerGroupsVServerGroupAttributes `json:"VServerGroupAttributes,omitempty" xml:"VServerGroupAttributes,omitempty" type:"Repeated"`
}

func (s CreateScalingGroupRequestVServerGroups) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupRequestVServerGroups) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestVServerGroups) SetLoadBalancerId(v string) *CreateScalingGroupRequestVServerGroups {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateScalingGroupRequestVServerGroups) SetVServerGroupAttributes(v []*CreateScalingGroupRequestVServerGroupsVServerGroupAttributes) *CreateScalingGroupRequestVServerGroups {
	s.VServerGroupAttributes = v
	return s
}

type CreateScalingGroupRequestVServerGroupsVServerGroupAttributes struct {
	Port           *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	Weight         *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateScalingGroupRequestVServerGroupsVServerGroupAttributes) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupRequestVServerGroupsVServerGroupAttributes) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestVServerGroupsVServerGroupAttributes) SetPort(v int32) *CreateScalingGroupRequestVServerGroupsVServerGroupAttributes {
	s.Port = &v
	return s
}

func (s *CreateScalingGroupRequestVServerGroupsVServerGroupAttributes) SetVServerGroupId(v string) *CreateScalingGroupRequestVServerGroupsVServerGroupAttributes {
	s.VServerGroupId = &v
	return s
}

func (s *CreateScalingGroupRequestVServerGroupsVServerGroupAttributes) SetWeight(v int32) *CreateScalingGroupRequestVServerGroupsVServerGroupAttributes {
	s.Weight = &v
	return s
}

type CreateScalingGroupResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s CreateScalingGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupResponseBody) SetRequestId(v string) *CreateScalingGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScalingGroupResponseBody) SetScalingGroupId(v string) *CreateScalingGroupResponseBody {
	s.ScalingGroupId = &v
	return s
}

type CreateScalingGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateScalingGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupResponse) SetHeaders(v map[string]*string) *CreateScalingGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateScalingGroupResponse) SetStatusCode(v int32) *CreateScalingGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScalingGroupResponse) SetBody(v *CreateScalingGroupResponseBody) *CreateScalingGroupResponse {
	s.Body = v
	return s
}

type CreateScalingRuleRequest struct {
	AdjustmentType           *string                                    `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	AdjustmentValue          *int32                                     `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	Cooldown                 *int32                                     `json:"Cooldown,omitempty" xml:"Cooldown,omitempty"`
	DisableScaleIn           *bool                                      `json:"DisableScaleIn,omitempty" xml:"DisableScaleIn,omitempty"`
	EstimatedInstanceWarmup  *int32                                     `json:"EstimatedInstanceWarmup,omitempty" xml:"EstimatedInstanceWarmup,omitempty"`
	InitialMaxSize           *int32                                     `json:"InitialMaxSize,omitempty" xml:"InitialMaxSize,omitempty"`
	MetricName               *string                                    `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	MinAdjustmentMagnitude   *int32                                     `json:"MinAdjustmentMagnitude,omitempty" xml:"MinAdjustmentMagnitude,omitempty"`
	OwnerAccount             *string                                    `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64                                     `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PredictiveScalingMode    *string                                    `json:"PredictiveScalingMode,omitempty" xml:"PredictiveScalingMode,omitempty"`
	PredictiveTaskBufferTime *int32                                     `json:"PredictiveTaskBufferTime,omitempty" xml:"PredictiveTaskBufferTime,omitempty"`
	PredictiveValueBehavior  *string                                    `json:"PredictiveValueBehavior,omitempty" xml:"PredictiveValueBehavior,omitempty"`
	PredictiveValueBuffer    *int32                                     `json:"PredictiveValueBuffer,omitempty" xml:"PredictiveValueBuffer,omitempty"`
	RegionId                 *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount     *string                                    `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScaleInEvaluationCount   *int32                                     `json:"ScaleInEvaluationCount,omitempty" xml:"ScaleInEvaluationCount,omitempty"`
	ScaleOutEvaluationCount  *int32                                     `json:"ScaleOutEvaluationCount,omitempty" xml:"ScaleOutEvaluationCount,omitempty"`
	ScalingGroupId           *string                                    `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScalingRuleName          *string                                    `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	ScalingRuleType          *string                                    `json:"ScalingRuleType,omitempty" xml:"ScalingRuleType,omitempty"`
	StepAdjustments          []*CreateScalingRuleRequestStepAdjustments `json:"StepAdjustments,omitempty" xml:"StepAdjustments,omitempty" type:"Repeated"`
	TargetValue              *float32                                   `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s CreateScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateScalingRuleRequest) SetAdjustmentType(v string) *CreateScalingRuleRequest {
	s.AdjustmentType = &v
	return s
}

func (s *CreateScalingRuleRequest) SetAdjustmentValue(v int32) *CreateScalingRuleRequest {
	s.AdjustmentValue = &v
	return s
}

func (s *CreateScalingRuleRequest) SetCooldown(v int32) *CreateScalingRuleRequest {
	s.Cooldown = &v
	return s
}

func (s *CreateScalingRuleRequest) SetDisableScaleIn(v bool) *CreateScalingRuleRequest {
	s.DisableScaleIn = &v
	return s
}

func (s *CreateScalingRuleRequest) SetEstimatedInstanceWarmup(v int32) *CreateScalingRuleRequest {
	s.EstimatedInstanceWarmup = &v
	return s
}

func (s *CreateScalingRuleRequest) SetInitialMaxSize(v int32) *CreateScalingRuleRequest {
	s.InitialMaxSize = &v
	return s
}

func (s *CreateScalingRuleRequest) SetMetricName(v string) *CreateScalingRuleRequest {
	s.MetricName = &v
	return s
}

func (s *CreateScalingRuleRequest) SetMinAdjustmentMagnitude(v int32) *CreateScalingRuleRequest {
	s.MinAdjustmentMagnitude = &v
	return s
}

func (s *CreateScalingRuleRequest) SetOwnerAccount(v string) *CreateScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateScalingRuleRequest) SetOwnerId(v int64) *CreateScalingRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateScalingRuleRequest) SetPredictiveScalingMode(v string) *CreateScalingRuleRequest {
	s.PredictiveScalingMode = &v
	return s
}

func (s *CreateScalingRuleRequest) SetPredictiveTaskBufferTime(v int32) *CreateScalingRuleRequest {
	s.PredictiveTaskBufferTime = &v
	return s
}

func (s *CreateScalingRuleRequest) SetPredictiveValueBehavior(v string) *CreateScalingRuleRequest {
	s.PredictiveValueBehavior = &v
	return s
}

func (s *CreateScalingRuleRequest) SetPredictiveValueBuffer(v int32) *CreateScalingRuleRequest {
	s.PredictiveValueBuffer = &v
	return s
}

func (s *CreateScalingRuleRequest) SetRegionId(v string) *CreateScalingRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateScalingRuleRequest) SetResourceOwnerAccount(v string) *CreateScalingRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateScalingRuleRequest) SetScaleInEvaluationCount(v int32) *CreateScalingRuleRequest {
	s.ScaleInEvaluationCount = &v
	return s
}

func (s *CreateScalingRuleRequest) SetScaleOutEvaluationCount(v int32) *CreateScalingRuleRequest {
	s.ScaleOutEvaluationCount = &v
	return s
}

func (s *CreateScalingRuleRequest) SetScalingGroupId(v string) *CreateScalingRuleRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *CreateScalingRuleRequest) SetScalingRuleName(v string) *CreateScalingRuleRequest {
	s.ScalingRuleName = &v
	return s
}

func (s *CreateScalingRuleRequest) SetScalingRuleType(v string) *CreateScalingRuleRequest {
	s.ScalingRuleType = &v
	return s
}

func (s *CreateScalingRuleRequest) SetStepAdjustments(v []*CreateScalingRuleRequestStepAdjustments) *CreateScalingRuleRequest {
	s.StepAdjustments = v
	return s
}

func (s *CreateScalingRuleRequest) SetTargetValue(v float32) *CreateScalingRuleRequest {
	s.TargetValue = &v
	return s
}

type CreateScalingRuleRequestStepAdjustments struct {
	MetricIntervalLowerBound *float32 `json:"MetricIntervalLowerBound,omitempty" xml:"MetricIntervalLowerBound,omitempty"`
	MetricIntervalUpperBound *float32 `json:"MetricIntervalUpperBound,omitempty" xml:"MetricIntervalUpperBound,omitempty"`
	ScalingAdjustment        *int32   `json:"ScalingAdjustment,omitempty" xml:"ScalingAdjustment,omitempty"`
}

func (s CreateScalingRuleRequestStepAdjustments) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingRuleRequestStepAdjustments) GoString() string {
	return s.String()
}

func (s *CreateScalingRuleRequestStepAdjustments) SetMetricIntervalLowerBound(v float32) *CreateScalingRuleRequestStepAdjustments {
	s.MetricIntervalLowerBound = &v
	return s
}

func (s *CreateScalingRuleRequestStepAdjustments) SetMetricIntervalUpperBound(v float32) *CreateScalingRuleRequestStepAdjustments {
	s.MetricIntervalUpperBound = &v
	return s
}

func (s *CreateScalingRuleRequestStepAdjustments) SetScalingAdjustment(v int32) *CreateScalingRuleRequestStepAdjustments {
	s.ScalingAdjustment = &v
	return s
}

type CreateScalingRuleResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingRuleAri *string `json:"ScalingRuleAri,omitempty" xml:"ScalingRuleAri,omitempty"`
	ScalingRuleId  *string `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty"`
}

func (s CreateScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScalingRuleResponseBody) SetRequestId(v string) *CreateScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScalingRuleResponseBody) SetScalingRuleAri(v string) *CreateScalingRuleResponseBody {
	s.ScalingRuleAri = &v
	return s
}

func (s *CreateScalingRuleResponseBody) SetScalingRuleId(v string) *CreateScalingRuleResponseBody {
	s.ScalingRuleId = &v
	return s
}

type CreateScalingRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateScalingRuleResponse) SetHeaders(v map[string]*string) *CreateScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateScalingRuleResponse) SetStatusCode(v int32) *CreateScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScalingRuleResponse) SetBody(v *CreateScalingRuleResponseBody) *CreateScalingRuleResponse {
	s.Body = v
	return s
}

type CreateScheduledTaskRequest struct {
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DesiredCapacity      *int32  `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	LaunchExpirationTime *int32  `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
	LaunchTime           *string `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	MaxValue             *int32  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinValue             *int32  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RecurrenceEndTime    *string `json:"RecurrenceEndTime,omitempty" xml:"RecurrenceEndTime,omitempty"`
	RecurrenceType       *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValue      *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScheduledAction      *string `json:"ScheduledAction,omitempty" xml:"ScheduledAction,omitempty"`
	ScheduledTaskName    *string `json:"ScheduledTaskName,omitempty" xml:"ScheduledTaskName,omitempty"`
	TaskEnabled          *bool   `json:"TaskEnabled,omitempty" xml:"TaskEnabled,omitempty"`
}

func (s CreateScheduledTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskRequest) SetDescription(v string) *CreateScheduledTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetDesiredCapacity(v int32) *CreateScheduledTaskRequest {
	s.DesiredCapacity = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetLaunchExpirationTime(v int32) *CreateScheduledTaskRequest {
	s.LaunchExpirationTime = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetLaunchTime(v string) *CreateScheduledTaskRequest {
	s.LaunchTime = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetMaxValue(v int32) *CreateScheduledTaskRequest {
	s.MaxValue = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetMinValue(v int32) *CreateScheduledTaskRequest {
	s.MinValue = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetOwnerAccount(v string) *CreateScheduledTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetOwnerId(v int64) *CreateScheduledTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetRecurrenceEndTime(v string) *CreateScheduledTaskRequest {
	s.RecurrenceEndTime = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetRecurrenceType(v string) *CreateScheduledTaskRequest {
	s.RecurrenceType = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetRecurrenceValue(v string) *CreateScheduledTaskRequest {
	s.RecurrenceValue = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetRegionId(v string) *CreateScheduledTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetResourceOwnerAccount(v string) *CreateScheduledTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetScalingGroupId(v string) *CreateScheduledTaskRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetScheduledAction(v string) *CreateScheduledTaskRequest {
	s.ScheduledAction = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetScheduledTaskName(v string) *CreateScheduledTaskRequest {
	s.ScheduledTaskName = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetTaskEnabled(v bool) *CreateScheduledTaskRequest {
	s.TaskEnabled = &v
	return s
}

type CreateScheduledTaskResponseBody struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScheduledTaskId *string `json:"ScheduledTaskId,omitempty" xml:"ScheduledTaskId,omitempty"`
}

func (s CreateScheduledTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskResponseBody) SetRequestId(v string) *CreateScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduledTaskResponseBody) SetScheduledTaskId(v string) *CreateScheduledTaskResponseBody {
	s.ScheduledTaskId = &v
	return s
}

type CreateScheduledTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateScheduledTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskResponse) SetHeaders(v map[string]*string) *CreateScheduledTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduledTaskResponse) SetStatusCode(v int32) *CreateScheduledTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScheduledTaskResponse) SetBody(v *CreateScheduledTaskResponseBody) *CreateScheduledTaskResponse {
	s.Body = v
	return s
}

type DeactivateScalingConfigurationRequest struct {
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
}

func (s DeactivateScalingConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeactivateScalingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeactivateScalingConfigurationRequest) SetOwnerAccount(v string) *DeactivateScalingConfigurationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeactivateScalingConfigurationRequest) SetOwnerId(v int64) *DeactivateScalingConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *DeactivateScalingConfigurationRequest) SetResourceOwnerAccount(v string) *DeactivateScalingConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeactivateScalingConfigurationRequest) SetScalingConfigurationId(v string) *DeactivateScalingConfigurationRequest {
	s.ScalingConfigurationId = &v
	return s
}

type DeactivateScalingConfigurationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeactivateScalingConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeactivateScalingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DeactivateScalingConfigurationResponseBody) SetRequestId(v string) *DeactivateScalingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type DeactivateScalingConfigurationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeactivateScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeactivateScalingConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeactivateScalingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeactivateScalingConfigurationResponse) SetHeaders(v map[string]*string) *DeactivateScalingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeactivateScalingConfigurationResponse) SetStatusCode(v int32) *DeactivateScalingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeactivateScalingConfigurationResponse) SetBody(v *DeactivateScalingConfigurationResponseBody) *DeactivateScalingConfigurationResponse {
	s.Body = v
	return s
}

type DeleteAlarmRequest struct {
	AlarmTaskId          *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DeleteAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlarmRequest) SetAlarmTaskId(v string) *DeleteAlarmRequest {
	s.AlarmTaskId = &v
	return s
}

func (s *DeleteAlarmRequest) SetOwnerId(v int64) *DeleteAlarmRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAlarmRequest) SetRegionId(v string) *DeleteAlarmRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAlarmRequest) SetResourceOwnerAccount(v string) *DeleteAlarmRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type DeleteAlarmResponseBody struct {
	AlarmTaskId *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlarmResponseBody) SetAlarmTaskId(v string) *DeleteAlarmResponseBody {
	s.AlarmTaskId = &v
	return s
}

func (s *DeleteAlarmResponseBody) SetRequestId(v string) *DeleteAlarmResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAlarmResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlarmResponse) SetHeaders(v map[string]*string) *DeleteAlarmResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlarmResponse) SetStatusCode(v int32) *DeleteAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlarmResponse) SetBody(v *DeleteAlarmResponseBody) *DeleteAlarmResponse {
	s.Body = v
	return s
}

type DeleteLifecycleHookRequest struct {
	LifecycleHookId      *string `json:"LifecycleHookId,omitempty" xml:"LifecycleHookId,omitempty"`
	LifecycleHookName    *string `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DeleteLifecycleHookRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLifecycleHookRequest) GoString() string {
	return s.String()
}

func (s *DeleteLifecycleHookRequest) SetLifecycleHookId(v string) *DeleteLifecycleHookRequest {
	s.LifecycleHookId = &v
	return s
}

func (s *DeleteLifecycleHookRequest) SetLifecycleHookName(v string) *DeleteLifecycleHookRequest {
	s.LifecycleHookName = &v
	return s
}

func (s *DeleteLifecycleHookRequest) SetOwnerAccount(v string) *DeleteLifecycleHookRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteLifecycleHookRequest) SetOwnerId(v int64) *DeleteLifecycleHookRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLifecycleHookRequest) SetRegionId(v string) *DeleteLifecycleHookRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLifecycleHookRequest) SetResourceOwnerAccount(v string) *DeleteLifecycleHookRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteLifecycleHookRequest) SetScalingGroupId(v string) *DeleteLifecycleHookRequest {
	s.ScalingGroupId = &v
	return s
}

type DeleteLifecycleHookResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLifecycleHookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLifecycleHookResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLifecycleHookResponseBody) SetRequestId(v string) *DeleteLifecycleHookResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLifecycleHookResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLifecycleHookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLifecycleHookResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLifecycleHookResponse) GoString() string {
	return s.String()
}

func (s *DeleteLifecycleHookResponse) SetHeaders(v map[string]*string) *DeleteLifecycleHookResponse {
	s.Headers = v
	return s
}

func (s *DeleteLifecycleHookResponse) SetStatusCode(v int32) *DeleteLifecycleHookResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLifecycleHookResponse) SetBody(v *DeleteLifecycleHookResponseBody) *DeleteLifecycleHookResponse {
	s.Body = v
	return s
}

type DeleteNotificationConfigurationRequest struct {
	NotificationArn      *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DeleteNotificationConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNotificationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeleteNotificationConfigurationRequest) SetNotificationArn(v string) *DeleteNotificationConfigurationRequest {
	s.NotificationArn = &v
	return s
}

func (s *DeleteNotificationConfigurationRequest) SetOwnerId(v int64) *DeleteNotificationConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteNotificationConfigurationRequest) SetRegionId(v string) *DeleteNotificationConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNotificationConfigurationRequest) SetResourceOwnerAccount(v string) *DeleteNotificationConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteNotificationConfigurationRequest) SetScalingGroupId(v string) *DeleteNotificationConfigurationRequest {
	s.ScalingGroupId = &v
	return s
}

type DeleteNotificationConfigurationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNotificationConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNotificationConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNotificationConfigurationResponseBody) SetRequestId(v string) *DeleteNotificationConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNotificationConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteNotificationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNotificationConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNotificationConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeleteNotificationConfigurationResponse) SetHeaders(v map[string]*string) *DeleteNotificationConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeleteNotificationConfigurationResponse) SetStatusCode(v int32) *DeleteNotificationConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNotificationConfigurationResponse) SetBody(v *DeleteNotificationConfigurationResponseBody) *DeleteNotificationConfigurationResponse {
	s.Body = v
	return s
}

type DeleteScalingConfigurationRequest struct {
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
}

func (s DeleteScalingConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeleteScalingConfigurationRequest) SetOwnerAccount(v string) *DeleteScalingConfigurationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteScalingConfigurationRequest) SetOwnerId(v int64) *DeleteScalingConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteScalingConfigurationRequest) SetResourceOwnerAccount(v string) *DeleteScalingConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteScalingConfigurationRequest) SetScalingConfigurationId(v string) *DeleteScalingConfigurationRequest {
	s.ScalingConfigurationId = &v
	return s
}

type DeleteScalingConfigurationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScalingConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScalingConfigurationResponseBody) SetRequestId(v string) *DeleteScalingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteScalingConfigurationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteScalingConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeleteScalingConfigurationResponse) SetHeaders(v map[string]*string) *DeleteScalingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeleteScalingConfigurationResponse) SetStatusCode(v int32) *DeleteScalingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScalingConfigurationResponse) SetBody(v *DeleteScalingConfigurationResponseBody) *DeleteScalingConfigurationResponse {
	s.Body = v
	return s
}

type DeleteScalingGroupRequest struct {
	ForceDelete          *bool   `json:"ForceDelete,omitempty" xml:"ForceDelete,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DeleteScalingGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteScalingGroupRequest) SetForceDelete(v bool) *DeleteScalingGroupRequest {
	s.ForceDelete = &v
	return s
}

func (s *DeleteScalingGroupRequest) SetOwnerAccount(v string) *DeleteScalingGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteScalingGroupRequest) SetOwnerId(v int64) *DeleteScalingGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteScalingGroupRequest) SetRegionId(v string) *DeleteScalingGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteScalingGroupRequest) SetResourceOwnerAccount(v string) *DeleteScalingGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteScalingGroupRequest) SetScalingGroupId(v string) *DeleteScalingGroupRequest {
	s.ScalingGroupId = &v
	return s
}

type DeleteScalingGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScalingGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScalingGroupResponseBody) SetRequestId(v string) *DeleteScalingGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteScalingGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteScalingGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteScalingGroupResponse) SetHeaders(v map[string]*string) *DeleteScalingGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteScalingGroupResponse) SetStatusCode(v int32) *DeleteScalingGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScalingGroupResponse) SetBody(v *DeleteScalingGroupResponseBody) *DeleteScalingGroupResponse {
	s.Body = v
	return s
}

type DeleteScalingRuleRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingRuleId        *string `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty"`
}

func (s DeleteScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteScalingRuleRequest) SetOwnerAccount(v string) *DeleteScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteScalingRuleRequest) SetOwnerId(v int64) *DeleteScalingRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteScalingRuleRequest) SetRegionId(v string) *DeleteScalingRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteScalingRuleRequest) SetResourceOwnerAccount(v string) *DeleteScalingRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteScalingRuleRequest) SetScalingRuleId(v string) *DeleteScalingRuleRequest {
	s.ScalingRuleId = &v
	return s
}

type DeleteScalingRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScalingRuleResponseBody) SetRequestId(v string) *DeleteScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteScalingRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteScalingRuleResponse) SetHeaders(v map[string]*string) *DeleteScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteScalingRuleResponse) SetStatusCode(v int32) *DeleteScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScalingRuleResponse) SetBody(v *DeleteScalingRuleResponseBody) *DeleteScalingRuleResponse {
	s.Body = v
	return s
}

type DeleteScheduledTaskRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScheduledTaskId      *string `json:"ScheduledTaskId,omitempty" xml:"ScheduledTaskId,omitempty"`
}

func (s DeleteScheduledTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduledTaskRequest) SetOwnerAccount(v string) *DeleteScheduledTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteScheduledTaskRequest) SetOwnerId(v int64) *DeleteScheduledTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteScheduledTaskRequest) SetRegionId(v string) *DeleteScheduledTaskRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteScheduledTaskRequest) SetResourceOwnerAccount(v string) *DeleteScheduledTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteScheduledTaskRequest) SetScheduledTaskId(v string) *DeleteScheduledTaskRequest {
	s.ScheduledTaskId = &v
	return s
}

type DeleteScheduledTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScheduledTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduledTaskResponseBody) SetRequestId(v string) *DeleteScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteScheduledTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteScheduledTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduledTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduledTaskResponse) SetHeaders(v map[string]*string) *DeleteScheduledTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduledTaskResponse) SetStatusCode(v int32) *DeleteScheduledTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScheduledTaskResponse) SetBody(v *DeleteScheduledTaskResponseBody) *DeleteScheduledTaskResponse {
	s.Body = v
	return s
}

type DescribeAlarmsRequest struct {
	AlarmTaskId          *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	IsEnable             *bool   `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
	MetricName           *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	MetricType           *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	State                *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeAlarmsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsRequest) SetAlarmTaskId(v string) *DescribeAlarmsRequest {
	s.AlarmTaskId = &v
	return s
}

func (s *DescribeAlarmsRequest) SetIsEnable(v bool) *DescribeAlarmsRequest {
	s.IsEnable = &v
	return s
}

func (s *DescribeAlarmsRequest) SetMetricName(v string) *DescribeAlarmsRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeAlarmsRequest) SetMetricType(v string) *DescribeAlarmsRequest {
	s.MetricType = &v
	return s
}

func (s *DescribeAlarmsRequest) SetOwnerId(v int64) *DescribeAlarmsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAlarmsRequest) SetPageNumber(v int32) *DescribeAlarmsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAlarmsRequest) SetPageSize(v int32) *DescribeAlarmsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlarmsRequest) SetRegionId(v string) *DescribeAlarmsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlarmsRequest) SetResourceOwnerAccount(v string) *DescribeAlarmsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAlarmsRequest) SetScalingGroupId(v string) *DescribeAlarmsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeAlarmsRequest) SetState(v string) *DescribeAlarmsRequest {
	s.State = &v
	return s
}

type DescribeAlarmsResponseBody struct {
	AlarmList  []*DescribeAlarmsResponseBodyAlarmList `json:"AlarmList,omitempty" xml:"AlarmList,omitempty" type:"Repeated"`
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAlarmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBody) SetAlarmList(v []*DescribeAlarmsResponseBodyAlarmList) *DescribeAlarmsResponseBody {
	s.AlarmList = v
	return s
}

func (s *DescribeAlarmsResponseBody) SetPageNumber(v int32) *DescribeAlarmsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetPageSize(v int32) *DescribeAlarmsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetRequestId(v string) *DescribeAlarmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetTotalCount(v int32) *DescribeAlarmsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAlarmsResponseBodyAlarmList struct {
	AlarmActions             []*string                                         `json:"AlarmActions,omitempty" xml:"AlarmActions,omitempty" type:"Repeated"`
	AlarmTaskId              *string                                           `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	ComparisonOperator       *string                                           `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Description              *string                                           `json:"Description,omitempty" xml:"Description,omitempty"`
	Dimensions               []*DescribeAlarmsResponseBodyAlarmListDimensions  `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	Effective                *string                                           `json:"Effective,omitempty" xml:"Effective,omitempty"`
	Enable                   *bool                                             `json:"Enable,omitempty" xml:"Enable,omitempty"`
	EvaluationCount          *int32                                            `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	Expressions              []*DescribeAlarmsResponseBodyAlarmListExpressions `json:"Expressions,omitempty" xml:"Expressions,omitempty" type:"Repeated"`
	ExpressionsLogicOperator *string                                           `json:"ExpressionsLogicOperator,omitempty" xml:"ExpressionsLogicOperator,omitempty"`
	MetricName               *string                                           `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	MetricType               *string                                           `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Name                     *string                                           `json:"Name,omitempty" xml:"Name,omitempty"`
	Period                   *int32                                            `json:"Period,omitempty" xml:"Period,omitempty"`
	ScalingGroupId           *string                                           `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	State                    *string                                           `json:"State,omitempty" xml:"State,omitempty"`
	Statistics               *string                                           `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold                *float32                                          `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeAlarmsResponseBodyAlarmList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponseBodyAlarmList) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetAlarmActions(v []*string) *DescribeAlarmsResponseBodyAlarmList {
	s.AlarmActions = v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetAlarmTaskId(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.AlarmTaskId = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetComparisonOperator(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetDescription(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.Description = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetDimensions(v []*DescribeAlarmsResponseBodyAlarmListDimensions) *DescribeAlarmsResponseBodyAlarmList {
	s.Dimensions = v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetEffective(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.Effective = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetEnable(v bool) *DescribeAlarmsResponseBodyAlarmList {
	s.Enable = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetEvaluationCount(v int32) *DescribeAlarmsResponseBodyAlarmList {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetExpressions(v []*DescribeAlarmsResponseBodyAlarmListExpressions) *DescribeAlarmsResponseBodyAlarmList {
	s.Expressions = v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetExpressionsLogicOperator(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.ExpressionsLogicOperator = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetMetricName(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.MetricName = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetMetricType(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.MetricType = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetName(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.Name = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetPeriod(v int32) *DescribeAlarmsResponseBodyAlarmList {
	s.Period = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetScalingGroupId(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetState(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.State = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetStatistics(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.Statistics = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetThreshold(v float32) *DescribeAlarmsResponseBodyAlarmList {
	s.Threshold = &v
	return s
}

type DescribeAlarmsResponseBodyAlarmListDimensions struct {
	DimensionKey   *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
}

func (s DescribeAlarmsResponseBodyAlarmListDimensions) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponseBodyAlarmListDimensions) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyAlarmListDimensions) SetDimensionKey(v string) *DescribeAlarmsResponseBodyAlarmListDimensions {
	s.DimensionKey = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListDimensions) SetDimensionValue(v string) *DescribeAlarmsResponseBodyAlarmListDimensions {
	s.DimensionValue = &v
	return s
}

type DescribeAlarmsResponseBodyAlarmListExpressions struct {
	ComparisonOperator *string  `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	MetricName         *string  `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Period             *int32   `json:"Period,omitempty" xml:"Period,omitempty"`
	Statistics         *string  `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeAlarmsResponseBodyAlarmListExpressions) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponseBodyAlarmListExpressions) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyAlarmListExpressions) SetComparisonOperator(v string) *DescribeAlarmsResponseBodyAlarmListExpressions {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListExpressions) SetMetricName(v string) *DescribeAlarmsResponseBodyAlarmListExpressions {
	s.MetricName = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListExpressions) SetPeriod(v int32) *DescribeAlarmsResponseBodyAlarmListExpressions {
	s.Period = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListExpressions) SetStatistics(v string) *DescribeAlarmsResponseBodyAlarmListExpressions {
	s.Statistics = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListExpressions) SetThreshold(v float32) *DescribeAlarmsResponseBodyAlarmListExpressions {
	s.Threshold = &v
	return s
}

type DescribeAlarmsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAlarmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAlarmsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponse) SetHeaders(v map[string]*string) *DescribeAlarmsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlarmsResponse) SetStatusCode(v int32) *DescribeAlarmsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlarmsResponse) SetBody(v *DescribeAlarmsResponseBody) *DescribeAlarmsResponse {
	s.Body = v
	return s
}

type DescribeEciScalingConfigurationsRequest struct {
	OwnerAccount              *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber                *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                  *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId                  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount      *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingConfigurationIds   []*string `json:"ScalingConfigurationIds,omitempty" xml:"ScalingConfigurationIds,omitempty" type:"Repeated"`
	ScalingConfigurationNames []*string `json:"ScalingConfigurationNames,omitempty" xml:"ScalingConfigurationNames,omitempty" type:"Repeated"`
	ScalingGroupId            *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeEciScalingConfigurationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciScalingConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsRequest) SetOwnerAccount(v string) *DescribeEciScalingConfigurationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeEciScalingConfigurationsRequest) SetOwnerId(v int64) *DescribeEciScalingConfigurationsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsRequest) SetPageNumber(v int32) *DescribeEciScalingConfigurationsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEciScalingConfigurationsRequest) SetPageSize(v int32) *DescribeEciScalingConfigurationsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEciScalingConfigurationsRequest) SetRegionId(v string) *DescribeEciScalingConfigurationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsRequest) SetResourceOwnerAccount(v string) *DescribeEciScalingConfigurationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEciScalingConfigurationsRequest) SetResourceOwnerId(v int64) *DescribeEciScalingConfigurationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsRequest) SetScalingConfigurationIds(v []*string) *DescribeEciScalingConfigurationsRequest {
	s.ScalingConfigurationIds = v
	return s
}

func (s *DescribeEciScalingConfigurationsRequest) SetScalingConfigurationNames(v []*string) *DescribeEciScalingConfigurationsRequest {
	s.ScalingConfigurationNames = v
	return s
}

func (s *DescribeEciScalingConfigurationsRequest) SetScalingGroupId(v string) *DescribeEciScalingConfigurationsRequest {
	s.ScalingGroupId = &v
	return s
}

type DescribeEciScalingConfigurationsResponseBody struct {
	PageNumber            *int32                                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32                                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId             *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingConfigurations []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurations `json:"ScalingConfigurations,omitempty" xml:"ScalingConfigurations,omitempty" type:"Repeated"`
	TotalCount            *int32                                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBody) SetPageNumber(v int32) *DescribeEciScalingConfigurationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBody) SetPageSize(v int32) *DescribeEciScalingConfigurationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBody) SetRequestId(v string) *DescribeEciScalingConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBody) SetScalingConfigurations(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) *DescribeEciScalingConfigurationsResponseBody {
	s.ScalingConfigurations = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBody) SetTotalCount(v int32) *DescribeEciScalingConfigurationsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurations struct {
	AcrRegistryInfos              []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos         `json:"AcrRegistryInfos,omitempty" xml:"AcrRegistryInfos,omitempty" type:"Repeated"`
	ActiveDeadlineSeconds         *int32                                                                                       `json:"ActiveDeadlineSeconds,omitempty" xml:"ActiveDeadlineSeconds,omitempty"`
	AutoCreateEip                 *bool                                                                                        `json:"AutoCreateEip,omitempty" xml:"AutoCreateEip,omitempty"`
	AutoMatchImageCache           *bool                                                                                        `json:"AutoMatchImageCache,omitempty" xml:"AutoMatchImageCache,omitempty"`
	ContainerGroupName            *string                                                                                      `json:"ContainerGroupName,omitempty" xml:"ContainerGroupName,omitempty"`
	Containers                    []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers               `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	CostOptimization              *bool                                                                                        `json:"CostOptimization,omitempty" xml:"CostOptimization,omitempty"`
	Cpu                           *float32                                                                                     `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CpuOptionsCore                *int32                                                                                       `json:"CpuOptionsCore,omitempty" xml:"CpuOptionsCore,omitempty"`
	CpuOptionsThreadsPerCore      *int32                                                                                       `json:"CpuOptionsThreadsPerCore,omitempty" xml:"CpuOptionsThreadsPerCore,omitempty"`
	CreationTime                  *string                                                                                      `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description                   *string                                                                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	DnsConfigNameServers          []*string                                                                                    `json:"DnsConfigNameServers,omitempty" xml:"DnsConfigNameServers,omitempty" type:"Repeated"`
	DnsConfigOptions              []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions         `json:"DnsConfigOptions,omitempty" xml:"DnsConfigOptions,omitempty" type:"Repeated"`
	DnsConfigSearches             []*string                                                                                    `json:"DnsConfigSearches,omitempty" xml:"DnsConfigSearches,omitempty" type:"Repeated"`
	DnsPolicy                     *string                                                                                      `json:"DnsPolicy,omitempty" xml:"DnsPolicy,omitempty"`
	EgressBandwidth               *int64                                                                                       `json:"EgressBandwidth,omitempty" xml:"EgressBandwidth,omitempty"`
	EipBandwidth                  *int32                                                                                       `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	EphemeralStorage              *int32                                                                                       `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
	HostAliases                   []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases              `json:"HostAliases,omitempty" xml:"HostAliases,omitempty" type:"Repeated"`
	HostName                      *string                                                                                      `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ImageRegistryCredentials      []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials `json:"ImageRegistryCredentials,omitempty" xml:"ImageRegistryCredentials,omitempty" type:"Repeated"`
	ImageSnapshotId               *string                                                                                      `json:"ImageSnapshotId,omitempty" xml:"ImageSnapshotId,omitempty"`
	IngressBandwidth              *int64                                                                                       `json:"IngressBandwidth,omitempty" xml:"IngressBandwidth,omitempty"`
	InitContainers                []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers           `json:"InitContainers,omitempty" xml:"InitContainers,omitempty" type:"Repeated"`
	InstanceFamilyLevel           *string                                                                                      `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	Ipv6AddressCount              *int32                                                                                       `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	LifecycleState                *string                                                                                      `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	LoadBalancerWeight            *int32                                                                                       `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	Memory                        *float32                                                                                     `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NtpServers                    []*string                                                                                    `json:"NtpServers,omitempty" xml:"NtpServers,omitempty" type:"Repeated"`
	RamRoleName                   *string                                                                                      `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	RegionId                      *string                                                                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId               *string                                                                                      `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RestartPolicy                 *string                                                                                      `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	ScalingConfigurationId        *string                                                                                      `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	ScalingConfigurationName      *string                                                                                      `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	ScalingGroupId                *string                                                                                      `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	SecurityContextSysCtls        []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls   `json:"SecurityContextSysCtls,omitempty" xml:"SecurityContextSysCtls,omitempty" type:"Repeated"`
	SecurityGroupId               *string                                                                                      `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SlsEnable                     *bool                                                                                        `json:"SlsEnable,omitempty" xml:"SlsEnable,omitempty"`
	SpotPriceLimit                *float32                                                                                     `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	SpotStrategy                  *string                                                                                      `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	Tags                          []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags                     `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TerminationGracePeriodSeconds *int32                                                                                       `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	Volumes                       []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes                  `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Repeated"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetAcrRegistryInfos(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.AcrRegistryInfos = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetActiveDeadlineSeconds(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.ActiveDeadlineSeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetAutoCreateEip(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.AutoCreateEip = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetAutoMatchImageCache(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.AutoMatchImageCache = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetContainerGroupName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.ContainerGroupName = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetContainers(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.Containers = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetCostOptimization(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.CostOptimization = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetCpu(v float32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.Cpu = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetCpuOptionsCore(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.CpuOptionsCore = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetCpuOptionsThreadsPerCore(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.CpuOptionsThreadsPerCore = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetCreationTime(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.CreationTime = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetDescription(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.Description = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetDnsConfigNameServers(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.DnsConfigNameServers = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetDnsConfigOptions(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.DnsConfigOptions = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetDnsConfigSearches(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.DnsConfigSearches = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetDnsPolicy(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.DnsPolicy = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetEgressBandwidth(v int64) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.EgressBandwidth = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetEipBandwidth(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.EipBandwidth = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetEphemeralStorage(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.EphemeralStorage = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetHostAliases(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.HostAliases = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetHostName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.HostName = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetImageRegistryCredentials(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.ImageRegistryCredentials = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetImageSnapshotId(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.ImageSnapshotId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetIngressBandwidth(v int64) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.IngressBandwidth = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetInitContainers(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.InitContainers = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetInstanceFamilyLevel(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetIpv6AddressCount(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.Ipv6AddressCount = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetLifecycleState(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.LifecycleState = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetLoadBalancerWeight(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.LoadBalancerWeight = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetMemory(v float32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.Memory = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetNtpServers(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.NtpServers = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetRamRoleName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.RamRoleName = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetRegionId(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.RegionId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetResourceGroupId(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetRestartPolicy(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.RestartPolicy = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetScalingConfigurationId(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetScalingConfigurationName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.ScalingConfigurationName = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetScalingGroupId(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetSecurityContextSysCtls(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.SecurityContextSysCtls = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetSecurityGroupId(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetSlsEnable(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.SlsEnable = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetSpotPriceLimit(v float32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.SpotPriceLimit = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetSpotStrategy(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetTags(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.Tags = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetTerminationGracePeriodSeconds(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetVolumes(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.Volumes = v
	return s
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos struct {
	Domains      []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos) SetDomains(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos {
	s.Domains = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos) SetInstanceId(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos {
	s.InstanceId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos) SetInstanceName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos {
	s.InstanceName = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos) SetRegionId(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos {
	s.RegionId = &v
	return s
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers struct {
	Args                                  []*string                                                                                     `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	Commands                              []*string                                                                                     `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	Cpu                                   *float32                                                                                      `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	EnvironmentVars                       []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	Gpu                                   *int32                                                                                        `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Image                                 *string                                                                                       `json:"Image,omitempty" xml:"Image,omitempty"`
	ImagePullPolicy                       *string                                                                                       `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	LivenessProbeExecCommands             []*string                                                                                     `json:"LivenessProbeExecCommands,omitempty" xml:"LivenessProbeExecCommands,omitempty" type:"Repeated"`
	LivenessProbeFailureThreshold         *int32                                                                                        `json:"LivenessProbeFailureThreshold,omitempty" xml:"LivenessProbeFailureThreshold,omitempty"`
	LivenessProbeHttpGetPath              *string                                                                                       `json:"LivenessProbeHttpGetPath,omitempty" xml:"LivenessProbeHttpGetPath,omitempty"`
	LivenessProbeHttpGetPort              *int32                                                                                        `json:"LivenessProbeHttpGetPort,omitempty" xml:"LivenessProbeHttpGetPort,omitempty"`
	LivenessProbeHttpGetScheme            *string                                                                                       `json:"LivenessProbeHttpGetScheme,omitempty" xml:"LivenessProbeHttpGetScheme,omitempty"`
	LivenessProbeInitialDelaySeconds      *int32                                                                                        `json:"LivenessProbeInitialDelaySeconds,omitempty" xml:"LivenessProbeInitialDelaySeconds,omitempty"`
	LivenessProbePeriodSeconds            *int32                                                                                        `json:"LivenessProbePeriodSeconds,omitempty" xml:"LivenessProbePeriodSeconds,omitempty"`
	LivenessProbeSuccessThreshold         *int32                                                                                        `json:"LivenessProbeSuccessThreshold,omitempty" xml:"LivenessProbeSuccessThreshold,omitempty"`
	LivenessProbeTcpSocketPort            *int32                                                                                        `json:"LivenessProbeTcpSocketPort,omitempty" xml:"LivenessProbeTcpSocketPort,omitempty"`
	LivenessProbeTimeoutSeconds           *int32                                                                                        `json:"LivenessProbeTimeoutSeconds,omitempty" xml:"LivenessProbeTimeoutSeconds,omitempty"`
	Memory                                *float32                                                                                      `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Name                                  *string                                                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	Ports                                 []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts           `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	ReadinessProbeExecCommands            []*string                                                                                     `json:"ReadinessProbeExecCommands,omitempty" xml:"ReadinessProbeExecCommands,omitempty" type:"Repeated"`
	ReadinessProbeFailureThreshold        *int32                                                                                        `json:"ReadinessProbeFailureThreshold,omitempty" xml:"ReadinessProbeFailureThreshold,omitempty"`
	ReadinessProbeHttpGetPath             *string                                                                                       `json:"ReadinessProbeHttpGetPath,omitempty" xml:"ReadinessProbeHttpGetPath,omitempty"`
	ReadinessProbeHttpGetPort             *int32                                                                                        `json:"ReadinessProbeHttpGetPort,omitempty" xml:"ReadinessProbeHttpGetPort,omitempty"`
	ReadinessProbeHttpGetScheme           *string                                                                                       `json:"ReadinessProbeHttpGetScheme,omitempty" xml:"ReadinessProbeHttpGetScheme,omitempty"`
	ReadinessProbeInitialDelaySeconds     *int32                                                                                        `json:"ReadinessProbeInitialDelaySeconds,omitempty" xml:"ReadinessProbeInitialDelaySeconds,omitempty"`
	ReadinessProbePeriodSeconds           *int32                                                                                        `json:"ReadinessProbePeriodSeconds,omitempty" xml:"ReadinessProbePeriodSeconds,omitempty"`
	ReadinessProbeSuccessThreshold        *int32                                                                                        `json:"ReadinessProbeSuccessThreshold,omitempty" xml:"ReadinessProbeSuccessThreshold,omitempty"`
	ReadinessProbeTcpSocketPort           *int32                                                                                        `json:"ReadinessProbeTcpSocketPort,omitempty" xml:"ReadinessProbeTcpSocketPort,omitempty"`
	ReadinessProbeTimeoutSeconds          *int32                                                                                        `json:"ReadinessProbeTimeoutSeconds,omitempty" xml:"ReadinessProbeTimeoutSeconds,omitempty"`
	SecurityContextCapabilityAdds         []*string                                                                                     `json:"SecurityContextCapabilityAdds,omitempty" xml:"SecurityContextCapabilityAdds,omitempty" type:"Repeated"`
	SecurityContextReadOnlyRootFilesystem *bool                                                                                         `json:"SecurityContextReadOnlyRootFilesystem,omitempty" xml:"SecurityContextReadOnlyRootFilesystem,omitempty"`
	SecurityContextRunAsUser              *int64                                                                                        `json:"SecurityContextRunAsUser,omitempty" xml:"SecurityContextRunAsUser,omitempty"`
	Stdin                                 *bool                                                                                         `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	StdinOnce                             *bool                                                                                         `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	Tty                                   *bool                                                                                         `json:"Tty,omitempty" xml:"Tty,omitempty"`
	VolumeMounts                          []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts    `json:"VolumeMounts,omitempty" xml:"VolumeMounts,omitempty" type:"Repeated"`
	WorkingDir                            *string                                                                                       `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetArgs(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.Args = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetCommands(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.Commands = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetCpu(v float32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.Cpu = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetEnvironmentVars(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.EnvironmentVars = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetGpu(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.Gpu = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetImage(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.Image = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetImagePullPolicy(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ImagePullPolicy = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLivenessProbeExecCommands(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LivenessProbeExecCommands = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLivenessProbeFailureThreshold(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LivenessProbeFailureThreshold = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLivenessProbeHttpGetPath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LivenessProbeHttpGetPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLivenessProbeHttpGetPort(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LivenessProbeHttpGetPort = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLivenessProbeHttpGetScheme(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LivenessProbeHttpGetScheme = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLivenessProbeInitialDelaySeconds(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LivenessProbeInitialDelaySeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLivenessProbePeriodSeconds(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LivenessProbePeriodSeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLivenessProbeSuccessThreshold(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LivenessProbeSuccessThreshold = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLivenessProbeTcpSocketPort(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LivenessProbeTcpSocketPort = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLivenessProbeTimeoutSeconds(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LivenessProbeTimeoutSeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetMemory(v float32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.Memory = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetPorts(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.Ports = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetReadinessProbeExecCommands(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ReadinessProbeExecCommands = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetReadinessProbeFailureThreshold(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ReadinessProbeFailureThreshold = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetReadinessProbeHttpGetPath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ReadinessProbeHttpGetPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetReadinessProbeHttpGetPort(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ReadinessProbeHttpGetPort = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetReadinessProbeHttpGetScheme(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ReadinessProbeHttpGetScheme = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetReadinessProbeInitialDelaySeconds(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ReadinessProbeInitialDelaySeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetReadinessProbePeriodSeconds(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ReadinessProbePeriodSeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetReadinessProbeSuccessThreshold(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ReadinessProbeSuccessThreshold = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetReadinessProbeTcpSocketPort(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ReadinessProbeTcpSocketPort = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetReadinessProbeTimeoutSeconds(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ReadinessProbeTimeoutSeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetSecurityContextCapabilityAdds(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.SecurityContextCapabilityAdds = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetSecurityContextReadOnlyRootFilesystem(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.SecurityContextReadOnlyRootFilesystem = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetSecurityContextRunAsUser(v int64) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.SecurityContextRunAsUser = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetStdin(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.Stdin = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetStdinOnce(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.StdinOnce = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetTty(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.Tty = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetVolumeMounts(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.VolumeMounts = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetWorkingDir(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.WorkingDir = &v
	return s
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars struct {
	FieldRefFieldPath *string `json:"FieldRefFieldPath,omitempty" xml:"FieldRefFieldPath,omitempty"`
	Key               *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value             *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars) SetFieldRefFieldPath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars {
	s.FieldRefFieldPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars) SetKey(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars {
	s.Key = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars) SetValue(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars {
	s.Value = &v
	return s
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts struct {
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts) SetPort(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts {
	s.Port = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts) SetProtocol(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts {
	s.Protocol = &v
	return s
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts struct {
	MountPath        *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ReadOnly         *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	SubPath          *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) SetMountPath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts {
	s.MountPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) SetMountPropagation(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts {
	s.MountPropagation = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) SetName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) SetReadOnly(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts {
	s.ReadOnly = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) SetSubPath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts {
	s.SubPath = &v
	return s
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions) SetName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions) SetValue(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions {
	s.Value = &v
	return s
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases struct {
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	Ip        *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases) SetHostnames(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases {
	s.Hostnames = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases) SetIp(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases {
	s.Ip = &v
	return s
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials) SetPassword(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials {
	s.Password = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials) SetServer(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials {
	s.Server = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials) SetUserName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials {
	s.UserName = &v
	return s
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers struct {
	Cpu                                   *float32                                                                                                       `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu                                   *int32                                                                                                         `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Image                                 *string                                                                                                        `json:"Image,omitempty" xml:"Image,omitempty"`
	ImagePullPolicy                       *string                                                                                                        `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	InitContainerArgs                     []*string                                                                                                      `json:"InitContainerArgs,omitempty" xml:"InitContainerArgs,omitempty" type:"Repeated"`
	InitContainerCommands                 []*string                                                                                                      `json:"InitContainerCommands,omitempty" xml:"InitContainerCommands,omitempty" type:"Repeated"`
	InitContainerEnvironmentVars          []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars `json:"InitContainerEnvironmentVars,omitempty" xml:"InitContainerEnvironmentVars,omitempty" type:"Repeated"`
	InitContainerPorts                    []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts           `json:"InitContainerPorts,omitempty" xml:"InitContainerPorts,omitempty" type:"Repeated"`
	InitContainerVolumeMounts             []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts    `json:"InitContainerVolumeMounts,omitempty" xml:"InitContainerVolumeMounts,omitempty" type:"Repeated"`
	Memory                                *float32                                                                                                       `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Name                                  *string                                                                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	SecurityContextCapabilityAdds         []*string                                                                                                      `json:"SecurityContextCapabilityAdds,omitempty" xml:"SecurityContextCapabilityAdds,omitempty" type:"Repeated"`
	SecurityContextReadOnlyRootFilesystem *bool                                                                                                          `json:"SecurityContextReadOnlyRootFilesystem,omitempty" xml:"SecurityContextReadOnlyRootFilesystem,omitempty"`
	SecurityContextRunAsUser              *string                                                                                                        `json:"SecurityContextRunAsUser,omitempty" xml:"SecurityContextRunAsUser,omitempty"`
	WorkingDir                            *string                                                                                                        `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetCpu(v float32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.Cpu = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetGpu(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.Gpu = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetImage(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.Image = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetImagePullPolicy(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.ImagePullPolicy = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetInitContainerArgs(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.InitContainerArgs = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetInitContainerCommands(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.InitContainerCommands = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetInitContainerEnvironmentVars(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.InitContainerEnvironmentVars = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetInitContainerPorts(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.InitContainerPorts = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetInitContainerVolumeMounts(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.InitContainerVolumeMounts = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetMemory(v float32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.Memory = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetSecurityContextCapabilityAdds(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.SecurityContextCapabilityAdds = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetSecurityContextReadOnlyRootFilesystem(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.SecurityContextReadOnlyRootFilesystem = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetSecurityContextRunAsUser(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.SecurityContextRunAsUser = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetWorkingDir(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.WorkingDir = &v
	return s
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars struct {
	FieldRefFieldPath *string `json:"FieldRefFieldPath,omitempty" xml:"FieldRefFieldPath,omitempty"`
	Key               *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value             *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars) SetFieldRefFieldPath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars {
	s.FieldRefFieldPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars) SetKey(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars {
	s.Key = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars) SetValue(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars {
	s.Value = &v
	return s
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts struct {
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts) SetPort(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts {
	s.Port = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts) SetProtocol(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts {
	s.Protocol = &v
	return s
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts struct {
	MountPath        *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ReadOnly         *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	SubPath          *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) SetMountPath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts {
	s.MountPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) SetMountPropagation(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts {
	s.MountPropagation = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) SetName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) SetReadOnly(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts {
	s.ReadOnly = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) SetSubPath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts {
	s.SubPath = &v
	return s
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls) SetName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls) SetValue(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls {
	s.Value = &v
	return s
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags) SetKey(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags {
	s.Key = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags) SetValue(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags {
	s.Value = &v
	return s
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes struct {
	ConfigFileVolumeConfigFileToPaths []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths `json:"ConfigFileVolumeConfigFileToPaths,omitempty" xml:"ConfigFileVolumeConfigFileToPaths,omitempty" type:"Repeated"`
	ConfigFileVolumeDefaultMode       *int32                                                                                                       `json:"ConfigFileVolumeDefaultMode,omitempty" xml:"ConfigFileVolumeDefaultMode,omitempty"`
	DiskVolumeDiskId                  *string                                                                                                      `json:"DiskVolumeDiskId,omitempty" xml:"DiskVolumeDiskId,omitempty"`
	DiskVolumeDiskSize                *int32                                                                                                       `json:"DiskVolumeDiskSize,omitempty" xml:"DiskVolumeDiskSize,omitempty"`
	DiskVolumeFsType                  *string                                                                                                      `json:"DiskVolumeFsType,omitempty" xml:"DiskVolumeFsType,omitempty"`
	EmptyDirVolumeMedium              *string                                                                                                      `json:"EmptyDirVolumeMedium,omitempty" xml:"EmptyDirVolumeMedium,omitempty"`
	FlexVolumeDriver                  *string                                                                                                      `json:"FlexVolumeDriver,omitempty" xml:"FlexVolumeDriver,omitempty"`
	FlexVolumeFsType                  *string                                                                                                      `json:"FlexVolumeFsType,omitempty" xml:"FlexVolumeFsType,omitempty"`
	FlexVolumeOptions                 *string                                                                                                      `json:"FlexVolumeOptions,omitempty" xml:"FlexVolumeOptions,omitempty"`
	NFSVolumePath                     *string                                                                                                      `json:"NFSVolumePath,omitempty" xml:"NFSVolumePath,omitempty"`
	NFSVolumeReadOnly                 *bool                                                                                                        `json:"NFSVolumeReadOnly,omitempty" xml:"NFSVolumeReadOnly,omitempty"`
	NFSVolumeServer                   *string                                                                                                      `json:"NFSVolumeServer,omitempty" xml:"NFSVolumeServer,omitempty"`
	Name                              *string                                                                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	Type                              *string                                                                                                      `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetConfigFileVolumeConfigFileToPaths(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.ConfigFileVolumeConfigFileToPaths = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetConfigFileVolumeDefaultMode(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.ConfigFileVolumeDefaultMode = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetDiskVolumeDiskId(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.DiskVolumeDiskId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetDiskVolumeDiskSize(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.DiskVolumeDiskSize = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetDiskVolumeFsType(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.DiskVolumeFsType = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetEmptyDirVolumeMedium(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.EmptyDirVolumeMedium = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetFlexVolumeDriver(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.FlexVolumeDriver = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetFlexVolumeFsType(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.FlexVolumeFsType = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetFlexVolumeOptions(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.FlexVolumeOptions = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetNFSVolumePath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.NFSVolumePath = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetNFSVolumeReadOnly(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.NFSVolumeReadOnly = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetNFSVolumeServer(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.NFSVolumeServer = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetType(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.Type = &v
	return s
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Mode    *int32  `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Path    *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths) SetContent(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths {
	s.Content = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths) SetMode(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths {
	s.Mode = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths) SetPath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths {
	s.Path = &v
	return s
}

type DescribeEciScalingConfigurationsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeEciScalingConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEciScalingConfigurationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponse) SetHeaders(v map[string]*string) *DescribeEciScalingConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponse) SetStatusCode(v int32) *DescribeEciScalingConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponse) SetBody(v *DescribeEciScalingConfigurationsResponseBody) *DescribeEciScalingConfigurationsResponse {
	s.Body = v
	return s
}

type DescribeLifecycleActionsRequest struct {
	LifecycleActionStatus *string `json:"LifecycleActionStatus,omitempty" xml:"LifecycleActionStatus,omitempty"`
	MaxResults            *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken             *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingActivityId     *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s DescribeLifecycleActionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecycleActionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleActionsRequest) SetLifecycleActionStatus(v string) *DescribeLifecycleActionsRequest {
	s.LifecycleActionStatus = &v
	return s
}

func (s *DescribeLifecycleActionsRequest) SetMaxResults(v int32) *DescribeLifecycleActionsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeLifecycleActionsRequest) SetNextToken(v string) *DescribeLifecycleActionsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeLifecycleActionsRequest) SetOwnerId(v int64) *DescribeLifecycleActionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLifecycleActionsRequest) SetRegionId(v string) *DescribeLifecycleActionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLifecycleActionsRequest) SetResourceOwnerAccount(v string) *DescribeLifecycleActionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLifecycleActionsRequest) SetScalingActivityId(v string) *DescribeLifecycleActionsRequest {
	s.ScalingActivityId = &v
	return s
}

type DescribeLifecycleActionsResponseBody struct {
	LifecycleActions []*DescribeLifecycleActionsResponseBodyLifecycleActions `json:"LifecycleActions,omitempty" xml:"LifecycleActions,omitempty" type:"Repeated"`
	MaxResults       *int32                                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string                                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount       *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLifecycleActionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecycleActionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleActionsResponseBody) SetLifecycleActions(v []*DescribeLifecycleActionsResponseBodyLifecycleActions) *DescribeLifecycleActionsResponseBody {
	s.LifecycleActions = v
	return s
}

func (s *DescribeLifecycleActionsResponseBody) SetMaxResults(v int32) *DescribeLifecycleActionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeLifecycleActionsResponseBody) SetNextToken(v string) *DescribeLifecycleActionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeLifecycleActionsResponseBody) SetRequestId(v string) *DescribeLifecycleActionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLifecycleActionsResponseBody) SetTotalCount(v int32) *DescribeLifecycleActionsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeLifecycleActionsResponseBodyLifecycleActions struct {
	InstanceIds           []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	LifecycleActionResult *string   `json:"LifecycleActionResult,omitempty" xml:"LifecycleActionResult,omitempty"`
	LifecycleActionStatus *string   `json:"LifecycleActionStatus,omitempty" xml:"LifecycleActionStatus,omitempty"`
	LifecycleActionToken  *string   `json:"LifecycleActionToken,omitempty" xml:"LifecycleActionToken,omitempty"`
	LifecycleHookId       *string   `json:"LifecycleHookId,omitempty" xml:"LifecycleHookId,omitempty"`
}

func (s DescribeLifecycleActionsResponseBodyLifecycleActions) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecycleActionsResponseBodyLifecycleActions) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActions) SetInstanceIds(v []*string) *DescribeLifecycleActionsResponseBodyLifecycleActions {
	s.InstanceIds = v
	return s
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActions) SetLifecycleActionResult(v string) *DescribeLifecycleActionsResponseBodyLifecycleActions {
	s.LifecycleActionResult = &v
	return s
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActions) SetLifecycleActionStatus(v string) *DescribeLifecycleActionsResponseBodyLifecycleActions {
	s.LifecycleActionStatus = &v
	return s
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActions) SetLifecycleActionToken(v string) *DescribeLifecycleActionsResponseBodyLifecycleActions {
	s.LifecycleActionToken = &v
	return s
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActions) SetLifecycleHookId(v string) *DescribeLifecycleActionsResponseBodyLifecycleActions {
	s.LifecycleHookId = &v
	return s
}

type DescribeLifecycleActionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLifecycleActionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLifecycleActionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecycleActionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleActionsResponse) SetHeaders(v map[string]*string) *DescribeLifecycleActionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLifecycleActionsResponse) SetStatusCode(v int32) *DescribeLifecycleActionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLifecycleActionsResponse) SetBody(v *DescribeLifecycleActionsResponseBody) *DescribeLifecycleActionsResponse {
	s.Body = v
	return s
}

type DescribeLifecycleHooksRequest struct {
	LifecycleHookIds     []*string `json:"LifecycleHookIds,omitempty" xml:"LifecycleHookIds,omitempty" type:"Repeated"`
	LifecycleHookName    *string   `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeLifecycleHooksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecycleHooksRequest) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleHooksRequest) SetLifecycleHookIds(v []*string) *DescribeLifecycleHooksRequest {
	s.LifecycleHookIds = v
	return s
}

func (s *DescribeLifecycleHooksRequest) SetLifecycleHookName(v string) *DescribeLifecycleHooksRequest {
	s.LifecycleHookName = &v
	return s
}

func (s *DescribeLifecycleHooksRequest) SetOwnerAccount(v string) *DescribeLifecycleHooksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLifecycleHooksRequest) SetOwnerId(v int64) *DescribeLifecycleHooksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLifecycleHooksRequest) SetPageNumber(v int32) *DescribeLifecycleHooksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLifecycleHooksRequest) SetPageSize(v int32) *DescribeLifecycleHooksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLifecycleHooksRequest) SetRegionId(v string) *DescribeLifecycleHooksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLifecycleHooksRequest) SetResourceOwnerAccount(v string) *DescribeLifecycleHooksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLifecycleHooksRequest) SetScalingGroupId(v string) *DescribeLifecycleHooksRequest {
	s.ScalingGroupId = &v
	return s
}

type DescribeLifecycleHooksResponseBody struct {
	LifecycleHooks []*DescribeLifecycleHooksResponseBodyLifecycleHooks `json:"LifecycleHooks,omitempty" xml:"LifecycleHooks,omitempty" type:"Repeated"`
	PageNumber     *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLifecycleHooksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecycleHooksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleHooksResponseBody) SetLifecycleHooks(v []*DescribeLifecycleHooksResponseBodyLifecycleHooks) *DescribeLifecycleHooksResponseBody {
	s.LifecycleHooks = v
	return s
}

func (s *DescribeLifecycleHooksResponseBody) SetPageNumber(v int32) *DescribeLifecycleHooksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBody) SetPageSize(v int32) *DescribeLifecycleHooksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBody) SetRequestId(v string) *DescribeLifecycleHooksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBody) SetTotalCount(v int32) *DescribeLifecycleHooksResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeLifecycleHooksResponseBodyLifecycleHooks struct {
	DefaultResult        *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	HeartbeatTimeout     *int32  `json:"HeartbeatTimeout,omitempty" xml:"HeartbeatTimeout,omitempty"`
	LifecycleHookId      *string `json:"LifecycleHookId,omitempty" xml:"LifecycleHookId,omitempty"`
	LifecycleHookName    *string `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	LifecycleTransition  *string `json:"LifecycleTransition,omitempty" xml:"LifecycleTransition,omitempty"`
	NotificationArn      *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	NotificationMetadata *string `json:"NotificationMetadata,omitempty" xml:"NotificationMetadata,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeLifecycleHooksResponseBodyLifecycleHooks) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecycleHooksResponseBodyLifecycleHooks) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) SetDefaultResult(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooks {
	s.DefaultResult = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) SetHeartbeatTimeout(v int32) *DescribeLifecycleHooksResponseBodyLifecycleHooks {
	s.HeartbeatTimeout = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) SetLifecycleHookId(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooks {
	s.LifecycleHookId = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) SetLifecycleHookName(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooks {
	s.LifecycleHookName = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) SetLifecycleTransition(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooks {
	s.LifecycleTransition = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) SetNotificationArn(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooks {
	s.NotificationArn = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) SetNotificationMetadata(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooks {
	s.NotificationMetadata = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) SetScalingGroupId(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooks {
	s.ScalingGroupId = &v
	return s
}

type DescribeLifecycleHooksResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLifecycleHooksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLifecycleHooksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecycleHooksResponse) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleHooksResponse) SetHeaders(v map[string]*string) *DescribeLifecycleHooksResponse {
	s.Headers = v
	return s
}

func (s *DescribeLifecycleHooksResponse) SetStatusCode(v int32) *DescribeLifecycleHooksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLifecycleHooksResponse) SetBody(v *DescribeLifecycleHooksResponseBody) *DescribeLifecycleHooksResponse {
	s.Body = v
	return s
}

type DescribeLimitationRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DescribeLimitationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLimitationRequest) GoString() string {
	return s.String()
}

func (s *DescribeLimitationRequest) SetOwnerId(v int64) *DescribeLimitationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLimitationRequest) SetResourceOwnerAccount(v string) *DescribeLimitationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type DescribeLimitationResponseBody struct {
	MaxNumberOfAlbServerGroup             *int32  `json:"MaxNumberOfAlbServerGroup,omitempty" xml:"MaxNumberOfAlbServerGroup,omitempty"`
	MaxNumberOfDBInstances                *int32  `json:"MaxNumberOfDBInstances,omitempty" xml:"MaxNumberOfDBInstances,omitempty"`
	MaxNumberOfLifecycleHooks             *int32  `json:"MaxNumberOfLifecycleHooks,omitempty" xml:"MaxNumberOfLifecycleHooks,omitempty"`
	MaxNumberOfLoadBalancers              *int32  `json:"MaxNumberOfLoadBalancers,omitempty" xml:"MaxNumberOfLoadBalancers,omitempty"`
	MaxNumberOfMaxSize                    *int32  `json:"MaxNumberOfMaxSize,omitempty" xml:"MaxNumberOfMaxSize,omitempty"`
	MaxNumberOfMinSize                    *int32  `json:"MaxNumberOfMinSize,omitempty" xml:"MaxNumberOfMinSize,omitempty"`
	MaxNumberOfNotificationConfigurations *int32  `json:"MaxNumberOfNotificationConfigurations,omitempty" xml:"MaxNumberOfNotificationConfigurations,omitempty"`
	MaxNumberOfScalingConfigurations      *int32  `json:"MaxNumberOfScalingConfigurations,omitempty" xml:"MaxNumberOfScalingConfigurations,omitempty"`
	MaxNumberOfScalingGroups              *int32  `json:"MaxNumberOfScalingGroups,omitempty" xml:"MaxNumberOfScalingGroups,omitempty"`
	MaxNumberOfScalingInstances           *int32  `json:"MaxNumberOfScalingInstances,omitempty" xml:"MaxNumberOfScalingInstances,omitempty"`
	MaxNumberOfScalingRules               *int32  `json:"MaxNumberOfScalingRules,omitempty" xml:"MaxNumberOfScalingRules,omitempty"`
	MaxNumberOfScheduledTasks             *int32  `json:"MaxNumberOfScheduledTasks,omitempty" xml:"MaxNumberOfScheduledTasks,omitempty"`
	MaxNumberOfVServerGroups              *int32  `json:"MaxNumberOfVServerGroups,omitempty" xml:"MaxNumberOfVServerGroups,omitempty"`
	RequestId                             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLimitationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLimitationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfAlbServerGroup(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfAlbServerGroup = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfDBInstances(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfDBInstances = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfLifecycleHooks(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfLifecycleHooks = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfLoadBalancers(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfLoadBalancers = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfMaxSize(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfMaxSize = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfMinSize(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfMinSize = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfNotificationConfigurations(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfNotificationConfigurations = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfScalingConfigurations(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfScalingConfigurations = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfScalingGroups(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfScalingGroups = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfScalingInstances(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfScalingInstances = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfScalingRules(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfScalingRules = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfScheduledTasks(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfScheduledTasks = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfVServerGroups(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfVServerGroups = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetRequestId(v string) *DescribeLimitationResponseBody {
	s.RequestId = &v
	return s
}

type DescribeLimitationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLimitationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLimitationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLimitationResponse) GoString() string {
	return s.String()
}

func (s *DescribeLimitationResponse) SetHeaders(v map[string]*string) *DescribeLimitationResponse {
	s.Headers = v
	return s
}

func (s *DescribeLimitationResponse) SetStatusCode(v int32) *DescribeLimitationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLimitationResponse) SetBody(v *DescribeLimitationResponseBody) *DescribeLimitationResponse {
	s.Body = v
	return s
}

type DescribeNotificationConfigurationsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeNotificationConfigurationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotificationConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeNotificationConfigurationsRequest) SetOwnerId(v int64) *DescribeNotificationConfigurationsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeNotificationConfigurationsRequest) SetRegionId(v string) *DescribeNotificationConfigurationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNotificationConfigurationsRequest) SetResourceOwnerAccount(v string) *DescribeNotificationConfigurationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeNotificationConfigurationsRequest) SetScalingGroupId(v string) *DescribeNotificationConfigurationsRequest {
	s.ScalingGroupId = &v
	return s
}

type DescribeNotificationConfigurationsResponseBody struct {
	NotificationConfigurationModels []*DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels `json:"NotificationConfigurationModels,omitempty" xml:"NotificationConfigurationModels,omitempty" type:"Repeated"`
	RequestId                       *string                                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNotificationConfigurationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotificationConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNotificationConfigurationsResponseBody) SetNotificationConfigurationModels(v []*DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) *DescribeNotificationConfigurationsResponseBody {
	s.NotificationConfigurationModels = v
	return s
}

func (s *DescribeNotificationConfigurationsResponseBody) SetRequestId(v string) *DescribeNotificationConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels struct {
	NotificationArn   *string   `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	NotificationTypes []*string `json:"NotificationTypes,omitempty" xml:"NotificationTypes,omitempty" type:"Repeated"`
	ScalingGroupId    *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) GoString() string {
	return s.String()
}

func (s *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) SetNotificationArn(v string) *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels {
	s.NotificationArn = &v
	return s
}

func (s *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) SetNotificationTypes(v []*string) *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels {
	s.NotificationTypes = v
	return s
}

func (s *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) SetScalingGroupId(v string) *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels {
	s.ScalingGroupId = &v
	return s
}

type DescribeNotificationConfigurationsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeNotificationConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNotificationConfigurationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotificationConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeNotificationConfigurationsResponse) SetHeaders(v map[string]*string) *DescribeNotificationConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeNotificationConfigurationsResponse) SetStatusCode(v int32) *DescribeNotificationConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNotificationConfigurationsResponse) SetBody(v *DescribeNotificationConfigurationsResponseBody) *DescribeNotificationConfigurationsResponse {
	s.Body = v
	return s
}

type DescribeNotificationTypesRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DescribeNotificationTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotificationTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNotificationTypesRequest) SetOwnerId(v int64) *DescribeNotificationTypesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeNotificationTypesRequest) SetResourceOwnerAccount(v string) *DescribeNotificationTypesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type DescribeNotificationTypesResponseBody struct {
	NotificationTypes []*string `json:"NotificationTypes,omitempty" xml:"NotificationTypes,omitempty" type:"Repeated"`
	RequestId         *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNotificationTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotificationTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNotificationTypesResponseBody) SetNotificationTypes(v []*string) *DescribeNotificationTypesResponseBody {
	s.NotificationTypes = v
	return s
}

func (s *DescribeNotificationTypesResponseBody) SetRequestId(v string) *DescribeNotificationTypesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNotificationTypesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeNotificationTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNotificationTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotificationTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNotificationTypesResponse) SetHeaders(v map[string]*string) *DescribeNotificationTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNotificationTypesResponse) SetStatusCode(v int32) *DescribeNotificationTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNotificationTypesResponse) SetBody(v *DescribeNotificationTypesResponseBody) *DescribeNotificationTypesResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	AcceptLanguage       *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s *DescribeRegionsRequest) SetOwnerId(v int64) *DescribeRegionsRequest {
	s.OwnerId = &v
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
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	ClassicUnavailable *bool   `json:"ClassicUnavailable,omitempty" xml:"ClassicUnavailable,omitempty"`
	LocalName          *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint     *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcUnavailable     *bool   `json:"VpcUnavailable,omitempty" xml:"VpcUnavailable,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetClassicUnavailable(v bool) *DescribeRegionsResponseBodyRegions {
	s.ClassicUnavailable = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetVpcUnavailable(v bool) *DescribeRegionsResponseBodyRegions {
	s.VpcUnavailable = &v
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

type DescribeScalingActivitiesRequest struct {
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingActivityIds   []*string `json:"ScalingActivityIds,omitempty" xml:"ScalingActivityIds,omitempty" type:"Repeated"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	StatusCode           *string   `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s DescribeScalingActivitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingActivitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivitiesRequest) SetOwnerAccount(v string) *DescribeScalingActivitiesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetOwnerId(v int64) *DescribeScalingActivitiesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetPageNumber(v int32) *DescribeScalingActivitiesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetPageSize(v int32) *DescribeScalingActivitiesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetRegionId(v string) *DescribeScalingActivitiesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetResourceOwnerAccount(v string) *DescribeScalingActivitiesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetResourceOwnerId(v int64) *DescribeScalingActivitiesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetScalingActivityIds(v []*string) *DescribeScalingActivitiesRequest {
	s.ScalingActivityIds = v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetScalingGroupId(v string) *DescribeScalingActivitiesRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetStatusCode(v string) *DescribeScalingActivitiesRequest {
	s.StatusCode = &v
	return s
}

type DescribeScalingActivitiesResponseBody struct {
	PageNumber        *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId         *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingActivities []*DescribeScalingActivitiesResponseBodyScalingActivities `json:"ScalingActivities,omitempty" xml:"ScalingActivities,omitempty" type:"Repeated"`
	TotalCount        *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScalingActivitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingActivitiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivitiesResponseBody) SetPageNumber(v int32) *DescribeScalingActivitiesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBody) SetPageSize(v int32) *DescribeScalingActivitiesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBody) SetRequestId(v string) *DescribeScalingActivitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBody) SetScalingActivities(v []*DescribeScalingActivitiesResponseBodyScalingActivities) *DescribeScalingActivitiesResponseBody {
	s.ScalingActivities = v
	return s
}

func (s *DescribeScalingActivitiesResponseBody) SetTotalCount(v int32) *DescribeScalingActivitiesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeScalingActivitiesResponseBodyScalingActivities struct {
	AttachedCapacity      *string `json:"AttachedCapacity,omitempty" xml:"AttachedCapacity,omitempty"`
	AutoCreatedCapacity   *string `json:"AutoCreatedCapacity,omitempty" xml:"AutoCreatedCapacity,omitempty"`
	Cause                 *string `json:"Cause,omitempty" xml:"Cause,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime               *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Progress              *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ScalingActivityId     *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	ScalingGroupId        *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScalingInstanceNumber *int32  `json:"ScalingInstanceNumber,omitempty" xml:"ScalingInstanceNumber,omitempty"`
	StartTime             *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatusCode            *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage         *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	TotalCapacity         *string `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
}

func (s DescribeScalingActivitiesResponseBodyScalingActivities) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingActivitiesResponseBodyScalingActivities) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetAttachedCapacity(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.AttachedCapacity = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetAutoCreatedCapacity(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.AutoCreatedCapacity = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetCause(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.Cause = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetDescription(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.Description = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetEndTime(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.EndTime = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetProgress(v int32) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.Progress = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetScalingActivityId(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.ScalingActivityId = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetScalingGroupId(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetScalingInstanceNumber(v int32) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.ScalingInstanceNumber = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetStartTime(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.StartTime = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetStatusCode(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetStatusMessage(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.StatusMessage = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetTotalCapacity(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.TotalCapacity = &v
	return s
}

type DescribeScalingActivitiesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScalingActivitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScalingActivitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingActivitiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivitiesResponse) SetHeaders(v map[string]*string) *DescribeScalingActivitiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingActivitiesResponse) SetStatusCode(v int32) *DescribeScalingActivitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingActivitiesResponse) SetBody(v *DescribeScalingActivitiesResponseBody) *DescribeScalingActivitiesResponse {
	s.Body = v
	return s
}

type DescribeScalingActivityDetailRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingActivityId    *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s DescribeScalingActivityDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingActivityDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivityDetailRequest) SetOwnerId(v int64) *DescribeScalingActivityDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScalingActivityDetailRequest) SetResourceOwnerAccount(v string) *DescribeScalingActivityDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScalingActivityDetailRequest) SetResourceOwnerId(v int64) *DescribeScalingActivityDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingActivityDetailRequest) SetScalingActivityId(v string) *DescribeScalingActivityDetailRequest {
	s.ScalingActivityId = &v
	return s
}

type DescribeScalingActivityDetailResponseBody struct {
	Detail            *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s DescribeScalingActivityDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingActivityDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivityDetailResponseBody) SetDetail(v string) *DescribeScalingActivityDetailResponseBody {
	s.Detail = &v
	return s
}

func (s *DescribeScalingActivityDetailResponseBody) SetRequestId(v string) *DescribeScalingActivityDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingActivityDetailResponseBody) SetScalingActivityId(v string) *DescribeScalingActivityDetailResponseBody {
	s.ScalingActivityId = &v
	return s
}

type DescribeScalingActivityDetailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScalingActivityDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScalingActivityDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingActivityDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivityDetailResponse) SetHeaders(v map[string]*string) *DescribeScalingActivityDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingActivityDetailResponse) SetStatusCode(v int32) *DescribeScalingActivityDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingActivityDetailResponse) SetBody(v *DescribeScalingActivityDetailResponseBody) *DescribeScalingActivityDetailResponse {
	s.Body = v
	return s
}

type DescribeScalingConfigurationsRequest struct {
	OwnerAccount              *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber                *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                  *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId                  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount      *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingConfigurationIds   []*string `json:"ScalingConfigurationIds,omitempty" xml:"ScalingConfigurationIds,omitempty" type:"Repeated"`
	ScalingConfigurationNames []*string `json:"ScalingConfigurationNames,omitempty" xml:"ScalingConfigurationNames,omitempty" type:"Repeated"`
	ScalingGroupId            *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeScalingConfigurationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsRequest) SetOwnerAccount(v string) *DescribeScalingConfigurationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetOwnerId(v int64) *DescribeScalingConfigurationsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetPageNumber(v int32) *DescribeScalingConfigurationsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetPageSize(v int32) *DescribeScalingConfigurationsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetRegionId(v string) *DescribeScalingConfigurationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetResourceOwnerAccount(v string) *DescribeScalingConfigurationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetResourceOwnerId(v int64) *DescribeScalingConfigurationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetScalingConfigurationIds(v []*string) *DescribeScalingConfigurationsRequest {
	s.ScalingConfigurationIds = v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetScalingConfigurationNames(v []*string) *DescribeScalingConfigurationsRequest {
	s.ScalingConfigurationNames = v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetScalingGroupId(v string) *DescribeScalingConfigurationsRequest {
	s.ScalingGroupId = &v
	return s
}

type DescribeScalingConfigurationsResponseBody struct {
	PageNumber            *int32                                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32                                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId             *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingConfigurations []*DescribeScalingConfigurationsResponseBodyScalingConfigurations `json:"ScalingConfigurations,omitempty" xml:"ScalingConfigurations,omitempty" type:"Repeated"`
	TotalCount            *int32                                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBody) SetPageNumber(v int32) *DescribeScalingConfigurationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBody) SetPageSize(v int32) *DescribeScalingConfigurationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBody) SetRequestId(v string) *DescribeScalingConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBody) SetScalingConfigurations(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurations) *DescribeScalingConfigurationsResponseBody {
	s.ScalingConfigurations = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBody) SetTotalCount(v int32) *DescribeScalingConfigurationsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurations struct {
	PrivatePoolOptions             *DescribeScalingConfigurationsResponseBodyScalingConfigurationsPrivatePoolOptions     `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" require:"true" type:"Struct"`
	Affinity                       *string                                                                               `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	Cpu                            *int32                                                                                `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreationTime                   *string                                                                               `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	CreditSpecification            *string                                                                               `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	DataDisks                      []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks            `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	DedicatedHostId                *string                                                                               `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	DeploymentSetId                *string                                                                               `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	HostName                       *string                                                                               `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HpcClusterId                   *string                                                                               `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	ImageFamily                    *string                                                                               `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	ImageId                        *string                                                                               `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName                      *string                                                                               `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	InstanceDescription            *string                                                                               `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	InstanceGeneration             *string                                                                               `json:"InstanceGeneration,omitempty" xml:"InstanceGeneration,omitempty"`
	InstanceName                   *string                                                                               `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstancePatternInfos           []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos `json:"InstancePatternInfos,omitempty" xml:"InstancePatternInfos,omitempty" type:"Repeated"`
	InstanceType                   *string                                                                               `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceTypes                  []*string                                                                             `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	InternetChargeType             *string                                                                               `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandwidthIn         *int32                                                                                `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	InternetMaxBandwidthOut        *int32                                                                                `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	IoOptimized                    *string                                                                               `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	Ipv6AddressCount               *int32                                                                                `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	KeyPairName                    *string                                                                               `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	LifecycleState                 *string                                                                               `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	LoadBalancerWeight             *int32                                                                                `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	Memory                         *int32                                                                                `json:"Memory,omitempty" xml:"Memory,omitempty"`
	PasswordInherit                *bool                                                                                 `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	RamRoleName                    *string                                                                               `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	ResourceGroupId                *string                                                                               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ScalingConfigurationId         *string                                                                               `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	ScalingConfigurationName       *string                                                                               `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	ScalingGroupId                 *string                                                                               `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	SchedulerOptions               *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSchedulerOptions       `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty" type:"Struct"`
	SecurityEnhancementStrategy    *string                                                                               `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	SecurityGroupId                *string                                                                               `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupIds               []*string                                                                             `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	SpotDuration                   *int32                                                                                `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	SpotInterruptionBehavior       *string                                                                               `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	SpotPriceLimits                []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits      `json:"SpotPriceLimits,omitempty" xml:"SpotPriceLimits,omitempty" type:"Repeated"`
	SpotStrategy                   *string                                                                               `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDiskAutoSnapshotPolicyId *string                                                                               `json:"SystemDiskAutoSnapshotPolicyId,omitempty" xml:"SystemDiskAutoSnapshotPolicyId,omitempty"`
	SystemDiskBurstingEnabled      *bool                                                                                 `json:"SystemDiskBurstingEnabled,omitempty" xml:"SystemDiskBurstingEnabled,omitempty"`
	SystemDiskCategories           []*string                                                                             `json:"SystemDiskCategories,omitempty" xml:"SystemDiskCategories,omitempty" type:"Repeated"`
	SystemDiskCategory             *string                                                                               `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	SystemDiskDescription          *string                                                                               `json:"SystemDiskDescription,omitempty" xml:"SystemDiskDescription,omitempty"`
	SystemDiskEncryptAlgorithm     *string                                                                               `json:"SystemDiskEncryptAlgorithm,omitempty" xml:"SystemDiskEncryptAlgorithm,omitempty"`
	SystemDiskEncrypted            *bool                                                                                 `json:"SystemDiskEncrypted,omitempty" xml:"SystemDiskEncrypted,omitempty"`
	SystemDiskKMSKeyId             *string                                                                               `json:"SystemDiskKMSKeyId,omitempty" xml:"SystemDiskKMSKeyId,omitempty"`
	SystemDiskName                 *string                                                                               `json:"SystemDiskName,omitempty" xml:"SystemDiskName,omitempty"`
	SystemDiskPerformanceLevel     *string                                                                               `json:"SystemDiskPerformanceLevel,omitempty" xml:"SystemDiskPerformanceLevel,omitempty"`
	SystemDiskProvisionedIops      *int64                                                                                `json:"SystemDiskProvisionedIops,omitempty" xml:"SystemDiskProvisionedIops,omitempty"`
	SystemDiskSize                 *int32                                                                                `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	Tags                           []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags                 `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Tenancy                        *string                                                                               `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	UserData                       *string                                                                               `json:"UserData,omitempty" xml:"UserData,omitempty"`
	WeightedCapacities             []*int32                                                                              `json:"WeightedCapacities,omitempty" xml:"WeightedCapacities,omitempty" type:"Repeated"`
	ZoneId                         *string                                                                               `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurations) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurations) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetPrivatePoolOptions(v *DescribeScalingConfigurationsResponseBodyScalingConfigurationsPrivatePoolOptions) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.PrivatePoolOptions = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetAffinity(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.Affinity = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetCpu(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.Cpu = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetCreationTime(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.CreationTime = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetCreditSpecification(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.CreditSpecification = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetDataDisks(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.DataDisks = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetDedicatedHostId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetDeploymentSetId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.DeploymentSetId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetHostName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.HostName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetHpcClusterId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.HpcClusterId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetImageFamily(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ImageFamily = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetImageId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ImageId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetImageName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ImageName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetInstanceDescription(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetInstanceGeneration(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.InstanceGeneration = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetInstanceName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.InstanceName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetInstancePatternInfos(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.InstancePatternInfos = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetInstanceType(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.InstanceType = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetInstanceTypes(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.InstanceTypes = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetInternetChargeType(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetInternetMaxBandwidthIn(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetInternetMaxBandwidthOut(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetIoOptimized(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.IoOptimized = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetIpv6AddressCount(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.Ipv6AddressCount = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetKeyPairName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.KeyPairName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetLifecycleState(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.LifecycleState = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetLoadBalancerWeight(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.LoadBalancerWeight = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetMemory(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.Memory = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetPasswordInherit(v bool) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.PasswordInherit = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetRamRoleName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.RamRoleName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetResourceGroupId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetScalingConfigurationId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetScalingConfigurationName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ScalingConfigurationName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetScalingGroupId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSchedulerOptions(v *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSchedulerOptions) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SchedulerOptions = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSecurityEnhancementStrategy(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSecurityGroupId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSecurityGroupIds(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSpotDuration(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SpotDuration = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSpotInterruptionBehavior(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSpotPriceLimits(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SpotPriceLimits = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSpotStrategy(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskAutoSnapshotPolicyId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskAutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskBurstingEnabled(v bool) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskBurstingEnabled = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskCategories(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskCategories = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskCategory(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskDescription(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskDescription = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskEncryptAlgorithm(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskEncryptAlgorithm = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskEncrypted(v bool) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskEncrypted = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskKMSKeyId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskKMSKeyId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskPerformanceLevel(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskProvisionedIops(v int64) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskProvisionedIops = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskSize(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetTags(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.Tags = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetTenancy(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.Tenancy = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetUserData(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.UserData = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetWeightedCapacities(v []*int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.WeightedCapacities = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetZoneId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ZoneId = &v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsPrivatePoolOptions struct {
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsPrivatePoolOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsPrivatePoolOptions) SetId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsPrivatePoolOptions) SetMatchCriteria(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks struct {
	AutoSnapshotPolicyId *string   `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	BurstingEnabled      *bool     `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	Categories           []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	Category             *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteWithInstance   *bool     `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	Description          *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Device               *string   `json:"Device,omitempty" xml:"Device,omitempty"`
	DiskName             *string   `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	Encrypted            *string   `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	KMSKeyId             *string   `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	PerformanceLevel     *string   `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	ProvisionedIops      *int64    `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	Size                 *int32    `json:"Size,omitempty" xml:"Size,omitempty"`
	SnapshotId           *string   `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetAutoSnapshotPolicyId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetBurstingEnabled(v bool) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.BurstingEnabled = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetCategories(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.Categories = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetCategory(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.Category = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetDeleteWithInstance(v bool) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetDescription(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.Description = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetDevice(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.Device = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetDiskName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.DiskName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetEncrypted(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.Encrypted = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetKMSKeyId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.KMSKeyId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetPerformanceLevel(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetProvisionedIops(v int64) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.ProvisionedIops = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetSize(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.Size = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetSnapshotId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.SnapshotId = &v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos struct {
	Architectures         []*string `json:"Architectures,omitempty" xml:"Architectures,omitempty" type:"Repeated"`
	BurstablePerformance  *string   `json:"BurstablePerformance,omitempty" xml:"BurstablePerformance,omitempty"`
	Cores                 *int32    `json:"Cores,omitempty" xml:"Cores,omitempty"`
	ExcludedInstanceTypes []*string `json:"ExcludedInstanceTypes,omitempty" xml:"ExcludedInstanceTypes,omitempty" type:"Repeated"`
	InstanceFamilyLevel   *string   `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	MaxPrice              *float32  `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	Memory                *float32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetArchitectures(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.Architectures = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetBurstablePerformance(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.BurstablePerformance = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetCores(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.Cores = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetExcludedInstanceTypes(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.ExcludedInstanceTypes = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetInstanceFamilyLevel(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetMaxPrice(v float32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.MaxPrice = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetMemory(v float32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.Memory = &v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsSchedulerOptions struct {
	ManagedPrivateSpaceId *string `json:"ManagedPrivateSpaceId,omitempty" xml:"ManagedPrivateSpaceId,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsSchedulerOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsSchedulerOptions) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSchedulerOptions) SetManagedPrivateSpaceId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSchedulerOptions {
	s.ManagedPrivateSpaceId = &v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits struct {
	InstanceType *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	PriceLimit   *float32 `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits) SetInstanceType(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits {
	s.InstanceType = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits) SetPriceLimit(v float32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits {
	s.PriceLimit = &v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags) SetKey(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags {
	s.Key = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags) SetValue(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags {
	s.Value = &v
	return s
}

type DescribeScalingConfigurationsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScalingConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScalingConfigurationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponse) SetHeaders(v map[string]*string) *DescribeScalingConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingConfigurationsResponse) SetStatusCode(v int32) *DescribeScalingConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingConfigurationsResponse) SetBody(v *DescribeScalingConfigurationsResponseBody) *DescribeScalingConfigurationsResponse {
	s.Body = v
	return s
}

type DescribeScalingGroupsRequest struct {
	GroupType            *string   `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupIds      []*string `json:"ScalingGroupIds,omitempty" xml:"ScalingGroupIds,omitempty" type:"Repeated"`
	ScalingGroupName     *string   `json:"ScalingGroupName,omitempty" xml:"ScalingGroupName,omitempty"`
	ScalingGroupNames    []*string `json:"ScalingGroupNames,omitempty" xml:"ScalingGroupNames,omitempty" type:"Repeated"`
}

func (s DescribeScalingGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsRequest) SetGroupType(v string) *DescribeScalingGroupsRequest {
	s.GroupType = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetOwnerAccount(v string) *DescribeScalingGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetOwnerId(v int64) *DescribeScalingGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetPageNumber(v int32) *DescribeScalingGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetPageSize(v int32) *DescribeScalingGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetRegionId(v string) *DescribeScalingGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetResourceOwnerAccount(v string) *DescribeScalingGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetResourceOwnerId(v int64) *DescribeScalingGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetScalingGroupIds(v []*string) *DescribeScalingGroupsRequest {
	s.ScalingGroupIds = v
	return s
}

func (s *DescribeScalingGroupsRequest) SetScalingGroupName(v string) *DescribeScalingGroupsRequest {
	s.ScalingGroupName = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetScalingGroupNames(v []*string) *DescribeScalingGroupsRequest {
	s.ScalingGroupNames = v
	return s
}

type DescribeScalingGroupsResponseBody struct {
	PageNumber    *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingGroups []*DescribeScalingGroupsResponseBodyScalingGroups `json:"ScalingGroups,omitempty" xml:"ScalingGroups,omitempty" type:"Repeated"`
	TotalCount    *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScalingGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBody) SetPageNumber(v int32) *DescribeScalingGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingGroupsResponseBody) SetPageSize(v int32) *DescribeScalingGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingGroupsResponseBody) SetRequestId(v string) *DescribeScalingGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBody) SetScalingGroups(v []*DescribeScalingGroupsResponseBodyScalingGroups) *DescribeScalingGroupsResponseBody {
	s.ScalingGroups = v
	return s
}

func (s *DescribeScalingGroupsResponseBody) SetTotalCount(v int32) *DescribeScalingGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeScalingGroupsResponseBodyScalingGroups struct {
	ActiveCapacity                      *int32                                                                   `json:"ActiveCapacity,omitempty" xml:"ActiveCapacity,omitempty"`
	ActiveScalingConfigurationId        *string                                                                  `json:"ActiveScalingConfigurationId,omitempty" xml:"ActiveScalingConfigurationId,omitempty"`
	AlbServerGroups                     []*DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups         `json:"AlbServerGroups,omitempty" xml:"AlbServerGroups,omitempty" type:"Repeated"`
	AllocationStrategy                  *string                                                                  `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	AzBalance                           *bool                                                                    `json:"AzBalance,omitempty" xml:"AzBalance,omitempty"`
	CompensateWithOnDemand              *bool                                                                    `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	CreationTime                        *string                                                                  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	CurrentHostName                     *string                                                                  `json:"CurrentHostName,omitempty" xml:"CurrentHostName,omitempty"`
	DBInstanceIds                       []*string                                                                `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty" type:"Repeated"`
	DefaultCooldown                     *int32                                                                   `json:"DefaultCooldown,omitempty" xml:"DefaultCooldown,omitempty"`
	DesiredCapacity                     *int32                                                                   `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	GroupDeletionProtection             *bool                                                                    `json:"GroupDeletionProtection,omitempty" xml:"GroupDeletionProtection,omitempty"`
	GroupType                           *string                                                                  `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	HealthCheckType                     *string                                                                  `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	IsElasticStrengthInAlarm            *bool                                                                    `json:"IsElasticStrengthInAlarm,omitempty" xml:"IsElasticStrengthInAlarm,omitempty"`
	LaunchTemplateId                    *string                                                                  `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	LaunchTemplateOverrides             []*DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides `json:"LaunchTemplateOverrides,omitempty" xml:"LaunchTemplateOverrides,omitempty" type:"Repeated"`
	LaunchTemplateVersion               *string                                                                  `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	LifecycleState                      *string                                                                  `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	LoadBalancerIds                     []*string                                                                `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	MaxInstanceLifetime                 *int32                                                                   `json:"MaxInstanceLifetime,omitempty" xml:"MaxInstanceLifetime,omitempty"`
	MaxSize                             *int32                                                                   `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	MinSize                             *int32                                                                   `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	ModificationTime                    *string                                                                  `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	MonitorGroupId                      *string                                                                  `json:"MonitorGroupId,omitempty" xml:"MonitorGroupId,omitempty"`
	MultiAZPolicy                       *string                                                                  `json:"MultiAZPolicy,omitempty" xml:"MultiAZPolicy,omitempty"`
	OnDemandBaseCapacity                *int32                                                                   `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	OnDemandPercentageAboveBaseCapacity *int32                                                                   `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	PendingCapacity                     *int32                                                                   `json:"PendingCapacity,omitempty" xml:"PendingCapacity,omitempty"`
	PendingWaitCapacity                 *int32                                                                   `json:"PendingWaitCapacity,omitempty" xml:"PendingWaitCapacity,omitempty"`
	ProtectedCapacity                   *int32                                                                   `json:"ProtectedCapacity,omitempty" xml:"ProtectedCapacity,omitempty"`
	RegionId                            *string                                                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RemovalPolicies                     []*string                                                                `json:"RemovalPolicies,omitempty" xml:"RemovalPolicies,omitempty" type:"Repeated"`
	RemovingCapacity                    *int32                                                                   `json:"RemovingCapacity,omitempty" xml:"RemovingCapacity,omitempty"`
	RemovingWaitCapacity                *int32                                                                   `json:"RemovingWaitCapacity,omitempty" xml:"RemovingWaitCapacity,omitempty"`
	ScalingGroupId                      *string                                                                  `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScalingGroupName                    *string                                                                  `json:"ScalingGroupName,omitempty" xml:"ScalingGroupName,omitempty"`
	ScalingPolicy                       *string                                                                  `json:"ScalingPolicy,omitempty" xml:"ScalingPolicy,omitempty"`
	SpotAllocationStrategy              *string                                                                  `json:"SpotAllocationStrategy,omitempty" xml:"SpotAllocationStrategy,omitempty"`
	SpotInstancePools                   *int32                                                                   `json:"SpotInstancePools,omitempty" xml:"SpotInstancePools,omitempty"`
	SpotInstanceRemedy                  *bool                                                                    `json:"SpotInstanceRemedy,omitempty" xml:"SpotInstanceRemedy,omitempty"`
	StandbyCapacity                     *int32                                                                   `json:"StandbyCapacity,omitempty" xml:"StandbyCapacity,omitempty"`
	StoppedCapacity                     *int32                                                                   `json:"StoppedCapacity,omitempty" xml:"StoppedCapacity,omitempty"`
	SuspendedProcesses                  []*string                                                                `json:"SuspendedProcesses,omitempty" xml:"SuspendedProcesses,omitempty" type:"Repeated"`
	SystemSuspended                     *bool                                                                    `json:"SystemSuspended,omitempty" xml:"SystemSuspended,omitempty"`
	TotalCapacity                       *int32                                                                   `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
	TotalInstanceCount                  *int32                                                                   `json:"TotalInstanceCount,omitempty" xml:"TotalInstanceCount,omitempty"`
	VServerGroups                       []*DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups           `json:"VServerGroups,omitempty" xml:"VServerGroups,omitempty" type:"Repeated"`
	VSwitchId                           *string                                                                  `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VSwitchIds                          []*string                                                                `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	VpcId                               *string                                                                  `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroups) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetActiveCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.ActiveCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetActiveScalingConfigurationId(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.ActiveScalingConfigurationId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetAlbServerGroups(v []*DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.AlbServerGroups = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetAllocationStrategy(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.AllocationStrategy = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetAzBalance(v bool) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.AzBalance = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetCompensateWithOnDemand(v bool) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.CompensateWithOnDemand = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetCreationTime(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.CreationTime = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetCurrentHostName(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.CurrentHostName = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetDBInstanceIds(v []*string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.DBInstanceIds = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetDefaultCooldown(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.DefaultCooldown = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetDesiredCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.DesiredCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetGroupDeletionProtection(v bool) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.GroupDeletionProtection = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetGroupType(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.GroupType = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetHealthCheckType(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.HealthCheckType = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetIsElasticStrengthInAlarm(v bool) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.IsElasticStrengthInAlarm = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetLaunchTemplateId(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.LaunchTemplateId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetLaunchTemplateOverrides(v []*DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.LaunchTemplateOverrides = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetLaunchTemplateVersion(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetLifecycleState(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.LifecycleState = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetLoadBalancerIds(v []*string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.LoadBalancerIds = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetMaxInstanceLifetime(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.MaxInstanceLifetime = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetMaxSize(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.MaxSize = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetMinSize(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.MinSize = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetModificationTime(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.ModificationTime = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetMonitorGroupId(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.MonitorGroupId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetMultiAZPolicy(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.MultiAZPolicy = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetOnDemandBaseCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.OnDemandBaseCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetOnDemandPercentageAboveBaseCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.OnDemandPercentageAboveBaseCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetPendingCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.PendingCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetPendingWaitCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.PendingWaitCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetProtectedCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.ProtectedCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetRegionId(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetRemovalPolicies(v []*string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.RemovalPolicies = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetRemovingCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.RemovingCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetRemovingWaitCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.RemovingWaitCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetScalingGroupId(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetScalingGroupName(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.ScalingGroupName = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetScalingPolicy(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.ScalingPolicy = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetSpotAllocationStrategy(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.SpotAllocationStrategy = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetSpotInstancePools(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.SpotInstancePools = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetSpotInstanceRemedy(v bool) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetStandbyCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.StandbyCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetStoppedCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.StoppedCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetSuspendedProcesses(v []*string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.SuspendedProcesses = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetSystemSuspended(v bool) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.SystemSuspended = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetTotalCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.TotalCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetTotalInstanceCount(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.TotalInstanceCount = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetVServerGroups(v []*DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.VServerGroups = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetVSwitchId(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.VSwitchId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetVSwitchIds(v []*string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.VSwitchIds = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetVpcId(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.VpcId = &v
	return s
}

type DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups struct {
	AlbServerGroupId *string `json:"AlbServerGroupId,omitempty" xml:"AlbServerGroupId,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Weight           *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups) SetAlbServerGroupId(v string) *DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups {
	s.AlbServerGroupId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups) SetPort(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups {
	s.Port = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups) SetWeight(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups {
	s.Weight = &v
	return s
}

type DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides struct {
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	WeightedCapacity *int32  `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides) SetInstanceType(v string) *DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides {
	s.InstanceType = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides) SetWeightedCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides {
	s.WeightedCapacity = &v
	return s
}

type DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups struct {
	LoadBalancerId         *string                                                                              `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	VServerGroupAttributes []*DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes `json:"VServerGroupAttributes,omitempty" xml:"VServerGroupAttributes,omitempty" type:"Repeated"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups) SetLoadBalancerId(v string) *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups) SetVServerGroupAttributes(v []*DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes) *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups {
	s.VServerGroupAttributes = v
	return s
}

type DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes struct {
	Port           *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	Weight         *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes) SetPort(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes {
	s.Port = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes) SetVServerGroupId(v string) *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes) SetWeight(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes {
	s.Weight = &v
	return s
}

type DescribeScalingGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScalingGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScalingGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponse) SetHeaders(v map[string]*string) *DescribeScalingGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingGroupsResponse) SetStatusCode(v int32) *DescribeScalingGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingGroupsResponse) SetBody(v *DescribeScalingGroupsResponseBody) *DescribeScalingGroupsResponse {
	s.Body = v
	return s
}

type DescribeScalingInstancesRequest struct {
	CreationType           *string   `json:"CreationType,omitempty" xml:"CreationType,omitempty"`
	HealthStatus           *string   `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	InstanceIds            []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	LifecycleState         *string   `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	OwnerAccount           *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber             *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize               *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId               *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount   *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingActivityId      *string   `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	ScalingConfigurationId *string   `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	ScalingGroupId         *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeScalingInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingInstancesRequest) SetCreationType(v string) *DescribeScalingInstancesRequest {
	s.CreationType = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetHealthStatus(v string) *DescribeScalingInstancesRequest {
	s.HealthStatus = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetInstanceIds(v []*string) *DescribeScalingInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeScalingInstancesRequest) SetLifecycleState(v string) *DescribeScalingInstancesRequest {
	s.LifecycleState = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetOwnerAccount(v string) *DescribeScalingInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetOwnerId(v int64) *DescribeScalingInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetPageNumber(v int32) *DescribeScalingInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetPageSize(v int32) *DescribeScalingInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetRegionId(v string) *DescribeScalingInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetResourceOwnerAccount(v string) *DescribeScalingInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetResourceOwnerId(v int64) *DescribeScalingInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetScalingActivityId(v string) *DescribeScalingInstancesRequest {
	s.ScalingActivityId = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetScalingConfigurationId(v string) *DescribeScalingInstancesRequest {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetScalingGroupId(v string) *DescribeScalingInstancesRequest {
	s.ScalingGroupId = &v
	return s
}

type DescribeScalingInstancesResponseBody struct {
	PageNumber       *int32                                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingInstances []*DescribeScalingInstancesResponseBodyScalingInstances `json:"ScalingInstances,omitempty" xml:"ScalingInstances,omitempty" type:"Repeated"`
	TotalCount       *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalSpotCount   *int32                                                  `json:"TotalSpotCount,omitempty" xml:"TotalSpotCount,omitempty"`
}

func (s DescribeScalingInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingInstancesResponseBody) SetPageNumber(v int32) *DescribeScalingInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingInstancesResponseBody) SetPageSize(v int32) *DescribeScalingInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingInstancesResponseBody) SetRequestId(v string) *DescribeScalingInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBody) SetScalingInstances(v []*DescribeScalingInstancesResponseBodyScalingInstances) *DescribeScalingInstancesResponseBody {
	s.ScalingInstances = v
	return s
}

func (s *DescribeScalingInstancesResponseBody) SetTotalCount(v int32) *DescribeScalingInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeScalingInstancesResponseBody) SetTotalSpotCount(v int32) *DescribeScalingInstancesResponseBody {
	s.TotalSpotCount = &v
	return s
}

type DescribeScalingInstancesResponseBodyScalingInstances struct {
	CreatedTime            *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	CreationTime           *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	CreationType           *string `json:"CreationType,omitempty" xml:"CreationType,omitempty"`
	Entrusted              *bool   `json:"Entrusted,omitempty" xml:"Entrusted,omitempty"`
	HealthStatus           *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LaunchTemplateId       *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	LaunchTemplateVersion  *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	LifecycleState         *string `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	LoadBalancerWeight     *int32  `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	ScalingActivityId      *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	ScalingGroupId         *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	SpotStrategy           *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	WarmupState            *string `json:"WarmupState,omitempty" xml:"WarmupState,omitempty"`
	WeightedCapacity       *int32  `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
	ZoneId                 *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeScalingInstancesResponseBodyScalingInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingInstancesResponseBodyScalingInstances) GoString() string {
	return s.String()
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetCreatedTime(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.CreatedTime = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetCreationTime(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.CreationTime = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetCreationType(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.CreationType = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetEntrusted(v bool) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.Entrusted = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetHealthStatus(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.HealthStatus = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetInstanceId(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetLaunchTemplateId(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.LaunchTemplateId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetLaunchTemplateVersion(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetLifecycleState(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.LifecycleState = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetLoadBalancerWeight(v int32) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.LoadBalancerWeight = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetScalingActivityId(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.ScalingActivityId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetScalingConfigurationId(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetScalingGroupId(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetSpotStrategy(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetWarmupState(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.WarmupState = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetWeightedCapacity(v int32) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.WeightedCapacity = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetZoneId(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.ZoneId = &v
	return s
}

type DescribeScalingInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScalingInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScalingInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingInstancesResponse) SetHeaders(v map[string]*string) *DescribeScalingInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingInstancesResponse) SetStatusCode(v int32) *DescribeScalingInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingInstancesResponse) SetBody(v *DescribeScalingInstancesResponseBody) *DescribeScalingInstancesResponse {
	s.Body = v
	return s
}

type DescribeScalingRulesRequest struct {
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScalingRuleAris      []*string `json:"ScalingRuleAris,omitempty" xml:"ScalingRuleAris,omitempty" type:"Repeated"`
	ScalingRuleIds       []*string `json:"ScalingRuleIds,omitempty" xml:"ScalingRuleIds,omitempty" type:"Repeated"`
	ScalingRuleNames     []*string `json:"ScalingRuleNames,omitempty" xml:"ScalingRuleNames,omitempty" type:"Repeated"`
	ScalingRuleType      *string   `json:"ScalingRuleType,omitempty" xml:"ScalingRuleType,omitempty"`
	ShowAlarmRules       *bool     `json:"ShowAlarmRules,omitempty" xml:"ShowAlarmRules,omitempty"`
}

func (s DescribeScalingRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesRequest) SetOwnerAccount(v string) *DescribeScalingRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetOwnerId(v int64) *DescribeScalingRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetPageNumber(v int32) *DescribeScalingRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetPageSize(v int32) *DescribeScalingRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetRegionId(v string) *DescribeScalingRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetResourceOwnerAccount(v string) *DescribeScalingRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetResourceOwnerId(v int64) *DescribeScalingRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetScalingGroupId(v string) *DescribeScalingRulesRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetScalingRuleAris(v []*string) *DescribeScalingRulesRequest {
	s.ScalingRuleAris = v
	return s
}

func (s *DescribeScalingRulesRequest) SetScalingRuleIds(v []*string) *DescribeScalingRulesRequest {
	s.ScalingRuleIds = v
	return s
}

func (s *DescribeScalingRulesRequest) SetScalingRuleNames(v []*string) *DescribeScalingRulesRequest {
	s.ScalingRuleNames = v
	return s
}

func (s *DescribeScalingRulesRequest) SetScalingRuleType(v string) *DescribeScalingRulesRequest {
	s.ScalingRuleType = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetShowAlarmRules(v bool) *DescribeScalingRulesRequest {
	s.ShowAlarmRules = &v
	return s
}

type DescribeScalingRulesResponseBody struct {
	PageNumber   *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingRules []*DescribeScalingRulesResponseBodyScalingRules `json:"ScalingRules,omitempty" xml:"ScalingRules,omitempty" type:"Repeated"`
	TotalCount   *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScalingRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBody) SetPageNumber(v int32) *DescribeScalingRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingRulesResponseBody) SetPageSize(v int32) *DescribeScalingRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingRulesResponseBody) SetRequestId(v string) *DescribeScalingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingRulesResponseBody) SetScalingRules(v []*DescribeScalingRulesResponseBodyScalingRules) *DescribeScalingRulesResponseBody {
	s.ScalingRules = v
	return s
}

func (s *DescribeScalingRulesResponseBody) SetTotalCount(v int32) *DescribeScalingRulesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeScalingRulesResponseBodyScalingRules struct {
	AdjustmentType           *string                                                        `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	AdjustmentValue          *int32                                                         `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	Alarms                   []*DescribeScalingRulesResponseBodyScalingRulesAlarms          `json:"Alarms,omitempty" xml:"Alarms,omitempty" type:"Repeated"`
	Cooldown                 *int32                                                         `json:"Cooldown,omitempty" xml:"Cooldown,omitempty"`
	DisableScaleIn           *bool                                                          `json:"DisableScaleIn,omitempty" xml:"DisableScaleIn,omitempty"`
	EstimatedInstanceWarmup  *int32                                                         `json:"EstimatedInstanceWarmup,omitempty" xml:"EstimatedInstanceWarmup,omitempty"`
	InitialMaxSize           *int32                                                         `json:"InitialMaxSize,omitempty" xml:"InitialMaxSize,omitempty"`
	MaxSize                  *int32                                                         `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	MetricName               *string                                                        `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	MinAdjustmentMagnitude   *int32                                                         `json:"MinAdjustmentMagnitude,omitempty" xml:"MinAdjustmentMagnitude,omitempty"`
	MinSize                  *int32                                                         `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	PredictiveScalingMode    *string                                                        `json:"PredictiveScalingMode,omitempty" xml:"PredictiveScalingMode,omitempty"`
	PredictiveTaskBufferTime *int32                                                         `json:"PredictiveTaskBufferTime,omitempty" xml:"PredictiveTaskBufferTime,omitempty"`
	PredictiveValueBehavior  *string                                                        `json:"PredictiveValueBehavior,omitempty" xml:"PredictiveValueBehavior,omitempty"`
	PredictiveValueBuffer    *int32                                                         `json:"PredictiveValueBuffer,omitempty" xml:"PredictiveValueBuffer,omitempty"`
	ScaleInEvaluationCount   *int32                                                         `json:"ScaleInEvaluationCount,omitempty" xml:"ScaleInEvaluationCount,omitempty"`
	ScaleOutEvaluationCount  *int32                                                         `json:"ScaleOutEvaluationCount,omitempty" xml:"ScaleOutEvaluationCount,omitempty"`
	ScalingGroupId           *string                                                        `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScalingRuleAri           *string                                                        `json:"ScalingRuleAri,omitempty" xml:"ScalingRuleAri,omitempty"`
	ScalingRuleId            *string                                                        `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty"`
	ScalingRuleName          *string                                                        `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	ScalingRuleType          *string                                                        `json:"ScalingRuleType,omitempty" xml:"ScalingRuleType,omitempty"`
	StepAdjustments          []*DescribeScalingRulesResponseBodyScalingRulesStepAdjustments `json:"StepAdjustments,omitempty" xml:"StepAdjustments,omitempty" type:"Repeated"`
	TargetValue              *float32                                                       `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s DescribeScalingRulesResponseBodyScalingRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRules) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetAdjustmentType(v string) *DescribeScalingRulesResponseBodyScalingRules {
	s.AdjustmentType = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetAdjustmentValue(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.AdjustmentValue = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetAlarms(v []*DescribeScalingRulesResponseBodyScalingRulesAlarms) *DescribeScalingRulesResponseBodyScalingRules {
	s.Alarms = v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetCooldown(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.Cooldown = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetDisableScaleIn(v bool) *DescribeScalingRulesResponseBodyScalingRules {
	s.DisableScaleIn = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetEstimatedInstanceWarmup(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.EstimatedInstanceWarmup = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetInitialMaxSize(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.InitialMaxSize = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetMaxSize(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.MaxSize = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetMetricName(v string) *DescribeScalingRulesResponseBodyScalingRules {
	s.MetricName = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetMinAdjustmentMagnitude(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.MinAdjustmentMagnitude = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetMinSize(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.MinSize = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetPredictiveScalingMode(v string) *DescribeScalingRulesResponseBodyScalingRules {
	s.PredictiveScalingMode = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetPredictiveTaskBufferTime(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.PredictiveTaskBufferTime = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetPredictiveValueBehavior(v string) *DescribeScalingRulesResponseBodyScalingRules {
	s.PredictiveValueBehavior = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetPredictiveValueBuffer(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.PredictiveValueBuffer = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetScaleInEvaluationCount(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.ScaleInEvaluationCount = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetScaleOutEvaluationCount(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.ScaleOutEvaluationCount = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetScalingGroupId(v string) *DescribeScalingRulesResponseBodyScalingRules {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetScalingRuleAri(v string) *DescribeScalingRulesResponseBodyScalingRules {
	s.ScalingRuleAri = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetScalingRuleId(v string) *DescribeScalingRulesResponseBodyScalingRules {
	s.ScalingRuleId = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetScalingRuleName(v string) *DescribeScalingRulesResponseBodyScalingRules {
	s.ScalingRuleName = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetScalingRuleType(v string) *DescribeScalingRulesResponseBodyScalingRules {
	s.ScalingRuleType = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetStepAdjustments(v []*DescribeScalingRulesResponseBodyScalingRulesStepAdjustments) *DescribeScalingRulesResponseBodyScalingRules {
	s.StepAdjustments = v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetTargetValue(v float32) *DescribeScalingRulesResponseBodyScalingRules {
	s.TargetValue = &v
	return s
}

type DescribeScalingRulesResponseBodyScalingRulesAlarms struct {
	AlarmTaskId        *string                                                         `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	AlarmTaskName      *string                                                         `json:"AlarmTaskName,omitempty" xml:"AlarmTaskName,omitempty"`
	ComparisonOperator *string                                                         `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Dimensions         []*DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	EvaluationCount    *int32                                                          `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	MetricName         *string                                                         `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	MetricType         *string                                                         `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Statistics         *string                                                         `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *float32                                                        `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeScalingRulesResponseBodyScalingRulesAlarms) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRulesAlarms) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) SetAlarmTaskId(v string) *DescribeScalingRulesResponseBodyScalingRulesAlarms {
	s.AlarmTaskId = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) SetAlarmTaskName(v string) *DescribeScalingRulesResponseBodyScalingRulesAlarms {
	s.AlarmTaskName = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) SetComparisonOperator(v string) *DescribeScalingRulesResponseBodyScalingRulesAlarms {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) SetDimensions(v []*DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions) *DescribeScalingRulesResponseBodyScalingRulesAlarms {
	s.Dimensions = v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) SetEvaluationCount(v int32) *DescribeScalingRulesResponseBodyScalingRulesAlarms {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) SetMetricName(v string) *DescribeScalingRulesResponseBodyScalingRulesAlarms {
	s.MetricName = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) SetMetricType(v string) *DescribeScalingRulesResponseBodyScalingRulesAlarms {
	s.MetricType = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) SetStatistics(v string) *DescribeScalingRulesResponseBodyScalingRulesAlarms {
	s.Statistics = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) SetThreshold(v float32) *DescribeScalingRulesResponseBodyScalingRulesAlarms {
	s.Threshold = &v
	return s
}

type DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions struct {
	DimensionKey   *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
}

func (s DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions) SetDimensionKey(v string) *DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions {
	s.DimensionKey = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions) SetDimensionValue(v string) *DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions {
	s.DimensionValue = &v
	return s
}

type DescribeScalingRulesResponseBodyScalingRulesStepAdjustments struct {
	MetricIntervalLowerBound *float32 `json:"MetricIntervalLowerBound,omitempty" xml:"MetricIntervalLowerBound,omitempty"`
	MetricIntervalUpperBound *float32 `json:"MetricIntervalUpperBound,omitempty" xml:"MetricIntervalUpperBound,omitempty"`
	ScalingAdjustment        *int32   `json:"ScalingAdjustment,omitempty" xml:"ScalingAdjustment,omitempty"`
}

func (s DescribeScalingRulesResponseBodyScalingRulesStepAdjustments) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRulesStepAdjustments) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRulesStepAdjustments) SetMetricIntervalLowerBound(v float32) *DescribeScalingRulesResponseBodyScalingRulesStepAdjustments {
	s.MetricIntervalLowerBound = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesStepAdjustments) SetMetricIntervalUpperBound(v float32) *DescribeScalingRulesResponseBodyScalingRulesStepAdjustments {
	s.MetricIntervalUpperBound = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesStepAdjustments) SetScalingAdjustment(v int32) *DescribeScalingRulesResponseBodyScalingRulesStepAdjustments {
	s.ScalingAdjustment = &v
	return s
}

type DescribeScalingRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScalingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScalingRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponse) SetHeaders(v map[string]*string) *DescribeScalingRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingRulesResponse) SetStatusCode(v int32) *DescribeScalingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingRulesResponse) SetBody(v *DescribeScalingRulesResponseBody) *DescribeScalingRulesResponse {
	s.Body = v
	return s
}

type DescribeScheduledTasksRequest struct {
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScheduledActions     []*string `json:"ScheduledActions,omitempty" xml:"ScheduledActions,omitempty" type:"Repeated"`
	ScheduledTaskIds     []*string `json:"ScheduledTaskIds,omitempty" xml:"ScheduledTaskIds,omitempty" type:"Repeated"`
	ScheduledTaskNames   []*string `json:"ScheduledTaskNames,omitempty" xml:"ScheduledTaskNames,omitempty" type:"Repeated"`
}

func (s DescribeScheduledTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduledTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksRequest) SetOwnerAccount(v string) *DescribeScheduledTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetOwnerId(v int64) *DescribeScheduledTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetPageNumber(v int32) *DescribeScheduledTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetPageSize(v int32) *DescribeScheduledTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetRegionId(v string) *DescribeScheduledTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetResourceOwnerAccount(v string) *DescribeScheduledTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetResourceOwnerId(v int64) *DescribeScheduledTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetScalingGroupId(v string) *DescribeScheduledTasksRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetScheduledActions(v []*string) *DescribeScheduledTasksRequest {
	s.ScheduledActions = v
	return s
}

func (s *DescribeScheduledTasksRequest) SetScheduledTaskIds(v []*string) *DescribeScheduledTasksRequest {
	s.ScheduledTaskIds = v
	return s
}

func (s *DescribeScheduledTasksRequest) SetScheduledTaskNames(v []*string) *DescribeScheduledTasksRequest {
	s.ScheduledTaskNames = v
	return s
}

type DescribeScheduledTasksResponseBody struct {
	PageNumber     *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScheduledTasks []*DescribeScheduledTasksResponseBodyScheduledTasks `json:"ScheduledTasks,omitempty" xml:"ScheduledTasks,omitempty" type:"Repeated"`
	TotalCount     *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScheduledTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduledTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksResponseBody) SetPageNumber(v int32) *DescribeScheduledTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetPageSize(v int32) *DescribeScheduledTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetRequestId(v string) *DescribeScheduledTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetScheduledTasks(v []*DescribeScheduledTasksResponseBodyScheduledTasks) *DescribeScheduledTasksResponseBody {
	s.ScheduledTasks = v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetTotalCount(v int32) *DescribeScheduledTasksResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeScheduledTasksResponseBodyScheduledTasks struct {
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DesiredCapacity      *int32  `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	LaunchExpirationTime *int32  `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
	LaunchTime           *string `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	MaxValue             *int32  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinValue             *int32  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	RecurrenceEndTime    *string `json:"RecurrenceEndTime,omitempty" xml:"RecurrenceEndTime,omitempty"`
	RecurrenceType       *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValue      *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScheduledAction      *string `json:"ScheduledAction,omitempty" xml:"ScheduledAction,omitempty"`
	ScheduledTaskId      *string `json:"ScheduledTaskId,omitempty" xml:"ScheduledTaskId,omitempty"`
	ScheduledTaskName    *string `json:"ScheduledTaskName,omitempty" xml:"ScheduledTaskName,omitempty"`
	TaskEnabled          *bool   `json:"TaskEnabled,omitempty" xml:"TaskEnabled,omitempty"`
}

func (s DescribeScheduledTasksResponseBodyScheduledTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduledTasksResponseBodyScheduledTasks) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetDescription(v string) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.Description = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetDesiredCapacity(v int32) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.DesiredCapacity = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetLaunchExpirationTime(v int32) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.LaunchExpirationTime = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetLaunchTime(v string) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.LaunchTime = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetMaxValue(v int32) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.MaxValue = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetMinValue(v int32) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.MinValue = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetRecurrenceEndTime(v string) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.RecurrenceEndTime = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetRecurrenceType(v string) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.RecurrenceType = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetRecurrenceValue(v string) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.RecurrenceValue = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetScalingGroupId(v string) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetScheduledAction(v string) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.ScheduledAction = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetScheduledTaskId(v string) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.ScheduledTaskId = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetScheduledTaskName(v string) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.ScheduledTaskName = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetTaskEnabled(v bool) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.TaskEnabled = &v
	return s
}

type DescribeScheduledTasksResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScheduledTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScheduledTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduledTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksResponse) SetHeaders(v map[string]*string) *DescribeScheduledTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeScheduledTasksResponse) SetStatusCode(v int32) *DescribeScheduledTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScheduledTasksResponse) SetBody(v *DescribeScheduledTasksResponseBody) *DescribeScheduledTasksResponse {
	s.Body = v
	return s
}

type DetachAlbServerGroupsRequest struct {
	AlbServerGroups      []*DetachAlbServerGroupsRequestAlbServerGroups `json:"AlbServerGroups,omitempty" xml:"AlbServerGroups,omitempty" type:"Repeated"`
	ClientToken          *string                                        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ForceDetach          *bool                                          `json:"ForceDetach,omitempty" xml:"ForceDetach,omitempty"`
	OwnerId              *int64                                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string                                        `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string                                        `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DetachAlbServerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachAlbServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *DetachAlbServerGroupsRequest) SetAlbServerGroups(v []*DetachAlbServerGroupsRequestAlbServerGroups) *DetachAlbServerGroupsRequest {
	s.AlbServerGroups = v
	return s
}

func (s *DetachAlbServerGroupsRequest) SetClientToken(v string) *DetachAlbServerGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachAlbServerGroupsRequest) SetForceDetach(v bool) *DetachAlbServerGroupsRequest {
	s.ForceDetach = &v
	return s
}

func (s *DetachAlbServerGroupsRequest) SetOwnerId(v int64) *DetachAlbServerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachAlbServerGroupsRequest) SetRegionId(v string) *DetachAlbServerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DetachAlbServerGroupsRequest) SetResourceOwnerAccount(v string) *DetachAlbServerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachAlbServerGroupsRequest) SetScalingGroupId(v string) *DetachAlbServerGroupsRequest {
	s.ScalingGroupId = &v
	return s
}

type DetachAlbServerGroupsRequestAlbServerGroups struct {
	AlbServerGroupId *string `json:"AlbServerGroupId,omitempty" xml:"AlbServerGroupId,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DetachAlbServerGroupsRequestAlbServerGroups) String() string {
	return tea.Prettify(s)
}

func (s DetachAlbServerGroupsRequestAlbServerGroups) GoString() string {
	return s.String()
}

func (s *DetachAlbServerGroupsRequestAlbServerGroups) SetAlbServerGroupId(v string) *DetachAlbServerGroupsRequestAlbServerGroups {
	s.AlbServerGroupId = &v
	return s
}

func (s *DetachAlbServerGroupsRequestAlbServerGroups) SetPort(v int32) *DetachAlbServerGroupsRequestAlbServerGroups {
	s.Port = &v
	return s
}

type DetachAlbServerGroupsResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s DetachAlbServerGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachAlbServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DetachAlbServerGroupsResponseBody) SetRequestId(v string) *DetachAlbServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachAlbServerGroupsResponseBody) SetScalingActivityId(v string) *DetachAlbServerGroupsResponseBody {
	s.ScalingActivityId = &v
	return s
}

type DetachAlbServerGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachAlbServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachAlbServerGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachAlbServerGroupsResponse) GoString() string {
	return s.String()
}

func (s *DetachAlbServerGroupsResponse) SetHeaders(v map[string]*string) *DetachAlbServerGroupsResponse {
	s.Headers = v
	return s
}

func (s *DetachAlbServerGroupsResponse) SetStatusCode(v int32) *DetachAlbServerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachAlbServerGroupsResponse) SetBody(v *DetachAlbServerGroupsResponseBody) *DetachAlbServerGroupsResponse {
	s.Body = v
	return s
}

type DetachDBInstancesRequest struct {
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstances          []*string `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	ForceDetach          *bool     `json:"ForceDetach,omitempty" xml:"ForceDetach,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DetachDBInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *DetachDBInstancesRequest) SetClientToken(v string) *DetachDBInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachDBInstancesRequest) SetDBInstances(v []*string) *DetachDBInstancesRequest {
	s.DBInstances = v
	return s
}

func (s *DetachDBInstancesRequest) SetForceDetach(v bool) *DetachDBInstancesRequest {
	s.ForceDetach = &v
	return s
}

func (s *DetachDBInstancesRequest) SetOwnerId(v int64) *DetachDBInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachDBInstancesRequest) SetRegionId(v string) *DetachDBInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DetachDBInstancesRequest) SetResourceOwnerAccount(v string) *DetachDBInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachDBInstancesRequest) SetScalingGroupId(v string) *DetachDBInstancesRequest {
	s.ScalingGroupId = &v
	return s
}

type DetachDBInstancesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachDBInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachDBInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DetachDBInstancesResponseBody) SetRequestId(v string) *DetachDBInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DetachDBInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachDBInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachDBInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachDBInstancesResponse) GoString() string {
	return s.String()
}

func (s *DetachDBInstancesResponse) SetHeaders(v map[string]*string) *DetachDBInstancesResponse {
	s.Headers = v
	return s
}

func (s *DetachDBInstancesResponse) SetStatusCode(v int32) *DetachDBInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachDBInstancesResponse) SetBody(v *DetachDBInstancesResponseBody) *DetachDBInstancesResponse {
	s.Body = v
	return s
}

type DetachInstancesRequest struct {
	DecreaseDesiredCapacity *bool     `json:"DecreaseDesiredCapacity,omitempty" xml:"DecreaseDesiredCapacity,omitempty"`
	DetachOption            *string   `json:"DetachOption,omitempty" xml:"DetachOption,omitempty"`
	InstanceIds             []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	LifecycleHook           *bool     `json:"LifecycleHook,omitempty" xml:"LifecycleHook,omitempty"`
	OwnerAccount            *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount    *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupId          *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DetachInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachInstancesRequest) GoString() string {
	return s.String()
}

func (s *DetachInstancesRequest) SetDecreaseDesiredCapacity(v bool) *DetachInstancesRequest {
	s.DecreaseDesiredCapacity = &v
	return s
}

func (s *DetachInstancesRequest) SetDetachOption(v string) *DetachInstancesRequest {
	s.DetachOption = &v
	return s
}

func (s *DetachInstancesRequest) SetInstanceIds(v []*string) *DetachInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *DetachInstancesRequest) SetLifecycleHook(v bool) *DetachInstancesRequest {
	s.LifecycleHook = &v
	return s
}

func (s *DetachInstancesRequest) SetOwnerAccount(v string) *DetachInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DetachInstancesRequest) SetOwnerId(v int64) *DetachInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachInstancesRequest) SetResourceOwnerAccount(v string) *DetachInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachInstancesRequest) SetResourceOwnerId(v int64) *DetachInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DetachInstancesRequest) SetScalingGroupId(v string) *DetachInstancesRequest {
	s.ScalingGroupId = &v
	return s
}

type DetachInstancesResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s DetachInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DetachInstancesResponseBody) SetRequestId(v string) *DetachInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachInstancesResponseBody) SetScalingActivityId(v string) *DetachInstancesResponseBody {
	s.ScalingActivityId = &v
	return s
}

type DetachInstancesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachInstancesResponse) GoString() string {
	return s.String()
}

func (s *DetachInstancesResponse) SetHeaders(v map[string]*string) *DetachInstancesResponse {
	s.Headers = v
	return s
}

func (s *DetachInstancesResponse) SetStatusCode(v int32) *DetachInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachInstancesResponse) SetBody(v *DetachInstancesResponseBody) *DetachInstancesResponse {
	s.Body = v
	return s
}

type DetachLoadBalancersRequest struct {
	Async                *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ForceDetach          *bool     `json:"ForceDetach,omitempty" xml:"ForceDetach,omitempty"`
	LoadBalancers        []*string `json:"LoadBalancers,omitempty" xml:"LoadBalancers,omitempty" type:"Repeated"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DetachLoadBalancersRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachLoadBalancersRequest) GoString() string {
	return s.String()
}

func (s *DetachLoadBalancersRequest) SetAsync(v bool) *DetachLoadBalancersRequest {
	s.Async = &v
	return s
}

func (s *DetachLoadBalancersRequest) SetClientToken(v string) *DetachLoadBalancersRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachLoadBalancersRequest) SetForceDetach(v bool) *DetachLoadBalancersRequest {
	s.ForceDetach = &v
	return s
}

func (s *DetachLoadBalancersRequest) SetLoadBalancers(v []*string) *DetachLoadBalancersRequest {
	s.LoadBalancers = v
	return s
}

func (s *DetachLoadBalancersRequest) SetOwnerId(v int64) *DetachLoadBalancersRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachLoadBalancersRequest) SetRegionId(v string) *DetachLoadBalancersRequest {
	s.RegionId = &v
	return s
}

func (s *DetachLoadBalancersRequest) SetResourceOwnerAccount(v string) *DetachLoadBalancersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachLoadBalancersRequest) SetScalingGroupId(v string) *DetachLoadBalancersRequest {
	s.ScalingGroupId = &v
	return s
}

type DetachLoadBalancersResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s DetachLoadBalancersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachLoadBalancersResponseBody) GoString() string {
	return s.String()
}

func (s *DetachLoadBalancersResponseBody) SetRequestId(v string) *DetachLoadBalancersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachLoadBalancersResponseBody) SetScalingActivityId(v string) *DetachLoadBalancersResponseBody {
	s.ScalingActivityId = &v
	return s
}

type DetachLoadBalancersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachLoadBalancersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachLoadBalancersResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachLoadBalancersResponse) GoString() string {
	return s.String()
}

func (s *DetachLoadBalancersResponse) SetHeaders(v map[string]*string) *DetachLoadBalancersResponse {
	s.Headers = v
	return s
}

func (s *DetachLoadBalancersResponse) SetStatusCode(v int32) *DetachLoadBalancersResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachLoadBalancersResponse) SetBody(v *DetachLoadBalancersResponseBody) *DetachLoadBalancersResponse {
	s.Body = v
	return s
}

type DetachVServerGroupsRequest struct {
	ClientToken          *string                                    `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ForceDetach          *bool                                      `json:"ForceDetach,omitempty" xml:"ForceDetach,omitempty"`
	OwnerId              *int64                                     `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string                                    `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string                                    `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	VServerGroups        []*DetachVServerGroupsRequestVServerGroups `json:"VServerGroups,omitempty" xml:"VServerGroups,omitempty" type:"Repeated"`
}

func (s DetachVServerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachVServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *DetachVServerGroupsRequest) SetClientToken(v string) *DetachVServerGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachVServerGroupsRequest) SetForceDetach(v bool) *DetachVServerGroupsRequest {
	s.ForceDetach = &v
	return s
}

func (s *DetachVServerGroupsRequest) SetOwnerId(v int64) *DetachVServerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachVServerGroupsRequest) SetRegionId(v string) *DetachVServerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DetachVServerGroupsRequest) SetResourceOwnerAccount(v string) *DetachVServerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachVServerGroupsRequest) SetScalingGroupId(v string) *DetachVServerGroupsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DetachVServerGroupsRequest) SetVServerGroups(v []*DetachVServerGroupsRequestVServerGroups) *DetachVServerGroupsRequest {
	s.VServerGroups = v
	return s
}

type DetachVServerGroupsRequestVServerGroups struct {
	LoadBalancerId         *string                                                          `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	VServerGroupAttributes []*DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes `json:"VServerGroupAttributes,omitempty" xml:"VServerGroupAttributes,omitempty" type:"Repeated"`
}

func (s DetachVServerGroupsRequestVServerGroups) String() string {
	return tea.Prettify(s)
}

func (s DetachVServerGroupsRequestVServerGroups) GoString() string {
	return s.String()
}

func (s *DetachVServerGroupsRequestVServerGroups) SetLoadBalancerId(v string) *DetachVServerGroupsRequestVServerGroups {
	s.LoadBalancerId = &v
	return s
}

func (s *DetachVServerGroupsRequestVServerGroups) SetVServerGroupAttributes(v []*DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes) *DetachVServerGroupsRequestVServerGroups {
	s.VServerGroupAttributes = v
	return s
}

type DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes struct {
	Port           *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes) String() string {
	return tea.Prettify(s)
}

func (s DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes) GoString() string {
	return s.String()
}

func (s *DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes) SetPort(v int32) *DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes {
	s.Port = &v
	return s
}

func (s *DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes) SetVServerGroupId(v string) *DetachVServerGroupsRequestVServerGroupsVServerGroupAttributes {
	s.VServerGroupId = &v
	return s
}

type DetachVServerGroupsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachVServerGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachVServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DetachVServerGroupsResponseBody) SetRequestId(v string) *DetachVServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

type DetachVServerGroupsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachVServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachVServerGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachVServerGroupsResponse) GoString() string {
	return s.String()
}

func (s *DetachVServerGroupsResponse) SetHeaders(v map[string]*string) *DetachVServerGroupsResponse {
	s.Headers = v
	return s
}

func (s *DetachVServerGroupsResponse) SetStatusCode(v int32) *DetachVServerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachVServerGroupsResponse) SetBody(v *DetachVServerGroupsResponseBody) *DetachVServerGroupsResponse {
	s.Body = v
	return s
}

type DisableAlarmRequest struct {
	AlarmTaskId          *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DisableAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableAlarmRequest) GoString() string {
	return s.String()
}

func (s *DisableAlarmRequest) SetAlarmTaskId(v string) *DisableAlarmRequest {
	s.AlarmTaskId = &v
	return s
}

func (s *DisableAlarmRequest) SetOwnerId(v int64) *DisableAlarmRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableAlarmRequest) SetRegionId(v string) *DisableAlarmRequest {
	s.RegionId = &v
	return s
}

func (s *DisableAlarmRequest) SetResourceOwnerAccount(v string) *DisableAlarmRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type DisableAlarmResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAlarmResponseBody) SetRequestId(v string) *DisableAlarmResponseBody {
	s.RequestId = &v
	return s
}

type DisableAlarmResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableAlarmResponse) GoString() string {
	return s.String()
}

func (s *DisableAlarmResponse) SetHeaders(v map[string]*string) *DisableAlarmResponse {
	s.Headers = v
	return s
}

func (s *DisableAlarmResponse) SetStatusCode(v int32) *DisableAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAlarmResponse) SetBody(v *DisableAlarmResponseBody) *DisableAlarmResponse {
	s.Body = v
	return s
}

type DisableScalingGroupRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DisableScalingGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *DisableScalingGroupRequest) SetOwnerAccount(v string) *DisableScalingGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DisableScalingGroupRequest) SetOwnerId(v int64) *DisableScalingGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableScalingGroupRequest) SetResourceOwnerAccount(v string) *DisableScalingGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisableScalingGroupRequest) SetResourceOwnerId(v int64) *DisableScalingGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisableScalingGroupRequest) SetScalingGroupId(v string) *DisableScalingGroupRequest {
	s.ScalingGroupId = &v
	return s
}

type DisableScalingGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableScalingGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableScalingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DisableScalingGroupResponseBody) SetRequestId(v string) *DisableScalingGroupResponseBody {
	s.RequestId = &v
	return s
}

type DisableScalingGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableScalingGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableScalingGroupResponse) GoString() string {
	return s.String()
}

func (s *DisableScalingGroupResponse) SetHeaders(v map[string]*string) *DisableScalingGroupResponse {
	s.Headers = v
	return s
}

func (s *DisableScalingGroupResponse) SetStatusCode(v int32) *DisableScalingGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableScalingGroupResponse) SetBody(v *DisableScalingGroupResponseBody) *DisableScalingGroupResponse {
	s.Body = v
	return s
}

type EnableAlarmRequest struct {
	AlarmTaskId          *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s EnableAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableAlarmRequest) GoString() string {
	return s.String()
}

func (s *EnableAlarmRequest) SetAlarmTaskId(v string) *EnableAlarmRequest {
	s.AlarmTaskId = &v
	return s
}

func (s *EnableAlarmRequest) SetOwnerId(v int64) *EnableAlarmRequest {
	s.OwnerId = &v
	return s
}

func (s *EnableAlarmRequest) SetRegionId(v string) *EnableAlarmRequest {
	s.RegionId = &v
	return s
}

func (s *EnableAlarmRequest) SetResourceOwnerAccount(v string) *EnableAlarmRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type EnableAlarmResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *EnableAlarmResponseBody) SetRequestId(v string) *EnableAlarmResponseBody {
	s.RequestId = &v
	return s
}

type EnableAlarmResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableAlarmResponse) GoString() string {
	return s.String()
}

func (s *EnableAlarmResponse) SetHeaders(v map[string]*string) *EnableAlarmResponse {
	s.Headers = v
	return s
}

func (s *EnableAlarmResponse) SetStatusCode(v int32) *EnableAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableAlarmResponse) SetBody(v *EnableAlarmResponseBody) *EnableAlarmResponse {
	s.Body = v
	return s
}

type EnableScalingGroupRequest struct {
	ActiveScalingConfigurationId *string                                             `json:"ActiveScalingConfigurationId,omitempty" xml:"ActiveScalingConfigurationId,omitempty"`
	InstanceIds                  []*string                                           `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	LaunchTemplateId             *string                                             `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	LaunchTemplateOverrides      []*EnableScalingGroupRequestLaunchTemplateOverrides `json:"LaunchTemplateOverrides,omitempty" xml:"LaunchTemplateOverrides,omitempty" type:"Repeated"`
	LaunchTemplateVersion        *string                                             `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	LoadBalancerWeights          []*int32                                            `json:"LoadBalancerWeights,omitempty" xml:"LoadBalancerWeights,omitempty" type:"Repeated"`
	OwnerAccount                 *string                                             `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                      *int64                                              `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                     *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount         *string                                             `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId              *int64                                              `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupId               *string                                             `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s EnableScalingGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *EnableScalingGroupRequest) SetActiveScalingConfigurationId(v string) *EnableScalingGroupRequest {
	s.ActiveScalingConfigurationId = &v
	return s
}

func (s *EnableScalingGroupRequest) SetInstanceIds(v []*string) *EnableScalingGroupRequest {
	s.InstanceIds = v
	return s
}

func (s *EnableScalingGroupRequest) SetLaunchTemplateId(v string) *EnableScalingGroupRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *EnableScalingGroupRequest) SetLaunchTemplateOverrides(v []*EnableScalingGroupRequestLaunchTemplateOverrides) *EnableScalingGroupRequest {
	s.LaunchTemplateOverrides = v
	return s
}

func (s *EnableScalingGroupRequest) SetLaunchTemplateVersion(v string) *EnableScalingGroupRequest {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *EnableScalingGroupRequest) SetLoadBalancerWeights(v []*int32) *EnableScalingGroupRequest {
	s.LoadBalancerWeights = v
	return s
}

func (s *EnableScalingGroupRequest) SetOwnerAccount(v string) *EnableScalingGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *EnableScalingGroupRequest) SetOwnerId(v int64) *EnableScalingGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *EnableScalingGroupRequest) SetRegionId(v string) *EnableScalingGroupRequest {
	s.RegionId = &v
	return s
}

func (s *EnableScalingGroupRequest) SetResourceOwnerAccount(v string) *EnableScalingGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *EnableScalingGroupRequest) SetResourceOwnerId(v int64) *EnableScalingGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *EnableScalingGroupRequest) SetScalingGroupId(v string) *EnableScalingGroupRequest {
	s.ScalingGroupId = &v
	return s
}

type EnableScalingGroupRequestLaunchTemplateOverrides struct {
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	WeightedCapacity *int32  `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s EnableScalingGroupRequestLaunchTemplateOverrides) String() string {
	return tea.Prettify(s)
}

func (s EnableScalingGroupRequestLaunchTemplateOverrides) GoString() string {
	return s.String()
}

func (s *EnableScalingGroupRequestLaunchTemplateOverrides) SetInstanceType(v string) *EnableScalingGroupRequestLaunchTemplateOverrides {
	s.InstanceType = &v
	return s
}

func (s *EnableScalingGroupRequestLaunchTemplateOverrides) SetWeightedCapacity(v int32) *EnableScalingGroupRequestLaunchTemplateOverrides {
	s.WeightedCapacity = &v
	return s
}

type EnableScalingGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableScalingGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableScalingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *EnableScalingGroupResponseBody) SetRequestId(v string) *EnableScalingGroupResponseBody {
	s.RequestId = &v
	return s
}

type EnableScalingGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableScalingGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableScalingGroupResponse) GoString() string {
	return s.String()
}

func (s *EnableScalingGroupResponse) SetHeaders(v map[string]*string) *EnableScalingGroupResponse {
	s.Headers = v
	return s
}

func (s *EnableScalingGroupResponse) SetStatusCode(v int32) *EnableScalingGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableScalingGroupResponse) SetBody(v *EnableScalingGroupResponseBody) *EnableScalingGroupResponse {
	s.Body = v
	return s
}

type EnterStandbyRequest struct {
	Async                *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceIds          []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s EnterStandbyRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterStandbyRequest) GoString() string {
	return s.String()
}

func (s *EnterStandbyRequest) SetAsync(v bool) *EnterStandbyRequest {
	s.Async = &v
	return s
}

func (s *EnterStandbyRequest) SetClientToken(v string) *EnterStandbyRequest {
	s.ClientToken = &v
	return s
}

func (s *EnterStandbyRequest) SetInstanceIds(v []*string) *EnterStandbyRequest {
	s.InstanceIds = v
	return s
}

func (s *EnterStandbyRequest) SetOwnerId(v int64) *EnterStandbyRequest {
	s.OwnerId = &v
	return s
}

func (s *EnterStandbyRequest) SetResourceOwnerAccount(v string) *EnterStandbyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *EnterStandbyRequest) SetScalingGroupId(v string) *EnterStandbyRequest {
	s.ScalingGroupId = &v
	return s
}

type EnterStandbyResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s EnterStandbyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnterStandbyResponseBody) GoString() string {
	return s.String()
}

func (s *EnterStandbyResponseBody) SetRequestId(v string) *EnterStandbyResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnterStandbyResponseBody) SetScalingActivityId(v string) *EnterStandbyResponseBody {
	s.ScalingActivityId = &v
	return s
}

type EnterStandbyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnterStandbyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnterStandbyResponse) String() string {
	return tea.Prettify(s)
}

func (s EnterStandbyResponse) GoString() string {
	return s.String()
}

func (s *EnterStandbyResponse) SetHeaders(v map[string]*string) *EnterStandbyResponse {
	s.Headers = v
	return s
}

func (s *EnterStandbyResponse) SetStatusCode(v int32) *EnterStandbyResponse {
	s.StatusCode = &v
	return s
}

func (s *EnterStandbyResponse) SetBody(v *EnterStandbyResponseBody) *EnterStandbyResponse {
	s.Body = v
	return s
}

type ExecuteScalingRuleRequest struct {
	BreachThreshold      *float32 `json:"BreachThreshold,omitempty" xml:"BreachThreshold,omitempty"`
	ClientToken          *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	MetricValue          *float32 `json:"MetricValue,omitempty" xml:"MetricValue,omitempty"`
	OwnerAccount         *string  `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string  `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64   `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingRuleAri       *string  `json:"ScalingRuleAri,omitempty" xml:"ScalingRuleAri,omitempty"`
}

func (s ExecuteScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *ExecuteScalingRuleRequest) SetBreachThreshold(v float32) *ExecuteScalingRuleRequest {
	s.BreachThreshold = &v
	return s
}

func (s *ExecuteScalingRuleRequest) SetClientToken(v string) *ExecuteScalingRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *ExecuteScalingRuleRequest) SetMetricValue(v float32) *ExecuteScalingRuleRequest {
	s.MetricValue = &v
	return s
}

func (s *ExecuteScalingRuleRequest) SetOwnerAccount(v string) *ExecuteScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ExecuteScalingRuleRequest) SetOwnerId(v int64) *ExecuteScalingRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *ExecuteScalingRuleRequest) SetRegionId(v string) *ExecuteScalingRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ExecuteScalingRuleRequest) SetResourceOwnerAccount(v string) *ExecuteScalingRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ExecuteScalingRuleRequest) SetResourceOwnerId(v int64) *ExecuteScalingRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ExecuteScalingRuleRequest) SetScalingRuleAri(v string) *ExecuteScalingRuleRequest {
	s.ScalingRuleAri = &v
	return s
}

type ExecuteScalingRuleResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s ExecuteScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteScalingRuleResponseBody) SetRequestId(v string) *ExecuteScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteScalingRuleResponseBody) SetScalingActivityId(v string) *ExecuteScalingRuleResponseBody {
	s.ScalingActivityId = &v
	return s
}

type ExecuteScalingRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExecuteScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *ExecuteScalingRuleResponse) SetHeaders(v map[string]*string) *ExecuteScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *ExecuteScalingRuleResponse) SetStatusCode(v int32) *ExecuteScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteScalingRuleResponse) SetBody(v *ExecuteScalingRuleResponseBody) *ExecuteScalingRuleResponse {
	s.Body = v
	return s
}

type ExitStandbyRequest struct {
	Async                *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceIds          []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s ExitStandbyRequest) String() string {
	return tea.Prettify(s)
}

func (s ExitStandbyRequest) GoString() string {
	return s.String()
}

func (s *ExitStandbyRequest) SetAsync(v bool) *ExitStandbyRequest {
	s.Async = &v
	return s
}

func (s *ExitStandbyRequest) SetClientToken(v string) *ExitStandbyRequest {
	s.ClientToken = &v
	return s
}

func (s *ExitStandbyRequest) SetInstanceIds(v []*string) *ExitStandbyRequest {
	s.InstanceIds = v
	return s
}

func (s *ExitStandbyRequest) SetOwnerId(v int64) *ExitStandbyRequest {
	s.OwnerId = &v
	return s
}

func (s *ExitStandbyRequest) SetRegionId(v string) *ExitStandbyRequest {
	s.RegionId = &v
	return s
}

func (s *ExitStandbyRequest) SetResourceOwnerAccount(v string) *ExitStandbyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ExitStandbyRequest) SetScalingGroupId(v string) *ExitStandbyRequest {
	s.ScalingGroupId = &v
	return s
}

type ExitStandbyResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s ExitStandbyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExitStandbyResponseBody) GoString() string {
	return s.String()
}

func (s *ExitStandbyResponseBody) SetRequestId(v string) *ExitStandbyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExitStandbyResponseBody) SetScalingActivityId(v string) *ExitStandbyResponseBody {
	s.ScalingActivityId = &v
	return s
}

type ExitStandbyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExitStandbyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExitStandbyResponse) String() string {
	return tea.Prettify(s)
}

func (s ExitStandbyResponse) GoString() string {
	return s.String()
}

func (s *ExitStandbyResponse) SetHeaders(v map[string]*string) *ExitStandbyResponse {
	s.Headers = v
	return s
}

func (s *ExitStandbyResponse) SetStatusCode(v int32) *ExitStandbyResponse {
	s.StatusCode = &v
	return s
}

func (s *ExitStandbyResponse) SetBody(v *ExitStandbyResponseBody) *ExitStandbyResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetOwnerId(v int64) *ListTagKeysRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int32) *ListTagKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceOwnerAccount(v string) *ListTagKeysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

type ListTagKeysResponseBody struct {
	Keys      []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	NextToken *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageSize  *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetKeys(v []*string) *ListTagKeysResponseBody {
	s.Keys = v
	return s
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetPageSize(v int32) *ListTagKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

type ListTagKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	NextToken            *string                        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId              *int64                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceIds          []*string                      `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                        `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceType         *string                        `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags                 []*ListTagResourcesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceIds(v []*string) *ListTagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTags(v []*ListTagResourcesRequestTags) *ListTagResourcesRequest {
	s.Tags = v
	return s
}

type ListTagResourcesRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTags) SetKey(v string) *ListTagResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTags) SetValue(v string) *ListTagResourcesRequestTags {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
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

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	Propagate    *bool   `json:"Propagate,omitempty" xml:"Propagate,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetPropagate(v bool) *ListTagResourcesResponseBodyTagResources {
	s.Propagate = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
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

type ListTagValuesRequest struct {
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetKey(v string) *ListTagValuesRequest {
	s.Key = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetOwnerId(v int64) *ListTagValuesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagValuesRequest) SetPageSize(v int32) *ListTagValuesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagValuesRequest) SetRegionId(v string) *ListTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceOwnerAccount(v string) *ListTagValuesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

type ListTagValuesResponseBody struct {
	NextToken *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageSize  *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Values    []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetPageSize(v int32) *ListTagValuesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetValues(v []*string) *ListTagValuesResponseBody {
	s.Values = v
	return s
}

type ListTagValuesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetStatusCode(v int32) *ListTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type ModifyAlarmRequest struct {
	AlarmActions             []*string                        `json:"AlarmActions,omitempty" xml:"AlarmActions,omitempty" type:"Repeated"`
	AlarmTaskId              *string                          `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	ComparisonOperator       *string                          `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Description              *string                          `json:"Description,omitempty" xml:"Description,omitempty"`
	Dimensions               []*ModifyAlarmRequestDimensions  `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	Effective                *string                          `json:"Effective,omitempty" xml:"Effective,omitempty"`
	EvaluationCount          *int32                           `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	Expressions              []*ModifyAlarmRequestExpressions `json:"Expressions,omitempty" xml:"Expressions,omitempty" type:"Repeated"`
	ExpressionsLogicOperator *string                          `json:"ExpressionsLogicOperator,omitempty" xml:"ExpressionsLogicOperator,omitempty"`
	GroupId                  *int32                           `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MetricName               *string                          `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	MetricType               *string                          `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Name                     *string                          `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId                  *int64                           `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Period                   *int32                           `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId                 *string                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount     *string                          `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	Statistics               *string                          `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold                *float32                         `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ModifyAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAlarmRequest) GoString() string {
	return s.String()
}

func (s *ModifyAlarmRequest) SetAlarmActions(v []*string) *ModifyAlarmRequest {
	s.AlarmActions = v
	return s
}

func (s *ModifyAlarmRequest) SetAlarmTaskId(v string) *ModifyAlarmRequest {
	s.AlarmTaskId = &v
	return s
}

func (s *ModifyAlarmRequest) SetComparisonOperator(v string) *ModifyAlarmRequest {
	s.ComparisonOperator = &v
	return s
}

func (s *ModifyAlarmRequest) SetDescription(v string) *ModifyAlarmRequest {
	s.Description = &v
	return s
}

func (s *ModifyAlarmRequest) SetDimensions(v []*ModifyAlarmRequestDimensions) *ModifyAlarmRequest {
	s.Dimensions = v
	return s
}

func (s *ModifyAlarmRequest) SetEffective(v string) *ModifyAlarmRequest {
	s.Effective = &v
	return s
}

func (s *ModifyAlarmRequest) SetEvaluationCount(v int32) *ModifyAlarmRequest {
	s.EvaluationCount = &v
	return s
}

func (s *ModifyAlarmRequest) SetExpressions(v []*ModifyAlarmRequestExpressions) *ModifyAlarmRequest {
	s.Expressions = v
	return s
}

func (s *ModifyAlarmRequest) SetExpressionsLogicOperator(v string) *ModifyAlarmRequest {
	s.ExpressionsLogicOperator = &v
	return s
}

func (s *ModifyAlarmRequest) SetGroupId(v int32) *ModifyAlarmRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyAlarmRequest) SetMetricName(v string) *ModifyAlarmRequest {
	s.MetricName = &v
	return s
}

func (s *ModifyAlarmRequest) SetMetricType(v string) *ModifyAlarmRequest {
	s.MetricType = &v
	return s
}

func (s *ModifyAlarmRequest) SetName(v string) *ModifyAlarmRequest {
	s.Name = &v
	return s
}

func (s *ModifyAlarmRequest) SetOwnerId(v int64) *ModifyAlarmRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAlarmRequest) SetPeriod(v int32) *ModifyAlarmRequest {
	s.Period = &v
	return s
}

func (s *ModifyAlarmRequest) SetRegionId(v string) *ModifyAlarmRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAlarmRequest) SetResourceOwnerAccount(v string) *ModifyAlarmRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAlarmRequest) SetStatistics(v string) *ModifyAlarmRequest {
	s.Statistics = &v
	return s
}

func (s *ModifyAlarmRequest) SetThreshold(v float32) *ModifyAlarmRequest {
	s.Threshold = &v
	return s
}

type ModifyAlarmRequestDimensions struct {
	DimensionKey   *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
}

func (s ModifyAlarmRequestDimensions) String() string {
	return tea.Prettify(s)
}

func (s ModifyAlarmRequestDimensions) GoString() string {
	return s.String()
}

func (s *ModifyAlarmRequestDimensions) SetDimensionKey(v string) *ModifyAlarmRequestDimensions {
	s.DimensionKey = &v
	return s
}

func (s *ModifyAlarmRequestDimensions) SetDimensionValue(v string) *ModifyAlarmRequestDimensions {
	s.DimensionValue = &v
	return s
}

type ModifyAlarmRequestExpressions struct {
	ComparisonOperator *string  `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	MetricName         *string  `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Period             *int32   `json:"Period,omitempty" xml:"Period,omitempty"`
	Statistics         *string  `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ModifyAlarmRequestExpressions) String() string {
	return tea.Prettify(s)
}

func (s ModifyAlarmRequestExpressions) GoString() string {
	return s.String()
}

func (s *ModifyAlarmRequestExpressions) SetComparisonOperator(v string) *ModifyAlarmRequestExpressions {
	s.ComparisonOperator = &v
	return s
}

func (s *ModifyAlarmRequestExpressions) SetMetricName(v string) *ModifyAlarmRequestExpressions {
	s.MetricName = &v
	return s
}

func (s *ModifyAlarmRequestExpressions) SetPeriod(v int32) *ModifyAlarmRequestExpressions {
	s.Period = &v
	return s
}

func (s *ModifyAlarmRequestExpressions) SetStatistics(v string) *ModifyAlarmRequestExpressions {
	s.Statistics = &v
	return s
}

func (s *ModifyAlarmRequestExpressions) SetThreshold(v float32) *ModifyAlarmRequestExpressions {
	s.Threshold = &v
	return s
}

type ModifyAlarmResponseBody struct {
	AlarmTaskId *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAlarmResponseBody) SetAlarmTaskId(v string) *ModifyAlarmResponseBody {
	s.AlarmTaskId = &v
	return s
}

func (s *ModifyAlarmResponseBody) SetRequestId(v string) *ModifyAlarmResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAlarmResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAlarmResponse) GoString() string {
	return s.String()
}

func (s *ModifyAlarmResponse) SetHeaders(v map[string]*string) *ModifyAlarmResponse {
	s.Headers = v
	return s
}

func (s *ModifyAlarmResponse) SetStatusCode(v int32) *ModifyAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAlarmResponse) SetBody(v *ModifyAlarmResponseBody) *ModifyAlarmResponse {
	s.Body = v
	return s
}

type ModifyEciScalingConfigurationRequest struct {
	AcrRegistryInfos              []*ModifyEciScalingConfigurationRequestAcrRegistryInfos         `json:"AcrRegistryInfos,omitempty" xml:"AcrRegistryInfos,omitempty" type:"Repeated"`
	ActiveDeadlineSeconds         *int64                                                          `json:"ActiveDeadlineSeconds,omitempty" xml:"ActiveDeadlineSeconds,omitempty"`
	AutoCreateEip                 *bool                                                           `json:"AutoCreateEip,omitempty" xml:"AutoCreateEip,omitempty"`
	AutoMatchImageCache           *bool                                                           `json:"AutoMatchImageCache,omitempty" xml:"AutoMatchImageCache,omitempty"`
	ContainerGroupName            *string                                                         `json:"ContainerGroupName,omitempty" xml:"ContainerGroupName,omitempty"`
	Containers                    []*ModifyEciScalingConfigurationRequestContainers               `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	CostOptimization              *bool                                                           `json:"CostOptimization,omitempty" xml:"CostOptimization,omitempty"`
	Cpu                           *float32                                                        `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CpuOptionsCore                *int32                                                          `json:"CpuOptionsCore,omitempty" xml:"CpuOptionsCore,omitempty"`
	CpuOptionsThreadsPerCore      *int32                                                          `json:"CpuOptionsThreadsPerCore,omitempty" xml:"CpuOptionsThreadsPerCore,omitempty"`
	Description                   *string                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	DnsConfigNameServers          []*string                                                       `json:"DnsConfigNameServers,omitempty" xml:"DnsConfigNameServers,omitempty" type:"Repeated"`
	DnsConfigOptions              []*ModifyEciScalingConfigurationRequestDnsConfigOptions         `json:"DnsConfigOptions,omitempty" xml:"DnsConfigOptions,omitempty" type:"Repeated"`
	DnsConfigSearchs              []*string                                                       `json:"DnsConfigSearchs,omitempty" xml:"DnsConfigSearchs,omitempty" type:"Repeated"`
	DnsPolicy                     *string                                                         `json:"DnsPolicy,omitempty" xml:"DnsPolicy,omitempty"`
	EgressBandwidth               *int64                                                          `json:"EgressBandwidth,omitempty" xml:"EgressBandwidth,omitempty"`
	EipBandwidth                  *int32                                                          `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	EnableSls                     *bool                                                           `json:"EnableSls,omitempty" xml:"EnableSls,omitempty"`
	EphemeralStorage              *int32                                                          `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
	HostAliases                   []*ModifyEciScalingConfigurationRequestHostAliases              `json:"HostAliases,omitempty" xml:"HostAliases,omitempty" type:"Repeated"`
	HostName                      *string                                                         `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ImageRegistryCredentials      []*ModifyEciScalingConfigurationRequestImageRegistryCredentials `json:"ImageRegistryCredentials,omitempty" xml:"ImageRegistryCredentials,omitempty" type:"Repeated"`
	ImageSnapshotId               *string                                                         `json:"ImageSnapshotId,omitempty" xml:"ImageSnapshotId,omitempty"`
	IngressBandwidth              *int64                                                          `json:"IngressBandwidth,omitempty" xml:"IngressBandwidth,omitempty"`
	InitContainers                []*ModifyEciScalingConfigurationRequestInitContainers           `json:"InitContainers,omitempty" xml:"InitContainers,omitempty" type:"Repeated"`
	InstanceFamilyLevel           *string                                                         `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	Ipv6AddressCount              *int32                                                          `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	LoadBalancerWeight            *int32                                                          `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	Memory                        *float32                                                        `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NtpServers                    []*string                                                       `json:"NtpServers,omitempty" xml:"NtpServers,omitempty" type:"Repeated"`
	OwnerId                       *int64                                                          `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RamRoleName                   *string                                                         `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	ResourceGroupId               *string                                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount          *string                                                         `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	RestartPolicy                 *string                                                         `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	ScalingConfigurationId        *string                                                         `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	ScalingConfigurationName      *string                                                         `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	SecurityContextSysCtls        []*ModifyEciScalingConfigurationRequestSecurityContextSysCtls   `json:"SecurityContextSysCtls,omitempty" xml:"SecurityContextSysCtls,omitempty" type:"Repeated"`
	SecurityGroupId               *string                                                         `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SpotPriceLimit                *float32                                                        `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	SpotStrategy                  *string                                                         `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	Tags                          []*ModifyEciScalingConfigurationRequestTags                     `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TerminationGracePeriodSeconds *int64                                                          `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	Volumes                       []*ModifyEciScalingConfigurationRequestVolumes                  `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Repeated"`
}

func (s ModifyEciScalingConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequest) SetAcrRegistryInfos(v []*ModifyEciScalingConfigurationRequestAcrRegistryInfos) *ModifyEciScalingConfigurationRequest {
	s.AcrRegistryInfos = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetActiveDeadlineSeconds(v int64) *ModifyEciScalingConfigurationRequest {
	s.ActiveDeadlineSeconds = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetAutoCreateEip(v bool) *ModifyEciScalingConfigurationRequest {
	s.AutoCreateEip = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetAutoMatchImageCache(v bool) *ModifyEciScalingConfigurationRequest {
	s.AutoMatchImageCache = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetContainerGroupName(v string) *ModifyEciScalingConfigurationRequest {
	s.ContainerGroupName = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetContainers(v []*ModifyEciScalingConfigurationRequestContainers) *ModifyEciScalingConfigurationRequest {
	s.Containers = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetCostOptimization(v bool) *ModifyEciScalingConfigurationRequest {
	s.CostOptimization = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetCpu(v float32) *ModifyEciScalingConfigurationRequest {
	s.Cpu = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetCpuOptionsCore(v int32) *ModifyEciScalingConfigurationRequest {
	s.CpuOptionsCore = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetCpuOptionsThreadsPerCore(v int32) *ModifyEciScalingConfigurationRequest {
	s.CpuOptionsThreadsPerCore = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetDescription(v string) *ModifyEciScalingConfigurationRequest {
	s.Description = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetDnsConfigNameServers(v []*string) *ModifyEciScalingConfigurationRequest {
	s.DnsConfigNameServers = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetDnsConfigOptions(v []*ModifyEciScalingConfigurationRequestDnsConfigOptions) *ModifyEciScalingConfigurationRequest {
	s.DnsConfigOptions = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetDnsConfigSearchs(v []*string) *ModifyEciScalingConfigurationRequest {
	s.DnsConfigSearchs = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetDnsPolicy(v string) *ModifyEciScalingConfigurationRequest {
	s.DnsPolicy = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetEgressBandwidth(v int64) *ModifyEciScalingConfigurationRequest {
	s.EgressBandwidth = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetEipBandwidth(v int32) *ModifyEciScalingConfigurationRequest {
	s.EipBandwidth = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetEnableSls(v bool) *ModifyEciScalingConfigurationRequest {
	s.EnableSls = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetEphemeralStorage(v int32) *ModifyEciScalingConfigurationRequest {
	s.EphemeralStorage = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetHostAliases(v []*ModifyEciScalingConfigurationRequestHostAliases) *ModifyEciScalingConfigurationRequest {
	s.HostAliases = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetHostName(v string) *ModifyEciScalingConfigurationRequest {
	s.HostName = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetImageRegistryCredentials(v []*ModifyEciScalingConfigurationRequestImageRegistryCredentials) *ModifyEciScalingConfigurationRequest {
	s.ImageRegistryCredentials = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetImageSnapshotId(v string) *ModifyEciScalingConfigurationRequest {
	s.ImageSnapshotId = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetIngressBandwidth(v int64) *ModifyEciScalingConfigurationRequest {
	s.IngressBandwidth = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetInitContainers(v []*ModifyEciScalingConfigurationRequestInitContainers) *ModifyEciScalingConfigurationRequest {
	s.InitContainers = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetInstanceFamilyLevel(v string) *ModifyEciScalingConfigurationRequest {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetIpv6AddressCount(v int32) *ModifyEciScalingConfigurationRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetLoadBalancerWeight(v int32) *ModifyEciScalingConfigurationRequest {
	s.LoadBalancerWeight = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetMemory(v float32) *ModifyEciScalingConfigurationRequest {
	s.Memory = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetNtpServers(v []*string) *ModifyEciScalingConfigurationRequest {
	s.NtpServers = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetOwnerId(v int64) *ModifyEciScalingConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetRamRoleName(v string) *ModifyEciScalingConfigurationRequest {
	s.RamRoleName = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetResourceGroupId(v string) *ModifyEciScalingConfigurationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetResourceOwnerAccount(v string) *ModifyEciScalingConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetRestartPolicy(v string) *ModifyEciScalingConfigurationRequest {
	s.RestartPolicy = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetScalingConfigurationId(v string) *ModifyEciScalingConfigurationRequest {
	s.ScalingConfigurationId = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetScalingConfigurationName(v string) *ModifyEciScalingConfigurationRequest {
	s.ScalingConfigurationName = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetSecurityContextSysCtls(v []*ModifyEciScalingConfigurationRequestSecurityContextSysCtls) *ModifyEciScalingConfigurationRequest {
	s.SecurityContextSysCtls = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetSecurityGroupId(v string) *ModifyEciScalingConfigurationRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetSpotPriceLimit(v float32) *ModifyEciScalingConfigurationRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetSpotStrategy(v string) *ModifyEciScalingConfigurationRequest {
	s.SpotStrategy = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetTags(v []*ModifyEciScalingConfigurationRequestTags) *ModifyEciScalingConfigurationRequest {
	s.Tags = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetTerminationGracePeriodSeconds(v int64) *ModifyEciScalingConfigurationRequest {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetVolumes(v []*ModifyEciScalingConfigurationRequestVolumes) *ModifyEciScalingConfigurationRequest {
	s.Volumes = v
	return s
}

type ModifyEciScalingConfigurationRequestAcrRegistryInfos struct {
	Domains      []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestAcrRegistryInfos) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestAcrRegistryInfos) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestAcrRegistryInfos) SetDomains(v []*string) *ModifyEciScalingConfigurationRequestAcrRegistryInfos {
	s.Domains = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestAcrRegistryInfos) SetInstanceId(v string) *ModifyEciScalingConfigurationRequestAcrRegistryInfos {
	s.InstanceId = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestAcrRegistryInfos) SetInstanceName(v string) *ModifyEciScalingConfigurationRequestAcrRegistryInfos {
	s.InstanceName = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestAcrRegistryInfos) SetRegionId(v string) *ModifyEciScalingConfigurationRequestAcrRegistryInfos {
	s.RegionId = &v
	return s
}

type ModifyEciScalingConfigurationRequestContainers struct {
	LivenessProbe   *ModifyEciScalingConfigurationRequestContainersLivenessProbe     `json:"LivenessProbe,omitempty" xml:"LivenessProbe,omitempty" require:"true" type:"Struct"`
	ReadinessProbe  *ModifyEciScalingConfigurationRequestContainersReadinessProbe    `json:"ReadinessProbe,omitempty" xml:"ReadinessProbe,omitempty" require:"true" type:"Struct"`
	SecurityContext *ModifyEciScalingConfigurationRequestContainersSecurityContext   `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" require:"true" type:"Struct"`
	Args            []*string                                                        `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	Commands        []*string                                                        `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	Cpu             *float32                                                         `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	EnvironmentVars []*ModifyEciScalingConfigurationRequestContainersEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	Gpu             *int32                                                           `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Image           *string                                                          `json:"Image,omitempty" xml:"Image,omitempty"`
	ImagePullPolicy *string                                                          `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	Memory          *float32                                                         `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Name            *string                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	Ports           []*ModifyEciScalingConfigurationRequestContainersPorts           `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	Stdin           *bool                                                            `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	StdinOnce       *bool                                                            `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	Tty             *bool                                                            `json:"Tty,omitempty" xml:"Tty,omitempty"`
	VolumeMounts    []*ModifyEciScalingConfigurationRequestContainersVolumeMounts    `json:"VolumeMounts,omitempty" xml:"VolumeMounts,omitempty" type:"Repeated"`
	WorkingDir      *string                                                          `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainers) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainers) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetLivenessProbe(v *ModifyEciScalingConfigurationRequestContainersLivenessProbe) *ModifyEciScalingConfigurationRequestContainers {
	s.LivenessProbe = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetReadinessProbe(v *ModifyEciScalingConfigurationRequestContainersReadinessProbe) *ModifyEciScalingConfigurationRequestContainers {
	s.ReadinessProbe = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetSecurityContext(v *ModifyEciScalingConfigurationRequestContainersSecurityContext) *ModifyEciScalingConfigurationRequestContainers {
	s.SecurityContext = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetArgs(v []*string) *ModifyEciScalingConfigurationRequestContainers {
	s.Args = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetCommands(v []*string) *ModifyEciScalingConfigurationRequestContainers {
	s.Commands = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetCpu(v float32) *ModifyEciScalingConfigurationRequestContainers {
	s.Cpu = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetEnvironmentVars(v []*ModifyEciScalingConfigurationRequestContainersEnvironmentVars) *ModifyEciScalingConfigurationRequestContainers {
	s.EnvironmentVars = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetGpu(v int32) *ModifyEciScalingConfigurationRequestContainers {
	s.Gpu = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetImage(v string) *ModifyEciScalingConfigurationRequestContainers {
	s.Image = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetImagePullPolicy(v string) *ModifyEciScalingConfigurationRequestContainers {
	s.ImagePullPolicy = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetMemory(v float32) *ModifyEciScalingConfigurationRequestContainers {
	s.Memory = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetName(v string) *ModifyEciScalingConfigurationRequestContainers {
	s.Name = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetPorts(v []*ModifyEciScalingConfigurationRequestContainersPorts) *ModifyEciScalingConfigurationRequestContainers {
	s.Ports = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetStdin(v bool) *ModifyEciScalingConfigurationRequestContainers {
	s.Stdin = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetStdinOnce(v bool) *ModifyEciScalingConfigurationRequestContainers {
	s.StdinOnce = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetTty(v bool) *ModifyEciScalingConfigurationRequestContainers {
	s.Tty = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetVolumeMounts(v []*ModifyEciScalingConfigurationRequestContainersVolumeMounts) *ModifyEciScalingConfigurationRequestContainers {
	s.VolumeMounts = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetWorkingDir(v string) *ModifyEciScalingConfigurationRequestContainers {
	s.WorkingDir = &v
	return s
}

type ModifyEciScalingConfigurationRequestContainersLivenessProbe struct {
	Exec                *ModifyEciScalingConfigurationRequestContainersLivenessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" require:"true" type:"Struct"`
	FailureThreshold    *int32                                                                `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" require:"true" type:"Struct"`
	InitialDelaySeconds *int32                                                                `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                                `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                                `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *ModifyEciScalingConfigurationRequestContainersLivenessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" require:"true" type:"Struct"`
	TimeoutSeconds      *int32                                                                `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersLivenessProbe) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersLivenessProbe) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) SetExec(v *ModifyEciScalingConfigurationRequestContainersLivenessProbeExec) *ModifyEciScalingConfigurationRequestContainersLivenessProbe {
	s.Exec = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) SetFailureThreshold(v int32) *ModifyEciScalingConfigurationRequestContainersLivenessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) SetHttpGet(v *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet) *ModifyEciScalingConfigurationRequestContainersLivenessProbe {
	s.HttpGet = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) SetInitialDelaySeconds(v int32) *ModifyEciScalingConfigurationRequestContainersLivenessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) SetPeriodSeconds(v int32) *ModifyEciScalingConfigurationRequestContainersLivenessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) SetSuccessThreshold(v int32) *ModifyEciScalingConfigurationRequestContainersLivenessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) SetTcpSocket(v *ModifyEciScalingConfigurationRequestContainersLivenessProbeTcpSocket) *ModifyEciScalingConfigurationRequestContainersLivenessProbe {
	s.TcpSocket = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) SetTimeoutSeconds(v int32) *ModifyEciScalingConfigurationRequestContainersLivenessProbe {
	s.TimeoutSeconds = &v
	return s
}

type ModifyEciScalingConfigurationRequestContainersLivenessProbeExec struct {
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
}

func (s ModifyEciScalingConfigurationRequestContainersLivenessProbeExec) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersLivenessProbeExec) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbeExec) SetCommands(v []*string) *ModifyEciScalingConfigurationRequestContainersLivenessProbeExec {
	s.Commands = v
	return s
}

type ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet struct {
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet) SetPath(v string) *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet) SetPort(v int32) *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet) SetScheme(v string) *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet {
	s.Scheme = &v
	return s
}

type ModifyEciScalingConfigurationRequestContainersLivenessProbeTcpSocket struct {
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersLivenessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersLivenessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbeTcpSocket) SetPort(v int32) *ModifyEciScalingConfigurationRequestContainersLivenessProbeTcpSocket {
	s.Port = &v
	return s
}

type ModifyEciScalingConfigurationRequestContainersReadinessProbe struct {
	Exec                *ModifyEciScalingConfigurationRequestContainersReadinessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" require:"true" type:"Struct"`
	FailureThreshold    *int32                                                                 `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" require:"true" type:"Struct"`
	InitialDelaySeconds *int32                                                                 `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                                 `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                                 `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *ModifyEciScalingConfigurationRequestContainersReadinessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" require:"true" type:"Struct"`
	TimeoutSeconds      *int32                                                                 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersReadinessProbe) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersReadinessProbe) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) SetExec(v *ModifyEciScalingConfigurationRequestContainersReadinessProbeExec) *ModifyEciScalingConfigurationRequestContainersReadinessProbe {
	s.Exec = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) SetFailureThreshold(v int32) *ModifyEciScalingConfigurationRequestContainersReadinessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) SetHttpGet(v *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet) *ModifyEciScalingConfigurationRequestContainersReadinessProbe {
	s.HttpGet = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) SetInitialDelaySeconds(v int32) *ModifyEciScalingConfigurationRequestContainersReadinessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) SetPeriodSeconds(v int32) *ModifyEciScalingConfigurationRequestContainersReadinessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) SetSuccessThreshold(v int32) *ModifyEciScalingConfigurationRequestContainersReadinessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) SetTcpSocket(v *ModifyEciScalingConfigurationRequestContainersReadinessProbeTcpSocket) *ModifyEciScalingConfigurationRequestContainersReadinessProbe {
	s.TcpSocket = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) SetTimeoutSeconds(v int32) *ModifyEciScalingConfigurationRequestContainersReadinessProbe {
	s.TimeoutSeconds = &v
	return s
}

type ModifyEciScalingConfigurationRequestContainersReadinessProbeExec struct {
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
}

func (s ModifyEciScalingConfigurationRequestContainersReadinessProbeExec) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersReadinessProbeExec) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbeExec) SetCommands(v []*string) *ModifyEciScalingConfigurationRequestContainersReadinessProbeExec {
	s.Commands = v
	return s
}

type ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet struct {
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet) SetPath(v string) *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet) SetPort(v int32) *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet) SetScheme(v string) *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet {
	s.Scheme = &v
	return s
}

type ModifyEciScalingConfigurationRequestContainersReadinessProbeTcpSocket struct {
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersReadinessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersReadinessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbeTcpSocket) SetPort(v int32) *ModifyEciScalingConfigurationRequestContainersReadinessProbeTcpSocket {
	s.Port = &v
	return s
}

type ModifyEciScalingConfigurationRequestContainersSecurityContext struct {
	Capability             *ModifyEciScalingConfigurationRequestContainersSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" require:"true" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                                    `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                                   `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersSecurityContext) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersSecurityContext) SetCapability(v *ModifyEciScalingConfigurationRequestContainersSecurityContextCapability) *ModifyEciScalingConfigurationRequestContainersSecurityContext {
	s.Capability = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersSecurityContext) SetReadOnlyRootFilesystem(v bool) *ModifyEciScalingConfigurationRequestContainersSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersSecurityContext) SetRunAsUser(v int64) *ModifyEciScalingConfigurationRequestContainersSecurityContext {
	s.RunAsUser = &v
	return s
}

type ModifyEciScalingConfigurationRequestContainersSecurityContextCapability struct {
	Adds []*string `json:"Adds,omitempty" xml:"Adds,omitempty" type:"Repeated"`
}

func (s ModifyEciScalingConfigurationRequestContainersSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersSecurityContextCapability) SetAdds(v []*string) *ModifyEciScalingConfigurationRequestContainersSecurityContextCapability {
	s.Adds = v
	return s
}

type ModifyEciScalingConfigurationRequestContainersEnvironmentVars struct {
	FieldRef *ModifyEciScalingConfigurationRequestContainersEnvironmentVarsFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" require:"true" type:"Struct"`
	Key      *string                                                                `json:"Key,omitempty" xml:"Key,omitempty"`
	Value    *string                                                                `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersEnvironmentVars) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersEnvironmentVars) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersEnvironmentVars) SetFieldRef(v *ModifyEciScalingConfigurationRequestContainersEnvironmentVarsFieldRef) *ModifyEciScalingConfigurationRequestContainersEnvironmentVars {
	s.FieldRef = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersEnvironmentVars) SetKey(v string) *ModifyEciScalingConfigurationRequestContainersEnvironmentVars {
	s.Key = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersEnvironmentVars) SetValue(v string) *ModifyEciScalingConfigurationRequestContainersEnvironmentVars {
	s.Value = &v
	return s
}

type ModifyEciScalingConfigurationRequestContainersEnvironmentVarsFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersEnvironmentVarsFieldRef) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersEnvironmentVarsFieldRef) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersEnvironmentVarsFieldRef) SetFieldPath(v string) *ModifyEciScalingConfigurationRequestContainersEnvironmentVarsFieldRef {
	s.FieldPath = &v
	return s
}

type ModifyEciScalingConfigurationRequestContainersPorts struct {
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersPorts) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersPorts) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersPorts) SetPort(v int32) *ModifyEciScalingConfigurationRequestContainersPorts {
	s.Port = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersPorts) SetProtocol(v string) *ModifyEciScalingConfigurationRequestContainersPorts {
	s.Protocol = &v
	return s
}

type ModifyEciScalingConfigurationRequestContainersVolumeMounts struct {
	MountPath        *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ReadOnly         *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	SubPath          *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersVolumeMounts) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersVolumeMounts) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersVolumeMounts) SetMountPath(v string) *ModifyEciScalingConfigurationRequestContainersVolumeMounts {
	s.MountPath = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersVolumeMounts) SetMountPropagation(v string) *ModifyEciScalingConfigurationRequestContainersVolumeMounts {
	s.MountPropagation = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersVolumeMounts) SetName(v string) *ModifyEciScalingConfigurationRequestContainersVolumeMounts {
	s.Name = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersVolumeMounts) SetReadOnly(v bool) *ModifyEciScalingConfigurationRequestContainersVolumeMounts {
	s.ReadOnly = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersVolumeMounts) SetSubPath(v string) *ModifyEciScalingConfigurationRequestContainersVolumeMounts {
	s.SubPath = &v
	return s
}

type ModifyEciScalingConfigurationRequestDnsConfigOptions struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestDnsConfigOptions) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestDnsConfigOptions) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestDnsConfigOptions) SetName(v string) *ModifyEciScalingConfigurationRequestDnsConfigOptions {
	s.Name = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestDnsConfigOptions) SetValue(v string) *ModifyEciScalingConfigurationRequestDnsConfigOptions {
	s.Value = &v
	return s
}

type ModifyEciScalingConfigurationRequestHostAliases struct {
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	Ip        *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestHostAliases) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestHostAliases) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestHostAliases) SetHostnames(v []*string) *ModifyEciScalingConfigurationRequestHostAliases {
	s.Hostnames = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestHostAliases) SetIp(v string) *ModifyEciScalingConfigurationRequestHostAliases {
	s.Ip = &v
	return s
}

type ModifyEciScalingConfigurationRequestImageRegistryCredentials struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestImageRegistryCredentials) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestImageRegistryCredentials) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestImageRegistryCredentials) SetPassword(v string) *ModifyEciScalingConfigurationRequestImageRegistryCredentials {
	s.Password = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestImageRegistryCredentials) SetServer(v string) *ModifyEciScalingConfigurationRequestImageRegistryCredentials {
	s.Server = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestImageRegistryCredentials) SetUserName(v string) *ModifyEciScalingConfigurationRequestImageRegistryCredentials {
	s.UserName = &v
	return s
}

type ModifyEciScalingConfigurationRequestInitContainers struct {
	SecurityContext              *ModifyEciScalingConfigurationRequestInitContainersSecurityContext                `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" require:"true" type:"Struct"`
	Args                         []*string                                                                         `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	Commands                     []*string                                                                         `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	Cpu                          *float32                                                                          `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu                          *int32                                                                            `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Image                        *string                                                                           `json:"Image,omitempty" xml:"Image,omitempty"`
	ImagePullPolicy              *string                                                                           `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	InitContainerEnvironmentVars []*ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars `json:"InitContainerEnvironmentVars,omitempty" xml:"InitContainerEnvironmentVars,omitempty" type:"Repeated"`
	InitContainerPorts           []*ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts           `json:"InitContainerPorts,omitempty" xml:"InitContainerPorts,omitempty" type:"Repeated"`
	InitContainerVolumeMounts    []*ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts    `json:"InitContainerVolumeMounts,omitempty" xml:"InitContainerVolumeMounts,omitempty" type:"Repeated"`
	Memory                       *float32                                                                          `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Name                         *string                                                                           `json:"Name,omitempty" xml:"Name,omitempty"`
	WorkingDir                   *string                                                                           `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestInitContainers) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestInitContainers) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetSecurityContext(v *ModifyEciScalingConfigurationRequestInitContainersSecurityContext) *ModifyEciScalingConfigurationRequestInitContainers {
	s.SecurityContext = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetArgs(v []*string) *ModifyEciScalingConfigurationRequestInitContainers {
	s.Args = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetCommands(v []*string) *ModifyEciScalingConfigurationRequestInitContainers {
	s.Commands = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetCpu(v float32) *ModifyEciScalingConfigurationRequestInitContainers {
	s.Cpu = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetGpu(v int32) *ModifyEciScalingConfigurationRequestInitContainers {
	s.Gpu = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetImage(v string) *ModifyEciScalingConfigurationRequestInitContainers {
	s.Image = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetImagePullPolicy(v string) *ModifyEciScalingConfigurationRequestInitContainers {
	s.ImagePullPolicy = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetInitContainerEnvironmentVars(v []*ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) *ModifyEciScalingConfigurationRequestInitContainers {
	s.InitContainerEnvironmentVars = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetInitContainerPorts(v []*ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts) *ModifyEciScalingConfigurationRequestInitContainers {
	s.InitContainerPorts = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetInitContainerVolumeMounts(v []*ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) *ModifyEciScalingConfigurationRequestInitContainers {
	s.InitContainerVolumeMounts = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetMemory(v float32) *ModifyEciScalingConfigurationRequestInitContainers {
	s.Memory = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetName(v string) *ModifyEciScalingConfigurationRequestInitContainers {
	s.Name = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetWorkingDir(v string) *ModifyEciScalingConfigurationRequestInitContainers {
	s.WorkingDir = &v
	return s
}

type ModifyEciScalingConfigurationRequestInitContainersSecurityContext struct {
	Capability             *ModifyEciScalingConfigurationRequestInitContainersSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" require:"true" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                                        `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                                       `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestInitContainersSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestInitContainersSecurityContext) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestInitContainersSecurityContext) SetCapability(v *ModifyEciScalingConfigurationRequestInitContainersSecurityContextCapability) *ModifyEciScalingConfigurationRequestInitContainersSecurityContext {
	s.Capability = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersSecurityContext) SetReadOnlyRootFilesystem(v bool) *ModifyEciScalingConfigurationRequestInitContainersSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersSecurityContext) SetRunAsUser(v int64) *ModifyEciScalingConfigurationRequestInitContainersSecurityContext {
	s.RunAsUser = &v
	return s
}

type ModifyEciScalingConfigurationRequestInitContainersSecurityContextCapability struct {
	Adds []*string `json:"Adds,omitempty" xml:"Adds,omitempty" type:"Repeated"`
}

func (s ModifyEciScalingConfigurationRequestInitContainersSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestInitContainersSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestInitContainersSecurityContextCapability) SetAdds(v []*string) *ModifyEciScalingConfigurationRequestInitContainersSecurityContextCapability {
	s.Adds = v
	return s
}

type ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars struct {
	FieldRef *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVarsFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" require:"true" type:"Struct"`
	Key      *string                                                                                 `json:"Key,omitempty" xml:"Key,omitempty"`
	Value    *string                                                                                 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) SetFieldRef(v *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVarsFieldRef) *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars {
	s.FieldRef = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) SetKey(v string) *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars {
	s.Key = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) SetValue(v string) *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars {
	s.Value = &v
	return s
}

type ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVarsFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVarsFieldRef) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVarsFieldRef) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVarsFieldRef) SetFieldPath(v string) *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVarsFieldRef {
	s.FieldPath = &v
	return s
}

type ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts struct {
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts) SetPort(v int32) *ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts {
	s.Port = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts) SetProtocol(v string) *ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts {
	s.Protocol = &v
	return s
}

type ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts struct {
	MountPath        *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ReadOnly         *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	SubPath          *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) SetMountPath(v string) *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	s.MountPath = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) SetMountPropagation(v string) *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	s.MountPropagation = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) SetName(v string) *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	s.Name = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) SetReadOnly(v bool) *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	s.ReadOnly = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) SetSubPath(v string) *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	s.SubPath = &v
	return s
}

type ModifyEciScalingConfigurationRequestSecurityContextSysCtls struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestSecurityContextSysCtls) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestSecurityContextSysCtls) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestSecurityContextSysCtls) SetName(v string) *ModifyEciScalingConfigurationRequestSecurityContextSysCtls {
	s.Name = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestSecurityContextSysCtls) SetValue(v string) *ModifyEciScalingConfigurationRequestSecurityContextSysCtls {
	s.Value = &v
	return s
}

type ModifyEciScalingConfigurationRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestTags) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestTags) SetKey(v string) *ModifyEciScalingConfigurationRequestTags {
	s.Key = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestTags) SetValue(v string) *ModifyEciScalingConfigurationRequestTags {
	s.Value = &v
	return s
}

type ModifyEciScalingConfigurationRequestVolumes struct {
	DiskVolume                       *ModifyEciScalingConfigurationRequestVolumesDiskVolume                         `json:"DiskVolume,omitempty" xml:"DiskVolume,omitempty" require:"true" type:"Struct"`
	EmptyDirVolume                   *ModifyEciScalingConfigurationRequestVolumesEmptyDirVolume                     `json:"EmptyDirVolume,omitempty" xml:"EmptyDirVolume,omitempty" require:"true" type:"Struct"`
	FlexVolume                       *ModifyEciScalingConfigurationRequestVolumesFlexVolume                         `json:"FlexVolume,omitempty" xml:"FlexVolume,omitempty" require:"true" type:"Struct"`
	HostPathVolume                   *ModifyEciScalingConfigurationRequestVolumesHostPathVolume                     `json:"HostPathVolume,omitempty" xml:"HostPathVolume,omitempty" require:"true" type:"Struct"`
	NFSVolume                        *ModifyEciScalingConfigurationRequestVolumesNFSVolume                          `json:"NFSVolume,omitempty" xml:"NFSVolume,omitempty" require:"true" type:"Struct"`
	ConfigFileVolumeConfigFileToPath []*ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath `json:"ConfigFileVolumeConfigFileToPath,omitempty" xml:"ConfigFileVolumeConfigFileToPath,omitempty" type:"Repeated"`
	ConfigFileVolumeDefaultMode      *int32                                                                         `json:"ConfigFileVolumeDefaultMode,omitempty" xml:"ConfigFileVolumeDefaultMode,omitempty"`
	Name                             *string                                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	Type                             *string                                                                        `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestVolumes) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestVolumes) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestVolumes) SetDiskVolume(v *ModifyEciScalingConfigurationRequestVolumesDiskVolume) *ModifyEciScalingConfigurationRequestVolumes {
	s.DiskVolume = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumes) SetEmptyDirVolume(v *ModifyEciScalingConfigurationRequestVolumesEmptyDirVolume) *ModifyEciScalingConfigurationRequestVolumes {
	s.EmptyDirVolume = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumes) SetFlexVolume(v *ModifyEciScalingConfigurationRequestVolumesFlexVolume) *ModifyEciScalingConfigurationRequestVolumes {
	s.FlexVolume = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumes) SetHostPathVolume(v *ModifyEciScalingConfigurationRequestVolumesHostPathVolume) *ModifyEciScalingConfigurationRequestVolumes {
	s.HostPathVolume = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumes) SetNFSVolume(v *ModifyEciScalingConfigurationRequestVolumesNFSVolume) *ModifyEciScalingConfigurationRequestVolumes {
	s.NFSVolume = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumes) SetConfigFileVolumeConfigFileToPath(v []*ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath) *ModifyEciScalingConfigurationRequestVolumes {
	s.ConfigFileVolumeConfigFileToPath = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumes) SetConfigFileVolumeDefaultMode(v int32) *ModifyEciScalingConfigurationRequestVolumes {
	s.ConfigFileVolumeDefaultMode = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumes) SetName(v string) *ModifyEciScalingConfigurationRequestVolumes {
	s.Name = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumes) SetType(v string) *ModifyEciScalingConfigurationRequestVolumes {
	s.Type = &v
	return s
}

type ModifyEciScalingConfigurationRequestVolumesDiskVolume struct {
	DiskId   *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	DiskSize *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	FsType   *string `json:"FsType,omitempty" xml:"FsType,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestVolumesDiskVolume) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestVolumesDiskVolume) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestVolumesDiskVolume) SetDiskId(v string) *ModifyEciScalingConfigurationRequestVolumesDiskVolume {
	s.DiskId = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesDiskVolume) SetDiskSize(v int32) *ModifyEciScalingConfigurationRequestVolumesDiskVolume {
	s.DiskSize = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesDiskVolume) SetFsType(v string) *ModifyEciScalingConfigurationRequestVolumesDiskVolume {
	s.FsType = &v
	return s
}

type ModifyEciScalingConfigurationRequestVolumesEmptyDirVolume struct {
	Medium *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestVolumesEmptyDirVolume) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestVolumesEmptyDirVolume) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestVolumesEmptyDirVolume) SetMedium(v string) *ModifyEciScalingConfigurationRequestVolumesEmptyDirVolume {
	s.Medium = &v
	return s
}

type ModifyEciScalingConfigurationRequestVolumesFlexVolume struct {
	Driver  *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
	FsType  *string `json:"FsType,omitempty" xml:"FsType,omitempty"`
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestVolumesFlexVolume) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestVolumesFlexVolume) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestVolumesFlexVolume) SetDriver(v string) *ModifyEciScalingConfigurationRequestVolumesFlexVolume {
	s.Driver = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesFlexVolume) SetFsType(v string) *ModifyEciScalingConfigurationRequestVolumesFlexVolume {
	s.FsType = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesFlexVolume) SetOptions(v string) *ModifyEciScalingConfigurationRequestVolumesFlexVolume {
	s.Options = &v
	return s
}

type ModifyEciScalingConfigurationRequestVolumesHostPathVolume struct {
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestVolumesHostPathVolume) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestVolumesHostPathVolume) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestVolumesHostPathVolume) SetPath(v string) *ModifyEciScalingConfigurationRequestVolumesHostPathVolume {
	s.Path = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesHostPathVolume) SetType(v string) *ModifyEciScalingConfigurationRequestVolumesHostPathVolume {
	s.Type = &v
	return s
}

type ModifyEciScalingConfigurationRequestVolumesNFSVolume struct {
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty"`
	ReadOnly *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestVolumesNFSVolume) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestVolumesNFSVolume) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestVolumesNFSVolume) SetPath(v string) *ModifyEciScalingConfigurationRequestVolumesNFSVolume {
	s.Path = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesNFSVolume) SetReadOnly(v bool) *ModifyEciScalingConfigurationRequestVolumesNFSVolume {
	s.ReadOnly = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesNFSVolume) SetServer(v string) *ModifyEciScalingConfigurationRequestVolumesNFSVolume {
	s.Server = &v
	return s
}

type ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Mode    *int32  `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Path    *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath) SetContent(v string) *ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath {
	s.Content = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath) SetMode(v int32) *ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath {
	s.Mode = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath) SetPath(v string) *ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath {
	s.Path = &v
	return s
}

type ModifyEciScalingConfigurationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEciScalingConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationResponseBody) SetRequestId(v string) *ModifyEciScalingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type ModifyEciScalingConfigurationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyEciScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyEciScalingConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyEciScalingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationResponse) SetHeaders(v map[string]*string) *ModifyEciScalingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ModifyEciScalingConfigurationResponse) SetStatusCode(v int32) *ModifyEciScalingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEciScalingConfigurationResponse) SetBody(v *ModifyEciScalingConfigurationResponseBody) *ModifyEciScalingConfigurationResponse {
	s.Body = v
	return s
}

type ModifyLifecycleHookRequest struct {
	DefaultResult        *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	HeartbeatTimeout     *int32  `json:"HeartbeatTimeout,omitempty" xml:"HeartbeatTimeout,omitempty"`
	LifecycleHookId      *string `json:"LifecycleHookId,omitempty" xml:"LifecycleHookId,omitempty"`
	LifecycleHookName    *string `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	LifecycleTransition  *string `json:"LifecycleTransition,omitempty" xml:"LifecycleTransition,omitempty"`
	NotificationArn      *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	NotificationMetadata *string `json:"NotificationMetadata,omitempty" xml:"NotificationMetadata,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s ModifyLifecycleHookRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLifecycleHookRequest) GoString() string {
	return s.String()
}

func (s *ModifyLifecycleHookRequest) SetDefaultResult(v string) *ModifyLifecycleHookRequest {
	s.DefaultResult = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetHeartbeatTimeout(v int32) *ModifyLifecycleHookRequest {
	s.HeartbeatTimeout = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetLifecycleHookId(v string) *ModifyLifecycleHookRequest {
	s.LifecycleHookId = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetLifecycleHookName(v string) *ModifyLifecycleHookRequest {
	s.LifecycleHookName = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetLifecycleTransition(v string) *ModifyLifecycleHookRequest {
	s.LifecycleTransition = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetNotificationArn(v string) *ModifyLifecycleHookRequest {
	s.NotificationArn = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetNotificationMetadata(v string) *ModifyLifecycleHookRequest {
	s.NotificationMetadata = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetOwnerAccount(v string) *ModifyLifecycleHookRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetOwnerId(v int64) *ModifyLifecycleHookRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetRegionId(v string) *ModifyLifecycleHookRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetResourceOwnerAccount(v string) *ModifyLifecycleHookRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetScalingGroupId(v string) *ModifyLifecycleHookRequest {
	s.ScalingGroupId = &v
	return s
}

type ModifyLifecycleHookResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLifecycleHookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLifecycleHookResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLifecycleHookResponseBody) SetRequestId(v string) *ModifyLifecycleHookResponseBody {
	s.RequestId = &v
	return s
}

type ModifyLifecycleHookResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyLifecycleHookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyLifecycleHookResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLifecycleHookResponse) GoString() string {
	return s.String()
}

func (s *ModifyLifecycleHookResponse) SetHeaders(v map[string]*string) *ModifyLifecycleHookResponse {
	s.Headers = v
	return s
}

func (s *ModifyLifecycleHookResponse) SetStatusCode(v int32) *ModifyLifecycleHookResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLifecycleHookResponse) SetBody(v *ModifyLifecycleHookResponseBody) *ModifyLifecycleHookResponse {
	s.Body = v
	return s
}

type ModifyNotificationConfigurationRequest struct {
	NotificationArn      *string   `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	NotificationTypes    []*string `json:"NotificationTypes,omitempty" xml:"NotificationTypes,omitempty" type:"Repeated"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s ModifyNotificationConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNotificationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ModifyNotificationConfigurationRequest) SetNotificationArn(v string) *ModifyNotificationConfigurationRequest {
	s.NotificationArn = &v
	return s
}

func (s *ModifyNotificationConfigurationRequest) SetNotificationTypes(v []*string) *ModifyNotificationConfigurationRequest {
	s.NotificationTypes = v
	return s
}

func (s *ModifyNotificationConfigurationRequest) SetOwnerId(v int64) *ModifyNotificationConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyNotificationConfigurationRequest) SetRegionId(v string) *ModifyNotificationConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNotificationConfigurationRequest) SetResourceOwnerAccount(v string) *ModifyNotificationConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyNotificationConfigurationRequest) SetScalingGroupId(v string) *ModifyNotificationConfigurationRequest {
	s.ScalingGroupId = &v
	return s
}

type ModifyNotificationConfigurationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNotificationConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNotificationConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNotificationConfigurationResponseBody) SetRequestId(v string) *ModifyNotificationConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type ModifyNotificationConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyNotificationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyNotificationConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNotificationConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ModifyNotificationConfigurationResponse) SetHeaders(v map[string]*string) *ModifyNotificationConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ModifyNotificationConfigurationResponse) SetStatusCode(v int32) *ModifyNotificationConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNotificationConfigurationResponse) SetBody(v *ModifyNotificationConfigurationResponseBody) *ModifyNotificationConfigurationResponse {
	s.Body = v
	return s
}

type ModifyScalingConfigurationRequest struct {
	PrivatePoolOptions       *ModifyScalingConfigurationRequestPrivatePoolOptions      `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	SystemDisk               *ModifyScalingConfigurationRequestSystemDisk              `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	Affinity                 *string                                                   `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	Cpu                      *int32                                                    `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreditSpecification      *string                                                   `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	DataDisks                []*ModifyScalingConfigurationRequestDataDisks             `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	DedicatedHostId          *string                                                   `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	DeploymentSetId          *string                                                   `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	HostName                 *string                                                   `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HpcClusterId             *string                                                   `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	ImageFamily              *string                                                   `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	ImageId                  *string                                                   `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName                *string                                                   `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	InstanceDescription      *string                                                   `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	InstanceName             *string                                                   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstancePatternInfos     []*ModifyScalingConfigurationRequestInstancePatternInfos  `json:"InstancePatternInfos,omitempty" xml:"InstancePatternInfos,omitempty" type:"Repeated"`
	InstanceTypeOverrides    []*ModifyScalingConfigurationRequestInstanceTypeOverrides `json:"InstanceTypeOverrides,omitempty" xml:"InstanceTypeOverrides,omitempty" type:"Repeated"`
	InstanceTypes            []*string                                                 `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	InternetChargeType       *string                                                   `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandwidthOut  *int32                                                    `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	IoOptimized              *string                                                   `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	Ipv6AddressCount         *int32                                                    `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	KeyPairName              *string                                                   `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	LoadBalancerWeight       *int32                                                    `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	Memory                   *int32                                                    `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Override                 *bool                                                     `json:"Override,omitempty" xml:"Override,omitempty"`
	OwnerAccount             *string                                                   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64                                                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PasswordInherit          *bool                                                     `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	RamRoleName              *string                                                   `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	ResourceGroupId          *string                                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount     *string                                                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingConfigurationId   *string                                                   `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	ScalingConfigurationName *string                                                   `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	SchedulerOptions         map[string]interface{}                                    `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty"`
	SecurityGroupId          *string                                                   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupIds         []*string                                                 `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	SpotDuration             *int32                                                    `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	SpotInterruptionBehavior *string                                                   `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	SpotPriceLimits          []*ModifyScalingConfigurationRequestSpotPriceLimits       `json:"SpotPriceLimits,omitempty" xml:"SpotPriceLimits,omitempty" type:"Repeated"`
	SpotStrategy             *string                                                   `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDiskCategories     []*string                                                 `json:"SystemDiskCategories,omitempty" xml:"SystemDiskCategories,omitempty" type:"Repeated"`
	Tags                     *string                                                   `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Tenancy                  *string                                                   `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	UserData                 *string                                                   `json:"UserData,omitempty" xml:"UserData,omitempty"`
	ZoneId                   *string                                                   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyScalingConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequest) SetPrivatePoolOptions(v *ModifyScalingConfigurationRequestPrivatePoolOptions) *ModifyScalingConfigurationRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSystemDisk(v *ModifyScalingConfigurationRequestSystemDisk) *ModifyScalingConfigurationRequest {
	s.SystemDisk = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetAffinity(v string) *ModifyScalingConfigurationRequest {
	s.Affinity = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetCpu(v int32) *ModifyScalingConfigurationRequest {
	s.Cpu = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetCreditSpecification(v string) *ModifyScalingConfigurationRequest {
	s.CreditSpecification = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetDataDisks(v []*ModifyScalingConfigurationRequestDataDisks) *ModifyScalingConfigurationRequest {
	s.DataDisks = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetDedicatedHostId(v string) *ModifyScalingConfigurationRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetDeploymentSetId(v string) *ModifyScalingConfigurationRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetHostName(v string) *ModifyScalingConfigurationRequest {
	s.HostName = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetHpcClusterId(v string) *ModifyScalingConfigurationRequest {
	s.HpcClusterId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetImageFamily(v string) *ModifyScalingConfigurationRequest {
	s.ImageFamily = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetImageId(v string) *ModifyScalingConfigurationRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetImageName(v string) *ModifyScalingConfigurationRequest {
	s.ImageName = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetInstanceDescription(v string) *ModifyScalingConfigurationRequest {
	s.InstanceDescription = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetInstanceName(v string) *ModifyScalingConfigurationRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetInstancePatternInfos(v []*ModifyScalingConfigurationRequestInstancePatternInfos) *ModifyScalingConfigurationRequest {
	s.InstancePatternInfos = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetInstanceTypeOverrides(v []*ModifyScalingConfigurationRequestInstanceTypeOverrides) *ModifyScalingConfigurationRequest {
	s.InstanceTypeOverrides = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetInstanceTypes(v []*string) *ModifyScalingConfigurationRequest {
	s.InstanceTypes = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetInternetChargeType(v string) *ModifyScalingConfigurationRequest {
	s.InternetChargeType = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetInternetMaxBandwidthOut(v int32) *ModifyScalingConfigurationRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetIoOptimized(v string) *ModifyScalingConfigurationRequest {
	s.IoOptimized = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetIpv6AddressCount(v int32) *ModifyScalingConfigurationRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetKeyPairName(v string) *ModifyScalingConfigurationRequest {
	s.KeyPairName = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetLoadBalancerWeight(v int32) *ModifyScalingConfigurationRequest {
	s.LoadBalancerWeight = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetMemory(v int32) *ModifyScalingConfigurationRequest {
	s.Memory = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetOverride(v bool) *ModifyScalingConfigurationRequest {
	s.Override = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetOwnerAccount(v string) *ModifyScalingConfigurationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetOwnerId(v int64) *ModifyScalingConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetPasswordInherit(v bool) *ModifyScalingConfigurationRequest {
	s.PasswordInherit = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetRamRoleName(v string) *ModifyScalingConfigurationRequest {
	s.RamRoleName = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetResourceGroupId(v string) *ModifyScalingConfigurationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetResourceOwnerAccount(v string) *ModifyScalingConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetScalingConfigurationId(v string) *ModifyScalingConfigurationRequest {
	s.ScalingConfigurationId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetScalingConfigurationName(v string) *ModifyScalingConfigurationRequest {
	s.ScalingConfigurationName = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSchedulerOptions(v map[string]interface{}) *ModifyScalingConfigurationRequest {
	s.SchedulerOptions = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSecurityGroupId(v string) *ModifyScalingConfigurationRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSecurityGroupIds(v []*string) *ModifyScalingConfigurationRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSpotDuration(v int32) *ModifyScalingConfigurationRequest {
	s.SpotDuration = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSpotInterruptionBehavior(v string) *ModifyScalingConfigurationRequest {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSpotPriceLimits(v []*ModifyScalingConfigurationRequestSpotPriceLimits) *ModifyScalingConfigurationRequest {
	s.SpotPriceLimits = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSpotStrategy(v string) *ModifyScalingConfigurationRequest {
	s.SpotStrategy = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSystemDiskCategories(v []*string) *ModifyScalingConfigurationRequest {
	s.SystemDiskCategories = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetTags(v string) *ModifyScalingConfigurationRequest {
	s.Tags = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetTenancy(v string) *ModifyScalingConfigurationRequest {
	s.Tenancy = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetUserData(v string) *ModifyScalingConfigurationRequest {
	s.UserData = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetZoneId(v string) *ModifyScalingConfigurationRequest {
	s.ZoneId = &v
	return s
}

type ModifyScalingConfigurationRequestPrivatePoolOptions struct {
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
}

func (s ModifyScalingConfigurationRequestPrivatePoolOptions) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestPrivatePoolOptions) SetId(v string) *ModifyScalingConfigurationRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *ModifyScalingConfigurationRequestPrivatePoolOptions) SetMatchCriteria(v string) *ModifyScalingConfigurationRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

type ModifyScalingConfigurationRequestSystemDisk struct {
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	BurstingEnabled      *bool   `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DiskName             *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	EncryptAlgorithm     *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	Encrypted            *bool   `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	KMSKeyId             *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	PerformanceLevel     *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	ProvisionedIops      *int64  `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	Size                 *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ModifyScalingConfigurationRequestSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetAutoSnapshotPolicyId(v string) *ModifyScalingConfigurationRequestSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetBurstingEnabled(v bool) *ModifyScalingConfigurationRequestSystemDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetCategory(v string) *ModifyScalingConfigurationRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetDescription(v string) *ModifyScalingConfigurationRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetDiskName(v string) *ModifyScalingConfigurationRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetEncryptAlgorithm(v string) *ModifyScalingConfigurationRequestSystemDisk {
	s.EncryptAlgorithm = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetEncrypted(v bool) *ModifyScalingConfigurationRequestSystemDisk {
	s.Encrypted = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetKMSKeyId(v string) *ModifyScalingConfigurationRequestSystemDisk {
	s.KMSKeyId = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetPerformanceLevel(v string) *ModifyScalingConfigurationRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetProvisionedIops(v int64) *ModifyScalingConfigurationRequestSystemDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetSize(v int32) *ModifyScalingConfigurationRequestSystemDisk {
	s.Size = &v
	return s
}

type ModifyScalingConfigurationRequestDataDisks struct {
	AutoSnapshotPolicyId *string   `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	BurstingEnabled      *bool     `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	Categories           []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	Category             *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteWithInstance   *bool     `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	Description          *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Device               *string   `json:"Device,omitempty" xml:"Device,omitempty"`
	DiskName             *string   `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	Encrypted            *string   `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	KMSKeyId             *string   `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	PerformanceLevel     *string   `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	ProvisionedIops      *int64    `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	Size                 *int32    `json:"Size,omitempty" xml:"Size,omitempty"`
	SnapshotId           *string   `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ModifyScalingConfigurationRequestDataDisks) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationRequestDataDisks) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetAutoSnapshotPolicyId(v string) *ModifyScalingConfigurationRequestDataDisks {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetBurstingEnabled(v bool) *ModifyScalingConfigurationRequestDataDisks {
	s.BurstingEnabled = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetCategories(v []*string) *ModifyScalingConfigurationRequestDataDisks {
	s.Categories = v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetCategory(v string) *ModifyScalingConfigurationRequestDataDisks {
	s.Category = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetDeleteWithInstance(v bool) *ModifyScalingConfigurationRequestDataDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetDescription(v string) *ModifyScalingConfigurationRequestDataDisks {
	s.Description = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetDevice(v string) *ModifyScalingConfigurationRequestDataDisks {
	s.Device = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetDiskName(v string) *ModifyScalingConfigurationRequestDataDisks {
	s.DiskName = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetEncrypted(v string) *ModifyScalingConfigurationRequestDataDisks {
	s.Encrypted = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetKMSKeyId(v string) *ModifyScalingConfigurationRequestDataDisks {
	s.KMSKeyId = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetPerformanceLevel(v string) *ModifyScalingConfigurationRequestDataDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetProvisionedIops(v int64) *ModifyScalingConfigurationRequestDataDisks {
	s.ProvisionedIops = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetSize(v int32) *ModifyScalingConfigurationRequestDataDisks {
	s.Size = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetSnapshotId(v string) *ModifyScalingConfigurationRequestDataDisks {
	s.SnapshotId = &v
	return s
}

type ModifyScalingConfigurationRequestInstancePatternInfos struct {
	Architectures         []*string `json:"Architectures,omitempty" xml:"Architectures,omitempty" type:"Repeated"`
	BurstablePerformance  *string   `json:"BurstablePerformance,omitempty" xml:"BurstablePerformance,omitempty"`
	Cores                 *int32    `json:"Cores,omitempty" xml:"Cores,omitempty"`
	ExcludedInstanceTypes []*string `json:"ExcludedInstanceTypes,omitempty" xml:"ExcludedInstanceTypes,omitempty" type:"Repeated"`
	InstanceFamilyLevel   *string   `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	MaxPrice              *float32  `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	Memory                *float32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ModifyScalingConfigurationRequestInstancePatternInfos) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationRequestInstancePatternInfos) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetArchitectures(v []*string) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.Architectures = v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetBurstablePerformance(v string) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.BurstablePerformance = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetCores(v int32) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.Cores = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetExcludedInstanceTypes(v []*string) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.ExcludedInstanceTypes = v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetInstanceFamilyLevel(v string) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetMaxPrice(v float32) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.MaxPrice = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetMemory(v float32) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.Memory = &v
	return s
}

type ModifyScalingConfigurationRequestInstanceTypeOverrides struct {
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	WeightedCapacity *int32  `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s ModifyScalingConfigurationRequestInstanceTypeOverrides) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationRequestInstanceTypeOverrides) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestInstanceTypeOverrides) SetInstanceType(v string) *ModifyScalingConfigurationRequestInstanceTypeOverrides {
	s.InstanceType = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstanceTypeOverrides) SetWeightedCapacity(v int32) *ModifyScalingConfigurationRequestInstanceTypeOverrides {
	s.WeightedCapacity = &v
	return s
}

type ModifyScalingConfigurationRequestSpotPriceLimits struct {
	InstanceType *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	PriceLimit   *float32 `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
}

func (s ModifyScalingConfigurationRequestSpotPriceLimits) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationRequestSpotPriceLimits) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestSpotPriceLimits) SetInstanceType(v string) *ModifyScalingConfigurationRequestSpotPriceLimits {
	s.InstanceType = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSpotPriceLimits) SetPriceLimit(v float32) *ModifyScalingConfigurationRequestSpotPriceLimits {
	s.PriceLimit = &v
	return s
}

type ModifyScalingConfigurationShrinkRequest struct {
	PrivatePoolOptions       *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions      `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	SystemDisk               *ModifyScalingConfigurationShrinkRequestSystemDisk              `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	Affinity                 *string                                                         `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	Cpu                      *int32                                                          `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreditSpecification      *string                                                         `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	DataDisks                []*ModifyScalingConfigurationShrinkRequestDataDisks             `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	DedicatedHostId          *string                                                         `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	DeploymentSetId          *string                                                         `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	HostName                 *string                                                         `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HpcClusterId             *string                                                         `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	ImageFamily              *string                                                         `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	ImageId                  *string                                                         `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName                *string                                                         `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	InstanceDescription      *string                                                         `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	InstanceName             *string                                                         `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstancePatternInfos     []*ModifyScalingConfigurationShrinkRequestInstancePatternInfos  `json:"InstancePatternInfos,omitempty" xml:"InstancePatternInfos,omitempty" type:"Repeated"`
	InstanceTypeOverrides    []*ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides `json:"InstanceTypeOverrides,omitempty" xml:"InstanceTypeOverrides,omitempty" type:"Repeated"`
	InstanceTypes            []*string                                                       `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	InternetChargeType       *string                                                         `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandwidthOut  *int32                                                          `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	IoOptimized              *string                                                         `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	Ipv6AddressCount         *int32                                                          `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	KeyPairName              *string                                                         `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	LoadBalancerWeight       *int32                                                          `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	Memory                   *int32                                                          `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Override                 *bool                                                           `json:"Override,omitempty" xml:"Override,omitempty"`
	OwnerAccount             *string                                                         `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64                                                          `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PasswordInherit          *bool                                                           `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	RamRoleName              *string                                                         `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	ResourceGroupId          *string                                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount     *string                                                         `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingConfigurationId   *string                                                         `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	ScalingConfigurationName *string                                                         `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	SchedulerOptionsShrink   *string                                                         `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty"`
	SecurityGroupId          *string                                                         `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupIds         []*string                                                       `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	SpotDuration             *int32                                                          `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	SpotInterruptionBehavior *string                                                         `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	SpotPriceLimits          []*ModifyScalingConfigurationShrinkRequestSpotPriceLimits       `json:"SpotPriceLimits,omitempty" xml:"SpotPriceLimits,omitempty" type:"Repeated"`
	SpotStrategy             *string                                                         `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDiskCategories     []*string                                                       `json:"SystemDiskCategories,omitempty" xml:"SystemDiskCategories,omitempty" type:"Repeated"`
	Tags                     *string                                                         `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Tenancy                  *string                                                         `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	UserData                 *string                                                         `json:"UserData,omitempty" xml:"UserData,omitempty"`
	ZoneId                   *string                                                         `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequest) SetPrivatePoolOptions(v *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions) *ModifyScalingConfigurationShrinkRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSystemDisk(v *ModifyScalingConfigurationShrinkRequestSystemDisk) *ModifyScalingConfigurationShrinkRequest {
	s.SystemDisk = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetAffinity(v string) *ModifyScalingConfigurationShrinkRequest {
	s.Affinity = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetCpu(v int32) *ModifyScalingConfigurationShrinkRequest {
	s.Cpu = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetCreditSpecification(v string) *ModifyScalingConfigurationShrinkRequest {
	s.CreditSpecification = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetDataDisks(v []*ModifyScalingConfigurationShrinkRequestDataDisks) *ModifyScalingConfigurationShrinkRequest {
	s.DataDisks = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetDedicatedHostId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetDeploymentSetId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetHostName(v string) *ModifyScalingConfigurationShrinkRequest {
	s.HostName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetHpcClusterId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.HpcClusterId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetImageFamily(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ImageFamily = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetImageId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetImageName(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ImageName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetInstanceDescription(v string) *ModifyScalingConfigurationShrinkRequest {
	s.InstanceDescription = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetInstanceName(v string) *ModifyScalingConfigurationShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetInstancePatternInfos(v []*ModifyScalingConfigurationShrinkRequestInstancePatternInfos) *ModifyScalingConfigurationShrinkRequest {
	s.InstancePatternInfos = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetInstanceTypeOverrides(v []*ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides) *ModifyScalingConfigurationShrinkRequest {
	s.InstanceTypeOverrides = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetInstanceTypes(v []*string) *ModifyScalingConfigurationShrinkRequest {
	s.InstanceTypes = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetInternetChargeType(v string) *ModifyScalingConfigurationShrinkRequest {
	s.InternetChargeType = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetInternetMaxBandwidthOut(v int32) *ModifyScalingConfigurationShrinkRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetIoOptimized(v string) *ModifyScalingConfigurationShrinkRequest {
	s.IoOptimized = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetIpv6AddressCount(v int32) *ModifyScalingConfigurationShrinkRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetKeyPairName(v string) *ModifyScalingConfigurationShrinkRequest {
	s.KeyPairName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetLoadBalancerWeight(v int32) *ModifyScalingConfigurationShrinkRequest {
	s.LoadBalancerWeight = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetMemory(v int32) *ModifyScalingConfigurationShrinkRequest {
	s.Memory = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetOverride(v bool) *ModifyScalingConfigurationShrinkRequest {
	s.Override = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetOwnerAccount(v string) *ModifyScalingConfigurationShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetOwnerId(v int64) *ModifyScalingConfigurationShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetPasswordInherit(v bool) *ModifyScalingConfigurationShrinkRequest {
	s.PasswordInherit = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetRamRoleName(v string) *ModifyScalingConfigurationShrinkRequest {
	s.RamRoleName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetResourceGroupId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetResourceOwnerAccount(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetScalingConfigurationId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ScalingConfigurationId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetScalingConfigurationName(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ScalingConfigurationName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSchedulerOptionsShrink(v string) *ModifyScalingConfigurationShrinkRequest {
	s.SchedulerOptionsShrink = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSecurityGroupId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSecurityGroupIds(v []*string) *ModifyScalingConfigurationShrinkRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSpotDuration(v int32) *ModifyScalingConfigurationShrinkRequest {
	s.SpotDuration = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSpotInterruptionBehavior(v string) *ModifyScalingConfigurationShrinkRequest {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSpotPriceLimits(v []*ModifyScalingConfigurationShrinkRequestSpotPriceLimits) *ModifyScalingConfigurationShrinkRequest {
	s.SpotPriceLimits = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSpotStrategy(v string) *ModifyScalingConfigurationShrinkRequest {
	s.SpotStrategy = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSystemDiskCategories(v []*string) *ModifyScalingConfigurationShrinkRequest {
	s.SystemDiskCategories = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetTags(v string) *ModifyScalingConfigurationShrinkRequest {
	s.Tags = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetTenancy(v string) *ModifyScalingConfigurationShrinkRequest {
	s.Tenancy = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetUserData(v string) *ModifyScalingConfigurationShrinkRequest {
	s.UserData = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetZoneId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ZoneId = &v
	return s
}

type ModifyScalingConfigurationShrinkRequestPrivatePoolOptions struct {
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestPrivatePoolOptions) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions) SetId(v string) *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions) SetMatchCriteria(v string) *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

type ModifyScalingConfigurationShrinkRequestSystemDisk struct {
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	BurstingEnabled      *bool   `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DiskName             *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	EncryptAlgorithm     *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	Encrypted            *bool   `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	KMSKeyId             *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	PerformanceLevel     *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	ProvisionedIops      *int64  `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	Size                 *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetAutoSnapshotPolicyId(v string) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetBurstingEnabled(v bool) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetCategory(v string) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetDescription(v string) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetDiskName(v string) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetEncryptAlgorithm(v string) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.EncryptAlgorithm = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetEncrypted(v bool) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.Encrypted = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetKMSKeyId(v string) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.KMSKeyId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetPerformanceLevel(v string) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetProvisionedIops(v int64) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetSize(v int32) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.Size = &v
	return s
}

type ModifyScalingConfigurationShrinkRequestDataDisks struct {
	AutoSnapshotPolicyId *string   `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	BurstingEnabled      *bool     `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	Categories           []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	Category             *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteWithInstance   *bool     `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	Description          *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Device               *string   `json:"Device,omitempty" xml:"Device,omitempty"`
	DiskName             *string   `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	Encrypted            *string   `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	KMSKeyId             *string   `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	PerformanceLevel     *string   `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	ProvisionedIops      *int64    `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	Size                 *int32    `json:"Size,omitempty" xml:"Size,omitempty"`
	SnapshotId           *string   `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestDataDisks) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestDataDisks) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetAutoSnapshotPolicyId(v string) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetBurstingEnabled(v bool) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.BurstingEnabled = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetCategories(v []*string) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.Categories = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetCategory(v string) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.Category = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetDeleteWithInstance(v bool) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetDescription(v string) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.Description = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetDevice(v string) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.Device = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetDiskName(v string) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.DiskName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetEncrypted(v string) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.Encrypted = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetKMSKeyId(v string) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.KMSKeyId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetPerformanceLevel(v string) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetProvisionedIops(v int64) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.ProvisionedIops = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetSize(v int32) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.Size = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetSnapshotId(v string) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.SnapshotId = &v
	return s
}

type ModifyScalingConfigurationShrinkRequestInstancePatternInfos struct {
	Architectures         []*string `json:"Architectures,omitempty" xml:"Architectures,omitempty" type:"Repeated"`
	BurstablePerformance  *string   `json:"BurstablePerformance,omitempty" xml:"BurstablePerformance,omitempty"`
	Cores                 *int32    `json:"Cores,omitempty" xml:"Cores,omitempty"`
	ExcludedInstanceTypes []*string `json:"ExcludedInstanceTypes,omitempty" xml:"ExcludedInstanceTypes,omitempty" type:"Repeated"`
	InstanceFamilyLevel   *string   `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	MaxPrice              *float32  `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	Memory                *float32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestInstancePatternInfos) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetArchitectures(v []*string) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.Architectures = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetBurstablePerformance(v string) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.BurstablePerformance = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetCores(v int32) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.Cores = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetExcludedInstanceTypes(v []*string) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.ExcludedInstanceTypes = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetInstanceFamilyLevel(v string) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetMaxPrice(v float32) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MaxPrice = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetMemory(v float32) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.Memory = &v
	return s
}

type ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides struct {
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	WeightedCapacity *int32  `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides) SetInstanceType(v string) *ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides {
	s.InstanceType = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides) SetWeightedCapacity(v int32) *ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides {
	s.WeightedCapacity = &v
	return s
}

type ModifyScalingConfigurationShrinkRequestSpotPriceLimits struct {
	InstanceType *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	PriceLimit   *float32 `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestSpotPriceLimits) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestSpotPriceLimits) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestSpotPriceLimits) SetInstanceType(v string) *ModifyScalingConfigurationShrinkRequestSpotPriceLimits {
	s.InstanceType = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSpotPriceLimits) SetPriceLimit(v float32) *ModifyScalingConfigurationShrinkRequestSpotPriceLimits {
	s.PriceLimit = &v
	return s
}

type ModifyScalingConfigurationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyScalingConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationResponseBody) SetRequestId(v string) *ModifyScalingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type ModifyScalingConfigurationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyScalingConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationResponse) SetHeaders(v map[string]*string) *ModifyScalingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ModifyScalingConfigurationResponse) SetStatusCode(v int32) *ModifyScalingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyScalingConfigurationResponse) SetBody(v *ModifyScalingConfigurationResponseBody) *ModifyScalingConfigurationResponse {
	s.Body = v
	return s
}

type ModifyScalingGroupRequest struct {
	ActiveScalingConfigurationId        *string                                             `json:"ActiveScalingConfigurationId,omitempty" xml:"ActiveScalingConfigurationId,omitempty"`
	AllocationStrategy                  *string                                             `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	AzBalance                           *bool                                               `json:"AzBalance,omitempty" xml:"AzBalance,omitempty"`
	CompensateWithOnDemand              *bool                                               `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	DefaultCooldown                     *int32                                              `json:"DefaultCooldown,omitempty" xml:"DefaultCooldown,omitempty"`
	DesiredCapacity                     *int32                                              `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	GroupDeletionProtection             *bool                                               `json:"GroupDeletionProtection,omitempty" xml:"GroupDeletionProtection,omitempty"`
	HealthCheckType                     *string                                             `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	LaunchTemplateId                    *string                                             `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	LaunchTemplateOverrides             []*ModifyScalingGroupRequestLaunchTemplateOverrides `json:"LaunchTemplateOverrides,omitempty" xml:"LaunchTemplateOverrides,omitempty" type:"Repeated"`
	LaunchTemplateVersion               *string                                             `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	MaxInstanceLifetime                 *int32                                              `json:"MaxInstanceLifetime,omitempty" xml:"MaxInstanceLifetime,omitempty"`
	MaxSize                             *int32                                              `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	MinSize                             *int32                                              `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	MultiAZPolicy                       *string                                             `json:"MultiAZPolicy,omitempty" xml:"MultiAZPolicy,omitempty"`
	OnDemandBaseCapacity                *int32                                              `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	OnDemandPercentageAboveBaseCapacity *int32                                              `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	OwnerAccount                        *string                                             `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                             *int64                                              `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RemovalPolicies                     []*string                                           `json:"RemovalPolicies,omitempty" xml:"RemovalPolicies,omitempty" type:"Repeated"`
	ResourceOwnerAccount                *string                                             `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                     *int64                                              `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupId                      *string                                             `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScalingGroupName                    *string                                             `json:"ScalingGroupName,omitempty" xml:"ScalingGroupName,omitempty"`
	SpotAllocationStrategy              *string                                             `json:"SpotAllocationStrategy,omitempty" xml:"SpotAllocationStrategy,omitempty"`
	SpotInstancePools                   *int32                                              `json:"SpotInstancePools,omitempty" xml:"SpotInstancePools,omitempty"`
	SpotInstanceRemedy                  *bool                                               `json:"SpotInstanceRemedy,omitempty" xml:"SpotInstanceRemedy,omitempty"`
	VSwitchIds                          []*string                                           `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s ModifyScalingGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyScalingGroupRequest) SetActiveScalingConfigurationId(v string) *ModifyScalingGroupRequest {
	s.ActiveScalingConfigurationId = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetAllocationStrategy(v string) *ModifyScalingGroupRequest {
	s.AllocationStrategy = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetAzBalance(v bool) *ModifyScalingGroupRequest {
	s.AzBalance = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetCompensateWithOnDemand(v bool) *ModifyScalingGroupRequest {
	s.CompensateWithOnDemand = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetDefaultCooldown(v int32) *ModifyScalingGroupRequest {
	s.DefaultCooldown = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetDesiredCapacity(v int32) *ModifyScalingGroupRequest {
	s.DesiredCapacity = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetGroupDeletionProtection(v bool) *ModifyScalingGroupRequest {
	s.GroupDeletionProtection = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetHealthCheckType(v string) *ModifyScalingGroupRequest {
	s.HealthCheckType = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetLaunchTemplateId(v string) *ModifyScalingGroupRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetLaunchTemplateOverrides(v []*ModifyScalingGroupRequestLaunchTemplateOverrides) *ModifyScalingGroupRequest {
	s.LaunchTemplateOverrides = v
	return s
}

func (s *ModifyScalingGroupRequest) SetLaunchTemplateVersion(v string) *ModifyScalingGroupRequest {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetMaxInstanceLifetime(v int32) *ModifyScalingGroupRequest {
	s.MaxInstanceLifetime = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetMaxSize(v int32) *ModifyScalingGroupRequest {
	s.MaxSize = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetMinSize(v int32) *ModifyScalingGroupRequest {
	s.MinSize = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetMultiAZPolicy(v string) *ModifyScalingGroupRequest {
	s.MultiAZPolicy = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetOnDemandBaseCapacity(v int32) *ModifyScalingGroupRequest {
	s.OnDemandBaseCapacity = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetOnDemandPercentageAboveBaseCapacity(v int32) *ModifyScalingGroupRequest {
	s.OnDemandPercentageAboveBaseCapacity = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetOwnerAccount(v string) *ModifyScalingGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetOwnerId(v int64) *ModifyScalingGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetRemovalPolicies(v []*string) *ModifyScalingGroupRequest {
	s.RemovalPolicies = v
	return s
}

func (s *ModifyScalingGroupRequest) SetResourceOwnerAccount(v string) *ModifyScalingGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetResourceOwnerId(v int64) *ModifyScalingGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetScalingGroupId(v string) *ModifyScalingGroupRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetScalingGroupName(v string) *ModifyScalingGroupRequest {
	s.ScalingGroupName = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetSpotAllocationStrategy(v string) *ModifyScalingGroupRequest {
	s.SpotAllocationStrategy = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetSpotInstancePools(v int32) *ModifyScalingGroupRequest {
	s.SpotInstancePools = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetSpotInstanceRemedy(v bool) *ModifyScalingGroupRequest {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetVSwitchIds(v []*string) *ModifyScalingGroupRequest {
	s.VSwitchIds = v
	return s
}

type ModifyScalingGroupRequestLaunchTemplateOverrides struct {
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	WeightedCapacity *int32  `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s ModifyScalingGroupRequestLaunchTemplateOverrides) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingGroupRequestLaunchTemplateOverrides) GoString() string {
	return s.String()
}

func (s *ModifyScalingGroupRequestLaunchTemplateOverrides) SetInstanceType(v string) *ModifyScalingGroupRequestLaunchTemplateOverrides {
	s.InstanceType = &v
	return s
}

func (s *ModifyScalingGroupRequestLaunchTemplateOverrides) SetWeightedCapacity(v int32) *ModifyScalingGroupRequestLaunchTemplateOverrides {
	s.WeightedCapacity = &v
	return s
}

type ModifyScalingGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyScalingGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScalingGroupResponseBody) SetRequestId(v string) *ModifyScalingGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyScalingGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyScalingGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyScalingGroupResponse) SetHeaders(v map[string]*string) *ModifyScalingGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyScalingGroupResponse) SetStatusCode(v int32) *ModifyScalingGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyScalingGroupResponse) SetBody(v *ModifyScalingGroupResponseBody) *ModifyScalingGroupResponse {
	s.Body = v
	return s
}

type ModifyScalingRuleRequest struct {
	AdjustmentType           *string                                    `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	AdjustmentValue          *int32                                     `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	Cooldown                 *int32                                     `json:"Cooldown,omitempty" xml:"Cooldown,omitempty"`
	DisableScaleIn           *bool                                      `json:"DisableScaleIn,omitempty" xml:"DisableScaleIn,omitempty"`
	EstimatedInstanceWarmup  *int32                                     `json:"EstimatedInstanceWarmup,omitempty" xml:"EstimatedInstanceWarmup,omitempty"`
	InitialMaxSize           *int32                                     `json:"InitialMaxSize,omitempty" xml:"InitialMaxSize,omitempty"`
	MetricName               *string                                    `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	MinAdjustmentMagnitude   *int32                                     `json:"MinAdjustmentMagnitude,omitempty" xml:"MinAdjustmentMagnitude,omitempty"`
	OwnerAccount             *string                                    `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64                                     `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PredictiveScalingMode    *string                                    `json:"PredictiveScalingMode,omitempty" xml:"PredictiveScalingMode,omitempty"`
	PredictiveTaskBufferTime *int32                                     `json:"PredictiveTaskBufferTime,omitempty" xml:"PredictiveTaskBufferTime,omitempty"`
	PredictiveValueBehavior  *string                                    `json:"PredictiveValueBehavior,omitempty" xml:"PredictiveValueBehavior,omitempty"`
	PredictiveValueBuffer    *int32                                     `json:"PredictiveValueBuffer,omitempty" xml:"PredictiveValueBuffer,omitempty"`
	ResourceOwnerAccount     *string                                    `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64                                     `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScaleInEvaluationCount   *int32                                     `json:"ScaleInEvaluationCount,omitempty" xml:"ScaleInEvaluationCount,omitempty"`
	ScaleOutEvaluationCount  *int32                                     `json:"ScaleOutEvaluationCount,omitempty" xml:"ScaleOutEvaluationCount,omitempty"`
	ScalingRuleId            *string                                    `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty"`
	ScalingRuleName          *string                                    `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	StepAdjustments          []*ModifyScalingRuleRequestStepAdjustments `json:"StepAdjustments,omitempty" xml:"StepAdjustments,omitempty" type:"Repeated"`
	TargetValue              *float32                                   `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s ModifyScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleRequest) SetAdjustmentType(v string) *ModifyScalingRuleRequest {
	s.AdjustmentType = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetAdjustmentValue(v int32) *ModifyScalingRuleRequest {
	s.AdjustmentValue = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetCooldown(v int32) *ModifyScalingRuleRequest {
	s.Cooldown = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetDisableScaleIn(v bool) *ModifyScalingRuleRequest {
	s.DisableScaleIn = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetEstimatedInstanceWarmup(v int32) *ModifyScalingRuleRequest {
	s.EstimatedInstanceWarmup = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetInitialMaxSize(v int32) *ModifyScalingRuleRequest {
	s.InitialMaxSize = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetMetricName(v string) *ModifyScalingRuleRequest {
	s.MetricName = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetMinAdjustmentMagnitude(v int32) *ModifyScalingRuleRequest {
	s.MinAdjustmentMagnitude = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetOwnerAccount(v string) *ModifyScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetOwnerId(v int64) *ModifyScalingRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetPredictiveScalingMode(v string) *ModifyScalingRuleRequest {
	s.PredictiveScalingMode = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetPredictiveTaskBufferTime(v int32) *ModifyScalingRuleRequest {
	s.PredictiveTaskBufferTime = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetPredictiveValueBehavior(v string) *ModifyScalingRuleRequest {
	s.PredictiveValueBehavior = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetPredictiveValueBuffer(v int32) *ModifyScalingRuleRequest {
	s.PredictiveValueBuffer = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetResourceOwnerAccount(v string) *ModifyScalingRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetResourceOwnerId(v int64) *ModifyScalingRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetScaleInEvaluationCount(v int32) *ModifyScalingRuleRequest {
	s.ScaleInEvaluationCount = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetScaleOutEvaluationCount(v int32) *ModifyScalingRuleRequest {
	s.ScaleOutEvaluationCount = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetScalingRuleId(v string) *ModifyScalingRuleRequest {
	s.ScalingRuleId = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetScalingRuleName(v string) *ModifyScalingRuleRequest {
	s.ScalingRuleName = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetStepAdjustments(v []*ModifyScalingRuleRequestStepAdjustments) *ModifyScalingRuleRequest {
	s.StepAdjustments = v
	return s
}

func (s *ModifyScalingRuleRequest) SetTargetValue(v float32) *ModifyScalingRuleRequest {
	s.TargetValue = &v
	return s
}

type ModifyScalingRuleRequestStepAdjustments struct {
	MetricIntervalLowerBound *float32 `json:"MetricIntervalLowerBound,omitempty" xml:"MetricIntervalLowerBound,omitempty"`
	MetricIntervalUpperBound *float32 `json:"MetricIntervalUpperBound,omitempty" xml:"MetricIntervalUpperBound,omitempty"`
	ScalingAdjustment        *int32   `json:"ScalingAdjustment,omitempty" xml:"ScalingAdjustment,omitempty"`
}

func (s ModifyScalingRuleRequestStepAdjustments) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingRuleRequestStepAdjustments) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleRequestStepAdjustments) SetMetricIntervalLowerBound(v float32) *ModifyScalingRuleRequestStepAdjustments {
	s.MetricIntervalLowerBound = &v
	return s
}

func (s *ModifyScalingRuleRequestStepAdjustments) SetMetricIntervalUpperBound(v float32) *ModifyScalingRuleRequestStepAdjustments {
	s.MetricIntervalUpperBound = &v
	return s
}

func (s *ModifyScalingRuleRequestStepAdjustments) SetScalingAdjustment(v int32) *ModifyScalingRuleRequestStepAdjustments {
	s.ScalingAdjustment = &v
	return s
}

type ModifyScalingRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleResponseBody) SetRequestId(v string) *ModifyScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

type ModifyScalingRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleResponse) SetHeaders(v map[string]*string) *ModifyScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyScalingRuleResponse) SetStatusCode(v int32) *ModifyScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyScalingRuleResponse) SetBody(v *ModifyScalingRuleResponseBody) *ModifyScalingRuleResponse {
	s.Body = v
	return s
}

type ModifyScheduledTaskRequest struct {
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DesiredCapacity      *int32  `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	LaunchExpirationTime *int32  `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
	LaunchTime           *string `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	MaxValue             *int32  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinValue             *int32  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RecurrenceEndTime    *string `json:"RecurrenceEndTime,omitempty" xml:"RecurrenceEndTime,omitempty"`
	RecurrenceType       *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValue      *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScheduledAction      *string `json:"ScheduledAction,omitempty" xml:"ScheduledAction,omitempty"`
	ScheduledTaskId      *string `json:"ScheduledTaskId,omitempty" xml:"ScheduledTaskId,omitempty"`
	ScheduledTaskName    *string `json:"ScheduledTaskName,omitempty" xml:"ScheduledTaskName,omitempty"`
	TaskEnabled          *bool   `json:"TaskEnabled,omitempty" xml:"TaskEnabled,omitempty"`
}

func (s ModifyScheduledTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskRequest) SetDescription(v string) *ModifyScheduledTaskRequest {
	s.Description = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetDesiredCapacity(v int32) *ModifyScheduledTaskRequest {
	s.DesiredCapacity = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetLaunchExpirationTime(v int32) *ModifyScheduledTaskRequest {
	s.LaunchExpirationTime = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetLaunchTime(v string) *ModifyScheduledTaskRequest {
	s.LaunchTime = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetMaxValue(v int32) *ModifyScheduledTaskRequest {
	s.MaxValue = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetMinValue(v int32) *ModifyScheduledTaskRequest {
	s.MinValue = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetOwnerAccount(v string) *ModifyScheduledTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetOwnerId(v int64) *ModifyScheduledTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetRecurrenceEndTime(v string) *ModifyScheduledTaskRequest {
	s.RecurrenceEndTime = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetRecurrenceType(v string) *ModifyScheduledTaskRequest {
	s.RecurrenceType = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetRecurrenceValue(v string) *ModifyScheduledTaskRequest {
	s.RecurrenceValue = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetResourceOwnerAccount(v string) *ModifyScheduledTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetResourceOwnerId(v int64) *ModifyScheduledTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetScalingGroupId(v string) *ModifyScheduledTaskRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetScheduledAction(v string) *ModifyScheduledTaskRequest {
	s.ScheduledAction = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetScheduledTaskId(v string) *ModifyScheduledTaskRequest {
	s.ScheduledTaskId = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetScheduledTaskName(v string) *ModifyScheduledTaskRequest {
	s.ScheduledTaskName = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetTaskEnabled(v bool) *ModifyScheduledTaskRequest {
	s.TaskEnabled = &v
	return s
}

type ModifyScheduledTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyScheduledTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskResponseBody) SetRequestId(v string) *ModifyScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

type ModifyScheduledTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyScheduledTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyScheduledTaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskResponse) SetHeaders(v map[string]*string) *ModifyScheduledTaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyScheduledTaskResponse) SetStatusCode(v int32) *ModifyScheduledTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyScheduledTaskResponse) SetBody(v *ModifyScheduledTaskResponseBody) *ModifyScheduledTaskResponse {
	s.Body = v
	return s
}

type RebalanceInstancesRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s RebalanceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s RebalanceInstancesRequest) GoString() string {
	return s.String()
}

func (s *RebalanceInstancesRequest) SetOwnerAccount(v string) *RebalanceInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RebalanceInstancesRequest) SetOwnerId(v int64) *RebalanceInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *RebalanceInstancesRequest) SetRegionId(v string) *RebalanceInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *RebalanceInstancesRequest) SetResourceOwnerAccount(v string) *RebalanceInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RebalanceInstancesRequest) SetResourceOwnerId(v int64) *RebalanceInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RebalanceInstancesRequest) SetScalingGroupId(v string) *RebalanceInstancesRequest {
	s.ScalingGroupId = &v
	return s
}

type RebalanceInstancesResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s RebalanceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebalanceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RebalanceInstancesResponseBody) SetRequestId(v string) *RebalanceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebalanceInstancesResponseBody) SetScalingActivityId(v string) *RebalanceInstancesResponseBody {
	s.ScalingActivityId = &v
	return s
}

type RebalanceInstancesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RebalanceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RebalanceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s RebalanceInstancesResponse) GoString() string {
	return s.String()
}

func (s *RebalanceInstancesResponse) SetHeaders(v map[string]*string) *RebalanceInstancesResponse {
	s.Headers = v
	return s
}

func (s *RebalanceInstancesResponse) SetStatusCode(v int32) *RebalanceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *RebalanceInstancesResponse) SetBody(v *RebalanceInstancesResponseBody) *RebalanceInstancesResponse {
	s.Body = v
	return s
}

type RecordLifecycleActionHeartbeatRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	HeartbeatTimeout     *int32  `json:"heartbeatTimeout,omitempty" xml:"heartbeatTimeout,omitempty"`
	LifecycleActionToken *string `json:"lifecycleActionToken,omitempty" xml:"lifecycleActionToken,omitempty"`
	LifecycleHookId      *string `json:"lifecycleHookId,omitempty" xml:"lifecycleHookId,omitempty"`
}

func (s RecordLifecycleActionHeartbeatRequest) String() string {
	return tea.Prettify(s)
}

func (s RecordLifecycleActionHeartbeatRequest) GoString() string {
	return s.String()
}

func (s *RecordLifecycleActionHeartbeatRequest) SetOwnerAccount(v string) *RecordLifecycleActionHeartbeatRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatRequest) SetOwnerId(v int64) *RecordLifecycleActionHeartbeatRequest {
	s.OwnerId = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatRequest) SetRegionId(v string) *RecordLifecycleActionHeartbeatRequest {
	s.RegionId = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatRequest) SetResourceOwnerAccount(v string) *RecordLifecycleActionHeartbeatRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatRequest) SetHeartbeatTimeout(v int32) *RecordLifecycleActionHeartbeatRequest {
	s.HeartbeatTimeout = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatRequest) SetLifecycleActionToken(v string) *RecordLifecycleActionHeartbeatRequest {
	s.LifecycleActionToken = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatRequest) SetLifecycleHookId(v string) *RecordLifecycleActionHeartbeatRequest {
	s.LifecycleHookId = &v
	return s
}

type RecordLifecycleActionHeartbeatResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecordLifecycleActionHeartbeatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecordLifecycleActionHeartbeatResponseBody) GoString() string {
	return s.String()
}

func (s *RecordLifecycleActionHeartbeatResponseBody) SetRequestId(v string) *RecordLifecycleActionHeartbeatResponseBody {
	s.RequestId = &v
	return s
}

type RecordLifecycleActionHeartbeatResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecordLifecycleActionHeartbeatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecordLifecycleActionHeartbeatResponse) String() string {
	return tea.Prettify(s)
}

func (s RecordLifecycleActionHeartbeatResponse) GoString() string {
	return s.String()
}

func (s *RecordLifecycleActionHeartbeatResponse) SetHeaders(v map[string]*string) *RecordLifecycleActionHeartbeatResponse {
	s.Headers = v
	return s
}

func (s *RecordLifecycleActionHeartbeatResponse) SetStatusCode(v int32) *RecordLifecycleActionHeartbeatResponse {
	s.StatusCode = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatResponse) SetBody(v *RecordLifecycleActionHeartbeatResponseBody) *RecordLifecycleActionHeartbeatResponse {
	s.Body = v
	return s
}

type RemoveInstancesRequest struct {
	DecreaseDesiredCapacity *bool     `json:"DecreaseDesiredCapacity,omitempty" xml:"DecreaseDesiredCapacity,omitempty"`
	InstanceIds             []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	OwnerAccount            *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RemovePolicy            *string   `json:"RemovePolicy,omitempty" xml:"RemovePolicy,omitempty"`
	ResourceOwnerAccount    *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupId          *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s RemoveInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveInstancesRequest) GoString() string {
	return s.String()
}

func (s *RemoveInstancesRequest) SetDecreaseDesiredCapacity(v bool) *RemoveInstancesRequest {
	s.DecreaseDesiredCapacity = &v
	return s
}

func (s *RemoveInstancesRequest) SetInstanceIds(v []*string) *RemoveInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *RemoveInstancesRequest) SetOwnerAccount(v string) *RemoveInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveInstancesRequest) SetOwnerId(v int64) *RemoveInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveInstancesRequest) SetRegionId(v string) *RemoveInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveInstancesRequest) SetRemovePolicy(v string) *RemoveInstancesRequest {
	s.RemovePolicy = &v
	return s
}

func (s *RemoveInstancesRequest) SetResourceOwnerAccount(v string) *RemoveInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveInstancesRequest) SetResourceOwnerId(v int64) *RemoveInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveInstancesRequest) SetScalingGroupId(v string) *RemoveInstancesRequest {
	s.ScalingGroupId = &v
	return s
}

type RemoveInstancesResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s RemoveInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveInstancesResponseBody) SetRequestId(v string) *RemoveInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveInstancesResponseBody) SetScalingActivityId(v string) *RemoveInstancesResponseBody {
	s.ScalingActivityId = &v
	return s
}

type RemoveInstancesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveInstancesResponse) GoString() string {
	return s.String()
}

func (s *RemoveInstancesResponse) SetHeaders(v map[string]*string) *RemoveInstancesResponse {
	s.Headers = v
	return s
}

func (s *RemoveInstancesResponse) SetStatusCode(v int32) *RemoveInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveInstancesResponse) SetBody(v *RemoveInstancesResponseBody) *RemoveInstancesResponse {
	s.Body = v
	return s
}

type ResumeProcessesRequest struct {
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Processes            []*string `json:"Processes,omitempty" xml:"Processes,omitempty" type:"Repeated"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s ResumeProcessesRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeProcessesRequest) GoString() string {
	return s.String()
}

func (s *ResumeProcessesRequest) SetClientToken(v string) *ResumeProcessesRequest {
	s.ClientToken = &v
	return s
}

func (s *ResumeProcessesRequest) SetOwnerId(v int64) *ResumeProcessesRequest {
	s.OwnerId = &v
	return s
}

func (s *ResumeProcessesRequest) SetProcesses(v []*string) *ResumeProcessesRequest {
	s.Processes = v
	return s
}

func (s *ResumeProcessesRequest) SetRegionId(v string) *ResumeProcessesRequest {
	s.RegionId = &v
	return s
}

func (s *ResumeProcessesRequest) SetResourceOwnerAccount(v string) *ResumeProcessesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResumeProcessesRequest) SetScalingGroupId(v string) *ResumeProcessesRequest {
	s.ScalingGroupId = &v
	return s
}

type ResumeProcessesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeProcessesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeProcessesResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeProcessesResponseBody) SetRequestId(v string) *ResumeProcessesResponseBody {
	s.RequestId = &v
	return s
}

type ResumeProcessesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResumeProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeProcessesResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeProcessesResponse) GoString() string {
	return s.String()
}

func (s *ResumeProcessesResponse) SetHeaders(v map[string]*string) *ResumeProcessesResponse {
	s.Headers = v
	return s
}

func (s *ResumeProcessesResponse) SetStatusCode(v int32) *ResumeProcessesResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeProcessesResponse) SetBody(v *ResumeProcessesResponseBody) *ResumeProcessesResponse {
	s.Body = v
	return s
}

type ScaleWithAdjustmentRequest struct {
	AdjustmentType         *string `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	AdjustmentValue        *int32  `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	ClientToken            *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	MinAdjustmentMagnitude *int32  `json:"MinAdjustmentMagnitude,omitempty" xml:"MinAdjustmentMagnitude,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId         *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	SyncActivity           *bool   `json:"SyncActivity,omitempty" xml:"SyncActivity,omitempty"`
}

func (s ScaleWithAdjustmentRequest) String() string {
	return tea.Prettify(s)
}

func (s ScaleWithAdjustmentRequest) GoString() string {
	return s.String()
}

func (s *ScaleWithAdjustmentRequest) SetAdjustmentType(v string) *ScaleWithAdjustmentRequest {
	s.AdjustmentType = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetAdjustmentValue(v int32) *ScaleWithAdjustmentRequest {
	s.AdjustmentValue = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetClientToken(v string) *ScaleWithAdjustmentRequest {
	s.ClientToken = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetMinAdjustmentMagnitude(v int32) *ScaleWithAdjustmentRequest {
	s.MinAdjustmentMagnitude = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetOwnerId(v int64) *ScaleWithAdjustmentRequest {
	s.OwnerId = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetResourceOwnerAccount(v string) *ScaleWithAdjustmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetScalingGroupId(v string) *ScaleWithAdjustmentRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetSyncActivity(v bool) *ScaleWithAdjustmentRequest {
	s.SyncActivity = &v
	return s
}

type ScaleWithAdjustmentResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s ScaleWithAdjustmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleWithAdjustmentResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleWithAdjustmentResponseBody) SetRequestId(v string) *ScaleWithAdjustmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleWithAdjustmentResponseBody) SetScalingActivityId(v string) *ScaleWithAdjustmentResponseBody {
	s.ScalingActivityId = &v
	return s
}

type ScaleWithAdjustmentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ScaleWithAdjustmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScaleWithAdjustmentResponse) String() string {
	return tea.Prettify(s)
}

func (s ScaleWithAdjustmentResponse) GoString() string {
	return s.String()
}

func (s *ScaleWithAdjustmentResponse) SetHeaders(v map[string]*string) *ScaleWithAdjustmentResponse {
	s.Headers = v
	return s
}

func (s *ScaleWithAdjustmentResponse) SetStatusCode(v int32) *ScaleWithAdjustmentResponse {
	s.StatusCode = &v
	return s
}

func (s *ScaleWithAdjustmentResponse) SetBody(v *ScaleWithAdjustmentResponseBody) *ScaleWithAdjustmentResponse {
	s.Body = v
	return s
}

type SetGroupDeletionProtectionRequest struct {
	GroupDeletionProtection *bool   `json:"GroupDeletionProtection,omitempty" xml:"GroupDeletionProtection,omitempty"`
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount    *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId          *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s SetGroupDeletionProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGroupDeletionProtectionRequest) GoString() string {
	return s.String()
}

func (s *SetGroupDeletionProtectionRequest) SetGroupDeletionProtection(v bool) *SetGroupDeletionProtectionRequest {
	s.GroupDeletionProtection = &v
	return s
}

func (s *SetGroupDeletionProtectionRequest) SetOwnerId(v int64) *SetGroupDeletionProtectionRequest {
	s.OwnerId = &v
	return s
}

func (s *SetGroupDeletionProtectionRequest) SetRegionId(v string) *SetGroupDeletionProtectionRequest {
	s.RegionId = &v
	return s
}

func (s *SetGroupDeletionProtectionRequest) SetResourceOwnerAccount(v string) *SetGroupDeletionProtectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetGroupDeletionProtectionRequest) SetScalingGroupId(v string) *SetGroupDeletionProtectionRequest {
	s.ScalingGroupId = &v
	return s
}

type SetGroupDeletionProtectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetGroupDeletionProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetGroupDeletionProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *SetGroupDeletionProtectionResponseBody) SetRequestId(v string) *SetGroupDeletionProtectionResponseBody {
	s.RequestId = &v
	return s
}

type SetGroupDeletionProtectionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetGroupDeletionProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetGroupDeletionProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s SetGroupDeletionProtectionResponse) GoString() string {
	return s.String()
}

func (s *SetGroupDeletionProtectionResponse) SetHeaders(v map[string]*string) *SetGroupDeletionProtectionResponse {
	s.Headers = v
	return s
}

func (s *SetGroupDeletionProtectionResponse) SetStatusCode(v int32) *SetGroupDeletionProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetGroupDeletionProtectionResponse) SetBody(v *SetGroupDeletionProtectionResponseBody) *SetGroupDeletionProtectionResponse {
	s.Body = v
	return s
}

type SetInstanceHealthRequest struct {
	HealthStatus         *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s SetInstanceHealthRequest) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceHealthRequest) GoString() string {
	return s.String()
}

func (s *SetInstanceHealthRequest) SetHealthStatus(v string) *SetInstanceHealthRequest {
	s.HealthStatus = &v
	return s
}

func (s *SetInstanceHealthRequest) SetInstanceId(v string) *SetInstanceHealthRequest {
	s.InstanceId = &v
	return s
}

func (s *SetInstanceHealthRequest) SetOwnerId(v int64) *SetInstanceHealthRequest {
	s.OwnerId = &v
	return s
}

func (s *SetInstanceHealthRequest) SetResourceOwnerAccount(v string) *SetInstanceHealthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type SetInstanceHealthResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetInstanceHealthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceHealthResponseBody) GoString() string {
	return s.String()
}

func (s *SetInstanceHealthResponseBody) SetRequestId(v string) *SetInstanceHealthResponseBody {
	s.RequestId = &v
	return s
}

type SetInstanceHealthResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetInstanceHealthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetInstanceHealthResponse) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceHealthResponse) GoString() string {
	return s.String()
}

func (s *SetInstanceHealthResponse) SetHeaders(v map[string]*string) *SetInstanceHealthResponse {
	s.Headers = v
	return s
}

func (s *SetInstanceHealthResponse) SetStatusCode(v int32) *SetInstanceHealthResponse {
	s.StatusCode = &v
	return s
}

func (s *SetInstanceHealthResponse) SetBody(v *SetInstanceHealthResponseBody) *SetInstanceHealthResponse {
	s.Body = v
	return s
}

type SetInstancesProtectionRequest struct {
	InstanceIds          []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProtectedFromScaleIn *bool     `json:"ProtectedFromScaleIn,omitempty" xml:"ProtectedFromScaleIn,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s SetInstancesProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetInstancesProtectionRequest) GoString() string {
	return s.String()
}

func (s *SetInstancesProtectionRequest) SetInstanceIds(v []*string) *SetInstancesProtectionRequest {
	s.InstanceIds = v
	return s
}

func (s *SetInstancesProtectionRequest) SetOwnerId(v int64) *SetInstancesProtectionRequest {
	s.OwnerId = &v
	return s
}

func (s *SetInstancesProtectionRequest) SetProtectedFromScaleIn(v bool) *SetInstancesProtectionRequest {
	s.ProtectedFromScaleIn = &v
	return s
}

func (s *SetInstancesProtectionRequest) SetResourceOwnerAccount(v string) *SetInstancesProtectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetInstancesProtectionRequest) SetScalingGroupId(v string) *SetInstancesProtectionRequest {
	s.ScalingGroupId = &v
	return s
}

type SetInstancesProtectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetInstancesProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetInstancesProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *SetInstancesProtectionResponseBody) SetRequestId(v string) *SetInstancesProtectionResponseBody {
	s.RequestId = &v
	return s
}

type SetInstancesProtectionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetInstancesProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetInstancesProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s SetInstancesProtectionResponse) GoString() string {
	return s.String()
}

func (s *SetInstancesProtectionResponse) SetHeaders(v map[string]*string) *SetInstancesProtectionResponse {
	s.Headers = v
	return s
}

func (s *SetInstancesProtectionResponse) SetStatusCode(v int32) *SetInstancesProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetInstancesProtectionResponse) SetBody(v *SetInstancesProtectionResponseBody) *SetInstancesProtectionResponse {
	s.Body = v
	return s
}

type SuspendProcessesRequest struct {
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Processes            []*string `json:"Processes,omitempty" xml:"Processes,omitempty" type:"Repeated"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s SuspendProcessesRequest) String() string {
	return tea.Prettify(s)
}

func (s SuspendProcessesRequest) GoString() string {
	return s.String()
}

func (s *SuspendProcessesRequest) SetClientToken(v string) *SuspendProcessesRequest {
	s.ClientToken = &v
	return s
}

func (s *SuspendProcessesRequest) SetOwnerId(v int64) *SuspendProcessesRequest {
	s.OwnerId = &v
	return s
}

func (s *SuspendProcessesRequest) SetProcesses(v []*string) *SuspendProcessesRequest {
	s.Processes = v
	return s
}

func (s *SuspendProcessesRequest) SetRegionId(v string) *SuspendProcessesRequest {
	s.RegionId = &v
	return s
}

func (s *SuspendProcessesRequest) SetResourceOwnerAccount(v string) *SuspendProcessesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SuspendProcessesRequest) SetScalingGroupId(v string) *SuspendProcessesRequest {
	s.ScalingGroupId = &v
	return s
}

type SuspendProcessesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SuspendProcessesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SuspendProcessesResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendProcessesResponseBody) SetRequestId(v string) *SuspendProcessesResponseBody {
	s.RequestId = &v
	return s
}

type SuspendProcessesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SuspendProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SuspendProcessesResponse) String() string {
	return tea.Prettify(s)
}

func (s SuspendProcessesResponse) GoString() string {
	return s.String()
}

func (s *SuspendProcessesResponse) SetHeaders(v map[string]*string) *SuspendProcessesResponse {
	s.Headers = v
	return s
}

func (s *SuspendProcessesResponse) SetStatusCode(v int32) *SuspendProcessesResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendProcessesResponse) SetBody(v *SuspendProcessesResponseBody) *SuspendProcessesResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerId              *int64                     `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceIds          []*string                  `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                    `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceType         *string                    `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags                 []*TagResourcesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceIds(v []*string) *TagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTags(v []*TagResourcesRequestTags) *TagResourcesRequest {
	s.Tags = v
	return s
}

type TagResourcesRequestTags struct {
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Propagate *bool   `json:"Propagate,omitempty" xml:"Propagate,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTags) SetKey(v string) *TagResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTags) SetPropagate(v bool) *TagResourcesRequestTags {
	s.Propagate = &v
	return s
}

func (s *TagResourcesRequestTags) SetValue(v string) *TagResourcesRequestTags {
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
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceIds          []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceType         *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKeys              []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
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

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceIds(v []*string) *UntagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerAccount(v string) *UntagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKeys(v []*string) *UntagResourcesRequest {
	s.TagKeys = v
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

type VerifyAuthenticationRequest struct {
	OnlyCheck            *bool   `json:"OnlyCheck,omitempty" xml:"OnlyCheck,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Uid                  *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s VerifyAuthenticationRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyAuthenticationRequest) GoString() string {
	return s.String()
}

func (s *VerifyAuthenticationRequest) SetOnlyCheck(v bool) *VerifyAuthenticationRequest {
	s.OnlyCheck = &v
	return s
}

func (s *VerifyAuthenticationRequest) SetOwnerId(v int64) *VerifyAuthenticationRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyAuthenticationRequest) SetResourceOwnerAccount(v string) *VerifyAuthenticationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *VerifyAuthenticationRequest) SetResourceOwnerId(v int64) *VerifyAuthenticationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *VerifyAuthenticationRequest) SetUid(v int64) *VerifyAuthenticationRequest {
	s.Uid = &v
	return s
}

type VerifyAuthenticationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyAuthenticationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyAuthenticationResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyAuthenticationResponseBody) SetRequestId(v string) *VerifyAuthenticationResponseBody {
	s.RequestId = &v
	return s
}

type VerifyAuthenticationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifyAuthenticationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyAuthenticationResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyAuthenticationResponse) GoString() string {
	return s.String()
}

func (s *VerifyAuthenticationResponse) SetHeaders(v map[string]*string) *VerifyAuthenticationResponse {
	s.Headers = v
	return s
}

func (s *VerifyAuthenticationResponse) SetStatusCode(v int32) *VerifyAuthenticationResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyAuthenticationResponse) SetBody(v *VerifyAuthenticationResponseBody) *VerifyAuthenticationResponse {
	s.Body = v
	return s
}

type VerifyUserRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s VerifyUserRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyUserRequest) GoString() string {
	return s.String()
}

func (s *VerifyUserRequest) SetOwnerId(v int64) *VerifyUserRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyUserRequest) SetResourceOwnerAccount(v string) *VerifyUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *VerifyUserRequest) SetResourceOwnerId(v int64) *VerifyUserRequest {
	s.ResourceOwnerId = &v
	return s
}

type VerifyUserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyUserResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyUserResponseBody) SetRequestId(v string) *VerifyUserResponseBody {
	s.RequestId = &v
	return s
}

type VerifyUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifyUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyUserResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyUserResponse) GoString() string {
	return s.String()
}

func (s *VerifyUserResponse) SetHeaders(v map[string]*string) *VerifyUserResponse {
	s.Headers = v
	return s
}

func (s *VerifyUserResponse) SetStatusCode(v int32) *VerifyUserResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyUserResponse) SetBody(v *VerifyUserResponseBody) *VerifyUserResponse {
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
		"cn-qingdao":                  tea.String("ess.aliyuncs.com"),
		"cn-beijing":                  tea.String("ess.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("ess.aliyuncs.com"),
		"cn-shanghai":                 tea.String("ess.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("ess.aliyuncs.com"),
		"cn-hongkong":                 tea.String("ess.aliyuncs.com"),
		"ap-southeast-1":              tea.String("ess.aliyuncs.com"),
		"us-east-1":                   tea.String("ess.aliyuncs.com"),
		"us-west-1":                   tea.String("ess.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("ess.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("ess.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("ess.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("ess.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("ess.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("ess.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("ess.aliyuncs.com"),
		"cn-edge-1":                   tea.String("ess.aliyuncs.com"),
		"cn-fujian":                   tea.String("ess.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("ess.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("ess.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("ess.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("ess.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("ess.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("ess.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("ess.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("ess.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("ess.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("ess.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("ess.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("ess.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("ess.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("ess.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("ess.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("ess.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("ess.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("ess.aliyuncs.com"),
		"cn-wuhan":                    tea.String("ess.aliyuncs.com"),
		"cn-yushanfang":               tea.String("ess.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("ess.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("ess.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("ess.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("ess.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("ess.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("ess.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ess"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AttachAlbServerGroupsWithOptions(request *AttachAlbServerGroupsRequest, runtime *util.RuntimeOptions) (_result *AttachAlbServerGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlbServerGroups)) {
		query["AlbServerGroups"] = request.AlbServerGroups
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ForceAttach)) {
		query["ForceAttach"] = request.ForceAttach
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachAlbServerGroups"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachAlbServerGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachAlbServerGroups(request *AttachAlbServerGroupsRequest) (_result *AttachAlbServerGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachAlbServerGroupsResponse{}
	_body, _err := client.AttachAlbServerGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachDBInstancesWithOptions(request *AttachDBInstancesRequest, runtime *util.RuntimeOptions) (_result *AttachDBInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstances)) {
		query["DBInstances"] = request.DBInstances
	}

	if !tea.BoolValue(util.IsUnset(request.ForceAttach)) {
		query["ForceAttach"] = request.ForceAttach
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachDBInstances"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachDBInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachDBInstances(request *AttachDBInstancesRequest) (_result *AttachDBInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachDBInstancesResponse{}
	_body, _err := client.AttachDBInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachInstancesWithOptions(request *AttachInstancesRequest, runtime *util.RuntimeOptions) (_result *AttachInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Entrusted)) {
		query["Entrusted"] = request.Entrusted
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleHook)) {
		query["LifecycleHook"] = request.LifecycleHook
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerWeights)) {
		query["LoadBalancerWeights"] = request.LoadBalancerWeights
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachInstances"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachInstances(request *AttachInstancesRequest) (_result *AttachInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachInstancesResponse{}
	_body, _err := client.AttachInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachLoadBalancersWithOptions(request *AttachLoadBalancersRequest, runtime *util.RuntimeOptions) (_result *AttachLoadBalancersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Async)) {
		query["Async"] = request.Async
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ForceAttach)) {
		query["ForceAttach"] = request.ForceAttach
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancers)) {
		query["LoadBalancers"] = request.LoadBalancers
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachLoadBalancers"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachLoadBalancersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachLoadBalancers(request *AttachLoadBalancersRequest) (_result *AttachLoadBalancersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachLoadBalancersResponse{}
	_body, _err := client.AttachLoadBalancersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachVServerGroupsWithOptions(request *AttachVServerGroupsRequest, runtime *util.RuntimeOptions) (_result *AttachVServerGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ForceAttach)) {
		query["ForceAttach"] = request.ForceAttach
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroups)) {
		query["VServerGroups"] = request.VServerGroups
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachVServerGroups"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachVServerGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachVServerGroups(request *AttachVServerGroupsRequest) (_result *AttachVServerGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachVServerGroupsResponse{}
	_body, _err := client.AttachVServerGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompleteLifecycleActionWithOptions(request *CompleteLifecycleActionRequest, runtime *util.RuntimeOptions) (_result *CompleteLifecycleActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleActionResult)) {
		query["LifecycleActionResult"] = request.LifecycleActionResult
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleActionToken)) {
		query["LifecycleActionToken"] = request.LifecycleActionToken
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleHookId)) {
		query["LifecycleHookId"] = request.LifecycleHookId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CompleteLifecycleAction"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CompleteLifecycleActionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompleteLifecycleAction(request *CompleteLifecycleActionRequest) (_result *CompleteLifecycleActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompleteLifecycleActionResponse{}
	_body, _err := client.CompleteLifecycleActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAlarmWithOptions(request *CreateAlarmRequest, runtime *util.RuntimeOptions) (_result *CreateAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmActions)) {
		query["AlarmActions"] = request.AlarmActions
	}

	if !tea.BoolValue(util.IsUnset(request.ComparisonOperator)) {
		query["ComparisonOperator"] = request.ComparisonOperator
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		query["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.Effective)) {
		query["Effective"] = request.Effective
	}

	if !tea.BoolValue(util.IsUnset(request.EvaluationCount)) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !tea.BoolValue(util.IsUnset(request.Expressions)) {
		query["Expressions"] = request.Expressions
	}

	if !tea.BoolValue(util.IsUnset(request.ExpressionsLogicOperator)) {
		query["ExpressionsLogicOperator"] = request.ExpressionsLogicOperator
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Statistics)) {
		query["Statistics"] = request.Statistics
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAlarm"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAlarm(request *CreateAlarmRequest) (_result *CreateAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAlarmResponse{}
	_body, _err := client.CreateAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEciScalingConfigurationWithOptions(request *CreateEciScalingConfigurationRequest, runtime *util.RuntimeOptions) (_result *CreateEciScalingConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcrRegistryInfos)) {
		query["AcrRegistryInfos"] = request.AcrRegistryInfos
	}

	if !tea.BoolValue(util.IsUnset(request.ActiveDeadlineSeconds)) {
		query["ActiveDeadlineSeconds"] = request.ActiveDeadlineSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.AutoCreateEip)) {
		query["AutoCreateEip"] = request.AutoCreateEip
	}

	if !tea.BoolValue(util.IsUnset(request.AutoMatchImageCache)) {
		query["AutoMatchImageCache"] = request.AutoMatchImageCache
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerGroupName)) {
		query["ContainerGroupName"] = request.ContainerGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Containers)) {
		query["Containers"] = request.Containers
	}

	if !tea.BoolValue(util.IsUnset(request.CostOptimization)) {
		query["CostOptimization"] = request.CostOptimization
	}

	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		query["Cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.CpuOptionsCore)) {
		query["CpuOptionsCore"] = request.CpuOptionsCore
	}

	if !tea.BoolValue(util.IsUnset(request.CpuOptionsThreadsPerCore)) {
		query["CpuOptionsThreadsPerCore"] = request.CpuOptionsThreadsPerCore
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DnsConfigNameServers)) {
		query["DnsConfigNameServers"] = request.DnsConfigNameServers
	}

	if !tea.BoolValue(util.IsUnset(request.DnsConfigOptions)) {
		query["DnsConfigOptions"] = request.DnsConfigOptions
	}

	if !tea.BoolValue(util.IsUnset(request.DnsConfigSearchs)) {
		query["DnsConfigSearchs"] = request.DnsConfigSearchs
	}

	if !tea.BoolValue(util.IsUnset(request.DnsPolicy)) {
		query["DnsPolicy"] = request.DnsPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.EgressBandwidth)) {
		query["EgressBandwidth"] = request.EgressBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.EipBandwidth)) {
		query["EipBandwidth"] = request.EipBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.EnableSls)) {
		query["EnableSls"] = request.EnableSls
	}

	if !tea.BoolValue(util.IsUnset(request.EphemeralStorage)) {
		query["EphemeralStorage"] = request.EphemeralStorage
	}

	if !tea.BoolValue(util.IsUnset(request.HostAliases)) {
		query["HostAliases"] = request.HostAliases
	}

	if !tea.BoolValue(util.IsUnset(request.HostName)) {
		query["HostName"] = request.HostName
	}

	if !tea.BoolValue(util.IsUnset(request.ImageRegistryCredentials)) {
		query["ImageRegistryCredentials"] = request.ImageRegistryCredentials
	}

	if !tea.BoolValue(util.IsUnset(request.ImageSnapshotId)) {
		query["ImageSnapshotId"] = request.ImageSnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.IngressBandwidth)) {
		query["IngressBandwidth"] = request.IngressBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.InitContainers)) {
		query["InitContainers"] = request.InitContainers
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceFamilyLevel)) {
		query["InstanceFamilyLevel"] = request.InstanceFamilyLevel
	}

	if !tea.BoolValue(util.IsUnset(request.Ipv6AddressCount)) {
		query["Ipv6AddressCount"] = request.Ipv6AddressCount
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerWeight)) {
		query["LoadBalancerWeight"] = request.LoadBalancerWeight
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		query["Memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.NtpServers)) {
		query["NtpServers"] = request.NtpServers
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RamRoleName)) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.RestartPolicy)) {
		query["RestartPolicy"] = request.RestartPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingConfigurationName)) {
		query["ScalingConfigurationName"] = request.ScalingConfigurationName
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityContextSysctls)) {
		query["SecurityContextSysctls"] = request.SecurityContextSysctls
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SpotPriceLimit)) {
		query["SpotPriceLimit"] = request.SpotPriceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.SpotStrategy)) {
		query["SpotStrategy"] = request.SpotStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.TerminationGracePeriodSeconds)) {
		query["TerminationGracePeriodSeconds"] = request.TerminationGracePeriodSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.Volumes)) {
		query["Volumes"] = request.Volumes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEciScalingConfiguration"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEciScalingConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEciScalingConfiguration(request *CreateEciScalingConfigurationRequest) (_result *CreateEciScalingConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEciScalingConfigurationResponse{}
	_body, _err := client.CreateEciScalingConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLifecycleHookWithOptions(request *CreateLifecycleHookRequest, runtime *util.RuntimeOptions) (_result *CreateLifecycleHookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefaultResult)) {
		query["DefaultResult"] = request.DefaultResult
	}

	if !tea.BoolValue(util.IsUnset(request.HeartbeatTimeout)) {
		query["HeartbeatTimeout"] = request.HeartbeatTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleHookName)) {
		query["LifecycleHookName"] = request.LifecycleHookName
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleTransition)) {
		query["LifecycleTransition"] = request.LifecycleTransition
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationArn)) {
		query["NotificationArn"] = request.NotificationArn
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationMetadata)) {
		query["NotificationMetadata"] = request.NotificationMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLifecycleHook"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLifecycleHookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLifecycleHook(request *CreateLifecycleHookRequest) (_result *CreateLifecycleHookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLifecycleHookResponse{}
	_body, _err := client.CreateLifecycleHookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateNotificationConfigurationWithOptions(request *CreateNotificationConfigurationRequest, runtime *util.RuntimeOptions) (_result *CreateNotificationConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NotificationArn)) {
		query["NotificationArn"] = request.NotificationArn
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationTypes)) {
		query["NotificationTypes"] = request.NotificationTypes
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNotificationConfiguration"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNotificationConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateNotificationConfiguration(request *CreateNotificationConfigurationRequest) (_result *CreateNotificationConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNotificationConfigurationResponse{}
	_body, _err := client.CreateNotificationConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateScalingConfigurationWithOptions(tmpReq *CreateScalingConfigurationRequest, runtime *util.RuntimeOptions) (_result *CreateScalingConfigurationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateScalingConfigurationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SchedulerOptions)) {
		request.SchedulerOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SchedulerOptions, tea.String("SchedulerOptions"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Affinity)) {
		query["Affinity"] = request.Affinity
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		query["Cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.CreditSpecification)) {
		query["CreditSpecification"] = request.CreditSpecification
	}

	if !tea.BoolValue(util.IsUnset(request.DataDisks)) {
		query["DataDisks"] = request.DataDisks
	}

	if !tea.BoolValue(util.IsUnset(request.DedicatedHostId)) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !tea.BoolValue(util.IsUnset(request.DeploymentSetId)) {
		query["DeploymentSetId"] = request.DeploymentSetId
	}

	if !tea.BoolValue(util.IsUnset(request.HostName)) {
		query["HostName"] = request.HostName
	}

	if !tea.BoolValue(util.IsUnset(request.HpcClusterId)) {
		query["HpcClusterId"] = request.HpcClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageFamily)) {
		query["ImageFamily"] = request.ImageFamily
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageName)) {
		query["ImageName"] = request.ImageName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceDescription)) {
		query["InstanceDescription"] = request.InstanceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.InstancePatternInfos)) {
		query["InstancePatternInfos"] = request.InstancePatternInfos
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceTypeOverrides)) {
		query["InstanceTypeOverrides"] = request.InstanceTypeOverrides
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceTypes)) {
		query["InstanceTypes"] = request.InstanceTypes
	}

	if !tea.BoolValue(util.IsUnset(request.InternetChargeType)) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.InternetMaxBandwidthIn)) {
		query["InternetMaxBandwidthIn"] = request.InternetMaxBandwidthIn
	}

	if !tea.BoolValue(util.IsUnset(request.InternetMaxBandwidthOut)) {
		query["InternetMaxBandwidthOut"] = request.InternetMaxBandwidthOut
	}

	if !tea.BoolValue(util.IsUnset(request.IoOptimized)) {
		query["IoOptimized"] = request.IoOptimized
	}

	if !tea.BoolValue(util.IsUnset(request.Ipv6AddressCount)) {
		query["Ipv6AddressCount"] = request.Ipv6AddressCount
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerWeight)) {
		query["LoadBalancerWeight"] = request.LoadBalancerWeight
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		query["Memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordInherit)) {
		query["PasswordInherit"] = request.PasswordInherit
	}

	if !tea.BoolValue(util.IsUnset(request.RamRoleName)) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingConfigurationName)) {
		query["ScalingConfigurationName"] = request.ScalingConfigurationName
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SchedulerOptionsShrink)) {
		query["SchedulerOptions"] = request.SchedulerOptionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityEnhancementStrategy)) {
		query["SecurityEnhancementStrategy"] = request.SecurityEnhancementStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupIds)) {
		query["SecurityGroupIds"] = request.SecurityGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.SpotDuration)) {
		query["SpotDuration"] = request.SpotDuration
	}

	if !tea.BoolValue(util.IsUnset(request.SpotInterruptionBehavior)) {
		query["SpotInterruptionBehavior"] = request.SpotInterruptionBehavior
	}

	if !tea.BoolValue(util.IsUnset(request.SpotPriceLimits)) {
		query["SpotPriceLimits"] = request.SpotPriceLimits
	}

	if !tea.BoolValue(util.IsUnset(request.SpotStrategy)) {
		query["SpotStrategy"] = request.SpotStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.SystemDiskCategories)) {
		query["SystemDiskCategories"] = request.SystemDiskCategories
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.Tenancy)) {
		query["Tenancy"] = request.Tenancy
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		query["UserData"] = request.UserData
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.PrivatePoolOptions))) {
		query["PrivatePoolOptions"] = request.PrivatePoolOptions
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SystemDisk))) {
		query["SystemDisk"] = request.SystemDisk
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateScalingConfiguration"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateScalingConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateScalingConfiguration(request *CreateScalingConfigurationRequest) (_result *CreateScalingConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScalingConfigurationResponse{}
	_body, _err := client.CreateScalingConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateScalingGroupWithOptions(request *CreateScalingGroupRequest, runtime *util.RuntimeOptions) (_result *CreateScalingGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlbServerGroups)) {
		query["AlbServerGroups"] = request.AlbServerGroups
	}

	if !tea.BoolValue(util.IsUnset(request.AllocationStrategy)) {
		query["AllocationStrategy"] = request.AllocationStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.AzBalance)) {
		query["AzBalance"] = request.AzBalance
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CompensateWithOnDemand)) {
		query["CompensateWithOnDemand"] = request.CompensateWithOnDemand
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstanceIds)) {
		query["DBInstanceIds"] = request.DBInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultCooldown)) {
		query["DefaultCooldown"] = request.DefaultCooldown
	}

	if !tea.BoolValue(util.IsUnset(request.DesiredCapacity)) {
		query["DesiredCapacity"] = request.DesiredCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.GroupDeletionProtection)) {
		query["GroupDeletionProtection"] = request.GroupDeletionProtection
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		query["GroupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckType)) {
		query["HealthCheckType"] = request.HealthCheckType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchTemplateId)) {
		query["LaunchTemplateId"] = request.LaunchTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchTemplateOverrides)) {
		query["LaunchTemplateOverrides"] = request.LaunchTemplateOverrides
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchTemplateVersion)) {
		query["LaunchTemplateVersion"] = request.LaunchTemplateVersion
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleHooks)) {
		query["LifecycleHooks"] = request.LifecycleHooks
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerIds)) {
		query["LoadBalancerIds"] = request.LoadBalancerIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxInstanceLifetime)) {
		query["MaxInstanceLifetime"] = request.MaxInstanceLifetime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSize)) {
		query["MaxSize"] = request.MaxSize
	}

	if !tea.BoolValue(util.IsUnset(request.MinSize)) {
		query["MinSize"] = request.MinSize
	}

	if !tea.BoolValue(util.IsUnset(request.MultiAZPolicy)) {
		query["MultiAZPolicy"] = request.MultiAZPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.OnDemandBaseCapacity)) {
		query["OnDemandBaseCapacity"] = request.OnDemandBaseCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.OnDemandPercentageAboveBaseCapacity)) {
		query["OnDemandPercentageAboveBaseCapacity"] = request.OnDemandPercentageAboveBaseCapacity
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

	if !tea.BoolValue(util.IsUnset(request.RemovalPolicies)) {
		query["RemovalPolicies"] = request.RemovalPolicies
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupName)) {
		query["ScalingGroupName"] = request.ScalingGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingPolicy)) {
		query["ScalingPolicy"] = request.ScalingPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.SpotAllocationStrategy)) {
		query["SpotAllocationStrategy"] = request.SpotAllocationStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.SpotInstancePools)) {
		query["SpotInstancePools"] = request.SpotInstancePools
	}

	if !tea.BoolValue(util.IsUnset(request.SpotInstanceRemedy)) {
		query["SpotInstanceRemedy"] = request.SpotInstanceRemedy
	}

	if !tea.BoolValue(util.IsUnset(request.SyncAlarmRuleToCms)) {
		query["SyncAlarmRuleToCms"] = request.SyncAlarmRuleToCms
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroups)) {
		query["VServerGroups"] = request.VServerGroups
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchIds)) {
		query["VSwitchIds"] = request.VSwitchIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateScalingGroup"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateScalingGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateScalingGroup(request *CreateScalingGroupRequest) (_result *CreateScalingGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScalingGroupResponse{}
	_body, _err := client.CreateScalingGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateScalingRuleWithOptions(request *CreateScalingRuleRequest, runtime *util.RuntimeOptions) (_result *CreateScalingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdjustmentType)) {
		query["AdjustmentType"] = request.AdjustmentType
	}

	if !tea.BoolValue(util.IsUnset(request.AdjustmentValue)) {
		query["AdjustmentValue"] = request.AdjustmentValue
	}

	if !tea.BoolValue(util.IsUnset(request.Cooldown)) {
		query["Cooldown"] = request.Cooldown
	}

	if !tea.BoolValue(util.IsUnset(request.DisableScaleIn)) {
		query["DisableScaleIn"] = request.DisableScaleIn
	}

	if !tea.BoolValue(util.IsUnset(request.EstimatedInstanceWarmup)) {
		query["EstimatedInstanceWarmup"] = request.EstimatedInstanceWarmup
	}

	if !tea.BoolValue(util.IsUnset(request.InitialMaxSize)) {
		query["InitialMaxSize"] = request.InitialMaxSize
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.MinAdjustmentMagnitude)) {
		query["MinAdjustmentMagnitude"] = request.MinAdjustmentMagnitude
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PredictiveScalingMode)) {
		query["PredictiveScalingMode"] = request.PredictiveScalingMode
	}

	if !tea.BoolValue(util.IsUnset(request.PredictiveTaskBufferTime)) {
		query["PredictiveTaskBufferTime"] = request.PredictiveTaskBufferTime
	}

	if !tea.BoolValue(util.IsUnset(request.PredictiveValueBehavior)) {
		query["PredictiveValueBehavior"] = request.PredictiveValueBehavior
	}

	if !tea.BoolValue(util.IsUnset(request.PredictiveValueBuffer)) {
		query["PredictiveValueBuffer"] = request.PredictiveValueBuffer
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleInEvaluationCount)) {
		query["ScaleInEvaluationCount"] = request.ScaleInEvaluationCount
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleOutEvaluationCount)) {
		query["ScaleOutEvaluationCount"] = request.ScaleOutEvaluationCount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleName)) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleType)) {
		query["ScalingRuleType"] = request.ScalingRuleType
	}

	if !tea.BoolValue(util.IsUnset(request.StepAdjustments)) {
		query["StepAdjustments"] = request.StepAdjustments
	}

	if !tea.BoolValue(util.IsUnset(request.TargetValue)) {
		query["TargetValue"] = request.TargetValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateScalingRule"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateScalingRule(request *CreateScalingRuleRequest) (_result *CreateScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScalingRuleResponse{}
	_body, _err := client.CreateScalingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateScheduledTaskWithOptions(request *CreateScheduledTaskRequest, runtime *util.RuntimeOptions) (_result *CreateScheduledTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DesiredCapacity)) {
		query["DesiredCapacity"] = request.DesiredCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchExpirationTime)) {
		query["LaunchExpirationTime"] = request.LaunchExpirationTime
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchTime)) {
		query["LaunchTime"] = request.LaunchTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxValue)) {
		query["MaxValue"] = request.MaxValue
	}

	if !tea.BoolValue(util.IsUnset(request.MinValue)) {
		query["MinValue"] = request.MinValue
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RecurrenceEndTime)) {
		query["RecurrenceEndTime"] = request.RecurrenceEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RecurrenceType)) {
		query["RecurrenceType"] = request.RecurrenceType
	}

	if !tea.BoolValue(util.IsUnset(request.RecurrenceValue)) {
		query["RecurrenceValue"] = request.RecurrenceValue
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledAction)) {
		query["ScheduledAction"] = request.ScheduledAction
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledTaskName)) {
		query["ScheduledTaskName"] = request.ScheduledTaskName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskEnabled)) {
		query["TaskEnabled"] = request.TaskEnabled
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateScheduledTask"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateScheduledTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateScheduledTask(request *CreateScheduledTaskRequest) (_result *CreateScheduledTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScheduledTaskResponse{}
	_body, _err := client.CreateScheduledTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeactivateScalingConfigurationWithOptions(request *DeactivateScalingConfigurationRequest, runtime *util.RuntimeOptions) (_result *DeactivateScalingConfigurationResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingConfigurationId)) {
		query["ScalingConfigurationId"] = request.ScalingConfigurationId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeactivateScalingConfiguration"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeactivateScalingConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeactivateScalingConfiguration(request *DeactivateScalingConfigurationRequest) (_result *DeactivateScalingConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeactivateScalingConfigurationResponse{}
	_body, _err := client.DeactivateScalingConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAlarmWithOptions(request *DeleteAlarmRequest, runtime *util.RuntimeOptions) (_result *DeleteAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmTaskId)) {
		query["AlarmTaskId"] = request.AlarmTaskId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAlarm"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAlarm(request *DeleteAlarmRequest) (_result *DeleteAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAlarmResponse{}
	_body, _err := client.DeleteAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLifecycleHookWithOptions(request *DeleteLifecycleHookRequest, runtime *util.RuntimeOptions) (_result *DeleteLifecycleHookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LifecycleHookId)) {
		query["LifecycleHookId"] = request.LifecycleHookId
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleHookName)) {
		query["LifecycleHookName"] = request.LifecycleHookName
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLifecycleHook"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLifecycleHookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLifecycleHook(request *DeleteLifecycleHookRequest) (_result *DeleteLifecycleHookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLifecycleHookResponse{}
	_body, _err := client.DeleteLifecycleHookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNotificationConfigurationWithOptions(request *DeleteNotificationConfigurationRequest, runtime *util.RuntimeOptions) (_result *DeleteNotificationConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NotificationArn)) {
		query["NotificationArn"] = request.NotificationArn
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNotificationConfiguration"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNotificationConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNotificationConfiguration(request *DeleteNotificationConfigurationRequest) (_result *DeleteNotificationConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNotificationConfigurationResponse{}
	_body, _err := client.DeleteNotificationConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteScalingConfigurationWithOptions(request *DeleteScalingConfigurationRequest, runtime *util.RuntimeOptions) (_result *DeleteScalingConfigurationResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingConfigurationId)) {
		query["ScalingConfigurationId"] = request.ScalingConfigurationId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScalingConfiguration"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteScalingConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteScalingConfiguration(request *DeleteScalingConfigurationRequest) (_result *DeleteScalingConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScalingConfigurationResponse{}
	_body, _err := client.DeleteScalingConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteScalingGroupWithOptions(request *DeleteScalingGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteScalingGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ForceDelete)) {
		query["ForceDelete"] = request.ForceDelete
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScalingGroup"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteScalingGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteScalingGroup(request *DeleteScalingGroupRequest) (_result *DeleteScalingGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScalingGroupResponse{}
	_body, _err := client.DeleteScalingGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteScalingRuleWithOptions(request *DeleteScalingRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteScalingRuleResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleId)) {
		query["ScalingRuleId"] = request.ScalingRuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScalingRule"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteScalingRule(request *DeleteScalingRuleRequest) (_result *DeleteScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScalingRuleResponse{}
	_body, _err := client.DeleteScalingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteScheduledTaskWithOptions(request *DeleteScheduledTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteScheduledTaskResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ScheduledTaskId)) {
		query["ScheduledTaskId"] = request.ScheduledTaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScheduledTask"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteScheduledTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteScheduledTask(request *DeleteScheduledTaskRequest) (_result *DeleteScheduledTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScheduledTaskResponse{}
	_body, _err := client.DeleteScheduledTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAlarmsWithOptions(request *DescribeAlarmsRequest, runtime *util.RuntimeOptions) (_result *DescribeAlarmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmTaskId)) {
		query["AlarmTaskId"] = request.AlarmTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.IsEnable)) {
		query["IsEnable"] = request.IsEnable
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlarms"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlarmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAlarms(request *DescribeAlarmsRequest) (_result *DescribeAlarmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlarmsResponse{}
	_body, _err := client.DescribeAlarmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEciScalingConfigurationsWithOptions(request *DescribeEciScalingConfigurationsRequest, runtime *util.RuntimeOptions) (_result *DescribeEciScalingConfigurationsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ScalingConfigurationIds)) {
		query["ScalingConfigurationIds"] = request.ScalingConfigurationIds
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingConfigurationNames)) {
		query["ScalingConfigurationNames"] = request.ScalingConfigurationNames
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEciScalingConfigurations"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEciScalingConfigurationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEciScalingConfigurations(request *DescribeEciScalingConfigurationsRequest) (_result *DescribeEciScalingConfigurationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEciScalingConfigurationsResponse{}
	_body, _err := client.DescribeEciScalingConfigurationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLifecycleActionsWithOptions(request *DescribeLifecycleActionsRequest, runtime *util.RuntimeOptions) (_result *DescribeLifecycleActionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LifecycleActionStatus)) {
		query["LifecycleActionStatus"] = request.LifecycleActionStatus
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
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

	if !tea.BoolValue(util.IsUnset(request.ScalingActivityId)) {
		query["ScalingActivityId"] = request.ScalingActivityId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLifecycleActions"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLifecycleActionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLifecycleActions(request *DescribeLifecycleActionsRequest) (_result *DescribeLifecycleActionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLifecycleActionsResponse{}
	_body, _err := client.DescribeLifecycleActionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLifecycleHooksWithOptions(request *DescribeLifecycleHooksRequest, runtime *util.RuntimeOptions) (_result *DescribeLifecycleHooksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LifecycleHookIds)) {
		query["LifecycleHookIds"] = request.LifecycleHookIds
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleHookName)) {
		query["LifecycleHookName"] = request.LifecycleHookName
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLifecycleHooks"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLifecycleHooksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLifecycleHooks(request *DescribeLifecycleHooksRequest) (_result *DescribeLifecycleHooksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLifecycleHooksResponse{}
	_body, _err := client.DescribeLifecycleHooksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLimitationWithOptions(request *DescribeLimitationRequest, runtime *util.RuntimeOptions) (_result *DescribeLimitationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLimitation"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLimitationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLimitation(request *DescribeLimitationRequest) (_result *DescribeLimitationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLimitationResponse{}
	_body, _err := client.DescribeLimitationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNotificationConfigurationsWithOptions(request *DescribeNotificationConfigurationsRequest, runtime *util.RuntimeOptions) (_result *DescribeNotificationConfigurationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNotificationConfigurations"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNotificationConfigurationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNotificationConfigurations(request *DescribeNotificationConfigurationsRequest) (_result *DescribeNotificationConfigurationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNotificationConfigurationsResponse{}
	_body, _err := client.DescribeNotificationConfigurationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNotificationTypesWithOptions(request *DescribeNotificationTypesRequest, runtime *util.RuntimeOptions) (_result *DescribeNotificationTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNotificationTypes"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNotificationTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNotificationTypes(request *DescribeNotificationTypesRequest) (_result *DescribeNotificationTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNotificationTypesResponse{}
	_body, _err := client.DescribeNotificationTypesWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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
		Version:     tea.String("2022-02-22"),
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

func (client *Client) DescribeScalingActivitiesWithOptions(request *DescribeScalingActivitiesRequest, runtime *util.RuntimeOptions) (_result *DescribeScalingActivitiesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ScalingActivityIds)) {
		query["ScalingActivityIds"] = request.ScalingActivityIds
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StatusCode)) {
		query["StatusCode"] = request.StatusCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScalingActivities"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScalingActivitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScalingActivities(request *DescribeScalingActivitiesRequest) (_result *DescribeScalingActivitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScalingActivitiesResponse{}
	_body, _err := client.DescribeScalingActivitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScalingActivityDetailWithOptions(request *DescribeScalingActivityDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeScalingActivityDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingActivityId)) {
		query["ScalingActivityId"] = request.ScalingActivityId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScalingActivityDetail"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScalingActivityDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScalingActivityDetail(request *DescribeScalingActivityDetailRequest) (_result *DescribeScalingActivityDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScalingActivityDetailResponse{}
	_body, _err := client.DescribeScalingActivityDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScalingConfigurationsWithOptions(request *DescribeScalingConfigurationsRequest, runtime *util.RuntimeOptions) (_result *DescribeScalingConfigurationsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ScalingConfigurationIds)) {
		query["ScalingConfigurationIds"] = request.ScalingConfigurationIds
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingConfigurationNames)) {
		query["ScalingConfigurationNames"] = request.ScalingConfigurationNames
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScalingConfigurations"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScalingConfigurationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScalingConfigurations(request *DescribeScalingConfigurationsRequest) (_result *DescribeScalingConfigurationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScalingConfigurationsResponse{}
	_body, _err := client.DescribeScalingConfigurationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScalingGroupsWithOptions(request *DescribeScalingGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeScalingGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		query["GroupType"] = request.GroupType
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupIds)) {
		query["ScalingGroupIds"] = request.ScalingGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupName)) {
		query["ScalingGroupName"] = request.ScalingGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupNames)) {
		query["ScalingGroupNames"] = request.ScalingGroupNames
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScalingGroups"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScalingGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScalingGroups(request *DescribeScalingGroupsRequest) (_result *DescribeScalingGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScalingGroupsResponse{}
	_body, _err := client.DescribeScalingGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScalingInstancesWithOptions(request *DescribeScalingInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeScalingInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreationType)) {
		query["CreationType"] = request.CreationType
	}

	if !tea.BoolValue(util.IsUnset(request.HealthStatus)) {
		query["HealthStatus"] = request.HealthStatus
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleState)) {
		query["LifecycleState"] = request.LifecycleState
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

	if !tea.BoolValue(util.IsUnset(request.ScalingActivityId)) {
		query["ScalingActivityId"] = request.ScalingActivityId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingConfigurationId)) {
		query["ScalingConfigurationId"] = request.ScalingConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScalingInstances"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScalingInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScalingInstances(request *DescribeScalingInstancesRequest) (_result *DescribeScalingInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScalingInstancesResponse{}
	_body, _err := client.DescribeScalingInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScalingRulesWithOptions(request *DescribeScalingRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeScalingRulesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleAris)) {
		query["ScalingRuleAris"] = request.ScalingRuleAris
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleIds)) {
		query["ScalingRuleIds"] = request.ScalingRuleIds
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleNames)) {
		query["ScalingRuleNames"] = request.ScalingRuleNames
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleType)) {
		query["ScalingRuleType"] = request.ScalingRuleType
	}

	if !tea.BoolValue(util.IsUnset(request.ShowAlarmRules)) {
		query["ShowAlarmRules"] = request.ShowAlarmRules
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScalingRules"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScalingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScalingRules(request *DescribeScalingRulesRequest) (_result *DescribeScalingRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScalingRulesResponse{}
	_body, _err := client.DescribeScalingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScheduledTasksWithOptions(request *DescribeScheduledTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeScheduledTasksResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledActions)) {
		query["ScheduledActions"] = request.ScheduledActions
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledTaskIds)) {
		query["ScheduledTaskIds"] = request.ScheduledTaskIds
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledTaskNames)) {
		query["ScheduledTaskNames"] = request.ScheduledTaskNames
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScheduledTasks"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScheduledTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScheduledTasks(request *DescribeScheduledTasksRequest) (_result *DescribeScheduledTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScheduledTasksResponse{}
	_body, _err := client.DescribeScheduledTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachAlbServerGroupsWithOptions(request *DetachAlbServerGroupsRequest, runtime *util.RuntimeOptions) (_result *DetachAlbServerGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlbServerGroups)) {
		query["AlbServerGroups"] = request.AlbServerGroups
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ForceDetach)) {
		query["ForceDetach"] = request.ForceDetach
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachAlbServerGroups"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachAlbServerGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachAlbServerGroups(request *DetachAlbServerGroupsRequest) (_result *DetachAlbServerGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachAlbServerGroupsResponse{}
	_body, _err := client.DetachAlbServerGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachDBInstancesWithOptions(request *DetachDBInstancesRequest, runtime *util.RuntimeOptions) (_result *DetachDBInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DBInstances)) {
		query["DBInstances"] = request.DBInstances
	}

	if !tea.BoolValue(util.IsUnset(request.ForceDetach)) {
		query["ForceDetach"] = request.ForceDetach
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachDBInstances"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachDBInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachDBInstances(request *DetachDBInstancesRequest) (_result *DetachDBInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachDBInstancesResponse{}
	_body, _err := client.DetachDBInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachInstancesWithOptions(request *DetachInstancesRequest, runtime *util.RuntimeOptions) (_result *DetachInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DecreaseDesiredCapacity)) {
		query["DecreaseDesiredCapacity"] = request.DecreaseDesiredCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.DetachOption)) {
		query["DetachOption"] = request.DetachOption
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleHook)) {
		query["LifecycleHook"] = request.LifecycleHook
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachInstances"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachInstances(request *DetachInstancesRequest) (_result *DetachInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachInstancesResponse{}
	_body, _err := client.DetachInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachLoadBalancersWithOptions(request *DetachLoadBalancersRequest, runtime *util.RuntimeOptions) (_result *DetachLoadBalancersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Async)) {
		query["Async"] = request.Async
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ForceDetach)) {
		query["ForceDetach"] = request.ForceDetach
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancers)) {
		query["LoadBalancers"] = request.LoadBalancers
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachLoadBalancers"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachLoadBalancersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachLoadBalancers(request *DetachLoadBalancersRequest) (_result *DetachLoadBalancersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachLoadBalancersResponse{}
	_body, _err := client.DetachLoadBalancersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachVServerGroupsWithOptions(request *DetachVServerGroupsRequest, runtime *util.RuntimeOptions) (_result *DetachVServerGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ForceDetach)) {
		query["ForceDetach"] = request.ForceDetach
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VServerGroups)) {
		query["VServerGroups"] = request.VServerGroups
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachVServerGroups"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachVServerGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachVServerGroups(request *DetachVServerGroupsRequest) (_result *DetachVServerGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachVServerGroupsResponse{}
	_body, _err := client.DetachVServerGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableAlarmWithOptions(request *DisableAlarmRequest, runtime *util.RuntimeOptions) (_result *DisableAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmTaskId)) {
		query["AlarmTaskId"] = request.AlarmTaskId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableAlarm"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableAlarm(request *DisableAlarmRequest) (_result *DisableAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableAlarmResponse{}
	_body, _err := client.DisableAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableScalingGroupWithOptions(request *DisableScalingGroupRequest, runtime *util.RuntimeOptions) (_result *DisableScalingGroupResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableScalingGroup"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableScalingGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableScalingGroup(request *DisableScalingGroupRequest) (_result *DisableScalingGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableScalingGroupResponse{}
	_body, _err := client.DisableScalingGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableAlarmWithOptions(request *EnableAlarmRequest, runtime *util.RuntimeOptions) (_result *EnableAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmTaskId)) {
		query["AlarmTaskId"] = request.AlarmTaskId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableAlarm"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableAlarm(request *EnableAlarmRequest) (_result *EnableAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableAlarmResponse{}
	_body, _err := client.EnableAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableScalingGroupWithOptions(request *EnableScalingGroupRequest, runtime *util.RuntimeOptions) (_result *EnableScalingGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActiveScalingConfigurationId)) {
		query["ActiveScalingConfigurationId"] = request.ActiveScalingConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchTemplateId)) {
		query["LaunchTemplateId"] = request.LaunchTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchTemplateOverrides)) {
		query["LaunchTemplateOverrides"] = request.LaunchTemplateOverrides
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchTemplateVersion)) {
		query["LaunchTemplateVersion"] = request.LaunchTemplateVersion
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerWeights)) {
		query["LoadBalancerWeights"] = request.LoadBalancerWeights
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableScalingGroup"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableScalingGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableScalingGroup(request *EnableScalingGroupRequest) (_result *EnableScalingGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableScalingGroupResponse{}
	_body, _err := client.EnableScalingGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnterStandbyWithOptions(request *EnterStandbyRequest, runtime *util.RuntimeOptions) (_result *EnterStandbyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Async)) {
		query["Async"] = request.Async
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnterStandby"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnterStandbyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnterStandby(request *EnterStandbyRequest) (_result *EnterStandbyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnterStandbyResponse{}
	_body, _err := client.EnterStandbyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteScalingRuleWithOptions(request *ExecuteScalingRuleRequest, runtime *util.RuntimeOptions) (_result *ExecuteScalingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BreachThreshold)) {
		query["BreachThreshold"] = request.BreachThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.MetricValue)) {
		query["MetricValue"] = request.MetricValue
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

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleAri)) {
		query["ScalingRuleAri"] = request.ScalingRuleAri
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteScalingRule"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteScalingRule(request *ExecuteScalingRuleRequest) (_result *ExecuteScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteScalingRuleResponse{}
	_body, _err := client.ExecuteScalingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExitStandbyWithOptions(request *ExitStandbyRequest, runtime *util.RuntimeOptions) (_result *ExitStandbyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Async)) {
		query["Async"] = request.Async
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExitStandby"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExitStandbyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExitStandby(request *ExitStandbyRequest) (_result *ExitStandbyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExitStandbyResponse{}
	_body, _err := client.ExitStandbyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2022-02-22"),
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

func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagValues"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAlarmWithOptions(request *ModifyAlarmRequest, runtime *util.RuntimeOptions) (_result *ModifyAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmActions)) {
		query["AlarmActions"] = request.AlarmActions
	}

	if !tea.BoolValue(util.IsUnset(request.AlarmTaskId)) {
		query["AlarmTaskId"] = request.AlarmTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.ComparisonOperator)) {
		query["ComparisonOperator"] = request.ComparisonOperator
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		query["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.Effective)) {
		query["Effective"] = request.Effective
	}

	if !tea.BoolValue(util.IsUnset(request.EvaluationCount)) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !tea.BoolValue(util.IsUnset(request.Expressions)) {
		query["Expressions"] = request.Expressions
	}

	if !tea.BoolValue(util.IsUnset(request.ExpressionsLogicOperator)) {
		query["ExpressionsLogicOperator"] = request.ExpressionsLogicOperator
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.Statistics)) {
		query["Statistics"] = request.Statistics
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAlarm"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAlarm(request *ModifyAlarmRequest) (_result *ModifyAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAlarmResponse{}
	_body, _err := client.ModifyAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyEciScalingConfigurationWithOptions(request *ModifyEciScalingConfigurationRequest, runtime *util.RuntimeOptions) (_result *ModifyEciScalingConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcrRegistryInfos)) {
		query["AcrRegistryInfos"] = request.AcrRegistryInfos
	}

	if !tea.BoolValue(util.IsUnset(request.ActiveDeadlineSeconds)) {
		query["ActiveDeadlineSeconds"] = request.ActiveDeadlineSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.AutoCreateEip)) {
		query["AutoCreateEip"] = request.AutoCreateEip
	}

	if !tea.BoolValue(util.IsUnset(request.AutoMatchImageCache)) {
		query["AutoMatchImageCache"] = request.AutoMatchImageCache
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerGroupName)) {
		query["ContainerGroupName"] = request.ContainerGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Containers)) {
		query["Containers"] = request.Containers
	}

	if !tea.BoolValue(util.IsUnset(request.CostOptimization)) {
		query["CostOptimization"] = request.CostOptimization
	}

	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		query["Cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.CpuOptionsCore)) {
		query["CpuOptionsCore"] = request.CpuOptionsCore
	}

	if !tea.BoolValue(util.IsUnset(request.CpuOptionsThreadsPerCore)) {
		query["CpuOptionsThreadsPerCore"] = request.CpuOptionsThreadsPerCore
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DnsConfigNameServers)) {
		query["DnsConfigNameServers"] = request.DnsConfigNameServers
	}

	if !tea.BoolValue(util.IsUnset(request.DnsConfigOptions)) {
		query["DnsConfigOptions"] = request.DnsConfigOptions
	}

	if !tea.BoolValue(util.IsUnset(request.DnsConfigSearchs)) {
		query["DnsConfigSearchs"] = request.DnsConfigSearchs
	}

	if !tea.BoolValue(util.IsUnset(request.DnsPolicy)) {
		query["DnsPolicy"] = request.DnsPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.EgressBandwidth)) {
		query["EgressBandwidth"] = request.EgressBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.EipBandwidth)) {
		query["EipBandwidth"] = request.EipBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.EnableSls)) {
		query["EnableSls"] = request.EnableSls
	}

	if !tea.BoolValue(util.IsUnset(request.EphemeralStorage)) {
		query["EphemeralStorage"] = request.EphemeralStorage
	}

	if !tea.BoolValue(util.IsUnset(request.HostAliases)) {
		query["HostAliases"] = request.HostAliases
	}

	if !tea.BoolValue(util.IsUnset(request.HostName)) {
		query["HostName"] = request.HostName
	}

	if !tea.BoolValue(util.IsUnset(request.ImageRegistryCredentials)) {
		query["ImageRegistryCredentials"] = request.ImageRegistryCredentials
	}

	if !tea.BoolValue(util.IsUnset(request.ImageSnapshotId)) {
		query["ImageSnapshotId"] = request.ImageSnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.IngressBandwidth)) {
		query["IngressBandwidth"] = request.IngressBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.InitContainers)) {
		query["InitContainers"] = request.InitContainers
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceFamilyLevel)) {
		query["InstanceFamilyLevel"] = request.InstanceFamilyLevel
	}

	if !tea.BoolValue(util.IsUnset(request.Ipv6AddressCount)) {
		query["Ipv6AddressCount"] = request.Ipv6AddressCount
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerWeight)) {
		query["LoadBalancerWeight"] = request.LoadBalancerWeight
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		query["Memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.NtpServers)) {
		query["NtpServers"] = request.NtpServers
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RamRoleName)) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.RestartPolicy)) {
		query["RestartPolicy"] = request.RestartPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingConfigurationId)) {
		query["ScalingConfigurationId"] = request.ScalingConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingConfigurationName)) {
		query["ScalingConfigurationName"] = request.ScalingConfigurationName
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityContextSysCtls)) {
		query["SecurityContextSysCtls"] = request.SecurityContextSysCtls
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SpotPriceLimit)) {
		query["SpotPriceLimit"] = request.SpotPriceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.SpotStrategy)) {
		query["SpotStrategy"] = request.SpotStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.TerminationGracePeriodSeconds)) {
		query["TerminationGracePeriodSeconds"] = request.TerminationGracePeriodSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.Volumes)) {
		query["Volumes"] = request.Volumes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyEciScalingConfiguration"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyEciScalingConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyEciScalingConfiguration(request *ModifyEciScalingConfigurationRequest) (_result *ModifyEciScalingConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyEciScalingConfigurationResponse{}
	_body, _err := client.ModifyEciScalingConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyLifecycleHookWithOptions(request *ModifyLifecycleHookRequest, runtime *util.RuntimeOptions) (_result *ModifyLifecycleHookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefaultResult)) {
		query["DefaultResult"] = request.DefaultResult
	}

	if !tea.BoolValue(util.IsUnset(request.HeartbeatTimeout)) {
		query["HeartbeatTimeout"] = request.HeartbeatTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleHookId)) {
		query["LifecycleHookId"] = request.LifecycleHookId
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleHookName)) {
		query["LifecycleHookName"] = request.LifecycleHookName
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleTransition)) {
		query["LifecycleTransition"] = request.LifecycleTransition
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationArn)) {
		query["NotificationArn"] = request.NotificationArn
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationMetadata)) {
		query["NotificationMetadata"] = request.NotificationMetadata
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyLifecycleHook"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyLifecycleHookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyLifecycleHook(request *ModifyLifecycleHookRequest) (_result *ModifyLifecycleHookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLifecycleHookResponse{}
	_body, _err := client.ModifyLifecycleHookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyNotificationConfigurationWithOptions(request *ModifyNotificationConfigurationRequest, runtime *util.RuntimeOptions) (_result *ModifyNotificationConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NotificationArn)) {
		query["NotificationArn"] = request.NotificationArn
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationTypes)) {
		query["NotificationTypes"] = request.NotificationTypes
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyNotificationConfiguration"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyNotificationConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyNotificationConfiguration(request *ModifyNotificationConfigurationRequest) (_result *ModifyNotificationConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNotificationConfigurationResponse{}
	_body, _err := client.ModifyNotificationConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyScalingConfigurationWithOptions(tmpReq *ModifyScalingConfigurationRequest, runtime *util.RuntimeOptions) (_result *ModifyScalingConfigurationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyScalingConfigurationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SchedulerOptions)) {
		request.SchedulerOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SchedulerOptions, tea.String("SchedulerOptions"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Affinity)) {
		query["Affinity"] = request.Affinity
	}

	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		query["Cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.CreditSpecification)) {
		query["CreditSpecification"] = request.CreditSpecification
	}

	if !tea.BoolValue(util.IsUnset(request.DataDisks)) {
		query["DataDisks"] = request.DataDisks
	}

	if !tea.BoolValue(util.IsUnset(request.DedicatedHostId)) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !tea.BoolValue(util.IsUnset(request.DeploymentSetId)) {
		query["DeploymentSetId"] = request.DeploymentSetId
	}

	if !tea.BoolValue(util.IsUnset(request.HostName)) {
		query["HostName"] = request.HostName
	}

	if !tea.BoolValue(util.IsUnset(request.HpcClusterId)) {
		query["HpcClusterId"] = request.HpcClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageFamily)) {
		query["ImageFamily"] = request.ImageFamily
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageName)) {
		query["ImageName"] = request.ImageName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceDescription)) {
		query["InstanceDescription"] = request.InstanceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.InstancePatternInfos)) {
		query["InstancePatternInfos"] = request.InstancePatternInfos
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceTypeOverrides)) {
		query["InstanceTypeOverrides"] = request.InstanceTypeOverrides
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceTypes)) {
		query["InstanceTypes"] = request.InstanceTypes
	}

	if !tea.BoolValue(util.IsUnset(request.InternetChargeType)) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.InternetMaxBandwidthOut)) {
		query["InternetMaxBandwidthOut"] = request.InternetMaxBandwidthOut
	}

	if !tea.BoolValue(util.IsUnset(request.IoOptimized)) {
		query["IoOptimized"] = request.IoOptimized
	}

	if !tea.BoolValue(util.IsUnset(request.Ipv6AddressCount)) {
		query["Ipv6AddressCount"] = request.Ipv6AddressCount
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerWeight)) {
		query["LoadBalancerWeight"] = request.LoadBalancerWeight
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		query["Memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.Override)) {
		query["Override"] = request.Override
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordInherit)) {
		query["PasswordInherit"] = request.PasswordInherit
	}

	if !tea.BoolValue(util.IsUnset(request.RamRoleName)) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingConfigurationId)) {
		query["ScalingConfigurationId"] = request.ScalingConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingConfigurationName)) {
		query["ScalingConfigurationName"] = request.ScalingConfigurationName
	}

	if !tea.BoolValue(util.IsUnset(request.SchedulerOptionsShrink)) {
		query["SchedulerOptions"] = request.SchedulerOptionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupIds)) {
		query["SecurityGroupIds"] = request.SecurityGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.SpotDuration)) {
		query["SpotDuration"] = request.SpotDuration
	}

	if !tea.BoolValue(util.IsUnset(request.SpotInterruptionBehavior)) {
		query["SpotInterruptionBehavior"] = request.SpotInterruptionBehavior
	}

	if !tea.BoolValue(util.IsUnset(request.SpotPriceLimits)) {
		query["SpotPriceLimits"] = request.SpotPriceLimits
	}

	if !tea.BoolValue(util.IsUnset(request.SpotStrategy)) {
		query["SpotStrategy"] = request.SpotStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.SystemDiskCategories)) {
		query["SystemDiskCategories"] = request.SystemDiskCategories
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.Tenancy)) {
		query["Tenancy"] = request.Tenancy
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		query["UserData"] = request.UserData
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.PrivatePoolOptions))) {
		query["PrivatePoolOptions"] = request.PrivatePoolOptions
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SystemDisk))) {
		query["SystemDisk"] = request.SystemDisk
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyScalingConfiguration"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyScalingConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyScalingConfiguration(request *ModifyScalingConfigurationRequest) (_result *ModifyScalingConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyScalingConfigurationResponse{}
	_body, _err := client.ModifyScalingConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyScalingGroupWithOptions(request *ModifyScalingGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyScalingGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActiveScalingConfigurationId)) {
		query["ActiveScalingConfigurationId"] = request.ActiveScalingConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.AllocationStrategy)) {
		query["AllocationStrategy"] = request.AllocationStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.AzBalance)) {
		query["AzBalance"] = request.AzBalance
	}

	if !tea.BoolValue(util.IsUnset(request.CompensateWithOnDemand)) {
		query["CompensateWithOnDemand"] = request.CompensateWithOnDemand
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultCooldown)) {
		query["DefaultCooldown"] = request.DefaultCooldown
	}

	if !tea.BoolValue(util.IsUnset(request.DesiredCapacity)) {
		query["DesiredCapacity"] = request.DesiredCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.GroupDeletionProtection)) {
		query["GroupDeletionProtection"] = request.GroupDeletionProtection
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckType)) {
		query["HealthCheckType"] = request.HealthCheckType
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchTemplateId)) {
		query["LaunchTemplateId"] = request.LaunchTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchTemplateOverrides)) {
		query["LaunchTemplateOverrides"] = request.LaunchTemplateOverrides
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchTemplateVersion)) {
		query["LaunchTemplateVersion"] = request.LaunchTemplateVersion
	}

	if !tea.BoolValue(util.IsUnset(request.MaxInstanceLifetime)) {
		query["MaxInstanceLifetime"] = request.MaxInstanceLifetime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSize)) {
		query["MaxSize"] = request.MaxSize
	}

	if !tea.BoolValue(util.IsUnset(request.MinSize)) {
		query["MinSize"] = request.MinSize
	}

	if !tea.BoolValue(util.IsUnset(request.MultiAZPolicy)) {
		query["MultiAZPolicy"] = request.MultiAZPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.OnDemandBaseCapacity)) {
		query["OnDemandBaseCapacity"] = request.OnDemandBaseCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.OnDemandPercentageAboveBaseCapacity)) {
		query["OnDemandPercentageAboveBaseCapacity"] = request.OnDemandPercentageAboveBaseCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RemovalPolicies)) {
		query["RemovalPolicies"] = request.RemovalPolicies
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupName)) {
		query["ScalingGroupName"] = request.ScalingGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.SpotAllocationStrategy)) {
		query["SpotAllocationStrategy"] = request.SpotAllocationStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.SpotInstancePools)) {
		query["SpotInstancePools"] = request.SpotInstancePools
	}

	if !tea.BoolValue(util.IsUnset(request.SpotInstanceRemedy)) {
		query["SpotInstanceRemedy"] = request.SpotInstanceRemedy
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchIds)) {
		query["VSwitchIds"] = request.VSwitchIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyScalingGroup"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyScalingGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyScalingGroup(request *ModifyScalingGroupRequest) (_result *ModifyScalingGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyScalingGroupResponse{}
	_body, _err := client.ModifyScalingGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyScalingRuleWithOptions(request *ModifyScalingRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyScalingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdjustmentType)) {
		query["AdjustmentType"] = request.AdjustmentType
	}

	if !tea.BoolValue(util.IsUnset(request.AdjustmentValue)) {
		query["AdjustmentValue"] = request.AdjustmentValue
	}

	if !tea.BoolValue(util.IsUnset(request.Cooldown)) {
		query["Cooldown"] = request.Cooldown
	}

	if !tea.BoolValue(util.IsUnset(request.DisableScaleIn)) {
		query["DisableScaleIn"] = request.DisableScaleIn
	}

	if !tea.BoolValue(util.IsUnset(request.EstimatedInstanceWarmup)) {
		query["EstimatedInstanceWarmup"] = request.EstimatedInstanceWarmup
	}

	if !tea.BoolValue(util.IsUnset(request.InitialMaxSize)) {
		query["InitialMaxSize"] = request.InitialMaxSize
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.MinAdjustmentMagnitude)) {
		query["MinAdjustmentMagnitude"] = request.MinAdjustmentMagnitude
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PredictiveScalingMode)) {
		query["PredictiveScalingMode"] = request.PredictiveScalingMode
	}

	if !tea.BoolValue(util.IsUnset(request.PredictiveTaskBufferTime)) {
		query["PredictiveTaskBufferTime"] = request.PredictiveTaskBufferTime
	}

	if !tea.BoolValue(util.IsUnset(request.PredictiveValueBehavior)) {
		query["PredictiveValueBehavior"] = request.PredictiveValueBehavior
	}

	if !tea.BoolValue(util.IsUnset(request.PredictiveValueBuffer)) {
		query["PredictiveValueBuffer"] = request.PredictiveValueBuffer
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleInEvaluationCount)) {
		query["ScaleInEvaluationCount"] = request.ScaleInEvaluationCount
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleOutEvaluationCount)) {
		query["ScaleOutEvaluationCount"] = request.ScaleOutEvaluationCount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleId)) {
		query["ScalingRuleId"] = request.ScalingRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleName)) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.StepAdjustments)) {
		query["StepAdjustments"] = request.StepAdjustments
	}

	if !tea.BoolValue(util.IsUnset(request.TargetValue)) {
		query["TargetValue"] = request.TargetValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyScalingRule"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyScalingRule(request *ModifyScalingRuleRequest) (_result *ModifyScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyScalingRuleResponse{}
	_body, _err := client.ModifyScalingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyScheduledTaskWithOptions(request *ModifyScheduledTaskRequest, runtime *util.RuntimeOptions) (_result *ModifyScheduledTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DesiredCapacity)) {
		query["DesiredCapacity"] = request.DesiredCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchExpirationTime)) {
		query["LaunchExpirationTime"] = request.LaunchExpirationTime
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchTime)) {
		query["LaunchTime"] = request.LaunchTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxValue)) {
		query["MaxValue"] = request.MaxValue
	}

	if !tea.BoolValue(util.IsUnset(request.MinValue)) {
		query["MinValue"] = request.MinValue
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RecurrenceEndTime)) {
		query["RecurrenceEndTime"] = request.RecurrenceEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RecurrenceType)) {
		query["RecurrenceType"] = request.RecurrenceType
	}

	if !tea.BoolValue(util.IsUnset(request.RecurrenceValue)) {
		query["RecurrenceValue"] = request.RecurrenceValue
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledAction)) {
		query["ScheduledAction"] = request.ScheduledAction
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledTaskId)) {
		query["ScheduledTaskId"] = request.ScheduledTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledTaskName)) {
		query["ScheduledTaskName"] = request.ScheduledTaskName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskEnabled)) {
		query["TaskEnabled"] = request.TaskEnabled
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyScheduledTask"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyScheduledTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyScheduledTask(request *ModifyScheduledTaskRequest) (_result *ModifyScheduledTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyScheduledTaskResponse{}
	_body, _err := client.ModifyScheduledTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebalanceInstancesWithOptions(request *RebalanceInstancesRequest, runtime *util.RuntimeOptions) (_result *RebalanceInstancesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RebalanceInstances"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RebalanceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebalanceInstances(request *RebalanceInstancesRequest) (_result *RebalanceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebalanceInstancesResponse{}
	_body, _err := client.RebalanceInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecordLifecycleActionHeartbeatWithOptions(request *RecordLifecycleActionHeartbeatRequest, runtime *util.RuntimeOptions) (_result *RecordLifecycleActionHeartbeatResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.HeartbeatTimeout)) {
		query["heartbeatTimeout"] = request.HeartbeatTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleActionToken)) {
		query["lifecycleActionToken"] = request.LifecycleActionToken
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleHookId)) {
		query["lifecycleHookId"] = request.LifecycleHookId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RecordLifecycleActionHeartbeat"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecordLifecycleActionHeartbeatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecordLifecycleActionHeartbeat(request *RecordLifecycleActionHeartbeatRequest) (_result *RecordLifecycleActionHeartbeatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecordLifecycleActionHeartbeatResponse{}
	_body, _err := client.RecordLifecycleActionHeartbeatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveInstancesWithOptions(request *RemoveInstancesRequest, runtime *util.RuntimeOptions) (_result *RemoveInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DecreaseDesiredCapacity)) {
		query["DecreaseDesiredCapacity"] = request.DecreaseDesiredCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
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

	if !tea.BoolValue(util.IsUnset(request.RemovePolicy)) {
		query["RemovePolicy"] = request.RemovePolicy
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveInstances"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveInstances(request *RemoveInstancesRequest) (_result *RemoveInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveInstancesResponse{}
	_body, _err := client.RemoveInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeProcessesWithOptions(request *ResumeProcessesRequest, runtime *util.RuntimeOptions) (_result *ResumeProcessesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Processes)) {
		query["Processes"] = request.Processes
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeProcesses"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeProcessesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeProcesses(request *ResumeProcessesRequest) (_result *ResumeProcessesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeProcessesResponse{}
	_body, _err := client.ResumeProcessesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScaleWithAdjustmentWithOptions(request *ScaleWithAdjustmentRequest, runtime *util.RuntimeOptions) (_result *ScaleWithAdjustmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdjustmentType)) {
		query["AdjustmentType"] = request.AdjustmentType
	}

	if !tea.BoolValue(util.IsUnset(request.AdjustmentValue)) {
		query["AdjustmentValue"] = request.AdjustmentValue
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.MinAdjustmentMagnitude)) {
		query["MinAdjustmentMagnitude"] = request.MinAdjustmentMagnitude
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SyncActivity)) {
		query["SyncActivity"] = request.SyncActivity
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ScaleWithAdjustment"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ScaleWithAdjustmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScaleWithAdjustment(request *ScaleWithAdjustmentRequest) (_result *ScaleWithAdjustmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScaleWithAdjustmentResponse{}
	_body, _err := client.ScaleWithAdjustmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetGroupDeletionProtectionWithOptions(request *SetGroupDeletionProtectionRequest, runtime *util.RuntimeOptions) (_result *SetGroupDeletionProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupDeletionProtection)) {
		query["GroupDeletionProtection"] = request.GroupDeletionProtection
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

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetGroupDeletionProtection"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetGroupDeletionProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetGroupDeletionProtection(request *SetGroupDeletionProtectionRequest) (_result *SetGroupDeletionProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetGroupDeletionProtectionResponse{}
	_body, _err := client.SetGroupDeletionProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetInstanceHealthWithOptions(request *SetInstanceHealthRequest, runtime *util.RuntimeOptions) (_result *SetInstanceHealthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HealthStatus)) {
		query["HealthStatus"] = request.HealthStatus
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetInstanceHealth"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetInstanceHealthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetInstanceHealth(request *SetInstanceHealthRequest) (_result *SetInstanceHealthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetInstanceHealthResponse{}
	_body, _err := client.SetInstanceHealthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetInstancesProtectionWithOptions(request *SetInstancesProtectionRequest, runtime *util.RuntimeOptions) (_result *SetInstancesProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProtectedFromScaleIn)) {
		query["ProtectedFromScaleIn"] = request.ProtectedFromScaleIn
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetInstancesProtection"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetInstancesProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetInstancesProtection(request *SetInstancesProtectionRequest) (_result *SetInstancesProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetInstancesProtectionResponse{}
	_body, _err := client.SetInstancesProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SuspendProcessesWithOptions(request *SuspendProcessesRequest, runtime *util.RuntimeOptions) (_result *SuspendProcessesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Processes)) {
		query["Processes"] = request.Processes
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SuspendProcesses"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SuspendProcessesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SuspendProcesses(request *SuspendProcessesRequest) (_result *SuspendProcessesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SuspendProcessesResponse{}
	_body, _err := client.SuspendProcessesWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2022-02-22"),
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

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeys)) {
		query["TagKeys"] = request.TagKeys
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2022-02-22"),
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

func (client *Client) VerifyAuthenticationWithOptions(request *VerifyAuthenticationRequest, runtime *util.RuntimeOptions) (_result *VerifyAuthenticationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OnlyCheck)) {
		query["OnlyCheck"] = request.OnlyCheck
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyAuthentication"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyAuthenticationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyAuthentication(request *VerifyAuthenticationRequest) (_result *VerifyAuthenticationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyAuthenticationResponse{}
	_body, _err := client.VerifyAuthenticationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyUserWithOptions(request *VerifyUserRequest, runtime *util.RuntimeOptions) (_result *VerifyUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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
		Action:      tea.String("VerifyUser"),
		Version:     tea.String("2022-02-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyUser(request *VerifyUserRequest) (_result *VerifyUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyUserResponse{}
	_body, _err := client.VerifyUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
