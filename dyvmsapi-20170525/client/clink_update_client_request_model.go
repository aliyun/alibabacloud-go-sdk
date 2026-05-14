// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkUpdateClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v int64) *ClinkUpdateClientRequest
	GetActive() *int64
	SetAreaCode(v string) *ClinkUpdateClientRequest
	GetAreaCode() *string
	SetAssignType(v int64) *ClinkUpdateClientRequest
	GetAssignType() *int64
	SetClid(v []*string) *ClinkUpdateClientRequest
	GetClid() []*string
	SetClidArea(v []*ClinkUpdateClientRequestClidArea) *ClinkUpdateClientRequest
	GetClidArea() []*ClinkUpdateClientRequestClidArea
	SetClidDefault(v []*string) *ClinkUpdateClientRequest
	GetClidDefault() []*string
	SetClidRule(v int64) *ClinkUpdateClientRequest
	GetClidRule() *int64
	SetClidType(v int64) *ClinkUpdateClientRequest
	GetClidType() *int64
	SetCloudNumberEnabled(v int64) *ClinkUpdateClientRequest
	GetCloudNumberEnabled() *int64
	SetCloudNumberModes(v []*int64) *ClinkUpdateClientRequest
	GetCloudNumberModes() []*int64
	SetCno(v int64) *ClinkUpdateClientRequest
	GetCno() *int64
	SetCrmId(v int64) *ClinkUpdateClientRequest
	GetCrmId() *int64
	SetDynamicTelGroupIdDefault(v int64) *ClinkUpdateClientRequest
	GetDynamicTelGroupIdDefault() *int64
	SetDynamicTelGroupName(v string) *ClinkUpdateClientRequest
	GetDynamicTelGroupName() *string
	SetEnterpriseId(v int64) *ClinkUpdateClientRequest
	GetEnterpriseId() *int64
	SetHiddenTel(v int64) *ClinkUpdateClientRequest
	GetHiddenTel() *int64
	SetIbWrapupTime(v int64) *ClinkUpdateClientRequest
	GetIbWrapupTime() *int64
	SetIbWrapupType(v int64) *ClinkUpdateClientRequest
	GetIbWrapupType() *int64
	SetName(v string) *ClinkUpdateClientRequest
	GetName() *string
	SetObClidDefaultType(v int64) *ClinkUpdateClientRequest
	GetObClidDefaultType() *int64
	SetObHangupSms(v int64) *ClinkUpdateClientRequest
	GetObHangupSms() *int64
	SetOwnerId(v int64) *ClinkUpdateClientRequest
	GetOwnerId() *int64
	SetPassword(v string) *ClinkUpdateClientRequest
	GetPassword() *string
	SetPauseLogin(v int64) *ClinkUpdateClientRequest
	GetPauseLogin() *int64
	SetPermission(v *ClinkUpdateClientRequestPermission) *ClinkUpdateClientRequest
	GetPermission() *ClinkUpdateClientRequestPermission
	SetQnos(v []*string) *ClinkUpdateClientRequest
	GetQnos() []*string
	SetRecurrentselectionType(v int64) *ClinkUpdateClientRequest
	GetRecurrentselectionType() *int64
	SetRecurrentselectionValue(v int64) *ClinkUpdateClientRequest
	GetRecurrentselectionValue() *int64
	SetResourceOwnerAccount(v string) *ClinkUpdateClientRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkUpdateClientRequest
	GetResourceOwnerId() *int64
	SetRole(v int64) *ClinkUpdateClientRequest
	GetRole() *int64
	SetServeArea(v []*string) *ClinkUpdateClientRequest
	GetServeArea() []*string
	SetWechatMiniProgramRtc(v int64) *ClinkUpdateClientRequest
	GetWechatMiniProgramRtc() *int64
	SetWrapupTime(v int64) *ClinkUpdateClientRequest
	GetWrapupTime() *int64
}

