// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudBatchCreateAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v int64) *CloudBatchCreateAgentRequest
	GetActive() *int64
	SetAreaCode(v string) *CloudBatchCreateAgentRequest
	GetAreaCode() *string
	SetCallPower(v int64) *CloudBatchCreateAgentRequest
	GetCallPower() *int64
	SetCno(v string) *CloudBatchCreateAgentRequest
	GetCno() *string
	SetComment(v string) *CloudBatchCreateAgentRequest
	GetComment() *string
	SetEndCno(v string) *CloudBatchCreateAgentRequest
	GetEndCno() *string
	SetEnterpriseId(v int64) *CloudBatchCreateAgentRequest
	GetEnterpriseId() *int64
	SetIbRecord(v int64) *CloudBatchCreateAgentRequest
	GetIbRecord() *int64
	SetIsAsr(v int64) *CloudBatchCreateAgentRequest
	GetIsAsr() *int64
	SetIsOb(v int64) *CloudBatchCreateAgentRequest
	GetIsOb() *int64
	SetIsQualityCheck(v int64) *CloudBatchCreateAgentRequest
	GetIsQualityCheck() *int64
	SetName(v string) *CloudBatchCreateAgentRequest
	GetName() *string
	SetObClid(v string) *CloudBatchCreateAgentRequest
	GetObClid() *string
	SetObClidProperty(v string) *CloudBatchCreateAgentRequest
	GetObClidProperty() *string
	SetObClidType(v int64) *CloudBatchCreateAgentRequest
	GetObClidType() *int64
	SetObRecord(v int64) *CloudBatchCreateAgentRequest
	GetObRecord() *int64
	SetOwnerId(v int64) *CloudBatchCreateAgentRequest
	GetOwnerId() *int64
	SetPower(v int64) *CloudBatchCreateAgentRequest
	GetPower() *int64
	SetResourceOwnerAccount(v string) *CloudBatchCreateAgentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudBatchCreateAgentRequest
	GetResourceOwnerId() *int64
	SetSkillIds(v string) *CloudBatchCreateAgentRequest
	GetSkillIds() *string
	SetSkillLevels(v string) *CloudBatchCreateAgentRequest
	GetSkillLevels() *string
	SetWebrtcUrlType(v int64) *CloudBatchCreateAgentRequest
	GetWebrtcUrlType() *int64
	SetWrapup(v int64) *CloudBatchCreateAgentRequest
	GetWrapup() *int64
}

type CloudBatchCreateAgentRequest struct {
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
	// 1
	CallPower *int64 `json:"CallPower,omitempty" xml:"CallPower,omitempty"`
	// 开始座席工号；正整数，取值3-10位数字
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 备注
	//
	// example:
	//
	// 示例值
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// 结束座席工号；正整数，取值3-10位数字，单次批量创建不能超过100个
	//
	// This parameter is required.
	//
	// example:
	//
	// 103
	EndCno *string `json:"EndCno,omitempty" xml:"EndCno,omitempty"`
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
	// 示例值
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 透传号码；可传入企业中继号码或设置好的客户侧外显号码，当obClidType值为2或3时必选
	//
	// example:
	//
	// 22223333
	ObClid *string `json:"ObClid,omitempty" xml:"ObClid,omitempty"`
	// 外显属性；取值：{"isMatchCapital":0,"areaCodeRule":1}；obClidType=4时，isMatchCapital表示是否匹配省会号码，1是 0否，areaCodeRule表示区号匹配规则，1座席区号 2客户号码区号
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
	// 座席权限；取值1或0，取值说明 1：班长席，0：普通座席，默认为0 普通座席
	//
	// example:
	//
	// 1
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
	// 1
	WebrtcUrlType *int64 `json:"WebrtcUrlType,omitempty" xml:"WebrtcUrlType,omitempty"`
	// 整理时间；单位秒数，默认10秒
	//
	// example:
	//
	// 20
	Wrapup *int64 `json:"Wrapup,omitempty" xml:"Wrapup,omitempty"`
}

func (s CloudBatchCreateAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudBatchCreateAgentRequest) GoString() string {
	return s.String()
}

func (s *CloudBatchCreateAgentRequest) GetActive() *int64 {
	return s.Active
}

func (s *CloudBatchCreateAgentRequest) GetAreaCode() *string {
	return s.AreaCode
}

func (s *CloudBatchCreateAgentRequest) GetCallPower() *int64 {
	return s.CallPower
}

