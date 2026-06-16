// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCaInstanceCrlAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCaInstanceCrlAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCaInstanceCrlAddressResponse
	GetStatusCode() *int32
	SetBody(v *GetCaInstanceCrlAddressResponseBody) *GetCaInstanceCrlAddressResponse
	GetBody() *GetCaInstanceCrlAddressResponseBody
}

type GetCaInstanceCrlAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCaInstanceCrlAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCaInstanceCrlAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCaInstanceCrlAddressResponse) GoString() string {
	return s.String()
}

func (s *GetCaInstanceCrlAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCaInstanceCrlAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCaInstanceCrlAddressResponse) GetBody() *GetCaInstanceCrlAddressResponseBody {
	return s.Body
}

func (s *GetCaInstanceCrlAddressResponse) SetHeaders(v map[string]*string) *GetCaInstanceCrlAddressResponse {
	s.Headers = v
	return s
}

func (s *GetCaInstanceCrlAddressResponse) SetStatusCode(v int32) *GetCaInstanceCrlAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCaInstanceCrlAddressResponse) SetBody(v *GetCaInstanceCrlAddressResponseBody) *GetCaInstanceCrlAddressResponse {
	s.Body = v
	return s
}

func (s *GetCaInstanceCrlAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
