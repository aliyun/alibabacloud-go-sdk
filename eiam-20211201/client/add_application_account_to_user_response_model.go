// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddApplicationAccountToUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddApplicationAccountToUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddApplicationAccountToUserResponse
	GetStatusCode() *int32
	SetBody(v *AddApplicationAccountToUserResponseBody) *AddApplicationAccountToUserResponse
	GetBody() *AddApplicationAccountToUserResponseBody
}

type AddApplicationAccountToUserResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddApplicationAccountToUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddApplicationAccountToUserResponse) String() string {
	return dara.Prettify(s)
}

func (s AddApplicationAccountToUserResponse) GoString() string {
	return s.String()
}

func (s *AddApplicationAccountToUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddApplicationAccountToUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddApplicationAccountToUserResponse) GetBody() *AddApplicationAccountToUserResponseBody {
	return s.Body
}

func (s *AddApplicationAccountToUserResponse) SetHeaders(v map[string]*string) *AddApplicationAccountToUserResponse {
	s.Headers = v
	return s
}

func (s *AddApplicationAccountToUserResponse) SetStatusCode(v int32) *AddApplicationAccountToUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AddApplicationAccountToUserResponse) SetBody(v *AddApplicationAccountToUserResponseBody) *AddApplicationAccountToUserResponse {
	s.Body = v
	return s
}

func (s *AddApplicationAccountToUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
