// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadVerifyRecordIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadVerifyRecordIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadVerifyRecordIntlResponse
	GetStatusCode() *int32
	SetBody(v *DownloadVerifyRecordIntlResponseBody) *DownloadVerifyRecordIntlResponse
	GetBody() *DownloadVerifyRecordIntlResponseBody
}

type DownloadVerifyRecordIntlResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadVerifyRecordIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadVerifyRecordIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadVerifyRecordIntlResponse) GoString() string {
	return s.String()
}

func (s *DownloadVerifyRecordIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadVerifyRecordIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadVerifyRecordIntlResponse) GetBody() *DownloadVerifyRecordIntlResponseBody {
	return s.Body
}

func (s *DownloadVerifyRecordIntlResponse) SetHeaders(v map[string]*string) *DownloadVerifyRecordIntlResponse {
	s.Headers = v
	return s
}

func (s *DownloadVerifyRecordIntlResponse) SetStatusCode(v int32) *DownloadVerifyRecordIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadVerifyRecordIntlResponse) SetBody(v *DownloadVerifyRecordIntlResponseBody) *DownloadVerifyRecordIntlResponse {
	s.Body = v
	return s
}

func (s *DownloadVerifyRecordIntlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
