// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteClientUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteClientUserResponse
	GetStatusCode() *int32
	SetBody(v *DeleteClientUserResponseBody) *DeleteClientUserResponse
	GetBody() *DeleteClientUserResponseBody
}

type DeleteClientUserResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClientUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClientUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteClientUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteClientUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteClientUserResponse) GetBody() *DeleteClientUserResponseBody {
	return s.Body
}

func (s *DeleteClientUserResponse) SetHeaders(v map[string]*string) *DeleteClientUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteClientUserResponse) SetStatusCode(v int32) *DeleteClientUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClientUserResponse) SetBody(v *DeleteClientUserResponseBody) *DeleteClientUserResponse {
	s.Body = v
	return s
}

func (s *DeleteClientUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
