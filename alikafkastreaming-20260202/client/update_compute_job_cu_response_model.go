// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeJobCuResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateComputeJobCuResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateComputeJobCuResponse
	GetStatusCode() *int32
	SetBody(v *UpdateComputeJobCuResponseBody) *UpdateComputeJobCuResponse
	GetBody() *UpdateComputeJobCuResponseBody
}

type UpdateComputeJobCuResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateComputeJobCuResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateComputeJobCuResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeJobCuResponse) GoString() string {
	return s.String()
}

func (s *UpdateComputeJobCuResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateComputeJobCuResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateComputeJobCuResponse) GetBody() *UpdateComputeJobCuResponseBody {
	return s.Body
}

func (s *UpdateComputeJobCuResponse) SetHeaders(v map[string]*string) *UpdateComputeJobCuResponse {
	s.Headers = v
	return s
}

func (s *UpdateComputeJobCuResponse) SetStatusCode(v int32) *UpdateComputeJobCuResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateComputeJobCuResponse) SetBody(v *UpdateComputeJobCuResponseBody) *UpdateComputeJobCuResponse {
	s.Body = v
	return s
}

func (s *UpdateComputeJobCuResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
