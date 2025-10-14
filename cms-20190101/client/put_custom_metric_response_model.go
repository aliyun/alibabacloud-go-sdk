// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutCustomMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutCustomMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutCustomMetricResponse
	GetStatusCode() *int32
	SetBody(v *PutCustomMetricResponseBody) *PutCustomMetricResponse
	GetBody() *PutCustomMetricResponseBody
}

type PutCustomMetricResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutCustomMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutCustomMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s PutCustomMetricResponse) GoString() string {
	return s.String()
}

func (s *PutCustomMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutCustomMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutCustomMetricResponse) GetBody() *PutCustomMetricResponseBody {
	return s.Body
}

func (s *PutCustomMetricResponse) SetHeaders(v map[string]*string) *PutCustomMetricResponse {
	s.Headers = v
	return s
}

func (s *PutCustomMetricResponse) SetStatusCode(v int32) *PutCustomMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *PutCustomMetricResponse) SetBody(v *PutCustomMetricResponseBody) *PutCustomMetricResponse {
	s.Body = v
	return s
}

func (s *PutCustomMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
