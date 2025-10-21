// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceClassResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceClassResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceClassResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceClassResponseBody) *ModifyDBInstanceClassResponse
	GetBody() *ModifyDBInstanceClassResponseBody
}

type ModifyDBInstanceClassResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceClassResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceClassResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceClassResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceClassResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceClassResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceClassResponse) GetBody() *ModifyDBInstanceClassResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceClassResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceClassResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceClassResponse) SetStatusCode(v int32) *ModifyDBInstanceClassResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceClassResponse) SetBody(v *ModifyDBInstanceClassResponseBody) *ModifyDBInstanceClassResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceClassResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
