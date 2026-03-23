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
	SetConfigSetId(v string) *GetTrackListByMailFromAndTagNameRequest
	GetConfigSetId() *string
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
	// The sender address.
	//
	// > If you leave this parameter empty, data for all addresses is returned. This parameter is required if you specify TagName.
	//
	// example:
	//
	// e-service@amegroups.cn
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The configuration set ID.
	//
	// example:
	//
	// xxx
	ConfigSetId *string `json:"ConfigSetId,omitempty" xml:"ConfigSetId,omitempty"`
	// The dedicated IP address. This parameter is available only for users of dedicated IPs.
	//
	// If you do not specify this parameter, data for all IPs is returned.
	//
	// example:
	//
	// xxx.xxx.xxx.xxx
	DedicatedIp *string `json:"DedicatedIp,omitempty" xml:"DedicatedIp,omitempty"`
	// The ID of the dedicated IP pool. This parameter is available only for users of dedicated IPs.
	//
	// If you do not specify this parameter, data for all IP pools is returned.
	//
	// example:
	//
	// xxx
	DedicatedIpPoolId *string `json:"DedicatedIpPoolId,omitempty" xml:"DedicatedIpPoolId,omitempty"`
	// The end time. The time span between the start time and end time cannot exceed 15 days. The format is yyyy-MM-dd.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-29
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The Email Service Provider (ESP). This parameter is available only for users of dedicated IPs. Valid values:
	//
	// - gmail.com
	//
	// - yahoo.com
	//
	// - outlook.com
	//
	// - icloud.com
	//
	// - others (data for ESPs other than the ones listed above)
	//
	// If you do not specify this parameter, data for all ESPs is returned.
	//
	// example:
	//
	// gmail.com
	Esp *string `json:"Esp,omitempty" xml:"Esp,omitempty"`
	// The value is 0 for the first query and 1 for subsequent queries. A value of 1 indicates a paged query in chronological order. (This field is deprecated)
	//
	// example:
	//
	// （本字段已废弃）
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// Used for paging. Do not set this parameter for the first query. For subsequent queries, set this parameter to the OffsetCreateTime value from the previous response. (This field is deprecated)
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
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The start time. The time cannot be earlier than 30 days ago. The format is yyyy-MM-dd.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-29
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The email tag. If you leave this parameter empty, data for all tags is returned.
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

func (s *GetTrackListByMailFromAndTagNameRequest) GetConfigSetId() *string {
	return s.ConfigSetId
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

func (s *GetTrackListByMailFromAndTagNameRequest) SetConfigSetId(v string) *GetTrackListByMailFromAndTagNameRequest {
	s.ConfigSetId = &v
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
