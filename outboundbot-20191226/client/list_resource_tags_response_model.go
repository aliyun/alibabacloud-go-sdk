// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceTagsResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceTagsResponseBody) *ListResourceTagsResponse
	GetBody() *ListResourceTagsResponseBody
}

type ListResourceTagsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTagsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceTagsResponse) GetBody() *ListResourceTagsResponseBody {
	return s.Body
}

func (s *ListResourceTagsResponse) SetHeaders(v map[string]*string) *ListResourceTagsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceTagsResponse) SetStatusCode(v int32) *ListResourceTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceTagsResponse) SetBody(v *ListResourceTagsResponseBody) *ListResourceTagsResponse {
	s.Body = v
	return s
}

func (s *ListResourceTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
