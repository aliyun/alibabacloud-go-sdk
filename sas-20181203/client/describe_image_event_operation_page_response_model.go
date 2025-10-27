// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageEventOperationPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageEventOperationPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageEventOperationPageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageEventOperationPageResponseBody) *DescribeImageEventOperationPageResponse
	GetBody() *DescribeImageEventOperationPageResponseBody
}

type DescribeImageEventOperationPageResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageEventOperationPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageEventOperationPageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageEventOperationPageResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageEventOperationPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageEventOperationPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageEventOperationPageResponse) GetBody() *DescribeImageEventOperationPageResponseBody {
	return s.Body
}

func (s *DescribeImageEventOperationPageResponse) SetHeaders(v map[string]*string) *DescribeImageEventOperationPageResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageEventOperationPageResponse) SetStatusCode(v int32) *DescribeImageEventOperationPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageEventOperationPageResponse) SetBody(v *DescribeImageEventOperationPageResponseBody) *DescribeImageEventOperationPageResponse {
	s.Body = v
	return s
}

func (s *DescribeImageEventOperationPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
