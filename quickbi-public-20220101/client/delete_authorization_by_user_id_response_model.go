// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuthorizationByUserIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAuthorizationByUserIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAuthorizationByUserIdResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAuthorizationByUserIdResponseBody) *DeleteAuthorizationByUserIdResponse
	GetBody() *DeleteAuthorizationByUserIdResponseBody
}

type DeleteAuthorizationByUserIdResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAuthorizationByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAuthorizationByUserIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuthorizationByUserIdResponse) GoString() string {
	return s.String()
}

func (s *DeleteAuthorizationByUserIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAuthorizationByUserIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAuthorizationByUserIdResponse) GetBody() *DeleteAuthorizationByUserIdResponseBody {
	return s.Body
}

func (s *DeleteAuthorizationByUserIdResponse) SetHeaders(v map[string]*string) *DeleteAuthorizationByUserIdResponse {
	s.Headers = v
	return s
}

func (s *DeleteAuthorizationByUserIdResponse) SetStatusCode(v int32) *DeleteAuthorizationByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAuthorizationByUserIdResponse) SetBody(v *DeleteAuthorizationByUserIdResponseBody) *DeleteAuthorizationByUserIdResponse {
	s.Body = v
	return s
}

func (s *DeleteAuthorizationByUserIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
