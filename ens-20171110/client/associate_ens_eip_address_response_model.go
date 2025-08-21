// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateEnsEipAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateEnsEipAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateEnsEipAddressResponse
	GetStatusCode() *int32
	SetBody(v *AssociateEnsEipAddressResponseBody) *AssociateEnsEipAddressResponse
	GetBody() *AssociateEnsEipAddressResponseBody
}

type AssociateEnsEipAddressResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateEnsEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateEnsEipAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateEnsEipAddressResponse) GoString() string {
	return s.String()
}

func (s *AssociateEnsEipAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateEnsEipAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateEnsEipAddressResponse) GetBody() *AssociateEnsEipAddressResponseBody {
	return s.Body
}

func (s *AssociateEnsEipAddressResponse) SetHeaders(v map[string]*string) *AssociateEnsEipAddressResponse {
	s.Headers = v
	return s
}

func (s *AssociateEnsEipAddressResponse) SetStatusCode(v int32) *AssociateEnsEipAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateEnsEipAddressResponse) SetBody(v *AssociateEnsEipAddressResponseBody) *AssociateEnsEipAddressResponse {
	s.Body = v
	return s
}

func (s *AssociateEnsEipAddressResponse) Validate() error {
	return dara.Validate(s)
}
