// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCategoriesByTaskIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCategoriesByTaskIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCategoriesByTaskIdResponse
	GetStatusCode() *int32
	SetBody(v *GetCategoriesByTaskIdResponseBody) *GetCategoriesByTaskIdResponse
	GetBody() *GetCategoriesByTaskIdResponseBody
}

type GetCategoriesByTaskIdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCategoriesByTaskIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCategoriesByTaskIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCategoriesByTaskIdResponse) GoString() string {
	return s.String()
}

func (s *GetCategoriesByTaskIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCategoriesByTaskIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCategoriesByTaskIdResponse) GetBody() *GetCategoriesByTaskIdResponseBody {
	return s.Body
}

func (s *GetCategoriesByTaskIdResponse) SetHeaders(v map[string]*string) *GetCategoriesByTaskIdResponse {
	s.Headers = v
	return s
}

func (s *GetCategoriesByTaskIdResponse) SetStatusCode(v int32) *GetCategoriesByTaskIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCategoriesByTaskIdResponse) SetBody(v *GetCategoriesByTaskIdResponseBody) *GetCategoriesByTaskIdResponse {
	s.Body = v
	return s
}

func (s *GetCategoriesByTaskIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
