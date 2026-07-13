// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteResourceCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteResourceCategoryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteResourceCategoryResponseBody) *DeleteResourceCategoryResponse
	GetBody() *DeleteResourceCategoryResponseBody
}

type DeleteResourceCategoryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceCategoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteResourceCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteResourceCategoryResponse) GetBody() *DeleteResourceCategoryResponseBody {
	return s.Body
}

func (s *DeleteResourceCategoryResponse) SetHeaders(v map[string]*string) *DeleteResourceCategoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceCategoryResponse) SetStatusCode(v int32) *DeleteResourceCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceCategoryResponse) SetBody(v *DeleteResourceCategoryResponseBody) *DeleteResourceCategoryResponse {
	s.Body = v
	return s
}

func (s *DeleteResourceCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
