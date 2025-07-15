// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFullNatEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFullNatEntries(v []*ListFullNatEntriesResponseBodyFullNatEntries) *ListFullNatEntriesResponseBody
	GetFullNatEntries() []*ListFullNatEntriesResponseBodyFullNatEntries
	SetFullNatTableId(v string) *ListFullNatEntriesResponseBody
	GetFullNatTableId() *string
	SetMaxResults(v int64) *ListFullNatEntriesResponseBody
	GetMaxResults() *int64
	SetNatGatewayId(v string) *ListFullNatEntriesResponseBody
	GetNatGatewayId() *string
	SetNextToken(v string) *ListFullNatEntriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListFullNatEntriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListFullNatEntriesResponseBody
	GetTotalCount() *int64
}

type ListFullNatEntriesResponseBody struct {
	// The information about the FULLNAT entries that are queried.
	FullNatEntries []*ListFullNatEntriesResponseBodyFullNatEntries `json:"FullNatEntries,omitempty" xml:"FullNatEntries,omitempty" type:"Repeated"`
	// The ID of the FULLNAT table to which the queried FULLNAT entries belong.
	//
	// example:
	//
	// fullnat-gw8fz23jezpbblf1j****
	FullNatTableId *string `json:"FullNatTableId,omitempty" xml:"FullNatTableId,omitempty"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 1
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The ID of the VPC NAT gateway.
	//
	// example:
	//
	// ngw-gw8054kn57y3hq3bv****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// Indicates whether the token for the next query exists. Valid values:
	//
	// 	- If the value of **NextToken*	- is empty, no next queries are sent.
	//
	// 	- If the value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F03E41F6-1A74-311F-8D98-124EEE9F37B8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of FULLNAT entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFullNatEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFullNatEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFullNatEntriesResponseBody) GetFullNatEntries() []*ListFullNatEntriesResponseBodyFullNatEntries {
	return s.FullNatEntries
}

func (s *ListFullNatEntriesResponseBody) GetFullNatTableId() *string {
	return s.FullNatTableId
}

func (s *ListFullNatEntriesResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListFullNatEntriesResponseBody) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ListFullNatEntriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFullNatEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFullNatEntriesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListFullNatEntriesResponseBody) SetFullNatEntries(v []*ListFullNatEntriesResponseBodyFullNatEntries) *ListFullNatEntriesResponseBody {
	s.FullNatEntries = v
	return s
}

func (s *ListFullNatEntriesResponseBody) SetFullNatTableId(v string) *ListFullNatEntriesResponseBody {
	s.FullNatTableId = &v
	return s
}

func (s *ListFullNatEntriesResponseBody) SetMaxResults(v int64) *ListFullNatEntriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListFullNatEntriesResponseBody) SetNatGatewayId(v string) *ListFullNatEntriesResponseBody {
	s.NatGatewayId = &v
	return s
}

func (s *ListFullNatEntriesResponseBody) SetNextToken(v string) *ListFullNatEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListFullNatEntriesResponseBody) SetRequestId(v string) *ListFullNatEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFullNatEntriesResponseBody) SetTotalCount(v int64) *ListFullNatEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFullNatEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFullNatEntriesResponseBodyFullNatEntries struct {
	// The backend IP address that is used for FULLNAT address translation in FULLNAT entries.
	//
	// example:
	//
	// 192.168.XX.XX
	AccessIp *string `json:"AccessIp,omitempty" xml:"AccessIp,omitempty"`
	// The backend port that is used for port mapping in FULLNAT entries. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 80
	AccessPort *string `json:"AccessPort,omitempty" xml:"AccessPort,omitempty"`
	// The time when the FULLNAT entry was created.
	//
	// example:
	//
	// 2021-10-27T02:44:40Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the FULLNAT entry.
	//
	// The name must be 2 to 128 characters in length. It must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// abc
	FullNatEntryDescription *string `json:"FullNatEntryDescription,omitempty" xml:"FullNatEntryDescription,omitempty"`
	// The ID of the FULLNAT entry.
	//
	// example:
	//
	// fullnat-gw8fz23jezpbblf1j****
	FullNatEntryId *string `json:"FullNatEntryId,omitempty" xml:"FullNatEntryId,omitempty"`
	// The name of the FULLNAT entry.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// test
	FullNatEntryName *string `json:"FullNatEntryName,omitempty" xml:"FullNatEntryName,omitempty"`
	// The status of the FULLNAT entry. Valid values:
	//
	// 	- **Pending**
	//
	// 	- **Available**
	//
	// 	- **Deleting**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Available
	FullNatEntryStatus *string `json:"FullNatEntryStatus,omitempty" xml:"FullNatEntryStatus,omitempty"`
	// The ID of the FULLNAT table to which the FULLNAT entry belongs.
	//
	// example:
	//
	// fulltb-gw88z7hhlv43rmb26****
	FullNatTableId *string `json:"FullNatTableId,omitempty" xml:"FullNatTableId,omitempty"`
	// The protocol of the packets that are forwarded. Valid values:
	//
	// 	- **TCP**
	//
	// 	- **UDP**
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The NAT IP address that is used for address translation in FULLNAT entries.
	//
	// example:
	//
	// 192.168.XX.XX
	NatIp *string `json:"NatIp,omitempty" xml:"NatIp,omitempty"`
	// The frontend port that is used for port mapping in FULLNAT entries. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 80
	NatIpPort *string `json:"NatIpPort,omitempty" xml:"NatIpPort,omitempty"`
	// The ID of the elastic network interface (ENI).
	//
	// example:
	//
	// eni-gw80wedm8pq0tpr2****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The type of the ENI. The value is set to **Endpoint**, which indicates a reverse endpoint.
	//
	// example:
	//
	// Endpoint
	NetworkInterfaceType *string `json:"NetworkInterfaceType,omitempty" xml:"NetworkInterfaceType,omitempty"`
}

func (s ListFullNatEntriesResponseBodyFullNatEntries) String() string {
	return dara.Prettify(s)
}

func (s ListFullNatEntriesResponseBodyFullNatEntries) GoString() string {
	return s.String()
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) GetAccessIp() *string {
	return s.AccessIp
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) GetAccessPort() *string {
	return s.AccessPort
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) GetFullNatEntryDescription() *string {
	return s.FullNatEntryDescription
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) GetFullNatEntryId() *string {
	return s.FullNatEntryId
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) GetFullNatEntryName() *string {
	return s.FullNatEntryName
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) GetFullNatEntryStatus() *string {
	return s.FullNatEntryStatus
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) GetFullNatTableId() *string {
	return s.FullNatTableId
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) GetNatIp() *string {
	return s.NatIp
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) GetNatIpPort() *string {
	return s.NatIpPort
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) GetNetworkInterfaceType() *string {
	return s.NetworkInterfaceType
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) SetAccessIp(v string) *ListFullNatEntriesResponseBodyFullNatEntries {
	s.AccessIp = &v
	return s
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) SetAccessPort(v string) *ListFullNatEntriesResponseBodyFullNatEntries {
	s.AccessPort = &v
	return s
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) SetCreationTime(v string) *ListFullNatEntriesResponseBodyFullNatEntries {
	s.CreationTime = &v
	return s
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) SetFullNatEntryDescription(v string) *ListFullNatEntriesResponseBodyFullNatEntries {
	s.FullNatEntryDescription = &v
	return s
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) SetFullNatEntryId(v string) *ListFullNatEntriesResponseBodyFullNatEntries {
	s.FullNatEntryId = &v
	return s
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) SetFullNatEntryName(v string) *ListFullNatEntriesResponseBodyFullNatEntries {
	s.FullNatEntryName = &v
	return s
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) SetFullNatEntryStatus(v string) *ListFullNatEntriesResponseBodyFullNatEntries {
	s.FullNatEntryStatus = &v
	return s
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) SetFullNatTableId(v string) *ListFullNatEntriesResponseBodyFullNatEntries {
	s.FullNatTableId = &v
	return s
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) SetIpProtocol(v string) *ListFullNatEntriesResponseBodyFullNatEntries {
	s.IpProtocol = &v
	return s
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) SetNatIp(v string) *ListFullNatEntriesResponseBodyFullNatEntries {
	s.NatIp = &v
	return s
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) SetNatIpPort(v string) *ListFullNatEntriesResponseBodyFullNatEntries {
	s.NatIpPort = &v
	return s
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) SetNetworkInterfaceId(v string) *ListFullNatEntriesResponseBodyFullNatEntries {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) SetNetworkInterfaceType(v string) *ListFullNatEntriesResponseBodyFullNatEntries {
	s.NetworkInterfaceType = &v
	return s
}

func (s *ListFullNatEntriesResponseBodyFullNatEntries) Validate() error {
	return dara.Validate(s)
}
