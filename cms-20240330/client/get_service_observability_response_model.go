// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceObservabilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceObservabilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceObservabilityResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceObservabilityResponseBody) *GetServiceObservabilityResponse
	GetBody() *GetServiceObservabilityResponseBody
}

type GetServiceObservabilityResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceObservabilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceObservabilityResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceObservabilityResponse) GoString() string {
	return s.String()
}

func (s *GetServiceObservabilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceObservabilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceObservabilityResponse) GetBody() *GetServiceObservabilityResponseBody {
	return s.Body
}

func (s *GetServiceObservabilityResponse) SetHeaders(v map[string]*string) *GetServiceObservabilityResponse {
	s.Headers = v
	return s
}

func (s *GetServiceObservabilityResponse) SetStatusCode(v int32) *GetServiceObservabilityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceObservabilityResponse) SetBody(v *GetServiceObservabilityResponseBody) *GetServiceObservabilityResponse {
	s.Body = v
	return s
}

func (s *GetServiceObservabilityResponse) Validate() error {
	return dara.Validate(s)
}
