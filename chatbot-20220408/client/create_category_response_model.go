// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCategoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateCategoryResponseBody) *CreateCategoryResponse
	GetBody() *CreateCategoryResponseBody
}

type CreateCategoryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCategoryResponse) GoString() string {
	return s.String()
}

func (s *CreateCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCategoryResponse) GetBody() *CreateCategoryResponseBody {
	return s.Body
}

func (s *CreateCategoryResponse) SetHeaders(v map[string]*string) *CreateCategoryResponse {
	s.Headers = v
	return s
}

func (s *CreateCategoryResponse) SetStatusCode(v int32) *CreateCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCategoryResponse) SetBody(v *CreateCategoryResponseBody) *CreateCategoryResponse {
	s.Body = v
	return s
}

func (s *CreateCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
