// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateResourceCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateResourceCategoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateResourceCategoryResponseBody) *CreateResourceCategoryResponse
	GetBody() *CreateResourceCategoryResponseBody
}

type CreateResourceCategoryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceCategoryResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateResourceCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateResourceCategoryResponse) GetBody() *CreateResourceCategoryResponseBody {
	return s.Body
}

func (s *CreateResourceCategoryResponse) SetHeaders(v map[string]*string) *CreateResourceCategoryResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceCategoryResponse) SetStatusCode(v int32) *CreateResourceCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceCategoryResponse) SetBody(v *CreateResourceCategoryResponseBody) *CreateResourceCategoryResponse {
	s.Body = v
	return s
}

func (s *CreateResourceCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
