// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDevicesSecureNetworkTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDevicesSecureNetworkTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDevicesSecureNetworkTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDevicesSecureNetworkTypeResponseBody) *ModifyDevicesSecureNetworkTypeResponse
	GetBody() *ModifyDevicesSecureNetworkTypeResponseBody
}

type ModifyDevicesSecureNetworkTypeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDevicesSecureNetworkTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDevicesSecureNetworkTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDevicesSecureNetworkTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDevicesSecureNetworkTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDevicesSecureNetworkTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDevicesSecureNetworkTypeResponse) GetBody() *ModifyDevicesSecureNetworkTypeResponseBody {
	return s.Body
}

func (s *ModifyDevicesSecureNetworkTypeResponse) SetHeaders(v map[string]*string) *ModifyDevicesSecureNetworkTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDevicesSecureNetworkTypeResponse) SetStatusCode(v int32) *ModifyDevicesSecureNetworkTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDevicesSecureNetworkTypeResponse) SetBody(v *ModifyDevicesSecureNetworkTypeResponseBody) *ModifyDevicesSecureNetworkTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyDevicesSecureNetworkTypeResponse) Validate() error {
	return dara.Validate(s)
}
