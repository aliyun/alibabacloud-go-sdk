// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoginTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoginTicket(v *DescribeLoginTicketResponseBodyLoginTicket) *DescribeLoginTicketResponseBody
	GetLoginTicket() *DescribeLoginTicketResponseBodyLoginTicket
	SetRequestId(v string) *DescribeLoginTicketResponseBody
	GetRequestId() *string
}

type DescribeLoginTicketResponseBody struct {
	LoginTicket *DescribeLoginTicketResponseBodyLoginTicket `json:"LoginTicket,omitempty" xml:"LoginTicket,omitempty" type:"Struct"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLoginTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoginTicketResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoginTicketResponseBody) GetLoginTicket() *DescribeLoginTicketResponseBodyLoginTicket {
	return s.LoginTicket
}

func (s *DescribeLoginTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoginTicketResponseBody) SetLoginTicket(v *DescribeLoginTicketResponseBodyLoginTicket) *DescribeLoginTicketResponseBody {
	s.LoginTicket = v
	return s
}

func (s *DescribeLoginTicketResponseBody) SetRequestId(v string) *DescribeLoginTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoginTicketResponseBody) Validate() error {
	if s.LoginTicket != nil {
		if err := s.LoginTicket.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLoginTicketResponseBodyLoginTicket struct {
	Certificate *string                                            `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	Ticket      *string                                            `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
	Zones       []*DescribeLoginTicketResponseBodyLoginTicketZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeLoginTicketResponseBodyLoginTicket) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoginTicketResponseBodyLoginTicket) GoString() string {
	return s.String()
}

func (s *DescribeLoginTicketResponseBodyLoginTicket) GetCertificate() *string {
	return s.Certificate
}

func (s *DescribeLoginTicketResponseBodyLoginTicket) GetTicket() *string {
	return s.Ticket
}

func (s *DescribeLoginTicketResponseBodyLoginTicket) GetZones() []*DescribeLoginTicketResponseBodyLoginTicketZones {
	return s.Zones
}

func (s *DescribeLoginTicketResponseBodyLoginTicket) SetCertificate(v string) *DescribeLoginTicketResponseBodyLoginTicket {
	s.Certificate = &v
	return s
}

func (s *DescribeLoginTicketResponseBodyLoginTicket) SetTicket(v string) *DescribeLoginTicketResponseBodyLoginTicket {
	s.Ticket = &v
	return s
}

func (s *DescribeLoginTicketResponseBodyLoginTicket) SetZones(v []*DescribeLoginTicketResponseBodyLoginTicketZones) *DescribeLoginTicketResponseBodyLoginTicket {
	s.Zones = v
	return s
}

func (s *DescribeLoginTicketResponseBodyLoginTicket) Validate() error {
	if s.Zones != nil {
		for _, item := range s.Zones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLoginTicketResponseBodyLoginTicketZones struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeLoginTicketResponseBodyLoginTicketZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoginTicketResponseBodyLoginTicketZones) GoString() string {
	return s.String()
}

func (s *DescribeLoginTicketResponseBodyLoginTicketZones) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeLoginTicketResponseBodyLoginTicketZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeLoginTicketResponseBodyLoginTicketZones) SetLocalName(v string) *DescribeLoginTicketResponseBodyLoginTicketZones {
	s.LocalName = &v
	return s
}

func (s *DescribeLoginTicketResponseBodyLoginTicketZones) SetZoneId(v string) *DescribeLoginTicketResponseBodyLoginTicketZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeLoginTicketResponseBodyLoginTicketZones) Validate() error {
	return dara.Validate(s)
}
