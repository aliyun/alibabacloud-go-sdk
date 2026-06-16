// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSpacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentSpacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentSpacesResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentSpacesResponseBody) *ListAgentSpacesResponse
	GetBody() *ListAgentSpacesResponseBody
}

type ListAgentSpacesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentSpacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentSpacesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSpacesResponse) GoString() string {
	return s.String()
}

func (s *ListAgentSpacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentSpacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentSpacesResponse) GetBody() *ListAgentSpacesResponseBody {
	return s.Body
}

func (s *ListAgentSpacesResponse) SetHeaders(v map[string]*string) *ListAgentSpacesResponse {
	s.Headers = v
	return s
}

func (s *ListAgentSpacesResponse) SetStatusCode(v int32) *ListAgentSpacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentSpacesResponse) SetBody(v *ListAgentSpacesResponseBody) *ListAgentSpacesResponse {
	s.Body = v
	return s
}

func (s *ListAgentSpacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
