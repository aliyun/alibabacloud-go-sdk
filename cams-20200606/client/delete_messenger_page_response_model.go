// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessengerPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMessengerPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMessengerPageResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMessengerPageResponseBody) *DeleteMessengerPageResponse
	GetBody() *DeleteMessengerPageResponseBody
}

type DeleteMessengerPageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMessengerPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMessengerPageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessengerPageResponse) GoString() string {
	return s.String()
}

func (s *DeleteMessengerPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMessengerPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMessengerPageResponse) GetBody() *DeleteMessengerPageResponseBody {
	return s.Body
}

func (s *DeleteMessengerPageResponse) SetHeaders(v map[string]*string) *DeleteMessengerPageResponse {
	s.Headers = v
	return s
}

func (s *DeleteMessengerPageResponse) SetStatusCode(v int32) *DeleteMessengerPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMessengerPageResponse) SetBody(v *DeleteMessengerPageResponseBody) *DeleteMessengerPageResponse {
	s.Body = v
	return s
}

func (s *DeleteMessengerPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
