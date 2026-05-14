// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkUpdateClientShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v int64) *ClinkUpdateClientShrinkRequest
	GetActive() *int64
	SetAreaCode(v string) *ClinkUpdateClientShrinkRequest
	GetAreaCode() *string
	SetAssignType(v int64) *ClinkUpdateClientShrinkRequest
	GetAssignType() *int64
	SetClidShrink(v string) *ClinkUpdateClientShrinkRequest
	GetClidShrink() *string
	SetClidAreaShrink(v string) *ClinkUpdateClientShrinkRequest
	GetClidAreaShrink() *string
	SetClidDefaultShrink(v string) *ClinkUpdateClientShrinkRequest
	GetClidDefaultShrink() *string
	SetClidRule(v int64) *ClinkUpdateClientShrinkRequest
	GetClidRule() *int64
	SetClidType(v int64) *ClinkUpdateClientShrinkRequest
	GetClidType() *int64
	SetCloudNumberEnabled(v int64) *ClinkUpdateClientShrinkRequest
	GetCloudNumberEnabled() *int64
	SetCloudNumberModesShrink(v string) *ClinkUpdateClientShrinkRequest
	GetCloudNumberModesShrink() *string
	SetCno(v int64) *ClinkUpdateClientShrinkRequest
	GetCno() *int64
	SetCrmId(v int64) *ClinkUpdateClientShrinkRequest
	GetCrmId() *int64
	SetDynamicTelGroupIdDefault(v int64) *ClinkUpdateClientShrinkRequest
	GetDynamicTelGroupIdDefault() *int64
	SetDynamicTelGroupName(v string) *ClinkUpdateClientShrinkRequest
	GetDynamicTelGroupName() *string
	SetEnterpriseId(v int64) *ClinkUpdateClientShrinkRequest
	GetEnterpriseId() *int64
	SetHiddenTel(v int64) *ClinkUpdateClientShrinkRequest
	GetHiddenTel() *int64
	SetIbWrapupTime(v int64) *ClinkUpdateClientShrinkRequest
	GetIbWrapupTime() *int64
	SetIbWrapupType(v int64) *ClinkUpdateClientShrinkRequest
	GetIbWrapupType() *int64
	SetName(v string) *ClinkUpdateClientShrinkRequest
	GetName() *string
	SetObClidDefaultType(v int64) *ClinkUpdateClientShrinkRequest
	GetObClidDefaultType() *int64
	SetObHangupSms(v int64) *ClinkUpdateClientShrinkRequest
	GetObHangupSms() *int64
	SetOwnerId(v int64) *ClinkUpdateClientShrinkRequest
	GetOwnerId() *int64
	SetPassword(v string) *ClinkUpdateClientShrinkRequest
	GetPassword() *string
	SetPauseLogin(v int64) *ClinkUpdateClientShrinkRequest
	GetPauseLogin() *int64
	SetPermissionShrink(v string) *ClinkUpdateClientShrinkRequest
	GetPermissionShrink() *string
	SetQnosShrink(v string) *ClinkUpdateClientShrinkRequest
	GetQnosShrink() *string
	SetRecurrentselectionType(v int64) *ClinkUpdateClientShrinkRequest
	GetRecurrentselectionType() *int64
	SetRecurrentselectionValue(v int64) *ClinkUpdateClientShrinkRequest
	GetRecurrentselectionValue() *int64
	SetResourceOwnerAccount(v string) *ClinkUpdateClientShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkUpdateClientShrinkRequest
	GetResourceOwnerId() *int64
	SetRole(v int64) *ClinkUpdateClientShrinkRequest
	GetRole() *int64
	SetServeAreaShrink(v string) *ClinkUpdateClientShrinkRequest
	GetServeAreaShrink() *string
	SetWechatMiniProgramRtc(v int64) *ClinkUpdateClientShrinkRequest
	GetWechatMiniProgramRtc() *int64
	SetWrapupTime(v int64) *ClinkUpdateClientShrinkRequest
	GetWrapupTime() *int64
}

