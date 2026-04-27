// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudUpdateAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v int64) *CloudUpdateAgentRequest
	GetActive() *int64
	SetAgentType(v int64) *CloudUpdateAgentRequest
	GetAgentType() *int64
	SetAreaCode(v string) *CloudUpdateAgentRequest
	GetAreaCode() *string
	SetCallPower(v int64) *CloudUpdateAgentRequest
	GetCallPower() *int64
	SetCno(v string) *CloudUpdateAgentRequest
	GetCno() *string
	SetComment(v string) *CloudUpdateAgentRequest
	GetComment() *string
	SetEnterpriseId(v int64) *CloudUpdateAgentRequest
	GetEnterpriseId() *int64
	SetIbRecord(v int64) *CloudUpdateAgentRequest
	GetIbRecord() *int64
	SetIsAsr(v int64) *CloudUpdateAgentRequest
	GetIsAsr() *int64
	SetIsOb(v int64) *CloudUpdateAgentRequest
	GetIsOb() *int64
	SetIsObRemember(v string) *CloudUpdateAgentRequest
	GetIsObRemember() *string
	SetIsQualityCheck(v int64) *CloudUpdateAgentRequest
	GetIsQualityCheck() *int64
	SetIsRandom(v string) *CloudUpdateAgentRequest
	GetIsRandom() *string
	SetName(v string) *CloudUpdateAgentRequest
	GetName() *string
	SetObClid(v string) *CloudUpdateAgentRequest
	GetObClid() *string
	SetObClidProperty(v string) *CloudUpdateAgentRequest
	GetObClidProperty() *string
	SetObClidType(v int64) *CloudUpdateAgentRequest
	GetObClidType() *int64
	SetObRecord(v int64) *CloudUpdateAgentRequest
	GetObRecord() *int64
	SetOwnerId(v int64) *CloudUpdateAgentRequest
	GetOwnerId() *int64
	SetPermitObPreviewTime(v string) *CloudUpdateAgentRequest
	GetPermitObPreviewTime() *string
	SetPower(v int64) *CloudUpdateAgentRequest
	GetPower() *int64
	SetResourceOwnerAccount(v string) *CloudUpdateAgentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudUpdateAgentRequest
	GetResourceOwnerId() *int64
	SetSkillIds(v string) *CloudUpdateAgentRequest
	GetSkillIds() *string
	SetSkillLevels(v string) *CloudUpdateAgentRequest
	GetSkillLevels() *string
	SetWebrtcUrlType(v int64) *CloudUpdateAgentRequest
	GetWebrtcUrlType() *int64
	SetWrapup(v int64) *CloudUpdateAgentRequest
	GetWrapup() *int64
}

