// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCategoryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCategoryResponseBody) *UpdateCategoryResponse
	GetBody() *UpdateCategoryResponseBody
}

type UpdateCategoryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCategoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCategoryResponse) GetBody() *UpdateCategoryResponseBody {
	return s.Body
}

func (s *UpdateCategoryResponse) SetHeaders(v map[string]*string) *UpdateCategoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateCategoryResponse) SetStatusCode(v int32) *UpdateCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCategoryResponse) SetBody(v *UpdateCategoryResponseBody) *UpdateCategoryResponse {
	s.Body = v
	return s
}

func (s *UpdateCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
