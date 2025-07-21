// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryChatappBindWabaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *QueryChatappBindWabaRequest
	GetCustSpaceId() *string
	SetIsvCode(v string) *QueryChatappBindWabaRequest
	GetIsvCode() *string
	SetOwnerId(v int64) *QueryChatappBindWabaRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryChatappBindWabaRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryChatappBindWabaRequest
	GetResourceOwnerId() *int64
}

type QueryChatappBindWabaRequest struct {
	// The space ID of the user under the ISV account.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The ISV verification code, which is used to verify whether the user is authorized by the ISV account.
	//
	// example:
	//
	// aksik93kdkkxmwol93939
	IsvCode              *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryChatappBindWabaRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryChatappBindWabaRequest) GoString() string {
	return s.String()
}

func (s *QueryChatappBindWabaRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *QueryChatappBindWabaRequest) GetIsvCode() *string {
	return s.IsvCode
}

func (s *QueryChatappBindWabaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryChatappBindWabaRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryChatappBindWabaRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryChatappBindWabaRequest) SetCustSpaceId(v string) *QueryChatappBindWabaRequest {
	s.CustSpaceId = &v
	return s
}

func (s *QueryChatappBindWabaRequest) SetIsvCode(v string) *QueryChatappBindWabaRequest {
	s.IsvCode = &v
	return s
}

func (s *QueryChatappBindWabaRequest) SetOwnerId(v int64) *QueryChatappBindWabaRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryChatappBindWabaRequest) SetResourceOwnerAccount(v string) *QueryChatappBindWabaRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryChatappBindWabaRequest) SetResourceOwnerId(v int64) *QueryChatappBindWabaRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryChatappBindWabaRequest) Validate() error {
	return dara.Validate(s)
}
