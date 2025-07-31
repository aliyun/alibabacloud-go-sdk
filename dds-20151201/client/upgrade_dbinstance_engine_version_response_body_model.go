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
	// The ID of the request.
	//
	// example:
	//
	// C4907B00-A208-4E0C-A636-AA85140E406C
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
