// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByUsernameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserIdByUsernameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserIdByUsernameResponse
	GetStatusCode() *int32
	SetBody(v *GetUserIdByUsernameResponseBody) *GetUserIdByUsernameResponse
	GetBody() *GetUserIdByUsernameResponseBody
}

type GetUserIdByUsernameResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserIdByUsernameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserIdByUsernameResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByUsernameResponse) GoString() string {
	return s.String()
}

func (s *GetUserIdByUsernameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserIdByUsernameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserIdByUsernameResponse) GetBody() *GetUserIdByUsernameResponseBody {
	return s.Body
}

func (s *GetUserIdByUsernameResponse) SetHeaders(v map[string]*string) *GetUserIdByUsernameResponse {
	s.Headers = v
	return s
}

func (s *GetUserIdByUsernameResponse) SetStatusCode(v int32) *GetUserIdByUsernameResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserIdByUsernameResponse) SetBody(v *GetUserIdByUsernameResponseBody) *GetUserIdByUsernameResponse {
	s.Body = v
	return s
}

func (s *GetUserIdByUsernameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
