// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQuaraFileDownloadInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeQuaraFileDownloadInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeQuaraFileDownloadInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeQuaraFileDownloadInfoResponseBody) *DescribeQuaraFileDownloadInfoResponse
	GetBody() *DescribeQuaraFileDownloadInfoResponseBody
}

type DescribeQuaraFileDownloadInfoResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQuaraFileDownloadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQuaraFileDownloadInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeQuaraFileDownloadInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeQuaraFileDownloadInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeQuaraFileDownloadInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeQuaraFileDownloadInfoResponse) GetBody() *DescribeQuaraFileDownloadInfoResponseBody {
	return s.Body
}

func (s *DescribeQuaraFileDownloadInfoResponse) SetHeaders(v map[string]*string) *DescribeQuaraFileDownloadInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeQuaraFileDownloadInfoResponse) SetStatusCode(v int32) *DescribeQuaraFileDownloadInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQuaraFileDownloadInfoResponse) SetBody(v *DescribeQuaraFileDownloadInfoResponseBody) *DescribeQuaraFileDownloadInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeQuaraFileDownloadInfoResponse) Validate() error {
	return dara.Validate(s)
}
