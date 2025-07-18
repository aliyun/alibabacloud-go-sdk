// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamingDataServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStreamingDataServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStreamingDataServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStreamingDataServiceResponseBody) *DescribeStreamingDataServiceResponse
	GetBody() *DescribeStreamingDataServiceResponseBody
}

type DescribeStreamingDataServiceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStreamingDataServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStreamingDataServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamingDataServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeStreamingDataServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStreamingDataServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStreamingDataServiceResponse) GetBody() *DescribeStreamingDataServiceResponseBody {
	return s.Body
}

func (s *DescribeStreamingDataServiceResponse) SetHeaders(v map[string]*string) *DescribeStreamingDataServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeStreamingDataServiceResponse) SetStatusCode(v int32) *DescribeStreamingDataServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStreamingDataServiceResponse) SetBody(v *DescribeStreamingDataServiceResponseBody) *DescribeStreamingDataServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeStreamingDataServiceResponse) Validate() error {
	return dara.Validate(s)
}
