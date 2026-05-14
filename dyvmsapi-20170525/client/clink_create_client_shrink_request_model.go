// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkCreateClientShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v int64) *ClinkCreateClientShrinkRequest
	GetActive() *int64
	SetAreaCode(v string) *ClinkCreateClientShrinkRequest
	GetAreaCode() *string
	SetAssignType(v int64) *ClinkCreateClientShrinkRequest
	GetAssignType() *int64
	SetClidShrink(v string) *ClinkCreateClientShrinkRequest
	GetClidShrink() *string
	SetClidAreaShrink(v string) *ClinkCreateClientShrinkRequest
	GetClidAreaShrink() *string
	SetClidDefaultShrink(v string) *ClinkCreateClientShrinkRequest
	GetClidDefaultShrink() *string
	SetClidRule(v int64) *ClinkCreateClientShrinkRequest
	GetClidRule() *int64
	SetClidType(v int64) *ClinkCreateClientShrinkRequest
	GetClidType() *int64
	SetCloudNumberEnabled(v int64) *ClinkCreateClientShrinkRequest
	GetCloudNumberEnabled() *int64
	SetCloudNumberModesShrink(v string) *ClinkCreateClientShrinkRequest
	GetCloudNumberModesShrink() *string
	SetCno(v int64) *ClinkCreateClientShrinkRequest
	GetCno() *int64
	SetCrmId(v int64) *ClinkCreateClientShrinkRequest
	GetCrmId() *int64
	SetDynamicTelGroupIdDefault(v int64) *ClinkCreateClientShrinkRequest
	GetDynamicTelGroupIdDefault() *int64
	SetDynamicTelGroupName(v string) *ClinkCreateClientShrinkRequest
	GetDynamicTelGroupName() *string
	SetEnterpriseId(v int64) *ClinkCreateClientShrinkRequest
	GetEnterpriseId() *int64
	SetHiddenTel(v int64) *ClinkCreateClientShrinkRequest
	GetHiddenTel() *int64
	SetName(v string) *ClinkCreateClientShrinkRequest
	GetName() *string
	SetObClidDefaultType(v int64) *ClinkCreateClientShrinkRequest
	GetObClidDefaultType() *int64
	SetObHangupSms(v int64) *ClinkCreateClientShrinkRequest
	GetObHangupSms() *int64
	SetOwnerId(v int64) *ClinkCreateClientShrinkRequest
	GetOwnerId() *int64
	SetPassword(v string) *ClinkCreateClientShrinkRequest
	GetPassword() *string
	SetPauseLogin(v int64) *ClinkCreateClientShrinkRequest
	GetPauseLogin() *int64
	SetPermissionShrink(v string) *ClinkCreateClientShrinkRequest
	GetPermissionShrink() *string
	SetQnosShrink(v string) *ClinkCreateClientShrinkRequest
	GetQnosShrink() *string
	SetRecurrentselectionType(v int64) *ClinkCreateClientShrinkRequest
	GetRecurrentselectionType() *int64
	SetRecurrentselectionValue(v int64) *ClinkCreateClientShrinkRequest
	GetRecurrentselectionValue() *int64
	SetResourceOwnerAccount(v string) *ClinkCreateClientShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkCreateClientShrinkRequest
	GetResourceOwnerId() *int64
	SetRole(v int64) *ClinkCreateClientShrinkRequest
	GetRole() *int64
	SetServeAreaShrink(v string) *ClinkCreateClientShrinkRequest
	GetServeAreaShrink() *string
	SetType(v int64) *ClinkCreateClientShrinkRequest
	GetType() *int64
	SetWechatMiniProgramRtc(v int64) *ClinkCreateClientShrinkRequest
	GetWechatMiniProgramRtc() *int64
	SetWrapupTime(v int64) *ClinkCreateClientShrinkRequest
	GetWrapupTime() *int64
}

