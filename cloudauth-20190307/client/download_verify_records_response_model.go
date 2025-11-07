// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadVerifyRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadVerifyRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadVerifyRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DownloadVerifyRecordsResponseBody) *DownloadVerifyRecordsResponse
	GetBody() *DownloadVerifyRecordsResponseBody
}

type DownloadVerifyRecordsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadVerifyRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadVerifyRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadVerifyRecordsResponse) GoString() string {
	return s.String()
}

func (s *DownloadVerifyRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadVerifyRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadVerifyRecordsResponse) GetBody() *DownloadVerifyRecordsResponseBody {
	return s.Body
}

func (s *DownloadVerifyRecordsResponse) SetHeaders(v map[string]*string) *DownloadVerifyRecordsResponse {
	s.Headers = v
	return s
}

func (s *DownloadVerifyRecordsResponse) SetStatusCode(v int32) *DownloadVerifyRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadVerifyRecordsResponse) SetBody(v *DownloadVerifyRecordsResponseBody) *DownloadVerifyRecordsResponse {
	s.Body = v
	return s
}

func (s *DownloadVerifyRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
