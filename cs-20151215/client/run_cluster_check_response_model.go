// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunClusterCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunClusterCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunClusterCheckResponse
	GetStatusCode() *int32
	SetBody(v *RunClusterCheckResponseBody) *RunClusterCheckResponse
	GetBody() *RunClusterCheckResponseBody
}

type RunClusterCheckResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunClusterCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunClusterCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s RunClusterCheckResponse) GoString() string {
	return s.String()
}

func (s *RunClusterCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunClusterCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunClusterCheckResponse) GetBody() *RunClusterCheckResponseBody {
	return s.Body
}

func (s *RunClusterCheckResponse) SetHeaders(v map[string]*string) *RunClusterCheckResponse {
	s.Headers = v
	return s
}

func (s *RunClusterCheckResponse) SetStatusCode(v int32) *RunClusterCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *RunClusterCheckResponse) SetBody(v *RunClusterCheckResponseBody) *RunClusterCheckResponse {
	s.Body = v
	return s
}

func (s *RunClusterCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
