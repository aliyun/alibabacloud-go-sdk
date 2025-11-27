// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceAutoUpgradeMinorVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceAutoUpgradeMinorVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceAutoUpgradeMinorVersionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceAutoUpgradeMinorVersionResponseBody) *ModifyDBInstanceAutoUpgradeMinorVersionResponse
	GetBody() *ModifyDBInstanceAutoUpgradeMinorVersionResponseBody
}

type ModifyDBInstanceAutoUpgradeMinorVersionResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceAutoUpgradeMinorVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceAutoUpgradeMinorVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceAutoUpgradeMinorVersionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionResponse) GetBody() *ModifyDBInstanceAutoUpgradeMinorVersionResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceAutoUpgradeMinorVersionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionResponse) SetStatusCode(v int32) *ModifyDBInstanceAutoUpgradeMinorVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionResponse) SetBody(v *ModifyDBInstanceAutoUpgradeMinorVersionResponseBody) *ModifyDBInstanceAutoUpgradeMinorVersionResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeMinorVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
