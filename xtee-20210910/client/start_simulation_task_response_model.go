// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSimulationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartSimulationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartSimulationTaskResponse
	GetStatusCode() *int32
	SetBody(v *StartSimulationTaskResponseBody) *StartSimulationTaskResponse
	GetBody() *StartSimulationTaskResponseBody
}

type StartSimulationTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartSimulationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartSimulationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StartSimulationTaskResponse) GoString() string {
	return s.String()
}

func (s *StartSimulationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartSimulationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartSimulationTaskResponse) GetBody() *StartSimulationTaskResponseBody {
	return s.Body
}

func (s *StartSimulationTaskResponse) SetHeaders(v map[string]*string) *StartSimulationTaskResponse {
	s.Headers = v
	return s
}

func (s *StartSimulationTaskResponse) SetStatusCode(v int32) *StartSimulationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartSimulationTaskResponse) SetBody(v *StartSimulationTaskResponseBody) *StartSimulationTaskResponse {
	s.Body = v
	return s
}

func (s *StartSimulationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
