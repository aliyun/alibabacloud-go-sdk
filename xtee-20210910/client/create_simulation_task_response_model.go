// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimulationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSimulationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSimulationTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateSimulationTaskResponseBody) *CreateSimulationTaskResponse
	GetBody() *CreateSimulationTaskResponseBody
}

type CreateSimulationTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSimulationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSimulationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSimulationTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSimulationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSimulationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSimulationTaskResponse) GetBody() *CreateSimulationTaskResponseBody {
	return s.Body
}

func (s *CreateSimulationTaskResponse) SetHeaders(v map[string]*string) *CreateSimulationTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateSimulationTaskResponse) SetStatusCode(v int32) *CreateSimulationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSimulationTaskResponse) SetBody(v *CreateSimulationTaskResponseBody) *CreateSimulationTaskResponse {
	s.Body = v
	return s
}

func (s *CreateSimulationTaskResponse) Validate() error {
	return dara.Validate(s)
}
