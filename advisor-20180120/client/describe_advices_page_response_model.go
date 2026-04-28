// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvicesPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAdvicesPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAdvicesPageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAdvicesPageResponseBody) *DescribeAdvicesPageResponse
	GetBody() *DescribeAdvicesPageResponseBody
}

type DescribeAdvicesPageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdvicesPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdvicesPageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvicesPageResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAdvicesPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAdvicesPageResponse) GetBody() *DescribeAdvicesPageResponseBody {
	return s.Body
}

func (s *DescribeAdvicesPageResponse) SetHeaders(v map[string]*string) *DescribeAdvicesPageResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdvicesPageResponse) SetStatusCode(v int32) *DescribeAdvicesPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdvicesPageResponse) SetBody(v *DescribeAdvicesPageResponseBody) *DescribeAdvicesPageResponse {
	s.Body = v
	return s
}

func (s *DescribeAdvicesPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
