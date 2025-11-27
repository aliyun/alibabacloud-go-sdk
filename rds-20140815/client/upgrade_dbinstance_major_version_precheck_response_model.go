// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceMajorVersionPrecheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeDBInstanceMajorVersionPrecheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeDBInstanceMajorVersionPrecheckResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeDBInstanceMajorVersionPrecheckResponseBody) *UpgradeDBInstanceMajorVersionPrecheckResponse
	GetBody() *UpgradeDBInstanceMajorVersionPrecheckResponseBody
}

type UpgradeDBInstanceMajorVersionPrecheckResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeDBInstanceMajorVersionPrecheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeDBInstanceMajorVersionPrecheckResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceMajorVersionPrecheckResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceMajorVersionPrecheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeDBInstanceMajorVersionPrecheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeDBInstanceMajorVersionPrecheckResponse) GetBody() *UpgradeDBInstanceMajorVersionPrecheckResponseBody {
	return s.Body
}

func (s *UpgradeDBInstanceMajorVersionPrecheckResponse) SetHeaders(v map[string]*string) *UpgradeDBInstanceMajorVersionPrecheckResponse {
	s.Headers = v
	return s
}

func (s *UpgradeDBInstanceMajorVersionPrecheckResponse) SetStatusCode(v int32) *UpgradeDBInstanceMajorVersionPrecheckResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionPrecheckResponse) SetBody(v *UpgradeDBInstanceMajorVersionPrecheckResponseBody) *UpgradeDBInstanceMajorVersionPrecheckResponse {
	s.Body = v
	return s
}

func (s *UpgradeDBInstanceMajorVersionPrecheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
