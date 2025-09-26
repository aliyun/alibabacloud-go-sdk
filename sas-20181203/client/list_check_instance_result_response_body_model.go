// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckInstanceResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBasicData(v []*ListCheckInstanceResultResponseBodyBasicData) *ListCheckInstanceResultResponseBody
	GetBasicData() []*ListCheckInstanceResultResponseBodyBasicData
	SetChecks(v []map[string]interface{}) *ListCheckInstanceResultResponseBody
	GetChecks() []map[string]interface{}
	SetColumns(v []*ListCheckInstanceResultResponseBodyColumns) *ListCheckInstanceResultResponseBody
	GetColumns() []*ListCheckInstanceResultResponseBodyColumns
	SetPageInfo(v *ListCheckInstanceResultResponseBodyPageInfo) *ListCheckInstanceResultResponseBody
	GetPageInfo() *ListCheckInstanceResultResponseBodyPageInfo
	SetRequestId(v string) *ListCheckInstanceResultResponseBody
	GetRequestId() *string
}

type ListCheckInstanceResultResponseBody struct {
	// The basic information about the affected instances.
	BasicData []*ListCheckInstanceResultResponseBodyBasicData `json:"BasicData,omitempty" xml:"BasicData,omitempty" type:"Repeated"`
	// The extended information about the instances.
	//
	// example:
	//
	// [{
	//
	// 	"SecurityGroupNameShow": {
	//
	// 		"value": "Sas_Malicious_Ip_Security_Group"
	//
	// 	},
	//
	// 	"InstanceIdShow": {
	//
	// 		"link": "https://ecs.console.aliyun.com/#/securityGroupDetail/region/ap-southeast-1/groupId/sg-t4nbk2aodzio52xvj00s/rule/intranetIngress",
	//
	// 		"value": "sg-t4nbk2aodzio52xv****"
	//
	// 	}
	//
	// }]
	Checks []map[string]interface{} `json:"Checks,omitempty" xml:"Checks,omitempty" type:"Repeated"`
	// The metadata information about the search conditions that can be used to filter instances.
	Columns []*ListCheckInstanceResultResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListCheckInstanceResultResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3AB18264-8A1B-52A6-A9AF-A886556E0F2E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCheckInstanceResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCheckInstanceResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListCheckInstanceResultResponseBody) GetBasicData() []*ListCheckInstanceResultResponseBodyBasicData {
	return s.BasicData
}

func (s *ListCheckInstanceResultResponseBody) GetChecks() []map[string]interface{} {
	return s.Checks
}

func (s *ListCheckInstanceResultResponseBody) GetColumns() []*ListCheckInstanceResultResponseBodyColumns {
	return s.Columns
}

func (s *ListCheckInstanceResultResponseBody) GetPageInfo() *ListCheckInstanceResultResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListCheckInstanceResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCheckInstanceResultResponseBody) SetBasicData(v []*ListCheckInstanceResultResponseBodyBasicData) *ListCheckInstanceResultResponseBody {
	s.BasicData = v
	return s
}

func (s *ListCheckInstanceResultResponseBody) SetChecks(v []map[string]interface{}) *ListCheckInstanceResultResponseBody {
	s.Checks = v
	return s
}

func (s *ListCheckInstanceResultResponseBody) SetColumns(v []*ListCheckInstanceResultResponseBodyColumns) *ListCheckInstanceResultResponseBody {
	s.Columns = v
	return s
}

