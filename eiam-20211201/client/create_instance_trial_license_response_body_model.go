// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceTrialLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLicenseId(v string) *CreateInstanceTrialLicenseResponseBody
	GetLicenseId() *string
	SetRequestId(v string) *CreateInstanceTrialLicenseResponseBody
	GetRequestId() *string
}

type CreateInstanceTrialLicenseResponseBody struct {
	// example:
	//
	// license_463hfmewi2njxxxx
	LicenseId *string `json:"LicenseId,omitempty" xml:"LicenseId,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceTrialLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceTrialLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceTrialLicenseResponseBody) GetLicenseId() *string {
	return s.LicenseId
}

func (s *CreateInstanceTrialLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceTrialLicenseResponseBody) SetLicenseId(v string) *CreateInstanceTrialLicenseResponseBody {
	s.LicenseId = &v
	return s
}

func (s *CreateInstanceTrialLicenseResponseBody) SetRequestId(v string) *CreateInstanceTrialLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceTrialLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}
