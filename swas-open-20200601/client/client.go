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

type AllocatePublicConnectionRequest struct {
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AllocatePublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocatePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocatePublicConnectionRequest) SetClientToken(v string) *AllocatePublicConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *AllocatePublicConnectionRequest) SetDatabaseInstanceId(v string) *AllocatePublicConnectionRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *AllocatePublicConnectionRequest) SetRegionId(v string) *AllocatePublicConnectionRequest {
	s.RegionId = &v
	return s
}

type AllocatePublicConnectionResponseBody struct {
	PublicConnection *string `json:"PublicConnection,omitempty" xml:"PublicConnection,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocatePublicConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocatePublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocatePublicConnectionResponseBody) SetPublicConnection(v string) *AllocatePublicConnectionResponseBody {
	s.PublicConnection = &v
	return s
}

func (s *AllocatePublicConnectionResponseBody) SetRequestId(v string) *AllocatePublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

type AllocatePublicConnectionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AllocatePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AllocatePublicConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocatePublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *AllocatePublicConnectionResponse) SetHeaders(v map[string]*string) *AllocatePublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *AllocatePublicConnectionResponse) SetStatusCode(v int32) *AllocatePublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocatePublicConnectionResponse) SetBody(v *AllocatePublicConnectionResponseBody) *AllocatePublicConnectionResponse {
	s.Body = v
	return s
}

type CreateCustomImageRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataSnapshotId   *string `json:"DataSnapshotId,omitempty" xml:"DataSnapshotId,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageName        *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SystemSnapshotId *string `json:"SystemSnapshotId,omitempty" xml:"SystemSnapshotId,omitempty"`
}

func (s CreateCustomImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomImageRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomImageRequest) SetClientToken(v string) *CreateCustomImageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCustomImageRequest) SetDataSnapshotId(v string) *CreateCustomImageRequest {
	s.DataSnapshotId = &v
	return s
}

func (s *CreateCustomImageRequest) SetDescription(v string) *CreateCustomImageRequest {
	s.Description = &v
	return s
}

func (s *CreateCustomImageRequest) SetImageName(v string) *CreateCustomImageRequest {
	s.ImageName = &v
	return s
}

func (s *CreateCustomImageRequest) SetInstanceId(v string) *CreateCustomImageRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCustomImageRequest) SetRegionId(v string) *CreateCustomImageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCustomImageRequest) SetSystemSnapshotId(v string) *CreateCustomImageRequest {
	s.SystemSnapshotId = &v
	return s
}

type CreateCustomImageResponseBody struct {
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomImageResponseBody) SetImageId(v string) *CreateCustomImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *CreateCustomImageResponseBody) SetRequestId(v string) *CreateCustomImageResponseBody {
	s.RequestId = &v
	return s
}

type CreateCustomImageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCustomImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCustomImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomImageResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomImageResponse) SetHeaders(v map[string]*string) *CreateCustomImageResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomImageResponse) SetStatusCode(v int32) *CreateCustomImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomImageResponse) SetBody(v *CreateCustomImageResponseBody) *CreateCustomImageResponse {
	s.Body = v
	return s
}

type CreateFirewallRuleRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
}

func (s CreateFirewallRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateFirewallRuleRequest) SetClientToken(v string) *CreateFirewallRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFirewallRuleRequest) SetInstanceId(v string) *CreateFirewallRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateFirewallRuleRequest) SetPort(v string) *CreateFirewallRuleRequest {
	s.Port = &v
	return s
}

func (s *CreateFirewallRuleRequest) SetRegionId(v string) *CreateFirewallRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFirewallRuleRequest) SetRemark(v string) *CreateFirewallRuleRequest {
	s.Remark = &v
	return s
}

func (s *CreateFirewallRuleRequest) SetRuleProtocol(v string) *CreateFirewallRuleRequest {
	s.RuleProtocol = &v
	return s
}

type CreateFirewallRuleResponseBody struct {
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFirewallRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFirewallRuleResponseBody) SetFirewallId(v string) *CreateFirewallRuleResponseBody {
	s.FirewallId = &v
	return s
}

func (s *CreateFirewallRuleResponseBody) SetRequestId(v string) *CreateFirewallRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateFirewallRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFirewallRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFirewallRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateFirewallRuleResponse) SetHeaders(v map[string]*string) *CreateFirewallRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateFirewallRuleResponse) SetStatusCode(v int32) *CreateFirewallRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFirewallRuleResponse) SetBody(v *CreateFirewallRuleResponseBody) *CreateFirewallRuleResponse {
	s.Body = v
	return s
}

type CreateInstancesRequest struct {
	Amount          *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AutoRenew       *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod *int32  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	ChargeType      *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataDiskSize    *int64  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	ImageId         *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Period          *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PlanId          *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstancesRequest) GoString() string {
	return s.String()
}

func (s *CreateInstancesRequest) SetAmount(v int32) *CreateInstancesRequest {
	s.Amount = &v
	return s
}

func (s *CreateInstancesRequest) SetAutoRenew(v bool) *CreateInstancesRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstancesRequest) SetAutoRenewPeriod(v int32) *CreateInstancesRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateInstancesRequest) SetChargeType(v string) *CreateInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstancesRequest) SetClientToken(v string) *CreateInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstancesRequest) SetDataDiskSize(v int64) *CreateInstancesRequest {
	s.DataDiskSize = &v
	return s
}

func (s *CreateInstancesRequest) SetImageId(v string) *CreateInstancesRequest {
	s.ImageId = &v
	return s
}

func (s *CreateInstancesRequest) SetPeriod(v int32) *CreateInstancesRequest {
	s.Period = &v
	return s
}

func (s *CreateInstancesRequest) SetPlanId(v string) *CreateInstancesRequest {
	s.PlanId = &v
	return s
}

func (s *CreateInstancesRequest) SetRegionId(v string) *CreateInstancesRequest {
	s.RegionId = &v
	return s
}

type CreateInstancesResponseBody struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstancesResponseBody) SetInstanceIds(v []*string) *CreateInstancesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *CreateInstancesResponseBody) SetRequestId(v string) *CreateInstancesResponseBody {
	s.RequestId = &v
	return s
}

type CreateInstancesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstancesResponse) GoString() string {
	return s.String()
}

func (s *CreateInstancesResponse) SetHeaders(v map[string]*string) *CreateInstancesResponse {
	s.Headers = v
	return s
}

func (s *CreateInstancesResponse) SetStatusCode(v int32) *CreateInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstancesResponse) SetBody(v *CreateInstancesResponseBody) *CreateInstancesResponse {
	s.Body = v
	return s
}

type CreateSnapshotRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DiskId       *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
}

func (s CreateSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) SetClientToken(v string) *CreateSnapshotRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSnapshotRequest) SetDiskId(v string) *CreateSnapshotRequest {
	s.DiskId = &v
	return s
}

func (s *CreateSnapshotRequest) SetRegionId(v string) *CreateSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSnapshotRequest) SetSnapshotName(v string) *CreateSnapshotRequest {
	s.SnapshotName = &v
	return s
}

type CreateSnapshotResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponseBody) SetRequestId(v string) *CreateSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSnapshotResponseBody) SetSnapshotId(v string) *CreateSnapshotResponseBody {
	s.SnapshotId = &v
	return s
}

type CreateSnapshotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponse) SetHeaders(v map[string]*string) *CreateSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CreateSnapshotResponse) SetStatusCode(v int32) *CreateSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSnapshotResponse) SetBody(v *CreateSnapshotResponseBody) *CreateSnapshotResponse {
	s.Body = v
	return s
}

type DeleteCustomImageRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCustomImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomImageRequest) SetClientToken(v string) *DeleteCustomImageRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCustomImageRequest) SetImageId(v string) *DeleteCustomImageRequest {
	s.ImageId = &v
	return s
}

func (s *DeleteCustomImageRequest) SetRegionId(v string) *DeleteCustomImageRequest {
	s.RegionId = &v
	return s
}

type DeleteCustomImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomImageResponseBody) SetRequestId(v string) *DeleteCustomImageResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCustomImageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCustomImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCustomImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomImageResponse) SetHeaders(v map[string]*string) *DeleteCustomImageResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomImageResponse) SetStatusCode(v int32) *DeleteCustomImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomImageResponse) SetBody(v *DeleteCustomImageResponseBody) *DeleteCustomImageResponse {
	s.Body = v
	return s
}

type DeleteFirewallRuleRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleId      *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteFirewallRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFirewallRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteFirewallRuleRequest) SetClientToken(v string) *DeleteFirewallRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFirewallRuleRequest) SetInstanceId(v string) *DeleteFirewallRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteFirewallRuleRequest) SetRegionId(v string) *DeleteFirewallRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFirewallRuleRequest) SetRuleId(v string) *DeleteFirewallRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteFirewallRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFirewallRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFirewallRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFirewallRuleResponseBody) SetRequestId(v string) *DeleteFirewallRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFirewallRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFirewallRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFirewallRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFirewallRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteFirewallRuleResponse) SetHeaders(v map[string]*string) *DeleteFirewallRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteFirewallRuleResponse) SetStatusCode(v int32) *DeleteFirewallRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFirewallRuleResponse) SetBody(v *DeleteFirewallRuleResponseBody) *DeleteFirewallRuleResponse {
	s.Body = v
	return s
}

type DeleteSnapshotRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotId  *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DeleteSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRequest) SetClientToken(v string) *DeleteSnapshotRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteSnapshotRequest) SetRegionId(v string) *DeleteSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSnapshotRequest) SetSnapshotId(v string) *DeleteSnapshotRequest {
	s.SnapshotId = &v
	return s
}

type DeleteSnapshotResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponseBody) SetRequestId(v string) *DeleteSnapshotResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSnapshotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponse) SetHeaders(v map[string]*string) *DeleteSnapshotResponse {
	s.Headers = v
	return s
}

func (s *DeleteSnapshotResponse) SetStatusCode(v int32) *DeleteSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSnapshotResponse) SetBody(v *DeleteSnapshotResponseBody) *DeleteSnapshotResponse {
	s.Body = v
	return s
}

type DescribeCloudAssistantStatusRequest struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	PageNumber  *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCloudAssistantStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudAssistantStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantStatusRequest) SetInstanceIds(v []*string) *DescribeCloudAssistantStatusRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) SetPageNumber(v int32) *DescribeCloudAssistantStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) SetPageSize(v int32) *DescribeCloudAssistantStatusRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) SetRegionId(v string) *DescribeCloudAssistantStatusRequest {
	s.RegionId = &v
	return s
}

type DescribeCloudAssistantStatusShrinkRequest struct {
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	PageNumber        *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCloudAssistantStatusShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudAssistantStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantStatusShrinkRequest) SetInstanceIdsShrink(v string) *DescribeCloudAssistantStatusShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *DescribeCloudAssistantStatusShrinkRequest) SetPageNumber(v int32) *DescribeCloudAssistantStatusShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudAssistantStatusShrinkRequest) SetPageSize(v int32) *DescribeCloudAssistantStatusShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudAssistantStatusShrinkRequest) SetRegionId(v string) *DescribeCloudAssistantStatusShrinkRequest {
	s.RegionId = &v
	return s
}

type DescribeCloudAssistantStatusResponseBody struct {
	CloudAssistantStatus []*DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus `json:"CloudAssistantStatus,omitempty" xml:"CloudAssistantStatus,omitempty" type:"Repeated"`
	PageNumber           *int32                                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId            *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount           *int32                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudAssistantStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudAssistantStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantStatusResponseBody) SetCloudAssistantStatus(v []*DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus) *DescribeCloudAssistantStatusResponseBody {
	s.CloudAssistantStatus = v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBody) SetPageNumber(v int32) *DescribeCloudAssistantStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBody) SetPageSize(v int32) *DescribeCloudAssistantStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBody) SetRequestId(v string) *DescribeCloudAssistantStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBody) SetTotalCount(v int32) *DescribeCloudAssistantStatusResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Status     *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus) SetInstanceId(v string) *DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus) SetStatus(v bool) *DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus {
	s.Status = &v
	return s
}

type DescribeCloudAssistantStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCloudAssistantStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCloudAssistantStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudAssistantStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantStatusResponse) SetHeaders(v map[string]*string) *DescribeCloudAssistantStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudAssistantStatusResponse) SetStatusCode(v int32) *DescribeCloudAssistantStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponse) SetBody(v *DescribeCloudAssistantStatusResponseBody) *DescribeCloudAssistantStatusResponse {
	s.Body = v
	return s
}

type DescribeDatabaseErrorLogsRequest struct {
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	EndTime            *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber         *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime          *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDatabaseErrorLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseErrorLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseErrorLogsRequest) SetDatabaseInstanceId(v string) *DescribeDatabaseErrorLogsRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *DescribeDatabaseErrorLogsRequest) SetEndTime(v string) *DescribeDatabaseErrorLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDatabaseErrorLogsRequest) SetPageNumber(v int32) *DescribeDatabaseErrorLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabaseErrorLogsRequest) SetPageSize(v int32) *DescribeDatabaseErrorLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabaseErrorLogsRequest) SetRegionId(v string) *DescribeDatabaseErrorLogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDatabaseErrorLogsRequest) SetStartTime(v string) *DescribeDatabaseErrorLogsRequest {
	s.StartTime = &v
	return s
}

type DescribeDatabaseErrorLogsResponseBody struct {
	ErrorLogs  []*DescribeDatabaseErrorLogsResponseBodyErrorLogs `json:"ErrorLogs,omitempty" xml:"ErrorLogs,omitempty" type:"Repeated"`
	PageNumber *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDatabaseErrorLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseErrorLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseErrorLogsResponseBody) SetErrorLogs(v []*DescribeDatabaseErrorLogsResponseBodyErrorLogs) *DescribeDatabaseErrorLogsResponseBody {
	s.ErrorLogs = v
	return s
}

func (s *DescribeDatabaseErrorLogsResponseBody) SetPageNumber(v int32) *DescribeDatabaseErrorLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabaseErrorLogsResponseBody) SetPageSize(v int32) *DescribeDatabaseErrorLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabaseErrorLogsResponseBody) SetRequestId(v string) *DescribeDatabaseErrorLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabaseErrorLogsResponseBody) SetTotalCount(v int32) *DescribeDatabaseErrorLogsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDatabaseErrorLogsResponseBodyErrorLogs struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ErrorInfo  *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
}

func (s DescribeDatabaseErrorLogsResponseBodyErrorLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseErrorLogsResponseBodyErrorLogs) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseErrorLogsResponseBodyErrorLogs) SetCreateTime(v string) *DescribeDatabaseErrorLogsResponseBodyErrorLogs {
	s.CreateTime = &v
	return s
}

func (s *DescribeDatabaseErrorLogsResponseBodyErrorLogs) SetErrorInfo(v string) *DescribeDatabaseErrorLogsResponseBodyErrorLogs {
	s.ErrorInfo = &v
	return s
}

type DescribeDatabaseErrorLogsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDatabaseErrorLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDatabaseErrorLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseErrorLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseErrorLogsResponse) SetHeaders(v map[string]*string) *DescribeDatabaseErrorLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabaseErrorLogsResponse) SetStatusCode(v int32) *DescribeDatabaseErrorLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatabaseErrorLogsResponse) SetBody(v *DescribeDatabaseErrorLogsResponseBody) *DescribeDatabaseErrorLogsResponse {
	s.Body = v
	return s
}

type DescribeDatabaseInstanceMetricDataRequest struct {
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	EndTime            *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MetricName         *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime          *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDatabaseInstanceMetricDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstanceMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceMetricDataRequest) SetDatabaseInstanceId(v string) *DescribeDatabaseInstanceMetricDataRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataRequest) SetEndTime(v string) *DescribeDatabaseInstanceMetricDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataRequest) SetMetricName(v string) *DescribeDatabaseInstanceMetricDataRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataRequest) SetRegionId(v string) *DescribeDatabaseInstanceMetricDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataRequest) SetStartTime(v string) *DescribeDatabaseInstanceMetricDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDatabaseInstanceMetricDataResponseBody struct {
	DataFormat *string `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	MetricData *string `json:"MetricData,omitempty" xml:"MetricData,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Unit       *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeDatabaseInstanceMetricDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstanceMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) SetDataFormat(v string) *DescribeDatabaseInstanceMetricDataResponseBody {
	s.DataFormat = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) SetMetricData(v string) *DescribeDatabaseInstanceMetricDataResponseBody {
	s.MetricData = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) SetMetricName(v string) *DescribeDatabaseInstanceMetricDataResponseBody {
	s.MetricName = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) SetRequestId(v string) *DescribeDatabaseInstanceMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) SetUnit(v string) *DescribeDatabaseInstanceMetricDataResponseBody {
	s.Unit = &v
	return s
}

type DescribeDatabaseInstanceMetricDataResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDatabaseInstanceMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDatabaseInstanceMetricDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstanceMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceMetricDataResponse) SetHeaders(v map[string]*string) *DescribeDatabaseInstanceMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataResponse) SetStatusCode(v int32) *DescribeDatabaseInstanceMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataResponse) SetBody(v *DescribeDatabaseInstanceMetricDataResponseBody) *DescribeDatabaseInstanceMetricDataResponse {
	s.Body = v
	return s
}

type DescribeDatabaseInstanceParametersRequest struct {
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDatabaseInstanceParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstanceParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceParametersRequest) SetDatabaseInstanceId(v string) *DescribeDatabaseInstanceParametersRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersRequest) SetRegionId(v string) *DescribeDatabaseInstanceParametersRequest {
	s.RegionId = &v
	return s
}

type DescribeDatabaseInstanceParametersResponseBody struct {
	ConfigParameters  []*DescribeDatabaseInstanceParametersResponseBodyConfigParameters  `json:"ConfigParameters,omitempty" xml:"ConfigParameters,omitempty" type:"Repeated"`
	Engine            *string                                                            `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion     *string                                                            `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	RequestId         *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RunningParameters []*DescribeDatabaseInstanceParametersResponseBodyRunningParameters `json:"RunningParameters,omitempty" xml:"RunningParameters,omitempty" type:"Repeated"`
}

func (s DescribeDatabaseInstanceParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstanceParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceParametersResponseBody) SetConfigParameters(v []*DescribeDatabaseInstanceParametersResponseBodyConfigParameters) *DescribeDatabaseInstanceParametersResponseBody {
	s.ConfigParameters = v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBody) SetEngine(v string) *DescribeDatabaseInstanceParametersResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBody) SetEngineVersion(v string) *DescribeDatabaseInstanceParametersResponseBody {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBody) SetRequestId(v string) *DescribeDatabaseInstanceParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBody) SetRunningParameters(v []*DescribeDatabaseInstanceParametersResponseBodyRunningParameters) *DescribeDatabaseInstanceParametersResponseBody {
	s.RunningParameters = v
	return s
}

type DescribeDatabaseInstanceParametersResponseBodyConfigParameters struct {
	CheckingCode         *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	ForceModify          *string `json:"ForceModify,omitempty" xml:"ForceModify,omitempty"`
	ForceRestart         *string `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue       *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeDatabaseInstanceParametersResponseBodyConfigParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstanceParametersResponseBodyConfigParameters) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) SetCheckingCode(v string) *DescribeDatabaseInstanceParametersResponseBodyConfigParameters {
	s.CheckingCode = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) SetForceModify(v string) *DescribeDatabaseInstanceParametersResponseBodyConfigParameters {
	s.ForceModify = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) SetForceRestart(v string) *DescribeDatabaseInstanceParametersResponseBodyConfigParameters {
	s.ForceRestart = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) SetParameterDescription(v string) *DescribeDatabaseInstanceParametersResponseBodyConfigParameters {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) SetParameterName(v string) *DescribeDatabaseInstanceParametersResponseBodyConfigParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) SetParameterValue(v string) *DescribeDatabaseInstanceParametersResponseBodyConfigParameters {
	s.ParameterValue = &v
	return s
}

type DescribeDatabaseInstanceParametersResponseBodyRunningParameters struct {
	CheckingCode         *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	ForceModify          *string `json:"ForceModify,omitempty" xml:"ForceModify,omitempty"`
	ForceRestart         *string `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue       *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeDatabaseInstanceParametersResponseBodyRunningParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstanceParametersResponseBodyRunningParameters) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) SetCheckingCode(v string) *DescribeDatabaseInstanceParametersResponseBodyRunningParameters {
	s.CheckingCode = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) SetForceModify(v string) *DescribeDatabaseInstanceParametersResponseBodyRunningParameters {
	s.ForceModify = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) SetForceRestart(v string) *DescribeDatabaseInstanceParametersResponseBodyRunningParameters {
	s.ForceRestart = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) SetParameterDescription(v string) *DescribeDatabaseInstanceParametersResponseBodyRunningParameters {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) SetParameterName(v string) *DescribeDatabaseInstanceParametersResponseBodyRunningParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) SetParameterValue(v string) *DescribeDatabaseInstanceParametersResponseBodyRunningParameters {
	s.ParameterValue = &v
	return s
}

type DescribeDatabaseInstanceParametersResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDatabaseInstanceParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDatabaseInstanceParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstanceParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceParametersResponse) SetHeaders(v map[string]*string) *DescribeDatabaseInstanceParametersResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponse) SetStatusCode(v int32) *DescribeDatabaseInstanceParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponse) SetBody(v *DescribeDatabaseInstanceParametersResponseBody) *DescribeDatabaseInstanceParametersResponse {
	s.Body = v
	return s
}

type DescribeDatabaseInstancesRequest struct {
	DatabaseInstanceIds *string `json:"DatabaseInstanceIds,omitempty" xml:"DatabaseInstanceIds,omitempty"`
	PageNumber          *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDatabaseInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstancesRequest) SetDatabaseInstanceIds(v string) *DescribeDatabaseInstancesRequest {
	s.DatabaseInstanceIds = &v
	return s
}

func (s *DescribeDatabaseInstancesRequest) SetPageNumber(v int32) *DescribeDatabaseInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabaseInstancesRequest) SetPageSize(v int32) *DescribeDatabaseInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabaseInstancesRequest) SetRegionId(v string) *DescribeDatabaseInstancesRequest {
	s.RegionId = &v
	return s
}

type DescribeDatabaseInstancesResponseBody struct {
	DatabaseInstances []*DescribeDatabaseInstancesResponseBodyDatabaseInstances `json:"DatabaseInstances,omitempty" xml:"DatabaseInstances,omitempty" type:"Repeated"`
	PageNumber        *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId         *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount        *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDatabaseInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstancesResponseBody) SetDatabaseInstances(v []*DescribeDatabaseInstancesResponseBodyDatabaseInstances) *DescribeDatabaseInstancesResponseBody {
	s.DatabaseInstances = v
	return s
}

