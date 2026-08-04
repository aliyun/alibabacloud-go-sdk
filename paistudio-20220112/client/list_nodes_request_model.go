// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorType(v string) *ListNodesRequest
	GetAcceleratorType() *string
	SetAvailabilityZone(v string) *ListNodesRequest
	GetAvailabilityZone() *string
	SetCliqueID(v string) *ListNodesRequest
	GetCliqueID() *string
	SetDiskPL(v string) *ListNodesRequest
	GetDiskPL() *string
	SetFilterByQuotaId(v string) *ListNodesRequest
	GetFilterByQuotaId() *string
	SetFilterByResourceGroupIds(v string) *ListNodesRequest
	GetFilterByResourceGroupIds() *string
	SetGPUType(v string) *ListNodesRequest
	GetGPUType() *string
	SetHealthCount(v *ListNodesRequestHealthCount) *ListNodesRequest
	GetHealthCount() *ListNodesRequestHealthCount
	SetHealthRate(v *ListNodesRequestHealthRate) *ListNodesRequest
	GetHealthRate() *ListNodesRequestHealthRate
	SetHyperNode(v string) *ListNodesRequest
	GetHyperNode() *string
	SetHyperZone(v string) *ListNodesRequest
	GetHyperZone() *string
	SetLayoutMode(v string) *ListNodesRequest
	GetLayoutMode() *string
	SetMachineGroupIds(v string) *ListNodesRequest
	GetMachineGroupIds() *string
	SetNodeNames(v string) *ListNodesRequest
	GetNodeNames() *string
	SetNodeStatuses(v string) *ListNodesRequest
	GetNodeStatuses() *string
	SetNodeTypes(v string) *ListNodesRequest
	GetNodeTypes() *string
	SetOrder(v string) *ListNodesRequest
	GetOrder() *string
	SetOrderInstanceIds(v string) *ListNodesRequest
	GetOrderInstanceIds() *string
	SetOrderStatuses(v string) *ListNodesRequest
	GetOrderStatuses() *string
	SetPageNumber(v int32) *ListNodesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNodesRequest
	GetPageSize() *int32
	SetPaymentType(v string) *ListNodesRequest
	GetPaymentType() *string
	SetPodNum(v int32) *ListNodesRequest
	GetPodNum() *int32
	SetQuotaId(v string) *ListNodesRequest
	GetQuotaId() *string
	SetReasonCodes(v string) *ListNodesRequest
	GetReasonCodes() *string
	SetResourceGroupIds(v string) *ListNodesRequest
	GetResourceGroupIds() *string
	SetResourceGroupName(v string) *ListNodesRequest
	GetResourceGroupName() *string
	SetSortBy(v string) *ListNodesRequest
	GetSortBy() *string
	SetVerbose(v bool) *ListNodesRequest
	GetVerbose() *bool
	SetWorkloadNum(v int32) *ListNodesRequest
	GetWorkloadNum() *int32
	SetWorkspaceId(v string) *ListNodesRequest
	GetWorkspaceId() *string
}

