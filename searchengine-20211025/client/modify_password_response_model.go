// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPasswordResponseBody) *ModifyPasswordResponse
	GetBody() *ModifyPasswordResponseBody
}

type ModifyPasswordResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPasswordResponse) GetBody() *ModifyPasswordResponseBody {
	return s.Body
}

func (s *ModifyPasswordResponse) SetHeaders(v map[string]*string) *ModifyPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyPasswordResponse) SetStatusCode(v int32) *ModifyPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPasswordResponse) SetBody(v *ModifyPasswordResponseBody) *ModifyPasswordResponse {
	s.Body = v
	return s
}

func (s *ModifyPasswordResponse) Validate() error {
	return dara.Validate(s)
}
