// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateAnycastEipAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateAnycastEipAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateAnycastEipAddressResponse
	GetStatusCode() *int32
	SetBody(v *AssociateAnycastEipAddressResponseBody) *AssociateAnycastEipAddressResponse
	GetBody() *AssociateAnycastEipAddressResponseBody
}

type AssociateAnycastEipAddressResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateAnycastEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateAnycastEipAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateAnycastEipAddressResponse) GoString() string {
	return s.String()
}

func (s *AssociateAnycastEipAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateAnycastEipAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateAnycastEipAddressResponse) GetBody() *AssociateAnycastEipAddressResponseBody {
	return s.Body
}

func (s *AssociateAnycastEipAddressResponse) SetHeaders(v map[string]*string) *AssociateAnycastEipAddressResponse {
	s.Headers = v
	return s
}

func (s *AssociateAnycastEipAddressResponse) SetStatusCode(v int32) *AssociateAnycastEipAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateAnycastEipAddressResponse) SetBody(v *AssociateAnycastEipAddressResponseBody) *AssociateAnycastEipAddressResponse {
	s.Body = v
	return s
}

func (s *AssociateAnycastEipAddressResponse) Validate() error {
	return dara.Validate(s)
}
