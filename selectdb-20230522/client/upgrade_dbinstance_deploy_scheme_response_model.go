// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceDeploySchemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeDBInstanceDeploySchemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeDBInstanceDeploySchemeResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeDBInstanceDeploySchemeResponseBody) *UpgradeDBInstanceDeploySchemeResponse
	GetBody() *UpgradeDBInstanceDeploySchemeResponseBody
}

type UpgradeDBInstanceDeploySchemeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeDBInstanceDeploySchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeDBInstanceDeploySchemeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceDeploySchemeResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceDeploySchemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeDBInstanceDeploySchemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeDBInstanceDeploySchemeResponse) GetBody() *UpgradeDBInstanceDeploySchemeResponseBody {
	return s.Body
}

func (s *UpgradeDBInstanceDeploySchemeResponse) SetHeaders(v map[string]*string) *UpgradeDBInstanceDeploySchemeResponse {
	s.Headers = v
	return s
}

func (s *UpgradeDBInstanceDeploySchemeResponse) SetStatusCode(v int32) *UpgradeDBInstanceDeploySchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeDBInstanceDeploySchemeResponse) SetBody(v *UpgradeDBInstanceDeploySchemeResponseBody) *UpgradeDBInstanceDeploySchemeResponse {
	s.Body = v
	return s
}

func (s *UpgradeDBInstanceDeploySchemeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
