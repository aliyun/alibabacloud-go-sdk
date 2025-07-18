// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceDeploymentModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceDeploymentModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceDeploymentModeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceDeploymentModeResponseBody) *ModifyDBInstanceDeploymentModeResponse
	GetBody() *ModifyDBInstanceDeploymentModeResponseBody
}

type ModifyDBInstanceDeploymentModeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceDeploymentModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceDeploymentModeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceDeploymentModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDeploymentModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceDeploymentModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceDeploymentModeResponse) GetBody() *ModifyDBInstanceDeploymentModeResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceDeploymentModeResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceDeploymentModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceDeploymentModeResponse) SetStatusCode(v int32) *ModifyDBInstanceDeploymentModeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceDeploymentModeResponse) SetBody(v *ModifyDBInstanceDeploymentModeResponseBody) *ModifyDBInstanceDeploymentModeResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceDeploymentModeResponse) Validate() error {
	return dara.Validate(s)
}
