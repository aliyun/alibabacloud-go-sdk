// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmsQualificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *DeleteSmsQualificationRequest
	GetOrderId() *int64
	SetOwnerId(v int64) *DeleteSmsQualificationRequest
	GetOwnerId() *int64
	SetQualificationGroupId(v int64) *DeleteSmsQualificationRequest
	GetQualificationGroupId() *int64
	SetResourceOwnerAccount(v string) *DeleteSmsQualificationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSmsQualificationRequest
	GetResourceOwnerId() *int64
}

type DeleteSmsQualificationRequest struct {
	// 工单ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 2001****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 资质组ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000****
	QualificationGroupId *int64  `json:"QualificationGroupId,omitempty" xml:"QualificationGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteSmsQualificationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmsQualificationRequest) GoString() string {
	return s.String()
}

func (s *DeleteSmsQualificationRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *DeleteSmsQualificationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSmsQualificationRequest) GetQualificationGroupId() *int64 {
	return s.QualificationGroupId
}

func (s *DeleteSmsQualificationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSmsQualificationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSmsQualificationRequest) SetOrderId(v int64) *DeleteSmsQualificationRequest {
	s.OrderId = &v
	return s
}

func (s *DeleteSmsQualificationRequest) SetOwnerId(v int64) *DeleteSmsQualificationRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSmsQualificationRequest) SetQualificationGroupId(v int64) *DeleteSmsQualificationRequest {
	s.QualificationGroupId = &v
	return s
}

func (s *DeleteSmsQualificationRequest) SetResourceOwnerAccount(v string) *DeleteSmsQualificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSmsQualificationRequest) SetResourceOwnerId(v int64) *DeleteSmsQualificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSmsQualificationRequest) Validate() error {
	return dara.Validate(s)
}
