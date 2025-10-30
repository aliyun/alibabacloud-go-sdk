// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeDBInstanceResponseBody) *UpgradeDBInstanceResponse
	GetBody() *UpgradeDBInstanceResponseBody
}

type UpgradeDBInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeDBInstanceResponse) GetBody() *UpgradeDBInstanceResponseBody {
	return s.Body
}

func (s *UpgradeDBInstanceResponse) SetHeaders(v map[string]*string) *UpgradeDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpgradeDBInstanceResponse) SetStatusCode(v int32) *UpgradeDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeDBInstanceResponse) SetBody(v *UpgradeDBInstanceResponseBody) *UpgradeDBInstanceResponse {
	s.Body = v
	return s
}

func (s *UpgradeDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
