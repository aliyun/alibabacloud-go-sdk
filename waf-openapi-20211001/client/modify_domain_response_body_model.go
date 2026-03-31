// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainInfo(v *ModifyDomainResponseBodyDomainInfo) *ModifyDomainResponseBody
	GetDomainInfo() *ModifyDomainResponseBodyDomainInfo
	SetRequestId(v string) *ModifyDomainResponseBody
	GetRequestId() *string
}

type ModifyDomainResponseBody struct {
	// The information about the domain name.
	DomainInfo *ModifyDomainResponseBodyDomainInfo `json:"DomainInfo,omitempty" xml:"DomainInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDomainResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDomainResponseBody) GetDomainInfo() *ModifyDomainResponseBodyDomainInfo {
	return s.DomainInfo
}

func (s *ModifyDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDomainResponseBody) SetDomainInfo(v *ModifyDomainResponseBodyDomainInfo) *ModifyDomainResponseBody {
	s.DomainInfo = v
	return s
}

func (s *ModifyDomainResponseBody) SetRequestId(v string) *ModifyDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDomainResponseBody) Validate() error {
	if s.DomainInfo != nil {
		if err := s.DomainInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyDomainResponseBodyDomainInfo struct {
	// The CNAME that is assigned by WAF to the domain name.
	//
	// example:
	//
	// xxxxxcvdaf.****.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The domain name whose access configurations you modified.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the domain name.
	//
	// example:
	//
	// www.aliyundoc.com-waf
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
}

func (s ModifyDomainResponseBodyDomainInfo) String() string {
	return dara.Prettify(s)
}

func (s ModifyDomainResponseBodyDomainInfo) GoString() string {
	return s.String()
}

func (s *ModifyDomainResponseBodyDomainInfo) GetCname() *string {
	return s.Cname
}

func (s *ModifyDomainResponseBodyDomainInfo) GetDomain() *string {
	return s.Domain
}

func (s *ModifyDomainResponseBodyDomainInfo) GetDomainId() *string {
	return s.DomainId
}

func (s *ModifyDomainResponseBodyDomainInfo) SetCname(v string) *ModifyDomainResponseBodyDomainInfo {
	s.Cname = &v
	return s
}

func (s *ModifyDomainResponseBodyDomainInfo) SetDomain(v string) *ModifyDomainResponseBodyDomainInfo {
	s.Domain = &v
	return s
}

func (s *ModifyDomainResponseBodyDomainInfo) SetDomainId(v string) *ModifyDomainResponseBodyDomainInfo {
	s.DomainId = &v
	return s
}

func (s *ModifyDomainResponseBodyDomainInfo) Validate() error {
	return dara.Validate(s)
}
