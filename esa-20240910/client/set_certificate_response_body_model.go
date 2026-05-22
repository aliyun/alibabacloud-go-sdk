// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *SetCertificateResponseBody
	GetId() *string
	SetRequestId(v string) *SetCertificateResponseBody
	GetRequestId() *string
}

type SetCertificateResponseBody struct {
	// The certificate ID.
	//
	// example:
	//
	// babaffe176ae44e2ad16d3e309b9e506
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A666D44F-19D6-490E-97CF-1A64AB962C57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetCertificateResponseBody) GetId() *string {
	return s.Id
}

func (s *SetCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetCertificateResponseBody) SetId(v string) *SetCertificateResponseBody {
	s.Id = &v
	return s
}

func (s *SetCertificateResponseBody) SetRequestId(v string) *SetCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
