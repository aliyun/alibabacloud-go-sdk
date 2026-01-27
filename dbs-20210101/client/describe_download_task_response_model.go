// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDownloadTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDownloadTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDownloadTaskResponseBody) *DescribeDownloadTaskResponse
	GetBody() *DescribeDownloadTaskResponseBody
}

type DescribeDownloadTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDownloadTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDownloadTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDownloadTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDownloadTaskResponse) GetBody() *DescribeDownloadTaskResponseBody {
	return s.Body
}

func (s *DescribeDownloadTaskResponse) SetHeaders(v map[string]*string) *DescribeDownloadTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeDownloadTaskResponse) SetStatusCode(v int32) *DescribeDownloadTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDownloadTaskResponse) SetBody(v *DescribeDownloadTaskResponseBody) *DescribeDownloadTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeDownloadTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
