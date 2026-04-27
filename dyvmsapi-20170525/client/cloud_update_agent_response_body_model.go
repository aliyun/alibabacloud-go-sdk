// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudUpdateAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudUpdateAgentResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudUpdateAgentResponseBody
	GetCode() *string
	SetData(v *CloudUpdateAgentResponseBodyData) *CloudUpdateAgentResponseBody
	GetData() *CloudUpdateAgentResponseBodyData
	SetMessage(v string) *CloudUpdateAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudUpdateAgentResponseBody
	GetRequestId() *string
}

type CloudUpdateAgentResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudUpdateAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudUpdateAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudUpdateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CloudUpdateAgentResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudUpdateAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudUpdateAgentResponseBody) GetData() *CloudUpdateAgentResponseBodyData {
	return s.Data
}

func (s *CloudUpdateAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudUpdateAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudUpdateAgentResponseBody) SetAccessDeniedDetail(v string) *CloudUpdateAgentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudUpdateAgentResponseBody) SetCode(v string) *CloudUpdateAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CloudUpdateAgentResponseBody) SetData(v *CloudUpdateAgentResponseBodyData) *CloudUpdateAgentResponseBody {
	s.Data = v
	return s
}

func (s *CloudUpdateAgentResponseBody) SetMessage(v string) *CloudUpdateAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CloudUpdateAgentResponseBody) SetRequestId(v string) *CloudUpdateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudUpdateAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudUpdateAgentResponseBodyData struct {
	// 座席配置信息
	Agent *CloudUpdateAgentResponseBodyDataAgent `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Struct"`
	// 座席所需技能数组
	AgentSkills []*CloudUpdateAgentResponseBodyDataAgentSkills `json:"AgentSkills,omitempty" xml:"AgentSkills,omitempty" type:"Repeated"`
}

func (s CloudUpdateAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudUpdateAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudUpdateAgentResponseBodyData) GetAgent() *CloudUpdateAgentResponseBodyDataAgent {
	return s.Agent
}

func (s *CloudUpdateAgentResponseBodyData) GetAgentSkills() []*CloudUpdateAgentResponseBodyDataAgentSkills {
	return s.AgentSkills
}

func (s *CloudUpdateAgentResponseBodyData) SetAgent(v *CloudUpdateAgentResponseBodyDataAgent) *CloudUpdateAgentResponseBodyData {
	s.Agent = v
	return s
}

func (s *CloudUpdateAgentResponseBodyData) SetAgentSkills(v []*CloudUpdateAgentResponseBodyDataAgentSkills) *CloudUpdateAgentResponseBodyData {
	s.AgentSkills = v
	return s
}