type ClinkUpdateClientShrinkRequest struct {
	// 启用状态， 0: 停用；1: 启用。
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
	// 可外显号码集合，当外显类型非全部时需指定此参数值。
	ClidShrink *string `json:"Clid,omitempty" xml:"Clid,omitempty"`
	// 按地区外显配置，当clidType（外显号码类型）值为3且clidRule（外显规则）值为2（自定义）时有意义，具体配置项见 clidArea 对象
	ClidAreaShrink *string `json:"ClidArea,omitempty" xml:"ClidArea,omitempty"`
	// 缺省外显号码集合，当clidType（外显号码类型）值为3且clidRule（外显规则）值为1（动态）时，支持配置缺省外显
	ClidDefaultShrink *string `json:"ClidDefault,omitempty" xml:"ClidDefault,omitempty"`
	// 外显规则， 当clidType（外显号码类型）值为0（全部）时，1:随机；2: 轮选。 当clidType（外显号码类型）值为3（智能外显）时，1:动态；2: 自定义。
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
	// 云号码四种呼叫模式；数组长度 为 4，依次对应云号码外呼的四种模式（实体卡、工作卡、两端呼、RTC)的开启状态； 例如：[0,0,0,1] 表示 座席只可使用 RTC 外呼模式。 注意：仅当企业开启了云手机外呼功能，才能为座席设置该属性，否则，座席将创建失败
	CloudNumberModesShrink *string `json:"CloudNumberModes,omitempty" xml:"CloudNumberModes,omitempty"`
	// 座席工号，4-11位数字，要求唯一
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
	// 95
	CrmId *int64 `json:"CrmId,omitempty" xml:"CrmId,omitempty"`
	// 缺省动态外呼组id，当obClidDefaultType值为1时为必填项
	//
	// example:
	//
	// 19
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
	// 号码隐藏类型，0: 不隐藏；1: 隐藏规则与全局设置保持一致。
	//
	// example:
	//
	// 0
	HiddenTel *int64 `json:"HiddenTel,omitempty" xml:"HiddenTel,omitempty"`
	// 呼入整理时长，客户来电座席接听，双方通话结束后座席的话后处理时长，此期间座席不接收新的客户来电。取值范围：0～3600秒。
	//
	// example:
	//
	// 30
	IbWrapupTime *int64 `json:"IbWrapupTime,omitempty" xml:"IbWrapupTime,omitempty"`
	// 呼入整理类型，1：队列；2：座席。
	//
	// example:
	//
	// 2
	IbWrapupType *int64 `json:"IbWrapupType,omitempty" xml:"IbWrapupType,omitempty"`
	// 座席名称，只允许输入中文，字母，数字，下划线，长度不超过 20 个字符
	//
	// example:
	//
	// 示例值示例值
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 缺省号码类型 0：号码 1：动态号码组 默认0
	//
	// example:
	//
	// 1
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
	// example:
	//
	// xxx
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// 是否可置忙登录，关闭后，座席不允许置忙登录，0：关闭，1：开启。
	//
	// example:
	//
	// 0
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
	// 座席角色，0: 普通座席；1: 班长座席。
	//
	// example:
	//
	// 1
	Role *int64 `json:"Role,omitempty" xml:"Role,omitempty"`
	// 座席服务地区，区号数组 配置座席服务的地区，可用于外呼任务中“按服务地区分配”的策略。
	ServeAreaShrink *string `json:"ServeArea,omitempty" xml:"ServeArea,omitempty"`
	// 微信小程序RTC开关 1：开启 0：关闭 默认关闭
	//
	// example:
	//
	// 1
	WechatMiniProgramRtc *int64 `json:"WechatMiniProgramRtc,omitempty" xml:"WechatMiniProgramRtc,omitempty"`
	// 整理时长，座席进行外呼操作后的整理时间，取值范围 0-300 秒。
	//
	// example:
	//
	// 51
	WrapupTime *int64 `json:"WrapupTime,omitempty" xml:"WrapupTime,omitempty"`
}

