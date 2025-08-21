// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeForwardTableEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExternalIp(v string) *DescribeForwardTableEntriesRequest
	GetExternalIp() *string
	SetForwardEntryId(v string) *DescribeForwardTableEntriesRequest
	GetForwardEntryId() *string
	SetForwardEntryName(v string) *DescribeForwardTableEntriesRequest
	GetForwardEntryName() *string
	SetInternalIp(v string) *DescribeForwardTableEntriesRequest
	GetInternalIp() *string
	SetIpProtocol(v string) *DescribeForwardTableEntriesRequest
	GetIpProtocol() *string
	SetNatGatewayId(v string) *DescribeForwardTableEntriesRequest
	GetNatGatewayId() *string
	SetPageNumber(v int32) *DescribeForwardTableEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeForwardTableEntriesRequest
	GetPageSize() *int32
}

type DescribeForwardTableEntriesRequest struct {
	// The EIP in the DNAT entry. The public IP address is used to access the Internet.
	//
	// example:
	//
	// 36.XXX.XXX.72
	ExternalIp *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	// The ID of the DNAT entry.
	//
	// example:
	//
	// fwd-5tfi6f0rutmd00xrhkag7****
	ForwardEntryId *string `json:"ForwardEntryId,omitempty" xml:"ForwardEntryId,omitempty"`
	// The name of the DNAT entry.
	//
	// example:
	//
	// test0
	ForwardEntryName *string `json:"ForwardEntryName,omitempty" xml:"ForwardEntryName,omitempty"`
	// The private IP address of the instance that uses the DNAT entry for Internet communication.
	//
	// example:
	//
	// 10.XXX.XXX.50
	InternalIp *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	// The protocol. Valid values:
	//
	// 	- **TCP**: forwards TCP packets.
	//
	// 	- **UDP**: forwards UDP packets.
	//
	// 	- **Any**: forwards all packets.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The ID of the NAT gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// nat-5t7nh1cfm6kxiszlttr38****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The page number. Pages start from page **1**.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **100**.
	//
	// Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeForwardTableEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeForwardTableEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeForwardTableEntriesRequest) GetExternalIp() *string {
	return s.ExternalIp
}

func (s *DescribeForwardTableEntriesRequest) GetForwardEntryId() *string {
	return s.ForwardEntryId
}

func (s *DescribeForwardTableEntriesRequest) GetForwardEntryName() *string {
	return s.ForwardEntryName
}

func (s *DescribeForwardTableEntriesRequest) GetInternalIp() *string {
	return s.InternalIp
}

func (s *DescribeForwardTableEntriesRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribeForwardTableEntriesRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeForwardTableEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeForwardTableEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeForwardTableEntriesRequest) SetExternalIp(v string) *DescribeForwardTableEntriesRequest {
	s.ExternalIp = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetForwardEntryId(v string) *DescribeForwardTableEntriesRequest {
	s.ForwardEntryId = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetForwardEntryName(v string) *DescribeForwardTableEntriesRequest {
	s.ForwardEntryName = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetInternalIp(v string) *DescribeForwardTableEntriesRequest {
	s.InternalIp = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetIpProtocol(v string) *DescribeForwardTableEntriesRequest {
	s.IpProtocol = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetNatGatewayId(v string) *DescribeForwardTableEntriesRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetPageNumber(v int32) *DescribeForwardTableEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetPageSize(v int32) *DescribeForwardTableEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) Validate() error {
	return dara.Validate(s)
}
