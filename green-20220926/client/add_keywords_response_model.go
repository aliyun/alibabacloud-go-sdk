// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKeywordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddKeywordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddKeywordsResponse
	GetStatusCode() *int32
	SetBody(v *AddKeywordsResponseBody) *AddKeywordsResponse
	GetBody() *AddKeywordsResponseBody
}

type AddKeywordsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddKeywordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddKeywordsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddKeywordsResponse) GoString() string {
	return s.String()
}

func (s *AddKeywordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddKeywordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddKeywordsResponse) GetBody() *AddKeywordsResponseBody {
	return s.Body
}

func (s *AddKeywordsResponse) SetHeaders(v map[string]*string) *AddKeywordsResponse {
	s.Headers = v
	return s
}

func (s *AddKeywordsResponse) SetStatusCode(v int32) *AddKeywordsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddKeywordsResponse) SetBody(v *AddKeywordsResponseBody) *AddKeywordsResponse {
	s.Body = v
	return s
}

func (s *AddKeywordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
