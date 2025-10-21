// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelCategoryResponse
	GetStatusCode() *int32
	SetBody(v *ListModelCategoryResponseBody) *ListModelCategoryResponse
	GetBody() *ListModelCategoryResponseBody
}

type ListModelCategoryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelCategoryResponse) GoString() string {
	return s.String()
}

func (s *ListModelCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelCategoryResponse) GetBody() *ListModelCategoryResponseBody {
	return s.Body
}

func (s *ListModelCategoryResponse) SetHeaders(v map[string]*string) *ListModelCategoryResponse {
	s.Headers = v
	return s
}

func (s *ListModelCategoryResponse) SetStatusCode(v int32) *ListModelCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelCategoryResponse) SetBody(v *ListModelCategoryResponseBody) *ListModelCategoryResponse {
	s.Body = v
	return s
}

func (s *ListModelCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
