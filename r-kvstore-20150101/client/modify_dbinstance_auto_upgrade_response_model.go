// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceAutoUpgradeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceAutoUpgradeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceAutoUpgradeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceAutoUpgradeResponseBody) *ModifyDBInstanceAutoUpgradeResponse
	GetBody() *ModifyDBInstanceAutoUpgradeResponseBody
}

type ModifyDBInstanceAutoUpgradeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceAutoUpgradeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceAutoUpgradeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceAutoUpgradeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceAutoUpgradeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceAutoUpgradeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceAutoUpgradeResponse) GetBody() *ModifyDBInstanceAutoUpgradeResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceAutoUpgradeResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceAutoUpgradeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeResponse) SetStatusCode(v int32) *ModifyDBInstanceAutoUpgradeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeResponse) SetBody(v *ModifyDBInstanceAutoUpgradeResponseBody) *ModifyDBInstanceAutoUpgradeResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
