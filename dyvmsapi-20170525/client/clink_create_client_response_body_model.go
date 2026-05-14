// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkCreateClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkCreateClientResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkCreateClientResponseBody
	GetCode() *string
	SetData(v *ClinkCreateClientResponseBodyData) *ClinkCreateClientResponseBody
	GetData() *ClinkCreateClientResponseBodyData
	SetMessage(v string) *ClinkCreateClientResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkCreateClientResponseBody
	GetRequestId() *string
}

type ClinkCreateClientResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkCreateClientResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkCreateClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateClientResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkCreateClientResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkCreateClientResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkCreateClientResponseBody) GetData() *ClinkCreateClientResponseBodyData {
	return s.Data
}

func (s *ClinkCreateClientResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkCreateClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkCreateClientResponseBody) SetAccessDeniedDetail(v string) *ClinkCreateClientResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkCreateClientResponseBody) SetCode(v string) *ClinkCreateClientResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkCreateClientResponseBody) SetData(v *ClinkCreateClientResponseBodyData) *ClinkCreateClientResponseBody {
	s.Data = v
	return s
}

func (s *ClinkCreateClientResponseBody) SetMessage(v string) *ClinkCreateClientResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkCreateClientResponseBody) SetRequestId(v string) *ClinkCreateClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkCreateClientResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkCreateClientResponseBodyData struct {
	Client *ClinkCreateClientResponseBodyDataClient `json:"Client,omitempty" xml:"Client,omitempty" type:"Struct"`
	// 请求 id
	//
	// example:
	//
	// 示例值
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
}

func (s ClinkCreateClientResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateClientResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkCreateClientResponseBodyData) GetClient() *ClinkCreateClientResponseBodyDataClient {
	return s.Client
}

func (s *ClinkCreateClientResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkCreateClientResponseBodyData) SetClient(v *ClinkCreateClientResponseBodyDataClient) *ClinkCreateClientResponseBodyData {
	s.Client = v
	return s
}

