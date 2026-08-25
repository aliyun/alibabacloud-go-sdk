// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveExternalSAMLIdPCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveExternalSAMLIdPCertificateResponseBody
	GetRequestId() *string
}

type RemoveExternalSAMLIdPCertificateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 400979BC-92EC-58B9-B47C-6913BD56A6FD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveExternalSAMLIdPCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveExternalSAMLIdPCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveExternalSAMLIdPCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveExternalSAMLIdPCertificateResponseBody) SetRequestId(v string) *RemoveExternalSAMLIdPCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveExternalSAMLIdPCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
