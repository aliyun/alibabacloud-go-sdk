// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageRepoCriteriaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageRepoCriteriaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageRepoCriteriaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageRepoCriteriaResponseBody) *DescribeImageRepoCriteriaResponse
	GetBody() *DescribeImageRepoCriteriaResponseBody
}

type DescribeImageRepoCriteriaResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageRepoCriteriaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageRepoCriteriaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageRepoCriteriaResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageRepoCriteriaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageRepoCriteriaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageRepoCriteriaResponse) GetBody() *DescribeImageRepoCriteriaResponseBody {
	return s.Body
}

func (s *DescribeImageRepoCriteriaResponse) SetHeaders(v map[string]*string) *DescribeImageRepoCriteriaResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageRepoCriteriaResponse) SetStatusCode(v int32) *DescribeImageRepoCriteriaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageRepoCriteriaResponse) SetBody(v *DescribeImageRepoCriteriaResponseBody) *DescribeImageRepoCriteriaResponse {
	s.Body = v
	return s
}

func (s *DescribeImageRepoCriteriaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
