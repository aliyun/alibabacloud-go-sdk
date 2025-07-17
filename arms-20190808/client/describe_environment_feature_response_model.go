// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnvironmentFeatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnvironmentFeatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnvironmentFeatureResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnvironmentFeatureResponseBody) *DescribeEnvironmentFeatureResponse
	GetBody() *DescribeEnvironmentFeatureResponseBody
}

type DescribeEnvironmentFeatureResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnvironmentFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnvironmentFeatureResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvironmentFeatureResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnvironmentFeatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnvironmentFeatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnvironmentFeatureResponse) GetBody() *DescribeEnvironmentFeatureResponseBody {
	return s.Body
}

func (s *DescribeEnvironmentFeatureResponse) SetHeaders(v map[string]*string) *DescribeEnvironmentFeatureResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnvironmentFeatureResponse) SetStatusCode(v int32) *DescribeEnvironmentFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnvironmentFeatureResponse) SetBody(v *DescribeEnvironmentFeatureResponseBody) *DescribeEnvironmentFeatureResponse {
	s.Body = v
	return s
}

func (s *DescribeEnvironmentFeatureResponse) Validate() error {
	return dara.Validate(s)
}
