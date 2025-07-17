// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDataTrackResultDownloadStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDataTrackResultDownloadStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDataTrackResultDownloadStatusResponse
	GetStatusCode() *int32
	SetBody(v *QueryDataTrackResultDownloadStatusResponseBody) *QueryDataTrackResultDownloadStatusResponse
	GetBody() *QueryDataTrackResultDownloadStatusResponseBody
}

type QueryDataTrackResultDownloadStatusResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDataTrackResultDownloadStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDataTrackResultDownloadStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDataTrackResultDownloadStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryDataTrackResultDownloadStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDataTrackResultDownloadStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDataTrackResultDownloadStatusResponse) GetBody() *QueryDataTrackResultDownloadStatusResponseBody {
	return s.Body
}

func (s *QueryDataTrackResultDownloadStatusResponse) SetHeaders(v map[string]*string) *QueryDataTrackResultDownloadStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryDataTrackResultDownloadStatusResponse) SetStatusCode(v int32) *QueryDataTrackResultDownloadStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDataTrackResultDownloadStatusResponse) SetBody(v *QueryDataTrackResultDownloadStatusResponseBody) *QueryDataTrackResultDownloadStatusResponse {
	s.Body = v
	return s
}

func (s *QueryDataTrackResultDownloadStatusResponse) Validate() error {
	return dara.Validate(s)
}