type ClinkCreateClientShrinkRequest struct {
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
	ClidShrink *string `json:"Clid,omitempty" xml:"Clid,omitempty"`
	// 按地区外显配置，当clidType（外显号码类型）值为3且clidRule（外显规则）值为2（自定义）时有意义，具体配置项见 clidArea 对象
	ClidAreaShrink *string `json:"ClidArea,omitempty" xml:"ClidArea,omitempty"`
	// 缺省外显号码集合，当clidType（外显号码类型）值为3且clidRule（外显规则）值为1（动态）时，支持配置缺省外显
	ClidDefaultShrink *string `json:"ClidDefault,omitempty" xml:"ClidDefault,omitempty"`
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
	CloudNumberModesShrink *string `json:"CloudNumberModes,omitempty" xml:"CloudNumberModes,omitempty"`
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
	PermissionShrink *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// 所属队列号集合
	QnosShrink *string `json:"Qnos,omitempty" xml:"Qnos,omitempty"`
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
	ServeAreaShrink *string `json:"ServeArea,omitempty" xml:"ServeArea,omitempty"`
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

func (s ClinkCreateClientShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateClientShrinkRequest) GoString() string {
	return s.String()
}

func (s *ClinkCreateClientShrinkRequest) GetActive() *int64 {
	return s.Active
}

func (s *ClinkCreateClientShrinkRequest) GetAreaCode() *string {
	return s.AreaCode
}

func (s *ClinkCreateClientShrinkRequest) GetAssignType() *int64 {
	return s.AssignType
}

func (s *ClinkCreateClientShrinkRequest) GetClidShrink() *string {
	return s.ClidShrink
}

func (s *ClinkCreateClientShrinkRequest) GetClidAreaShrink() *string {
	return s.ClidAreaShrink
}

func (s *ClinkCreateClientShrinkRequest) GetClidDefaultShrink() *string {
	return s.ClidDefaultShrink
}

func (s *ClinkCreateClientShrinkRequest) GetClidRule() *int64 {
	return s.ClidRule
}

func (s *ClinkCreateClientShrinkRequest) GetClidType() *int64 {
	return s.ClidType
}

func (s *ClinkCreateClientShrinkRequest) GetCloudNumberEnabled() *int64 {
	return s.CloudNumberEnabled
}

func (s *ClinkCreateClientShrinkRequest) GetCloudNumberModesShrink() *string {
	return s.CloudNumberModesShrink
}

func (s *ClinkCreateClientShrinkRequest) GetCno() *int64 {
	return s.Cno
}

func (s *ClinkCreateClientShrinkRequest) GetCrmId() *int64 {
	return s.CrmId
}

func (s *ClinkCreateClientShrinkRequest) GetDynamicTelGroupIdDefault() *int64 {
	return s.DynamicTelGroupIdDefault
}

func (s *ClinkCreateClientShrinkRequest) GetDynamicTelGroupName() *string {
	return s.DynamicTelGroupName
}

func (s *ClinkCreateClientShrinkRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkCreateClientShrinkRequest) GetHiddenTel() *int64 {
	return s.HiddenTel
}

func (s *ClinkCreateClientShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ClinkCreateClientShrinkRequest) GetObClidDefaultType() *int64 {
	return s.ObClidDefaultType
}

func (s *ClinkCreateClientShrinkRequest) GetObHangupSms() *int64 {
	return s.ObHangupSms
}

func (s *ClinkCreateClientShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkCreateClientShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *ClinkCreateClientShrinkRequest) GetPauseLogin() *int64 {
	return s.PauseLogin
}

func (s *ClinkCreateClientShrinkRequest) GetPermissionShrink() *string {
	return s.PermissionShrink
}

func (s *ClinkCreateClientShrinkRequest) GetQnosShrink() *string {
	return s.QnosShrink
}

func (s *ClinkCreateClientShrinkRequest) GetRecurrentselectionType() *int64 {
	return s.RecurrentselectionType
}

func (s *ClinkCreateClientShrinkRequest) GetRecurrentselectionValue() *int64 {
	return s.RecurrentselectionValue
}

func (s *ClinkCreateClientShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkCreateClientShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkCreateClientShrinkRequest) GetRole() *int64 {
	return s.Role
}

func (s *ClinkCreateClientShrinkRequest) GetServeAreaShrink() *string {
	return s.ServeAreaShrink
}

func (s *ClinkCreateClientShrinkRequest) GetType() *int64 {
	return s.Type
}

func (s *ClinkCreateClientShrinkRequest) GetWechatMiniProgramRtc() *int64 {
	return s.WechatMiniProgramRtc
}

func (s *ClinkCreateClientShrinkRequest) GetWrapupTime() *int64 {
	return s.WrapupTime
}

func (s *ClinkCreateClientShrinkRequest) SetActive(v int64) *ClinkCreateClientShrinkRequest {
	s.Active = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetAreaCode(v string) *ClinkCreateClientShrinkRequest {
	s.AreaCode = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetAssignType(v int64) *ClinkCreateClientShrinkRequest {
	s.AssignType = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetClidShrink(v string) *ClinkCreateClientShrinkRequest {
	s.ClidShrink = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetClidAreaShrink(v string) *ClinkCreateClientShrinkRequest {
	s.ClidAreaShrink = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetClidDefaultShrink(v string) *ClinkCreateClientShrinkRequest {
	s.ClidDefaultShrink = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetClidRule(v int64) *ClinkCreateClientShrinkRequest {
	s.ClidRule = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetClidType(v int64) *ClinkCreateClientShrinkRequest {
	s.ClidType = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetCloudNumberEnabled(v int64) *ClinkCreateClientShrinkRequest {
	s.CloudNumberEnabled = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetCloudNumberModesShrink(v string) *ClinkCreateClientShrinkRequest {
	s.CloudNumberModesShrink = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetCno(v int64) *ClinkCreateClientShrinkRequest {
	s.Cno = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetCrmId(v int64) *ClinkCreateClientShrinkRequest {
	s.CrmId = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetDynamicTelGroupIdDefault(v int64) *ClinkCreateClientShrinkRequest {
	s.DynamicTelGroupIdDefault = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetDynamicTelGroupName(v string) *ClinkCreateClientShrinkRequest {
	s.DynamicTelGroupName = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetEnterpriseId(v int64) *ClinkCreateClientShrinkRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetHiddenTel(v int64) *ClinkCreateClientShrinkRequest {
	s.HiddenTel = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetName(v string) *ClinkCreateClientShrinkRequest {
	s.Name = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetObClidDefaultType(v int64) *ClinkCreateClientShrinkRequest {
	s.ObClidDefaultType = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetObHangupSms(v int64) *ClinkCreateClientShrinkRequest {
	s.ObHangupSms = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetOwnerId(v int64) *ClinkCreateClientShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetPassword(v string) *ClinkCreateClientShrinkRequest {
	s.Password = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetPauseLogin(v int64) *ClinkCreateClientShrinkRequest {
	s.PauseLogin = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetPermissionShrink(v string) *ClinkCreateClientShrinkRequest {
	s.PermissionShrink = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetQnosShrink(v string) *ClinkCreateClientShrinkRequest {
	s.QnosShrink = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetRecurrentselectionType(v int64) *ClinkCreateClientShrinkRequest {
	s.RecurrentselectionType = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetRecurrentselectionValue(v int64) *ClinkCreateClientShrinkRequest {
	s.RecurrentselectionValue = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetResourceOwnerAccount(v string) *ClinkCreateClientShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetResourceOwnerId(v int64) *ClinkCreateClientShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetRole(v int64) *ClinkCreateClientShrinkRequest {
	s.Role = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetServeAreaShrink(v string) *ClinkCreateClientShrinkRequest {
	s.ServeAreaShrink = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetType(v int64) *ClinkCreateClientShrinkRequest {
	s.Type = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetWechatMiniProgramRtc(v int64) *ClinkCreateClientShrinkRequest {
	s.WechatMiniProgramRtc = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) SetWrapupTime(v int64) *ClinkCreateClientShrinkRequest {
	s.WrapupTime = &v
	return s
}

func (s *ClinkCreateClientShrinkRequest) Validate() error {
	return dara.Validate(s)
}
