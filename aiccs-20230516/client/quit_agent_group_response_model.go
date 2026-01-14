// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuitAgentGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuitAgentGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuitAgentGroupResponse
	GetStatusCode() *int32
	SetBody(v *QuitAgentGroupResponseBody) *QuitAgentGroupResponse
	GetBody() *QuitAgentGroupResponseBody
}

type QuitAgentGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuitAgentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuitAgentGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s QuitAgentGroupResponse) GoString() string {
	return s.String()
}

func (s *QuitAgentGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuitAgentGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuitAgentGroupResponse) GetBody() *QuitAgentGroupResponseBody {
	return s.Body
}

func (s *QuitAgentGroupResponse) SetHeaders(v map[string]*string) *QuitAgentGroupResponse {
	s.Headers = v
	return s
}

func (s *QuitAgentGroupResponse) SetStatusCode(v int32) *QuitAgentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *QuitAgentGroupResponse) SetBody(v *QuitAgentGroupResponseBody) *QuitAgentGroupResponse {
	s.Body = v
	return s
}

func (s *QuitAgentGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