func (s ClinkUpdateClientShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkUpdateClientShrinkRequest) GoString() string {
	return s.String()
}

func (s *ClinkUpdateClientShrinkRequest) GetActive() *int64 {
	return s.Active
}

func (s *ClinkUpdateClientShrinkRequest) GetAreaCode() *string {
	return s.AreaCode
}

func (s *ClinkUpdateClientShrinkRequest) GetAssignType() *int64 {
	return s.AssignType
}

func (s *ClinkUpdateClientShrinkRequest) GetClidShrink() *string {
	return s.ClidShrink
}

func (s *ClinkUpdateClientShrinkRequest) GetClidAreaShrink() *string {
	return s.ClidAreaShrink
}

func (s *ClinkUpdateClientShrinkRequest) GetClidDefaultShrink() *string {
	return s.ClidDefaultShrink
}

func (s *ClinkUpdateClientShrinkRequest) GetClidRule() *int64 {
	return s.ClidRule
}

func (s *ClinkUpdateClientShrinkRequest) GetClidType() *int64 {
	return s.ClidType
}

func (s *ClinkUpdateClientShrinkRequest) GetCloudNumberEnabled() *int64 {
	return s.CloudNumberEnabled
}

func (s *ClinkUpdateClientShrinkRequest) GetCloudNumberModesShrink() *string {
	return s.CloudNumberModesShrink
}

func (s *ClinkUpdateClientShrinkRequest) GetCno() *int64 {
	return s.Cno
}

func (s *ClinkUpdateClientShrinkRequest) GetCrmId() *int64 {
	return s.CrmId
}

func (s *ClinkUpdateClientShrinkRequest) GetDynamicTelGroupIdDefault() *int64 {
	return s.DynamicTelGroupIdDefault
}

func (s *ClinkUpdateClientShrinkRequest) GetDynamicTelGroupName() *string {
	return s.DynamicTelGroupName
}

func (s *ClinkUpdateClientShrinkRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkUpdateClientShrinkRequest) GetHiddenTel() *int64 {
	return s.HiddenTel
}

func (s *ClinkUpdateClientShrinkRequest) GetIbWrapupTime() *int64 {
	return s.IbWrapupTime
}

func (s *ClinkUpdateClientShrinkRequest) GetIbWrapupType() *int64 {
	return s.IbWrapupType
}

func (s *ClinkUpdateClientShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ClinkUpdateClientShrinkRequest) GetObClidDefaultType() *int64 {
	return s.ObClidDefaultType
}

func (s *ClinkUpdateClientShrinkRequest) GetObHangupSms() *int64 {
	return s.ObHangupSms
}

func (s *ClinkUpdateClientShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkUpdateClientShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *ClinkUpdateClientShrinkRequest) GetPauseLogin() *int64 {
	return s.PauseLogin
}

func (s *ClinkUpdateClientShrinkRequest) GetPermissionShrink() *string {
	return s.PermissionShrink
}

func (s *ClinkUpdateClientShrinkRequest) GetQnosShrink() *string {
	return s.QnosShrink
}

func (s *ClinkUpdateClientShrinkRequest) GetRecurrentselectionType() *int64 {
	return s.RecurrentselectionType
}

func (s *ClinkUpdateClientShrinkRequest) GetRecurrentselectionValue() *int64 {
	return s.RecurrentselectionValue
}

func (s *ClinkUpdateClientShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkUpdateClientShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkUpdateClientShrinkRequest) GetRole() *int64 {
	return s.Role
}

func (s *ClinkUpdateClientShrinkRequest) GetServeAreaShrink() *string {
	return s.ServeAreaShrink
}

func (s *ClinkUpdateClientShrinkRequest) GetWechatMiniProgramRtc() *int64 {
	return s.WechatMiniProgramRtc
}

