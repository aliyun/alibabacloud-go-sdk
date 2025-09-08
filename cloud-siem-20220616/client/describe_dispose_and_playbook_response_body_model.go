// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisposeAndPlaybookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeDisposeAndPlaybookResponseBody
	GetCode() *int32
	SetData(v *DescribeDisposeAndPlaybookResponseBodyData) *DescribeDisposeAndPlaybookResponseBody
	GetData() *DescribeDisposeAndPlaybookResponseBodyData
	SetMessage(v string) *DescribeDisposeAndPlaybookResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDisposeAndPlaybookResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDisposeAndPlaybookResponseBody
	GetSuccess() *bool
}

type DescribeDisposeAndPlaybookResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeDisposeAndPlaybookResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDisposeAndPlaybookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisposeAndPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeDisposeAndPlaybookResponseBody) GetData() *DescribeDisposeAndPlaybookResponseBodyData {
	return s.Data
}

func (s *DescribeDisposeAndPlaybookResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDisposeAndPlaybookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDisposeAndPlaybookResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDisposeAndPlaybookResponseBody) SetCode(v int32) *DescribeDisposeAndPlaybookResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBody) SetData(v *DescribeDisposeAndPlaybookResponseBodyData) *DescribeDisposeAndPlaybookResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBody) SetMessage(v string) *DescribeDisposeAndPlaybookResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBody) SetRequestId(v string) *DescribeDisposeAndPlaybookResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBody) SetSuccess(v bool) *DescribeDisposeAndPlaybookResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDisposeAndPlaybookResponseBodyData struct {
	// The pagination information.
	PageInfo *DescribeDisposeAndPlaybookResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*DescribeDisposeAndPlaybookResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s DescribeDisposeAndPlaybookResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisposeAndPlaybookResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookResponseBodyData) GetPageInfo() *DescribeDisposeAndPlaybookResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *DescribeDisposeAndPlaybookResponseBodyData) GetResponseData() []*DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	return s.ResponseData
}

func (s *DescribeDisposeAndPlaybookResponseBodyData) SetPageInfo(v *DescribeDisposeAndPlaybookResponseBodyDataPageInfo) *DescribeDisposeAndPlaybookResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyData) SetResponseData(v []*DescribeDisposeAndPlaybookResponseBodyDataResponseData) *DescribeDisposeAndPlaybookResponseBodyData {
	s.ResponseData = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeDisposeAndPlaybookResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDisposeAndPlaybookResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisposeAndPlaybookResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataPageInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeDisposeAndPlaybookResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeDisposeAndPlaybookResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataPageInfo) SetTotalCount(v int64) *DescribeDisposeAndPlaybookResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeDisposeAndPlaybookResponseBodyDataResponseData struct {
	// The number of alerts that are associated with the entity.
	//
	// example:
	//
	// 1
	AlertNum *int32 `json:"AlertNum,omitempty" xml:"AlertNum,omitempty"`
	// The object for handling.
	//
	// example:
	//
	// 192.168.1.1
	Dispose *string `json:"Dispose,omitempty" xml:"Dispose,omitempty"`
	// The entity ID
	//
	// example:
	//
	// 12345
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The entity information.
	//
	// example:
	//
	// {"file_path": "c:/www/leixi.jsp","file_hash": "aa0ca926ad948cd820e0a3d9a18c09d0","host_uuid": "efed2cf7-0b77-45d9-a97b-d2cf246bcbb3","malware_type": "${aliyun.siem.sas.alert_tag.webshell}","host_name": "launch-advisor-20230531"}
	EntityInfo map[string]interface{} `json:"EntityInfo,omitempty" xml:"EntityInfo,omitempty"`
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The key-value pairs each of which consists of opcode and oplevel.
	//
	// example:
	//
	// 12345
	OpcodeMap map[string]*string `json:"OpcodeMap,omitempty" xml:"OpcodeMap,omitempty"`
	// The codes of the playbooks that are recommended for entity handling.
	//
	// example:
	//
	// [1,3]
	OpcodeSet []*string `json:"OpcodeSet,omitempty" xml:"OpcodeSet,omitempty" type:"Repeated"`
	// The playbooks that can handle the entity.
	PlaybookList []*DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList `json:"PlaybookList,omitempty" xml:"PlaybookList,omitempty" type:"Repeated"`
	// The IDs of the users who can handle objects.
	//
	// example:
	//
	// 176618589410****
	Scope []interface{} `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Repeated"`
}