type CloudUpdateAgentRequest struct {
	// 是否激活；取值0或1，取值说明 0：不激活，1激活，默认激活
	//
	// example:
	//
	// 1
	Active *int64 `json:"Active,omitempty" xml:"Active,omitempty"`
	// 座席类型；取值1或2，取值说明 1：电话座席，2：电脑座席，默认电话座席
	//
	// example:
	//
	// 1
	AgentType *int64 `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// 座席所属区号；区号格式
	//
	// example:
	//
	// 010
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// 呼叫权限；取值说明 0：无限制，1：国内长途，2：国内本市，3：内部呼叫，默认无限制
	//
	// example:
	//
	// 1
	CallPower *int64 `json:"CallPower,omitempty" xml:"CallPower,omitempty"`
	// 座席工号；正整数，取值3-10位数字
	//
	// This parameter is required.
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 备注
	//
	// example:
	//
	// 示例值
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 呼入是否录音；取值说明 0：不录用，1：录音，默认录音
	//
	// example:
	//
	// 1
	IbRecord *int64 `json:"IbRecord,omitempty" xml:"IbRecord,omitempty"`
	// 是否开启ASR转写；取值说明：0：不开启，1：开启，默认不开启
	//
	// example:
	//
	// 1
	IsAsr *int64 `json:"IsAsr,omitempty" xml:"IsAsr,omitempty"`
	// 是否允许外呼；取值说明 0：不允许，1：可以，默认允许
	//
	// example:
	//
	// 1
	IsOb *int64 `json:"IsOb,omitempty" xml:"IsOb,omitempty"`
	// 外呼主叫记忆开关；取值说明：0：关闭 1：开启；默认开启
	//
	// example:
	//
	// 1
	IsObRemember *string `json:"IsObRemember,omitempty" xml:"IsObRemember,omitempty"`
	// 是否开启座席质检；取值说明：0：不开启，1：开启，默认开启
	//
	// example:
	//
	// 1
	IsQualityCheck *int64 `json:"IsQualityCheck,omitempty" xml:"IsQualityCheck,omitempty"`
	// 随机外显；取值说明 0:否，1:是。当obClidType值为4时需要传值，默认值为0
	//
	// example:
	//
	// 1
	IsRandom *string `json:"IsRandom,omitempty" xml:"IsRandom,omitempty"`
	// 座席姓名；需进行UTF-8格式的URLEncode编码
	//
	// example:
	//
	// 示例值示例值示例值
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 透传号码；可传入企业中继号码或设置好的客户侧外显号码，当obClidType值为2或3时必选
	//
	// example:
	//
	// 22223333
	ObClid *string `json:"ObClid,omitempty" xml:"ObClid,omitempty"`
	// 外显属性；取值：{"isMatchCapital":0,"areaCodeRule":1}；obClidType=4时，isMatchCapital表示是否匹配省会号码，1是 0否，areaCodeRule表示区号匹配规则，1座席区号 2客户号码区号（注意：json格式数据需要URLencode）
	//
	// example:
	//
	// {"isMatchCapital":0,"areaCodeRule":1}
	ObClidProperty *string `json:"ObClidProperty,omitempty" xml:"ObClidProperty,omitempty"`
	// 外显规则；取值：1:默认 2:随机 3:按区号 4 动态外显 ，默认值为1
	//
	// example:
	//
	// 1
	ObClidType *int64 `json:"ObClidType,omitempty" xml:"ObClidType,omitempty"`
	// 外呼是否录音；取值说明 0：不录音，1：录音，默认录音
	//
	// example:
	//
	// 1
	ObRecord *int64 `json:"ObRecord,omitempty" xml:"ObRecord,omitempty"`
	OwnerId  *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 可外呼时间段；格式：08:00,20:00
	//
	// example:
	//
	// 08:00,20:00
	PermitObPreviewTime *string `json:"PermitObPreviewTime,omitempty" xml:"PermitObPreviewTime,omitempty"`
	// 座席权限；取值1或0，取值说明 1：班长席，0：普通座席，默认为0 普通座席
	//
	// example:
	//
	// 1
	Power                *int64  `json:"Power,omitempty" xml:"Power,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 所选的技能id；可选择多个，多个之间使用英文标点逗号间隔.例如:"1,2,3"；skillIds为"0"时代表清空座席所有技能，不为"0"时需要同时传入skillLevels
	//
	// example:
	//
	// 1,2,3
	SkillIds *string `json:"SkillIds,omitempty" xml:"SkillIds,omitempty"`
	// 所选的技能的等级；值越小技能越高，多个间用英文标点逗号间隔。(与技能id相对应) 例如:"5,1,1"表示技能id为1的等级是5; 技能id为2的等级是1; 技能id为3的等级是1.
	//
	// example:
	//
	// 5,1,1
	SkillLevels *string `json:"SkillLevels,omitempty" xml:"SkillLevels,omitempty"`
	// webrtc软电话返回地址；取值说明：0：企业默认 1：公网域名 2：专线域名 3：公网IP 4：专线IP
	//
	// example:
	//
	// 1
	WebrtcUrlType *int64 `json:"WebrtcUrlType,omitempty" xml:"WebrtcUrlType,omitempty"`
	// 整理时间；单位秒数，默认10秒
	//
	// example:
	//
	// 22
	Wrapup *int64 `json:"Wrapup,omitempty" xml:"Wrapup,omitempty"`
}

func (s CloudUpdateAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudUpdateAgentRequest) GoString() string {
	return s.String()
}

func (s *CloudUpdateAgentRequest) GetActive() *int64 {
	return s.Active
}

func (s *CloudUpdateAgentRequest) GetAgentType() *int64 {
	return s.AgentType
}

func (s *CloudUpdateAgentRequest) GetAreaCode() *string {
	return s.AreaCode
}

func (s *CloudUpdateAgentRequest) GetCallPower() *int64 {
	return s.CallPower
}

func (s *CloudUpdateAgentRequest) GetCno() *string {
	return s.Cno
}

func (s *CloudUpdateAgentRequest) GetComment() *string {
	return s.Comment
}

func (s *CloudUpdateAgentRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudUpdateAgentRequest) GetIbRecord() *int64 {
	return s.IbRecord
}

