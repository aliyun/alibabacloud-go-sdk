// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateExtenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllow(v string) *CloudCreateExtenRequest
	GetAllow() *string
	SetAreaCode(v string) *CloudCreateExtenRequest
	GetAreaCode() *string
	SetCallPower(v string) *CloudCreateExtenRequest
	GetCallPower() *string
	SetDenoise(v string) *CloudCreateExtenRequest
	GetDenoise() *string
	SetEnterpriseId(v int64) *CloudCreateExtenRequest
	GetEnterpriseId() *int64
	SetExten(v string) *CloudCreateExtenRequest
	GetExten() *string
	SetIadName(v string) *CloudCreateExtenRequest
	GetIadName() *string
	SetIbRecord(v int64) *CloudCreateExtenRequest
	GetIbRecord() *int64
	SetIsDirect(v int64) *CloudCreateExtenRequest
	GetIsDirect() *int64
	SetIsOb(v string) *CloudCreateExtenRequest
	GetIsOb() *string
	SetJitterBuffer(v int64) *CloudCreateExtenRequest
	GetJitterBuffer() *int64
	SetObRecord(v int64) *CloudCreateExtenRequest
	GetObRecord() *int64
	SetOwnerId(v int64) *CloudCreateExtenRequest
	GetOwnerId() *int64
	SetPassword(v string) *CloudCreateExtenRequest
	GetPassword() *string
	SetProperty(v string) *CloudCreateExtenRequest
	GetProperty() *string
	SetResourceOwnerAccount(v string) *CloudCreateExtenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudCreateExtenRequest
	GetResourceOwnerId() *int64
	SetType(v int64) *CloudCreateExtenRequest
	GetType() *int64
}

type CloudCreateExtenRequest struct {
	// 允许的语音编码,支持格式: 1. g729 2. g729,alaw,ulaw 3. alaw,ulaw,g729 4. alaw,ulaw 5. myopus,alaw,ulaw5. myopus,alaw,ulaw6. myopus,g729,alaw,ulaw 7. myopus,g729 公网+远程话机支持配置1/2/3;专线+远程话机支持配置1/2;公网+软电话支持配置4;专线+软电话支持配置5;
	//
	// example:
	//
	// 1
	Allow *string `json:"Allow,omitempty" xml:"Allow,omitempty"`
	// 区号
	//
	// This parameter is required.
	//
	// example:
	//
	// 010
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// 呼叫权限；0：无限制，1：国内长途，2：国内本市，3：内部呼叫，默认无限制
	//
	// example:
	//
	// 0
	CallPower *string `json:"CallPower,omitempty" xml:"CallPower,omitempty"`
	// 降噪开关；0:关闭 1:开启 (该参数只有在类型为注册到webrtc才有效)
	//
	// example:
	//
	// 0
	Denoise *string `json:"Denoise,omitempty" xml:"Denoise,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 分机号,3-11位
	//
	// This parameter is required.
	//
	// example:
	//
	// 9000
	Exten *string `json:"Exten,omitempty" xml:"Exten,omitempty"`
	// iad网关；该参数只有在类型为，注册到IAD分机，才有效
	//
	// example:
	//
	// ""
	IadName *string `json:"IadName,omitempty" xml:"IadName,omitempty"`
	// 呼入是否录音；0：不录用，1：录音，默认录音
	//
	// example:
	//
	// 1
	IbRecord *int64 `json:"IbRecord,omitempty" xml:"IbRecord,omitempty"`
	// 是否允许摘机外呼，0：不允许，1：可以，默认允许
	//
	// example:
	//
	// 1
	IsDirect *int64 `json:"IsDirect,omitempty" xml:"IsDirect,omitempty"`
	// 是否允许外呼；0：不允许，1：可以，默认允许
	//
	// example:
	//
	// 1
	IsOb *string `json:"IsOb,omitempty" xml:"IsOb,omitempty"`
	// 网络防抖；0：关闭，1：开启，默认关闭
	//
	// example:
	//
	// 0
	JitterBuffer *int64 `json:"JitterBuffer,omitempty" xml:"JitterBuffer,omitempty"`
	// 外呼是否录音；0：不录音，1：录音，默认录音
	//
	// example:
	//
	// 1
	ObRecord *int64 `json:"ObRecord,omitempty" xml:"ObRecord,omitempty"`
	OwnerId  *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 密码
	//
	// This parameter is required.
	//
	// example:
	//
	// Aa248236
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// 属性
	//
	// example:
	//
	// {}
	Property             *string `json:"Property,omitempty" xml:"Property,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 类型；1. 注册到IAD分机 2.注册到webrtc 3.注册到远程话机
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CloudCreateExtenRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateExtenRequest) GoString() string {
	return s.String()
}

