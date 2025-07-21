// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserSuppressionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *ListUserSuppressionRequest
	GetAddress() *string
	SetEndBounceTime(v int32) *ListUserSuppressionRequest
	GetEndBounceTime() *int32
	SetEndCreateTime(v int32) *ListUserSuppressionRequest
	GetEndCreateTime() *int32
	SetOwnerId(v int64) *ListUserSuppressionRequest
	GetOwnerId() *int64
	SetPageNo(v int32) *ListUserSuppressionRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListUserSuppressionRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *ListUserSuppressionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListUserSuppressionRequest
	GetResourceOwnerId() *int64
	SetStartBounceTime(v int32) *ListUserSuppressionRequest
	GetStartBounceTime() *int32
	SetStartCreateTime(v int32) *ListUserSuppressionRequest
	GetStartCreateTime() *int32
}

type ListUserSuppressionRequest struct {
	// Email address or domain name
	//
	// example:
	//
	// test@example.net
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// End time of the last bounce hit, timestamp, accurate to the second. The span between start and end times cannot exceed 7 days.
	//
	// example:
	//
	// 1715669077
	EndBounceTime *int32 `json:"EndBounceTime,omitempty" xml:"EndBounceTime,omitempty"`
	// End creation time, timestamp, accurate to the second. The span between start and end times cannot exceed 7 days.
	//
	// example:
	//
	// 1715669077
	EndCreateTime *int32 `json:"EndCreateTime,omitempty" xml:"EndCreateTime,omitempty"`
	OwnerId       *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Current page number
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Start time of the last bounce hit, timestamp, accurate to the second.
	//
	// example:
	//
	// 1715668852
	StartBounceTime *int32 `json:"StartBounceTime,omitempty" xml:"StartBounceTime,omitempty"`
	// Start creation time, timestamp, accurate to the second.
	//
	// example:
	//
	// 1715668852
	StartCreateTime *int32 `json:"StartCreateTime,omitempty" xml:"StartCreateTime,omitempty"`
}

func (s ListUserSuppressionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserSuppressionRequest) GoString() string {
	return s.String()
}

func (s *ListUserSuppressionRequest) GetAddress() *string {
	return s.Address
}

func (s *ListUserSuppressionRequest) GetEndBounceTime() *int32 {
	return s.EndBounceTime
}

func (s *ListUserSuppressionRequest) GetEndCreateTime() *int32 {
	return s.EndCreateTime
}

func (s *ListUserSuppressionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListUserSuppressionRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListUserSuppressionRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserSuppressionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListUserSuppressionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListUserSuppressionRequest) GetStartBounceTime() *int32 {
	return s.StartBounceTime
}

func (s *ListUserSuppressionRequest) GetStartCreateTime() *int32 {
	return s.StartCreateTime
}

func (s *ListUserSuppressionRequest) SetAddress(v string) *ListUserSuppressionRequest {
	s.Address = &v
	return s
}

func (s *ListUserSuppressionRequest) SetEndBounceTime(v int32) *ListUserSuppressionRequest {
	s.EndBounceTime = &v
	return s
}

func (s *ListUserSuppressionRequest) SetEndCreateTime(v int32) *ListUserSuppressionRequest {
	s.EndCreateTime = &v
	return s
}

func (s *ListUserSuppressionRequest) SetOwnerId(v int64) *ListUserSuppressionRequest {
	s.OwnerId = &v
	return s
}

func (s *ListUserSuppressionRequest) SetPageNo(v int32) *ListUserSuppressionRequest {
	s.PageNo = &v
	return s
}

func (s *ListUserSuppressionRequest) SetPageSize(v int32) *ListUserSuppressionRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserSuppressionRequest) SetResourceOwnerAccount(v string) *ListUserSuppressionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListUserSuppressionRequest) SetResourceOwnerId(v int64) *ListUserSuppressionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListUserSuppressionRequest) SetStartBounceTime(v int32) *ListUserSuppressionRequest {
	s.StartBounceTime = &v
	return s
}

func (s *ListUserSuppressionRequest) SetStartCreateTime(v int32) *ListUserSuppressionRequest {
	s.StartCreateTime = &v
	return s
}

func (s *ListUserSuppressionRequest) Validate() error {
	return dara.Validate(s)
}
