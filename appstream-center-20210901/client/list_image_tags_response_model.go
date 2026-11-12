// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListImageTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListImageTagsResponse
	GetStatusCode() *int32
	SetBody(v *ListImageTagsResponseBody) *ListImageTagsResponse
	GetBody() *ListImageTagsResponseBody
}

type ListImageTagsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImageTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImageTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListImageTagsResponse) GoString() string {
	return s.String()
}

func (s *ListImageTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListImageTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListImageTagsResponse) GetBody() *ListImageTagsResponseBody {
	return s.Body
}

func (s *ListImageTagsResponse) SetHeaders(v map[string]*string) *ListImageTagsResponse {
	s.Headers = v
	return s
}

func (s *ListImageTagsResponse) SetStatusCode(v int32) *ListImageTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImageTagsResponse) SetBody(v *ListImageTagsResponseBody) *ListImageTagsResponse {
	s.Body = v
	return s
}

func (s *ListImageTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
