// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAccountPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAccountPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAccountPasswordResponseBody) *ModifyAccountPasswordResponse
	GetBody() *ModifyAccountPasswordResponseBody
}

type ModifyAccountPasswordResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAccountPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAccountPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAccountPasswordResponse) GetBody() *ModifyAccountPasswordResponseBody {
	return s.Body
}

func (s *ModifyAccountPasswordResponse) SetHeaders(v map[string]*string) *ModifyAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountPasswordResponse) SetStatusCode(v int32) *ModifyAccountPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountPasswordResponse) SetBody(v *ModifyAccountPasswordResponseBody) *ModifyAccountPasswordResponse {
	s.Body = v
	return s
}

func (s *ModifyAccountPasswordResponse) Validate() error {
	return dara.Validate(s)
}
