// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsFirstBbsTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCna(v string) *IsFirstBbsTicketRequest
	GetCna() *string
	SetDistributionChannel(v string) *IsFirstBbsTicketRequest
	GetDistributionChannel() *string
	SetReferrer(v string) *IsFirstBbsTicketRequest
	GetReferrer() *string
	SetSubDistributionChannel(v string) *IsFirstBbsTicketRequest
	GetSubDistributionChannel() *string
	SetXGatewayExtendInfo(v string) *IsFirstBbsTicketRequest
	GetXGatewayExtendInfo() *string
}

type IsFirstBbsTicketRequest struct {
	Cna                    *string `json:"Cna,omitempty" xml:"Cna,omitempty"`
	DistributionChannel    *string `json:"DistributionChannel,omitempty" xml:"DistributionChannel,omitempty"`
	Referrer               *string `json:"Referrer,omitempty" xml:"Referrer,omitempty"`
	SubDistributionChannel *string `json:"SubDistributionChannel,omitempty" xml:"SubDistributionChannel,omitempty"`
	XGatewayExtendInfo     *string `json:"XGatewayExtendInfo,omitempty" xml:"XGatewayExtendInfo,omitempty"`
}

func (s IsFirstBbsTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s IsFirstBbsTicketRequest) GoString() string {
	return s.String()
}

func (s *IsFirstBbsTicketRequest) GetCna() *string {
	return s.Cna
}

func (s *IsFirstBbsTicketRequest) GetDistributionChannel() *string {
	return s.DistributionChannel
}

func (s *IsFirstBbsTicketRequest) GetReferrer() *string {
	return s.Referrer
}

func (s *IsFirstBbsTicketRequest) GetSubDistributionChannel() *string {
	return s.SubDistributionChannel
}

func (s *IsFirstBbsTicketRequest) GetXGatewayExtendInfo() *string {
	return s.XGatewayExtendInfo
}

func (s *IsFirstBbsTicketRequest) SetCna(v string) *IsFirstBbsTicketRequest {
	s.Cna = &v
	return s
}

func (s *IsFirstBbsTicketRequest) SetDistributionChannel(v string) *IsFirstBbsTicketRequest {
	s.DistributionChannel = &v
	return s
}

func (s *IsFirstBbsTicketRequest) SetReferrer(v string) *IsFirstBbsTicketRequest {
	s.Referrer = &v
	return s
}

func (s *IsFirstBbsTicketRequest) SetSubDistributionChannel(v string) *IsFirstBbsTicketRequest {
	s.SubDistributionChannel = &v
	return s
}

func (s *IsFirstBbsTicketRequest) SetXGatewayExtendInfo(v string) *IsFirstBbsTicketRequest {
	s.XGatewayExtendInfo = &v
	return s
}

func (s *IsFirstBbsTicketRequest) Validate() error {
	return dara.Validate(s)
}
