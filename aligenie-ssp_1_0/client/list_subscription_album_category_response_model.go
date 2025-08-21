// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubscriptionAlbumCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSubscriptionAlbumCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSubscriptionAlbumCategoryResponse
	GetStatusCode() *int32
	SetBody(v *ListSubscriptionAlbumCategoryResponseBody) *ListSubscriptionAlbumCategoryResponse
	GetBody() *ListSubscriptionAlbumCategoryResponseBody
}

type ListSubscriptionAlbumCategoryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubscriptionAlbumCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubscriptionAlbumCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionAlbumCategoryResponse) GoString() string {
	return s.String()
}

func (s *ListSubscriptionAlbumCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSubscriptionAlbumCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSubscriptionAlbumCategoryResponse) GetBody() *ListSubscriptionAlbumCategoryResponseBody {
	return s.Body
}

func (s *ListSubscriptionAlbumCategoryResponse) SetHeaders(v map[string]*string) *ListSubscriptionAlbumCategoryResponse {
	s.Headers = v
	return s
}

func (s *ListSubscriptionAlbumCategoryResponse) SetStatusCode(v int32) *ListSubscriptionAlbumCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubscriptionAlbumCategoryResponse) SetBody(v *ListSubscriptionAlbumCategoryResponseBody) *ListSubscriptionAlbumCategoryResponse {
	s.Body = v
	return s
}

func (s *ListSubscriptionAlbumCategoryResponse) Validate() error {
	return dara.Validate(s)
}
