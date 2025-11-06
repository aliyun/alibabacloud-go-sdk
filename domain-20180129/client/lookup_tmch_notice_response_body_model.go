// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLookupTmchNoticeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClaims(v *LookupTmchNoticeResponseBodyClaims) *LookupTmchNoticeResponseBody
	GetClaims() *LookupTmchNoticeResponseBodyClaims
	SetId(v int64) *LookupTmchNoticeResponseBody
	GetId() *int64
	SetLabel(v string) *LookupTmchNoticeResponseBody
	GetLabel() *string
	SetNotAfter(v string) *LookupTmchNoticeResponseBody
	GetNotAfter() *string
	SetNotBefore(v string) *LookupTmchNoticeResponseBody
	GetNotBefore() *string
	SetRequestId(v string) *LookupTmchNoticeResponseBody
	GetRequestId() *string
}

type LookupTmchNoticeResponseBody struct {
	Claims *LookupTmchNoticeResponseBodyClaims `json:"Claims,omitempty" xml:"Claims,omitempty" type:"Struct"`
	// example:
	//
	// 586608000000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// noted
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// 2018-10-15T00:00:00.0Z
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// example:
	//
	// 2018-10-13T00:00:00.0Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// example:
	//
	// 01C10C8E-0468-468C-BCD9-E709BDD0AE8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LookupTmchNoticeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LookupTmchNoticeResponseBody) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBody) GetClaims() *LookupTmchNoticeResponseBodyClaims {
	return s.Claims
}

func (s *LookupTmchNoticeResponseBody) GetId() *int64 {
	return s.Id
}

func (s *LookupTmchNoticeResponseBody) GetLabel() *string {
	return s.Label
}

func (s *LookupTmchNoticeResponseBody) GetNotAfter() *string {
	return s.NotAfter
}

func (s *LookupTmchNoticeResponseBody) GetNotBefore() *string {
	return s.NotBefore
}

func (s *LookupTmchNoticeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LookupTmchNoticeResponseBody) SetClaims(v *LookupTmchNoticeResponseBodyClaims) *LookupTmchNoticeResponseBody {
	s.Claims = v
	return s
}

func (s *LookupTmchNoticeResponseBody) SetId(v int64) *LookupTmchNoticeResponseBody {
	s.Id = &v
	return s
}

func (s *LookupTmchNoticeResponseBody) SetLabel(v string) *LookupTmchNoticeResponseBody {
	s.Label = &v
	return s
}

func (s *LookupTmchNoticeResponseBody) SetNotAfter(v string) *LookupTmchNoticeResponseBody {
	s.NotAfter = &v
	return s
}

func (s *LookupTmchNoticeResponseBody) SetNotBefore(v string) *LookupTmchNoticeResponseBody {
	s.NotBefore = &v
	return s
}

func (s *LookupTmchNoticeResponseBody) SetRequestId(v string) *LookupTmchNoticeResponseBody {
	s.RequestId = &v
	return s
}

