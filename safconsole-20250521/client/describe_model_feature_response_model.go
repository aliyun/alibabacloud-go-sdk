// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelFeatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeModelFeatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeModelFeatureResponse
	GetStatusCode() *int32
	SetBody(v *DescribeModelFeatureResponseBody) *DescribeModelFeatureResponse
	GetBody() *DescribeModelFeatureResponseBody
}

type DescribeModelFeatureResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeModelFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeModelFeatureResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelFeatureResponse) GoString() string {
	return s.String()
}

func (s *DescribeModelFeatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeModelFeatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeModelFeatureResponse) GetBody() *DescribeModelFeatureResponseBody {
	return s.Body
}

func (s *DescribeModelFeatureResponse) SetHeaders(v map[string]*string) *DescribeModelFeatureResponse {
	s.Headers = v
	return s
}

func (s *DescribeModelFeatureResponse) SetStatusCode(v int32) *DescribeModelFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModelFeatureResponse) SetBody(v *DescribeModelFeatureResponseBody) *DescribeModelFeatureResponse {
	s.Body = v
	return s
}

func (s *DescribeModelFeatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
