// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLinkInstanceCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LinkInstanceCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LinkInstanceCategoryResponse
	GetStatusCode() *int32
	SetBody(v *LinkInstanceCategoryResponseBody) *LinkInstanceCategoryResponse
	GetBody() *LinkInstanceCategoryResponseBody
}

type LinkInstanceCategoryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LinkInstanceCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LinkInstanceCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s LinkInstanceCategoryResponse) GoString() string {
	return s.String()
}

func (s *LinkInstanceCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LinkInstanceCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LinkInstanceCategoryResponse) GetBody() *LinkInstanceCategoryResponseBody {
	return s.Body
}

func (s *LinkInstanceCategoryResponse) SetHeaders(v map[string]*string) *LinkInstanceCategoryResponse {
	s.Headers = v
	return s
}

func (s *LinkInstanceCategoryResponse) SetStatusCode(v int32) *LinkInstanceCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *LinkInstanceCategoryResponse) SetBody(v *LinkInstanceCategoryResponseBody) *LinkInstanceCategoryResponse {
	s.Body = v
	return s
}

func (s *LinkInstanceCategoryResponse) Validate() error {
	return dara.Validate(s)
}
