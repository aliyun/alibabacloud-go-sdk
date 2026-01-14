// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinAgentGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *JoinAgentGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *JoinAgentGroupResponse
	GetStatusCode() *int32
	SetBody(v *JoinAgentGroupResponseBody) *JoinAgentGroupResponse
	GetBody() *JoinAgentGroupResponseBody
}

type JoinAgentGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *JoinAgentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s JoinAgentGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s JoinAgentGroupResponse) GoString() string {
	return s.String()
}

func (s *JoinAgentGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *JoinAgentGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *JoinAgentGroupResponse) GetBody() *JoinAgentGroupResponseBody {
	return s.Body
}

func (s *JoinAgentGroupResponse) SetHeaders(v map[string]*string) *JoinAgentGroupResponse {
	s.Headers = v
	return s
}

func (s *JoinAgentGroupResponse) SetStatusCode(v int32) *JoinAgentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinAgentGroupResponse) SetBody(v *JoinAgentGroupResponseBody) *JoinAgentGroupResponse {
	s.Body = v
	return s
}

func (s *JoinAgentGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
