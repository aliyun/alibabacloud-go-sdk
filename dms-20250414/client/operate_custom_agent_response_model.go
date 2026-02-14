// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateCustomAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateCustomAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateCustomAgentResponse
	GetStatusCode() *int32
	SetBody(v *OperateCustomAgentResponseBody) *OperateCustomAgentResponse
	GetBody() *OperateCustomAgentResponseBody
}

type OperateCustomAgentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateCustomAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateCustomAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateCustomAgentResponse) GoString() string {
	return s.String()
}

func (s *OperateCustomAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateCustomAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateCustomAgentResponse) GetBody() *OperateCustomAgentResponseBody {
	return s.Body
}

func (s *OperateCustomAgentResponse) SetHeaders(v map[string]*string) *OperateCustomAgentResponse {
	s.Headers = v
	return s
}

func (s *OperateCustomAgentResponse) SetStatusCode(v int32) *OperateCustomAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateCustomAgentResponse) SetBody(v *OperateCustomAgentResponseBody) *OperateCustomAgentResponse {
	s.Body = v
	return s
}

func (s *OperateCustomAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
