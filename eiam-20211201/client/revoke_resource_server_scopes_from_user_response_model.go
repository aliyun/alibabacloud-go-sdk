// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeResourceServerScopesFromUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeResourceServerScopesFromUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeResourceServerScopesFromUserResponse
	GetStatusCode() *int32
	SetBody(v *RevokeResourceServerScopesFromUserResponseBody) *RevokeResourceServerScopesFromUserResponse
	GetBody() *RevokeResourceServerScopesFromUserResponseBody
}

type RevokeResourceServerScopesFromUserResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeResourceServerScopesFromUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeResourceServerScopesFromUserResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeResourceServerScopesFromUserResponse) GoString() string {
	return s.String()
}

func (s *RevokeResourceServerScopesFromUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeResourceServerScopesFromUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeResourceServerScopesFromUserResponse) GetBody() *RevokeResourceServerScopesFromUserResponseBody {
	return s.Body
}

func (s *RevokeResourceServerScopesFromUserResponse) SetHeaders(v map[string]*string) *RevokeResourceServerScopesFromUserResponse {
	s.Headers = v
	return s
}

func (s *RevokeResourceServerScopesFromUserResponse) SetStatusCode(v int32) *RevokeResourceServerScopesFromUserResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeResourceServerScopesFromUserResponse) SetBody(v *RevokeResourceServerScopesFromUserResponseBody) *RevokeResourceServerScopesFromUserResponse {
	s.Body = v
	return s
}

func (s *RevokeResourceServerScopesFromUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
