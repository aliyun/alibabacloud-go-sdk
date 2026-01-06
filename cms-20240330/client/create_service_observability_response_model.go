// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceObservabilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceObservabilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceObservabilityResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceObservabilityResponseBody) *CreateServiceObservabilityResponse
	GetBody() *CreateServiceObservabilityResponseBody
}

type CreateServiceObservabilityResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceObservabilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceObservabilityResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceObservabilityResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceObservabilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceObservabilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceObservabilityResponse) GetBody() *CreateServiceObservabilityResponseBody {
	return s.Body
}

func (s *CreateServiceObservabilityResponse) SetHeaders(v map[string]*string) *CreateServiceObservabilityResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceObservabilityResponse) SetStatusCode(v int32) *CreateServiceObservabilityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceObservabilityResponse) SetBody(v *CreateServiceObservabilityResponseBody) *CreateServiceObservabilityResponse {
	s.Body = v
	return s
}

func (s *CreateServiceObservabilityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
