// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrackListByMailFromAndTagNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GetTrackListByMailFromAndTagNameRequest
	GetAccountName() *string
	SetDedicatedIp(v string) *GetTrackListByMailFromAndTagNameRequest
	GetDedicatedIp() *string
	SetDedicatedIpPoolId(v string) *GetTrackListByMailFromAndTagNameRequest
	GetDedicatedIpPoolId() *string
	SetEndTime(v string) *GetTrackListByMailFromAndTagNameRequest
	GetEndTime() *string
	SetEsp(v string) *GetTrackListByMailFromAndTagNameRequest
	GetEsp() *string
	SetOffset(v string) *GetTrackListByMailFromAndTagNameRequest
	GetOffset() *string
	SetOffsetCreateTime(v string) *GetTrackListByMailFromAndTagNameRequest
	GetOffsetCreateTime() *string
	SetOffsetCreateTimeDesc(v string) *GetTrackListByMailFromAndTagNameRequest
	GetOffsetCreateTimeDesc() *string
	SetOwnerId(v int64) *GetTrackListByMailFromAndTagNameRequest
	GetOwnerId() *int64
	SetPageNumber(v string) *GetTrackListByMailFromAndTagNameRequest
	GetPageNumber() *string
	SetPageSize(v string) *GetTrackListByMailFromAndTagNameRequest
	GetPageSize() *string
	SetResourceOwnerAccount(v string) *GetTrackListByMailFromAndTagNameRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetTrackListByMailFromAndTagNameRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *GetTrackListByMailFromAndTagNameRequest
	GetStartTime() *string
	SetTagName(v string) *GetTrackListByMailFromAndTagNameRequest
	GetTagName() *string
	SetTotal(v string) *GetTrackListByMailFromAndTagNameRequest
	GetTotal() *string
}

type GetTrackListByMailFromAndTagNameRequest struct {
	// Sender address.
	//
	// > If not filled, it represents all addresses; if there is a TagName, this parameter must not be empty.
	//
	// example:
	//
	// e-service@amegroups.cn
	AccountName       *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DedicatedIp       *string `json:"DedicatedIp,omitempty" xml:"DedicatedIp,omitempty"`
	DedicatedIpPoolId *string `json:"DedicatedIpPoolId,omitempty" xml:"DedicatedIpPoolId,omitempty"`
	// End time, with a span from the start time that cannot exceed 15 days. Format: yyyy-MM-dd.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-29
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Esp     *string `json:"Esp,omitempty" xml:"Esp,omitempty"`
	// For the first query, set to 0; for subsequent queries, fixed at 1. 1 indicates pagination in ascending order by time. (This field is deprecated)
	//
	// example:
	//
	// （本字段已废弃）
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// Used for pagination. Not set for the first query; for subsequent queries, set to the value of OffsetCreateTime from the previous response. (This field is deprecated)
	//
	// example:
	//
	// （本字段已废弃）
	OffsetCreateTime *string `json:"OffsetCreateTime,omitempty" xml:"OffsetCreateTime,omitempty"`
	// (This field is deprecated)
	//
	// example:
	//
	// （本字段已废弃）
	OffsetCreateTimeDesc *string `json:"OffsetCreateTimeDesc,omitempty" xml:"OffsetCreateTimeDesc,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Current page number
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Start time, which cannot be earlier than 30 days. Format: yyyy-MM-dd.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-29
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Email tag. If not filled, it represents all tags.
	//
	// example:
	//
	// Subscription
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// (This field is deprecated)
	//
	// example:
	//
	// （本字段已废弃）
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetTrackListByMailFromAndTagNameRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTrackListByMailFromAndTagNameRequest) GoString() string {
	return s.String()
}

func (s *GetTrackListByMailFromAndTagNameRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GetTrackListByMailFromAndTagNameRequest) GetDedicatedIp() *string {
	return s.DedicatedIp
}

func (s *GetTrackListByMailFromAndTagNameRequest) GetDedicatedIpPoolId() *string {
	return s.DedicatedIpPoolId
}

func (s *GetTrackListByMailFromAndTagNameRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetTrackListByMailFromAndTagNameRequest) GetEsp() *string {
	return s.Esp
}

func (s *GetTrackListByMailFromAndTagNameRequest) GetOffset() *string {
	return s.Offset
}

func (s *GetTrackListByMailFromAndTagNameRequest) GetOffsetCreateTime() *string {
	return s.OffsetCreateTime
}

func (s *GetTrackListByMailFromAndTagNameRequest) GetOffsetCreateTimeDesc() *string {
	return s.OffsetCreateTimeDesc
}

func (s *GetTrackListByMailFromAndTagNameRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetTrackListByMailFromAndTagNameRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *GetTrackListByMailFromAndTagNameRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *GetTrackListByMailFromAndTagNameRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetTrackListByMailFromAndTagNameRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetTrackListByMailFromAndTagNameRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetTrackListByMailFromAndTagNameRequest) GetTagName() *string {
	return s.TagName
}

func (s *GetTrackListByMailFromAndTagNameRequest) GetTotal() *string {
	return s.Total
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetAccountName(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.AccountName = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetDedicatedIp(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.DedicatedIp = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetDedicatedIpPoolId(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.DedicatedIpPoolId = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetEndTime(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.EndTime = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetEsp(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.Esp = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetOffset(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.Offset = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetOffsetCreateTime(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.OffsetCreateTime = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetOffsetCreateTimeDesc(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.OffsetCreateTimeDesc = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetOwnerId(v int64) *GetTrackListByMailFromAndTagNameRequest {
	s.OwnerId = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetPageNumber(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.PageNumber = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetPageSize(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.PageSize = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetResourceOwnerAccount(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetResourceOwnerId(v int64) *GetTrackListByMailFromAndTagNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetStartTime(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.StartTime = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetTagName(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.TagName = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) SetTotal(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.Total = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameRequest) Validate() error {
	return dara.Validate(s)
}
