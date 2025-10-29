// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPipelinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunPipelinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunPipelinesResponse
	GetStatusCode() *int32
	SetBody(v *RunPipelinesResponseBody) *RunPipelinesResponse
	GetBody() *RunPipelinesResponseBody
}

type RunPipelinesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunPipelinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunPipelinesResponse) String() string {
	return dara.Prettify(s)
}

func (s RunPipelinesResponse) GoString() string {
	return s.String()
}

func (s *RunPipelinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunPipelinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunPipelinesResponse) GetBody() *RunPipelinesResponseBody {
	return s.Body
}

func (s *RunPipelinesResponse) SetHeaders(v map[string]*string) *RunPipelinesResponse {
	s.Headers = v
	return s
}

func (s *RunPipelinesResponse) SetStatusCode(v int32) *RunPipelinesResponse {
	s.StatusCode = &v
	return s
}

func (s *RunPipelinesResponse) SetBody(v *RunPipelinesResponseBody) *RunPipelinesResponse {
	s.Body = v
	return s
}

func (s *RunPipelinesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
