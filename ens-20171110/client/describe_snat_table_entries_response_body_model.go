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
	SetSnatTableEntries(v []*DescribeSnatTableEntriesResponseBodySnatTableEntries) *DescribeSnatTableEntriesResponseBody
	GetSnatTableEntries() []*DescribeSnatTableEntriesResponseBodySnatTableEntries
	SetTotalCount(v int32) *DescribeSnatTableEntriesResponseBody
	GetTotalCount() *int32
}

type DescribeSnatTableEntriesResponseBody struct {
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
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the SNAT entries.
	SnatTableEntries []*DescribeSnatTableEntriesResponseBodySnatTableEntries `json:"SnatTableEntries,omitempty" xml:"SnatTableEntries,omitempty" type:"Repeated"`
	// The number of SNAT entries that are returned.
	//
	// example:
	//
	// 7
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

func (s *DescribeSnatTableEntriesResponseBody) GetSnatTableEntries() []*DescribeSnatTableEntriesResponseBodySnatTableEntries {
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

func (s *DescribeSnatTableEntriesResponseBody) SetSnatTableEntries(v []*DescribeSnatTableEntriesResponseBodySnatTableEntries) *DescribeSnatTableEntriesResponseBody {
	s.SnatTableEntries = v
	return s
}

func (s *DescribeSnatTableEntriesResponseBody) SetTotalCount(v int32) *DescribeSnatTableEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSnatTableEntriesResponseBodySnatTableEntries struct {
	// Specifies whether to enable EIP affinity. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// **
	//
	// **Description*	- After you enable EIP affinity, if multiple EIPs are associated with an SNAT entry, each client uses one EIP to access the Internet. If EIP affinity is disabled, each client uses a random EIP to access the Internet.
	//
	// example:
	//
	// false
	EipAffinity *bool `json:"EipAffinity,omitempty" xml:"EipAffinity,omitempty"`
	// The timeout period for idle connections. Valid values: **1*	- to **86400**. Unit: seconds.
	//
	// example:
	//
	// 900
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// Whether to enable operator affinity. Value taking:
	//
	// - false:Do not open.
	//
	// - true:Open.
	//
	// example:
	//
	// true
	IspAffinity *bool `json:"IspAffinity,omitempty" xml:"IspAffinity,omitempty"`
	// The ID of the NAT gateway.
	//
	// example:
	//
	// nat-5t7nh1cfm6kxiszlttr38****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The ID of the SNAT entry.
	//
	// example:
	//
	// snat-5tfjp3537mi6iokl59g5c****
	SnatEntryId *string `json:"SnatEntryId,omitempty" xml:"SnatEntryId,omitempty"`
	// The name of the SNAT entry.
	//
	// example:
	//
	// test0
	SnatEntryName *string `json:"SnatEntryName,omitempty" xml:"SnatEntryName,omitempty"`
	// The EIP specified in the SNAT entry.
	//
	// example:
	//
	// 120.XXX.XXX.71
	SnatIp *string `json:"SnatIp,omitempty" xml:"SnatIp,omitempty"`
	// The source CIDR block specified in the SNAT entry.
	//
	// example:
	//
	// 10.0.0.13/32
	SourceCIDR *string `json:"SourceCIDR,omitempty" xml:"SourceCIDR,omitempty"`
	// The secondary EIP. Multiple EIPs are separated by commas (,).
	//
	// example:
	//
	// 101.XXX.XXX.7
	StandbySnatIp *string `json:"StandbySnatIp,omitempty" xml:"StandbySnatIp,omitempty"`
	// The status of the secondary EIP. Valid values:
	//
	// 	- Running
	//
	// 	- Stopping
	//
	// 	- Stopped
	//
	// 	- Starting
	//
	// example:
	//
	// Stopped
	StandbyStatus *string `json:"StandbyStatus,omitempty" xml:"StandbyStatus,omitempty"`
	// The status of the SNAT entry. Valid values:
	//
	// 	- Pending: The SNAT entry is being created or modified.
	//
	// 	- Available: The SNAT entry is available.
	//
	// 	- Deleting: The SNAT entry is being deleted.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSnatTableEntriesResponseBodySnatTableEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnatTableEntriesResponseBodySnatTableEntries) GoString() string {
	return s.String()
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetEipAffinity() *bool {
	return s.EipAffinity
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetIspAffinity() *bool {
	return s.IspAffinity
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetSnatEntryId() *string {
	return s.SnatEntryId
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetSnatEntryName() *string {
	return s.SnatEntryName
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetSnatIp() *string {
	return s.SnatIp
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetSourceCIDR() *string {
	return s.SourceCIDR
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetStandbySnatIp() *string {
	return s.StandbySnatIp
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetStandbyStatus() *string {
	return s.StandbyStatus
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) GetStatus() *string {
	return s.Status
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetEipAffinity(v bool) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.EipAffinity = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetIdleTimeout(v int32) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.IdleTimeout = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetIspAffinity(v bool) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.IspAffinity = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetNatGatewayId(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetSnatEntryId(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.SnatEntryId = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetSnatEntryName(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.SnatEntryName = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetSnatIp(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.SnatIp = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetSourceCIDR(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.SourceCIDR = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetStandbySnatIp(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.StandbySnatIp = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetStandbyStatus(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.StandbyStatus = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) SetStatus(v string) *DescribeSnatTableEntriesResponseBodySnatTableEntries {
	s.Status = &v
	return s
}

func (s *DescribeSnatTableEntriesResponseBodySnatTableEntries) Validate() error {
	return dara.Validate(s)
}
