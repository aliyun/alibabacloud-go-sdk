// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorType(v string) *ListNodesShrinkRequest
	GetAcceleratorType() *string
	SetAvailabilityZone(v string) *ListNodesShrinkRequest
	GetAvailabilityZone() *string
	SetCliqueID(v string) *ListNodesShrinkRequest
	GetCliqueID() *string
	SetDiskPL(v string) *ListNodesShrinkRequest
	GetDiskPL() *string
	SetFilterByQuotaId(v string) *ListNodesShrinkRequest
	GetFilterByQuotaId() *string
	SetFilterByResourceGroupIds(v string) *ListNodesShrinkRequest
	GetFilterByResourceGroupIds() *string
	SetGPUType(v string) *ListNodesShrinkRequest
	GetGPUType() *string
	SetHealthCountShrink(v string) *ListNodesShrinkRequest
	GetHealthCountShrink() *string
	SetHealthRateShrink(v string) *ListNodesShrinkRequest
	GetHealthRateShrink() *string
	SetHyperNode(v string) *ListNodesShrinkRequest
	GetHyperNode() *string
	SetHyperZone(v string) *ListNodesShrinkRequest
	GetHyperZone() *string
	SetLayoutMode(v string) *ListNodesShrinkRequest
	GetLayoutMode() *string
	SetMachineGroupIds(v string) *ListNodesShrinkRequest
	GetMachineGroupIds() *string
	SetNodeNames(v string) *ListNodesShrinkRequest
	GetNodeNames() *string
	SetNodeStatuses(v string) *ListNodesShrinkRequest
	GetNodeStatuses() *string
	SetNodeTypes(v string) *ListNodesShrinkRequest
	GetNodeTypes() *string
	SetOrder(v string) *ListNodesShrinkRequest
	GetOrder() *string
	SetOrderInstanceIds(v string) *ListNodesShrinkRequest
	GetOrderInstanceIds() *string
	SetOrderStatuses(v string) *ListNodesShrinkRequest
	GetOrderStatuses() *string
	SetPageNumber(v int32) *ListNodesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNodesShrinkRequest
	GetPageSize() *int32
	SetPaymentType(v string) *ListNodesShrinkRequest
	GetPaymentType() *string
	SetPodNum(v int32) *ListNodesShrinkRequest
	GetPodNum() *int32
	SetQuotaId(v string) *ListNodesShrinkRequest
	GetQuotaId() *string
	SetReasonCodes(v string) *ListNodesShrinkRequest
	GetReasonCodes() *string
	SetResourceGroupIds(v string) *ListNodesShrinkRequest
	GetResourceGroupIds() *string
	SetResourceGroupName(v string) *ListNodesShrinkRequest
	GetResourceGroupName() *string
	SetSortBy(v string) *ListNodesShrinkRequest
	GetSortBy() *string
	SetVerbose(v bool) *ListNodesShrinkRequest
	GetVerbose() *bool
	SetWorkloadNum(v int32) *ListNodesShrinkRequest
	GetWorkloadNum() *int32
	SetWorkspaceId(v string) *ListNodesShrinkRequest
	GetWorkspaceId() *string
}

