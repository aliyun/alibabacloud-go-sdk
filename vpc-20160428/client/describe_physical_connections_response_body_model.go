// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhysicalConnectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribePhysicalConnectionsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePhysicalConnectionsResponseBody
	GetPageSize() *int32
	SetPhysicalConnectionSet(v *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet) *DescribePhysicalConnectionsResponseBody
	GetPhysicalConnectionSet() *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet
	SetRequestId(v string) *DescribePhysicalConnectionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribePhysicalConnectionsResponseBody
	GetTotalCount() *int32
}

type DescribePhysicalConnectionsResponseBody struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 10
	PageSize              *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PhysicalConnectionSet *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet `json:"PhysicalConnectionSet,omitempty" xml:"PhysicalConnectionSet,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0E6D0EC4-7C91-53E2-9F65-64BF713114B0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePhysicalConnectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePhysicalConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhysicalConnectionsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePhysicalConnectionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePhysicalConnectionsResponseBody) GetPhysicalConnectionSet() *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet {
	return s.PhysicalConnectionSet
}

func (s *DescribePhysicalConnectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePhysicalConnectionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePhysicalConnectionsResponseBody) SetPageNumber(v int32) *DescribePhysicalConnectionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBody) SetPageSize(v int32) *DescribePhysicalConnectionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBody) SetPhysicalConnectionSet(v *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet) *DescribePhysicalConnectionsResponseBody {
	s.PhysicalConnectionSet = v
	return s
}

func (s *DescribePhysicalConnectionsResponseBody) SetRequestId(v string) *DescribePhysicalConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBody) SetTotalCount(v int32) *DescribePhysicalConnectionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBody) Validate() error {
	if s.PhysicalConnectionSet != nil {
		if err := s.PhysicalConnectionSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet struct {
	PhysicalConnectionType []*DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType `json:"PhysicalConnectionType,omitempty" xml:"PhysicalConnectionType,omitempty" type:"Repeated"`
}

func (s DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet) String() string {
	return dara.Prettify(s)
}

func (s DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet) GoString() string {
	return s.String()
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet) GetPhysicalConnectionType() []*DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	return s.PhysicalConnectionType
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet) SetPhysicalConnectionType(v []*DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet {
	s.PhysicalConnectionType = v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet) Validate() error {
	if s.PhysicalConnectionType != nil {
		for _, item := range s.PhysicalConnectionType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType struct {
	AccessPointId                  *string                                                                                 `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	AccessPointType                *string                                                                                 `json:"AccessPointType,omitempty" xml:"AccessPointType,omitempty"`
	AdDetailLocation               *string                                                                                 `json:"AdDetailLocation,omitempty" xml:"AdDetailLocation,omitempty"`
	AdLocation                     *string                                                                                 `json:"AdLocation,omitempty" xml:"AdLocation,omitempty"`
	Bandwidth                      *int64                                                                                  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BusinessStatus                 *string                                                                                 `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	ChargeType                     *string                                                                                 `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CircuitCode                    *string                                                                                 `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	CreationTime                   *string                                                                                 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description                    *string                                                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	EnabledTime                    *string                                                                                 `json:"EnabledTime,omitempty" xml:"EnabledTime,omitempty"`
	EndTime                        *string                                                                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExpectSpec                     *string                                                                                 `json:"ExpectSpec,omitempty" xml:"ExpectSpec,omitempty"`
	HasReservationData             *string                                                                                 `json:"HasReservationData,omitempty" xml:"HasReservationData,omitempty"`
	LineOperator                   *string                                                                                 `json:"LineOperator,omitempty" xml:"LineOperator,omitempty"`
	LoaStatus                      *string                                                                                 `json:"LoaStatus,omitempty" xml:"LoaStatus,omitempty"`
	Name                           *string                                                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	OpticalModuleModel             *string                                                                                 `json:"OpticalModuleModel,omitempty" xml:"OpticalModuleModel,omitempty"`
	OrderMode                      *string                                                                                 `json:"OrderMode,omitempty" xml:"OrderMode,omitempty"`
	ParentPhysicalConnectionAliUid *int64                                                                                  `json:"ParentPhysicalConnectionAliUid,omitempty" xml:"ParentPhysicalConnectionAliUid,omitempty"`
	ParentPhysicalConnectionId     *string                                                                                 `json:"ParentPhysicalConnectionId,omitempty" xml:"ParentPhysicalConnectionId,omitempty"`
	PeerLocation                   *string                                                                                 `json:"PeerLocation,omitempty" xml:"PeerLocation,omitempty"`
	PhysicalConnectionId           *string                                                                                 `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	PortNumber                     *string                                                                                 `json:"PortNumber,omitempty" xml:"PortNumber,omitempty"`
	PortType                       *string                                                                                 `json:"PortType,omitempty" xml:"PortType,omitempty"`
	ProductType                    *string                                                                                 `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	QosId                          *string                                                                                 `json:"QosId,omitempty" xml:"QosId,omitempty"`
	RedundantPhysicalConnectionId  *string                                                                                 `json:"RedundantPhysicalConnectionId,omitempty" xml:"RedundantPhysicalConnectionId,omitempty"`
	ReservationActiveTime          *string                                                                                 `json:"ReservationActiveTime,omitempty" xml:"ReservationActiveTime,omitempty"`
	ReservationInternetChargeType  *string                                                                                 `json:"ReservationInternetChargeType,omitempty" xml:"ReservationInternetChargeType,omitempty"`
	ReservationOrderType           *string                                                                                 `json:"ReservationOrderType,omitempty" xml:"ReservationOrderType,omitempty"`
	ResourceGroupId                *string                                                                                 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Spec                           *string                                                                                 `json:"Spec,omitempty" xml:"Spec,omitempty"`
	Status                         *string                                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                           *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Type                           *string                                                                                 `json:"Type,omitempty" xml:"Type,omitempty"`
	VirtualPhysicalConnectionCount *int32                                                                                  `json:"VirtualPhysicalConnectionCount,omitempty" xml:"VirtualPhysicalConnectionCount,omitempty"`
	VlanId                         *string                                                                                 `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
	VpconnStatus                   *string                                                                                 `json:"VpconnStatus,omitempty" xml:"VpconnStatus,omitempty"`
}

