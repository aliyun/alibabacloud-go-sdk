// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTaskFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTaskFlowResponse
	GetStatusCode() *int32
	SetBody(v *CreateTaskFlowResponseBody) *CreateTaskFlowResponse
	GetBody() *CreateTaskFlowResponseBody
}

type CreateTaskFlowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTaskFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTaskFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskFlowResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTaskFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTaskFlowResponse) GetBody() *CreateTaskFlowResponseBody {
	return s.Body
}

func (s *CreateTaskFlowResponse) SetHeaders(v map[string]*string) *CreateTaskFlowResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskFlowResponse) SetStatusCode(v int32) *CreateTaskFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTaskFlowResponse) SetBody(v *CreateTaskFlowResponseBody) *CreateTaskFlowResponse {
	s.Body = v
	return s
}

func (s *CreateTaskFlowResponse) Validate() error {
	return dara.Validate(s)
}
