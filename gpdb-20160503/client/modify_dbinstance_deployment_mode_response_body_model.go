// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceDeploymentModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBInstanceDeploymentModeResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceDeploymentModeResponseBody struct {
	// The unique ID of the request.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceDeploymentModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceDeploymentModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDeploymentModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceDeploymentModeResponseBody) SetRequestId(v string) *ModifyDBInstanceDeploymentModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceDeploymentModeResponseBody) Validate() error {
	return dara.Validate(s)
}
