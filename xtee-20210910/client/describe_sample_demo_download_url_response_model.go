// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleDemoDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSampleDemoDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSampleDemoDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSampleDemoDownloadUrlResponseBody) *DescribeSampleDemoDownloadUrlResponse
	GetBody() *DescribeSampleDemoDownloadUrlResponseBody
}

type DescribeSampleDemoDownloadUrlResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSampleDemoDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSampleDemoDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDemoDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeSampleDemoDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSampleDemoDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSampleDemoDownloadUrlResponse) GetBody() *DescribeSampleDemoDownloadUrlResponseBody {
	return s.Body
}

func (s *DescribeSampleDemoDownloadUrlResponse) SetHeaders(v map[string]*string) *DescribeSampleDemoDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeSampleDemoDownloadUrlResponse) SetStatusCode(v int32) *DescribeSampleDemoDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSampleDemoDownloadUrlResponse) SetBody(v *DescribeSampleDemoDownloadUrlResponseBody) *DescribeSampleDemoDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *DescribeSampleDemoDownloadUrlResponse) Validate() error {
	return dara.Validate(s)
}
