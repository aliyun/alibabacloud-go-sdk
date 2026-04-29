// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateExtenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudCreateExtenResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudCreateExtenResponseBody
	GetCode() *string
	SetData(v *CloudCreateExtenResponseBodyData) *CloudCreateExtenResponseBody
	GetData() *CloudCreateExtenResponseBodyData
	SetMessage(v string) *CloudCreateExtenResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudCreateExtenResponseBody
	GetRequestId() *string
}

type CloudCreateExtenResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudCreateExtenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudCreateExtenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateExtenResponseBody) GoString() string {
	return s.String()
}

func (s *CloudCreateExtenResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudCreateExtenResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudCreateExtenResponseBody) GetData() *CloudCreateExtenResponseBodyData {
	return s.Data
}

func (s *CloudCreateExtenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudCreateExtenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudCreateExtenResponseBody) SetAccessDeniedDetail(v string) *CloudCreateExtenResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudCreateExtenResponseBody) SetCode(v string) *CloudCreateExtenResponseBody {
	s.Code = &v
	return s
}

func (s *CloudCreateExtenResponseBody) SetData(v *CloudCreateExtenResponseBodyData) *CloudCreateExtenResponseBody {
	s.Data = v
	return s
}

func (s *CloudCreateExtenResponseBody) SetMessage(v string) *CloudCreateExtenResponseBody {
	s.Message = &v
	return s
}

func (s *CloudCreateExtenResponseBody) SetRequestId(v string) *CloudCreateExtenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudCreateExtenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudCreateExtenResponseBodyData struct {
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
	// 28
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
	// 分机号,3-11位
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
	IsDirect *int64 `json:"IsDirect,omitempty" xml:"IsDirect,omitempty"`
	// 是否允许外呼，0：不允许，1：可以，默认允许
	//
	// example:
	//
	// 1
	IsOb *string `json:"IsOb,omitempty" xml:"IsOb,omitempty"`
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

func (s CloudCreateExtenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateExtenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudCreateExtenResponseBodyData) GetAllow() *string {
	return s.Allow
}

func (s *CloudCreateExtenResponseBodyData) GetAreaCode() *string {
	return s.AreaCode
}

func (s *CloudCreateExtenResponseBodyData) GetBindGatewayId() *int64 {
	return s.BindGatewayId
}

func (s *CloudCreateExtenResponseBodyData) GetCallPower() *string {
	return s.CallPower
}

func (s *CloudCreateExtenResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudCreateExtenResponseBodyData) GetDenoise() *int64 {
	return s.Denoise
}

func (s *CloudCreateExtenResponseBodyData) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudCreateExtenResponseBodyData) GetExten() *string {
	return s.Exten
}

func (s *CloudCreateExtenResponseBodyData) GetIbRecord() *int64 {
	return s.IbRecord
}

func (s *CloudCreateExtenResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *CloudCreateExtenResponseBodyData) GetIsDirect() *int64 {
	return s.IsDirect
}

func (s *CloudCreateExtenResponseBodyData) GetIsOb() *string {
	return s.IsOb
}

func (s *CloudCreateExtenResponseBodyData) GetObRecord() *int64 {
	return s.ObRecord
}

func (s *CloudCreateExtenResponseBodyData) GetPassword() *string {
	return s.Password
}

func (s *CloudCreateExtenResponseBodyData) GetType() *int64 {
	return s.Type
}

func (s *CloudCreateExtenResponseBodyData) SetAllow(v string) *CloudCreateExtenResponseBodyData {
	s.Allow = &v
	return s
}

func (s *CloudCreateExtenResponseBodyData) SetAreaCode(v string) *CloudCreateExtenResponseBodyData {
	s.AreaCode = &v
	return s
}

func (s *CloudCreateExtenResponseBodyData) SetBindGatewayId(v int64) *CloudCreateExtenResponseBodyData {
	s.BindGatewayId = &v
	return s
}

func (s *CloudCreateExtenResponseBodyData) SetCallPower(v string) *CloudCreateExtenResponseBodyData {
	s.CallPower = &v
	return s
}

func (s *CloudCreateExtenResponseBodyData) SetCreateTime(v string) *CloudCreateExtenResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CloudCreateExtenResponseBodyData) SetDenoise(v int64) *CloudCreateExtenResponseBodyData {
	s.Denoise = &v
	return s
}

func (s *CloudCreateExtenResponseBodyData) SetEnterpriseId(v string) *CloudCreateExtenResponseBodyData {
	s.EnterpriseId = &v
	return s
}

func (s *CloudCreateExtenResponseBodyData) SetExten(v string) *CloudCreateExtenResponseBodyData {
	s.Exten = &v
	return s
}

func (s *CloudCreateExtenResponseBodyData) SetIbRecord(v int64) *CloudCreateExtenResponseBodyData {
	s.IbRecord = &v
	return s
}

func (s *CloudCreateExtenResponseBodyData) SetId(v int64) *CloudCreateExtenResponseBodyData {
	s.Id = &v
	return s
}

func (s *CloudCreateExtenResponseBodyData) SetIsDirect(v int64) *CloudCreateExtenResponseBodyData {
	s.IsDirect = &v
	return s
}

func (s *CloudCreateExtenResponseBodyData) SetIsOb(v string) *CloudCreateExtenResponseBodyData {
	s.IsOb = &v
	return s
}

func (s *CloudCreateExtenResponseBodyData) SetObRecord(v int64) *CloudCreateExtenResponseBodyData {
	s.ObRecord = &v
	return s
}

func (s *CloudCreateExtenResponseBodyData) SetPassword(v string) *CloudCreateExtenResponseBodyData {
	s.Password = &v
	return s
}

func (s *CloudCreateExtenResponseBodyData) SetType(v int64) *CloudCreateExtenResponseBodyData {
	s.Type = &v
	return s
}

func (s *CloudCreateExtenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