func (s *DescribeDatabaseInstancesResponseBody) SetPageNumber(v int32) *DescribeDatabaseInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBody) SetPageSize(v int32) *DescribeDatabaseInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBody) SetRequestId(v string) *DescribeDatabaseInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBody) SetTotalCount(v int32) *DescribeDatabaseInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDatabaseInstancesResponseBodyDatabaseInstances struct {
	BusinessStatus          *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	ChargeType              *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Cpu                     *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreationTime            *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DatabaseInstanceEdition *string `json:"DatabaseInstanceEdition,omitempty" xml:"DatabaseInstanceEdition,omitempty"`
	DatabaseInstanceId      *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	DatabaseInstanceName    *string `json:"DatabaseInstanceName,omitempty" xml:"DatabaseInstanceName,omitempty"`
	DatabaseInstanceStatus  *string `json:"DatabaseInstanceStatus,omitempty" xml:"DatabaseInstanceStatus,omitempty"`
	DatabaseVersion         *string `json:"DatabaseVersion,omitempty" xml:"DatabaseVersion,omitempty"`
	ExpiredTime             *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Memory                  *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	PrivateConnection       *string `json:"PrivateConnection,omitempty" xml:"PrivateConnection,omitempty"`
	PublicConnection        *string `json:"PublicConnection,omitempty" xml:"PublicConnection,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Storage                 *int32  `json:"Storage,omitempty" xml:"Storage,omitempty"`
	SuperAccountName        *string `json:"SuperAccountName,omitempty" xml:"SuperAccountName,omitempty"`
}

func (s DescribeDatabaseInstancesResponseBodyDatabaseInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstancesResponseBodyDatabaseInstances) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetBusinessStatus(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetChargeType(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.ChargeType = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetCpu(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.Cpu = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetCreationTime(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.CreationTime = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetDatabaseInstanceEdition(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.DatabaseInstanceEdition = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetDatabaseInstanceId(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.DatabaseInstanceId = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetDatabaseInstanceName(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.DatabaseInstanceName = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetDatabaseInstanceStatus(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.DatabaseInstanceStatus = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetDatabaseVersion(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.DatabaseVersion = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetExpiredTime(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetMemory(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.Memory = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetPrivateConnection(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.PrivateConnection = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetPublicConnection(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.PublicConnection = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetRegionId(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetStorage(v int32) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.Storage = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetSuperAccountName(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.SuperAccountName = &v
	return s
}

type DescribeDatabaseInstancesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDatabaseInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDatabaseInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstancesResponse) SetHeaders(v map[string]*string) *DescribeDatabaseInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabaseInstancesResponse) SetStatusCode(v int32) *DescribeDatabaseInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatabaseInstancesResponse) SetBody(v *DescribeDatabaseInstancesResponseBody) *DescribeDatabaseInstancesResponse {
	s.Body = v
	return s
}

type DescribeDatabaseSlowLogRecordsRequest struct {
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	EndTime            *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber         *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime          *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDatabaseSlowLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseSlowLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseSlowLogRecordsRequest) SetDatabaseInstanceId(v string) *DescribeDatabaseSlowLogRecordsRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsRequest) SetEndTime(v string) *DescribeDatabaseSlowLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsRequest) SetPageNumber(v int32) *DescribeDatabaseSlowLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsRequest) SetPageSize(v int32) *DescribeDatabaseSlowLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsRequest) SetRegionId(v string) *DescribeDatabaseSlowLogRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsRequest) SetStartTime(v string) *DescribeDatabaseSlowLogRecordsRequest {
	s.StartTime = &v
	return s
}

type DescribeDatabaseSlowLogRecordsResponseBody struct {
	Engine         *string                                               `json:"Engine,omitempty" xml:"Engine,omitempty"`
	PageNumber     *int32                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PhysicalIORead *int64                                                `json:"PhysicalIORead,omitempty" xml:"PhysicalIORead,omitempty"`
	RequestId      *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlowLogs       []*DescribeDatabaseSlowLogRecordsResponseBodySlowLogs `json:"SlowLogs,omitempty" xml:"SlowLogs,omitempty" type:"Repeated"`
	TotalCount     *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDatabaseSlowLogRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseSlowLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetEngine(v string) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetPageNumber(v int32) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetPageSize(v int32) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetPhysicalIORead(v int64) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.PhysicalIORead = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetRequestId(v string) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetSlowLogs(v []*DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.SlowLogs = v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetTotalCount(v int32) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDatabaseSlowLogRecordsResponseBodySlowLogs struct {
	DBName             *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" xml:"ExecutionStartTime,omitempty"`
	HostAddress        *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	LockTimes          *int64  `json:"LockTimes,omitempty" xml:"LockTimes,omitempty"`
	ParseRowCounts     *int64  `json:"ParseRowCounts,omitempty" xml:"ParseRowCounts,omitempty"`
	QueryTimeMS        *int64  `json:"QueryTimeMS,omitempty" xml:"QueryTimeMS,omitempty"`
	QueryTimes         *int64  `json:"QueryTimes,omitempty" xml:"QueryTimes,omitempty"`
	ReturnRowCounts    *int64  `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	SQLText            *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
}

func (s DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetDBName(v string) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.DBName = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetExecutionStartTime(v string) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.ExecutionStartTime = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetHostAddress(v string) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.HostAddress = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetLockTimes(v int64) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.LockTimes = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetParseRowCounts(v int64) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.ParseRowCounts = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetQueryTimeMS(v int64) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.QueryTimeMS = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetQueryTimes(v int64) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.QueryTimes = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetReturnRowCounts(v int64) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetSQLText(v string) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.SQLText = &v
	return s
}

type DescribeDatabaseSlowLogRecordsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDatabaseSlowLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDatabaseSlowLogRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseSlowLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseSlowLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeDatabaseSlowLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponse) SetStatusCode(v int32) *DescribeDatabaseSlowLogRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponse) SetBody(v *DescribeDatabaseSlowLogRecordsResponseBody) *DescribeDatabaseSlowLogRecordsResponse {
	s.Body = v
	return s
}

type DescribeInvocationResultRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InvokeId   *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInvocationResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultRequest) SetInstanceId(v string) *DescribeInvocationResultRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInvocationResultRequest) SetInvokeId(v string) *DescribeInvocationResultRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationResultRequest) SetRegionId(v string) *DescribeInvocationResultRequest {
	s.RegionId = &v
	return s
}

type DescribeInvocationResultResponseBody struct {
	InvocationResult *DescribeInvocationResultResponseBodyInvocationResult `json:"InvocationResult,omitempty" xml:"InvocationResult,omitempty" type:"Struct"`
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInvocationResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultResponseBody) SetInvocationResult(v *DescribeInvocationResultResponseBodyInvocationResult) *DescribeInvocationResultResponseBody {
	s.InvocationResult = v
	return s
}

func (s *DescribeInvocationResultResponseBody) SetRequestId(v string) *DescribeInvocationResultResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInvocationResultResponseBodyInvocationResult struct {
	ErrorCode          *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorInfo          *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	ExitCode           *int64  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	FinishedTime       *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InvocationStatus   *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	InvokeId           *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	InvokeRecordStatus *string `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	InvokeUser         *string `json:"InvokeUser,omitempty" xml:"InvokeUser,omitempty"`
	Output             *string `json:"Output,omitempty" xml:"Output,omitempty"`
	StartTime          *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInvocationResultResponseBodyInvocationResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationResultResponseBodyInvocationResult) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetErrorCode(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.ErrorCode = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetErrorInfo(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetExitCode(v int64) *DescribeInvocationResultResponseBodyInvocationResult {
	s.ExitCode = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetFinishedTime(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.FinishedTime = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetInstanceId(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetInvocationStatus(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetInvokeId(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetInvokeRecordStatus(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.InvokeRecordStatus = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetInvokeUser(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.InvokeUser = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetOutput(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.Output = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetStartTime(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.StartTime = &v
	return s
}

type DescribeInvocationResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInvocationResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInvocationResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultResponse) SetHeaders(v map[string]*string) *DescribeInvocationResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvocationResultResponse) SetStatusCode(v int32) *DescribeInvocationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInvocationResultResponse) SetBody(v *DescribeInvocationResultResponseBody) *DescribeInvocationResultResponse {
	s.Body = v
	return s
}

type DescribeInvocationsRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InvokeStatus *string `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInvocationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsRequest) SetInstanceId(v string) *DescribeInvocationsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInvokeStatus(v string) *DescribeInvocationsRequest {
	s.InvokeStatus = &v
	return s
}

func (s *DescribeInvocationsRequest) SetPageNumber(v int32) *DescribeInvocationsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInvocationsRequest) SetPageSize(v int32) *DescribeInvocationsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInvocationsRequest) SetRegionId(v string) *DescribeInvocationsRequest {
	s.RegionId = &v
	return s
}

type DescribeInvocationsResponseBody struct {
	Invocations []*DescribeInvocationsResponseBodyInvocations `json:"Invocations,omitempty" xml:"Invocations,omitempty" type:"Repeated"`
	PageNumber  *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInvocationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBody) SetInvocations(v []*DescribeInvocationsResponseBodyInvocations) *DescribeInvocationsResponseBody {
	s.Invocations = v
	return s
}

func (s *DescribeInvocationsResponseBody) SetPageNumber(v int32) *DescribeInvocationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInvocationsResponseBody) SetPageSize(v int32) *DescribeInvocationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInvocationsResponseBody) SetRequestId(v string) *DescribeInvocationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInvocationsResponseBody) SetTotalCount(v int32) *DescribeInvocationsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeInvocationsResponseBodyInvocations struct {
	CommandContent   *string                `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	CommandName      *string                `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	CommandType      *string                `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	CreationTime     *string                `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	InvocationStatus *string                `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	InvokeId         *string                `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	InvokeStatus     *string                `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	Parameters       map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s DescribeInvocationsResponseBodyInvocations) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocations) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocations) SetCommandContent(v string) *DescribeInvocationsResponseBodyInvocations {
	s.CommandContent = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetCommandName(v string) *DescribeInvocationsResponseBodyInvocations {
	s.CommandName = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetCommandType(v string) *DescribeInvocationsResponseBodyInvocations {
	s.CommandType = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetCreationTime(v string) *DescribeInvocationsResponseBodyInvocations {
	s.CreationTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvocationStatus(v string) *DescribeInvocationsResponseBodyInvocations {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvokeId(v string) *DescribeInvocationsResponseBodyInvocations {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvokeStatus(v string) *DescribeInvocationsResponseBodyInvocations {
	s.InvokeStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetParameters(v map[string]interface{}) *DescribeInvocationsResponseBodyInvocations {
	s.Parameters = v
	return s
}

type DescribeInvocationsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInvocationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInvocationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponse) SetHeaders(v map[string]*string) *DescribeInvocationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvocationsResponse) SetStatusCode(v int32) *DescribeInvocationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInvocationsResponse) SetBody(v *DescribeInvocationsResponseBody) *DescribeInvocationsResponse {
	s.Body = v
	return s
}

type InstallCloudAssistantRequest struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InstallCloudAssistantRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallCloudAssistantRequest) GoString() string {
	return s.String()
}

func (s *InstallCloudAssistantRequest) SetInstanceIds(v []*string) *InstallCloudAssistantRequest {
	s.InstanceIds = v
	return s
}

func (s *InstallCloudAssistantRequest) SetRegionId(v string) *InstallCloudAssistantRequest {
	s.RegionId = &v
	return s
}

type InstallCloudAssistantShrinkRequest struct {
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InstallCloudAssistantShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallCloudAssistantShrinkRequest) GoString() string {
	return s.String()
}

func (s *InstallCloudAssistantShrinkRequest) SetInstanceIdsShrink(v string) *InstallCloudAssistantShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *InstallCloudAssistantShrinkRequest) SetRegionId(v string) *InstallCloudAssistantShrinkRequest {
	s.RegionId = &v
	return s
}

type InstallCloudAssistantResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallCloudAssistantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallCloudAssistantResponseBody) GoString() string {
	return s.String()
}

func (s *InstallCloudAssistantResponseBody) SetRequestId(v string) *InstallCloudAssistantResponseBody {
	s.RequestId = &v
	return s
}

type InstallCloudAssistantResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InstallCloudAssistantResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InstallCloudAssistantResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallCloudAssistantResponse) GoString() string {
	return s.String()
}

func (s *InstallCloudAssistantResponse) SetHeaders(v map[string]*string) *InstallCloudAssistantResponse {
	s.Headers = v
	return s
}

func (s *InstallCloudAssistantResponse) SetStatusCode(v int32) *InstallCloudAssistantResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallCloudAssistantResponse) SetBody(v *InstallCloudAssistantResponseBody) *InstallCloudAssistantResponse {
	s.Body = v
	return s
}

type ListDisksRequest struct {
	DiskIds    *string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListDisksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDisksRequest) GoString() string {
	return s.String()
}

func (s *ListDisksRequest) SetDiskIds(v string) *ListDisksRequest {
	s.DiskIds = &v
	return s
}

func (s *ListDisksRequest) SetInstanceId(v string) *ListDisksRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDisksRequest) SetPageNumber(v int32) *ListDisksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDisksRequest) SetPageSize(v int32) *ListDisksRequest {
	s.PageSize = &v
	return s
}

func (s *ListDisksRequest) SetRegionId(v string) *ListDisksRequest {
	s.RegionId = &v
	return s
}

type ListDisksResponseBody struct {
	Disks      []*ListDisksResponseBodyDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	PageNumber *int32                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDisksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDisksResponseBody) GoString() string {
	return s.String()
}

func (s *ListDisksResponseBody) SetDisks(v []*ListDisksResponseBodyDisks) *ListDisksResponseBody {
	s.Disks = v
	return s
}

func (s *ListDisksResponseBody) SetPageNumber(v int32) *ListDisksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDisksResponseBody) SetPageSize(v int32) *ListDisksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDisksResponseBody) SetRequestId(v string) *ListDisksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDisksResponseBody) SetTotalCount(v int32) *ListDisksResponseBody {
	s.TotalCount = &v
	return s
}

