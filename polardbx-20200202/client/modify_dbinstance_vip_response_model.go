// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceVipResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceVipResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceVipResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceVipResponseBody) *ModifyDBInstanceVipResponse
	GetBody() *ModifyDBInstanceVipResponseBody
}

type ModifyDBInstanceVipResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceVipResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceVipResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceVipResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceVipResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceVipResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceVipResponse) GetBody() *ModifyDBInstanceVipResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceVipResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceVipResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceVipResponse) SetStatusCode(v int32) *ModifyDBInstanceVipResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceVipResponse) SetBody(v *ModifyDBInstanceVipResponseBody) *ModifyDBInstanceVipResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceVipResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
