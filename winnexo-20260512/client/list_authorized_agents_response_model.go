// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedAgentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorizedAgentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorizedAgentsResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthorizedAgentsResponseBody) *ListAuthorizedAgentsResponse
	GetBody() *ListAuthorizedAgentsResponseBody
}

type ListAuthorizedAgentsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizedAgentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizedAgentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedAgentsResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizedAgentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorizedAgentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorizedAgentsResponse) GetBody() *ListAuthorizedAgentsResponseBody {
	return s.Body
}

func (s *ListAuthorizedAgentsResponse) SetHeaders(v map[string]*string) *ListAuthorizedAgentsResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizedAgentsResponse) SetStatusCode(v int32) *ListAuthorizedAgentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizedAgentsResponse) SetBody(v *ListAuthorizedAgentsResponseBody) *ListAuthorizedAgentsResponse {
	s.Body = v
	return s
}

func (s *ListAuthorizedAgentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