type ListDisksResponseBodyDisks struct {
	Category       *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CreationTime   *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Device         *string `json:"Device,omitempty" xml:"Device,omitempty"`
	DiskChargeType *string `json:"DiskChargeType,omitempty" xml:"DiskChargeType,omitempty"`
	DiskId         *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	DiskName       *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	DiskType       *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Size           *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDisksResponseBodyDisks) String() string {
	return tea.Prettify(s)
}

func (s ListDisksResponseBodyDisks) GoString() string {
	return s.String()
}

func (s *ListDisksResponseBodyDisks) SetCategory(v string) *ListDisksResponseBodyDisks {
	s.Category = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetCreationTime(v string) *ListDisksResponseBodyDisks {
	s.CreationTime = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetDevice(v string) *ListDisksResponseBodyDisks {
	s.Device = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetDiskChargeType(v string) *ListDisksResponseBodyDisks {
	s.DiskChargeType = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetDiskId(v string) *ListDisksResponseBodyDisks {
	s.DiskId = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetDiskName(v string) *ListDisksResponseBodyDisks {
	s.DiskName = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetDiskType(v string) *ListDisksResponseBodyDisks {
	s.DiskType = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetInstanceId(v string) *ListDisksResponseBodyDisks {
	s.InstanceId = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetRegionId(v string) *ListDisksResponseBodyDisks {
	s.RegionId = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetSize(v int32) *ListDisksResponseBodyDisks {
	s.Size = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetStatus(v string) *ListDisksResponseBodyDisks {
	s.Status = &v
	return s
}

type ListDisksResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDisksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDisksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDisksResponse) GoString() string {
	return s.String()
}

func (s *ListDisksResponse) SetHeaders(v map[string]*string) *ListDisksResponse {
	s.Headers = v
	return s
}

func (s *ListDisksResponse) SetStatusCode(v int32) *ListDisksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDisksResponse) SetBody(v *ListDisksResponseBody) *ListDisksResponse {
	s.Body = v
	return s
}

type ListFirewallRulesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListFirewallRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFirewallRulesRequest) GoString() string {
	return s.String()
}

func (s *ListFirewallRulesRequest) SetInstanceId(v string) *ListFirewallRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFirewallRulesRequest) SetPageNumber(v int32) *ListFirewallRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFirewallRulesRequest) SetPageSize(v int32) *ListFirewallRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListFirewallRulesRequest) SetRegionId(v string) *ListFirewallRulesRequest {
	s.RegionId = &v
	return s
}

type ListFirewallRulesResponseBody struct {
	FirewallRules []*ListFirewallRulesResponseBodyFirewallRules `json:"FirewallRules,omitempty" xml:"FirewallRules,omitempty" type:"Repeated"`
	PageNumber    *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFirewallRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFirewallRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFirewallRulesResponseBody) SetFirewallRules(v []*ListFirewallRulesResponseBodyFirewallRules) *ListFirewallRulesResponseBody {
	s.FirewallRules = v
	return s
}

func (s *ListFirewallRulesResponseBody) SetPageNumber(v int32) *ListFirewallRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFirewallRulesResponseBody) SetPageSize(v int32) *ListFirewallRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFirewallRulesResponseBody) SetRequestId(v string) *ListFirewallRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFirewallRulesResponseBody) SetTotalCount(v int32) *ListFirewallRulesResponseBody {
	s.TotalCount = &v
	return s
}

type ListFirewallRulesResponseBodyFirewallRules struct {
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RuleId       *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
}

func (s ListFirewallRulesResponseBodyFirewallRules) String() string {
	return tea.Prettify(s)
}

func (s ListFirewallRulesResponseBodyFirewallRules) GoString() string {
	return s.String()
}

func (s *ListFirewallRulesResponseBodyFirewallRules) SetPort(v string) *ListFirewallRulesResponseBodyFirewallRules {
	s.Port = &v
	return s
}

func (s *ListFirewallRulesResponseBodyFirewallRules) SetRemark(v string) *ListFirewallRulesResponseBodyFirewallRules {
	s.Remark = &v
	return s
}

func (s *ListFirewallRulesResponseBodyFirewallRules) SetRuleId(v string) *ListFirewallRulesResponseBodyFirewallRules {
	s.RuleId = &v
	return s
}

func (s *ListFirewallRulesResponseBodyFirewallRules) SetRuleProtocol(v string) *ListFirewallRulesResponseBodyFirewallRules {
	s.RuleProtocol = &v
	return s
}

type ListFirewallRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFirewallRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFirewallRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFirewallRulesResponse) GoString() string {
	return s.String()
}

func (s *ListFirewallRulesResponse) SetHeaders(v map[string]*string) *ListFirewallRulesResponse {
	s.Headers = v
	return s
}

func (s *ListFirewallRulesResponse) SetStatusCode(v int32) *ListFirewallRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFirewallRulesResponse) SetBody(v *ListFirewallRulesResponseBody) *ListFirewallRulesResponse {
	s.Body = v
	return s
}

type ListImagesRequest struct {
	ImageIds  *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetImageIds(v string) *ListImagesRequest {
	s.ImageIds = &v
	return s
}

func (s *ListImagesRequest) SetImageType(v string) *ListImagesRequest {
	s.ImageType = &v
	return s
}

func (s *ListImagesRequest) SetRegionId(v string) *ListImagesRequest {
	s.RegionId = &v
	return s
}

type ListImagesResponseBody struct {
	Images    []*ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

type ListImagesResponseBodyImages struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName   *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageType   *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	Platform    *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s ListImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImages) SetDescription(v string) *ListImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageId(v string) *ListImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageName(v string) *ListImagesResponseBodyImages {
	s.ImageName = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageType(v string) *ListImagesResponseBodyImages {
	s.ImageType = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetPlatform(v string) *ListImagesResponseBodyImages {
	s.Platform = &v
	return s
}

type ListImagesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponse) GoString() string {
	return s.String()
}

func (s *ListImagesResponse) SetHeaders(v map[string]*string) *ListImagesResponse {
	s.Headers = v
	return s
}

func (s *ListImagesResponse) SetStatusCode(v int32) *ListImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

type ListInstancePlansModificationRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstancePlansModificationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePlansModificationRequest) GoString() string {
	return s.String()
}

func (s *ListInstancePlansModificationRequest) SetInstanceId(v string) *ListInstancePlansModificationRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancePlansModificationRequest) SetRegionId(v string) *ListInstancePlansModificationRequest {
	s.RegionId = &v
	return s
}

type ListInstancePlansModificationResponseBody struct {
	Plans     []*ListInstancePlansModificationResponseBodyPlans `json:"Plans,omitempty" xml:"Plans,omitempty" type:"Repeated"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancePlansModificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePlansModificationResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancePlansModificationResponseBody) SetPlans(v []*ListInstancePlansModificationResponseBodyPlans) *ListInstancePlansModificationResponseBody {
	s.Plans = v
	return s
}

func (s *ListInstancePlansModificationResponseBody) SetRequestId(v string) *ListInstancePlansModificationResponseBody {
	s.RequestId = &v
	return s
}

type ListInstancePlansModificationResponseBodyPlans struct {
	Bandwidth       *int32   `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Core            *int32   `json:"Core,omitempty" xml:"Core,omitempty"`
	Currency        *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiskSize        *int32   `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	DiskType        *string  `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	Flow            *int32   `json:"Flow,omitempty" xml:"Flow,omitempty"`
	Memory          *int32   `json:"Memory,omitempty" xml:"Memory,omitempty"`
	OriginPrice     *float64 `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty"`
	PlanId          *string  `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	SupportPlatform *string  `json:"SupportPlatform,omitempty" xml:"SupportPlatform,omitempty"`
}

func (s ListInstancePlansModificationResponseBodyPlans) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePlansModificationResponseBodyPlans) GoString() string {
	return s.String()
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetBandwidth(v int32) *ListInstancePlansModificationResponseBodyPlans {
	s.Bandwidth = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetCore(v int32) *ListInstancePlansModificationResponseBodyPlans {
	s.Core = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetCurrency(v string) *ListInstancePlansModificationResponseBodyPlans {
	s.Currency = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetDiskSize(v int32) *ListInstancePlansModificationResponseBodyPlans {
	s.DiskSize = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetDiskType(v string) *ListInstancePlansModificationResponseBodyPlans {
	s.DiskType = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetFlow(v int32) *ListInstancePlansModificationResponseBodyPlans {
	s.Flow = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetMemory(v int32) *ListInstancePlansModificationResponseBodyPlans {
	s.Memory = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetOriginPrice(v float64) *ListInstancePlansModificationResponseBodyPlans {
	s.OriginPrice = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetPlanId(v string) *ListInstancePlansModificationResponseBodyPlans {
	s.PlanId = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetSupportPlatform(v string) *ListInstancePlansModificationResponseBodyPlans {
	s.SupportPlatform = &v
	return s
}

type ListInstancePlansModificationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstancePlansModificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstancePlansModificationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePlansModificationResponse) GoString() string {
	return s.String()
}

func (s *ListInstancePlansModificationResponse) SetHeaders(v map[string]*string) *ListInstancePlansModificationResponse {
	s.Headers = v
	return s
}

func (s *ListInstancePlansModificationResponse) SetStatusCode(v int32) *ListInstancePlansModificationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancePlansModificationResponse) SetBody(v *ListInstancePlansModificationResponseBody) *ListInstancePlansModificationResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	ChargeType        *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	InstanceIds       *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	PageNumber        *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PublicIpAddresses *string `json:"PublicIpAddresses,omitempty" xml:"PublicIpAddresses,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetChargeType(v string) *ListInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceIds(v string) *ListInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int32) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetPublicIpAddresses(v string) *ListInstancesRequest {
	s.PublicIpAddresses = &v
	return s
}

func (s *ListInstancesRequest) SetRegionId(v string) *ListInstancesRequest {
	s.RegionId = &v
	return s
}

