// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBusinessCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBusinessCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBusinessCategoryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBusinessCategoryResponseBody) *DeleteBusinessCategoryResponse
	GetBody() *DeleteBusinessCategoryResponseBody
}

type DeleteBusinessCategoryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBusinessCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBusinessCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBusinessCategoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteBusinessCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBusinessCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBusinessCategoryResponse) GetBody() *DeleteBusinessCategoryResponseBody {
	return s.Body
}

func (s *DeleteBusinessCategoryResponse) SetHeaders(v map[string]*string) *DeleteBusinessCategoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteBusinessCategoryResponse) SetStatusCode(v int32) *DeleteBusinessCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBusinessCategoryResponse) SetBody(v *DeleteBusinessCategoryResponseBody) *DeleteBusinessCategoryResponse {
	s.Body = v
	return s
}

func (s *DeleteBusinessCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
