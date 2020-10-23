// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetResourceComplianceTimelineRequest struct {
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit        *int    `json:"Limit,omitempty" xml:"Limit,omitempty"`
	MultiAccount *bool   `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	MemberId     *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetResourceComplianceTimelineRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceTimelineRequest) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceTimelineRequest) SetResourceType(v string) *GetResourceComplianceTimelineRequest {
	s.ResourceType = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetResourceId(v string) *GetResourceComplianceTimelineRequest {
	s.ResourceId = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetStartTime(v int64) *GetResourceComplianceTimelineRequest {
	s.StartTime = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetEndTime(v int64) *GetResourceComplianceTimelineRequest {
	s.EndTime = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetLimit(v int) *GetResourceComplianceTimelineRequest {
	s.Limit = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetMultiAccount(v bool) *GetResourceComplianceTimelineRequest {
	s.MultiAccount = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetMemberId(v string) *GetResourceComplianceTimelineRequest {
	s.MemberId = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetRegion(v string) *GetResourceComplianceTimelineRequest {
	s.Region = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetNextToken(v string) *GetResourceComplianceTimelineRequest {
	s.NextToken = &v
	return s
}

type GetResourceComplianceTimelineResponse struct {
	RequestId                  *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResourceComplianceTimeline *GetResourceComplianceTimelineResponseResourceComplianceTimeline `json:"ResourceComplianceTimeline,omitempty" xml:"ResourceComplianceTimeline,omitempty" require:"true" type:"Struct"`
}

func (s GetResourceComplianceTimelineResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceTimelineResponse) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceTimelineResponse) SetRequestId(v string) *GetResourceComplianceTimelineResponse {
	s.RequestId = &v
	return s
}

func (s *GetResourceComplianceTimelineResponse) SetResourceComplianceTimeline(v *GetResourceComplianceTimelineResponseResourceComplianceTimeline) *GetResourceComplianceTimelineResponse {
	s.ResourceComplianceTimeline = v
	return s
}

type GetResourceComplianceTimelineResponseResourceComplianceTimeline struct {
	Limit          *int                                                                             `json:"Limit,omitempty" xml:"Limit,omitempty" require:"true"`
	TotalCount     *int64                                                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	NextToken      *string                                                                          `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	ComplianceList []*GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList `json:"ComplianceList,omitempty" xml:"ComplianceList,omitempty" require:"true" type:"Repeated"`
}

func (s GetResourceComplianceTimelineResponseResourceComplianceTimeline) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceTimelineResponseResourceComplianceTimeline) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceTimelineResponseResourceComplianceTimeline) SetLimit(v int) *GetResourceComplianceTimelineResponseResourceComplianceTimeline {
	s.Limit = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseResourceComplianceTimeline) SetTotalCount(v int64) *GetResourceComplianceTimelineResponseResourceComplianceTimeline {
	s.TotalCount = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseResourceComplianceTimeline) SetNextToken(v string) *GetResourceComplianceTimelineResponseResourceComplianceTimeline {
	s.NextToken = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseResourceComplianceTimeline) SetComplianceList(v []*GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList) *GetResourceComplianceTimelineResponseResourceComplianceTimeline {
	s.ComplianceList = v
	return s
}

type GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList struct {
	AccountId          *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	AvailabilityZone   *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty" require:"true"`
	CaptureTime        *int64  `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty" require:"true"`
	Configuration      *string `json:"Configuration,omitempty" xml:"Configuration,omitempty" require:"true"`
	ConfigurationDiff  *string `json:"ConfigurationDiff,omitempty" xml:"ConfigurationDiff,omitempty" require:"true"`
	Region             *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	ResourceCreateTime *int64  `json:"ResourceCreateTime,omitempty" xml:"ResourceCreateTime,omitempty" require:"true"`
	ResourceId         *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	ResourceName       *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty" require:"true"`
	ResourceStatus     *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty" require:"true"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	Tags               *string `json:"Tags,omitempty" xml:"Tags,omitempty" require:"true"`
}

func (s GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList) SetAccountId(v string) *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList {
	s.AccountId = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList) SetAvailabilityZone(v string) *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList {
	s.AvailabilityZone = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList) SetCaptureTime(v int64) *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList {
	s.CaptureTime = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList) SetConfiguration(v string) *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList {
	s.Configuration = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList) SetConfigurationDiff(v string) *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList {
	s.ConfigurationDiff = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList) SetRegion(v string) *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList {
	s.Region = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList) SetResourceCreateTime(v int64) *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList {
	s.ResourceCreateTime = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList) SetResourceId(v string) *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList {
	s.ResourceId = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList) SetResourceName(v string) *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList {
	s.ResourceName = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList) SetResourceStatus(v string) *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList {
	s.ResourceStatus = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList) SetResourceType(v string) *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList {
	s.ResourceType = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList) SetTags(v string) *GetResourceComplianceTimelineResponseResourceComplianceTimelineComplianceList {
	s.Tags = &v
	return s
}

type GetResourceConfigurationTimelineRequest struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit        *int    `json:"Limit,omitempty" xml:"Limit,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	MultiAccount *bool   `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	MemberId     *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetResourceConfigurationTimelineRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationTimelineRequest) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationTimelineRequest) SetResourceId(v string) *GetResourceConfigurationTimelineRequest {
	s.ResourceId = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetStartTime(v int64) *GetResourceConfigurationTimelineRequest {
	s.StartTime = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetEndTime(v int64) *GetResourceConfigurationTimelineRequest {
	s.EndTime = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetLimit(v int) *GetResourceConfigurationTimelineRequest {
	s.Limit = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetResourceType(v string) *GetResourceConfigurationTimelineRequest {
	s.ResourceType = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetRegion(v string) *GetResourceConfigurationTimelineRequest {
	s.Region = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetMultiAccount(v bool) *GetResourceConfigurationTimelineRequest {
	s.MultiAccount = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetMemberId(v int64) *GetResourceConfigurationTimelineRequest {
	s.MemberId = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetNextToken(v string) *GetResourceConfigurationTimelineRequest {
	s.NextToken = &v
	return s
}

type GetResourceConfigurationTimelineResponse struct {
	RequestId                     *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResourceConfigurationTimeline *GetResourceConfigurationTimelineResponseResourceConfigurationTimeline `json:"ResourceConfigurationTimeline,omitempty" xml:"ResourceConfigurationTimeline,omitempty" require:"true" type:"Struct"`
}

func (s GetResourceConfigurationTimelineResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationTimelineResponse) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationTimelineResponse) SetRequestId(v string) *GetResourceConfigurationTimelineResponse {
	s.RequestId = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponse) SetResourceConfigurationTimeline(v *GetResourceConfigurationTimelineResponseResourceConfigurationTimeline) *GetResourceConfigurationTimelineResponse {
	s.ResourceConfigurationTimeline = v
	return s
}

type GetResourceConfigurationTimelineResponseResourceConfigurationTimeline struct {
	NextToken         *string                                                                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	Limit             *int                                                                                      `json:"Limit,omitempty" xml:"Limit,omitempty" require:"true"`
	TotalCount        *int64                                                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	ConfigurationList []*GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList `json:"ConfigurationList,omitempty" xml:"ConfigurationList,omitempty" require:"true" type:"Repeated"`
}

func (s GetResourceConfigurationTimelineResponseResourceConfigurationTimeline) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationTimelineResponseResourceConfigurationTimeline) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationTimelineResponseResourceConfigurationTimeline) SetNextToken(v string) *GetResourceConfigurationTimelineResponseResourceConfigurationTimeline {
	s.NextToken = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseResourceConfigurationTimeline) SetLimit(v int) *GetResourceConfigurationTimelineResponseResourceConfigurationTimeline {
	s.Limit = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseResourceConfigurationTimeline) SetTotalCount(v int64) *GetResourceConfigurationTimelineResponseResourceConfigurationTimeline {
	s.TotalCount = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseResourceConfigurationTimeline) SetConfigurationList(v []*GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList) *GetResourceConfigurationTimelineResponseResourceConfigurationTimeline {
	s.ConfigurationList = v
	return s
}

type GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList struct {
	AccountId          *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	AvailabilityZone   *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty" require:"true"`
	CaptureTime        *string `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty" require:"true"`
	ConfigurationDiff  *string `json:"ConfigurationDiff,omitempty" xml:"ConfigurationDiff,omitempty" require:"true"`
	Region             *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	Relationship       *string `json:"Relationship,omitempty" xml:"Relationship,omitempty" require:"true"`
	RelationshipDiff   *string `json:"RelationshipDiff,omitempty" xml:"RelationshipDiff,omitempty" require:"true"`
	ResourceCreateTime *string `json:"ResourceCreateTime,omitempty" xml:"ResourceCreateTime,omitempty" require:"true"`
	ResourceId         *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	ResourceName       *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty" require:"true"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	Tags               *string `json:"Tags,omitempty" xml:"Tags,omitempty" require:"true"`
}

func (s GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList) SetAccountId(v int64) *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList {
	s.AccountId = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList) SetAvailabilityZone(v string) *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList {
	s.AvailabilityZone = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList) SetCaptureTime(v string) *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList {
	s.CaptureTime = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList) SetConfigurationDiff(v string) *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList {
	s.ConfigurationDiff = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList) SetRegion(v string) *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList {
	s.Region = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList) SetRelationship(v string) *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList {
	s.Relationship = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList) SetRelationshipDiff(v string) *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList {
	s.RelationshipDiff = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList) SetResourceCreateTime(v string) *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList {
	s.ResourceCreateTime = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList) SetResourceId(v string) *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList {
	s.ResourceId = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList) SetResourceName(v string) *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList {
	s.ResourceName = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList) SetResourceType(v string) *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList {
	s.ResourceType = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList) SetTags(v string) *GetResourceConfigurationTimelineResponseResourceConfigurationTimelineConfigurationList {
	s.Tags = &v
	return s
}

type DescribeDeliveryChannelsRequest struct {
	DeliveryChannelIds *string `json:"DeliveryChannelIds,omitempty" xml:"DeliveryChannelIds,omitempty"`
}

func (s DescribeDeliveryChannelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryChannelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryChannelsRequest) SetDeliveryChannelIds(v string) *DescribeDeliveryChannelsRequest {
	s.DeliveryChannelIds = &v
	return s
}

type DescribeDeliveryChannelsResponse struct {
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DeliveryChannels []*DescribeDeliveryChannelsResponseDeliveryChannels `json:"DeliveryChannels,omitempty" xml:"DeliveryChannels,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDeliveryChannelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryChannelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryChannelsResponse) SetRequestId(v string) *DescribeDeliveryChannelsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDeliveryChannelsResponse) SetDeliveryChannels(v []*DescribeDeliveryChannelsResponseDeliveryChannels) *DescribeDeliveryChannelsResponse {
	s.DeliveryChannels = v
	return s
}

type DescribeDeliveryChannelsResponseDeliveryChannels struct {
	DeliveryChannelId            *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty" require:"true"`
	DeliveryChannelName          *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty" require:"true"`
	DeliveryChannelType          *string `json:"DeliveryChannelType,omitempty" xml:"DeliveryChannelType,omitempty" require:"true"`
	DeliveryChannelTargetArn     *string `json:"DeliveryChannelTargetArn,omitempty" xml:"DeliveryChannelTargetArn,omitempty" require:"true"`
	DeliveryChannelAssumeRoleArn *string `json:"DeliveryChannelAssumeRoleArn,omitempty" xml:"DeliveryChannelAssumeRoleArn,omitempty" require:"true"`
	DeliveryChannelCondition     *string `json:"DeliveryChannelCondition,omitempty" xml:"DeliveryChannelCondition,omitempty" require:"true"`
	Description                  *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	Status                       *int    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s DescribeDeliveryChannelsResponseDeliveryChannels) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryChannelsResponseDeliveryChannels) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryChannelsResponseDeliveryChannels) SetDeliveryChannelId(v string) *DescribeDeliveryChannelsResponseDeliveryChannels {
	s.DeliveryChannelId = &v
	return s
}

func (s *DescribeDeliveryChannelsResponseDeliveryChannels) SetDeliveryChannelName(v string) *DescribeDeliveryChannelsResponseDeliveryChannels {
	s.DeliveryChannelName = &v
	return s
}

func (s *DescribeDeliveryChannelsResponseDeliveryChannels) SetDeliveryChannelType(v string) *DescribeDeliveryChannelsResponseDeliveryChannels {
	s.DeliveryChannelType = &v
	return s
}

func (s *DescribeDeliveryChannelsResponseDeliveryChannels) SetDeliveryChannelTargetArn(v string) *DescribeDeliveryChannelsResponseDeliveryChannels {
	s.DeliveryChannelTargetArn = &v
	return s
}

func (s *DescribeDeliveryChannelsResponseDeliveryChannels) SetDeliveryChannelAssumeRoleArn(v string) *DescribeDeliveryChannelsResponseDeliveryChannels {
	s.DeliveryChannelAssumeRoleArn = &v
	return s
}

func (s *DescribeDeliveryChannelsResponseDeliveryChannels) SetDeliveryChannelCondition(v string) *DescribeDeliveryChannelsResponseDeliveryChannels {
	s.DeliveryChannelCondition = &v
	return s
}

func (s *DescribeDeliveryChannelsResponseDeliveryChannels) SetDescription(v string) *DescribeDeliveryChannelsResponseDeliveryChannels {
	s.Description = &v
	return s
}

func (s *DescribeDeliveryChannelsResponseDeliveryChannels) SetStatus(v int) *DescribeDeliveryChannelsResponseDeliveryChannels {
	s.Status = &v
	return s
}

type PutConfigurationRecorderRequest struct {
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" require:"true"`
}

func (s PutConfigurationRecorderRequest) String() string {
	return tea.Prettify(s)
}

func (s PutConfigurationRecorderRequest) GoString() string {
	return s.String()
}

func (s *PutConfigurationRecorderRequest) SetResourceTypes(v string) *PutConfigurationRecorderRequest {
	s.ResourceTypes = &v
	return s
}

type PutConfigurationRecorderResponse struct {
	RequestId             *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ConfigurationRecorder *PutConfigurationRecorderResponseConfigurationRecorder `json:"ConfigurationRecorder,omitempty" xml:"ConfigurationRecorder,omitempty" require:"true" type:"Struct"`
}

func (s PutConfigurationRecorderResponse) String() string {
	return tea.Prettify(s)
}

func (s PutConfigurationRecorderResponse) GoString() string {
	return s.String()
}

func (s *PutConfigurationRecorderResponse) SetRequestId(v string) *PutConfigurationRecorderResponse {
	s.RequestId = &v
	return s
}

func (s *PutConfigurationRecorderResponse) SetConfigurationRecorder(v *PutConfigurationRecorderResponseConfigurationRecorder) *PutConfigurationRecorderResponse {
	s.ConfigurationRecorder = v
	return s
}

type PutConfigurationRecorderResponseConfigurationRecorder struct {
	AccountId                   *int64    `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	ConfigurationRecorderStatus *string   `json:"ConfigurationRecorderStatus,omitempty" xml:"ConfigurationRecorderStatus,omitempty" require:"true"`
	OrganizationMasterId        *int64    `json:"OrganizationMasterId,omitempty" xml:"OrganizationMasterId,omitempty" require:"true"`
	OrganizationEnableStatus    *string   `json:"OrganizationEnableStatus,omitempty" xml:"OrganizationEnableStatus,omitempty" require:"true"`
	ResourceTypes               []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" require:"true" type:"Repeated"`
}

func (s PutConfigurationRecorderResponseConfigurationRecorder) String() string {
	return tea.Prettify(s)
}

func (s PutConfigurationRecorderResponseConfigurationRecorder) GoString() string {
	return s.String()
}

func (s *PutConfigurationRecorderResponseConfigurationRecorder) SetAccountId(v int64) *PutConfigurationRecorderResponseConfigurationRecorder {
	s.AccountId = &v
	return s
}

func (s *PutConfigurationRecorderResponseConfigurationRecorder) SetConfigurationRecorderStatus(v string) *PutConfigurationRecorderResponseConfigurationRecorder {
	s.ConfigurationRecorderStatus = &v
	return s
}

func (s *PutConfigurationRecorderResponseConfigurationRecorder) SetOrganizationMasterId(v int64) *PutConfigurationRecorderResponseConfigurationRecorder {
	s.OrganizationMasterId = &v
	return s
}

func (s *PutConfigurationRecorderResponseConfigurationRecorder) SetOrganizationEnableStatus(v string) *PutConfigurationRecorderResponseConfigurationRecorder {
	s.OrganizationEnableStatus = &v
	return s
}

func (s *PutConfigurationRecorderResponseConfigurationRecorder) SetResourceTypes(v []*string) *PutConfigurationRecorderResponseConfigurationRecorder {
	s.ResourceTypes = v
	return s
}

type GetDiscoveredResourceSummaryRequest struct {
	MultiAccount *bool  `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	MemberId     *int64 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s GetDiscoveredResourceSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceSummaryRequest) SetMultiAccount(v bool) *GetDiscoveredResourceSummaryRequest {
	s.MultiAccount = &v
	return s
}

func (s *GetDiscoveredResourceSummaryRequest) SetMemberId(v int64) *GetDiscoveredResourceSummaryRequest {
	s.MemberId = &v
	return s
}

type GetDiscoveredResourceSummaryResponse struct {
	RequestId                 *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DiscoveredResourceSummary *GetDiscoveredResourceSummaryResponseDiscoveredResourceSummary `json:"DiscoveredResourceSummary,omitempty" xml:"DiscoveredResourceSummary,omitempty" require:"true" type:"Struct"`
}

func (s GetDiscoveredResourceSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceSummaryResponse) SetRequestId(v string) *GetDiscoveredResourceSummaryResponse {
	s.RequestId = &v
	return s
}

func (s *GetDiscoveredResourceSummaryResponse) SetDiscoveredResourceSummary(v *GetDiscoveredResourceSummaryResponseDiscoveredResourceSummary) *GetDiscoveredResourceSummaryResponse {
	s.DiscoveredResourceSummary = v
	return s
}

type GetDiscoveredResourceSummaryResponseDiscoveredResourceSummary struct {
	ResourceCount     *int `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty" require:"true"`
	ResourceTypeCount *int `json:"ResourceTypeCount,omitempty" xml:"ResourceTypeCount,omitempty" require:"true"`
	RegionCount       *int `json:"RegionCount,omitempty" xml:"RegionCount,omitempty" require:"true"`
}

func (s GetDiscoveredResourceSummaryResponseDiscoveredResourceSummary) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceSummaryResponseDiscoveredResourceSummary) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceSummaryResponseDiscoveredResourceSummary) SetResourceCount(v int) *GetDiscoveredResourceSummaryResponseDiscoveredResourceSummary {
	s.ResourceCount = &v
	return s
}

func (s *GetDiscoveredResourceSummaryResponseDiscoveredResourceSummary) SetResourceTypeCount(v int) *GetDiscoveredResourceSummaryResponseDiscoveredResourceSummary {
	s.ResourceTypeCount = &v
	return s
}

func (s *GetDiscoveredResourceSummaryResponseDiscoveredResourceSummary) SetRegionCount(v int) *GetDiscoveredResourceSummaryResponseDiscoveredResourceSummary {
	s.RegionCount = &v
	return s
}

type ActiveConfigRulesRequest struct {
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty" require:"true"`
}

func (s ActiveConfigRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ActiveConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *ActiveConfigRulesRequest) SetConfigRuleIds(v string) *ActiveConfigRulesRequest {
	s.ConfigRuleIds = &v
	return s
}

type ActiveConfigRulesResponse struct {
	RequestId         *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	OperateRuleResult *ActiveConfigRulesResponseOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" require:"true" type:"Struct"`
}

func (s ActiveConfigRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ActiveConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *ActiveConfigRulesResponse) SetRequestId(v string) *ActiveConfigRulesResponse {
	s.RequestId = &v
	return s
}

func (s *ActiveConfigRulesResponse) SetOperateRuleResult(v *ActiveConfigRulesResponseOperateRuleResult) *ActiveConfigRulesResponse {
	s.OperateRuleResult = v
	return s
}

type ActiveConfigRulesResponseOperateRuleResult struct {
	OperateRuleItemList []*ActiveConfigRulesResponseOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" require:"true" type:"Repeated"`
}

func (s ActiveConfigRulesResponseOperateRuleResult) String() string {
	return tea.Prettify(s)
}

func (s ActiveConfigRulesResponseOperateRuleResult) GoString() string {
	return s.String()
}

func (s *ActiveConfigRulesResponseOperateRuleResult) SetOperateRuleItemList(v []*ActiveConfigRulesResponseOperateRuleResultOperateRuleItemList) *ActiveConfigRulesResponseOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

type ActiveConfigRulesResponseOperateRuleResultOperateRuleItemList struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty" require:"true"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty" require:"true"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s ActiveConfigRulesResponseOperateRuleResultOperateRuleItemList) String() string {
	return tea.Prettify(s)
}

func (s ActiveConfigRulesResponseOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *ActiveConfigRulesResponseOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *ActiveConfigRulesResponseOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *ActiveConfigRulesResponseOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *ActiveConfigRulesResponseOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *ActiveConfigRulesResponseOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *ActiveConfigRulesResponseOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

type StopConfigRulesRequest struct {
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty" require:"true"`
}

func (s StopConfigRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s StopConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *StopConfigRulesRequest) SetConfigRuleIds(v string) *StopConfigRulesRequest {
	s.ConfigRuleIds = &v
	return s
}

type StopConfigRulesResponse struct {
	RequestId         *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	OperateRuleResult *StopConfigRulesResponseOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" require:"true" type:"Struct"`
}

func (s StopConfigRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s StopConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *StopConfigRulesResponse) SetRequestId(v string) *StopConfigRulesResponse {
	s.RequestId = &v
	return s
}

func (s *StopConfigRulesResponse) SetOperateRuleResult(v *StopConfigRulesResponseOperateRuleResult) *StopConfigRulesResponse {
	s.OperateRuleResult = v
	return s
}

type StopConfigRulesResponseOperateRuleResult struct {
	OperateRuleItemList []*StopConfigRulesResponseOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" require:"true" type:"Repeated"`
}

func (s StopConfigRulesResponseOperateRuleResult) String() string {
	return tea.Prettify(s)
}

func (s StopConfigRulesResponseOperateRuleResult) GoString() string {
	return s.String()
}

func (s *StopConfigRulesResponseOperateRuleResult) SetOperateRuleItemList(v []*StopConfigRulesResponseOperateRuleResultOperateRuleItemList) *StopConfigRulesResponseOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

type StopConfigRulesResponseOperateRuleResultOperateRuleItemList struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty" require:"true"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty" require:"true"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s StopConfigRulesResponseOperateRuleResultOperateRuleItemList) String() string {
	return tea.Prettify(s)
}

func (s StopConfigRulesResponseOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *StopConfigRulesResponseOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *StopConfigRulesResponseOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *StopConfigRulesResponseOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *StopConfigRulesResponseOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *StopConfigRulesResponseOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *StopConfigRulesResponseOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

type DescribeComplianceSummaryRequest struct {
	MultiAccount *bool  `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	MemberId     *int64 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s DescribeComplianceSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeComplianceSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeComplianceSummaryRequest) SetMultiAccount(v bool) *DescribeComplianceSummaryRequest {
	s.MultiAccount = &v
	return s
}

func (s *DescribeComplianceSummaryRequest) SetMemberId(v int64) *DescribeComplianceSummaryRequest {
	s.MemberId = &v
	return s
}

type DescribeComplianceSummaryResponse struct {
	RequestId         *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ComplianceSummary *DescribeComplianceSummaryResponseComplianceSummary `json:"ComplianceSummary,omitempty" xml:"ComplianceSummary,omitempty" require:"true" type:"Struct"`
}

func (s DescribeComplianceSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeComplianceSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeComplianceSummaryResponse) SetRequestId(v string) *DescribeComplianceSummaryResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeComplianceSummaryResponse) SetComplianceSummary(v *DescribeComplianceSummaryResponseComplianceSummary) *DescribeComplianceSummaryResponse {
	s.ComplianceSummary = v
	return s
}

type DescribeComplianceSummaryResponseComplianceSummary struct {
	ComplianceSummaryByConfigRule *DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByConfigRule `json:"ComplianceSummaryByConfigRule,omitempty" xml:"ComplianceSummaryByConfigRule,omitempty" require:"true" type:"Struct"`
	ComplianceSummaryByResource   *DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByResource   `json:"ComplianceSummaryByResource,omitempty" xml:"ComplianceSummaryByResource,omitempty" require:"true" type:"Struct"`
}

func (s DescribeComplianceSummaryResponseComplianceSummary) String() string {
	return tea.Prettify(s)
}

func (s DescribeComplianceSummaryResponseComplianceSummary) GoString() string {
	return s.String()
}

func (s *DescribeComplianceSummaryResponseComplianceSummary) SetComplianceSummaryByConfigRule(v *DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByConfigRule) *DescribeComplianceSummaryResponseComplianceSummary {
	s.ComplianceSummaryByConfigRule = v
	return s
}

func (s *DescribeComplianceSummaryResponseComplianceSummary) SetComplianceSummaryByResource(v *DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByResource) *DescribeComplianceSummaryResponseComplianceSummary {
	s.ComplianceSummaryByResource = v
	return s
}

type DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByConfigRule struct {
	ComplianceSummaryTimestamp *int64 `json:"ComplianceSummaryTimestamp,omitempty" xml:"ComplianceSummaryTimestamp,omitempty" require:"true"`
	CompliantCount             *int   `json:"CompliantCount,omitempty" xml:"CompliantCount,omitempty" require:"true"`
	NonCompliantCount          *int   `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty" require:"true"`
	TotalCount                 *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
}

func (s DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByConfigRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByConfigRule) GoString() string {
	return s.String()
}

func (s *DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByConfigRule) SetComplianceSummaryTimestamp(v int64) *DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByConfigRule {
	s.ComplianceSummaryTimestamp = &v
	return s
}

func (s *DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByConfigRule) SetCompliantCount(v int) *DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByConfigRule {
	s.CompliantCount = &v
	return s
}

func (s *DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByConfigRule) SetNonCompliantCount(v int) *DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByConfigRule {
	s.NonCompliantCount = &v
	return s
}

func (s *DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByConfigRule) SetTotalCount(v int64) *DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByConfigRule {
	s.TotalCount = &v
	return s
}

type DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByResource struct {
	ComplianceSummaryTimestamp *int64 `json:"ComplianceSummaryTimestamp,omitempty" xml:"ComplianceSummaryTimestamp,omitempty" require:"true"`
	CompliantCount             *int   `json:"CompliantCount,omitempty" xml:"CompliantCount,omitempty" require:"true"`
	NonCompliantCount          *int   `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty" require:"true"`
	TotalCount                 *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
}

func (s DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByResource) GoString() string {
	return s.String()
}

func (s *DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByResource) SetComplianceSummaryTimestamp(v int64) *DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByResource {
	s.ComplianceSummaryTimestamp = &v
	return s
}

func (s *DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByResource) SetCompliantCount(v int) *DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByResource {
	s.CompliantCount = &v
	return s
}

func (s *DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByResource) SetNonCompliantCount(v int) *DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByResource {
	s.NonCompliantCount = &v
	return s
}

func (s *DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByResource) SetTotalCount(v int64) *DescribeComplianceSummaryResponseComplianceSummaryComplianceSummaryByResource {
	s.TotalCount = &v
	return s
}

type ListConfigRulesRequest struct {
	ConfigRuleState *string `json:"ConfigRuleState,omitempty" xml:"ConfigRuleState,omitempty"`
	ComplianceType  *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	RiskLevel       *int    `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	MessageType     *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	PageNumber      *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	MultiAccount    *bool   `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	MemberId        *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s ListConfigRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *ListConfigRulesRequest) SetConfigRuleState(v string) *ListConfigRulesRequest {
	s.ConfigRuleState = &v
	return s
}

func (s *ListConfigRulesRequest) SetComplianceType(v string) *ListConfigRulesRequest {
	s.ComplianceType = &v
	return s
}

func (s *ListConfigRulesRequest) SetRiskLevel(v int) *ListConfigRulesRequest {
	s.RiskLevel = &v
	return s
}

func (s *ListConfigRulesRequest) SetMessageType(v string) *ListConfigRulesRequest {
	s.MessageType = &v
	return s
}

func (s *ListConfigRulesRequest) SetPageNumber(v int) *ListConfigRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConfigRulesRequest) SetPageSize(v int) *ListConfigRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListConfigRulesRequest) SetMultiAccount(v bool) *ListConfigRulesRequest {
	s.MultiAccount = &v
	return s
}

