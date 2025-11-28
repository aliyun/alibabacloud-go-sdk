// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSlowSQLRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadSlowSQLRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadSlowSQLRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DownloadSlowSQLRecordsResponseBody) *DownloadSlowSQLRecordsResponse
	GetBody() *DownloadSlowSQLRecordsResponseBody
}

type DownloadSlowSQLRecordsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadSlowSQLRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadSlowSQLRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadSlowSQLRecordsResponse) GoString() string {
	return s.String()
}

func (s *DownloadSlowSQLRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadSlowSQLRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadSlowSQLRecordsResponse) GetBody() *DownloadSlowSQLRecordsResponseBody {
	return s.Body
}

func (s *DownloadSlowSQLRecordsResponse) SetHeaders(v map[string]*string) *DownloadSlowSQLRecordsResponse {
	s.Headers = v
	return s
}

func (s *DownloadSlowSQLRecordsResponse) SetStatusCode(v int32) *DownloadSlowSQLRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadSlowSQLRecordsResponse) SetBody(v *DownloadSlowSQLRecordsResponseBody) *DownloadSlowSQLRecordsResponse {
	s.Body = v
	return s
}

func (s *DownloadSlowSQLRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
