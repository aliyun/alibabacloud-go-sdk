// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterMigrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterMigrationResponseBody
	GetRequestId() *string
}

type ModifyDBClusterMigrationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A1B303A5-653F-4AEE-A598-023FF9******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterMigrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterMigrationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMigrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterMigrationResponseBody) SetRequestId(v string) *ModifyDBClusterMigrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterMigrationResponseBody) Validate() error {
	return dara.Validate(s)
}
