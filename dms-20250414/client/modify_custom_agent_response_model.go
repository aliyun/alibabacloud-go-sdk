// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCustomAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCustomAgentResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCustomAgentResponseBody) *ModifyCustomAgentResponse
	GetBody() *ModifyCustomAgentResponseBody
}

type ModifyCustomAgentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCustomAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCustomAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomAgentResponse) GoString() string {
	return s.String()
}

func (s *ModifyCustomAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCustomAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCustomAgentResponse) GetBody() *ModifyCustomAgentResponseBody {
	return s.Body
}

func (s *ModifyCustomAgentResponse) SetHeaders(v map[string]*string) *ModifyCustomAgentResponse {
	s.Headers = v
	return s
}

func (s *ModifyCustomAgentResponse) SetStatusCode(v int32) *ModifyCustomAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCustomAgentResponse) SetBody(v *ModifyCustomAgentResponseBody) *ModifyCustomAgentResponse {
	s.Body = v
	return s
}

func (s *ModifyCustomAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