func (s *CloudUpdateAgentResponseBodyData) Validate() error {
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

type CloudUpdateAgentResponseBodyDataAgent struct {
	// 是否启用，0：停用，1：启用，默认启用
	//
	// example:
	//
	// 1
	Active *string `json:"Active,omitempty" xml:"Active,omitempty"`
	// 座席类型，1：电话座席，2：电脑座席，默认电话座席
	//
	// example:
	//
	// 1
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
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
	// 22223333
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
	// 1
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
	// 2026-04-20 10:00:00
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
	// 100012
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 是否开启ASR转写：0：不开启，1：开启，默认不开启
	//
	// example:
	//
	// 1
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
	// 示例值示例值示例值
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 外显号码
	//
	// example:
	//
	// 22223333
	ObClid *string `json:"ObClid,omitempty" xml:"ObClid,omitempty"`
	// 外显规则属性，JSON格式
	//
	// example:
	//
	// {"isMatchCapital":0,"areaCodeRule":1}
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
	// 1
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
	// 1
	WebrtcUrlType *string `json:"WebrtcUrlType,omitempty" xml:"WebrtcUrlType,omitempty"`
	// 整理时间，秒数，默认10秒
	//
	// example:
	//
	// 22
	Wrapup *string `json:"Wrapup,omitempty" xml:"Wrapup,omitempty"`
}

func (s CloudUpdateAgentResponseBodyDataAgent) String() string {
	return dara.Prettify(s)
}

func (s CloudUpdateAgentResponseBodyDataAgent) GoString() string {
	return s.String()
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetActive() *string {
	return s.Active
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetAgentType() *string {
	return s.AgentType
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetAreaCode() *string {
	return s.AreaCode
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetBindTel() *string {
	return s.BindTel
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetBindTelType() *string {
	return s.BindTelType
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetCallPower() *string {
	return s.CallPower
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetCno() *string {
	return s.Cno
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetIbRecord() *string {
	return s.IbRecord
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetId() *string {
	return s.Id
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetIsAsr() *string {
	return s.IsAsr
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetIsOb() *string {
	return s.IsOb
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetIsQualityCheck() *string {
	return s.IsQualityCheck
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetName() *string {
	return s.Name
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetObClid() *string {
	return s.ObClid
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetObClidProperty() *string {
	return s.ObClidProperty
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetObClidType() *string {
	return s.ObClidType
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetObRecord() *string {
	return s.ObRecord
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetPower() *string {
	return s.Power
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetStatus() *string {
	return s.Status
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetWebrtcUrlType() *string {
	return s.WebrtcUrlType
}

func (s *CloudUpdateAgentResponseBodyDataAgent) GetWrapup() *string {
	return s.Wrapup
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetActive(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.Active = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetAgentType(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.AgentType = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetAreaCode(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.AreaCode = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetBindTel(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.BindTel = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetBindTelType(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.BindTelType = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetCallPower(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.CallPower = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetCno(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.Cno = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetCreateTime(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.CreateTime = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetEnterpriseId(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.EnterpriseId = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetIbRecord(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.IbRecord = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetId(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.Id = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetIsAsr(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.IsAsr = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetIsOb(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.IsOb = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetIsQualityCheck(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.IsQualityCheck = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetName(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.Name = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetObClid(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.ObClid = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetObClidProperty(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.ObClidProperty = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetObClidType(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.ObClidType = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetObRecord(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.ObRecord = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetPower(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.Power = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetStatus(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.Status = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetWebrtcUrlType(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.WebrtcUrlType = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) SetWrapup(v string) *CloudUpdateAgentResponseBodyDataAgent {
	s.Wrapup = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgent) Validate() error {
	return dara.Validate(s)
}

type CloudUpdateAgentResponseBodyDataAgentSkills struct {
	// 座席id
	//
	// example:
	//
	// 1234
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 创建时间，格式: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-04-20 10:00:00
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
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// skill的id
	//
	// example:
	//
	// 111
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// 技能值
	//
	// example:
	//
	// 2
	SkillLevel *string `json:"SkillLevel,omitempty" xml:"SkillLevel,omitempty"`
}

func (s CloudUpdateAgentResponseBodyDataAgentSkills) String() string {
	return dara.Prettify(s)
}

func (s CloudUpdateAgentResponseBodyDataAgentSkills) GoString() string {
	return s.String()
}

func (s *CloudUpdateAgentResponseBodyDataAgentSkills) GetAgentId() *string {
	return s.AgentId
}

func (s *CloudUpdateAgentResponseBodyDataAgentSkills) GetCno() *string {
	return s.Cno
}

func (s *CloudUpdateAgentResponseBodyDataAgentSkills) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudUpdateAgentResponseBodyDataAgentSkills) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudUpdateAgentResponseBodyDataAgentSkills) GetId() *string {
	return s.Id
}

func (s *CloudUpdateAgentResponseBodyDataAgentSkills) GetSkillId() *string {
	return s.SkillId
}

func (s *CloudUpdateAgentResponseBodyDataAgentSkills) GetSkillLevel() *string {
	return s.SkillLevel
}

func (s *CloudUpdateAgentResponseBodyDataAgentSkills) SetAgentId(v string) *CloudUpdateAgentResponseBodyDataAgentSkills {
	s.AgentId = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgentSkills) SetCno(v string) *CloudUpdateAgentResponseBodyDataAgentSkills {
	s.Cno = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgentSkills) SetCreateTime(v string) *CloudUpdateAgentResponseBodyDataAgentSkills {
	s.CreateTime = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgentSkills) SetEnterpriseId(v string) *CloudUpdateAgentResponseBodyDataAgentSkills {
	s.EnterpriseId = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgentSkills) SetId(v string) *CloudUpdateAgentResponseBodyDataAgentSkills {
	s.Id = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgentSkills) SetSkillId(v string) *CloudUpdateAgentResponseBodyDataAgentSkills {
	s.SkillId = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgentSkills) SetSkillLevel(v string) *CloudUpdateAgentResponseBodyDataAgentSkills {
	s.SkillLevel = &v
	return s
}

func (s *CloudUpdateAgentResponseBodyDataAgentSkills) Validate() error {
	return dara.Validate(s)
}
