// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallInPoolTransferConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryCallInPoolTransferConfigRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *QueryCallInPoolTransferConfigRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *QueryCallInPoolTransferConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryCallInPoolTransferConfigRequest
	GetResourceOwnerId() *int64
}

type QueryCallInPoolTransferConfigRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The China 400 number used to transfer the call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 400****
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryCallInPoolTransferConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCallInPoolTransferConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryCallInPoolTransferConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryCallInPoolTransferConfigRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *QueryCallInPoolTransferConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryCallInPoolTransferConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryCallInPoolTransferConfigRequest) SetOwnerId(v int64) *QueryCallInPoolTransferConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCallInPoolTransferConfigRequest) SetPhoneNumber(v string) *QueryCallInPoolTransferConfigRequest {
	s.PhoneNumber = &v
	return s
}

func (s *QueryCallInPoolTransferConfigRequest) SetResourceOwnerAccount(v string) *QueryCallInPoolTransferConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCallInPoolTransferConfigRequest) SetResourceOwnerId(v int64) *QueryCallInPoolTransferConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryCallInPoolTransferConfigRequest) Validate() error {
	return dara.Validate(s)
}