type ListInstancesResponseBody struct {
	Instances  []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetPageNumber(v int32) *ListInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesResponseBody) SetPageSize(v int32) *ListInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int32) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstancesResponseBodyInstances struct {
	BusinessStatus  *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	ChargeType      *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreationTime    *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DdosStatus      *string `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	ExpiredTime     *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ImageId         *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InnerIpAddress  *string `json:"InnerIpAddress,omitempty" xml:"InnerIpAddress,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PlanId          *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) SetBusinessStatus(v string) *ListInstancesResponseBodyInstances {
	s.BusinessStatus = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetChargeType(v string) *ListInstancesResponseBodyInstances {
	s.ChargeType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetCreationTime(v string) *ListInstancesResponseBodyInstances {
	s.CreationTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDdosStatus(v string) *ListInstancesResponseBodyInstances {
	s.DdosStatus = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetExpiredTime(v string) *ListInstancesResponseBodyInstances {
	s.ExpiredTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetImageId(v string) *ListInstancesResponseBodyInstances {
	s.ImageId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInnerIpAddress(v string) *ListInstancesResponseBodyInstances {
	s.InnerIpAddress = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceName(v string) *ListInstancesResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetPlanId(v string) *ListInstancesResponseBodyInstances {
	s.PlanId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetPublicIpAddress(v string) *ListInstancesResponseBodyInstances {
	s.PublicIpAddress = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetRegionId(v string) *ListInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStatus(v string) *ListInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type ListInstancesTrafficPackagesRequest struct {
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstancesTrafficPackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesTrafficPackagesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesTrafficPackagesRequest) SetInstanceIds(v string) *ListInstancesTrafficPackagesRequest {
	s.InstanceIds = &v
	return s
}

func (s *ListInstancesTrafficPackagesRequest) SetRegionId(v string) *ListInstancesTrafficPackagesRequest {
	s.RegionId = &v
	return s
}

type ListInstancesTrafficPackagesResponseBody struct {
	InstanceTrafficPackageUsages []*ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages `json:"InstanceTrafficPackageUsages,omitempty" xml:"InstanceTrafficPackageUsages,omitempty" type:"Repeated"`
	RequestId                    *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancesTrafficPackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesTrafficPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesTrafficPackagesResponseBody) SetInstanceTrafficPackageUsages(v []*ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) *ListInstancesTrafficPackagesResponseBody {
	s.InstanceTrafficPackageUsages = v
	return s
}

func (s *ListInstancesTrafficPackagesResponseBody) SetRequestId(v string) *ListInstancesTrafficPackagesResponseBody {
	s.RequestId = &v
	return s
}

type ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages struct {
	InstanceId              *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TrafficOverflow         *int64  `json:"TrafficOverflow,omitempty" xml:"TrafficOverflow,omitempty"`
	TrafficPackageRemaining *int64  `json:"TrafficPackageRemaining,omitempty" xml:"TrafficPackageRemaining,omitempty"`
	TrafficPackageTotal     *int64  `json:"TrafficPackageTotal,omitempty" xml:"TrafficPackageTotal,omitempty"`
	TrafficUsed             *int64  `json:"TrafficUsed,omitempty" xml:"TrafficUsed,omitempty"`
}

func (s ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) GoString() string {
	return s.String()
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) SetInstanceId(v string) *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) SetTrafficOverflow(v int64) *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages {
	s.TrafficOverflow = &v
	return s
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) SetTrafficPackageRemaining(v int64) *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages {
	s.TrafficPackageRemaining = &v
	return s
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) SetTrafficPackageTotal(v int64) *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages {
	s.TrafficPackageTotal = &v
	return s
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) SetTrafficUsed(v int64) *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages {
	s.TrafficUsed = &v
	return s
}

type ListInstancesTrafficPackagesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstancesTrafficPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstancesTrafficPackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesTrafficPackagesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesTrafficPackagesResponse) SetHeaders(v map[string]*string) *ListInstancesTrafficPackagesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesTrafficPackagesResponse) SetStatusCode(v int32) *ListInstancesTrafficPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesTrafficPackagesResponse) SetBody(v *ListInstancesTrafficPackagesResponseBody) *ListInstancesTrafficPackagesResponse {
	s.Body = v
	return s
}

type ListPlansRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPlansRequest) GoString() string {
	return s.String()
}

func (s *ListPlansRequest) SetRegionId(v string) *ListPlansRequest {
	s.RegionId = &v
	return s
}

type ListPlansResponseBody struct {
	Plans     []*ListPlansResponseBodyPlans `json:"Plans,omitempty" xml:"Plans,omitempty" type:"Repeated"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPlansResponseBody) GoString() string {
	return s.String()
}

func (s *ListPlansResponseBody) SetPlans(v []*ListPlansResponseBodyPlans) *ListPlansResponseBody {
	s.Plans = v
	return s
}

func (s *ListPlansResponseBody) SetRequestId(v string) *ListPlansResponseBody {
	s.RequestId = &v
	return s
}

type ListPlansResponseBodyPlans struct {
	Bandwidth       *int32   `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Core            *int32   `json:"Core,omitempty" xml:"Core,omitempty"`
	Currency        *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiskSize        *int32   `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	DiskType        *string  `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	Flow            *int32   `json:"Flow,omitempty" xml:"Flow,omitempty"`
	Memory          *int32   `json:"Memory,omitempty" xml:"Memory,omitempty"`
	OriginPrice     *float64 `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty"`
	PlanId          *string  `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	SupportPlatform *string  `json:"SupportPlatform,omitempty" xml:"SupportPlatform,omitempty"`
}

func (s ListPlansResponseBodyPlans) String() string {
	return tea.Prettify(s)
}

func (s ListPlansResponseBodyPlans) GoString() string {
	return s.String()
}

func (s *ListPlansResponseBodyPlans) SetBandwidth(v int32) *ListPlansResponseBodyPlans {
	s.Bandwidth = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetCore(v int32) *ListPlansResponseBodyPlans {
	s.Core = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetCurrency(v string) *ListPlansResponseBodyPlans {
	s.Currency = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetDiskSize(v int32) *ListPlansResponseBodyPlans {
	s.DiskSize = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetDiskType(v string) *ListPlansResponseBodyPlans {
	s.DiskType = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetFlow(v int32) *ListPlansResponseBodyPlans {
	s.Flow = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetMemory(v int32) *ListPlansResponseBodyPlans {
	s.Memory = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetOriginPrice(v float64) *ListPlansResponseBodyPlans {
	s.OriginPrice = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetPlanId(v string) *ListPlansResponseBodyPlans {
	s.PlanId = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetSupportPlatform(v string) *ListPlansResponseBodyPlans {
	s.SupportPlatform = &v
	return s
}

type ListPlansResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPlansResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPlansResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPlansResponse) GoString() string {
	return s.String()
}

func (s *ListPlansResponse) SetHeaders(v map[string]*string) *ListPlansResponse {
	s.Headers = v
	return s
}

func (s *ListPlansResponse) SetStatusCode(v int32) *ListPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPlansResponse) SetBody(v *ListPlansResponseBody) *ListPlansResponse {
	s.Body = v
	return s
}

type ListRegionsResponseBody struct {
	Regions   []*ListRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetRegions(v []*ListRegionsResponseBodyRegions) *ListRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

type ListRegionsResponseBodyRegions struct {
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegions) SetLocalName(v string) *ListRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionEndpoint(v string) *ListRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionId(v string) *ListRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type ListRegionsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetStatusCode(v int32) *ListRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ListSnapshotsRequest struct {
	DiskId      *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotIds *string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty"`
}

func (s ListSnapshotsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *ListSnapshotsRequest) SetDiskId(v string) *ListSnapshotsRequest {
	s.DiskId = &v
	return s
}

func (s *ListSnapshotsRequest) SetInstanceId(v string) *ListSnapshotsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSnapshotsRequest) SetPageNumber(v int32) *ListSnapshotsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSnapshotsRequest) SetPageSize(v int32) *ListSnapshotsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSnapshotsRequest) SetRegionId(v string) *ListSnapshotsRequest {
	s.RegionId = &v
	return s
}

func (s *ListSnapshotsRequest) SetSnapshotIds(v string) *ListSnapshotsRequest {
	s.SnapshotIds = &v
	return s
}

type ListSnapshotsResponseBody struct {
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Snapshots  []*ListSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSnapshotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBody) SetPageNumber(v int32) *ListSnapshotsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSnapshotsResponseBody) SetPageSize(v int32) *ListSnapshotsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSnapshotsResponseBody) SetRequestId(v string) *ListSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSnapshotsResponseBody) SetSnapshots(v []*ListSnapshotsResponseBodySnapshots) *ListSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *ListSnapshotsResponseBody) SetTotalCount(v int32) *ListSnapshotsResponseBody {
	s.TotalCount = &v
	return s
}

type ListSnapshotsResponseBodySnapshots struct {
	CreationTime   *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Progress       *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark         *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SnapshotId     *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SnapshotName   *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	SourceDiskId   *string `json:"SourceDiskId,omitempty" xml:"SourceDiskId,omitempty"`
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSnapshotsResponseBodySnapshots) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBodySnapshots) SetCreationTime(v string) *ListSnapshotsResponseBodySnapshots {
	s.CreationTime = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetProgress(v string) *ListSnapshotsResponseBodySnapshots {
	s.Progress = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetRegionId(v string) *ListSnapshotsResponseBodySnapshots {
	s.RegionId = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetRemark(v string) *ListSnapshotsResponseBodySnapshots {
	s.Remark = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetSnapshotId(v string) *ListSnapshotsResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetSnapshotName(v string) *ListSnapshotsResponseBodySnapshots {
	s.SnapshotName = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetSourceDiskId(v string) *ListSnapshotsResponseBodySnapshots {
	s.SourceDiskId = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetSourceDiskType(v string) *ListSnapshotsResponseBodySnapshots {
	s.SourceDiskType = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetStatus(v string) *ListSnapshotsResponseBodySnapshots {
	s.Status = &v
	return s
}

type ListSnapshotsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSnapshotsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponse) SetHeaders(v map[string]*string) *ListSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *ListSnapshotsResponse) SetStatusCode(v int32) *ListSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSnapshotsResponse) SetBody(v *ListSnapshotsResponseBody) *ListSnapshotsResponse {
	s.Body = v
	return s
}

type LoginInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Password   *string `json:"Password,omitempty" xml:"Password,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s LoginInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s LoginInstanceRequest) GoString() string {
	return s.String()
}

func (s *LoginInstanceRequest) SetInstanceId(v string) *LoginInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *LoginInstanceRequest) SetPassword(v string) *LoginInstanceRequest {
	s.Password = &v
	return s
}

func (s *LoginInstanceRequest) SetRegionId(v string) *LoginInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *LoginInstanceRequest) SetUsername(v string) *LoginInstanceRequest {
	s.Username = &v
	return s
}

type LoginInstanceResponseBody struct {
	RedirectUrl *string `json:"RedirectUrl,omitempty" xml:"RedirectUrl,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LoginInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LoginInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *LoginInstanceResponseBody) SetRedirectUrl(v string) *LoginInstanceResponseBody {
	s.RedirectUrl = &v
	return s
}

func (s *LoginInstanceResponseBody) SetRequestId(v string) *LoginInstanceResponseBody {
	s.RequestId = &v
	return s
}

type LoginInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LoginInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LoginInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s LoginInstanceResponse) GoString() string {
	return s.String()
}

func (s *LoginInstanceResponse) SetHeaders(v map[string]*string) *LoginInstanceResponse {
	s.Headers = v
	return s
}

func (s *LoginInstanceResponse) SetStatusCode(v int32) *LoginInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *LoginInstanceResponse) SetBody(v *LoginInstanceResponseBody) *LoginInstanceResponse {
	s.Body = v
	return s
}

type ModifyDatabaseInstanceDescriptionRequest struct {
	ClientToken                 *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DatabaseInstanceDescription *string `json:"DatabaseInstanceDescription,omitempty" xml:"DatabaseInstanceDescription,omitempty"`
	DatabaseInstanceId          *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	RegionId                    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDatabaseInstanceDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseInstanceDescriptionRequest) SetClientToken(v string) *ModifyDatabaseInstanceDescriptionRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDatabaseInstanceDescriptionRequest) SetDatabaseInstanceDescription(v string) *ModifyDatabaseInstanceDescriptionRequest {
	s.DatabaseInstanceDescription = &v
	return s
}

func (s *ModifyDatabaseInstanceDescriptionRequest) SetDatabaseInstanceId(v string) *ModifyDatabaseInstanceDescriptionRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *ModifyDatabaseInstanceDescriptionRequest) SetRegionId(v string) *ModifyDatabaseInstanceDescriptionRequest {
	s.RegionId = &v
	return s
}

type ModifyDatabaseInstanceDescriptionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDatabaseInstanceDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseInstanceDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseInstanceDescriptionResponseBody) SetRequestId(v string) *ModifyDatabaseInstanceDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDatabaseInstanceDescriptionResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDatabaseInstanceDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDatabaseInstanceDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseInstanceDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseInstanceDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDatabaseInstanceDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDatabaseInstanceDescriptionResponse) SetStatusCode(v int32) *ModifyDatabaseInstanceDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDatabaseInstanceDescriptionResponse) SetBody(v *ModifyDatabaseInstanceDescriptionResponseBody) *ModifyDatabaseInstanceDescriptionResponse {
	s.Body = v
	return s
}

