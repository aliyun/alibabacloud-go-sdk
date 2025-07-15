// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDesktopNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDesktopNameResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDesktopNameResponseBody) *ModifyDesktopNameResponse
	GetBody() *ModifyDesktopNameResponseBody
}

type ModifyDesktopNameResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDesktopNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDesktopNameResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDesktopNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDesktopNameResponse) GetBody() *ModifyDesktopNameResponseBody {
	return s.Body
}

func (s *ModifyDesktopNameResponse) SetHeaders(v map[string]*string) *ModifyDesktopNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopNameResponse) SetStatusCode(v int32) *ModifyDesktopNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDesktopNameResponse) SetBody(v *ModifyDesktopNameResponseBody) *ModifyDesktopNameResponse {
	s.Body = v
	return s
}

func (s *ModifyDesktopNameResponse) Validate() error {
	return dara.Validate(s)
}
