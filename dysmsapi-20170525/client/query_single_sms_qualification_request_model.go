// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySingleSmsQualificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *QuerySingleSmsQualificationRequest
	GetOrderId() *int64
	SetOwnerId(v int64) *QuerySingleSmsQualificationRequest
	GetOwnerId() *int64
	SetQualificationGroupId(v int64) *QuerySingleSmsQualificationRequest
	GetQualificationGroupId() *int64
	SetResourceOwnerAccount(v string) *QuerySingleSmsQualificationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySingleSmsQualificationRequest
	GetResourceOwnerId() *int64
}

type QuerySingleSmsQualificationRequest struct {
	// 工单id
	//
	// example:
	//
	// 2001****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 资质id
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

func (s QuerySingleSmsQualificationRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySingleSmsQualificationRequest) GoString() string {
	return s.String()
}

func (s *QuerySingleSmsQualificationRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *QuerySingleSmsQualificationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySingleSmsQualificationRequest) GetQualificationGroupId() *int64 {
	return s.QualificationGroupId
}

func (s *QuerySingleSmsQualificationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySingleSmsQualificationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySingleSmsQualificationRequest) SetOrderId(v int64) *QuerySingleSmsQualificationRequest {
	s.OrderId = &v
	return s
}

func (s *QuerySingleSmsQualificationRequest) SetOwnerId(v int64) *QuerySingleSmsQualificationRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySingleSmsQualificationRequest) SetQualificationGroupId(v int64) *QuerySingleSmsQualificationRequest {
	s.QualificationGroupId = &v
	return s
}

func (s *QuerySingleSmsQualificationRequest) SetResourceOwnerAccount(v string) *QuerySingleSmsQualificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySingleSmsQualificationRequest) SetResourceOwnerId(v int64) *QuerySingleSmsQualificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySingleSmsQualificationRequest) Validate() error {
	return dara.Validate(s)
}
