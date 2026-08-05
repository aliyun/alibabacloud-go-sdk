// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCapabilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCapabilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCapabilityResponse
	GetStatusCode() *int32
	SetBody(v *CreateCapabilityResponseBody) *CreateCapabilityResponse
	GetBody() *CreateCapabilityResponseBody
}

type CreateCapabilityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCapabilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCapabilityResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCapabilityResponse) GoString() string {
	return s.String()
}

func (s *CreateCapabilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCapabilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCapabilityResponse) GetBody() *CreateCapabilityResponseBody {
	return s.Body
}

func (s *CreateCapabilityResponse) SetHeaders(v map[string]*string) *CreateCapabilityResponse {
	s.Headers = v
	return s
}

func (s *CreateCapabilityResponse) SetStatusCode(v int32) *CreateCapabilityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCapabilityResponse) SetBody(v *CreateCapabilityResponseBody) *CreateCapabilityResponse {
	s.Body = v
	return s
}

func (s *CreateCapabilityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
