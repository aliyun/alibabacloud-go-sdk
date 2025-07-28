// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainInfo(v *CreateDomainResponseBodyDomainInfo) *CreateDomainResponseBody
	GetDomainInfo() *CreateDomainResponseBodyDomainInfo
	SetRequestId(v string) *CreateDomainResponseBody
	GetRequestId() *string
}

type CreateDomainResponseBody struct {
	// The information about the domain name that is added.
	DomainInfo *CreateDomainResponseBodyDomainInfo `json:"DomainInfo,omitempty" xml:"DomainInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDomainResponseBody) GetDomainInfo() *CreateDomainResponseBodyDomainInfo {
	return s.DomainInfo
}

func (s *CreateDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDomainResponseBody) SetDomainInfo(v *CreateDomainResponseBodyDomainInfo) *CreateDomainResponseBody {
	s.DomainInfo = v
	return s
}

func (s *CreateDomainResponseBody) SetRequestId(v string) *CreateDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDomainResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateDomainResponseBodyDomainInfo struct {
	// The CNAME that is assigned by WAF to the domain name.
	//
	// example:
	//
	// xxxxxwww.****.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The domain name that you added to WAF.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The name of the protected object that is generated.
	//
	// example:
	//
	// www.aliyundoc.com-waf
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
}

func (s CreateDomainResponseBodyDomainInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainResponseBodyDomainInfo) GoString() string {
	return s.String()
}

func (s *CreateDomainResponseBodyDomainInfo) GetCname() *string {
	return s.Cname
}

func (s *CreateDomainResponseBodyDomainInfo) GetDomain() *string {
	return s.Domain
}

func (s *CreateDomainResponseBodyDomainInfo) GetDomainId() *string {
	return s.DomainId
}

func (s *CreateDomainResponseBodyDomainInfo) SetCname(v string) *CreateDomainResponseBodyDomainInfo {
	s.Cname = &v
	return s
}

func (s *CreateDomainResponseBodyDomainInfo) SetDomain(v string) *CreateDomainResponseBodyDomainInfo {
	s.Domain = &v
	return s
}

func (s *CreateDomainResponseBodyDomainInfo) SetDomainId(v string) *CreateDomainResponseBodyDomainInfo {
	s.DomainId = &v
	return s
}

func (s *CreateDomainResponseBodyDomainInfo) Validate() error {
	return dara.Validate(s)
}
