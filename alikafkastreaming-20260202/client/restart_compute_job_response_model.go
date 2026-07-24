// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartComputeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartComputeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartComputeJobResponse
	GetStatusCode() *int32
	SetBody(v *RestartComputeJobResponseBody) *RestartComputeJobResponse
	GetBody() *RestartComputeJobResponseBody
}

type RestartComputeJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartComputeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartComputeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartComputeJobResponse) GoString() string {
	return s.String()
}

func (s *RestartComputeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartComputeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartComputeJobResponse) GetBody() *RestartComputeJobResponseBody {
	return s.Body
}

func (s *RestartComputeJobResponse) SetHeaders(v map[string]*string) *RestartComputeJobResponse {
	s.Headers = v
	return s
}

func (s *RestartComputeJobResponse) SetStatusCode(v int32) *RestartComputeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartComputeJobResponse) SetBody(v *RestartComputeJobResponseBody) *RestartComputeJobResponse {
	s.Body = v
	return s
}

func (s *RestartComputeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
