// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeChatAgentStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeChatAgentStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeChatAgentStatusResponse
	GetStatusCode() *int32
	SetBody(v *ChangeChatAgentStatusResponseBody) *ChangeChatAgentStatusResponse
	GetBody() *ChangeChatAgentStatusResponseBody
}

type ChangeChatAgentStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeChatAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeChatAgentStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeChatAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *ChangeChatAgentStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeChatAgentStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeChatAgentStatusResponse) GetBody() *ChangeChatAgentStatusResponseBody {
	return s.Body
}

func (s *ChangeChatAgentStatusResponse) SetHeaders(v map[string]*string) *ChangeChatAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *ChangeChatAgentStatusResponse) SetStatusCode(v int32) *ChangeChatAgentStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeChatAgentStatusResponse) SetBody(v *ChangeChatAgentStatusResponseBody) *ChangeChatAgentStatusResponse {
	s.Body = v
	return s
}

func (s *ChangeChatAgentStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
