// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMediaCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMediaCategoryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMediaCategoryResponseBody) *UpdateMediaCategoryResponse
	GetBody() *UpdateMediaCategoryResponseBody
}

type UpdateMediaCategoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMediaCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMediaCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaCategoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateMediaCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMediaCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMediaCategoryResponse) GetBody() *UpdateMediaCategoryResponseBody {
	return s.Body
}

func (s *UpdateMediaCategoryResponse) SetHeaders(v map[string]*string) *UpdateMediaCategoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateMediaCategoryResponse) SetStatusCode(v int32) *UpdateMediaCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMediaCategoryResponse) SetBody(v *UpdateMediaCategoryResponseBody) *UpdateMediaCategoryResponse {
	s.Body = v
	return s
}

func (s *UpdateMediaCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
