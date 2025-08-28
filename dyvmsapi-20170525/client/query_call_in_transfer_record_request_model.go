// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallInTransferRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallInCaller(v string) *QueryCallInTransferRecordRequest
	GetCallInCaller() *string
	SetOwnerId(v int64) *QueryCallInTransferRecordRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *QueryCallInTransferRecordRequest
	GetPageNo() *int64
	SetPageSize(v int64) *QueryCallInTransferRecordRequest
	GetPageSize() *int64
	SetPhoneNumber(v string) *QueryCallInTransferRecordRequest
	GetPhoneNumber() *string
	SetQueryDate(v string) *QueryCallInTransferRecordRequest
	GetQueryDate() *string
	SetResourceOwnerAccount(v string) *QueryCallInTransferRecordRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryCallInTransferRecordRequest
	GetResourceOwnerId() *int64
}

type QueryCallInTransferRecordRequest struct {
	// The calling number of the inbound call.
	//
	// example:
	//
	// 150****0000
	CallInCaller *string `json:"CallInCaller,omitempty" xml:"CallInCaller,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: 1 to 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The phone number to which a call is transferred.
	//
	// This parameter is required.
	//
	// example:
	//
	// 400****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The time at which call transfer records are queried, in the YYYY-MM-DD hh:mm:ss format.
	//
	// > The query result is all the call transfer records of the specified day.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-06-28 00:00:00
	QueryDate            *string `json:"QueryDate,omitempty" xml:"QueryDate,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryCallInTransferRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCallInTransferRecordRequest) GoString() string {
	return s.String()
}

func (s *QueryCallInTransferRecordRequest) GetCallInCaller() *string {
	return s.CallInCaller
}

func (s *QueryCallInTransferRecordRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryCallInTransferRecordRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QueryCallInTransferRecordRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryCallInTransferRecordRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *QueryCallInTransferRecordRequest) GetQueryDate() *string {
	return s.QueryDate
}

func (s *QueryCallInTransferRecordRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryCallInTransferRecordRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryCallInTransferRecordRequest) SetCallInCaller(v string) *QueryCallInTransferRecordRequest {
	s.CallInCaller = &v
	return s
}

func (s *QueryCallInTransferRecordRequest) SetOwnerId(v int64) *QueryCallInTransferRecordRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCallInTransferRecordRequest) SetPageNo(v int64) *QueryCallInTransferRecordRequest {
	s.PageNo = &v
	return s
}

func (s *QueryCallInTransferRecordRequest) SetPageSize(v int64) *QueryCallInTransferRecordRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCallInTransferRecordRequest) SetPhoneNumber(v string) *QueryCallInTransferRecordRequest {
	s.PhoneNumber = &v
	return s
}

func (s *QueryCallInTransferRecordRequest) SetQueryDate(v string) *QueryCallInTransferRecordRequest {
	s.QueryDate = &v
	return s
}

func (s *QueryCallInTransferRecordRequest) SetResourceOwnerAccount(v string) *QueryCallInTransferRecordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCallInTransferRecordRequest) SetResourceOwnerId(v int64) *QueryCallInTransferRecordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryCallInTransferRecordRequest) Validate() error {
	return dara.Validate(s)
}
