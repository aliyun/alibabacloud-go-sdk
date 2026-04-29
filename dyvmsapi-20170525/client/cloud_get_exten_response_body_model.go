// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetExtenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudGetExtenResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudGetExtenResponseBody
	GetCode() *string
	SetData(v *CloudGetExtenResponseBodyData) *CloudGetExtenResponseBody
	GetData() *CloudGetExtenResponseBodyData
	SetMessage(v string) *CloudGetExtenResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudGetExtenResponseBody
	GetRequestId() *string
}

type CloudGetExtenResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudGetExtenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudGetExtenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudGetExtenResponseBody) GoString() string {
	return s.String()
}

func (s *CloudGetExtenResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudGetExtenResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudGetExtenResponseBody) GetData() *CloudGetExtenResponseBodyData {
	return s.Data
}

func (s *CloudGetExtenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudGetExtenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudGetExtenResponseBody) SetAccessDeniedDetail(v string) *CloudGetExtenResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudGetExtenResponseBody) SetCode(v string) *CloudGetExtenResponseBody {
	s.Code = &v
	return s
}

func (s *CloudGetExtenResponseBody) SetData(v *CloudGetExtenResponseBodyData) *CloudGetExtenResponseBody {
	s.Data = v
	return s
}

func (s *CloudGetExtenResponseBody) SetMessage(v string) *CloudGetExtenResponseBody {
	s.Message = &v
	return s
}