type ClinkUpdateClientRequest struct {
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
	Clid []*string `json:"Clid,omitempty" xml:"Clid,omitempty" type:"Repeated"`
	// 按地区外显配置，当clidType（外显号码类型）值为3且clidRule（外显规则）值为2（自定义）时有意义，具体配置项见 clidArea 对象
	ClidArea []*ClinkUpdateClientRequestClidArea `json:"ClidArea,omitempty" xml:"ClidArea,omitempty" type:"Repeated"`
	// 缺省外显号码集合，当clidType（外显号码类型）值为3且clidRule（外显规则）值为1（动态）时，支持配置缺省外显
	ClidDefault []*string `json:"ClidDefault,omitempty" xml:"ClidDefault,omitempty" type:"Repeated"`
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
	CloudNumberModes []*int64 `json:"CloudNumberModes,omitempty" xml:"CloudNumberModes,omitempty" type:"Repeated"`
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
	Permission *ClinkUpdateClientRequestPermission `json:"Permission,omitempty" xml:"Permission,omitempty" type:"Struct"`
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
	// 座席角色，0: 普通座席；1: 班长座席。
	//
	// example:
	//
	// 1
	Role *int64 `json:"Role,omitempty" xml:"Role,omitempty"`
	// 座席服务地区，区号数组 配置座席服务的地区，可用于外呼任务中“按服务地区分配”的策略。
	ServeArea []*string `json:"ServeArea,omitempty" xml:"ServeArea,omitempty" type:"Repeated"`
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

func (s ClinkUpdateClientRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkUpdateClientRequest) GoString() string {
	return s.String()
}

func (s *ClinkUpdateClientRequest) GetActive() *int64 {
	return s.Active
}

func (s *ClinkUpdateClientRequest) GetAreaCode() *string {
	return s.AreaCode
}

func (s *ClinkUpdateClientRequest) GetAssignType() *int64 {
	return s.AssignType
}

func (s *ClinkUpdateClientRequest) GetClid() []*string {
	return s.Clid
}

func (s *ClinkUpdateClientRequest) GetClidArea() []*ClinkUpdateClientRequestClidArea {
	return s.ClidArea
}

func (s *ClinkUpdateClientRequest) GetClidDefault() []*string {
	return s.ClidDefault
}

func (s *ClinkUpdateClientRequest) GetClidRule() *int64 {
	return s.ClidRule
}

func (s *ClinkUpdateClientRequest) GetClidType() *int64 {
	return s.ClidType
}

func (s *ClinkUpdateClientRequest) GetCloudNumberEnabled() *int64 {
	return s.CloudNumberEnabled
}

func (s *ClinkUpdateClientRequest) GetCloudNumberModes() []*int64 {
	return s.CloudNumberModes
}

func (s *ClinkUpdateClientRequest) GetCno() *int64 {
	return s.Cno
}

func (s *ClinkUpdateClientRequest) GetCrmId() *int64 {
	return s.CrmId
}

func (s *ClinkUpdateClientRequest) GetDynamicTelGroupIdDefault() *int64 {
	return s.DynamicTelGroupIdDefault
}

func (s *ClinkUpdateClientRequest) GetDynamicTelGroupName() *string {
	return s.DynamicTelGroupName
}

func (s *ClinkUpdateClientRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkUpdateClientRequest) GetHiddenTel() *int64 {
	return s.HiddenTel
}

func (s *ClinkUpdateClientRequest) GetIbWrapupTime() *int64 {
	return s.IbWrapupTime
}

func (s *ClinkUpdateClientRequest) GetIbWrapupType() *int64 {
	return s.IbWrapupType
}

func (s *ClinkUpdateClientRequest) GetName() *string {
	return s.Name
}

func (s *ClinkUpdateClientRequest) GetObClidDefaultType() *int64 {
	return s.ObClidDefaultType
}

func (s *ClinkUpdateClientRequest) GetObHangupSms() *int64 {
	return s.ObHangupSms
}

func (s *ClinkUpdateClientRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkUpdateClientRequest) GetPassword() *string {
	return s.Password
}

func (s *ClinkUpdateClientRequest) GetPauseLogin() *int64 {
	return s.PauseLogin
}

func (s *ClinkUpdateClientRequest) GetPermission() *ClinkUpdateClientRequestPermission {
	return s.Permission
}

func (s *ClinkUpdateClientRequest) GetQnos() []*string {
	return s.Qnos
}

func (s *ClinkUpdateClientRequest) GetRecurrentselectionType() *int64 {
	return s.RecurrentselectionType
}

func (s *ClinkUpdateClientRequest) GetRecurrentselectionValue() *int64 {
	return s.RecurrentselectionValue
}

func (s *ClinkUpdateClientRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkUpdateClientRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkUpdateClientRequest) GetRole() *int64 {
	return s.Role
}

func (s *ClinkUpdateClientRequest) GetServeArea() []*string {
	return s.ServeArea
}

func (s *ClinkUpdateClientRequest) GetWechatMiniProgramRtc() *int64 {
	return s.WechatMiniProgramRtc
}

func (s *ClinkUpdateClientRequest) GetWrapupTime() *int64 {
	return s.WrapupTime
}

func (s *ClinkUpdateClientRequest) SetActive(v int64) *ClinkUpdateClientRequest {
	s.Active = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetAreaCode(v string) *ClinkUpdateClientRequest {
	s.AreaCode = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetAssignType(v int64) *ClinkUpdateClientRequest {
	s.AssignType = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetClid(v []*string) *ClinkUpdateClientRequest {
	s.Clid = v
	return s
}

func (s *ClinkUpdateClientRequest) SetClidArea(v []*ClinkUpdateClientRequestClidArea) *ClinkUpdateClientRequest {
	s.ClidArea = v
	return s
}

func (s *ClinkUpdateClientRequest) SetClidDefault(v []*string) *ClinkUpdateClientRequest {
	s.ClidDefault = v
	return s
}

func (s *ClinkUpdateClientRequest) SetClidRule(v int64) *ClinkUpdateClientRequest {
	s.ClidRule = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetClidType(v int64) *ClinkUpdateClientRequest {
	s.ClidType = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetCloudNumberEnabled(v int64) *ClinkUpdateClientRequest {
	s.CloudNumberEnabled = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetCloudNumberModes(v []*int64) *ClinkUpdateClientRequest {
	s.CloudNumberModes = v
	return s
}

func (s *ClinkUpdateClientRequest) SetCno(v int64) *ClinkUpdateClientRequest {
	s.Cno = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetCrmId(v int64) *ClinkUpdateClientRequest {
	s.CrmId = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetDynamicTelGroupIdDefault(v int64) *ClinkUpdateClientRequest {
	s.DynamicTelGroupIdDefault = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetDynamicTelGroupName(v string) *ClinkUpdateClientRequest {
	s.DynamicTelGroupName = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetEnterpriseId(v int64) *ClinkUpdateClientRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetHiddenTel(v int64) *ClinkUpdateClientRequest {
	s.HiddenTel = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetIbWrapupTime(v int64) *ClinkUpdateClientRequest {
	s.IbWrapupTime = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetIbWrapupType(v int64) *ClinkUpdateClientRequest {
	s.IbWrapupType = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetName(v string) *ClinkUpdateClientRequest {
	s.Name = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetObClidDefaultType(v int64) *ClinkUpdateClientRequest {
	s.ObClidDefaultType = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetObHangupSms(v int64) *ClinkUpdateClientRequest {
	s.ObHangupSms = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetOwnerId(v int64) *ClinkUpdateClientRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetPassword(v string) *ClinkUpdateClientRequest {
	s.Password = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetPauseLogin(v int64) *ClinkUpdateClientRequest {
	s.PauseLogin = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetPermission(v *ClinkUpdateClientRequestPermission) *ClinkUpdateClientRequest {
	s.Permission = v
	return s
}

func (s *ClinkUpdateClientRequest) SetQnos(v []*string) *ClinkUpdateClientRequest {
	s.Qnos = v
	return s
}

func (s *ClinkUpdateClientRequest) SetRecurrentselectionType(v int64) *ClinkUpdateClientRequest {
	s.RecurrentselectionType = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetRecurrentselectionValue(v int64) *ClinkUpdateClientRequest {
	s.RecurrentselectionValue = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetResourceOwnerAccount(v string) *ClinkUpdateClientRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetResourceOwnerId(v int64) *ClinkUpdateClientRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetRole(v int64) *ClinkUpdateClientRequest {
	s.Role = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetServeArea(v []*string) *ClinkUpdateClientRequest {
	s.ServeArea = v
	return s
}

func (s *ClinkUpdateClientRequest) SetWechatMiniProgramRtc(v int64) *ClinkUpdateClientRequest {
	s.WechatMiniProgramRtc = &v
	return s
}

func (s *ClinkUpdateClientRequest) SetWrapupTime(v int64) *ClinkUpdateClientRequest {
	s.WrapupTime = &v
	return s
}

func (s *ClinkUpdateClientRequest) Validate() error {
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

type ClinkUpdateClientRequestClidArea struct {
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

func (s ClinkUpdateClientRequestClidArea) String() string {
	return dara.Prettify(s)
}

func (s ClinkUpdateClientRequestClidArea) GoString() string {
	return s.String()
}

func (s *ClinkUpdateClientRequestClidArea) GetAreaGroupName() *string {
	return s.AreaGroupName
}

func (s *ClinkUpdateClientRequestClidArea) GetAssignType() *int64 {
	return s.AssignType
}

func (s *ClinkUpdateClientRequestClidArea) GetDynamicTelGroupName() *string {
	return s.DynamicTelGroupName
}

func (s *ClinkUpdateClientRequestClidArea) GetObClids() []*string {
	return s.ObClids
}

func (s *ClinkUpdateClientRequestClidArea) SetAreaGroupName(v string) *ClinkUpdateClientRequestClidArea {
	s.AreaGroupName = &v
	return s
}

func (s *ClinkUpdateClientRequestClidArea) SetAssignType(v int64) *ClinkUpdateClientRequestClidArea {
	s.AssignType = &v
	return s
}

func (s *ClinkUpdateClientRequestClidArea) SetDynamicTelGroupName(v string) *ClinkUpdateClientRequestClidArea {
	s.DynamicTelGroupName = &v
	return s
}

func (s *ClinkUpdateClientRequestClidArea) SetObClids(v []*string) *ClinkUpdateClientRequestClidArea {
	s.ObClids = v
	return s
}

func (s *ClinkUpdateClientRequestClidArea) Validate() error {
	return dara.Validate(s)
}

type ClinkUpdateClientRequestPermission struct {
	// 语音转写，0: 关闭；1: 呼入开启；2 外呼开启；3 全部开启；默认值为 0
	//
	// example:
	//
	// 1
	Asr *int64 `json:"Asr,omitempty" xml:"Asr,omitempty"`
	// 外呼权限，0:关闭；1: 无限制；2: 国内长途；3: 国内本地
	//
	// example:
	//
	// 1
	Call *int64 `json:"Call,omitempty" xml:"Call,omitempty"`
	// 通话记录查看权限，1: 全部；2: 所属队列；3: 本座席
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
	// 通话录音权限，0: 关闭；1: 呼入；2: 外呼；3: 全部
	//
	// example:
	//
	// 1
	Record *int64 `json:"Record,omitempty" xml:"Record,omitempty"`
	// 录音试听下载权限， 0: 关闭；1: 试听；2: 试听下载
	//
	// example:
	//
	// 1
	RecordDownload *int64 `json:"RecordDownload,omitempty" xml:"RecordDownload,omitempty"`
	// 短信发送权限，0: 关闭；1: 开启。默认值为 0
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

func (s ClinkUpdateClientRequestPermission) String() string {
	return dara.Prettify(s)
}

func (s ClinkUpdateClientRequestPermission) GoString() string {
	return s.String()
}

func (s *ClinkUpdateClientRequestPermission) GetAsr() *int64 {
	return s.Asr
}

func (s *ClinkUpdateClientRequestPermission) GetCall() *int64 {
	return s.Call
}

func (s *ClinkUpdateClientRequestPermission) GetCdr() *int64 {
	return s.Cdr
}

func (s *ClinkUpdateClientRequestPermission) GetChat() *int64 {
	return s.Chat
}

func (s *ClinkUpdateClientRequestPermission) GetOtherData() *int64 {
	return s.OtherData
}

func (s *ClinkUpdateClientRequestPermission) GetRecord() *int64 {
	return s.Record
}

func (s *ClinkUpdateClientRequestPermission) GetRecordDownload() *int64 {
	return s.RecordDownload
}

func (s *ClinkUpdateClientRequestPermission) GetSms() *int64 {
	return s.Sms
}

func (s *ClinkUpdateClientRequestPermission) GetTaskInventory() *int64 {
	return s.TaskInventory
}

func (s *ClinkUpdateClientRequestPermission) GetTransfer() *int64 {
	return s.Transfer
}

func (s *ClinkUpdateClientRequestPermission) SetAsr(v int64) *ClinkUpdateClientRequestPermission {
	s.Asr = &v
	return s
}

func (s *ClinkUpdateClientRequestPermission) SetCall(v int64) *ClinkUpdateClientRequestPermission {
	s.Call = &v
	return s
}

func (s *ClinkUpdateClientRequestPermission) SetCdr(v int64) *ClinkUpdateClientRequestPermission {
	s.Cdr = &v
	return s
}

func (s *ClinkUpdateClientRequestPermission) SetChat(v int64) *ClinkUpdateClientRequestPermission {
	s.Chat = &v
	return s
}

func (s *ClinkUpdateClientRequestPermission) SetOtherData(v int64) *ClinkUpdateClientRequestPermission {
	s.OtherData = &v
	return s
}

func (s *ClinkUpdateClientRequestPermission) SetRecord(v int64) *ClinkUpdateClientRequestPermission {
	s.Record = &v
	return s
}

func (s *ClinkUpdateClientRequestPermission) SetRecordDownload(v int64) *ClinkUpdateClientRequestPermission {
	s.RecordDownload = &v
	return s
}

func (s *ClinkUpdateClientRequestPermission) SetSms(v int64) *ClinkUpdateClientRequestPermission {
	s.Sms = &v
	return s
}

func (s *ClinkUpdateClientRequestPermission) SetTaskInventory(v int64) *ClinkUpdateClientRequestPermission {
	s.TaskInventory = &v
	return s
}

func (s *ClinkUpdateClientRequestPermission) SetTransfer(v int64) *ClinkUpdateClientRequestPermission {
	s.Transfer = &v
	return s
}

func (s *ClinkUpdateClientRequestPermission) Validate() error {
	return dara.Validate(s)
}
