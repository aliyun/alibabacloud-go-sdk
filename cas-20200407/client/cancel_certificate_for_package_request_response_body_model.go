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
	// The request ID.
	//
	// example:
	//
	// 082FAB35-6AB9-4FD5-8750-D36673548E76
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
