// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFeatureOptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFeatureOptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFeatureOptionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFeatureOptionResponseBody) *DescribeFeatureOptionResponse
	GetBody() *DescribeFeatureOptionResponseBody
}

type DescribeFeatureOptionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFeatureOptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFeatureOptionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFeatureOptionResponse) GoString() string {
	return s.String()
}

func (s *DescribeFeatureOptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFeatureOptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFeatureOptionResponse) GetBody() *DescribeFeatureOptionResponseBody {
	return s.Body
}

func (s *DescribeFeatureOptionResponse) SetHeaders(v map[string]*string) *DescribeFeatureOptionResponse {
	s.Headers = v
	return s
}

func (s *DescribeFeatureOptionResponse) SetStatusCode(v int32) *DescribeFeatureOptionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFeatureOptionResponse) SetBody(v *DescribeFeatureOptionResponseBody) *DescribeFeatureOptionResponse {
	s.Body = v
	return s
}

func (s *DescribeFeatureOptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
