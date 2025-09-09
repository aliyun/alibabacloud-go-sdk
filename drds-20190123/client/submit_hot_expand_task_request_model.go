// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHotExpandTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *SubmitHotExpandTaskRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *SubmitHotExpandTaskRequest
	GetDrdsInstanceId() *string
	SetExtendedMapping(v []*SubmitHotExpandTaskRequestExtendedMapping) *SubmitHotExpandTaskRequest
	GetExtendedMapping() []*SubmitHotExpandTaskRequestExtendedMapping
	SetInstanceDbMapping(v []*SubmitHotExpandTaskRequestInstanceDbMapping) *SubmitHotExpandTaskRequest
	GetInstanceDbMapping() []*SubmitHotExpandTaskRequestInstanceDbMapping
	SetMapping(v []*SubmitHotExpandTaskRequestMapping) *SubmitHotExpandTaskRequest
	GetMapping() []*SubmitHotExpandTaskRequestMapping
	SetSupperAccountMapping(v []*SubmitHotExpandTaskRequestSupperAccountMapping) *SubmitHotExpandTaskRequest
	GetSupperAccountMapping() []*SubmitHotExpandTaskRequestSupperAccountMapping
	SetTaskDesc(v string) *SubmitHotExpandTaskRequest
	GetTaskDesc() *string
	SetTaskName(v string) *SubmitHotExpandTaskRequest
	GetTaskName() *string
}

