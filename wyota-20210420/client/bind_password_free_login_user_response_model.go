// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindPasswordFreeLoginUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindPasswordFreeLoginUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindPasswordFreeLoginUserResponse
	GetStatusCode() *int32
	SetBody(v *BindPasswordFreeLoginUserResponseBody) *BindPasswordFreeLoginUserResponse
	GetBody() *BindPasswordFreeLoginUserResponseBody
}

type BindPasswordFreeLoginUserResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindPasswordFreeLoginUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindPasswordFreeLoginUserResponse) String() string {
	return dara.Prettify(s)
}

func (s BindPasswordFreeLoginUserResponse) GoString() string {
	return s.String()
}

func (s *BindPasswordFreeLoginUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindPasswordFreeLoginUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindPasswordFreeLoginUserResponse) GetBody() *BindPasswordFreeLoginUserResponseBody {
	return s.Body
}

func (s *BindPasswordFreeLoginUserResponse) SetHeaders(v map[string]*string) *BindPasswordFreeLoginUserResponse {
	s.Headers = v
	return s
}

func (s *BindPasswordFreeLoginUserResponse) SetStatusCode(v int32) *BindPasswordFreeLoginUserResponse {
	s.StatusCode = &v
	return s
}

func (s *BindPasswordFreeLoginUserResponse) SetBody(v *BindPasswordFreeLoginUserResponseBody) *BindPasswordFreeLoginUserResponse {
	s.Body = v
	return s
}

func (s *BindPasswordFreeLoginUserResponse) Validate() error {
	return dara.Validate(s)
}
