// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserPoolResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserPoolResponseBody) *DeleteUserPoolResponse
	GetBody() *DeleteUserPoolResponseBody
}

type DeleteUserPoolResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserPoolResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserPoolResponse) GetBody() *DeleteUserPoolResponseBody {
	return s.Body
}

func (s *DeleteUserPoolResponse) SetHeaders(v map[string]*string) *DeleteUserPoolResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserPoolResponse) SetStatusCode(v int32) *DeleteUserPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserPoolResponse) SetBody(v *DeleteUserPoolResponseBody) *DeleteUserPoolResponse {
	s.Body = v
	return s
}

func (s *DeleteUserPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
