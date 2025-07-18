// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceNetworkTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceNetworkTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceNetworkTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceNetworkTypeResponseBody) *ModifyDBInstanceNetworkTypeResponse
	GetBody() *ModifyDBInstanceNetworkTypeResponseBody
}

type ModifyDBInstanceNetworkTypeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceNetworkTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceNetworkTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceNetworkTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetworkTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceNetworkTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceNetworkTypeResponse) GetBody() *ModifyDBInstanceNetworkTypeResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceNetworkTypeResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceNetworkTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceNetworkTypeResponse) SetStatusCode(v int32) *ModifyDBInstanceNetworkTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeResponse) SetBody(v *ModifyDBInstanceNetworkTypeResponseBody) *ModifyDBInstanceNetworkTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceNetworkTypeResponse) Validate() error {
	return dara.Validate(s)
}
