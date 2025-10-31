// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateAnycastEipAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateAnycastEipAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateAnycastEipAddressResponse
	GetStatusCode() *int32
	SetBody(v *AllocateAnycastEipAddressResponseBody) *AllocateAnycastEipAddressResponse
	GetBody() *AllocateAnycastEipAddressResponseBody
}

type AllocateAnycastEipAddressResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateAnycastEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateAnycastEipAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateAnycastEipAddressResponse) GoString() string {
	return s.String()
}

func (s *AllocateAnycastEipAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateAnycastEipAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateAnycastEipAddressResponse) GetBody() *AllocateAnycastEipAddressResponseBody {
	return s.Body
}

func (s *AllocateAnycastEipAddressResponse) SetHeaders(v map[string]*string) *AllocateAnycastEipAddressResponse {
	s.Headers = v
	return s
}

func (s *AllocateAnycastEipAddressResponse) SetStatusCode(v int32) *AllocateAnycastEipAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateAnycastEipAddressResponse) SetBody(v *AllocateAnycastEipAddressResponseBody) *AllocateAnycastEipAddressResponse {
	s.Body = v
	return s
}

func (s *AllocateAnycastEipAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
