// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTicketTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *GetTicketTemplateRequest
	GetCategoryId() *int64
	SetCna(v string) *GetTicketTemplateRequest
	GetCna() *string
	SetDistributionChannel(v string) *GetTicketTemplateRequest
	GetDistributionChannel() *string
	SetReferrer(v string) *GetTicketTemplateRequest
	GetReferrer() *string
	SetSubDistributionChannel(v string) *GetTicketTemplateRequest
	GetSubDistributionChannel() *string
	SetXGatewayExtendInfo(v string) *GetTicketTemplateRequest
	GetXGatewayExtendInfo() *string
}

type GetTicketTemplateRequest struct {
	// This parameter is required.
	CategoryId             *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Cna                    *string `json:"Cna,omitempty" xml:"Cna,omitempty"`
	DistributionChannel    *string `json:"DistributionChannel,omitempty" xml:"DistributionChannel,omitempty"`
	Referrer               *string `json:"Referrer,omitempty" xml:"Referrer,omitempty"`
	SubDistributionChannel *string `json:"SubDistributionChannel,omitempty" xml:"SubDistributionChannel,omitempty"`
	XGatewayExtendInfo     *string `json:"XGatewayExtendInfo,omitempty" xml:"XGatewayExtendInfo,omitempty"`
}

func (s GetTicketTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTicketTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTicketTemplateRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *GetTicketTemplateRequest) GetCna() *string {
	return s.Cna
}

func (s *GetTicketTemplateRequest) GetDistributionChannel() *string {
	return s.DistributionChannel
}

func (s *GetTicketTemplateRequest) GetReferrer() *string {
	return s.Referrer
}

func (s *GetTicketTemplateRequest) GetSubDistributionChannel() *string {
	return s.SubDistributionChannel
}

func (s *GetTicketTemplateRequest) GetXGatewayExtendInfo() *string {
	return s.XGatewayExtendInfo
}

func (s *GetTicketTemplateRequest) SetCategoryId(v int64) *GetTicketTemplateRequest {
	s.CategoryId = &v
	return s
}

func (s *GetTicketTemplateRequest) SetCna(v string) *GetTicketTemplateRequest {
	s.Cna = &v
	return s
}

func (s *GetTicketTemplateRequest) SetDistributionChannel(v string) *GetTicketTemplateRequest {
	s.DistributionChannel = &v
	return s
}

func (s *GetTicketTemplateRequest) SetReferrer(v string) *GetTicketTemplateRequest {
	s.Referrer = &v
	return s
}

func (s *GetTicketTemplateRequest) SetSubDistributionChannel(v string) *GetTicketTemplateRequest {
	s.SubDistributionChannel = &v
	return s
}

func (s *GetTicketTemplateRequest) SetXGatewayExtendInfo(v string) *GetTicketTemplateRequest {
	s.XGatewayExtendInfo = &v
	return s
}

func (s *GetTicketTemplateRequest) Validate() error {
	return dara.Validate(s)
}
