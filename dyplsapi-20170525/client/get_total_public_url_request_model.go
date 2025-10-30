// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTotalPublicUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *GetTotalPublicUrlRequest
	GetCallId() *string
	SetCallTime(v string) *GetTotalPublicUrlRequest
	GetCallTime() *string
	SetCheckSubs(v bool) *GetTotalPublicUrlRequest
	GetCheckSubs() *bool
	SetOwnerId(v int64) *GetTotalPublicUrlRequest
	GetOwnerId() *int64
	SetPartnerKey(v string) *GetTotalPublicUrlRequest
	GetPartnerKey() *string
	SetResourceOwnerAccount(v string) *GetTotalPublicUrlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetTotalPublicUrlRequest
	GetResourceOwnerId() *int64
}

type GetTotalPublicUrlRequest struct {
	// The ID of the call record.
	//
	// Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view **Call Record ID*	- on the **Call Record Query*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2568900****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The call initiation time in the call record.
	//
	// Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account). View **Call Initiated At*	- on the **Call Record Query*	- page, or view the call_time field in the Call Detail Record (CDR) receipt.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-03-05 12:00:00
	CallTime *string `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	// Specifies whether the verification on the binding ID is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	CheckSubs *bool  `json:"CheckSubs,omitempty" xml:"CheckSubs,omitempty"`
	OwnerId   *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC12256****
	PartnerKey           *string `json:"PartnerKey,omitempty" xml:"PartnerKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetTotalPublicUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTotalPublicUrlRequest) GoString() string {
	return s.String()
}

func (s *GetTotalPublicUrlRequest) GetCallId() *string {
	return s.CallId
}

func (s *GetTotalPublicUrlRequest) GetCallTime() *string {
	return s.CallTime
}

func (s *GetTotalPublicUrlRequest) GetCheckSubs() *bool {
	return s.CheckSubs
}

func (s *GetTotalPublicUrlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetTotalPublicUrlRequest) GetPartnerKey() *string {
	return s.PartnerKey
}

func (s *GetTotalPublicUrlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetTotalPublicUrlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetTotalPublicUrlRequest) SetCallId(v string) *GetTotalPublicUrlRequest {
	s.CallId = &v
	return s
}

func (s *GetTotalPublicUrlRequest) SetCallTime(v string) *GetTotalPublicUrlRequest {
	s.CallTime = &v
	return s
}

func (s *GetTotalPublicUrlRequest) SetCheckSubs(v bool) *GetTotalPublicUrlRequest {
	s.CheckSubs = &v
	return s
}

func (s *GetTotalPublicUrlRequest) SetOwnerId(v int64) *GetTotalPublicUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *GetTotalPublicUrlRequest) SetPartnerKey(v string) *GetTotalPublicUrlRequest {
	s.PartnerKey = &v
	return s
}

func (s *GetTotalPublicUrlRequest) SetResourceOwnerAccount(v string) *GetTotalPublicUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetTotalPublicUrlRequest) SetResourceOwnerId(v int64) *GetTotalPublicUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetTotalPublicUrlRequest) Validate() error {
	return dara.Validate(s)
}
