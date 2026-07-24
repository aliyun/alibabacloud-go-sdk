// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateComputeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateComputeJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateComputeJobResponseBody) *CreateComputeJobResponse
	GetBody() *CreateComputeJobResponseBody
}

type CreateComputeJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateComputeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateComputeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeJobResponse) GoString() string {
	return s.String()
}

func (s *CreateComputeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateComputeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateComputeJobResponse) GetBody() *CreateComputeJobResponseBody {
	return s.Body
}

func (s *CreateComputeJobResponse) SetHeaders(v map[string]*string) *CreateComputeJobResponse {
	s.Headers = v
	return s
}

func (s *CreateComputeJobResponse) SetStatusCode(v int32) *CreateComputeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateComputeJobResponse) SetBody(v *CreateComputeJobResponseBody) *CreateComputeJobResponse {
	s.Body = v
	return s
}

func (s *CreateComputeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
