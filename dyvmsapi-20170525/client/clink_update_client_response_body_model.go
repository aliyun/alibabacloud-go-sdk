// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkUpdateClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkUpdateClientResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkUpdateClientResponseBody
	GetCode() *string
	SetData(v *ClinkUpdateClientResponseBodyData) *ClinkUpdateClientResponseBody
	GetData() *ClinkUpdateClientResponseBodyData
	SetMessage(v string) *ClinkUpdateClientResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkUpdateClientResponseBody
	GetRequestId() *string
}

type ClinkUpdateClientResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkUpdateClientResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkUpdateClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkUpdateClientResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkUpdateClientResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkUpdateClientResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkUpdateClientResponseBody) GetData() *ClinkUpdateClientResponseBodyData {
	return s.Data
}

func (s *ClinkUpdateClientResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkUpdateClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkUpdateClientResponseBody) SetAccessDeniedDetail(v string) *ClinkUpdateClientResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkUpdateClientResponseBody) SetCode(v string) *ClinkUpdateClientResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkUpdateClientResponseBody) SetData(v *ClinkUpdateClientResponseBodyData) *ClinkUpdateClientResponseBody {
	s.Data = v
	return s
}

func (s *ClinkUpdateClientResponseBody) SetMessage(v string) *ClinkUpdateClientResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkUpdateClientResponseBody) SetRequestId(v string) *ClinkUpdateClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkUpdateClientResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkUpdateClientResponseBodyData struct {
	Client *ClinkUpdateClientResponseBodyDataClient `json:"Client,omitempty" xml:"Client,omitempty" type:"Struct"`
	// 请求 id
	//
	// example:
	//
	// 示例值示例值
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
}

func (s ClinkUpdateClientResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkUpdateClientResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkUpdateClientResponseBodyData) GetClient() *ClinkUpdateClientResponseBodyDataClient {
	return s.Client
}

func (s *ClinkUpdateClientResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkUpdateClientResponseBodyData) SetClient(v *ClinkUpdateClientResponseBodyDataClient) *ClinkUpdateClientResponseBodyData {
	s.Client = v
	return s
}

