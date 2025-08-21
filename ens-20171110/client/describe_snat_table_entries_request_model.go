// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnatTableEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNatGatewayId(v string) *DescribeSnatTableEntriesRequest
	GetNatGatewayId() *string
	SetPageNumber(v int32) *DescribeSnatTableEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSnatTableEntriesRequest
	GetPageSize() *int32
	SetSnatEntryId(v string) *DescribeSnatTableEntriesRequest
	GetSnatEntryId() *string
	SetSnatEntryName(v string) *DescribeSnatTableEntriesRequest
	GetSnatEntryName() *string
	SetSnatIp(v string) *DescribeSnatTableEntriesRequest
	GetSnatIp() *string
	SetSnatIps(v []*string) *DescribeSnatTableEntriesRequest
	GetSnatIps() []*string
	SetSourceCIDR(v string) *DescribeSnatTableEntriesRequest
	GetSourceCIDR() *string
}

type DescribeSnatTableEntriesRequest struct {
	// The ID of the Network Address Translation (NAT) gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// nat-5tawjw5j7sgd2deujxuk0****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The page number. Pages start from page **1**.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The maximum value is **100**.
	//
	// Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the SNAT entry.
	//
	// example:
	//
	// snat-5tfjp36fsrb36zs36faj0****
	SnatEntryId *string `json:"SnatEntryId,omitempty" xml:"SnatEntryId,omitempty"`
	// The name of the SNAT entry.
	//
	// example:
	//
	// test0
	SnatEntryName *string `json:"SnatEntryName,omitempty" xml:"SnatEntryName,omitempty"`
	// The elastic IP address (EIP) specified in the SNAT entry.
	//
	// example:
	//
	// 58.XXXX.XXX.29
	SnatIp *string `json:"SnatIp,omitempty" xml:"SnatIp,omitempty"`
	// The information about the EIP specified in the SNAT entry.
	SnatIps []*string `json:"SnatIps,omitempty" xml:"SnatIps,omitempty" type:"Repeated"`
	// The source CIDR block specified in the SNAT entry.
	//
	// example:
	//
	// 10.1.0.50/32
	SourceCIDR *string `json:"SourceCIDR,omitempty" xml:"SourceCIDR,omitempty"`
}

func (s DescribeSnatTableEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnatTableEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnatTableEntriesRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeSnatTableEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSnatTableEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSnatTableEntriesRequest) GetSnatEntryId() *string {
	return s.SnatEntryId
}

func (s *DescribeSnatTableEntriesRequest) GetSnatEntryName() *string {
	return s.SnatEntryName
}

func (s *DescribeSnatTableEntriesRequest) GetSnatIp() *string {
	return s.SnatIp
}

func (s *DescribeSnatTableEntriesRequest) GetSnatIps() []*string {
	return s.SnatIps
}

func (s *DescribeSnatTableEntriesRequest) GetSourceCIDR() *string {
	return s.SourceCIDR
}

func (s *DescribeSnatTableEntriesRequest) SetNatGatewayId(v string) *DescribeSnatTableEntriesRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetPageNumber(v int32) *DescribeSnatTableEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetPageSize(v int32) *DescribeSnatTableEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetSnatEntryId(v string) *DescribeSnatTableEntriesRequest {
	s.SnatEntryId = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetSnatEntryName(v string) *DescribeSnatTableEntriesRequest {
	s.SnatEntryName = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetSnatIp(v string) *DescribeSnatTableEntriesRequest {
	s.SnatIp = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetSnatIps(v []*string) *DescribeSnatTableEntriesRequest {
	s.SnatIps = v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetSourceCIDR(v string) *DescribeSnatTableEntriesRequest {
	s.SourceCIDR = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) Validate() error {
	return dara.Validate(s)
}
