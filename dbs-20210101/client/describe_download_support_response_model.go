// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadSupportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDownloadSupportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDownloadSupportResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDownloadSupportResponseBody) *DescribeDownloadSupportResponse
	GetBody() *DescribeDownloadSupportResponseBody
}

type DescribeDownloadSupportResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDownloadSupportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDownloadSupportResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadSupportResponse) GoString() string {
	return s.String()
}

func (s *DescribeDownloadSupportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDownloadSupportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDownloadSupportResponse) GetBody() *DescribeDownloadSupportResponseBody {
	return s.Body
}

func (s *DescribeDownloadSupportResponse) SetHeaders(v map[string]*string) *DescribeDownloadSupportResponse {
	s.Headers = v
	return s
}

func (s *DescribeDownloadSupportResponse) SetStatusCode(v int32) *DescribeDownloadSupportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDownloadSupportResponse) SetBody(v *DescribeDownloadSupportResponseBody) *DescribeDownloadSupportResponse {
	s.Body = v
	return s
}

func (s *DescribeDownloadSupportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
