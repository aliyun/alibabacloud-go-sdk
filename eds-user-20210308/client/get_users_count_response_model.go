// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUsersCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUsersCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUsersCountResponse
	GetStatusCode() *int32
	SetBody(v *GetUsersCountResponseBody) *GetUsersCountResponse
	GetBody() *GetUsersCountResponseBody
}

type GetUsersCountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUsersCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUsersCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUsersCountResponse) GoString() string {
	return s.String()
}

func (s *GetUsersCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUsersCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUsersCountResponse) GetBody() *GetUsersCountResponseBody {
	return s.Body
}

func (s *GetUsersCountResponse) SetHeaders(v map[string]*string) *GetUsersCountResponse {
	s.Headers = v
	return s
}

func (s *GetUsersCountResponse) SetStatusCode(v int32) *GetUsersCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUsersCountResponse) SetBody(v *GetUsersCountResponseBody) *GetUsersCountResponse {
	s.Body = v
	return s
}

func (s *GetUsersCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
