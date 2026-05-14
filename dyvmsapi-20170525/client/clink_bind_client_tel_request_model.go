// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkBindClientTelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCno(v string) *ClinkBindClientTelRequest
	GetCno() *string
	SetEnterpriseId(v int64) *ClinkBindClientTelRequest
	GetEnterpriseId() *int64
	SetIsBind(v int64) *ClinkBindClientTelRequest
	GetIsBind() *int64
	SetIsReserveTel(v int64) *ClinkBindClientTelRequest
	GetIsReserveTel() *int64
	SetOwnerId(v int64) *ClinkBindClientTelRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ClinkBindClientTelRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkBindClientTelRequest
	GetResourceOwnerId() *int64
	SetTel(v string) *ClinkBindClientTelRequest
	GetTel() *string
}

type ClinkBindClientTelRequest struct {
	// 座席号
	//
	// This parameter is required.
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 8004970
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 是否绑定 1: 是，0: 否
	//
	// example:
	//
	// 1
	IsBind *int64 `json:"IsBind,omitempty" xml:"IsBind,omitempty"`
	// 是否同步为备用手机 1: 是，0: 否
	//
	// example:
	//
	// 1
	IsReserveTel         *int64  `json:"IsReserveTel,omitempty" xml:"IsReserveTel,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 电话号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 177xxxx7750
	Tel *string `json:"Tel,omitempty" xml:"Tel,omitempty"`
}

func (s ClinkBindClientTelRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkBindClientTelRequest) GoString() string {
	return s.String()
}

func (s *ClinkBindClientTelRequest) GetCno() *string {
	return s.Cno
}

func (s *ClinkBindClientTelRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkBindClientTelRequest) GetIsBind() *int64 {
	return s.IsBind
}

func (s *ClinkBindClientTelRequest) GetIsReserveTel() *int64 {
	return s.IsReserveTel
}

func (s *ClinkBindClientTelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkBindClientTelRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkBindClientTelRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkBindClientTelRequest) GetTel() *string {
	return s.Tel
}

func (s *ClinkBindClientTelRequest) SetCno(v string) *ClinkBindClientTelRequest {
	s.Cno = &v
	return s
}

func (s *ClinkBindClientTelRequest) SetEnterpriseId(v int64) *ClinkBindClientTelRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkBindClientTelRequest) SetIsBind(v int64) *ClinkBindClientTelRequest {
	s.IsBind = &v
	return s
}

func (s *ClinkBindClientTelRequest) SetIsReserveTel(v int64) *ClinkBindClientTelRequest {
	s.IsReserveTel = &v
	return s
}

func (s *ClinkBindClientTelRequest) SetOwnerId(v int64) *ClinkBindClientTelRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkBindClientTelRequest) SetResourceOwnerAccount(v string) *ClinkBindClientTelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkBindClientTelRequest) SetResourceOwnerId(v int64) *ClinkBindClientTelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkBindClientTelRequest) SetTel(v string) *ClinkBindClientTelRequest {
	s.Tel = &v
	return s
}

func (s *ClinkBindClientTelRequest) Validate() error {
	return dara.Validate(s)
}