type SubmitHotExpandTaskRequest struct {
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdshbga1138****
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The information about the database on which you want to perform hot-spot scale-out.
	//
	// This parameter is required.
	ExtendedMapping []*SubmitHotExpandTaskRequestExtendedMapping `json:"ExtendedMapping,omitempty" xml:"ExtendedMapping,omitempty" type:"Repeated"`
	// The information about the instance to which the hot-spot database belongs.
	//
	// This parameter is required.
	InstanceDbMapping []*SubmitHotExpandTaskRequestInstanceDbMapping `json:"InstanceDbMapping,omitempty" xml:"InstanceDbMapping,omitempty" type:"Repeated"`
	// The information about the hot-spot database.
	//
	// This parameter is required.
	Mapping []*SubmitHotExpandTaskRequestMapping `json:"Mapping,omitempty" xml:"Mapping,omitempty" type:"Repeated"`
	// The information about the privileged account.
	SupperAccountMapping []*SubmitHotExpandTaskRequestSupperAccountMapping `json:"SupperAccountMapping,omitempty" xml:"SupperAccountMapping,omitempty" type:"Repeated"`
	// The description of the task.
	//
	// example:
	//
	// test
	TaskDesc *string `json:"TaskDesc,omitempty" xml:"TaskDesc,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// test
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s SubmitHotExpandTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotExpandTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskRequest) GetDbName() *string {
	return s.DbName
}

func (s *SubmitHotExpandTaskRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *SubmitHotExpandTaskRequest) GetExtendedMapping() []*SubmitHotExpandTaskRequestExtendedMapping {
	return s.ExtendedMapping
}

func (s *SubmitHotExpandTaskRequest) GetInstanceDbMapping() []*SubmitHotExpandTaskRequestInstanceDbMapping {
	return s.InstanceDbMapping
}

func (s *SubmitHotExpandTaskRequest) GetMapping() []*SubmitHotExpandTaskRequestMapping {
	return s.Mapping
}

func (s *SubmitHotExpandTaskRequest) GetSupperAccountMapping() []*SubmitHotExpandTaskRequestSupperAccountMapping {
	return s.SupperAccountMapping
}

func (s *SubmitHotExpandTaskRequest) GetTaskDesc() *string {
	return s.TaskDesc
}

func (s *SubmitHotExpandTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *SubmitHotExpandTaskRequest) SetDbName(v string) *SubmitHotExpandTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetDrdsInstanceId(v string) *SubmitHotExpandTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetExtendedMapping(v []*SubmitHotExpandTaskRequestExtendedMapping) *SubmitHotExpandTaskRequest {
	s.ExtendedMapping = v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetInstanceDbMapping(v []*SubmitHotExpandTaskRequestInstanceDbMapping) *SubmitHotExpandTaskRequest {
	s.InstanceDbMapping = v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetMapping(v []*SubmitHotExpandTaskRequestMapping) *SubmitHotExpandTaskRequest {
	s.Mapping = v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetSupperAccountMapping(v []*SubmitHotExpandTaskRequestSupperAccountMapping) *SubmitHotExpandTaskRequest {
	s.SupperAccountMapping = v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetTaskDesc(v string) *SubmitHotExpandTaskRequest {
	s.TaskDesc = &v
	return s
}

func (s *SubmitHotExpandTaskRequest) SetTaskName(v string) *SubmitHotExpandTaskRequest {
	s.TaskName = &v
	return s
}

func (s *SubmitHotExpandTaskRequest) Validate() error {
	return dara.Validate(s)
}

type SubmitHotExpandTaskRequestExtendedMapping struct {
	// The name of the source physical database.
	//
	// example:
	//
	// test
	SrcDb *string `json:"SrcDb,omitempty" xml:"SrcDb,omitempty"`
	// The ID of the ApsaraDB RDS instance to which the source physical database belongs.
	//
	// example:
	//
	// rm-bp1t1mk5a5bdj****
	SrcInstanceId *string `json:"SrcInstanceId,omitempty" xml:"SrcInstanceId,omitempty"`
}

func (s SubmitHotExpandTaskRequestExtendedMapping) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotExpandTaskRequestExtendedMapping) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskRequestExtendedMapping) GetSrcDb() *string {
	return s.SrcDb
}

func (s *SubmitHotExpandTaskRequestExtendedMapping) GetSrcInstanceId() *string {
	return s.SrcInstanceId
}

func (s *SubmitHotExpandTaskRequestExtendedMapping) SetSrcDb(v string) *SubmitHotExpandTaskRequestExtendedMapping {
	s.SrcDb = &v
	return s
}

func (s *SubmitHotExpandTaskRequestExtendedMapping) SetSrcInstanceId(v string) *SubmitHotExpandTaskRequestExtendedMapping {
	s.SrcInstanceId = &v
	return s
}

func (s *SubmitHotExpandTaskRequestExtendedMapping) Validate() error {
	return dara.Validate(s)
}

type SubmitHotExpandTaskRequestInstanceDbMapping struct {
	// The name of the hot-spot database.
	//
	// This parameter is required.
	//
	// example:
	//
	// hot_test_****_****
	DbList *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	// The name of the ApsaraDB RDS instance to which the hot-spot database belongs.
	//
	// example:
	//
	// rm-bp1t1mk5a5bdj****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s SubmitHotExpandTaskRequestInstanceDbMapping) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotExpandTaskRequestInstanceDbMapping) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskRequestInstanceDbMapping) GetDbList() *string {
	return s.DbList
}

func (s *SubmitHotExpandTaskRequestInstanceDbMapping) GetInstanceName() *string {
	return s.InstanceName
}

func (s *SubmitHotExpandTaskRequestInstanceDbMapping) SetDbList(v string) *SubmitHotExpandTaskRequestInstanceDbMapping {
	s.DbList = &v
	return s
}

func (s *SubmitHotExpandTaskRequestInstanceDbMapping) SetInstanceName(v string) *SubmitHotExpandTaskRequestInstanceDbMapping {
	s.InstanceName = &v
	return s
}

func (s *SubmitHotExpandTaskRequestInstanceDbMapping) Validate() error {
	return dara.Validate(s)
}

type SubmitHotExpandTaskRequestMapping struct {
	// The shard key used to split the database to which the associated table belongs.
	//
	// example:
	//
	// platform
	DbShardColumn *string `json:"DbShardColumn,omitempty" xml:"DbShardColumn,omitempty"`
	// The name of the hot-spot database.
	//
	// example:
	//
	// hot_test_****_****
	HotDbName *string `json:"HotDbName,omitempty" xml:"HotDbName,omitempty"`
	// The name of the hot-spot table. The name must be prefixed with the name of a logical table.
	//
	// example:
	//
	// test_table_*****
	HotTableName *string `json:"HotTableName,omitempty" xml:"HotTableName,omitempty"`
	// The name of the logical table on which you want to perform hot-spot scale-out.
	//
	// example:
	//
	// test_table
	LogicTable *string `json:"LogicTable,omitempty" xml:"LogicTable,omitempty"`
	// The value of the shard key used to split a database.
	//
	// example:
	//
	// test
	ShardDbValue *string `json:"ShardDbValue,omitempty" xml:"ShardDbValue,omitempty"`
	// The value of the shard key used to split a table.
	//
	// example:
	//
	// test
	ShardTbValue *string `json:"ShardTbValue,omitempty" xml:"ShardTbValue,omitempty"`
	// The shard key used to split an associated table.
	//
	// example:
	//
	// platform
	TbShardColumn *string `json:"TbShardColumn,omitempty" xml:"TbShardColumn,omitempty"`
}

func (s SubmitHotExpandTaskRequestMapping) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotExpandTaskRequestMapping) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskRequestMapping) GetDbShardColumn() *string {
	return s.DbShardColumn
}

func (s *SubmitHotExpandTaskRequestMapping) GetHotDbName() *string {
	return s.HotDbName
}

func (s *SubmitHotExpandTaskRequestMapping) GetHotTableName() *string {
	return s.HotTableName
}

func (s *SubmitHotExpandTaskRequestMapping) GetLogicTable() *string {
	return s.LogicTable
}

func (s *SubmitHotExpandTaskRequestMapping) GetShardDbValue() *string {
	return s.ShardDbValue
}

func (s *SubmitHotExpandTaskRequestMapping) GetShardTbValue() *string {
	return s.ShardTbValue
}

func (s *SubmitHotExpandTaskRequestMapping) GetTbShardColumn() *string {
	return s.TbShardColumn
}

func (s *SubmitHotExpandTaskRequestMapping) SetDbShardColumn(v string) *SubmitHotExpandTaskRequestMapping {
	s.DbShardColumn = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetHotDbName(v string) *SubmitHotExpandTaskRequestMapping {
	s.HotDbName = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetHotTableName(v string) *SubmitHotExpandTaskRequestMapping {
	s.HotTableName = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetLogicTable(v string) *SubmitHotExpandTaskRequestMapping {
	s.LogicTable = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetShardDbValue(v string) *SubmitHotExpandTaskRequestMapping {
	s.ShardDbValue = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetShardTbValue(v string) *SubmitHotExpandTaskRequestMapping {
	s.ShardTbValue = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) SetTbShardColumn(v string) *SubmitHotExpandTaskRequestMapping {
	s.TbShardColumn = &v
	return s
}

func (s *SubmitHotExpandTaskRequestMapping) Validate() error {
	return dara.Validate(s)
}

type SubmitHotExpandTaskRequestSupperAccountMapping struct {
	// The ID of the ApsaraDB RDS instance that has the privileged account.
	//
	// example:
	//
	// rm-bp1t1mk5a5bdj****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The name of the privileged account of the ApsaraDB RDS instance.
	//
	// example:
	//
	// test
	SupperAccount *string `json:"SupperAccount,omitempty" xml:"SupperAccount,omitempty"`
	// The password of the privileged account of the ApsaraDB RDS instance.
	//
	// example:
	//
	// 11111111
	SupperPassword *string `json:"SupperPassword,omitempty" xml:"SupperPassword,omitempty"`
}

func (s SubmitHotExpandTaskRequestSupperAccountMapping) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotExpandTaskRequestSupperAccountMapping) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskRequestSupperAccountMapping) GetInstanceName() *string {
	return s.InstanceName
}

func (s *SubmitHotExpandTaskRequestSupperAccountMapping) GetSupperAccount() *string {
	return s.SupperAccount
}

func (s *SubmitHotExpandTaskRequestSupperAccountMapping) GetSupperPassword() *string {
	return s.SupperPassword
}

func (s *SubmitHotExpandTaskRequestSupperAccountMapping) SetInstanceName(v string) *SubmitHotExpandTaskRequestSupperAccountMapping {
	s.InstanceName = &v
	return s
}

func (s *SubmitHotExpandTaskRequestSupperAccountMapping) SetSupperAccount(v string) *SubmitHotExpandTaskRequestSupperAccountMapping {
	s.SupperAccount = &v
	return s
}

func (s *SubmitHotExpandTaskRequestSupperAccountMapping) SetSupperPassword(v string) *SubmitHotExpandTaskRequestSupperAccountMapping {
	s.SupperPassword = &v
	return s
}

func (s *SubmitHotExpandTaskRequestSupperAccountMapping) Validate() error {
	return dara.Validate(s)
}
