// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAgentStorage2VpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindAgentStorage2VpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindAgentStorage2VpcResponse
	GetStatusCode() *int32
	SetBody(v *BindAgentStorage2VpcResponseBody) *BindAgentStorage2VpcResponse
	GetBody() *BindAgentStorage2VpcResponseBody
}

type BindAgentStorage2VpcResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAgentStorage2VpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAgentStorage2VpcResponse) String() string {
	return dara.Prettify(s)
}

func (s BindAgentStorage2VpcResponse) GoString() string {
	return s.String()
}

func (s *BindAgentStorage2VpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindAgentStorage2VpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindAgentStorage2VpcResponse) GetBody() *BindAgentStorage2VpcResponseBody {
	return s.Body
}

func (s *BindAgentStorage2VpcResponse) SetHeaders(v map[string]*string) *BindAgentStorage2VpcResponse {
	s.Headers = v
	return s
}

func (s *BindAgentStorage2VpcResponse) SetStatusCode(v int32) *BindAgentStorage2VpcResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAgentStorage2VpcResponse) SetBody(v *BindAgentStorage2VpcResponseBody) *BindAgentStorage2VpcResponse {
	s.Body = v
	return s
}

func (s *BindAgentStorage2VpcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
