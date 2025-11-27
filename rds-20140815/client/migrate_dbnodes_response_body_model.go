// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateDBNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *MigrateDBNodesResponseBody
	GetDBInstanceId() *string
	SetRequestId(v string) *MigrateDBNodesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MigrateDBNodesResponseBody
	GetSuccess() *bool
}

type MigrateDBNodesResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-uf64oq9381l03w1qp
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8B993DA9-5272-5414-94E3-4CA8BA0146C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MigrateDBNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MigrateDBNodesResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateDBNodesResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *MigrateDBNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MigrateDBNodesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MigrateDBNodesResponseBody) SetDBInstanceId(v string) *MigrateDBNodesResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *MigrateDBNodesResponseBody) SetRequestId(v string) *MigrateDBNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateDBNodesResponseBody) SetSuccess(v bool) *MigrateDBNodesResponseBody {
	s.Success = &v
	return s
}

func (s *MigrateDBNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
