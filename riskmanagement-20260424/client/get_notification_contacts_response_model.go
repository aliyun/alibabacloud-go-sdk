// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotificationContactsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNotificationContactsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNotificationContactsResponse
	GetStatusCode() *int32
	SetBody(v *GetNotificationContactsResponseBody) *GetNotificationContactsResponse
	GetBody() *GetNotificationContactsResponseBody
}

type GetNotificationContactsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNotificationContactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNotificationContactsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNotificationContactsResponse) GoString() string {
	return s.String()
}

func (s *GetNotificationContactsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNotificationContactsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNotificationContactsResponse) GetBody() *GetNotificationContactsResponseBody {
	return s.Body
}

func (s *GetNotificationContactsResponse) SetHeaders(v map[string]*string) *GetNotificationContactsResponse {
	s.Headers = v
	return s
}

func (s *GetNotificationContactsResponse) SetStatusCode(v int32) *GetNotificationContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNotificationContactsResponse) SetBody(v *GetNotificationContactsResponseBody) *GetNotificationContactsResponse {
	s.Body = v
	return s
}

func (s *GetNotificationContactsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
