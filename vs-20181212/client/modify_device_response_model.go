// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDeviceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDeviceResponseBody) *ModifyDeviceResponse
	GetBody() *ModifyDeviceResponseBody
}

type ModifyDeviceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeviceResponse) GoString() string {
	return s.String()
}

func (s *ModifyDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDeviceResponse) GetBody() *ModifyDeviceResponseBody {
	return s.Body
}

func (s *ModifyDeviceResponse) SetHeaders(v map[string]*string) *ModifyDeviceResponse {
	s.Headers = v
	return s
}

func (s *ModifyDeviceResponse) SetStatusCode(v int32) *ModifyDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDeviceResponse) SetBody(v *ModifyDeviceResponseBody) *ModifyDeviceResponse {
	s.Body = v
	return s
}

func (s *ModifyDeviceResponse) Validate() error {
	return dara.Validate(s)
}