func (s *LookupTmchNoticeResponseBody) Validate() error {
	if s.Claims != nil {
		if err := s.Claims.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LookupTmchNoticeResponseBodyClaims struct {
	Claim []*LookupTmchNoticeResponseBodyClaimsClaim `json:"Claim,omitempty" xml:"Claim,omitempty" type:"Repeated"`
}

func (s LookupTmchNoticeResponseBodyClaims) String() string {
	return dara.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaims) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaims) GetClaim() []*LookupTmchNoticeResponseBodyClaimsClaim {
	return s.Claim
}

func (s *LookupTmchNoticeResponseBodyClaims) SetClaim(v []*LookupTmchNoticeResponseBodyClaimsClaim) *LookupTmchNoticeResponseBodyClaims {
	s.Claim = v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaims) Validate() error {
	if s.Claim != nil {
		for _, item := range s.Claim {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type LookupTmchNoticeResponseBodyClaimsClaim struct {
	ClassDescs *LookupTmchNoticeResponseBodyClaimsClaimClassDescs `json:"ClassDescs,omitempty" xml:"ClassDescs,omitempty" type:"Struct"`
	Contacts   *LookupTmchNoticeResponseBodyClaimsClaimContacts   `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Struct"`
	// example:
	//
	// Class 9: Calculators; bags, coverings,containers, carriers and holders for mobile phones, personal handheld computers and notebooks; headphones; speakers; blank storage media;batteries. Class 16: Paper
	GoodsAndServices *string                                         `json:"GoodsAndServices,omitempty" xml:"GoodsAndServices,omitempty"`
	Holders          *LookupTmchNoticeResponseBodyClaimsClaimHolders `json:"Holders,omitempty" xml:"Holders,omitempty" type:"Struct"`
	JurDesc          *LookupTmchNoticeResponseBodyClaimsClaimJurDesc `json:"JurDesc,omitempty" xml:"JurDesc,omitempty" type:"Struct"`
	// example:
	//
	// POTED
	MarkName *string `json:"MarkName,omitempty" xml:"MarkName,omitempty"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaim) String() string {
	return dara.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaim) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaim) GetClassDescs() *LookupTmchNoticeResponseBodyClaimsClaimClassDescs {
	return s.ClassDescs
}

func (s *LookupTmchNoticeResponseBodyClaimsClaim) GetContacts() *LookupTmchNoticeResponseBodyClaimsClaimContacts {
	return s.Contacts
}

func (s *LookupTmchNoticeResponseBodyClaimsClaim) GetGoodsAndServices() *string {
	return s.GoodsAndServices
}

func (s *LookupTmchNoticeResponseBodyClaimsClaim) GetHolders() *LookupTmchNoticeResponseBodyClaimsClaimHolders {
	return s.Holders
}

func (s *LookupTmchNoticeResponseBodyClaimsClaim) GetJurDesc() *LookupTmchNoticeResponseBodyClaimsClaimJurDesc {
	return s.JurDesc
}

func (s *LookupTmchNoticeResponseBodyClaimsClaim) GetMarkName() *string {
	return s.MarkName
}

func (s *LookupTmchNoticeResponseBodyClaimsClaim) SetClassDescs(v *LookupTmchNoticeResponseBodyClaimsClaimClassDescs) *LookupTmchNoticeResponseBodyClaimsClaim {
	s.ClassDescs = v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaim) SetContacts(v *LookupTmchNoticeResponseBodyClaimsClaimContacts) *LookupTmchNoticeResponseBodyClaimsClaim {
	s.Contacts = v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaim) SetGoodsAndServices(v string) *LookupTmchNoticeResponseBodyClaimsClaim {
	s.GoodsAndServices = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaim) SetHolders(v *LookupTmchNoticeResponseBodyClaimsClaimHolders) *LookupTmchNoticeResponseBodyClaimsClaim {
	s.Holders = v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaim) SetJurDesc(v *LookupTmchNoticeResponseBodyClaimsClaimJurDesc) *LookupTmchNoticeResponseBodyClaimsClaim {
	s.JurDesc = v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaim) SetMarkName(v string) *LookupTmchNoticeResponseBodyClaimsClaim {
	s.MarkName = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaim) Validate() error {
	if s.ClassDescs != nil {
		if err := s.ClassDescs.Validate(); err != nil {
			return err
		}
	}
	if s.Contacts != nil {
		if err := s.Contacts.Validate(); err != nil {
			return err
		}
	}
	if s.Holders != nil {
		if err := s.Holders.Validate(); err != nil {
			return err
		}
	}
	if s.JurDesc != nil {
		if err := s.JurDesc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LookupTmchNoticeResponseBodyClaimsClaimClassDescs struct {
	ClassDesc []*LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc `json:"ClassDesc,omitempty" xml:"ClassDesc,omitempty" type:"Repeated"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimClassDescs) String() string {
	return dara.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimClassDescs) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimClassDescs) GetClassDesc() []*LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc {
	return s.ClassDesc
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimClassDescs) SetClassDesc(v []*LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc) *LookupTmchNoticeResponseBodyClaimsClaimClassDescs {
	s.ClassDesc = v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimClassDescs) Validate() error {
	if s.ClassDesc != nil {
		for _, item := range s.ClassDesc {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc struct {
	// example:
	//
	// 18
	ClassNum *int32 `json:"ClassNum,omitempty" xml:"ClassNum,omitempty"`
	// example:
	//
	// New Zealand
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc) String() string {
	return dara.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc) GetClassNum() *int32 {
	return s.ClassNum
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc) GetDesc() *string {
	return s.Desc
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc) SetClassNum(v int32) *LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc {
	s.ClassNum = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc) SetDesc(v string) *LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc {
	s.Desc = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc) Validate() error {
	return dara.Validate(s)
}

type LookupTmchNoticeResponseBodyClaimsClaimContacts struct {
	Contact []*LookupTmchNoticeResponseBodyClaimsClaimContactsContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Repeated"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimContacts) String() string {
	return dara.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimContacts) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContacts) GetContact() []*LookupTmchNoticeResponseBodyClaimsClaimContactsContact {
	return s.Contact
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContacts) SetContact(v []*LookupTmchNoticeResponseBodyClaimsClaimContactsContact) *LookupTmchNoticeResponseBodyClaimsClaimContacts {
	s.Contact = v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContacts) Validate() error {
	if s.Contact != nil {
		for _, item := range s.Contact {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type LookupTmchNoticeResponseBodyClaimsClaimContactsContact struct {
	Addr *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr `json:"Addr,omitempty" xml:"Addr,omitempty" type:"Struct"`
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 4472335**8
	Fax *string `json:"Fax,omitempty" xml:"Fax,omitempty"`
	// example:
	//
	// Tom
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Tom
	Org *string `json:"Org,omitempty" xml:"Org,omitempty"`
	// example:
	//
	// agent
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1390000****
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimContactsContact) String() string {
	return dara.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimContactsContact) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) GetAddr() *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr {
	return s.Addr
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) GetEmail() *string {
	return s.Email
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) GetFax() *string {
	return s.Fax
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) GetName() *string {
	return s.Name
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) GetOrg() *string {
	return s.Org
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) GetType() *string {
	return s.Type
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) GetVoice() *string {
	return s.Voice
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) SetAddr(v *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) *LookupTmchNoticeResponseBodyClaimsClaimContactsContact {
	s.Addr = v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) SetEmail(v string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContact {
	s.Email = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) SetFax(v string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContact {
	s.Fax = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) SetName(v string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContact {
	s.Name = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) SetOrg(v string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContact {
	s.Org = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) SetType(v string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContact {
	s.Type = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) SetVoice(v string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContact {
	s.Voice = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) Validate() error {
	if s.Addr != nil {
		if err := s.Addr.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr struct {
	// example:
	//
	// NZ
	Cc *string `json:"Cc,omitempty" xml:"Cc,omitempty"`
	// example:
	//
	// Auckland
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// 1010
	Pc *string `json:"Pc,omitempty" xml:"Pc,omitempty"`
	// example:
	//
	// Auckland
	Sp     *string                                                           `json:"Sp,omitempty" xml:"Sp,omitempty"`
	Street *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddrStreet `json:"Street,omitempty" xml:"Street,omitempty" type:"Struct"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) String() string {
	return dara.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) GetCc() *string {
	return s.Cc
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) GetCity() *string {
	return s.City
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) GetPc() *string {
	return s.Pc
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) GetSp() *string {
	return s.Sp
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) GetStreet() *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddrStreet {
	return s.Street
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) SetCc(v string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr {
	s.Cc = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) SetCity(v string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr {
	s.City = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) SetPc(v string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr {
	s.Pc = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) SetSp(v string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr {
	s.Sp = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) SetStreet(v *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddrStreet) *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr {
	s.Street = v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) Validate() error {
	if s.Street != nil {
		if err := s.Street.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddrStreet struct {
	Street []*string `json:"Street,omitempty" xml:"Street,omitempty" type:"Repeated"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddrStreet) String() string {
	return dara.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddrStreet) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddrStreet) GetStreet() []*string {
	return s.Street
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddrStreet) SetStreet(v []*string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddrStreet {
	s.Street = v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddrStreet) Validate() error {
	return dara.Validate(s)
}

type LookupTmchNoticeResponseBodyClaimsClaimHolders struct {
	Holder []*LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder `json:"Holder,omitempty" xml:"Holder,omitempty" type:"Repeated"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimHolders) String() string {
	return dara.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimHolders) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHolders) GetHolder() []*LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder {
	return s.Holder
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHolders) SetHolder(v []*LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder) *LookupTmchNoticeResponseBodyClaimsClaimHolders {
	s.Holder = v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHolders) Validate() error {
	if s.Holder != nil {
		for _, item := range s.Holder {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder struct {
	Addr *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr `json:"Addr,omitempty" xml:"Addr,omitempty" type:"Struct"`
	// example:
	//
	// owner
	Entitlement *string `json:"Entitlement,omitempty" xml:"Entitlement,omitempty"`
	// example:
	//
	// Whitcoulls 2011 Limited
	Org *string `json:"Org,omitempty" xml:"Org,omitempty"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder) String() string {
	return dara.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder) GetAddr() *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr {
	return s.Addr
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder) GetEntitlement() *string {
	return s.Entitlement
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder) GetOrg() *string {
	return s.Org
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder) SetAddr(v *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder {
	s.Addr = v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder) SetEntitlement(v string) *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder {
	s.Entitlement = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder) SetOrg(v string) *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder {
	s.Org = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder) Validate() error {
	if s.Addr != nil {
		if err := s.Addr.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr struct {
	// example:
	//
	// NZ
	Cc *string `json:"Cc,omitempty" xml:"Cc,omitempty"`
	// example:
	//
	// Wellington
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// 6011
	Pc *string `json:"Pc,omitempty" xml:"Pc,omitempty"`
	// example:
	//
	// Wellington
	Sp     *string                                                         `json:"Sp,omitempty" xml:"Sp,omitempty"`
	Street *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddrStreet `json:"Street,omitempty" xml:"Street,omitempty" type:"Struct"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) String() string {
	return dara.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) GetCc() *string {
	return s.Cc
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) GetCity() *string {
	return s.City
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) GetPc() *string {
	return s.Pc
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) GetSp() *string {
	return s.Sp
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) GetStreet() *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddrStreet {
	return s.Street
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) SetCc(v string) *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr {
	s.Cc = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) SetCity(v string) *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr {
	s.City = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) SetPc(v string) *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr {
	s.Pc = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) SetSp(v string) *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr {
	s.Sp = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) SetStreet(v *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddrStreet) *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr {
	s.Street = v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) Validate() error {
	if s.Street != nil {
		if err := s.Street.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddrStreet struct {
	Street []*string `json:"Street,omitempty" xml:"Street,omitempty" type:"Repeated"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddrStreet) String() string {
	return dara.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddrStreet) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddrStreet) GetStreet() []*string {
	return s.Street
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddrStreet) SetStreet(v []*string) *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddrStreet {
	s.Street = v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddrStreet) Validate() error {
	return dara.Validate(s)
}

type LookupTmchNoticeResponseBodyClaimsClaimJurDesc struct {
	// example:
	//
	// New Zealand
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// example:
	//
	// NZ
	JurCC *string `json:"JurCC,omitempty" xml:"JurCC,omitempty"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimJurDesc) String() string {
	return dara.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimJurDesc) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimJurDesc) GetDesc() *string {
	return s.Desc
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimJurDesc) GetJurCC() *string {
	return s.JurCC
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimJurDesc) SetDesc(v string) *LookupTmchNoticeResponseBodyClaimsClaimJurDesc {
	s.Desc = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimJurDesc) SetJurCC(v string) *LookupTmchNoticeResponseBodyClaimsClaimJurDesc {
	s.JurCC = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimJurDesc) Validate() error {
	return dara.Validate(s)
}
