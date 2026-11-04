// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRecordDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRecordDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *QueryRecordDownloadUrlResponseBody) *QueryRecordDownloadUrlResponse
	GetBody() *QueryRecordDownloadUrlResponseBody
}

type QueryRecordDownloadUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRecordDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRecordDownloadUrlResponse) GetBody() *QueryRecordDownloadUrlResponseBody {
	return s.Body
}

func (s *QueryRecordDownloadUrlResponse) SetHeaders(v map[string]*string) *QueryRecordDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordDownloadUrlResponse) SetStatusCode(v int32) *QueryRecordDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordDownloadUrlResponse) SetBody(v *QueryRecordDownloadUrlResponseBody) *QueryRecordDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *QueryRecordDownloadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
