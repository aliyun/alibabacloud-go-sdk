// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExternalAgentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExternalAgentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExternalAgentsResponse
	GetStatusCode() *int32
	SetBody(v *ListExternalAgentsResponseBody) *ListExternalAgentsResponse
	GetBody() *ListExternalAgentsResponseBody
}

type ListExternalAgentsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExternalAgentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExternalAgentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExternalAgentsResponse) GoString() string {
	return s.String()
}

func (s *ListExternalAgentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExternalAgentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExternalAgentsResponse) GetBody() *ListExternalAgentsResponseBody {
	return s.Body
}

func (s *ListExternalAgentsResponse) SetHeaders(v map[string]*string) *ListExternalAgentsResponse {
	s.Headers = v
	return s
}

func (s *ListExternalAgentsResponse) SetStatusCode(v int32) *ListExternalAgentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExternalAgentsResponse) SetBody(v *ListExternalAgentsResponseBody) *ListExternalAgentsResponse {
	s.Body = v
	return s
}

func (s *ListExternalAgentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
