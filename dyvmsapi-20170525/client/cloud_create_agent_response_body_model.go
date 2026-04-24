// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudCreateAgentResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudCreateAgentResponseBody
	GetCode() *string
	SetData(v *CloudCreateAgentResponseBodyData) *CloudCreateAgentResponseBody
	GetData() *CloudCreateAgentResponseBodyData
	SetMessage(v string) *CloudCreateAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudCreateAgentResponseBody
	GetRequestId() *string
}

type CloudCreateAgentResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudCreateAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudCreateAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CloudCreateAgentResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudCreateAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudCreateAgentResponseBody) GetData() *CloudCreateAgentResponseBodyData {
	return s.Data
}

func (s *CloudCreateAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudCreateAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudCreateAgentResponseBody) SetAccessDeniedDetail(v string) *CloudCreateAgentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudCreateAgentResponseBody) SetCode(v string) *CloudCreateAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CloudCreateAgentResponseBody) SetData(v *CloudCreateAgentResponseBodyData) *CloudCreateAgentResponseBody {
	s.Data = v
	return s
}

func (s *CloudCreateAgentResponseBody) SetMessage(v string) *CloudCreateAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CloudCreateAgentResponseBody) SetRequestId(v string) *CloudCreateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudCreateAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudCreateAgentResponseBodyData struct {
	// 座席配置信息
	Agent *CloudCreateAgentResponseBodyDataAgent `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Struct"`
	// 座席所需技能数组
	AgentSkills []*CloudCreateAgentResponseBodyDataAgentSkills `json:"AgentSkills,omitempty" xml:"AgentSkills,omitempty" type:"Repeated"`
}

func (s CloudCreateAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudCreateAgentResponseBodyData) GetAgent() *CloudCreateAgentResponseBodyDataAgent {
	return s.Agent
}

func (s *CloudCreateAgentResponseBodyData) GetAgentSkills() []*CloudCreateAgentResponseBodyDataAgentSkills {
	return s.AgentSkills
}

func (s *CloudCreateAgentResponseBodyData) SetAgent(v *CloudCreateAgentResponseBodyDataAgent) *CloudCreateAgentResponseBodyData {
	s.Agent = v
	return s
}

func (s *CloudCreateAgentResponseBodyData) SetAgentSkills(v []*CloudCreateAgentResponseBodyDataAgentSkills) *CloudCreateAgentResponseBodyData {
	s.AgentSkills = v
	return s
}

