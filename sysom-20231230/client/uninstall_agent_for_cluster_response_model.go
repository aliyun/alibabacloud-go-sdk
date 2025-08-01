// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallAgentForClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallAgentForClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallAgentForClusterResponse
	GetStatusCode() *int32
	SetBody(v *UninstallAgentForClusterResponseBody) *UninstallAgentForClusterResponse
	GetBody() *UninstallAgentForClusterResponseBody
}

type UninstallAgentForClusterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallAgentForClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallAgentForClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallAgentForClusterResponse) GoString() string {
	return s.String()
}

func (s *UninstallAgentForClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallAgentForClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallAgentForClusterResponse) GetBody() *UninstallAgentForClusterResponseBody {
	return s.Body
}

func (s *UninstallAgentForClusterResponse) SetHeaders(v map[string]*string) *UninstallAgentForClusterResponse {
	s.Headers = v
	return s
}

func (s *UninstallAgentForClusterResponse) SetStatusCode(v int32) *UninstallAgentForClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallAgentForClusterResponse) SetBody(v *UninstallAgentForClusterResponseBody) *UninstallAgentForClusterResponse {
	s.Body = v
	return s
}

func (s *UninstallAgentForClusterResponse) Validate() error {
	return dara.Validate(s)
}
