// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappPhoneNumberMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetChatappPhoneNumberMetricRequest
	GetCustSpaceId() *string
	SetEnd(v int64) *GetChatappPhoneNumberMetricRequest
	GetEnd() *int64
	SetGranularity(v string) *GetChatappPhoneNumberMetricRequest
	GetGranularity() *string
	SetIsvCode(v string) *GetChatappPhoneNumberMetricRequest
	GetIsvCode() *string
	SetOwnerId(v int64) *GetChatappPhoneNumberMetricRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *GetChatappPhoneNumberMetricRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *GetChatappPhoneNumberMetricRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetChatappPhoneNumberMetricRequest
	GetResourceOwnerId() *int64
	SetStart(v int64) *GetChatappPhoneNumberMetricRequest
	GetStart() *int64
}

type GetChatappPhoneNumberMetricRequest struct {
	// The space ID of the RAM user within the ISV account.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The end of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1693407714687
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// The granularity of the metric.
	//
	// Valid values:
	//
	// 	- DAILY
	//
	// 	- HALF_HOUR
	//
	// example:
	//
	// DAILY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The independent software vendor (ISV) verification code, which is used to verify whether the RAM user is authorized by the ISV account.
	//
	// example:
	//
	// skdi3kksloslikdkkdk
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The business phone number.
	//
	// example:
	//
	// 861380000
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1693107714687
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetChatappPhoneNumberMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChatappPhoneNumberMetricRequest) GoString() string {
	return s.String()
}

func (s *GetChatappPhoneNumberMetricRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetChatappPhoneNumberMetricRequest) GetEnd() *int64 {
	return s.End
}

func (s *GetChatappPhoneNumberMetricRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *GetChatappPhoneNumberMetricRequest) GetIsvCode() *string {
	return s.IsvCode
}

func (s *GetChatappPhoneNumberMetricRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetChatappPhoneNumberMetricRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetChatappPhoneNumberMetricRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetChatappPhoneNumberMetricRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetChatappPhoneNumberMetricRequest) GetStart() *int64 {
	return s.Start
}

func (s *GetChatappPhoneNumberMetricRequest) SetCustSpaceId(v string) *GetChatappPhoneNumberMetricRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetChatappPhoneNumberMetricRequest) SetEnd(v int64) *GetChatappPhoneNumberMetricRequest {
	s.End = &v
	return s
}

func (s *GetChatappPhoneNumberMetricRequest) SetGranularity(v string) *GetChatappPhoneNumberMetricRequest {
	s.Granularity = &v
	return s
}

func (s *GetChatappPhoneNumberMetricRequest) SetIsvCode(v string) *GetChatappPhoneNumberMetricRequest {
	s.IsvCode = &v
	return s
}

func (s *GetChatappPhoneNumberMetricRequest) SetOwnerId(v int64) *GetChatappPhoneNumberMetricRequest {
	s.OwnerId = &v
	return s
}

func (s *GetChatappPhoneNumberMetricRequest) SetPhoneNumber(v string) *GetChatappPhoneNumberMetricRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetChatappPhoneNumberMetricRequest) SetResourceOwnerAccount(v string) *GetChatappPhoneNumberMetricRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetChatappPhoneNumberMetricRequest) SetResourceOwnerId(v int64) *GetChatappPhoneNumberMetricRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetChatappPhoneNumberMetricRequest) SetStart(v int64) *GetChatappPhoneNumberMetricRequest {
	s.Start = &v
	return s
}

func (s *GetChatappPhoneNumberMetricRequest) Validate() error {
	return dara.Validate(s)
}
