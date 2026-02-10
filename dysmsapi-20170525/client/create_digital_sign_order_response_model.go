// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDigitalSignOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDigitalSignOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDigitalSignOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateDigitalSignOrderResponseBody) *CreateDigitalSignOrderResponse
	GetBody() *CreateDigitalSignOrderResponseBody
}

type CreateDigitalSignOrderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDigitalSignOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDigitalSignOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalSignOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateDigitalSignOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDigitalSignOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDigitalSignOrderResponse) GetBody() *CreateDigitalSignOrderResponseBody {
	return s.Body
}

func (s *CreateDigitalSignOrderResponse) SetHeaders(v map[string]*string) *CreateDigitalSignOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateDigitalSignOrderResponse) SetStatusCode(v int32) *CreateDigitalSignOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDigitalSignOrderResponse) SetBody(v *CreateDigitalSignOrderResponseBody) *CreateDigitalSignOrderResponse {
	s.Body = v
	return s
}

func (s *CreateDigitalSignOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
