// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAvailabilityMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAvailabilityMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAvailabilityMetricResponse
	GetStatusCode() *int32
	SetBody(v *GetAvailabilityMetricResponseBody) *GetAvailabilityMetricResponse
	GetBody() *GetAvailabilityMetricResponseBody
}

type GetAvailabilityMetricResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAvailabilityMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAvailabilityMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAvailabilityMetricResponse) GoString() string {
	return s.String()
}

func (s *GetAvailabilityMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAvailabilityMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAvailabilityMetricResponse) GetBody() *GetAvailabilityMetricResponseBody {
	return s.Body
}

func (s *GetAvailabilityMetricResponse) SetHeaders(v map[string]*string) *GetAvailabilityMetricResponse {
	s.Headers = v
	return s
}

func (s *GetAvailabilityMetricResponse) SetStatusCode(v int32) *GetAvailabilityMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAvailabilityMetricResponse) SetBody(v *GetAvailabilityMetricResponseBody) *GetAvailabilityMetricResponse {
	s.Body = v
	return s
}

func (s *GetAvailabilityMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
