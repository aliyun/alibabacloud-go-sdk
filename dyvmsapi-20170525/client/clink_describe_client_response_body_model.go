// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkDescribeClientResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkDescribeClientResponseBody
	GetCode() *string
	SetData(v *ClinkDescribeClientResponseBodyData) *ClinkDescribeClientResponseBody
	GetData() *ClinkDescribeClientResponseBodyData
	SetMessage(v string) *ClinkDescribeClientResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkDescribeClientResponseBody
	GetRequestId() *string
}

type ClinkDescribeClientResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkDescribeClientResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// ED815433-724A-4357-9991-A54AD2FF09FD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkDescribeClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeClientResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkDescribeClientResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkDescribeClientResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkDescribeClientResponseBody) GetData() *ClinkDescribeClientResponseBodyData {
	return s.Data
}

func (s *ClinkDescribeClientResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkDescribeClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkDescribeClientResponseBody) SetAccessDeniedDetail(v string) *ClinkDescribeClientResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkDescribeClientResponseBody) SetCode(v string) *ClinkDescribeClientResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkDescribeClientResponseBody) SetData(v *ClinkDescribeClientResponseBodyData) *ClinkDescribeClientResponseBody {
	s.Data = v
	return s
}

func (s *ClinkDescribeClientResponseBody) SetMessage(v string) *ClinkDescribeClientResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkDescribeClientResponseBody) SetRequestId(v string) *ClinkDescribeClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkDescribeClientResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkDescribeClientResponseBodyData struct {
	Client *ClinkDescribeClientResponseBodyDataClient `json:"Client,omitempty" xml:"Client,omitempty" type:"Struct"`
	// 请求 id
	//
	// example:
	//
	// 示例值示例值示例值
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
}

func (s ClinkDescribeClientResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeClientResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkDescribeClientResponseBodyData) GetClient() *ClinkDescribeClientResponseBodyDataClient {
	return s.Client
}

func (s *ClinkDescribeClientResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkDescribeClientResponseBodyData) SetClient(v *ClinkDescribeClientResponseBodyDataClient) *ClinkDescribeClientResponseBodyData {
	s.Client = v
	return s
}

