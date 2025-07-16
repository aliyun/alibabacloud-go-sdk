// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProcessDefinitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFormUuid(v string) *GetProcessDefinitionResponseBody
	GetFormUuid() *string
	SetOriginator(v *GetProcessDefinitionResponseBodyOriginator) *GetProcessDefinitionResponseBody
	GetOriginator() *GetProcessDefinitionResponseBodyOriginator
	SetOutResult(v string) *GetProcessDefinitionResponseBody
	GetOutResult() *string
	SetOwners(v []*GetProcessDefinitionResponseBodyOwners) *GetProcessDefinitionResponseBody
	GetOwners() []*GetProcessDefinitionResponseBodyOwners
	SetProcessId(v string) *GetProcessDefinitionResponseBody
	GetProcessId() *string
	SetProcessInstanceId(v string) *GetProcessDefinitionResponseBody
	GetProcessInstanceId() *string
	SetRequestId(v string) *GetProcessDefinitionResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetProcessDefinitionResponseBody
	GetStatus() *string
	SetTasks(v []*GetProcessDefinitionResponseBodyTasks) *GetProcessDefinitionResponseBody
	GetTasks() []*GetProcessDefinitionResponseBodyTasks
	SetTitle(v string) *GetProcessDefinitionResponseBody
	GetTitle() *string
	SetVariables(v map[string]interface{}) *GetProcessDefinitionResponseBody
	GetVariables() map[string]interface{}
	SetVendorRequestId(v string) *GetProcessDefinitionResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetProcessDefinitionResponseBody
	GetVendorType() *string
}

type GetProcessDefinitionResponseBody struct {
	// example:
	//
	// FORM-EF6Y4xxx
	FormUuid   *string                                     `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	Originator *GetProcessDefinitionResponseBodyOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// example:
	//
	// agree
	OutResult *string                                   `json:"outResult,omitempty" xml:"outResult,omitempty"`
	Owners    []*GetProcessDefinitionResponseBodyOwners `json:"owners,omitempty" xml:"owners,omitempty" type:"Repeated"`
	// example:
	//
	// proc-123
	ProcessId *string `json:"processId,omitempty" xml:"processId,omitempty"`
	// example:
	//
	// f30233fb-72e1-4xxx
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// running
	Status *string                                  `json:"status,omitempty" xml:"status,omitempty"`
	Tasks  []*GetProcessDefinitionResponseBodyTasks `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
	// example:
	//
	// 李四发起的请购单
	Title     *string                `json:"title,omitempty" xml:"title,omitempty"`
	Variables map[string]interface{} `json:"variables,omitempty" xml:"variables,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetProcessDefinitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProcessDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponseBody) GetFormUuid() *string {
	return s.FormUuid
}

func (s *GetProcessDefinitionResponseBody) GetOriginator() *GetProcessDefinitionResponseBodyOriginator {
	return s.Originator
}

func (s *GetProcessDefinitionResponseBody) GetOutResult() *string {
	return s.OutResult
}

func (s *GetProcessDefinitionResponseBody) GetOwners() []*GetProcessDefinitionResponseBodyOwners {
	return s.Owners
}

func (s *GetProcessDefinitionResponseBody) GetProcessId() *string {
	return s.ProcessId
}

func (s *GetProcessDefinitionResponseBody) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *GetProcessDefinitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProcessDefinitionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetProcessDefinitionResponseBody) GetTasks() []*GetProcessDefinitionResponseBodyTasks {
	return s.Tasks
}

func (s *GetProcessDefinitionResponseBody) GetTitle() *string {
	return s.Title
}

func (s *GetProcessDefinitionResponseBody) GetVariables() map[string]interface{} {
	return s.Variables
}

func (s *GetProcessDefinitionResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetProcessDefinitionResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetProcessDefinitionResponseBody) SetFormUuid(v string) *GetProcessDefinitionResponseBody {
	s.FormUuid = &v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetOriginator(v *GetProcessDefinitionResponseBodyOriginator) *GetProcessDefinitionResponseBody {
	s.Originator = v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetOutResult(v string) *GetProcessDefinitionResponseBody {
	s.OutResult = &v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetOwners(v []*GetProcessDefinitionResponseBodyOwners) *GetProcessDefinitionResponseBody {
	s.Owners = v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetProcessId(v string) *GetProcessDefinitionResponseBody {
	s.ProcessId = &v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetProcessInstanceId(v string) *GetProcessDefinitionResponseBody {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetRequestId(v string) *GetProcessDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetStatus(v string) *GetProcessDefinitionResponseBody {
	s.Status = &v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetTasks(v []*GetProcessDefinitionResponseBodyTasks) *GetProcessDefinitionResponseBody {
	s.Tasks = v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetTitle(v string) *GetProcessDefinitionResponseBody {
	s.Title = &v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetVariables(v map[string]interface{}) *GetProcessDefinitionResponseBody {
	s.Variables = v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetVendorRequestId(v string) *GetProcessDefinitionResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetVendorType(v string) *GetProcessDefinitionResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetProcessDefinitionResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetProcessDefinitionResponseBodyOriginator struct {
	// example:
	//
	// 开发部成立于2000年
	DepartmentDescription *string `json:"DepartmentDescription,omitempty" xml:"DepartmentDescription,omitempty"`
	// example:
	//
	// ZhangSan
	DisplayEnName *string `json:"DisplayEnName,omitempty" xml:"DisplayEnName,omitempty"`
	// example:
	//
	// 测试应用
	DisplayName           *string                                                            `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	MasterDataDepartments []*GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments `json:"MasterDataDepartments,omitempty" xml:"MasterDataDepartments,omitempty" type:"Repeated"`
	// example:
	//
	// o-YDJKINSxxx
	OrderNumber *string `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	// example:
	//
	// https://abc.com/a.png
	PersonalPhoto *string `json:"PersonalPhoto,omitempty" xml:"PersonalPhoto,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// wang123
	TbWang *string `json:"TbWang,omitempty" xml:"TbWang,omitempty"`
	// example:
	//
	// manager123
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 张三
	UserInfo *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetProcessDefinitionResponseBodyOriginator) String() string {
	return dara.Prettify(s)
}

