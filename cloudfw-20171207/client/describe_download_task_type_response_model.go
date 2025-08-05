// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadTaskTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDownloadTaskTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDownloadTaskTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDownloadTaskTypeResponseBody) *DescribeDownloadTaskTypeResponse
	GetBody() *DescribeDownloadTaskTypeResponseBody
}

type DescribeDownloadTaskTypeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDownloadTaskTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDownloadTaskTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadTaskTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDownloadTaskTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDownloadTaskTypeResponse) GetBody() *DescribeDownloadTaskTypeResponseBody {
	return s.Body
}

func (s *DescribeDownloadTaskTypeResponse) SetHeaders(v map[string]*string) *DescribeDownloadTaskTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDownloadTaskTypeResponse) SetStatusCode(v int32) *DescribeDownloadTaskTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDownloadTaskTypeResponse) SetBody(v *DescribeDownloadTaskTypeResponseBody) *DescribeDownloadTaskTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeDownloadTaskTypeResponse) Validate() error {
	return dara.Validate(s)
}
