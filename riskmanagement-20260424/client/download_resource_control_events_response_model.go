// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadResourceControlEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadResourceControlEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadResourceControlEventsResponse
	GetStatusCode() *int32
	SetBody(v *DownloadResourceControlEventsResponseBody) *DownloadResourceControlEventsResponse
	GetBody() *DownloadResourceControlEventsResponseBody
}

type DownloadResourceControlEventsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadResourceControlEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadResourceControlEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadResourceControlEventsResponse) GoString() string {
	return s.String()
}

func (s *DownloadResourceControlEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadResourceControlEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadResourceControlEventsResponse) GetBody() *DownloadResourceControlEventsResponseBody {
	return s.Body
}

func (s *DownloadResourceControlEventsResponse) SetHeaders(v map[string]*string) *DownloadResourceControlEventsResponse {
	s.Headers = v
	return s
}

func (s *DownloadResourceControlEventsResponse) SetStatusCode(v int32) *DownloadResourceControlEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadResourceControlEventsResponse) SetBody(v *DownloadResourceControlEventsResponseBody) *DownloadResourceControlEventsResponse {
	s.Body = v
	return s
}

func (s *DownloadResourceControlEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
