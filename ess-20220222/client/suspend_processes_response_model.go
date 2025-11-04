// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendProcessesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendProcessesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendProcessesResponse
	GetStatusCode() *int32
	SetBody(v *SuspendProcessesResponseBody) *SuspendProcessesResponse
	GetBody() *SuspendProcessesResponseBody
}

type SuspendProcessesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendProcessesResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendProcessesResponse) GoString() string {
	return s.String()
}

func (s *SuspendProcessesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendProcessesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendProcessesResponse) GetBody() *SuspendProcessesResponseBody {
	return s.Body
}

func (s *SuspendProcessesResponse) SetHeaders(v map[string]*string) *SuspendProcessesResponse {
	s.Headers = v
	return s
}

func (s *SuspendProcessesResponse) SetStatusCode(v int32) *SuspendProcessesResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendProcessesResponse) SetBody(v *SuspendProcessesResponseBody) *SuspendProcessesResponse {
	s.Body = v
	return s
}

func (s *SuspendProcessesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
