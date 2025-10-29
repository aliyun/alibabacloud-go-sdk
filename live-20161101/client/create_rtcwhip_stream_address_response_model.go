// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRTCWhipStreamAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRTCWhipStreamAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRTCWhipStreamAddressResponse
	GetStatusCode() *int32
	SetBody(v *CreateRTCWhipStreamAddressResponseBody) *CreateRTCWhipStreamAddressResponse
	GetBody() *CreateRTCWhipStreamAddressResponseBody
}

type CreateRTCWhipStreamAddressResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRTCWhipStreamAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRTCWhipStreamAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRTCWhipStreamAddressResponse) GoString() string {
	return s.String()
}

func (s *CreateRTCWhipStreamAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRTCWhipStreamAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRTCWhipStreamAddressResponse) GetBody() *CreateRTCWhipStreamAddressResponseBody {
	return s.Body
}

func (s *CreateRTCWhipStreamAddressResponse) SetHeaders(v map[string]*string) *CreateRTCWhipStreamAddressResponse {
	s.Headers = v
	return s
}

func (s *CreateRTCWhipStreamAddressResponse) SetStatusCode(v int32) *CreateRTCWhipStreamAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRTCWhipStreamAddressResponse) SetBody(v *CreateRTCWhipStreamAddressResponseBody) *CreateRTCWhipStreamAddressResponse {
	s.Body = v
	return s
}

func (s *CreateRTCWhipStreamAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
