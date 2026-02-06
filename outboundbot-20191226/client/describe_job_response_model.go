// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeJobResponse
	GetStatusCode() *int32
	SetBody(v *DescribeJobResponseBody) *DescribeJobResponse
	GetBody() *DescribeJobResponseBody
}

type DescribeJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeJobResponse) GetBody() *DescribeJobResponseBody {
	return s.Body
}

func (s *DescribeJobResponse) SetHeaders(v map[string]*string) *DescribeJobResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobResponse) SetStatusCode(v int32) *DescribeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJobResponse) SetBody(v *DescribeJobResponseBody) *DescribeJobResponse {
	s.Body = v
	return s
}

func (s *DescribeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