func (s DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) String() string {
	return dara.Prettify(s)
}

func (s DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GoString() string {
	return s.String()
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetAccessPointType() *string {
	return s.AccessPointType
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetAdDetailLocation() *string {
	return s.AdDetailLocation
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetAdLocation() *string {
	return s.AdLocation
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetCircuitCode() *string {
	return s.CircuitCode
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetDescription() *string {
	return s.Description
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetEnabledTime() *string {
	return s.EnabledTime
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetExpectSpec() *string {
	return s.ExpectSpec
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetHasReservationData() *string {
	return s.HasReservationData
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetLineOperator() *string {
	return s.LineOperator
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetLoaStatus() *string {
	return s.LoaStatus
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetName() *string {
	return s.Name
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetOpticalModuleModel() *string {
	return s.OpticalModuleModel
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetOrderMode() *string {
	return s.OrderMode
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetParentPhysicalConnectionAliUid() *int64 {
	return s.ParentPhysicalConnectionAliUid
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetParentPhysicalConnectionId() *string {
	return s.ParentPhysicalConnectionId
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetPeerLocation() *string {
	return s.PeerLocation
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetPortNumber() *string {
	return s.PortNumber
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetPortType() *string {
	return s.PortType
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetProductType() *string {
	return s.ProductType
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetQosId() *string {
	return s.QosId
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetRedundantPhysicalConnectionId() *string {
	return s.RedundantPhysicalConnectionId
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetReservationActiveTime() *string {
	return s.ReservationActiveTime
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetReservationInternetChargeType() *string {
	return s.ReservationInternetChargeType
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetReservationOrderType() *string {
	return s.ReservationOrderType
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetSpec() *string {
	return s.Spec
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetStatus() *string {
	return s.Status
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetTags() *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTags {
	return s.Tags
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetType() *string {
	return s.Type
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetVirtualPhysicalConnectionCount() *int32 {
	return s.VirtualPhysicalConnectionCount
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetVlanId() *string {
	return s.VlanId
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetVpconnStatus() *string {
	return s.VpconnStatus
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetAccessPointId(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.AccessPointId = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetAccessPointType(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.AccessPointType = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetAdDetailLocation(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.AdDetailLocation = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetAdLocation(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.AdLocation = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetBandwidth(v int64) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.Bandwidth = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetBusinessStatus(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.BusinessStatus = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetChargeType(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.ChargeType = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetCircuitCode(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.CircuitCode = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetCreationTime(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.CreationTime = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetDescription(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.Description = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetEnabledTime(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.EnabledTime = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetEndTime(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.EndTime = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetExpectSpec(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.ExpectSpec = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetHasReservationData(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.HasReservationData = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetLineOperator(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.LineOperator = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetLoaStatus(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.LoaStatus = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetName(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.Name = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetOpticalModuleModel(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.OpticalModuleModel = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetOrderMode(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.OrderMode = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetParentPhysicalConnectionAliUid(v int64) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.ParentPhysicalConnectionAliUid = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetParentPhysicalConnectionId(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.ParentPhysicalConnectionId = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetPeerLocation(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.PeerLocation = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetPhysicalConnectionId(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.PhysicalConnectionId = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetPortNumber(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.PortNumber = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetPortType(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.PortType = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetProductType(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.ProductType = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetQosId(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.QosId = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetRedundantPhysicalConnectionId(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.RedundantPhysicalConnectionId = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetReservationActiveTime(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.ReservationActiveTime = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetReservationInternetChargeType(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.ReservationInternetChargeType = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetReservationOrderType(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.ReservationOrderType = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetResourceGroupId(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetSpec(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.Spec = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetStatus(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.Status = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetTags(v *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTags) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.Tags = v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetType(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.Type = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetVirtualPhysicalConnectionCount(v int32) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.VirtualPhysicalConnectionCount = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetVlanId(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.VlanId = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetVpconnStatus(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.VpconnStatus = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTags struct {
	Tags []*DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTagsTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTags) String() string {
	return dara.Prettify(s)
}

func (s DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTags) GoString() string {
	return s.String()
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTags) GetTags() []*DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTagsTags {
	return s.Tags
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTags) SetTags(v []*DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTagsTags) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTags {
	s.Tags = v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTags) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTagsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTagsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTagsTags) GoString() string {
	return s.String()
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTagsTags) GetKey() *string {
	return s.Key
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTagsTags) GetValue() *string {
	return s.Value
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTagsTags) SetKey(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTagsTags {
	s.Key = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTagsTags) SetValue(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTagsTags {
	s.Value = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTagsTags) Validate() error {
	return dara.Validate(s)
}
