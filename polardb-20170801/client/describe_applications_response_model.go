// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationsResponseBody) *DescribeApplicationsResponse
	GetBody() *DescribeApplicationsResponseBody
}

type DescribeApplicationsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationsResponse) GetBody() *DescribeApplicationsResponseBody {
	return s.Body
}

func (s *DescribeApplicationsResponse) SetHeaders(v map[string]*string) *DescribeApplicationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationsResponse) SetStatusCode(v int32) *DescribeApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationsResponse) SetBody(v *DescribeApplicationsResponseBody) *DescribeApplicationsResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
