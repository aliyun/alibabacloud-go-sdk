// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadDataTrackResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadDataTrackResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadDataTrackResultResponse
	GetStatusCode() *int32
	SetBody(v *DownloadDataTrackResultResponseBody) *DownloadDataTrackResultResponse
	GetBody() *DownloadDataTrackResultResponseBody
}

type DownloadDataTrackResultResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadDataTrackResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadDataTrackResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadDataTrackResultResponse) GoString() string {
	return s.String()
}

func (s *DownloadDataTrackResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadDataTrackResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadDataTrackResultResponse) GetBody() *DownloadDataTrackResultResponseBody {
	return s.Body
}

func (s *DownloadDataTrackResultResponse) SetHeaders(v map[string]*string) *DownloadDataTrackResultResponse {
	s.Headers = v
	return s
}

func (s *DownloadDataTrackResultResponse) SetStatusCode(v int32) *DownloadDataTrackResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadDataTrackResultResponse) SetBody(v *DownloadDataTrackResultResponseBody) *DownloadDataTrackResultResponse {
	s.Body = v
	return s
}

func (s *DownloadDataTrackResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
