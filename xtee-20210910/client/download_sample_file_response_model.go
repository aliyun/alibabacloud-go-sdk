// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSampleFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadSampleFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadSampleFileResponse
	GetStatusCode() *int32
	SetBody(v *DownloadSampleFileResponseBody) *DownloadSampleFileResponse
	GetBody() *DownloadSampleFileResponseBody
}

type DownloadSampleFileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadSampleFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadSampleFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadSampleFileResponse) GoString() string {
	return s.String()
}

func (s *DownloadSampleFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadSampleFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadSampleFileResponse) GetBody() *DownloadSampleFileResponseBody {
	return s.Body
}

func (s *DownloadSampleFileResponse) SetHeaders(v map[string]*string) *DownloadSampleFileResponse {
	s.Headers = v
	return s
}

func (s *DownloadSampleFileResponse) SetStatusCode(v int32) *DownloadSampleFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadSampleFileResponse) SetBody(v *DownloadSampleFileResponseBody) *DownloadSampleFileResponse {
	s.Body = v
	return s
}

func (s *DownloadSampleFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
