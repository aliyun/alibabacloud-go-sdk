// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActiveAnnouncementsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListActiveAnnouncementsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListActiveAnnouncementsResponse
	GetStatusCode() *int32
	SetBody(v *ListActiveAnnouncementsResponseBody) *ListActiveAnnouncementsResponse
	GetBody() *ListActiveAnnouncementsResponseBody
}

type ListActiveAnnouncementsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListActiveAnnouncementsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListActiveAnnouncementsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListActiveAnnouncementsResponse) GoString() string {
	return s.String()
}

func (s *ListActiveAnnouncementsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListActiveAnnouncementsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListActiveAnnouncementsResponse) GetBody() *ListActiveAnnouncementsResponseBody {
	return s.Body
}

func (s *ListActiveAnnouncementsResponse) SetHeaders(v map[string]*string) *ListActiveAnnouncementsResponse {
	s.Headers = v
	return s
}

func (s *ListActiveAnnouncementsResponse) SetStatusCode(v int32) *ListActiveAnnouncementsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListActiveAnnouncementsResponse) SetBody(v *ListActiveAnnouncementsResponseBody) *ListActiveAnnouncementsResponse {
	s.Body = v
	return s
}

func (s *ListActiveAnnouncementsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
