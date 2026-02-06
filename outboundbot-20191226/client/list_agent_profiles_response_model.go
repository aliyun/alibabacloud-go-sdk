// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentProfilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentProfilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentProfilesResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentProfilesResponseBody) *ListAgentProfilesResponse
	GetBody() *ListAgentProfilesResponseBody
}

type ListAgentProfilesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentProfilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentProfilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentProfilesResponse) GoString() string {
	return s.String()
}

func (s *ListAgentProfilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentProfilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentProfilesResponse) GetBody() *ListAgentProfilesResponseBody {
	return s.Body
}

func (s *ListAgentProfilesResponse) SetHeaders(v map[string]*string) *ListAgentProfilesResponse {
	s.Headers = v
	return s
}

func (s *ListAgentProfilesResponse) SetStatusCode(v int32) *ListAgentProfilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentProfilesResponse) SetBody(v *ListAgentProfilesResponseBody) *ListAgentProfilesResponse {
	s.Body = v
	return s
}

func (s *ListAgentProfilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
