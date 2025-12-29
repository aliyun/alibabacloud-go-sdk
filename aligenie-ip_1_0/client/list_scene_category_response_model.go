// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSceneCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSceneCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSceneCategoryResponse
	GetStatusCode() *int32
	SetBody(v *ListSceneCategoryResponseBody) *ListSceneCategoryResponse
	GetBody() *ListSceneCategoryResponseBody
}

type ListSceneCategoryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSceneCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSceneCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSceneCategoryResponse) GoString() string {
	return s.String()
}

func (s *ListSceneCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSceneCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSceneCategoryResponse) GetBody() *ListSceneCategoryResponseBody {
	return s.Body
}

func (s *ListSceneCategoryResponse) SetHeaders(v map[string]*string) *ListSceneCategoryResponse {
	s.Headers = v
	return s
}

func (s *ListSceneCategoryResponse) SetStatusCode(v int32) *ListSceneCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSceneCategoryResponse) SetBody(v *ListSceneCategoryResponseBody) *ListSceneCategoryResponse {
	s.Body = v
	return s
}

func (s *ListSceneCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