func (s *CloudGetExtenResponseBody) SetRequestId(v string) *CloudGetExtenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudGetExtenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudGetExtenResponseBodyData struct {
	// 语音编码
	//
	// example:
	//
	// alaw,ulaw
	Allow *string `json:"Allow,omitempty" xml:"Allow,omitempty"`
	// 区号
	//
	// example:
	//
	// 010
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// 绑定的agent-gateway
	//
	// example:
	//
	// 22
	BindGatewayId *int64 `json:"BindGatewayId,omitempty" xml:"BindGatewayId,omitempty"`
	// 呼叫权限，0：无限制，1：国内长途，2：国内本市，3：内部呼叫，默认无限制
	//
	// example:
	//
	// 0
	CallPower *string `json:"CallPower,omitempty" xml:"CallPower,omitempty"`
	// 创建时间
	//
	// example:
	//
	// 2026-03-30 06:09:04
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 降噪开关 0:关闭 1:开启 默认关闭(该参数只有在类型为注册到webrtc才有效)
	//
	// example:
	//
	// 0
	Denoise *int64 `json:"Denoise,omitempty" xml:"Denoise,omitempty"`
	// 企业编号
	//
	// example:
	//
	// 7000002
	EnterpriseId *string `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 分机号，3-11位
	//
	// example:
	//
	// 9000
	Exten *string `json:"Exten,omitempty" xml:"Exten,omitempty"`
	// 呼入是否录音，0：不录用，1：录音，默认录音
	//
	// example:
	//
	// 1
	IbRecord *int64 `json:"IbRecord,omitempty" xml:"IbRecord,omitempty"`
	// 分机号id
	//
	// example:
	//
	// 336450
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 是否允许摘机外呼，0：不允许，1：可以，默认允许
	//
	// example:
	//
	// 1
	IsDirect *string `json:"IsDirect,omitempty" xml:"IsDirect,omitempty"`
	// 是否允许外呼，0：不允许，1：可以，默认允许
	//
	// example:
	//
	// 1
	IsOb *string `json:"IsOb,omitempty" xml:"IsOb,omitempty"`
	// 网络防抖
	//
	// example:
	//
	// 0
	JitterBuffer *int64 `json:"JitterBuffer,omitempty" xml:"JitterBuffer,omitempty"`
	// 外呼是否录音，0：不录音，1：录音，默认录音
	//
	// example:
	//
	// 1
	ObRecord *int64 `json:"ObRecord,omitempty" xml:"ObRecord,omitempty"`
	// 密码
	//
	// example:
	//
	// Aa248236
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// 类型，1. 注册到IAD分机 2.注册到webrtc 3.注册到远程话机
	//
	// example:
	//
	// 2
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CloudGetExtenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudGetExtenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudGetExtenResponseBodyData) GetAllow() *string {
	return s.Allow
}

func (s *CloudGetExtenResponseBodyData) GetAreaCode() *string {
	return s.AreaCode
}

func (s *CloudGetExtenResponseBodyData) GetBindGatewayId() *int64 {
	return s.BindGatewayId
}

func (s *CloudGetExtenResponseBodyData) GetCallPower() *string {
	return s.CallPower
}

func (s *CloudGetExtenResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudGetExtenResponseBodyData) GetDenoise() *int64 {
	return s.Denoise
}

func (s *CloudGetExtenResponseBodyData) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudGetExtenResponseBodyData) GetExten() *string {
	return s.Exten
}

func (s *CloudGetExtenResponseBodyData) GetIbRecord() *int64 {
	return s.IbRecord
}

func (s *CloudGetExtenResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *CloudGetExtenResponseBodyData) GetIsDirect() *string {
	return s.IsDirect
}

func (s *CloudGetExtenResponseBodyData) GetIsOb() *string {
	return s.IsOb
}

func (s *CloudGetExtenResponseBodyData) GetJitterBuffer() *int64 {
	return s.JitterBuffer
}

func (s *CloudGetExtenResponseBodyData) GetObRecord() *int64 {
	return s.ObRecord
}

func (s *CloudGetExtenResponseBodyData) GetPassword() *string {
	return s.Password
}

func (s *CloudGetExtenResponseBodyData) GetType() *int64 {
	return s.Type
}

func (s *CloudGetExtenResponseBodyData) SetAllow(v string) *CloudGetExtenResponseBodyData {
	s.Allow = &v
	return s
}

func (s *CloudGetExtenResponseBodyData) SetAreaCode(v string) *CloudGetExtenResponseBodyData {
	s.AreaCode = &v
	return s
}

func (s *CloudGetExtenResponseBodyData) SetBindGatewayId(v int64) *CloudGetExtenResponseBodyData {
	s.BindGatewayId = &v
	return s
}

func (s *CloudGetExtenResponseBodyData) SetCallPower(v string) *CloudGetExtenResponseBodyData {
	s.CallPower = &v
	return s
}

func (s *CloudGetExtenResponseBodyData) SetCreateTime(v string) *CloudGetExtenResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CloudGetExtenResponseBodyData) SetDenoise(v int64) *CloudGetExtenResponseBodyData {
	s.Denoise = &v
	return s
}

func (s *CloudGetExtenResponseBodyData) SetEnterpriseId(v string) *CloudGetExtenResponseBodyData {
	s.EnterpriseId = &v
	return s
}

func (s *CloudGetExtenResponseBodyData) SetExten(v string) *CloudGetExtenResponseBodyData {
	s.Exten = &v
	return s
}

func (s *CloudGetExtenResponseBodyData) SetIbRecord(v int64) *CloudGetExtenResponseBodyData {
	s.IbRecord = &v
	return s
}

func (s *CloudGetExtenResponseBodyData) SetId(v int64) *CloudGetExtenResponseBodyData {
	s.Id = &v
	return s
}

func (s *CloudGetExtenResponseBodyData) SetIsDirect(v string) *CloudGetExtenResponseBodyData {
	s.IsDirect = &v
	return s
}

func (s *CloudGetExtenResponseBodyData) SetIsOb(v string) *CloudGetExtenResponseBodyData {
	s.IsOb = &v
	return s
}

func (s *CloudGetExtenResponseBodyData) SetJitterBuffer(v int64) *CloudGetExtenResponseBodyData {
	s.JitterBuffer = &v
	return s
}

func (s *CloudGetExtenResponseBodyData) SetObRecord(v int64) *CloudGetExtenResponseBodyData {
	s.ObRecord = &v
	return s
}

func (s *CloudGetExtenResponseBodyData) SetPassword(v string) *CloudGetExtenResponseBodyData {
	s.Password = &v
	return s
}

func (s *CloudGetExtenResponseBodyData) SetType(v int64) *CloudGetExtenResponseBodyData {
	s.Type = &v
	return s
}

func (s *CloudGetExtenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
