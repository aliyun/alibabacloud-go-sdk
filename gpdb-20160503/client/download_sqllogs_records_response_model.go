// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSQLLogsRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadSQLLogsRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadSQLLogsRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DownloadSQLLogsRecordsResponseBody) *DownloadSQLLogsRecordsResponse
	GetBody() *DownloadSQLLogsRecordsResponseBody
}

type DownloadSQLLogsRecordsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadSQLLogsRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadSQLLogsRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadSQLLogsRecordsResponse) GoString() string {
	return s.String()
}

func (s *DownloadSQLLogsRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadSQLLogsRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadSQLLogsRecordsResponse) GetBody() *DownloadSQLLogsRecordsResponseBody {
	return s.Body
}

func (s *DownloadSQLLogsRecordsResponse) SetHeaders(v map[string]*string) *DownloadSQLLogsRecordsResponse {
	s.Headers = v
	return s
}

func (s *DownloadSQLLogsRecordsResponse) SetStatusCode(v int32) *DownloadSQLLogsRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadSQLLogsRecordsResponse) SetBody(v *DownloadSQLLogsRecordsResponseBody) *DownloadSQLLogsRecordsResponse {
	s.Body = v
	return s
}

func (s *DownloadSQLLogsRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