type ListNodesRequest struct {
	// The accelerator type. Valid values:
	//
	// - CPU
	//
	// - GPU
	//
	// If omitted, this operation returns nodes of all accelerator types.
	//
	// example:
	//
	// CPU
	AcceleratorType  *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	CliqueID         *string `json:"CliqueID,omitempty" xml:"CliqueID,omitempty"`
	DiskPL           *string `json:"DiskPL,omitempty" xml:"DiskPL,omitempty"`
	// When used with `ResourceGroupIds`, this parameter further filters the results to include only nodes from the specified resource quota.
	//
	// example:
	//
	// quotamtl37ge7gkvdz
	FilterByQuotaId *string `json:"FilterByQuotaId,omitempty" xml:"FilterByQuotaId,omitempty"`
	// When used with `QuotaId`, this parameter further filters the results to include only nodes from the specified resource groups.
	//
	// example:
	//
	// rg69rj0leslwdnbe
	FilterByResourceGroupIds *string `json:"FilterByResourceGroupIds,omitempty" xml:"FilterByResourceGroupIds,omitempty"`
	// The GPU type. Fuzzy matching is supported.
	//
	// example:
	//
	// T4
	GPUType         *string                      `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	HealthCount     *ListNodesRequestHealthCount `json:"HealthCount,omitempty" xml:"HealthCount,omitempty" type:"Struct"`
	HealthRate      *ListNodesRequestHealthRate  `json:"HealthRate,omitempty" xml:"HealthRate,omitempty" type:"Struct"`
	HyperNode       *string                      `json:"HyperNode,omitempty" xml:"HyperNode,omitempty"`
	HyperZone       *string                      `json:"HyperZone,omitempty" xml:"HyperZone,omitempty"`
	LayoutMode      *string                      `json:"LayoutMode,omitempty" xml:"LayoutMode,omitempty"`
	MachineGroupIds *string                      `json:"MachineGroupIds,omitempty" xml:"MachineGroupIds,omitempty"`
	// A comma-separated list of node names. Only nodes with names that match this list are returned.
	//
	// example:
	//
	// lingjxxxx
	NodeNames *string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty"`
	// A comma-separated list of node statuses. If this parameter is omitted, this operation returns nodes of all statuses.
	//
	// example:
	//
	// Ready
	NodeStatuses *string `json:"NodeStatuses,omitempty" xml:"NodeStatuses,omitempty"`
	// A comma-separated list of node specifications. If this parameter is omitted, this operation returns nodes of all specifications.
	//
	// example:
	//
	// ecs.c6.xlarge
	NodeTypes *string `json:"NodeTypes,omitempty" xml:"NodeTypes,omitempty"`
	// The sort order. Valid values:
	//
	// - `desc`: Descending
	//
	// - `asc`: Ascending
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// A comma-separated list of order IDs.
	//
	// example:
	//
	// 260590501560397
	OrderInstanceIds *string `json:"OrderInstanceIds,omitempty" xml:"OrderInstanceIds,omitempty"`
	// A comma-separated list of order statuses.
	//
	// example:
	//
	// Ready
	OrderStatuses *string `json:"OrderStatuses,omitempty" xml:"OrderStatuses,omitempty"`
	// The page number. The first page is 1.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page.
	//
	// example:
	//
	// 10
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	PodNum      *int32  `json:"PodNum,omitempty" xml:"PodNum,omitempty"`
	// The ID of the resource quota that contains the nodes.
	//
	// example:
	//
	// quotamtl37ge7gkvdz
	QuotaId     *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	ReasonCodes *string `json:"ReasonCodes,omitempty" xml:"ReasonCodes,omitempty"`
	// A comma-separated list of resource group IDs. You must specify either this parameter or `QuotaId`.
	//
	// Constraints:
	//
	// 1. The user ID of the request must match the user ID associated with the specified resource groups.
	//
	// 2. All specified resource groups must be of the same type.
	//
	// 3. All specified resource groups must be in the same VPC.
	//
	// example:
	//
	// rg69rj0leslwdnbe
	ResourceGroupIds  *string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty"`
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The field by which to sort the results.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// Specifies whether to return resource usage information. This parameter applies only when `QuotaId` is specified.
	//
	// example:
	//
	// false
	Verbose     *bool   `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	WorkloadNum *int32  `json:"WorkloadNum,omitempty" xml:"WorkloadNum,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) GetAcceleratorType() *string {
	return s.AcceleratorType
}

func (s *ListNodesRequest) GetAvailabilityZone() *string {
	return s.AvailabilityZone
}

func (s *ListNodesRequest) GetCliqueID() *string {
	return s.CliqueID
}

func (s *ListNodesRequest) GetDiskPL() *string {
	return s.DiskPL
}

func (s *ListNodesRequest) GetFilterByQuotaId() *string {
	return s.FilterByQuotaId
}

