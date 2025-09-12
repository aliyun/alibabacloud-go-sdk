// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDomainInfoForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDomainInfoForPartnerResponseBodyData) *GetDomainInfoForPartnerResponseBody
	GetData() *GetDomainInfoForPartnerResponseBodyData
	SetRequestId(v string) *GetDomainInfoForPartnerResponseBody
	GetRequestId() *string
}

type GetDomainInfoForPartnerResponseBody struct {
	Data *GetDomainInfoForPartnerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDomainInfoForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDomainInfoForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainInfoForPartnerResponseBody) GetData() *GetDomainInfoForPartnerResponseBodyData {
	return s.Data
}

func (s *GetDomainInfoForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDomainInfoForPartnerResponseBody) SetData(v *GetDomainInfoForPartnerResponseBodyData) *GetDomainInfoForPartnerResponseBody {
	s.Data = v
	return s
}

func (s *GetDomainInfoForPartnerResponseBody) SetRequestId(v string) *GetDomainInfoForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDomainInfoForPartnerResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDomainInfoForPartnerResponseBodyData struct {
	// example:
	//
	// playnew-alilive.daotantan.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// NS3.ALIYUN.COM,NS4.ALIYUN.COM
	NameServers *string                                           `json:"NameServers,omitempty" xml:"NameServers,omitempty"`
	Ownership   *GetDomainInfoForPartnerResponseBodyDataOwnership `json:"Ownership,omitempty" xml:"Ownership,omitempty" type:"Struct"`
	// example:
	//
	// aliyun
	Registrar *string `json:"Registrar,omitempty" xml:"Registrar,omitempty"`
}

func (s GetDomainInfoForPartnerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDomainInfoForPartnerResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDomainInfoForPartnerResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *GetDomainInfoForPartnerResponseBodyData) GetNameServers() *string {
	return s.NameServers
}

func (s *GetDomainInfoForPartnerResponseBodyData) GetOwnership() *GetDomainInfoForPartnerResponseBodyDataOwnership {
	return s.Ownership
}

func (s *GetDomainInfoForPartnerResponseBodyData) GetRegistrar() *string {
	return s.Registrar
}

func (s *GetDomainInfoForPartnerResponseBodyData) SetDomainName(v string) *GetDomainInfoForPartnerResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *GetDomainInfoForPartnerResponseBodyData) SetNameServers(v string) *GetDomainInfoForPartnerResponseBodyData {
	s.NameServers = &v
	return s
}

func (s *GetDomainInfoForPartnerResponseBodyData) SetOwnership(v *GetDomainInfoForPartnerResponseBodyDataOwnership) *GetDomainInfoForPartnerResponseBodyData {
	s.Ownership = v
	return s
}

func (s *GetDomainInfoForPartnerResponseBodyData) SetRegistrar(v string) *GetDomainInfoForPartnerResponseBodyData {
	s.Registrar = &v
	return s
}

func (s *GetDomainInfoForPartnerResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDomainInfoForPartnerResponseBodyDataOwnership struct {
	// example:
	//
	// 1189245564569485
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// example:
	//
	// pai
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
}

func (s GetDomainInfoForPartnerResponseBodyDataOwnership) String() string {
	return dara.Prettify(s)
}

func (s GetDomainInfoForPartnerResponseBodyDataOwnership) GoString() string {
	return s.String()
}

func (s *GetDomainInfoForPartnerResponseBodyDataOwnership) GetAccount() *string {
	return s.Account
}

func (s *GetDomainInfoForPartnerResponseBodyDataOwnership) GetProvider() *string {
	return s.Provider
}

func (s *GetDomainInfoForPartnerResponseBodyDataOwnership) SetAccount(v string) *GetDomainInfoForPartnerResponseBodyDataOwnership {
	s.Account = &v
	return s
}

func (s *GetDomainInfoForPartnerResponseBodyDataOwnership) SetProvider(v string) *GetDomainInfoForPartnerResponseBodyDataOwnership {
	s.Provider = &v
	return s
}

func (s *GetDomainInfoForPartnerResponseBodyDataOwnership) Validate() error {
	return dara.Validate(s)
}
