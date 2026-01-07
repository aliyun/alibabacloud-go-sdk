// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentWorkspaceMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataAgentWorkspaceMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataAgentWorkspaceMemberResponse
	GetStatusCode() *int32
	SetBody(v *ListDataAgentWorkspaceMemberResponseBody) *ListDataAgentWorkspaceMemberResponse
	GetBody() *ListDataAgentWorkspaceMemberResponseBody
}

type ListDataAgentWorkspaceMemberResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataAgentWorkspaceMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataAgentWorkspaceMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentWorkspaceMemberResponse) GoString() string {
	return s.String()
}

func (s *ListDataAgentWorkspaceMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataAgentWorkspaceMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataAgentWorkspaceMemberResponse) GetBody() *ListDataAgentWorkspaceMemberResponseBody {
	return s.Body
}

func (s *ListDataAgentWorkspaceMemberResponse) SetHeaders(v map[string]*string) *ListDataAgentWorkspaceMemberResponse {
	s.Headers = v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponse) SetStatusCode(v int32) *ListDataAgentWorkspaceMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponse) SetBody(v *ListDataAgentWorkspaceMemberResponseBody) *ListDataAgentWorkspaceMemberResponse {
	s.Body = v
	return s
}

func (s *ListDataAgentWorkspaceMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
