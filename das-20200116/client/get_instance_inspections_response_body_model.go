// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceInspectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceInspectionsResponseBody
	GetCode() *string
	SetData(v *GetInstanceInspectionsResponseBodyData) *GetInstanceInspectionsResponseBody
	GetData() *GetInstanceInspectionsResponseBodyData
	SetMessage(v string) *GetInstanceInspectionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceInspectionsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetInstanceInspectionsResponseBody
	GetSuccess() *string
}

type GetInstanceInspectionsResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details.
	Data *GetInstanceInspectionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, Successful is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceInspectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceInspectionsResponseBody) GetData() *GetInstanceInspectionsResponseBodyData {
	return s.Data
}

func (s *GetInstanceInspectionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceInspectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceInspectionsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetInstanceInspectionsResponseBody) SetCode(v string) *GetInstanceInspectionsResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceInspectionsResponseBody) SetData(v *GetInstanceInspectionsResponseBodyData) *GetInstanceInspectionsResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceInspectionsResponseBody) SetMessage(v string) *GetInstanceInspectionsResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceInspectionsResponseBody) SetRequestId(v string) *GetInstanceInspectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceInspectionsResponseBody) SetSuccess(v string) *GetInstanceInspectionsResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceInspectionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionsResponseBodyData struct {
	// The detailed information.
	List []*GetInstanceInspectionsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number. The value returned is a positive integer. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetInstanceInspectionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionsResponseBodyData) GetList() []*GetInstanceInspectionsResponseBodyDataList {
	return s.List
}

func (s *GetInstanceInspectionsResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetInstanceInspectionsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetInstanceInspectionsResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetInstanceInspectionsResponseBodyData) SetList(v []*GetInstanceInspectionsResponseBodyDataList) *GetInstanceInspectionsResponseBodyData {
	s.List = v
	return s
}

func (s *GetInstanceInspectionsResponseBodyData) SetPageNo(v int64) *GetInstanceInspectionsResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyData) SetPageSize(v int64) *GetInstanceInspectionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyData) SetTotal(v int64) *GetInstanceInspectionsResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionsResponseBodyDataList struct {
	// Indicates whether the autonomy service is enabled.
	AutoFunction *GetInstanceInspectionsResponseBodyDataListAutoFunction `json:"AutoFunction,omitempty" xml:"AutoFunction,omitempty" type:"Struct"`
	// The data returned.
	//
	// example:
	//
	// "data": { "hasDeadLock": false, "exceptionTableMap": {}, "bigTransactionCount": 0, "cpu": 4, "isRds": true, "rdsEnable": true, "enable": false, "activeSessions": [], "bigTransactionList": [], "bigSessionList": [ { "blockDuration": 0, "active": false, "Time": 0, "db": "" },
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// Indicates whether DAS Enterprise Edition is enabled. Valid values:
	//
	// 	- **0**: disabled.
	//
	// 	- **1**: enabled.
	//
	// 	- **2**: not supported.
	//
	// example:
	//
	// 0
	EnableDasPro *int32 `json:"EnableDasPro,omitempty" xml:"EnableDasPro,omitempty"`
	// The end time of the inspection and scoring task. The value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The end time must be later than the start time.
	//
	// example:
	//
	// 1608888296001
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the task was created. The value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1603247192000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The information about the instance.
	Instance *GetInstanceInspectionsResponseBodyDataListInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// The inspection score of the instance.
	//
	// example:
	//
	// 100
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The scores that are deducted for the instance.
	ScoreMap map[string]interface{} `json:"ScoreMap,omitempty" xml:"ScoreMap,omitempty"`
	// The start time of the inspection and scoring task. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1608888296000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the inspection and scoring task. Valid values:
	//
	// 	- **0**: The task is waiting for execution.
	//
	// 	- **1**: The task is in progress.
	//
	// 	- **2**: The task is complete.
	//
	// example:
	//
	// 2
	State *int32 `json:"State,omitempty" xml:"State,omitempty"`
	// The mode in which the inspection and scoring task was initiated. Valid values:
	//
	// 	- **0**: automatic mode.
	//
	// 	- **1**: manual mode.
	//
	// example:
	//
	// 0
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetInstanceInspectionsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionsResponseBodyDataList) GetAutoFunction() *GetInstanceInspectionsResponseBodyDataListAutoFunction {
	return s.AutoFunction
}

func (s *GetInstanceInspectionsResponseBodyDataList) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetInstanceInspectionsResponseBodyDataList) GetEnableDasPro() *int32 {
	return s.EnableDasPro
}

