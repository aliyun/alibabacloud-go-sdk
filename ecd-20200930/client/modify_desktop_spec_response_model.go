// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDesktopSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDesktopSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDesktopSpecResponseBody) *ModifyDesktopSpecResponse
	GetBody() *ModifyDesktopSpecResponseBody
}

type ModifyDesktopSpecResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDesktopSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDesktopSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDesktopSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDesktopSpecResponse) GetBody() *ModifyDesktopSpecResponseBody {
	return s.Body
}

func (s *ModifyDesktopSpecResponse) SetHeaders(v map[string]*string) *ModifyDesktopSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopSpecResponse) SetStatusCode(v int32) *ModifyDesktopSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDesktopSpecResponse) SetBody(v *ModifyDesktopSpecResponseBody) *ModifyDesktopSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyDesktopSpecResponse) Validate() error {
	return dara.Validate(s)
}
