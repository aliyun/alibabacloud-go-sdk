// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAttackAppCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAttackAppCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAttackAppCategoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAttackAppCategoryResponseBody) *DescribeAttackAppCategoryResponse
	GetBody() *DescribeAttackAppCategoryResponseBody
}

type DescribeAttackAppCategoryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAttackAppCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAttackAppCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttackAppCategoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeAttackAppCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAttackAppCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAttackAppCategoryResponse) GetBody() *DescribeAttackAppCategoryResponseBody {
	return s.Body
}

func (s *DescribeAttackAppCategoryResponse) SetHeaders(v map[string]*string) *DescribeAttackAppCategoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeAttackAppCategoryResponse) SetStatusCode(v int32) *DescribeAttackAppCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAttackAppCategoryResponse) SetBody(v *DescribeAttackAppCategoryResponseBody) *DescribeAttackAppCategoryResponse {
	s.Body = v
	return s
}

func (s *DescribeAttackAppCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