type ListNodesShrinkRequest struct {
	// The accelerator type. Valid values:
	//
	// - CPU
	//
	// - GPU
	//
	// If this parameter is left empty, all types are queried.
	//
	// example:
	//
	// CPU
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// The zone where the nodes reside.
	//
	// example:
	//
	// C0
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	// The clique ID.
	CliqueID *string `json:"CliqueID,omitempty" xml:"CliqueID,omitempty"`
	// The performance level (PL) of the cloud disk.
	//
	// example:
	//
	// PL2
	DiskPL *string `json:"DiskPL,omitempty" xml:"DiskPL,omitempty"`
	// Used together with ResourceGroupIds to display nodes in the resource group that are bound to specific quotas.
	//
	// example:
	//
	// quotamtl37ge7gkvdz
	FilterByQuotaId *string `json:"FilterByQuotaId,omitempty" xml:"FilterByQuotaId,omitempty"`
	// Used together with QuotaId to display nodes in the quota that belong to specific resource groups.
	//
	// example:
	//
	// rg69rj0leslwdnbe
	FilterByResourceGroupIds *string `json:"FilterByResourceGroupIds,omitempty" xml:"FilterByResourceGroupIds,omitempty"`
	// The GPU card type. Fuzzy match is supported.
	//
	// example:
	//
	// T4
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// The number of healthy nodes.
	HealthCountShrink *string `json:"HealthCount,omitempty" xml:"HealthCount,omitempty"`
	// The health rate of the nodes, which is the proportion of healthy nodes.
	HealthRateShrink *string `json:"HealthRate,omitempty" xml:"HealthRate,omitempty"`
	// The name of the hypernode.
	HyperNode *string `json:"HyperNode,omitempty" xml:"HyperNode,omitempty"`
	// The hyperzone where the nodes reside.
	//
	// example:
	//
	// A0
	HyperZone *string `json:"HyperZone,omitempty" xml:"HyperZone,omitempty"`
	// The layout mode.
	//
	// example:
	//
	// Tree
	LayoutMode *string `json:"LayoutMode,omitempty" xml:"LayoutMode,omitempty"`
	// The IDs of the machine groups to which the nodes belong.
	//
	// example:
	//
	// mg124
	MachineGroupIds *string `json:"MachineGroupIds,omitempty" xml:"MachineGroupIds,omitempty"`
	// The names of the nodes to display. Separate multiple names with commas.
	//
	// example:
	//
	// lingjxxxx
	NodeNames *string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty"`
	// The status of the node. Separate multiple statuses with commas. If you do not specify this parameter, the system returns a list of nodes in all statuses.
	//
	// example:
	//
	// Ready
	NodeStatuses *string `json:"NodeStatuses,omitempty" xml:"NodeStatuses,omitempty"`
	// The list of resource node specifications. Separate multiple specifications with commas. If you do not specify this parameter, the system returns a list of machines of all resource node specifications.
	//
	// example:
	//
	// ecs.c6.xlarge
	NodeTypes *string `json:"NodeTypes,omitempty" xml:"NodeTypes,omitempty"`
	// The sort order. Valid values:
	//
	// - desc: descending order.
	//
	// - asc: ascending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The order ID information.
	//
	// example:
	//
	// 260590501560397
	OrderInstanceIds *string `json:"OrderInstanceIds,omitempty" xml:"OrderInstanceIds,omitempty"`
	// The order status of the node.
	//
	// example:
	//
	// Ready
	OrderStatuses *string `json:"OrderStatuses,omitempty" xml:"OrderStatuses,omitempty"`
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page in a paged query. This parameter is used for paging.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The billing method of the node order.
	//
	// example:
	//
	// PREPAY
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The number of pods on the node.
	//
	// example:
	//
	// 1
	PodNum *int32 `json:"PodNum,omitempty" xml:"PodNum,omitempty"`
	// The ID of the resource quota to which the node belongs.
	//
	// example:
	//
	// quotamtl37ge7gkvdz
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// The reason codes for the current status of the nodes.
	//
	// example:
	//
	// CordonByUser
	ReasonCodes *string `json:"ReasonCodes,omitempty" xml:"ReasonCodes,omitempty"`
	// The list of resource group IDs. Separate multiple IDs with commas. You must specify either this parameter or QuotaId.
	//
	// Constraints:
	//
	// 1. The user ID of the specified resource group must match the user ID of the request.
	//
	// 2. The resource groups in the list must be of the same type.
	//
	// 3. The VPCs associated with the resource groups in the list must be the same.
	//
	// example:
	//
	// rg69rj0leslwdnbe
	ResourceGroupIds *string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// rg-dlc
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The field based on which the results are sorted.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// Specifies whether to return resource usage information. This parameter takes effect only when QuotaId is specified.
	//
	// example:
	//
	// false
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// The number of workloads on the node.
	//
	// example:
	//
	// 1
	WorkloadNum *int32 `json:"WorkloadNum,omitempty" xml:"WorkloadNum,omitempty"`
	// The ID of the workspace to which the nodes belong.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListNodesShrinkRequest) GetAcceleratorType() *string {
	return s.AcceleratorType
}

func (s *ListNodesShrinkRequest) GetAvailabilityZone() *string {
	return s.AvailabilityZone
}

func (s *ListNodesShrinkRequest) GetCliqueID() *string {
	return s.CliqueID
}

func (s *ListNodesShrinkRequest) GetDiskPL() *string {
	return s.DiskPL
}

func (s *ListNodesShrinkRequest) GetFilterByQuotaId() *string {
	return s.FilterByQuotaId
}

func (s *ListNodesShrinkRequest) GetFilterByResourceGroupIds() *string {
	return s.FilterByResourceGroupIds
}

func (s *ListNodesShrinkRequest) GetGPUType() *string {
	return s.GPUType
}

func (s *ListNodesShrinkRequest) GetHealthCountShrink() *string {
	return s.HealthCountShrink
}

func (s *ListNodesShrinkRequest) GetHealthRateShrink() *string {
	return s.HealthRateShrink
}

func (s *ListNodesShrinkRequest) GetHyperNode() *string {
	return s.HyperNode
}

func (s *ListNodesShrinkRequest) GetHyperZone() *string {
	return s.HyperZone
}

func (s *ListNodesShrinkRequest) GetLayoutMode() *string {
	return s.LayoutMode
}

func (s *ListNodesShrinkRequest) GetMachineGroupIds() *string {
	return s.MachineGroupIds
}

func (s *ListNodesShrinkRequest) GetNodeNames() *string {
	return s.NodeNames
}

func (s *ListNodesShrinkRequest) GetNodeStatuses() *string {
	return s.NodeStatuses
}

func (s *ListNodesShrinkRequest) GetNodeTypes() *string {
	return s.NodeTypes
}

func (s *ListNodesShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListNodesShrinkRequest) GetOrderInstanceIds() *string {
	return s.OrderInstanceIds
}

func (s *ListNodesShrinkRequest) GetOrderStatuses() *string {
	return s.OrderStatuses
}

func (s *ListNodesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNodesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNodesShrinkRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListNodesShrinkRequest) GetPodNum() *int32 {
	return s.PodNum
}

