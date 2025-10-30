// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStreamingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStreamingJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateStreamingJobResponseBody) *CreateStreamingJobResponse
	GetBody() *CreateStreamingJobResponseBody
}

type CreateStreamingJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStreamingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStreamingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamingJobResponse) GoString() string {
	return s.String()
}

func (s *CreateStreamingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStreamingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStreamingJobResponse) GetBody() *CreateStreamingJobResponseBody {
	return s.Body
}

func (s *CreateStreamingJobResponse) SetHeaders(v map[string]*string) *CreateStreamingJobResponse {
	s.Headers = v
	return s
}

func (s *CreateStreamingJobResponse) SetStatusCode(v int32) *CreateStreamingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStreamingJobResponse) SetBody(v *CreateStreamingJobResponseBody) *CreateStreamingJobResponse {
	s.Body = v
	return s
}

func (s *CreateStreamingJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