func (s *ListCheckInstanceResultResponseBody) SetPageInfo(v *ListCheckInstanceResultResponseBodyPageInfo) *ListCheckInstanceResultResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListCheckInstanceResultResponseBody) SetRequestId(v string) *ListCheckInstanceResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCheckInstanceResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCheckInstanceResultResponseBodyBasicData struct {
	// The ID of the check result for the instance.
	//
	// example:
	//
	// 300054
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance ID of the server.
	//
	// example:
	//
	// rm-m5es7ch1s62i4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The information about the instance on which the check item is used.
	InstanceInfo *ListCheckInstanceResultResponseBodyBasicDataInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Struct"`
	// The instance name of the server.
	//
	// example:
	//
	// sg-t4nbk2aodzio52xv****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// ap-southeast-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The states of check items. Multiple states are separated with commas (,). Valid values:
	//
	// 	- **PASS**: passed
	//
	// 	- **NOT_PASS**: failed
	//
	// 	- **CHECKING**: being checked
	//
	// 	- **NOT_CHECK**: not checked
	//
	// 	- **WHITELIST**: added to the whitelist
	//
	// example:
	//
	// NOT_PASS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The exception message of the check item.
	//
	// example:
	//
	// Task is failed.
	StatusMessage  *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	VendorUserName *string `json:"VendorUserName,omitempty" xml:"VendorUserName,omitempty"`
}

func (s ListCheckInstanceResultResponseBodyBasicData) String() string {
	return dara.Prettify(s)
}

func (s ListCheckInstanceResultResponseBodyBasicData) GoString() string {
	return s.String()
}

func (s *ListCheckInstanceResultResponseBodyBasicData) GetId() *int64 {
	return s.Id
}

func (s *ListCheckInstanceResultResponseBodyBasicData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCheckInstanceResultResponseBodyBasicData) GetInstanceInfo() *ListCheckInstanceResultResponseBodyBasicDataInstanceInfo {
	return s.InstanceInfo
}

func (s *ListCheckInstanceResultResponseBodyBasicData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListCheckInstanceResultResponseBodyBasicData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCheckInstanceResultResponseBodyBasicData) GetStatus() *string {
	return s.Status
}

func (s *ListCheckInstanceResultResponseBodyBasicData) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *ListCheckInstanceResultResponseBodyBasicData) GetVendorUserName() *string {
	return s.VendorUserName
}

func (s *ListCheckInstanceResultResponseBodyBasicData) SetId(v int64) *ListCheckInstanceResultResponseBodyBasicData {
	s.Id = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyBasicData) SetInstanceId(v string) *ListCheckInstanceResultResponseBodyBasicData {
	s.InstanceId = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyBasicData) SetInstanceInfo(v *ListCheckInstanceResultResponseBodyBasicDataInstanceInfo) *ListCheckInstanceResultResponseBodyBasicData {
	s.InstanceInfo = v
	return s
}

func (s *ListCheckInstanceResultResponseBodyBasicData) SetInstanceName(v string) *ListCheckInstanceResultResponseBodyBasicData {
	s.InstanceName = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyBasicData) SetRegionId(v string) *ListCheckInstanceResultResponseBodyBasicData {
	s.RegionId = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyBasicData) SetStatus(v string) *ListCheckInstanceResultResponseBodyBasicData {
	s.Status = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyBasicData) SetStatusMessage(v string) *ListCheckInstanceResultResponseBodyBasicData {
	s.StatusMessage = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyBasicData) SetVendorUserName(v string) *ListCheckInstanceResultResponseBodyBasicData {
	s.VendorUserName = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyBasicData) Validate() error {
	return dara.Validate(s)
}

type ListCheckInstanceResultResponseBodyBasicDataInstanceInfo struct {
	// The information about the configuration item whose risks are fixed for the instance.
	Config []*ListCheckInstanceResultResponseBodyBasicDataInstanceInfoConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Repeated"`
	// The time of the first check.
	//
	// example:
	//
	// 1716447535531
	FirstUpdateTime *int64 `json:"FirstUpdateTime,omitempty" xml:"FirstUpdateTime,omitempty"`
	// The time of the last check.
	//
	// example:
	//
	// 1716447535531
	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
}

func (s ListCheckInstanceResultResponseBodyBasicDataInstanceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCheckInstanceResultResponseBodyBasicDataInstanceInfo) GoString() string {
	return s.String()
}

