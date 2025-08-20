// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeControllerDetectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeControllerDetectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeControllerDetectionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeControllerDetectionResponseBody) *DescribeControllerDetectionResponse
	GetBody() *DescribeControllerDetectionResponseBody
}

type DescribeControllerDetectionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeControllerDetectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeControllerDetectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeControllerDetectionResponse) GoString() string {
	return s.String()
}

func (s *DescribeControllerDetectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeControllerDetectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeControllerDetectionResponse) GetBody() *DescribeControllerDetectionResponseBody {
	return s.Body
}

func (s *DescribeControllerDetectionResponse) SetHeaders(v map[string]*string) *DescribeControllerDetectionResponse {
	s.Headers = v
	return s
}

func (s *DescribeControllerDetectionResponse) SetStatusCode(v int32) *DescribeControllerDetectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeControllerDetectionResponse) SetBody(v *DescribeControllerDetectionResponseBody) *DescribeControllerDetectionResponse {
	s.Body = v
	return s
}

func (s *DescribeControllerDetectionResponse) Validate() error {
	return dara.Validate(s)
}
