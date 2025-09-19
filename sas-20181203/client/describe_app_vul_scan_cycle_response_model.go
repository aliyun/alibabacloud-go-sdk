// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppVulScanCycleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppVulScanCycleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppVulScanCycleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppVulScanCycleResponseBody) *DescribeAppVulScanCycleResponse
	GetBody() *DescribeAppVulScanCycleResponseBody
}

type DescribeAppVulScanCycleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppVulScanCycleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppVulScanCycleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppVulScanCycleResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppVulScanCycleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppVulScanCycleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppVulScanCycleResponse) GetBody() *DescribeAppVulScanCycleResponseBody {
	return s.Body
}

func (s *DescribeAppVulScanCycleResponse) SetHeaders(v map[string]*string) *DescribeAppVulScanCycleResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppVulScanCycleResponse) SetStatusCode(v int32) *DescribeAppVulScanCycleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppVulScanCycleResponse) SetBody(v *DescribeAppVulScanCycleResponseBody) *DescribeAppVulScanCycleResponse {
	s.Body = v
	return s
}

func (s *DescribeAppVulScanCycleResponse) Validate() error {
	return dara.Validate(s)
}
