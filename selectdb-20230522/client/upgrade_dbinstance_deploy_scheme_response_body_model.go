// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceDeploySchemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeDBInstanceDeploySchemeResponseBody
	GetRequestId() *string
}

type UpgradeDBInstanceDeploySchemeResponseBody struct {
	// example:
	//
	// ADF42B18-43FD-5100-83A9-BE81AB70C863
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeDBInstanceDeploySchemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceDeploySchemeResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceDeploySchemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeDBInstanceDeploySchemeResponseBody) SetRequestId(v string) *UpgradeDBInstanceDeploySchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeDBInstanceDeploySchemeResponseBody) Validate() error {
	return dara.Validate(s)
}
