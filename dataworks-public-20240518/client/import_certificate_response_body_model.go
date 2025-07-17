// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ImportCertificateResponseBody
	GetId() *int64
	SetRequestId(v string) *ImportCertificateResponseBody
	GetRequestId() *string
}

type ImportCertificateResponseBody struct {
	// example:
	//
	// 676303114031776
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 7C352CB7-CD88-50CF-9D0D-E81BDF020E7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ImportCertificateResponseBody) GetId() *int64 {
	return s.Id
}

func (s *ImportCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportCertificateResponseBody) SetId(v int64) *ImportCertificateResponseBody {
	s.Id = &v
	return s
}

func (s *ImportCertificateResponseBody) SetRequestId(v string) *ImportCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
