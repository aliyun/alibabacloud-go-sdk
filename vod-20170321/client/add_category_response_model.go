// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCategoryResponse
	GetStatusCode() *int32
	SetBody(v *AddCategoryResponseBody) *AddCategoryResponse
	GetBody() *AddCategoryResponseBody
}

type AddCategoryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCategoryResponse) GoString() string {
	return s.String()
}

func (s *AddCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCategoryResponse) GetBody() *AddCategoryResponseBody {
	return s.Body
}

func (s *AddCategoryResponse) SetHeaders(v map[string]*string) *AddCategoryResponse {
	s.Headers = v
	return s
}

func (s *AddCategoryResponse) SetStatusCode(v int32) *AddCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCategoryResponse) SetBody(v *AddCategoryResponseBody) *AddCategoryResponse {
	s.Body = v
	return s
}

func (s *AddCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
