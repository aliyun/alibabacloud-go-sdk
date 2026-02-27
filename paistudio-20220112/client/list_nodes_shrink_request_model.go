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
	// example:
	//
	// CPU
	AcceleratorType  *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	CliqueID         *string `json:"CliqueID,omitempty" xml:"CliqueID,omitempty"`
	DiskPL           *string `json:"DiskPL,omitempty" xml:"DiskPL,omitempty"`
	// example:
	//
	// quotamtl37ge7gkvdz
	FilterByQuotaId *string `json:"FilterByQuotaId,omitempty" xml:"FilterByQuotaId,omitempty"`
	// example:
	//
	// rg69rj0leslwdnbe
	FilterByResourceGroupIds *string `json:"FilterByResourceGroupIds,omitempty" xml:"FilterByResourceGroupIds,omitempty"`
	// example:
	//
	// T4
	GPUType           *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	HealthCountShrink *string `json:"HealthCount,omitempty" xml:"HealthCount,omitempty"`
	HealthRateShrink  *string `json:"HealthRate,omitempty" xml:"HealthRate,omitempty"`
	HyperNode         *string `json:"HyperNode,omitempty" xml:"HyperNode,omitempty"`
	HyperZone         *string `json:"HyperZone,omitempty" xml:"HyperZone,omitempty"`
	LayoutMode        *string `json:"LayoutMode,omitempty" xml:"LayoutMode,omitempty"`
	MachineGroupIds   *string `json:"MachineGroupIds,omitempty" xml:"MachineGroupIds,omitempty"`
	// example:
	//
	// lingjxxxx
	NodeNames *string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty"`
	// example:
	//
	// Ready
	NodeStatuses *string `json:"NodeStatuses,omitempty" xml:"NodeStatuses,omitempty"`
	// example:
	//
	// ecs.c6.xlarge
	NodeTypes *string `json:"NodeTypes,omitempty" xml:"NodeTypes,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 260590501560397
	OrderInstanceIds *string `json:"OrderInstanceIds,omitempty" xml:"OrderInstanceIds,omitempty"`
	// example:
	//
	// Ready
	OrderStatuses *string `json:"OrderStatuses,omitempty" xml:"OrderStatuses,omitempty"`
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	PodNum      *int32  `json:"PodNum,omitempty" xml:"PodNum,omitempty"`
	// example:
	//
	// quotamtl37ge7gkvdz
	QuotaId     *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	ReasonCodes *string `json:"ReasonCodes,omitempty" xml:"ReasonCodes,omitempty"`
	// example:
	//
	// rg69rj0leslwdnbe
	ResourceGroupIds  *string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty"`
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// false
	Verbose     *bool   `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	WorkloadNum *int32  `json:"WorkloadNum,omitempty" xml:"WorkloadNum,omitempty"`
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
