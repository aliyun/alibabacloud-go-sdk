// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UploadCertificateResponseBody
	GetRequestId() *string
}

type UploadCertificateResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 15a735a1-8fe6-45cc-a64c-3c4ff839334e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *UploadCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadCertificateResponseBody) SetRequestId(v string) *UploadCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
