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
	SetFilterByQuotaId(v string) *ListNodesRequest
	GetFilterByQuotaId() *string
	SetFilterByResourceGroupIds(v string) *ListNodesRequest
	GetFilterByResourceGroupIds() *string
	SetGPUType(v string) *ListNodesRequest
	GetGPUType() *string
	SetHyperZone(v string) *ListNodesRequest
	GetHyperZone() *string
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
}

type ListNodesRequest struct {
	// example:
	//
	// CPU
	AcceleratorType  *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
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
	GPUType         *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	HyperZone       *string `json:"HyperZone,omitempty" xml:"HyperZone,omitempty"`
	MachineGroupIds *string `json:"MachineGroupIds,omitempty" xml:"MachineGroupIds,omitempty"`
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
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
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

func (s *ListNodesRequest) GetFilterByQuotaId() *string {
	return s.FilterByQuotaId
}

func (s *ListNodesRequest) GetFilterByResourceGroupIds() *string {
	return s.FilterByResourceGroupIds
}

func (s *ListNodesRequest) GetGPUType() *string {
	return s.GPUType
}

func (s *ListNodesRequest) GetHyperZone() *string {
	return s.HyperZone
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

func (s *ListNodesRequest) SetAcceleratorType(v string) *ListNodesRequest {
	s.AcceleratorType = &v
	return s
}

func (s *ListNodesRequest) SetAvailabilityZone(v string) *ListNodesRequest {
	s.AvailabilityZone = &v
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

func (s *ListNodesRequest) SetHyperZone(v string) *ListNodesRequest {
	s.HyperZone = &v
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

func (s *ListNodesRequest) Validate() error {
	return dara.Validate(s)
}
