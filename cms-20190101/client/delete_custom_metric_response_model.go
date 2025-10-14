// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomMetricResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomMetricResponseBody) *DeleteCustomMetricResponse
	GetBody() *DeleteCustomMetricResponseBody
}

type DeleteCustomMetricResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomMetricResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomMetricResponse) GetBody() *DeleteCustomMetricResponseBody {
	return s.Body
}

func (s *DeleteCustomMetricResponse) SetHeaders(v map[string]*string) *DeleteCustomMetricResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomMetricResponse) SetStatusCode(v int32) *DeleteCustomMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomMetricResponse) SetBody(v *DeleteCustomMetricResponseBody) *DeleteCustomMetricResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
