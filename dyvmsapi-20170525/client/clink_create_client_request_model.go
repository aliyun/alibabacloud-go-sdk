// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkCreateClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v int64) *ClinkCreateClientRequest
	GetActive() *int64
	SetAreaCode(v string) *ClinkCreateClientRequest
	GetAreaCode() *string
	SetAssignType(v int64) *ClinkCreateClientRequest
	GetAssignType() *int64
	SetClid(v []*string) *ClinkCreateClientRequest
	GetClid() []*string
	SetClidArea(v []*ClinkCreateClientRequestClidArea) *ClinkCreateClientRequest
	GetClidArea() []*ClinkCreateClientRequestClidArea
	SetClidDefault(v []*string) *ClinkCreateClientRequest
	GetClidDefault() []*string
	SetClidRule(v int64) *ClinkCreateClientRequest
	GetClidRule() *int64
	SetClidType(v int64) *ClinkCreateClientRequest
	GetClidType() *int64
	SetCloudNumberEnabled(v int64) *ClinkCreateClientRequest
	GetCloudNumberEnabled() *int64
	SetCloudNumberModes(v []*int64) *ClinkCreateClientRequest
	GetCloudNumberModes() []*int64
	SetCno(v int64) *ClinkCreateClientRequest
	GetCno() *int64
	SetCrmId(v int64) *ClinkCreateClientRequest
	GetCrmId() *int64
	SetDynamicTelGroupIdDefault(v int64) *ClinkCreateClientRequest
	GetDynamicTelGroupIdDefault() *int64
	SetDynamicTelGroupName(v string) *ClinkCreateClientRequest
	GetDynamicTelGroupName() *string
	SetEnterpriseId(v int64) *ClinkCreateClientRequest
	GetEnterpriseId() *int64
	SetHiddenTel(v int64) *ClinkCreateClientRequest
	GetHiddenTel() *int64
	SetName(v string) *ClinkCreateClientRequest
	GetName() *string
	SetObClidDefaultType(v int64) *ClinkCreateClientRequest
	GetObClidDefaultType() *int64
	SetObHangupSms(v int64) *ClinkCreateClientRequest
	GetObHangupSms() *int64
	SetOwnerId(v int64) *ClinkCreateClientRequest
	GetOwnerId() *int64
	SetPassword(v string) *ClinkCreateClientRequest
	GetPassword() *string
	SetPauseLogin(v int64) *ClinkCreateClientRequest
	GetPauseLogin() *int64
	SetPermission(v *ClinkCreateClientRequestPermission) *ClinkCreateClientRequest
	GetPermission() *ClinkCreateClientRequestPermission
	SetQnos(v []*string) *ClinkCreateClientRequest
	GetQnos() []*string
	SetRecurrentselectionType(v int64) *ClinkCreateClientRequest
	GetRecurrentselectionType() *int64
	SetRecurrentselectionValue(v int64) *ClinkCreateClientRequest
	GetRecurrentselectionValue() *int64
	SetResourceOwnerAccount(v string) *ClinkCreateClientRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkCreateClientRequest
	GetResourceOwnerId() *int64
	SetRole(v int64) *ClinkCreateClientRequest
	GetRole() *int64
	SetServeArea(v []*string) *ClinkCreateClientRequest
	GetServeArea() []*string
	SetType(v int64) *ClinkCreateClientRequest
	GetType() *int64
	SetWechatMiniProgramRtc(v int64) *ClinkCreateClientRequest
	GetWechatMiniProgramRtc() *int64
	SetWrapupTime(v int64) *ClinkCreateClientRequest
	GetWrapupTime() *int64
}

