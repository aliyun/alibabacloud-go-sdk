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
	// The domain part of the email address parsed from syntax validation (lowercased).
	//
	// example:
	//
	// yyy.com
	DomainPart *string `json:"DomainPart,omitempty" xml:"DomainPart,omitempty"`
	// Indicates whether the address is a free mailbox.
	//
	// example:
	//
	// true
	IsFreeMail *bool `json:"IsFreeMail,omitempty" xml:"IsFreeMail,omitempty"`
	// The local part of the email address parsed from syntax validation (lowercased and with the plus-sign portion removed).
	//
	// example:
	//
	// xxx
	LocalPart *string `json:"LocalPart,omitempty" xml:"LocalPart,omitempty"`
	// The email provider classification of the address.
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
	// The email address status obtained from validation.
	//
	// This parameter is required.
	//
	// example:
	//
	// VALID
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The email address sub-status obtained from validation, which provides a detailed description of the status.
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
