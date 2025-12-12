// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBusinessCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddBusinessCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddBusinessCategoryResponse
	GetStatusCode() *int32
	SetBody(v *AddBusinessCategoryResponseBody) *AddBusinessCategoryResponse
	GetBody() *AddBusinessCategoryResponseBody
}

type AddBusinessCategoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddBusinessCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddBusinessCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s AddBusinessCategoryResponse) GoString() string {
	return s.String()
}

func (s *AddBusinessCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddBusinessCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddBusinessCategoryResponse) GetBody() *AddBusinessCategoryResponseBody {
	return s.Body
}

func (s *AddBusinessCategoryResponse) SetHeaders(v map[string]*string) *AddBusinessCategoryResponse {
	s.Headers = v
	return s
}

func (s *AddBusinessCategoryResponse) SetStatusCode(v int32) *AddBusinessCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddBusinessCategoryResponse) SetBody(v *AddBusinessCategoryResponseBody) *AddBusinessCategoryResponse {
	s.Body = v
	return s
}

func (s *AddBusinessCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