func (s *ListConfigRulesRequest) SetMemberId(v int64) *ListConfigRulesRequest {
	s.MemberId = &v
	return s
}

type ListConfigRulesResponse struct {
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ConfigRules *ListConfigRulesResponseConfigRules `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" require:"true" type:"Struct"`
}

func (s ListConfigRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *ListConfigRulesResponse) SetRequestId(v string) *ListConfigRulesResponse {
	s.RequestId = &v
	return s
}

func (s *ListConfigRulesResponse) SetConfigRules(v *ListConfigRulesResponseConfigRules) *ListConfigRulesResponse {
	s.ConfigRules = v
	return s
}

type ListConfigRulesResponseConfigRules struct {
	PageNumber     *int                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize       *int                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount     *int64                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	ConfigRuleList []*ListConfigRulesResponseConfigRulesConfigRuleList `json:"ConfigRuleList,omitempty" xml:"ConfigRuleList,omitempty" require:"true" type:"Repeated"`
}

func (s ListConfigRulesResponseConfigRules) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRulesResponseConfigRules) GoString() string {
	return s.String()
}

func (s *ListConfigRulesResponseConfigRules) SetPageNumber(v int) *ListConfigRulesResponseConfigRules {
	s.PageNumber = &v
	return s
}

func (s *ListConfigRulesResponseConfigRules) SetPageSize(v int) *ListConfigRulesResponseConfigRules {
	s.PageSize = &v
	return s
}

func (s *ListConfigRulesResponseConfigRules) SetTotalCount(v int64) *ListConfigRulesResponseConfigRules {
	s.TotalCount = &v
	return s
}

func (s *ListConfigRulesResponseConfigRules) SetConfigRuleList(v []*ListConfigRulesResponseConfigRulesConfigRuleList) *ListConfigRulesResponseConfigRules {
	s.ConfigRuleList = v
	return s
}

type ListConfigRulesResponseConfigRulesConfigRuleList struct {
	AccountId        *int64                                                      `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	ConfigRuleArn    *string                                                     `json:"ConfigRuleArn,omitempty" xml:"ConfigRuleArn,omitempty" require:"true"`
	ConfigRuleId     *string                                                     `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty" require:"true"`
	ConfigRuleName   *string                                                     `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty" require:"true"`
	ConfigRuleState  *string                                                     `json:"ConfigRuleState,omitempty" xml:"ConfigRuleState,omitempty" require:"true"`
	Description      *string                                                     `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	RiskLevel        *int                                                        `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" require:"true"`
	SourceIdentifier *string                                                     `json:"SourceIdentifier,omitempty" xml:"SourceIdentifier,omitempty" require:"true"`
	SourceOwner      *string                                                     `json:"SourceOwner,omitempty" xml:"SourceOwner,omitempty" require:"true"`
	AutomationType   *string                                                     `json:"AutomationType,omitempty" xml:"AutomationType,omitempty" require:"true"`
	Compliance       *ListConfigRulesResponseConfigRulesConfigRuleListCompliance `json:"Compliance,omitempty" xml:"Compliance,omitempty" require:"true" type:"Struct"`
	CreateBy         *ListConfigRulesResponseConfigRulesConfigRuleListCreateBy   `json:"CreateBy,omitempty" xml:"CreateBy,omitempty" require:"true" type:"Struct"`
}

func (s ListConfigRulesResponseConfigRulesConfigRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRulesResponseConfigRulesConfigRuleList) GoString() string {
	return s.String()
}

func (s *ListConfigRulesResponseConfigRulesConfigRuleList) SetAccountId(v int64) *ListConfigRulesResponseConfigRulesConfigRuleList {
	s.AccountId = &v
	return s
}

func (s *ListConfigRulesResponseConfigRulesConfigRuleList) SetConfigRuleArn(v string) *ListConfigRulesResponseConfigRulesConfigRuleList {
	s.ConfigRuleArn = &v
	return s
}

func (s *ListConfigRulesResponseConfigRulesConfigRuleList) SetConfigRuleId(v string) *ListConfigRulesResponseConfigRulesConfigRuleList {
	s.ConfigRuleId = &v
	return s
}

func (s *ListConfigRulesResponseConfigRulesConfigRuleList) SetConfigRuleName(v string) *ListConfigRulesResponseConfigRulesConfigRuleList {
	s.ConfigRuleName = &v
	return s
}

func (s *ListConfigRulesResponseConfigRulesConfigRuleList) SetConfigRuleState(v string) *ListConfigRulesResponseConfigRulesConfigRuleList {
	s.ConfigRuleState = &v
	return s
}

func (s *ListConfigRulesResponseConfigRulesConfigRuleList) SetDescription(v string) *ListConfigRulesResponseConfigRulesConfigRuleList {
	s.Description = &v
	return s
}

func (s *ListConfigRulesResponseConfigRulesConfigRuleList) SetRiskLevel(v int) *ListConfigRulesResponseConfigRulesConfigRuleList {
	s.RiskLevel = &v
	return s
}

func (s *ListConfigRulesResponseConfigRulesConfigRuleList) SetSourceIdentifier(v string) *ListConfigRulesResponseConfigRulesConfigRuleList {
	s.SourceIdentifier = &v
	return s
}

func (s *ListConfigRulesResponseConfigRulesConfigRuleList) SetSourceOwner(v string) *ListConfigRulesResponseConfigRulesConfigRuleList {
	s.SourceOwner = &v
	return s
}

func (s *ListConfigRulesResponseConfigRulesConfigRuleList) SetAutomationType(v string) *ListConfigRulesResponseConfigRulesConfigRuleList {
	s.AutomationType = &v
	return s
}

func (s *ListConfigRulesResponseConfigRulesConfigRuleList) SetCompliance(v *ListConfigRulesResponseConfigRulesConfigRuleListCompliance) *ListConfigRulesResponseConfigRulesConfigRuleList {
	s.Compliance = v
	return s
}

func (s *ListConfigRulesResponseConfigRulesConfigRuleList) SetCreateBy(v *ListConfigRulesResponseConfigRulesConfigRuleListCreateBy) *ListConfigRulesResponseConfigRulesConfigRuleList {
	s.CreateBy = v
	return s
}

type ListConfigRulesResponseConfigRulesConfigRuleListCompliance struct {
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty" require:"true"`
	Count          *int    `json:"Count,omitempty" xml:"Count,omitempty" require:"true"`
}

func (s ListConfigRulesResponseConfigRulesConfigRuleListCompliance) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRulesResponseConfigRulesConfigRuleListCompliance) GoString() string {
	return s.String()
}

func (s *ListConfigRulesResponseConfigRulesConfigRuleListCompliance) SetComplianceType(v string) *ListConfigRulesResponseConfigRulesConfigRuleListCompliance {
	s.ComplianceType = &v
	return s
}

func (s *ListConfigRulesResponseConfigRulesConfigRuleListCompliance) SetCount(v int) *ListConfigRulesResponseConfigRulesConfigRuleListCompliance {
	s.Count = &v
	return s
}

type ListConfigRulesResponseConfigRulesConfigRuleListCreateBy struct {
	CreatorId           *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty" require:"true"`
	CreatorName         *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty" require:"true"`
	CreatorType         *string `json:"CreatorType,omitempty" xml:"CreatorType,omitempty" require:"true"`
	ConfigRuleSceneId   *string `json:"ConfigRuleSceneId,omitempty" xml:"ConfigRuleSceneId,omitempty" require:"true"`
	ConfigRuleSceneName *string `json:"ConfigRuleSceneName,omitempty" xml:"ConfigRuleSceneName,omitempty" require:"true"`
}

func (s ListConfigRulesResponseConfigRulesConfigRuleListCreateBy) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRulesResponseConfigRulesConfigRuleListCreateBy) GoString() string {
	return s.String()
}

func (s *ListConfigRulesResponseConfigRulesConfigRuleListCreateBy) SetCreatorId(v string) *ListConfigRulesResponseConfigRulesConfigRuleListCreateBy {
	s.CreatorId = &v
	return s
}

func (s *ListConfigRulesResponseConfigRulesConfigRuleListCreateBy) SetCreatorName(v string) *ListConfigRulesResponseConfigRulesConfigRuleListCreateBy {
	s.CreatorName = &v
	return s
}

func (s *ListConfigRulesResponseConfigRulesConfigRuleListCreateBy) SetCreatorType(v string) *ListConfigRulesResponseConfigRulesConfigRuleListCreateBy {
	s.CreatorType = &v
	return s
}

func (s *ListConfigRulesResponseConfigRulesConfigRuleListCreateBy) SetConfigRuleSceneId(v string) *ListConfigRulesResponseConfigRulesConfigRuleListCreateBy {
	s.ConfigRuleSceneId = &v
	return s
}

func (s *ListConfigRulesResponseConfigRulesConfigRuleListCreateBy) SetConfigRuleSceneName(v string) *ListConfigRulesResponseConfigRulesConfigRuleListCreateBy {
	s.ConfigRuleSceneName = &v
	return s
}

type PutConfigRuleRequest struct {
	ConfigRuleId                    *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName                  *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty" require:"true"`
	Description                     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InputParameters                 *string `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	SourceOwner                     *string `json:"SourceOwner,omitempty" xml:"SourceOwner,omitempty" require:"true"`
	SourceIdentifier                *string `json:"SourceIdentifier,omitempty" xml:"SourceIdentifier,omitempty" require:"true"`
	SourceDetailMessageType         *string `json:"SourceDetailMessageType,omitempty" xml:"SourceDetailMessageType,omitempty" require:"true"`
	SourceMaximumExecutionFrequency *string `json:"SourceMaximumExecutionFrequency,omitempty" xml:"SourceMaximumExecutionFrequency,omitempty"`
	ScopeComplianceResourceId       *string `json:"ScopeComplianceResourceId,omitempty" xml:"ScopeComplianceResourceId,omitempty"`
	ScopeComplianceResourceTypes    *string `json:"ScopeComplianceResourceTypes,omitempty" xml:"ScopeComplianceResourceTypes,omitempty" require:"true"`
	RiskLevel                       *int    `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" require:"true"`
	ClientToken                     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	MultiAccount                    *bool   `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	MemberId                        *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s PutConfigRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s PutConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *PutConfigRuleRequest) SetConfigRuleId(v string) *PutConfigRuleRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *PutConfigRuleRequest) SetConfigRuleName(v string) *PutConfigRuleRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *PutConfigRuleRequest) SetDescription(v string) *PutConfigRuleRequest {
	s.Description = &v
	return s
}

func (s *PutConfigRuleRequest) SetInputParameters(v string) *PutConfigRuleRequest {
	s.InputParameters = &v
	return s
}

func (s *PutConfigRuleRequest) SetSourceOwner(v string) *PutConfigRuleRequest {
	s.SourceOwner = &v
	return s
}

func (s *PutConfigRuleRequest) SetSourceIdentifier(v string) *PutConfigRuleRequest {
	s.SourceIdentifier = &v
	return s
}

func (s *PutConfigRuleRequest) SetSourceDetailMessageType(v string) *PutConfigRuleRequest {
	s.SourceDetailMessageType = &v
	return s
}

func (s *PutConfigRuleRequest) SetSourceMaximumExecutionFrequency(v string) *PutConfigRuleRequest {
	s.SourceMaximumExecutionFrequency = &v
	return s
}

func (s *PutConfigRuleRequest) SetScopeComplianceResourceId(v string) *PutConfigRuleRequest {
	s.ScopeComplianceResourceId = &v
	return s
}

func (s *PutConfigRuleRequest) SetScopeComplianceResourceTypes(v string) *PutConfigRuleRequest {
	s.ScopeComplianceResourceTypes = &v
	return s
}

func (s *PutConfigRuleRequest) SetRiskLevel(v int) *PutConfigRuleRequest {
	s.RiskLevel = &v
	return s
}

func (s *PutConfigRuleRequest) SetClientToken(v string) *PutConfigRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *PutConfigRuleRequest) SetMultiAccount(v bool) *PutConfigRuleRequest {
	s.MultiAccount = &v
	return s
}

func (s *PutConfigRuleRequest) SetMemberId(v int64) *PutConfigRuleRequest {
	s.MemberId = &v
	return s
}

type PutConfigRuleResponse struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty" require:"true"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s PutConfigRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s PutConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *PutConfigRuleResponse) SetConfigRuleId(v string) *PutConfigRuleResponse {
	s.ConfigRuleId = &v
	return s
}

func (s *PutConfigRuleResponse) SetRequestId(v string) *PutConfigRuleResponse {
	s.RequestId = &v
	return s
}

type DescribeEvaluationResultsRequest struct {
	ResourceType   *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId     *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	PageNumber     *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ConfigRuleId   *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	MultiAccount   *bool   `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	MemberId       *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s DescribeEvaluationResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEvaluationResultsRequest) SetResourceType(v string) *DescribeEvaluationResultsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeEvaluationResultsRequest) SetResourceId(v string) *DescribeEvaluationResultsRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeEvaluationResultsRequest) SetComplianceType(v string) *DescribeEvaluationResultsRequest {
	s.ComplianceType = &v
	return s
}

