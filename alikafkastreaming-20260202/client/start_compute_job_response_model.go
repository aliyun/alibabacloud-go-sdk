// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartComputeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartComputeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartComputeJobResponse
	GetStatusCode() *int32
	SetBody(v *StartComputeJobResponseBody) *StartComputeJobResponse
	GetBody() *StartComputeJobResponseBody
}

type StartComputeJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartComputeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartComputeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StartComputeJobResponse) GoString() string {
	return s.String()
}

func (s *StartComputeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartComputeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartComputeJobResponse) GetBody() *StartComputeJobResponseBody {
	return s.Body
}

func (s *StartComputeJobResponse) SetHeaders(v map[string]*string) *StartComputeJobResponse {
	s.Headers = v
	return s
}

func (s *StartComputeJobResponse) SetStatusCode(v int32) *StartComputeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StartComputeJobResponse) SetBody(v *StartComputeJobResponseBody) *StartComputeJobResponse {
	s.Body = v
	return s
}

func (s *StartComputeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
