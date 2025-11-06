// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDispatchConsoleAPIForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLiveToken(v string) *DispatchConsoleAPIForPartnerRequest
	GetLiveToken() *string
	SetOperation(v string) *DispatchConsoleAPIForPartnerRequest
	GetOperation() *string
	SetParams(v string) *DispatchConsoleAPIForPartnerRequest
	GetParams() *string
	SetProduct(v string) *DispatchConsoleAPIForPartnerRequest
	GetProduct() *string
	SetSiteHost(v string) *DispatchConsoleAPIForPartnerRequest
	GetSiteHost() *string
}

type DispatchConsoleAPIForPartnerRequest struct {
	// This parameter is required.
	LiveToken *string `json:"LiveToken,omitempty" xml:"LiveToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// disable
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// [\\"\\"]
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbaseue
	Product  *string `json:"Product,omitempty" xml:"Product,omitempty"`
	SiteHost *string `json:"SiteHost,omitempty" xml:"SiteHost,omitempty"`
}

func (s DispatchConsoleAPIForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s DispatchConsoleAPIForPartnerRequest) GoString() string {
	return s.String()
}

func (s *DispatchConsoleAPIForPartnerRequest) GetLiveToken() *string {
	return s.LiveToken
}

func (s *DispatchConsoleAPIForPartnerRequest) GetOperation() *string {
	return s.Operation
}

func (s *DispatchConsoleAPIForPartnerRequest) GetParams() *string {
	return s.Params
}

func (s *DispatchConsoleAPIForPartnerRequest) GetProduct() *string {
	return s.Product
}

func (s *DispatchConsoleAPIForPartnerRequest) GetSiteHost() *string {
	return s.SiteHost
}

func (s *DispatchConsoleAPIForPartnerRequest) SetLiveToken(v string) *DispatchConsoleAPIForPartnerRequest {
	s.LiveToken = &v
	return s
}

func (s *DispatchConsoleAPIForPartnerRequest) SetOperation(v string) *DispatchConsoleAPIForPartnerRequest {
	s.Operation = &v
	return s
}

func (s *DispatchConsoleAPIForPartnerRequest) SetParams(v string) *DispatchConsoleAPIForPartnerRequest {
	s.Params = &v
	return s
}

func (s *DispatchConsoleAPIForPartnerRequest) SetProduct(v string) *DispatchConsoleAPIForPartnerRequest {
	s.Product = &v
	return s
}

func (s *DispatchConsoleAPIForPartnerRequest) SetSiteHost(v string) *DispatchConsoleAPIForPartnerRequest {
	s.SiteHost = &v
	return s
}

func (s *DispatchConsoleAPIForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
