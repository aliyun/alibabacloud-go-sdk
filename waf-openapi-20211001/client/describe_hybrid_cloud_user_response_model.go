// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridCloudUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridCloudUserResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridCloudUserResponseBody) *DescribeHybridCloudUserResponse
	GetBody() *DescribeHybridCloudUserResponseBody
}

type DescribeHybridCloudUserResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridCloudUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridCloudUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudUserResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridCloudUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridCloudUserResponse) GetBody() *DescribeHybridCloudUserResponseBody {
	return s.Body
}

func (s *DescribeHybridCloudUserResponse) SetHeaders(v map[string]*string) *DescribeHybridCloudUserResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridCloudUserResponse) SetStatusCode(v int32) *DescribeHybridCloudUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridCloudUserResponse) SetBody(v *DescribeHybridCloudUserResponseBody) *DescribeHybridCloudUserResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridCloudUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
