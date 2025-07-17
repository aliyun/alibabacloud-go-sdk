// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnvCustomJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEnvCustomJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEnvCustomJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEnvCustomJobResponseBody) *DeleteEnvCustomJobResponse
	GetBody() *DeleteEnvCustomJobResponseBody
}

type DeleteEnvCustomJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEnvCustomJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEnvCustomJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnvCustomJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnvCustomJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEnvCustomJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEnvCustomJobResponse) GetBody() *DeleteEnvCustomJobResponseBody {
	return s.Body
}

func (s *DeleteEnvCustomJobResponse) SetHeaders(v map[string]*string) *DeleteEnvCustomJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnvCustomJobResponse) SetStatusCode(v int32) *DeleteEnvCustomJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEnvCustomJobResponse) SetBody(v *DeleteEnvCustomJobResponseBody) *DeleteEnvCustomJobResponse {
	s.Body = v
	return s
}

func (s *DeleteEnvCustomJobResponse) Validate() error {
	return dara.Validate(s)
}
