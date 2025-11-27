// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceMaintainTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceMaintainTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceMaintainTimeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceMaintainTimeResponseBody) *ModifyDBInstanceMaintainTimeResponse
	GetBody() *ModifyDBInstanceMaintainTimeResponseBody
}

type ModifyDBInstanceMaintainTimeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceMaintainTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceMaintainTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceMaintainTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMaintainTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceMaintainTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceMaintainTimeResponse) GetBody() *ModifyDBInstanceMaintainTimeResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceMaintainTimeResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceMaintainTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceMaintainTimeResponse) SetStatusCode(v int32) *ModifyDBInstanceMaintainTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeResponse) SetBody(v *ModifyDBInstanceMaintainTimeResponseBody) *ModifyDBInstanceMaintainTimeResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceMaintainTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
