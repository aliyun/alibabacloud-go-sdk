// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSaseUserTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSaseUserTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSaseUserTagsResponse
	GetStatusCode() *int32
	SetBody(v *ListSaseUserTagsResponseBody) *ListSaseUserTagsResponse
	GetBody() *ListSaseUserTagsResponseBody
}

type ListSaseUserTagsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSaseUserTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSaseUserTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSaseUserTagsResponse) GoString() string {
	return s.String()
}

func (s *ListSaseUserTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSaseUserTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSaseUserTagsResponse) GetBody() *ListSaseUserTagsResponseBody {
	return s.Body
}

func (s *ListSaseUserTagsResponse) SetHeaders(v map[string]*string) *ListSaseUserTagsResponse {
	s.Headers = v
	return s
}

func (s *ListSaseUserTagsResponse) SetStatusCode(v int32) *ListSaseUserTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSaseUserTagsResponse) SetBody(v *ListSaseUserTagsResponseBody) *ListSaseUserTagsResponse {
	s.Body = v
	return s
}

func (s *ListSaseUserTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
