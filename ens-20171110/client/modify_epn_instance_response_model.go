// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEpnInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyEpnInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyEpnInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyEpnInstanceResponseBody) *ModifyEpnInstanceResponse
	GetBody() *ModifyEpnInstanceResponseBody
}

type ModifyEpnInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEpnInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyEpnInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyEpnInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyEpnInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyEpnInstanceResponse) GetBody() *ModifyEpnInstanceResponseBody {
	return s.Body
}

func (s *ModifyEpnInstanceResponse) SetHeaders(v map[string]*string) *ModifyEpnInstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyEpnInstanceResponse) SetStatusCode(v int32) *ModifyEpnInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEpnInstanceResponse) SetBody(v *ModifyEpnInstanceResponseBody) *ModifyEpnInstanceResponse {
	s.Body = v
	return s
}

func (s *ModifyEpnInstanceResponse) Validate() error {
	return dara.Validate(s)
}