func (s *ListNodesShrinkRequest) GetQuotaId() *string {
	return s.QuotaId
}

func (s *ListNodesShrinkRequest) GetReasonCodes() *string {
	return s.ReasonCodes
}

func (s *ListNodesShrinkRequest) GetResourceGroupIds() *string {
	return s.ResourceGroupIds
}

func (s *ListNodesShrinkRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ListNodesShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListNodesShrinkRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *ListNodesShrinkRequest) GetWorkloadNum() *int32 {
	return s.WorkloadNum
}

func (s *ListNodesShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListNodesShrinkRequest) SetAcceleratorType(v string) *ListNodesShrinkRequest {
	s.AcceleratorType = &v
	return s
}

func (s *ListNodesShrinkRequest) SetAvailabilityZone(v string) *ListNodesShrinkRequest {
	s.AvailabilityZone = &v
	return s
}

func (s *ListNodesShrinkRequest) SetCliqueID(v string) *ListNodesShrinkRequest {
	s.CliqueID = &v
	return s
}

func (s *ListNodesShrinkRequest) SetDiskPL(v string) *ListNodesShrinkRequest {
	s.DiskPL = &v
	return s
}

func (s *ListNodesShrinkRequest) SetFilterByQuotaId(v string) *ListNodesShrinkRequest {
	s.FilterByQuotaId = &v
	return s
}

func (s *ListNodesShrinkRequest) SetFilterByResourceGroupIds(v string) *ListNodesShrinkRequest {
	s.FilterByResourceGroupIds = &v
	return s
}

func (s *ListNodesShrinkRequest) SetGPUType(v string) *ListNodesShrinkRequest {
	s.GPUType = &v
	return s
}

func (s *ListNodesShrinkRequest) SetHealthCountShrink(v string) *ListNodesShrinkRequest {
	s.HealthCountShrink = &v
	return s
}

func (s *ListNodesShrinkRequest) SetHealthRateShrink(v string) *ListNodesShrinkRequest {
	s.HealthRateShrink = &v
	return s
}

func (s *ListNodesShrinkRequest) SetHyperNode(v string) *ListNodesShrinkRequest {
	s.HyperNode = &v
	return s
}

func (s *ListNodesShrinkRequest) SetHyperZone(v string) *ListNodesShrinkRequest {
	s.HyperZone = &v
	return s
}

func (s *ListNodesShrinkRequest) SetLayoutMode(v string) *ListNodesShrinkRequest {
	s.LayoutMode = &v
	return s
}

func (s *ListNodesShrinkRequest) SetMachineGroupIds(v string) *ListNodesShrinkRequest {
	s.MachineGroupIds = &v
	return s
}

func (s *ListNodesShrinkRequest) SetNodeNames(v string) *ListNodesShrinkRequest {
	s.NodeNames = &v
	return s
}

func (s *ListNodesShrinkRequest) SetNodeStatuses(v string) *ListNodesShrinkRequest {
	s.NodeStatuses = &v
	return s
}

func (s *ListNodesShrinkRequest) SetNodeTypes(v string) *ListNodesShrinkRequest {
	s.NodeTypes = &v
	return s
}

func (s *ListNodesShrinkRequest) SetOrder(v string) *ListNodesShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListNodesShrinkRequest) SetOrderInstanceIds(v string) *ListNodesShrinkRequest {
	s.OrderInstanceIds = &v
	return s
}

func (s *ListNodesShrinkRequest) SetOrderStatuses(v string) *ListNodesShrinkRequest {
	s.OrderStatuses = &v
	return s
}

func (s *ListNodesShrinkRequest) SetPageNumber(v int32) *ListNodesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodesShrinkRequest) SetPageSize(v int32) *ListNodesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodesShrinkRequest) SetPaymentType(v string) *ListNodesShrinkRequest {
	s.PaymentType = &v
	return s
}

func (s *ListNodesShrinkRequest) SetPodNum(v int32) *ListNodesShrinkRequest {
	s.PodNum = &v
	return s
}

func (s *ListNodesShrinkRequest) SetQuotaId(v string) *ListNodesShrinkRequest {
	s.QuotaId = &v
	return s
}

func (s *ListNodesShrinkRequest) SetReasonCodes(v string) *ListNodesShrinkRequest {
	s.ReasonCodes = &v
	return s
}

func (s *ListNodesShrinkRequest) SetResourceGroupIds(v string) *ListNodesShrinkRequest {
	s.ResourceGroupIds = &v
	return s
}

func (s *ListNodesShrinkRequest) SetResourceGroupName(v string) *ListNodesShrinkRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *ListNodesShrinkRequest) SetSortBy(v string) *ListNodesShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListNodesShrinkRequest) SetVerbose(v bool) *ListNodesShrinkRequest {
	s.Verbose = &v
	return s
}

func (s *ListNodesShrinkRequest) SetWorkloadNum(v int32) *ListNodesShrinkRequest {
	s.WorkloadNum = &v
	return s
}

func (s *ListNodesShrinkRequest) SetWorkspaceId(v string) *ListNodesShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
