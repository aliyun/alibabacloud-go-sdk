// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExtendfilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateExtendfilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateExtendfilesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateExtendfilesResponseBody) *UpdateExtendfilesResponse
	GetBody() *UpdateExtendfilesResponseBody
}

type UpdateExtendfilesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExtendfilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExtendfilesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateExtendfilesResponse) GoString() string {
	return s.String()
}

func (s *UpdateExtendfilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateExtendfilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateExtendfilesResponse) GetBody() *UpdateExtendfilesResponseBody {
	return s.Body
}

func (s *UpdateExtendfilesResponse) SetHeaders(v map[string]*string) *UpdateExtendfilesResponse {
	s.Headers = v
	return s
}

func (s *UpdateExtendfilesResponse) SetStatusCode(v int32) *UpdateExtendfilesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExtendfilesResponse) SetBody(v *UpdateExtendfilesResponseBody) *UpdateExtendfilesResponse {
	s.Body = v
	return s
}

func (s *UpdateExtendfilesResponse) Validate() error {
	return dara.Validate(s)
}
