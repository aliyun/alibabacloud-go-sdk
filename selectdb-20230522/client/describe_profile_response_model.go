// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProfileResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProfileResponseBody) *DescribeProfileResponse
	GetBody() *DescribeProfileResponseBody
}

type DescribeProfileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProfileResponse) GoString() string {
	return s.String()
}

func (s *DescribeProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProfileResponse) GetBody() *DescribeProfileResponseBody {
	return s.Body
}

func (s *DescribeProfileResponse) SetHeaders(v map[string]*string) *DescribeProfileResponse {
	s.Headers = v
	return s
}

func (s *DescribeProfileResponse) SetStatusCode(v int32) *DescribeProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProfileResponse) SetBody(v *DescribeProfileResponseBody) *DescribeProfileResponse {
	s.Body = v
	return s
}

func (s *DescribeProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
