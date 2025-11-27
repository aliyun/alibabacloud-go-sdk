// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceReplicationSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceReplicationSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceReplicationSwitchResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceReplicationSwitchResponseBody) *ModifyDBInstanceReplicationSwitchResponse
	GetBody() *ModifyDBInstanceReplicationSwitchResponseBody
}

type ModifyDBInstanceReplicationSwitchResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceReplicationSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceReplicationSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceReplicationSwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceReplicationSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceReplicationSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceReplicationSwitchResponse) GetBody() *ModifyDBInstanceReplicationSwitchResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceReplicationSwitchResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceReplicationSwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceReplicationSwitchResponse) SetStatusCode(v int32) *ModifyDBInstanceReplicationSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceReplicationSwitchResponse) SetBody(v *ModifyDBInstanceReplicationSwitchResponseBody) *ModifyDBInstanceReplicationSwitchResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceReplicationSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
