// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineCodeDeploymentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRoutineCodeDeploymentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRoutineCodeDeploymentResponse
	GetStatusCode() *int32
	SetBody(v *CreateRoutineCodeDeploymentResponseBody) *CreateRoutineCodeDeploymentResponse
	GetBody() *CreateRoutineCodeDeploymentResponseBody
}

type CreateRoutineCodeDeploymentResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRoutineCodeDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRoutineCodeDeploymentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineCodeDeploymentResponse) GoString() string {
	return s.String()
}

func (s *CreateRoutineCodeDeploymentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRoutineCodeDeploymentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRoutineCodeDeploymentResponse) GetBody() *CreateRoutineCodeDeploymentResponseBody {
	return s.Body
}

func (s *CreateRoutineCodeDeploymentResponse) SetHeaders(v map[string]*string) *CreateRoutineCodeDeploymentResponse {
	s.Headers = v
	return s
}

func (s *CreateRoutineCodeDeploymentResponse) SetStatusCode(v int32) *CreateRoutineCodeDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRoutineCodeDeploymentResponse) SetBody(v *CreateRoutineCodeDeploymentResponseBody) *CreateRoutineCodeDeploymentResponse {
	s.Body = v
	return s
}

func (s *CreateRoutineCodeDeploymentResponse) Validate() error {
	return dara.Validate(s)
}
