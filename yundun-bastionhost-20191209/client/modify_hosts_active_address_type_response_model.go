// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostsActiveAddressTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHostsActiveAddressTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHostsActiveAddressTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHostsActiveAddressTypeResponseBody) *ModifyHostsActiveAddressTypeResponse
	GetBody() *ModifyHostsActiveAddressTypeResponseBody
}

type ModifyHostsActiveAddressTypeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHostsActiveAddressTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHostsActiveAddressTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostsActiveAddressTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyHostsActiveAddressTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHostsActiveAddressTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHostsActiveAddressTypeResponse) GetBody() *ModifyHostsActiveAddressTypeResponseBody {
	return s.Body
}

func (s *ModifyHostsActiveAddressTypeResponse) SetHeaders(v map[string]*string) *ModifyHostsActiveAddressTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyHostsActiveAddressTypeResponse) SetStatusCode(v int32) *ModifyHostsActiveAddressTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeResponse) SetBody(v *ModifyHostsActiveAddressTypeResponseBody) *ModifyHostsActiveAddressTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyHostsActiveAddressTypeResponse) Validate() error {
	return dara.Validate(s)
}
