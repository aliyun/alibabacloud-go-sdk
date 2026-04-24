// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v int64) *CloudCreateAgentRequest
	GetActive() *int64
	SetAreaCode(v string) *CloudCreateAgentRequest
	GetAreaCode() *string
	SetCallPower(v int64) *CloudCreateAgentRequest
	GetCallPower() *int64
	SetCno(v string) *CloudCreateAgentRequest
	GetCno() *string
	SetComment(v string) *CloudCreateAgentRequest
	GetComment() *string
	SetEnterpriseId(v int64) *CloudCreateAgentRequest
	GetEnterpriseId() *int64
	SetIbRecord(v int64) *CloudCreateAgentRequest
	GetIbRecord() *int64
	SetIsAsr(v int64) *CloudCreateAgentRequest
	GetIsAsr() *int64
	SetIsOb(v int64) *CloudCreateAgentRequest
	GetIsOb() *int64
	SetIsObRemember(v string) *CloudCreateAgentRequest
	GetIsObRemember() *string
	SetIsQualityCheck(v int64) *CloudCreateAgentRequest
	GetIsQualityCheck() *int64
	SetName(v string) *CloudCreateAgentRequest
	GetName() *string
	SetObClid(v string) *CloudCreateAgentRequest
	GetObClid() *string
	SetObClidProperty(v string) *CloudCreateAgentRequest
	GetObClidProperty() *string
	SetObClidType(v int64) *CloudCreateAgentRequest
	GetObClidType() *int64
	SetObRecord(v int64) *CloudCreateAgentRequest
	GetObRecord() *int64
	SetOwnerId(v int64) *CloudCreateAgentRequest
	GetOwnerId() *int64
	SetPermitObPreviewTime(v string) *CloudCreateAgentRequest
	GetPermitObPreviewTime() *string
	SetPower(v int64) *CloudCreateAgentRequest
	GetPower() *int64
	SetResourceOwnerAccount(v string) *CloudCreateAgentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudCreateAgentRequest
	GetResourceOwnerId() *int64
	SetSkillIds(v string) *CloudCreateAgentRequest
	GetSkillIds() *string
	SetSkillLevels(v string) *CloudCreateAgentRequest
	GetSkillLevels() *string
	SetWebrtcUrlType(v int64) *CloudCreateAgentRequest
	GetWebrtcUrlType() *int64
	SetWrapup(v int64) *CloudCreateAgentRequest
	GetWrapup() *int64
}