func (s *DescribeEvaluationResultsRequest) SetPageNumber(v int) *DescribeEvaluationResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEvaluationResultsRequest) SetPageSize(v int) *DescribeEvaluationResultsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEvaluationResultsRequest) SetConfigRuleId(v string) *DescribeEvaluationResultsRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *DescribeEvaluationResultsRequest) SetMultiAccount(v bool) *DescribeEvaluationResultsRequest {
	s.MultiAccount = &v
	return s
}

func (s *DescribeEvaluationResultsRequest) SetMemberId(v int64) *DescribeEvaluationResultsRequest {
	s.MemberId = &v
	return s
}

type DescribeEvaluationResultsResponse struct {
	RequestId         *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	EvaluationResults *DescribeEvaluationResultsResponseEvaluationResults `json:"EvaluationResults,omitempty" xml:"EvaluationResults,omitempty" require:"true" type:"Struct"`
}

func (s DescribeEvaluationResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluationResultsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEvaluationResultsResponse) SetRequestId(v string) *DescribeEvaluationResultsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeEvaluationResultsResponse) SetEvaluationResults(v *DescribeEvaluationResultsResponseEvaluationResults) *DescribeEvaluationResultsResponse {
	s.EvaluationResults = v
	return s
}

type DescribeEvaluationResultsResponseEvaluationResults struct {
	PageNumber           *int                                                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize             *int                                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount           *int64                                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	EvaluationResultList []*DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultList `json:"EvaluationResultList,omitempty" xml:"EvaluationResultList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeEvaluationResultsResponseEvaluationResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluationResultsResponseEvaluationResults) GoString() string {
	return s.String()
}

func (s *DescribeEvaluationResultsResponseEvaluationResults) SetPageNumber(v int) *DescribeEvaluationResultsResponseEvaluationResults {
	s.PageNumber = &v
	return s
}

func (s *DescribeEvaluationResultsResponseEvaluationResults) SetPageSize(v int) *DescribeEvaluationResultsResponseEvaluationResults {
	s.PageSize = &v
	return s
}

func (s *DescribeEvaluationResultsResponseEvaluationResults) SetTotalCount(v int64) *DescribeEvaluationResultsResponseEvaluationResults {
	s.TotalCount = &v
	return s
}

func (s *DescribeEvaluationResultsResponseEvaluationResults) SetEvaluationResultList(v []*DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultList) *DescribeEvaluationResultsResponseEvaluationResults {
	s.EvaluationResultList = v
	return s
}

type DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultList struct {
	Annotation                 *string                                                                                           `json:"Annotation,omitempty" xml:"Annotation,omitempty" require:"true"`
	ComplianceType             *string                                                                                           `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty" require:"true"`
	ConfigRuleInvokedTimestamp *int64                                                                                            `json:"ConfigRuleInvokedTimestamp,omitempty" xml:"ConfigRuleInvokedTimestamp,omitempty" require:"true"`
	InvokingEventMessageType   *string                                                                                           `json:"InvokingEventMessageType,omitempty" xml:"InvokingEventMessageType,omitempty" require:"true"`
	ResultRecordedTimestamp    *int64                                                                                            `json:"ResultRecordedTimestamp,omitempty" xml:"ResultRecordedTimestamp,omitempty" require:"true"`
	RiskLevel                  *int                                                                                              `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" require:"true"`
	EvaluationResultIdentifier *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifier `json:"EvaluationResultIdentifier,omitempty" xml:"EvaluationResultIdentifier,omitempty" require:"true" type:"Struct"`
}

func (s DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultList) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultList) GoString() string {
	return s.String()
}

func (s *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultList) SetAnnotation(v string) *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultList {
	s.Annotation = &v
	return s
}

func (s *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultList) SetComplianceType(v string) *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultList {
	s.ComplianceType = &v
	return s
}

func (s *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultList) SetConfigRuleInvokedTimestamp(v int64) *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultList {
	s.ConfigRuleInvokedTimestamp = &v
	return s
}

func (s *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultList) SetInvokingEventMessageType(v string) *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultList {
	s.InvokingEventMessageType = &v
	return s
}

func (s *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultList) SetResultRecordedTimestamp(v int64) *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultList {
	s.ResultRecordedTimestamp = &v
	return s
}

func (s *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultList) SetRiskLevel(v int) *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultList) SetEvaluationResultIdentifier(v *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifier) *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultList {
	s.EvaluationResultIdentifier = v
	return s
}

type DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifier struct {
	OrderingTimestamp         *int64                                                                                                                     `json:"OrderingTimestamp,omitempty" xml:"OrderingTimestamp,omitempty" require:"true"`
	EvaluationResultQualifier *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier `json:"EvaluationResultQualifier,omitempty" xml:"EvaluationResultQualifier,omitempty" require:"true" type:"Struct"`
}

func (s DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifier) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifier) GoString() string {
	return s.String()
}

func (s *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifier) SetOrderingTimestamp(v int64) *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	s.OrderingTimestamp = &v
	return s
}

func (s *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifier) SetEvaluationResultQualifier(v *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	s.EvaluationResultQualifier = v
	return s
}

type DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier struct {
	ConfigRuleArn  *string `json:"ConfigRuleArn,omitempty" xml:"ConfigRuleArn,omitempty" require:"true"`
	ConfigRuleId   *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty" require:"true"`
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty" require:"true"`
	ResourceId     *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	ResourceType   *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GoString() string {
	return s.String()
}

func (s *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleArn(v string) *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleArn = &v
	return s
}

func (s *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleId(v string) *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleId = &v
	return s
}

func (s *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleName(v string) *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleName = &v
	return s
}

func (s *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceId(v string) *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceId = &v
	return s
}

func (s *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceType(v string) *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceType = &v
	return s
}

func (s *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetRegionId(v string) *DescribeEvaluationResultsResponseEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.RegionId = &v
	return s
}

type DeleteConfigRulesRequest struct {
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty" require:"true"`
}

func (s DeleteConfigRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteConfigRulesRequest) SetConfigRuleIds(v string) *DeleteConfigRulesRequest {
	s.ConfigRuleIds = &v
	return s
}

type DeleteConfigRulesResponse struct {
	RequestId         *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	OperateRuleResult *DeleteConfigRulesResponseOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" require:"true" type:"Struct"`
}

func (s DeleteConfigRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteConfigRulesResponse) SetRequestId(v string) *DeleteConfigRulesResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteConfigRulesResponse) SetOperateRuleResult(v *DeleteConfigRulesResponseOperateRuleResult) *DeleteConfigRulesResponse {
	s.OperateRuleResult = v
	return s
}

type DeleteConfigRulesResponseOperateRuleResult struct {
	OperateRuleItemList []*DeleteConfigRulesResponseOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteConfigRulesResponseOperateRuleResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigRulesResponseOperateRuleResult) GoString() string {
	return s.String()
}

func (s *DeleteConfigRulesResponseOperateRuleResult) SetOperateRuleItemList(v []*DeleteConfigRulesResponseOperateRuleResultOperateRuleItemList) *DeleteConfigRulesResponseOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

type DeleteConfigRulesResponseOperateRuleResultOperateRuleItemList struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty" require:"true"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty" require:"true"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s DeleteConfigRulesResponseOperateRuleResultOperateRuleItemList) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigRulesResponseOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *DeleteConfigRulesResponseOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *DeleteConfigRulesResponseOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *DeleteConfigRulesResponseOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *DeleteConfigRulesResponseOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *DeleteConfigRulesResponseOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *DeleteConfigRulesResponseOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

type DescribeComplianceRequest struct {
	ResourceType   *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId     *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	ConfigRuleId   *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	MultiAccount   *bool   `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	MemberId       *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s DescribeComplianceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeComplianceRequest) GoString() string {
	return s.String()
}

func (s *DescribeComplianceRequest) SetResourceType(v string) *DescribeComplianceRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeComplianceRequest) SetResourceId(v string) *DescribeComplianceRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeComplianceRequest) SetComplianceType(v string) *DescribeComplianceRequest {
	s.ComplianceType = &v
	return s
}

func (s *DescribeComplianceRequest) SetConfigRuleId(v string) *DescribeComplianceRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *DescribeComplianceRequest) SetMultiAccount(v bool) *DescribeComplianceRequest {
	s.MultiAccount = &v
	return s
}

func (s *DescribeComplianceRequest) SetMemberId(v int64) *DescribeComplianceRequest {
	s.MemberId = &v
	return s
}

type DescribeComplianceResponse struct {
	RequestId        *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ComplianceResult *DescribeComplianceResponseComplianceResult `json:"ComplianceResult,omitempty" xml:"ComplianceResult,omitempty" require:"true" type:"Struct"`
}

func (s DescribeComplianceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeComplianceResponse) GoString() string {
	return s.String()
}

func (s *DescribeComplianceResponse) SetRequestId(v string) *DescribeComplianceResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeComplianceResponse) SetComplianceResult(v *DescribeComplianceResponseComplianceResult) *DescribeComplianceResponse {
	s.ComplianceResult = v
	return s
}

type DescribeComplianceResponseComplianceResult struct {
	TotalCount  *int64                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Compliances []*DescribeComplianceResponseComplianceResultCompliances `json:"Compliances,omitempty" xml:"Compliances,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeComplianceResponseComplianceResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeComplianceResponseComplianceResult) GoString() string {
	return s.String()
}

func (s *DescribeComplianceResponseComplianceResult) SetTotalCount(v int64) *DescribeComplianceResponseComplianceResult {
	s.TotalCount = &v
	return s
}

func (s *DescribeComplianceResponseComplianceResult) SetCompliances(v []*DescribeComplianceResponseComplianceResultCompliances) *DescribeComplianceResponseComplianceResult {
	s.Compliances = v
	return s
}

type DescribeComplianceResponseComplianceResultCompliances struct {
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty" require:"true"`
	Count          *int    `json:"Count,omitempty" xml:"Count,omitempty" require:"true"`
}

func (s DescribeComplianceResponseComplianceResultCompliances) String() string {
	return tea.Prettify(s)
}

func (s DescribeComplianceResponseComplianceResultCompliances) GoString() string {
	return s.String()
}

func (s *DescribeComplianceResponseComplianceResultCompliances) SetComplianceType(v string) *DescribeComplianceResponseComplianceResultCompliances {
	s.ComplianceType = &v
	return s
}

func (s *DescribeComplianceResponseComplianceResultCompliances) SetCount(v int) *DescribeComplianceResponseComplianceResultCompliances {
	s.Count = &v
	return s
}

