// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHostAvailabilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHostAvailabilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHostAvailabilityResponse
	GetStatusCode() *int32
	SetBody(v *CreateHostAvailabilityResponseBody) *CreateHostAvailabilityResponse
	GetBody() *CreateHostAvailabilityResponseBody
}

type CreateHostAvailabilityResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHostAvailabilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHostAvailabilityResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHostAvailabilityResponse) GoString() string {
	return s.String()
}

func (s *CreateHostAvailabilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHostAvailabilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHostAvailabilityResponse) GetBody() *CreateHostAvailabilityResponseBody {
	return s.Body
}

func (s *CreateHostAvailabilityResponse) SetHeaders(v map[string]*string) *CreateHostAvailabilityResponse {
	s.Headers = v
	return s
}

func (s *CreateHostAvailabilityResponse) SetStatusCode(v int32) *CreateHostAvailabilityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHostAvailabilityResponse) SetBody(v *CreateHostAvailabilityResponseBody) *CreateHostAvailabilityResponse {
	s.Body = v
	return s
}

func (s *CreateHostAvailabilityResponse) Validate() error {
	return dara.Validate(s)
}
