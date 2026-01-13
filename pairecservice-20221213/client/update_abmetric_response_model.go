// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateABMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateABMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateABMetricResponse
	GetStatusCode() *int32
	SetBody(v *UpdateABMetricResponseBody) *UpdateABMetricResponse
	GetBody() *UpdateABMetricResponseBody
}

type UpdateABMetricResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateABMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateABMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateABMetricResponse) GoString() string {
	return s.String()
}

func (s *UpdateABMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateABMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateABMetricResponse) GetBody() *UpdateABMetricResponseBody {
	return s.Body
}

func (s *UpdateABMetricResponse) SetHeaders(v map[string]*string) *UpdateABMetricResponse {
	s.Headers = v
	return s
}

func (s *UpdateABMetricResponse) SetStatusCode(v int32) *UpdateABMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateABMetricResponse) SetBody(v *UpdateABMetricResponseBody) *UpdateABMetricResponse {
	s.Body = v
	return s
}

func (s *UpdateABMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
