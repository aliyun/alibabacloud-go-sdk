// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKernelReleaseNotesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKernelReleaseNotesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKernelReleaseNotesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKernelReleaseNotesResponseBody) *DescribeKernelReleaseNotesResponse
	GetBody() *DescribeKernelReleaseNotesResponseBody
}

type DescribeKernelReleaseNotesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKernelReleaseNotesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKernelReleaseNotesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKernelReleaseNotesResponse) GoString() string {
	return s.String()
}

func (s *DescribeKernelReleaseNotesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKernelReleaseNotesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKernelReleaseNotesResponse) GetBody() *DescribeKernelReleaseNotesResponseBody {
	return s.Body
}

func (s *DescribeKernelReleaseNotesResponse) SetHeaders(v map[string]*string) *DescribeKernelReleaseNotesResponse {
	s.Headers = v
	return s
}

func (s *DescribeKernelReleaseNotesResponse) SetStatusCode(v int32) *DescribeKernelReleaseNotesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKernelReleaseNotesResponse) SetBody(v *DescribeKernelReleaseNotesResponseBody) *DescribeKernelReleaseNotesResponse {
	s.Body = v
	return s
}

func (s *DescribeKernelReleaseNotesResponse) Validate() error {
	return dara.Validate(s)
}