type ModifyDatabaseInstanceParameterRequest struct {
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	ForceRestart       *bool   `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	Parameters         *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDatabaseInstanceParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseInstanceParameterRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseInstanceParameterRequest) SetClientToken(v string) *ModifyDatabaseInstanceParameterRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDatabaseInstanceParameterRequest) SetDatabaseInstanceId(v string) *ModifyDatabaseInstanceParameterRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *ModifyDatabaseInstanceParameterRequest) SetForceRestart(v bool) *ModifyDatabaseInstanceParameterRequest {
	s.ForceRestart = &v
	return s
}

func (s *ModifyDatabaseInstanceParameterRequest) SetParameters(v string) *ModifyDatabaseInstanceParameterRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyDatabaseInstanceParameterRequest) SetRegionId(v string) *ModifyDatabaseInstanceParameterRequest {
	s.RegionId = &v
	return s
}

type ModifyDatabaseInstanceParameterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDatabaseInstanceParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseInstanceParameterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseInstanceParameterResponseBody) SetRequestId(v string) *ModifyDatabaseInstanceParameterResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDatabaseInstanceParameterResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDatabaseInstanceParameterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDatabaseInstanceParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseInstanceParameterResponse) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseInstanceParameterResponse) SetHeaders(v map[string]*string) *ModifyDatabaseInstanceParameterResponse {
	s.Headers = v
	return s
}

func (s *ModifyDatabaseInstanceParameterResponse) SetStatusCode(v int32) *ModifyDatabaseInstanceParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDatabaseInstanceParameterResponse) SetBody(v *ModifyDatabaseInstanceParameterResponseBody) *ModifyDatabaseInstanceParameterResponse {
	s.Body = v
	return s
}

type ModifyImageShareStatusRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Operation   *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyImageShareStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageShareStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageShareStatusRequest) SetClientToken(v string) *ModifyImageShareStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyImageShareStatusRequest) SetImageId(v string) *ModifyImageShareStatusRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyImageShareStatusRequest) SetOperation(v string) *ModifyImageShareStatusRequest {
	s.Operation = &v
	return s
}

func (s *ModifyImageShareStatusRequest) SetRegionId(v string) *ModifyImageShareStatusRequest {
	s.RegionId = &v
	return s
}

type ModifyImageShareStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyImageShareStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageShareStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyImageShareStatusResponseBody) SetRequestId(v string) *ModifyImageShareStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyImageShareStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyImageShareStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyImageShareStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageShareStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyImageShareStatusResponse) SetHeaders(v map[string]*string) *ModifyImageShareStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyImageShareStatusResponse) SetStatusCode(v int32) *ModifyImageShareStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyImageShareStatusResponse) SetBody(v *ModifyImageShareStatusResponseBody) *ModifyImageShareStatusResponse {
	s.Body = v
	return s
}

type RebootInstanceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RebootInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RebootInstanceRequest) GoString() string {
	return s.String()
}

func (s *RebootInstanceRequest) SetClientToken(v string) *RebootInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RebootInstanceRequest) SetInstanceId(v string) *RebootInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RebootInstanceRequest) SetRegionId(v string) *RebootInstanceRequest {
	s.RegionId = &v
	return s
}

type RebootInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebootInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RebootInstanceResponseBody) SetRequestId(v string) *RebootInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RebootInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RebootInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RebootInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RebootInstanceResponse) GoString() string {
	return s.String()
}

func (s *RebootInstanceResponse) SetHeaders(v map[string]*string) *RebootInstanceResponse {
	s.Headers = v
	return s
}

func (s *RebootInstanceResponse) SetStatusCode(v int32) *RebootInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootInstanceResponse) SetBody(v *RebootInstanceResponseBody) *RebootInstanceResponse {
	s.Body = v
	return s
}

type ReleasePublicConnectionRequest struct {
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReleasePublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleasePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleasePublicConnectionRequest) SetClientToken(v string) *ReleasePublicConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *ReleasePublicConnectionRequest) SetDatabaseInstanceId(v string) *ReleasePublicConnectionRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *ReleasePublicConnectionRequest) SetRegionId(v string) *ReleasePublicConnectionRequest {
	s.RegionId = &v
	return s
}

type ReleasePublicConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleasePublicConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleasePublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleasePublicConnectionResponseBody) SetRequestId(v string) *ReleasePublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

type ReleasePublicConnectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReleasePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleasePublicConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleasePublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *ReleasePublicConnectionResponse) SetHeaders(v map[string]*string) *ReleasePublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *ReleasePublicConnectionResponse) SetStatusCode(v int32) *ReleasePublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleasePublicConnectionResponse) SetBody(v *ReleasePublicConnectionResponseBody) *ReleasePublicConnectionResponse {
	s.Body = v
	return s
}

type RenewInstanceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Period      *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) SetClientToken(v string) *RenewInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewInstanceRequest) SetInstanceId(v string) *RenewInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewInstanceRequest) SetPeriod(v int32) *RenewInstanceRequest {
	s.Period = &v
	return s
}

func (s *RenewInstanceRequest) SetRegionId(v string) *RenewInstanceRequest {
	s.RegionId = &v
	return s
}

type RenewInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBody) SetRequestId(v string) *RenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RenewInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RenewInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponse) SetHeaders(v map[string]*string) *RenewInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewInstanceResponse) SetStatusCode(v int32) *RenewInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewInstanceResponse) SetBody(v *RenewInstanceResponseBody) *RenewInstanceResponse {
	s.Body = v
	return s
}

type ResetDatabaseAccountPasswordRequest struct {
	AccountPassword    *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetDatabaseAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetDatabaseAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetDatabaseAccountPasswordRequest) SetAccountPassword(v string) *ResetDatabaseAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ResetDatabaseAccountPasswordRequest) SetClientToken(v string) *ResetDatabaseAccountPasswordRequest {
	s.ClientToken = &v
	return s
}

func (s *ResetDatabaseAccountPasswordRequest) SetDatabaseInstanceId(v string) *ResetDatabaseAccountPasswordRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *ResetDatabaseAccountPasswordRequest) SetRegionId(v string) *ResetDatabaseAccountPasswordRequest {
	s.RegionId = &v
	return s
}

type ResetDatabaseAccountPasswordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetDatabaseAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetDatabaseAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetDatabaseAccountPasswordResponseBody) SetRequestId(v string) *ResetDatabaseAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ResetDatabaseAccountPasswordResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetDatabaseAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetDatabaseAccountPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetDatabaseAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetDatabaseAccountPasswordResponse) SetHeaders(v map[string]*string) *ResetDatabaseAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetDatabaseAccountPasswordResponse) SetStatusCode(v int32) *ResetDatabaseAccountPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetDatabaseAccountPasswordResponse) SetBody(v *ResetDatabaseAccountPasswordResponseBody) *ResetDatabaseAccountPasswordResponse {
	s.Body = v
	return s
}

type ResetDiskRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DiskId      *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotId  *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ResetDiskRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetDiskRequest) GoString() string {
	return s.String()
}

func (s *ResetDiskRequest) SetClientToken(v string) *ResetDiskRequest {
	s.ClientToken = &v
	return s
}

func (s *ResetDiskRequest) SetDiskId(v string) *ResetDiskRequest {
	s.DiskId = &v
	return s
}

func (s *ResetDiskRequest) SetRegionId(v string) *ResetDiskRequest {
	s.RegionId = &v
	return s
}

func (s *ResetDiskRequest) SetSnapshotId(v string) *ResetDiskRequest {
	s.SnapshotId = &v
	return s
}

type ResetDiskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetDiskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetDiskResponseBody) GoString() string {
	return s.String()
}

func (s *ResetDiskResponseBody) SetRequestId(v string) *ResetDiskResponseBody {
	s.RequestId = &v
	return s
}

type ResetDiskResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetDiskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetDiskResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetDiskResponse) GoString() string {
	return s.String()
}

func (s *ResetDiskResponse) SetHeaders(v map[string]*string) *ResetDiskResponse {
	s.Headers = v
	return s
}

func (s *ResetDiskResponse) SetStatusCode(v int32) *ResetDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetDiskResponse) SetBody(v *ResetDiskResponseBody) *ResetDiskResponse {
	s.Body = v
	return s
}

type ResetSystemRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetSystemRequest) GoString() string {
	return s.String()
}

func (s *ResetSystemRequest) SetClientToken(v string) *ResetSystemRequest {
	s.ClientToken = &v
	return s
}

func (s *ResetSystemRequest) SetImageId(v string) *ResetSystemRequest {
	s.ImageId = &v
	return s
}

func (s *ResetSystemRequest) SetInstanceId(v string) *ResetSystemRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetSystemRequest) SetRegionId(v string) *ResetSystemRequest {
	s.RegionId = &v
	return s
}

type ResetSystemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetSystemResponseBody) GoString() string {
	return s.String()
}

func (s *ResetSystemResponseBody) SetRequestId(v string) *ResetSystemResponseBody {
	s.RequestId = &v
	return s
}

type ResetSystemResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetSystemResponse) GoString() string {
	return s.String()
}

func (s *ResetSystemResponse) SetHeaders(v map[string]*string) *ResetSystemResponse {
	s.Headers = v
	return s
}

func (s *ResetSystemResponse) SetStatusCode(v int32) *ResetSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetSystemResponse) SetBody(v *ResetSystemResponseBody) *ResetSystemResponse {
	s.Body = v
	return s
}

type RestartDatabaseInstanceRequest struct {
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RestartDatabaseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartDatabaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartDatabaseInstanceRequest) SetClientToken(v string) *RestartDatabaseInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartDatabaseInstanceRequest) SetDatabaseInstanceId(v string) *RestartDatabaseInstanceRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *RestartDatabaseInstanceRequest) SetRegionId(v string) *RestartDatabaseInstanceRequest {
	s.RegionId = &v
	return s
}

type RestartDatabaseInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartDatabaseInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartDatabaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDatabaseInstanceResponseBody) SetRequestId(v string) *RestartDatabaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RestartDatabaseInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RestartDatabaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartDatabaseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartDatabaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartDatabaseInstanceResponse) SetHeaders(v map[string]*string) *RestartDatabaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartDatabaseInstanceResponse) SetStatusCode(v int32) *RestartDatabaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartDatabaseInstanceResponse) SetBody(v *RestartDatabaseInstanceResponseBody) *RestartDatabaseInstanceResponse {
	s.Body = v
	return s
}

type RunCommandRequest struct {
	CommandContent      *string                `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	EnableParameter     *bool                  `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	InstanceId          *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name                *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Parameters          map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	RegionId            *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Timeout             *int32                 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	Type                *string                `json:"Type,omitempty" xml:"Type,omitempty"`
	WindowsPasswordName *string                `json:"WindowsPasswordName,omitempty" xml:"WindowsPasswordName,omitempty"`
	WorkingDir          *string                `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
	WorkingUser         *string                `json:"WorkingUser,omitempty" xml:"WorkingUser,omitempty"`
}

func (s RunCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCommandRequest) GoString() string {
	return s.String()
}

func (s *RunCommandRequest) SetCommandContent(v string) *RunCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *RunCommandRequest) SetEnableParameter(v bool) *RunCommandRequest {
	s.EnableParameter = &v
	return s
}

func (s *RunCommandRequest) SetInstanceId(v string) *RunCommandRequest {
	s.InstanceId = &v
	return s
}

func (s *RunCommandRequest) SetName(v string) *RunCommandRequest {
	s.Name = &v
	return s
}

func (s *RunCommandRequest) SetParameters(v map[string]interface{}) *RunCommandRequest {
	s.Parameters = v
	return s
}

