// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetViberByRequestNoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetViberByRequestNoRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *GetViberByRequestNoRequest
	GetOwnerId() *int64
	SetRequestNo(v string) *GetViberByRequestNoRequest
	GetRequestNo() *string
	SetResourceOwnerAccount(v string) *GetViberByRequestNoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetViberByRequestNoRequest
	GetResourceOwnerId() *int64
}

type GetViberByRequestNoRequest struct {
	// example:
	//
	// 28251486512358****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 86512358****
	RequestNo            *string `json:"RequestNo,omitempty" xml:"RequestNo,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetViberByRequestNoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetViberByRequestNoRequest) GoString() string {
	return s.String()
}

func (s *GetViberByRequestNoRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetViberByRequestNoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetViberByRequestNoRequest) GetRequestNo() *string {
	return s.RequestNo
}

func (s *GetViberByRequestNoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetViberByRequestNoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetViberByRequestNoRequest) SetCustSpaceId(v string) *GetViberByRequestNoRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetViberByRequestNoRequest) SetOwnerId(v int64) *GetViberByRequestNoRequest {
	s.OwnerId = &v
	return s
}

func (s *GetViberByRequestNoRequest) SetRequestNo(v string) *GetViberByRequestNoRequest {
	s.RequestNo = &v
	return s
}

func (s *GetViberByRequestNoRequest) SetResourceOwnerAccount(v string) *GetViberByRequestNoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetViberByRequestNoRequest) SetResourceOwnerId(v int64) *GetViberByRequestNoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetViberByRequestNoRequest) Validate() error {
	return dara.Validate(s)
}