type GetDiscoveredResourceCountsRequest struct {
	GroupByKey   *string `json:"GroupByKey,omitempty" xml:"GroupByKey,omitempty"`
	MultiAccount *bool   `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	MemberId     *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s GetDiscoveredResourceCountsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceCountsRequest) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsRequest) SetGroupByKey(v string) *GetDiscoveredResourceCountsRequest {
	s.GroupByKey = &v
	return s
}

func (s *GetDiscoveredResourceCountsRequest) SetMultiAccount(v bool) *GetDiscoveredResourceCountsRequest {
	s.MultiAccount = &v
	return s
}

func (s *GetDiscoveredResourceCountsRequest) SetMemberId(v int64) *GetDiscoveredResourceCountsRequest {
	s.MemberId = &v
	return s
}

type GetDiscoveredResourceCountsResponse struct {
	RequestId             *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	GroupedResourceCounts *GetDiscoveredResourceCountsResponseGroupedResourceCounts `json:"GroupedResourceCounts,omitempty" xml:"GroupedResourceCounts,omitempty" require:"true" type:"Struct"`
}

func (s GetDiscoveredResourceCountsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceCountsResponse) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsResponse) SetRequestId(v string) *GetDiscoveredResourceCountsResponse {
	s.RequestId = &v
	return s
}

func (s *GetDiscoveredResourceCountsResponse) SetGroupedResourceCounts(v *GetDiscoveredResourceCountsResponseGroupedResourceCounts) *GetDiscoveredResourceCountsResponse {
	s.GroupedResourceCounts = v
	return s
}

type GetDiscoveredResourceCountsResponseGroupedResourceCounts struct {
	GroupByKey               *string                                                                             `json:"GroupByKey,omitempty" xml:"GroupByKey,omitempty" require:"true"`
	GroupedResourceCountList []*GetDiscoveredResourceCountsResponseGroupedResourceCountsGroupedResourceCountList `json:"GroupedResourceCountList,omitempty" xml:"GroupedResourceCountList,omitempty" require:"true" type:"Repeated"`
}

func (s GetDiscoveredResourceCountsResponseGroupedResourceCounts) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceCountsResponseGroupedResourceCounts) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsResponseGroupedResourceCounts) SetGroupByKey(v string) *GetDiscoveredResourceCountsResponseGroupedResourceCounts {
	s.GroupByKey = &v
	return s
}

func (s *GetDiscoveredResourceCountsResponseGroupedResourceCounts) SetGroupedResourceCountList(v []*GetDiscoveredResourceCountsResponseGroupedResourceCountsGroupedResourceCountList) *GetDiscoveredResourceCountsResponseGroupedResourceCounts {
	s.GroupedResourceCountList = v
	return s
}

type GetDiscoveredResourceCountsResponseGroupedResourceCountsGroupedResourceCountList struct {
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty" require:"true"`
	ResourceCount *int64  `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty" require:"true"`
}

func (s GetDiscoveredResourceCountsResponseGroupedResourceCountsGroupedResourceCountList) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceCountsResponseGroupedResourceCountsGroupedResourceCountList) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsResponseGroupedResourceCountsGroupedResourceCountList) SetGroupName(v string) *GetDiscoveredResourceCountsResponseGroupedResourceCountsGroupedResourceCountList {
	s.GroupName = &v
	return s
}

func (s *GetDiscoveredResourceCountsResponseGroupedResourceCountsGroupedResourceCountList) SetResourceCount(v int64) *GetDiscoveredResourceCountsResponseGroupedResourceCountsGroupedResourceCountList {
	s.ResourceCount = &v
	return s
}

type ListDiscoveredResourcesRequest struct {
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceDeleted *int    `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	PageNumber      *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	ResourceTypes   *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
	Regions         *string `json:"Regions,omitempty" xml:"Regions,omitempty"`
	ComplianceType  *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	MultiAccount    *bool   `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	MemberId        *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s ListDiscoveredResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDiscoveredResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListDiscoveredResourcesRequest) SetResourceId(v string) *ListDiscoveredResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetResourceDeleted(v int) *ListDiscoveredResourcesRequest {
	s.ResourceDeleted = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetPageSize(v int) *ListDiscoveredResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetPageNumber(v int) *ListDiscoveredResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetResourceTypes(v string) *ListDiscoveredResourcesRequest {
	s.ResourceTypes = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetRegions(v string) *ListDiscoveredResourcesRequest {
	s.Regions = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetComplianceType(v string) *ListDiscoveredResourcesRequest {
	s.ComplianceType = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetMultiAccount(v bool) *ListDiscoveredResourcesRequest {
	s.MultiAccount = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetMemberId(v int64) *ListDiscoveredResourcesRequest {
	s.MemberId = &v
	return s
}

type ListDiscoveredResourcesResponse struct {
	RequestId                  *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DiscoveredResourceProfiles *ListDiscoveredResourcesResponseDiscoveredResourceProfiles `json:"DiscoveredResourceProfiles,omitempty" xml:"DiscoveredResourceProfiles,omitempty" require:"true" type:"Struct"`
}

func (s ListDiscoveredResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDiscoveredResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListDiscoveredResourcesResponse) SetRequestId(v string) *ListDiscoveredResourcesResponse {
	s.RequestId = &v
	return s
}

func (s *ListDiscoveredResourcesResponse) SetDiscoveredResourceProfiles(v *ListDiscoveredResourcesResponseDiscoveredResourceProfiles) *ListDiscoveredResourcesResponse {
	s.DiscoveredResourceProfiles = v
	return s
}

type ListDiscoveredResourcesResponseDiscoveredResourceProfiles struct {
	PageNumber                    *int                                                                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize                      *int                                                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount                    *int                                                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	DiscoveredResourceProfileList []*ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList `json:"DiscoveredResourceProfileList,omitempty" xml:"DiscoveredResourceProfileList,omitempty" require:"true" type:"Repeated"`
}

func (s ListDiscoveredResourcesResponseDiscoveredResourceProfiles) String() string {
	return tea.Prettify(s)
}

func (s ListDiscoveredResourcesResponseDiscoveredResourceProfiles) GoString() string {
	return s.String()
}

func (s *ListDiscoveredResourcesResponseDiscoveredResourceProfiles) SetPageNumber(v int) *ListDiscoveredResourcesResponseDiscoveredResourceProfiles {
	s.PageNumber = &v
	return s
}

func (s *ListDiscoveredResourcesResponseDiscoveredResourceProfiles) SetPageSize(v int) *ListDiscoveredResourcesResponseDiscoveredResourceProfiles {
	s.PageSize = &v
	return s
}

func (s *ListDiscoveredResourcesResponseDiscoveredResourceProfiles) SetTotalCount(v int) *ListDiscoveredResourcesResponseDiscoveredResourceProfiles {
	s.TotalCount = &v
	return s
}

func (s *ListDiscoveredResourcesResponseDiscoveredResourceProfiles) SetDiscoveredResourceProfileList(v []*ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList) *ListDiscoveredResourcesResponseDiscoveredResourceProfiles {
	s.DiscoveredResourceProfileList = v
	return s
}

type ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList struct {
	AccountId            *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	ResourceCreationTime *int64  `json:"ResourceCreationTime,omitempty" xml:"ResourceCreationTime,omitempty" require:"true"`
	ResourceDeleted      *int    `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty" require:"true"`
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	ResourceName         *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty" require:"true"`
	ResourceStatus       *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty" require:"true"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty" require:"true"`
}

func (s ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList) String() string {
	return tea.Prettify(s)
}

func (s ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList) GoString() string {
	return s.String()
}

func (s *ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList) SetAccountId(v int64) *ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.AccountId = &v
	return s
}

func (s *ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList) SetRegion(v string) *ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.Region = &v
	return s
}

func (s *ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceCreationTime(v int64) *ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceCreationTime = &v
	return s
}

func (s *ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceDeleted(v int) *ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceDeleted = &v
	return s
}

func (s *ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceId(v string) *ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceId = &v
	return s
}

func (s *ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceName(v string) *ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceName = &v
	return s
}

func (s *ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceStatus(v string) *ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceStatus = &v
	return s
}

func (s *ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceType(v string) *ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceType = &v
	return s
}

func (s *ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList) SetTags(v string) *ListDiscoveredResourcesResponseDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.Tags = &v
	return s
}

type DescribeConfigurationRecorderRequest struct {
}

func (s DescribeConfigurationRecorderRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigurationRecorderRequest) GoString() string {
	return s.String()
}

type DescribeConfigurationRecorderResponse struct {
	RequestId             *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ConfigurationRecorder *DescribeConfigurationRecorderResponseConfigurationRecorder `json:"ConfigurationRecorder,omitempty" xml:"ConfigurationRecorder,omitempty" require:"true" type:"Struct"`
}

func (s DescribeConfigurationRecorderResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigurationRecorderResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationRecorderResponse) SetRequestId(v string) *DescribeConfigurationRecorderResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeConfigurationRecorderResponse) SetConfigurationRecorder(v *DescribeConfigurationRecorderResponseConfigurationRecorder) *DescribeConfigurationRecorderResponse {
	s.ConfigurationRecorder = v
	return s
}

type DescribeConfigurationRecorderResponseConfigurationRecorder struct {
	AccountId                   *int64    `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	ConfigurationRecorderStatus *string   `json:"ConfigurationRecorderStatus,omitempty" xml:"ConfigurationRecorderStatus,omitempty" require:"true"`
	OrganizationMasterId        *int64    `json:"OrganizationMasterId,omitempty" xml:"OrganizationMasterId,omitempty" require:"true"`
	OrganizationEnableStatus    *string   `json:"OrganizationEnableStatus,omitempty" xml:"OrganizationEnableStatus,omitempty" require:"true"`
	ResourceTypes               []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeConfigurationRecorderResponseConfigurationRecorder) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigurationRecorderResponseConfigurationRecorder) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationRecorderResponseConfigurationRecorder) SetAccountId(v int64) *DescribeConfigurationRecorderResponseConfigurationRecorder {
	s.AccountId = &v
	return s
}

func (s *DescribeConfigurationRecorderResponseConfigurationRecorder) SetConfigurationRecorderStatus(v string) *DescribeConfigurationRecorderResponseConfigurationRecorder {
	s.ConfigurationRecorderStatus = &v
	return s
}

func (s *DescribeConfigurationRecorderResponseConfigurationRecorder) SetOrganizationMasterId(v int64) *DescribeConfigurationRecorderResponseConfigurationRecorder {
	s.OrganizationMasterId = &v
	return s
}

func (s *DescribeConfigurationRecorderResponseConfigurationRecorder) SetOrganizationEnableStatus(v string) *DescribeConfigurationRecorderResponseConfigurationRecorder {
	s.OrganizationEnableStatus = &v
	return s
}

func (s *DescribeConfigurationRecorderResponseConfigurationRecorder) SetResourceTypes(v []*string) *DescribeConfigurationRecorderResponseConfigurationRecorder {
	s.ResourceTypes = v
	return s
}

type DescribeDiscoveredResourceRequest struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	MultiAccount *bool   `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	MemberId     *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s DescribeDiscoveredResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiscoveredResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiscoveredResourceRequest) SetResourceId(v string) *DescribeDiscoveredResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeDiscoveredResourceRequest) SetResourceType(v string) *DescribeDiscoveredResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeDiscoveredResourceRequest) SetRegion(v string) *DescribeDiscoveredResourceRequest {
	s.Region = &v
	return s
}

func (s *DescribeDiscoveredResourceRequest) SetMultiAccount(v bool) *DescribeDiscoveredResourceRequest {
	s.MultiAccount = &v
	return s
}

func (s *DescribeDiscoveredResourceRequest) SetMemberId(v int64) *DescribeDiscoveredResourceRequest {
	s.MemberId = &v
	return s
}

type DescribeDiscoveredResourceResponse struct {
	RequestId                *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DiscoveredResourceDetail *DescribeDiscoveredResourceResponseDiscoveredResourceDetail `json:"DiscoveredResourceDetail,omitempty" xml:"DiscoveredResourceDetail,omitempty" require:"true" type:"Struct"`
}

func (s DescribeDiscoveredResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiscoveredResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiscoveredResourceResponse) SetRequestId(v string) *DescribeDiscoveredResourceResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDiscoveredResourceResponse) SetDiscoveredResourceDetail(v *DescribeDiscoveredResourceResponseDiscoveredResourceDetail) *DescribeDiscoveredResourceResponse {
	s.DiscoveredResourceDetail = v
	return s
}

type DescribeDiscoveredResourceResponseDiscoveredResourceDetail struct {
	AccountId            *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceName         *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty" require:"true"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	AvailabilityZone     *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty" require:"true"`
	ResourceCreationTime *int64  `json:"ResourceCreationTime,omitempty" xml:"ResourceCreationTime,omitempty" require:"true"`
	ResourceStatus       *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty" require:"true"`
	ResourceDeleted      *int    `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty" require:"true"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty" require:"true"`
	Configuration        *string `json:"Configuration,omitempty" xml:"Configuration,omitempty" require:"true"`
}

func (s DescribeDiscoveredResourceResponseDiscoveredResourceDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiscoveredResourceResponseDiscoveredResourceDetail) GoString() string {
	return s.String()
}

func (s *DescribeDiscoveredResourceResponseDiscoveredResourceDetail) SetAccountId(v int64) *DescribeDiscoveredResourceResponseDiscoveredResourceDetail {
	s.AccountId = &v
	return s
}

func (s *DescribeDiscoveredResourceResponseDiscoveredResourceDetail) SetResourceId(v string) *DescribeDiscoveredResourceResponseDiscoveredResourceDetail {
	s.ResourceId = &v
	return s
}

func (s *DescribeDiscoveredResourceResponseDiscoveredResourceDetail) SetResourceType(v string) *DescribeDiscoveredResourceResponseDiscoveredResourceDetail {
	s.ResourceType = &v
	return s
}

func (s *DescribeDiscoveredResourceResponseDiscoveredResourceDetail) SetResourceName(v string) *DescribeDiscoveredResourceResponseDiscoveredResourceDetail {
	s.ResourceName = &v
	return s
}

func (s *DescribeDiscoveredResourceResponseDiscoveredResourceDetail) SetRegion(v string) *DescribeDiscoveredResourceResponseDiscoveredResourceDetail {
	s.Region = &v
	return s
}

func (s *DescribeDiscoveredResourceResponseDiscoveredResourceDetail) SetAvailabilityZone(v string) *DescribeDiscoveredResourceResponseDiscoveredResourceDetail {
	s.AvailabilityZone = &v
	return s
}

func (s *DescribeDiscoveredResourceResponseDiscoveredResourceDetail) SetResourceCreationTime(v int64) *DescribeDiscoveredResourceResponseDiscoveredResourceDetail {
	s.ResourceCreationTime = &v
	return s
}

func (s *DescribeDiscoveredResourceResponseDiscoveredResourceDetail) SetResourceStatus(v string) *DescribeDiscoveredResourceResponseDiscoveredResourceDetail {
	s.ResourceStatus = &v
	return s
}

func (s *DescribeDiscoveredResourceResponseDiscoveredResourceDetail) SetResourceDeleted(v int) *DescribeDiscoveredResourceResponseDiscoveredResourceDetail {
	s.ResourceDeleted = &v
	return s
}

func (s *DescribeDiscoveredResourceResponseDiscoveredResourceDetail) SetTags(v string) *DescribeDiscoveredResourceResponseDiscoveredResourceDetail {
	s.Tags = &v
	return s
}

func (s *DescribeDiscoveredResourceResponseDiscoveredResourceDetail) SetConfiguration(v string) *DescribeDiscoveredResourceResponseDiscoveredResourceDetail {
	s.Configuration = &v
	return s
}

type StartConfigurationRecorderRequest struct {
	EnterpriseEdition *bool `json:"EnterpriseEdition,omitempty" xml:"EnterpriseEdition,omitempty"`
}

func (s StartConfigurationRecorderRequest) String() string {
	return tea.Prettify(s)
}

func (s StartConfigurationRecorderRequest) GoString() string {
	return s.String()
}

func (s *StartConfigurationRecorderRequest) SetEnterpriseEdition(v bool) *StartConfigurationRecorderRequest {
	s.EnterpriseEdition = &v
	return s
}

type StartConfigurationRecorderResponse struct {
	RequestId             *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ConfigurationRecorder *StartConfigurationRecorderResponseConfigurationRecorder `json:"ConfigurationRecorder,omitempty" xml:"ConfigurationRecorder,omitempty" require:"true" type:"Struct"`
}

func (s StartConfigurationRecorderResponse) String() string {
	return tea.Prettify(s)
}

func (s StartConfigurationRecorderResponse) GoString() string {
	return s.String()
}

func (s *StartConfigurationRecorderResponse) SetRequestId(v string) *StartConfigurationRecorderResponse {
	s.RequestId = &v
	return s
}

func (s *StartConfigurationRecorderResponse) SetConfigurationRecorder(v *StartConfigurationRecorderResponseConfigurationRecorder) *StartConfigurationRecorderResponse {
	s.ConfigurationRecorder = v
	return s
}

type StartConfigurationRecorderResponseConfigurationRecorder struct {
	AccountId                   *int64    `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	ConfigurationRecorderStatus *string   `json:"ConfigurationRecorderStatus,omitempty" xml:"ConfigurationRecorderStatus,omitempty" require:"true"`
	OrganizationMasterId        *int64    `json:"OrganizationMasterId,omitempty" xml:"OrganizationMasterId,omitempty" require:"true"`
	OrganizationEnableStatus    *string   `json:"OrganizationEnableStatus,omitempty" xml:"OrganizationEnableStatus,omitempty" require:"true"`
	ResourceTypes               []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" require:"true" type:"Repeated"`
}

func (s StartConfigurationRecorderResponseConfigurationRecorder) String() string {
	return tea.Prettify(s)
}

func (s StartConfigurationRecorderResponseConfigurationRecorder) GoString() string {
	return s.String()
}

func (s *StartConfigurationRecorderResponseConfigurationRecorder) SetAccountId(v int64) *StartConfigurationRecorderResponseConfigurationRecorder {
	s.AccountId = &v
	return s
}

func (s *StartConfigurationRecorderResponseConfigurationRecorder) SetConfigurationRecorderStatus(v string) *StartConfigurationRecorderResponseConfigurationRecorder {
	s.ConfigurationRecorderStatus = &v
	return s
}

func (s *StartConfigurationRecorderResponseConfigurationRecorder) SetOrganizationMasterId(v int64) *StartConfigurationRecorderResponseConfigurationRecorder {
	s.OrganizationMasterId = &v
	return s
}

func (s *StartConfigurationRecorderResponseConfigurationRecorder) SetOrganizationEnableStatus(v string) *StartConfigurationRecorderResponseConfigurationRecorder {
	s.OrganizationEnableStatus = &v
	return s
}

func (s *StartConfigurationRecorderResponseConfigurationRecorder) SetResourceTypes(v []*string) *StartConfigurationRecorderResponseConfigurationRecorder {
	s.ResourceTypes = v
	return s
}

type DescribeConfigRuleRequest struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty" require:"true"`
	MultiAccount *bool   `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	MemberId     *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s DescribeConfigRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleRequest) SetConfigRuleId(v string) *DescribeConfigRuleRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *DescribeConfigRuleRequest) SetMultiAccount(v bool) *DescribeConfigRuleRequest {
	s.MultiAccount = &v
	return s
}

func (s *DescribeConfigRuleRequest) SetMemberId(v int64) *DescribeConfigRuleRequest {
	s.MemberId = &v
	return s
}

type DescribeConfigRuleResponse struct {
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ConfigRule *DescribeConfigRuleResponseConfigRule `json:"ConfigRule,omitempty" xml:"ConfigRule,omitempty" require:"true" type:"Struct"`
}

func (s DescribeConfigRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleResponse) SetRequestId(v string) *DescribeConfigRuleResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeConfigRuleResponse) SetConfigRule(v *DescribeConfigRuleResponseConfigRule) *DescribeConfigRuleResponse {
	s.ConfigRule = v
	return s
}

type DescribeConfigRuleResponseConfigRule struct {
	ConfigRuleArn              *string                                                         `json:"ConfigRuleArn,omitempty" xml:"ConfigRuleArn,omitempty" require:"true"`
	ConfigRuleId               *string                                                         `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty" require:"true"`
	ConfigRuleName             *string                                                         `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty" require:"true"`
	ConfigRuleState            *string                                                         `json:"ConfigRuleState,omitempty" xml:"ConfigRuleState,omitempty" require:"true"`
	CreateTimestamp            *int64                                                          `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty" require:"true"`
	Description                *string                                                         `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	InputParameters            map[string]interface{}                                          `json:"InputParameters,omitempty" xml:"InputParameters,omitempty" require:"true"`
	ModifiedTimestamp          *int64                                                          `json:"ModifiedTimestamp,omitempty" xml:"ModifiedTimestamp,omitempty" require:"true"`
	RiskLevel                  *int                                                            `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" require:"true"`
	MaximumExecutionFrequency  *string                                                         `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty" require:"true"`
	CreateBy                   *DescribeConfigRuleResponseConfigRuleCreateBy                   `json:"CreateBy,omitempty" xml:"CreateBy,omitempty" require:"true" type:"Struct"`
	ConfigRuleEvaluationStatus *DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus `json:"ConfigRuleEvaluationStatus,omitempty" xml:"ConfigRuleEvaluationStatus,omitempty" require:"true" type:"Struct"`
	ManagedRule                *DescribeConfigRuleResponseConfigRuleManagedRule                `json:"ManagedRule,omitempty" xml:"ManagedRule,omitempty" require:"true" type:"Struct"`
	Source                     *DescribeConfigRuleResponseConfigRuleSource                     `json:"Source,omitempty" xml:"Source,omitempty" require:"true" type:"Struct"`
	Scope                      *DescribeConfigRuleResponseConfigRuleScope                      `json:"Scope,omitempty" xml:"Scope,omitempty" require:"true" type:"Struct"`
}

func (s DescribeConfigRuleResponseConfigRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleResponseConfigRule) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleResponseConfigRule) SetConfigRuleArn(v string) *DescribeConfigRuleResponseConfigRule {
	s.ConfigRuleArn = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRule) SetConfigRuleId(v string) *DescribeConfigRuleResponseConfigRule {
	s.ConfigRuleId = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRule) SetConfigRuleName(v string) *DescribeConfigRuleResponseConfigRule {
	s.ConfigRuleName = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRule) SetConfigRuleState(v string) *DescribeConfigRuleResponseConfigRule {
	s.ConfigRuleState = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRule) SetCreateTimestamp(v int64) *DescribeConfigRuleResponseConfigRule {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRule) SetDescription(v string) *DescribeConfigRuleResponseConfigRule {
	s.Description = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRule) SetInputParameters(v map[string]interface{}) *DescribeConfigRuleResponseConfigRule {
	s.InputParameters = v
	return s
}

