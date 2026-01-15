// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnatTableEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSnatTableEntriesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSnatTableEntriesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSnatTableEntriesResponseBody
	GetRequestId() *string
	SetSnatTableEntries(v *DescribeSnatTableEntriesResponseBodySnatTableEntries) *DescribeSnatTableEntriesResponseBody
	GetSnatTableEntries() *DescribeSnatTableEntriesResponseBodySnatTableEntries
	SetTotalCount(v int32) *DescribeSnatTableEntriesResponseBody
	GetTotalCount() *int32
}

type DescribeSnatTableEntriesResponseBody struct {
	// The number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6D7E89B1-1C5B-412B-8585-4908E222EED5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of SNAT entries.
	SnatTableEntries *DescribeSnatTableEntriesResponseBodySnatTableEntries `json:"SnatTableEntries,omitempty" xml:"SnatTableEntries,omitempty" type:"Struct"`
	// The number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSnatTableEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnatTableEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnatTableEntriesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSnatTableEntriesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSnatTableEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSnatTableEntriesResponseBody) GetSnatTableEntries() *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	return s.SnatTableEntries
}

func (s *DescribeSnatTableEntriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSnatTableEntriesResponseBody) SetPageNumber(v int32) *DescribeSnatTableEntriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBody) SetPageSize(v int32) *DescribeSnatTableEntriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBody) SetRequestId(v string) *DescribeSnatTableEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBody) SetSnatTableEntries(v *DescribeSnatTableEntriesResponseBodySnatTableEntries) *DescribeSnatTableEntriesResponseBody {
	s.SnatTableEntries = v
	return s
}

func (s *DescribeSnatTableEntriesResponseBody) SetTotalCount(v int32) *DescribeSnatTableEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBody) Validate() error {
	if s.SnatTableEntries != nil {
		if err := s.SnatTableEntries.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSnatTableEntriesResponseBodySnatTableEntries struct {
	SnatTableEntry []*DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry `json:"SnatTableEntry,omitempty" xml:"SnatTableEntry,omitempty" type:"Repeated"`
}

func (s DescribeSnatTableEntriesResponseBodySnatTableEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnatTableEntriesResponseBodySnatTableEntries) GoString() string {
	return s.String()
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetSnatTableEntry() []*DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry {
	return s.SnatTableEntry
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetSnatTableEntry(v []*DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.SnatTableEntry = v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) Validate() error {
	if s.SnatTableEntry != nil {
		for _, item := range s.SnatTableEntry {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry struct {
	// Whether to enable IP affinity. Values:
	//
	// - **0**: Disable IP affinity. - **1**: Enable IP affinity.
	//
	// > After enabling the IP affinity switch, if an SNAT entry is bound to multiple EIPs or NAT IPs, the same client will use the same EIP or NAT IP for access. Otherwise, the client will randomly select from the bound EIPs or NAT IPs for access.
	//
	// example:
	//
	// 1
	EipAffinity *string `json:"EipAffinity,omitempty" xml:"EipAffinity,omitempty"`
	// The ID of the NAT gateway to which the SNAT entry belongs.
	//
	// example:
	//
	// ngw-bp1uewa15k4iy5770****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// Elastic Network Interface ID
	//
	// example:
	//
	// eni-gw8g131ef2dnbu3k****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The ID of the SNAT entry.
	//
	// example:
	//
	// snat-kmd6nv8fy****
	SnatEntryId *string `json:"SnatEntryId,omitempty" xml:"SnatEntryId,omitempty"`
	// The name of the SNAT entry.
	//
	// example:
	//
	// SnatEntry-1
	SnatEntryName *string `json:"SnatEntryName,omitempty" xml:"SnatEntryName,omitempty"`
	// 	- When you query SNAT entries of Internet NAT gateways, this parameter indicates the EIP in an SNAT entry.
	//
	// 	- When you query SNAT entries of VPC NAT gateways, this parameter indicates the NAT IP address in an SNAT entry.
	//
	// example:
	//
	// 116.22.XX.XX
	SnatIp *string `json:"SnatIp,omitempty" xml:"SnatIp,omitempty"`
	// The ID of the SNAT table to which the SNAT entry belongs.
	//
	// example:
	//
	// stb-gz3r3odawdgffde****
	SnatTableId *string `json:"SnatTableId,omitempty" xml:"SnatTableId,omitempty"`
	// The source CIDR block specified in the SNAT entry.
	//
	// example:
	//
	// 116.22.XX.XX/24
	SourceCIDR *string `json:"SourceCIDR,omitempty" xml:"SourceCIDR,omitempty"`
	// 	- When you query SNAT entries of Internet NAT gateways, this parameter indicates the ID of the vSwitch that uses SNAT to access the Internet.
	//
	// 	- When you query SNAT entries of VPC NAT gateways, this parameter indicates the ID of the vSwitch that uses SNAT to access external networks.
	//
	// example:
	//
	// vsw-3xbdsffvfgdfds****
	SourceVSwitchId *string `json:"SourceVSwitchId,omitempty" xml:"SourceVSwitchId,omitempty"`
	// The status of the SNAT entry. Valid values:
	//
	// 	- **Pending**
	//
	// 	- **Available**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) GoString() string {
	return s.String()
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) GetEipAffinity() *string {
	return s.EipAffinity
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) GetSnatEntryId() *string {
	return s.SnatEntryId
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) GetSnatEntryName() *string {
	return s.SnatEntryName
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) GetSnatIp() *string {
	return s.SnatIp
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) GetSnatTableId() *string {
	return s.SnatTableId
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) GetSourceCIDR() *string {
	return s.SourceCIDR
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) GetSourceVSwitchId() *string {
	return s.SourceVSwitchId
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) GetStatus() *string {
	return s.Status
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) SetEipAffinity(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry {
	s.EipAffinity = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) SetNatGatewayId(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) SetNetworkInterfaceId(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) SetSnatEntryId(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry {
	s.SnatEntryId = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) SetSnatEntryName(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry {
	s.SnatEntryName = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) SetSnatIp(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry {
	s.SnatIp = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) SetSnatTableId(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry {
	s.SnatTableId = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) SetSourceCIDR(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry {
	s.SourceCIDR = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) SetSourceVSwitchId(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry {
	s.SourceVSwitchId = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) SetStatus(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry {
	s.Status = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntriesSnatTableEntry) Validate() error {
	return dara.Validate(s)
}
