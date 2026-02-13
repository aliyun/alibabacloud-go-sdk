// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataAgentSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataAgentSessionResponse
	GetStatusCode() *int32
	SetBody(v *ListDataAgentSessionResponseBody) *ListDataAgentSessionResponse
	GetBody() *ListDataAgentSessionResponseBody
}

type ListDataAgentSessionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataAgentSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataAgentSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentSessionResponse) GoString() string {
	return s.String()
}

func (s *ListDataAgentSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataAgentSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataAgentSessionResponse) GetBody() *ListDataAgentSessionResponseBody {
	return s.Body
}

func (s *ListDataAgentSessionResponse) SetHeaders(v map[string]*string) *ListDataAgentSessionResponse {
	s.Headers = v
	return s
}

func (s *ListDataAgentSessionResponse) SetStatusCode(v int32) *ListDataAgentSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataAgentSessionResponse) SetBody(v *ListDataAgentSessionResponseBody) *ListDataAgentSessionResponse {
	s.Body = v
	return s
}

func (s *ListDataAgentSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