func (s *ClinkCreateClientResponseBodyData) SetClinkRequestId(v string) *ClinkCreateClientResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkCreateClientResponseBodyData) Validate() error {
	if s.Client != nil {
		if err := s.Client.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkCreateClientResponseBodyDataClient struct {
	// 启用状态，0: 停用；1: 启用
	//
	// example:
	//
	// 0
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
	ClidArea []*ClinkCreateClientResponseBodyDataClientClidArea `json:"ClidArea,omitempty" xml:"ClidArea,omitempty" type:"Repeated"`
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
	// 20
	CrmId *int64 `json:"CrmId,omitempty" xml:"CrmId,omitempty"`
	// 动态号码组名称，当assignType（号码类型）值为1（动态号码组）时为必填项
	//
	// example:
	//
	// xxx
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
	// 示例值
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 座席密码
	//
	// example:
	//
	// xxx
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// permission
	Permission *ClinkCreateClientResponseBodyDataClientPermission `json:"Permission,omitempty" xml:"Permission,omitempty" type:"Struct"`
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
	// 42
	WrapupTime *int64 `json:"WrapupTime,omitempty" xml:"WrapupTime,omitempty"`
}

func (s ClinkCreateClientResponseBodyDataClient) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateClientResponseBodyDataClient) GoString() string {
	return s.String()
}

func (s *ClinkCreateClientResponseBodyDataClient) GetActive() *int64 {
	return s.Active
}

func (s *ClinkCreateClientResponseBodyDataClient) GetAreaCode() *string {
	return s.AreaCode
}

func (s *ClinkCreateClientResponseBodyDataClient) GetAssignType() *int64 {
	return s.AssignType
}

func (s *ClinkCreateClientResponseBodyDataClient) GetClid() []*string {
	return s.Clid
}

func (s *ClinkCreateClientResponseBodyDataClient) GetClidArea() []*ClinkCreateClientResponseBodyDataClientClidArea {
	return s.ClidArea
}

func (s *ClinkCreateClientResponseBodyDataClient) GetClidDefault() []*string {
	return s.ClidDefault
}

func (s *ClinkCreateClientResponseBodyDataClient) GetClidRule() *int64 {
	return s.ClidRule
}

func (s *ClinkCreateClientResponseBodyDataClient) GetClidType() *int64 {
	return s.ClidType
}

func (s *ClinkCreateClientResponseBodyDataClient) GetCloudNumberEnabled() *int64 {
	return s.CloudNumberEnabled
}

func (s *ClinkCreateClientResponseBodyDataClient) GetCloudNumberModes() []*int64 {
	return s.CloudNumberModes
}

func (s *ClinkCreateClientResponseBodyDataClient) GetCno() *int64 {
	return s.Cno
}

func (s *ClinkCreateClientResponseBodyDataClient) GetCrmId() *int64 {
	return s.CrmId
}

func (s *ClinkCreateClientResponseBodyDataClient) GetDynamicTelGroupName() *string {
	return s.DynamicTelGroupName
}

func (s *ClinkCreateClientResponseBodyDataClient) GetHiddenTel() *int64 {
	return s.HiddenTel
}

func (s *ClinkCreateClientResponseBodyDataClient) GetName() *string {
	return s.Name
}

func (s *ClinkCreateClientResponseBodyDataClient) GetPassword() *string {
	return s.Password
}

func (s *ClinkCreateClientResponseBodyDataClient) GetPermission() *ClinkCreateClientResponseBodyDataClientPermission {
	return s.Permission
}

func (s *ClinkCreateClientResponseBodyDataClient) GetQnos() []*string {
	return s.Qnos
}

func (s *ClinkCreateClientResponseBodyDataClient) GetRecurrentselectionType() *int64 {
	return s.RecurrentselectionType
}

func (s *ClinkCreateClientResponseBodyDataClient) GetRecurrentselectionValue() *int64 {
	return s.RecurrentselectionValue
}

func (s *ClinkCreateClientResponseBodyDataClient) GetRole() *int64 {
	return s.Role
}

func (s *ClinkCreateClientResponseBodyDataClient) GetType() *int64 {
	return s.Type
}

func (s *ClinkCreateClientResponseBodyDataClient) GetWrapupTime() *int64 {
	return s.WrapupTime
}

func (s *ClinkCreateClientResponseBodyDataClient) SetActive(v int64) *ClinkCreateClientResponseBodyDataClient {
	s.Active = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetAreaCode(v string) *ClinkCreateClientResponseBodyDataClient {
	s.AreaCode = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetAssignType(v int64) *ClinkCreateClientResponseBodyDataClient {
	s.AssignType = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetClid(v []*string) *ClinkCreateClientResponseBodyDataClient {
	s.Clid = v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetClidArea(v []*ClinkCreateClientResponseBodyDataClientClidArea) *ClinkCreateClientResponseBodyDataClient {
	s.ClidArea = v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetClidDefault(v []*string) *ClinkCreateClientResponseBodyDataClient {
	s.ClidDefault = v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetClidRule(v int64) *ClinkCreateClientResponseBodyDataClient {
	s.ClidRule = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetClidType(v int64) *ClinkCreateClientResponseBodyDataClient {
	s.ClidType = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetCloudNumberEnabled(v int64) *ClinkCreateClientResponseBodyDataClient {
	s.CloudNumberEnabled = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetCloudNumberModes(v []*int64) *ClinkCreateClientResponseBodyDataClient {
	s.CloudNumberModes = v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetCno(v int64) *ClinkCreateClientResponseBodyDataClient {
	s.Cno = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetCrmId(v int64) *ClinkCreateClientResponseBodyDataClient {
	s.CrmId = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetDynamicTelGroupName(v string) *ClinkCreateClientResponseBodyDataClient {
	s.DynamicTelGroupName = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetHiddenTel(v int64) *ClinkCreateClientResponseBodyDataClient {
	s.HiddenTel = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetName(v string) *ClinkCreateClientResponseBodyDataClient {
	s.Name = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetPassword(v string) *ClinkCreateClientResponseBodyDataClient {
	s.Password = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetPermission(v *ClinkCreateClientResponseBodyDataClientPermission) *ClinkCreateClientResponseBodyDataClient {
	s.Permission = v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetQnos(v []*string) *ClinkCreateClientResponseBodyDataClient {
	s.Qnos = v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetRecurrentselectionType(v int64) *ClinkCreateClientResponseBodyDataClient {
	s.RecurrentselectionType = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetRecurrentselectionValue(v int64) *ClinkCreateClientResponseBodyDataClient {
	s.RecurrentselectionValue = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetRole(v int64) *ClinkCreateClientResponseBodyDataClient {
	s.Role = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetType(v int64) *ClinkCreateClientResponseBodyDataClient {
	s.Type = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) SetWrapupTime(v int64) *ClinkCreateClientResponseBodyDataClient {
	s.WrapupTime = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClient) Validate() error {
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

type ClinkCreateClientResponseBodyDataClientClidArea struct {
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

func (s ClinkCreateClientResponseBodyDataClientClidArea) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateClientResponseBodyDataClientClidArea) GoString() string {
	return s.String()
}

func (s *ClinkCreateClientResponseBodyDataClientClidArea) GetAreaGroupName() *string {
	return s.AreaGroupName
}

func (s *ClinkCreateClientResponseBodyDataClientClidArea) GetAssignType() *int64 {
	return s.AssignType
}

func (s *ClinkCreateClientResponseBodyDataClientClidArea) GetDynamicTelGroupName() *string {
	return s.DynamicTelGroupName
}

func (s *ClinkCreateClientResponseBodyDataClientClidArea) GetObClids() []*string {
	return s.ObClids
}

func (s *ClinkCreateClientResponseBodyDataClientClidArea) SetAreaGroupName(v string) *ClinkCreateClientResponseBodyDataClientClidArea {
	s.AreaGroupName = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClientClidArea) SetAssignType(v int64) *ClinkCreateClientResponseBodyDataClientClidArea {
	s.AssignType = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClientClidArea) SetDynamicTelGroupName(v string) *ClinkCreateClientResponseBodyDataClientClidArea {
	s.DynamicTelGroupName = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClientClidArea) SetObClids(v []*string) *ClinkCreateClientResponseBodyDataClientClidArea {
	s.ObClids = v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClientClidArea) Validate() error {
	return dara.Validate(s)
}

type ClinkCreateClientResponseBodyDataClientPermission struct {
	// 语音转写，0: 关闭；1: 呼入开启；2 外呼开启；3 全部开启；默认值为 0
	//
	// example:
	//
	// 0
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
	// 0
	Chat *int64 `json:"Chat,omitempty" xml:"Chat,omitempty"`
	// 其他数据查看权限：1：全部 2：所属员工组 3：自己 4:指定员工组 ,默认所属员工组
	//
	// example:
	//
	// 2
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
	// 0
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

func (s ClinkCreateClientResponseBodyDataClientPermission) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateClientResponseBodyDataClientPermission) GoString() string {
	return s.String()
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) GetAsr() *int64 {
	return s.Asr
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) GetCall() *int64 {
	return s.Call
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) GetCdr() *int64 {
	return s.Cdr
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) GetChat() *int64 {
	return s.Chat
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) GetOtherData() *int64 {
	return s.OtherData
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) GetRecord() *int64 {
	return s.Record
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) GetRecordDownload() *int64 {
	return s.RecordDownload
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) GetSms() *int64 {
	return s.Sms
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) GetTaskInventory() *int64 {
	return s.TaskInventory
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) GetTransfer() *int64 {
	return s.Transfer
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) SetAsr(v int64) *ClinkCreateClientResponseBodyDataClientPermission {
	s.Asr = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) SetCall(v int64) *ClinkCreateClientResponseBodyDataClientPermission {
	s.Call = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) SetCdr(v int64) *ClinkCreateClientResponseBodyDataClientPermission {
	s.Cdr = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) SetChat(v int64) *ClinkCreateClientResponseBodyDataClientPermission {
	s.Chat = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) SetOtherData(v int64) *ClinkCreateClientResponseBodyDataClientPermission {
	s.OtherData = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) SetRecord(v int64) *ClinkCreateClientResponseBodyDataClientPermission {
	s.Record = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) SetRecordDownload(v int64) *ClinkCreateClientResponseBodyDataClientPermission {
	s.RecordDownload = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) SetSms(v int64) *ClinkCreateClientResponseBodyDataClientPermission {
	s.Sms = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) SetTaskInventory(v int64) *ClinkCreateClientResponseBodyDataClientPermission {
	s.TaskInventory = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) SetTransfer(v int64) *ClinkCreateClientResponseBodyDataClientPermission {
	s.Transfer = &v
	return s
}

func (s *ClinkCreateClientResponseBodyDataClientPermission) Validate() error {
	return dara.Validate(s)
}