func (s *ClinkUpdateClientShrinkRequest) GetWrapupTime() *int64 {
	return s.WrapupTime
}

func (s *ClinkUpdateClientShrinkRequest) SetActive(v int64) *ClinkUpdateClientShrinkRequest {
	s.Active = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetAreaCode(v string) *ClinkUpdateClientShrinkRequest {
	s.AreaCode = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetAssignType(v int64) *ClinkUpdateClientShrinkRequest {
	s.AssignType = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetClidShrink(v string) *ClinkUpdateClientShrinkRequest {
	s.ClidShrink = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetClidAreaShrink(v string) *ClinkUpdateClientShrinkRequest {
	s.ClidAreaShrink = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetClidDefaultShrink(v string) *ClinkUpdateClientShrinkRequest {
	s.ClidDefaultShrink = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetClidRule(v int64) *ClinkUpdateClientShrinkRequest {
	s.ClidRule = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetClidType(v int64) *ClinkUpdateClientShrinkRequest {
	s.ClidType = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetCloudNumberEnabled(v int64) *ClinkUpdateClientShrinkRequest {
	s.CloudNumberEnabled = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetCloudNumberModesShrink(v string) *ClinkUpdateClientShrinkRequest {
	s.CloudNumberModesShrink = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetCno(v int64) *ClinkUpdateClientShrinkRequest {
	s.Cno = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetCrmId(v int64) *ClinkUpdateClientShrinkRequest {
	s.CrmId = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetDynamicTelGroupIdDefault(v int64) *ClinkUpdateClientShrinkRequest {
	s.DynamicTelGroupIdDefault = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetDynamicTelGroupName(v string) *ClinkUpdateClientShrinkRequest {
	s.DynamicTelGroupName = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetEnterpriseId(v int64) *ClinkUpdateClientShrinkRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetHiddenTel(v int64) *ClinkUpdateClientShrinkRequest {
	s.HiddenTel = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetIbWrapupTime(v int64) *ClinkUpdateClientShrinkRequest {
	s.IbWrapupTime = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetIbWrapupType(v int64) *ClinkUpdateClientShrinkRequest {
	s.IbWrapupType = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetName(v string) *ClinkUpdateClientShrinkRequest {
	s.Name = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetObClidDefaultType(v int64) *ClinkUpdateClientShrinkRequest {
	s.ObClidDefaultType = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetObHangupSms(v int64) *ClinkUpdateClientShrinkRequest {
	s.ObHangupSms = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetOwnerId(v int64) *ClinkUpdateClientShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetPassword(v string) *ClinkUpdateClientShrinkRequest {
	s.Password = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetPauseLogin(v int64) *ClinkUpdateClientShrinkRequest {
	s.PauseLogin = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetPermissionShrink(v string) *ClinkUpdateClientShrinkRequest {
	s.PermissionShrink = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetQnosShrink(v string) *ClinkUpdateClientShrinkRequest {
	s.QnosShrink = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetRecurrentselectionType(v int64) *ClinkUpdateClientShrinkRequest {
	s.RecurrentselectionType = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetRecurrentselectionValue(v int64) *ClinkUpdateClientShrinkRequest {
	s.RecurrentselectionValue = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetResourceOwnerAccount(v string) *ClinkUpdateClientShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetResourceOwnerId(v int64) *ClinkUpdateClientShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetRole(v int64) *ClinkUpdateClientShrinkRequest {
	s.Role = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetServeAreaShrink(v string) *ClinkUpdateClientShrinkRequest {
	s.ServeAreaShrink = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetWechatMiniProgramRtc(v int64) *ClinkUpdateClientShrinkRequest {
	s.WechatMiniProgramRtc = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) SetWrapupTime(v int64) *ClinkUpdateClientShrinkRequest {
	s.WrapupTime = &v
	return s
}

func (s *ClinkUpdateClientShrinkRequest) Validate() error {
	return dara.Validate(s)
}
