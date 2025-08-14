// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSampleDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSampleDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSampleDownloadUrlResponseBody) *DescribeSampleDownloadUrlResponse
	GetBody() *DescribeSampleDownloadUrlResponseBody
}

type DescribeSampleDownloadUrlResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSampleDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSampleDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeSampleDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSampleDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSampleDownloadUrlResponse) GetBody() *DescribeSampleDownloadUrlResponseBody {
	return s.Body
}

func (s *DescribeSampleDownloadUrlResponse) SetHeaders(v map[string]*string) *DescribeSampleDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeSampleDownloadUrlResponse) SetStatusCode(v int32) *DescribeSampleDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSampleDownloadUrlResponse) SetBody(v *DescribeSampleDownloadUrlResponseBody) *DescribeSampleDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *DescribeSampleDownloadUrlResponse) Validate() error {
	return dara.Validate(s)
}
