// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeApplicationToGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeApplicationToGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeApplicationToGroupsResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeApplicationToGroupsResponseBody) *AuthorizeApplicationToGroupsResponse
	GetBody() *AuthorizeApplicationToGroupsResponseBody
}

type AuthorizeApplicationToGroupsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeApplicationToGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeApplicationToGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeApplicationToGroupsResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationToGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeApplicationToGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeApplicationToGroupsResponse) GetBody() *AuthorizeApplicationToGroupsResponseBody {
	return s.Body
}

func (s *AuthorizeApplicationToGroupsResponse) SetHeaders(v map[string]*string) *AuthorizeApplicationToGroupsResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeApplicationToGroupsResponse) SetStatusCode(v int32) *AuthorizeApplicationToGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeApplicationToGroupsResponse) SetBody(v *AuthorizeApplicationToGroupsResponseBody) *AuthorizeApplicationToGroupsResponse {
	s.Body = v
	return s
}

func (s *AuthorizeApplicationToGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
