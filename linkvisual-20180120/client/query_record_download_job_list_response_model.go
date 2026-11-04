// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordDownloadJobListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRecordDownloadJobListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRecordDownloadJobListResponse
	GetStatusCode() *int32
	SetBody(v *QueryRecordDownloadJobListResponseBody) *QueryRecordDownloadJobListResponse
	GetBody() *QueryRecordDownloadJobListResponseBody
}

type QueryRecordDownloadJobListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordDownloadJobListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordDownloadJobListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordDownloadJobListResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordDownloadJobListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRecordDownloadJobListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRecordDownloadJobListResponse) GetBody() *QueryRecordDownloadJobListResponseBody {
	return s.Body
}

func (s *QueryRecordDownloadJobListResponse) SetHeaders(v map[string]*string) *QueryRecordDownloadJobListResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordDownloadJobListResponse) SetStatusCode(v int32) *QueryRecordDownloadJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordDownloadJobListResponse) SetBody(v *QueryRecordDownloadJobListResponseBody) *QueryRecordDownloadJobListResponse {
	s.Body = v
	return s
}

func (s *QueryRecordDownloadJobListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