func (s *DescribeConfigRuleResponseConfigRule) SetModifiedTimestamp(v int64) *DescribeConfigRuleResponseConfigRule {
	s.ModifiedTimestamp = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRule) SetRiskLevel(v int) *DescribeConfigRuleResponseConfigRule {
	s.RiskLevel = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRule) SetMaximumExecutionFrequency(v string) *DescribeConfigRuleResponseConfigRule {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRule) SetCreateBy(v *DescribeConfigRuleResponseConfigRuleCreateBy) *DescribeConfigRuleResponseConfigRule {
	s.CreateBy = v
	return s
}

func (s *DescribeConfigRuleResponseConfigRule) SetConfigRuleEvaluationStatus(v *DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus) *DescribeConfigRuleResponseConfigRule {
	s.ConfigRuleEvaluationStatus = v
	return s
}

func (s *DescribeConfigRuleResponseConfigRule) SetManagedRule(v *DescribeConfigRuleResponseConfigRuleManagedRule) *DescribeConfigRuleResponseConfigRule {
	s.ManagedRule = v
	return s
}

func (s *DescribeConfigRuleResponseConfigRule) SetSource(v *DescribeConfigRuleResponseConfigRuleSource) *DescribeConfigRuleResponseConfigRule {
	s.Source = v
	return s
}

func (s *DescribeConfigRuleResponseConfigRule) SetScope(v *DescribeConfigRuleResponseConfigRuleScope) *DescribeConfigRuleResponseConfigRule {
	s.Scope = v
	return s
}

type DescribeConfigRuleResponseConfigRuleCreateBy struct {
	CreatorType         *string `json:"CreatorType,omitempty" xml:"CreatorType,omitempty" require:"true"`
	CreatorId           *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty" require:"true"`
	CreatorName         *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty" require:"true"`
	ConfigRuleSceneId   *string `json:"ConfigRuleSceneId,omitempty" xml:"ConfigRuleSceneId,omitempty" require:"true"`
	ConfigRuleSceneName *string `json:"ConfigRuleSceneName,omitempty" xml:"ConfigRuleSceneName,omitempty" require:"true"`
}

func (s DescribeConfigRuleResponseConfigRuleCreateBy) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleResponseConfigRuleCreateBy) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleResponseConfigRuleCreateBy) SetCreatorType(v string) *DescribeConfigRuleResponseConfigRuleCreateBy {
	s.CreatorType = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleCreateBy) SetCreatorId(v string) *DescribeConfigRuleResponseConfigRuleCreateBy {
	s.CreatorId = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleCreateBy) SetCreatorName(v string) *DescribeConfigRuleResponseConfigRuleCreateBy {
	s.CreatorName = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleCreateBy) SetConfigRuleSceneId(v string) *DescribeConfigRuleResponseConfigRuleCreateBy {
	s.ConfigRuleSceneId = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleCreateBy) SetConfigRuleSceneName(v string) *DescribeConfigRuleResponseConfigRuleCreateBy {
	s.ConfigRuleSceneName = &v
	return s
}

type DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus struct {
	FirstActivatedTimestamp           *int64  `json:"FirstActivatedTimestamp,omitempty" xml:"FirstActivatedTimestamp,omitempty" require:"true"`
	FirstEvaluationStarted            *bool   `json:"FirstEvaluationStarted,omitempty" xml:"FirstEvaluationStarted,omitempty" require:"true"`
	LastErrorCode                     *string `json:"LastErrorCode,omitempty" xml:"LastErrorCode,omitempty" require:"true"`
	LastErrorMessage                  *string `json:"LastErrorMessage,omitempty" xml:"LastErrorMessage,omitempty" require:"true"`
	LastFailedEvaluationTimestamp     *int64  `json:"LastFailedEvaluationTimestamp,omitempty" xml:"LastFailedEvaluationTimestamp,omitempty" require:"true"`
	LastFailedInvocationTimestamp     *int64  `json:"LastFailedInvocationTimestamp,omitempty" xml:"LastFailedInvocationTimestamp,omitempty" require:"true"`
	LastSuccessfulEvaluationTimestamp *int64  `json:"LastSuccessfulEvaluationTimestamp,omitempty" xml:"LastSuccessfulEvaluationTimestamp,omitempty" require:"true"`
	LastSuccessfulInvocationTimestamp *int64  `json:"LastSuccessfulInvocationTimestamp,omitempty" xml:"LastSuccessfulInvocationTimestamp,omitempty" require:"true"`
}

