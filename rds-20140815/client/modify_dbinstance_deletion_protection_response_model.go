// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceDeletionProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceDeletionProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceDeletionProtectionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceDeletionProtectionResponseBody) *ModifyDBInstanceDeletionProtectionResponse
	GetBody() *ModifyDBInstanceDeletionProtectionResponseBody
}

type ModifyDBInstanceDeletionProtectionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceDeletionProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceDeletionProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceDeletionProtectionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDeletionProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceDeletionProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceDeletionProtectionResponse) GetBody() *ModifyDBInstanceDeletionProtectionResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceDeletionProtectionResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceDeletionProtectionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceDeletionProtectionResponse) SetStatusCode(v int32) *ModifyDBInstanceDeletionProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceDeletionProtectionResponse) SetBody(v *ModifyDBInstanceDeletionProtectionResponseBody) *ModifyDBInstanceDeletionProtectionResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceDeletionProtectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
