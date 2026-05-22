// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSiteNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CheckSiteNameResponseBody
	GetDescription() *string
	SetIsSubSite(v bool) *CheckSiteNameResponseBody
	GetIsSubSite() *bool
	SetMesseage(v string) *CheckSiteNameResponseBody
	GetMesseage() *string
	SetPassed(v bool) *CheckSiteNameResponseBody
	GetPassed() *bool
	SetRequestId(v string) *CheckSiteNameResponseBody
	GetRequestId() *string
}

type CheckSiteNameResponseBody struct {
	// The description of the verification result. Valid values:
	//
	// 	- **success**: The verification is successful.
	//
	// 	- **Site.AlreadyExist**: The website domain name has already been added.
	//
	// 	- **Site.InvalidName**: Invalid website domain name.
	//
	// 	- **Site.SubSiteUnavailable**: Subdomains are not allowed.
	//
	// 	- **Site.InternalError**: An internal error occurs.
	//
	// example:
	//
	// success
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether a subdomain is specified. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IsSubSite *bool `json:"IsSubSite,omitempty" xml:"IsSubSite,omitempty"`
	// The verification message.
	Messeage *string `json:"Messeage,omitempty" xml:"Messeage,omitempty"`
	// Indicates whether the verification passed.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Passed *bool `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-280B-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckSiteNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckSiteNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSiteNameResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CheckSiteNameResponseBody) GetIsSubSite() *bool {
	return s.IsSubSite
}

func (s *CheckSiteNameResponseBody) GetMesseage() *string {
	return s.Messeage
}

func (s *CheckSiteNameResponseBody) GetPassed() *bool {
	return s.Passed
}

func (s *CheckSiteNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckSiteNameResponseBody) SetDescription(v string) *CheckSiteNameResponseBody {
	s.Description = &v
	return s
}

func (s *CheckSiteNameResponseBody) SetIsSubSite(v bool) *CheckSiteNameResponseBody {
	s.IsSubSite = &v
	return s
}

func (s *CheckSiteNameResponseBody) SetMesseage(v string) *CheckSiteNameResponseBody {
	s.Messeage = &v
	return s
}

func (s *CheckSiteNameResponseBody) SetPassed(v bool) *CheckSiteNameResponseBody {
	s.Passed = &v
	return s
}

func (s *CheckSiteNameResponseBody) SetRequestId(v string) *CheckSiteNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckSiteNameResponseBody) Validate() error {
	return dara.Validate(s)
}
