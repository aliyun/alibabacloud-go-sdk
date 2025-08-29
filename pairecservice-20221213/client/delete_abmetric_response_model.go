// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteABMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteABMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteABMetricResponse
	GetStatusCode() *int32
	SetBody(v *DeleteABMetricResponseBody) *DeleteABMetricResponse
	GetBody() *DeleteABMetricResponseBody
}

type DeleteABMetricResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteABMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteABMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteABMetricResponse) GoString() string {
	return s.String()
}

func (s *DeleteABMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteABMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteABMetricResponse) GetBody() *DeleteABMetricResponseBody {
	return s.Body
}

func (s *DeleteABMetricResponse) SetHeaders(v map[string]*string) *DeleteABMetricResponse {
	s.Headers = v
	return s
}

func (s *DeleteABMetricResponse) SetStatusCode(v int32) *DeleteABMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteABMetricResponse) SetBody(v *DeleteABMetricResponseBody) *DeleteABMetricResponse {
	s.Body = v
	return s
}

func (s *DeleteABMetricResponse) Validate() error {
	return dara.Validate(s)
}
