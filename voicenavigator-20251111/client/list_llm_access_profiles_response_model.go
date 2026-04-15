// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLlmAccessProfilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLlmAccessProfilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLlmAccessProfilesResponse
	GetStatusCode() *int32
	SetBody(v *ListLlmAccessProfilesResponseBody) *ListLlmAccessProfilesResponse
	GetBody() *ListLlmAccessProfilesResponseBody
}

type ListLlmAccessProfilesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLlmAccessProfilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLlmAccessProfilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLlmAccessProfilesResponse) GoString() string {
	return s.String()
}

func (s *ListLlmAccessProfilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLlmAccessProfilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLlmAccessProfilesResponse) GetBody() *ListLlmAccessProfilesResponseBody {
	return s.Body
}

func (s *ListLlmAccessProfilesResponse) SetHeaders(v map[string]*string) *ListLlmAccessProfilesResponse {
	s.Headers = v
	return s
}

func (s *ListLlmAccessProfilesResponse) SetStatusCode(v int32) *ListLlmAccessProfilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLlmAccessProfilesResponse) SetBody(v *ListLlmAccessProfilesResponseBody) *ListLlmAccessProfilesResponse {
	s.Body = v
	return s
}

func (s *ListLlmAccessProfilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
