// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMetricResponse
	GetStatusCode() *int32
	SetBody(v *QueryMetricResponseBody) *QueryMetricResponse
	GetBody() *QueryMetricResponseBody
}

type QueryMetricResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMetricResponse) GoString() string {
	return s.String()
}

func (s *QueryMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMetricResponse) GetBody() *QueryMetricResponseBody {
	return s.Body
}

func (s *QueryMetricResponse) SetHeaders(v map[string]*string) *QueryMetricResponse {
	s.Headers = v
	return s
}

func (s *QueryMetricResponse) SetStatusCode(v int32) *QueryMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMetricResponse) SetBody(v *QueryMetricResponseBody) *QueryMetricResponse {
	s.Body = v
	return s
}

func (s *QueryMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
