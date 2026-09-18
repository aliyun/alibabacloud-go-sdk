// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCategoryDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCategoryDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCategoryDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetCategoryDetailResponseBody) *GetCategoryDetailResponse
	GetBody() *GetCategoryDetailResponseBody
}

type GetCategoryDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCategoryDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCategoryDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCategoryDetailResponse) GoString() string {
	return s.String()
}

func (s *GetCategoryDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCategoryDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCategoryDetailResponse) GetBody() *GetCategoryDetailResponseBody {
	return s.Body
}

func (s *GetCategoryDetailResponse) SetHeaders(v map[string]*string) *GetCategoryDetailResponse {
	s.Headers = v
	return s
}

func (s *GetCategoryDetailResponse) SetStatusCode(v int32) *GetCategoryDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCategoryDetailResponse) SetBody(v *GetCategoryDetailResponseBody) *GetCategoryDetailResponse {
	s.Body = v
	return s
}

func (s *GetCategoryDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