func (s DescribeDisposeAndPlaybookResponseBodyDataResponseData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisposeAndPlaybookResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) GetAlertNum() *int32 {
	return s.AlertNum
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) GetDispose() *string {
	return s.Dispose
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) GetEntityId() *int64 {
	return s.EntityId
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) GetEntityInfo() map[string]interface{} {
	return s.EntityInfo
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) GetEntityType() *string {
	return s.EntityType
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) GetOpcodeMap() map[string]*string {
	return s.OpcodeMap
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) GetOpcodeSet() []*string {
	return s.OpcodeSet
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) GetPlaybookList() []*DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	return s.PlaybookList
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) GetScope() []interface{} {
	return s.Scope
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetAlertNum(v int32) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.AlertNum = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetDispose(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.Dispose = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetEntityId(v int64) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.EntityId = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetEntityInfo(v map[string]interface{}) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.EntityInfo = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetEntityType(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.EntityType = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetOpcodeMap(v map[string]*string) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.OpcodeMap = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetOpcodeSet(v []*string) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.OpcodeSet = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetPlaybookList(v []*DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.PlaybookList = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetScope(v []interface{}) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.Scope = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) Validate() error {
	return dara.Validate(s)
}

type DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList struct {
	Available *string `json:"Available,omitempty" xml:"Available,omitempty"`
	// The playbook description.
	//
	// example:
	//
	// WafBlockIP
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the playbook.
	//
	// example:
	//
	// WafBlockIP
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The playbook name, which is the unique identifier of the playbook.
	//
	// example:
	//
	// kill_process_isolate_file
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The opcode of the playbook, which corresponds to the opcode of the playbook recommended for entity handling.
	//
	// example:
	//
	// 7
	OpCode *string `json:"OpCode,omitempty" xml:"OpCode,omitempty"`
	// Indicates whether quick event handling is selected by default. Valid values:
	//
	// 	- 2: Quick event handling is selected.
	//
	// 	- 1: Quick event handling is displayed but not selected.
	//
	// example:
	//
	// 2
	OpLevel *string `json:"OpLevel,omitempty" xml:"OpLevel,omitempty"`
	// The playbook parameters and the corresponding properties.
	ParamConfig []interface{} `json:"ParamConfig,omitempty" xml:"ParamConfig,omitempty" type:"Repeated"`
	// The opcode configuration.
	//
	// example:
	//
	// {"opCode":"3"}
	TaskConfig *string `json:"TaskConfig,omitempty" xml:"TaskConfig,omitempty"`
	// example:
	//
	// kill_process_isolate_file
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// Indicates whether the playbook is intended for Web Application Firewall (WAF). Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	WafPlaybook *bool `json:"WafPlaybook,omitempty" xml:"WafPlaybook,omitempty"`
}

func (s DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) GetAvailable() *string {
	return s.Available
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) GetDescription() *string {
	return s.Description
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) GetName() *string {
	return s.Name
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) GetOpCode() *string {
	return s.OpCode
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) GetOpLevel() *string {
	return s.OpLevel
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) GetParamConfig() []interface{} {
	return s.ParamConfig
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) GetTaskConfig() *string {
	return s.TaskConfig
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) GetWafPlaybook() *bool {
	return s.WafPlaybook
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetAvailable(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.Available = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetDescription(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.Description = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetDisplayName(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.DisplayName = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetName(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.Name = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetOpCode(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.OpCode = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetOpLevel(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.OpLevel = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetParamConfig(v []interface{}) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.ParamConfig = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetTaskConfig(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.TaskConfig = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetUuid(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.Uuid = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetWafPlaybook(v bool) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.WafPlaybook = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) Validate() error {
	return dara.Validate(s)
}
