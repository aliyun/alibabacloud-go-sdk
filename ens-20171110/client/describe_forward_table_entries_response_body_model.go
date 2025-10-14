// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeForwardTableEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForwardTableEntries(v []*DescribeForwardTableEntriesResponseBodyForwardTableEntries) *DescribeForwardTableEntriesResponseBody
	GetForwardTableEntries() []*DescribeForwardTableEntriesResponseBodyForwardTableEntries
	SetPageNumber(v string) *DescribeForwardTableEntriesResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeForwardTableEntriesResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeForwardTableEntriesResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeForwardTableEntriesResponseBody
	GetTotalCount() *string
}

type DescribeForwardTableEntriesResponseBody struct {
	// Details of DNAT entries.
	ForwardTableEntries []*DescribeForwardTableEntriesResponseBodyForwardTableEntries `json:"ForwardTableEntries,omitempty" xml:"ForwardTableEntries,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeForwardTableEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeForwardTableEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeForwardTableEntriesResponseBody) GetForwardTableEntries() []*DescribeForwardTableEntriesResponseBodyForwardTableEntries {
	return s.ForwardTableEntries
}

func (s *DescribeForwardTableEntriesResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeForwardTableEntriesResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeForwardTableEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeForwardTableEntriesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeForwardTableEntriesResponseBody) SetForwardTableEntries(v []*DescribeForwardTableEntriesResponseBodyForwardTableEntries) *DescribeForwardTableEntriesResponseBody {
	s.ForwardTableEntries = v
	return s
}

func (s *DescribeForwardTableEntriesResponseBody) SetPageNumber(v string) *DescribeForwardTableEntriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBody) SetPageSize(v string) *DescribeForwardTableEntriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBody) SetRequestId(v string) *DescribeForwardTableEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBody) SetTotalCount(v string) *DescribeForwardTableEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBody) Validate() error {
	if s.ForwardTableEntries != nil {
		for _, item := range s.ForwardTableEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeForwardTableEntriesResponseBodyForwardTableEntries struct {
	// The EIP in the DNAT entry. The public IP address is used to access the Internet.
	//
	// example:
	//
	// 120.XXX.XXX.70
	ExternalIp *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	// The external port or port range that is used in port forwarding.
	//
	// example:
	//
	// 22
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// The ID of the DNAT entry.
	//
	// example:
	//
	// fwd-5tf66679oi2uoxcvlg0g2****
	ForwardEntryId *string `json:"ForwardEntryId,omitempty" xml:"ForwardEntryId,omitempty"`
	// The name of the DNAT entry.
	//
	// example:
	//
	// test0
	ForwardEntryName *string `json:"ForwardEntryName,omitempty" xml:"ForwardEntryName,omitempty"`
	// The probe port of DNAT.
	//
	// example:
	//
	// 80
	HealthCheckPort *string `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	// The private IP address of the instance that uses the DNAT entry for Internet communication.
	//
	// example:
	//
	// 10.XXX.XXX.3
	InternalIp *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	// The internal port or port range that is used for port forwarding.
	//
	// example:
	//
	// 22
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
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
	// Any
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The ID of the NAT gateway.
	//
	// example:
	//
	// nat-5tawjw5j7sgd2deujxuk0****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The secondary EIP that is used to access the Internet.
	//
	// example:
	//
	// 101.XXX.XXX.7
	StandbyExternalIp *string `json:"StandbyExternalIp,omitempty" xml:"StandbyExternalIp,omitempty"`
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
	// The status of the DNAT entry. Valid values:
	//
	// 	- Pending: The DNAT entry is being created or modified.
	//
	// 	- Available: The DNAT entry is available.
	//
	// 	- Deleting: The DNAT entry is being deleted.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeForwardTableEntriesResponseBodyForwardTableEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeForwardTableEntriesResponseBodyForwardTableEntries) GoString() string {
	return s.String()
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) GetExternalIp() *string {
	return s.ExternalIp
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) GetForwardEntryId() *string {
	return s.ForwardEntryId
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) GetForwardEntryName() *string {
	return s.ForwardEntryName
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) GetHealthCheckPort() *string {
	return s.HealthCheckPort
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) GetInternalIp() *string {
	return s.InternalIp
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) GetInternalPort() *string {
	return s.InternalPort
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) GetStandbyExternalIp() *string {
	return s.StandbyExternalIp
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) GetStandbyStatus() *string {
	return s.StandbyStatus
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) GetStatus() *string {
	return s.Status
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) SetExternalIp(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntries {
	s.ExternalIp = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) SetExternalPort(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntries {
	s.ExternalPort = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) SetForwardEntryId(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntries {
	s.ForwardEntryId = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) SetForwardEntryName(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntries {
	s.ForwardEntryName = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) SetHealthCheckPort(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntries {
	s.HealthCheckPort = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) SetInternalIp(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntries {
	s.InternalIp = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) SetInternalPort(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntries {
	s.InternalPort = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) SetIpProtocol(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntries {
	s.IpProtocol = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) SetNatGatewayId(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntries {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) SetStandbyExternalIp(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntries {
	s.StandbyExternalIp = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) SetStandbyStatus(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntries {
	s.StandbyStatus = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) SetStatus(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntries {
	s.Status = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) Validate() error {
	return dara.Validate(s)
}
