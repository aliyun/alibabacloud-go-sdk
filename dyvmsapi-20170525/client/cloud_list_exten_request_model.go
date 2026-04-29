// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListExtenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAreaCode(v string) *CloudListExtenRequest
	GetAreaCode() *string
	SetCallPower(v string) *CloudListExtenRequest
	GetCallPower() *string
	SetEnterpriseId(v int64) *CloudListExtenRequest
	GetEnterpriseId() *int64
	SetExten(v string) *CloudListExtenRequest
	GetExten() *string
	SetIbRecord(v int64) *CloudListExtenRequest
	GetIbRecord() *int64
	SetIsBind(v int64) *CloudListExtenRequest
	GetIsBind() *int64
	SetIsOb(v string) *CloudListExtenRequest
	GetIsOb() *string
	SetJitterBuffer(v int64) *CloudListExtenRequest
	GetJitterBuffer() *int64
	SetLimit(v int64) *CloudListExtenRequest
	GetLimit() *int64
	SetObRecord(v int64) *CloudListExtenRequest
	GetObRecord() *int64
	SetOffset(v int64) *CloudListExtenRequest
	GetOffset() *int64
	SetOwnerId(v int64) *CloudListExtenRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudListExtenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudListExtenRequest
	GetResourceOwnerId() *int64
	SetType(v int64) *CloudListExtenRequest
	GetType() *int64
}

type CloudListExtenRequest struct {
	// 区号
	//
	// example:
	//
	// 010
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// 呼叫权限，0：无限制，1：国内长途，2：国内本市，3：内部呼叫，默认无限制
	//
	// example:
	//
	// 0
	CallPower *string `json:"CallPower,omitempty" xml:"CallPower,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 分机号
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
	// 是否被座席绑定，1 绑定 0未绑定
	//
	// example:
	//
	// 0
	IsBind *int64 `json:"IsBind,omitempty" xml:"IsBind,omitempty"`
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
	// 条数；正整数 大于0 小于等于500 默认为10条
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// 外呼是否录音，0：不录音，1：录音，默认录音
	//
	// example:
	//
	// 1
	ObRecord *int64 `json:"ObRecord,omitempty" xml:"ObRecord,omitempty"`
	// 从第几条开始；正整数 默认为0
	//
	// example:
	//
	// 0
	Offset               *int64  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 类型，1. 注册到IAD分机 2.注册到webrtc 3.注册到远程话机
	//
	// example:
	//
	// 2
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CloudListExtenRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudListExtenRequest) GoString() string {
	return s.String()
}

func (s *CloudListExtenRequest) GetAreaCode() *string {
	return s.AreaCode
}

func (s *CloudListExtenRequest) GetCallPower() *string {
	return s.CallPower
}

func (s *CloudListExtenRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudListExtenRequest) GetExten() *string {
	return s.Exten
}

func (s *CloudListExtenRequest) GetIbRecord() *int64 {
	return s.IbRecord
}

func (s *CloudListExtenRequest) GetIsBind() *int64 {
	return s.IsBind
}

func (s *CloudListExtenRequest) GetIsOb() *string {
	return s.IsOb
}

func (s *CloudListExtenRequest) GetJitterBuffer() *int64 {
	return s.JitterBuffer
}

func (s *CloudListExtenRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *CloudListExtenRequest) GetObRecord() *int64 {
	return s.ObRecord
}

func (s *CloudListExtenRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *CloudListExtenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudListExtenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudListExtenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudListExtenRequest) GetType() *int64 {
	return s.Type
}

func (s *CloudListExtenRequest) SetAreaCode(v string) *CloudListExtenRequest {
	s.AreaCode = &v
	return s
}

func (s *CloudListExtenRequest) SetCallPower(v string) *CloudListExtenRequest {
	s.CallPower = &v
	return s
}

func (s *CloudListExtenRequest) SetEnterpriseId(v int64) *CloudListExtenRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListExtenRequest) SetExten(v string) *CloudListExtenRequest {
	s.Exten = &v
	return s
}

func (s *CloudListExtenRequest) SetIbRecord(v int64) *CloudListExtenRequest {
	s.IbRecord = &v
	return s
}

func (s *CloudListExtenRequest) SetIsBind(v int64) *CloudListExtenRequest {
	s.IsBind = &v
	return s
}

func (s *CloudListExtenRequest) SetIsOb(v string) *CloudListExtenRequest {
	s.IsOb = &v
	return s
}

func (s *CloudListExtenRequest) SetJitterBuffer(v int64) *CloudListExtenRequest {
	s.JitterBuffer = &v
	return s
}

func (s *CloudListExtenRequest) SetLimit(v int64) *CloudListExtenRequest {
	s.Limit = &v
	return s
}

func (s *CloudListExtenRequest) SetObRecord(v int64) *CloudListExtenRequest {
	s.ObRecord = &v
	return s
}

func (s *CloudListExtenRequest) SetOffset(v int64) *CloudListExtenRequest {
	s.Offset = &v
	return s
}

func (s *CloudListExtenRequest) SetOwnerId(v int64) *CloudListExtenRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudListExtenRequest) SetResourceOwnerAccount(v string) *CloudListExtenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudListExtenRequest) SetResourceOwnerId(v int64) *CloudListExtenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudListExtenRequest) SetType(v int64) *CloudListExtenRequest {
	s.Type = &v
	return s
}

func (s *CloudListExtenRequest) Validate() error {
	return dara.Validate(s)
}
