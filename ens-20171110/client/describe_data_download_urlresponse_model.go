// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataDownloadURLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataDownloadURLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataDownloadURLResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataDownloadURLResponseBody) *DescribeDataDownloadURLResponse
	GetBody() *DescribeDataDownloadURLResponseBody
}

type DescribeDataDownloadURLResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataDownloadURLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataDownloadURLResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataDownloadURLResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataDownloadURLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataDownloadURLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataDownloadURLResponse) GetBody() *DescribeDataDownloadURLResponseBody {
	return s.Body
}

func (s *DescribeDataDownloadURLResponse) SetHeaders(v map[string]*string) *DescribeDataDownloadURLResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataDownloadURLResponse) SetStatusCode(v int32) *DescribeDataDownloadURLResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataDownloadURLResponse) SetBody(v *DescribeDataDownloadURLResponseBody) *DescribeDataDownloadURLResponse {
	s.Body = v
	return s
}

func (s *DescribeDataDownloadURLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
