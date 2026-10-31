// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPhysicalInstanceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetPhysicalInstanceResponseBody
	GetHttpStatusCode() *int32
	SetInstance(v *GetPhysicalInstanceResponseBodyInstance) *GetPhysicalInstanceResponseBody
	GetInstance() *GetPhysicalInstanceResponseBodyInstance
	SetMessage(v string) *GetPhysicalInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPhysicalInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPhysicalInstanceResponseBody
	GetSuccess() *bool
}

type GetPhysicalInstanceResponseBody struct {
	// The error code. A value of OK indicates that the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The instance details.
	Instance *GetPhysicalInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPhysicalInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPhysicalInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPhysicalInstanceResponseBody) GetInstance() *GetPhysicalInstanceResponseBodyInstance {
	return s.Instance
}

func (s *GetPhysicalInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPhysicalInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPhysicalInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPhysicalInstanceResponseBody) SetCode(v string) *GetPhysicalInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhysicalInstanceResponseBody) SetHttpStatusCode(v int32) *GetPhysicalInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPhysicalInstanceResponseBody) SetInstance(v *GetPhysicalInstanceResponseBodyInstance) *GetPhysicalInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *GetPhysicalInstanceResponseBody) SetMessage(v string) *GetPhysicalInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhysicalInstanceResponseBody) SetRequestId(v string) *GetPhysicalInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhysicalInstanceResponseBody) SetSuccess(v bool) *GetPhysicalInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *GetPhysicalInstanceResponseBody) Validate() error {
	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPhysicalInstanceResponseBodyInstance struct {
	// The business date.
	//
	// example:
	//
	// 2023-06-25
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The scheduled time.
	//
	// example:
	//
	// 2023-06-27 00:30:00
	DueTime *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// The execution duration.
	//
	// example:
	//
	// 3600s
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The end time of the execution.
	//
	// example:
	//
	// 2023-06-27 02:30:00
	EndExecuteTime *int64 `json:"EndExecuteTime,omitempty" xml:"EndExecuteTime,omitempty"`
	// The extended information.
	//
	// example:
	//
	// xx
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// t_23231
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The sequence number of the instance.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The node details.
	NodeInfo *GetPhysicalInstanceResponseBodyInstanceNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	// The start time of the execution.
	//
	// example:
	//
	// 2023-06-27 01:30:00
	StartExecuteTime *int64 `json:"StartExecuteTime,omitempty" xml:"StartExecuteTime,omitempty"`
	// The status list.
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s GetPhysicalInstanceResponseBodyInstance) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceResponseBodyInstance) GetBizDate() *string {
	return s.BizDate
}

func (s *GetPhysicalInstanceResponseBodyInstance) GetDueTime() *string {
	return s.DueTime
}

func (s *GetPhysicalInstanceResponseBodyInstance) GetDuration() *string {
	return s.Duration
}

func (s *GetPhysicalInstanceResponseBodyInstance) GetEndExecuteTime() *int64 {
	return s.EndExecuteTime
}

func (s *GetPhysicalInstanceResponseBodyInstance) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *GetPhysicalInstanceResponseBodyInstance) GetId() *string {
	return s.Id
}

func (s *GetPhysicalInstanceResponseBodyInstance) GetIndex() *int32 {
	return s.Index
}

func (s *GetPhysicalInstanceResponseBodyInstance) GetNodeInfo() *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	return s.NodeInfo
}

func (s *GetPhysicalInstanceResponseBodyInstance) GetStartExecuteTime() *int64 {
	return s.StartExecuteTime
}

func (s *GetPhysicalInstanceResponseBodyInstance) GetStatusList() []*string {
	return s.StatusList
}