func (s DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus) SetFirstActivatedTimestamp(v int64) *DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus {
	s.FirstActivatedTimestamp = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus) SetFirstEvaluationStarted(v bool) *DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus {
	s.FirstEvaluationStarted = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus) SetLastErrorCode(v string) *DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus {
	s.LastErrorCode = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus) SetLastErrorMessage(v string) *DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus {
	s.LastErrorMessage = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus) SetLastFailedEvaluationTimestamp(v int64) *DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus {
	s.LastFailedEvaluationTimestamp = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus) SetLastFailedInvocationTimestamp(v int64) *DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus {
	s.LastFailedInvocationTimestamp = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus) SetLastSuccessfulEvaluationTimestamp(v int64) *DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus {
	s.LastSuccessfulEvaluationTimestamp = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus) SetLastSuccessfulInvocationTimestamp(v int64) *DescribeConfigRuleResponseConfigRuleConfigRuleEvaluationStatus {
	s.LastSuccessfulInvocationTimestamp = &v
	return s
}

type DescribeConfigRuleResponseConfigRuleManagedRule struct {
	ManagedRuleName                 *string                                                         `json:"ManagedRuleName,omitempty" xml:"ManagedRuleName,omitempty" require:"true"`
	Description                     *string                                                         `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	Identifier                      *string                                                         `json:"Identifier,omitempty" xml:"Identifier,omitempty" require:"true"`
	CompulsoryInputParameterDetails map[string]interface{}                                          `json:"CompulsoryInputParameterDetails,omitempty" xml:"CompulsoryInputParameterDetails,omitempty" require:"true"`
	OptionalInputParameterDetails   map[string]interface{}                                          `json:"OptionalInputParameterDetails,omitempty" xml:"OptionalInputParameterDetails,omitempty" require:"true"`
	HelpUrl                         *string                                                         `json:"HelpUrl,omitempty" xml:"HelpUrl,omitempty" require:"true"`
	SourceDetails                   []*DescribeConfigRuleResponseConfigRuleManagedRuleSourceDetails `json:"SourceDetails,omitempty" xml:"SourceDetails,omitempty" require:"true" type:"Repeated"`
	Labels                          []*string                                                       `json:"Labels,omitempty" xml:"Labels,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeConfigRuleResponseConfigRuleManagedRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleResponseConfigRuleManagedRule) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleResponseConfigRuleManagedRule) SetManagedRuleName(v string) *DescribeConfigRuleResponseConfigRuleManagedRule {
	s.ManagedRuleName = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleManagedRule) SetDescription(v string) *DescribeConfigRuleResponseConfigRuleManagedRule {
	s.Description = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleManagedRule) SetIdentifier(v string) *DescribeConfigRuleResponseConfigRuleManagedRule {
	s.Identifier = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleManagedRule) SetCompulsoryInputParameterDetails(v map[string]interface{}) *DescribeConfigRuleResponseConfigRuleManagedRule {
	s.CompulsoryInputParameterDetails = v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleManagedRule) SetOptionalInputParameterDetails(v map[string]interface{}) *DescribeConfigRuleResponseConfigRuleManagedRule {
	s.OptionalInputParameterDetails = v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleManagedRule) SetHelpUrl(v string) *DescribeConfigRuleResponseConfigRuleManagedRule {
	s.HelpUrl = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleManagedRule) SetSourceDetails(v []*DescribeConfigRuleResponseConfigRuleManagedRuleSourceDetails) *DescribeConfigRuleResponseConfigRuleManagedRule {
	s.SourceDetails = v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleManagedRule) SetLabels(v []*string) *DescribeConfigRuleResponseConfigRuleManagedRule {
	s.Labels = v
	return s
}

type DescribeConfigRuleResponseConfigRuleManagedRuleSourceDetails struct {
	EventSource               *string `json:"EventSource,omitempty" xml:"EventSource,omitempty" require:"true"`
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty" require:"true"`
	MessageType               *string `json:"MessageType,omitempty" xml:"MessageType,omitempty" require:"true"`
}

func (s DescribeConfigRuleResponseConfigRuleManagedRuleSourceDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleResponseConfigRuleManagedRuleSourceDetails) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleResponseConfigRuleManagedRuleSourceDetails) SetEventSource(v string) *DescribeConfigRuleResponseConfigRuleManagedRuleSourceDetails {
	s.EventSource = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleManagedRuleSourceDetails) SetMaximumExecutionFrequency(v string) *DescribeConfigRuleResponseConfigRuleManagedRuleSourceDetails {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleManagedRuleSourceDetails) SetMessageType(v string) *DescribeConfigRuleResponseConfigRuleManagedRuleSourceDetails {
	s.MessageType = &v
	return s
}

type DescribeConfigRuleResponseConfigRuleSource struct {
	Identifier       *string                                                       `json:"Identifier,omitempty" xml:"Identifier,omitempty" require:"true"`
	Owner            *string                                                       `json:"Owner,omitempty" xml:"Owner,omitempty" require:"true"`
	SourceConditions []*DescribeConfigRuleResponseConfigRuleSourceSourceConditions `json:"SourceConditions,omitempty" xml:"SourceConditions,omitempty" require:"true" type:"Repeated"`
	SourceDetails    []*DescribeConfigRuleResponseConfigRuleSourceSourceDetails    `json:"SourceDetails,omitempty" xml:"SourceDetails,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeConfigRuleResponseConfigRuleSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleResponseConfigRuleSource) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleResponseConfigRuleSource) SetIdentifier(v string) *DescribeConfigRuleResponseConfigRuleSource {
	s.Identifier = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleSource) SetOwner(v string) *DescribeConfigRuleResponseConfigRuleSource {
	s.Owner = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleSource) SetSourceConditions(v []*DescribeConfigRuleResponseConfigRuleSourceSourceConditions) *DescribeConfigRuleResponseConfigRuleSource {
	s.SourceConditions = v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleSource) SetSourceDetails(v []*DescribeConfigRuleResponseConfigRuleSourceSourceDetails) *DescribeConfigRuleResponseConfigRuleSource {
	s.SourceDetails = v
	return s
}

type DescribeConfigRuleResponseConfigRuleSourceSourceConditions struct {
	DesiredValue *string `json:"DesiredValue,omitempty" xml:"DesiredValue,omitempty" require:"true"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Operator     *string `json:"Operator,omitempty" xml:"Operator,omitempty" require:"true"`
	Required     *bool   `json:"Required,omitempty" xml:"Required,omitempty" require:"true"`
	SelectPath   *string `json:"SelectPath,omitempty" xml:"SelectPath,omitempty" require:"true"`
	Tips         *string `json:"Tips,omitempty" xml:"Tips,omitempty" require:"true"`
}

func (s DescribeConfigRuleResponseConfigRuleSourceSourceConditions) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleResponseConfigRuleSourceSourceConditions) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleResponseConfigRuleSourceSourceConditions) SetDesiredValue(v string) *DescribeConfigRuleResponseConfigRuleSourceSourceConditions {
	s.DesiredValue = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleSourceSourceConditions) SetName(v string) *DescribeConfigRuleResponseConfigRuleSourceSourceConditions {
	s.Name = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleSourceSourceConditions) SetOperator(v string) *DescribeConfigRuleResponseConfigRuleSourceSourceConditions {
	s.Operator = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleSourceSourceConditions) SetRequired(v bool) *DescribeConfigRuleResponseConfigRuleSourceSourceConditions {
	s.Required = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleSourceSourceConditions) SetSelectPath(v string) *DescribeConfigRuleResponseConfigRuleSourceSourceConditions {
	s.SelectPath = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleSourceSourceConditions) SetTips(v string) *DescribeConfigRuleResponseConfigRuleSourceSourceConditions {
	s.Tips = &v
	return s
}

type DescribeConfigRuleResponseConfigRuleSourceSourceDetails struct {
	EventSource               *string `json:"EventSource,omitempty" xml:"EventSource,omitempty" require:"true"`
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty" require:"true"`
	MessageType               *string `json:"MessageType,omitempty" xml:"MessageType,omitempty" require:"true"`
}

func (s DescribeConfigRuleResponseConfigRuleSourceSourceDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleResponseConfigRuleSourceSourceDetails) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleResponseConfigRuleSourceSourceDetails) SetEventSource(v string) *DescribeConfigRuleResponseConfigRuleSourceSourceDetails {
	s.EventSource = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleSourceSourceDetails) SetMaximumExecutionFrequency(v string) *DescribeConfigRuleResponseConfigRuleSourceSourceDetails {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleSourceSourceDetails) SetMessageType(v string) *DescribeConfigRuleResponseConfigRuleSourceSourceDetails {
	s.MessageType = &v
	return s
}

type DescribeConfigRuleResponseConfigRuleScope struct {
	ComplianceResourceId    *string   `json:"ComplianceResourceId,omitempty" xml:"ComplianceResourceId,omitempty" require:"true"`
	ComplianceResourceTypes []*string `json:"ComplianceResourceTypes,omitempty" xml:"ComplianceResourceTypes,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeConfigRuleResponseConfigRuleScope) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleResponseConfigRuleScope) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleResponseConfigRuleScope) SetComplianceResourceId(v string) *DescribeConfigRuleResponseConfigRuleScope {
	s.ComplianceResourceId = &v
	return s
}

func (s *DescribeConfigRuleResponseConfigRuleScope) SetComplianceResourceTypes(v []*string) *DescribeConfigRuleResponseConfigRuleScope {
	s.ComplianceResourceTypes = v
	return s
}

type GetSupportedResourceTypesRequest struct {
}

func (s GetSupportedResourceTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSupportedResourceTypesRequest) GoString() string {
	return s.String()
}

type GetSupportedResourceTypesResponse struct {
	RequestId     *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" require:"true" type:"Repeated"`
}

func (s GetSupportedResourceTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSupportedResourceTypesResponse) GoString() string {
	return s.String()
}

func (s *GetSupportedResourceTypesResponse) SetRequestId(v string) *GetSupportedResourceTypesResponse {
	s.RequestId = &v
	return s
}

func (s *GetSupportedResourceTypesResponse) SetResourceTypes(v []*string) *GetSupportedResourceTypesResponse {
	s.ResourceTypes = v
	return s
}

type PutDeliveryChannelRequest struct {
	ClientToken                  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DeliveryChannelId            *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	DeliveryChannelName          *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	DeliveryChannelType          *string `json:"DeliveryChannelType,omitempty" xml:"DeliveryChannelType,omitempty" require:"true"`
	DeliveryChannelTargetArn     *string `json:"DeliveryChannelTargetArn,omitempty" xml:"DeliveryChannelTargetArn,omitempty" require:"true"`
	DeliveryChannelAssumeRoleArn *string `json:"DeliveryChannelAssumeRoleArn,omitempty" xml:"DeliveryChannelAssumeRoleArn,omitempty" require:"true"`
	DeliveryChannelCondition     *string `json:"DeliveryChannelCondition,omitempty" xml:"DeliveryChannelCondition,omitempty"`
	Description                  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Status                       *int    `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PutDeliveryChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s PutDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *PutDeliveryChannelRequest) SetClientToken(v string) *PutDeliveryChannelRequest {
	s.ClientToken = &v
	return s
}

func (s *PutDeliveryChannelRequest) SetDeliveryChannelId(v string) *PutDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

func (s *PutDeliveryChannelRequest) SetDeliveryChannelName(v string) *PutDeliveryChannelRequest {
	s.DeliveryChannelName = &v
	return s
}

func (s *PutDeliveryChannelRequest) SetDeliveryChannelType(v string) *PutDeliveryChannelRequest {
	s.DeliveryChannelType = &v
	return s
}

func (s *PutDeliveryChannelRequest) SetDeliveryChannelTargetArn(v string) *PutDeliveryChannelRequest {
	s.DeliveryChannelTargetArn = &v
	return s
}

func (s *PutDeliveryChannelRequest) SetDeliveryChannelAssumeRoleArn(v string) *PutDeliveryChannelRequest {
	s.DeliveryChannelAssumeRoleArn = &v
	return s
}

func (s *PutDeliveryChannelRequest) SetDeliveryChannelCondition(v string) *PutDeliveryChannelRequest {
	s.DeliveryChannelCondition = &v
	return s
}

func (s *PutDeliveryChannelRequest) SetDescription(v string) *PutDeliveryChannelRequest {
	s.Description = &v
	return s
}

func (s *PutDeliveryChannelRequest) SetStatus(v int) *PutDeliveryChannelRequest {
	s.Status = &v
	return s
}

type PutDeliveryChannelResponse struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty" require:"true"`
}

func (s PutDeliveryChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s PutDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *PutDeliveryChannelResponse) SetRequestId(v string) *PutDeliveryChannelResponse {
	s.RequestId = &v
	return s
}

func (s *PutDeliveryChannelResponse) SetDeliveryChannelId(v string) *PutDeliveryChannelResponse {
	s.DeliveryChannelId = &v
	return s
}

type PutEvaluationsRequest struct {
	ResultToken *string `json:"ResultToken,omitempty" xml:"ResultToken,omitempty" require:"true"`
	Evaluations *string `json:"Evaluations,omitempty" xml:"Evaluations,omitempty"`
}

func (s PutEvaluationsRequest) String() string {
	return tea.Prettify(s)
}

func (s PutEvaluationsRequest) GoString() string {
	return s.String()
}

