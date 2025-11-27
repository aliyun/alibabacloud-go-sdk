// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceNetworkExpireTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceNetworkExpireTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceNetworkExpireTimeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceNetworkExpireTimeResponseBody) *ModifyDBInstanceNetworkExpireTimeResponse
	GetBody() *ModifyDBInstanceNetworkExpireTimeResponseBody
}

type ModifyDBInstanceNetworkExpireTimeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceNetworkExpireTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceNetworkExpireTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceNetworkExpireTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetworkExpireTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceNetworkExpireTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceNetworkExpireTimeResponse) GetBody() *ModifyDBInstanceNetworkExpireTimeResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceNetworkExpireTimeResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceNetworkExpireTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceNetworkExpireTimeResponse) SetStatusCode(v int32) *ModifyDBInstanceNetworkExpireTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceNetworkExpireTimeResponse) SetBody(v *ModifyDBInstanceNetworkExpireTimeResponseBody) *ModifyDBInstanceNetworkExpireTimeResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceNetworkExpireTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