func (s *ListCheckInstanceResultResponseBodyBasicDataInstanceInfo) GetConfig() []*ListCheckInstanceResultResponseBodyBasicDataInstanceInfoConfig {
	return s.Config
}

func (s *ListCheckInstanceResultResponseBodyBasicDataInstanceInfo) GetFirstUpdateTime() *int64 {
	return s.FirstUpdateTime
}

func (s *ListCheckInstanceResultResponseBodyBasicDataInstanceInfo) GetLastUpdateTime() *int64 {
	return s.LastUpdateTime
}

func (s *ListCheckInstanceResultResponseBodyBasicDataInstanceInfo) SetConfig(v []*ListCheckInstanceResultResponseBodyBasicDataInstanceInfoConfig) *ListCheckInstanceResultResponseBodyBasicDataInstanceInfo {
	s.Config = v
	return s
}

func (s *ListCheckInstanceResultResponseBodyBasicDataInstanceInfo) SetFirstUpdateTime(v int64) *ListCheckInstanceResultResponseBodyBasicDataInstanceInfo {
	s.FirstUpdateTime = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyBasicDataInstanceInfo) SetLastUpdateTime(v int64) *ListCheckInstanceResultResponseBodyBasicDataInstanceInfo {
	s.LastUpdateTime = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyBasicDataInstanceInfo) Validate() error {
	return dara.Validate(s)
}

type ListCheckInstanceResultResponseBodyBasicDataInstanceInfoConfig struct {
	// The name of the configuration item, which is unique.
	//
	// example:
	//
	// prot
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The display name of the configuration item for internationalization.
	//
	// example:
	//
	// prot
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// The value of the configuration item specified for the instance.
	//
	// example:
	//
	// 8080
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListCheckInstanceResultResponseBodyBasicDataInstanceInfoConfig) String() string {
	return dara.Prettify(s)
}

func (s ListCheckInstanceResultResponseBodyBasicDataInstanceInfoConfig) GoString() string {
	return s.String()
}

func (s *ListCheckInstanceResultResponseBodyBasicDataInstanceInfoConfig) GetName() *string {
	return s.Name
}

func (s *ListCheckInstanceResultResponseBodyBasicDataInstanceInfoConfig) GetShowName() *string {
	return s.ShowName
}

func (s *ListCheckInstanceResultResponseBodyBasicDataInstanceInfoConfig) GetValue() *string {
	return s.Value
}

func (s *ListCheckInstanceResultResponseBodyBasicDataInstanceInfoConfig) SetName(v string) *ListCheckInstanceResultResponseBodyBasicDataInstanceInfoConfig {
	s.Name = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyBasicDataInstanceInfoConfig) SetShowName(v string) *ListCheckInstanceResultResponseBodyBasicDataInstanceInfoConfig {
	s.ShowName = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyBasicDataInstanceInfoConfig) SetValue(v string) *ListCheckInstanceResultResponseBodyBasicDataInstanceInfoConfig {
	s.Value = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyBasicDataInstanceInfoConfig) Validate() error {
	return dara.Validate(s)
}

type ListCheckInstanceResultResponseBodyColumns struct {
	// The metadata information about the details of the instance.
	Grids []*ListCheckInstanceResultResponseBodyColumnsGrids `json:"Grids,omitempty" xml:"Grids,omitempty" type:"Repeated"`
	// The search condition.
	//
	// example:
	//
	// RegionIdShow
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Indicates whether the search condition is used. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Search *bool `json:"Search,omitempty" xml:"Search,omitempty"`
	// The search key.
	//
	// example:
	//
	// InstanceIdKey
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The display name of the search condition.
	//
	// example:
	//
	// Region
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// The type of the check result for the instance. Valid values:
	//
	// 	- **text**
	//
	// 	- **link**
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCheckInstanceResultResponseBodyColumns) String() string {
	return dara.Prettify(s)
}

