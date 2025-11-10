// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErrorCodeSolutionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetErrorCodeSolutionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetErrorCodeSolutionsResponse
	GetStatusCode() *int32
	SetBody(v *GetErrorCodeSolutionsResponseBody) *GetErrorCodeSolutionsResponse
	GetBody() *GetErrorCodeSolutionsResponseBody
}

type GetErrorCodeSolutionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetErrorCodeSolutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetErrorCodeSolutionsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetErrorCodeSolutionsResponse) GoString() string {
	return s.String()
}

func (s *GetErrorCodeSolutionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetErrorCodeSolutionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetErrorCodeSolutionsResponse) GetBody() *GetErrorCodeSolutionsResponseBody {
	return s.Body
}

func (s *GetErrorCodeSolutionsResponse) SetHeaders(v map[string]*string) *GetErrorCodeSolutionsResponse {
	s.Headers = v
	return s
}

func (s *GetErrorCodeSolutionsResponse) SetStatusCode(v int32) *GetErrorCodeSolutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetErrorCodeSolutionsResponse) SetBody(v *GetErrorCodeSolutionsResponseBody) *GetErrorCodeSolutionsResponse {
	s.Body = v
	return s
}

func (s *GetErrorCodeSolutionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
