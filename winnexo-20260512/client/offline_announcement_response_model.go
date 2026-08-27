// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineAnnouncementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OfflineAnnouncementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OfflineAnnouncementResponse
	GetStatusCode() *int32
	SetBody(v *OfflineAnnouncementResponseBody) *OfflineAnnouncementResponse
	GetBody() *OfflineAnnouncementResponseBody
}

type OfflineAnnouncementResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineAnnouncementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineAnnouncementResponse) String() string {
	return dara.Prettify(s)
}

func (s OfflineAnnouncementResponse) GoString() string {
	return s.String()
}

func (s *OfflineAnnouncementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OfflineAnnouncementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OfflineAnnouncementResponse) GetBody() *OfflineAnnouncementResponseBody {
	return s.Body
}

func (s *OfflineAnnouncementResponse) SetHeaders(v map[string]*string) *OfflineAnnouncementResponse {
	s.Headers = v
	return s
}

func (s *OfflineAnnouncementResponse) SetStatusCode(v int32) *OfflineAnnouncementResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineAnnouncementResponse) SetBody(v *OfflineAnnouncementResponseBody) *OfflineAnnouncementResponse {
	s.Body = v
	return s
}

func (s *OfflineAnnouncementResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
