// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetApAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetApAddressResponse
	GetStatusCode() *int32
	SetBody(v *SetApAddressResponseBody) *SetApAddressResponse
	GetBody() *SetApAddressResponseBody
}

type SetApAddressResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetApAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetApAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s SetApAddressResponse) GoString() string {
	return s.String()
}

func (s *SetApAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetApAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetApAddressResponse) GetBody() *SetApAddressResponseBody {
	return s.Body
}

func (s *SetApAddressResponse) SetHeaders(v map[string]*string) *SetApAddressResponse {
	s.Headers = v
	return s
}

func (s *SetApAddressResponse) SetStatusCode(v int32) *SetApAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *SetApAddressResponse) SetBody(v *SetApAddressResponseBody) *SetApAddressResponse {
	s.Body = v
	return s
}

func (s *SetApAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
