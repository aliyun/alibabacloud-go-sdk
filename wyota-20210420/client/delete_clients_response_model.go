// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteClientsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteClientsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteClientsResponseBody) *DeleteClientsResponse
	GetBody() *DeleteClientsResponseBody
}

type DeleteClientsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClientsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClientsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientsResponse) GoString() string {
	return s.String()
}

func (s *DeleteClientsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteClientsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteClientsResponse) GetBody() *DeleteClientsResponseBody {
	return s.Body
}

func (s *DeleteClientsResponse) SetHeaders(v map[string]*string) *DeleteClientsResponse {
	s.Headers = v
	return s
}

func (s *DeleteClientsResponse) SetStatusCode(v int32) *DeleteClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClientsResponse) SetBody(v *DeleteClientsResponseBody) *DeleteClientsResponse {
	s.Body = v
	return s
}

func (s *DeleteClientsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
