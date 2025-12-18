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
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of Express Connect circuits.
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
	// The ID of the Express Connect circuit.
	//
	// example:
	//
	// ap-cn-hangzhou-finance-yh-E
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// The type of the access point.
	//
	// example:
	//
	// VPC
	AccessPointType *string `json:"AccessPointType,omitempty" xml:"AccessPointType,omitempty"`
	// The information about the data center and rack.
	//
	// example:
	//
	// Position 30, Server Rack JXX, Booth ET135ET135-XX-2, Room XX, Building 10, XX Road, XX Town, XX District, Hangzhou, Zhejiang Province
	AdDetailLocation *string `json:"AdDetailLocation,omitempty" xml:"AdDetailLocation,omitempty"`
	// The location of the access point.
	//
	// example:
	//
	// Number 10, XX Road, XX Town, XX District, Hangzhou City, Zhejiang Province.
	AdLocation *string `json:"AdLocation,omitempty" xml:"AdLocation,omitempty"`
	// The maximum bandwidth of the Express Connect circuit.
	//
	// Unit: Mbit/s.
	//
	// example:
	//
	// 10
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The status of the Express Connect circuit. Valid values:
	//
	// 	- **Normal**: enabled
	//
	// 	- **FinancialLocked**: locked due to overdue payments
	//
	// 	- **SecurityLocked**: locked for security reasons
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The billing method of the Express Connect circuit.
	//
	// If **Prepaid*	- is returned, it indicates that the Express Connect circuit is billed on a subscription basis.
	//
	// example:
	//
	// Prepaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The circuit code of the Express Connect circuit. The circuit code is provided by the connectivity provider.
	//
	// example:
	//
	// longtel001
	CircuitCode *string `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	// The time when the Express Connect circuit was created.
	//
	// example:
	//
	// 2021-08-24T07:30:58Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the Express Connect circuit.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the Express Connect circuit was enabled.
	//
	// example:
	//
	// 2021-08-24T07:33:18Z
	EnabledTime *string `json:"EnabledTime,omitempty" xml:"EnabledTime,omitempty"`
	// The time when the Express Connect circuit expires.
	//
	// example:
	//
	// 2022-04-24T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The estimated maximum bandwidth of the shared Express Connect circuit. The estimated bandwidth takes effect after you complete the payment.
	//
	// Unit: **M*	- (Mbit/s) and **G*	- (Gbit/s).
	//
	// example:
	//
	// 50M
	ExpectSpec *string `json:"ExpectSpec,omitempty" xml:"ExpectSpec,omitempty"`
	// Indicates whether the data about pending orders is returned. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	HasReservationData *string `json:"HasReservationData,omitempty" xml:"HasReservationData,omitempty"`
	// The connectivity provider of the Express Connect circuit. Valid values:
	//
	// 	- **CT**: China Telecom.
	//
	// 	- **CU**: China Unicom.
	//
	// 	- **CM**: China Mobile.
	//
	// 	- **CO**: other connectivity providers in the Chinese mainland.
	//
	// 	- **Equinix**: Equinix.
	//
	// 	- **Other**: other connectivity providers outside the Chinese mainland.
	//
	// example:
	//
	// CT
	LineOperator *string `json:"LineOperator,omitempty" xml:"LineOperator,omitempty"`
	// The status of the letter of authorization (LOA). Valid values:
	//
	// 	- **Applying**: The LOA is pending for approval.
	//
	// 	- **Accept**: The LOA is approved.
	//
	// 	- **Available**: The LOA is available.
	//
	// 	- **Rejected**: The LOA is rejected.
	//
	// 	- **Completing**: The Express Connect circuit is under construction.
	//
	// 	- **Complete**: The Express Connect circuit is installed.
	//
	// 	- **Deleted**: The LOA is deleted.
	//
	// example:
	//
	// Available
	LoaStatus *string `json:"LoaStatus,omitempty" xml:"LoaStatus,omitempty"`
	// The name of the Express Connect circuit.
	//
	// example:
	//
	// nametest
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OpticalModuleModel *string `json:"OpticalModuleModel,omitempty" xml:"OpticalModuleModel,omitempty"`
	// The payer for the hosted connection. Valid values:
	//
	// 	- **PayByPhysicalConnectionOwner**: The partner pays for the shared Express Connect circuit.
	//
	// 	- **PayByVirtualPhysicalConnectionOwner**: The tenant pays for the shared Express Connect circuit.
	//
	// example:
	//
	// PayByPhysicalConnectionOwner
	OrderMode *string `json:"OrderMode,omitempty" xml:"OrderMode,omitempty"`
	// The ID of the Alibaba Cloud account to which the parent Express Connect circuit belongs.
	//
	// example:
	//
	// 283117732402483989
	ParentPhysicalConnectionAliUid *int64 `json:"ParentPhysicalConnectionAliUid,omitempty" xml:"ParentPhysicalConnectionAliUid,omitempty"`
	// The ID of the parent Express Connect circuit.
	//
	// example:
	//
	// pc-bp1ciz7ekd2grn1as****
	ParentPhysicalConnectionId *string `json:"ParentPhysicalConnectionId,omitempty" xml:"ParentPhysicalConnectionId,omitempty"`
	// The geographical location of the data center.
	//
	// example:
	//
	// XX Number, XX Road, XX Town, XX District, Hangzhou City, Zhejiang Province.
	PeerLocation *string `json:"PeerLocation,omitempty" xml:"PeerLocation,omitempty"`
	// The ID of the Express Connect circuit.
	//
	// example:
	//
	// pc-bp1ciz7ekd2grn1as****
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// The ID of the port on the access device.
	//
	// example:
	//
	// 1/1/1
	PortNumber *string `json:"PortNumber,omitempty" xml:"PortNumber,omitempty"`
	// The port type of the Express Connect circuit. Valid values:
	//
	// 	- **100Base-T**: 100 Mbit/s copper Ethernet port
	//
	// 	- **1000Base-T**: 1,000 Mbit/s copper Ethernet port
	//
	// 	- **1000Base-LX**: 1,000 Mbit/s single-mode optical port (10 km)
	//
	// 	- **10GBase-T**: 10,000 Mbit/s copper Ethernet port
	//
	// 	- **10GBase-LR**: 10,000 Mbit/s single-mode optical port (10 km)
	//
	// 	- **40GBase-LR**: 40,000 Mbit/s single-mode optical port
	//
	// 	- **100GBase-LR**: 100,000 Mbit/s single-mode optical port
	//
	// > Whether 40GBase-LR and 100GBase-LR ports can be created depends on resource supplies. For more information, contact your account manager.
	//
	// example:
	//
	// 10GBase-LR
	PortType *string `json:"PortType,omitempty" xml:"PortType,omitempty"`
	// The type of the Express Connect circuit. Valid values:
	//
	// 	- **VirtualPhysicalConnection**: shared Express Connect circuit
	//
	// 	- **PhysicalConnection**: dedicated Express Connect circuit
	//
	// example:
	//
	// PhysicalConnection
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The ID of the QoS policy.
	//
	// example:
	//
	// qos-bp10s3szn8rgnxuw7****
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The ID of the standby Express Connect circuit.
	//
	// example:
	//
	// pc-119mfjzm****
	RedundantPhysicalConnectionId *string `json:"RedundantPhysicalConnectionId,omitempty" xml:"RedundantPhysicalConnectionId,omitempty"`
	// The time when the pending order takes effect.
	//
	// example:
	//
	// 2022-02-25T11:01:04Z
	ReservationActiveTime *string `json:"ReservationActiveTime,omitempty" xml:"ReservationActiveTime,omitempty"`
	// The billing method of the pending order.
	//
	// If **PayByBandwidth*	- is returned, it indicates that the Express Connect circuit is billed on a pay-by-bandwidth basis.
	//
	// example:
	//
	// PayByBandwidth
	ReservationInternetChargeType *string `json:"ReservationInternetChargeType,omitempty" xml:"ReservationInternetChargeType,omitempty"`
	// The type of the pending order.
	//
	// If the value is **RENEW**, it indicates that the order is placed for service renewal.
	//
	// example:
	//
	// RENEW
	ReservationOrderType *string `json:"ReservationOrderType,omitempty" xml:"ReservationOrderType,omitempty"`
	// The resource group ID to which the instance belongs.
	//
	// example:
	//
	// rg-acfmwu3k52prgdi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The specification of the Express Connect circuit.
	//
	// Unit: **G*	- (Gbit/s).
	//
	// example:
	//
	// 10G
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The status of the Express Connect circuit. Valid values:
	//
	// 	- **Initial**
	//
	// 	- **Approved**
	//
	// 	- **Allocating**
	//
	// 	- **Allocated**
	//
	// 	- **Confirmed**
	//
	// 	- **Enabled**
	//
	// 	- **Rejected**
	//
	// 	- **Canceled**
	//
	// 	- **Allocation Failed**
	//
	// 	- **Terminating**
	//
	// 	- **Terminated**
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the cluster.
	Tags *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionTypeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The type of resource to which the Express Connect circuit is connected. Only **VPC*	- may be returned.
	//
	// example:
	//
	// VPC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The number of Express Connect circuits that are established.
	//
	// example:
	//
	// 0
	VirtualPhysicalConnectionCount *int32 `json:"VirtualPhysicalConnectionCount,omitempty" xml:"VirtualPhysicalConnectionCount,omitempty"`
	// The VLAN ID of the shared Express Connect circuit.
	//
	// example:
	//
	// 10
	VlanId *string `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
	// The status of the shared Express Connect circuit. Valid values:
	//
	// 	- **Confirmed**
	//
	// 	- **UnConfirmed**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Confirmed
	VpconnStatus *string `json:"VpconnStatus,omitempty" xml:"VpconnStatus,omitempty"`
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
	// The key of tag N added to the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N added to the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
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
