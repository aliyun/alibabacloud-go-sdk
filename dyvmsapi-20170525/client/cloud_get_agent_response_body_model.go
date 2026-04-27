// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudGetAgentResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudGetAgentResponseBody
	GetCode() *string
	SetData(v *CloudGetAgentResponseBodyData) *CloudGetAgentResponseBody
	GetData() *CloudGetAgentResponseBodyData
	SetMessage(v string) *CloudGetAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudGetAgentResponseBody
	GetRequestId() *string
}

type CloudGetAgentResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudGetAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 53D0F0Fe-cbbB-De28-6FCd-DdbBcefA46dD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudGetAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CloudGetAgentResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudGetAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudGetAgentResponseBody) GetData() *CloudGetAgentResponseBodyData {
	return s.Data
}

func (s *CloudGetAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudGetAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudGetAgentResponseBody) SetAccessDeniedDetail(v string) *CloudGetAgentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudGetAgentResponseBody) SetCode(v string) *CloudGetAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CloudGetAgentResponseBody) SetData(v *CloudGetAgentResponseBodyData) *CloudGetAgentResponseBody {
	s.Data = v
	return s
}

func (s *CloudGetAgentResponseBody) SetMessage(v string) *CloudGetAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CloudGetAgentResponseBody) SetRequestId(v string) *CloudGetAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudGetAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudGetAgentResponseBodyData struct {
	// 座席列表数组
	Agent []*CloudGetAgentResponseBodyDataAgent `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Repeated"`
}

func (s CloudGetAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudGetAgentResponseBodyData) GetAgent() []*CloudGetAgentResponseBodyDataAgent {
	return s.Agent
}

func (s *CloudGetAgentResponseBodyData) SetAgent(v []*CloudGetAgentResponseBodyDataAgent) *CloudGetAgentResponseBodyData {
	s.Agent = v
	return s
}

