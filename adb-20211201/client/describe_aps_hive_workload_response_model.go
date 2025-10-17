// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsHiveWorkloadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApsHiveWorkloadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApsHiveWorkloadResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApsHiveWorkloadResponseBody) *DescribeApsHiveWorkloadResponse
	GetBody() *DescribeApsHiveWorkloadResponseBody
}

type DescribeApsHiveWorkloadResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApsHiveWorkloadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApsHiveWorkloadResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsHiveWorkloadResponse) GoString() string {
	return s.String()
}

func (s *DescribeApsHiveWorkloadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApsHiveWorkloadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApsHiveWorkloadResponse) GetBody() *DescribeApsHiveWorkloadResponseBody {
	return s.Body
}

func (s *DescribeApsHiveWorkloadResponse) SetHeaders(v map[string]*string) *DescribeApsHiveWorkloadResponse {
	s.Headers = v
	return s
}

func (s *DescribeApsHiveWorkloadResponse) SetStatusCode(v int32) *DescribeApsHiveWorkloadResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponse) SetBody(v *DescribeApsHiveWorkloadResponseBody) *DescribeApsHiveWorkloadResponse {
	s.Body = v
	return s
}

func (s *DescribeApsHiveWorkloadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
