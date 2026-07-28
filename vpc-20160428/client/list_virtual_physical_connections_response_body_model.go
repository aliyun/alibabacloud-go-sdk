// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirtualPhysicalConnectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListVirtualPhysicalConnectionsResponseBody
	GetCount() *int32
	SetNextToken(v string) *ListVirtualPhysicalConnectionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVirtualPhysicalConnectionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListVirtualPhysicalConnectionsResponseBody
	GetTotalCount() *int32
	SetVirtualPhysicalConnections(v []*ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) *ListVirtualPhysicalConnectionsResponseBody
	GetVirtualPhysicalConnections() []*ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections
}

type ListVirtualPhysicalConnectionsResponseBody struct {
	// The number of entries returned in the current query.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Indicates whether a next query token is available. Valid values:
	//
	// - If **NextToken*	- is empty, no next query is available.
	//
	// - If **NextToken*	- is returned, the value indicates the token for the next query.
	//
	// example:
	//
	// dd20****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2A55F69E-EE3D-5CBE-8805-734F7D5B46B9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of shared Express Connect circuits.
	VirtualPhysicalConnections []*ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections `json:"VirtualPhysicalConnections,omitempty" xml:"VirtualPhysicalConnections,omitempty" type:"Repeated"`
}

func (s ListVirtualPhysicalConnectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualPhysicalConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirtualPhysicalConnectionsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListVirtualPhysicalConnectionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVirtualPhysicalConnectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVirtualPhysicalConnectionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVirtualPhysicalConnectionsResponseBody) GetVirtualPhysicalConnections() []*ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	return s.VirtualPhysicalConnections
}

