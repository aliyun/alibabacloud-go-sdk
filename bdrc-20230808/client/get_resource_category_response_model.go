// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceCategoryResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceCategoryResponseBody) *GetResourceCategoryResponse
	GetBody() *GetResourceCategoryResponseBody
}

type GetResourceCategoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceCategoryResponse) GoString() string {
	return s.String()
}

func (s *GetResourceCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceCategoryResponse) GetBody() *GetResourceCategoryResponseBody {
	return s.Body
}

func (s *GetResourceCategoryResponse) SetHeaders(v map[string]*string) *GetResourceCategoryResponse {
	s.Headers = v
	return s
}

func (s *GetResourceCategoryResponse) SetStatusCode(v int32) *GetResourceCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceCategoryResponse) SetBody(v *GetResourceCategoryResponseBody) *GetResourceCategoryResponse {
	s.Body = v
	return s
}

func (s *GetResourceCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
