// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateComputeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateComputeJobResponse
	GetStatusCode() *int32
	SetBody(v *UpdateComputeJobResponseBody) *UpdateComputeJobResponse
	GetBody() *UpdateComputeJobResponseBody
}

type UpdateComputeJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateComputeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateComputeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateComputeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateComputeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateComputeJobResponse) GetBody() *UpdateComputeJobResponseBody {
	return s.Body
}

func (s *UpdateComputeJobResponse) SetHeaders(v map[string]*string) *UpdateComputeJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateComputeJobResponse) SetStatusCode(v int32) *UpdateComputeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateComputeJobResponse) SetBody(v *UpdateComputeJobResponseBody) *UpdateComputeJobResponse {
	s.Body = v
	return s
}

func (s *UpdateComputeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
