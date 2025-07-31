// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUsersResponse
	GetStatusCode() *int32
	SetBody(v *GetUsersResponseBody) *GetUsersResponse
	GetBody() *GetUsersResponseBody
}

type GetUsersResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUsersResponse) GoString() string {
	return s.String()
}

func (s *GetUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUsersResponse) GetBody() *GetUsersResponseBody {
	return s.Body
}

func (s *GetUsersResponse) SetHeaders(v map[string]*string) *GetUsersResponse {
	s.Headers = v
	return s
}

func (s *GetUsersResponse) SetStatusCode(v int32) *GetUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUsersResponse) SetBody(v *GetUsersResponseBody) *GetUsersResponse {
	s.Body = v
	return s
}

func (s *GetUsersResponse) Validate() error {
	return dara.Validate(s)
}
