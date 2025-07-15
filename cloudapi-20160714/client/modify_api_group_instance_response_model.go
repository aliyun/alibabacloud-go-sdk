// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiGroupInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApiGroupInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApiGroupInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApiGroupInstanceResponseBody) *ModifyApiGroupInstanceResponse
	GetBody() *ModifyApiGroupInstanceResponseBody
}

type ModifyApiGroupInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApiGroupInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApiGroupInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiGroupInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApiGroupInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApiGroupInstanceResponse) GetBody() *ModifyApiGroupInstanceResponseBody {
	return s.Body
}

func (s *ModifyApiGroupInstanceResponse) SetHeaders(v map[string]*string) *ModifyApiGroupInstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyApiGroupInstanceResponse) SetStatusCode(v int32) *ModifyApiGroupInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApiGroupInstanceResponse) SetBody(v *ModifyApiGroupInstanceResponseBody) *ModifyApiGroupInstanceResponse {
	s.Body = v
	return s
}

func (s *ModifyApiGroupInstanceResponse) Validate() error {
	return dara.Validate(s)
}
