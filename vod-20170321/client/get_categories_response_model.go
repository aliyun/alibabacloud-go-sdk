// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCategoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCategoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCategoriesResponse
	GetStatusCode() *int32
	SetBody(v *GetCategoriesResponseBody) *GetCategoriesResponse
	GetBody() *GetCategoriesResponseBody
}

type GetCategoriesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCategoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCategoriesResponse) GoString() string {
	return s.String()
}

func (s *GetCategoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCategoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCategoriesResponse) GetBody() *GetCategoriesResponseBody {
	return s.Body
}

func (s *GetCategoriesResponse) SetHeaders(v map[string]*string) *GetCategoriesResponse {
	s.Headers = v
	return s
}

func (s *GetCategoriesResponse) SetStatusCode(v int32) *GetCategoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCategoriesResponse) SetBody(v *GetCategoriesResponseBody) *GetCategoriesResponse {
	s.Body = v
	return s
}

func (s *GetCategoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
