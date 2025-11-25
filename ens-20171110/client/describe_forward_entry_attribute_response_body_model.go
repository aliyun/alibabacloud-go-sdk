// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeForwardEntryAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreationTime(v string) *DescribeForwardEntryAttributeResponseBody
	GetCreationTime() *string
	SetExternalIp(v string) *DescribeForwardEntryAttributeResponseBody
	GetExternalIp() *string
	SetExternalPort(v string) *DescribeForwardEntryAttributeResponseBody
	GetExternalPort() *string
	SetForwardEntryId(v string) *DescribeForwardEntryAttributeResponseBody
	GetForwardEntryId() *string
	SetForwardEntryName(v string) *DescribeForwardEntryAttributeResponseBody
	GetForwardEntryName() *string
	SetHealthCheckPort(v string) *DescribeForwardEntryAttributeResponseBody
	GetHealthCheckPort() *string
	SetInternalIp(v string) *DescribeForwardEntryAttributeResponseBody
	GetInternalIp() *string
	SetInternalPort(v string) *DescribeForwardEntryAttributeResponseBody
	GetInternalPort() *string
	SetIpProtocol(v string) *DescribeForwardEntryAttributeResponseBody
	GetIpProtocol() *string
	SetNatGatewayId(v string) *DescribeForwardEntryAttributeResponseBody
	GetNatGatewayId() *string
	SetRequestId(v string) *DescribeForwardEntryAttributeResponseBody
	GetRequestId() *string
	SetStandbyExternalIp(v string) *DescribeForwardEntryAttributeResponseBody
	GetStandbyExternalIp() *string
	SetStandbyStatus(v string) *DescribeForwardEntryAttributeResponseBody
	GetStandbyStatus() *string
	SetStatus(v string) *DescribeForwardEntryAttributeResponseBody
	GetStatus() *string
}

type DescribeForwardEntryAttributeResponseBody struct {
	// example:
	//
	// 2020-04-26T15:38:27Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// 36.XXX.XXX.72
	ExternalIp *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	// example:
	//
	// 22
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// example:
	//
	// fwd-5tfi6f0rutmd00xrhkag7****
	ForwardEntryId *string `json:"ForwardEntryId,omitempty" xml:"ForwardEntryId,omitempty"`
	// example:
	//
	// test0
	ForwardEntryName *string `json:"ForwardEntryName,omitempty" xml:"ForwardEntryName,omitempty"`
	// example:
	//
	// 80
	HealthCheckPort *string `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	// example:
	//
	// 10.XXX.XXX.50
	InternalIp *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	// example:
	//
	// 22
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	// example:
	//
	// Any
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// example:
	//
	// nat-5t7nh1cfm6kxiszlttr38****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// example:
	//
	// 6666C5A5-75ED-422E-A022-7121FA18C968
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 101.XXX.XXX.4
	StandbyExternalIp *string `json:"StandbyExternalIp,omitempty" xml:"StandbyExternalIp,omitempty"`
	// example:
	//
	// Stopped
	StandbyStatus *string `json:"StandbyStatus,omitempty" xml:"StandbyStatus,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeForwardEntryAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeForwardEntryAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeForwardEntryAttributeResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeForwardEntryAttributeResponseBody) GetExternalIp() *string {
	return s.ExternalIp
}

func (s *DescribeForwardEntryAttributeResponseBody) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *DescribeForwardEntryAttributeResponseBody) GetForwardEntryId() *string {
	return s.ForwardEntryId
}

func (s *DescribeForwardEntryAttributeResponseBody) GetForwardEntryName() *string {
	return s.ForwardEntryName
}

func (s *DescribeForwardEntryAttributeResponseBody) GetHealthCheckPort() *string {
	return s.HealthCheckPort
}

func (s *DescribeForwardEntryAttributeResponseBody) GetInternalIp() *string {
	return s.InternalIp
}

func (s *DescribeForwardEntryAttributeResponseBody) GetInternalPort() *string {
	return s.InternalPort
}

func (s *DescribeForwardEntryAttributeResponseBody) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribeForwardEntryAttributeResponseBody) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeForwardEntryAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeForwardEntryAttributeResponseBody) GetStandbyExternalIp() *string {
	return s.StandbyExternalIp
}

func (s *DescribeForwardEntryAttributeResponseBody) GetStandbyStatus() *string {
	return s.StandbyStatus
}

func (s *DescribeForwardEntryAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeForwardEntryAttributeResponseBody) SetCreationTime(v string) *DescribeForwardEntryAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeForwardEntryAttributeResponseBody) SetExternalIp(v string) *DescribeForwardEntryAttributeResponseBody {
	s.ExternalIp = &v
	return s
}

func (s *DescribeForwardEntryAttributeResponseBody) SetExternalPort(v string) *DescribeForwardEntryAttributeResponseBody {
	s.ExternalPort = &v
	return s
}

func (s *DescribeForwardEntryAttributeResponseBody) SetForwardEntryId(v string) *DescribeForwardEntryAttributeResponseBody {
	s.ForwardEntryId = &v
	return s
}

func (s *DescribeForwardEntryAttributeResponseBody) SetForwardEntryName(v string) *DescribeForwardEntryAttributeResponseBody {
	s.ForwardEntryName = &v
	return s
}

func (s *DescribeForwardEntryAttributeResponseBody) SetHealthCheckPort(v string) *DescribeForwardEntryAttributeResponseBody {
	s.HealthCheckPort = &v
	return s
}

func (s *DescribeForwardEntryAttributeResponseBody) SetInternalIp(v string) *DescribeForwardEntryAttributeResponseBody {
	s.InternalIp = &v
	return s
}

func (s *DescribeForwardEntryAttributeResponseBody) SetInternalPort(v string) *DescribeForwardEntryAttributeResponseBody {
	s.InternalPort = &v
	return s
}

func (s *DescribeForwardEntryAttributeResponseBody) SetIpProtocol(v string) *DescribeForwardEntryAttributeResponseBody {
	s.IpProtocol = &v
	return s
}

func (s *DescribeForwardEntryAttributeResponseBody) SetNatGatewayId(v string) *DescribeForwardEntryAttributeResponseBody {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeForwardEntryAttributeResponseBody) SetRequestId(v string) *DescribeForwardEntryAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeForwardEntryAttributeResponseBody) SetStandbyExternalIp(v string) *DescribeForwardEntryAttributeResponseBody {
	s.StandbyExternalIp = &v
	return s
}

func (s *DescribeForwardEntryAttributeResponseBody) SetStandbyStatus(v string) *DescribeForwardEntryAttributeResponseBody {
	s.StandbyStatus = &v
	return s
}

func (s *DescribeForwardEntryAttributeResponseBody) SetStatus(v string) *DescribeForwardEntryAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeForwardEntryAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
