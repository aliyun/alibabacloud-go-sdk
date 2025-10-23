// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInvalidAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryInvalidAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryInvalidAddressResponse
	GetStatusCode() *int32
	SetBody(v *QueryInvalidAddressResponseBody) *QueryInvalidAddressResponse
	GetBody() *QueryInvalidAddressResponseBody
}

type QueryInvalidAddressResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInvalidAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInvalidAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryInvalidAddressResponse) GoString() string {
	return s.String()
}

func (s *QueryInvalidAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryInvalidAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryInvalidAddressResponse) GetBody() *QueryInvalidAddressResponseBody {
	return s.Body
}

func (s *QueryInvalidAddressResponse) SetHeaders(v map[string]*string) *QueryInvalidAddressResponse {
	s.Headers = v
	return s
}

func (s *QueryInvalidAddressResponse) SetStatusCode(v int32) *QueryInvalidAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInvalidAddressResponse) SetBody(v *QueryInvalidAddressResponseBody) *QueryInvalidAddressResponse {
	s.Body = v
	return s
}

func (s *QueryInvalidAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
