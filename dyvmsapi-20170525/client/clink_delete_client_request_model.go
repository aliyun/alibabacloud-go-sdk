// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDeleteClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCno(v int64) *ClinkDeleteClientRequest
	GetCno() *int64
	SetEnterpriseId(v int64) *ClinkDeleteClientRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *ClinkDeleteClientRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ClinkDeleteClientRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkDeleteClientRequest
	GetResourceOwnerId() *int64
}

type ClinkDeleteClientRequest struct {
	// 座席工号，长度为4-11个字符，必须全部为数字
	//
	// This parameter is required.
	//
	// example:
	//
	// 14
	Cno *int64 `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 28
	EnterpriseId         *int64  `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ClinkDeleteClientRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkDeleteClientRequest) GoString() string {
	return s.String()
}

func (s *ClinkDeleteClientRequest) GetCno() *int64 {
	return s.Cno
}

func (s *ClinkDeleteClientRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkDeleteClientRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkDeleteClientRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkDeleteClientRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkDeleteClientRequest) SetCno(v int64) *ClinkDeleteClientRequest {
	s.Cno = &v
	return s
}

func (s *ClinkDeleteClientRequest) SetEnterpriseId(v int64) *ClinkDeleteClientRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkDeleteClientRequest) SetOwnerId(v int64) *ClinkDeleteClientRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkDeleteClientRequest) SetResourceOwnerAccount(v string) *ClinkDeleteClientRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkDeleteClientRequest) SetResourceOwnerId(v int64) *ClinkDeleteClientRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkDeleteClientRequest) Validate() error {
	return dara.Validate(s)
}
