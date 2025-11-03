// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceNetExpireTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceNetExpireTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceNetExpireTimeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceNetExpireTimeResponseBody) *ModifyDBInstanceNetExpireTimeResponse
	GetBody() *ModifyDBInstanceNetExpireTimeResponseBody
}

type ModifyDBInstanceNetExpireTimeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceNetExpireTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceNetExpireTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceNetExpireTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetExpireTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceNetExpireTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceNetExpireTimeResponse) GetBody() *ModifyDBInstanceNetExpireTimeResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceNetExpireTimeResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceNetExpireTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceNetExpireTimeResponse) SetStatusCode(v int32) *ModifyDBInstanceNetExpireTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceNetExpireTimeResponse) SetBody(v *ModifyDBInstanceNetExpireTimeResponseBody) *ModifyDBInstanceNetExpireTimeResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceNetExpireTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