func (s *GetInstanceInspectionsResponseBodyDataList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetInstanceInspectionsResponseBodyDataList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetInstanceInspectionsResponseBodyDataList) GetInstance() *GetInstanceInspectionsResponseBodyDataListInstance {
	return s.Instance
}

func (s *GetInstanceInspectionsResponseBodyDataList) GetScore() *int32 {
	return s.Score
}

func (s *GetInstanceInspectionsResponseBodyDataList) GetScoreMap() map[string]interface{} {
	return s.ScoreMap
}

func (s *GetInstanceInspectionsResponseBodyDataList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetInstanceInspectionsResponseBodyDataList) GetState() *int32 {
	return s.State
}

func (s *GetInstanceInspectionsResponseBodyDataList) GetTaskType() *int32 {
	return s.TaskType
}

func (s *GetInstanceInspectionsResponseBodyDataList) SetAutoFunction(v *GetInstanceInspectionsResponseBodyDataListAutoFunction) *GetInstanceInspectionsResponseBodyDataList {
	s.AutoFunction = v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataList) SetData(v map[string]interface{}) *GetInstanceInspectionsResponseBodyDataList {
	s.Data = v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataList) SetEnableDasPro(v int32) *GetInstanceInspectionsResponseBodyDataList {
	s.EnableDasPro = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataList) SetEndTime(v int64) *GetInstanceInspectionsResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataList) SetGmtCreate(v int64) *GetInstanceInspectionsResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataList) SetInstance(v *GetInstanceInspectionsResponseBodyDataListInstance) *GetInstanceInspectionsResponseBodyDataList {
	s.Instance = v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataList) SetScore(v int32) *GetInstanceInspectionsResponseBodyDataList {
	s.Score = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataList) SetScoreMap(v map[string]interface{}) *GetInstanceInspectionsResponseBodyDataList {
	s.ScoreMap = v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataList) SetStartTime(v int64) *GetInstanceInspectionsResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataList) SetState(v int32) *GetInstanceInspectionsResponseBodyDataList {
	s.State = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataList) SetTaskType(v int32) *GetInstanceInspectionsResponseBodyDataList {
	s.TaskType = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionsResponseBodyDataListAutoFunction struct {
	// Indicates whether the feature of automatically creating and deleting indexes is enabled. Valid values:
	//
	// 	- **0**: disabled.
	//
	// 	- **1**: enabled.
	//
	// 	- **2**: not supported.
	//
	// example:
	//
	// 2
	AutoIndex *int32 `json:"AutoIndex,omitempty" xml:"AutoIndex,omitempty"`
	// Indicates whether the automatic throttling feature is enabled. Valid values:
	//
	// 	- **0**: disabled.
	//
	// 	- **1**: enabled.
	//
	// 	- **2**: not supported.
	//
	// example:
	//
	// 2
	AutoLimitedSql *int32 `json:"AutoLimitedSql,omitempty" xml:"AutoLimitedSql,omitempty"`
	// Indicates whether the automatic fragment recycling feature is enabled. Valid values:
	//
	// 	- **0**: disabled.
	//
	// 	- **1**: enabled.
	//
	// 	- **2**: not supported.
	//
	// example:
	//
	// 0
	AutoResourceOptimize *int32 `json:"AutoResourceOptimize,omitempty" xml:"AutoResourceOptimize,omitempty"`
	// Indicates whether the auto scaling feature is enabled. Valid values:
	//
	// 	- **0**: disabled.
	//
	// 	- **1**: enabled.
	//
	// 	- **2**: not supported.
	//
	// example:
	//
	// 0
	AutoScale *int32 `json:"AutoScale,omitempty" xml:"AutoScale,omitempty"`
	// Indicates whether the event subscription feature is enabled. Valid values:
	//
	// 	- **0**: disabled.
	//
	// 	- **1**: enabled.
	//
	// 	- **2**: not supported.
	//
	// example:
	//
	// 0
	EventSubscription *int32 `json:"EventSubscription,omitempty" xml:"EventSubscription,omitempty"`
}

func (s GetInstanceInspectionsResponseBodyDataListAutoFunction) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionsResponseBodyDataListAutoFunction) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionsResponseBodyDataListAutoFunction) GetAutoIndex() *int32 {
	return s.AutoIndex
}

func (s *GetInstanceInspectionsResponseBodyDataListAutoFunction) GetAutoLimitedSql() *int32 {
	return s.AutoLimitedSql
}

func (s *GetInstanceInspectionsResponseBodyDataListAutoFunction) GetAutoResourceOptimize() *int32 {
	return s.AutoResourceOptimize
}

