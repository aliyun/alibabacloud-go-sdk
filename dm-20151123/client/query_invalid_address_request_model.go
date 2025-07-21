// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInvalidAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *QueryInvalidAddressRequest
	GetEndTime() *string
	SetKeyWord(v string) *QueryInvalidAddressRequest
	GetKeyWord() *string
	SetLength(v int32) *QueryInvalidAddressRequest
	GetLength() *int32
	SetNextStart(v string) *QueryInvalidAddressRequest
	GetNextStart() *string
	SetOwnerId(v int64) *QueryInvalidAddressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryInvalidAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryInvalidAddressRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *QueryInvalidAddressRequest
	GetStartTime() *string
}

type QueryInvalidAddressRequest struct {
	// End time, with a span from the start time that cannot exceed 30 days, in the format yyyy-MM-dd.
	//
	// example:
	//
	// 2019-09-29
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Keyword. If not provided, it represents all invalid addresses.
	//
	// example:
	//
	// info
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// Number of items per request.
	//
	// example:
	//
	// 100
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// Request starting position.
	//
	// example:
	//
	// ***
	NextStart            *string `json:"NextStart,omitempty" xml:"NextStart,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Start time, which cannot be earlier than 30 days ago, in the format yyyy-MM-dd.
	//
	// example:
	//
	// 2019-09-29
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryInvalidAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInvalidAddressRequest) GoString() string {
	return s.String()
}

func (s *QueryInvalidAddressRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryInvalidAddressRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *QueryInvalidAddressRequest) GetLength() *int32 {
	return s.Length
}

func (s *QueryInvalidAddressRequest) GetNextStart() *string {
	return s.NextStart
}

func (s *QueryInvalidAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryInvalidAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryInvalidAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryInvalidAddressRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryInvalidAddressRequest) SetEndTime(v string) *QueryInvalidAddressRequest {
	s.EndTime = &v
	return s
}

func (s *QueryInvalidAddressRequest) SetKeyWord(v string) *QueryInvalidAddressRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryInvalidAddressRequest) SetLength(v int32) *QueryInvalidAddressRequest {
	s.Length = &v
	return s
}

func (s *QueryInvalidAddressRequest) SetNextStart(v string) *QueryInvalidAddressRequest {
	s.NextStart = &v
	return s
}

func (s *QueryInvalidAddressRequest) SetOwnerId(v int64) *QueryInvalidAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryInvalidAddressRequest) SetResourceOwnerAccount(v string) *QueryInvalidAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryInvalidAddressRequest) SetResourceOwnerId(v int64) *QueryInvalidAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryInvalidAddressRequest) SetStartTime(v string) *QueryInvalidAddressRequest {
	s.StartTime = &v
	return s
}

func (s *QueryInvalidAddressRequest) Validate() error {
	return dara.Validate(s)
}
