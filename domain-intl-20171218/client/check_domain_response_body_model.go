// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvail(v string) *CheckDomainResponseBody
	GetAvail() *string
	SetDomainName(v string) *CheckDomainResponseBody
	GetDomainName() *string
	SetDynamicCheck(v bool) *CheckDomainResponseBody
	GetDynamicCheck() *bool
	SetPremium(v string) *CheckDomainResponseBody
	GetPremium() *string
	SetPrice(v int64) *CheckDomainResponseBody
	GetPrice() *int64
	SetReason(v string) *CheckDomainResponseBody
	GetReason() *string
	SetRequestId(v string) *CheckDomainResponseBody
	GetRequestId() *string
}

type CheckDomainResponseBody struct {
	// example:
	//
	// 1
	Avail *string `json:"Avail,omitempty" xml:"Avail,omitempty"`
	// example:
	//
	// nvgtntights.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// false
	DynamicCheck *bool `json:"DynamicCheck,omitempty" xml:"DynamicCheck,omitempty"`
	// example:
	//
	// true
	Premium *string `json:"Premium,omitempty" xml:"Premium,omitempty"`
	// example:
	//
	// 2000
	Price *int64 `json:"Price,omitempty" xml:"Price,omitempty"`
	// example:
	//
	// In use
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// AF7D4DCE-0776-47F2-A9B2-6FB85A87AA60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDomainResponseBody) GetAvail() *string {
	return s.Avail
}

func (s *CheckDomainResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *CheckDomainResponseBody) GetDynamicCheck() *bool {
	return s.DynamicCheck
}

func (s *CheckDomainResponseBody) GetPremium() *string {
	return s.Premium
}

func (s *CheckDomainResponseBody) GetPrice() *int64 {
	return s.Price
}

func (s *CheckDomainResponseBody) GetReason() *string {
	return s.Reason
}

func (s *CheckDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDomainResponseBody) SetAvail(v string) *CheckDomainResponseBody {
	s.Avail = &v
	return s
}

func (s *CheckDomainResponseBody) SetDomainName(v string) *CheckDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *CheckDomainResponseBody) SetDynamicCheck(v bool) *CheckDomainResponseBody {
	s.DynamicCheck = &v
	return s
}

func (s *CheckDomainResponseBody) SetPremium(v string) *CheckDomainResponseBody {
	s.Premium = &v
	return s
}

func (s *CheckDomainResponseBody) SetPrice(v int64) *CheckDomainResponseBody {
	s.Price = &v
	return s
}

func (s *CheckDomainResponseBody) SetReason(v string) *CheckDomainResponseBody {
	s.Reason = &v
	return s
}

func (s *CheckDomainResponseBody) SetRequestId(v string) *CheckDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
