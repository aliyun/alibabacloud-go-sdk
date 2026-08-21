// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProhibitedTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProhibitedTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProhibitedTagsResponse
	GetStatusCode() *int32
	SetBody(v *ListProhibitedTagsResponseBody) *ListProhibitedTagsResponse
	GetBody() *ListProhibitedTagsResponseBody
}

type ListProhibitedTagsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProhibitedTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProhibitedTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedTagsResponse) GoString() string {
	return s.String()
}

func (s *ListProhibitedTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProhibitedTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProhibitedTagsResponse) GetBody() *ListProhibitedTagsResponseBody {
	return s.Body
}

func (s *ListProhibitedTagsResponse) SetHeaders(v map[string]*string) *ListProhibitedTagsResponse {
	s.Headers = v
	return s
}

func (s *ListProhibitedTagsResponse) SetStatusCode(v int32) *ListProhibitedTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProhibitedTagsResponse) SetBody(v *ListProhibitedTagsResponseBody) *ListProhibitedTagsResponse {
	s.Body = v
	return s
}

func (s *ListProhibitedTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
