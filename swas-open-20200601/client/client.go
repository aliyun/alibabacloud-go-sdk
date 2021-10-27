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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCustomImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFirewallRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCustomImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFirewallRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteSnapshotResponse) SetBody(v *DeleteSnapshotResponseBody) *DeleteSnapshotResponse {
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDisksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFirewallRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstancePlansModificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstancesTrafficPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPlansResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListSnapshotsResponse) SetBody(v *ListSnapshotsResponseBody) *ListSnapshotsResponse {
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyImageShareStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RebootInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RebootInstanceResponse) SetBody(v *RebootInstanceResponseBody) *RebootInstanceResponse {
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenewInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RenewInstanceResponse) SetBody(v *RenewInstanceResponseBody) *RenewInstanceResponse {
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetDiskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ResetSystemResponse) SetBody(v *ResetSystemResponseBody) *ResetSystemResponse {
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StartInstanceResponse) SetBody(v *StartInstanceResponseBody) *StartInstanceResponse {
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (client *Client) CreateCustomImageWithOptions(request *CreateCustomImageRequest, runtime *util.RuntimeOptions) (_result *CreateCustomImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCustomImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCustomImage"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFirewallRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFirewallRule"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateInstances"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSnapshot"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCustomImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCustomImage"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFirewallRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFirewallRule"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSnapshot"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListDisksWithOptions(request *ListDisksRequest, runtime *util.RuntimeOptions) (_result *ListDisksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDisksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDisks"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFirewallRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFirewallRules"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListImages"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListInstancePlansModificationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInstancePlansModification"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInstances"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListInstancesTrafficPackagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInstancesTrafficPackages"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPlansResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPlans"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_result = &ListRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRegions"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSnapshotsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSnapshots"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyImageShareStatusWithOptions(request *ModifyImageShareStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyImageShareStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyImageShareStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyImageShareStatus"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RebootInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RebootInstance"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RenewInstanceWithOptions(request *RenewInstanceRequest, runtime *util.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RenewInstance"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ResetDiskWithOptions(request *ResetDiskRequest, runtime *util.RuntimeOptions) (_result *ResetDiskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetDiskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetDisk"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetSystemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetSystem"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) StartInstanceWithOptions(request *StartInstanceRequest, runtime *util.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartInstance"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) StopInstanceWithOptions(request *StopInstanceRequest, runtime *util.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopInstance"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateInstanceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateInstanceAttribute"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeInstance"), tea.String("2020-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
