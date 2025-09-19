// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoginBaseConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLoginBaseConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLoginBaseConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLoginBaseConfigResponseBody) *ModifyLoginBaseConfigResponse
	GetBody() *ModifyLoginBaseConfigResponseBody
}

type ModifyLoginBaseConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLoginBaseConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLoginBaseConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoginBaseConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyLoginBaseConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLoginBaseConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLoginBaseConfigResponse) GetBody() *ModifyLoginBaseConfigResponseBody {
	return s.Body
}

func (s *ModifyLoginBaseConfigResponse) SetHeaders(v map[string]*string) *ModifyLoginBaseConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyLoginBaseConfigResponse) SetStatusCode(v int32) *ModifyLoginBaseConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLoginBaseConfigResponse) SetBody(v *ModifyLoginBaseConfigResponseBody) *ModifyLoginBaseConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyLoginBaseConfigResponse) Validate() error {
	return dara.Validate(s)
}
