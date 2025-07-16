// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMeCorpSubmissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetMeCorpSubmissionResponseBodyData) *GetMeCorpSubmissionResponseBody
	GetData() []*GetMeCorpSubmissionResponseBodyData
	SetPageNumber(v int64) *GetMeCorpSubmissionResponseBody
	GetPageNumber() *int64
	SetRequestId(v string) *GetMeCorpSubmissionResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *GetMeCorpSubmissionResponseBody
	GetTotalCount() *int64
	SetVendorRequestId(v string) *GetMeCorpSubmissionResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetMeCorpSubmissionResponseBody
	GetVendorType() *string
}

type GetMeCorpSubmissionResponseBody struct {
	Data []*GetMeCorpSubmissionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetMeCorpSubmissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMeCorpSubmissionResponseBody) GoString() string {
	return s.String()
}

func (s *GetMeCorpSubmissionResponseBody) GetData() []*GetMeCorpSubmissionResponseBodyData {
	return s.Data
}

func (s *GetMeCorpSubmissionResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetMeCorpSubmissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMeCorpSubmissionResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetMeCorpSubmissionResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetMeCorpSubmissionResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetMeCorpSubmissionResponseBody) SetData(v []*GetMeCorpSubmissionResponseBodyData) *GetMeCorpSubmissionResponseBody {
	s.Data = v
	return s
}