func (s *CloudCreateAgentResponseBodyData) Validate() error {
	if s.Agent != nil {
		if err := s.Agent.Validate(); err != nil {
			return err
		}
	}
	if s.AgentSkills != nil {
		for _, item := range s.AgentSkills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloudCreateAgentResponseBodyDataAgent struct {
	// 是否启用，0：停用，1：启用，默认启用
	//
	// example:
	//
	// 1
	Active *string `json:"Active,omitempty" xml:"Active,omitempty"`
	// 区号格式
	//
	// example:
	//
	// 010
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// 座席绑定电话
	//
	// example:
	//
	// 2222333
	BindTel *string `json:"BindTel,omitempty" xml:"BindTel,omitempty"`
	// 电话类型，1:固话 2:手机 3:分机 4:软电话
	//
	// example:
	//
	// 1
	BindTelType *string `json:"BindTelType,omitempty" xml:"BindTelType,omitempty"`
	// 呼叫权限，0：无限制，1：国内长途，2：国内本市，3：内部呼叫，默认无限制
	//
	// example:
	//
	// 0
	CallPower *string `json:"CallPower,omitempty" xml:"CallPower,omitempty"`
	// 座席工号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 创建时间，格式: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-01-30 08:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 企业编号
	//
	// example:
	//
	// 7000002
	EnterpriseId *string `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 呼入是否录音，0：不录用，1：录音，默认录音
	//
	// example:
	//
	// 1
	IbRecord *string `json:"IbRecord,omitempty" xml:"IbRecord,omitempty"`
	// 座席id
	//
	// example:
	//
	// 355
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 是否开启ASR转写：0：不开启，1：开启，默认不开启
	//
	// example:
	//
	// 0
	IsAsr *string `json:"IsAsr,omitempty" xml:"IsAsr,omitempty"`
	// 是否允许外呼，0：不允许，1：可以，默认允许
	//
	// example:
	//
	// 1
	IsOb *string `json:"IsOb,omitempty" xml:"IsOb,omitempty"`
	// example:
	//
	// 1
	IsQualityCheck *string `json:"IsQualityCheck,omitempty" xml:"IsQualityCheck,omitempty"`
	// 座席姓名
	//
	// example:
	//
	// 示例值
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 外显号码
	//
	// example:
	//
	// xxxxxxxx
	ObClid *string `json:"ObClid,omitempty" xml:"ObClid,omitempty"`
	// 外显规则属性，JSON格式
	//
	// example:
	//
	// {"isMatchCapital":0,"areaCodeRule":1,"isRandom":1}
	ObClidProperty *string `json:"ObClidProperty,omitempty" xml:"ObClidProperty,omitempty"`
	// 外显规则 1：企业默认 2：随机 3：按区号 4：动态外显
	//
	// example:
	//
	// 1
	ObClidType *string `json:"ObClidType,omitempty" xml:"ObClidType,omitempty"`
	// 外呼是否录音，0：不录音，1：录音，默认录音
	//
	// example:
	//
	// 1
	ObRecord *string `json:"ObRecord,omitempty" xml:"ObRecord,omitempty"`
	// 1：班长席，0：普通座席，默认普通座席
	//
	// example:
	//
	// 0
	Power *string `json:"Power,omitempty" xml:"Power,omitempty"`
	// 座席状态，0:离线，1：在线
	//
	// example:
	//
	// 0
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// webrtc软电话返回地址，0：企业默认 1：公网域名 2：专线域名 3：公网IP 4：专线IP
	//
	// example:
	//
	// 0
	WebrtcUrlType *string `json:"WebrtcUrlType,omitempty" xml:"WebrtcUrlType,omitempty"`
	// 整理时间，秒数，默认10秒
	//
	// example:
	//
	// 10
	Wrapup *string `json:"Wrapup,omitempty" xml:"Wrapup,omitempty"`
}

func (s CloudCreateAgentResponseBodyDataAgent) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateAgentResponseBodyDataAgent) GoString() string {
	return s.String()
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetActive() *string {
	return s.Active
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetAreaCode() *string {
	return s.AreaCode
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetBindTel() *string {
	return s.BindTel
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetBindTelType() *string {
	return s.BindTelType
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetCallPower() *string {
	return s.CallPower
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetCno() *string {
	return s.Cno
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetIbRecord() *string {
	return s.IbRecord
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetId() *string {
	return s.Id
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetIsAsr() *string {
	return s.IsAsr
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetIsOb() *string {
	return s.IsOb
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetIsQualityCheck() *string {
	return s.IsQualityCheck
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetName() *string {
	return s.Name
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetObClid() *string {
	return s.ObClid
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetObClidProperty() *string {
	return s.ObClidProperty
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetObClidType() *string {
	return s.ObClidType
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetObRecord() *string {
	return s.ObRecord
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetPower() *string {
	return s.Power
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetStatus() *string {
	return s.Status
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetWebrtcUrlType() *string {
	return s.WebrtcUrlType
}

func (s *CloudCreateAgentResponseBodyDataAgent) GetWrapup() *string {
	return s.Wrapup
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetActive(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.Active = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetAreaCode(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.AreaCode = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetBindTel(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.BindTel = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetBindTelType(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.BindTelType = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetCallPower(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.CallPower = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetCno(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.Cno = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetCreateTime(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.CreateTime = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetEnterpriseId(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.EnterpriseId = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetIbRecord(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.IbRecord = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetId(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.Id = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetIsAsr(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.IsAsr = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetIsOb(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.IsOb = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetIsQualityCheck(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.IsQualityCheck = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetName(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.Name = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetObClid(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.ObClid = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetObClidProperty(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.ObClidProperty = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetObClidType(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.ObClidType = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetObRecord(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.ObRecord = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetPower(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.Power = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetStatus(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.Status = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetWebrtcUrlType(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.WebrtcUrlType = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) SetWrapup(v string) *CloudCreateAgentResponseBodyDataAgent {
	s.Wrapup = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgent) Validate() error {
	return dara.Validate(s)
}

type CloudCreateAgentResponseBodyDataAgentSkills struct {
	// 座席id
	//
	// example:
	//
	// 2333
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 创建时间，格式: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-01-30 08:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 企业编号
	//
	// example:
	//
	// 7000002
	EnterpriseId *string `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// queueSkill关系表中id
	//
	// example:
	//
	// 355
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// skill的id
	//
	// example:
	//
	// 233
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// 技能值
	//
	// example:
	//
	// 2
	SkillLevel *string `json:"SkillLevel,omitempty" xml:"SkillLevel,omitempty"`
}

func (s CloudCreateAgentResponseBodyDataAgentSkills) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateAgentResponseBodyDataAgentSkills) GoString() string {
	return s.String()
}

func (s *CloudCreateAgentResponseBodyDataAgentSkills) GetAgentId() *string {
	return s.AgentId
}

func (s *CloudCreateAgentResponseBodyDataAgentSkills) GetCno() *string {
	return s.Cno
}

func (s *CloudCreateAgentResponseBodyDataAgentSkills) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudCreateAgentResponseBodyDataAgentSkills) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudCreateAgentResponseBodyDataAgentSkills) GetId() *string {
	return s.Id
}

func (s *CloudCreateAgentResponseBodyDataAgentSkills) GetSkillId() *string {
	return s.SkillId
}

func (s *CloudCreateAgentResponseBodyDataAgentSkills) GetSkillLevel() *string {
	return s.SkillLevel
}

func (s *CloudCreateAgentResponseBodyDataAgentSkills) SetAgentId(v string) *CloudCreateAgentResponseBodyDataAgentSkills {
	s.AgentId = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgentSkills) SetCno(v string) *CloudCreateAgentResponseBodyDataAgentSkills {
	s.Cno = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgentSkills) SetCreateTime(v string) *CloudCreateAgentResponseBodyDataAgentSkills {
	s.CreateTime = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgentSkills) SetEnterpriseId(v string) *CloudCreateAgentResponseBodyDataAgentSkills {
	s.EnterpriseId = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgentSkills) SetId(v string) *CloudCreateAgentResponseBodyDataAgentSkills {
	s.Id = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgentSkills) SetSkillId(v string) *CloudCreateAgentResponseBodyDataAgentSkills {
	s.SkillId = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgentSkills) SetSkillLevel(v string) *CloudCreateAgentResponseBodyDataAgentSkills {
	s.SkillLevel = &v
	return s
}

func (s *CloudCreateAgentResponseBodyDataAgentSkills) Validate() error {
	return dara.Validate(s)
}
