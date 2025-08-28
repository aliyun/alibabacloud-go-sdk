// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineQualificationByOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *GetHotlineQualificationByOrderRequest
	GetOrderId() *string
	SetOwnerId(v int64) *GetHotlineQualificationByOrderRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetHotlineQualificationByOrderRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetHotlineQualificationByOrderRequest
	GetResourceOwnerId() *int64
}

type GetHotlineQualificationByOrderRequest struct {
	// The ticket ID.
	//
	// You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Qualification\\&Communication Script Management*	- > **Qualification Management**, and then click the **400 Qualifications*	- tab to view the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22456****
	OrderId              *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetHotlineQualificationByOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineQualificationByOrderRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineQualificationByOrderRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *GetHotlineQualificationByOrderRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetHotlineQualificationByOrderRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetHotlineQualificationByOrderRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetHotlineQualificationByOrderRequest) SetOrderId(v string) *GetHotlineQualificationByOrderRequest {
	s.OrderId = &v
	return s
}

func (s *GetHotlineQualificationByOrderRequest) SetOwnerId(v int64) *GetHotlineQualificationByOrderRequest {
	s.OwnerId = &v
	return s
}

func (s *GetHotlineQualificationByOrderRequest) SetResourceOwnerAccount(v string) *GetHotlineQualificationByOrderRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetHotlineQualificationByOrderRequest) SetResourceOwnerId(v int64) *GetHotlineQualificationByOrderRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetHotlineQualificationByOrderRequest) Validate() error {
	return dara.Validate(s)
}
