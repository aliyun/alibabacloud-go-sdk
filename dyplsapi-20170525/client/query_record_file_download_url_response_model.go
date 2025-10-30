// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordFileDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRecordFileDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRecordFileDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *QueryRecordFileDownloadUrlResponseBody) *QueryRecordFileDownloadUrlResponse
	GetBody() *QueryRecordFileDownloadUrlResponseBody
}

type QueryRecordFileDownloadUrlResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordFileDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordFileDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordFileDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordFileDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRecordFileDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRecordFileDownloadUrlResponse) GetBody() *QueryRecordFileDownloadUrlResponseBody {
	return s.Body
}

func (s *QueryRecordFileDownloadUrlResponse) SetHeaders(v map[string]*string) *QueryRecordFileDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordFileDownloadUrlResponse) SetStatusCode(v int32) *QueryRecordFileDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordFileDownloadUrlResponse) SetBody(v *QueryRecordFileDownloadUrlResponseBody) *QueryRecordFileDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *QueryRecordFileDownloadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
