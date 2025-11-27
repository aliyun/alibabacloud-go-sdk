// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceDelayedReplicationTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceDelayedReplicationTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceDelayedReplicationTimeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceDelayedReplicationTimeResponseBody) *ModifyDBInstanceDelayedReplicationTimeResponse
	GetBody() *ModifyDBInstanceDelayedReplicationTimeResponseBody
}

type ModifyDBInstanceDelayedReplicationTimeResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceDelayedReplicationTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceDelayedReplicationTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceDelayedReplicationTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDelayedReplicationTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceDelayedReplicationTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceDelayedReplicationTimeResponse) GetBody() *ModifyDBInstanceDelayedReplicationTimeResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceDelayedReplicationTimeResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceDelayedReplicationTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceDelayedReplicationTimeResponse) SetStatusCode(v int32) *ModifyDBInstanceDelayedReplicationTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceDelayedReplicationTimeResponse) SetBody(v *ModifyDBInstanceDelayedReplicationTimeResponseBody) *ModifyDBInstanceDelayedReplicationTimeResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceDelayedReplicationTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
