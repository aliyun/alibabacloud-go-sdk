// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClientUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClientUserResponse
	GetStatusCode() *int32
	SetBody(v *GetClientUserResponseBody) *GetClientUserResponse
	GetBody() *GetClientUserResponseBody
}

type GetClientUserResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClientUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClientUserResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClientUserResponse) GoString() string {
	return s.String()
}

func (s *GetClientUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClientUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClientUserResponse) GetBody() *GetClientUserResponseBody {
	return s.Body
}

func (s *GetClientUserResponse) SetHeaders(v map[string]*string) *GetClientUserResponse {
	s.Headers = v
	return s
}

func (s *GetClientUserResponse) SetStatusCode(v int32) *GetClientUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClientUserResponse) SetBody(v *GetClientUserResponseBody) *GetClientUserResponse {
	s.Body = v
	return s
}

func (s *GetClientUserResponse) Validate() error {
	return dara.Validate(s)
}
