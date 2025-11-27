// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceMajorVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeDBInstanceMajorVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeDBInstanceMajorVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeDBInstanceMajorVersionResponseBody) *UpgradeDBInstanceMajorVersionResponse
	GetBody() *UpgradeDBInstanceMajorVersionResponseBody
}

type UpgradeDBInstanceMajorVersionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeDBInstanceMajorVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeDBInstanceMajorVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceMajorVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceMajorVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeDBInstanceMajorVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeDBInstanceMajorVersionResponse) GetBody() *UpgradeDBInstanceMajorVersionResponseBody {
	return s.Body
}

func (s *UpgradeDBInstanceMajorVersionResponse) SetHeaders(v map[string]*string) *UpgradeDBInstanceMajorVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeDBInstanceMajorVersionResponse) SetStatusCode(v int32) *UpgradeDBInstanceMajorVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionResponse) SetBody(v *UpgradeDBInstanceMajorVersionResponseBody) *UpgradeDBInstanceMajorVersionResponse {
	s.Body = v
	return s
}

func (s *UpgradeDBInstanceMajorVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
