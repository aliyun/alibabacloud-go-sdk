// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSDKDownloadListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSDKDownloadListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSDKDownloadListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSDKDownloadListResponseBody) *DescribeSDKDownloadListResponse
	GetBody() *DescribeSDKDownloadListResponseBody
}

type DescribeSDKDownloadListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSDKDownloadListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSDKDownloadListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDKDownloadListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSDKDownloadListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSDKDownloadListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSDKDownloadListResponse) GetBody() *DescribeSDKDownloadListResponseBody {
	return s.Body
}

func (s *DescribeSDKDownloadListResponse) SetHeaders(v map[string]*string) *DescribeSDKDownloadListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSDKDownloadListResponse) SetStatusCode(v int32) *DescribeSDKDownloadListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSDKDownloadListResponse) SetBody(v *DescribeSDKDownloadListResponseBody) *DescribeSDKDownloadListResponse {
	s.Body = v
	return s
}

func (s *DescribeSDKDownloadListResponse) Validate() error {
	return dara.Validate(s)
}
