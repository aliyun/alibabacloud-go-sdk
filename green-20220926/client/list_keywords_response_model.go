// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKeywordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKeywordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKeywordsResponse
	GetStatusCode() *int32
	SetBody(v *ListKeywordsResponseBody) *ListKeywordsResponse
	GetBody() *ListKeywordsResponseBody
}

type ListKeywordsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKeywordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKeywordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKeywordsResponse) GoString() string {
	return s.String()
}

func (s *ListKeywordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKeywordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKeywordsResponse) GetBody() *ListKeywordsResponseBody {
	return s.Body
}

func (s *ListKeywordsResponse) SetHeaders(v map[string]*string) *ListKeywordsResponse {
	s.Headers = v
	return s
}

func (s *ListKeywordsResponse) SetStatusCode(v int32) *ListKeywordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKeywordsResponse) SetBody(v *ListKeywordsResponseBody) *ListKeywordsResponse {
	s.Body = v
	return s
}

func (s *ListKeywordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
