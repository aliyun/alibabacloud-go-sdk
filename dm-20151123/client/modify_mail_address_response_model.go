// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMailAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMailAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMailAddressResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMailAddressResponseBody) *ModifyMailAddressResponse
	GetBody() *ModifyMailAddressResponseBody
}

type ModifyMailAddressResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMailAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMailAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMailAddressResponse) GoString() string {
	return s.String()
}

func (s *ModifyMailAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMailAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMailAddressResponse) GetBody() *ModifyMailAddressResponseBody {
	return s.Body
}

func (s *ModifyMailAddressResponse) SetHeaders(v map[string]*string) *ModifyMailAddressResponse {
	s.Headers = v
	return s
}

func (s *ModifyMailAddressResponse) SetStatusCode(v int32) *ModifyMailAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMailAddressResponse) SetBody(v *ModifyMailAddressResponseBody) *ModifyMailAddressResponse {
	s.Body = v
	return s
}

func (s *ModifyMailAddressResponse) Validate() error {
	return dara.Validate(s)
}
