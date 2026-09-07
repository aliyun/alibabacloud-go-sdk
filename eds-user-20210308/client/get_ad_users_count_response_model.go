// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdUsersCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAdUsersCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAdUsersCountResponse
	GetStatusCode() *int32
	SetBody(v *GetAdUsersCountResponseBody) *GetAdUsersCountResponse
	GetBody() *GetAdUsersCountResponseBody
}

type GetAdUsersCountResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAdUsersCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAdUsersCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAdUsersCountResponse) GoString() string {
	return s.String()
}

func (s *GetAdUsersCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAdUsersCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAdUsersCountResponse) GetBody() *GetAdUsersCountResponseBody {
	return s.Body
}

func (s *GetAdUsersCountResponse) SetHeaders(v map[string]*string) *GetAdUsersCountResponse {
	s.Headers = v
	return s
}

func (s *GetAdUsersCountResponse) SetStatusCode(v int32) *GetAdUsersCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAdUsersCountResponse) SetBody(v *GetAdUsersCountResponseBody) *GetAdUsersCountResponse {
	s.Body = v
	return s
}

func (s *GetAdUsersCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
