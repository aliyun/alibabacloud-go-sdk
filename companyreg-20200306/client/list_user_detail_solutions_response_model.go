// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDetailSolutionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserDetailSolutionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserDetailSolutionsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserDetailSolutionsResponseBody) *ListUserDetailSolutionsResponse
	GetBody() *ListUserDetailSolutionsResponseBody
}

type ListUserDetailSolutionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserDetailSolutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserDetailSolutionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserDetailSolutionsResponse) GoString() string {
	return s.String()
}

func (s *ListUserDetailSolutionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserDetailSolutionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserDetailSolutionsResponse) GetBody() *ListUserDetailSolutionsResponseBody {
	return s.Body
}

func (s *ListUserDetailSolutionsResponse) SetHeaders(v map[string]*string) *ListUserDetailSolutionsResponse {
	s.Headers = v
	return s
}

func (s *ListUserDetailSolutionsResponse) SetStatusCode(v int32) *ListUserDetailSolutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserDetailSolutionsResponse) SetBody(v *ListUserDetailSolutionsResponseBody) *ListUserDetailSolutionsResponse {
	s.Body = v
	return s
}

func (s *ListUserDetailSolutionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
