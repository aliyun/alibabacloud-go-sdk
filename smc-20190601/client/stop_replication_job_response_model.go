// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopReplicationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopReplicationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopReplicationJobResponse
	GetStatusCode() *int32
	SetBody(v *StopReplicationJobResponseBody) *StopReplicationJobResponse
	GetBody() *StopReplicationJobResponseBody
}

type StopReplicationJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopReplicationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopReplicationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StopReplicationJobResponse) GoString() string {
	return s.String()
}

func (s *StopReplicationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopReplicationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopReplicationJobResponse) GetBody() *StopReplicationJobResponseBody {
	return s.Body
}

func (s *StopReplicationJobResponse) SetHeaders(v map[string]*string) *StopReplicationJobResponse {
	s.Headers = v
	return s
}

func (s *StopReplicationJobResponse) SetStatusCode(v int32) *StopReplicationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopReplicationJobResponse) SetBody(v *StopReplicationJobResponseBody) *StopReplicationJobResponse {
	s.Body = v
	return s
}

func (s *StopReplicationJobResponse) Validate() error {
	return dara.Validate(s)
}
