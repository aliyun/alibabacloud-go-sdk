// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAgentStorage2VpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindAgentStorage2VpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindAgentStorage2VpcResponse
	GetStatusCode() *int32
	SetBody(v *UnbindAgentStorage2VpcResponseBody) *UnbindAgentStorage2VpcResponse
	GetBody() *UnbindAgentStorage2VpcResponseBody
}

type UnbindAgentStorage2VpcResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindAgentStorage2VpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindAgentStorage2VpcResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindAgentStorage2VpcResponse) GoString() string {
	return s.String()
}

func (s *UnbindAgentStorage2VpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindAgentStorage2VpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindAgentStorage2VpcResponse) GetBody() *UnbindAgentStorage2VpcResponseBody {
	return s.Body
}

func (s *UnbindAgentStorage2VpcResponse) SetHeaders(v map[string]*string) *UnbindAgentStorage2VpcResponse {
	s.Headers = v
	return s
}

func (s *UnbindAgentStorage2VpcResponse) SetStatusCode(v int32) *UnbindAgentStorage2VpcResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindAgentStorage2VpcResponse) SetBody(v *UnbindAgentStorage2VpcResponseBody) *UnbindAgentStorage2VpcResponse {
	s.Body = v
	return s
}

func (s *UnbindAgentStorage2VpcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
