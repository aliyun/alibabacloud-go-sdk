// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterMigrationEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBClusterMigrationEndpointResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *ModifyDBClusterMigrationEndpointResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ModifyDBClusterMigrationEndpointResponseBody
	GetTaskId() *string
}

type ModifyDBClusterMigrationEndpointResponseBody struct {
	// example:
	//
	// pc-k2j1qqukj583di7n9
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F901FB05-8109-547F-A0B9-9C4FF7F4927A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 21498490
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyDBClusterMigrationEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterMigrationEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMigrationEndpointResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterMigrationEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterMigrationEndpointResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyDBClusterMigrationEndpointResponseBody) SetDBClusterId(v string) *ModifyDBClusterMigrationEndpointResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterMigrationEndpointResponseBody) SetRequestId(v string) *ModifyDBClusterMigrationEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterMigrationEndpointResponseBody) SetTaskId(v string) *ModifyDBClusterMigrationEndpointResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyDBClusterMigrationEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
