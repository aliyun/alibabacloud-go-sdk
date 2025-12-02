// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKeywordLibsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKeywordLibsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKeywordLibsResponse
	GetStatusCode() *int32
	SetBody(v *ListKeywordLibsResponseBody) *ListKeywordLibsResponse
	GetBody() *ListKeywordLibsResponseBody
}

type ListKeywordLibsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKeywordLibsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKeywordLibsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKeywordLibsResponse) GoString() string {
	return s.String()
}

func (s *ListKeywordLibsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKeywordLibsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKeywordLibsResponse) GetBody() *ListKeywordLibsResponseBody {
	return s.Body
}

func (s *ListKeywordLibsResponse) SetHeaders(v map[string]*string) *ListKeywordLibsResponse {
	s.Headers = v
	return s
}

func (s *ListKeywordLibsResponse) SetStatusCode(v int32) *ListKeywordLibsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKeywordLibsResponse) SetBody(v *ListKeywordLibsResponseBody) *ListKeywordLibsResponse {
	s.Body = v
	return s
}

func (s *ListKeywordLibsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
