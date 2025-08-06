// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *CheckLicenseRequest
	GetBusinessType() *string
	SetFilter(v string) *CheckLicenseRequest
	GetFilter() *string
	SetNonce(v string) *CheckLicenseRequest
	GetNonce() *string
	SetSign(v string) *CheckLicenseRequest
	GetSign() *string
	SetTime(v string) *CheckLicenseRequest
	GetTime() *string
}

type CheckLicenseRequest struct {
	// This parameter is required.
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// This parameter is required.
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// This parameter is required.
	Nonce *string `json:"Nonce,omitempty" xml:"Nonce,omitempty"`
	// This parameter is required.
	Sign *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
	// This parameter is required.
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s CheckLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckLicenseRequest) GoString() string {
	return s.String()
}

func (s *CheckLicenseRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *CheckLicenseRequest) GetFilter() *string {
	return s.Filter
}

func (s *CheckLicenseRequest) GetNonce() *string {
	return s.Nonce
}

func (s *CheckLicenseRequest) GetSign() *string {
	return s.Sign
}

func (s *CheckLicenseRequest) GetTime() *string {
	return s.Time
}

func (s *CheckLicenseRequest) SetBusinessType(v string) *CheckLicenseRequest {
	s.BusinessType = &v
	return s
}

func (s *CheckLicenseRequest) SetFilter(v string) *CheckLicenseRequest {
	s.Filter = &v
	return s
}

func (s *CheckLicenseRequest) SetNonce(v string) *CheckLicenseRequest {
	s.Nonce = &v
	return s
}

func (s *CheckLicenseRequest) SetSign(v string) *CheckLicenseRequest {
	s.Sign = &v
	return s
}

func (s *CheckLicenseRequest) SetTime(v string) *CheckLicenseRequest {
	s.Time = &v
	return s
}

func (s *CheckLicenseRequest) Validate() error {
	return dara.Validate(s)
}
