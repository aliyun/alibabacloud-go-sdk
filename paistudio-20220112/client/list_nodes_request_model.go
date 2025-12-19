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
	SetQuotaId(v string) *ListNodesRequest
	GetQuotaId() *string
	SetReasonCodes(v string) *ListNodesRequest
	GetReasonCodes() *string
	SetResourceGroupIds(v string) *ListNodesRequest
	GetResourceGroupIds() *string
	SetSortBy(v string) *ListNodesRequest
	GetSortBy() *string
	SetVerbose(v bool) *ListNodesRequest
	GetVerbose() *bool
	SetWorkspaceId(v string) *ListNodesRequest
	GetWorkspaceId() *string
}

type ListNodesRequest struct {
	// example:
	//
	// CPU
	AcceleratorType  *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	CliqueID         *string `json:"CliqueID,omitempty" xml:"CliqueID,omitempty"`
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
	GPUType         *string                      `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	HealthCount     *ListNodesRequestHealthCount `json:"HealthCount,omitempty" xml:"HealthCount,omitempty" type:"Struct"`
	HealthRate      *ListNodesRequestHealthRate  `json:"HealthRate,omitempty" xml:"HealthRate,omitempty" type:"Struct"`
	HyperNode       *string                      `json:"HyperNode,omitempty" xml:"HyperNode,omitempty"`
	HyperZone       *string                      `json:"HyperZone,omitempty" xml:"HyperZone,omitempty"`
	LayoutMode      *string                      `json:"LayoutMode,omitempty" xml:"LayoutMode,omitempty"`
	MachineGroupIds *string                      `json:"MachineGroupIds,omitempty" xml:"MachineGroupIds,omitempty"`
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
	// example:
	//
	// quotamtl37ge7gkvdz
	QuotaId     *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	ReasonCodes *string `json:"ReasonCodes,omitempty" xml:"ReasonCodes,omitempty"`
	// example:
	//
	// rg69rj0leslwdnbe
	ResourceGroupIds *string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// false
	Verbose     *bool   `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
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

func (s *ListNodesRequest) GetQuotaId() *string {
	return s.QuotaId
}

func (s *ListNodesRequest) GetReasonCodes() *string {
	return s.ReasonCodes
}

func (s *ListNodesRequest) GetResourceGroupIds() *string {
	return s.ResourceGroupIds
}

func (s *ListNodesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListNodesRequest) GetVerbose() *bool {
	return s.Verbose
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

func (s *ListNodesRequest) SetSortBy(v string) *ListNodesRequest {
	s.SortBy = &v
	return s
}

func (s *ListNodesRequest) SetVerbose(v bool) *ListNodesRequest {
	s.Verbose = &v
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