func (s *PutEvaluationsRequest) SetResultToken(v string) *PutEvaluationsRequest {
	s.ResultToken = &v
	return s
}

func (s *PutEvaluationsRequest) SetEvaluations(v string) *PutEvaluationsRequest {
	s.Evaluations = &v
	return s
}

type PutEvaluationsResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty" require:"true"`
}

func (s PutEvaluationsResponse) String() string {
	return tea.Prettify(s)
}

func (s PutEvaluationsResponse) GoString() string {
	return s.String()
}

func (s *PutEvaluationsResponse) SetRequestId(v string) *PutEvaluationsResponse {
	s.RequestId = &v
	return s
}

func (s *PutEvaluationsResponse) SetResult(v bool) *PutEvaluationsResponse {
	s.Result = &v
	return s
}

type StartConfigRuleEvaluationRequest struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty" require:"true"`
	MultiAccount *bool   `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	MemberId     *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s StartConfigRuleEvaluationRequest) String() string {
	return tea.Prettify(s)
}

func (s StartConfigRuleEvaluationRequest) GoString() string {
	return s.String()
}

func (s *StartConfigRuleEvaluationRequest) SetConfigRuleId(v string) *StartConfigRuleEvaluationRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *StartConfigRuleEvaluationRequest) SetMultiAccount(v bool) *StartConfigRuleEvaluationRequest {
	s.MultiAccount = &v
	return s
}

func (s *StartConfigRuleEvaluationRequest) SetMemberId(v int64) *StartConfigRuleEvaluationRequest {
	s.MemberId = &v
	return s
}

type StartConfigRuleEvaluationResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty" require:"true"`
}

func (s StartConfigRuleEvaluationResponse) String() string {
	return tea.Prettify(s)
}

func (s StartConfigRuleEvaluationResponse) GoString() string {
	return s.String()
}

func (s *StartConfigRuleEvaluationResponse) SetRequestId(v string) *StartConfigRuleEvaluationResponse {
	s.RequestId = &v
	return s
}

func (s *StartConfigRuleEvaluationResponse) SetResult(v bool) *StartConfigRuleEvaluationResponse {
	s.Result = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("central")
	client.EndpointMap = map[string]*string{
		"cn-shanghai":    tea.String("config.cn-shanghai.aliyuncs.com"),
		"ap-southeast-1": tea.String("config.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("config"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetResourceComplianceTimelineWithOptions(request *GetResourceComplianceTimelineRequest, runtime *util.RuntimeOptions) (_result *GetResourceComplianceTimelineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetResourceComplianceTimelineResponse{}
	_body, _err := client.DoRequest(tea.String("GetResourceComplianceTimeline"), tea.String("HTTPS"), tea.String("GET"), tea.String("2019-01-08"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceComplianceTimeline(request *GetResourceComplianceTimelineRequest) (_result *GetResourceComplianceTimelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceComplianceTimelineResponse{}
	_body, _err := client.GetResourceComplianceTimelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourceConfigurationTimelineWithOptions(request *GetResourceConfigurationTimelineRequest, runtime *util.RuntimeOptions) (_result *GetResourceConfigurationTimelineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetResourceConfigurationTimelineResponse{}
	_body, _err := client.DoRequest(tea.String("GetResourceConfigurationTimeline"), tea.String("HTTPS"), tea.String("GET"), tea.String("2019-01-08"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceConfigurationTimeline(request *GetResourceConfigurationTimelineRequest) (_result *GetResourceConfigurationTimelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceConfigurationTimelineResponse{}
	_body, _err := client.GetResourceConfigurationTimelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeliveryChannelsWithOptions(request *DescribeDeliveryChannelsRequest, runtime *util.RuntimeOptions) (_result *DescribeDeliveryChannelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDeliveryChannelsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDeliveryChannels"), tea.String("HTTPS"), tea.String("GET"), tea.String("2019-01-08"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeliveryChannels(request *DescribeDeliveryChannelsRequest) (_result *DescribeDeliveryChannelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeliveryChannelsResponse{}
	_body, _err := client.DescribeDeliveryChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutConfigurationRecorderWithOptions(request *PutConfigurationRecorderRequest, runtime *util.RuntimeOptions) (_result *PutConfigurationRecorderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PutConfigurationRecorderResponse{}
	_body, _err := client.DoRequest(tea.String("PutConfigurationRecorder"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-01-08"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutConfigurationRecorder(request *PutConfigurationRecorderRequest) (_result *PutConfigurationRecorderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutConfigurationRecorderResponse{}
	_body, _err := client.PutConfigurationRecorderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDiscoveredResourceSummaryWithOptions(request *GetDiscoveredResourceSummaryRequest, runtime *util.RuntimeOptions) (_result *GetDiscoveredResourceSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetDiscoveredResourceSummaryResponse{}
	_body, _err := client.DoRequest(tea.String("GetDiscoveredResourceSummary"), tea.String("HTTPS"), tea.String("GET"), tea.String("2019-01-08"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDiscoveredResourceSummary(request *GetDiscoveredResourceSummaryRequest) (_result *GetDiscoveredResourceSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDiscoveredResourceSummaryResponse{}
	_body, _err := client.GetDiscoveredResourceSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ActiveConfigRulesWithOptions(request *ActiveConfigRulesRequest, runtime *util.RuntimeOptions) (_result *ActiveConfigRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ActiveConfigRulesResponse{}
	_body, _err := client.DoRequest(tea.String("ActiveConfigRules"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-01-08"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActiveConfigRules(request *ActiveConfigRulesRequest) (_result *ActiveConfigRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActiveConfigRulesResponse{}
	_body, _err := client.ActiveConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopConfigRulesWithOptions(request *StopConfigRulesRequest, runtime *util.RuntimeOptions) (_result *StopConfigRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StopConfigRulesResponse{}
	_body, _err := client.DoRequest(tea.String("StopConfigRules"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-01-08"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopConfigRules(request *StopConfigRulesRequest) (_result *StopConfigRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopConfigRulesResponse{}
	_body, _err := client.StopConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeComplianceSummaryWithOptions(request *DescribeComplianceSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeComplianceSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeComplianceSummaryResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeComplianceSummary"), tea.String("HTTPS"), tea.String("GET"), tea.String("2019-01-08"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeComplianceSummary(request *DescribeComplianceSummaryRequest) (_result *DescribeComplianceSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeComplianceSummaryResponse{}
	_body, _err := client.DescribeComplianceSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConfigRulesWithOptions(request *ListConfigRulesRequest, runtime *util.RuntimeOptions) (_result *ListConfigRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListConfigRulesResponse{}
	_body, _err := client.DoRequest(tea.String("ListConfigRules"), tea.String("HTTPS"), tea.String("GET"), tea.String("2019-01-08"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConfigRules(request *ListConfigRulesRequest) (_result *ListConfigRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConfigRulesResponse{}
	_body, _err := client.ListConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutConfigRuleWithOptions(request *PutConfigRuleRequest, runtime *util.RuntimeOptions) (_result *PutConfigRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PutConfigRuleResponse{}
	_body, _err := client.DoRequest(tea.String("PutConfigRule"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-01-08"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutConfigRule(request *PutConfigRuleRequest) (_result *PutConfigRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutConfigRuleResponse{}
	_body, _err := client.PutConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEvaluationResultsWithOptions(request *DescribeEvaluationResultsRequest, runtime *util.RuntimeOptions) (_result *DescribeEvaluationResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeEvaluationResultsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeEvaluationResults"), tea.String("HTTPS"), tea.String("GET"), tea.String("2019-01-08"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEvaluationResults(request *DescribeEvaluationResultsRequest) (_result *DescribeEvaluationResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEvaluationResultsResponse{}
	_body, _err := client.DescribeEvaluationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConfigRulesWithOptions(request *DeleteConfigRulesRequest, runtime *util.RuntimeOptions) (_result *DeleteConfigRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteConfigRulesResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteConfigRules"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-01-08"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConfigRules(request *DeleteConfigRulesRequest) (_result *DeleteConfigRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteConfigRulesResponse{}
	_body, _err := client.DeleteConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeComplianceWithOptions(request *DescribeComplianceRequest, runtime *util.RuntimeOptions) (_result *DescribeComplianceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeComplianceResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeCompliance"), tea.String("HTTPS"), tea.String("GET"), tea.String("2019-01-08"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCompliance(request *DescribeComplianceRequest) (_result *DescribeComplianceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeComplianceResponse{}
	_body, _err := client.DescribeComplianceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDiscoveredResourceCountsWithOptions(request *GetDiscoveredResourceCountsRequest, runtime *util.RuntimeOptions) (_result *GetDiscoveredResourceCountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetDiscoveredResourceCountsResponse{}
	_body, _err := client.DoRequest(tea.String("GetDiscoveredResourceCounts"), tea.String("HTTPS"), tea.String("GET"), tea.String("2019-01-08"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDiscoveredResourceCounts(request *GetDiscoveredResourceCountsRequest) (_result *GetDiscoveredResourceCountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDiscoveredResourceCountsResponse{}
	_body, _err := client.GetDiscoveredResourceCountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDiscoveredResourcesWithOptions(request *ListDiscoveredResourcesRequest, runtime *util.RuntimeOptions) (_result *ListDiscoveredResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListDiscoveredResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("ListDiscoveredResources"), tea.String("HTTPS"), tea.String("GET"), tea.String("2019-01-08"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDiscoveredResources(request *ListDiscoveredResourcesRequest) (_result *ListDiscoveredResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDiscoveredResourcesResponse{}
	_body, _err := client.ListDiscoveredResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConfigurationRecorderWithOptions(request *DescribeConfigurationRecorderRequest, runtime *util.RuntimeOptions) (_result *DescribeConfigurationRecorderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeConfigurationRecorderResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeConfigurationRecorder"), tea.String("HTTPS"), tea.String("GET"), tea.String("2019-01-08"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConfigurationRecorder(request *DescribeConfigurationRecorderRequest) (_result *DescribeConfigurationRecorderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConfigurationRecorderResponse{}
	_body, _err := client.DescribeConfigurationRecorderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiscoveredResourceWithOptions(request *DescribeDiscoveredResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeDiscoveredResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDiscoveredResourceResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDiscoveredResource"), tea.String("HTTPS"), tea.String("GET"), tea.String("2019-01-08"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiscoveredResource(request *DescribeDiscoveredResourceRequest) (_result *DescribeDiscoveredResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiscoveredResourceResponse{}
	_body, _err := client.DescribeDiscoveredResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartConfigurationRecorderWithOptions(request *StartConfigurationRecorderRequest, runtime *util.RuntimeOptions) (_result *StartConfigurationRecorderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StartConfigurationRecorderResponse{}
	_body, _err := client.DoRequest(tea.String("StartConfigurationRecorder"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-01-08"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartConfigurationRecorder(request *StartConfigurationRecorderRequest) (_result *StartConfigurationRecorderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartConfigurationRecorderResponse{}
	_body, _err := client.StartConfigurationRecorderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConfigRuleWithOptions(request *DescribeConfigRuleRequest, runtime *util.RuntimeOptions) (_result *DescribeConfigRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeConfigRuleResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeConfigRule"), tea.String("HTTPS"), tea.String("GET"), tea.String("2019-01-08"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConfigRule(request *DescribeConfigRuleRequest) (_result *DescribeConfigRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConfigRuleResponse{}
	_body, _err := client.DescribeConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSupportedResourceTypesWithOptions(request *GetSupportedResourceTypesRequest, runtime *util.RuntimeOptions) (_result *GetSupportedResourceTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetSupportedResourceTypesResponse{}
	_body, _err := client.DoRequest(tea.String("GetSupportedResourceTypes"), tea.String("HTTPS"), tea.String("GET"), tea.String("2019-01-08"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSupportedResourceTypes(request *GetSupportedResourceTypesRequest) (_result *GetSupportedResourceTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSupportedResourceTypesResponse{}
	_body, _err := client.GetSupportedResourceTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutDeliveryChannelWithOptions(request *PutDeliveryChannelRequest, runtime *util.RuntimeOptions) (_result *PutDeliveryChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PutDeliveryChannelResponse{}
	_body, _err := client.DoRequest(tea.String("PutDeliveryChannel"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-01-08"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutDeliveryChannel(request *PutDeliveryChannelRequest) (_result *PutDeliveryChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutDeliveryChannelResponse{}
	_body, _err := client.PutDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutEvaluationsWithOptions(request *PutEvaluationsRequest, runtime *util.RuntimeOptions) (_result *PutEvaluationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PutEvaluationsResponse{}
	_body, _err := client.DoRequest(tea.String("PutEvaluations"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-01-08"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutEvaluations(request *PutEvaluationsRequest) (_result *PutEvaluationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutEvaluationsResponse{}
	_body, _err := client.PutEvaluationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartConfigRuleEvaluationWithOptions(request *StartConfigRuleEvaluationRequest, runtime *util.RuntimeOptions) (_result *StartConfigRuleEvaluationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StartConfigRuleEvaluationResponse{}
	_body, _err := client.DoRequest(tea.String("StartConfigRuleEvaluation"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-01-08"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartConfigRuleEvaluation(request *StartConfigRuleEvaluationRequest) (_result *StartConfigRuleEvaluationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartConfigRuleEvaluationResponse{}
	_body, _err := client.StartConfigRuleEvaluationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
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