func (s *CloudBatchCreateAgentRequest) GetCno() *string {
	return s.Cno
}

func (s *CloudBatchCreateAgentRequest) GetComment() *string {
	return s.Comment
}

func (s *CloudBatchCreateAgentRequest) GetEndCno() *string {
	return s.EndCno
}

func (s *CloudBatchCreateAgentRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudBatchCreateAgentRequest) GetIbRecord() *int64 {
	return s.IbRecord
}

func (s *CloudBatchCreateAgentRequest) GetIsAsr() *int64 {
	return s.IsAsr
}

func (s *CloudBatchCreateAgentRequest) GetIsOb() *int64 {
	return s.IsOb
}

func (s *CloudBatchCreateAgentRequest) GetIsQualityCheck() *int64 {
	return s.IsQualityCheck
}

func (s *CloudBatchCreateAgentRequest) GetName() *string {
	return s.Name
}

func (s *CloudBatchCreateAgentRequest) GetObClid() *string {
	return s.ObClid
}

func (s *CloudBatchCreateAgentRequest) GetObClidProperty() *string {
	return s.ObClidProperty
}

func (s *CloudBatchCreateAgentRequest) GetObClidType() *int64 {
	return s.ObClidType
}

func (s *CloudBatchCreateAgentRequest) GetObRecord() *int64 {
	return s.ObRecord
}

func (s *CloudBatchCreateAgentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudBatchCreateAgentRequest) GetPower() *int64 {
	return s.Power
}

func (s *CloudBatchCreateAgentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudBatchCreateAgentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudBatchCreateAgentRequest) GetSkillIds() *string {
	return s.SkillIds
}

func (s *CloudBatchCreateAgentRequest) GetSkillLevels() *string {
	return s.SkillLevels
}

func (s *CloudBatchCreateAgentRequest) GetWebrtcUrlType() *int64 {
	return s.WebrtcUrlType
}

func (s *CloudBatchCreateAgentRequest) GetWrapup() *int64 {
	return s.Wrapup
}

func (s *CloudBatchCreateAgentRequest) SetActive(v int64) *CloudBatchCreateAgentRequest {
	s.Active = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetAreaCode(v string) *CloudBatchCreateAgentRequest {
	s.AreaCode = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetCallPower(v int64) *CloudBatchCreateAgentRequest {
	s.CallPower = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetCno(v string) *CloudBatchCreateAgentRequest {
	s.Cno = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetComment(v string) *CloudBatchCreateAgentRequest {
	s.Comment = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetEndCno(v string) *CloudBatchCreateAgentRequest {
	s.EndCno = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetEnterpriseId(v int64) *CloudBatchCreateAgentRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetIbRecord(v int64) *CloudBatchCreateAgentRequest {
	s.IbRecord = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetIsAsr(v int64) *CloudBatchCreateAgentRequest {
	s.IsAsr = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetIsOb(v int64) *CloudBatchCreateAgentRequest {
	s.IsOb = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetIsQualityCheck(v int64) *CloudBatchCreateAgentRequest {
	s.IsQualityCheck = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetName(v string) *CloudBatchCreateAgentRequest {
	s.Name = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetObClid(v string) *CloudBatchCreateAgentRequest {
	s.ObClid = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetObClidProperty(v string) *CloudBatchCreateAgentRequest {
	s.ObClidProperty = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetObClidType(v int64) *CloudBatchCreateAgentRequest {
	s.ObClidType = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetObRecord(v int64) *CloudBatchCreateAgentRequest {
	s.ObRecord = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetOwnerId(v int64) *CloudBatchCreateAgentRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetPower(v int64) *CloudBatchCreateAgentRequest {
	s.Power = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetResourceOwnerAccount(v string) *CloudBatchCreateAgentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetResourceOwnerId(v int64) *CloudBatchCreateAgentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetSkillIds(v string) *CloudBatchCreateAgentRequest {
	s.SkillIds = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetSkillLevels(v string) *CloudBatchCreateAgentRequest {
	s.SkillLevels = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetWebrtcUrlType(v int64) *CloudBatchCreateAgentRequest {
	s.WebrtcUrlType = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) SetWrapup(v int64) *CloudBatchCreateAgentRequest {
	s.Wrapup = &v
	return s
}

func (s *CloudBatchCreateAgentRequest) Validate() error {
	return dara.Validate(s)
}
