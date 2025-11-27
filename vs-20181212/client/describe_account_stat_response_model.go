// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountStatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccountStatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccountStatResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccountStatResponseBody) *DescribeAccountStatResponse
	GetBody() *DescribeAccountStatResponseBody
}

type DescribeAccountStatResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountStatResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountStatResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountStatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccountStatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccountStatResponse) GetBody() *DescribeAccountStatResponseBody {
	return s.Body
}

func (s *DescribeAccountStatResponse) SetHeaders(v map[string]*string) *DescribeAccountStatResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountStatResponse) SetStatusCode(v int32) *DescribeAccountStatResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountStatResponse) SetBody(v *DescribeAccountStatResponseBody) *DescribeAccountStatResponse {
	s.Body = v
	return s
}

func (s *DescribeAccountStatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
