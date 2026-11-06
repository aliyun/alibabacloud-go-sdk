// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationAvailableVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationAvailableVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationAvailableVersionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationAvailableVersionResponseBody) *DescribeApplicationAvailableVersionResponse
	GetBody() *DescribeApplicationAvailableVersionResponseBody
}

type DescribeApplicationAvailableVersionResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationAvailableVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationAvailableVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationAvailableVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationAvailableVersionResponse) GetBody() *DescribeApplicationAvailableVersionResponseBody {
	return s.Body
}

func (s *DescribeApplicationAvailableVersionResponse) SetHeaders(v map[string]*string) *DescribeApplicationAvailableVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationAvailableVersionResponse) SetStatusCode(v int32) *DescribeApplicationAvailableVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationAvailableVersionResponse) SetBody(v *DescribeApplicationAvailableVersionResponseBody) *DescribeApplicationAvailableVersionResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationAvailableVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
