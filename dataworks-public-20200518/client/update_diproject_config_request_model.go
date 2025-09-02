// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDIProjectConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationType(v string) *UpdateDIProjectConfigRequest
	GetDestinationType() *string
	SetProjectConfig(v string) *UpdateDIProjectConfigRequest
	GetProjectConfig() *string
	SetProjectId(v int64) *UpdateDIProjectConfigRequest
	GetProjectId() *int64
	SetSourceType(v string) *UpdateDIProjectConfigRequest
	GetSourceType() *string
}

type UpdateDIProjectConfigRequest struct {
	// The type of the destinations of the synchronization solutions. This parameter cannot be left empty. Valid values: analyticdb_for_mysql, odps, elasticsearch, holo, mysql, and polardb.
	//
	// This parameter is required.
	//
	// example:
	//
	// odps
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The new default global configuration of the synchronization solutions. The value indicates the processing rules of different types of DDL messages. The value must be in the JSON format. Example: {"RENAMECOLUMN":"WARNING","DROPTABLE":"WARNING","CREATETABLE":"WARNING","MODIFYCOLUMN":"WARNING","TRUNCATETABLE":"WARNING","DROPCOLUMN":"WARNING","ADDCOLUMN":"WARNING","RENAMETABLE":"WARNING"}.
	//
	// Field description:
	//
	// 	- RENAMECOLUMN: renames a column.
	//
	// 	- DROPTABLE: deletes a table.
	//
	// 	- CREATETABLE: creates a table.
	//
	// 	- MODIFYCOLUMN: changes the data type of a column.
	//
	// 	- TRUNCATETABLE: clears a table.
	//
	// 	- DROPCOLUMN: deletes a column.
	//
	// 	- ADDCOLUMN: creates a column.
	//
	// 	- RENAMETABLE: renames a table.
	//
	// DataWorks processes a DDL message of a specific type based on the following rules:
	//
	// 	- WARNING: ignores the message and records an alert in real-time synchronization logs. The alert contains information about the situation that the message is ignored because of an execution error.
	//
	// 	- IGNORE: discards the message and does not send it to the destination.
	//
	// 	- CRITICAL: terminates the real-time synchronization task and sets the node status to Failed.
	//
	// 	- NORMAL: sends the message to the destination to process the message. Each destination processes DDL messages based on its own business logic. If DataWorks adopts the NORMAL policy, DataWorks only forwards DDL messages.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"RENAMECOLUMN":"WARNING","DROPTABLE":"WARNING","CREATETABLE":"WARNING","MODIFYCOLUMN":"WARNING","TRUNCATETABLE":"WARNING","DROPCOLUMN":"WARNING","ADDCOLUMN":"WARNING","RENAMETABLE":"WARNING"}
	ProjectConfig *string `json:"ProjectConfig,omitempty" xml:"ProjectConfig,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The type of the sources of the synchronization solutions. Valid values: oracle, mysql, polardb, datahub, drds, and analyticdb_for_mysql. If you do not configure this parameter, DataWorks applies the default global configuration to all sources.
	//
	// example:
	//
	// mysql
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s UpdateDIProjectConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIProjectConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateDIProjectConfigRequest) GetDestinationType() *string {
	return s.DestinationType
}

func (s *UpdateDIProjectConfigRequest) GetProjectConfig() *string {
	return s.ProjectConfig
}

func (s *UpdateDIProjectConfigRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateDIProjectConfigRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateDIProjectConfigRequest) SetDestinationType(v string) *UpdateDIProjectConfigRequest {
	s.DestinationType = &v
	return s
}

func (s *UpdateDIProjectConfigRequest) SetProjectConfig(v string) *UpdateDIProjectConfigRequest {
	s.ProjectConfig = &v
	return s
}

func (s *UpdateDIProjectConfigRequest) SetProjectId(v int64) *UpdateDIProjectConfigRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDIProjectConfigRequest) SetSourceType(v string) *UpdateDIProjectConfigRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateDIProjectConfigRequest) Validate() error {
	return dara.Validate(s)
}