func (s *ClinkDescribeClientResponseBodyData) SetClinkRequestId(v string) *ClinkDescribeClientResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyData) Validate() error {
	if s.Client != nil {
		if err := s.Client.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkDescribeClientResponseBodyDataClient struct {
	// 是否激活，0: 否；1: 是
	//
	// example:
	//
	// 1
	Active *int64 `json:"Active,omitempty" xml:"Active,omitempty"`
	// 区号
	//
	// example:
	//
	// 010
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// 座席满意度自动执行，1：开启；0：关闭，默认开启
	//
	// example:
	//
	// 1
	AutoInvestigation *int64 `json:"AutoInvestigation,omitempty" xml:"AutoInvestigation,omitempty"`
	// 座席绑定电话
	//
	// example:
	//
	// xxx
	BindTel *string `json:"BindTel,omitempty" xml:"BindTel,omitempty"`
	// 外显号码数组
	Clid []*string `json:"Clid,omitempty" xml:"Clid,omitempty" type:"Repeated"`
	// 按地区外显配置，当clidType（外显号码类型）值为3且clidRule（外显规则）值为2（自定义）时有意义，具体配置项见 clidArea 对象
	ClidArea []*ClinkDescribeClientResponseBodyDataClientClidArea `json:"ClidArea,omitempty" xml:"ClidArea,omitempty" type:"Repeated"`
	// 外显规则 当clidType（外显号码类型）值为0（全部）时，1:随机；2: 轮选。 当clidType（外显号码类型）值为3（智能外显）时，1:动态；2: 自定义。
	//
	// example:
	//
	// 1
	ClidRule *int64 `json:"ClidRule,omitempty" xml:"ClidRule,omitempty"`
	// 外显号码类型，0: 全部; 2: 座席指定; 3: 智能外显
	//
	// example:
	//
	// 1
	ClidType *int64 `json:"ClidType,omitempty" xml:"ClidType,omitempty"`
	// 云号码外呼开关，0-关，1-开
	//
	// example:
	//
	// 0
	CloudNumberEnabled *int64 `json:"CloudNumberEnabled,omitempty" xml:"CloudNumberEnabled,omitempty"`
	// 云号码四种呼叫模式；数组长度 为 4，依次对应云号码外呼的四种模式（实体卡、工作卡、两端呼、RTC)的开启状态；
	CloudNumberModes []*int64 `json:"CloudNumberModes,omitempty" xml:"CloudNumberModes,omitempty" type:"Repeated"`
	// 座席号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 座席crm id
	//
	// example:
	//
	// 123
	CrmId *string `json:"CrmId,omitempty" xml:"CrmId,omitempty"`
	// 号码隐藏类型，0: 不隐藏；1: 隐藏规则与全局设置保持一致
	//
	// example:
	//
	// 0
	HiddenTel *string `json:"HiddenTel,omitempty" xml:"HiddenTel,omitempty"`
	// 呼入整理时长，客户来电座席接听，双方通话结束后座席的话后处理时长，此期间座席不接收新的客户来电。取值范围：0～3600秒。
	//
	// example:
	//
	// 20
	IbWrapupTime *int64 `json:"IbWrapupTime,omitempty" xml:"IbWrapupTime,omitempty"`
	// 呼入整理类型，1：队列；2：座席。
	//
	// example:
	//
	// 1
	IbWrapupType *int64 `json:"IbWrapupType,omitempty" xml:"IbWrapupType,omitempty"`
	// 座席名称
	//
	// example:
	//
	// xclient
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 外呼挂机短信开关 0：关闭， 1：开启。开启后，当座席发起外呼时，系统会依据【短信中心 - 短信设置 - 外呼双方接听通知 / 外呼客户未接听通知】中的配置条件发送短信。
	//
	// example:
	//
	// 0
	ObHangupSms *int64 `json:"ObHangupSms,omitempty" xml:"ObHangupSms,omitempty"`
	// 是否可置忙登录，关闭后，座席不允许置忙登录，0：关闭，1：开启，默认1。
	//
	// example:
	//
	// 1
	PauseLogin *int64 `json:"PauseLogin,omitempty" xml:"PauseLogin,omitempty"`
	// permission
	Permission *ClinkDescribeClientResponseBodyDataClientPermission `json:"Permission,omitempty" xml:"Permission,omitempty" type:"Struct"`
	// 座席所属队列的队列号数组
	Qnos []*string `json:"Qnos,omitempty" xml:"Qnos,omitempty" type:"Repeated"`
	// 座席所属队列详情
	Queues []*ClinkDescribeClientResponseBodyDataClientQueues `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Repeated"`
	// 轮选方式，1: 按天轮选；2: 按次轮选，当外显号码类型为全部、外显规则为轮选时配置才生效
	//
	// example:
	//
	// 1
	RecurrentselectionType *int64 `json:"RecurrentselectionType,omitempty" xml:"RecurrentselectionType,omitempty"`
	// 轮选值设置，当外显号码类型为全部、外显规则为轮选时配置才生效 1: 按天轮选，每 n 天外呼更换一次外显号码，可设置 1-30 天 2: 按次轮选，每 n 次外呼更换一次外显号码，可设置 1-30 次
	//
	// example:
	//
	// 1
	RecurrentselectionValue *int64 `json:"RecurrentselectionValue,omitempty" xml:"RecurrentselectionValue,omitempty"`
	// 座席备用手机号
	//
	// example:
	//
	// xxx
	ReserveTel *string `json:"ReserveTel,omitempty" xml:"ReserveTel,omitempty"`
	// 座席角色，0: 普通座席；1: 班长座席
	//
	// example:
	//
	// 1
	Role *int64 `json:"Role,omitempty" xml:"Role,omitempty"`
	// 座席服务地区，区号数组 配置座席服务的地区，可用于外呼任务中“按服务地区分配”的策略。
	ServeArea []*string `json:"ServeArea,omitempty" xml:"ServeArea,omitempty" type:"Repeated"`
	// 座席状态，0: 离线；1: 在线
	//
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 电话绑定类型，电话类型， 1: 固话；2: 手机；3:IP话机；4:软电话
	//
	// example:
	//
	// 1
	TelType *int64 `json:"TelType,omitempty" xml:"TelType,omitempty"`
	// 座席类型，1：全渠道、2：呼叫中心、3：在线客服
	//
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
	// 整理时长，座席进行外呼操作后的整理时间，取值范围 0-300 秒。
	//
	// example:
	//
	// 59
	WrapupTime *int64 `json:"WrapupTime,omitempty" xml:"WrapupTime,omitempty"`
}

func (s ClinkDescribeClientResponseBodyDataClient) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeClientResponseBodyDataClient) GoString() string {
	return s.String()
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetActive() *int64 {
	return s.Active
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetAreaCode() *string {
	return s.AreaCode
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetAutoInvestigation() *int64 {
	return s.AutoInvestigation
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetBindTel() *string {
	return s.BindTel
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetClid() []*string {
	return s.Clid
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetClidArea() []*ClinkDescribeClientResponseBodyDataClientClidArea {
	return s.ClidArea
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetClidRule() *int64 {
	return s.ClidRule
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetClidType() *int64 {
	return s.ClidType
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetCloudNumberEnabled() *int64 {
	return s.CloudNumberEnabled
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetCloudNumberModes() []*int64 {
	return s.CloudNumberModes
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetCno() *string {
	return s.Cno
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetCrmId() *string {
	return s.CrmId
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetHiddenTel() *string {
	return s.HiddenTel
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetIbWrapupTime() *int64 {
	return s.IbWrapupTime
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetIbWrapupType() *int64 {
	return s.IbWrapupType
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetName() *string {
	return s.Name
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetObHangupSms() *int64 {
	return s.ObHangupSms
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetPauseLogin() *int64 {
	return s.PauseLogin
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetPermission() *ClinkDescribeClientResponseBodyDataClientPermission {
	return s.Permission
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetQnos() []*string {
	return s.Qnos
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetQueues() []*ClinkDescribeClientResponseBodyDataClientQueues {
	return s.Queues
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetRecurrentselectionType() *int64 {
	return s.RecurrentselectionType
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetRecurrentselectionValue() *int64 {
	return s.RecurrentselectionValue
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetReserveTel() *string {
	return s.ReserveTel
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetRole() *int64 {
	return s.Role
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetServeArea() []*string {
	return s.ServeArea
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetStatus() *int64 {
	return s.Status
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetTelType() *int64 {
	return s.TelType
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetType() *int64 {
	return s.Type
}

func (s *ClinkDescribeClientResponseBodyDataClient) GetWrapupTime() *int64 {
	return s.WrapupTime
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetActive(v int64) *ClinkDescribeClientResponseBodyDataClient {
	s.Active = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetAreaCode(v string) *ClinkDescribeClientResponseBodyDataClient {
	s.AreaCode = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetAutoInvestigation(v int64) *ClinkDescribeClientResponseBodyDataClient {
	s.AutoInvestigation = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetBindTel(v string) *ClinkDescribeClientResponseBodyDataClient {
	s.BindTel = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetClid(v []*string) *ClinkDescribeClientResponseBodyDataClient {
	s.Clid = v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetClidArea(v []*ClinkDescribeClientResponseBodyDataClientClidArea) *ClinkDescribeClientResponseBodyDataClient {
	s.ClidArea = v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetClidRule(v int64) *ClinkDescribeClientResponseBodyDataClient {
	s.ClidRule = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetClidType(v int64) *ClinkDescribeClientResponseBodyDataClient {
	s.ClidType = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetCloudNumberEnabled(v int64) *ClinkDescribeClientResponseBodyDataClient {
	s.CloudNumberEnabled = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetCloudNumberModes(v []*int64) *ClinkDescribeClientResponseBodyDataClient {
	s.CloudNumberModes = v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetCno(v string) *ClinkDescribeClientResponseBodyDataClient {
	s.Cno = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetCrmId(v string) *ClinkDescribeClientResponseBodyDataClient {
	s.CrmId = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetHiddenTel(v string) *ClinkDescribeClientResponseBodyDataClient {
	s.HiddenTel = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetIbWrapupTime(v int64) *ClinkDescribeClientResponseBodyDataClient {
	s.IbWrapupTime = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetIbWrapupType(v int64) *ClinkDescribeClientResponseBodyDataClient {
	s.IbWrapupType = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetName(v string) *ClinkDescribeClientResponseBodyDataClient {
	s.Name = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetObHangupSms(v int64) *ClinkDescribeClientResponseBodyDataClient {
	s.ObHangupSms = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetPauseLogin(v int64) *ClinkDescribeClientResponseBodyDataClient {
	s.PauseLogin = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetPermission(v *ClinkDescribeClientResponseBodyDataClientPermission) *ClinkDescribeClientResponseBodyDataClient {
	s.Permission = v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetQnos(v []*string) *ClinkDescribeClientResponseBodyDataClient {
	s.Qnos = v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetQueues(v []*ClinkDescribeClientResponseBodyDataClientQueues) *ClinkDescribeClientResponseBodyDataClient {
	s.Queues = v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetRecurrentselectionType(v int64) *ClinkDescribeClientResponseBodyDataClient {
	s.RecurrentselectionType = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetRecurrentselectionValue(v int64) *ClinkDescribeClientResponseBodyDataClient {
	s.RecurrentselectionValue = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetReserveTel(v string) *ClinkDescribeClientResponseBodyDataClient {
	s.ReserveTel = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetRole(v int64) *ClinkDescribeClientResponseBodyDataClient {
	s.Role = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetServeArea(v []*string) *ClinkDescribeClientResponseBodyDataClient {
	s.ServeArea = v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetStatus(v int64) *ClinkDescribeClientResponseBodyDataClient {
	s.Status = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetTelType(v int64) *ClinkDescribeClientResponseBodyDataClient {
	s.TelType = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetType(v int64) *ClinkDescribeClientResponseBodyDataClient {
	s.Type = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) SetWrapupTime(v int64) *ClinkDescribeClientResponseBodyDataClient {
	s.WrapupTime = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClient) Validate() error {
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
	if s.Queues != nil {
		for _, item := range s.Queues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClinkDescribeClientResponseBodyDataClientClidArea struct {
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
	// xxx
	DynamicTelGroupName *string `json:"DynamicTelGroupName,omitempty" xml:"DynamicTelGroupName,omitempty"`
	// 外显号码
	ObClids []*string `json:"ObClids,omitempty" xml:"ObClids,omitempty" type:"Repeated"`
}

func (s ClinkDescribeClientResponseBodyDataClientClidArea) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeClientResponseBodyDataClientClidArea) GoString() string {
	return s.String()
}

func (s *ClinkDescribeClientResponseBodyDataClientClidArea) GetAreaGroupName() *string {
	return s.AreaGroupName
}

func (s *ClinkDescribeClientResponseBodyDataClientClidArea) GetAssignType() *int64 {
	return s.AssignType
}

func (s *ClinkDescribeClientResponseBodyDataClientClidArea) GetDynamicTelGroupName() *string {
	return s.DynamicTelGroupName
}

func (s *ClinkDescribeClientResponseBodyDataClientClidArea) GetObClids() []*string {
	return s.ObClids
}

func (s *ClinkDescribeClientResponseBodyDataClientClidArea) SetAreaGroupName(v string) *ClinkDescribeClientResponseBodyDataClientClidArea {
	s.AreaGroupName = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClientClidArea) SetAssignType(v int64) *ClinkDescribeClientResponseBodyDataClientClidArea {
	s.AssignType = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClientClidArea) SetDynamicTelGroupName(v string) *ClinkDescribeClientResponseBodyDataClientClidArea {
	s.DynamicTelGroupName = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClientClidArea) SetObClids(v []*string) *ClinkDescribeClientResponseBodyDataClientClidArea {
	s.ObClids = v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClientClidArea) Validate() error {
	return dara.Validate(s)
}

type ClinkDescribeClientResponseBodyDataClientPermission struct {
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
	// 1
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

func (s ClinkDescribeClientResponseBodyDataClientPermission) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeClientResponseBodyDataClientPermission) GoString() string {
	return s.String()
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) GetAsr() *int64 {
	return s.Asr
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) GetCall() *int64 {
	return s.Call
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) GetCdr() *int64 {
	return s.Cdr
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) GetChat() *int64 {
	return s.Chat
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) GetOtherData() *int64 {
	return s.OtherData
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) GetRecord() *int64 {
	return s.Record
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) GetRecordDownload() *int64 {
	return s.RecordDownload
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) GetSms() *int64 {
	return s.Sms
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) GetTaskInventory() *int64 {
	return s.TaskInventory
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) GetTransfer() *int64 {
	return s.Transfer
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) SetAsr(v int64) *ClinkDescribeClientResponseBodyDataClientPermission {
	s.Asr = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) SetCall(v int64) *ClinkDescribeClientResponseBodyDataClientPermission {
	s.Call = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) SetCdr(v int64) *ClinkDescribeClientResponseBodyDataClientPermission {
	s.Cdr = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) SetChat(v int64) *ClinkDescribeClientResponseBodyDataClientPermission {
	s.Chat = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) SetOtherData(v int64) *ClinkDescribeClientResponseBodyDataClientPermission {
	s.OtherData = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) SetRecord(v int64) *ClinkDescribeClientResponseBodyDataClientPermission {
	s.Record = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) SetRecordDownload(v int64) *ClinkDescribeClientResponseBodyDataClientPermission {
	s.RecordDownload = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) SetSms(v int64) *ClinkDescribeClientResponseBodyDataClientPermission {
	s.Sms = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) SetTaskInventory(v int64) *ClinkDescribeClientResponseBodyDataClientPermission {
	s.TaskInventory = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) SetTransfer(v int64) *ClinkDescribeClientResponseBodyDataClientPermission {
	s.Transfer = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClientPermission) Validate() error {
	return dara.Validate(s)
}

type ClinkDescribeClientResponseBodyDataClientQueues struct {
	// 队列名称
	//
	// example:
	//
	// xx1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 优先级
	//
	// example:
	//
	// 8
	Penalty *int64 `json:"Penalty,omitempty" xml:"Penalty,omitempty"`
	// 队列号
	//
	// example:
	//
	// 1122
	Qno *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
}

func (s ClinkDescribeClientResponseBodyDataClientQueues) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeClientResponseBodyDataClientQueues) GoString() string {
	return s.String()
}

func (s *ClinkDescribeClientResponseBodyDataClientQueues) GetName() *string {
	return s.Name
}

func (s *ClinkDescribeClientResponseBodyDataClientQueues) GetPenalty() *int64 {
	return s.Penalty
}

func (s *ClinkDescribeClientResponseBodyDataClientQueues) GetQno() *string {
	return s.Qno
}

func (s *ClinkDescribeClientResponseBodyDataClientQueues) SetName(v string) *ClinkDescribeClientResponseBodyDataClientQueues {
	s.Name = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClientQueues) SetPenalty(v int64) *ClinkDescribeClientResponseBodyDataClientQueues {
	s.Penalty = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClientQueues) SetQno(v string) *ClinkDescribeClientResponseBodyDataClientQueues {
	s.Qno = &v
	return s
}

func (s *ClinkDescribeClientResponseBodyDataClientQueues) Validate() error {
	return dara.Validate(s)
}