func (s *ListNodesRequest) GetFilterByResourceGroupIds() *string {
	return s.FilterByResourceGroupIds
}

func (s *ListNodesRequest) GetGPUType() *string {
	return s.GPUType
}

func (s *ListNodesRequest) GetHealthCount() *ListNodesRequestHealthCount {
	return s.HealthCount
}

func (s *ListNodesRequest) GetHealthRate() *ListNodesRequestHealthRate {
	return s.HealthRate
}

func (s *ListNodesRequest) GetHyperNode() *string {
	return s.HyperNode
}

func (s *ListNodesRequest) GetHyperZone() *string {
	return s.HyperZone
}

func (s *ListNodesRequest) GetLayoutMode() *string {
	return s.LayoutMode
}

func (s *ListNodesRequest) GetMachineGroupIds() *string {
	return s.MachineGroupIds
}

func (s *ListNodesRequest) GetNodeNames() *string {
	return s.NodeNames
}

func (s *ListNodesRequest) GetNodeStatuses() *string {
	return s.NodeStatuses
}

func (s *ListNodesRequest) GetNodeTypes() *string {
	return s.NodeTypes
}

func (s *ListNodesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListNodesRequest) GetOrderInstanceIds() *string {
	return s.OrderInstanceIds
}

func (s *ListNodesRequest) GetOrderStatuses() *string {
	return s.OrderStatuses
}

func (s *ListNodesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNodesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNodesRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListNodesRequest) GetPodNum() *int32 {
	return s.PodNum
}

func (s *ListNodesRequest) GetQuotaId() *string {
	return s.QuotaId
}

func (s *ListNodesRequest) GetReasonCodes() *string {
	return s.ReasonCodes
}

func (s *ListNodesRequest) GetResourceGroupIds() *string {
	return s.ResourceGroupIds
}

func (s *ListNodesRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ListNodesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListNodesRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *ListNodesRequest) GetWorkloadNum() *int32 {
	return s.WorkloadNum
}

func (s *ListNodesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListNodesRequest) SetAcceleratorType(v string) *ListNodesRequest {
	s.AcceleratorType = &v
	return s
}

func (s *ListNodesRequest) SetAvailabilityZone(v string) *ListNodesRequest {
	s.AvailabilityZone = &v
	return s
}

func (s *ListNodesRequest) SetCliqueID(v string) *ListNodesRequest {
	s.CliqueID = &v
	return s
}

func (s *ListNodesRequest) SetDiskPL(v string) *ListNodesRequest {
	s.DiskPL = &v
	return s
}

func (s *ListNodesRequest) SetFilterByQuotaId(v string) *ListNodesRequest {
	s.FilterByQuotaId = &v
	return s
}

func (s *ListNodesRequest) SetFilterByResourceGroupIds(v string) *ListNodesRequest {
	s.FilterByResourceGroupIds = &v
	return s
}

func (s *ListNodesRequest) SetGPUType(v string) *ListNodesRequest {
	s.GPUType = &v
	return s
}

func (s *ListNodesRequest) SetHealthCount(v *ListNodesRequestHealthCount) *ListNodesRequest {
	s.HealthCount = v
	return s
}

func (s *ListNodesRequest) SetHealthRate(v *ListNodesRequestHealthRate) *ListNodesRequest {
	s.HealthRate = v
	return s
}

func (s *ListNodesRequest) SetHyperNode(v string) *ListNodesRequest {
	s.HyperNode = &v
	return s
}

func (s *ListNodesRequest) SetHyperZone(v string) *ListNodesRequest {
	s.HyperZone = &v
	return s
}

func (s *ListNodesRequest) SetLayoutMode(v string) *ListNodesRequest {
	s.LayoutMode = &v
	return s
}

func (s *ListNodesRequest) SetMachineGroupIds(v string) *ListNodesRequest {
	s.MachineGroupIds = &v
	return s
}

func (s *ListNodesRequest) SetNodeNames(v string) *ListNodesRequest {
	s.NodeNames = &v
	return s
}

