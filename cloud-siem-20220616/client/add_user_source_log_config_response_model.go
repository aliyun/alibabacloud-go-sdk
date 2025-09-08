// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserSourceLogConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUserSourceLogConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUserSourceLogConfigResponse
	GetStatusCode() *int32
	SetBody(v *AddUserSourceLogConfigResponseBody) *AddUserSourceLogConfigResponse
	GetBody() *AddUserSourceLogConfigResponseBody
}

type AddUserSourceLogConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserSourceLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserSourceLogConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUserSourceLogConfigResponse) GoString() string {
	return s.String()
}

func (s *AddUserSourceLogConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUserSourceLogConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUserSourceLogConfigResponse) GetBody() *AddUserSourceLogConfigResponseBody {
	return s.Body
}

func (s *AddUserSourceLogConfigResponse) SetHeaders(v map[string]*string) *AddUserSourceLogConfigResponse {
	s.Headers = v
	return s
}

func (s *AddUserSourceLogConfigResponse) SetStatusCode(v int32) *AddUserSourceLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserSourceLogConfigResponse) SetBody(v *AddUserSourceLogConfigResponseBody) *AddUserSourceLogConfigResponse {
	s.Body = v
	return s
}

func (s *AddUserSourceLogConfigResponse) Validate() error {
	return dara.Validate(s)
}
