// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudPreviewoutcallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupTels(v string) *CloudPreviewoutcallRequest
	GetBackupTels() *string
	SetCallVariables(v string) *CloudPreviewoutcallRequest
	GetCallVariables() *string
	SetCdrIsAsr(v int64) *CloudPreviewoutcallRequest
	GetCdrIsAsr() *int64
	SetClidList(v string) *CloudPreviewoutcallRequest
	GetClidList() *string
	SetCno(v string) *CloudPreviewoutcallRequest
	GetCno() *string
	SetCrnId(v string) *CloudPreviewoutcallRequest
	GetCrnId() *string
	SetDialTelTimeout(v int64) *CloudPreviewoutcallRequest
	GetDialTelTimeout() *int64
	SetEnterpriseId(v int64) *CloudPreviewoutcallRequest
	GetEnterpriseId() *int64
	SetIsInvestigation(v int64) *CloudPreviewoutcallRequest
	GetIsInvestigation() *int64
	SetObClid(v string) *CloudPreviewoutcallRequest
	GetObClid() *string
	SetObClidAreaCode(v string) *CloudPreviewoutcallRequest
	GetObClidAreaCode() *string
	SetObClidGroup(v string) *CloudPreviewoutcallRequest
	GetObClidGroup() *string
	SetRequestUniqueId(v string) *CloudPreviewoutcallRequest
	GetRequestUniqueId() *string
	SetTel(v string) *CloudPreviewoutcallRequest
	GetTel() *string
	SetTimeout(v int64) *CloudPreviewoutcallRequest
	GetTimeout() *int64
}

