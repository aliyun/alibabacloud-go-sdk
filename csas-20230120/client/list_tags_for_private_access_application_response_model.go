// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsForPrivateAccessApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTagsForPrivateAccessApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTagsForPrivateAccessApplicationResponse
	GetStatusCode() *int32
	SetBody(v *ListTagsForPrivateAccessApplicationResponseBody) *ListTagsForPrivateAccessApplicationResponse
	GetBody() *ListTagsForPrivateAccessApplicationResponseBody
}

type ListTagsForPrivateAccessApplicationResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagsForPrivateAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagsForPrivateAccessApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTagsForPrivateAccessApplicationResponse) GoString() string {
	return s.String()
}

func (s *ListTagsForPrivateAccessApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTagsForPrivateAccessApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTagsForPrivateAccessApplicationResponse) GetBody() *ListTagsForPrivateAccessApplicationResponseBody {
	return s.Body
}

func (s *ListTagsForPrivateAccessApplicationResponse) SetHeaders(v map[string]*string) *ListTagsForPrivateAccessApplicationResponse {
	s.Headers = v
	return s
}

func (s *ListTagsForPrivateAccessApplicationResponse) SetStatusCode(v int32) *ListTagsForPrivateAccessApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagsForPrivateAccessApplicationResponse) SetBody(v *ListTagsForPrivateAccessApplicationResponseBody) *ListTagsForPrivateAccessApplicationResponse {
	s.Body = v
	return s
}

func (s *ListTagsForPrivateAccessApplicationResponse) Validate() error {
	return dara.Validate(s)
}
