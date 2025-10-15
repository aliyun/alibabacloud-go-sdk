// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationTokensResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationTokensResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationTokensResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationTokensResponseBody) *ListApplicationTokensResponse
	GetBody() *ListApplicationTokensResponseBody
}

type ListApplicationTokensResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationTokensResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationTokensResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationTokensResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationTokensResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationTokensResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationTokensResponse) GetBody() *ListApplicationTokensResponseBody {
	return s.Body
}

func (s *ListApplicationTokensResponse) SetHeaders(v map[string]*string) *ListApplicationTokensResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationTokensResponse) SetStatusCode(v int32) *ListApplicationTokensResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationTokensResponse) SetBody(v *ListApplicationTokensResponseBody) *ListApplicationTokensResponse {
	s.Body = v
	return s
}

func (s *ListApplicationTokensResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
