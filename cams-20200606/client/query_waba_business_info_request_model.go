// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWabaBusinessInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *QueryWabaBusinessInfoRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *QueryWabaBusinessInfoRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryWabaBusinessInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryWabaBusinessInfoRequest
	GetResourceOwnerId() *int64
	SetWabaId(v string) *QueryWabaBusinessInfoRequest
	GetWabaId() *string
}

type QueryWabaBusinessInfoRequest struct {
	// The space ID of the RAM user within the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493****
	CustSpaceId          *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the WhatsApp Business Account (WABA).
	//
	// This parameter is required.
	//
	// example:
	//
	// 293848822333
	WabaId *string `json:"WabaId,omitempty" xml:"WabaId,omitempty"`
}

func (s QueryWabaBusinessInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryWabaBusinessInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryWabaBusinessInfoRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *QueryWabaBusinessInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryWabaBusinessInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryWabaBusinessInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryWabaBusinessInfoRequest) GetWabaId() *string {
	return s.WabaId
}

func (s *QueryWabaBusinessInfoRequest) SetCustSpaceId(v string) *QueryWabaBusinessInfoRequest {
	s.CustSpaceId = &v
	return s
}

func (s *QueryWabaBusinessInfoRequest) SetOwnerId(v int64) *QueryWabaBusinessInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryWabaBusinessInfoRequest) SetResourceOwnerAccount(v string) *QueryWabaBusinessInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryWabaBusinessInfoRequest) SetResourceOwnerId(v int64) *QueryWabaBusinessInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryWabaBusinessInfoRequest) SetWabaId(v string) *QueryWabaBusinessInfoRequest {
	s.WabaId = &v
	return s
}

func (s *QueryWabaBusinessInfoRequest) Validate() error {
	return dara.Validate(s)
}
