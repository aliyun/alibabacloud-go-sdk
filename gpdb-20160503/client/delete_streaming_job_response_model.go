// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStreamingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStreamingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStreamingJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStreamingJobResponseBody) *DeleteStreamingJobResponse
	GetBody() *DeleteStreamingJobResponseBody
}

type DeleteStreamingJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStreamingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStreamingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStreamingJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteStreamingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStreamingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStreamingJobResponse) GetBody() *DeleteStreamingJobResponseBody {
	return s.Body
}

func (s *DeleteStreamingJobResponse) SetHeaders(v map[string]*string) *DeleteStreamingJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteStreamingJobResponse) SetStatusCode(v int32) *DeleteStreamingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStreamingJobResponse) SetBody(v *DeleteStreamingJobResponseBody) *DeleteStreamingJobResponse {
	s.Body = v
	return s
}

func (s *DeleteStreamingJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