func (s *ListVirtualPhysicalConnectionsResponseBody) SetCount(v int32) *ListVirtualPhysicalConnectionsResponseBody {
	s.Count = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBody) SetNextToken(v string) *ListVirtualPhysicalConnectionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBody) SetRequestId(v string) *ListVirtualPhysicalConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBody) SetTotalCount(v int32) *ListVirtualPhysicalConnectionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBody) SetVirtualPhysicalConnections(v []*ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) *ListVirtualPhysicalConnectionsResponseBody {
	s.VirtualPhysicalConnections = v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBody) Validate() error {
	if s.VirtualPhysicalConnections != nil {
		for _, item := range s.VirtualPhysicalConnections {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections struct {
	// The access point ID of the Express Connect circuit.
	//
	// example:
	//
	// ap-cn-hangzhou-finance-yh-E
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// The physical location of the access device for the Express Connect circuit.
	//
	// example:
	//
	// 余杭经济开发区XXX交叉口, 余杭XX机房, E***包间
	AdLocation *string `json:"AdLocation,omitempty" xml:"AdLocation,omitempty"`
	// The Alibaba Cloud account ID of the shared Express Connect circuits owner.
	//
	// example:
	//
	// 15346073170691****
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The bandwidth of the Express Connect circuit. Unit: Mbit/s.
	//
	// example:
	//
	// 50
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The business status of the Express Connect circuit. Valid values:
	//
	// - **Normal**: Enabled.
	//
	// - **FinancialLocked**: Financial lock.
	//
	// - **SecurityLocked**: Locked for security reasons.
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The billing method of the Express Connect circuit.
	//
	// Valid values: **Prepaid**, which indicates subscription.
	//
	// example:
	//
	// Prepaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The circuit code provided by the carrier for the Express Connect circuit.
	//
	// example:
	//
	// longtel001
	CircuitCode *string `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	// The time when the Express Connect circuit was created.
	//
	// example:
	//
	// 2021-06-08T12:20:55
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
	// 2021-10-08T10:44Z
	EnabledTime *string `json:"EnabledTime,omitempty" xml:"EnabledTime,omitempty"`
	// The expiration time of the shared Express Connect circuits.
	//
	// The time is displayed in the ISO 8601 standard in UTC. Format: YYYY-MM-DDThh:mm:ssZ.
	//
	// example:
	//
	// 2021-11-08T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The expected bandwidth of the shared Express Connect circuits. The expected bandwidth takes effect only after payment is complete.
	//
	// Unit: **M*	- indicates Mbit/s, and **G*	- indicates Gbit/s.
	//
	// example:
	//
	// 50M
	ExpectSpec *string `json:"ExpectSpec,omitempty" xml:"ExpectSpec,omitempty"`
	// The carrier that provides the physical connection. Valid values:
	//
	// - **CT**: China Telecom.
	//
	// - **CU**: China Unicom.
	//
	// - **CM**: China Mobile.
	//
	// - **CO**: Other carriers in the Chinese mainland.
	//
	// - **Equinix**: Equinix.
	//
	// - **Other**: Other carriers outside the Chinese mainland.
	//
	// example:
	//
	// CU
	LineOperator *string `json:"LineOperator,omitempty" xml:"LineOperator,omitempty"`
	// The status of the Letter of Authorization (LOA). Valid values:
	//
	// - **Applying**: The LOA application is being processed.
	//
	// - **Accept**: The LOA application is approved.
	//
	// - **Available**: The LOA is available.
	//
	// - **Rejected**: The LOA application is rejected.
	//
	// - **Completing**: The Express Connect circuit is under construction.
	//
	// - **Complete**: The Express Connect circuit construction is complete.
	//
	// - **Deleted**: The LOA is deleted.
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
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The payer of the shared Express Connect circuits. Valid values:
	//
	// - **PayByPhysicalConnectionOwner**: The owner of the Express Connect circuit associated with the shared Express Connect circuits is the payer.
	//
	// - **PayByVirtualPhysicalConnectionOwner**: The owner of the shared Express Connect circuits is the payer.
	//
	// example:
	//
	// PayByPhysicalConnectionOwner
	OrderMode *string `json:"OrderMode,omitempty" xml:"OrderMode,omitempty"`
	// The Alibaba Cloud account ID to which the Express Connect circuit belongs.
	//
	// example:
	//
	// 18311773240248****
	ParentPhysicalConnectionAliUid *string `json:"ParentPhysicalConnectionAliUid,omitempty" xml:"ParentPhysicalConnectionAliUid,omitempty"`
	// The instance ID of the Express Connect circuit.
	//
	// example:
	//
	// pc-bp1ciz7ekd2grn1as****
	ParentPhysicalConnectionId *string `json:"ParentPhysicalConnectionId,omitempty" xml:"ParentPhysicalConnectionId,omitempty"`
	// The geographic location of the on-premises data center.
	//
	// example:
	//
	// 浙江省杭州市XX区XX街道XX号
	PeerLocation *string `json:"PeerLocation,omitempty" xml:"PeerLocation,omitempty"`
	// The instance ID of the shared Express Connect circuits.
	//
	// example:
	//
	// pc-bp1mrgfbtmc9brre7****
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// The port number of the access device for the Express Connect circuit.
	//
	// example:
	//
	// 80
	PortNumber *string `json:"PortNumber,omitempty" xml:"PortNumber,omitempty"`
	// The port type of the Express Connect circuit. Valid values:
	//
	// - **100Base-T**: 100M Ethernet port.
	//
	// - **1000Base-T**: 1 GE electrical port.
	//
	// - **1000Base-LX**: GE single-mode optical port (10 km).
	//
	// - **10GBase-T**: 10 GE electrical port.
	//
	// - **10GBase-LR**: 10 GE single-mode optical port (10 km).
	//
	// - **40GBase-LR**: 40 GE single-mode optical port.
	//
	// - **100GBase-LR**: 100 GE single-mode optical port.
	//
	// example:
	//
	// 10GBase-LR
	PortType *string `json:"PortType,omitempty" xml:"PortType,omitempty"`
	// The type of the Express Connect circuit. Valid values:
	//
	// - **VirtualPhysicalConnection**: Shared Express Connect circuits.
	//
	// - **PhysicalConnection**: Dedicated Express Connect circuit.
	//
	// example:
	//
	// VirtualPhysicalConnection
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The ID of the redundant Express Connect circuit.
	//
	// example:
	//
	// pc-119mfjzm****
	RedundantPhysicalConnectionId *string `json:"RedundantPhysicalConnectionId,omitempty" xml:"RedundantPhysicalConnectionId,omitempty"`
	// The ID of the resource group to which the shared Express Connect circuits belong.
	//
	// example:
	//
	// rg-acfm3wmsyui****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The bandwidth of the shared Express Connect circuits.
	//
	// Unit: **M*	- indicates Mbit/s, and **G*	- indicates Gbit/s.
	//
	// example:
	//
	// 50M
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The status of the Express Connect circuit. Valid values:
	//
	// - **Initial**: The application is under review.
	//
	// - **Approved**: The application is approved.
	//
	// - **Allocating**: Resources are being allocated.
	//
	// - **Allocated**: Under construction.
	//
	// - **Confirmed**: Pending confirmation.
	//
	// - **Enabled**: Enabled.
	//
	// - **Rejected**: The application is rejected.
	//
	// - **Canceled**: Canceled.
	//
	// - **Allocation Failed**: Resource allocation failed.
	//
	// - **Terminated**: Terminated.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag list.
	Tags []*ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnectionsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the Express Connect circuit. Default value: **VPC**.
	//
	// example:
	//
	// VPC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The business status of the shared Express Connect circuits. Valid values:
	//
	// - **Confirmed**: The shared Express Connect circuits have been accepted.
	//
	// - **UnConfirmed**: The shared Express Connect circuits have not been accepted.
	//
	// - **Deleted**: The shared Express Connect circuits have been deleted.
	//
	// example:
	//
	// Confirmed
	VirtualPhysicalConnectionStatus *string `json:"VirtualPhysicalConnectionStatus,omitempty" xml:"VirtualPhysicalConnectionStatus,omitempty"`
	// The VLAN ID of the shared Express Connect circuits.
	//
	// example:
	//
	// 10
	VlanId *string `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
}

func (s ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GoString() string {
	return s.String()
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetAdLocation() *string {
	return s.AdLocation
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetAliUid() *string {
	return s.AliUid
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetCircuitCode() *string {
	return s.CircuitCode
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetDescription() *string {
	return s.Description
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetEnabledTime() *string {
	return s.EnabledTime
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetEndTime() *string {
	return s.EndTime
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetExpectSpec() *string {
	return s.ExpectSpec
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetLineOperator() *string {
	return s.LineOperator
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetLoaStatus() *string {
	return s.LoaStatus
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetName() *string {
	return s.Name
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetOrderMode() *string {
	return s.OrderMode
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetParentPhysicalConnectionAliUid() *string {
	return s.ParentPhysicalConnectionAliUid
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetParentPhysicalConnectionId() *string {
	return s.ParentPhysicalConnectionId
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetPeerLocation() *string {
	return s.PeerLocation
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetPortNumber() *string {
	return s.PortNumber
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetPortType() *string {
	return s.PortType
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetProductType() *string {
	return s.ProductType
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetRedundantPhysicalConnectionId() *string {
	return s.RedundantPhysicalConnectionId
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetSpec() *string {
	return s.Spec
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetStatus() *string {
	return s.Status
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetTags() []*ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnectionsTags {
	return s.Tags
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetType() *string {
	return s.Type
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetVirtualPhysicalConnectionStatus() *string {
	return s.VirtualPhysicalConnectionStatus
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) GetVlanId() *string {
	return s.VlanId
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetAccessPointId(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.AccessPointId = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetAdLocation(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.AdLocation = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetAliUid(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.AliUid = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetBandwidth(v int64) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.Bandwidth = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetBusinessStatus(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.BusinessStatus = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetChargeType(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.ChargeType = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetCircuitCode(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.CircuitCode = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetCreationTime(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.CreationTime = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetDescription(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.Description = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetEnabledTime(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.EnabledTime = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetEndTime(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.EndTime = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetExpectSpec(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.ExpectSpec = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetLineOperator(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.LineOperator = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetLoaStatus(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.LoaStatus = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetName(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.Name = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetOrderMode(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.OrderMode = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetParentPhysicalConnectionAliUid(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.ParentPhysicalConnectionAliUid = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetParentPhysicalConnectionId(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.ParentPhysicalConnectionId = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetPeerLocation(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.PeerLocation = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetPhysicalConnectionId(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.PhysicalConnectionId = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetPortNumber(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.PortNumber = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetPortType(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.PortType = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetProductType(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.ProductType = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetRedundantPhysicalConnectionId(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.RedundantPhysicalConnectionId = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetResourceGroupId(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetSpec(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.Spec = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetStatus(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.Status = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetTags(v []*ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnectionsTags) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.Tags = v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetType(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.Type = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetVirtualPhysicalConnectionStatus(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.VirtualPhysicalConnectionStatus = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) SetVlanId(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections {
	s.VlanId = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnections) Validate() error {
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

type ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnectionsTags struct {
	// The tag key of the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The tag key cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The tag value cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnectionsTags) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnectionsTags) GoString() string {
	return s.String()
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnectionsTags) GetKey() *string {
	return s.Key
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnectionsTags) GetValue() *string {
	return s.Value
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnectionsTags) SetKey(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnectionsTags {
	s.Key = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnectionsTags) SetValue(v string) *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnectionsTags {
	s.Value = &v
	return s
}

func (s *ListVirtualPhysicalConnectionsResponseBodyVirtualPhysicalConnectionsTags) Validate() error {
	return dara.Validate(s)
}
