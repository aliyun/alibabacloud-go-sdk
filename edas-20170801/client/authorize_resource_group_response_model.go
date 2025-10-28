// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeResourceGroupResponseBody) *AuthorizeResourceGroupResponse
	GetBody() *AuthorizeResourceGroupResponseBody
}

type AuthorizeResourceGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeResourceGroupResponse) GetBody() *AuthorizeResourceGroupResponseBody {
	return s.Body
}

func (s *AuthorizeResourceGroupResponse) SetHeaders(v map[string]*string) *AuthorizeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeResourceGroupResponse) SetStatusCode(v int32) *AuthorizeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeResourceGroupResponse) SetBody(v *AuthorizeResourceGroupResponseBody) *AuthorizeResourceGroupResponse {
	s.Body = v
	return s
}

func (s *AuthorizeResourceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