func (s *GetInstanceInspectionsResponseBodyDataListAutoFunction) GetAutoScale() *int32 {
	return s.AutoScale
}

func (s *GetInstanceInspectionsResponseBodyDataListAutoFunction) GetEventSubscription() *int32 {
	return s.EventSubscription
}

func (s *GetInstanceInspectionsResponseBodyDataListAutoFunction) SetAutoIndex(v int32) *GetInstanceInspectionsResponseBodyDataListAutoFunction {
	s.AutoIndex = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListAutoFunction) SetAutoLimitedSql(v int32) *GetInstanceInspectionsResponseBodyDataListAutoFunction {
	s.AutoLimitedSql = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListAutoFunction) SetAutoResourceOptimize(v int32) *GetInstanceInspectionsResponseBodyDataListAutoFunction {
	s.AutoResourceOptimize = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListAutoFunction) SetAutoScale(v int32) *GetInstanceInspectionsResponseBodyDataListAutoFunction {
	s.AutoScale = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListAutoFunction) SetEventSubscription(v int32) *GetInstanceInspectionsResponseBodyDataListAutoFunction {
	s.EventSubscription = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListAutoFunction) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionsResponseBodyDataListInstance struct {
	// The account ID. You can view the ID of the logon account by moving the pointer over the profile in the Alibaba Cloud management console.
	//
	// example:
	//
	// 108398049688****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The connection mode of the instance. Valid values:
	//
	// 	- **standard**: standard mode.
	//
	// 	- **safe**: database proxy mode.
	//
	// example:
	//
	// standard
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The CPU specification of the instance. For example, if a value of 8 is returned, the instance has eight CPU cores.
	//
	// example:
	//
	// 8
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The database engine. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **Redis**
	//
	// 	- **PolarDBMySQL**
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The version number of the database engine.
	//
	// example:
	//
	// 8.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The instance name.
	//
	// example:
	//
	// test-01
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// The type of the instance on which the database is deployed. Valid values:
	//
	// 	- **RDS**: an Alibaba Cloud database instance.
	//
	// 	- **ECS**: an Elastic Compute Service (ECS) instance on which a self-managed database is deployed.
	//
	// 	- **IDC**: a self-managed database instance that is not deployed on Alibaba Cloud.
	//
	// >  The value IDC indicates that the instance is deployed in a data center.
	//
	// example:
	//
	// RDS
	InstanceArea *string `json:"InstanceArea,omitempty" xml:"InstanceArea,omitempty"`
	// The instance type.
	//
	// example:
	//
	// rds.mysql.s2.xlarge
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-bp10usoc1erj7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The memory capacity of the database that is deployed on the instance. Unit: MB.
	//
	// example:
	//
	// 32768
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The network type of the instance.
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The ID of the node on the instance.
	//
	// example:
	//
	// rm-bp10usoc1erj7****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The storage space of the instance. Unit: GB.
	//
	// example:
	//
	// 150
	Storage *int32 `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// The unique identifier of the instance.
	//
	// example:
	//
	// hdm_3063db6792965c080a4bcb6e6304****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the instance is deployed.
	//
	// example:
	//
	// vpc-bp1knt7m55z9exoo7****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetInstanceInspectionsResponseBodyDataListInstance) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionsResponseBodyDataListInstance) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) GetAccountId() *string {
	return s.AccountId
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) GetCategory() *string {
	return s.Category
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) GetCpu() *string {
	return s.Cpu
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) GetEngine() *string {
	return s.Engine
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) GetInstanceArea() *string {
	return s.InstanceArea
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) GetMemory() *int32 {
	return s.Memory
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) GetNodeId() *string {
	return s.NodeId
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) GetRegion() *string {
	return s.Region
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) GetStorage() *int32 {
	return s.Storage
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) GetUuid() *string {
	return s.Uuid
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) GetVpcId() *string {
	return s.VpcId
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetAccountId(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.AccountId = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetCategory(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.Category = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetCpu(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.Cpu = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetEngine(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.Engine = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetEngineVersion(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.EngineVersion = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetInstanceAlias(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.InstanceAlias = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetInstanceArea(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.InstanceArea = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetInstanceClass(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.InstanceClass = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetInstanceId(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetMemory(v int32) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.Memory = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetNetworkType(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.NetworkType = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetNodeId(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.NodeId = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetRegion(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.Region = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetStorage(v int32) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.Storage = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetUuid(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.Uuid = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) SetVpcId(v string) *GetInstanceInspectionsResponseBodyDataListInstance {
	s.VpcId = &v
	return s
}

func (s *GetInstanceInspectionsResponseBodyDataListInstance) Validate() error {
	return dara.Validate(s)
}
