// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeviceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDeviceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDeviceInfoResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDeviceInfoResponseBody) *ModifyDeviceInfoResponse
	GetBody() *ModifyDeviceInfoResponseBody
}

type ModifyDeviceInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDeviceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyDeviceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDeviceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDeviceInfoResponse) GetBody() *ModifyDeviceInfoResponseBody {
	return s.Body
}

func (s *ModifyDeviceInfoResponse) SetHeaders(v map[string]*string) *ModifyDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *ModifyDeviceInfoResponse) SetStatusCode(v int32) *ModifyDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDeviceInfoResponse) SetBody(v *ModifyDeviceInfoResponseBody) *ModifyDeviceInfoResponse {
	s.Body = v
	return s
}

func (s *ModifyDeviceInfoResponse) Validate() error {
	return dara.Validate(s)
}
