// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignOutGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SignOutGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SignOutGroupResponse
	GetStatusCode() *int32
	SetBody(v *SignOutGroupResponseBody) *SignOutGroupResponse
	GetBody() *SignOutGroupResponseBody
}

type SignOutGroupResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SignOutGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SignOutGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s SignOutGroupResponse) GoString() string {
	return s.String()
}

func (s *SignOutGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SignOutGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SignOutGroupResponse) GetBody() *SignOutGroupResponseBody {
	return s.Body
}

func (s *SignOutGroupResponse) SetHeaders(v map[string]*string) *SignOutGroupResponse {
	s.Headers = v
	return s
}

func (s *SignOutGroupResponse) SetStatusCode(v int32) *SignOutGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *SignOutGroupResponse) SetBody(v *SignOutGroupResponseBody) *SignOutGroupResponse {
	s.Body = v
	return s
}

func (s *SignOutGroupResponse) Validate() error {
	return dara.Validate(s)
}