func (s GetProcessDefinitionResponseBodyOriginator) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponseBodyOriginator) GetDepartmentDescription() *string {
	return s.DepartmentDescription
}

func (s *GetProcessDefinitionResponseBodyOriginator) GetDisplayEnName() *string {
	return s.DisplayEnName
}

func (s *GetProcessDefinitionResponseBodyOriginator) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetProcessDefinitionResponseBodyOriginator) GetMasterDataDepartments() []*GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments {
	return s.MasterDataDepartments
}

func (s *GetProcessDefinitionResponseBodyOriginator) GetOrderNumber() *string {
	return s.OrderNumber
}

func (s *GetProcessDefinitionResponseBodyOriginator) GetPersonalPhoto() *string {
	return s.PersonalPhoto
}

func (s *GetProcessDefinitionResponseBodyOriginator) GetStatus() *string {
	return s.Status
}

func (s *GetProcessDefinitionResponseBodyOriginator) GetTbWang() *string {
	return s.TbWang
}

func (s *GetProcessDefinitionResponseBodyOriginator) GetUserId() *string {
	return s.UserId
}

func (s *GetProcessDefinitionResponseBodyOriginator) GetUserInfo() *string {
	return s.UserInfo
}

func (s *GetProcessDefinitionResponseBodyOriginator) SetDepartmentDescription(v string) *GetProcessDefinitionResponseBodyOriginator {
	s.DepartmentDescription = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginator) SetDisplayEnName(v string) *GetProcessDefinitionResponseBodyOriginator {
	s.DisplayEnName = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginator) SetDisplayName(v string) *GetProcessDefinitionResponseBodyOriginator {
	s.DisplayName = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginator) SetMasterDataDepartments(v []*GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) *GetProcessDefinitionResponseBodyOriginator {
	s.MasterDataDepartments = v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginator) SetOrderNumber(v string) *GetProcessDefinitionResponseBodyOriginator {
	s.OrderNumber = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginator) SetPersonalPhoto(v string) *GetProcessDefinitionResponseBodyOriginator {
	s.PersonalPhoto = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginator) SetStatus(v string) *GetProcessDefinitionResponseBodyOriginator {
	s.Status = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginator) SetTbWang(v string) *GetProcessDefinitionResponseBodyOriginator {
	s.TbWang = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginator) SetUserId(v string) *GetProcessDefinitionResponseBodyOriginator {
	s.UserId = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginator) SetUserInfo(v string) *GetProcessDefinitionResponseBodyOriginator {
	s.UserInfo = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginator) Validate() error {
	return dara.Validate(s)
}

type GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments struct {
	// example:
	//
	// 开发部
	DeptName *string `json:"DeptName,omitempty" xml:"DeptName,omitempty"`
	// example:
	//
	// development department
	DeptNameInEnglish *string `json:"DeptNameInEnglish,omitempty" xml:"DeptNameInEnglish,omitempty"`
	// example:
	//
	// develop-A
	DeptNo *string `json:"DeptNo,omitempty" xml:"DeptNo,omitempty"`
	// example:
	//
	// 总部-开发部
	DeptPath *string `json:"DeptPath,omitempty" xml:"DeptPath,omitempty"`
	// example:
	//
	// xxafafaf
	HumanSourceGroupOrderNumber *string `json:"HumanSourceGroupOrderNumber,omitempty" xml:"HumanSourceGroupOrderNumber,omitempty"`
	// example:
	//
	// 123311221
	HumanSourceGroupWorkNo *string `json:"HumanSourceGroupWorkNo,omitempty" xml:"HumanSourceGroupWorkNo,omitempty"`
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1732245789
	MasterWorkNo *string `json:"MasterWorkNo,omitempty" xml:"MasterWorkNo,omitempty"`
}

func (s GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) String() string {
	return dara.Prettify(s)
}

func (s GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) GetDeptName() *string {
	return s.DeptName
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) GetDeptNameInEnglish() *string {
	return s.DeptNameInEnglish
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) GetDeptNo() *string {
	return s.DeptNo
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) GetDeptPath() *string {
	return s.DeptPath
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) GetHumanSourceGroupOrderNumber() *string {
	return s.HumanSourceGroupOrderNumber
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) GetHumanSourceGroupWorkNo() *string {
	return s.HumanSourceGroupWorkNo
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) GetId() *int64 {
	return s.Id
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) GetMasterWorkNo() *string {
	return s.MasterWorkNo
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) SetDeptName(v string) *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments {
	s.DeptName = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) SetDeptNameInEnglish(v string) *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments {
	s.DeptNameInEnglish = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) SetDeptNo(v string) *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments {
	s.DeptNo = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) SetDeptPath(v string) *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments {
	s.DeptPath = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) SetHumanSourceGroupOrderNumber(v string) *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments {
	s.HumanSourceGroupOrderNumber = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) SetHumanSourceGroupWorkNo(v string) *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments {
	s.HumanSourceGroupWorkNo = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) SetId(v int64) *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments {
	s.Id = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) SetMasterWorkNo(v string) *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments {
	s.MasterWorkNo = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) Validate() error {
	return dara.Validate(s)
}

