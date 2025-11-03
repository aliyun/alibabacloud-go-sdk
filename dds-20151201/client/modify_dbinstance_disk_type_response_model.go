// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceDiskTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceDiskTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceDiskTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceDiskTypeResponseBody) *ModifyDBInstanceDiskTypeResponse
	GetBody() *ModifyDBInstanceDiskTypeResponseBody
}

type ModifyDBInstanceDiskTypeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceDiskTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceDiskTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceDiskTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDiskTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceDiskTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceDiskTypeResponse) GetBody() *ModifyDBInstanceDiskTypeResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceDiskTypeResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceDiskTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceDiskTypeResponse) SetStatusCode(v int32) *ModifyDBInstanceDiskTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceDiskTypeResponse) SetBody(v *ModifyDBInstanceDiskTypeResponseBody) *ModifyDBInstanceDiskTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceDiskTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
