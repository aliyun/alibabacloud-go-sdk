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
	SetMaxResults(v int32) *DescribeForwardTableEntriesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeForwardTableEntriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeForwardTableEntriesResponseBody
	GetRequestId() *string
}

type DescribeForwardTableEntriesResponseBody struct {
	ForwardTableEntries []*DescribeForwardTableEntriesResponseBodyForwardTableEntries `json:"ForwardTableEntries,omitempty" xml:"ForwardTableEntries,omitempty" type:"Repeated"`
	MaxResults          *int32                                                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken           *string                                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId           *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *DescribeForwardTableEntriesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeForwardTableEntriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeForwardTableEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeForwardTableEntriesResponseBody) SetForwardTableEntries(v []*DescribeForwardTableEntriesResponseBodyForwardTableEntries) *DescribeForwardTableEntriesResponseBody {
	s.ForwardTableEntries = v
	return s
}

func (s *DescribeForwardTableEntriesResponseBody) SetMaxResults(v int32) *DescribeForwardTableEntriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBody) SetNextToken(v string) *DescribeForwardTableEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBody) SetRequestId(v string) *DescribeForwardTableEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeForwardTableEntriesResponseBodyForwardTableEntries struct {
	ExternalIp       *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	ExternalPort     *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	ForwardEntryId   *string `json:"ForwardEntryId,omitempty" xml:"ForwardEntryId,omitempty"`
	ForwardEntryName *string `json:"ForwardEntryName,omitempty" xml:"ForwardEntryName,omitempty"`
	ForwardTableId   *string `json:"ForwardTableId,omitempty" xml:"ForwardTableId,omitempty"`
	InternalIp       *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	InternalPort     *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	IpProtocol       *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	NatGatewayId     *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) GetForwardTableId() *string {
	return s.ForwardTableId
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

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) SetForwardTableId(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntries {
	s.ForwardTableId = &v
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

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) SetStatus(v string) *DescribeForwardTableEntriesResponseBodyForwardTableEntries {
	s.Status = &v
	return s
}

func (s *DescribeForwardTableEntriesResponseBodyForwardTableEntries) Validate() error {
	return dara.Validate(s)
}