type CloudCreateAgentRequest struct {
	// 是否激活；取值0或1，取值说明 0：不激活，1激活，默认激活
	//
	// example:
	//
	// 1
	Active *int64 `json:"Active,omitempty" xml:"Active,omitempty"`
	// 座席所属区号；区号格式
	//
	// This parameter is required.
	//
	// example:
	//
	// 010
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// 呼叫权限；取值说明 0：无限制，1：国内长途，2：国内本市，3：内部呼叫，默认无限制
	//
	// example:
	//
	// 0
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
	// 示例值示例值
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
	// 0
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
	// 座席姓名；需进行UTF-8格式的URLEncode编码
	//
	// This parameter is required.
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
	// 外显属性；取值：{"isMatchCapital":0,"areaCodeRule":1,"isRandom":1}；obClidType=4时，isMatchCapital表示是否匹配省会号码，1是 0否，areaCodeRule表示区号匹配规则，1座席区号 2客户号码区号 ;isRandom 随机外显
	//
	// example:
	//
	// {"isMatchCapital":0,"areaCodeRule":1,"isRandom":1}
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
	// 0
	Power                *int64  `json:"Power,omitempty" xml:"Power,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 所选的技能id；可选择多个，多个之间使用英文标点逗号间隔.例如:"1,2,3"；传入skillIds时，需要同时传入skillLevels
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
	// 0
	WebrtcUrlType *int64 `json:"WebrtcUrlType,omitempty" xml:"WebrtcUrlType,omitempty"`
	// 整理时间；单位秒数，默认10秒
	//
	// example:
	//
	// 20
	Wrapup *int64 `json:"Wrapup,omitempty" xml:"Wrapup,omitempty"`
}

func (s CloudCreateAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateAgentRequest) GoString() string {
	return s.String()
}

func (s *CloudCreateAgentRequest) GetActive() *int64 {
	return s.Active
}

func (s *CloudCreateAgentRequest) GetAreaCode() *string {
	return s.AreaCode
}

func (s *CloudCreateAgentRequest) GetCallPower() *int64 {
	return s.CallPower
}

func (s *CloudCreateAgentRequest) GetCno() *string {
	return s.Cno
}

func (s *CloudCreateAgentRequest) GetComment() *string {
	return s.Comment
}

func (s *CloudCreateAgentRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudCreateAgentRequest) GetIbRecord() *int64 {
	return s.IbRecord
}

func (s *CloudCreateAgentRequest) GetIsAsr() *int64 {
	return s.IsAsr
}

func (s *CloudCreateAgentRequest) GetIsOb() *int64 {
	return s.IsOb
}

func (s *CloudCreateAgentRequest) GetIsObRemember() *string {
	return s.IsObRemember
}

func (s *CloudCreateAgentRequest) GetIsQualityCheck() *int64 {
	return s.IsQualityCheck
}

func (s *CloudCreateAgentRequest) GetName() *string {
	return s.Name
}

func (s *CloudCreateAgentRequest) GetObClid() *string {
	return s.ObClid
}

func (s *CloudCreateAgentRequest) GetObClidProperty() *string {
	return s.ObClidProperty
}

func (s *CloudCreateAgentRequest) GetObClidType() *int64 {
	return s.ObClidType
}

func (s *CloudCreateAgentRequest) GetObRecord() *int64 {
	return s.ObRecord
}

func (s *CloudCreateAgentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudCreateAgentRequest) GetPermitObPreviewTime() *string {
	return s.PermitObPreviewTime
}

func (s *CloudCreateAgentRequest) GetPower() *int64 {
	return s.Power
}

func (s *CloudCreateAgentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudCreateAgentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudCreateAgentRequest) GetSkillIds() *string {
	return s.SkillIds
}

func (s *CloudCreateAgentRequest) GetSkillLevels() *string {
	return s.SkillLevels
}

func (s *CloudCreateAgentRequest) GetWebrtcUrlType() *int64 {
	return s.WebrtcUrlType
}

func (s *CloudCreateAgentRequest) GetWrapup() *int64 {
	return s.Wrapup
}

func (s *CloudCreateAgentRequest) SetActive(v int64) *CloudCreateAgentRequest {
	s.Active = &v
	return s
}

func (s *CloudCreateAgentRequest) SetAreaCode(v string) *CloudCreateAgentRequest {
	s.AreaCode = &v
	return s
}

func (s *CloudCreateAgentRequest) SetCallPower(v int64) *CloudCreateAgentRequest {
	s.CallPower = &v
	return s
}

func (s *CloudCreateAgentRequest) SetCno(v string) *CloudCreateAgentRequest {
	s.Cno = &v
	return s
}

func (s *CloudCreateAgentRequest) SetComment(v string) *CloudCreateAgentRequest {
	s.Comment = &v
	return s
}

func (s *CloudCreateAgentRequest) SetEnterpriseId(v int64) *CloudCreateAgentRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudCreateAgentRequest) SetIbRecord(v int64) *CloudCreateAgentRequest {
	s.IbRecord = &v
	return s
}

func (s *CloudCreateAgentRequest) SetIsAsr(v int64) *CloudCreateAgentRequest {
	s.IsAsr = &v
	return s
}

func (s *CloudCreateAgentRequest) SetIsOb(v int64) *CloudCreateAgentRequest {
	s.IsOb = &v
	return s
}

func (s *CloudCreateAgentRequest) SetIsObRemember(v string) *CloudCreateAgentRequest {
	s.IsObRemember = &v
	return s
}

func (s *CloudCreateAgentRequest) SetIsQualityCheck(v int64) *CloudCreateAgentRequest {
	s.IsQualityCheck = &v
	return s
}

func (s *CloudCreateAgentRequest) SetName(v string) *CloudCreateAgentRequest {
	s.Name = &v
	return s
}

func (s *CloudCreateAgentRequest) SetObClid(v string) *CloudCreateAgentRequest {
	s.ObClid = &v
	return s
}

func (s *CloudCreateAgentRequest) SetObClidProperty(v string) *CloudCreateAgentRequest {
	s.ObClidProperty = &v
	return s
}

func (s *CloudCreateAgentRequest) SetObClidType(v int64) *CloudCreateAgentRequest {
	s.ObClidType = &v
	return s
}

func (s *CloudCreateAgentRequest) SetObRecord(v int64) *CloudCreateAgentRequest {
	s.ObRecord = &v
	return s
}

func (s *CloudCreateAgentRequest) SetOwnerId(v int64) *CloudCreateAgentRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudCreateAgentRequest) SetPermitObPreviewTime(v string) *CloudCreateAgentRequest {
	s.PermitObPreviewTime = &v
	return s
}

func (s *CloudCreateAgentRequest) SetPower(v int64) *CloudCreateAgentRequest {
	s.Power = &v
	return s
}

func (s *CloudCreateAgentRequest) SetResourceOwnerAccount(v string) *CloudCreateAgentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudCreateAgentRequest) SetResourceOwnerId(v int64) *CloudCreateAgentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudCreateAgentRequest) SetSkillIds(v string) *CloudCreateAgentRequest {
	s.SkillIds = &v
	return s
}

func (s *CloudCreateAgentRequest) SetSkillLevels(v string) *CloudCreateAgentRequest {
	s.SkillLevels = &v
	return s
}

func (s *CloudCreateAgentRequest) SetWebrtcUrlType(v int64) *CloudCreateAgentRequest {
	s.WebrtcUrlType = &v
	return s
}

func (s *CloudCreateAgentRequest) SetWrapup(v int64) *CloudCreateAgentRequest {
	s.Wrapup = &v
	return s
}

func (s *CloudCreateAgentRequest) Validate() error {
	return dara.Validate(s)
}
