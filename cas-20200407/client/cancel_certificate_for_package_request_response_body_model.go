// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCertificateForPackageRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelCertificateForPackageRequestResponseBody
	GetRequestId() *string
}

type CancelCertificateForPackageRequestResponseBody struct {
	// The unique ID of the request. Alibaba Cloud generates this ID for each request. Use this ID to troubleshoot issues.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelCertificateForPackageRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelCertificateForPackageRequestResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCertificateForPackageRequestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelCertificateForPackageRequestResponseBody) SetRequestId(v string) *CancelCertificateForPackageRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelCertificateForPackageRequestResponseBody) Validate() error {
	return dara.Validate(s)
}
