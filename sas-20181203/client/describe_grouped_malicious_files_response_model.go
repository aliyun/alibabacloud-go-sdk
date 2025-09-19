// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupedMaliciousFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupedMaliciousFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupedMaliciousFilesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupedMaliciousFilesResponseBody) *DescribeGroupedMaliciousFilesResponse
	GetBody() *DescribeGroupedMaliciousFilesResponseBody
}

type DescribeGroupedMaliciousFilesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupedMaliciousFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupedMaliciousFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedMaliciousFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupedMaliciousFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupedMaliciousFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupedMaliciousFilesResponse) GetBody() *DescribeGroupedMaliciousFilesResponseBody {
	return s.Body
}

func (s *DescribeGroupedMaliciousFilesResponse) SetHeaders(v map[string]*string) *DescribeGroupedMaliciousFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponse) SetStatusCode(v int32) *DescribeGroupedMaliciousFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponse) SetBody(v *DescribeGroupedMaliciousFilesResponseBody) *DescribeGroupedMaliciousFilesResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponse) Validate() error {
	return dara.Validate(s)
}
