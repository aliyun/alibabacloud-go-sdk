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

type AttachAlbServerGroupsRequest struct {
	OwnerId              *int64                                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	RegionId             *string                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScalingGroupId       *string                                       `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ClientToken          *string                                       `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ForceAttach          *bool                                         `json:"ForceAttach,omitempty" xml:"ForceAttach,omitempty"`
	AlbServerGroup       []*AttachAlbServerGroupsRequestAlbServerGroup `json:"AlbServerGroup,omitempty" xml:"AlbServerGroup,omitempty" type:"Repeated"`
}

func (s AttachAlbServerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachAlbServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *AttachAlbServerGroupsRequest) SetOwnerId(v int64) *AttachAlbServerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachAlbServerGroupsRequest) SetResourceOwnerAccount(v string) *AttachAlbServerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachAlbServerGroupsRequest) SetRegionId(v string) *AttachAlbServerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *AttachAlbServerGroupsRequest) SetScalingGroupId(v string) *AttachAlbServerGroupsRequest {
	s.ScalingGroupId = &v
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

func (s *AttachAlbServerGroupsRequest) SetAlbServerGroup(v []*AttachAlbServerGroupsRequestAlbServerGroup) *AttachAlbServerGroupsRequest {
	s.AlbServerGroup = v
	return s
}

type AttachAlbServerGroupsRequestAlbServerGroup struct {
	AlbServerGroupId *string `json:"AlbServerGroupId,omitempty" xml:"AlbServerGroupId,omitempty"`
	Weight           *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s AttachAlbServerGroupsRequestAlbServerGroup) String() string {
	return tea.Prettify(s)
}

func (s AttachAlbServerGroupsRequestAlbServerGroup) GoString() string {
	return s.String()
}

func (s *AttachAlbServerGroupsRequestAlbServerGroup) SetAlbServerGroupId(v string) *AttachAlbServerGroupsRequestAlbServerGroup {
	s.AlbServerGroupId = &v
	return s
}

func (s *AttachAlbServerGroupsRequestAlbServerGroup) SetWeight(v int32) *AttachAlbServerGroupsRequestAlbServerGroup {
	s.Weight = &v
	return s
}

func (s *AttachAlbServerGroupsRequestAlbServerGroup) SetPort(v int32) *AttachAlbServerGroupsRequestAlbServerGroup {
	s.Port = &v
	return s
}

type AttachAlbServerGroupsResponseBody struct {
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachAlbServerGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachAlbServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *AttachAlbServerGroupsResponseBody) SetScalingActivityId(v string) *AttachAlbServerGroupsResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *AttachAlbServerGroupsResponseBody) SetRequestId(v string) *AttachAlbServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

type AttachAlbServerGroupsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachAlbServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AttachAlbServerGroupsResponse) SetBody(v *AttachAlbServerGroupsResponseBody) *AttachAlbServerGroupsResponse {
	s.Body = v
	return s
}

type AttachDBInstancesRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ForceAttach          *bool     `json:"ForceAttach,omitempty" xml:"ForceAttach,omitempty"`
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstance           []*string `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Repeated"`
}

func (s AttachDBInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *AttachDBInstancesRequest) SetOwnerId(v int64) *AttachDBInstancesRequest {
	s.OwnerId = &v
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

func (s *AttachDBInstancesRequest) SetForceAttach(v bool) *AttachDBInstancesRequest {
	s.ForceAttach = &v
	return s
}

func (s *AttachDBInstancesRequest) SetClientToken(v string) *AttachDBInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachDBInstancesRequest) SetDBInstance(v []*string) *AttachDBInstancesRequest {
	s.DBInstance = v
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachDBInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AttachDBInstancesResponse) SetBody(v *AttachDBInstancesResponseBody) *AttachDBInstancesResponse {
	s.Body = v
	return s
}

type AttachInstancesRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	Entrusted            *bool     `json:"Entrusted,omitempty" xml:"Entrusted,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	InstanceId           []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	LoadBalancerWeight   []*int    `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty" type:"Repeated"`
}

func (s AttachInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesRequest) GoString() string {
	return s.String()
}

func (s *AttachInstancesRequest) SetOwnerId(v int64) *AttachInstancesRequest {
	s.OwnerId = &v
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

func (s *AttachInstancesRequest) SetEntrusted(v bool) *AttachInstancesRequest {
	s.Entrusted = &v
	return s
}

func (s *AttachInstancesRequest) SetOwnerAccount(v string) *AttachInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AttachInstancesRequest) SetInstanceId(v []*string) *AttachInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *AttachInstancesRequest) SetLoadBalancerWeight(v []*int) *AttachInstancesRequest {
	s.LoadBalancerWeight = v
	return s
}

type AttachInstancesResponseBody struct {
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *AttachInstancesResponseBody) SetScalingActivityId(v string) *AttachInstancesResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *AttachInstancesResponseBody) SetRequestId(v string) *AttachInstancesResponseBody {
	s.RequestId = &v
	return s
}

type AttachInstancesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AttachInstancesResponse) SetBody(v *AttachInstancesResponseBody) *AttachInstancesResponse {
	s.Body = v
	return s
}

type AttachLoadBalancersRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ForceAttach          *bool     `json:"ForceAttach,omitempty" xml:"ForceAttach,omitempty"`
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Async                *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
	LoadBalancer         []*string `json:"LoadBalancer,omitempty" xml:"LoadBalancer,omitempty" type:"Repeated"`
}

func (s AttachLoadBalancersRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachLoadBalancersRequest) GoString() string {
	return s.String()
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

func (s *AttachLoadBalancersRequest) SetForceAttach(v bool) *AttachLoadBalancersRequest {
	s.ForceAttach = &v
	return s
}

func (s *AttachLoadBalancersRequest) SetClientToken(v string) *AttachLoadBalancersRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachLoadBalancersRequest) SetAsync(v bool) *AttachLoadBalancersRequest {
	s.Async = &v
	return s
}

func (s *AttachLoadBalancersRequest) SetLoadBalancer(v []*string) *AttachLoadBalancersRequest {
	s.LoadBalancer = v
	return s
}

type AttachLoadBalancersResponseBody struct {
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachLoadBalancersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachLoadBalancersResponseBody) GoString() string {
	return s.String()
}

func (s *AttachLoadBalancersResponseBody) SetScalingActivityId(v string) *AttachLoadBalancersResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *AttachLoadBalancersResponseBody) SetRequestId(v string) *AttachLoadBalancersResponseBody {
	s.RequestId = &v
	return s
}

type AttachLoadBalancersResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachLoadBalancersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AttachLoadBalancersResponse) SetBody(v *AttachLoadBalancersResponseBody) *AttachLoadBalancersResponse {
	s.Body = v
	return s
}

type AttachVServerGroupsRequest struct {
	OwnerId              *int64                                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	RegionId             *string                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScalingGroupId       *string                                   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ClientToken          *string                                   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ForceAttach          *bool                                     `json:"ForceAttach,omitempty" xml:"ForceAttach,omitempty"`
	VServerGroup         []*AttachVServerGroupsRequestVServerGroup `json:"VServerGroup,omitempty" xml:"VServerGroup,omitempty" type:"Repeated"`
}

func (s AttachVServerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachVServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *AttachVServerGroupsRequest) SetOwnerId(v int64) *AttachVServerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachVServerGroupsRequest) SetResourceOwnerAccount(v string) *AttachVServerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachVServerGroupsRequest) SetRegionId(v string) *AttachVServerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *AttachVServerGroupsRequest) SetScalingGroupId(v string) *AttachVServerGroupsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *AttachVServerGroupsRequest) SetClientToken(v string) *AttachVServerGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachVServerGroupsRequest) SetForceAttach(v bool) *AttachVServerGroupsRequest {
	s.ForceAttach = &v
	return s
}

func (s *AttachVServerGroupsRequest) SetVServerGroup(v []*AttachVServerGroupsRequestVServerGroup) *AttachVServerGroupsRequest {
	s.VServerGroup = v
	return s
}

type AttachVServerGroupsRequestVServerGroup struct {
	VServerGroupAttribute []*AttachVServerGroupsRequestVServerGroupVServerGroupAttribute `json:"VServerGroupAttribute,omitempty" xml:"VServerGroupAttribute,omitempty" type:"Repeated"`
	LoadBalancerId        *string                                                        `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s AttachVServerGroupsRequestVServerGroup) String() string {
	return tea.Prettify(s)
}

func (s AttachVServerGroupsRequestVServerGroup) GoString() string {
	return s.String()
}

func (s *AttachVServerGroupsRequestVServerGroup) SetVServerGroupAttribute(v []*AttachVServerGroupsRequestVServerGroupVServerGroupAttribute) *AttachVServerGroupsRequestVServerGroup {
	s.VServerGroupAttribute = v
	return s
}

func (s *AttachVServerGroupsRequestVServerGroup) SetLoadBalancerId(v string) *AttachVServerGroupsRequestVServerGroup {
	s.LoadBalancerId = &v
	return s
}

type AttachVServerGroupsRequestVServerGroupVServerGroupAttribute struct {
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	Weight         *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
	Port           *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s AttachVServerGroupsRequestVServerGroupVServerGroupAttribute) String() string {
	return tea.Prettify(s)
}

func (s AttachVServerGroupsRequestVServerGroupVServerGroupAttribute) GoString() string {
	return s.String()
}

func (s *AttachVServerGroupsRequestVServerGroupVServerGroupAttribute) SetVServerGroupId(v string) *AttachVServerGroupsRequestVServerGroupVServerGroupAttribute {
	s.VServerGroupId = &v
	return s
}

func (s *AttachVServerGroupsRequestVServerGroupVServerGroupAttribute) SetWeight(v int32) *AttachVServerGroupsRequestVServerGroupVServerGroupAttribute {
	s.Weight = &v
	return s
}

func (s *AttachVServerGroupsRequestVServerGroupVServerGroupAttribute) SetPort(v int32) *AttachVServerGroupsRequestVServerGroupVServerGroupAttribute {
	s.Port = &v
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachVServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AttachVServerGroupsResponse) SetBody(v *AttachVServerGroupsResponseBody) *AttachVServerGroupsResponse {
	s.Body = v
	return s
}

type CompleteLifecycleActionRequest struct {
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	LifecycleHookId       *string `json:"LifecycleHookId,omitempty" xml:"LifecycleHookId,omitempty"`
	LifecycleActionToken  *string `json:"LifecycleActionToken,omitempty" xml:"LifecycleActionToken,omitempty"`
	LifecycleActionResult *string `json:"LifecycleActionResult,omitempty" xml:"LifecycleActionResult,omitempty"`
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CompleteLifecycleActionRequest) String() string {
	return tea.Prettify(s)
}

func (s CompleteLifecycleActionRequest) GoString() string {
	return s.String()
}

func (s *CompleteLifecycleActionRequest) SetOwnerId(v int64) *CompleteLifecycleActionRequest {
	s.OwnerId = &v
	return s
}

func (s *CompleteLifecycleActionRequest) SetResourceOwnerAccount(v string) *CompleteLifecycleActionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CompleteLifecycleActionRequest) SetOwnerAccount(v string) *CompleteLifecycleActionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CompleteLifecycleActionRequest) SetLifecycleHookId(v string) *CompleteLifecycleActionRequest {
	s.LifecycleHookId = &v
	return s
}

func (s *CompleteLifecycleActionRequest) SetLifecycleActionToken(v string) *CompleteLifecycleActionRequest {
	s.LifecycleActionToken = &v
	return s
}

func (s *CompleteLifecycleActionRequest) SetLifecycleActionResult(v string) *CompleteLifecycleActionRequest {
	s.LifecycleActionResult = &v
	return s
}

func (s *CompleteLifecycleActionRequest) SetClientToken(v string) *CompleteLifecycleActionRequest {
	s.ClientToken = &v
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CompleteLifecycleActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CompleteLifecycleActionResponse) SetBody(v *CompleteLifecycleActionResponseBody) *CompleteLifecycleActionResponse {
	s.Body = v
	return s
}

type CreateAlarmRequest struct {
	OwnerId              *int64                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                        `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	RegionId             *string                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name                 *string                        `json:"Name,omitempty" xml:"Name,omitempty"`
	Description          *string                        `json:"Description,omitempty" xml:"Description,omitempty"`
	ScalingGroupId       *string                        `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	MetricName           *string                        `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	MetricType           *string                        `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Period               *int32                         `json:"Period,omitempty" xml:"Period,omitempty"`
	Statistics           *string                        `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold            *float32                       `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	ComparisonOperator   *string                        `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	EvaluationCount      *int32                         `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	GroupId              *int32                         `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Effective            *string                        `json:"Effective,omitempty" xml:"Effective,omitempty"`
	AlarmAction          []*string                      `json:"AlarmAction,omitempty" xml:"AlarmAction,omitempty" type:"Repeated"`
	Dimension            []*CreateAlarmRequestDimension `json:"Dimension,omitempty" xml:"Dimension,omitempty" type:"Repeated"`
}

func (s CreateAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmRequest) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequest) SetOwnerId(v int64) *CreateAlarmRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAlarmRequest) SetResourceOwnerAccount(v string) *CreateAlarmRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAlarmRequest) SetRegionId(v string) *CreateAlarmRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAlarmRequest) SetName(v string) *CreateAlarmRequest {
	s.Name = &v
	return s
}

func (s *CreateAlarmRequest) SetDescription(v string) *CreateAlarmRequest {
	s.Description = &v
	return s
}

func (s *CreateAlarmRequest) SetScalingGroupId(v string) *CreateAlarmRequest {
	s.ScalingGroupId = &v
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

func (s *CreateAlarmRequest) SetPeriod(v int32) *CreateAlarmRequest {
	s.Period = &v
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

func (s *CreateAlarmRequest) SetComparisonOperator(v string) *CreateAlarmRequest {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateAlarmRequest) SetEvaluationCount(v int32) *CreateAlarmRequest {
	s.EvaluationCount = &v
	return s
}

func (s *CreateAlarmRequest) SetGroupId(v int32) *CreateAlarmRequest {
	s.GroupId = &v
	return s
}

func (s *CreateAlarmRequest) SetEffective(v string) *CreateAlarmRequest {
	s.Effective = &v
	return s
}

func (s *CreateAlarmRequest) SetAlarmAction(v []*string) *CreateAlarmRequest {
	s.AlarmAction = v
	return s
}

func (s *CreateAlarmRequest) SetDimension(v []*CreateAlarmRequestDimension) *CreateAlarmRequest {
	s.Dimension = v
	return s
}

type CreateAlarmRequestDimension struct {
	DimensionKey   *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
}

func (s CreateAlarmRequestDimension) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmRequestDimension) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestDimension) SetDimensionKey(v string) *CreateAlarmRequestDimension {
	s.DimensionKey = &v
	return s
}

func (s *CreateAlarmRequestDimension) SetDimensionValue(v string) *CreateAlarmRequestDimension {
	s.DimensionValue = &v
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateAlarmResponse) SetBody(v *CreateAlarmResponseBody) *CreateAlarmResponse {
	s.Body = v
	return s
}

type CreateLifecycleHookRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	LifecycleHookName    *string `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	LifecycleTransition  *string `json:"LifecycleTransition,omitempty" xml:"LifecycleTransition,omitempty"`
	DefaultResult        *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	HeartbeatTimeout     *int32  `json:"HeartbeatTimeout,omitempty" xml:"HeartbeatTimeout,omitempty"`
	NotificationMetadata *string `json:"NotificationMetadata,omitempty" xml:"NotificationMetadata,omitempty"`
	NotificationArn      *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
}

func (s CreateLifecycleHookRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLifecycleHookRequest) GoString() string {
	return s.String()
}

func (s *CreateLifecycleHookRequest) SetOwnerId(v int64) *CreateLifecycleHookRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetResourceOwnerAccount(v string) *CreateLifecycleHookRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetOwnerAccount(v string) *CreateLifecycleHookRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetScalingGroupId(v string) *CreateLifecycleHookRequest {
	s.ScalingGroupId = &v
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

func (s *CreateLifecycleHookRequest) SetDefaultResult(v string) *CreateLifecycleHookRequest {
	s.DefaultResult = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetHeartbeatTimeout(v int32) *CreateLifecycleHookRequest {
	s.HeartbeatTimeout = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetNotificationMetadata(v string) *CreateLifecycleHookRequest {
	s.NotificationMetadata = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetNotificationArn(v string) *CreateLifecycleHookRequest {
	s.NotificationArn = &v
	return s
}

type CreateLifecycleHookResponseBody struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LifecycleHookId *string `json:"LifecycleHookId,omitempty" xml:"LifecycleHookId,omitempty"`
}

func (s CreateLifecycleHookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLifecycleHookResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLifecycleHookResponseBody) SetRequestId(v string) *CreateLifecycleHookResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLifecycleHookResponseBody) SetLifecycleHookId(v string) *CreateLifecycleHookResponseBody {
	s.LifecycleHookId = &v
	return s
}

type CreateLifecycleHookResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLifecycleHookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateLifecycleHookResponse) SetBody(v *CreateLifecycleHookResponseBody) *CreateLifecycleHookResponse {
	s.Body = v
	return s
}

type CreateNotificationConfigurationRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	NotificationArn      *string   `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	NotificationType     []*string `json:"NotificationType,omitempty" xml:"NotificationType,omitempty" type:"Repeated"`
}

