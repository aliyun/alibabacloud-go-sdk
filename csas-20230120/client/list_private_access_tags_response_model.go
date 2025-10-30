// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateAccessTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrivateAccessTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrivateAccessTagsResponse
	GetStatusCode() *int32
	SetBody(v *ListPrivateAccessTagsResponseBody) *ListPrivateAccessTagsResponse
	GetBody() *ListPrivateAccessTagsResponseBody
}

type ListPrivateAccessTagsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrivateAccessTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivateAccessTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessTagsResponse) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrivateAccessTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrivateAccessTagsResponse) GetBody() *ListPrivateAccessTagsResponseBody {
	return s.Body
}

func (s *ListPrivateAccessTagsResponse) SetHeaders(v map[string]*string) *ListPrivateAccessTagsResponse {
	s.Headers = v
	return s
}

func (s *ListPrivateAccessTagsResponse) SetStatusCode(v int32) *ListPrivateAccessTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivateAccessTagsResponse) SetBody(v *ListPrivateAccessTagsResponseBody) *ListPrivateAccessTagsResponse {
	s.Body = v
	return s
}

func (s *ListPrivateAccessTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
