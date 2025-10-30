// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSampleFileDownloadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SampleFileDownloadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SampleFileDownloadResponse
	GetStatusCode() *int32
	SetBody(v *SampleFileDownloadResponseBody) *SampleFileDownloadResponse
	GetBody() *SampleFileDownloadResponseBody
}

type SampleFileDownloadResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SampleFileDownloadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SampleFileDownloadResponse) String() string {
	return dara.Prettify(s)
}

func (s SampleFileDownloadResponse) GoString() string {
	return s.String()
}

func (s *SampleFileDownloadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SampleFileDownloadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SampleFileDownloadResponse) GetBody() *SampleFileDownloadResponseBody {
	return s.Body
}

func (s *SampleFileDownloadResponse) SetHeaders(v map[string]*string) *SampleFileDownloadResponse {
	s.Headers = v
	return s
}

func (s *SampleFileDownloadResponse) SetStatusCode(v int32) *SampleFileDownloadResponse {
	s.StatusCode = &v
	return s
}

func (s *SampleFileDownloadResponse) SetBody(v *SampleFileDownloadResponseBody) *SampleFileDownloadResponse {
	s.Body = v
	return s
}

func (s *SampleFileDownloadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
