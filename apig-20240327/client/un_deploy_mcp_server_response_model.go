// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnDeployMcpServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnDeployMcpServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnDeployMcpServerResponse
	GetStatusCode() *int32
	SetBody(v *UnDeployMcpServerResponseBody) *UnDeployMcpServerResponse
	GetBody() *UnDeployMcpServerResponseBody
}

type UnDeployMcpServerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnDeployMcpServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnDeployMcpServerResponse) String() string {
	return dara.Prettify(s)
}

func (s UnDeployMcpServerResponse) GoString() string {
	return s.String()
}

func (s *UnDeployMcpServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnDeployMcpServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnDeployMcpServerResponse) GetBody() *UnDeployMcpServerResponseBody {
	return s.Body
}

func (s *UnDeployMcpServerResponse) SetHeaders(v map[string]*string) *UnDeployMcpServerResponse {
	s.Headers = v
	return s
}

func (s *UnDeployMcpServerResponse) SetStatusCode(v int32) *UnDeployMcpServerResponse {
	s.StatusCode = &v
	return s
}

func (s *UnDeployMcpServerResponse) SetBody(v *UnDeployMcpServerResponseBody) *UnDeployMcpServerResponse {
	s.Body = v
	return s
}

func (s *UnDeployMcpServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