func (s *CloudGetAgentResponseBodyData) Validate() error {
	if s.Agent != nil {
		for _, item := range s.Agent {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloudGetAgentResponseBodyDataAgent struct {
	// 是否启用，0：停用，1：启用，默认启用
	//
	// example:
	//
	// 1
	Active *int64 `json:"Active,omitempty" xml:"Active,omitempty"`
	// 座席类型，1：电话座席，2：电脑座席，默认电话座席
	//
	// example:
	//
	// 1
	AgentType *int64 `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// 区号格式
	//
	// example:
	//
	// 010
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// 允许语音识别的通话类型，1:呼入 4：预览外呼 9：主叫外呼 5：预测外呼 2：webcall
	//
	// example:
	//
	// 2
	AsrCallType *string `json:"AsrCallType,omitempty" xml:"AsrCallType,omitempty"`
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
	BindTelType *int64 `json:"BindTelType,omitempty" xml:"BindTelType,omitempty"`
	// 呼叫权限，0：无限制，1：国内长途，2：国内本市，3：内部呼叫，默认无限制
	//
	// example:
	//
	// 0
	CallPower *int64 `json:"CallPower,omitempty" xml:"CallPower,omitempty"`
	// 座席工号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 备注
	//
	// example:
	//
	// 示例值示例值示例值
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
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
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 呼入是否录音，0：不录用，1：录音，默认录音
	//
	// example:
	//
	// 1
	IbRecord *int64 `json:"IbRecord,omitempty" xml:"IbRecord,omitempty"`
	// 座席id
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 是否开启ASR转写：0：不开启，1：开启，默认不开启
	//
	// example:
	//
	// 1
	IsAsr *int64 `json:"IsAsr,omitempty" xml:"IsAsr,omitempty"`
	// 是否允许axb外呼，默认1开启 0关闭
	//
	// example:
	//
	// 0
	IsAxbCall *string `json:"IsAxbCall,omitempty" xml:"IsAxbCall,omitempty"`
	// 是否允许外呼，0：不允许，1：可以，默认允许
	//
	// example:
	//
	// 1
	IsOb *int64 `json:"IsOb,omitempty" xml:"IsOb,omitempty"`
	// 外呼主叫记忆 1：开启 0：关闭
	//
	// example:
	//
	// 1
	IsObRemember *string `json:"IsObRemember,omitempty" xml:"IsObRemember,omitempty"`
	// 坐席最后一次登陆的时间
	//
	// example:
	//
	// 1774821736
	LoginTime *string `json:"LoginTime,omitempty" xml:"LoginTime,omitempty"`
	// 座席axb外呼是否使用mrcp方式推送语音流配置JSON Str格式数据
	//
	// example:
	//
	// {}
	MrcpProperty *string `json:"MrcpProperty,omitempty" xml:"MrcpProperty,omitempty"`
	// 座席姓名
	//
	// example:
	//
	// 示例值示例值
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
	// {"isMatchCapital":0,"areaCodeRule":1,"isRandom":1}
	ObClidProperty *string `json:"ObClidProperty,omitempty" xml:"ObClidProperty,omitempty"`
	// 外显规则 1：企业默认 2：随机 3：按区号 4：动态外显
	//
	// example:
	//
	// 1
	ObClidType *int64 `json:"ObClidType,omitempty" xml:"ObClidType,omitempty"`
	// 外呼是否录音，0：不录音，1：录音，默认录音
	//
	// example:
	//
	// 1
	ObRecord *int64 `json:"ObRecord,omitempty" xml:"ObRecord,omitempty"`
	// 可外呼时间段用,号分割
	//
	// example:
	//
	// 08:00,20:00
	PermitObPreviewTime *string `json:"PermitObPreviewTime,omitempty" xml:"PermitObPreviewTime,omitempty"`
	// 1：班长席，0：普通座席，默认普通座席
	//
	// example:
	//
	// 1
	Power *int64 `json:"Power,omitempty" xml:"Power,omitempty"`
	// 座席主动挂机配置Json Str 格式数据 unLink 是否允许主动挂断, 0:关, 1:开，prohibitDuration禁止时长
	//
	// example:
	//
	// 0
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// 坐席所属队列
	QueueList []*string `json:"QueueList,omitempty" xml:"QueueList,omitempty" type:"Repeated"`
	// 座席状态，0:离线，1：在线
	//
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 更新时间，格式: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-04-20 10:00:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	// 20
	Wrapup *int64 `json:"Wrapup,omitempty" xml:"Wrapup,omitempty"`
}

func (s CloudGetAgentResponseBodyDataAgent) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAgentResponseBodyDataAgent) GoString() string {
	return s.String()
}

func (s *CloudGetAgentResponseBodyDataAgent) GetActive() *int64 {
	return s.Active
}

func (s *CloudGetAgentResponseBodyDataAgent) GetAgentType() *int64 {
	return s.AgentType
}

func (s *CloudGetAgentResponseBodyDataAgent) GetAreaCode() *string {
	return s.AreaCode
}

func (s *CloudGetAgentResponseBodyDataAgent) GetAsrCallType() *string {
	return s.AsrCallType
}

func (s *CloudGetAgentResponseBodyDataAgent) GetBindTel() *string {
	return s.BindTel
}

func (s *CloudGetAgentResponseBodyDataAgent) GetBindTelType() *int64 {
	return s.BindTelType
}

func (s *CloudGetAgentResponseBodyDataAgent) GetCallPower() *int64 {
	return s.CallPower
}

func (s *CloudGetAgentResponseBodyDataAgent) GetCno() *string {
	return s.Cno
}

func (s *CloudGetAgentResponseBodyDataAgent) GetComment() *string {
	return s.Comment
}

func (s *CloudGetAgentResponseBodyDataAgent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudGetAgentResponseBodyDataAgent) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudGetAgentResponseBodyDataAgent) GetIbRecord() *int64 {
	return s.IbRecord
}

func (s *CloudGetAgentResponseBodyDataAgent) GetId() *int64 {
	return s.Id
}

func (s *CloudGetAgentResponseBodyDataAgent) GetIsAsr() *int64 {
	return s.IsAsr
}

func (s *CloudGetAgentResponseBodyDataAgent) GetIsAxbCall() *string {
	return s.IsAxbCall
}

func (s *CloudGetAgentResponseBodyDataAgent) GetIsOb() *int64 {
	return s.IsOb
}

func (s *CloudGetAgentResponseBodyDataAgent) GetIsObRemember() *string {
	return s.IsObRemember
}

func (s *CloudGetAgentResponseBodyDataAgent) GetLoginTime() *string {
	return s.LoginTime
}

func (s *CloudGetAgentResponseBodyDataAgent) GetMrcpProperty() *string {
	return s.MrcpProperty
}

func (s *CloudGetAgentResponseBodyDataAgent) GetName() *string {
	return s.Name
}

func (s *CloudGetAgentResponseBodyDataAgent) GetObClid() *string {
	return s.ObClid
}

func (s *CloudGetAgentResponseBodyDataAgent) GetObClidProperty() *string {
	return s.ObClidProperty
}

func (s *CloudGetAgentResponseBodyDataAgent) GetObClidType() *int64 {
	return s.ObClidType
}

func (s *CloudGetAgentResponseBodyDataAgent) GetObRecord() *int64 {
	return s.ObRecord
}

func (s *CloudGetAgentResponseBodyDataAgent) GetPermitObPreviewTime() *string {
	return s.PermitObPreviewTime
}

func (s *CloudGetAgentResponseBodyDataAgent) GetPower() *int64 {
	return s.Power
}

func (s *CloudGetAgentResponseBodyDataAgent) GetProperty() *string {
	return s.Property
}

func (s *CloudGetAgentResponseBodyDataAgent) GetQueueList() []*string {
	return s.QueueList
}

func (s *CloudGetAgentResponseBodyDataAgent) GetStatus() *int64 {
	return s.Status
}

func (s *CloudGetAgentResponseBodyDataAgent) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CloudGetAgentResponseBodyDataAgent) GetWebrtcUrlType() *string {
	return s.WebrtcUrlType
}

func (s *CloudGetAgentResponseBodyDataAgent) GetWrapup() *int64 {
	return s.Wrapup
}

func (s *CloudGetAgentResponseBodyDataAgent) SetActive(v int64) *CloudGetAgentResponseBodyDataAgent {
	s.Active = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetAgentType(v int64) *CloudGetAgentResponseBodyDataAgent {
	s.AgentType = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetAreaCode(v string) *CloudGetAgentResponseBodyDataAgent {
	s.AreaCode = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetAsrCallType(v string) *CloudGetAgentResponseBodyDataAgent {
	s.AsrCallType = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetBindTel(v string) *CloudGetAgentResponseBodyDataAgent {
	s.BindTel = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetBindTelType(v int64) *CloudGetAgentResponseBodyDataAgent {
	s.BindTelType = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetCallPower(v int64) *CloudGetAgentResponseBodyDataAgent {
	s.CallPower = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetCno(v string) *CloudGetAgentResponseBodyDataAgent {
	s.Cno = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetComment(v string) *CloudGetAgentResponseBodyDataAgent {
	s.Comment = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetCreateTime(v string) *CloudGetAgentResponseBodyDataAgent {
	s.CreateTime = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetEnterpriseId(v int64) *CloudGetAgentResponseBodyDataAgent {
	s.EnterpriseId = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetIbRecord(v int64) *CloudGetAgentResponseBodyDataAgent {
	s.IbRecord = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetId(v int64) *CloudGetAgentResponseBodyDataAgent {
	s.Id = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetIsAsr(v int64) *CloudGetAgentResponseBodyDataAgent {
	s.IsAsr = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetIsAxbCall(v string) *CloudGetAgentResponseBodyDataAgent {
	s.IsAxbCall = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetIsOb(v int64) *CloudGetAgentResponseBodyDataAgent {
	s.IsOb = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetIsObRemember(v string) *CloudGetAgentResponseBodyDataAgent {
	s.IsObRemember = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetLoginTime(v string) *CloudGetAgentResponseBodyDataAgent {
	s.LoginTime = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetMrcpProperty(v string) *CloudGetAgentResponseBodyDataAgent {
	s.MrcpProperty = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetName(v string) *CloudGetAgentResponseBodyDataAgent {
	s.Name = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetObClid(v string) *CloudGetAgentResponseBodyDataAgent {
	s.ObClid = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetObClidProperty(v string) *CloudGetAgentResponseBodyDataAgent {
	s.ObClidProperty = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetObClidType(v int64) *CloudGetAgentResponseBodyDataAgent {
	s.ObClidType = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetObRecord(v int64) *CloudGetAgentResponseBodyDataAgent {
	s.ObRecord = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetPermitObPreviewTime(v string) *CloudGetAgentResponseBodyDataAgent {
	s.PermitObPreviewTime = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetPower(v int64) *CloudGetAgentResponseBodyDataAgent {
	s.Power = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetProperty(v string) *CloudGetAgentResponseBodyDataAgent {
	s.Property = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetQueueList(v []*string) *CloudGetAgentResponseBodyDataAgent {
	s.QueueList = v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetStatus(v int64) *CloudGetAgentResponseBodyDataAgent {
	s.Status = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetUpdateTime(v string) *CloudGetAgentResponseBodyDataAgent {
	s.UpdateTime = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetWebrtcUrlType(v string) *CloudGetAgentResponseBodyDataAgent {
	s.WebrtcUrlType = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) SetWrapup(v int64) *CloudGetAgentResponseBodyDataAgent {
	s.Wrapup = &v
	return s
}

func (s *CloudGetAgentResponseBodyDataAgent) Validate() error {
	return dara.Validate(s)
}
