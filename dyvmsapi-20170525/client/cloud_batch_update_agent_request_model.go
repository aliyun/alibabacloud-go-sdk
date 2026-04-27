// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudBatchUpdateAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v int64) *CloudBatchUpdateAgentRequest
	GetActive() *int64
	SetAgentType(v int64) *CloudBatchUpdateAgentRequest
	GetAgentType() *int64
	SetAreaCode(v string) *CloudBatchUpdateAgentRequest
	GetAreaCode() *string
	SetCallPower(v int64) *CloudBatchUpdateAgentRequest
	GetCallPower() *int64
	SetCnos(v string) *CloudBatchUpdateAgentRequest
	GetCnos() *string
	SetComment(v string) *CloudBatchUpdateAgentRequest
	GetComment() *string
	SetEnterpriseId(v int64) *CloudBatchUpdateAgentRequest
	GetEnterpriseId() *int64
	SetIbRecord(v int64) *CloudBatchUpdateAgentRequest
	GetIbRecord() *int64
	SetIsAsr(v int64) *CloudBatchUpdateAgentRequest
	GetIsAsr() *int64
	SetIsOb(v int64) *CloudBatchUpdateAgentRequest
	GetIsOb() *int64
	SetIsObRemember(v string) *CloudBatchUpdateAgentRequest
	GetIsObRemember() *string
	SetIsQualityCheck(v int64) *CloudBatchUpdateAgentRequest
	GetIsQualityCheck() *int64
	SetIsRandom(v string) *CloudBatchUpdateAgentRequest
	GetIsRandom() *string
	SetName(v string) *CloudBatchUpdateAgentRequest
	GetName() *string
	SetObClid(v string) *CloudBatchUpdateAgentRequest
	GetObClid() *string
	SetObClidProperty(v string) *CloudBatchUpdateAgentRequest
	GetObClidProperty() *string
	SetObClidType(v int64) *CloudBatchUpdateAgentRequest
	GetObClidType() *int64
	SetObRecord(v int64) *CloudBatchUpdateAgentRequest
	GetObRecord() *int64
	SetOwnerId(v int64) *CloudBatchUpdateAgentRequest
	GetOwnerId() *int64
	SetPermitObPreviewTime(v string) *CloudBatchUpdateAgentRequest
	GetPermitObPreviewTime() *string
	SetPower(v int64) *CloudBatchUpdateAgentRequest
	GetPower() *int64
	SetResourceOwnerAccount(v string) *CloudBatchUpdateAgentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudBatchUpdateAgentRequest
	GetResourceOwnerId() *int64
	SetWebrtcUrlType(v int64) *CloudBatchUpdateAgentRequest
	GetWebrtcUrlType() *int64
	SetWrapup(v int64) *CloudBatchUpdateAgentRequest
	GetWrapup() *int64
}

type CloudBatchUpdateAgentRequest struct {
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
	// 0
	CallPower *int64 `json:"CallPower,omitempty" xml:"CallPower,omitempty"`
	// 座席工号列表；需要批量更新座席的工号列表，多个之间使用英文标点逗号间隔，最多支持100个座席的批量更新。例如："0001,0002,00103"；座席工号，正整数，取值3-10位数字
	//
	// This parameter is required.
	//
	// example:
	//
	// 0001,0002,00103
	Cnos *string `json:"Cnos,omitempty" xml:"Cnos,omitempty"`
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
	// 0
	IsRandom *string `json:"IsRandom,omitempty" xml:"IsRandom,omitempty"`
	// 座席姓名；需进行UTF-8格式的URLEncode编码
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
	// 10
	Wrapup *int64 `json:"Wrapup,omitempty" xml:"Wrapup,omitempty"`
}

func (s CloudBatchUpdateAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudBatchUpdateAgentRequest) GoString() string {
	return s.String()
}

func (s *CloudBatchUpdateAgentRequest) GetActive() *int64 {
	return s.Active
}

func (s *CloudBatchUpdateAgentRequest) GetAgentType() *int64 {
	return s.AgentType
}

func (s *CloudBatchUpdateAgentRequest) GetAreaCode() *string {
	return s.AreaCode
}

func (s *CloudBatchUpdateAgentRequest) GetCallPower() *int64 {
	return s.CallPower
}

func (s *CloudBatchUpdateAgentRequest) GetCnos() *string {
	return s.Cnos
}

func (s *CloudBatchUpdateAgentRequest) GetComment() *string {
	return s.Comment
}

func (s *CloudBatchUpdateAgentRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudBatchUpdateAgentRequest) GetIbRecord() *int64 {
	return s.IbRecord
}

func (s *CloudBatchUpdateAgentRequest) GetIsAsr() *int64 {
	return s.IsAsr
}

func (s *CloudBatchUpdateAgentRequest) GetIsOb() *int64 {
	return s.IsOb
}

func (s *CloudBatchUpdateAgentRequest) GetIsObRemember() *string {
	return s.IsObRemember
}

func (s *CloudBatchUpdateAgentRequest) GetIsQualityCheck() *int64 {
	return s.IsQualityCheck
}

func (s *CloudBatchUpdateAgentRequest) GetIsRandom() *string {
	return s.IsRandom
}

func (s *CloudBatchUpdateAgentRequest) GetName() *string {
	return s.Name
}

func (s *CloudBatchUpdateAgentRequest) GetObClid() *string {
	return s.ObClid
}

func (s *CloudBatchUpdateAgentRequest) GetObClidProperty() *string {
	return s.ObClidProperty
}

func (s *CloudBatchUpdateAgentRequest) GetObClidType() *int64 {
	return s.ObClidType
}

func (s *CloudBatchUpdateAgentRequest) GetObRecord() *int64 {
	return s.ObRecord
}

func (s *CloudBatchUpdateAgentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudBatchUpdateAgentRequest) GetPermitObPreviewTime() *string {
	return s.PermitObPreviewTime
}

func (s *CloudBatchUpdateAgentRequest) GetPower() *int64 {
	return s.Power
}

func (s *CloudBatchUpdateAgentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudBatchUpdateAgentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudBatchUpdateAgentRequest) GetWebrtcUrlType() *int64 {
	return s.WebrtcUrlType
}

func (s *CloudBatchUpdateAgentRequest) GetWrapup() *int64 {
	return s.Wrapup
}

func (s *CloudBatchUpdateAgentRequest) SetActive(v int64) *CloudBatchUpdateAgentRequest {
	s.Active = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetAgentType(v int64) *CloudBatchUpdateAgentRequest {
	s.AgentType = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetAreaCode(v string) *CloudBatchUpdateAgentRequest {
	s.AreaCode = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetCallPower(v int64) *CloudBatchUpdateAgentRequest {
	s.CallPower = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetCnos(v string) *CloudBatchUpdateAgentRequest {
	s.Cnos = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetComment(v string) *CloudBatchUpdateAgentRequest {
	s.Comment = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetEnterpriseId(v int64) *CloudBatchUpdateAgentRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetIbRecord(v int64) *CloudBatchUpdateAgentRequest {
	s.IbRecord = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetIsAsr(v int64) *CloudBatchUpdateAgentRequest {
	s.IsAsr = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetIsOb(v int64) *CloudBatchUpdateAgentRequest {
	s.IsOb = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetIsObRemember(v string) *CloudBatchUpdateAgentRequest {
	s.IsObRemember = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetIsQualityCheck(v int64) *CloudBatchUpdateAgentRequest {
	s.IsQualityCheck = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetIsRandom(v string) *CloudBatchUpdateAgentRequest {
	s.IsRandom = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetName(v string) *CloudBatchUpdateAgentRequest {
	s.Name = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetObClid(v string) *CloudBatchUpdateAgentRequest {
	s.ObClid = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetObClidProperty(v string) *CloudBatchUpdateAgentRequest {
	s.ObClidProperty = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetObClidType(v int64) *CloudBatchUpdateAgentRequest {
	s.ObClidType = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetObRecord(v int64) *CloudBatchUpdateAgentRequest {
	s.ObRecord = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetOwnerId(v int64) *CloudBatchUpdateAgentRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetPermitObPreviewTime(v string) *CloudBatchUpdateAgentRequest {
	s.PermitObPreviewTime = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetPower(v int64) *CloudBatchUpdateAgentRequest {
	s.Power = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetResourceOwnerAccount(v string) *CloudBatchUpdateAgentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetResourceOwnerId(v int64) *CloudBatchUpdateAgentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetWebrtcUrlType(v int64) *CloudBatchUpdateAgentRequest {
	s.WebrtcUrlType = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) SetWrapup(v int64) *CloudBatchUpdateAgentRequest {
	s.Wrapup = &v
	return s
}

func (s *CloudBatchUpdateAgentRequest) Validate() error {
	return dara.Validate(s)
}
