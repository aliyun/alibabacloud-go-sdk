// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCategoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCategoryResponseBody) *DescribeCategoryResponse
	GetBody() *DescribeCategoryResponseBody
}

type DescribeCategoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCategoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCategoryResponse) GetBody() *DescribeCategoryResponseBody {
	return s.Body
}

func (s *DescribeCategoryResponse) SetHeaders(v map[string]*string) *DescribeCategoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeCategoryResponse) SetStatusCode(v int32) *DescribeCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCategoryResponse) SetBody(v *DescribeCategoryResponseBody) *DescribeCategoryResponse {
	s.Body = v
	return s
}

func (s *DescribeCategoryResponse) Validate() error {
	return dara.Validate(s)
}