type GetProcessDefinitionResponseBodyOwners struct {
	// example:
	//
	// 开发部成立于2000年
	DepartmentDescription *string `json:"DepartmentDescription,omitempty" xml:"DepartmentDescription,omitempty"`
	// example:
	//
	// ZhangSan
	DisplayEnName *string `json:"DisplayEnName,omitempty" xml:"DisplayEnName,omitempty"`
	// example:
	//
	// 测试应用
	DisplayName           *string                                                        `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	MasterDataDepartments []*GetProcessDefinitionResponseBodyOwnersMasterDataDepartments `json:"MasterDataDepartments,omitempty" xml:"MasterDataDepartments,omitempty" type:"Repeated"`
	// example:
	//
	// o-YDJKINSxxx
	OrderNumber *string `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	// example:
	//
	// https://abc.com/a.png
	PersonalPhoto *string `json:"PersonalPhoto,omitempty" xml:"PersonalPhoto,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// wang123
	TbWang *string `json:"TbWang,omitempty" xml:"TbWang,omitempty"`
	// example:
	//
	// manager123
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 张三
	UserInfo *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetProcessDefinitionResponseBodyOwners) String() string {
	return dara.Prettify(s)
}

func (s GetProcessDefinitionResponseBodyOwners) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponseBodyOwners) GetDepartmentDescription() *string {
	return s.DepartmentDescription
}

func (s *GetProcessDefinitionResponseBodyOwners) GetDisplayEnName() *string {
	return s.DisplayEnName
}

func (s *GetProcessDefinitionResponseBodyOwners) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetProcessDefinitionResponseBodyOwners) GetMasterDataDepartments() []*GetProcessDefinitionResponseBodyOwnersMasterDataDepartments {
	return s.MasterDataDepartments
}

func (s *GetProcessDefinitionResponseBodyOwners) GetOrderNumber() *string {
	return s.OrderNumber
}

func (s *GetProcessDefinitionResponseBodyOwners) GetPersonalPhoto() *string {
	return s.PersonalPhoto
}

func (s *GetProcessDefinitionResponseBodyOwners) GetStatus() *string {
	return s.Status
}

func (s *GetProcessDefinitionResponseBodyOwners) GetTbWang() *string {
	return s.TbWang
}

func (s *GetProcessDefinitionResponseBodyOwners) GetUserId() *string {
	return s.UserId
}

func (s *GetProcessDefinitionResponseBodyOwners) GetUserInfo() *string {
	return s.UserInfo
}

func (s *GetProcessDefinitionResponseBodyOwners) SetDepartmentDescription(v string) *GetProcessDefinitionResponseBodyOwners {
	s.DepartmentDescription = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwners) SetDisplayEnName(v string) *GetProcessDefinitionResponseBodyOwners {
	s.DisplayEnName = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwners) SetDisplayName(v string) *GetProcessDefinitionResponseBodyOwners {
	s.DisplayName = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwners) SetMasterDataDepartments(v []*GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) *GetProcessDefinitionResponseBodyOwners {
	s.MasterDataDepartments = v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwners) SetOrderNumber(v string) *GetProcessDefinitionResponseBodyOwners {
	s.OrderNumber = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwners) SetPersonalPhoto(v string) *GetProcessDefinitionResponseBodyOwners {
	s.PersonalPhoto = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwners) SetStatus(v string) *GetProcessDefinitionResponseBodyOwners {
	s.Status = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwners) SetTbWang(v string) *GetProcessDefinitionResponseBodyOwners {
	s.TbWang = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwners) SetUserId(v string) *GetProcessDefinitionResponseBodyOwners {
	s.UserId = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwners) SetUserInfo(v string) *GetProcessDefinitionResponseBodyOwners {
	s.UserInfo = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwners) Validate() error {
	return dara.Validate(s)
}

type GetProcessDefinitionResponseBodyOwnersMasterDataDepartments struct {
	// example:
	//
	// 开发部
	DeptName *string `json:"DeptName,omitempty" xml:"DeptName,omitempty"`
	// example:
	//
	// development department
	DeptNameInEnglish *string `json:"DeptNameInEnglish,omitempty" xml:"DeptNameInEnglish,omitempty"`
	// example:
	//
	// develop-A
	DeptNo *string `json:"DeptNo,omitempty" xml:"DeptNo,omitempty"`
	// example:
	//
	// 总部-开发部
	DeptPath *string `json:"DeptPath,omitempty" xml:"DeptPath,omitempty"`
	// example:
	//
	// xxafafaf
	HumanSourceGroupOrderNumber *string `json:"HumanSourceGroupOrderNumber,omitempty" xml:"HumanSourceGroupOrderNumber,omitempty"`
	// example:
	//
	// 123311221
	HumanSourceGroupWorkNo *string `json:"HumanSourceGroupWorkNo,omitempty" xml:"HumanSourceGroupWorkNo,omitempty"`
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1732245789
	MasterWorkNo *string `json:"MasterWorkNo,omitempty" xml:"MasterWorkNo,omitempty"`
}

func (s GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) String() string {
	return dara.Prettify(s)
}

func (s GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) GetDeptName() *string {
	return s.DeptName
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) GetDeptNameInEnglish() *string {
	return s.DeptNameInEnglish
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) GetDeptNo() *string {
	return s.DeptNo
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) GetDeptPath() *string {
	return s.DeptPath
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) GetHumanSourceGroupOrderNumber() *string {
	return s.HumanSourceGroupOrderNumber
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) GetHumanSourceGroupWorkNo() *string {
	return s.HumanSourceGroupWorkNo
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) GetId() *int64 {
	return s.Id
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) GetMasterWorkNo() *string {
	return s.MasterWorkNo
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) SetDeptName(v string) *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments {
	s.DeptName = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) SetDeptNameInEnglish(v string) *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments {
	s.DeptNameInEnglish = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) SetDeptNo(v string) *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments {
	s.DeptNo = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) SetDeptPath(v string) *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments {
	s.DeptPath = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) SetHumanSourceGroupOrderNumber(v string) *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments {
	s.HumanSourceGroupOrderNumber = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) SetHumanSourceGroupWorkNo(v string) *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments {
	s.HumanSourceGroupWorkNo = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) SetId(v int64) *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments {
	s.Id = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) SetMasterWorkNo(v string) *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments {
	s.MasterWorkNo = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) Validate() error {
	return dara.Validate(s)
}

type GetProcessDefinitionResponseBodyTasks struct {
	// example:
	//
	// manager123
	ActionerId *string                                        `json:"ActionerId,omitempty" xml:"ActionerId,omitempty"`
	Activity   *GetProcessDefinitionResponseBodyTasksActivity `json:"Activity,omitempty" xml:"Activity,omitempty" type:"Struct"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 792
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetProcessDefinitionResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s GetProcessDefinitionResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponseBodyTasks) GetActionerId() *string {
	return s.ActionerId
}

