// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserSolutionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserSolutionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserSolutionsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserSolutionsResponseBody) *ListUserSolutionsResponse
	GetBody() *ListUserSolutionsResponseBody
}

type ListUserSolutionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserSolutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserSolutionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserSolutionsResponse) GoString() string {
	return s.String()
}

func (s *ListUserSolutionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserSolutionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserSolutionsResponse) GetBody() *ListUserSolutionsResponseBody {
	return s.Body
}

func (s *ListUserSolutionsResponse) SetHeaders(v map[string]*string) *ListUserSolutionsResponse {
	s.Headers = v
	return s
}

func (s *ListUserSolutionsResponse) SetStatusCode(v int32) *ListUserSolutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserSolutionsResponse) SetBody(v *ListUserSolutionsResponseBody) *ListUserSolutionsResponse {
	s.Body = v
	return s
}

func (s *ListUserSolutionsResponse) Validate() error {
	return dara.Validate(s)
}