type CloudPreviewoutcallRequest struct {
	// 备用外呼号码；tel没呼通时呼叫备用号码，最多支持5个号码，多个之间用逗号分隔
	//
	// example:
	//
	// 41008502
	BackupTels *string `json:"BackupTels,omitempty" xml:"BackupTels,omitempty"`
	// 通道变量可以在以下场景中使用：1. 事件推送，作为推送字段使用2. 通话记录中，将字段存储在自定义字段3. 通过SIP-Header方式将字段传到呼叫的下游链路格式如: [{"name":"mainUniqueId","value":"cdr_main_unique_id","type":"2"},{"name":"callType","value":"cdr_call_type","type":"2"}]name：变量名称value：变量值type：变量类型，1.普通变量 2.PJSIP_HEADER变量（用于将变量通过SIP-Header方式传到客户侧，最多支持5个）
	//
	// example:
	//
	// [{"name":"mainUniqueId","value":"cdr_main_unique_id","type":"2"},{"name":"callType","value":"cdr_call_type","type":"2"}]
	CallVariables *string `json:"CallVariables,omitempty" xml:"CallVariables,omitempty"`
	// 此次通话录音是否进行ASR转写；0.不进行 1.进行 默认：1
	//
	// example:
	//
	// 1
	CdrIsAsr *int64 `json:"CdrIsAsr,omitempty" xml:"CdrIsAsr,omitempty"`
	// 外显号码集合, 格式如: [{"clid":"1708587xxxx","priority":1},{"clid":"1304412xxxx","priority":2},{"clid":"0107990xxxx","priority":3}] clid：外显号码（String类型）（必填）priority：优先级（Integer类型）（非必填） 特别说明： 一次呼叫最多允许传10个指定号码；若指定了号码优先级，要求所有传入的号码都必传
	//
	// example:
	//
	// [{"clid":"1708587xxxx","priority":1},{"clid":"1304412xxxx","priority":2},{"clid":"0107990xxxx","priority":3}]
	ClidList *string `json:"ClidList,omitempty" xml:"ClidList,omitempty"`
	// 座席工号；3-10位数字
	//
	// This parameter is required.
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 外显导航标识
	//
	// example:
	//
	// 12333
	CrnId *string `json:"CrnId,omitempty" xml:"CrnId,omitempty"`
	// 呼叫客户侧超时时间；取值范围 5<=dialTelTimeout<=60，默认45(单位:s)
	//
	// example:
	//
	// 26
	DialTelTimeout *int64 `json:"DialTelTimeout,omitempty" xml:"DialTelTimeout,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 是否满意度调查；0.不进行满意度调查 1.进行满意度调查，默认：取企业配置
	//
	// example:
	//
	// 0
	IsInvestigation *int64 `json:"IsInvestigation,omitempty" xml:"IsInvestigation,omitempty"`
	// 可传入企业中继号码或设置好的客户侧外显号码
	//
	// example:
	//
	// 41008502
	ObClid *string `json:"ObClid,omitempty" xml:"ObClid,omitempty"`
	// 指定外显区号；传入区号，在指定区号下外显号码；如果obClid和obClidAreaCode都传的情况下，按照obClid外显
	//
	// example:
	//
	// 010
	ObClidAreaCode *string `json:"ObClidAreaCode,omitempty" xml:"ObClidAreaCode,omitempty"`
	// 客户侧外显号码组；使用clidGroup需要账号支持按标识路由，使用此参数时clid参数无效
	//
	// example:
	//
	// 333
	ObClidGroup *string `json:"ObClidGroup,omitempty" xml:"ObClidGroup,omitempty"`
	// 请求唯一id；取值：如果没有传入值则系统会生产。如果是加密的号码，需要URLEncode后传入
	//
	// example:
	//
	// req1234567
	RequestUniqueId *string `json:"RequestUniqueId,omitempty" xml:"RequestUniqueId,omitempty"`
	// 需要进行呼叫的客户方电话号码，固话带区号，手机不加0。 当开启号码隐藏设置时，可从弹屏事件中获取customerNumberKey的值，进行外呼；[加密外呼](../字段定义/接口部分/加密外呼号码加密规则.md)；如果是加密的号码，需要URLEncode后传入；当外呼相关配置-支持分机号外呼开启后，此字段可支持传手机号-分机号格式
	//
	// This parameter is required.
	//
	// example:
	//
	// 1774821736
	Tel *string `json:"Tel,omitempty" xml:"Tel,omitempty"`
	// 呼叫座席侧超时时间；取值范围 5<=timeout<=60，默认30(单位:s)
	//
	// example:
	//
	// 16
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CloudPreviewoutcallRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudPreviewoutcallRequest) GoString() string {
	return s.String()
}

func (s *CloudPreviewoutcallRequest) GetBackupTels() *string {
	return s.BackupTels
}

func (s *CloudPreviewoutcallRequest) GetCallVariables() *string {
	return s.CallVariables
}

func (s *CloudPreviewoutcallRequest) GetCdrIsAsr() *int64 {
	return s.CdrIsAsr
}

func (s *CloudPreviewoutcallRequest) GetClidList() *string {
	return s.ClidList
}

func (s *CloudPreviewoutcallRequest) GetCno() *string {
	return s.Cno
}

func (s *CloudPreviewoutcallRequest) GetCrnId() *string {
	return s.CrnId
}

func (s *CloudPreviewoutcallRequest) GetDialTelTimeout() *int64 {
	return s.DialTelTimeout
}

func (s *CloudPreviewoutcallRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudPreviewoutcallRequest) GetIsInvestigation() *int64 {
	return s.IsInvestigation
}

func (s *CloudPreviewoutcallRequest) GetObClid() *string {
	return s.ObClid
}

func (s *CloudPreviewoutcallRequest) GetObClidAreaCode() *string {
	return s.ObClidAreaCode
}

func (s *CloudPreviewoutcallRequest) GetObClidGroup() *string {
	return s.ObClidGroup
}

func (s *CloudPreviewoutcallRequest) GetRequestUniqueId() *string {
	return s.RequestUniqueId
}

func (s *CloudPreviewoutcallRequest) GetTel() *string {
	return s.Tel
}

func (s *CloudPreviewoutcallRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *CloudPreviewoutcallRequest) SetBackupTels(v string) *CloudPreviewoutcallRequest {
	s.BackupTels = &v
	return s
}

func (s *CloudPreviewoutcallRequest) SetCallVariables(v string) *CloudPreviewoutcallRequest {
	s.CallVariables = &v
	return s
}

func (s *CloudPreviewoutcallRequest) SetCdrIsAsr(v int64) *CloudPreviewoutcallRequest {
	s.CdrIsAsr = &v
	return s
}

func (s *CloudPreviewoutcallRequest) SetClidList(v string) *CloudPreviewoutcallRequest {
	s.ClidList = &v
	return s
}

func (s *CloudPreviewoutcallRequest) SetCno(v string) *CloudPreviewoutcallRequest {
	s.Cno = &v
	return s
}

func (s *CloudPreviewoutcallRequest) SetCrnId(v string) *CloudPreviewoutcallRequest {
	s.CrnId = &v
	return s
}

func (s *CloudPreviewoutcallRequest) SetDialTelTimeout(v int64) *CloudPreviewoutcallRequest {
	s.DialTelTimeout = &v
	return s
}

func (s *CloudPreviewoutcallRequest) SetEnterpriseId(v int64) *CloudPreviewoutcallRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudPreviewoutcallRequest) SetIsInvestigation(v int64) *CloudPreviewoutcallRequest {
	s.IsInvestigation = &v
	return s
}

func (s *CloudPreviewoutcallRequest) SetObClid(v string) *CloudPreviewoutcallRequest {
	s.ObClid = &v
	return s
}

func (s *CloudPreviewoutcallRequest) SetObClidAreaCode(v string) *CloudPreviewoutcallRequest {
	s.ObClidAreaCode = &v
	return s
}

func (s *CloudPreviewoutcallRequest) SetObClidGroup(v string) *CloudPreviewoutcallRequest {
	s.ObClidGroup = &v
	return s
}

func (s *CloudPreviewoutcallRequest) SetRequestUniqueId(v string) *CloudPreviewoutcallRequest {
	s.RequestUniqueId = &v
	return s
}

func (s *CloudPreviewoutcallRequest) SetTel(v string) *CloudPreviewoutcallRequest {
	s.Tel = &v
	return s
}

func (s *CloudPreviewoutcallRequest) SetTimeout(v int64) *CloudPreviewoutcallRequest {
	s.Timeout = &v
	return s
}

func (s *CloudPreviewoutcallRequest) Validate() error {
	return dara.Validate(s)
}
