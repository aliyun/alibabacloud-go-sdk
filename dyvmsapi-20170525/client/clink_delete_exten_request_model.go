// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDeleteExtenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *ClinkDeleteExtenRequest
	GetEnterpriseId() *int64
	SetExtenNumber(v string) *ClinkDeleteExtenRequest
	GetExtenNumber() *string
	SetOwnerId(v int64) *ClinkDeleteExtenRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ClinkDeleteExtenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkDeleteExtenRequest
	GetResourceOwnerId() *int64
}

type ClinkDeleteExtenRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 8004970
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 话机号码。3-6 位数字
	//
	// This parameter is required.
	//
	// example:
	//
	// 333
	ExtenNumber          *string `json:"ExtenNumber,omitempty" xml:"ExtenNumber,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ClinkDeleteExtenRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkDeleteExtenRequest) GoString() string {
	return s.String()
}

func (s *ClinkDeleteExtenRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkDeleteExtenRequest) GetExtenNumber() *string {
	return s.ExtenNumber
}

func (s *ClinkDeleteExtenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkDeleteExtenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkDeleteExtenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkDeleteExtenRequest) SetEnterpriseId(v int64) *ClinkDeleteExtenRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkDeleteExtenRequest) SetExtenNumber(v string) *ClinkDeleteExtenRequest {
	s.ExtenNumber = &v
	return s
}

func (s *ClinkDeleteExtenRequest) SetOwnerId(v int64) *ClinkDeleteExtenRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkDeleteExtenRequest) SetResourceOwnerAccount(v string) *ClinkDeleteExtenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkDeleteExtenRequest) SetResourceOwnerId(v int64) *ClinkDeleteExtenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkDeleteExtenRequest) Validate() error {
	return dara.Validate(s)
}
