// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetOwnersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetOwnersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetOwnersResponse
	GetStatusCode() *int32
	SetBody(v *SetOwnersResponseBody) *SetOwnersResponse
	GetBody() *SetOwnersResponseBody
}

type SetOwnersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetOwnersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetOwnersResponse) String() string {
	return dara.Prettify(s)
}

func (s SetOwnersResponse) GoString() string {
	return s.String()
}

func (s *SetOwnersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetOwnersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetOwnersResponse) GetBody() *SetOwnersResponseBody {
	return s.Body
}

func (s *SetOwnersResponse) SetHeaders(v map[string]*string) *SetOwnersResponse {
	s.Headers = v
	return s
}

func (s *SetOwnersResponse) SetStatusCode(v int32) *SetOwnersResponse {
	s.StatusCode = &v
	return s
}

func (s *SetOwnersResponse) SetBody(v *SetOwnersResponseBody) *SetOwnersResponse {
	s.Body = v
	return s
}

func (s *SetOwnersResponse) Validate() error {
	return dara.Validate(s)
}