func (s *GetProcessDefinitionResponseBodyTasks) GetActivity() *GetProcessDefinitionResponseBodyTasksActivity {
	return s.Activity
}

func (s *GetProcessDefinitionResponseBodyTasks) GetStatus() *string {
	return s.Status
}

func (s *GetProcessDefinitionResponseBodyTasks) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetProcessDefinitionResponseBodyTasks) SetActionerId(v string) *GetProcessDefinitionResponseBodyTasks {
	s.ActionerId = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyTasks) SetActivity(v *GetProcessDefinitionResponseBodyTasksActivity) *GetProcessDefinitionResponseBodyTasks {
	s.Activity = v
	return s
}

func (s *GetProcessDefinitionResponseBodyTasks) SetStatus(v string) *GetProcessDefinitionResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyTasks) SetTaskId(v int64) *GetProcessDefinitionResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}

type GetProcessDefinitionResponseBodyTasksActivity struct {
	// example:
	//
	// act-xxaanfaf
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// running
	ActivityInstanceStatus *string `json:"ActivityInstanceStatus,omitempty" xml:"ActivityInstanceStatus,omitempty"`
	// example:
	//
	// activity-124
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// example:
	//
	// redirect task
	ActivityNameInEnglish *string `json:"ActivityNameInEnglish,omitempty" xml:"ActivityNameInEnglish,omitempty"`
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetProcessDefinitionResponseBodyTasksActivity) String() string {
	return dara.Prettify(s)
}

func (s GetProcessDefinitionResponseBodyTasksActivity) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponseBodyTasksActivity) GetActivityId() *string {
	return s.ActivityId
}

func (s *GetProcessDefinitionResponseBodyTasksActivity) GetActivityInstanceStatus() *string {
	return s.ActivityInstanceStatus
}

func (s *GetProcessDefinitionResponseBodyTasksActivity) GetActivityName() *string {
	return s.ActivityName
}

func (s *GetProcessDefinitionResponseBodyTasksActivity) GetActivityNameInEnglish() *string {
	return s.ActivityNameInEnglish
}

func (s *GetProcessDefinitionResponseBodyTasksActivity) GetId() *int64 {
	return s.Id
}

func (s *GetProcessDefinitionResponseBodyTasksActivity) SetActivityId(v string) *GetProcessDefinitionResponseBodyTasksActivity {
	s.ActivityId = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyTasksActivity) SetActivityInstanceStatus(v string) *GetProcessDefinitionResponseBodyTasksActivity {
	s.ActivityInstanceStatus = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyTasksActivity) SetActivityName(v string) *GetProcessDefinitionResponseBodyTasksActivity {
	s.ActivityName = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyTasksActivity) SetActivityNameInEnglish(v string) *GetProcessDefinitionResponseBodyTasksActivity {
	s.ActivityNameInEnglish = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyTasksActivity) SetId(v int64) *GetProcessDefinitionResponseBodyTasksActivity {
	s.Id = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyTasksActivity) Validate() error {
	return dara.Validate(s)
}