func (s CreateNotificationConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNotificationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateNotificationConfigurationRequest) SetOwnerId(v int64) *CreateNotificationConfigurationRequest {
	s.OwnerId = &v
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

func (s *CreateNotificationConfigurationRequest) SetNotificationArn(v string) *CreateNotificationConfigurationRequest {
	s.NotificationArn = &v
	return s
}

func (s *CreateNotificationConfigurationRequest) SetNotificationType(v []*string) *CreateNotificationConfigurationRequest {
	s.NotificationType = v
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
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateNotificationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateNotificationConfigurationResponse) SetBody(v *CreateNotificationConfigurationResponseBody) *CreateNotificationConfigurationResponse {
	s.Body = v
	return s
}

type CreateScalingConfigurationRequest struct {
	SystemDisk                  *CreateScalingConfigurationRequestSystemDisk             `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	PrivatePoolOptions          *CreateScalingConfigurationRequestPrivatePoolOptions     `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	OwnerId                     *int64                                                   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount        *string                                                  `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId              *string                                                  `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ImageId                     *string                                                  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName                   *string                                                  `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	InstanceType                *string                                                  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Cpu                         *int32                                                   `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Memory                      *int32                                                   `json:"Memory,omitempty" xml:"Memory,omitempty"`
	DeploymentSetId             *string                                                  `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	SecurityGroupId             *string                                                  `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	IoOptimized                 *string                                                  `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	InternetChargeType          *string                                                  `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandwidthIn      *int32                                                   `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	InternetMaxBandwidthOut     *int32                                                   `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	ScalingConfigurationName    *string                                                  `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	LoadBalancerWeight          *int32                                                   `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	OwnerAccount                *string                                                  `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	Tags                        *string                                                  `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData                    *string                                                  `json:"UserData,omitempty" xml:"UserData,omitempty"`
	KeyPairName                 *string                                                  `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	RamRoleName                 *string                                                  `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	SecurityEnhancementStrategy *string                                                  `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	InstanceName                *string                                                  `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	HostName                    *string                                                  `json:"HostName,omitempty" xml:"HostName,omitempty"`
	SpotStrategy                *string                                                  `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	PasswordInherit             *bool                                                    `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	Password                    *string                                                  `json:"Password,omitempty" xml:"Password,omitempty"`
	ResourceGroupId             *string                                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	HpcClusterId                *string                                                  `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	InstanceDescription         *string                                                  `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	ClientToken                 *string                                                  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Ipv6AddressCount            *int32                                                   `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	CreditSpecification         *string                                                  `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	ImageFamily                 *string                                                  `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	ZoneId                      *string                                                  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DedicatedHostId             *string                                                  `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	Affinity                    *string                                                  `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	Tenancy                     *string                                                  `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	SchedulerOptions            map[string]interface{}                                   `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty"`
	SpotDuration                *int32                                                   `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	SpotInterruptionBehavior    *string                                                  `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	InstanceTypes               []*string                                                `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	InstanceTypeOverride        []*CreateScalingConfigurationRequestInstanceTypeOverride `json:"InstanceTypeOverride,omitempty" xml:"InstanceTypeOverride,omitempty" type:"Repeated"`
	DataDisk                    []*CreateScalingConfigurationRequestDataDisk             `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	SpotPriceLimit              []*CreateScalingConfigurationRequestSpotPriceLimit       `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty" type:"Repeated"`
	SecurityGroupIds            []*string                                                `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	InstancePatternInfo         []*CreateScalingConfigurationRequestInstancePatternInfo  `json:"InstancePatternInfo,omitempty" xml:"InstancePatternInfo,omitempty" type:"Repeated"`
	SystemDiskCategory          []*string                                                `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty" type:"Repeated"`
}

func (s CreateScalingConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequest) SetSystemDisk(v *CreateScalingConfigurationRequestSystemDisk) *CreateScalingConfigurationRequest {
	s.SystemDisk = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetPrivatePoolOptions(v *CreateScalingConfigurationRequestPrivatePoolOptions) *CreateScalingConfigurationRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetOwnerId(v int64) *CreateScalingConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetResourceOwnerAccount(v string) *CreateScalingConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetScalingGroupId(v string) *CreateScalingConfigurationRequest {
	s.ScalingGroupId = &v
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

func (s *CreateScalingConfigurationRequest) SetInstanceType(v string) *CreateScalingConfigurationRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetCpu(v int32) *CreateScalingConfigurationRequest {
	s.Cpu = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetMemory(v int32) *CreateScalingConfigurationRequest {
	s.Memory = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetDeploymentSetId(v string) *CreateScalingConfigurationRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSecurityGroupId(v string) *CreateScalingConfigurationRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetIoOptimized(v string) *CreateScalingConfigurationRequest {
	s.IoOptimized = &v
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

func (s *CreateScalingConfigurationRequest) SetScalingConfigurationName(v string) *CreateScalingConfigurationRequest {
	s.ScalingConfigurationName = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetLoadBalancerWeight(v int32) *CreateScalingConfigurationRequest {
	s.LoadBalancerWeight = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetOwnerAccount(v string) *CreateScalingConfigurationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetTags(v string) *CreateScalingConfigurationRequest {
	s.Tags = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetUserData(v string) *CreateScalingConfigurationRequest {
	s.UserData = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetKeyPairName(v string) *CreateScalingConfigurationRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetRamRoleName(v string) *CreateScalingConfigurationRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSecurityEnhancementStrategy(v string) *CreateScalingConfigurationRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInstanceName(v string) *CreateScalingConfigurationRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetHostName(v string) *CreateScalingConfigurationRequest {
	s.HostName = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSpotStrategy(v string) *CreateScalingConfigurationRequest {
	s.SpotStrategy = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetPasswordInherit(v bool) *CreateScalingConfigurationRequest {
	s.PasswordInherit = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetPassword(v string) *CreateScalingConfigurationRequest {
	s.Password = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetResourceGroupId(v string) *CreateScalingConfigurationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetHpcClusterId(v string) *CreateScalingConfigurationRequest {
	s.HpcClusterId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInstanceDescription(v string) *CreateScalingConfigurationRequest {
	s.InstanceDescription = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetClientToken(v string) *CreateScalingConfigurationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetIpv6AddressCount(v int32) *CreateScalingConfigurationRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetCreditSpecification(v string) *CreateScalingConfigurationRequest {
	s.CreditSpecification = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetImageFamily(v string) *CreateScalingConfigurationRequest {
	s.ImageFamily = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetZoneId(v string) *CreateScalingConfigurationRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetDedicatedHostId(v string) *CreateScalingConfigurationRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetAffinity(v string) *CreateScalingConfigurationRequest {
	s.Affinity = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetTenancy(v string) *CreateScalingConfigurationRequest {
	s.Tenancy = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSchedulerOptions(v map[string]interface{}) *CreateScalingConfigurationRequest {
	s.SchedulerOptions = v
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

func (s *CreateScalingConfigurationRequest) SetInstanceTypes(v []*string) *CreateScalingConfigurationRequest {
	s.InstanceTypes = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInstanceTypeOverride(v []*CreateScalingConfigurationRequestInstanceTypeOverride) *CreateScalingConfigurationRequest {
	s.InstanceTypeOverride = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetDataDisk(v []*CreateScalingConfigurationRequestDataDisk) *CreateScalingConfigurationRequest {
	s.DataDisk = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSpotPriceLimit(v []*CreateScalingConfigurationRequestSpotPriceLimit) *CreateScalingConfigurationRequest {
	s.SpotPriceLimit = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSecurityGroupIds(v []*string) *CreateScalingConfigurationRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInstancePatternInfo(v []*CreateScalingConfigurationRequestInstancePatternInfo) *CreateScalingConfigurationRequest {
	s.InstancePatternInfo = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSystemDiskCategory(v []*string) *CreateScalingConfigurationRequest {
	s.SystemDiskCategory = v
	return s
}

type CreateScalingConfigurationRequestSystemDisk struct {
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Size                 *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	DiskName             *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	PerformanceLevel     *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
}

func (s CreateScalingConfigurationRequestSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetCategory(v string) *CreateScalingConfigurationRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetSize(v int32) *CreateScalingConfigurationRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetDiskName(v string) *CreateScalingConfigurationRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetDescription(v string) *CreateScalingConfigurationRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetAutoSnapshotPolicyId(v string) *CreateScalingConfigurationRequestSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetPerformanceLevel(v string) *CreateScalingConfigurationRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

type CreateScalingConfigurationRequestPrivatePoolOptions struct {
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateScalingConfigurationRequestPrivatePoolOptions) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestPrivatePoolOptions) SetMatchCriteria(v string) *CreateScalingConfigurationRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *CreateScalingConfigurationRequestPrivatePoolOptions) SetId(v string) *CreateScalingConfigurationRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

type CreateScalingConfigurationRequestInstanceTypeOverride struct {
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	WeightedCapacity *int32  `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s CreateScalingConfigurationRequestInstanceTypeOverride) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationRequestInstanceTypeOverride) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestInstanceTypeOverride) SetInstanceType(v string) *CreateScalingConfigurationRequestInstanceTypeOverride {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstanceTypeOverride) SetWeightedCapacity(v int32) *CreateScalingConfigurationRequestInstanceTypeOverride {
	s.WeightedCapacity = &v
	return s
}

type CreateScalingConfigurationRequestDataDisk struct {
	Categorys            []*string `json:"Categorys,omitempty" xml:"Categorys,omitempty" type:"Repeated"`
	PerformanceLevel     *string   `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	AutoSnapshotPolicyId *string   `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	Encrypted            *string   `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	Description          *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	SnapshotId           *string   `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	Size                 *int32    `json:"Size,omitempty" xml:"Size,omitempty"`
	Device               *string   `json:"Device,omitempty" xml:"Device,omitempty"`
	DiskName             *string   `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	Category             *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteWithInstance   *bool     `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	KMSKeyId             *string   `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
}

func (s CreateScalingConfigurationRequestDataDisk) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationRequestDataDisk) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestDataDisk) SetCategorys(v []*string) *CreateScalingConfigurationRequestDataDisk {
	s.Categorys = v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisk) SetPerformanceLevel(v string) *CreateScalingConfigurationRequestDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisk) SetAutoSnapshotPolicyId(v string) *CreateScalingConfigurationRequestDataDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisk) SetEncrypted(v string) *CreateScalingConfigurationRequestDataDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisk) SetDescription(v string) *CreateScalingConfigurationRequestDataDisk {
	s.Description = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisk) SetSnapshotId(v string) *CreateScalingConfigurationRequestDataDisk {
	s.SnapshotId = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisk) SetSize(v int32) *CreateScalingConfigurationRequestDataDisk {
	s.Size = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisk) SetDevice(v string) *CreateScalingConfigurationRequestDataDisk {
	s.Device = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisk) SetDiskName(v string) *CreateScalingConfigurationRequestDataDisk {
	s.DiskName = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisk) SetCategory(v string) *CreateScalingConfigurationRequestDataDisk {
	s.Category = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisk) SetDeleteWithInstance(v bool) *CreateScalingConfigurationRequestDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisk) SetKMSKeyId(v string) *CreateScalingConfigurationRequestDataDisk {
	s.KMSKeyId = &v
	return s
}

type CreateScalingConfigurationRequestSpotPriceLimit struct {
	PriceLimit   *float32 `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
	InstanceType *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateScalingConfigurationRequestSpotPriceLimit) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationRequestSpotPriceLimit) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestSpotPriceLimit) SetPriceLimit(v float32) *CreateScalingConfigurationRequestSpotPriceLimit {
	s.PriceLimit = &v
	return s
}

func (s *CreateScalingConfigurationRequestSpotPriceLimit) SetInstanceType(v string) *CreateScalingConfigurationRequestSpotPriceLimit {
	s.InstanceType = &v
	return s
}

type CreateScalingConfigurationRequestInstancePatternInfo struct {
	Cores               *int32   `json:"Cores,omitempty" xml:"Cores,omitempty"`
	InstanceFamilyLevel *string  `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	MaxPrice            *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	Memory              *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s CreateScalingConfigurationRequestInstancePatternInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationRequestInstancePatternInfo) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestInstancePatternInfo) SetCores(v int32) *CreateScalingConfigurationRequestInstancePatternInfo {
	s.Cores = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfo) SetInstanceFamilyLevel(v string) *CreateScalingConfigurationRequestInstancePatternInfo {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfo) SetMaxPrice(v float32) *CreateScalingConfigurationRequestInstancePatternInfo {
	s.MaxPrice = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfo) SetMemory(v float32) *CreateScalingConfigurationRequestInstancePatternInfo {
	s.Memory = &v
	return s
}

type CreateScalingConfigurationShrinkRequest struct {
	SystemDisk                  *CreateScalingConfigurationShrinkRequestSystemDisk             `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	PrivatePoolOptions          *CreateScalingConfigurationShrinkRequestPrivatePoolOptions     `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	OwnerId                     *int64                                                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount        *string                                                        `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId              *string                                                        `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ImageId                     *string                                                        `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName                   *string                                                        `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	InstanceType                *string                                                        `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Cpu                         *int32                                                         `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Memory                      *int32                                                         `json:"Memory,omitempty" xml:"Memory,omitempty"`
	DeploymentSetId             *string                                                        `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	SecurityGroupId             *string                                                        `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	IoOptimized                 *string                                                        `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	InternetChargeType          *string                                                        `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandwidthIn      *int32                                                         `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	InternetMaxBandwidthOut     *int32                                                         `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	ScalingConfigurationName    *string                                                        `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	LoadBalancerWeight          *int32                                                         `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	OwnerAccount                *string                                                        `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	Tags                        *string                                                        `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserData                    *string                                                        `json:"UserData,omitempty" xml:"UserData,omitempty"`
	KeyPairName                 *string                                                        `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	RamRoleName                 *string                                                        `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	SecurityEnhancementStrategy *string                                                        `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	InstanceName                *string                                                        `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	HostName                    *string                                                        `json:"HostName,omitempty" xml:"HostName,omitempty"`
	SpotStrategy                *string                                                        `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	PasswordInherit             *bool                                                          `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	Password                    *string                                                        `json:"Password,omitempty" xml:"Password,omitempty"`
	ResourceGroupId             *string                                                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	HpcClusterId                *string                                                        `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	InstanceDescription         *string                                                        `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	ClientToken                 *string                                                        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Ipv6AddressCount            *int32                                                         `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	CreditSpecification         *string                                                        `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	ImageFamily                 *string                                                        `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	ZoneId                      *string                                                        `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DedicatedHostId             *string                                                        `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	Affinity                    *string                                                        `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	Tenancy                     *string                                                        `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	SchedulerOptionsShrink      *string                                                        `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty"`
	SpotDuration                *int32                                                         `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	SpotInterruptionBehavior    *string                                                        `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	InstanceTypes               []*string                                                      `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	InstanceTypeOverride        []*CreateScalingConfigurationShrinkRequestInstanceTypeOverride `json:"InstanceTypeOverride,omitempty" xml:"InstanceTypeOverride,omitempty" type:"Repeated"`
	DataDisk                    []*CreateScalingConfigurationShrinkRequestDataDisk             `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	SpotPriceLimit              []*CreateScalingConfigurationShrinkRequestSpotPriceLimit       `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty" type:"Repeated"`
	SecurityGroupIds            []*string                                                      `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	InstancePatternInfo         []*CreateScalingConfigurationShrinkRequestInstancePatternInfo  `json:"InstancePatternInfo,omitempty" xml:"InstancePatternInfo,omitempty" type:"Repeated"`
	SystemDiskCategory          []*string                                                      `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty" type:"Repeated"`
}

func (s CreateScalingConfigurationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequest) SetSystemDisk(v *CreateScalingConfigurationShrinkRequestSystemDisk) *CreateScalingConfigurationShrinkRequest {
	s.SystemDisk = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetPrivatePoolOptions(v *CreateScalingConfigurationShrinkRequestPrivatePoolOptions) *CreateScalingConfigurationShrinkRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetOwnerId(v int64) *CreateScalingConfigurationShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetResourceOwnerAccount(v string) *CreateScalingConfigurationShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetScalingGroupId(v string) *CreateScalingConfigurationShrinkRequest {
	s.ScalingGroupId = &v
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

func (s *CreateScalingConfigurationShrinkRequest) SetInstanceType(v string) *CreateScalingConfigurationShrinkRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetCpu(v int32) *CreateScalingConfigurationShrinkRequest {
	s.Cpu = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetMemory(v int32) *CreateScalingConfigurationShrinkRequest {
	s.Memory = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetDeploymentSetId(v string) *CreateScalingConfigurationShrinkRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSecurityGroupId(v string) *CreateScalingConfigurationShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetIoOptimized(v string) *CreateScalingConfigurationShrinkRequest {
	s.IoOptimized = &v
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

func (s *CreateScalingConfigurationShrinkRequest) SetScalingConfigurationName(v string) *CreateScalingConfigurationShrinkRequest {
	s.ScalingConfigurationName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetLoadBalancerWeight(v int32) *CreateScalingConfigurationShrinkRequest {
	s.LoadBalancerWeight = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetOwnerAccount(v string) *CreateScalingConfigurationShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetTags(v string) *CreateScalingConfigurationShrinkRequest {
	s.Tags = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetUserData(v string) *CreateScalingConfigurationShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetKeyPairName(v string) *CreateScalingConfigurationShrinkRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetRamRoleName(v string) *CreateScalingConfigurationShrinkRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSecurityEnhancementStrategy(v string) *CreateScalingConfigurationShrinkRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInstanceName(v string) *CreateScalingConfigurationShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetHostName(v string) *CreateScalingConfigurationShrinkRequest {
	s.HostName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSpotStrategy(v string) *CreateScalingConfigurationShrinkRequest {
	s.SpotStrategy = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetPasswordInherit(v bool) *CreateScalingConfigurationShrinkRequest {
	s.PasswordInherit = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetPassword(v string) *CreateScalingConfigurationShrinkRequest {
	s.Password = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetResourceGroupId(v string) *CreateScalingConfigurationShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetHpcClusterId(v string) *CreateScalingConfigurationShrinkRequest {
	s.HpcClusterId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInstanceDescription(v string) *CreateScalingConfigurationShrinkRequest {
	s.InstanceDescription = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetClientToken(v string) *CreateScalingConfigurationShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetIpv6AddressCount(v int32) *CreateScalingConfigurationShrinkRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetCreditSpecification(v string) *CreateScalingConfigurationShrinkRequest {
	s.CreditSpecification = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetImageFamily(v string) *CreateScalingConfigurationShrinkRequest {
	s.ImageFamily = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetZoneId(v string) *CreateScalingConfigurationShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetDedicatedHostId(v string) *CreateScalingConfigurationShrinkRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetAffinity(v string) *CreateScalingConfigurationShrinkRequest {
	s.Affinity = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetTenancy(v string) *CreateScalingConfigurationShrinkRequest {
	s.Tenancy = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSchedulerOptionsShrink(v string) *CreateScalingConfigurationShrinkRequest {
	s.SchedulerOptionsShrink = &v
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

func (s *CreateScalingConfigurationShrinkRequest) SetInstanceTypes(v []*string) *CreateScalingConfigurationShrinkRequest {
	s.InstanceTypes = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInstanceTypeOverride(v []*CreateScalingConfigurationShrinkRequestInstanceTypeOverride) *CreateScalingConfigurationShrinkRequest {
	s.InstanceTypeOverride = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetDataDisk(v []*CreateScalingConfigurationShrinkRequestDataDisk) *CreateScalingConfigurationShrinkRequest {
	s.DataDisk = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSpotPriceLimit(v []*CreateScalingConfigurationShrinkRequestSpotPriceLimit) *CreateScalingConfigurationShrinkRequest {
	s.SpotPriceLimit = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSecurityGroupIds(v []*string) *CreateScalingConfigurationShrinkRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInstancePatternInfo(v []*CreateScalingConfigurationShrinkRequestInstancePatternInfo) *CreateScalingConfigurationShrinkRequest {
	s.InstancePatternInfo = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSystemDiskCategory(v []*string) *CreateScalingConfigurationShrinkRequest {
	s.SystemDiskCategory = v
	return s
}

type CreateScalingConfigurationShrinkRequestSystemDisk struct {
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Size                 *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	DiskName             *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	PerformanceLevel     *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
}

func (s CreateScalingConfigurationShrinkRequestSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetCategory(v string) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetSize(v int32) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetDiskName(v string) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetDescription(v string) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetAutoSnapshotPolicyId(v string) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetPerformanceLevel(v string) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

type CreateScalingConfigurationShrinkRequestPrivatePoolOptions struct {
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateScalingConfigurationShrinkRequestPrivatePoolOptions) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestPrivatePoolOptions) SetMatchCriteria(v string) *CreateScalingConfigurationShrinkRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestPrivatePoolOptions) SetId(v string) *CreateScalingConfigurationShrinkRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

type CreateScalingConfigurationShrinkRequestInstanceTypeOverride struct {
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	WeightedCapacity *int32  `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s CreateScalingConfigurationShrinkRequestInstanceTypeOverride) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestInstanceTypeOverride) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestInstanceTypeOverride) SetInstanceType(v string) *CreateScalingConfigurationShrinkRequestInstanceTypeOverride {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstanceTypeOverride) SetWeightedCapacity(v int32) *CreateScalingConfigurationShrinkRequestInstanceTypeOverride {
	s.WeightedCapacity = &v
	return s
}

type CreateScalingConfigurationShrinkRequestDataDisk struct {
	Categorys            []*string `json:"Categorys,omitempty" xml:"Categorys,omitempty" type:"Repeated"`
	PerformanceLevel     *string   `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	AutoSnapshotPolicyId *string   `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	Encrypted            *string   `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	Description          *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	SnapshotId           *string   `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	Size                 *int32    `json:"Size,omitempty" xml:"Size,omitempty"`
	Device               *string   `json:"Device,omitempty" xml:"Device,omitempty"`
	DiskName             *string   `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	Category             *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteWithInstance   *bool     `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	KMSKeyId             *string   `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
}

func (s CreateScalingConfigurationShrinkRequestDataDisk) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestDataDisk) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestDataDisk) SetCategorys(v []*string) *CreateScalingConfigurationShrinkRequestDataDisk {
	s.Categorys = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisk) SetPerformanceLevel(v string) *CreateScalingConfigurationShrinkRequestDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisk) SetAutoSnapshotPolicyId(v string) *CreateScalingConfigurationShrinkRequestDataDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisk) SetEncrypted(v string) *CreateScalingConfigurationShrinkRequestDataDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisk) SetDescription(v string) *CreateScalingConfigurationShrinkRequestDataDisk {
	s.Description = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisk) SetSnapshotId(v string) *CreateScalingConfigurationShrinkRequestDataDisk {
	s.SnapshotId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisk) SetSize(v int32) *CreateScalingConfigurationShrinkRequestDataDisk {
	s.Size = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisk) SetDevice(v string) *CreateScalingConfigurationShrinkRequestDataDisk {
	s.Device = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisk) SetDiskName(v string) *CreateScalingConfigurationShrinkRequestDataDisk {
	s.DiskName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisk) SetCategory(v string) *CreateScalingConfigurationShrinkRequestDataDisk {
	s.Category = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisk) SetDeleteWithInstance(v bool) *CreateScalingConfigurationShrinkRequestDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisk) SetKMSKeyId(v string) *CreateScalingConfigurationShrinkRequestDataDisk {
	s.KMSKeyId = &v
	return s
}

type CreateScalingConfigurationShrinkRequestSpotPriceLimit struct {
	PriceLimit   *float32 `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
	InstanceType *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateScalingConfigurationShrinkRequestSpotPriceLimit) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestSpotPriceLimit) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestSpotPriceLimit) SetPriceLimit(v float32) *CreateScalingConfigurationShrinkRequestSpotPriceLimit {
	s.PriceLimit = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSpotPriceLimit) SetInstanceType(v string) *CreateScalingConfigurationShrinkRequestSpotPriceLimit {
	s.InstanceType = &v
	return s
}

type CreateScalingConfigurationShrinkRequestInstancePatternInfo struct {
	Cores               *int32   `json:"Cores,omitempty" xml:"Cores,omitempty"`
	InstanceFamilyLevel *string  `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	MaxPrice            *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	Memory              *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s CreateScalingConfigurationShrinkRequestInstancePatternInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestInstancePatternInfo) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfo) SetCores(v int32) *CreateScalingConfigurationShrinkRequestInstancePatternInfo {
	s.Cores = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfo) SetInstanceFamilyLevel(v string) *CreateScalingConfigurationShrinkRequestInstancePatternInfo {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfo) SetMaxPrice(v float32) *CreateScalingConfigurationShrinkRequestInstancePatternInfo {
	s.MaxPrice = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfo) SetMemory(v float32) *CreateScalingConfigurationShrinkRequestInstancePatternInfo {
	s.Memory = &v
	return s
}

type CreateScalingConfigurationResponseBody struct {
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateScalingConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationResponseBody) SetScalingConfigurationId(v string) *CreateScalingConfigurationResponseBody {
	s.ScalingConfigurationId = &v
	return s
}

func (s *CreateScalingConfigurationResponseBody) SetRequestId(v string) *CreateScalingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type CreateScalingConfigurationResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateScalingConfigurationResponse) SetBody(v *CreateScalingConfigurationResponseBody) *CreateScalingConfigurationResponse {
	s.Body = v
	return s
}

type CreateScalingGroupRequest struct {
	RemovalPolicy                       []*string                                          `json:"RemovalPolicy,omitempty" xml:"RemovalPolicy,omitempty" type:"Repeated"`
	OwnerId                             *int64                                             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount                *string                                            `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupName                    *string                                            `json:"ScalingGroupName,omitempty" xml:"ScalingGroupName,omitempty"`
	LaunchTemplateId                    *string                                            `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	LaunchTemplateVersion               *string                                            `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	InstanceId                          *string                                            `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId                            *string                                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MinSize                             *int32                                             `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	MaxSize                             *int32                                             `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	DefaultCooldown                     *int32                                             `json:"DefaultCooldown,omitempty" xml:"DefaultCooldown,omitempty"`
	LoadBalancerIds                     *string                                            `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty"`
	DBInstanceIds                       *string                                            `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	OwnerAccount                        *string                                            `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	VSwitchId                           *string                                            `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	MultiAZPolicy                       *string                                            `json:"MultiAZPolicy,omitempty" xml:"MultiAZPolicy,omitempty"`
	HealthCheckType                     *string                                            `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	ScalingPolicy                       *string                                            `json:"ScalingPolicy,omitempty" xml:"ScalingPolicy,omitempty"`
	ClientToken                         *string                                            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OnDemandBaseCapacity                *int32                                             `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	OnDemandPercentageAboveBaseCapacity *int32                                             `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	SpotInstanceRemedy                  *bool                                              `json:"SpotInstanceRemedy,omitempty" xml:"SpotInstanceRemedy,omitempty"`
	CompensateWithOnDemand              *bool                                              `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	SpotInstancePools                   *int32                                             `json:"SpotInstancePools,omitempty" xml:"SpotInstancePools,omitempty"`
	DesiredCapacity                     *int32                                             `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	GroupDeletionProtection             *bool                                              `json:"GroupDeletionProtection,omitempty" xml:"GroupDeletionProtection,omitempty"`
	VSwitchIds                          []*string                                          `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	LifecycleHook                       []*CreateScalingGroupRequestLifecycleHook          `json:"LifecycleHook,omitempty" xml:"LifecycleHook,omitempty" type:"Repeated"`
	VServerGroup                        []*CreateScalingGroupRequestVServerGroup           `json:"VServerGroup,omitempty" xml:"VServerGroup,omitempty" type:"Repeated"`
	Tag                                 []*CreateScalingGroupRequestTag                    `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	LaunchTemplateOverride              []*CreateScalingGroupRequestLaunchTemplateOverride `json:"LaunchTemplateOverride,omitempty" xml:"LaunchTemplateOverride,omitempty" type:"Repeated"`
	AlbServerGroup                      []*CreateScalingGroupRequestAlbServerGroup         `json:"AlbServerGroup,omitempty" xml:"AlbServerGroup,omitempty" type:"Repeated"`
}

func (s CreateScalingGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequest) SetRemovalPolicy(v []*string) *CreateScalingGroupRequest {
	s.RemovalPolicy = v
	return s
}

func (s *CreateScalingGroupRequest) SetOwnerId(v int64) *CreateScalingGroupRequest {
	s.OwnerId = &v
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

func (s *CreateScalingGroupRequest) SetLaunchTemplateId(v string) *CreateScalingGroupRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetLaunchTemplateVersion(v string) *CreateScalingGroupRequest {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *CreateScalingGroupRequest) SetInstanceId(v string) *CreateScalingGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetRegionId(v string) *CreateScalingGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetMinSize(v int32) *CreateScalingGroupRequest {
	s.MinSize = &v
	return s
}

func (s *CreateScalingGroupRequest) SetMaxSize(v int32) *CreateScalingGroupRequest {
	s.MaxSize = &v
	return s
}

func (s *CreateScalingGroupRequest) SetDefaultCooldown(v int32) *CreateScalingGroupRequest {
	s.DefaultCooldown = &v
	return s
}

func (s *CreateScalingGroupRequest) SetLoadBalancerIds(v string) *CreateScalingGroupRequest {
	s.LoadBalancerIds = &v
	return s
}

func (s *CreateScalingGroupRequest) SetDBInstanceIds(v string) *CreateScalingGroupRequest {
	s.DBInstanceIds = &v
	return s
}

func (s *CreateScalingGroupRequest) SetOwnerAccount(v string) *CreateScalingGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateScalingGroupRequest) SetVSwitchId(v string) *CreateScalingGroupRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetMultiAZPolicy(v string) *CreateScalingGroupRequest {
	s.MultiAZPolicy = &v
	return s
}

func (s *CreateScalingGroupRequest) SetHealthCheckType(v string) *CreateScalingGroupRequest {
	s.HealthCheckType = &v
	return s
}

func (s *CreateScalingGroupRequest) SetScalingPolicy(v string) *CreateScalingGroupRequest {
	s.ScalingPolicy = &v
	return s
}

func (s *CreateScalingGroupRequest) SetClientToken(v string) *CreateScalingGroupRequest {
	s.ClientToken = &v
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

func (s *CreateScalingGroupRequest) SetSpotInstanceRemedy(v bool) *CreateScalingGroupRequest {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *CreateScalingGroupRequest) SetCompensateWithOnDemand(v bool) *CreateScalingGroupRequest {
	s.CompensateWithOnDemand = &v
	return s
}

func (s *CreateScalingGroupRequest) SetSpotInstancePools(v int32) *CreateScalingGroupRequest {
	s.SpotInstancePools = &v
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

func (s *CreateScalingGroupRequest) SetVSwitchIds(v []*string) *CreateScalingGroupRequest {
	s.VSwitchIds = v
	return s
}

func (s *CreateScalingGroupRequest) SetLifecycleHook(v []*CreateScalingGroupRequestLifecycleHook) *CreateScalingGroupRequest {
	s.LifecycleHook = v
	return s
}

func (s *CreateScalingGroupRequest) SetVServerGroup(v []*CreateScalingGroupRequestVServerGroup) *CreateScalingGroupRequest {
	s.VServerGroup = v
	return s
}

func (s *CreateScalingGroupRequest) SetTag(v []*CreateScalingGroupRequestTag) *CreateScalingGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateScalingGroupRequest) SetLaunchTemplateOverride(v []*CreateScalingGroupRequestLaunchTemplateOverride) *CreateScalingGroupRequest {
	s.LaunchTemplateOverride = v
	return s
}

func (s *CreateScalingGroupRequest) SetAlbServerGroup(v []*CreateScalingGroupRequestAlbServerGroup) *CreateScalingGroupRequest {
	s.AlbServerGroup = v
	return s
}

type CreateScalingGroupRequestLifecycleHook struct {
	DefaultResult        *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	LifecycleHookName    *string `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	LifecycleTransition  *string `json:"LifecycleTransition,omitempty" xml:"LifecycleTransition,omitempty"`
	NotificationMetadata *string `json:"NotificationMetadata,omitempty" xml:"NotificationMetadata,omitempty"`
	NotificationArn      *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	HeartbeatTimeout     *int32  `json:"HeartbeatTimeout,omitempty" xml:"HeartbeatTimeout,omitempty"`
}

func (s CreateScalingGroupRequestLifecycleHook) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupRequestLifecycleHook) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestLifecycleHook) SetDefaultResult(v string) *CreateScalingGroupRequestLifecycleHook {
	s.DefaultResult = &v
	return s
}

func (s *CreateScalingGroupRequestLifecycleHook) SetLifecycleHookName(v string) *CreateScalingGroupRequestLifecycleHook {
	s.LifecycleHookName = &v
	return s
}

func (s *CreateScalingGroupRequestLifecycleHook) SetLifecycleTransition(v string) *CreateScalingGroupRequestLifecycleHook {
	s.LifecycleTransition = &v
	return s
}

func (s *CreateScalingGroupRequestLifecycleHook) SetNotificationMetadata(v string) *CreateScalingGroupRequestLifecycleHook {
	s.NotificationMetadata = &v
	return s
}

func (s *CreateScalingGroupRequestLifecycleHook) SetNotificationArn(v string) *CreateScalingGroupRequestLifecycleHook {
	s.NotificationArn = &v
	return s
}

func (s *CreateScalingGroupRequestLifecycleHook) SetHeartbeatTimeout(v int32) *CreateScalingGroupRequestLifecycleHook {
	s.HeartbeatTimeout = &v
	return s
}

type CreateScalingGroupRequestVServerGroup struct {
	VServerGroupAttribute []*CreateScalingGroupRequestVServerGroupVServerGroupAttribute `json:"VServerGroupAttribute,omitempty" xml:"VServerGroupAttribute,omitempty" type:"Repeated"`
	LoadBalancerId        *string                                                       `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s CreateScalingGroupRequestVServerGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupRequestVServerGroup) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestVServerGroup) SetVServerGroupAttribute(v []*CreateScalingGroupRequestVServerGroupVServerGroupAttribute) *CreateScalingGroupRequestVServerGroup {
	s.VServerGroupAttribute = v
	return s
}

func (s *CreateScalingGroupRequestVServerGroup) SetLoadBalancerId(v string) *CreateScalingGroupRequestVServerGroup {
	s.LoadBalancerId = &v
	return s
}

type CreateScalingGroupRequestVServerGroupVServerGroupAttribute struct {
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	Weight         *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
	Port           *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s CreateScalingGroupRequestVServerGroupVServerGroupAttribute) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupRequestVServerGroupVServerGroupAttribute) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestVServerGroupVServerGroupAttribute) SetVServerGroupId(v string) *CreateScalingGroupRequestVServerGroupVServerGroupAttribute {
	s.VServerGroupId = &v
	return s
}

func (s *CreateScalingGroupRequestVServerGroupVServerGroupAttribute) SetWeight(v int32) *CreateScalingGroupRequestVServerGroupVServerGroupAttribute {
	s.Weight = &v
	return s
}

func (s *CreateScalingGroupRequestVServerGroupVServerGroupAttribute) SetPort(v int32) *CreateScalingGroupRequestVServerGroupVServerGroupAttribute {
	s.Port = &v
	return s
}

type CreateScalingGroupRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateScalingGroupRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestTag) SetKey(v string) *CreateScalingGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateScalingGroupRequestTag) SetValue(v string) *CreateScalingGroupRequestTag {
	s.Value = &v
	return s
}

type CreateScalingGroupRequestLaunchTemplateOverride struct {
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	WeightedCapacity *int32  `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s CreateScalingGroupRequestLaunchTemplateOverride) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupRequestLaunchTemplateOverride) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestLaunchTemplateOverride) SetInstanceType(v string) *CreateScalingGroupRequestLaunchTemplateOverride {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingGroupRequestLaunchTemplateOverride) SetWeightedCapacity(v int32) *CreateScalingGroupRequestLaunchTemplateOverride {
	s.WeightedCapacity = &v
	return s
}

type CreateScalingGroupRequestAlbServerGroup struct {
	AlbServerGroupId *string `json:"AlbServerGroupId,omitempty" xml:"AlbServerGroupId,omitempty"`
	Weight           *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s CreateScalingGroupRequestAlbServerGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupRequestAlbServerGroup) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestAlbServerGroup) SetAlbServerGroupId(v string) *CreateScalingGroupRequestAlbServerGroup {
	s.AlbServerGroupId = &v
	return s
}

func (s *CreateScalingGroupRequestAlbServerGroup) SetWeight(v int32) *CreateScalingGroupRequestAlbServerGroup {
	s.Weight = &v
	return s
}

func (s *CreateScalingGroupRequestAlbServerGroup) SetPort(v int32) *CreateScalingGroupRequestAlbServerGroup {
	s.Port = &v
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateScalingGroupResponse) SetBody(v *CreateScalingGroupResponseBody) *CreateScalingGroupResponse {
	s.Body = v
	return s
}

type CreateScalingRuleRequest struct {
	OwnerId                  *int64                                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount     *string                                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId           *string                                   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScalingRuleName          *string                                   `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	Cooldown                 *int32                                    `json:"Cooldown,omitempty" xml:"Cooldown,omitempty"`
	MinAdjustmentMagnitude   *int32                                    `json:"MinAdjustmentMagnitude,omitempty" xml:"MinAdjustmentMagnitude,omitempty"`
	AdjustmentType           *string                                   `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	AdjustmentValue          *int32                                    `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	ScalingRuleType          *string                                   `json:"ScalingRuleType,omitempty" xml:"ScalingRuleType,omitempty"`
	EstimatedInstanceWarmup  *int32                                    `json:"EstimatedInstanceWarmup,omitempty" xml:"EstimatedInstanceWarmup,omitempty"`
	MetricName               *string                                   `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	TargetValue              *float32                                  `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
	DisableScaleIn           *bool                                     `json:"DisableScaleIn,omitempty" xml:"DisableScaleIn,omitempty"`
	ScaleInEvaluationCount   *int32                                    `json:"ScaleInEvaluationCount,omitempty" xml:"ScaleInEvaluationCount,omitempty"`
	ScaleOutEvaluationCount  *int32                                    `json:"ScaleOutEvaluationCount,omitempty" xml:"ScaleOutEvaluationCount,omitempty"`
	OwnerAccount             *string                                   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	PredictiveScalingMode    *string                                   `json:"PredictiveScalingMode,omitempty" xml:"PredictiveScalingMode,omitempty"`
	PredictiveValueBehavior  *string                                   `json:"PredictiveValueBehavior,omitempty" xml:"PredictiveValueBehavior,omitempty"`
	PredictiveValueBuffer    *int32                                    `json:"PredictiveValueBuffer,omitempty" xml:"PredictiveValueBuffer,omitempty"`
	PredictiveTaskBufferTime *int32                                    `json:"PredictiveTaskBufferTime,omitempty" xml:"PredictiveTaskBufferTime,omitempty"`
	InitialMaxSize           *int32                                    `json:"InitialMaxSize,omitempty" xml:"InitialMaxSize,omitempty"`
	StepAdjustment           []*CreateScalingRuleRequestStepAdjustment `json:"StepAdjustment,omitempty" xml:"StepAdjustment,omitempty" type:"Repeated"`
}

func (s CreateScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateScalingRuleRequest) SetOwnerId(v int64) *CreateScalingRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateScalingRuleRequest) SetResourceOwnerAccount(v string) *CreateScalingRuleRequest {
	s.ResourceOwnerAccount = &v
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

func (s *CreateScalingRuleRequest) SetCooldown(v int32) *CreateScalingRuleRequest {
	s.Cooldown = &v
	return s
}

func (s *CreateScalingRuleRequest) SetMinAdjustmentMagnitude(v int32) *CreateScalingRuleRequest {
	s.MinAdjustmentMagnitude = &v
	return s
}

func (s *CreateScalingRuleRequest) SetAdjustmentType(v string) *CreateScalingRuleRequest {
	s.AdjustmentType = &v
	return s
}

func (s *CreateScalingRuleRequest) SetAdjustmentValue(v int32) *CreateScalingRuleRequest {
	s.AdjustmentValue = &v
	return s
}

func (s *CreateScalingRuleRequest) SetScalingRuleType(v string) *CreateScalingRuleRequest {
	s.ScalingRuleType = &v
	return s
}

func (s *CreateScalingRuleRequest) SetEstimatedInstanceWarmup(v int32) *CreateScalingRuleRequest {
	s.EstimatedInstanceWarmup = &v
	return s
}

func (s *CreateScalingRuleRequest) SetMetricName(v string) *CreateScalingRuleRequest {
	s.MetricName = &v
	return s
}

func (s *CreateScalingRuleRequest) SetTargetValue(v float32) *CreateScalingRuleRequest {
	s.TargetValue = &v
	return s
}

func (s *CreateScalingRuleRequest) SetDisableScaleIn(v bool) *CreateScalingRuleRequest {
	s.DisableScaleIn = &v
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

func (s *CreateScalingRuleRequest) SetOwnerAccount(v string) *CreateScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateScalingRuleRequest) SetPredictiveScalingMode(v string) *CreateScalingRuleRequest {
	s.PredictiveScalingMode = &v
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

func (s *CreateScalingRuleRequest) SetPredictiveTaskBufferTime(v int32) *CreateScalingRuleRequest {
	s.PredictiveTaskBufferTime = &v
	return s
}

func (s *CreateScalingRuleRequest) SetInitialMaxSize(v int32) *CreateScalingRuleRequest {
	s.InitialMaxSize = &v
	return s
}

func (s *CreateScalingRuleRequest) SetStepAdjustment(v []*CreateScalingRuleRequestStepAdjustment) *CreateScalingRuleRequest {
	s.StepAdjustment = v
	return s
}

type CreateScalingRuleRequestStepAdjustment struct {
	MetricIntervalUpperBound *float32 `json:"MetricIntervalUpperBound,omitempty" xml:"MetricIntervalUpperBound,omitempty"`
	ScalingAdjustment        *int32   `json:"ScalingAdjustment,omitempty" xml:"ScalingAdjustment,omitempty"`
	MetricIntervalLowerBound *float32 `json:"MetricIntervalLowerBound,omitempty" xml:"MetricIntervalLowerBound,omitempty"`
}

func (s CreateScalingRuleRequestStepAdjustment) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingRuleRequestStepAdjustment) GoString() string {
	return s.String()
}

func (s *CreateScalingRuleRequestStepAdjustment) SetMetricIntervalUpperBound(v float32) *CreateScalingRuleRequestStepAdjustment {
	s.MetricIntervalUpperBound = &v
	return s
}

func (s *CreateScalingRuleRequestStepAdjustment) SetScalingAdjustment(v int32) *CreateScalingRuleRequestStepAdjustment {
	s.ScalingAdjustment = &v
	return s
}

func (s *CreateScalingRuleRequestStepAdjustment) SetMetricIntervalLowerBound(v float32) *CreateScalingRuleRequestStepAdjustment {
	s.MetricIntervalLowerBound = &v
	return s
}

type CreateScalingRuleResponseBody struct {
	ScalingRuleAri *string `json:"ScalingRuleAri,omitempty" xml:"ScalingRuleAri,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingRuleId  *string `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty"`
}

func (s CreateScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScalingRuleResponseBody) SetScalingRuleAri(v string) *CreateScalingRuleResponseBody {
	s.ScalingRuleAri = &v
	return s
}

func (s *CreateScalingRuleResponseBody) SetRequestId(v string) *CreateScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScalingRuleResponseBody) SetScalingRuleId(v string) *CreateScalingRuleResponseBody {
	s.ScalingRuleId = &v
	return s
}

type CreateScalingRuleResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateScalingRuleResponse) SetBody(v *CreateScalingRuleResponseBody) *CreateScalingRuleResponse {
	s.Body = v
	return s
}

type CreateScheduledTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScheduledTaskName    *string `json:"ScheduledTaskName,omitempty" xml:"ScheduledTaskName,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ScheduledAction      *string `json:"ScheduledAction,omitempty" xml:"ScheduledAction,omitempty"`
	RecurrenceEndTime    *string `json:"RecurrenceEndTime,omitempty" xml:"RecurrenceEndTime,omitempty"`
	LaunchTime           *string `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	RecurrenceType       *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValue      *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	TaskEnabled          *bool   `json:"TaskEnabled,omitempty" xml:"TaskEnabled,omitempty"`
	LaunchExpirationTime *int32  `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	MinValue             *int32  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	MaxValue             *int32  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	DesiredCapacity      *int32  `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s CreateScheduledTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskRequest) SetOwnerId(v int64) *CreateScheduledTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetResourceOwnerAccount(v string) *CreateScheduledTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetRegionId(v string) *CreateScheduledTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetScheduledTaskName(v string) *CreateScheduledTaskRequest {
	s.ScheduledTaskName = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetDescription(v string) *CreateScheduledTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetScheduledAction(v string) *CreateScheduledTaskRequest {
	s.ScheduledAction = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetRecurrenceEndTime(v string) *CreateScheduledTaskRequest {
	s.RecurrenceEndTime = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetLaunchTime(v string) *CreateScheduledTaskRequest {
	s.LaunchTime = &v
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

func (s *CreateScheduledTaskRequest) SetTaskEnabled(v bool) *CreateScheduledTaskRequest {
	s.TaskEnabled = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetLaunchExpirationTime(v int32) *CreateScheduledTaskRequest {
	s.LaunchExpirationTime = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetOwnerAccount(v string) *CreateScheduledTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetMinValue(v int32) *CreateScheduledTaskRequest {
	s.MinValue = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetMaxValue(v int32) *CreateScheduledTaskRequest {
	s.MaxValue = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetDesiredCapacity(v int32) *CreateScheduledTaskRequest {
	s.DesiredCapacity = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetScalingGroupId(v string) *CreateScheduledTaskRequest {
	s.ScalingGroupId = &v
	return s
}

type CreateScheduledTaskResponseBody struct {
	ScheduledTaskId *string `json:"ScheduledTaskId,omitempty" xml:"ScheduledTaskId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateScheduledTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskResponseBody) SetScheduledTaskId(v string) *CreateScheduledTaskResponseBody {
	s.ScheduledTaskId = &v
	return s
}

func (s *CreateScheduledTaskResponseBody) SetRequestId(v string) *CreateScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateScheduledTaskResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateScheduledTaskResponse) SetBody(v *CreateScheduledTaskResponseBody) *CreateScheduledTaskResponse {
	s.Body = v
	return s
}

type DeactivateScalingConfigurationRequest struct {
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s DeactivateScalingConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeactivateScalingConfigurationRequest) GoString() string {
	return s.String()
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

func (s *DeactivateScalingConfigurationRequest) SetOwnerAccount(v string) *DeactivateScalingConfigurationRequest {
	s.OwnerAccount = &v
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeactivateScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeactivateScalingConfigurationResponse) SetBody(v *DeactivateScalingConfigurationResponseBody) *DeactivateScalingConfigurationResponse {
	s.Body = v
	return s
}

type DeleteAlarmRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AlarmTaskId          *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
}

func (s DeleteAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlarmRequest) SetOwnerId(v int64) *DeleteAlarmRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAlarmRequest) SetResourceOwnerAccount(v string) *DeleteAlarmRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAlarmRequest) SetRegionId(v string) *DeleteAlarmRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAlarmRequest) SetAlarmTaskId(v string) *DeleteAlarmRequest {
	s.AlarmTaskId = &v
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteAlarmResponse) SetBody(v *DeleteAlarmResponseBody) *DeleteAlarmResponse {
	s.Body = v
	return s
}

type DeleteLifecycleHookRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	LifecycleHookId      *string `json:"LifecycleHookId,omitempty" xml:"LifecycleHookId,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	LifecycleHookName    *string `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
}

func (s DeleteLifecycleHookRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLifecycleHookRequest) GoString() string {
	return s.String()
}

func (s *DeleteLifecycleHookRequest) SetOwnerId(v int64) *DeleteLifecycleHookRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLifecycleHookRequest) SetResourceOwnerAccount(v string) *DeleteLifecycleHookRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteLifecycleHookRequest) SetOwnerAccount(v string) *DeleteLifecycleHookRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteLifecycleHookRequest) SetLifecycleHookId(v string) *DeleteLifecycleHookRequest {
	s.LifecycleHookId = &v
	return s
}

func (s *DeleteLifecycleHookRequest) SetScalingGroupId(v string) *DeleteLifecycleHookRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DeleteLifecycleHookRequest) SetLifecycleHookName(v string) *DeleteLifecycleHookRequest {
	s.LifecycleHookName = &v
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteLifecycleHookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteLifecycleHookResponse) SetBody(v *DeleteLifecycleHookResponseBody) *DeleteLifecycleHookResponse {
	s.Body = v
	return s
}

type DeleteNotificationConfigurationRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	NotificationArn      *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
}

func (s DeleteNotificationConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNotificationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeleteNotificationConfigurationRequest) SetOwnerId(v int64) *DeleteNotificationConfigurationRequest {
	s.OwnerId = &v
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

func (s *DeleteNotificationConfigurationRequest) SetNotificationArn(v string) *DeleteNotificationConfigurationRequest {
	s.NotificationArn = &v
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
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteNotificationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteNotificationConfigurationResponse) SetBody(v *DeleteNotificationConfigurationResponseBody) *DeleteNotificationConfigurationResponse {
	s.Body = v
	return s
}

type DeleteScalingConfigurationRequest struct {
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s DeleteScalingConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingConfigurationRequest) GoString() string {
	return s.String()
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

func (s *DeleteScalingConfigurationRequest) SetOwnerAccount(v string) *DeleteScalingConfigurationRequest {
	s.OwnerAccount = &v
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteScalingConfigurationResponse) SetBody(v *DeleteScalingConfigurationResponseBody) *DeleteScalingConfigurationResponse {
	s.Body = v
	return s
}

type DeleteScalingGroupRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ForceDelete          *bool   `json:"ForceDelete,omitempty" xml:"ForceDelete,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s DeleteScalingGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteScalingGroupRequest) SetOwnerId(v int64) *DeleteScalingGroupRequest {
	s.OwnerId = &v
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

func (s *DeleteScalingGroupRequest) SetForceDelete(v bool) *DeleteScalingGroupRequest {
	s.ForceDelete = &v
	return s
}

func (s *DeleteScalingGroupRequest) SetOwnerAccount(v string) *DeleteScalingGroupRequest {
	s.OwnerAccount = &v
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteScalingGroupResponse) SetBody(v *DeleteScalingGroupResponseBody) *DeleteScalingGroupResponse {
	s.Body = v
	return s
}

type DeleteScalingRuleRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingRuleId        *string `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s DeleteScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteScalingRuleRequest) SetOwnerId(v int64) *DeleteScalingRuleRequest {
	s.OwnerId = &v
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

func (s *DeleteScalingRuleRequest) SetOwnerAccount(v string) *DeleteScalingRuleRequest {
	s.OwnerAccount = &v
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteScalingRuleResponse) SetBody(v *DeleteScalingRuleResponseBody) *DeleteScalingRuleResponse {
	s.Body = v
	return s
}

type DeleteScheduledTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScheduledTaskId      *string `json:"ScheduledTaskId,omitempty" xml:"ScheduledTaskId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s DeleteScheduledTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduledTaskRequest) SetOwnerId(v int64) *DeleteScheduledTaskRequest {
	s.OwnerId = &v
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

func (s *DeleteScheduledTaskRequest) SetOwnerAccount(v string) *DeleteScheduledTaskRequest {
	s.OwnerAccount = &v
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteScheduledTaskResponse) SetBody(v *DeleteScheduledTaskResponseBody) *DeleteScheduledTaskResponse {
	s.Body = v
	return s
}

type DescribeAlarmsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	AlarmTaskId          *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	State                *string `json:"State,omitempty" xml:"State,omitempty"`
	IsEnable             *bool   `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
	MetricType           *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	MetricName           *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeAlarmsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsRequest) SetOwnerId(v int64) *DescribeAlarmsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAlarmsRequest) SetResourceOwnerAccount(v string) *DescribeAlarmsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAlarmsRequest) SetRegionId(v string) *DescribeAlarmsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlarmsRequest) SetScalingGroupId(v string) *DescribeAlarmsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeAlarmsRequest) SetAlarmTaskId(v string) *DescribeAlarmsRequest {
	s.AlarmTaskId = &v
	return s
}

func (s *DescribeAlarmsRequest) SetState(v string) *DescribeAlarmsRequest {
	s.State = &v
	return s
}

func (s *DescribeAlarmsRequest) SetIsEnable(v bool) *DescribeAlarmsRequest {
	s.IsEnable = &v
	return s
}

func (s *DescribeAlarmsRequest) SetMetricType(v string) *DescribeAlarmsRequest {
	s.MetricType = &v
	return s
}

func (s *DescribeAlarmsRequest) SetMetricName(v string) *DescribeAlarmsRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeAlarmsRequest) SetPageSize(v int32) *DescribeAlarmsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlarmsRequest) SetPageNumber(v int32) *DescribeAlarmsRequest {
	s.PageNumber = &v
	return s
}

type DescribeAlarmsResponseBody struct {
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	AlarmList  *DescribeAlarmsResponseBodyAlarmList `json:"AlarmList,omitempty" xml:"AlarmList,omitempty" type:"Struct"`
}

func (s DescribeAlarmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBody) SetPageSize(v int32) *DescribeAlarmsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetRequestId(v string) *DescribeAlarmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetPageNumber(v int32) *DescribeAlarmsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetTotalCount(v int32) *DescribeAlarmsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetAlarmList(v *DescribeAlarmsResponseBodyAlarmList) *DescribeAlarmsResponseBody {
	s.AlarmList = v
	return s
}

type DescribeAlarmsResponseBodyAlarmList struct {
	Alarm []*DescribeAlarmsResponseBodyAlarmListAlarm `json:"Alarm,omitempty" xml:"Alarm,omitempty" type:"Repeated"`
}

func (s DescribeAlarmsResponseBodyAlarmList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponseBodyAlarmList) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetAlarm(v []*DescribeAlarmsResponseBodyAlarmListAlarm) *DescribeAlarmsResponseBodyAlarmList {
	s.Alarm = v
	return s
}

type DescribeAlarmsResponseBodyAlarmListAlarm struct {
	AlarmTaskId        *string                                               `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	MetricName         *string                                               `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	EvaluationCount    *int32                                                `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	State              *string                                               `json:"State,omitempty" xml:"State,omitempty"`
	Period             *int32                                                `json:"Period,omitempty" xml:"Period,omitempty"`
	ScalingGroupId     *string                                               `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ComparisonOperator *string                                               `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Effective          *string                                               `json:"Effective,omitempty" xml:"Effective,omitempty"`
	Description        *string                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	MetricType         *string                                               `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Name               *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Threshold          *float32                                              `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Enable             *bool                                                 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Statistics         *string                                               `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Dimensions         *DescribeAlarmsResponseBodyAlarmListAlarmDimensions   `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Struct"`
	AlarmActions       *DescribeAlarmsResponseBodyAlarmListAlarmAlarmActions `json:"AlarmActions,omitempty" xml:"AlarmActions,omitempty" type:"Struct"`
}

func (s DescribeAlarmsResponseBodyAlarmListAlarm) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponseBodyAlarmListAlarm) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyAlarmListAlarm) SetAlarmTaskId(v string) *DescribeAlarmsResponseBodyAlarmListAlarm {
	s.AlarmTaskId = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListAlarm) SetMetricName(v string) *DescribeAlarmsResponseBodyAlarmListAlarm {
	s.MetricName = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListAlarm) SetEvaluationCount(v int32) *DescribeAlarmsResponseBodyAlarmListAlarm {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListAlarm) SetState(v string) *DescribeAlarmsResponseBodyAlarmListAlarm {
	s.State = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListAlarm) SetPeriod(v int32) *DescribeAlarmsResponseBodyAlarmListAlarm {
	s.Period = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListAlarm) SetScalingGroupId(v string) *DescribeAlarmsResponseBodyAlarmListAlarm {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListAlarm) SetComparisonOperator(v string) *DescribeAlarmsResponseBodyAlarmListAlarm {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListAlarm) SetEffective(v string) *DescribeAlarmsResponseBodyAlarmListAlarm {
	s.Effective = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListAlarm) SetDescription(v string) *DescribeAlarmsResponseBodyAlarmListAlarm {
	s.Description = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListAlarm) SetMetricType(v string) *DescribeAlarmsResponseBodyAlarmListAlarm {
	s.MetricType = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListAlarm) SetName(v string) *DescribeAlarmsResponseBodyAlarmListAlarm {
	s.Name = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListAlarm) SetThreshold(v float32) *DescribeAlarmsResponseBodyAlarmListAlarm {
	s.Threshold = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListAlarm) SetEnable(v bool) *DescribeAlarmsResponseBodyAlarmListAlarm {
	s.Enable = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListAlarm) SetStatistics(v string) *DescribeAlarmsResponseBodyAlarmListAlarm {
	s.Statistics = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListAlarm) SetDimensions(v *DescribeAlarmsResponseBodyAlarmListAlarmDimensions) *DescribeAlarmsResponseBodyAlarmListAlarm {
	s.Dimensions = v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListAlarm) SetAlarmActions(v *DescribeAlarmsResponseBodyAlarmListAlarmAlarmActions) *DescribeAlarmsResponseBodyAlarmListAlarm {
	s.AlarmActions = v
	return s
}

type DescribeAlarmsResponseBodyAlarmListAlarmDimensions struct {
	Dimension []*DescribeAlarmsResponseBodyAlarmListAlarmDimensionsDimension `json:"Dimension,omitempty" xml:"Dimension,omitempty" type:"Repeated"`
}

func (s DescribeAlarmsResponseBodyAlarmListAlarmDimensions) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponseBodyAlarmListAlarmDimensions) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyAlarmListAlarmDimensions) SetDimension(v []*DescribeAlarmsResponseBodyAlarmListAlarmDimensionsDimension) *DescribeAlarmsResponseBodyAlarmListAlarmDimensions {
	s.Dimension = v
	return s
}

type DescribeAlarmsResponseBodyAlarmListAlarmDimensionsDimension struct {
	DimensionKey   *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
}

func (s DescribeAlarmsResponseBodyAlarmListAlarmDimensionsDimension) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponseBodyAlarmListAlarmDimensionsDimension) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyAlarmListAlarmDimensionsDimension) SetDimensionKey(v string) *DescribeAlarmsResponseBodyAlarmListAlarmDimensionsDimension {
	s.DimensionKey = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListAlarmDimensionsDimension) SetDimensionValue(v string) *DescribeAlarmsResponseBodyAlarmListAlarmDimensionsDimension {
	s.DimensionValue = &v
	return s
}

type DescribeAlarmsResponseBodyAlarmListAlarmAlarmActions struct {
	AlarmAction []*string `json:"AlarmAction,omitempty" xml:"AlarmAction,omitempty" type:"Repeated"`
}

func (s DescribeAlarmsResponseBodyAlarmListAlarmAlarmActions) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponseBodyAlarmListAlarmAlarmActions) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyAlarmListAlarmAlarmActions) SetAlarmAction(v []*string) *DescribeAlarmsResponseBodyAlarmListAlarmAlarmActions {
	s.AlarmAction = v
	return s
}

type DescribeAlarmsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAlarmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeAlarmsResponse) SetBody(v *DescribeAlarmsResponseBody) *DescribeAlarmsResponse {
	s.Body = v
	return s
}

type DescribeLifecycleActionsRequest struct {
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingActivityId     *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	LifecycleActionStatus *string `json:"LifecycleActionStatus,omitempty" xml:"LifecycleActionStatus,omitempty"`
	NextToken             *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults            *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s DescribeLifecycleActionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecycleActionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleActionsRequest) SetOwnerId(v int64) *DescribeLifecycleActionsRequest {
	s.OwnerId = &v
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

func (s *DescribeLifecycleActionsRequest) SetLifecycleActionStatus(v string) *DescribeLifecycleActionsRequest {
	s.LifecycleActionStatus = &v
	return s
}

func (s *DescribeLifecycleActionsRequest) SetNextToken(v string) *DescribeLifecycleActionsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeLifecycleActionsRequest) SetMaxResults(v int32) *DescribeLifecycleActionsRequest {
	s.MaxResults = &v
	return s
}

type DescribeLifecycleActionsResponseBody struct {
	NextToken        *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount       *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	MaxResults       *int32                                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	LifecycleActions *DescribeLifecycleActionsResponseBodyLifecycleActions `json:"LifecycleActions,omitempty" xml:"LifecycleActions,omitempty" type:"Struct"`
}

func (s DescribeLifecycleActionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecycleActionsResponseBody) GoString() string {
	return s.String()
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

func (s *DescribeLifecycleActionsResponseBody) SetMaxResults(v int32) *DescribeLifecycleActionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeLifecycleActionsResponseBody) SetLifecycleActions(v *DescribeLifecycleActionsResponseBodyLifecycleActions) *DescribeLifecycleActionsResponseBody {
	s.LifecycleActions = v
	return s
}

type DescribeLifecycleActionsResponseBodyLifecycleActions struct {
	LifecycleAction []*DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleAction `json:"LifecycleAction,omitempty" xml:"LifecycleAction,omitempty" type:"Repeated"`
}

func (s DescribeLifecycleActionsResponseBodyLifecycleActions) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecycleActionsResponseBodyLifecycleActions) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActions) SetLifecycleAction(v []*DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleAction) *DescribeLifecycleActionsResponseBodyLifecycleActions {
	s.LifecycleAction = v
	return s
}

type DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleAction struct {
	LifecycleHookId       *string                                                                         `json:"LifecycleHookId,omitempty" xml:"LifecycleHookId,omitempty"`
	LifecycleActionToken  *string                                                                         `json:"LifecycleActionToken,omitempty" xml:"LifecycleActionToken,omitempty"`
	LifecycleActionStatus *string                                                                         `json:"LifecycleActionStatus,omitempty" xml:"LifecycleActionStatus,omitempty"`
	LifecycleActionResult *string                                                                         `json:"LifecycleActionResult,omitempty" xml:"LifecycleActionResult,omitempty"`
	InstanceIds           *DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleActionInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
}

func (s DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleAction) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleAction) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleAction) SetLifecycleHookId(v string) *DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleAction {
	s.LifecycleHookId = &v
	return s
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleAction) SetLifecycleActionToken(v string) *DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleAction {
	s.LifecycleActionToken = &v
	return s
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleAction) SetLifecycleActionStatus(v string) *DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleAction {
	s.LifecycleActionStatus = &v
	return s
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleAction) SetLifecycleActionResult(v string) *DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleAction {
	s.LifecycleActionResult = &v
	return s
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleAction) SetInstanceIds(v *DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleActionInstanceIds) *DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleAction {
	s.InstanceIds = v
	return s
}

type DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleActionInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleActionInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleActionInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleActionInstanceIds) SetInstanceId(v []*string) *DescribeLifecycleActionsResponseBodyLifecycleActionsLifecycleActionInstanceIds {
	s.InstanceId = v
	return s
}

type DescribeLifecycleActionsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLifecycleActionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeLifecycleActionsResponse) SetBody(v *DescribeLifecycleActionsResponseBody) *DescribeLifecycleActionsResponse {
	s.Body = v
	return s
}

type DescribeLifecycleHooksRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	LifecycleHookName    *string   `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	PageNumber           *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	LifecycleHookId      []*string `json:"LifecycleHookId,omitempty" xml:"LifecycleHookId,omitempty" type:"Repeated"`
}

func (s DescribeLifecycleHooksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecycleHooksRequest) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleHooksRequest) SetOwnerId(v int64) *DescribeLifecycleHooksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLifecycleHooksRequest) SetResourceOwnerAccount(v string) *DescribeLifecycleHooksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLifecycleHooksRequest) SetOwnerAccount(v string) *DescribeLifecycleHooksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLifecycleHooksRequest) SetScalingGroupId(v string) *DescribeLifecycleHooksRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeLifecycleHooksRequest) SetLifecycleHookName(v string) *DescribeLifecycleHooksRequest {
	s.LifecycleHookName = &v
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

func (s *DescribeLifecycleHooksRequest) SetLifecycleHookId(v []*string) *DescribeLifecycleHooksRequest {
	s.LifecycleHookId = v
	return s
}

type DescribeLifecycleHooksResponseBody struct {
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber     *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	LifecycleHooks *DescribeLifecycleHooksResponseBodyLifecycleHooks `json:"LifecycleHooks,omitempty" xml:"LifecycleHooks,omitempty" type:"Struct"`
}

func (s DescribeLifecycleHooksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecycleHooksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleHooksResponseBody) SetRequestId(v string) *DescribeLifecycleHooksResponseBody {
	s.RequestId = &v
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

func (s *DescribeLifecycleHooksResponseBody) SetTotalCount(v int32) *DescribeLifecycleHooksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBody) SetLifecycleHooks(v *DescribeLifecycleHooksResponseBodyLifecycleHooks) *DescribeLifecycleHooksResponseBody {
	s.LifecycleHooks = v
	return s
}

type DescribeLifecycleHooksResponseBodyLifecycleHooks struct {
	LifecycleHook []*DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook `json:"LifecycleHook,omitempty" xml:"LifecycleHook,omitempty" type:"Repeated"`
}

func (s DescribeLifecycleHooksResponseBodyLifecycleHooks) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecycleHooksResponseBodyLifecycleHooks) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) SetLifecycleHook(v []*DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook) *DescribeLifecycleHooksResponseBodyLifecycleHooks {
	s.LifecycleHook = v
	return s
}

type DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook struct {
	DefaultResult        *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	LifecycleHookId      *string `json:"LifecycleHookId,omitempty" xml:"LifecycleHookId,omitempty"`
	LifecycleHookName    *string `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	LifecycleTransition  *string `json:"LifecycleTransition,omitempty" xml:"LifecycleTransition,omitempty"`
	NotificationMetadata *string `json:"NotificationMetadata,omitempty" xml:"NotificationMetadata,omitempty"`
	NotificationArn      *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	HeartbeatTimeout     *int32  `json:"HeartbeatTimeout,omitempty" xml:"HeartbeatTimeout,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook) SetDefaultResult(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook {
	s.DefaultResult = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook) SetLifecycleHookId(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook {
	s.LifecycleHookId = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook) SetLifecycleHookName(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook {
	s.LifecycleHookName = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook) SetLifecycleTransition(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook {
	s.LifecycleTransition = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook) SetNotificationMetadata(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook {
	s.NotificationMetadata = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook) SetNotificationArn(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook {
	s.NotificationArn = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook) SetHeartbeatTimeout(v int32) *DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook {
	s.HeartbeatTimeout = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook) SetScalingGroupId(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooksLifecycleHook {
	s.ScalingGroupId = &v
	return s
}

type DescribeLifecycleHooksResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLifecycleHooksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MaxNumberOfScheduledTasks             *int32  `json:"MaxNumberOfScheduledTasks,omitempty" xml:"MaxNumberOfScheduledTasks,omitempty"`
	MaxNumberOfLoadBalancers              *int32  `json:"MaxNumberOfLoadBalancers,omitempty" xml:"MaxNumberOfLoadBalancers,omitempty"`
	MaxNumberOfMaxSize                    *int32  `json:"MaxNumberOfMaxSize,omitempty" xml:"MaxNumberOfMaxSize,omitempty"`
	MaxNumberOfAlbServerGroup             *int32  `json:"MaxNumberOfAlbServerGroup,omitempty" xml:"MaxNumberOfAlbServerGroup,omitempty"`
	MaxNumberOfDBInstances                *int32  `json:"MaxNumberOfDBInstances,omitempty" xml:"MaxNumberOfDBInstances,omitempty"`
	MaxNumberOfScalingConfigurations      *int32  `json:"MaxNumberOfScalingConfigurations,omitempty" xml:"MaxNumberOfScalingConfigurations,omitempty"`
	MaxNumberOfMinSize                    *int32  `json:"MaxNumberOfMinSize,omitempty" xml:"MaxNumberOfMinSize,omitempty"`
	MaxNumberOfLifecycleHooks             *int32  `json:"MaxNumberOfLifecycleHooks,omitempty" xml:"MaxNumberOfLifecycleHooks,omitempty"`
	MaxNumberOfScalingInstances           *int32  `json:"MaxNumberOfScalingInstances,omitempty" xml:"MaxNumberOfScalingInstances,omitempty"`
	MaxNumberOfScalingGroups              *int32  `json:"MaxNumberOfScalingGroups,omitempty" xml:"MaxNumberOfScalingGroups,omitempty"`
	RequestId                             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxNumberOfNotificationConfigurations *int32  `json:"MaxNumberOfNotificationConfigurations,omitempty" xml:"MaxNumberOfNotificationConfigurations,omitempty"`
	MaxNumberOfVServerGroups              *int32  `json:"MaxNumberOfVServerGroups,omitempty" xml:"MaxNumberOfVServerGroups,omitempty"`
	MaxNumberOfScalingRules               *int32  `json:"MaxNumberOfScalingRules,omitempty" xml:"MaxNumberOfScalingRules,omitempty"`
}

func (s DescribeLimitationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLimitationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfScheduledTasks(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfScheduledTasks = &v
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

func (s *DescribeLimitationResponseBody) SetMaxNumberOfAlbServerGroup(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfAlbServerGroup = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfDBInstances(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfDBInstances = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfScalingConfigurations(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfScalingConfigurations = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfMinSize(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfMinSize = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfLifecycleHooks(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfLifecycleHooks = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfScalingInstances(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfScalingInstances = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfScalingGroups(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfScalingGroups = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetRequestId(v string) *DescribeLimitationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfNotificationConfigurations(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfNotificationConfigurations = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfVServerGroups(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfVServerGroups = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfScalingRules(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfScalingRules = &v
	return s
}

type DescribeLimitationResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLimitationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeLimitationResponse) SetBody(v *DescribeLimitationResponseBody) *DescribeLimitationResponse {
	s.Body = v
	return s
}

type DescribeNotificationConfigurationsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s *DescribeNotificationConfigurationsRequest) SetResourceOwnerAccount(v string) *DescribeNotificationConfigurationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeNotificationConfigurationsRequest) SetScalingGroupId(v string) *DescribeNotificationConfigurationsRequest {
	s.ScalingGroupId = &v
	return s
}

type DescribeNotificationConfigurationsResponseBody struct {
	RequestId                       *string                                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NotificationConfigurationModels *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels `json:"NotificationConfigurationModels,omitempty" xml:"NotificationConfigurationModels,omitempty" type:"Struct"`
}

func (s DescribeNotificationConfigurationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotificationConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNotificationConfigurationsResponseBody) SetRequestId(v string) *DescribeNotificationConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNotificationConfigurationsResponseBody) SetNotificationConfigurationModels(v *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) *DescribeNotificationConfigurationsResponseBody {
	s.NotificationConfigurationModels = v
	return s
}

type DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels struct {
	NotificationConfigurationModel []*DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModelsNotificationConfigurationModel `json:"NotificationConfigurationModel,omitempty" xml:"NotificationConfigurationModel,omitempty" type:"Repeated"`
}

func (s DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) GoString() string {
	return s.String()
}

func (s *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) SetNotificationConfigurationModel(v []*DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModelsNotificationConfigurationModel) *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels {
	s.NotificationConfigurationModel = v
	return s
}

type DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModelsNotificationConfigurationModel struct {
	NotificationArn   *string                                                                                                                       `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	ScalingGroupId    *string                                                                                                                       `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	NotificationTypes *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModelsNotificationConfigurationModelNotificationTypes `json:"NotificationTypes,omitempty" xml:"NotificationTypes,omitempty" type:"Struct"`
}

func (s DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModelsNotificationConfigurationModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModelsNotificationConfigurationModel) GoString() string {
	return s.String()
}

func (s *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModelsNotificationConfigurationModel) SetNotificationArn(v string) *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModelsNotificationConfigurationModel {
	s.NotificationArn = &v
	return s
}

func (s *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModelsNotificationConfigurationModel) SetScalingGroupId(v string) *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModelsNotificationConfigurationModel {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModelsNotificationConfigurationModel) SetNotificationTypes(v *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModelsNotificationConfigurationModelNotificationTypes) *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModelsNotificationConfigurationModel {
	s.NotificationTypes = v
	return s
}

type DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModelsNotificationConfigurationModelNotificationTypes struct {
	NotificationType []*string `json:"NotificationType,omitempty" xml:"NotificationType,omitempty" type:"Repeated"`
}

func (s DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModelsNotificationConfigurationModelNotificationTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModelsNotificationConfigurationModelNotificationTypes) GoString() string {
	return s.String()
}

func (s *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModelsNotificationConfigurationModelNotificationTypes) SetNotificationType(v []*string) *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModelsNotificationConfigurationModelNotificationTypes {
	s.NotificationType = v
	return s
}

type DescribeNotificationConfigurationsResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNotificationConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NotificationTypes *DescribeNotificationTypesResponseBodyNotificationTypes `json:"NotificationTypes,omitempty" xml:"NotificationTypes,omitempty" type:"Struct"`
}

func (s DescribeNotificationTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotificationTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNotificationTypesResponseBody) SetRequestId(v string) *DescribeNotificationTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNotificationTypesResponseBody) SetNotificationTypes(v *DescribeNotificationTypesResponseBodyNotificationTypes) *DescribeNotificationTypesResponseBody {
	s.NotificationTypes = v
	return s
}

type DescribeNotificationTypesResponseBodyNotificationTypes struct {
	NotificationType []*string `json:"NotificationType,omitempty" xml:"NotificationType,omitempty" type:"Repeated"`
}

func (s DescribeNotificationTypesResponseBodyNotificationTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotificationTypesResponseBodyNotificationTypes) GoString() string {
	return s.String()
}

func (s *DescribeNotificationTypesResponseBodyNotificationTypes) SetNotificationType(v []*string) *DescribeNotificationTypesResponseBodyNotificationTypes {
	s.NotificationType = v
	return s
}

type DescribeNotificationTypesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNotificationTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeNotificationTypesResponse) SetBody(v *DescribeNotificationTypesResponseBody) *DescribeNotificationTypesResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	AcceptLanguage       *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
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

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
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
	ClassicUnavailable *bool   `json:"ClassicUnavailable,omitempty" xml:"ClassicUnavailable,omitempty"`
	RegionEndpoint     *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	LocalName          *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	VpcUnavailable     *bool   `json:"VpcUnavailable,omitempty" xml:"VpcUnavailable,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetClassicUnavailable(v bool) *DescribeRegionsResponseBodyRegionsRegion {
	s.ClassicUnavailable = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetVpcUnavailable(v bool) *DescribeRegionsResponseBodyRegionsRegion {
	s.VpcUnavailable = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeScalingActivitiesRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber           *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	StatusCode           *string   `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	ScalingActivityId    []*string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty" type:"Repeated"`
}

func (s DescribeScalingActivitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingActivitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivitiesRequest) SetOwnerId(v int64) *DescribeScalingActivitiesRequest {
	s.OwnerId = &v
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

func (s *DescribeScalingActivitiesRequest) SetRegionId(v string) *DescribeScalingActivitiesRequest {
	s.RegionId = &v
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

func (s *DescribeScalingActivitiesRequest) SetScalingGroupId(v string) *DescribeScalingActivitiesRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetStatusCode(v string) *DescribeScalingActivitiesRequest {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetOwnerAccount(v string) *DescribeScalingActivitiesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetScalingActivityId(v []*string) *DescribeScalingActivitiesRequest {
	s.ScalingActivityId = v
	return s
}

type DescribeScalingActivitiesResponseBody struct {
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber        *int32                                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount        *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ScalingActivities *DescribeScalingActivitiesResponseBodyScalingActivities `json:"ScalingActivities,omitempty" xml:"ScalingActivities,omitempty" type:"Struct"`
}

func (s DescribeScalingActivitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingActivitiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivitiesResponseBody) SetRequestId(v string) *DescribeScalingActivitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBody) SetPageNumber(v int32) *DescribeScalingActivitiesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBody) SetPageSize(v int32) *DescribeScalingActivitiesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBody) SetTotalCount(v int32) *DescribeScalingActivitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBody) SetScalingActivities(v *DescribeScalingActivitiesResponseBodyScalingActivities) *DescribeScalingActivitiesResponseBody {
	s.ScalingActivities = v
	return s
}

type DescribeScalingActivitiesResponseBodyScalingActivities struct {
	ScalingActivity []*DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity `json:"ScalingActivity,omitempty" xml:"ScalingActivity,omitempty" type:"Repeated"`
}

func (s DescribeScalingActivitiesResponseBodyScalingActivities) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingActivitiesResponseBodyScalingActivities) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetScalingActivity(v []*DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.ScalingActivity = v
	return s
}

type DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity struct {
	Progress              *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ScalingInstanceNumber *int32  `json:"ScalingInstanceNumber,omitempty" xml:"ScalingInstanceNumber,omitempty"`
	AttachedCapacity      *string `json:"AttachedCapacity,omitempty" xml:"AttachedCapacity,omitempty"`
	TotalCapacity         *string `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
	ScalingGroupId        *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	AutoCreatedCapacity   *string `json:"AutoCreatedCapacity,omitempty" xml:"AutoCreatedCapacity,omitempty"`
	EndTime               *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime             *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	StatusCode            *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	Cause                 *string `json:"Cause,omitempty" xml:"Cause,omitempty"`
	ScalingActivityId     *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	StatusMessage         *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetProgress(v int32) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.Progress = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetScalingInstanceNumber(v int32) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.ScalingInstanceNumber = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetAttachedCapacity(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.AttachedCapacity = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetTotalCapacity(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.TotalCapacity = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetScalingGroupId(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetAutoCreatedCapacity(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.AutoCreatedCapacity = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetEndTime(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.EndTime = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetStartTime(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.StartTime = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetDescription(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.Description = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetStatusCode(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetCause(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.Cause = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetScalingActivityId(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.ScalingActivityId = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetStatusMessage(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.StatusMessage = &v
	return s
}

type DescribeScalingActivitiesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScalingActivitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	Detail            *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
}

func (s DescribeScalingActivityDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingActivityDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivityDetailResponseBody) SetScalingActivityId(v string) *DescribeScalingActivityDetailResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *DescribeScalingActivityDetailResponseBody) SetDetail(v string) *DescribeScalingActivityDetailResponseBody {
	s.Detail = &v
	return s
}

type DescribeScalingActivityDetailResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScalingActivityDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeScalingActivityDetailResponse) SetBody(v *DescribeScalingActivityDetailResponseBody) *DescribeScalingActivityDetailResponse {
	s.Body = v
	return s
}

type DescribeScalingConfigurationsRequest struct {
	OwnerId                  *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount     *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId                 *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber               *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                 *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ScalingGroupId           *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	OwnerAccount             *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	ScalingConfigurationId   []*string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty" type:"Repeated"`
	ScalingConfigurationName []*string `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty" type:"Repeated"`
}

func (s DescribeScalingConfigurationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsRequest) SetOwnerId(v int64) *DescribeScalingConfigurationsRequest {
	s.OwnerId = &v
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

func (s *DescribeScalingConfigurationsRequest) SetRegionId(v string) *DescribeScalingConfigurationsRequest {
	s.RegionId = &v
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

func (s *DescribeScalingConfigurationsRequest) SetScalingGroupId(v string) *DescribeScalingConfigurationsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetOwnerAccount(v string) *DescribeScalingConfigurationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetScalingConfigurationId(v []*string) *DescribeScalingConfigurationsRequest {
	s.ScalingConfigurationId = v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetScalingConfigurationName(v []*string) *DescribeScalingConfigurationsRequest {
	s.ScalingConfigurationName = v
	return s
}

type DescribeScalingConfigurationsResponseBody struct {
	RequestId             *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber            *int32                                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount            *int32                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ScalingConfigurations *DescribeScalingConfigurationsResponseBodyScalingConfigurations `json:"ScalingConfigurations,omitempty" xml:"ScalingConfigurations,omitempty" type:"Struct"`
}

func (s DescribeScalingConfigurationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBody) SetRequestId(v string) *DescribeScalingConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBody) SetPageNumber(v int32) *DescribeScalingConfigurationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBody) SetPageSize(v int32) *DescribeScalingConfigurationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBody) SetTotalCount(v int32) *DescribeScalingConfigurationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBody) SetScalingConfigurations(v *DescribeScalingConfigurationsResponseBodyScalingConfigurations) *DescribeScalingConfigurationsResponseBody {
	s.ScalingConfigurations = v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurations struct {
	ScalingConfiguration []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration `json:"ScalingConfiguration,omitempty" xml:"ScalingConfiguration,omitempty" type:"Repeated"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurations) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurations) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetScalingConfiguration(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ScalingConfiguration = v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration struct {
	PrivatePoolOptions             *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationPrivatePoolOptions   `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" require:"true" type:"Struct"`
	DeploymentSetId                *string                                                                                                 `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	CreationTime                   *string                                                                                                 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ScalingConfigurationName       *string                                                                                                 `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	SystemDiskDescription          *string                                                                                                 `json:"SystemDiskDescription,omitempty" xml:"SystemDiskDescription,omitempty"`
	KeyPairName                    *string                                                                                                 `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	SecurityGroupId                *string                                                                                                 `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SystemDiskAutoSnapshotPolicyId *string                                                                                                 `json:"SystemDiskAutoSnapshotPolicyId,omitempty" xml:"SystemDiskAutoSnapshotPolicyId,omitempty"`
	SpotStrategy                   *string                                                                                                 `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	ScalingGroupId                 *string                                                                                                 `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	Affinity                       *string                                                                                                 `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	Tenancy                        *string                                                                                                 `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	SystemDiskSize                 *int32                                                                                                  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	Ipv6AddressCount               *int32                                                                                                  `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	SpotDuration                   *int32                                                                                                  `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	LifecycleState                 *string                                                                                                 `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	InstanceName                   *string                                                                                                 `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	SecurityEnhancementStrategy    *string                                                                                                 `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	UserData                       *string                                                                                                 `json:"UserData,omitempty" xml:"UserData,omitempty"`
	DedicatedHostId                *string                                                                                                 `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	InstanceGeneration             *string                                                                                                 `json:"InstanceGeneration,omitempty" xml:"InstanceGeneration,omitempty"`
	HpcClusterId                   *string                                                                                                 `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	PasswordInherit                *bool                                                                                                   `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	Memory                         *int32                                                                                                  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	ImageId                        *string                                                                                                 `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageFamily                    *string                                                                                                 `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	LoadBalancerWeight             *int32                                                                                                  `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	SystemDiskCategory             *string                                                                                                 `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	HostName                       *string                                                                                                 `json:"HostName,omitempty" xml:"HostName,omitempty"`
	SystemDiskName                 *string                                                                                                 `json:"SystemDiskName,omitempty" xml:"SystemDiskName,omitempty"`
	InternetMaxBandwidthOut        *int32                                                                                                  `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	InternetMaxBandwidthIn         *int32                                                                                                  `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	InstanceType                   *string                                                                                                 `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceDescription            *string                                                                                                 `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	IoOptimized                    *string                                                                                                 `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	RamRoleName                    *string                                                                                                 `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	SystemDiskPerformanceLevel     *string                                                                                                 `json:"SystemDiskPerformanceLevel,omitempty" xml:"SystemDiskPerformanceLevel,omitempty"`
	Cpu                            *int32                                                                                                  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	ResourceGroupId                *string                                                                                                 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ZoneId                         *string                                                                                                 `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	InternetChargeType             *string                                                                                                 `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	ImageName                      *string                                                                                                 `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ScalingConfigurationId         *string                                                                                                 `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	CreditSpecification            *string                                                                                                 `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	SpotInterruptionBehavior       *string                                                                                                 `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	DataDisks                      *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisks            `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Struct"`
	Tags                           *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationTags                 `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	SpotPriceLimit                 *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSpotPriceLimit       `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty" type:"Struct"`
	InstancePatternInfos           *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstancePatternInfos `json:"InstancePatternInfos,omitempty" xml:"InstancePatternInfos,omitempty" type:"Struct"`
	SystemDiskCategories           *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSystemDiskCategories `json:"SystemDiskCategories,omitempty" xml:"SystemDiskCategories,omitempty" type:"Struct"`
	WeightedCapacities             *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationWeightedCapacities   `json:"WeightedCapacities,omitempty" xml:"WeightedCapacities,omitempty" type:"Struct"`
	InstanceTypes                  *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstanceTypes        `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Struct"`
	SecurityGroupIds               *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSecurityGroupIds     `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Struct"`
	SchedulerOptions               *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSchedulerOptions     `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty" type:"Struct"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetPrivatePoolOptions(v *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationPrivatePoolOptions) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.PrivatePoolOptions = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetDeploymentSetId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.DeploymentSetId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetCreationTime(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.CreationTime = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetScalingConfigurationName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.ScalingConfigurationName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetSystemDiskDescription(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.SystemDiskDescription = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetKeyPairName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.KeyPairName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetSecurityGroupId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetSystemDiskAutoSnapshotPolicyId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.SystemDiskAutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetSpotStrategy(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetScalingGroupId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetAffinity(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.Affinity = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetTenancy(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.Tenancy = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetSystemDiskSize(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetIpv6AddressCount(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.Ipv6AddressCount = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetSpotDuration(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.SpotDuration = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetLifecycleState(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.LifecycleState = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetInstanceName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.InstanceName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetSecurityEnhancementStrategy(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetUserData(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.UserData = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetDedicatedHostId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetInstanceGeneration(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.InstanceGeneration = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetHpcClusterId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.HpcClusterId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetPasswordInherit(v bool) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.PasswordInherit = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetMemory(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.Memory = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetImageId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.ImageId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetImageFamily(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.ImageFamily = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetLoadBalancerWeight(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.LoadBalancerWeight = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetSystemDiskCategory(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetHostName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.HostName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetSystemDiskName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.SystemDiskName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetInternetMaxBandwidthOut(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetInternetMaxBandwidthIn(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetInstanceType(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.InstanceType = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetInstanceDescription(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetIoOptimized(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.IoOptimized = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetRamRoleName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.RamRoleName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetSystemDiskPerformanceLevel(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetCpu(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.Cpu = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetResourceGroupId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetZoneId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.ZoneId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetInternetChargeType(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetImageName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.ImageName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetScalingConfigurationId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetCreditSpecification(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.CreditSpecification = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetSpotInterruptionBehavior(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetDataDisks(v *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisks) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.DataDisks = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetTags(v *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationTags) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.Tags = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetSpotPriceLimit(v *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSpotPriceLimit) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.SpotPriceLimit = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetInstancePatternInfos(v *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstancePatternInfos) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.InstancePatternInfos = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetSystemDiskCategories(v *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSystemDiskCategories) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.SystemDiskCategories = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetWeightedCapacities(v *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationWeightedCapacities) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.WeightedCapacities = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetInstanceTypes(v *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstanceTypes) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.InstanceTypes = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetSecurityGroupIds(v *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSecurityGroupIds) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetSchedulerOptions(v *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSchedulerOptions) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.SchedulerOptions = v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationPrivatePoolOptions struct {
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationPrivatePoolOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationPrivatePoolOptions) SetId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationPrivatePoolOptions) SetMatchCriteria(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisks struct {
	DataDisk []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisks) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisks) SetDataDisk(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisks {
	s.DataDisk = v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk struct {
	PerformanceLevel     *string                                                                                                        `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Description          *string                                                                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	SnapshotId           *string                                                                                                        `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	Device               *string                                                                                                        `json:"Device,omitempty" xml:"Device,omitempty"`
	Size                 *int32                                                                                                         `json:"Size,omitempty" xml:"Size,omitempty"`
	DiskName             *string                                                                                                        `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	AutoSnapshotPolicyId *string                                                                                                        `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	Category             *string                                                                                                        `json:"Category,omitempty" xml:"Category,omitempty"`
	KMSKeyId             *string                                                                                                        `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	DeleteWithInstance   *bool                                                                                                          `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	Encrypted            *string                                                                                                        `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	Categories           *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDiskCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Struct"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) SetPerformanceLevel(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) SetDescription(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk {
	s.Description = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) SetSnapshotId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk {
	s.SnapshotId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) SetDevice(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk {
	s.Device = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) SetSize(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk {
	s.Size = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) SetDiskName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk {
	s.DiskName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) SetAutoSnapshotPolicyId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) SetCategory(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk {
	s.Category = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) SetKMSKeyId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk {
	s.KMSKeyId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) SetDeleteWithInstance(v bool) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) SetEncrypted(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk {
	s.Encrypted = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) SetCategories(v *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDiskCategories) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk {
	s.Categories = v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDiskCategories struct {
	Category []*string `json:"Category,omitempty" xml:"Category,omitempty" type:"Repeated"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDiskCategories) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDiskCategories) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDiskCategories) SetCategory(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDiskCategories {
	s.Category = v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationTags struct {
	Tag []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationTags) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationTags) SetTag(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationTagsTag) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationTags {
	s.Tag = v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationTagsTag) SetKey(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationTagsTag) SetValue(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationTagsTag {
	s.Value = &v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSpotPriceLimit struct {
	SpotPriceModel []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSpotPriceLimitSpotPriceModel `json:"SpotPriceModel,omitempty" xml:"SpotPriceModel,omitempty" type:"Repeated"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSpotPriceLimit) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSpotPriceLimit) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSpotPriceLimit) SetSpotPriceModel(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSpotPriceLimitSpotPriceModel) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSpotPriceLimit {
	s.SpotPriceModel = v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSpotPriceLimitSpotPriceModel struct {
	InstanceType *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	PriceLimit   *float32 `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSpotPriceLimitSpotPriceModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSpotPriceLimitSpotPriceModel) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSpotPriceLimitSpotPriceModel) SetInstanceType(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSpotPriceLimitSpotPriceModel {
	s.InstanceType = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSpotPriceLimitSpotPriceModel) SetPriceLimit(v float32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSpotPriceLimitSpotPriceModel {
	s.PriceLimit = &v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstancePatternInfos struct {
	InstancePatternInfo []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstancePatternInfosInstancePatternInfo `json:"InstancePatternInfo,omitempty" xml:"InstancePatternInfo,omitempty" type:"Repeated"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstancePatternInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstancePatternInfos) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstancePatternInfos) SetInstancePatternInfo(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstancePatternInfosInstancePatternInfo) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstancePatternInfos {
	s.InstancePatternInfo = v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstancePatternInfosInstancePatternInfo struct {
	MaxPrice            *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	Cores               *int32   `json:"Cores,omitempty" xml:"Cores,omitempty"`
	Memory              *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	InstanceFamilyLevel *string  `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstancePatternInfosInstancePatternInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstancePatternInfosInstancePatternInfo) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstancePatternInfosInstancePatternInfo) SetMaxPrice(v float32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstancePatternInfosInstancePatternInfo {
	s.MaxPrice = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstancePatternInfosInstancePatternInfo) SetCores(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstancePatternInfosInstancePatternInfo {
	s.Cores = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstancePatternInfosInstancePatternInfo) SetMemory(v float32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstancePatternInfosInstancePatternInfo {
	s.Memory = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstancePatternInfosInstancePatternInfo) SetInstanceFamilyLevel(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstancePatternInfosInstancePatternInfo {
	s.InstanceFamilyLevel = &v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSystemDiskCategories struct {
	SystemDiskCategory []*string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty" type:"Repeated"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSystemDiskCategories) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSystemDiskCategories) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSystemDiskCategories) SetSystemDiskCategory(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSystemDiskCategories {
	s.SystemDiskCategory = v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationWeightedCapacities struct {
	WeightedCapacity []*string `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty" type:"Repeated"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationWeightedCapacities) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationWeightedCapacities) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationWeightedCapacities) SetWeightedCapacity(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationWeightedCapacities {
	s.WeightedCapacity = v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstanceTypes struct {
	InstanceType []*string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Repeated"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstanceTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstanceTypes) SetInstanceType(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationInstanceTypes {
	s.InstanceType = v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSecurityGroupIds struct {
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSecurityGroupIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSecurityGroupIds) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSecurityGroupIds) SetSecurityGroupId(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSecurityGroupIds {
	s.SecurityGroupId = v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSchedulerOptions struct {
	ManagedPrivateSpaceId *string `json:"ManagedPrivateSpaceId,omitempty" xml:"ManagedPrivateSpaceId,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSchedulerOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSchedulerOptions) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSchedulerOptions) SetManagedPrivateSpaceId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationSchedulerOptions {
	s.ManagedPrivateSpaceId = &v
	return s
}

type DescribeScalingConfigurationsResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScalingConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeScalingConfigurationsResponse) SetBody(v *DescribeScalingConfigurationsResponseBody) *DescribeScalingConfigurationsResponse {
	s.Body = v
	return s
}

type DescribeScalingInstancesRequest struct {
	OwnerId                *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId               *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScalingGroupId         *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScalingConfigurationId *string   `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	HealthStatus           *string   `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	LifecycleState         *string   `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	CreationType           *string   `json:"CreationType,omitempty" xml:"CreationType,omitempty"`
	PageNumber             *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize               *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	OwnerAccount           *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	ScalingActivityId      *string   `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	InstanceId             []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s DescribeScalingInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingInstancesRequest) SetOwnerId(v int64) *DescribeScalingInstancesRequest {
	s.OwnerId = &v
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

func (s *DescribeScalingInstancesRequest) SetRegionId(v string) *DescribeScalingInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetScalingGroupId(v string) *DescribeScalingInstancesRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetScalingConfigurationId(v string) *DescribeScalingInstancesRequest {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetHealthStatus(v string) *DescribeScalingInstancesRequest {
	s.HealthStatus = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetLifecycleState(v string) *DescribeScalingInstancesRequest {
	s.LifecycleState = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetCreationType(v string) *DescribeScalingInstancesRequest {
	s.CreationType = &v
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

func (s *DescribeScalingInstancesRequest) SetOwnerAccount(v string) *DescribeScalingInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetScalingActivityId(v string) *DescribeScalingInstancesRequest {
	s.ScalingActivityId = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetInstanceId(v []*string) *DescribeScalingInstancesRequest {
	s.InstanceId = v
	return s
}

type DescribeScalingInstancesResponseBody struct {
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize         *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber       *int32                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalSpotCount   *int32                                                `json:"TotalSpotCount,omitempty" xml:"TotalSpotCount,omitempty"`
	TotalCount       *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ScalingInstances *DescribeScalingInstancesResponseBodyScalingInstances `json:"ScalingInstances,omitempty" xml:"ScalingInstances,omitempty" type:"Struct"`
}

func (s DescribeScalingInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingInstancesResponseBody) SetRequestId(v string) *DescribeScalingInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBody) SetPageSize(v int32) *DescribeScalingInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingInstancesResponseBody) SetPageNumber(v int32) *DescribeScalingInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingInstancesResponseBody) SetTotalSpotCount(v int32) *DescribeScalingInstancesResponseBody {
	s.TotalSpotCount = &v
	return s
}

func (s *DescribeScalingInstancesResponseBody) SetTotalCount(v int32) *DescribeScalingInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeScalingInstancesResponseBody) SetScalingInstances(v *DescribeScalingInstancesResponseBodyScalingInstances) *DescribeScalingInstancesResponseBody {
	s.ScalingInstances = v
	return s
}

type DescribeScalingInstancesResponseBodyScalingInstances struct {
	ScalingInstance []*DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance `json:"ScalingInstance,omitempty" xml:"ScalingInstance,omitempty" type:"Repeated"`
}

func (s DescribeScalingInstancesResponseBodyScalingInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingInstancesResponseBodyScalingInstances) GoString() string {
	return s.String()
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetScalingInstance(v []*DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.ScalingInstance = v
	return s
}

type DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance struct {
	CreationTime           *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	LoadBalancerWeight     *int32  `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	LaunchTemplateId       *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SpotStrategy           *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	LaunchTemplateVersion  *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	HealthStatus           *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	ScalingGroupId         *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	WarmupState            *string `json:"WarmupState,omitempty" xml:"WarmupState,omitempty"`
	LifecycleState         *string `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	CreationType           *string `json:"CreationType,omitempty" xml:"CreationType,omitempty"`
	ZoneId                 *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	Entrusted              *bool   `json:"Entrusted,omitempty" xml:"Entrusted,omitempty"`
	WeightedCapacity       *int32  `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
	CreatedTime            *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	ScalingActivityId      *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) GoString() string {
	return s.String()
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetCreationTime(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetLoadBalancerWeight(v int32) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.LoadBalancerWeight = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetLaunchTemplateId(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.LaunchTemplateId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetInstanceId(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetSpotStrategy(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetLaunchTemplateVersion(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetHealthStatus(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.HealthStatus = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetScalingGroupId(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetWarmupState(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.WarmupState = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetLifecycleState(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.LifecycleState = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetCreationType(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.CreationType = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetZoneId(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetScalingConfigurationId(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetEntrusted(v bool) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.Entrusted = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetWeightedCapacity(v int32) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.WeightedCapacity = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetCreatedTime(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.CreatedTime = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetScalingActivityId(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.ScalingActivityId = &v
	return s
}

type DescribeScalingInstancesResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScalingInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeScalingInstancesResponse) SetBody(v *DescribeScalingInstancesResponseBody) *DescribeScalingInstancesResponse {
	s.Body = v
	return s
}

type DescribeScalingRulesRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber           *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScalingRuleType      *string   `json:"ScalingRuleType,omitempty" xml:"ScalingRuleType,omitempty"`
	ShowAlarmRules       *bool     `json:"ShowAlarmRules,omitempty" xml:"ShowAlarmRules,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	ScalingRuleId        []*string `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty" type:"Repeated"`
	ScalingRuleName      []*string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty" type:"Repeated"`
	ScalingRuleAri       []*string `json:"ScalingRuleAri,omitempty" xml:"ScalingRuleAri,omitempty" type:"Repeated"`
}

func (s DescribeScalingRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesRequest) SetOwnerId(v int64) *DescribeScalingRulesRequest {
	s.OwnerId = &v
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

func (s *DescribeScalingRulesRequest) SetRegionId(v string) *DescribeScalingRulesRequest {
	s.RegionId = &v
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

func (s *DescribeScalingRulesRequest) SetScalingGroupId(v string) *DescribeScalingRulesRequest {
	s.ScalingGroupId = &v
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

func (s *DescribeScalingRulesRequest) SetOwnerAccount(v string) *DescribeScalingRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetScalingRuleId(v []*string) *DescribeScalingRulesRequest {
	s.ScalingRuleId = v
	return s
}

func (s *DescribeScalingRulesRequest) SetScalingRuleName(v []*string) *DescribeScalingRulesRequest {
	s.ScalingRuleName = v
	return s
}

func (s *DescribeScalingRulesRequest) SetScalingRuleAri(v []*string) *DescribeScalingRulesRequest {
	s.ScalingRuleAri = v
	return s
}

type DescribeScalingRulesResponseBody struct {
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber   *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount   *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ScalingRules *DescribeScalingRulesResponseBodyScalingRules `json:"ScalingRules,omitempty" xml:"ScalingRules,omitempty" type:"Struct"`
}

func (s DescribeScalingRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBody) SetRequestId(v string) *DescribeScalingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingRulesResponseBody) SetPageNumber(v int32) *DescribeScalingRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingRulesResponseBody) SetPageSize(v int32) *DescribeScalingRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingRulesResponseBody) SetTotalCount(v int32) *DescribeScalingRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeScalingRulesResponseBody) SetScalingRules(v *DescribeScalingRulesResponseBodyScalingRules) *DescribeScalingRulesResponseBody {
	s.ScalingRules = v
	return s
}

type DescribeScalingRulesResponseBodyScalingRules struct {
	ScalingRule []*DescribeScalingRulesResponseBodyScalingRulesScalingRule `json:"ScalingRule,omitempty" xml:"ScalingRule,omitempty" type:"Repeated"`
}

func (s DescribeScalingRulesResponseBodyScalingRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRules) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetScalingRule(v []*DescribeScalingRulesResponseBodyScalingRulesScalingRule) *DescribeScalingRulesResponseBodyScalingRules {
	s.ScalingRule = v
	return s
}

type DescribeScalingRulesResponseBodyScalingRulesScalingRule struct {
	MetricName               *string                                                                 `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	AdjustmentType           *string                                                                 `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	InitialMaxSize           *int32                                                                  `json:"InitialMaxSize,omitempty" xml:"InitialMaxSize,omitempty"`
	EstimatedInstanceWarmup  *int32                                                                  `json:"EstimatedInstanceWarmup,omitempty" xml:"EstimatedInstanceWarmup,omitempty"`
	ScaleOutEvaluationCount  *int32                                                                  `json:"ScaleOutEvaluationCount,omitempty" xml:"ScaleOutEvaluationCount,omitempty"`
	PredictiveScalingMode    *string                                                                 `json:"PredictiveScalingMode,omitempty" xml:"PredictiveScalingMode,omitempty"`
	MinAdjustmentMagnitude   *int32                                                                  `json:"MinAdjustmentMagnitude,omitempty" xml:"MinAdjustmentMagnitude,omitempty"`
	ScalingRuleAri           *string                                                                 `json:"ScalingRuleAri,omitempty" xml:"ScalingRuleAri,omitempty"`
	PredictiveTaskBufferTime *int32                                                                  `json:"PredictiveTaskBufferTime,omitempty" xml:"PredictiveTaskBufferTime,omitempty"`
	MinSize                  *int32                                                                  `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	ScalingGroupId           *string                                                                 `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	PredictiveValueBehavior  *string                                                                 `json:"PredictiveValueBehavior,omitempty" xml:"PredictiveValueBehavior,omitempty"`
	TargetValue              *float32                                                                `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
	Cooldown                 *int32                                                                  `json:"Cooldown,omitempty" xml:"Cooldown,omitempty"`
	MaxSize                  *int32                                                                  `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	PredictiveValueBuffer    *int32                                                                  `json:"PredictiveValueBuffer,omitempty" xml:"PredictiveValueBuffer,omitempty"`
	ScalingRuleType          *string                                                                 `json:"ScalingRuleType,omitempty" xml:"ScalingRuleType,omitempty"`
	AdjustmentValue          *int32                                                                  `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	ScaleInEvaluationCount   *int32                                                                  `json:"ScaleInEvaluationCount,omitempty" xml:"ScaleInEvaluationCount,omitempty"`
	DisableScaleIn           *bool                                                                   `json:"DisableScaleIn,omitempty" xml:"DisableScaleIn,omitempty"`
	ScalingRuleName          *string                                                                 `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	ScalingRuleId            *string                                                                 `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty"`
	Alarms                   *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarms          `json:"Alarms,omitempty" xml:"Alarms,omitempty" type:"Struct"`
	StepAdjustments          *DescribeScalingRulesResponseBodyScalingRulesScalingRuleStepAdjustments `json:"StepAdjustments,omitempty" xml:"StepAdjustments,omitempty" type:"Struct"`
}

func (s DescribeScalingRulesResponseBodyScalingRulesScalingRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRulesScalingRule) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetMetricName(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.MetricName = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetAdjustmentType(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.AdjustmentType = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetInitialMaxSize(v int32) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.InitialMaxSize = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetEstimatedInstanceWarmup(v int32) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.EstimatedInstanceWarmup = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetScaleOutEvaluationCount(v int32) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.ScaleOutEvaluationCount = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetPredictiveScalingMode(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.PredictiveScalingMode = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetMinAdjustmentMagnitude(v int32) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.MinAdjustmentMagnitude = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetScalingRuleAri(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.ScalingRuleAri = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetPredictiveTaskBufferTime(v int32) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.PredictiveTaskBufferTime = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetMinSize(v int32) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.MinSize = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetScalingGroupId(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetPredictiveValueBehavior(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.PredictiveValueBehavior = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetTargetValue(v float32) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.TargetValue = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetCooldown(v int32) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.Cooldown = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetMaxSize(v int32) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.MaxSize = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetPredictiveValueBuffer(v int32) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.PredictiveValueBuffer = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetScalingRuleType(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.ScalingRuleType = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetAdjustmentValue(v int32) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.AdjustmentValue = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetScaleInEvaluationCount(v int32) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.ScaleInEvaluationCount = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetDisableScaleIn(v bool) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.DisableScaleIn = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetScalingRuleName(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.ScalingRuleName = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetScalingRuleId(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.ScalingRuleId = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetAlarms(v *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarms) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.Alarms = v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetStepAdjustments(v *DescribeScalingRulesResponseBodyScalingRulesScalingRuleStepAdjustments) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.StepAdjustments = v
	return s
}

type DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarms struct {
	Alarm []*DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm `json:"Alarm,omitempty" xml:"Alarm,omitempty" type:"Repeated"`
}

func (s DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarms) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarms) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarms) SetAlarm(v []*DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm) *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarms {
	s.Alarm = v
	return s
}

type DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm struct {
	AlarmTaskId        *string                                                                       `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	ComparisonOperator *string                                                                       `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	MetricName         *string                                                                       `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	EvaluationCount    *int32                                                                        `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	AlarmTaskName      *string                                                                       `json:"AlarmTaskName,omitempty" xml:"AlarmTaskName,omitempty"`
	MetricType         *string                                                                       `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Threshold          *float32                                                                      `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Statistics         *string                                                                       `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Dimensions         *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarmDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Struct"`
}

func (s DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm) SetAlarmTaskId(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm {
	s.AlarmTaskId = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm) SetComparisonOperator(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm) SetMetricName(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm {
	s.MetricName = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm) SetEvaluationCount(v int32) *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm) SetAlarmTaskName(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm {
	s.AlarmTaskName = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm) SetMetricType(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm {
	s.MetricType = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm) SetThreshold(v float32) *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm {
	s.Threshold = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm) SetStatistics(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm {
	s.Statistics = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm) SetDimensions(v *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarmDimensions) *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarm {
	s.Dimensions = v
	return s
}

type DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarmDimensions struct {
	Dimension []*DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarmDimensionsDimension `json:"Dimension,omitempty" xml:"Dimension,omitempty" type:"Repeated"`
}

func (s DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarmDimensions) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarmDimensions) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarmDimensions) SetDimension(v []*DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarmDimensionsDimension) *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarmDimensions {
	s.Dimension = v
	return s
}

type DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarmDimensionsDimension struct {
	DimensionKey   *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
}

func (s DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarmDimensionsDimension) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarmDimensionsDimension) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarmDimensionsDimension) SetDimensionKey(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarmDimensionsDimension {
	s.DimensionKey = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarmDimensionsDimension) SetDimensionValue(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRuleAlarmsAlarmDimensionsDimension {
	s.DimensionValue = &v
	return s
}

type DescribeScalingRulesResponseBodyScalingRulesScalingRuleStepAdjustments struct {
	StepAdjustment []*DescribeScalingRulesResponseBodyScalingRulesScalingRuleStepAdjustmentsStepAdjustment `json:"StepAdjustment,omitempty" xml:"StepAdjustment,omitempty" type:"Repeated"`
}

func (s DescribeScalingRulesResponseBodyScalingRulesScalingRuleStepAdjustments) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRulesScalingRuleStepAdjustments) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRuleStepAdjustments) SetStepAdjustment(v []*DescribeScalingRulesResponseBodyScalingRulesScalingRuleStepAdjustmentsStepAdjustment) *DescribeScalingRulesResponseBodyScalingRulesScalingRuleStepAdjustments {
	s.StepAdjustment = v
	return s
}

type DescribeScalingRulesResponseBodyScalingRulesScalingRuleStepAdjustmentsStepAdjustment struct {
	MetricIntervalUpperBound *float32 `json:"MetricIntervalUpperBound,omitempty" xml:"MetricIntervalUpperBound,omitempty"`
	ScalingAdjustment        *int32   `json:"ScalingAdjustment,omitempty" xml:"ScalingAdjustment,omitempty"`
	MetricIntervalLowerBound *float32 `json:"MetricIntervalLowerBound,omitempty" xml:"MetricIntervalLowerBound,omitempty"`
}

func (s DescribeScalingRulesResponseBodyScalingRulesScalingRuleStepAdjustmentsStepAdjustment) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRulesScalingRuleStepAdjustmentsStepAdjustment) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRuleStepAdjustmentsStepAdjustment) SetMetricIntervalUpperBound(v float32) *DescribeScalingRulesResponseBodyScalingRulesScalingRuleStepAdjustmentsStepAdjustment {
	s.MetricIntervalUpperBound = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRuleStepAdjustmentsStepAdjustment) SetScalingAdjustment(v int32) *DescribeScalingRulesResponseBodyScalingRulesScalingRuleStepAdjustmentsStepAdjustment {
	s.ScalingAdjustment = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRuleStepAdjustmentsStepAdjustment) SetMetricIntervalLowerBound(v float32) *DescribeScalingRulesResponseBodyScalingRulesScalingRuleStepAdjustmentsStepAdjustment {
	s.MetricIntervalLowerBound = &v
	return s
}

type DescribeScalingRulesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScalingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeScalingRulesResponse) SetBody(v *DescribeScalingRulesResponseBody) *DescribeScalingRulesResponse {
	s.Body = v
	return s
}

type DescribeScheduledTasksRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber           *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScheduledAction      []*string `json:"ScheduledAction,omitempty" xml:"ScheduledAction,omitempty" type:"Repeated"`
	ScheduledTaskId      []*string `json:"ScheduledTaskId,omitempty" xml:"ScheduledTaskId,omitempty" type:"Repeated"`
	ScheduledTaskName    []*string `json:"ScheduledTaskName,omitempty" xml:"ScheduledTaskName,omitempty" type:"Repeated"`
}

func (s DescribeScheduledTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduledTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksRequest) SetOwnerId(v int64) *DescribeScheduledTasksRequest {
	s.OwnerId = &v
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

func (s *DescribeScheduledTasksRequest) SetRegionId(v string) *DescribeScheduledTasksRequest {
	s.RegionId = &v
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

func (s *DescribeScheduledTasksRequest) SetOwnerAccount(v string) *DescribeScheduledTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetScalingGroupId(v string) *DescribeScheduledTasksRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetScheduledAction(v []*string) *DescribeScheduledTasksRequest {
	s.ScheduledAction = v
	return s
}

func (s *DescribeScheduledTasksRequest) SetScheduledTaskId(v []*string) *DescribeScheduledTasksRequest {
	s.ScheduledTaskId = v
	return s
}

func (s *DescribeScheduledTasksRequest) SetScheduledTaskName(v []*string) *DescribeScheduledTasksRequest {
	s.ScheduledTaskName = v
	return s
}

type DescribeScheduledTasksResponseBody struct {
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber     *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ScheduledTasks *DescribeScheduledTasksResponseBodyScheduledTasks `json:"ScheduledTasks,omitempty" xml:"ScheduledTasks,omitempty" type:"Struct"`
}

func (s DescribeScheduledTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduledTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksResponseBody) SetRequestId(v string) *DescribeScheduledTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetPageNumber(v int32) *DescribeScheduledTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetPageSize(v int32) *DescribeScheduledTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetTotalCount(v int32) *DescribeScheduledTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetScheduledTasks(v *DescribeScheduledTasksResponseBodyScheduledTasks) *DescribeScheduledTasksResponseBody {
	s.ScheduledTasks = v
	return s
}

type DescribeScheduledTasksResponseBodyScheduledTasks struct {
	ScheduledTask []*DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask `json:"ScheduledTask,omitempty" xml:"ScheduledTask,omitempty" type:"Repeated"`
}

func (s DescribeScheduledTasksResponseBodyScheduledTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduledTasksResponseBodyScheduledTasks) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetScheduledTask(v []*DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.ScheduledTask = v
	return s
}

type DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask struct {
	TaskEnabled          *bool   `json:"TaskEnabled,omitempty" xml:"TaskEnabled,omitempty"`
	RecurrenceValue      *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	RecurrenceType       *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	MaxValue             *int32  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	ScheduledTaskName    *string `json:"ScheduledTaskName,omitempty" xml:"ScheduledTaskName,omitempty"`
	RecurrenceEndTime    *string `json:"RecurrenceEndTime,omitempty" xml:"RecurrenceEndTime,omitempty"`
	DesiredCapacity      *int32  `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	ScheduledTaskId      *string `json:"ScheduledTaskId,omitempty" xml:"ScheduledTaskId,omitempty"`
	MinValue             *int32  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	LaunchExpirationTime *int32  `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ScheduledAction      *string `json:"ScheduledAction,omitempty" xml:"ScheduledAction,omitempty"`
	LaunchTime           *string `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
}

func (s DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetTaskEnabled(v bool) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.TaskEnabled = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetRecurrenceValue(v string) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.RecurrenceValue = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetRecurrenceType(v string) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.RecurrenceType = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetMaxValue(v int32) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.MaxValue = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetScheduledTaskName(v string) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.ScheduledTaskName = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetRecurrenceEndTime(v string) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.RecurrenceEndTime = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetDesiredCapacity(v int32) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.DesiredCapacity = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetScheduledTaskId(v string) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.ScheduledTaskId = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetMinValue(v int32) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.MinValue = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetScalingGroupId(v string) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetLaunchExpirationTime(v int32) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.LaunchExpirationTime = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetDescription(v string) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.Description = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetScheduledAction(v string) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.ScheduledAction = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetLaunchTime(v string) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.LaunchTime = &v
	return s
}

type DescribeScheduledTasksResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScheduledTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeScheduledTasksResponse) SetBody(v *DescribeScheduledTasksResponseBody) *DescribeScheduledTasksResponse {
	s.Body = v
	return s
}

type DetachAlbServerGroupsRequest struct {
	OwnerId              *int64                                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	RegionId             *string                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScalingGroupId       *string                                       `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ClientToken          *string                                       `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ForceDetach          *bool                                         `json:"ForceDetach,omitempty" xml:"ForceDetach,omitempty"`
	AlbServerGroup       []*DetachAlbServerGroupsRequestAlbServerGroup `json:"AlbServerGroup,omitempty" xml:"AlbServerGroup,omitempty" type:"Repeated"`
}

func (s DetachAlbServerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachAlbServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *DetachAlbServerGroupsRequest) SetOwnerId(v int64) *DetachAlbServerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachAlbServerGroupsRequest) SetResourceOwnerAccount(v string) *DetachAlbServerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachAlbServerGroupsRequest) SetRegionId(v string) *DetachAlbServerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DetachAlbServerGroupsRequest) SetScalingGroupId(v string) *DetachAlbServerGroupsRequest {
	s.ScalingGroupId = &v
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

func (s *DetachAlbServerGroupsRequest) SetAlbServerGroup(v []*DetachAlbServerGroupsRequestAlbServerGroup) *DetachAlbServerGroupsRequest {
	s.AlbServerGroup = v
	return s
}

type DetachAlbServerGroupsRequestAlbServerGroup struct {
	AlbServerGroupId *string `json:"AlbServerGroupId,omitempty" xml:"AlbServerGroupId,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DetachAlbServerGroupsRequestAlbServerGroup) String() string {
	return tea.Prettify(s)
}

func (s DetachAlbServerGroupsRequestAlbServerGroup) GoString() string {
	return s.String()
}

func (s *DetachAlbServerGroupsRequestAlbServerGroup) SetAlbServerGroupId(v string) *DetachAlbServerGroupsRequestAlbServerGroup {
	s.AlbServerGroupId = &v
	return s
}

func (s *DetachAlbServerGroupsRequestAlbServerGroup) SetPort(v int32) *DetachAlbServerGroupsRequestAlbServerGroup {
	s.Port = &v
	return s
}

type DetachAlbServerGroupsResponseBody struct {
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachAlbServerGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachAlbServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DetachAlbServerGroupsResponseBody) SetScalingActivityId(v string) *DetachAlbServerGroupsResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *DetachAlbServerGroupsResponseBody) SetRequestId(v string) *DetachAlbServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

type DetachAlbServerGroupsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachAlbServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetachAlbServerGroupsResponse) SetBody(v *DetachAlbServerGroupsResponseBody) *DetachAlbServerGroupsResponse {
	s.Body = v
	return s
}

type DetachDBInstancesRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ForceDetach          *bool     `json:"ForceDetach,omitempty" xml:"ForceDetach,omitempty"`
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstance           []*string `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Repeated"`
}

func (s DetachDBInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *DetachDBInstancesRequest) SetOwnerId(v int64) *DetachDBInstancesRequest {
	s.OwnerId = &v
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

func (s *DetachDBInstancesRequest) SetForceDetach(v bool) *DetachDBInstancesRequest {
	s.ForceDetach = &v
	return s
}

func (s *DetachDBInstancesRequest) SetClientToken(v string) *DetachDBInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachDBInstancesRequest) SetDBInstance(v []*string) *DetachDBInstancesRequest {
	s.DBInstance = v
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachDBInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetachDBInstancesResponse) SetBody(v *DetachDBInstancesResponseBody) *DetachDBInstancesResponse {
	s.Body = v
	return s
}

type DetachInstancesRequest struct {
	OwnerId                 *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount    *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupId          *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	OwnerAccount            *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DecreaseDesiredCapacity *bool     `json:"DecreaseDesiredCapacity,omitempty" xml:"DecreaseDesiredCapacity,omitempty"`
	DetachOption            *string   `json:"DetachOption,omitempty" xml:"DetachOption,omitempty"`
	InstanceId              []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s DetachInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachInstancesRequest) GoString() string {
	return s.String()
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

func (s *DetachInstancesRequest) SetOwnerAccount(v string) *DetachInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DetachInstancesRequest) SetDecreaseDesiredCapacity(v bool) *DetachInstancesRequest {
	s.DecreaseDesiredCapacity = &v
	return s
}

func (s *DetachInstancesRequest) SetDetachOption(v string) *DetachInstancesRequest {
	s.DetachOption = &v
	return s
}

func (s *DetachInstancesRequest) SetInstanceId(v []*string) *DetachInstancesRequest {
	s.InstanceId = v
	return s
}

type DetachInstancesResponseBody struct {
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DetachInstancesResponseBody) SetScalingActivityId(v string) *DetachInstancesResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *DetachInstancesResponseBody) SetRequestId(v string) *DetachInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DetachInstancesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetachInstancesResponse) SetBody(v *DetachInstancesResponseBody) *DetachInstancesResponse {
	s.Body = v
	return s
}

type DetachLoadBalancersRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ForceDetach          *bool     `json:"ForceDetach,omitempty" xml:"ForceDetach,omitempty"`
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Async                *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
	LoadBalancer         []*string `json:"LoadBalancer,omitempty" xml:"LoadBalancer,omitempty" type:"Repeated"`
}

func (s DetachLoadBalancersRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachLoadBalancersRequest) GoString() string {
	return s.String()
}

func (s *DetachLoadBalancersRequest) SetOwnerId(v int64) *DetachLoadBalancersRequest {
	s.OwnerId = &v
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

func (s *DetachLoadBalancersRequest) SetForceDetach(v bool) *DetachLoadBalancersRequest {
	s.ForceDetach = &v
	return s
}

func (s *DetachLoadBalancersRequest) SetClientToken(v string) *DetachLoadBalancersRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachLoadBalancersRequest) SetAsync(v bool) *DetachLoadBalancersRequest {
	s.Async = &v
	return s
}

func (s *DetachLoadBalancersRequest) SetLoadBalancer(v []*string) *DetachLoadBalancersRequest {
	s.LoadBalancer = v
	return s
}

type DetachLoadBalancersResponseBody struct {
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachLoadBalancersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachLoadBalancersResponseBody) GoString() string {
	return s.String()
}

func (s *DetachLoadBalancersResponseBody) SetScalingActivityId(v string) *DetachLoadBalancersResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *DetachLoadBalancersResponseBody) SetRequestId(v string) *DetachLoadBalancersResponseBody {
	s.RequestId = &v
	return s
}

type DetachLoadBalancersResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachLoadBalancersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetachLoadBalancersResponse) SetBody(v *DetachLoadBalancersResponseBody) *DetachLoadBalancersResponse {
	s.Body = v
	return s
}

type DetachVServerGroupsRequest struct {
	OwnerId              *int64                                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	RegionId             *string                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScalingGroupId       *string                                   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ClientToken          *string                                   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ForceDetach          *bool                                     `json:"ForceDetach,omitempty" xml:"ForceDetach,omitempty"`
	VServerGroup         []*DetachVServerGroupsRequestVServerGroup `json:"VServerGroup,omitempty" xml:"VServerGroup,omitempty" type:"Repeated"`
}

func (s DetachVServerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachVServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *DetachVServerGroupsRequest) SetOwnerId(v int64) *DetachVServerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachVServerGroupsRequest) SetResourceOwnerAccount(v string) *DetachVServerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachVServerGroupsRequest) SetRegionId(v string) *DetachVServerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DetachVServerGroupsRequest) SetScalingGroupId(v string) *DetachVServerGroupsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DetachVServerGroupsRequest) SetClientToken(v string) *DetachVServerGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachVServerGroupsRequest) SetForceDetach(v bool) *DetachVServerGroupsRequest {
	s.ForceDetach = &v
	return s
}

func (s *DetachVServerGroupsRequest) SetVServerGroup(v []*DetachVServerGroupsRequestVServerGroup) *DetachVServerGroupsRequest {
	s.VServerGroup = v
	return s
}

type DetachVServerGroupsRequestVServerGroup struct {
	VServerGroupAttribute []*DetachVServerGroupsRequestVServerGroupVServerGroupAttribute `json:"VServerGroupAttribute,omitempty" xml:"VServerGroupAttribute,omitempty" type:"Repeated"`
	LoadBalancerId        *string                                                        `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s DetachVServerGroupsRequestVServerGroup) String() string {
	return tea.Prettify(s)
}

func (s DetachVServerGroupsRequestVServerGroup) GoString() string {
	return s.String()
}

func (s *DetachVServerGroupsRequestVServerGroup) SetVServerGroupAttribute(v []*DetachVServerGroupsRequestVServerGroupVServerGroupAttribute) *DetachVServerGroupsRequestVServerGroup {
	s.VServerGroupAttribute = v
	return s
}

func (s *DetachVServerGroupsRequestVServerGroup) SetLoadBalancerId(v string) *DetachVServerGroupsRequestVServerGroup {
	s.LoadBalancerId = &v
	return s
}

type DetachVServerGroupsRequestVServerGroupVServerGroupAttribute struct {
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	Port           *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DetachVServerGroupsRequestVServerGroupVServerGroupAttribute) String() string {
	return tea.Prettify(s)
}

func (s DetachVServerGroupsRequestVServerGroupVServerGroupAttribute) GoString() string {
	return s.String()
}

func (s *DetachVServerGroupsRequestVServerGroupVServerGroupAttribute) SetVServerGroupId(v string) *DetachVServerGroupsRequestVServerGroupVServerGroupAttribute {
	s.VServerGroupId = &v
	return s
}

func (s *DetachVServerGroupsRequestVServerGroupVServerGroupAttribute) SetPort(v int32) *DetachVServerGroupsRequestVServerGroupVServerGroupAttribute {
	s.Port = &v
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachVServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetachVServerGroupsResponse) SetBody(v *DetachVServerGroupsResponseBody) *DetachVServerGroupsResponse {
	s.Body = v
	return s
}

type DisableAlarmRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AlarmTaskId          *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
}

func (s DisableAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableAlarmRequest) GoString() string {
	return s.String()
}

func (s *DisableAlarmRequest) SetOwnerId(v int64) *DisableAlarmRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableAlarmRequest) SetResourceOwnerAccount(v string) *DisableAlarmRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisableAlarmRequest) SetRegionId(v string) *DisableAlarmRequest {
	s.RegionId = &v
	return s
}

func (s *DisableAlarmRequest) SetAlarmTaskId(v string) *DisableAlarmRequest {
	s.AlarmTaskId = &v
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DisableAlarmResponse) SetBody(v *DisableAlarmResponseBody) *DisableAlarmResponse {
	s.Body = v
	return s
}

type DisableScalingGroupRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s DisableScalingGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableScalingGroupRequest) GoString() string {
	return s.String()
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

func (s *DisableScalingGroupRequest) SetOwnerAccount(v string) *DisableScalingGroupRequest {
	s.OwnerAccount = &v
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DisableScalingGroupResponse) SetBody(v *DisableScalingGroupResponseBody) *DisableScalingGroupResponse {
	s.Body = v
	return s
}

type EnableAlarmRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AlarmTaskId          *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
}

func (s EnableAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableAlarmRequest) GoString() string {
	return s.String()
}

func (s *EnableAlarmRequest) SetOwnerId(v int64) *EnableAlarmRequest {
	s.OwnerId = &v
	return s
}

func (s *EnableAlarmRequest) SetResourceOwnerAccount(v string) *EnableAlarmRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *EnableAlarmRequest) SetRegionId(v string) *EnableAlarmRequest {
	s.RegionId = &v
	return s
}

func (s *EnableAlarmRequest) SetAlarmTaskId(v string) *EnableAlarmRequest {
	s.AlarmTaskId = &v
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *EnableAlarmResponse) SetBody(v *EnableAlarmResponseBody) *EnableAlarmResponse {
	s.Body = v
	return s
}

type EnableScalingGroupRequest struct {
	OwnerId                      *int64                                             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount         *string                                            `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId              *int64                                             `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupId               *string                                            `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ActiveScalingConfigurationId *string                                            `json:"ActiveScalingConfigurationId,omitempty" xml:"ActiveScalingConfigurationId,omitempty"`
	OwnerAccount                 *string                                            `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	LaunchTemplateId             *string                                            `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	LaunchTemplateVersion        *string                                            `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	InstanceId                   []*string                                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	LoadBalancerWeight           []*int                                             `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty" type:"Repeated"`
	LaunchTemplateOverride       []*EnableScalingGroupRequestLaunchTemplateOverride `json:"LaunchTemplateOverride,omitempty" xml:"LaunchTemplateOverride,omitempty" type:"Repeated"`
}

func (s EnableScalingGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *EnableScalingGroupRequest) SetOwnerId(v int64) *EnableScalingGroupRequest {
	s.OwnerId = &v
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

func (s *EnableScalingGroupRequest) SetActiveScalingConfigurationId(v string) *EnableScalingGroupRequest {
	s.ActiveScalingConfigurationId = &v
	return s
}

func (s *EnableScalingGroupRequest) SetOwnerAccount(v string) *EnableScalingGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *EnableScalingGroupRequest) SetLaunchTemplateId(v string) *EnableScalingGroupRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *EnableScalingGroupRequest) SetLaunchTemplateVersion(v string) *EnableScalingGroupRequest {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *EnableScalingGroupRequest) SetInstanceId(v []*string) *EnableScalingGroupRequest {
	s.InstanceId = v
	return s
}

func (s *EnableScalingGroupRequest) SetLoadBalancerWeight(v []*int) *EnableScalingGroupRequest {
	s.LoadBalancerWeight = v
	return s
}

func (s *EnableScalingGroupRequest) SetLaunchTemplateOverride(v []*EnableScalingGroupRequestLaunchTemplateOverride) *EnableScalingGroupRequest {
	s.LaunchTemplateOverride = v
	return s
}

type EnableScalingGroupRequestLaunchTemplateOverride struct {
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	WeightedCapacity *int32  `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s EnableScalingGroupRequestLaunchTemplateOverride) String() string {
	return tea.Prettify(s)
}

func (s EnableScalingGroupRequestLaunchTemplateOverride) GoString() string {
	return s.String()
}

func (s *EnableScalingGroupRequestLaunchTemplateOverride) SetInstanceType(v string) *EnableScalingGroupRequestLaunchTemplateOverride {
	s.InstanceType = &v
	return s
}

func (s *EnableScalingGroupRequestLaunchTemplateOverride) SetWeightedCapacity(v int32) *EnableScalingGroupRequestLaunchTemplateOverride {
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *EnableScalingGroupResponse) SetBody(v *EnableScalingGroupResponseBody) *EnableScalingGroupResponse {
	s.Body = v
	return s
}

type EnterStandbyRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId           []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s EnterStandbyRequest) String() string {
	return tea.Prettify(s)
}

func (s EnterStandbyRequest) GoString() string {
	return s.String()
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

func (s *EnterStandbyRequest) SetClientToken(v string) *EnterStandbyRequest {
	s.ClientToken = &v
	return s
}

func (s *EnterStandbyRequest) SetInstanceId(v []*string) *EnterStandbyRequest {
	s.InstanceId = v
	return s
}

type EnterStandbyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type EnterStandbyResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnterStandbyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *EnterStandbyResponse) SetBody(v *EnterStandbyResponseBody) *EnterStandbyResponse {
	s.Body = v
	return s
}

type ExecuteScalingRuleRequest struct {
	OwnerId              *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string  `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64   `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingRuleAri       *string  `json:"ScalingRuleAri,omitempty" xml:"ScalingRuleAri,omitempty"`
	ClientToken          *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	BreachThreshold      *float32 `json:"BreachThreshold,omitempty" xml:"BreachThreshold,omitempty"`
	MetricValue          *float32 `json:"MetricValue,omitempty" xml:"MetricValue,omitempty"`
	OwnerAccount         *string  `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s ExecuteScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *ExecuteScalingRuleRequest) SetOwnerId(v int64) *ExecuteScalingRuleRequest {
	s.OwnerId = &v
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

func (s *ExecuteScalingRuleRequest) SetClientToken(v string) *ExecuteScalingRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *ExecuteScalingRuleRequest) SetBreachThreshold(v float32) *ExecuteScalingRuleRequest {
	s.BreachThreshold = &v
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

type ExecuteScalingRuleResponseBody struct {
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteScalingRuleResponseBody) SetScalingActivityId(v string) *ExecuteScalingRuleResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *ExecuteScalingRuleResponseBody) SetRequestId(v string) *ExecuteScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

type ExecuteScalingRuleResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ExecuteScalingRuleResponse) SetBody(v *ExecuteScalingRuleResponseBody) *ExecuteScalingRuleResponse {
	s.Body = v
	return s
}

type ExitStandbyRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId           []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s ExitStandbyRequest) String() string {
	return tea.Prettify(s)
}

func (s ExitStandbyRequest) GoString() string {
	return s.String()
}

func (s *ExitStandbyRequest) SetOwnerId(v int64) *ExitStandbyRequest {
	s.OwnerId = &v
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

func (s *ExitStandbyRequest) SetClientToken(v string) *ExitStandbyRequest {
	s.ClientToken = &v
	return s
}

func (s *ExitStandbyRequest) SetInstanceId(v []*string) *ExitStandbyRequest {
	s.InstanceId = v
	return s
}

type ExitStandbyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type ExitStandbyResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExitStandbyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ExitStandbyResponse) SetBody(v *ExitStandbyResponseBody) *ExitStandbyResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetOwnerId(v int64) *ListTagKeysRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceOwnerAccount(v string) *ListTagKeysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int32) *ListTagKeysRequest {
	s.PageSize = &v
	return s
}

type ListTagKeysResponseBody struct {
	NextToken *string                      `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize  *int32                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Keys      *ListTagKeysResponseBodyKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetPageSize(v int32) *ListTagKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysResponseBody) SetKeys(v *ListTagKeysResponseBodyKeys) *ListTagKeysResponseBody {
	s.Keys = v
	return s
}

type ListTagKeysResponseBodyKeys struct {
	Key []*string `json:"Key,omitempty" xml:"Key,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBodyKeys) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBodyKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyKeys) SetKey(v []*string) *ListTagKeysResponseBodyKeys {
	s.Key = v
	return s
}

type ListTagKeysResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	OwnerId              *int64                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	RegionId             *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType         *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken            *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId           []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag                  []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
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
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListTagValuesRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetOwnerId(v int64) *ListTagValuesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceOwnerAccount(v string) *ListTagValuesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagValuesRequest) SetRegionId(v string) *ListTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetPageSize(v int32) *ListTagValuesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagValuesRequest) SetKey(v string) *ListTagValuesRequest {
	s.Key = &v
	return s
}

type ListTagValuesResponseBody struct {
	NextToken *string                          `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize  *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Values    *ListTagValuesResponseBodyValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
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

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetPageSize(v int32) *ListTagValuesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagValuesResponseBody) SetValues(v *ListTagValuesResponseBodyValues) *ListTagValuesResponseBody {
	s.Values = v
	return s
}

type ListTagValuesResponseBodyValues struct {
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBodyValues) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBodyValues) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBodyValues) SetValue(v []*string) *ListTagValuesResponseBodyValues {
	s.Value = v
	return s
}

type ListTagValuesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type ModifyAlarmRequest struct {
	OwnerId              *int64                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                        `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	RegionId             *string                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AlarmTaskId          *string                        `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	Name                 *string                        `json:"Name,omitempty" xml:"Name,omitempty"`
	Description          *string                        `json:"Description,omitempty" xml:"Description,omitempty"`
	MetricName           *string                        `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	MetricType           *string                        `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Period               *int32                         `json:"Period,omitempty" xml:"Period,omitempty"`
	Statistics           *string                        `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold            *float32                       `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	ComparisonOperator   *string                        `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	EvaluationCount      *int32                         `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	GroupId              *int32                         `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Effective            *string                        `json:"Effective,omitempty" xml:"Effective,omitempty"`
	AlarmAction          []*string                      `json:"AlarmAction,omitempty" xml:"AlarmAction,omitempty" type:"Repeated"`
	Dimension            []*ModifyAlarmRequestDimension `json:"Dimension,omitempty" xml:"Dimension,omitempty" type:"Repeated"`
}

func (s ModifyAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAlarmRequest) GoString() string {
	return s.String()
}

func (s *ModifyAlarmRequest) SetOwnerId(v int64) *ModifyAlarmRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAlarmRequest) SetResourceOwnerAccount(v string) *ModifyAlarmRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAlarmRequest) SetRegionId(v string) *ModifyAlarmRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAlarmRequest) SetAlarmTaskId(v string) *ModifyAlarmRequest {
	s.AlarmTaskId = &v
	return s
}

func (s *ModifyAlarmRequest) SetName(v string) *ModifyAlarmRequest {
	s.Name = &v
	return s
}

func (s *ModifyAlarmRequest) SetDescription(v string) *ModifyAlarmRequest {
	s.Description = &v
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

func (s *ModifyAlarmRequest) SetPeriod(v int32) *ModifyAlarmRequest {
	s.Period = &v
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

func (s *ModifyAlarmRequest) SetComparisonOperator(v string) *ModifyAlarmRequest {
	s.ComparisonOperator = &v
	return s
}

func (s *ModifyAlarmRequest) SetEvaluationCount(v int32) *ModifyAlarmRequest {
	s.EvaluationCount = &v
	return s
}

func (s *ModifyAlarmRequest) SetGroupId(v int32) *ModifyAlarmRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyAlarmRequest) SetEffective(v string) *ModifyAlarmRequest {
	s.Effective = &v
	return s
}

func (s *ModifyAlarmRequest) SetAlarmAction(v []*string) *ModifyAlarmRequest {
	s.AlarmAction = v
	return s
}

func (s *ModifyAlarmRequest) SetDimension(v []*ModifyAlarmRequestDimension) *ModifyAlarmRequest {
	s.Dimension = v
	return s
}

type ModifyAlarmRequestDimension struct {
	DimensionKey   *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
}

func (s ModifyAlarmRequestDimension) String() string {
	return tea.Prettify(s)
}

func (s ModifyAlarmRequestDimension) GoString() string {
	return s.String()
}

func (s *ModifyAlarmRequestDimension) SetDimensionKey(v string) *ModifyAlarmRequestDimension {
	s.DimensionKey = &v
	return s
}

func (s *ModifyAlarmRequestDimension) SetDimensionValue(v string) *ModifyAlarmRequestDimension {
	s.DimensionValue = &v
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyAlarmResponse) SetBody(v *ModifyAlarmResponseBody) *ModifyAlarmResponse {
	s.Body = v
	return s
}

type ModifyLifecycleHookRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	LifecycleHookId      *string `json:"LifecycleHookId,omitempty" xml:"LifecycleHookId,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	LifecycleHookName    *string `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	DefaultResult        *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	HeartbeatTimeout     *int32  `json:"HeartbeatTimeout,omitempty" xml:"HeartbeatTimeout,omitempty"`
	LifecycleTransition  *string `json:"LifecycleTransition,omitempty" xml:"LifecycleTransition,omitempty"`
	NotificationMetadata *string `json:"NotificationMetadata,omitempty" xml:"NotificationMetadata,omitempty"`
	NotificationArn      *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
}

func (s ModifyLifecycleHookRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLifecycleHookRequest) GoString() string {
	return s.String()
}

func (s *ModifyLifecycleHookRequest) SetOwnerId(v int64) *ModifyLifecycleHookRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetResourceOwnerAccount(v string) *ModifyLifecycleHookRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetOwnerAccount(v string) *ModifyLifecycleHookRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetLifecycleHookId(v string) *ModifyLifecycleHookRequest {
	s.LifecycleHookId = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetScalingGroupId(v string) *ModifyLifecycleHookRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetLifecycleHookName(v string) *ModifyLifecycleHookRequest {
	s.LifecycleHookName = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetDefaultResult(v string) *ModifyLifecycleHookRequest {
	s.DefaultResult = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetHeartbeatTimeout(v int32) *ModifyLifecycleHookRequest {
	s.HeartbeatTimeout = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetLifecycleTransition(v string) *ModifyLifecycleHookRequest {
	s.LifecycleTransition = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetNotificationMetadata(v string) *ModifyLifecycleHookRequest {
	s.NotificationMetadata = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetNotificationArn(v string) *ModifyLifecycleHookRequest {
	s.NotificationArn = &v
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyLifecycleHookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyLifecycleHookResponse) SetBody(v *ModifyLifecycleHookResponseBody) *ModifyLifecycleHookResponse {
	s.Body = v
	return s
}

type ModifyNotificationConfigurationRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	NotificationArn      *string   `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	NotificationType     []*string `json:"NotificationType,omitempty" xml:"NotificationType,omitempty" type:"Repeated"`
}

func (s ModifyNotificationConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNotificationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ModifyNotificationConfigurationRequest) SetOwnerId(v int64) *ModifyNotificationConfigurationRequest {
	s.OwnerId = &v
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

func (s *ModifyNotificationConfigurationRequest) SetNotificationArn(v string) *ModifyNotificationConfigurationRequest {
	s.NotificationArn = &v
	return s
}

func (s *ModifyNotificationConfigurationRequest) SetNotificationType(v []*string) *ModifyNotificationConfigurationRequest {
	s.NotificationType = v
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
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyNotificationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyNotificationConfigurationResponse) SetBody(v *ModifyNotificationConfigurationResponseBody) *ModifyNotificationConfigurationResponse {
	s.Body = v
	return s
}

type ModifyScalingConfigurationRequest struct {
	SystemDisk               *ModifyScalingConfigurationRequestSystemDisk             `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	PrivatePoolOptions       *ModifyScalingConfigurationRequestPrivatePoolOptions     `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	OwnerId                  *int64                                                   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount     *string                                                  `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	OwnerAccount             *string                                                  `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	ScalingConfigurationId   *string                                                  `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	IoOptimized              *string                                                  `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	SpotStrategy             *string                                                  `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	ScalingConfigurationName *string                                                  `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	InstanceName             *string                                                  `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	HostName                 *string                                                  `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ImageId                  *string                                                  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName                *string                                                  `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	Cpu                      *int32                                                   `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Memory                   *int32                                                   `json:"Memory,omitempty" xml:"Memory,omitempty"`
	InternetChargeType       *string                                                  `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandwidthOut  *int32                                                   `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	LoadBalancerWeight       *int32                                                   `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	UserData                 *string                                                  `json:"UserData,omitempty" xml:"UserData,omitempty"`
	KeyPairName              *string                                                  `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	RamRoleName              *string                                                  `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	PasswordInherit          *bool                                                    `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	Tags                     *string                                                  `json:"Tags,omitempty" xml:"Tags,omitempty"`
	DeploymentSetId          *string                                                  `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	SecurityGroupId          *string                                                  `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Override                 *bool                                                    `json:"Override,omitempty" xml:"Override,omitempty"`
	ResourceGroupId          *string                                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	HpcClusterId             *string                                                  `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	InstanceDescription      *string                                                  `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	Ipv6AddressCount         *int32                                                   `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	CreditSpecification      *string                                                  `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	ImageFamily              *string                                                  `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	ZoneId                   *string                                                  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DedicatedHostId          *string                                                  `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	Affinity                 *string                                                  `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	Tenancy                  *string                                                  `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	SchedulerOptions         map[string]interface{}                                   `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty"`
	SpotDuration             *int32                                                   `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	SpotInterruptionBehavior *string                                                  `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	DataDisk                 []*ModifyScalingConfigurationRequestDataDisk             `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	SpotPriceLimit           []*ModifyScalingConfigurationRequestSpotPriceLimit       `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty" type:"Repeated"`
	InstanceTypes            []*string                                                `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	InstanceTypeOverride     []*ModifyScalingConfigurationRequestInstanceTypeOverride `json:"InstanceTypeOverride,omitempty" xml:"InstanceTypeOverride,omitempty" type:"Repeated"`
	SecurityGroupIds         []*string                                                `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	InstancePatternInfo      []*ModifyScalingConfigurationRequestInstancePatternInfo  `json:"InstancePatternInfo,omitempty" xml:"InstancePatternInfo,omitempty" type:"Repeated"`
	SystemDiskCategory       []*string                                                `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty" type:"Repeated"`
}

func (s ModifyScalingConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequest) SetSystemDisk(v *ModifyScalingConfigurationRequestSystemDisk) *ModifyScalingConfigurationRequest {
	s.SystemDisk = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetPrivatePoolOptions(v *ModifyScalingConfigurationRequestPrivatePoolOptions) *ModifyScalingConfigurationRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetOwnerId(v int64) *ModifyScalingConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetResourceOwnerAccount(v string) *ModifyScalingConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetOwnerAccount(v string) *ModifyScalingConfigurationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetScalingConfigurationId(v string) *ModifyScalingConfigurationRequest {
	s.ScalingConfigurationId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetIoOptimized(v string) *ModifyScalingConfigurationRequest {
	s.IoOptimized = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSpotStrategy(v string) *ModifyScalingConfigurationRequest {
	s.SpotStrategy = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetScalingConfigurationName(v string) *ModifyScalingConfigurationRequest {
	s.ScalingConfigurationName = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetInstanceName(v string) *ModifyScalingConfigurationRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetHostName(v string) *ModifyScalingConfigurationRequest {
	s.HostName = &v
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

func (s *ModifyScalingConfigurationRequest) SetCpu(v int32) *ModifyScalingConfigurationRequest {
	s.Cpu = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetMemory(v int32) *ModifyScalingConfigurationRequest {
	s.Memory = &v
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

func (s *ModifyScalingConfigurationRequest) SetLoadBalancerWeight(v int32) *ModifyScalingConfigurationRequest {
	s.LoadBalancerWeight = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetUserData(v string) *ModifyScalingConfigurationRequest {
	s.UserData = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetKeyPairName(v string) *ModifyScalingConfigurationRequest {
	s.KeyPairName = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetRamRoleName(v string) *ModifyScalingConfigurationRequest {
	s.RamRoleName = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetPasswordInherit(v bool) *ModifyScalingConfigurationRequest {
	s.PasswordInherit = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetTags(v string) *ModifyScalingConfigurationRequest {
	s.Tags = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetDeploymentSetId(v string) *ModifyScalingConfigurationRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSecurityGroupId(v string) *ModifyScalingConfigurationRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetOverride(v bool) *ModifyScalingConfigurationRequest {
	s.Override = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetResourceGroupId(v string) *ModifyScalingConfigurationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetHpcClusterId(v string) *ModifyScalingConfigurationRequest {
	s.HpcClusterId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetInstanceDescription(v string) *ModifyScalingConfigurationRequest {
	s.InstanceDescription = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetIpv6AddressCount(v int32) *ModifyScalingConfigurationRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetCreditSpecification(v string) *ModifyScalingConfigurationRequest {
	s.CreditSpecification = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetImageFamily(v string) *ModifyScalingConfigurationRequest {
	s.ImageFamily = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetZoneId(v string) *ModifyScalingConfigurationRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetDedicatedHostId(v string) *ModifyScalingConfigurationRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetAffinity(v string) *ModifyScalingConfigurationRequest {
	s.Affinity = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetTenancy(v string) *ModifyScalingConfigurationRequest {
	s.Tenancy = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSchedulerOptions(v map[string]interface{}) *ModifyScalingConfigurationRequest {
	s.SchedulerOptions = v
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

func (s *ModifyScalingConfigurationRequest) SetDataDisk(v []*ModifyScalingConfigurationRequestDataDisk) *ModifyScalingConfigurationRequest {
	s.DataDisk = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSpotPriceLimit(v []*ModifyScalingConfigurationRequestSpotPriceLimit) *ModifyScalingConfigurationRequest {
	s.SpotPriceLimit = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetInstanceTypes(v []*string) *ModifyScalingConfigurationRequest {
	s.InstanceTypes = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetInstanceTypeOverride(v []*ModifyScalingConfigurationRequestInstanceTypeOverride) *ModifyScalingConfigurationRequest {
	s.InstanceTypeOverride = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSecurityGroupIds(v []*string) *ModifyScalingConfigurationRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetInstancePatternInfo(v []*ModifyScalingConfigurationRequestInstancePatternInfo) *ModifyScalingConfigurationRequest {
	s.InstancePatternInfo = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSystemDiskCategory(v []*string) *ModifyScalingConfigurationRequest {
	s.SystemDiskCategory = v
	return s
}

type ModifyScalingConfigurationRequestSystemDisk struct {
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Size                 *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	DiskName             *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	PerformanceLevel     *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
}

func (s ModifyScalingConfigurationRequestSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetCategory(v string) *ModifyScalingConfigurationRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetSize(v int32) *ModifyScalingConfigurationRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetDiskName(v string) *ModifyScalingConfigurationRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetDescription(v string) *ModifyScalingConfigurationRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetAutoSnapshotPolicyId(v string) *ModifyScalingConfigurationRequestSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetPerformanceLevel(v string) *ModifyScalingConfigurationRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

type ModifyScalingConfigurationRequestPrivatePoolOptions struct {
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ModifyScalingConfigurationRequestPrivatePoolOptions) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestPrivatePoolOptions) SetMatchCriteria(v string) *ModifyScalingConfigurationRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *ModifyScalingConfigurationRequestPrivatePoolOptions) SetId(v string) *ModifyScalingConfigurationRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

type ModifyScalingConfigurationRequestDataDisk struct {
	Categorys            []*string `json:"Categorys,omitempty" xml:"Categorys,omitempty" type:"Repeated"`
	PerformanceLevel     *string   `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	AutoSnapshotPolicyId *string   `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	Encrypted            *string   `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	Description          *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	SnapshotId           *string   `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	Size                 *int32    `json:"Size,omitempty" xml:"Size,omitempty"`
	Device               *string   `json:"Device,omitempty" xml:"Device,omitempty"`
	DiskName             *string   `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	Category             *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteWithInstance   *bool     `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	KMSKeyId             *string   `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
}

func (s ModifyScalingConfigurationRequestDataDisk) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationRequestDataDisk) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestDataDisk) SetCategorys(v []*string) *ModifyScalingConfigurationRequestDataDisk {
	s.Categorys = v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisk) SetPerformanceLevel(v string) *ModifyScalingConfigurationRequestDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisk) SetAutoSnapshotPolicyId(v string) *ModifyScalingConfigurationRequestDataDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisk) SetEncrypted(v string) *ModifyScalingConfigurationRequestDataDisk {
	s.Encrypted = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisk) SetDescription(v string) *ModifyScalingConfigurationRequestDataDisk {
	s.Description = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisk) SetSnapshotId(v string) *ModifyScalingConfigurationRequestDataDisk {
	s.SnapshotId = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisk) SetSize(v int32) *ModifyScalingConfigurationRequestDataDisk {
	s.Size = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisk) SetDevice(v string) *ModifyScalingConfigurationRequestDataDisk {
	s.Device = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisk) SetDiskName(v string) *ModifyScalingConfigurationRequestDataDisk {
	s.DiskName = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisk) SetCategory(v string) *ModifyScalingConfigurationRequestDataDisk {
	s.Category = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisk) SetDeleteWithInstance(v bool) *ModifyScalingConfigurationRequestDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisk) SetKMSKeyId(v string) *ModifyScalingConfigurationRequestDataDisk {
	s.KMSKeyId = &v
	return s
}

type ModifyScalingConfigurationRequestSpotPriceLimit struct {
	PriceLimit   *float32 `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
	InstanceType *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s ModifyScalingConfigurationRequestSpotPriceLimit) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationRequestSpotPriceLimit) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestSpotPriceLimit) SetPriceLimit(v float32) *ModifyScalingConfigurationRequestSpotPriceLimit {
	s.PriceLimit = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSpotPriceLimit) SetInstanceType(v string) *ModifyScalingConfigurationRequestSpotPriceLimit {
	s.InstanceType = &v
	return s
}

type ModifyScalingConfigurationRequestInstanceTypeOverride struct {
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	WeightedCapacity *int32  `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s ModifyScalingConfigurationRequestInstanceTypeOverride) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationRequestInstanceTypeOverride) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestInstanceTypeOverride) SetInstanceType(v string) *ModifyScalingConfigurationRequestInstanceTypeOverride {
	s.InstanceType = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstanceTypeOverride) SetWeightedCapacity(v int32) *ModifyScalingConfigurationRequestInstanceTypeOverride {
	s.WeightedCapacity = &v
	return s
}

type ModifyScalingConfigurationRequestInstancePatternInfo struct {
	Cores               *int32   `json:"Cores,omitempty" xml:"Cores,omitempty"`
	InstanceFamilyLevel *string  `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	MaxPrice            *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	Memory              *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ModifyScalingConfigurationRequestInstancePatternInfo) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationRequestInstancePatternInfo) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfo) SetCores(v int32) *ModifyScalingConfigurationRequestInstancePatternInfo {
	s.Cores = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfo) SetInstanceFamilyLevel(v string) *ModifyScalingConfigurationRequestInstancePatternInfo {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfo) SetMaxPrice(v float32) *ModifyScalingConfigurationRequestInstancePatternInfo {
	s.MaxPrice = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfo) SetMemory(v float32) *ModifyScalingConfigurationRequestInstancePatternInfo {
	s.Memory = &v
	return s
}

type ModifyScalingConfigurationShrinkRequest struct {
	SystemDisk               *ModifyScalingConfigurationShrinkRequestSystemDisk             `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	PrivatePoolOptions       *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions     `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	OwnerId                  *int64                                                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount     *string                                                        `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	OwnerAccount             *string                                                        `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	ScalingConfigurationId   *string                                                        `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	IoOptimized              *string                                                        `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	SpotStrategy             *string                                                        `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	ScalingConfigurationName *string                                                        `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	InstanceName             *string                                                        `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	HostName                 *string                                                        `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ImageId                  *string                                                        `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName                *string                                                        `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	Cpu                      *int32                                                         `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Memory                   *int32                                                         `json:"Memory,omitempty" xml:"Memory,omitempty"`
	InternetChargeType       *string                                                        `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandwidthOut  *int32                                                         `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	LoadBalancerWeight       *int32                                                         `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	UserData                 *string                                                        `json:"UserData,omitempty" xml:"UserData,omitempty"`
	KeyPairName              *string                                                        `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	RamRoleName              *string                                                        `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	PasswordInherit          *bool                                                          `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	Tags                     *string                                                        `json:"Tags,omitempty" xml:"Tags,omitempty"`
	DeploymentSetId          *string                                                        `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	SecurityGroupId          *string                                                        `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Override                 *bool                                                          `json:"Override,omitempty" xml:"Override,omitempty"`
	ResourceGroupId          *string                                                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	HpcClusterId             *string                                                        `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	InstanceDescription      *string                                                        `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	Ipv6AddressCount         *int32                                                         `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	CreditSpecification      *string                                                        `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	ImageFamily              *string                                                        `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	ZoneId                   *string                                                        `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DedicatedHostId          *string                                                        `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	Affinity                 *string                                                        `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	Tenancy                  *string                                                        `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	SchedulerOptionsShrink   *string                                                        `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty"`
	SpotDuration             *int32                                                         `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	SpotInterruptionBehavior *string                                                        `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	DataDisk                 []*ModifyScalingConfigurationShrinkRequestDataDisk             `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	SpotPriceLimit           []*ModifyScalingConfigurationShrinkRequestSpotPriceLimit       `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty" type:"Repeated"`
	InstanceTypes            []*string                                                      `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	InstanceTypeOverride     []*ModifyScalingConfigurationShrinkRequestInstanceTypeOverride `json:"InstanceTypeOverride,omitempty" xml:"InstanceTypeOverride,omitempty" type:"Repeated"`
	SecurityGroupIds         []*string                                                      `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	InstancePatternInfo      []*ModifyScalingConfigurationShrinkRequestInstancePatternInfo  `json:"InstancePatternInfo,omitempty" xml:"InstancePatternInfo,omitempty" type:"Repeated"`
	SystemDiskCategory       []*string                                                      `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty" type:"Repeated"`
}

func (s ModifyScalingConfigurationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSystemDisk(v *ModifyScalingConfigurationShrinkRequestSystemDisk) *ModifyScalingConfigurationShrinkRequest {
	s.SystemDisk = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetPrivatePoolOptions(v *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions) *ModifyScalingConfigurationShrinkRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetOwnerId(v int64) *ModifyScalingConfigurationShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetResourceOwnerAccount(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetOwnerAccount(v string) *ModifyScalingConfigurationShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetScalingConfigurationId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ScalingConfigurationId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetIoOptimized(v string) *ModifyScalingConfigurationShrinkRequest {
	s.IoOptimized = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSpotStrategy(v string) *ModifyScalingConfigurationShrinkRequest {
	s.SpotStrategy = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetScalingConfigurationName(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ScalingConfigurationName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetInstanceName(v string) *ModifyScalingConfigurationShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetHostName(v string) *ModifyScalingConfigurationShrinkRequest {
	s.HostName = &v
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

func (s *ModifyScalingConfigurationShrinkRequest) SetCpu(v int32) *ModifyScalingConfigurationShrinkRequest {
	s.Cpu = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetMemory(v int32) *ModifyScalingConfigurationShrinkRequest {
	s.Memory = &v
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

func (s *ModifyScalingConfigurationShrinkRequest) SetLoadBalancerWeight(v int32) *ModifyScalingConfigurationShrinkRequest {
	s.LoadBalancerWeight = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetUserData(v string) *ModifyScalingConfigurationShrinkRequest {
	s.UserData = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetKeyPairName(v string) *ModifyScalingConfigurationShrinkRequest {
	s.KeyPairName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetRamRoleName(v string) *ModifyScalingConfigurationShrinkRequest {
	s.RamRoleName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetPasswordInherit(v bool) *ModifyScalingConfigurationShrinkRequest {
	s.PasswordInherit = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetTags(v string) *ModifyScalingConfigurationShrinkRequest {
	s.Tags = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetDeploymentSetId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSecurityGroupId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetOverride(v bool) *ModifyScalingConfigurationShrinkRequest {
	s.Override = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetResourceGroupId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetHpcClusterId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.HpcClusterId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetInstanceDescription(v string) *ModifyScalingConfigurationShrinkRequest {
	s.InstanceDescription = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetIpv6AddressCount(v int32) *ModifyScalingConfigurationShrinkRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetCreditSpecification(v string) *ModifyScalingConfigurationShrinkRequest {
	s.CreditSpecification = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetImageFamily(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ImageFamily = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetZoneId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetDedicatedHostId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetAffinity(v string) *ModifyScalingConfigurationShrinkRequest {
	s.Affinity = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetTenancy(v string) *ModifyScalingConfigurationShrinkRequest {
	s.Tenancy = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSchedulerOptionsShrink(v string) *ModifyScalingConfigurationShrinkRequest {
	s.SchedulerOptionsShrink = &v
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

func (s *ModifyScalingConfigurationShrinkRequest) SetDataDisk(v []*ModifyScalingConfigurationShrinkRequestDataDisk) *ModifyScalingConfigurationShrinkRequest {
	s.DataDisk = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSpotPriceLimit(v []*ModifyScalingConfigurationShrinkRequestSpotPriceLimit) *ModifyScalingConfigurationShrinkRequest {
	s.SpotPriceLimit = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetInstanceTypes(v []*string) *ModifyScalingConfigurationShrinkRequest {
	s.InstanceTypes = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetInstanceTypeOverride(v []*ModifyScalingConfigurationShrinkRequestInstanceTypeOverride) *ModifyScalingConfigurationShrinkRequest {
	s.InstanceTypeOverride = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSecurityGroupIds(v []*string) *ModifyScalingConfigurationShrinkRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetInstancePatternInfo(v []*ModifyScalingConfigurationShrinkRequestInstancePatternInfo) *ModifyScalingConfigurationShrinkRequest {
	s.InstancePatternInfo = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSystemDiskCategory(v []*string) *ModifyScalingConfigurationShrinkRequest {
	s.SystemDiskCategory = v
	return s
}

type ModifyScalingConfigurationShrinkRequestSystemDisk struct {
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Size                 *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	DiskName             *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	PerformanceLevel     *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetCategory(v string) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetSize(v int32) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetDiskName(v string) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetDescription(v string) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetAutoSnapshotPolicyId(v string) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetPerformanceLevel(v string) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

type ModifyScalingConfigurationShrinkRequestPrivatePoolOptions struct {
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestPrivatePoolOptions) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions) SetMatchCriteria(v string) *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions) SetId(v string) *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

type ModifyScalingConfigurationShrinkRequestDataDisk struct {
	Categorys            []*string `json:"Categorys,omitempty" xml:"Categorys,omitempty" type:"Repeated"`
	PerformanceLevel     *string   `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	AutoSnapshotPolicyId *string   `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	Encrypted            *string   `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	Description          *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	SnapshotId           *string   `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	Size                 *int32    `json:"Size,omitempty" xml:"Size,omitempty"`
	Device               *string   `json:"Device,omitempty" xml:"Device,omitempty"`
	DiskName             *string   `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	Category             *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteWithInstance   *bool     `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	KMSKeyId             *string   `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestDataDisk) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestDataDisk) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisk) SetCategorys(v []*string) *ModifyScalingConfigurationShrinkRequestDataDisk {
	s.Categorys = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisk) SetPerformanceLevel(v string) *ModifyScalingConfigurationShrinkRequestDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisk) SetAutoSnapshotPolicyId(v string) *ModifyScalingConfigurationShrinkRequestDataDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisk) SetEncrypted(v string) *ModifyScalingConfigurationShrinkRequestDataDisk {
	s.Encrypted = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisk) SetDescription(v string) *ModifyScalingConfigurationShrinkRequestDataDisk {
	s.Description = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisk) SetSnapshotId(v string) *ModifyScalingConfigurationShrinkRequestDataDisk {
	s.SnapshotId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisk) SetSize(v int32) *ModifyScalingConfigurationShrinkRequestDataDisk {
	s.Size = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisk) SetDevice(v string) *ModifyScalingConfigurationShrinkRequestDataDisk {
	s.Device = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisk) SetDiskName(v string) *ModifyScalingConfigurationShrinkRequestDataDisk {
	s.DiskName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisk) SetCategory(v string) *ModifyScalingConfigurationShrinkRequestDataDisk {
	s.Category = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisk) SetDeleteWithInstance(v bool) *ModifyScalingConfigurationShrinkRequestDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisk) SetKMSKeyId(v string) *ModifyScalingConfigurationShrinkRequestDataDisk {
	s.KMSKeyId = &v
	return s
}

type ModifyScalingConfigurationShrinkRequestSpotPriceLimit struct {
	PriceLimit   *float32 `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
	InstanceType *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestSpotPriceLimit) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestSpotPriceLimit) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestSpotPriceLimit) SetPriceLimit(v float32) *ModifyScalingConfigurationShrinkRequestSpotPriceLimit {
	s.PriceLimit = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSpotPriceLimit) SetInstanceType(v string) *ModifyScalingConfigurationShrinkRequestSpotPriceLimit {
	s.InstanceType = &v
	return s
}

type ModifyScalingConfigurationShrinkRequestInstanceTypeOverride struct {
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	WeightedCapacity *int32  `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestInstanceTypeOverride) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestInstanceTypeOverride) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestInstanceTypeOverride) SetInstanceType(v string) *ModifyScalingConfigurationShrinkRequestInstanceTypeOverride {
	s.InstanceType = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstanceTypeOverride) SetWeightedCapacity(v int32) *ModifyScalingConfigurationShrinkRequestInstanceTypeOverride {
	s.WeightedCapacity = &v
	return s
}

type ModifyScalingConfigurationShrinkRequestInstancePatternInfo struct {
	Cores               *int32   `json:"Cores,omitempty" xml:"Cores,omitempty"`
	InstanceFamilyLevel *string  `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	MaxPrice            *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	Memory              *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestInstancePatternInfo) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestInstancePatternInfo) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfo) SetCores(v int32) *ModifyScalingConfigurationShrinkRequestInstancePatternInfo {
	s.Cores = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfo) SetInstanceFamilyLevel(v string) *ModifyScalingConfigurationShrinkRequestInstancePatternInfo {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfo) SetMaxPrice(v float32) *ModifyScalingConfigurationShrinkRequestInstancePatternInfo {
	s.MaxPrice = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfo) SetMemory(v float32) *ModifyScalingConfigurationShrinkRequestInstancePatternInfo {
	s.Memory = &v
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyScalingConfigurationResponse) SetBody(v *ModifyScalingConfigurationResponseBody) *ModifyScalingConfigurationResponse {
	s.Body = v
	return s
}

type ModifyScalingGroupRequest struct {
	RemovalPolicy                       []*string                                          `json:"RemovalPolicy,omitempty" xml:"RemovalPolicy,omitempty" type:"Repeated"`
	OwnerId                             *int64                                             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount                *string                                            `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                     *int64                                             `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupId                      *string                                            `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScalingGroupName                    *string                                            `json:"ScalingGroupName,omitempty" xml:"ScalingGroupName,omitempty"`
	MinSize                             *int32                                             `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	MaxSize                             *int32                                             `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	DefaultCooldown                     *int32                                             `json:"DefaultCooldown,omitempty" xml:"DefaultCooldown,omitempty"`
	ActiveScalingConfigurationId        *string                                            `json:"ActiveScalingConfigurationId,omitempty" xml:"ActiveScalingConfigurationId,omitempty"`
	OwnerAccount                        *string                                            `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	HealthCheckType                     *string                                            `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	LaunchTemplateId                    *string                                            `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	LaunchTemplateVersion               *string                                            `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	OnDemandBaseCapacity                *int32                                             `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	OnDemandPercentageAboveBaseCapacity *int32                                             `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	SpotInstanceRemedy                  *bool                                              `json:"SpotInstanceRemedy,omitempty" xml:"SpotInstanceRemedy,omitempty"`
	CompensateWithOnDemand              *bool                                              `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	SpotInstancePools                   *int32                                             `json:"SpotInstancePools,omitempty" xml:"SpotInstancePools,omitempty"`
	DesiredCapacity                     *int32                                             `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	GroupDeletionProtection             *bool                                              `json:"GroupDeletionProtection,omitempty" xml:"GroupDeletionProtection,omitempty"`
	MultiAZPolicy                       *string                                            `json:"MultiAZPolicy,omitempty" xml:"MultiAZPolicy,omitempty"`
	VSwitchIds                          []*string                                          `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	LaunchTemplateOverride              []*ModifyScalingGroupRequestLaunchTemplateOverride `json:"LaunchTemplateOverride,omitempty" xml:"LaunchTemplateOverride,omitempty" type:"Repeated"`
}

func (s ModifyScalingGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyScalingGroupRequest) SetRemovalPolicy(v []*string) *ModifyScalingGroupRequest {
	s.RemovalPolicy = v
	return s
}

func (s *ModifyScalingGroupRequest) SetOwnerId(v int64) *ModifyScalingGroupRequest {
	s.OwnerId = &v
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

func (s *ModifyScalingGroupRequest) SetMinSize(v int32) *ModifyScalingGroupRequest {
	s.MinSize = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetMaxSize(v int32) *ModifyScalingGroupRequest {
	s.MaxSize = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetDefaultCooldown(v int32) *ModifyScalingGroupRequest {
	s.DefaultCooldown = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetActiveScalingConfigurationId(v string) *ModifyScalingGroupRequest {
	s.ActiveScalingConfigurationId = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetOwnerAccount(v string) *ModifyScalingGroupRequest {
	s.OwnerAccount = &v
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

func (s *ModifyScalingGroupRequest) SetLaunchTemplateVersion(v string) *ModifyScalingGroupRequest {
	s.LaunchTemplateVersion = &v
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

func (s *ModifyScalingGroupRequest) SetSpotInstanceRemedy(v bool) *ModifyScalingGroupRequest {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetCompensateWithOnDemand(v bool) *ModifyScalingGroupRequest {
	s.CompensateWithOnDemand = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetSpotInstancePools(v int32) *ModifyScalingGroupRequest {
	s.SpotInstancePools = &v
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

func (s *ModifyScalingGroupRequest) SetMultiAZPolicy(v string) *ModifyScalingGroupRequest {
	s.MultiAZPolicy = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetVSwitchIds(v []*string) *ModifyScalingGroupRequest {
	s.VSwitchIds = v
	return s
}

func (s *ModifyScalingGroupRequest) SetLaunchTemplateOverride(v []*ModifyScalingGroupRequestLaunchTemplateOverride) *ModifyScalingGroupRequest {
	s.LaunchTemplateOverride = v
	return s
}

type ModifyScalingGroupRequestLaunchTemplateOverride struct {
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	WeightedCapacity *int32  `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s ModifyScalingGroupRequestLaunchTemplateOverride) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingGroupRequestLaunchTemplateOverride) GoString() string {
	return s.String()
}

func (s *ModifyScalingGroupRequestLaunchTemplateOverride) SetInstanceType(v string) *ModifyScalingGroupRequestLaunchTemplateOverride {
	s.InstanceType = &v
	return s
}

func (s *ModifyScalingGroupRequestLaunchTemplateOverride) SetWeightedCapacity(v int32) *ModifyScalingGroupRequestLaunchTemplateOverride {
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyScalingGroupResponse) SetBody(v *ModifyScalingGroupResponseBody) *ModifyScalingGroupResponse {
	s.Body = v
	return s
}

type ModifyScalingRuleRequest struct {
	OwnerId                  *int64                                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount     *string                                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64                                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingRuleId            *string                                   `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty"`
	ScalingRuleName          *string                                   `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	Cooldown                 *int32                                    `json:"Cooldown,omitempty" xml:"Cooldown,omitempty"`
	MinAdjustmentMagnitude   *int32                                    `json:"MinAdjustmentMagnitude,omitempty" xml:"MinAdjustmentMagnitude,omitempty"`
	AdjustmentType           *string                                   `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	AdjustmentValue          *int32                                    `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	EstimatedInstanceWarmup  *int32                                    `json:"EstimatedInstanceWarmup,omitempty" xml:"EstimatedInstanceWarmup,omitempty"`
	MetricName               *string                                   `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	TargetValue              *float32                                  `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
	DisableScaleIn           *bool                                     `json:"DisableScaleIn,omitempty" xml:"DisableScaleIn,omitempty"`
	ScaleInEvaluationCount   *int32                                    `json:"ScaleInEvaluationCount,omitempty" xml:"ScaleInEvaluationCount,omitempty"`
	ScaleOutEvaluationCount  *int32                                    `json:"ScaleOutEvaluationCount,omitempty" xml:"ScaleOutEvaluationCount,omitempty"`
	OwnerAccount             *string                                   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	PredictiveScalingMode    *string                                   `json:"PredictiveScalingMode,omitempty" xml:"PredictiveScalingMode,omitempty"`
	PredictiveValueBehavior  *string                                   `json:"PredictiveValueBehavior,omitempty" xml:"PredictiveValueBehavior,omitempty"`
	PredictiveValueBuffer    *int32                                    `json:"PredictiveValueBuffer,omitempty" xml:"PredictiveValueBuffer,omitempty"`
	PredictiveTaskBufferTime *int32                                    `json:"PredictiveTaskBufferTime,omitempty" xml:"PredictiveTaskBufferTime,omitempty"`
	InitialMaxSize           *int32                                    `json:"InitialMaxSize,omitempty" xml:"InitialMaxSize,omitempty"`
	StepAdjustment           []*ModifyScalingRuleRequestStepAdjustment `json:"StepAdjustment,omitempty" xml:"StepAdjustment,omitempty" type:"Repeated"`
}

func (s ModifyScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleRequest) SetOwnerId(v int64) *ModifyScalingRuleRequest {
	s.OwnerId = &v
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

func (s *ModifyScalingRuleRequest) SetScalingRuleId(v string) *ModifyScalingRuleRequest {
	s.ScalingRuleId = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetScalingRuleName(v string) *ModifyScalingRuleRequest {
	s.ScalingRuleName = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetCooldown(v int32) *ModifyScalingRuleRequest {
	s.Cooldown = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetMinAdjustmentMagnitude(v int32) *ModifyScalingRuleRequest {
	s.MinAdjustmentMagnitude = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetAdjustmentType(v string) *ModifyScalingRuleRequest {
	s.AdjustmentType = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetAdjustmentValue(v int32) *ModifyScalingRuleRequest {
	s.AdjustmentValue = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetEstimatedInstanceWarmup(v int32) *ModifyScalingRuleRequest {
	s.EstimatedInstanceWarmup = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetMetricName(v string) *ModifyScalingRuleRequest {
	s.MetricName = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetTargetValue(v float32) *ModifyScalingRuleRequest {
	s.TargetValue = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetDisableScaleIn(v bool) *ModifyScalingRuleRequest {
	s.DisableScaleIn = &v
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

func (s *ModifyScalingRuleRequest) SetOwnerAccount(v string) *ModifyScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetPredictiveScalingMode(v string) *ModifyScalingRuleRequest {
	s.PredictiveScalingMode = &v
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

func (s *ModifyScalingRuleRequest) SetPredictiveTaskBufferTime(v int32) *ModifyScalingRuleRequest {
	s.PredictiveTaskBufferTime = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetInitialMaxSize(v int32) *ModifyScalingRuleRequest {
	s.InitialMaxSize = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetStepAdjustment(v []*ModifyScalingRuleRequestStepAdjustment) *ModifyScalingRuleRequest {
	s.StepAdjustment = v
	return s
}

type ModifyScalingRuleRequestStepAdjustment struct {
	MetricIntervalUpperBound *float32 `json:"MetricIntervalUpperBound,omitempty" xml:"MetricIntervalUpperBound,omitempty"`
	ScalingAdjustment        *int32   `json:"ScalingAdjustment,omitempty" xml:"ScalingAdjustment,omitempty"`
	MetricIntervalLowerBound *float32 `json:"MetricIntervalLowerBound,omitempty" xml:"MetricIntervalLowerBound,omitempty"`
}

func (s ModifyScalingRuleRequestStepAdjustment) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingRuleRequestStepAdjustment) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleRequestStepAdjustment) SetMetricIntervalUpperBound(v float32) *ModifyScalingRuleRequestStepAdjustment {
	s.MetricIntervalUpperBound = &v
	return s
}

func (s *ModifyScalingRuleRequestStepAdjustment) SetScalingAdjustment(v int32) *ModifyScalingRuleRequestStepAdjustment {
	s.ScalingAdjustment = &v
	return s
}

func (s *ModifyScalingRuleRequestStepAdjustment) SetMetricIntervalLowerBound(v float32) *ModifyScalingRuleRequestStepAdjustment {
	s.MetricIntervalLowerBound = &v
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyScalingRuleResponse) SetBody(v *ModifyScalingRuleResponseBody) *ModifyScalingRuleResponse {
	s.Body = v
	return s
}

type ModifyScheduledTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScheduledTaskId      *string `json:"ScheduledTaskId,omitempty" xml:"ScheduledTaskId,omitempty"`
	ScheduledTaskName    *string `json:"ScheduledTaskName,omitempty" xml:"ScheduledTaskName,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ScheduledAction      *string `json:"ScheduledAction,omitempty" xml:"ScheduledAction,omitempty"`
	RecurrenceEndTime    *string `json:"RecurrenceEndTime,omitempty" xml:"RecurrenceEndTime,omitempty"`
	LaunchTime           *string `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	RecurrenceType       *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValue      *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	TaskEnabled          *bool   `json:"TaskEnabled,omitempty" xml:"TaskEnabled,omitempty"`
	LaunchExpirationTime *int32  `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	MinValue             *int32  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	MaxValue             *int32  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	DesiredCapacity      *int32  `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s ModifyScheduledTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskRequest) SetOwnerId(v int64) *ModifyScheduledTaskRequest {
	s.OwnerId = &v
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

func (s *ModifyScheduledTaskRequest) SetScheduledTaskId(v string) *ModifyScheduledTaskRequest {
	s.ScheduledTaskId = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetScheduledTaskName(v string) *ModifyScheduledTaskRequest {
	s.ScheduledTaskName = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetDescription(v string) *ModifyScheduledTaskRequest {
	s.Description = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetScheduledAction(v string) *ModifyScheduledTaskRequest {
	s.ScheduledAction = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetRecurrenceEndTime(v string) *ModifyScheduledTaskRequest {
	s.RecurrenceEndTime = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetLaunchTime(v string) *ModifyScheduledTaskRequest {
	s.LaunchTime = &v
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

func (s *ModifyScheduledTaskRequest) SetTaskEnabled(v bool) *ModifyScheduledTaskRequest {
	s.TaskEnabled = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetLaunchExpirationTime(v int32) *ModifyScheduledTaskRequest {
	s.LaunchExpirationTime = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetOwnerAccount(v string) *ModifyScheduledTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetMinValue(v int32) *ModifyScheduledTaskRequest {
	s.MinValue = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetMaxValue(v int32) *ModifyScheduledTaskRequest {
	s.MaxValue = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetDesiredCapacity(v int32) *ModifyScheduledTaskRequest {
	s.DesiredCapacity = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetScalingGroupId(v string) *ModifyScheduledTaskRequest {
	s.ScalingGroupId = &v
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyScheduledTaskResponse) SetBody(v *ModifyScheduledTaskResponseBody) *ModifyScheduledTaskResponse {
	s.Body = v
	return s
}

type RebalanceInstancesRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s RebalanceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s RebalanceInstancesRequest) GoString() string {
	return s.String()
}

func (s *RebalanceInstancesRequest) SetOwnerId(v int64) *RebalanceInstancesRequest {
	s.OwnerId = &v
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

func (s *RebalanceInstancesRequest) SetOwnerAccount(v string) *RebalanceInstancesRequest {
	s.OwnerAccount = &v
	return s
}

type RebalanceInstancesResponseBody struct {
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebalanceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebalanceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RebalanceInstancesResponseBody) SetScalingActivityId(v string) *RebalanceInstancesResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *RebalanceInstancesResponseBody) SetRequestId(v string) *RebalanceInstancesResponseBody {
	s.RequestId = &v
	return s
}

type RebalanceInstancesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RebalanceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RebalanceInstancesResponse) SetBody(v *RebalanceInstancesResponseBody) *RebalanceInstancesResponse {
	s.Body = v
	return s
}

type RecordLifecycleActionHeartbeatRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	LifecycleHookId      *string `json:"lifecycleHookId,omitempty" xml:"lifecycleHookId,omitempty"`
	LifecycleActionToken *string `json:"lifecycleActionToken,omitempty" xml:"lifecycleActionToken,omitempty"`
	HeartbeatTimeout     *int32  `json:"heartbeatTimeout,omitempty" xml:"heartbeatTimeout,omitempty"`
}

func (s RecordLifecycleActionHeartbeatRequest) String() string {
	return tea.Prettify(s)
}

func (s RecordLifecycleActionHeartbeatRequest) GoString() string {
	return s.String()
}

func (s *RecordLifecycleActionHeartbeatRequest) SetOwnerId(v int64) *RecordLifecycleActionHeartbeatRequest {
	s.OwnerId = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatRequest) SetResourceOwnerAccount(v string) *RecordLifecycleActionHeartbeatRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatRequest) SetOwnerAccount(v string) *RecordLifecycleActionHeartbeatRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatRequest) SetLifecycleHookId(v string) *RecordLifecycleActionHeartbeatRequest {
	s.LifecycleHookId = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatRequest) SetLifecycleActionToken(v string) *RecordLifecycleActionHeartbeatRequest {
	s.LifecycleActionToken = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatRequest) SetHeartbeatTimeout(v int32) *RecordLifecycleActionHeartbeatRequest {
	s.HeartbeatTimeout = &v
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecordLifecycleActionHeartbeatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecordLifecycleActionHeartbeatResponse) SetBody(v *RecordLifecycleActionHeartbeatResponseBody) *RecordLifecycleActionHeartbeatResponse {
	s.Body = v
	return s
}

type RemoveInstancesRequest struct {
	OwnerId                 *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount    *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupId          *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	OwnerAccount            *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RemovePolicy            *string   `json:"RemovePolicy,omitempty" xml:"RemovePolicy,omitempty"`
	DecreaseDesiredCapacity *bool     `json:"DecreaseDesiredCapacity,omitempty" xml:"DecreaseDesiredCapacity,omitempty"`
	InstanceId              []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s RemoveInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveInstancesRequest) GoString() string {
	return s.String()
}

func (s *RemoveInstancesRequest) SetOwnerId(v int64) *RemoveInstancesRequest {
	s.OwnerId = &v
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

func (s *RemoveInstancesRequest) SetOwnerAccount(v string) *RemoveInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveInstancesRequest) SetRemovePolicy(v string) *RemoveInstancesRequest {
	s.RemovePolicy = &v
	return s
}

func (s *RemoveInstancesRequest) SetDecreaseDesiredCapacity(v bool) *RemoveInstancesRequest {
	s.DecreaseDesiredCapacity = &v
	return s
}

func (s *RemoveInstancesRequest) SetInstanceId(v []*string) *RemoveInstancesRequest {
	s.InstanceId = v
	return s
}

type RemoveInstancesResponseBody struct {
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveInstancesResponseBody) SetScalingActivityId(v string) *RemoveInstancesResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *RemoveInstancesResponseBody) SetRequestId(v string) *RemoveInstancesResponseBody {
	s.RequestId = &v
	return s
}

type RemoveInstancesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveInstancesResponse) SetBody(v *RemoveInstancesResponseBody) *RemoveInstancesResponse {
	s.Body = v
	return s
}

type ResumeProcessesRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Process              []*string `json:"Process,omitempty" xml:"Process,omitempty" type:"Repeated"`
}

func (s ResumeProcessesRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeProcessesRequest) GoString() string {
	return s.String()
}

func (s *ResumeProcessesRequest) SetOwnerId(v int64) *ResumeProcessesRequest {
	s.OwnerId = &v
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

func (s *ResumeProcessesRequest) SetClientToken(v string) *ResumeProcessesRequest {
	s.ClientToken = &v
	return s
}

func (s *ResumeProcessesRequest) SetProcess(v []*string) *ResumeProcessesRequest {
	s.Process = v
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResumeProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ResumeProcessesResponse) SetBody(v *ResumeProcessesResponseBody) *ResumeProcessesResponse {
	s.Body = v
	return s
}

type ScaleWithAdjustmentRequest struct {
	ScalingGroupId         *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	AdjustmentType         *string `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	AdjustmentValue        *int32  `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	MinAdjustmentMagnitude *int32  `json:"MinAdjustmentMagnitude,omitempty" xml:"MinAdjustmentMagnitude,omitempty"`
	ClientToken            *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s ScaleWithAdjustmentRequest) String() string {
	return tea.Prettify(s)
}

func (s ScaleWithAdjustmentRequest) GoString() string {
	return s.String()
}

func (s *ScaleWithAdjustmentRequest) SetScalingGroupId(v string) *ScaleWithAdjustmentRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetAdjustmentType(v string) *ScaleWithAdjustmentRequest {
	s.AdjustmentType = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetAdjustmentValue(v int32) *ScaleWithAdjustmentRequest {
	s.AdjustmentValue = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetMinAdjustmentMagnitude(v int32) *ScaleWithAdjustmentRequest {
	s.MinAdjustmentMagnitude = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetClientToken(v string) *ScaleWithAdjustmentRequest {
	s.ClientToken = &v
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

type ScaleWithAdjustmentResponseBody struct {
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScaleWithAdjustmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleWithAdjustmentResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleWithAdjustmentResponseBody) SetScalingActivityId(v string) *ScaleWithAdjustmentResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *ScaleWithAdjustmentResponseBody) SetRequestId(v string) *ScaleWithAdjustmentResponseBody {
	s.RequestId = &v
	return s
}

type ScaleWithAdjustmentResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ScaleWithAdjustmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ScaleWithAdjustmentResponse) SetBody(v *ScaleWithAdjustmentResponseBody) *ScaleWithAdjustmentResponse {
	s.Body = v
	return s
}

type SetGroupDeletionProtectionRequest struct {
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount    *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	GroupDeletionProtection *bool   `json:"GroupDeletionProtection,omitempty" xml:"GroupDeletionProtection,omitempty"`
	ScalingGroupId          *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s SetGroupDeletionProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGroupDeletionProtectionRequest) GoString() string {
	return s.String()
}

func (s *SetGroupDeletionProtectionRequest) SetOwnerId(v int64) *SetGroupDeletionProtectionRequest {
	s.OwnerId = &v
	return s
}

func (s *SetGroupDeletionProtectionRequest) SetResourceOwnerAccount(v string) *SetGroupDeletionProtectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetGroupDeletionProtectionRequest) SetGroupDeletionProtection(v bool) *SetGroupDeletionProtectionRequest {
	s.GroupDeletionProtection = &v
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetGroupDeletionProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetGroupDeletionProtectionResponse) SetBody(v *SetGroupDeletionProtectionResponseBody) *SetGroupDeletionProtectionResponse {
	s.Body = v
	return s
}

type SetInstanceHealthRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	HealthStatus         *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
}

func (s SetInstanceHealthRequest) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceHealthRequest) GoString() string {
	return s.String()
}

func (s *SetInstanceHealthRequest) SetOwnerId(v int64) *SetInstanceHealthRequest {
	s.OwnerId = &v
	return s
}

func (s *SetInstanceHealthRequest) SetResourceOwnerAccount(v string) *SetInstanceHealthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetInstanceHealthRequest) SetInstanceId(v string) *SetInstanceHealthRequest {
	s.InstanceId = &v
	return s
}

func (s *SetInstanceHealthRequest) SetHealthStatus(v string) *SetInstanceHealthRequest {
	s.HealthStatus = &v
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetInstanceHealthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetInstanceHealthResponse) SetBody(v *SetInstanceHealthResponseBody) *SetInstanceHealthResponse {
	s.Body = v
	return s
}

type SetInstancesProtectionRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ProtectedFromScaleIn *bool     `json:"ProtectedFromScaleIn,omitempty" xml:"ProtectedFromScaleIn,omitempty"`
	InstanceId           []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s SetInstancesProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetInstancesProtectionRequest) GoString() string {
	return s.String()
}

func (s *SetInstancesProtectionRequest) SetOwnerId(v int64) *SetInstancesProtectionRequest {
	s.OwnerId = &v
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

func (s *SetInstancesProtectionRequest) SetProtectedFromScaleIn(v bool) *SetInstancesProtectionRequest {
	s.ProtectedFromScaleIn = &v
	return s
}

func (s *SetInstancesProtectionRequest) SetInstanceId(v []*string) *SetInstancesProtectionRequest {
	s.InstanceId = v
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetInstancesProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetInstancesProtectionResponse) SetBody(v *SetInstancesProtectionResponseBody) *SetInstancesProtectionResponse {
	s.Body = v
	return s
}

type SuspendProcessesRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupId       *string   `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Process              []*string `json:"Process,omitempty" xml:"Process,omitempty" type:"Repeated"`
}

func (s SuspendProcessesRequest) String() string {
	return tea.Prettify(s)
}

func (s SuspendProcessesRequest) GoString() string {
	return s.String()
}

func (s *SuspendProcessesRequest) SetOwnerId(v int64) *SuspendProcessesRequest {
	s.OwnerId = &v
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

func (s *SuspendProcessesRequest) SetClientToken(v string) *SuspendProcessesRequest {
	s.ClientToken = &v
	return s
}

func (s *SuspendProcessesRequest) SetProcess(v []*string) *SuspendProcessesRequest {
	s.Process = v
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SuspendProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SuspendProcessesResponse) SetBody(v *SuspendProcessesResponseBody) *SuspendProcessesResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerId              *int64                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	RegionId             *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType         *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId           []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag                  []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType         *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All                  *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey               []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerAccount(v string) *UntagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type VerifyAuthenticationRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Uid                  *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
	OnlyCheck            *bool   `json:"OnlyCheck,omitempty" xml:"OnlyCheck,omitempty"`
}

func (s VerifyAuthenticationRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyAuthenticationRequest) GoString() string {
	return s.String()
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

func (s *VerifyAuthenticationRequest) SetOnlyCheck(v bool) *VerifyAuthenticationRequest {
	s.OnlyCheck = &v
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *VerifyAuthenticationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type VerifyUserResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
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
		"us-west-1":                   tea.String("ess.aliyuncs.com"),
		"us-east-1":                   tea.String("ess.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("ess.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("ess.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("ess.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("ess.ap-northeast-1.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("ess.aliyuncs.com"),
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
		"cn-zhangbei-na61-b01":        tea.String("ess.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("ess.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("ess.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("ess.ap-northeast-1.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("ess.ap-northeast-1.aliyuncs.com"),
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachAlbServerGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachAlbServerGroups"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachDBInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachDBInstances"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachInstances"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachLoadBalancersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachLoadBalancers"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachVServerGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachVServerGroups"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CompleteLifecycleActionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CompleteLifecycleAction"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAlarmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAlarm"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateLifecycleHookWithOptions(request *CreateLifecycleHookRequest, runtime *util.RuntimeOptions) (_result *CreateLifecycleHookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateLifecycleHookResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateLifecycleHook"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateNotificationConfigurationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateNotificationConfiguration"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateScalingConfigurationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateScalingConfiguration"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateScalingGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateScalingGroup"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateScalingRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateScalingRule"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateScheduledTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateScheduledTask"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeactivateScalingConfigurationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeactivateScalingConfiguration"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAlarmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAlarm"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteLifecycleHookResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteLifecycleHook"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteNotificationConfigurationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteNotificationConfiguration"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteScalingConfigurationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteScalingConfiguration"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteScalingGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteScalingGroup"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteScalingRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteScalingRule"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteScheduledTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteScheduledTask"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAlarmsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAlarms"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeLifecycleActionsWithOptions(request *DescribeLifecycleActionsRequest, runtime *util.RuntimeOptions) (_result *DescribeLifecycleActionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeLifecycleActionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLifecycleActions"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeLifecycleHooksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLifecycleHooks"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeLimitationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLimitation"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeNotificationConfigurationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeNotificationConfigurations"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeNotificationTypesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeNotificationTypes"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeScalingActivitiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeScalingActivities"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeScalingActivityDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeScalingActivityDetail"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeScalingConfigurationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeScalingConfigurations"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeScalingInstancesWithOptions(request *DescribeScalingInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeScalingInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeScalingInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeScalingInstances"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeScalingRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeScalingRules"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeScheduledTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeScheduledTasks"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetachAlbServerGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetachAlbServerGroups"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetachDBInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetachDBInstances"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetachInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetachInstances"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetachLoadBalancersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetachLoadBalancers"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetachVServerGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetachVServerGroups"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableAlarmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableAlarm"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableScalingGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableScalingGroup"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableAlarmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableAlarm"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableScalingGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableScalingGroup"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnterStandbyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnterStandby"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExecuteScalingRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExecuteScalingRule"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExitStandbyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExitStandby"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagKeys"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagValues"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAlarmResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAlarm"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyLifecycleHookWithOptions(request *ModifyLifecycleHookRequest, runtime *util.RuntimeOptions) (_result *ModifyLifecycleHookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyLifecycleHookResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyLifecycleHook"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyNotificationConfigurationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyNotificationConfiguration"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyScalingConfigurationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyScalingConfiguration"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyScalingGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyScalingGroup"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyScalingRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyScalingRule"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyScheduledTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyScheduledTask"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RebalanceInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RebalanceInstances"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecordLifecycleActionHeartbeatResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecordLifecycleActionHeartbeat"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveInstances"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResumeProcessesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResumeProcesses"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ScaleWithAdjustmentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ScaleWithAdjustment"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetGroupDeletionProtectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetGroupDeletionProtection"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetInstanceHealthResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetInstanceHealth"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetInstancesProtectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetInstancesProtection"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SuspendProcessesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SuspendProcesses"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &VerifyAuthenticationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("VerifyAuthentication"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &VerifyUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("VerifyUser"), tea.String("2014-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("none"), req, runtime)
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
