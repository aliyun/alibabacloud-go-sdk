// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkCreateExtenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllow(v int64) *ClinkCreateExtenRequest
	GetAllow() *int64
	SetAreaCode(v string) *ClinkCreateExtenRequest
	GetAreaCode() *string
	SetEnterpriseId(v int64) *ClinkCreateExtenRequest
	GetEnterpriseId() *int64
	SetExtenNumber(v string) *ClinkCreateExtenRequest
	GetExtenNumber() *string
	SetIsDirect(v int64) *ClinkCreateExtenRequest
	GetIsDirect() *int64
	SetJittBuffer(v int64) *ClinkCreateExtenRequest
	GetJittBuffer() *int64
	SetOwnerId(v int64) *ClinkCreateExtenRequest
	GetOwnerId() *int64
	SetPassword(v string) *ClinkCreateExtenRequest
	GetPassword() *string
	SetResourceOwnerAccount(v string) *ClinkCreateExtenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkCreateExtenRequest
	GetResourceOwnerId() *int64
	SetType(v int64) *ClinkCreateExtenRequest
	GetType() *int64
}

type ClinkCreateExtenRequest struct {
	// 语音编码。 IP话机支持的语音编码：1、2、3、4、5 软电话支持的语音编码：5、6 1：g729(已弃用) 2：g729,alaw,ulaw 3：alaw,ulaw,g729 4：myopus,alaw,ulaw,g729 5：alaw,ulaw 6：myopus,alaw,ulaw
	//
	// This parameter is required.
	//
	// example:
	//
	// 6
	Allow *int64 `json:"Allow,omitempty" xml:"Allow,omitempty"`
	// 话机区号。以 0 开头 3-4 位数字
	//
	// This parameter is required.
	//
	// example:
	//
	// 010
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 8004970
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 话机号码。3-6 位数字，要求唯一
	//
	// This parameter is required.
	//
	// example:
	//
	// 12235
	ExtenNumber *string `json:"ExtenNumber,omitempty" xml:"ExtenNumber,omitempty"`
	// 是否允许主叫外呼，1：允许；0：不允许。话机类型为IP话机时，可以开启主叫外呼功能，直接通过IP话机进行外呼。若要使用主叫外呼功能，需要开启企业级开关。
	//
	// example:
	//
	// 1
	IsDirect *int64 `json:"IsDirect,omitempty" xml:"IsDirect,omitempty"`
	// 网络防抖，0：关闭；1：开启，默认关闭，当开启时，语音编码类型默认为5
	//
	// example:
	//
	// 70
	JittBuffer *int64 `json:"JittBuffer,omitempty" xml:"JittBuffer,omitempty"`
	OwnerId    *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 话机密码，密码规则参见企业配置
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	Password             *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 话机类型。1: IP话机， 2: 软电话
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ClinkCreateExtenRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateExtenRequest) GoString() string {
	return s.String()
}

func (s *ClinkCreateExtenRequest) GetAllow() *int64 {
	return s.Allow
}

func (s *ClinkCreateExtenRequest) GetAreaCode() *string {
	return s.AreaCode
}

func (s *ClinkCreateExtenRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkCreateExtenRequest) GetExtenNumber() *string {
	return s.ExtenNumber
}

func (s *ClinkCreateExtenRequest) GetIsDirect() *int64 {
	return s.IsDirect
}

func (s *ClinkCreateExtenRequest) GetJittBuffer() *int64 {
	return s.JittBuffer
}

func (s *ClinkCreateExtenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkCreateExtenRequest) GetPassword() *string {
	return s.Password
}

func (s *ClinkCreateExtenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkCreateExtenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkCreateExtenRequest) GetType() *int64 {
	return s.Type
}

func (s *ClinkCreateExtenRequest) SetAllow(v int64) *ClinkCreateExtenRequest {
	s.Allow = &v
	return s
}

func (s *ClinkCreateExtenRequest) SetAreaCode(v string) *ClinkCreateExtenRequest {
	s.AreaCode = &v
	return s
}

func (s *ClinkCreateExtenRequest) SetEnterpriseId(v int64) *ClinkCreateExtenRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkCreateExtenRequest) SetExtenNumber(v string) *ClinkCreateExtenRequest {
	s.ExtenNumber = &v
	return s
}

func (s *ClinkCreateExtenRequest) SetIsDirect(v int64) *ClinkCreateExtenRequest {
	s.IsDirect = &v
	return s
}

func (s *ClinkCreateExtenRequest) SetJittBuffer(v int64) *ClinkCreateExtenRequest {
	s.JittBuffer = &v
	return s
}

func (s *ClinkCreateExtenRequest) SetOwnerId(v int64) *ClinkCreateExtenRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkCreateExtenRequest) SetPassword(v string) *ClinkCreateExtenRequest {
	s.Password = &v
	return s
}

func (s *ClinkCreateExtenRequest) SetResourceOwnerAccount(v string) *ClinkCreateExtenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkCreateExtenRequest) SetResourceOwnerId(v int64) *ClinkCreateExtenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkCreateExtenRequest) SetType(v int64) *ClinkCreateExtenRequest {
	s.Type = &v
	return s
}

func (s *ClinkCreateExtenRequest) Validate() error {
	return dara.Validate(s)
}
