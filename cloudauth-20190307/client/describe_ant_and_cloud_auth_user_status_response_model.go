// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAntAndCloudAuthUserStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAntAndCloudAuthUserStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAntAndCloudAuthUserStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAntAndCloudAuthUserStatusResponseBody) *DescribeAntAndCloudAuthUserStatusResponse
	GetBody() *DescribeAntAndCloudAuthUserStatusResponseBody
}

type DescribeAntAndCloudAuthUserStatusResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAntAndCloudAuthUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAntAndCloudAuthUserStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAntAndCloudAuthUserStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntAndCloudAuthUserStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAntAndCloudAuthUserStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAntAndCloudAuthUserStatusResponse) GetBody() *DescribeAntAndCloudAuthUserStatusResponseBody {
	return s.Body
}

func (s *DescribeAntAndCloudAuthUserStatusResponse) SetHeaders(v map[string]*string) *DescribeAntAndCloudAuthUserStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntAndCloudAuthUserStatusResponse) SetStatusCode(v int32) *DescribeAntAndCloudAuthUserStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAntAndCloudAuthUserStatusResponse) SetBody(v *DescribeAntAndCloudAuthUserStatusResponseBody) *DescribeAntAndCloudAuthUserStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeAntAndCloudAuthUserStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
