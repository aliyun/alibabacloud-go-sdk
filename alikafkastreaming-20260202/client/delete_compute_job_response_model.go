// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComputeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteComputeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteComputeJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteComputeJobResponseBody) *DeleteComputeJobResponse
	GetBody() *DeleteComputeJobResponseBody
}

type DeleteComputeJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteComputeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteComputeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteComputeJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteComputeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteComputeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteComputeJobResponse) GetBody() *DeleteComputeJobResponseBody {
	return s.Body
}

func (s *DeleteComputeJobResponse) SetHeaders(v map[string]*string) *DeleteComputeJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteComputeJobResponse) SetStatusCode(v int32) *DeleteComputeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteComputeJobResponse) SetBody(v *DeleteComputeJobResponseBody) *DeleteComputeJobResponse {
	s.Body = v
	return s
}

func (s *DeleteComputeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
