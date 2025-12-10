// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAddressResponse
	GetStatusCode() *int32
	SetBody(v *GetAddressResponseBody) *GetAddressResponse
	GetBody() *GetAddressResponseBody
}

type GetAddressResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAddressResponse) GoString() string {
	return s.String()
}

func (s *GetAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAddressResponse) GetBody() *GetAddressResponseBody {
	return s.Body
}

func (s *GetAddressResponse) SetHeaders(v map[string]*string) *GetAddressResponse {
	s.Headers = v
	return s
}

func (s *GetAddressResponse) SetStatusCode(v int32) *GetAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAddressResponse) SetBody(v *GetAddressResponseBody) *GetAddressResponse {
	s.Body = v
	return s
}

func (s *GetAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
