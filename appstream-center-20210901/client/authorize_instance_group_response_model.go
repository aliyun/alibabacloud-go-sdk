// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeInstanceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeInstanceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeInstanceGroupResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeInstanceGroupResponseBody) *AuthorizeInstanceGroupResponse
	GetBody() *AuthorizeInstanceGroupResponseBody
}

type AuthorizeInstanceGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeInstanceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeInstanceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeInstanceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeInstanceGroupResponse) GetBody() *AuthorizeInstanceGroupResponseBody {
	return s.Body
}

func (s *AuthorizeInstanceGroupResponse) SetHeaders(v map[string]*string) *AuthorizeInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeInstanceGroupResponse) SetStatusCode(v int32) *AuthorizeInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeInstanceGroupResponse) SetBody(v *AuthorizeInstanceGroupResponseBody) *AuthorizeInstanceGroupResponse {
	s.Body = v
	return s
}

func (s *AuthorizeInstanceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
