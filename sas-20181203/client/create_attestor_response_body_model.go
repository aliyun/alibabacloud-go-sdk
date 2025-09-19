// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAttestorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAttestorResponseBody
	GetRequestId() *string
}

type CreateAttestorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B44EA7F0-497A-5F10-B5A8-87291356****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAttestorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAttestorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAttestorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAttestorResponseBody) SetRequestId(v string) *CreateAttestorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAttestorResponseBody) Validate() error {
	return dara.Validate(s)
}
