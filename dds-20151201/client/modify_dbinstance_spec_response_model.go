// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceSpecResponseBody) *ModifyDBInstanceSpecResponse
	GetBody() *ModifyDBInstanceSpecResponseBody
}

type ModifyDBInstanceSpecResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceSpecResponse) GetBody() *ModifyDBInstanceSpecResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceSpecResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceSpecResponse) SetStatusCode(v int32) *ModifyDBInstanceSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceSpecResponse) SetBody(v *ModifyDBInstanceSpecResponseBody) *ModifyDBInstanceSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
