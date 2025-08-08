// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeNebulaTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcubeNebulaTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcubeNebulaTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcubeNebulaTaskResponseBody) *CreateMcubeNebulaTaskResponse
	GetBody() *CreateMcubeNebulaTaskResponseBody
}

type CreateMcubeNebulaTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcubeNebulaTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcubeNebulaTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeNebulaTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateMcubeNebulaTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcubeNebulaTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcubeNebulaTaskResponse) GetBody() *CreateMcubeNebulaTaskResponseBody {
	return s.Body
}

func (s *CreateMcubeNebulaTaskResponse) SetHeaders(v map[string]*string) *CreateMcubeNebulaTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateMcubeNebulaTaskResponse) SetStatusCode(v int32) *CreateMcubeNebulaTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcubeNebulaTaskResponse) SetBody(v *CreateMcubeNebulaTaskResponseBody) *CreateMcubeNebulaTaskResponse {
	s.Body = v
	return s
}

func (s *CreateMcubeNebulaTaskResponse) Validate() error {
	return dara.Validate(s)
}