func (s *GetMeCorpSubmissionResponseBody) SetPageNumber(v int64) *GetMeCorpSubmissionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBody) SetRequestId(v string) *GetMeCorpSubmissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBody) SetTotalCount(v int64) *GetMeCorpSubmissionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBody) SetVendorRequestId(v string) *GetMeCorpSubmissionResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBody) SetVendorType(v string) *GetMeCorpSubmissionResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMeCorpSubmissionResponseBodyData struct {
	Actioner     []*GetMeCorpSubmissionResponseBodyDataActioner `json:"Actioner,omitempty" xml:"Actioner,omitempty" type:"Repeated"`
	ActionerId   []*string                                      `json:"ActionerId,omitempty" xml:"ActionerId,omitempty" type:"Repeated"`
	ActionerName []*string                                      `json:"ActionerName,omitempty" xml:"ActionerName,omitempty" type:"Repeated"`
	// example:
	//
	// APP_PBKT0xxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// 2021-01-01
	CreateTimeGMT            *string                                                        `json:"CreateTimeGMT,omitempty" xml:"CreateTimeGMT,omitempty"`
	CurrentActivityInstances []*GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances `json:"CurrentActivityInstances,omitempty" xml:"CurrentActivityInstances,omitempty" type:"Repeated"`
	DataMap                  map[string]interface{}                                         `json:"DataMap,omitempty" xml:"DataMap,omitempty"`
	// example:
	//
	// edit
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// 2021-01-01
	FinishTimeGMT *string `json:"FinishTimeGMT,omitempty" xml:"FinishTimeGMT,omitempty"`
	// example:
	//
	// FINST-NJYJxxx
	FormInstanceId *string `json:"FormInstanceId,omitempty" xml:"FormInstanceId,omitempty"`
	// example:
	//
	// FORM-EF6xxx
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// example:
	//
	// 符合宜搭表单实例格式的json数据
	InstanceValue *string `json:"InstanceValue,omitempty" xml:"InstanceValue,omitempty"`
	// example:
	//
	// 2021-01-01
	ModifiedTimeGMT *string `json:"ModifiedTimeGMT,omitempty" xml:"ModifiedTimeGMT,omitempty"`
	// example:
	//
	// zhangsan@mediaId
	OriginatorAvatar *string `json:"OriginatorAvatar,omitempty" xml:"OriginatorAvatar,omitempty"`
	// example:
	//
	// 张三
	OriginatorDisplayName *string `json:"OriginatorDisplayName,omitempty" xml:"OriginatorDisplayName,omitempty"`
	// example:
	//
	// manager123
	OriginatorId *string `json:"OriginatorId,omitempty" xml:"OriginatorId,omitempty"`
	// example:
	//
	// 同意
	ProcessApprovedResult *string `json:"ProcessApprovedResult,omitempty" xml:"ProcessApprovedResult,omitempty"`
	// example:
	//
	// 通过
	ProcessApprovedResultText *string `json:"ProcessApprovedResultText,omitempty" xml:"ProcessApprovedResultText,omitempty"`
	// example:
	//
	// TPROC--X1Gxxx
	ProcessCode *string `json:"ProcessCode,omitempty" xml:"ProcessCode,omitempty"`
	// example:
	//
	// 52330
	ProcessId *int64 `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// example:
	//
	// f30233fb-72e1-xxx
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	// example:
	//
	// finished
	ProcessInstanceStatus *string `json:"ProcessInstanceStatus,omitempty" xml:"ProcessInstanceStatus,omitempty"`
	// example:
	//
	// 已同意
	ProcessInstanceStatusText *string `json:"ProcessInstanceStatusText,omitempty" xml:"ProcessInstanceStatusText,omitempty"`
	// example:
	//
	// 小红的单子
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// example:
	//
	// 小红发起的请购单
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetMeCorpSubmissionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMeCorpSubmissionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMeCorpSubmissionResponseBodyData) GetActioner() []*GetMeCorpSubmissionResponseBodyDataActioner {
	return s.Actioner
}

func (s *GetMeCorpSubmissionResponseBodyData) GetActionerId() []*string {
	return s.ActionerId
}

func (s *GetMeCorpSubmissionResponseBodyData) GetActionerName() []*string {
	return s.ActionerName
}

func (s *GetMeCorpSubmissionResponseBodyData) GetAppType() *string {
	return s.AppType
}

func (s *GetMeCorpSubmissionResponseBodyData) GetCreateTimeGMT() *string {
	return s.CreateTimeGMT
}

func (s *GetMeCorpSubmissionResponseBodyData) GetCurrentActivityInstances() []*GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances {
	return s.CurrentActivityInstances
}

func (s *GetMeCorpSubmissionResponseBodyData) GetDataMap() map[string]interface{} {
	return s.DataMap
}

func (s *GetMeCorpSubmissionResponseBodyData) GetDataType() *string {
	return s.DataType
}

func (s *GetMeCorpSubmissionResponseBodyData) GetFinishTimeGMT() *string {
	return s.FinishTimeGMT
}

func (s *GetMeCorpSubmissionResponseBodyData) GetFormInstanceId() *string {
	return s.FormInstanceId
}

func (s *GetMeCorpSubmissionResponseBodyData) GetFormUuid() *string {
	return s.FormUuid
}

func (s *GetMeCorpSubmissionResponseBodyData) GetInstanceValue() *string {
	return s.InstanceValue
}

func (s *GetMeCorpSubmissionResponseBodyData) GetModifiedTimeGMT() *string {
	return s.ModifiedTimeGMT
}

func (s *GetMeCorpSubmissionResponseBodyData) GetOriginatorAvatar() *string {
	return s.OriginatorAvatar
}

func (s *GetMeCorpSubmissionResponseBodyData) GetOriginatorDisplayName() *string {
	return s.OriginatorDisplayName
}

func (s *GetMeCorpSubmissionResponseBodyData) GetOriginatorId() *string {
	return s.OriginatorId
}

func (s *GetMeCorpSubmissionResponseBodyData) GetProcessApprovedResult() *string {
	return s.ProcessApprovedResult
}

func (s *GetMeCorpSubmissionResponseBodyData) GetProcessApprovedResultText() *string {
	return s.ProcessApprovedResultText
}

func (s *GetMeCorpSubmissionResponseBodyData) GetProcessCode() *string {
	return s.ProcessCode
}

func (s *GetMeCorpSubmissionResponseBodyData) GetProcessId() *int64 {
	return s.ProcessId
}

func (s *GetMeCorpSubmissionResponseBodyData) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *GetMeCorpSubmissionResponseBodyData) GetProcessInstanceStatus() *string {
	return s.ProcessInstanceStatus
}

func (s *GetMeCorpSubmissionResponseBodyData) GetProcessInstanceStatusText() *string {
	return s.ProcessInstanceStatusText
}

func (s *GetMeCorpSubmissionResponseBodyData) GetProcessName() *string {
	return s.ProcessName
}

func (s *GetMeCorpSubmissionResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetMeCorpSubmissionResponseBodyData) GetVersion() *int64 {
	return s.Version
}

func (s *GetMeCorpSubmissionResponseBodyData) SetActioner(v []*GetMeCorpSubmissionResponseBodyDataActioner) *GetMeCorpSubmissionResponseBodyData {
	s.Actioner = v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetActionerId(v []*string) *GetMeCorpSubmissionResponseBodyData {
	s.ActionerId = v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetActionerName(v []*string) *GetMeCorpSubmissionResponseBodyData {
	s.ActionerName = v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetAppType(v string) *GetMeCorpSubmissionResponseBodyData {
	s.AppType = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetCreateTimeGMT(v string) *GetMeCorpSubmissionResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetCurrentActivityInstances(v []*GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) *GetMeCorpSubmissionResponseBodyData {
	s.CurrentActivityInstances = v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetDataMap(v map[string]interface{}) *GetMeCorpSubmissionResponseBodyData {
	s.DataMap = v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetDataType(v string) *GetMeCorpSubmissionResponseBodyData {
	s.DataType = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetFinishTimeGMT(v string) *GetMeCorpSubmissionResponseBodyData {
	s.FinishTimeGMT = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetFormInstanceId(v string) *GetMeCorpSubmissionResponseBodyData {
	s.FormInstanceId = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetFormUuid(v string) *GetMeCorpSubmissionResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetInstanceValue(v string) *GetMeCorpSubmissionResponseBodyData {
	s.InstanceValue = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetModifiedTimeGMT(v string) *GetMeCorpSubmissionResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetOriginatorAvatar(v string) *GetMeCorpSubmissionResponseBodyData {
	s.OriginatorAvatar = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetOriginatorDisplayName(v string) *GetMeCorpSubmissionResponseBodyData {
	s.OriginatorDisplayName = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetOriginatorId(v string) *GetMeCorpSubmissionResponseBodyData {
	s.OriginatorId = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetProcessApprovedResult(v string) *GetMeCorpSubmissionResponseBodyData {
	s.ProcessApprovedResult = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetProcessApprovedResultText(v string) *GetMeCorpSubmissionResponseBodyData {
	s.ProcessApprovedResultText = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetProcessCode(v string) *GetMeCorpSubmissionResponseBodyData {
	s.ProcessCode = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetProcessId(v int64) *GetMeCorpSubmissionResponseBodyData {
	s.ProcessId = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetProcessInstanceId(v string) *GetMeCorpSubmissionResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetProcessInstanceStatus(v string) *GetMeCorpSubmissionResponseBodyData {
	s.ProcessInstanceStatus = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetProcessInstanceStatusText(v string) *GetMeCorpSubmissionResponseBodyData {
	s.ProcessInstanceStatusText = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetProcessName(v string) *GetMeCorpSubmissionResponseBodyData {
	s.ProcessName = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetTitle(v string) *GetMeCorpSubmissionResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetVersion(v int64) *GetMeCorpSubmissionResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetMeCorpSubmissionResponseBodyDataActioner struct {
	// example:
	//
	// 某研究部
	BuName *string `json:"BuName,omitempty" xml:"BuName,omitempty"`
	// example:
	//
	// abc@alimail.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 正式
	EmployeeType *string `json:"EmployeeType,omitempty" xml:"EmployeeType,omitempty"`
	// example:
	//
	// official
	EmployeeTypeInformation *string `json:"EmployeeTypeInformation,omitempty" xml:"EmployeeTypeInformation,omitempty"`
	// example:
	//
	// 123311221
	HumanResourceGroupWorkNumber *string `json:"HumanResourceGroupWorkNumber,omitempty" xml:"HumanResourceGroupWorkNumber,omitempty"`
	// example:
	//
	// true
	IsSystemAdmin *bool `json:"IsSystemAdmin,omitempty" xml:"IsSystemAdmin,omitempty"`
	// example:
	//
	// P7
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 请购单
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 与心
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
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
	// https://oss/zhangsan.png
	PersonalPhotoUrl *string `json:"PersonalPhotoUrl,omitempty" xml:"PersonalPhotoUrl,omitempty"`
	// example:
	//
	// XIAOHONG
	PinyinNameAll *string `json:"PinyinNameAll,omitempty" xml:"PinyinNameAll,omitempty"`
	// example:
	//
	// xiaohong
	PinyinNickName *string `json:"PinyinNickName,omitempty" xml:"PinyinNickName,omitempty"`
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// manager123
	SuperUserId *string `json:"SuperUserId,omitempty" xml:"SuperUserId,omitempty"`
	// example:
	//
	// wang123
	TbWang *string `json:"TbWang,omitempty" xml:"TbWang,omitempty"`
	// example:
	//
	// manager123
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetMeCorpSubmissionResponseBodyDataActioner) String() string {
	return dara.Prettify(s)
}

func (s GetMeCorpSubmissionResponseBodyDataActioner) GoString() string {
	return s.String()
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) GetBuName() *string {
	return s.BuName
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) GetEmail() *string {
	return s.Email
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) GetEmployeeType() *string {
	return s.EmployeeType
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) GetEmployeeTypeInformation() *string {
	return s.EmployeeTypeInformation
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) GetHumanResourceGroupWorkNumber() *string {
	return s.HumanResourceGroupWorkNumber
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) GetIsSystemAdmin() *bool {
	return s.IsSystemAdmin
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) GetLevel() *string {
	return s.Level
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) GetName() *string {
	return s.Name
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) GetNickName() *string {
	return s.NickName
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) GetOrderNumber() *string {
	return s.OrderNumber
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) GetPersonalPhoto() *string {
	return s.PersonalPhoto
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) GetPersonalPhotoUrl() *string {
	return s.PersonalPhotoUrl
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) GetPinyinNameAll() *string {
	return s.PinyinNameAll
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) GetPinyinNickName() *string {
	return s.PinyinNickName
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) GetState() *string {
	return s.State
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) GetSuperUserId() *string {
	return s.SuperUserId
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) GetTbWang() *string {
	return s.TbWang
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) GetUserId() *string {
	return s.UserId
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetBuName(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.BuName = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetEmail(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.Email = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetEmployeeType(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.EmployeeType = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetEmployeeTypeInformation(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.EmployeeTypeInformation = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetHumanResourceGroupWorkNumber(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.HumanResourceGroupWorkNumber = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetIsSystemAdmin(v bool) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.IsSystemAdmin = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetLevel(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.Level = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetName(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.Name = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetNickName(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.NickName = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetOrderNumber(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.OrderNumber = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetPersonalPhoto(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.PersonalPhoto = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetPersonalPhotoUrl(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.PersonalPhotoUrl = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetPinyinNameAll(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.PinyinNameAll = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetPinyinNickName(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.PinyinNickName = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetState(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.State = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetSuperUserId(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.SuperUserId = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetTbWang(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.TbWang = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetUserId(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.UserId = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) Validate() error {
	return dara.Validate(s)
}

type GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances struct {
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
	ActivityNameEn *string `json:"ActivityNameEn,omitempty" xml:"ActivityNameEn,omitempty"`
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) String() string {
	return dara.Prettify(s)
}

func (s GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) GoString() string {
	return s.String()
}

func (s *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) GetActivityId() *string {
	return s.ActivityId
}

func (s *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) GetActivityInstanceStatus() *string {
	return s.ActivityInstanceStatus
}

func (s *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) GetActivityName() *string {
	return s.ActivityName
}

func (s *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) GetActivityNameEn() *string {
	return s.ActivityNameEn
}

func (s *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) GetId() *int64 {
	return s.Id
}

func (s *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) SetActivityId(v string) *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances {
	s.ActivityId = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) SetActivityInstanceStatus(v string) *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances {
	s.ActivityInstanceStatus = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) SetActivityName(v string) *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances {
	s.ActivityName = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) SetActivityNameEn(v string) *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances {
	s.ActivityNameEn = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) SetId(v int64) *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances {
	s.Id = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) Validate() error {
	return dara.Validate(s)
}
