// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteExtenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudDeleteExtenRequest
	GetEnterpriseId() *int64
	SetExten(v int64) *CloudDeleteExtenRequest
	GetExten() *int64
	SetOwnerId(v int64) *CloudDeleteExtenRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudDeleteExtenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudDeleteExtenRequest
	GetResourceOwnerId() *int64
}

type CloudDeleteExtenRequest struct {
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
	Exten                *int64  `json:"Exten,omitempty" xml:"Exten,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloudDeleteExtenRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteExtenRequest) GoString() string {
	return s.String()
}

func (s *CloudDeleteExtenRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudDeleteExtenRequest) GetExten() *int64 {
	return s.Exten
}

func (s *CloudDeleteExtenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudDeleteExtenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudDeleteExtenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudDeleteExtenRequest) SetEnterpriseId(v int64) *CloudDeleteExtenRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudDeleteExtenRequest) SetExten(v int64) *CloudDeleteExtenRequest {
	s.Exten = &v
	return s
}

func (s *CloudDeleteExtenRequest) SetOwnerId(v int64) *CloudDeleteExtenRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudDeleteExtenRequest) SetResourceOwnerAccount(v string) *CloudDeleteExtenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudDeleteExtenRequest) SetResourceOwnerId(v int64) *CloudDeleteExtenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudDeleteExtenRequest) Validate() error {
	return dara.Validate(s)
}
