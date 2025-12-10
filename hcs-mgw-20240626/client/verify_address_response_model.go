// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyAddressResponse
	GetStatusCode() *int32
	SetBody(v *VerifyAddressResponseBody) *VerifyAddressResponse
	GetBody() *VerifyAddressResponseBody
}

type VerifyAddressResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyAddressResponse) GoString() string {
	return s.String()
}

func (s *VerifyAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyAddressResponse) GetBody() *VerifyAddressResponseBody {
	return s.Body
}

func (s *VerifyAddressResponse) SetHeaders(v map[string]*string) *VerifyAddressResponse {
	s.Headers = v
	return s
}

func (s *VerifyAddressResponse) SetStatusCode(v int32) *VerifyAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyAddressResponse) SetBody(v *VerifyAddressResponseBody) *VerifyAddressResponse {
	s.Body = v
	return s
}

func (s *VerifyAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
