// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllPublicMediaTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllPublicMediaTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllPublicMediaTagsResponse
	GetStatusCode() *int32
	SetBody(v *ListAllPublicMediaTagsResponseBody) *ListAllPublicMediaTagsResponse
	GetBody() *ListAllPublicMediaTagsResponseBody
}

type ListAllPublicMediaTagsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllPublicMediaTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllPublicMediaTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllPublicMediaTagsResponse) GoString() string {
	return s.String()
}

func (s *ListAllPublicMediaTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllPublicMediaTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllPublicMediaTagsResponse) GetBody() *ListAllPublicMediaTagsResponseBody {
	return s.Body
}

func (s *ListAllPublicMediaTagsResponse) SetHeaders(v map[string]*string) *ListAllPublicMediaTagsResponse {
	s.Headers = v
	return s
}

func (s *ListAllPublicMediaTagsResponse) SetStatusCode(v int32) *ListAllPublicMediaTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllPublicMediaTagsResponse) SetBody(v *ListAllPublicMediaTagsResponseBody) *ListAllPublicMediaTagsResponse {
	s.Body = v
	return s
}

func (s *ListAllPublicMediaTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
