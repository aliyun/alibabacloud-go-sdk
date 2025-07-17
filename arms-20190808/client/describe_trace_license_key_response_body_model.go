// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTraceLicenseKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLicenseKey(v string) *DescribeTraceLicenseKeyResponseBody
	GetLicenseKey() *string
	SetRequestId(v string) *DescribeTraceLicenseKeyResponseBody
	GetRequestId() *string
}

type DescribeTraceLicenseKeyResponseBody struct {
	// The license key for the application.
	//
	// example:
	//
	// b590lhguqs@3a75d95f218****
	LicenseKey *string `json:"LicenseKey,omitempty" xml:"LicenseKey,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 29053944-6FE5-4240-8927-10095ECE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTraceLicenseKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceLicenseKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTraceLicenseKeyResponseBody) GetLicenseKey() *string {
	return s.LicenseKey
}

func (s *DescribeTraceLicenseKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTraceLicenseKeyResponseBody) SetLicenseKey(v string) *DescribeTraceLicenseKeyResponseBody {
	s.LicenseKey = &v
	return s
}

func (s *DescribeTraceLicenseKeyResponseBody) SetRequestId(v string) *DescribeTraceLicenseKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTraceLicenseKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
