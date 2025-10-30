// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppKeyPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppKeyPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppKeyPageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppKeyPageResponseBody) *DescribeAppKeyPageResponse
	GetBody() *DescribeAppKeyPageResponseBody
}

type DescribeAppKeyPageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppKeyPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppKeyPageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppKeyPageResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppKeyPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppKeyPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppKeyPageResponse) GetBody() *DescribeAppKeyPageResponseBody {
	return s.Body
}

func (s *DescribeAppKeyPageResponse) SetHeaders(v map[string]*string) *DescribeAppKeyPageResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppKeyPageResponse) SetStatusCode(v int32) *DescribeAppKeyPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppKeyPageResponse) SetBody(v *DescribeAppKeyPageResponseBody) *DescribeAppKeyPageResponse {
	s.Body = v
	return s
}

func (s *DescribeAppKeyPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