func (s *ListNodesRequest) SetNodeStatuses(v string) *ListNodesRequest {
	s.NodeStatuses = &v
	return s
}

func (s *ListNodesRequest) SetNodeTypes(v string) *ListNodesRequest {
	s.NodeTypes = &v
	return s
}

func (s *ListNodesRequest) SetOrder(v string) *ListNodesRequest {
	s.Order = &v
	return s
}

func (s *ListNodesRequest) SetOrderInstanceIds(v string) *ListNodesRequest {
	s.OrderInstanceIds = &v
	return s
}

func (s *ListNodesRequest) SetOrderStatuses(v string) *ListNodesRequest {
	s.OrderStatuses = &v
	return s
}

func (s *ListNodesRequest) SetPageNumber(v int32) *ListNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodesRequest) SetPageSize(v int32) *ListNodesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodesRequest) SetPaymentType(v string) *ListNodesRequest {
	s.PaymentType = &v
	return s
}

func (s *ListNodesRequest) SetPodNum(v int32) *ListNodesRequest {
	s.PodNum = &v
	return s
}

func (s *ListNodesRequest) SetQuotaId(v string) *ListNodesRequest {
	s.QuotaId = &v
	return s
}

func (s *ListNodesRequest) SetReasonCodes(v string) *ListNodesRequest {
	s.ReasonCodes = &v
	return s
}

func (s *ListNodesRequest) SetResourceGroupIds(v string) *ListNodesRequest {
	s.ResourceGroupIds = &v
	return s
}

func (s *ListNodesRequest) SetResourceGroupName(v string) *ListNodesRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *ListNodesRequest) SetSortBy(v string) *ListNodesRequest {
	s.SortBy = &v
	return s
}

func (s *ListNodesRequest) SetVerbose(v bool) *ListNodesRequest {
	s.Verbose = &v
	return s
}

func (s *ListNodesRequest) SetWorkloadNum(v int32) *ListNodesRequest {
	s.WorkloadNum = &v
	return s
}

func (s *ListNodesRequest) SetWorkspaceId(v string) *ListNodesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListNodesRequest) Validate() error {
	if s.HealthCount != nil {
		if err := s.HealthCount.Validate(); err != nil {
			return err
		}
	}
	if s.HealthRate != nil {
		if err := s.HealthRate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNodesRequestHealthCount struct {
	Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
	Value     *int32  `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListNodesRequestHealthCount) String() string {
	return dara.Prettify(s)
}

func (s ListNodesRequestHealthCount) GoString() string {
	return s.String()
}

func (s *ListNodesRequestHealthCount) GetOperation() *string {
	return s.Operation
}

func (s *ListNodesRequestHealthCount) GetValue() *int32 {
	return s.Value
}

func (s *ListNodesRequestHealthCount) SetOperation(v string) *ListNodesRequestHealthCount {
	s.Operation = &v
	return s
}

func (s *ListNodesRequestHealthCount) SetValue(v int32) *ListNodesRequestHealthCount {
	s.Value = &v
	return s
}

func (s *ListNodesRequestHealthCount) Validate() error {
	return dara.Validate(s)
}

type ListNodesRequestHealthRate struct {
	Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
	Value     *int32  `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListNodesRequestHealthRate) String() string {
	return dara.Prettify(s)
}

func (s ListNodesRequestHealthRate) GoString() string {
	return s.String()
}

func (s *ListNodesRequestHealthRate) GetOperation() *string {
	return s.Operation
}

func (s *ListNodesRequestHealthRate) GetValue() *int32 {
	return s.Value
}

func (s *ListNodesRequestHealthRate) SetOperation(v string) *ListNodesRequestHealthRate {
	s.Operation = &v
	return s
}

func (s *ListNodesRequestHealthRate) SetValue(v int32) *ListNodesRequestHealthRate {
	s.Value = &v
	return s
}

func (s *ListNodesRequestHealthRate) Validate() error {
	return dara.Validate(s)
}