func (s *ClinkUpdateClientResponseBodyData) SetClinkRequestId(v string) *ClinkUpdateClientResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyData) Validate() error {
	if s.Client != nil {
		if err := s.Client.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkUpdateClientResponseBodyDataClient struct {
	// 启用状态，0: 停用；1: 启用
	//
	// example:
	//
	// 1
	Active *int64 `json:"Active,omitempty" xml:"Active,omitempty"`
	// 所属区号
	//
	// example:
	//
	// 010
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// 号码类型，默认值为0，0: 号码；1: 动态号码组
	//
	// example:
	//
	// 0
	AssignType *int64 `json:"AssignType,omitempty" xml:"AssignType,omitempty"`
	// 可外显号码集合
	Clid []*string `json:"Clid,omitempty" xml:"Clid,omitempty" type:"Repeated"`
	// 给地区指定可外显号码
	ClidArea []*ClinkUpdateClientResponseBodyDataClientClidArea `json:"ClidArea,omitempty" xml:"ClidArea,omitempty" type:"Repeated"`
	// 缺省外显号码集合，当clidType（外显号码类型）值为3且clidRule（外显规则）值为1时，支持配置缺省外显
	ClidDefault []*string `json:"ClidDefault,omitempty" xml:"ClidDefault,omitempty" type:"Repeated"`
	// 外显规则，根据外显号码类型值来定义，1: 随机；2: 轮选 或 1: 动态；2: 自定义
	//
	// example:
	//
	// 1
	ClidRule *int64 `json:"ClidRule,omitempty" xml:"ClidRule,omitempty"`
	// 外显号码类型，0: 全部；2: 座席指定；3: 智能外显
	//
	// example:
	//
	// 0
	ClidType *int64 `json:"ClidType,omitempty" xml:"ClidType,omitempty"`
	// 云号码外呼开关，0-关，1-开
	//
	// example:
	//
	// 0
	CloudNumberEnabled *int64 `json:"CloudNumberEnabled,omitempty" xml:"CloudNumberEnabled,omitempty"`
	// 云号码四种呼叫模式；数组长度 为 4，依次对应云号码外呼的四种模式（实体卡、工作卡、两端呼、RTC)的开启状态；
	CloudNumberModes []*int64 `json:"CloudNumberModes,omitempty" xml:"CloudNumberModes,omitempty" type:"Repeated"`
	// 座席工号
	//
	// example:
	//
	// 1111
	Cno *int64 `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// CRM 编号，与第三方 CRM 系统对接时，可作为唯一标识
	//
	// example:
	//
	// 51
	CrmId *int64 `json:"CrmId,omitempty" xml:"CrmId,omitempty"`
	// 动态号码组名称，当assignType（号码类型）值为1（动态号码组）时为必填项
	//
	// example:
	//
	// 示例值示例值示例值
	DynamicTelGroupName *string `json:"DynamicTelGroupName,omitempty" xml:"DynamicTelGroupName,omitempty"`
	// 号码隐藏类型，0: 不隐藏；1: 隐藏规则与全局设置保持一致
	//
	// example:
	//
	// 0
	HiddenTel *int64 `json:"HiddenTel,omitempty" xml:"HiddenTel,omitempty"`
	// 座席名称
	//
	// example:
	//
	// 示例值示例值
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 座席密码
	//
	// example:
	//
	// xxx
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// permission
	Permission *ClinkUpdateClientResponseBodyDataClientPermission `json:"Permission,omitempty" xml:"Permission,omitempty" type:"Struct"`
	// 所属队列号集合
	Qnos []*string `json:"Qnos,omitempty" xml:"Qnos,omitempty" type:"Repeated"`
	// 轮选方式，1: 按天轮选；2: 按次轮选
	//
	// example:
	//
	// 1
	RecurrentselectionType *int64 `json:"RecurrentselectionType,omitempty" xml:"RecurrentselectionType,omitempty"`
	// 轮选值设置 1: 按天轮选，每 n 天外呼更换一次外显号码，可设置 1-30 天 2: 按次轮选，每 n 次外呼更换一次外显号码，可设置 1-30 次
	//
	// example:
	//
	// 1
	RecurrentselectionValue *int64 `json:"RecurrentselectionValue,omitempty" xml:"RecurrentselectionValue,omitempty"`
	// 座席角色，0: 普通座席；1: 班长座席
	//
	// example:
	//
	// 0
	Role *int64 `json:"Role,omitempty" xml:"Role,omitempty"`
	// 座席类型，1：全渠道、2：呼叫中心、3：在线客服
	//
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
	// 整理时长，座席进行外呼操作后的整理时间
	//
	// example:
	//
	// 79
	WrapupTime *int64 `json:"WrapupTime,omitempty" xml:"WrapupTime,omitempty"`
}

func (s ClinkUpdateClientResponseBodyDataClient) String() string {
	return dara.Prettify(s)
}

func (s ClinkUpdateClientResponseBodyDataClient) GoString() string {
	return s.String()
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetActive() *int64 {
	return s.Active
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetAreaCode() *string {
	return s.AreaCode
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetAssignType() *int64 {
	return s.AssignType
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetClid() []*string {
	return s.Clid
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetClidArea() []*ClinkUpdateClientResponseBodyDataClientClidArea {
	return s.ClidArea
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetClidDefault() []*string {
	return s.ClidDefault
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetClidRule() *int64 {
	return s.ClidRule
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetClidType() *int64 {
	return s.ClidType
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetCloudNumberEnabled() *int64 {
	return s.CloudNumberEnabled
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetCloudNumberModes() []*int64 {
	return s.CloudNumberModes
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetCno() *int64 {
	return s.Cno
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetCrmId() *int64 {
	return s.CrmId
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetDynamicTelGroupName() *string {
	return s.DynamicTelGroupName
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetHiddenTel() *int64 {
	return s.HiddenTel
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetName() *string {
	return s.Name
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetPassword() *string {
	return s.Password
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetPermission() *ClinkUpdateClientResponseBodyDataClientPermission {
	return s.Permission
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetQnos() []*string {
	return s.Qnos
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetRecurrentselectionType() *int64 {
	return s.RecurrentselectionType
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetRecurrentselectionValue() *int64 {
	return s.RecurrentselectionValue
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetRole() *int64 {
	return s.Role
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetType() *int64 {
	return s.Type
}

func (s *ClinkUpdateClientResponseBodyDataClient) GetWrapupTime() *int64 {
	return s.WrapupTime
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetActive(v int64) *ClinkUpdateClientResponseBodyDataClient {
	s.Active = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetAreaCode(v string) *ClinkUpdateClientResponseBodyDataClient {
	s.AreaCode = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetAssignType(v int64) *ClinkUpdateClientResponseBodyDataClient {
	s.AssignType = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetClid(v []*string) *ClinkUpdateClientResponseBodyDataClient {
	s.Clid = v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetClidArea(v []*ClinkUpdateClientResponseBodyDataClientClidArea) *ClinkUpdateClientResponseBodyDataClient {
	s.ClidArea = v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetClidDefault(v []*string) *ClinkUpdateClientResponseBodyDataClient {
	s.ClidDefault = v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetClidRule(v int64) *ClinkUpdateClientResponseBodyDataClient {
	s.ClidRule = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetClidType(v int64) *ClinkUpdateClientResponseBodyDataClient {
	s.ClidType = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetCloudNumberEnabled(v int64) *ClinkUpdateClientResponseBodyDataClient {
	s.CloudNumberEnabled = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetCloudNumberModes(v []*int64) *ClinkUpdateClientResponseBodyDataClient {
	s.CloudNumberModes = v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetCno(v int64) *ClinkUpdateClientResponseBodyDataClient {
	s.Cno = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetCrmId(v int64) *ClinkUpdateClientResponseBodyDataClient {
	s.CrmId = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetDynamicTelGroupName(v string) *ClinkUpdateClientResponseBodyDataClient {
	s.DynamicTelGroupName = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetHiddenTel(v int64) *ClinkUpdateClientResponseBodyDataClient {
	s.HiddenTel = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetName(v string) *ClinkUpdateClientResponseBodyDataClient {
	s.Name = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetPassword(v string) *ClinkUpdateClientResponseBodyDataClient {
	s.Password = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetPermission(v *ClinkUpdateClientResponseBodyDataClientPermission) *ClinkUpdateClientResponseBodyDataClient {
	s.Permission = v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetQnos(v []*string) *ClinkUpdateClientResponseBodyDataClient {
	s.Qnos = v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetRecurrentselectionType(v int64) *ClinkUpdateClientResponseBodyDataClient {
	s.RecurrentselectionType = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetRecurrentselectionValue(v int64) *ClinkUpdateClientResponseBodyDataClient {
	s.RecurrentselectionValue = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetRole(v int64) *ClinkUpdateClientResponseBodyDataClient {
	s.Role = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetType(v int64) *ClinkUpdateClientResponseBodyDataClient {
	s.Type = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) SetWrapupTime(v int64) *ClinkUpdateClientResponseBodyDataClient {
	s.WrapupTime = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClient) Validate() error {
	if s.ClidArea != nil {
		for _, item := range s.ClidArea {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Permission != nil {
		if err := s.Permission.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkUpdateClientResponseBodyDataClientClidArea struct {
	// 地区组名称
	//
	// example:
	//
	// 示例值示例值
	AreaGroupName *string `json:"AreaGroupName,omitempty" xml:"AreaGroupName,omitempty"`
	// 号码类型，默认值为0，0: 号码；1: 动态号码组
	//
	// example:
	//
	// 0
	AssignType *int64 `json:"AssignType,omitempty" xml:"AssignType,omitempty"`
	// 动态号码组名称，当assignType（号码类型）值为1（动态号码组）时为必填项
	//
	// example:
	//
	// 示例值示例值
	DynamicTelGroupName *string `json:"DynamicTelGroupName,omitempty" xml:"DynamicTelGroupName,omitempty"`
	// 外显号码
	ObClids []*string `json:"ObClids,omitempty" xml:"ObClids,omitempty" type:"Repeated"`
}

func (s ClinkUpdateClientResponseBodyDataClientClidArea) String() string {
	return dara.Prettify(s)
}

func (s ClinkUpdateClientResponseBodyDataClientClidArea) GoString() string {
	return s.String()
}

func (s *ClinkUpdateClientResponseBodyDataClientClidArea) GetAreaGroupName() *string {
	return s.AreaGroupName
}

func (s *ClinkUpdateClientResponseBodyDataClientClidArea) GetAssignType() *int64 {
	return s.AssignType
}

func (s *ClinkUpdateClientResponseBodyDataClientClidArea) GetDynamicTelGroupName() *string {
	return s.DynamicTelGroupName
}

func (s *ClinkUpdateClientResponseBodyDataClientClidArea) GetObClids() []*string {
	return s.ObClids
}

func (s *ClinkUpdateClientResponseBodyDataClientClidArea) SetAreaGroupName(v string) *ClinkUpdateClientResponseBodyDataClientClidArea {
	s.AreaGroupName = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClientClidArea) SetAssignType(v int64) *ClinkUpdateClientResponseBodyDataClientClidArea {
	s.AssignType = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClientClidArea) SetDynamicTelGroupName(v string) *ClinkUpdateClientResponseBodyDataClientClidArea {
	s.DynamicTelGroupName = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClientClidArea) SetObClids(v []*string) *ClinkUpdateClientResponseBodyDataClientClidArea {
	s.ObClids = v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClientClidArea) Validate() error {
	return dara.Validate(s)
}

type ClinkUpdateClientResponseBodyDataClientPermission struct {
	// 语音转写，0: 关闭；1: 呼入开启；2 外呼开启；3 全部开启；默认值为 0
	//
	// example:
	//
	// 1
	Asr *int64 `json:"Asr,omitempty" xml:"Asr,omitempty"`
	// 外呼权限，0: 关闭；1: 无限制；2: 国内长途；3: 国内本地，默认值为 1
	//
	// example:
	//
	// 1
	Call *int64 `json:"Call,omitempty" xml:"Call,omitempty"`
	// 通话记录查看权限，1: 全部；2: 所属队列；3:本座席，默认值为 1
	//
	// example:
	//
	// 1
	Cdr *int64 `json:"Cdr,omitempty" xml:"Cdr,omitempty"`
	// 在线客服查看会话记录权限 ，0：全部、1：所属队列、2：本座席； 默认值为 0
	//
	// example:
	//
	// 1
	Chat *int64 `json:"Chat,omitempty" xml:"Chat,omitempty"`
	// 其他数据查看权限：1：全部 2：所属员工组 3：自己 4:指定员工组 ,默认所属员工组
	//
	// example:
	//
	// 1
	OtherData *int64 `json:"OtherData,omitempty" xml:"OtherData,omitempty"`
	// 通话录音权限，0: 关闭；1: 呼入；2: 外呼；3: 全部，默认值为 3
	//
	// example:
	//
	// 3
	Record *int64 `json:"Record,omitempty" xml:"Record,omitempty"`
	// 录音试听下载权限，0: 关闭；1: 试听下载；2：试听，默认值 1
	//
	// example:
	//
	// 1
	RecordDownload *int64 `json:"RecordDownload,omitempty" xml:"RecordDownload,omitempty"`
	// 短信发送权限，0: 关闭；1: 开启，默认值为 0
	//
	// example:
	//
	// 1
	Sms *int64 `json:"Sms,omitempty" xml:"Sms,omitempty"`
	// 外呼任务查看权限，1：全部 3：自己，默认全部
	//
	// example:
	//
	// 1
	TaskInventory *int64 `json:"TaskInventory,omitempty" xml:"TaskInventory,omitempty"`
	// 通话转移/咨询权限，可选范围，0：全部，1：所属员工组, 默认所属员工组
	//
	// example:
	//
	// 1
	Transfer *int64 `json:"Transfer,omitempty" xml:"Transfer,omitempty"`
}

func (s ClinkUpdateClientResponseBodyDataClientPermission) String() string {
	return dara.Prettify(s)
}

func (s ClinkUpdateClientResponseBodyDataClientPermission) GoString() string {
	return s.String()
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) GetAsr() *int64 {
	return s.Asr
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) GetCall() *int64 {
	return s.Call
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) GetCdr() *int64 {
	return s.Cdr
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) GetChat() *int64 {
	return s.Chat
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) GetOtherData() *int64 {
	return s.OtherData
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) GetRecord() *int64 {
	return s.Record
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) GetRecordDownload() *int64 {
	return s.RecordDownload
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) GetSms() *int64 {
	return s.Sms
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) GetTaskInventory() *int64 {
	return s.TaskInventory
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) GetTransfer() *int64 {
	return s.Transfer
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) SetAsr(v int64) *ClinkUpdateClientResponseBodyDataClientPermission {
	s.Asr = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) SetCall(v int64) *ClinkUpdateClientResponseBodyDataClientPermission {
	s.Call = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) SetCdr(v int64) *ClinkUpdateClientResponseBodyDataClientPermission {
	s.Cdr = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) SetChat(v int64) *ClinkUpdateClientResponseBodyDataClientPermission {
	s.Chat = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) SetOtherData(v int64) *ClinkUpdateClientResponseBodyDataClientPermission {
	s.OtherData = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) SetRecord(v int64) *ClinkUpdateClientResponseBodyDataClientPermission {
	s.Record = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) SetRecordDownload(v int64) *ClinkUpdateClientResponseBodyDataClientPermission {
	s.RecordDownload = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) SetSms(v int64) *ClinkUpdateClientResponseBodyDataClientPermission {
	s.Sms = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) SetTaskInventory(v int64) *ClinkUpdateClientResponseBodyDataClientPermission {
	s.TaskInventory = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) SetTransfer(v int64) *ClinkUpdateClientResponseBodyDataClientPermission {
	s.Transfer = &v
	return s
}

func (s *ClinkUpdateClientResponseBodyDataClientPermission) Validate() error {
	return dara.Validate(s)
}
