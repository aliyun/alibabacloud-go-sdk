// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDISyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v int64) *UpdateDISyncTaskRequest
	GetFileId() *int64
	SetProjectId(v int64) *UpdateDISyncTaskRequest
	GetProjectId() *int64
	SetTaskContent(v string) *UpdateDISyncTaskRequest
	GetTaskContent() *string
	SetTaskParam(v string) *UpdateDISyncTaskRequest
	GetTaskParam() *string
	SetTaskType(v string) *UpdateDISyncTaskRequest
	GetTaskType() *string
}

type UpdateDISyncTaskRequest struct {
	// The ID of the data synchronization task. You can call the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000000
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the workspace ID. You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The updated configurations of the data synchronization task. Calling this API operation to update a data synchronization task is equivalent to updating a data synchronization task by using the code editor in the DataWorks console. For more information, see [Create a synchronization task by using the code editor](https://help.aliyun.com/document_detail/137717.html). You can call the UpdateDISyncTask operation to update only batch synchronization tasks. If you do not need to update the configurations of the data synchronization task, leave this parameter empty.
	//
	// example:
	//
	// {"type":"job","version":"2.0","steps":[{"stepType":"mysql","parameter":{"envType":1,"datasource":"mysql_pub","column":["id","name","create_time","age","score","t_01"],"connection":[{"datasource":"mysql_pub","table":["u_pk"]}],"where":"","splitPk":"id","encoding":"UTF-8"},"name":"Reader","category":"reader"},{"stepType":"odps","parameter":{"partition":"pt=${bizdate}","truncate":true,"datasource":"odps_first","envType":1,"column":["id","name","create_time","age","score","t_01"],"emptyAsNull":false,"tableComment":"null","table":"u_pk"},"name":"Writer","category":"writer"}],"setting":{"executeMode":null,"errorLimit":{"record":""},"speed":{"concurrent":2,"throttle":false}},"order":{"hops":[{"from":"Reader","to":"Writer"}]}}
	TaskContent *string `json:"TaskContent,omitempty" xml:"TaskContent,omitempty"`
	// The configuration parameters of the data synchronization task. You must configure this parameter in the JSON format.
	//
	// 	- ResourceGroup: the identifier of the resource group for Data Integration that is used by the data synchronization task. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/173913.html) operation to query the identifier of the resource group.
	//
	// 	- Cu: the specifications occupied by the data synchronization task in the serverless resource group. The value of this parameter must be a multiple of 0.5.
	//
	// example:
	//
	// {"ResourceGroup":"S_res_group_XXX_XXXX"}
	TaskParam *string `json:"TaskParam,omitempty" xml:"TaskParam,omitempty"`
	// The type of the data synchronization task. Set the value to DI_OFFLINE. You can call the UpdateDISyncTask operation to update only batch synchronization tasks.
	//
	// This parameter is required.
	//
	// example:
	//
	// DI_OFFLINE
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s UpdateDISyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDISyncTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateDISyncTaskRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *UpdateDISyncTaskRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateDISyncTaskRequest) GetTaskContent() *string {
	return s.TaskContent
}

func (s *UpdateDISyncTaskRequest) GetTaskParam() *string {
	return s.TaskParam
}

func (s *UpdateDISyncTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *UpdateDISyncTaskRequest) SetFileId(v int64) *UpdateDISyncTaskRequest {
	s.FileId = &v
	return s
}

func (s *UpdateDISyncTaskRequest) SetProjectId(v int64) *UpdateDISyncTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDISyncTaskRequest) SetTaskContent(v string) *UpdateDISyncTaskRequest {
	s.TaskContent = &v
	return s
}

func (s *UpdateDISyncTaskRequest) SetTaskParam(v string) *UpdateDISyncTaskRequest {
	s.TaskParam = &v
	return s
}

func (s *UpdateDISyncTaskRequest) SetTaskType(v string) *UpdateDISyncTaskRequest {
	s.TaskType = &v
	return s
}

func (s *UpdateDISyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