func (s *RunCommandRequest) SetRegionId(v string) *RunCommandRequest {
	s.RegionId = &v
	return s
}

func (s *RunCommandRequest) SetTimeout(v int32) *RunCommandRequest {
	s.Timeout = &v
	return s
}

func (s *RunCommandRequest) SetType(v string) *RunCommandRequest {
	s.Type = &v
	return s
}

func (s *RunCommandRequest) SetWindowsPasswordName(v string) *RunCommandRequest {
	s.WindowsPasswordName = &v
	return s
}

func (s *RunCommandRequest) SetWorkingDir(v string) *RunCommandRequest {
	s.WorkingDir = &v
	return s
}

func (s *RunCommandRequest) SetWorkingUser(v string) *RunCommandRequest {
	s.WorkingUser = &v
	return s
}

type RunCommandShrinkRequest struct {
	CommandContent      *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	EnableParameter     *bool   `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParametersShrink    *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Timeout             *int32  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WindowsPasswordName *string `json:"WindowsPasswordName,omitempty" xml:"WindowsPasswordName,omitempty"`
	WorkingDir          *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
	WorkingUser         *string `json:"WorkingUser,omitempty" xml:"WorkingUser,omitempty"`
}

func (s RunCommandShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCommandShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunCommandShrinkRequest) SetCommandContent(v string) *RunCommandShrinkRequest {
	s.CommandContent = &v
	return s
}

func (s *RunCommandShrinkRequest) SetEnableParameter(v bool) *RunCommandShrinkRequest {
	s.EnableParameter = &v
	return s
}

func (s *RunCommandShrinkRequest) SetInstanceId(v string) *RunCommandShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *RunCommandShrinkRequest) SetName(v string) *RunCommandShrinkRequest {
	s.Name = &v
	return s
}

func (s *RunCommandShrinkRequest) SetParametersShrink(v string) *RunCommandShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *RunCommandShrinkRequest) SetRegionId(v string) *RunCommandShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RunCommandShrinkRequest) SetTimeout(v int32) *RunCommandShrinkRequest {
	s.Timeout = &v
	return s
}

func (s *RunCommandShrinkRequest) SetType(v string) *RunCommandShrinkRequest {
	s.Type = &v
	return s
}

func (s *RunCommandShrinkRequest) SetWindowsPasswordName(v string) *RunCommandShrinkRequest {
	s.WindowsPasswordName = &v
	return s
}

func (s *RunCommandShrinkRequest) SetWorkingDir(v string) *RunCommandShrinkRequest {
	s.WorkingDir = &v
	return s
}

func (s *RunCommandShrinkRequest) SetWorkingUser(v string) *RunCommandShrinkRequest {
	s.WorkingUser = &v
	return s
}

type RunCommandResponseBody struct {
	InvokeId  *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunCommandResponseBody) GoString() string {
	return s.String()
}

func (s *RunCommandResponseBody) SetInvokeId(v string) *RunCommandResponseBody {
	s.InvokeId = &v
	return s
}

func (s *RunCommandResponseBody) SetRequestId(v string) *RunCommandResponseBody {
	s.RequestId = &v
	return s
}

type RunCommandResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RunCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCommandResponse) GoString() string {
	return s.String()
}

func (s *RunCommandResponse) SetHeaders(v map[string]*string) *RunCommandResponse {
	s.Headers = v
	return s
}

func (s *RunCommandResponse) SetStatusCode(v int32) *RunCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCommandResponse) SetBody(v *RunCommandResponseBody) *RunCommandResponse {
	s.Body = v
	return s
}

type StartDatabaseInstanceRequest struct {
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartDatabaseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDatabaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartDatabaseInstanceRequest) SetClientToken(v string) *StartDatabaseInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StartDatabaseInstanceRequest) SetDatabaseInstanceId(v string) *StartDatabaseInstanceRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *StartDatabaseInstanceRequest) SetRegionId(v string) *StartDatabaseInstanceRequest {
	s.RegionId = &v
	return s
}

type StartDatabaseInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDatabaseInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDatabaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartDatabaseInstanceResponseBody) SetRequestId(v string) *StartDatabaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StartDatabaseInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartDatabaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartDatabaseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDatabaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartDatabaseInstanceResponse) SetHeaders(v map[string]*string) *StartDatabaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartDatabaseInstanceResponse) SetStatusCode(v int32) *StartDatabaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDatabaseInstanceResponse) SetBody(v *StartDatabaseInstanceResponseBody) *StartDatabaseInstanceResponse {
	s.Body = v
	return s
}

type StartInstanceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) SetClientToken(v string) *StartInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StartInstanceRequest) SetInstanceId(v string) *StartInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StartInstanceRequest) SetRegionId(v string) *StartInstanceRequest {
	s.RegionId = &v
	return s
}

type StartInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) SetRequestId(v string) *StartInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StartInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartInstanceResponse) SetHeaders(v map[string]*string) *StartInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartInstanceResponse) SetStatusCode(v int32) *StartInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartInstanceResponse) SetBody(v *StartInstanceResponseBody) *StartInstanceResponse {
	s.Body = v
	return s
}

type StopDatabaseInstanceRequest struct {
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopDatabaseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDatabaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopDatabaseInstanceRequest) SetClientToken(v string) *StopDatabaseInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StopDatabaseInstanceRequest) SetDatabaseInstanceId(v string) *StopDatabaseInstanceRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *StopDatabaseInstanceRequest) SetRegionId(v string) *StopDatabaseInstanceRequest {
	s.RegionId = &v
	return s
}

type StopDatabaseInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDatabaseInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDatabaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopDatabaseInstanceResponseBody) SetRequestId(v string) *StopDatabaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StopDatabaseInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopDatabaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopDatabaseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDatabaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopDatabaseInstanceResponse) SetHeaders(v map[string]*string) *StopDatabaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopDatabaseInstanceResponse) SetStatusCode(v int32) *StopDatabaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDatabaseInstanceResponse) SetBody(v *StopDatabaseInstanceResponseBody) *StopDatabaseInstanceResponse {
	s.Body = v
	return s
}

type StopInstanceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopInstanceRequest) SetClientToken(v string) *StopInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StopInstanceRequest) SetInstanceId(v string) *StopInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StopInstanceRequest) SetRegionId(v string) *StopInstanceRequest {
	s.RegionId = &v
	return s
}

type StopInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstanceResponseBody) SetRequestId(v string) *StopInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StopInstanceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopInstanceResponse) SetHeaders(v map[string]*string) *StopInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopInstanceResponse) SetStatusCode(v int32) *StopInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopInstanceResponse) SetBody(v *StopInstanceResponseBody) *StopInstanceResponse {
	s.Body = v
	return s
}

type UpdateInstanceAttributeRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Password     *string `json:"Password,omitempty" xml:"Password,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAttributeRequest) SetClientToken(v string) *UpdateInstanceAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetInstanceId(v string) *UpdateInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetInstanceName(v string) *UpdateInstanceAttributeRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetPassword(v string) *UpdateInstanceAttributeRequest {
	s.Password = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetRegionId(v string) *UpdateInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

type UpdateInstanceAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAttributeResponseBody) SetRequestId(v string) *UpdateInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstanceAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAttributeResponse) SetHeaders(v map[string]*string) *UpdateInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceAttributeResponse) SetStatusCode(v int32) *UpdateInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceAttributeResponse) SetBody(v *UpdateInstanceAttributeResponseBody) *UpdateInstanceAttributeResponse {
	s.Body = v
	return s
}

type UpgradeInstanceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PlanId      *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpgradeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceRequest) SetClientToken(v string) *UpgradeInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeInstanceRequest) SetInstanceId(v string) *UpgradeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeInstanceRequest) SetPlanId(v string) *UpgradeInstanceRequest {
	s.PlanId = &v
	return s
}

func (s *UpgradeInstanceRequest) SetRegionId(v string) *UpgradeInstanceRequest {
	s.RegionId = &v
	return s
}

type UpgradeInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceResponseBody) SetRequestId(v string) *UpgradeInstanceResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceResponse) SetHeaders(v map[string]*string) *UpgradeInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpgradeInstanceResponse) SetStatusCode(v int32) *UpgradeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeInstanceResponse) SetBody(v *UpgradeInstanceResponseBody) *UpgradeInstanceResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("swas-open"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AllocatePublicConnectionWithOptions(request *AllocatePublicConnectionRequest, runtime *util.RuntimeOptions) (_result *AllocatePublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AllocatePublicConnection"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AllocatePublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AllocatePublicConnection(request *AllocatePublicConnectionRequest) (_result *AllocatePublicConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocatePublicConnectionResponse{}
	_body, _err := client.AllocatePublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCustomImageWithOptions(request *CreateCustomImageRequest, runtime *util.RuntimeOptions) (_result *CreateCustomImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataSnapshotId)) {
		query["DataSnapshotId"] = request.DataSnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ImageName)) {
		query["ImageName"] = request.ImageName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SystemSnapshotId)) {
		query["SystemSnapshotId"] = request.SystemSnapshotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCustomImage"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCustomImage(request *CreateCustomImageRequest) (_result *CreateCustomImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCustomImageResponse{}
	_body, _err := client.CreateCustomImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFirewallRuleWithOptions(request *CreateFirewallRuleRequest, runtime *util.RuntimeOptions) (_result *CreateFirewallRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.RuleProtocol)) {
		query["RuleProtocol"] = request.RuleProtocol
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFirewallRule"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFirewallRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFirewallRule(request *CreateFirewallRuleRequest) (_result *CreateFirewallRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFirewallRuleResponse{}
	_body, _err := client.CreateFirewallRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstancesWithOptions(request *CreateInstancesRequest, runtime *util.RuntimeOptions) (_result *CreateInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		query["Amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewPeriod)) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataDiskSize)) {
		query["DataDiskSize"] = request.DataDiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstances"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstances(request *CreateInstancesRequest) (_result *CreateInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstancesResponse{}
	_body, _err := client.CreateInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSnapshotWithOptions(request *CreateSnapshotRequest, runtime *util.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DiskId)) {
		query["DiskId"] = request.DiskId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotName)) {
		query["SnapshotName"] = request.SnapshotName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSnapshot"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSnapshot(request *CreateSnapshotRequest) (_result *CreateSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CreateSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCustomImageWithOptions(request *DeleteCustomImageRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomImage"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCustomImage(request *DeleteCustomImageRequest) (_result *DeleteCustomImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomImageResponse{}
	_body, _err := client.DeleteCustomImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFirewallRuleWithOptions(request *DeleteFirewallRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteFirewallRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFirewallRule"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFirewallRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFirewallRule(request *DeleteFirewallRuleRequest) (_result *DeleteFirewallRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFirewallRuleResponse{}
	_body, _err := client.DeleteFirewallRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSnapshotWithOptions(request *DeleteSnapshotRequest, runtime *util.RuntimeOptions) (_result *DeleteSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSnapshot"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSnapshot(request *DeleteSnapshotRequest) (_result *DeleteSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DeleteSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCloudAssistantStatusWithOptions(tmpReq *DescribeCloudAssistantStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudAssistantStatusResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeCloudAssistantStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InstanceIds)) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, tea.String("InstanceIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIdsShrink)) {
		query["InstanceIds"] = request.InstanceIdsShrink
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCloudAssistantStatus"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCloudAssistantStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCloudAssistantStatus(request *DescribeCloudAssistantStatusRequest) (_result *DescribeCloudAssistantStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudAssistantStatusResponse{}
	_body, _err := client.DescribeCloudAssistantStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDatabaseErrorLogsWithOptions(request *DescribeDatabaseErrorLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeDatabaseErrorLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
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

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDatabaseErrorLogs"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDatabaseErrorLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDatabaseErrorLogs(request *DescribeDatabaseErrorLogsRequest) (_result *DescribeDatabaseErrorLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDatabaseErrorLogsResponse{}
	_body, _err := client.DescribeDatabaseErrorLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDatabaseInstanceMetricDataWithOptions(request *DescribeDatabaseInstanceMetricDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDatabaseInstanceMetricDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDatabaseInstanceMetricData"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDatabaseInstanceMetricDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDatabaseInstanceMetricData(request *DescribeDatabaseInstanceMetricDataRequest) (_result *DescribeDatabaseInstanceMetricDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDatabaseInstanceMetricDataResponse{}
	_body, _err := client.DescribeDatabaseInstanceMetricDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDatabaseInstanceParametersWithOptions(request *DescribeDatabaseInstanceParametersRequest, runtime *util.RuntimeOptions) (_result *DescribeDatabaseInstanceParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDatabaseInstanceParameters"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDatabaseInstanceParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDatabaseInstanceParameters(request *DescribeDatabaseInstanceParametersRequest) (_result *DescribeDatabaseInstanceParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDatabaseInstanceParametersResponse{}
	_body, _err := client.DescribeDatabaseInstanceParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDatabaseInstancesWithOptions(request *DescribeDatabaseInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDatabaseInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceIds)) {
		query["DatabaseInstanceIds"] = request.DatabaseInstanceIds
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDatabaseInstances"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDatabaseInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDatabaseInstances(request *DescribeDatabaseInstancesRequest) (_result *DescribeDatabaseInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDatabaseInstancesResponse{}
	_body, _err := client.DescribeDatabaseInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDatabaseSlowLogRecordsWithOptions(request *DescribeDatabaseSlowLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeDatabaseSlowLogRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
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

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDatabaseSlowLogRecords"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDatabaseSlowLogRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDatabaseSlowLogRecords(request *DescribeDatabaseSlowLogRecordsRequest) (_result *DescribeDatabaseSlowLogRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDatabaseSlowLogRecordsResponse{}
	_body, _err := client.DescribeDatabaseSlowLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInvocationResultWithOptions(request *DescribeInvocationResultRequest, runtime *util.RuntimeOptions) (_result *DescribeInvocationResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InvokeId)) {
		query["InvokeId"] = request.InvokeId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInvocationResult"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInvocationResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInvocationResult(request *DescribeInvocationResultRequest) (_result *DescribeInvocationResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInvocationResultResponse{}
	_body, _err := client.DescribeInvocationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInvocationsWithOptions(request *DescribeInvocationsRequest, runtime *util.RuntimeOptions) (_result *DescribeInvocationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InvokeStatus)) {
		query["InvokeStatus"] = request.InvokeStatus
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInvocations"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInvocations(request *DescribeInvocationsRequest) (_result *DescribeInvocationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.DescribeInvocationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallCloudAssistantWithOptions(tmpReq *InstallCloudAssistantRequest, runtime *util.RuntimeOptions) (_result *InstallCloudAssistantResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &InstallCloudAssistantShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InstanceIds)) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, tea.String("InstanceIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIdsShrink)) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallCloudAssistant"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallCloudAssistantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstallCloudAssistant(request *InstallCloudAssistantRequest) (_result *InstallCloudAssistantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallCloudAssistantResponse{}
	_body, _err := client.InstallCloudAssistantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDisksWithOptions(request *ListDisksRequest, runtime *util.RuntimeOptions) (_result *ListDisksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiskIds)) {
		query["DiskIds"] = request.DiskIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDisks"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDisksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDisks(request *ListDisksRequest) (_result *ListDisksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDisksResponse{}
	_body, _err := client.ListDisksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFirewallRulesWithOptions(request *ListFirewallRulesRequest, runtime *util.RuntimeOptions) (_result *ListFirewallRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFirewallRules"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFirewallRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFirewallRules(request *ListFirewallRulesRequest) (_result *ListFirewallRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFirewallRulesResponse{}
	_body, _err := client.ListFirewallRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListImagesWithOptions(request *ListImagesRequest, runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageIds)) {
		query["ImageIds"] = request.ImageIds
	}

	if !tea.BoolValue(util.IsUnset(request.ImageType)) {
		query["ImageType"] = request.ImageType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListImages"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstancePlansModificationWithOptions(request *ListInstancePlansModificationRequest, runtime *util.RuntimeOptions) (_result *ListInstancePlansModificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstancePlansModification"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancePlansModificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstancePlansModification(request *ListInstancePlansModificationRequest) (_result *ListInstancePlansModificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancePlansModificationResponse{}
	_body, _err := client.ListInstancePlansModificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PublicIpAddresses)) {
		query["PublicIpAddresses"] = request.PublicIpAddresses
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstancesTrafficPackagesWithOptions(request *ListInstancesTrafficPackagesRequest, runtime *util.RuntimeOptions) (_result *ListInstancesTrafficPackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstancesTrafficPackages"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesTrafficPackagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstancesTrafficPackages(request *ListInstancesTrafficPackagesRequest) (_result *ListInstancesTrafficPackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancesTrafficPackagesResponse{}
	_body, _err := client.ListInstancesTrafficPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPlansWithOptions(request *ListPlansRequest, runtime *util.RuntimeOptions) (_result *ListPlansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPlans"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPlans(request *ListPlansRequest) (_result *ListPlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPlansResponse{}
	_body, _err := client.ListPlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRegionsWithOptions(runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListRegions"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSnapshotsWithOptions(request *ListSnapshotsRequest, runtime *util.RuntimeOptions) (_result *ListSnapshotsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiskId)) {
		query["DiskId"] = request.DiskId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
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

	if !tea.BoolValue(util.IsUnset(request.SnapshotIds)) {
		query["SnapshotIds"] = request.SnapshotIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSnapshots"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSnapshotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSnapshots(request *ListSnapshotsRequest) (_result *ListSnapshotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSnapshotsResponse{}
	_body, _err := client.ListSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LoginInstanceWithOptions(request *LoginInstanceRequest, runtime *util.RuntimeOptions) (_result *LoginInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LoginInstance"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LoginInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LoginInstance(request *LoginInstanceRequest) (_result *LoginInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LoginInstanceResponse{}
	_body, _err := client.LoginInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDatabaseInstanceDescriptionWithOptions(request *ModifyDatabaseInstanceDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDatabaseInstanceDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceDescription)) {
		query["DatabaseInstanceDescription"] = request.DatabaseInstanceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDatabaseInstanceDescription"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDatabaseInstanceDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDatabaseInstanceDescription(request *ModifyDatabaseInstanceDescriptionRequest) (_result *ModifyDatabaseInstanceDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDatabaseInstanceDescriptionResponse{}
	_body, _err := client.ModifyDatabaseInstanceDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDatabaseInstanceParameterWithOptions(request *ModifyDatabaseInstanceParameterRequest, runtime *util.RuntimeOptions) (_result *ModifyDatabaseInstanceParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ForceRestart)) {
		query["ForceRestart"] = request.ForceRestart
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDatabaseInstanceParameter"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDatabaseInstanceParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDatabaseInstanceParameter(request *ModifyDatabaseInstanceParameterRequest) (_result *ModifyDatabaseInstanceParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDatabaseInstanceParameterResponse{}
	_body, _err := client.ModifyDatabaseInstanceParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyImageShareStatusWithOptions(request *ModifyImageShareStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyImageShareStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		query["Operation"] = request.Operation
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyImageShareStatus"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyImageShareStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyImageShareStatus(request *ModifyImageShareStatusRequest) (_result *ModifyImageShareStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyImageShareStatusResponse{}
	_body, _err := client.ModifyImageShareStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebootInstanceWithOptions(request *RebootInstanceRequest, runtime *util.RuntimeOptions) (_result *RebootInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RebootInstance"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RebootInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebootInstance(request *RebootInstanceRequest) (_result *RebootInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebootInstanceResponse{}
	_body, _err := client.RebootInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleasePublicConnectionWithOptions(request *ReleasePublicConnectionRequest, runtime *util.RuntimeOptions) (_result *ReleasePublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleasePublicConnection"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleasePublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleasePublicConnection(request *ReleasePublicConnectionRequest) (_result *ReleasePublicConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleasePublicConnectionResponse{}
	_body, _err := client.ReleasePublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewInstanceWithOptions(request *RenewInstanceRequest, runtime *util.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewInstance"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewInstance(request *RenewInstanceRequest) (_result *RenewInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewInstanceResponse{}
	_body, _err := client.RenewInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetDatabaseAccountPasswordWithOptions(request *ResetDatabaseAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetDatabaseAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountPassword)) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetDatabaseAccountPassword"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetDatabaseAccountPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetDatabaseAccountPassword(request *ResetDatabaseAccountPasswordRequest) (_result *ResetDatabaseAccountPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetDatabaseAccountPasswordResponse{}
	_body, _err := client.ResetDatabaseAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetDiskWithOptions(request *ResetDiskRequest, runtime *util.RuntimeOptions) (_result *ResetDiskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DiskId)) {
		query["DiskId"] = request.DiskId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetDisk"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetDiskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetDisk(request *ResetDiskRequest) (_result *ResetDiskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetDiskResponse{}
	_body, _err := client.ResetDiskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetSystemWithOptions(request *ResetSystemRequest, runtime *util.RuntimeOptions) (_result *ResetSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetSystem"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetSystem(request *ResetSystemRequest) (_result *ResetSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetSystemResponse{}
	_body, _err := client.ResetSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartDatabaseInstanceWithOptions(request *RestartDatabaseInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartDatabaseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartDatabaseInstance"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartDatabaseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartDatabaseInstance(request *RestartDatabaseInstanceRequest) (_result *RestartDatabaseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartDatabaseInstanceResponse{}
	_body, _err := client.RestartDatabaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunCommandWithOptions(tmpReq *RunCommandRequest, runtime *util.RuntimeOptions) (_result *RunCommandResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunCommandShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommandContent)) {
		query["CommandContent"] = request.CommandContent
	}

	if !tea.BoolValue(util.IsUnset(request.EnableParameter)) {
		query["EnableParameter"] = request.EnableParameter
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		query["Timeout"] = request.Timeout
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.WindowsPasswordName)) {
		query["WindowsPasswordName"] = request.WindowsPasswordName
	}

	if !tea.BoolValue(util.IsUnset(request.WorkingDir)) {
		query["WorkingDir"] = request.WorkingDir
	}

	if !tea.BoolValue(util.IsUnset(request.WorkingUser)) {
		query["WorkingUser"] = request.WorkingUser
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RunCommand"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunCommand(request *RunCommandRequest) (_result *RunCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunCommandResponse{}
	_body, _err := client.RunCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartDatabaseInstanceWithOptions(request *StartDatabaseInstanceRequest, runtime *util.RuntimeOptions) (_result *StartDatabaseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartDatabaseInstance"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartDatabaseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartDatabaseInstance(request *StartDatabaseInstanceRequest) (_result *StartDatabaseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDatabaseInstanceResponse{}
	_body, _err := client.StartDatabaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartInstanceWithOptions(request *StartInstanceRequest, runtime *util.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartInstance"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartInstance(request *StartInstanceRequest) (_result *StartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartInstanceResponse{}
	_body, _err := client.StartInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopDatabaseInstanceWithOptions(request *StopDatabaseInstanceRequest, runtime *util.RuntimeOptions) (_result *StopDatabaseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopDatabaseInstance"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopDatabaseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopDatabaseInstance(request *StopDatabaseInstanceRequest) (_result *StopDatabaseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDatabaseInstanceResponse{}
	_body, _err := client.StopDatabaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopInstanceWithOptions(request *StopInstanceRequest, runtime *util.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopInstance"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopInstance(request *StopInstanceRequest) (_result *StopInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopInstanceResponse{}
	_body, _err := client.StopInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstanceAttributeWithOptions(request *UpdateInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceAttribute"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstanceAttribute(request *UpdateInstanceAttributeRequest) (_result *UpdateInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceAttributeResponse{}
	_body, _err := client.UpdateInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeInstanceWithOptions(request *UpgradeInstanceRequest, runtime *util.RuntimeOptions) (_result *UpgradeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeInstance"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeInstance(request *UpgradeInstanceRequest) (_result *UpgradeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeInstanceResponse{}
	_body, _err := client.UpgradeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