func (s *CloudUpdateAgentRequest) GetIsAsr() *int64 {
	return s.IsAsr
}

func (s *CloudUpdateAgentRequest) GetIsOb() *int64 {
	return s.IsOb
}

func (s *CloudUpdateAgentRequest) GetIsObRemember() *string {
	return s.IsObRemember
}

func (s *CloudUpdateAgentRequest) GetIsQualityCheck() *int64 {
	return s.IsQualityCheck
}

func (s *CloudUpdateAgentRequest) GetIsRandom() *string {
	return s.IsRandom
}

func (s *CloudUpdateAgentRequest) GetName() *string {
	return s.Name
}

func (s *CloudUpdateAgentRequest) GetObClid() *string {
	return s.ObClid
}

func (s *CloudUpdateAgentRequest) GetObClidProperty() *string {
	return s.ObClidProperty
}

func (s *CloudUpdateAgentRequest) GetObClidType() *int64 {
	return s.ObClidType
}

func (s *CloudUpdateAgentRequest) GetObRecord() *int64 {
	return s.ObRecord
}

func (s *CloudUpdateAgentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudUpdateAgentRequest) GetPermitObPreviewTime() *string {
	return s.PermitObPreviewTime
}

func (s *CloudUpdateAgentRequest) GetPower() *int64 {
	return s.Power
}

func (s *CloudUpdateAgentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudUpdateAgentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudUpdateAgentRequest) GetSkillIds() *string {
	return s.SkillIds
}

func (s *CloudUpdateAgentRequest) GetSkillLevels() *string {
	return s.SkillLevels
}

func (s *CloudUpdateAgentRequest) GetWebrtcUrlType() *int64 {
	return s.WebrtcUrlType
}

func (s *CloudUpdateAgentRequest) GetWrapup() *int64 {
	return s.Wrapup
}

func (s *CloudUpdateAgentRequest) SetActive(v int64) *CloudUpdateAgentRequest {
	s.Active = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetAgentType(v int64) *CloudUpdateAgentRequest {
	s.AgentType = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetAreaCode(v string) *CloudUpdateAgentRequest {
	s.AreaCode = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetCallPower(v int64) *CloudUpdateAgentRequest {
	s.CallPower = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetCno(v string) *CloudUpdateAgentRequest {
	s.Cno = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetComment(v string) *CloudUpdateAgentRequest {
	s.Comment = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetEnterpriseId(v int64) *CloudUpdateAgentRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetIbRecord(v int64) *CloudUpdateAgentRequest {
	s.IbRecord = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetIsAsr(v int64) *CloudUpdateAgentRequest {
	s.IsAsr = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetIsOb(v int64) *CloudUpdateAgentRequest {
	s.IsOb = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetIsObRemember(v string) *CloudUpdateAgentRequest {
	s.IsObRemember = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetIsQualityCheck(v int64) *CloudUpdateAgentRequest {
	s.IsQualityCheck = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetIsRandom(v string) *CloudUpdateAgentRequest {
	s.IsRandom = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetName(v string) *CloudUpdateAgentRequest {
	s.Name = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetObClid(v string) *CloudUpdateAgentRequest {
	s.ObClid = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetObClidProperty(v string) *CloudUpdateAgentRequest {
	s.ObClidProperty = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetObClidType(v int64) *CloudUpdateAgentRequest {
	s.ObClidType = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetObRecord(v int64) *CloudUpdateAgentRequest {
	s.ObRecord = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetOwnerId(v int64) *CloudUpdateAgentRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetPermitObPreviewTime(v string) *CloudUpdateAgentRequest {
	s.PermitObPreviewTime = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetPower(v int64) *CloudUpdateAgentRequest {
	s.Power = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetResourceOwnerAccount(v string) *CloudUpdateAgentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetResourceOwnerId(v int64) *CloudUpdateAgentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetSkillIds(v string) *CloudUpdateAgentRequest {
	s.SkillIds = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetSkillLevels(v string) *CloudUpdateAgentRequest {
	s.SkillLevels = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetWebrtcUrlType(v int64) *CloudUpdateAgentRequest {
	s.WebrtcUrlType = &v
	return s
}

func (s *CloudUpdateAgentRequest) SetWrapup(v int64) *CloudUpdateAgentRequest {
	s.Wrapup = &v
	return s
}

func (s *CloudUpdateAgentRequest) Validate() error {
	return dara.Validate(s)
}
