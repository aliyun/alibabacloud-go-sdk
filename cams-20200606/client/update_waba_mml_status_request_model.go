// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWabaMmlStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateWabaMmlStatusRequest
	GetCode() *string
	SetCustSpaceId(v string) *UpdateWabaMmlStatusRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *UpdateWabaMmlStatusRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateWabaMmlStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateWabaMmlStatusRequest
	GetResourceOwnerId() *int64
	SetWabaId(v string) *UpdateWabaMmlStatusRequest
	GetWabaId() *string
}

type UpdateWabaMmlStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	CustSpaceId          *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// waba Id。
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	WabaId *string `json:"WabaId,omitempty" xml:"WabaId,omitempty"`
}

func (s UpdateWabaMmlStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWabaMmlStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateWabaMmlStatusRequest) GetCode() *string {
	return s.Code
}

func (s *UpdateWabaMmlStatusRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *UpdateWabaMmlStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateWabaMmlStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateWabaMmlStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateWabaMmlStatusRequest) GetWabaId() *string {
	return s.WabaId
}

func (s *UpdateWabaMmlStatusRequest) SetCode(v string) *UpdateWabaMmlStatusRequest {
	s.Code = &v
	return s
}

func (s *UpdateWabaMmlStatusRequest) SetCustSpaceId(v string) *UpdateWabaMmlStatusRequest {
	s.CustSpaceId = &v
	return s
}

func (s *UpdateWabaMmlStatusRequest) SetOwnerId(v int64) *UpdateWabaMmlStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateWabaMmlStatusRequest) SetResourceOwnerAccount(v string) *UpdateWabaMmlStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateWabaMmlStatusRequest) SetResourceOwnerId(v int64) *UpdateWabaMmlStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateWabaMmlStatusRequest) SetWabaId(v string) *UpdateWabaMmlStatusRequest {
	s.WabaId = &v
	return s
}

func (s *UpdateWabaMmlStatusRequest) Validate() error {
	return dara.Validate(s)
}
