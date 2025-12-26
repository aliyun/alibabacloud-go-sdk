// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClearAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClearAddressResponse
	GetStatusCode() *int32
	SetBody(v *ClearAddressResponseBody) *ClearAddressResponse
	GetBody() *ClearAddressResponseBody
}

type ClearAddressResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ClearAddressResponse) GoString() string {
	return s.String()
}

func (s *ClearAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClearAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClearAddressResponse) GetBody() *ClearAddressResponseBody {
	return s.Body
}

func (s *ClearAddressResponse) SetHeaders(v map[string]*string) *ClearAddressResponse {
	s.Headers = v
	return s
}

func (s *ClearAddressResponse) SetStatusCode(v int32) *ClearAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearAddressResponse) SetBody(v *ClearAddressResponseBody) *ClearAddressResponse {
	s.Body = v
	return s
}

func (s *ClearAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
