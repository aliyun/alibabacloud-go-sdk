// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudRecordStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudRecordStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudRecordStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudRecordStatusResponseBody) *DescribeCloudRecordStatusResponse
	GetBody() *DescribeCloudRecordStatusResponseBody
}

type DescribeCloudRecordStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudRecordStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudRecordStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudRecordStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudRecordStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudRecordStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudRecordStatusResponse) GetBody() *DescribeCloudRecordStatusResponseBody {
	return s.Body
}

func (s *DescribeCloudRecordStatusResponse) SetHeaders(v map[string]*string) *DescribeCloudRecordStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudRecordStatusResponse) SetStatusCode(v int32) *DescribeCloudRecordStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudRecordStatusResponse) SetBody(v *DescribeCloudRecordStatusResponseBody) *DescribeCloudRecordStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudRecordStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
