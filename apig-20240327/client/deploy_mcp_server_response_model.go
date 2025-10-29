// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployMcpServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeployMcpServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeployMcpServerResponse
	GetStatusCode() *int32
	SetBody(v *DeployMcpServerResponseBody) *DeployMcpServerResponse
	GetBody() *DeployMcpServerResponseBody
}

type DeployMcpServerResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployMcpServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployMcpServerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeployMcpServerResponse) GoString() string {
	return s.String()
}

func (s *DeployMcpServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeployMcpServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeployMcpServerResponse) GetBody() *DeployMcpServerResponseBody {
	return s.Body
}

func (s *DeployMcpServerResponse) SetHeaders(v map[string]*string) *DeployMcpServerResponse {
	s.Headers = v
	return s
}

func (s *DeployMcpServerResponse) SetStatusCode(v int32) *DeployMcpServerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployMcpServerResponse) SetBody(v *DeployMcpServerResponseBody) *DeployMcpServerResponse {
	s.Body = v
	return s
}

func (s *DeployMcpServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
