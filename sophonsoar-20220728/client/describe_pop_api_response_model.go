// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePopApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePopApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePopApiResponse
	GetStatusCode() *int32
	SetBody(v *DescribePopApiResponseBody) *DescribePopApiResponse
	GetBody() *DescribePopApiResponseBody
}

type DescribePopApiResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePopApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePopApiResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePopApiResponse) GoString() string {
	return s.String()
}

func (s *DescribePopApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePopApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePopApiResponse) GetBody() *DescribePopApiResponseBody {
	return s.Body
}

func (s *DescribePopApiResponse) SetHeaders(v map[string]*string) *DescribePopApiResponse {
	s.Headers = v
	return s
}

func (s *DescribePopApiResponse) SetStatusCode(v int32) *DescribePopApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePopApiResponse) SetBody(v *DescribePopApiResponseBody) *DescribePopApiResponse {
	s.Body = v
	return s
}

func (s *DescribePopApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
