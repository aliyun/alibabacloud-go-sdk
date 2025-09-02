// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDISyncInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v int64) *StartDISyncInstanceRequest
	GetFileId() *int64
	SetProjectId(v int64) *StartDISyncInstanceRequest
	GetProjectId() *int64
	SetStartParam(v string) *StartDISyncInstanceRequest
	GetStartParam() *string
	SetTaskType(v string) *StartDISyncInstanceRequest
	GetTaskType() *string
}

type StartDISyncInstanceRequest struct {
	// 	- If you set TaskType to DI_REALTIME, set this parameter to the ID of the real-time synchronization task that you want to start.
	//
	// 	- If you set TaskType to DI_SOLUTION, set this parameter to the ID of the data synchronization solution that you want to start.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the workspace ID. You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 	- If you set TaskType to DI_REALTIME, the StartParam parameter specifies the startup parameters for the real-time synchronization task. The startup parameters include failover-related parameters, the parameter that specifies the number of dirty data records allowed, and the parameters in the data definition language (DDL) statements.
	//
	// 	- If you set TaskType to DI_SOLUTION, the StartParam parameter does not take effect.
	//
	// example:
	//
	// {"failoverLimit":{"count":10,"interval":30},"errorLimit":{"record":0},"ddlMarkMap":{"RENAMECOLUMN":"WARNING","DROPTABLE":"WARNING","CREATETABLE":"IGNORE","MODIFYCOLUMN":"WARNING","TRUNCATETABLE":"NORMAL","DROPCOLUMN":"IGNORE","ADDCOLUMN":"NORMAL","RENAMETABLE":"CRITICAL"}}
	StartParam *string `json:"StartParam,omitempty" xml:"StartParam,omitempty"`
	// The type of the object that you want to start. Valid values:
	//
	// 	- DI_REALTIME: real-time synchronization task
	//
	// 	- DI_SOLUTION: data synchronization solution
	//
	// This parameter is required.
	//
	// example:
	//
	// DI_REALTIME
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s StartDISyncInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDISyncInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartDISyncInstanceRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *StartDISyncInstanceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *StartDISyncInstanceRequest) GetStartParam() *string {
	return s.StartParam
}

func (s *StartDISyncInstanceRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *StartDISyncInstanceRequest) SetFileId(v int64) *StartDISyncInstanceRequest {
	s.FileId = &v
	return s
}

func (s *StartDISyncInstanceRequest) SetProjectId(v int64) *StartDISyncInstanceRequest {
	s.ProjectId = &v
	return s
}

func (s *StartDISyncInstanceRequest) SetStartParam(v string) *StartDISyncInstanceRequest {
	s.StartParam = &v
	return s
}

func (s *StartDISyncInstanceRequest) SetTaskType(v string) *StartDISyncInstanceRequest {
	s.TaskType = &v
	return s
}

func (s *StartDISyncInstanceRequest) Validate() error {
	return dara.Validate(s)
}
