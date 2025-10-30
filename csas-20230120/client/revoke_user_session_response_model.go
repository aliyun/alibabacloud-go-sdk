// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeUserSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeUserSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeUserSessionResponse
	GetStatusCode() *int32
	SetBody(v *RevokeUserSessionResponseBody) *RevokeUserSessionResponse
	GetBody() *RevokeUserSessionResponseBody
}

type RevokeUserSessionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeUserSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeUserSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeUserSessionResponse) GoString() string {
	return s.String()
}

func (s *RevokeUserSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeUserSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeUserSessionResponse) GetBody() *RevokeUserSessionResponseBody {
	return s.Body
}

func (s *RevokeUserSessionResponse) SetHeaders(v map[string]*string) *RevokeUserSessionResponse {
	s.Headers = v
	return s
}

func (s *RevokeUserSessionResponse) SetStatusCode(v int32) *RevokeUserSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeUserSessionResponse) SetBody(v *RevokeUserSessionResponseBody) *RevokeUserSessionResponse {
	s.Body = v
	return s
}

func (s *RevokeUserSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
