// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnnouncementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAnnouncementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAnnouncementResponse
	GetStatusCode() *int32
	SetBody(v *CreateAnnouncementResponseBody) *CreateAnnouncementResponse
	GetBody() *CreateAnnouncementResponseBody
}

type CreateAnnouncementResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAnnouncementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAnnouncementResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAnnouncementResponse) GoString() string {
	return s.String()
}

func (s *CreateAnnouncementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAnnouncementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAnnouncementResponse) GetBody() *CreateAnnouncementResponseBody {
	return s.Body
}

func (s *CreateAnnouncementResponse) SetHeaders(v map[string]*string) *CreateAnnouncementResponse {
	s.Headers = v
	return s
}

func (s *CreateAnnouncementResponse) SetStatusCode(v int32) *CreateAnnouncementResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAnnouncementResponse) SetBody(v *CreateAnnouncementResponseBody) *CreateAnnouncementResponse {
	s.Body = v
	return s
}

func (s *CreateAnnouncementResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
