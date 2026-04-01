// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceResponseBody) *ModifyDBInstanceResponse
	GetBody() *ModifyDBInstanceResponseBody
}

type ModifyDBInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceResponse) GetBody() *ModifyDBInstanceResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceResponse) SetStatusCode(v int32) *ModifyDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceResponse) SetBody(v *ModifyDBInstanceResponseBody) *ModifyDBInstanceResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
