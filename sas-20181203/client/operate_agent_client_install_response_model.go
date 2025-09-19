// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAgentClientInstallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateAgentClientInstallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateAgentClientInstallResponse
	GetStatusCode() *int32
	SetBody(v *OperateAgentClientInstallResponseBody) *OperateAgentClientInstallResponse
	GetBody() *OperateAgentClientInstallResponseBody
}

type OperateAgentClientInstallResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateAgentClientInstallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateAgentClientInstallResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateAgentClientInstallResponse) GoString() string {
	return s.String()
}

func (s *OperateAgentClientInstallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateAgentClientInstallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateAgentClientInstallResponse) GetBody() *OperateAgentClientInstallResponseBody {
	return s.Body
}

func (s *OperateAgentClientInstallResponse) SetHeaders(v map[string]*string) *OperateAgentClientInstallResponse {
	s.Headers = v
	return s
}

func (s *OperateAgentClientInstallResponse) SetStatusCode(v int32) *OperateAgentClientInstallResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateAgentClientInstallResponse) SetBody(v *OperateAgentClientInstallResponseBody) *OperateAgentClientInstallResponse {
	s.Body = v
	return s
}

func (s *OperateAgentClientInstallResponse) Validate() error {
	return dara.Validate(s)
}