func (s ListCheckInstanceResultResponseBodyColumns) GoString() string {
	return s.String()
}

func (s *ListCheckInstanceResultResponseBodyColumns) GetGrids() []*ListCheckInstanceResultResponseBodyColumnsGrids {
	return s.Grids
}

func (s *ListCheckInstanceResultResponseBodyColumns) GetKey() *string {
	return s.Key
}

func (s *ListCheckInstanceResultResponseBodyColumns) GetSearch() *bool {
	return s.Search
}

func (s *ListCheckInstanceResultResponseBodyColumns) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListCheckInstanceResultResponseBodyColumns) GetShowName() *string {
	return s.ShowName
}

func (s *ListCheckInstanceResultResponseBodyColumns) GetType() *string {
	return s.Type
}

func (s *ListCheckInstanceResultResponseBodyColumns) SetGrids(v []*ListCheckInstanceResultResponseBodyColumnsGrids) *ListCheckInstanceResultResponseBodyColumns {
	s.Grids = v
	return s
}

func (s *ListCheckInstanceResultResponseBodyColumns) SetKey(v string) *ListCheckInstanceResultResponseBodyColumns {
	s.Key = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyColumns) SetSearch(v bool) *ListCheckInstanceResultResponseBodyColumns {
	s.Search = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyColumns) SetSearchKey(v string) *ListCheckInstanceResultResponseBodyColumns {
	s.SearchKey = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyColumns) SetShowName(v string) *ListCheckInstanceResultResponseBodyColumns {
	s.ShowName = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyColumns) SetType(v string) *ListCheckInstanceResultResponseBodyColumns {
	s.Type = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyColumns) Validate() error {
	return dara.Validate(s)
}

type ListCheckInstanceResultResponseBodyColumnsGrids struct {
	// The search condition.
	//
	// example:
	//
	// RegionIdShow
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The display name of the search condition.
	//
	// example:
	//
	// Region
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// The format of the check result for the instance. Valid values:
	//
	// 	- **text**
	//
	// 	- **link**
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCheckInstanceResultResponseBodyColumnsGrids) String() string {
	return dara.Prettify(s)
}

func (s ListCheckInstanceResultResponseBodyColumnsGrids) GoString() string {
	return s.String()
}

func (s *ListCheckInstanceResultResponseBodyColumnsGrids) GetKey() *string {
	return s.Key
}

func (s *ListCheckInstanceResultResponseBodyColumnsGrids) GetShowName() *string {
	return s.ShowName
}

func (s *ListCheckInstanceResultResponseBodyColumnsGrids) GetType() *string {
	return s.Type
}

func (s *ListCheckInstanceResultResponseBodyColumnsGrids) SetKey(v string) *ListCheckInstanceResultResponseBodyColumnsGrids {
	s.Key = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyColumnsGrids) SetShowName(v string) *ListCheckInstanceResultResponseBodyColumnsGrids {
	s.ShowName = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyColumnsGrids) SetType(v string) *ListCheckInstanceResultResponseBodyColumnsGrids {
	s.Type = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyColumnsGrids) Validate() error {
	return dara.Validate(s)
}

type ListCheckInstanceResultResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCheckInstanceResultResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCheckInstanceResultResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListCheckInstanceResultResponseBodyPageInfo) GetCount() *string {
	return s.Count
}

func (s *ListCheckInstanceResultResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCheckInstanceResultResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCheckInstanceResultResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCheckInstanceResultResponseBodyPageInfo) SetCount(v string) *ListCheckInstanceResultResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyPageInfo) SetCurrentPage(v int32) *ListCheckInstanceResultResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyPageInfo) SetPageSize(v int32) *ListCheckInstanceResultResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyPageInfo) SetTotalCount(v int32) *ListCheckInstanceResultResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCheckInstanceResultResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
