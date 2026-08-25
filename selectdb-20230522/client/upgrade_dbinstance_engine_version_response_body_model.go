// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceEngineVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeDBInstanceEngineVersionResponseBody
	GetRequestId() *string
}

type UpgradeDBInstanceEngineVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4773E4EC-025D-509F-AEA9-D53123FDFB0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeDBInstanceEngineVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceEngineVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceEngineVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeDBInstanceEngineVersionResponseBody) SetRequestId(v string) *UpgradeDBInstanceEngineVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
