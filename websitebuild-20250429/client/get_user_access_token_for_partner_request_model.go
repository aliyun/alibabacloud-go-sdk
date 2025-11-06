// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserAccessTokenForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteHost(v string) *GetUserAccessTokenForPartnerRequest
	GetSiteHost() *string
	SetTicket(v string) *GetUserAccessTokenForPartnerRequest
	GetTicket() *string
}

type GetUserAccessTokenForPartnerRequest struct {
	SiteHost *string `json:"SiteHost,omitempty" xml:"SiteHost,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2f68fe1e-d9d5-4803-94d0-2fc81032e91d
	Ticket *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s GetUserAccessTokenForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserAccessTokenForPartnerRequest) GoString() string {
	return s.String()
}

func (s *GetUserAccessTokenForPartnerRequest) GetSiteHost() *string {
	return s.SiteHost
}

func (s *GetUserAccessTokenForPartnerRequest) GetTicket() *string {
	return s.Ticket
}

func (s *GetUserAccessTokenForPartnerRequest) SetSiteHost(v string) *GetUserAccessTokenForPartnerRequest {
	s.SiteHost = &v
	return s
}

func (s *GetUserAccessTokenForPartnerRequest) SetTicket(v string) *GetUserAccessTokenForPartnerRequest {
	s.Ticket = &v
	return s
}

func (s *GetUserAccessTokenForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
