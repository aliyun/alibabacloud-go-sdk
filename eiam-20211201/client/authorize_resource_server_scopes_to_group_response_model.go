// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeResourceServerScopesToGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeResourceServerScopesToGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeResourceServerScopesToGroupResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeResourceServerScopesToGroupResponseBody) *AuthorizeResourceServerScopesToGroupResponse
	GetBody() *AuthorizeResourceServerScopesToGroupResponseBody
}

type AuthorizeResourceServerScopesToGroupResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeResourceServerScopesToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeResourceServerScopesToGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeResourceServerScopesToGroupResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceServerScopesToGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeResourceServerScopesToGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeResourceServerScopesToGroupResponse) GetBody() *AuthorizeResourceServerScopesToGroupResponseBody {
	return s.Body
}

func (s *AuthorizeResourceServerScopesToGroupResponse) SetHeaders(v map[string]*string) *AuthorizeResourceServerScopesToGroupResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeResourceServerScopesToGroupResponse) SetStatusCode(v int32) *AuthorizeResourceServerScopesToGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeResourceServerScopesToGroupResponse) SetBody(v *AuthorizeResourceServerScopesToGroupResponseBody) *AuthorizeResourceServerScopesToGroupResponse {
	s.Body = v
	return s
}

func (s *AuthorizeResourceServerScopesToGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