func (s *CloudCreateExtenRequest) GetAllow() *string {
	return s.Allow
}

func (s *CloudCreateExtenRequest) GetAreaCode() *string {
	return s.AreaCode
}

func (s *CloudCreateExtenRequest) GetCallPower() *string {
	return s.CallPower
}

func (s *CloudCreateExtenRequest) GetDenoise() *string {
	return s.Denoise
}

func (s *CloudCreateExtenRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudCreateExtenRequest) GetExten() *string {
	return s.Exten
}

func (s *CloudCreateExtenRequest) GetIadName() *string {
	return s.IadName
}

func (s *CloudCreateExtenRequest) GetIbRecord() *int64 {
	return s.IbRecord
}

func (s *CloudCreateExtenRequest) GetIsDirect() *int64 {
	return s.IsDirect
}

func (s *CloudCreateExtenRequest) GetIsOb() *string {
	return s.IsOb
}

func (s *CloudCreateExtenRequest) GetJitterBuffer() *int64 {
	return s.JitterBuffer
}

func (s *CloudCreateExtenRequest) GetObRecord() *int64 {
	return s.ObRecord
}

func (s *CloudCreateExtenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudCreateExtenRequest) GetPassword() *string {
	return s.Password
}

func (s *CloudCreateExtenRequest) GetProperty() *string {
	return s.Property
}

func (s *CloudCreateExtenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudCreateExtenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudCreateExtenRequest) GetType() *int64 {
	return s.Type
}

func (s *CloudCreateExtenRequest) SetAllow(v string) *CloudCreateExtenRequest {
	s.Allow = &v
	return s
}

func (s *CloudCreateExtenRequest) SetAreaCode(v string) *CloudCreateExtenRequest {
	s.AreaCode = &v
	return s
}

func (s *CloudCreateExtenRequest) SetCallPower(v string) *CloudCreateExtenRequest {
	s.CallPower = &v
	return s
}

func (s *CloudCreateExtenRequest) SetDenoise(v string) *CloudCreateExtenRequest {
	s.Denoise = &v
	return s
}

func (s *CloudCreateExtenRequest) SetEnterpriseId(v int64) *CloudCreateExtenRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudCreateExtenRequest) SetExten(v string) *CloudCreateExtenRequest {
	s.Exten = &v
	return s
}

func (s *CloudCreateExtenRequest) SetIadName(v string) *CloudCreateExtenRequest {
	s.IadName = &v
	return s
}

func (s *CloudCreateExtenRequest) SetIbRecord(v int64) *CloudCreateExtenRequest {
	s.IbRecord = &v
	return s
}

func (s *CloudCreateExtenRequest) SetIsDirect(v int64) *CloudCreateExtenRequest {
	s.IsDirect = &v
	return s
}

func (s *CloudCreateExtenRequest) SetIsOb(v string) *CloudCreateExtenRequest {
	s.IsOb = &v
	return s
}

func (s *CloudCreateExtenRequest) SetJitterBuffer(v int64) *CloudCreateExtenRequest {
	s.JitterBuffer = &v
	return s
}

func (s *CloudCreateExtenRequest) SetObRecord(v int64) *CloudCreateExtenRequest {
	s.ObRecord = &v
	return s
}

func (s *CloudCreateExtenRequest) SetOwnerId(v int64) *CloudCreateExtenRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudCreateExtenRequest) SetPassword(v string) *CloudCreateExtenRequest {
	s.Password = &v
	return s
}

func (s *CloudCreateExtenRequest) SetProperty(v string) *CloudCreateExtenRequest {
	s.Property = &v
	return s
}

func (s *CloudCreateExtenRequest) SetResourceOwnerAccount(v string) *CloudCreateExtenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudCreateExtenRequest) SetResourceOwnerId(v int64) *CloudCreateExtenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudCreateExtenRequest) SetType(v int64) *CloudCreateExtenRequest {
	s.Type = &v
	return s
}

func (s *CloudCreateExtenRequest) Validate() error {
	return dara.Validate(s)
}