type ClinkCreateClientRequest struct {
	// 启用状态，0: 停用；1: 启用，默认值为 1
	//
	// example:
	//
	// 1
	Active *int64 `json:"Active,omitempty" xml:"Active,omitempty"`
	// 所属区号
	//
	// This parameter is required.
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
	// 可外显号码集合，当clidType（外显号码类型）值为非0时为必填项
	Clid []*string `json:"Clid,omitempty" xml:"Clid,omitempty" type:"Repeated"`
	// 按地区外显配置，当clidType（外显号码类型）值为3且clidRule（外显规则）值为2（自定义）时有意义，具体配置项见 clidArea 对象
	ClidArea []*ClinkCreateClientRequestClidArea `json:"ClidArea,omitempty" xml:"ClidArea,omitempty" type:"Repeated"`
	// 缺省外显号码集合，当clidType（外显号码类型）值为3且clidRule（外显规则）值为1（动态）时，支持配置缺省外显
	ClidDefault []*string `json:"ClidDefault,omitempty" xml:"ClidDefault,omitempty" type:"Repeated"`
	// 外显规则，默认值为 1 当clidType（外显号码类型）值为0（全部）时，1:随机；2: 轮选。 当clidType（外显号码类型）值为3（智能外显）时，1:动态；2: 自定义。
	//
	// example:
	//
	// 1
	ClidRule *int64 `json:"ClidRule,omitempty" xml:"ClidRule,omitempty"`
	// 外显号码类型，0: 全部；2: 座席指定；3: 智能外显；默认值为 0
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
	// 云号码四种呼叫模式；数组长度 为 4，依次对应云号码外呼的四种模式（实体卡、工作卡、两端呼、RTC)的开启状态； 例如：[0,0,0,1] 表示 座席只可使用 RTC 外呼模式。 注意：仅当企业开启了云手机外呼功能，才能为座席设置该属性，否则，座席将创建失败
	CloudNumberModes []*int64 `json:"CloudNumberModes,omitempty" xml:"CloudNumberModes,omitempty" type:"Repeated"`
	// 座席工号，长度为4-11个字符，必须全部为数字，要求唯一
	//
	// This parameter is required.
	//
	// example:
	//
	// 1111
	Cno *int64 `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// CRM 编号，与第三方 CRM 系统对接时，可作为唯一标识
	//
	// example:
	//
	// 87
	CrmId *int64 `json:"CrmId,omitempty" xml:"CrmId,omitempty"`
	// 缺省动态外呼组id，当obClidDefaultType值为1时为必填项
	//
	// example:
	//
	// 17
	DynamicTelGroupIdDefault *int64 `json:"DynamicTelGroupIdDefault,omitempty" xml:"DynamicTelGroupIdDefault,omitempty"`
	// 动态号码组名称，当assignType（号码类型）值为1（动态号码组）时为必填项
	//
	// example:
	//
	// 示例值示例值
	DynamicTelGroupName *string `json:"DynamicTelGroupName,omitempty" xml:"DynamicTelGroupName,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 8004970
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 号码隐藏类型，0: 不隐藏；1: 隐藏规则与全局设置保持一致，默认值为 0
	//
	// example:
	//
	// 0
	HiddenTel *int64 `json:"HiddenTel,omitempty" xml:"HiddenTel,omitempty"`
	// 座席名称，只允许输入中文、字母、数字、下划线，长度不超过 20 个字符
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 缺省号码类型 0：号码 1：动态号码组 默认0
	//
	// example:
	//
	// 0
	ObClidDefaultType *int64 `json:"ObClidDefaultType,omitempty" xml:"ObClidDefaultType,omitempty"`
	// 外呼挂机短信开关 0：关闭， 1：开启。开启后，当座席发起外呼时，系统会依据【短信中心 - 短信设置 - 外呼双方接听通知 / 外呼客户未接听通知】中的配置条件发送短信。
	//
	// example:
	//
	// 0
	ObHangupSms *int64 `json:"ObHangupSms,omitempty" xml:"ObHangupSms,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 座席密码，采用 AES加密 ，默认长度为 8 位。如果企业开启了密码安全设置，则需要按照设置的规则设置
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// 是否可置忙登录，关闭后，座席不允许置忙登录，0：关闭，1：开启，默认1。
	//
	// example:
	//
	// 1
	PauseLogin *int64 `json:"PauseLogin,omitempty" xml:"PauseLogin,omitempty"`
	// permission
	//
	// This parameter is required.
	Permission *ClinkCreateClientRequestPermission `json:"Permission,omitempty" xml:"Permission,omitempty" type:"Struct"`
	// 所属队列号集合
	Qnos []*string `json:"Qnos,omitempty" xml:"Qnos,omitempty" type:"Repeated"`
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
	RecurrentselectionValue *int64  `json:"RecurrentselectionValue,omitempty" xml:"RecurrentselectionValue,omitempty"`
	ResourceOwnerAccount    *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 座席角色，0: 普通座席；1: 班长座席，默认值为 0
	//
	// example:
	//
	// 0
	Role *int64 `json:"Role,omitempty" xml:"Role,omitempty"`
	// 座席服务地区，区号数组 配置座席服务的地区，可用于外呼任务中“按服务地区分配”的策略。
	ServeArea []*string `json:"ServeArea,omitempty" xml:"ServeArea,omitempty" type:"Repeated"`
	// 座席类型，1：全渠道、2：呼叫中心、3：在线客服；默认值为 2
	//
	// example:
	//
	// 2
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
	// 微信小程序RTC开关 1：开启 0：关闭 默认关闭
	//
	// example:
	//
	// 0
	WechatMiniProgramRtc *int64 `json:"WechatMiniProgramRtc,omitempty" xml:"WechatMiniProgramRtc,omitempty"`
	// 整理时长，座席进行外呼操作后的整理时间。取值范围 0-300 秒，默认值为 30
	//
	// example:
	//
	// 24
	WrapupTime *int64 `json:"WrapupTime,omitempty" xml:"WrapupTime,omitempty"`
}

func (s ClinkCreateClientRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateClientRequest) GoString() string {
	return s.String()
}

func (s *ClinkCreateClientRequest) GetActive() *int64 {
	return s.Active
}

func (s *ClinkCreateClientRequest) GetAreaCode() *string {
	return s.AreaCode
}

func (s *ClinkCreateClientRequest) GetAssignType() *int64 {
	return s.AssignType
}

func (s *ClinkCreateClientRequest) GetClid() []*string {
	return s.Clid
}

func (s *ClinkCreateClientRequest) GetClidArea() []*ClinkCreateClientRequestClidArea {
	return s.ClidArea
}

func (s *ClinkCreateClientRequest) GetClidDefault() []*string {
	return s.ClidDefault
}

func (s *ClinkCreateClientRequest) GetClidRule() *int64 {
	return s.ClidRule
}

func (s *ClinkCreateClientRequest) GetClidType() *int64 {
	return s.ClidType
}

func (s *ClinkCreateClientRequest) GetCloudNumberEnabled() *int64 {
	return s.CloudNumberEnabled
}

func (s *ClinkCreateClientRequest) GetCloudNumberModes() []*int64 {
	return s.CloudNumberModes
}

func (s *ClinkCreateClientRequest) GetCno() *int64 {
	return s.Cno
}

func (s *ClinkCreateClientRequest) GetCrmId() *int64 {
	return s.CrmId
}

func (s *ClinkCreateClientRequest) GetDynamicTelGroupIdDefault() *int64 {
	return s.DynamicTelGroupIdDefault
}

func (s *ClinkCreateClientRequest) GetDynamicTelGroupName() *string {
	return s.DynamicTelGroupName
}

func (s *ClinkCreateClientRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkCreateClientRequest) GetHiddenTel() *int64 {
	return s.HiddenTel
}

func (s *ClinkCreateClientRequest) GetName() *string {
	return s.Name
}

func (s *ClinkCreateClientRequest) GetObClidDefaultType() *int64 {
	return s.ObClidDefaultType
}

func (s *ClinkCreateClientRequest) GetObHangupSms() *int64 {
	return s.ObHangupSms
}

func (s *ClinkCreateClientRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkCreateClientRequest) GetPassword() *string {
	return s.Password
}

func (s *ClinkCreateClientRequest) GetPauseLogin() *int64 {
	return s.PauseLogin
}

func (s *ClinkCreateClientRequest) GetPermission() *ClinkCreateClientRequestPermission {
	return s.Permission
}

func (s *ClinkCreateClientRequest) GetQnos() []*string {
	return s.Qnos
}

func (s *ClinkCreateClientRequest) GetRecurrentselectionType() *int64 {
	return s.RecurrentselectionType
}

func (s *ClinkCreateClientRequest) GetRecurrentselectionValue() *int64 {
	return s.RecurrentselectionValue
}

func (s *ClinkCreateClientRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkCreateClientRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkCreateClientRequest) GetRole() *int64 {
	return s.Role
}

func (s *ClinkCreateClientRequest) GetServeArea() []*string {
	return s.ServeArea
}

func (s *ClinkCreateClientRequest) GetType() *int64 {
	return s.Type
}

func (s *ClinkCreateClientRequest) GetWechatMiniProgramRtc() *int64 {
	return s.WechatMiniProgramRtc
}

func (s *ClinkCreateClientRequest) GetWrapupTime() *int64 {
	return s.WrapupTime
}

func (s *ClinkCreateClientRequest) SetActive(v int64) *ClinkCreateClientRequest {
	s.Active = &v
	return s
}

func (s *ClinkCreateClientRequest) SetAreaCode(v string) *ClinkCreateClientRequest {
	s.AreaCode = &v
	return s
}

func (s *ClinkCreateClientRequest) SetAssignType(v int64) *ClinkCreateClientRequest {
	s.AssignType = &v
	return s
}

func (s *ClinkCreateClientRequest) SetClid(v []*string) *ClinkCreateClientRequest {
	s.Clid = v
	return s
}

func (s *ClinkCreateClientRequest) SetClidArea(v []*ClinkCreateClientRequestClidArea) *ClinkCreateClientRequest {
	s.ClidArea = v
	return s
}

func (s *ClinkCreateClientRequest) SetClidDefault(v []*string) *ClinkCreateClientRequest {
	s.ClidDefault = v
	return s
}

func (s *ClinkCreateClientRequest) SetClidRule(v int64) *ClinkCreateClientRequest {
	s.ClidRule = &v
	return s
}

func (s *ClinkCreateClientRequest) SetClidType(v int64) *ClinkCreateClientRequest {
	s.ClidType = &v
	return s
}

func (s *ClinkCreateClientRequest) SetCloudNumberEnabled(v int64) *ClinkCreateClientRequest {
	s.CloudNumberEnabled = &v
	return s
}

func (s *ClinkCreateClientRequest) SetCloudNumberModes(v []*int64) *ClinkCreateClientRequest {
	s.CloudNumberModes = v
	return s
}

func (s *ClinkCreateClientRequest) SetCno(v int64) *ClinkCreateClientRequest {
	s.Cno = &v
	return s
}

func (s *ClinkCreateClientRequest) SetCrmId(v int64) *ClinkCreateClientRequest {
	s.CrmId = &v
	return s
}

func (s *ClinkCreateClientRequest) SetDynamicTelGroupIdDefault(v int64) *ClinkCreateClientRequest {
	s.DynamicTelGroupIdDefault = &v
	return s
}

func (s *ClinkCreateClientRequest) SetDynamicTelGroupName(v string) *ClinkCreateClientRequest {
	s.DynamicTelGroupName = &v
	return s
}

func (s *ClinkCreateClientRequest) SetEnterpriseId(v int64) *ClinkCreateClientRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkCreateClientRequest) SetHiddenTel(v int64) *ClinkCreateClientRequest {
	s.HiddenTel = &v
	return s
}

func (s *ClinkCreateClientRequest) SetName(v string) *ClinkCreateClientRequest {
	s.Name = &v
	return s
}

func (s *ClinkCreateClientRequest) SetObClidDefaultType(v int64) *ClinkCreateClientRequest {
	s.ObClidDefaultType = &v
	return s
}

func (s *ClinkCreateClientRequest) SetObHangupSms(v int64) *ClinkCreateClientRequest {
	s.ObHangupSms = &v
	return s
}

func (s *ClinkCreateClientRequest) SetOwnerId(v int64) *ClinkCreateClientRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkCreateClientRequest) SetPassword(v string) *ClinkCreateClientRequest {
	s.Password = &v
	return s
}

func (s *ClinkCreateClientRequest) SetPauseLogin(v int64) *ClinkCreateClientRequest {
	s.PauseLogin = &v
	return s
}

func (s *ClinkCreateClientRequest) SetPermission(v *ClinkCreateClientRequestPermission) *ClinkCreateClientRequest {
	s.Permission = v
	return s
}

func (s *ClinkCreateClientRequest) SetQnos(v []*string) *ClinkCreateClientRequest {
	s.Qnos = v
	return s
}

func (s *ClinkCreateClientRequest) SetRecurrentselectionType(v int64) *ClinkCreateClientRequest {
	s.RecurrentselectionType = &v
	return s
}

func (s *ClinkCreateClientRequest) SetRecurrentselectionValue(v int64) *ClinkCreateClientRequest {
	s.RecurrentselectionValue = &v
	return s
}

func (s *ClinkCreateClientRequest) SetResourceOwnerAccount(v string) *ClinkCreateClientRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkCreateClientRequest) SetResourceOwnerId(v int64) *ClinkCreateClientRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkCreateClientRequest) SetRole(v int64) *ClinkCreateClientRequest {
	s.Role = &v
	return s
}

func (s *ClinkCreateClientRequest) SetServeArea(v []*string) *ClinkCreateClientRequest {
	s.ServeArea = v
	return s
}

func (s *ClinkCreateClientRequest) SetType(v int64) *ClinkCreateClientRequest {
	s.Type = &v
	return s
}

func (s *ClinkCreateClientRequest) SetWechatMiniProgramRtc(v int64) *ClinkCreateClientRequest {
	s.WechatMiniProgramRtc = &v
	return s
}

func (s *ClinkCreateClientRequest) SetWrapupTime(v int64) *ClinkCreateClientRequest {
	s.WrapupTime = &v
	return s
}

func (s *ClinkCreateClientRequest) Validate() error {
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

type ClinkCreateClientRequestClidArea struct {
	// 地区组名称
	//
	// example:
	//
	// 示例值示例值示例值
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

func (s ClinkCreateClientRequestClidArea) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateClientRequestClidArea) GoString() string {
	return s.String()
}

func (s *ClinkCreateClientRequestClidArea) GetAreaGroupName() *string {
	return s.AreaGroupName
}

func (s *ClinkCreateClientRequestClidArea) GetAssignType() *int64 {
	return s.AssignType
}

func (s *ClinkCreateClientRequestClidArea) GetDynamicTelGroupName() *string {
	return s.DynamicTelGroupName
}

func (s *ClinkCreateClientRequestClidArea) GetObClids() []*string {
	return s.ObClids
}

func (s *ClinkCreateClientRequestClidArea) SetAreaGroupName(v string) *ClinkCreateClientRequestClidArea {
	s.AreaGroupName = &v
	return s
}

func (s *ClinkCreateClientRequestClidArea) SetAssignType(v int64) *ClinkCreateClientRequestClidArea {
	s.AssignType = &v
	return s
}

func (s *ClinkCreateClientRequestClidArea) SetDynamicTelGroupName(v string) *ClinkCreateClientRequestClidArea {
	s.DynamicTelGroupName = &v
	return s
}

func (s *ClinkCreateClientRequestClidArea) SetObClids(v []*string) *ClinkCreateClientRequestClidArea {
	s.ObClids = v
	return s
}

func (s *ClinkCreateClientRequestClidArea) Validate() error {
	return dara.Validate(s)
}

type ClinkCreateClientRequestPermission struct {
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
	// 0
	Transfer *int64 `json:"Transfer,omitempty" xml:"Transfer,omitempty"`
}

func (s ClinkCreateClientRequestPermission) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateClientRequestPermission) GoString() string {
	return s.String()
}

func (s *ClinkCreateClientRequestPermission) GetAsr() *int64 {
	return s.Asr
}

func (s *ClinkCreateClientRequestPermission) GetCall() *int64 {
	return s.Call
}

func (s *ClinkCreateClientRequestPermission) GetCdr() *int64 {
	return s.Cdr
}

func (s *ClinkCreateClientRequestPermission) GetChat() *int64 {
	return s.Chat
}

func (s *ClinkCreateClientRequestPermission) GetOtherData() *int64 {
	return s.OtherData
}

func (s *ClinkCreateClientRequestPermission) GetRecord() *int64 {
	return s.Record
}

func (s *ClinkCreateClientRequestPermission) GetRecordDownload() *int64 {
	return s.RecordDownload
}

func (s *ClinkCreateClientRequestPermission) GetSms() *int64 {
	return s.Sms
}

func (s *ClinkCreateClientRequestPermission) GetTaskInventory() *int64 {
	return s.TaskInventory
}

func (s *ClinkCreateClientRequestPermission) GetTransfer() *int64 {
	return s.Transfer
}

func (s *ClinkCreateClientRequestPermission) SetAsr(v int64) *ClinkCreateClientRequestPermission {
	s.Asr = &v
	return s
}

func (s *ClinkCreateClientRequestPermission) SetCall(v int64) *ClinkCreateClientRequestPermission {
	s.Call = &v
	return s
}

func (s *ClinkCreateClientRequestPermission) SetCdr(v int64) *ClinkCreateClientRequestPermission {
	s.Cdr = &v
	return s
}

func (s *ClinkCreateClientRequestPermission) SetChat(v int64) *ClinkCreateClientRequestPermission {
	s.Chat = &v
	return s
}

func (s *ClinkCreateClientRequestPermission) SetOtherData(v int64) *ClinkCreateClientRequestPermission {
	s.OtherData = &v
	return s
}

func (s *ClinkCreateClientRequestPermission) SetRecord(v int64) *ClinkCreateClientRequestPermission {
	s.Record = &v
	return s
}

func (s *ClinkCreateClientRequestPermission) SetRecordDownload(v int64) *ClinkCreateClientRequestPermission {
	s.RecordDownload = &v
	return s
}

func (s *ClinkCreateClientRequestPermission) SetSms(v int64) *ClinkCreateClientRequestPermission {
	s.Sms = &v
	return s
}

func (s *ClinkCreateClientRequestPermission) SetTaskInventory(v int64) *ClinkCreateClientRequestPermission {
	s.TaskInventory = &v
	return s
}

func (s *ClinkCreateClientRequestPermission) SetTransfer(v int64) *ClinkCreateClientRequestPermission {
	s.Transfer = &v
	return s
}

func (s *ClinkCreateClientRequestPermission) Validate() error {
	return dara.Validate(s)
}
