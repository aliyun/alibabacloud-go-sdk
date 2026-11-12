// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSharedDiskRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSharedDiskRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSharedDiskRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSharedDiskRecordsResponseBody) *DescribeSharedDiskRecordsResponse
	GetBody() *DescribeSharedDiskRecordsResponseBody
}

type DescribeSharedDiskRecordsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSharedDiskRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSharedDiskRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSharedDiskRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSharedDiskRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSharedDiskRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSharedDiskRecordsResponse) GetBody() *DescribeSharedDiskRecordsResponseBody {
	return s.Body
}

func (s *DescribeSharedDiskRecordsResponse) SetHeaders(v map[string]*string) *DescribeSharedDiskRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSharedDiskRecordsResponse) SetStatusCode(v int32) *DescribeSharedDiskRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSharedDiskRecordsResponse) SetBody(v *DescribeSharedDiskRecordsResponseBody) *DescribeSharedDiskRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeSharedDiskRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
