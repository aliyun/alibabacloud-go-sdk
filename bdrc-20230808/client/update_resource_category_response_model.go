// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateResourceCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateResourceCategoryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateResourceCategoryResponseBody) *UpdateResourceCategoryResponse
	GetBody() *UpdateResourceCategoryResponseBody
}

type UpdateResourceCategoryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceCategoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateResourceCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateResourceCategoryResponse) GetBody() *UpdateResourceCategoryResponseBody {
	return s.Body
}

func (s *UpdateResourceCategoryResponse) SetHeaders(v map[string]*string) *UpdateResourceCategoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceCategoryResponse) SetStatusCode(v int32) *UpdateResourceCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceCategoryResponse) SetBody(v *UpdateResourceCategoryResponseBody) *UpdateResourceCategoryResponse {
	s.Body = v
	return s
}

func (s *UpdateResourceCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
