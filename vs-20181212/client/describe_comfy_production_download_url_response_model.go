// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyProductionDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeComfyProductionDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeComfyProductionDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeComfyProductionDownloadUrlResponseBody) *DescribeComfyProductionDownloadUrlResponse
	GetBody() *DescribeComfyProductionDownloadUrlResponseBody
}

type DescribeComfyProductionDownloadUrlResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComfyProductionDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComfyProductionDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyProductionDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeComfyProductionDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeComfyProductionDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeComfyProductionDownloadUrlResponse) GetBody() *DescribeComfyProductionDownloadUrlResponseBody {
	return s.Body
}

func (s *DescribeComfyProductionDownloadUrlResponse) SetHeaders(v map[string]*string) *DescribeComfyProductionDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeComfyProductionDownloadUrlResponse) SetStatusCode(v int32) *DescribeComfyProductionDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComfyProductionDownloadUrlResponse) SetBody(v *DescribeComfyProductionDownloadUrlResponseBody) *DescribeComfyProductionDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *DescribeComfyProductionDownloadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
