// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPickUpWaybillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCancelDesc(v string) *CancelPickUpWaybillRequest
	GetCancelDesc() *string
	SetOuterOrderCode(v string) *CancelPickUpWaybillRequest
	GetOuterOrderCode() *string
	SetOwnerId(v int64) *CancelPickUpWaybillRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CancelPickUpWaybillRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelPickUpWaybillRequest
	GetResourceOwnerId() *int64
}

type CancelPickUpWaybillRequest struct {
	// The cancellation reason.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"action\\":\\"UPDATE_DESC\\",\\"value\\":\\"The courier is unable to pick up the package.\\"}
	CancelDesc *string `json:"CancelDesc,omitempty" xml:"CancelDesc,omitempty"`
	// The ID of the external order.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1145678823****
	OuterOrderCode       *string `json:"OuterOrderCode,omitempty" xml:"OuterOrderCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CancelPickUpWaybillRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelPickUpWaybillRequest) GoString() string {
	return s.String()
}

func (s *CancelPickUpWaybillRequest) GetCancelDesc() *string {
	return s.CancelDesc
}

func (s *CancelPickUpWaybillRequest) GetOuterOrderCode() *string {
	return s.OuterOrderCode
}

func (s *CancelPickUpWaybillRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelPickUpWaybillRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelPickUpWaybillRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelPickUpWaybillRequest) SetCancelDesc(v string) *CancelPickUpWaybillRequest {
	s.CancelDesc = &v
	return s
}

func (s *CancelPickUpWaybillRequest) SetOuterOrderCode(v string) *CancelPickUpWaybillRequest {
	s.OuterOrderCode = &v
	return s
}

func (s *CancelPickUpWaybillRequest) SetOwnerId(v int64) *CancelPickUpWaybillRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelPickUpWaybillRequest) SetResourceOwnerAccount(v string) *CancelPickUpWaybillRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelPickUpWaybillRequest) SetResourceOwnerId(v int64) *CancelPickUpWaybillRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelPickUpWaybillRequest) Validate() error {
	return dara.Validate(s)
}
