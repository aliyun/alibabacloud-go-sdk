// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMailAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMailAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMailAddressResponse
	GetStatusCode() *int32
	SetBody(v *CreateMailAddressResponseBody) *CreateMailAddressResponse
	GetBody() *CreateMailAddressResponseBody
}

type CreateMailAddressResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMailAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMailAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMailAddressResponse) GoString() string {
	return s.String()
}

func (s *CreateMailAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMailAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMailAddressResponse) GetBody() *CreateMailAddressResponseBody {
	return s.Body
}

func (s *CreateMailAddressResponse) SetHeaders(v map[string]*string) *CreateMailAddressResponse {
	s.Headers = v
	return s
}

func (s *CreateMailAddressResponse) SetStatusCode(v int32) *CreateMailAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMailAddressResponse) SetBody(v *CreateMailAddressResponseBody) *CreateMailAddressResponse {
	s.Body = v
	return s
}

func (s *CreateMailAddressResponse) Validate() error {
	return dara.Validate(s)
}
