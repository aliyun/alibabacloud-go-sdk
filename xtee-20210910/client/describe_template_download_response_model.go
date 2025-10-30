// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateDownloadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTemplateDownloadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTemplateDownloadResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTemplateDownloadResponseBody) *DescribeTemplateDownloadResponse
	GetBody() *DescribeTemplateDownloadResponseBody
}

type DescribeTemplateDownloadResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTemplateDownloadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTemplateDownloadResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateDownloadResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplateDownloadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTemplateDownloadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTemplateDownloadResponse) GetBody() *DescribeTemplateDownloadResponseBody {
	return s.Body
}

func (s *DescribeTemplateDownloadResponse) SetHeaders(v map[string]*string) *DescribeTemplateDownloadResponse {
	s.Headers = v
	return s
}

func (s *DescribeTemplateDownloadResponse) SetStatusCode(v int32) *DescribeTemplateDownloadResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTemplateDownloadResponse) SetBody(v *DescribeTemplateDownloadResponseBody) *DescribeTemplateDownloadResponse {
	s.Body = v
	return s
}

func (s *DescribeTemplateDownloadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
