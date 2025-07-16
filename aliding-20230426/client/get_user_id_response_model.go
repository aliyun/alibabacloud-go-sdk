// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserIdResponse
	GetStatusCode() *int32
	SetBody(v *GetUserIdResponseBody) *GetUserIdResponse
	GetBody() *GetUserIdResponseBody
}

type GetUserIdResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdResponse) GoString() string {
	return s.String()
}

func (s *GetUserIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserIdResponse) GetBody() *GetUserIdResponseBody {
	return s.Body
}

func (s *GetUserIdResponse) SetHeaders(v map[string]*string) *GetUserIdResponse {
	s.Headers = v
	return s
}

func (s *GetUserIdResponse) SetStatusCode(v int32) *GetUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserIdResponse) SetBody(v *GetUserIdResponseBody) *GetUserIdResponse {
	s.Body = v
	return s
}

func (s *GetUserIdResponse) Validate() error {
	return dara.Validate(s)
}
