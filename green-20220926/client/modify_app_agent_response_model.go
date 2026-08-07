// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppAgentResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppAgentResponseBody) *ModifyAppAgentResponse
	GetBody() *ModifyAppAgentResponseBody
}

type ModifyAppAgentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppAgentResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppAgentResponse) GetBody() *ModifyAppAgentResponseBody {
	return s.Body
}

func (s *ModifyAppAgentResponse) SetHeaders(v map[string]*string) *ModifyAppAgentResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppAgentResponse) SetStatusCode(v int32) *ModifyAppAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppAgentResponse) SetBody(v *ModifyAppAgentResponseBody) *ModifyAppAgentResponse {
	s.Body = v
	return s
}

func (s *ModifyAppAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
