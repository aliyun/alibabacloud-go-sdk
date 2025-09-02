// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployDISyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeployDISyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeployDISyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeployDISyncTaskResponseBody) *DeployDISyncTaskResponse
	GetBody() *DeployDISyncTaskResponseBody
}

type DeployDISyncTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployDISyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployDISyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeployDISyncTaskResponse) GoString() string {
	return s.String()
}

func (s *DeployDISyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeployDISyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeployDISyncTaskResponse) GetBody() *DeployDISyncTaskResponseBody {
	return s.Body
}

func (s *DeployDISyncTaskResponse) SetHeaders(v map[string]*string) *DeployDISyncTaskResponse {
	s.Headers = v
	return s
}

func (s *DeployDISyncTaskResponse) SetStatusCode(v int32) *DeployDISyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployDISyncTaskResponse) SetBody(v *DeployDISyncTaskResponseBody) *DeployDISyncTaskResponse {
	s.Body = v
	return s
}

func (s *DeployDISyncTaskResponse) Validate() error {
	return dara.Validate(s)
}
