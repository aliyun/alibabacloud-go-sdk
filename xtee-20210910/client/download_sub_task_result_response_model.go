// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSubTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadSubTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadSubTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *DownloadSubTaskResultResponseBody) *DownloadSubTaskResultResponse
	GetBody() *DownloadSubTaskResultResponseBody
}

type DownloadSubTaskResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadSubTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadSubTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadSubTaskResultResponse) GoString() string {
	return s.String()
}

func (s *DownloadSubTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadSubTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadSubTaskResultResponse) GetBody() *DownloadSubTaskResultResponseBody {
	return s.Body
}

func (s *DownloadSubTaskResultResponse) SetHeaders(v map[string]*string) *DownloadSubTaskResultResponse {
	s.Headers = v
	return s
}

func (s *DownloadSubTaskResultResponse) SetStatusCode(v int32) *DownloadSubTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadSubTaskResultResponse) SetBody(v *DownloadSubTaskResultResponseBody) *DownloadSubTaskResultResponse {
	s.Body = v
	return s
}

func (s *DownloadSubTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