func (s *GetPhysicalInstanceResponseBodyInstance) SetBizDate(v string) *GetPhysicalInstanceResponseBodyInstance {
	s.BizDate = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstance) SetDueTime(v string) *GetPhysicalInstanceResponseBodyInstance {
	s.DueTime = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstance) SetDuration(v string) *GetPhysicalInstanceResponseBodyInstance {
	s.Duration = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstance) SetEndExecuteTime(v int64) *GetPhysicalInstanceResponseBodyInstance {
	s.EndExecuteTime = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstance) SetExtendInfo(v string) *GetPhysicalInstanceResponseBodyInstance {
	s.ExtendInfo = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstance) SetId(v string) *GetPhysicalInstanceResponseBodyInstance {
	s.Id = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstance) SetIndex(v int32) *GetPhysicalInstanceResponseBodyInstance {
	s.Index = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstance) SetNodeInfo(v *GetPhysicalInstanceResponseBodyInstanceNodeInfo) *GetPhysicalInstanceResponseBodyInstance {
	s.NodeInfo = v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstance) SetStartExecuteTime(v int64) *GetPhysicalInstanceResponseBodyInstance {
	s.StartExecuteTime = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstance) SetStatusList(v []*string) *GetPhysicalInstanceResponseBodyInstance {
	s.StatusList = v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstance) Validate() error {
	if s.NodeInfo != nil {
		if err := s.NodeInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPhysicalInstanceResponseBodyInstanceNodeInfo struct {
	// The business unit name.
	//
	// example:
	//
	// xx
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2023-02-02 23:53:17
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator.
	Creator *GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// xx test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether a dry run is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The source system of the node.
	//
	// example:
	//
	// DATA_PROCESS
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// Indicates whether a development environment is available.
	//
	// example:
	//
	// true
	HasDev *bool `json:"HasDev,omitempty" xml:"HasDev,omitempty"`
	// Indicates whether a production environment is available.
	//
	// example:
	//
	// true
	HasProd *bool `json:"HasProd,omitempty" xml:"HasProd,omitempty"`
	// The node ID.
	//
	// example:
	//
	// n_3232312
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The last modified time.
	//
	// example:
	//
	// 2024-02-02 23:53:17
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The modifier.
	Modifier *GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier `json:"Modifier,omitempty" xml:"Modifier,omitempty" type:"Struct"`
	// The node name.
	//
	// example:
	//
	// xx test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owners.
	OwnerList []*GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	// The priority. Valid values:
	//
	// - HIGHEST
	//
	// - HIGH
	//
	// - MIDDLE
	//
	// - LOW
	//
	// - LOWEST
	PriorityList []*string `json:"PriorityList,omitempty" xml:"PriorityList,omitempty" type:"Repeated"`
	// The resource group.
	ResourceGroupList []*string `json:"ResourceGroupList,omitempty" xml:"ResourceGroupList,omitempty" type:"Repeated"`
	// Indicates whether scheduling is paused.
	//
	// example:
	//
	// false
	SchedulePaused *bool `json:"SchedulePaused,omitempty" xml:"SchedulePaused,omitempty"`
	// The scheduling cycle. Valid values:
	//
	// - YEARLY
	//
	// - MONTHLY
	//
	// - WEEKLY
	//
	// - DAILY
	//
	// - HOURLY
	//
	// - MINUTELY
	SchedulePeriodList []*string `json:"SchedulePeriodList,omitempty" xml:"SchedulePeriodList,omitempty" type:"Repeated"`
	// The sub-business type. Valid values:
	//
	// - MAX_COMPUTE_SQL
	//
	// - HIVE_SQL
	//
	// - SHELL
	//
	// - PYTHON
	//
	// - ONE_SERVICE_SQL
	//
	// - DATABASE_SQL
	//
	// example:
	//
	// SHELL
	SubDetailType *string `json:"SubDetailType,omitempty" xml:"SubDetailType,omitempty"`
	// The node type. Valid values:
	//
	// - BBOX_LOGIC_TABLE_NODE: logical table node
	//
	// - BBOX_LOGIC_FIELD_NODE: logical field node
	//
	// - BBOX_LOGIC_FIELD_GROUP_NODE: field group node
	//
	// - BBOX_INNER_TEMP_NODE: logical table temporary node
	//
	// - DATA_PROCESS: code task
	//
	// - STREAM_TASK_NODE: real-time task
	//
	// - PIPELINE_NODE: pipeline node
	//
	// - FLINK_BATCH: Flink batch task
	//
	// - ODM_NODE: ODM node
	//
	// example:
	//
	// DATA_PROCESS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetPhysicalInstanceResponseBodyInstanceNodeInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalInstanceResponseBodyInstanceNodeInfo) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) GetCreator() *GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator {
	return s.Creator
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) GetDescription() *string {
	return s.Description
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) GetDryRun() *bool {
	return s.DryRun
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) GetFrom() *string {
	return s.From
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) GetHasDev() *bool {
	return s.HasDev
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) GetHasProd() *bool {
	return s.HasProd
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) GetId() *string {
	return s.Id
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) GetModifier() *GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier {
	return s.Modifier
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) GetName() *string {
	return s.Name
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) GetOwnerList() []*GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList {
	return s.OwnerList
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) GetPriorityList() []*string {
	return s.PriorityList
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) GetResourceGroupList() []*string {
	return s.ResourceGroupList
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) GetSchedulePaused() *bool {
	return s.SchedulePaused
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) GetSchedulePeriodList() []*string {
	return s.SchedulePeriodList
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) GetSubDetailType() *string {
	return s.SubDetailType
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) GetType() *string {
	return s.Type
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetBizUnitName(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.BizUnitName = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetCreateTime(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.CreateTime = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetCreator(v *GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.Creator = v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetDescription(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.Description = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetDryRun(v bool) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.DryRun = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetFrom(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.From = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetHasDev(v bool) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.HasDev = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetHasProd(v bool) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.HasProd = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetId(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.Id = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetLastModifiedTime(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.LastModifiedTime = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetModifier(v *GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.Modifier = v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetName(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.Name = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetOwnerList(v []*GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.OwnerList = v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetPriorityList(v []*string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.PriorityList = v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetResourceGroupList(v []*string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.ResourceGroupList = v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetSchedulePaused(v bool) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.SchedulePaused = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetSchedulePeriodList(v []*string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.SchedulePeriodList = v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetSubDetailType(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.SubDetailType = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetType(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.Type = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) Validate() error {
	if s.Creator != nil {
		if err := s.Creator.Validate(); err != nil {
			return err
		}
	}
	if s.Modifier != nil {
		if err := s.Modifier.Validate(); err != nil {
			return err
		}
	}
	if s.OwnerList != nil {
		for _, item := range s.OwnerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator struct {
	// The user ID.
	//
	// example:
	//
	// 2323111
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The username.
	//
	// example:
	//
	// zhangsan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator) GetId() *string {
	return s.Id
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator) GetName() *string {
	return s.Name
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator) SetId(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator {
	s.Id = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator) SetName(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator {
	s.Name = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator) Validate() error {
	return dara.Validate(s)
}

type GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier struct {
	// The user ID.
	//
	// example:
	//
	// 2323111
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The username.
	//
	// example:
	//
	// zhangsan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier) GetId() *string {
	return s.Id
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier) GetName() *string {
	return s.Name
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier) SetId(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier {
	s.Id = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier) SetName(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier {
	s.Name = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier) Validate() error {
	return dara.Validate(s)
}

type GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList struct {
	// The user ID.
	//
	// example:
	//
	// 2323111
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The username.
	//
	// example:
	//
	// zhangsan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList) GetId() *string {
	return s.Id
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList) GetName() *string {
	return s.Name
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList) SetId(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList {
	s.Id = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList) SetName(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList {
	s.Name = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList) Validate() error {
	return dara.Validate(s)
}
