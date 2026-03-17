// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKickOutClientsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *KickOutClientsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *KickOutClientsResponse
	GetStatusCode() *int32
	SetBody(v *KickOutClientsResponseBody) *KickOutClientsResponse
	GetBody() *KickOutClientsResponseBody
}

type KickOutClientsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KickOutClientsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KickOutClientsResponse) String() string {
	return dara.Prettify(s)
}

func (s KickOutClientsResponse) GoString() string {
	return s.String()
}

func (s *KickOutClientsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *KickOutClientsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *KickOutClientsResponse) GetBody() *KickOutClientsResponseBody {
	return s.Body
}

func (s *KickOutClientsResponse) SetHeaders(v map[string]*string) *KickOutClientsResponse {
	s.Headers = v
	return s
}

func (s *KickOutClientsResponse) SetStatusCode(v int32) *KickOutClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *KickOutClientsResponse) SetBody(v *KickOutClientsResponseBody) *KickOutClientsResponse {
	s.Body = v
	return s
}

func (s *KickOutClientsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
