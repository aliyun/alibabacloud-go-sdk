// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeForwardTableEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForwardTableEntries(v *DescribeForwardTableEntriesResponseBodyForwardTableEntries) *DescribeForwardTableEntriesResponseBody
	GetForwardTableEntries() *DescribeForwardTableEntriesResponseBodyForwardTableEntries
	SetPageNumber(v int32) *DescribeForwardTableEntriesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeForwardTableEntriesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeForwardTableEntriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeForwardTableEntriesResponseBody
	GetTotalCount() *int32
}

type DescribeForwardTableEntriesResponseBody struct {
	// The details of DNAT entries.
	ForwardTableEntries *DescribeForwardTableEntriesResponseBodyForwardTableEntries `json:"ForwardTableEntries,omitempty" xml:"ForwardTableEntries,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A6C4A8B1-7561-4509-949C-20DEB40D71E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeForwardTableEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeForwardTableEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeForwardTableEntriesResponseBody) GetForwardTableEntries() *DescribeForwardTableEntriesResponseBodyForwardTableEntries {
	return s.ForwardTableEntries
}

func (s *DescribeForwardTableEntriesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeForwardTableEntriesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeForwardTableEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeForwardTableEntriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeForwardTableEntriesResponseBody) SetForwardTableEntries(v *DescribeForwardTableEntriesResponseBodyForwardTableEntries) *DescribeForwardTableEntriesResponseBody {
	s.ForwardTableEntries = v
	return s
}

func (s *DescribeForwardTableEntriesResponseBody) SetPageNumber(v int32) *DescribeForwardTableEntriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBody) SetPageSize(v int32) *DescribeForwardTableEntriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBody) SetRequestId(v string) *DescribeForwardTableEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBody) SetTotalCount(v int32) *DescribeForwardTableEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeForwardTableEntriesResponseBodyForwardTableEntries struct {
	ForwardTableEntry []*DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry `json:"ForwardTableEntry,omitempty" xml:"ForwardTableEntry,omitempty" type:"Repeated"`
}

func (s DescribeForwardTableEntriesResponseBodyForwardTableEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeForwardTableEntriesResponseBodyForwardTableEntries) GoString() string {
	return s.String()
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) GetForwardTableEntry() []*DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry {
	return s.ForwardTableEntry
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) SetForwardTableEntry(v []*DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) *DescribeForwardTableEntriesResponseBodyForwardTableEntries {
	s.ForwardTableEntry = v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) Validate() error {
	return dara.Validate(s)
}

type DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry struct {
	// 	- The EIPs that can be accessed over the Internet when you query DNAT entries of Internet NAT gateways.
	//
	// 	- The NAT IP addresses that can be accessed by external networks when you query DNAT entries of VPC NAT gateways.
	//
	// example:
	//
	// 139.79.XX.XX
	ExternalIp *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	// 	- The external port or port range that is used for port forwarding when you query DNAT entries of Internet NAT gateways.
	//
	// 	- The port that is used when the NAT IP address can be accessed by external networks when you query DNAT entries of VPC NAT gateways.
	//
	// example:
	//
	// 80
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// The ID of the DNAT entry.
	//
	// example:
	//
	// fwd-119smw5tk****
	ForwardEntryId *string `json:"ForwardEntryId,omitempty" xml:"ForwardEntryId,omitempty"`
	// The name of the DNAT entry.
	//
	// example:
	//
	// ForwardEntry-1
	ForwardEntryName *string `json:"ForwardEntryName,omitempty" xml:"ForwardEntryName,omitempty"`
	// The ID of the DNAT table to which the DNAT entry belongs.
	//
	// example:
	//
	// ftb-11tc6xgmv****
	ForwardTableId *string `json:"ForwardTableId,omitempty" xml:"ForwardTableId,omitempty"`
	// The private IP address.
	//
	// 	- The private IP address of the ECS instance that uses DNAT entries to communicate with the Internet when you query DNAT entries of Internet NAT gateways.
	//
	// 	- The private IP address that uses DNAT entries when you query DNAT entries of VPC NAT gateways.
	//
	// example:
	//
	// 192.168.XX.XX
	InternalIp *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	// 	- The internal port or port range that is used for port forwarding when you query DNAT entries of Internet NAT gateways.
	//
	// 	- The destination ECS instance port to be mapped when you query DNAT entries of VPC NAT gateways.
	//
	// example:
	//
	// 25
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	// The protocol. Valid values:
	//
	// 	- **TCP**
	//
	// 	- **UDP**
	//
	// 	- **Any**
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The ID of the NAT gateway to which the DNAT entry belongs.
	//
	// example:
	//
	// ngw-bp1uewa15k4iy5770****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The status of the DNAT entry. Valid values:
	//
	// 	- **Pending**
	//
	// 	- **Available**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) GoString() string {
	return s.String()
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) GetExternalIp() *string {
	return s.ExternalIp
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) GetForwardEntryId() *string {
	return s.ForwardEntryId
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) GetForwardEntryName() *string {
	return s.ForwardEntryName
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) GetForwardTableId() *string {
	return s.ForwardTableId
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) GetInternalIp() *string {
	return s.InternalIp
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) GetInternalPort() *string {
	return s.InternalPort
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) GetStatus() *string {
	return s.Status
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) SetExternalIp(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry {
	s.ExternalIp = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) SetExternalPort(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry {
	s.ExternalPort = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) SetForwardEntryId(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry {
	s.ForwardEntryId = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) SetForwardEntryName(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry {
	s.ForwardEntryName = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) SetForwardTableId(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry {
	s.ForwardTableId = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) SetInternalIp(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry {
	s.InternalIp = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) SetInternalPort(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry {
	s.InternalPort = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) SetIpProtocol(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry {
	s.IpProtocol = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) SetNatGatewayId(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) SetStatus(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry {
	s.Status = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntriesForwardTableEntry) Validate() error {
	return dara.Validate(s)
}
