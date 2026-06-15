// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateEmailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainPart(v string) *ValidateEmailResponseBody
	GetDomainPart() *string
	SetIsFreeMail(v bool) *ValidateEmailResponseBody
	GetIsFreeMail() *bool
	SetLocalPart(v string) *ValidateEmailResponseBody
	GetLocalPart() *string
	SetProvider(v string) *ValidateEmailResponseBody
	GetProvider() *string
	SetRequestId(v string) *ValidateEmailResponseBody
	GetRequestId() *string
	SetStatus(v string) *ValidateEmailResponseBody
	GetStatus() *string
	SetSubStatus(v string) *ValidateEmailResponseBody
	GetSubStatus() *string
}

type ValidateEmailResponseBody struct {
	// The domain part of the email address parsed from the syntax check. The domain part is converted to lowercase.
	//
	// example:
	//
	// yyy.com
	DomainPart *string `json:"DomainPart,omitempty" xml:"DomainPart,omitempty"`
	// Indicates whether the address is from a free email service.
	//
	// example:
	//
	// true
	IsFreeMail *bool `json:"IsFreeMail,omitempty" xml:"IsFreeMail,omitempty"`
	// The local part of the email address parsed from the syntax check. The local part is converted to lowercase and the content after the plus sign (+) is removed.
	//
	// example:
	//
	// xxx
	LocalPart *string `json:"LocalPart,omitempty" xml:"LocalPart,omitempty"`
	// The email service provider of the address.
	//
	// example:
	//
	// Gmail
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The validation status of the email address.
	//
	// This parameter is required.
	//
	// example:
	//
	// VALID
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The detailed validation status of the email address. This provides more information about the Status.
	//
	// This parameter is required.
	//
	// example:
	//
	// UNSPECIFIED
	SubStatus *string `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
}

func (s ValidateEmailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateEmailResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateEmailResponseBody) GetDomainPart() *string {
	return s.DomainPart
}

func (s *ValidateEmailResponseBody) GetIsFreeMail() *bool {
	return s.IsFreeMail
}

func (s *ValidateEmailResponseBody) GetLocalPart() *string {
	return s.LocalPart
}

func (s *ValidateEmailResponseBody) GetProvider() *string {
	return s.Provider
}

func (s *ValidateEmailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateEmailResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ValidateEmailResponseBody) GetSubStatus() *string {
	return s.SubStatus
}

func (s *ValidateEmailResponseBody) SetDomainPart(v string) *ValidateEmailResponseBody {
	s.DomainPart = &v
	return s
}

func (s *ValidateEmailResponseBody) SetIsFreeMail(v bool) *ValidateEmailResponseBody {
	s.IsFreeMail = &v
	return s
}

func (s *ValidateEmailResponseBody) SetLocalPart(v string) *ValidateEmailResponseBody {
	s.LocalPart = &v
	return s
}

func (s *ValidateEmailResponseBody) SetProvider(v string) *ValidateEmailResponseBody {
	s.Provider = &v
	return s
}

func (s *ValidateEmailResponseBody) SetRequestId(v string) *ValidateEmailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateEmailResponseBody) SetStatus(v string) *ValidateEmailResponseBody {
	s.Status = &v
	return s
}

func (s *ValidateEmailResponseBody) SetSubStatus(v string) *ValidateEmailResponseBody {
	s.SubStatus = &v
	return s
}

func (s *ValidateEmailResponseBody) Validate() error {
	return dara.Validate(s)
}
