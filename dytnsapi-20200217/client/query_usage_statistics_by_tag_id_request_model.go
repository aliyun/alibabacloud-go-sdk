// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUsageStatisticsByTagIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v string) *QueryUsageStatisticsByTagIdRequest
	GetBeginTime() *string
	SetEndTime(v string) *QueryUsageStatisticsByTagIdRequest
	GetEndTime() *string
	SetOwnerId(v int64) *QueryUsageStatisticsByTagIdRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *QueryUsageStatisticsByTagIdRequest
	GetPageNo() *int64
	SetPageSize(v int64) *QueryUsageStatisticsByTagIdRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *QueryUsageStatisticsByTagIdRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryUsageStatisticsByTagIdRequest
	GetResourceOwnerId() *int64
	SetTagId(v int64) *QueryUsageStatisticsByTagIdRequest
	GetTagId() *int64
}

type QueryUsageStatisticsByTagIdRequest struct {
	// The beginning of the time range to query.
	//
	// example:
	//
	// 20230308
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end of the time range to query.
	//
	// example:
	//
	// 20230406
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag ID.
	//
	// example:
	//
	// 14
	TagId *int64 `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s QueryUsageStatisticsByTagIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUsageStatisticsByTagIdRequest) GoString() string {
	return s.String()
}

func (s *QueryUsageStatisticsByTagIdRequest) GetBeginTime() *string {
	return s.BeginTime
}

func (s *QueryUsageStatisticsByTagIdRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryUsageStatisticsByTagIdRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryUsageStatisticsByTagIdRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QueryUsageStatisticsByTagIdRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryUsageStatisticsByTagIdRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryUsageStatisticsByTagIdRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryUsageStatisticsByTagIdRequest) GetTagId() *int64 {
	return s.TagId
}

func (s *QueryUsageStatisticsByTagIdRequest) SetBeginTime(v string) *QueryUsageStatisticsByTagIdRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdRequest) SetEndTime(v string) *QueryUsageStatisticsByTagIdRequest {
	s.EndTime = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdRequest) SetOwnerId(v int64) *QueryUsageStatisticsByTagIdRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdRequest) SetPageNo(v int64) *QueryUsageStatisticsByTagIdRequest {
	s.PageNo = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdRequest) SetPageSize(v int64) *QueryUsageStatisticsByTagIdRequest {
	s.PageSize = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdRequest) SetResourceOwnerAccount(v string) *QueryUsageStatisticsByTagIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdRequest) SetResourceOwnerId(v int64) *QueryUsageStatisticsByTagIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdRequest) SetTagId(v int64) *QueryUsageStatisticsByTagIdRequest {
	s.TagId = &v
	return s
}

func (s *QueryUsageStatisticsByTagIdRequest) Validate() error {
	return dara.Validate(s)
}
