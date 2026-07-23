// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTrafficControlFlinkTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopTrafficControlFlinkTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopTrafficControlFlinkTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopTrafficControlFlinkTaskResponseBody) *StopTrafficControlFlinkTaskResponse
	GetBody() *StopTrafficControlFlinkTaskResponseBody
}

type StopTrafficControlFlinkTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopTrafficControlFlinkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTrafficControlFlinkTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopTrafficControlFlinkTaskResponse) GoString() string {
	return s.String()
}

func (s *StopTrafficControlFlinkTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopTrafficControlFlinkTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopTrafficControlFlinkTaskResponse) GetBody() *StopTrafficControlFlinkTaskResponseBody {
	return s.Body
}

func (s *StopTrafficControlFlinkTaskResponse) SetHeaders(v map[string]*string) *StopTrafficControlFlinkTaskResponse {
	s.Headers = v
	return s
}

func (s *StopTrafficControlFlinkTaskResponse) SetStatusCode(v int32) *StopTrafficControlFlinkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTrafficControlFlinkTaskResponse) SetBody(v *StopTrafficControlFlinkTaskResponseBody) *StopTrafficControlFlinkTaskResponse {
	s.Body = v
	return s
}

func (s *StopTrafficControlFlinkTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
