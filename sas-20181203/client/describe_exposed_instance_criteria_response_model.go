// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExposedInstanceCriteriaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExposedInstanceCriteriaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExposedInstanceCriteriaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExposedInstanceCriteriaResponseBody) *DescribeExposedInstanceCriteriaResponse
	GetBody() *DescribeExposedInstanceCriteriaResponseBody
}

type DescribeExposedInstanceCriteriaResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExposedInstanceCriteriaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExposedInstanceCriteriaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedInstanceCriteriaResponse) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceCriteriaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExposedInstanceCriteriaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExposedInstanceCriteriaResponse) GetBody() *DescribeExposedInstanceCriteriaResponseBody {
	return s.Body
}

func (s *DescribeExposedInstanceCriteriaResponse) SetHeaders(v map[string]*string) *DescribeExposedInstanceCriteriaResponse {
	s.Headers = v
	return s
}

func (s *DescribeExposedInstanceCriteriaResponse) SetStatusCode(v int32) *DescribeExposedInstanceCriteriaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExposedInstanceCriteriaResponse) SetBody(v *DescribeExposedInstanceCriteriaResponseBody) *DescribeExposedInstanceCriteriaResponse {
	s.Body = v
	return s
}

func (s *DescribeExposedInstanceCriteriaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
