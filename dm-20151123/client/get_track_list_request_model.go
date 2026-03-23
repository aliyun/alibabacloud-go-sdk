// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrackListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GetTrackListRequest
	GetAccountName() *string
	SetConfigSetId(v string) *GetTrackListRequest
	GetConfigSetId() *string
	SetDedicatedIp(v string) *GetTrackListRequest
	GetDedicatedIp() *string
	SetDedicatedIpPoolId(v string) *GetTrackListRequest
	GetDedicatedIpPoolId() *string
	SetDomain(v string) *GetTrackListRequest
	GetDomain() *string
	SetEndTime(v string) *GetTrackListRequest
	GetEndTime() *string
	SetEsp(v string) *GetTrackListRequest
	GetEsp() *string
	SetOffset(v string) *GetTrackListRequest
	GetOffset() *string
	SetOffsetCreateTime(v string) *GetTrackListRequest
	GetOffsetCreateTime() *string
	SetOffsetCreateTimeDesc(v string) *GetTrackListRequest
	GetOffsetCreateTimeDesc() *string
	SetOwnerId(v int64) *GetTrackListRequest
	GetOwnerId() *int64
	SetPageNumber(v string) *GetTrackListRequest
	GetPageNumber() *string
	SetPageSize(v string) *GetTrackListRequest
	GetPageSize() *string
	SetResourceOwnerAccount(v string) *GetTrackListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetTrackListRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *GetTrackListRequest
	GetStartTime() *string
	SetTagName(v string) *GetTrackListRequest
	GetTagName() *string
	SetTotal(v string) *GetTrackListRequest
	GetTotal() *string
}

type GetTrackListRequest struct {
	// The sender address.
	//
	// > If you omit this parameter, the query returns data for all sender addresses. This parameter is required if you specify the `TagName` parameter.
	//
	// example:
	//
	// test@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the configuration set.
	//
	// example:
	//
	// xxx
	ConfigSetId *string `json:"ConfigSetId,omitempty" xml:"ConfigSetId,omitempty"`
	// The dedicated IP address to query.
	//
	// If this parameter is omitted, data for all dedicated IPs is returned.
	//
	// example:
	//
	// xxx.xxx.xxx.xxx
	DedicatedIp *string `json:"DedicatedIp,omitempty" xml:"DedicatedIp,omitempty"`
	// The ID of the dedicated IP pool to query.
	//
	// If this parameter is omitted, data for all IP pools is returned.
	//
	// example:
	//
	// xxx
	DedicatedIpPoolId *string `json:"DedicatedIpPoolId,omitempty" xml:"DedicatedIpPoolId,omitempty"`
	// example:
	//
	// dmdomain.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end date of the query. The duration between the StartTime and EndTime cannot exceed 7 days. The format is `yyyy-MM-dd`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-29
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The Email Service Provider (ESP) to query. Valid values are:
	//
	// - gmail.com
	//
	// - yahoo.com
	//
	// - outlook.com
	//
	// - icloud.com
	//
	// - Others: Any ESP not listed above.
	//
	// If you omit this parameter, the query returns data for all ESPs.
	//
	// example:
	//
	// gmail.com
	Esp *string `json:"Esp,omitempty" xml:"Esp,omitempty"`
	// Set this to 0 for the first query. For subsequent queries, set it to 1 to perform a paged query in chronological order. (This field is deprecated)
	//
	// example:
	//
	// （本字段已废弃）
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// Used for pagination. Do not set this parameter for the first query. For subsequent queries, set this parameter to the `OffsetCreateTime` value returned in the previous response. (This field is deprecated)
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
	// The page number to return.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The start date of the query. The date must be within the last 30 days. The format is `yyyy-MM-dd`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-09-29
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The tag name.
	//
	// example:
	//
	// tagname
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// (This field is deprecated)
	//
	// example:
	//
	// （本字段已废弃）
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetTrackListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTrackListRequest) GoString() string {
	return s.String()
}

func (s *GetTrackListRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GetTrackListRequest) GetConfigSetId() *string {
	return s.ConfigSetId
}

func (s *GetTrackListRequest) GetDedicatedIp() *string {
	return s.DedicatedIp
}

func (s *GetTrackListRequest) GetDedicatedIpPoolId() *string {
	return s.DedicatedIpPoolId
}

func (s *GetTrackListRequest) GetDomain() *string {
	return s.Domain
}

func (s *GetTrackListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetTrackListRequest) GetEsp() *string {
	return s.Esp
}

func (s *GetTrackListRequest) GetOffset() *string {
	return s.Offset
}

func (s *GetTrackListRequest) GetOffsetCreateTime() *string {
	return s.OffsetCreateTime
}

func (s *GetTrackListRequest) GetOffsetCreateTimeDesc() *string {
	return s.OffsetCreateTimeDesc
}

func (s *GetTrackListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetTrackListRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *GetTrackListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *GetTrackListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetTrackListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetTrackListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetTrackListRequest) GetTagName() *string {
	return s.TagName
}

func (s *GetTrackListRequest) GetTotal() *string {
	return s.Total
}

func (s *GetTrackListRequest) SetAccountName(v string) *GetTrackListRequest {
	s.AccountName = &v
	return s
}

func (s *GetTrackListRequest) SetConfigSetId(v string) *GetTrackListRequest {
	s.ConfigSetId = &v
	return s
}

func (s *GetTrackListRequest) SetDedicatedIp(v string) *GetTrackListRequest {
	s.DedicatedIp = &v
	return s
}

func (s *GetTrackListRequest) SetDedicatedIpPoolId(v string) *GetTrackListRequest {
	s.DedicatedIpPoolId = &v
	return s
}

func (s *GetTrackListRequest) SetDomain(v string) *GetTrackListRequest {
	s.Domain = &v
	return s
}

func (s *GetTrackListRequest) SetEndTime(v string) *GetTrackListRequest {
	s.EndTime = &v
	return s
}

func (s *GetTrackListRequest) SetEsp(v string) *GetTrackListRequest {
	s.Esp = &v
	return s
}

func (s *GetTrackListRequest) SetOffset(v string) *GetTrackListRequest {
	s.Offset = &v
	return s
}

func (s *GetTrackListRequest) SetOffsetCreateTime(v string) *GetTrackListRequest {
	s.OffsetCreateTime = &v
	return s
}

func (s *GetTrackListRequest) SetOffsetCreateTimeDesc(v string) *GetTrackListRequest {
	s.OffsetCreateTimeDesc = &v
	return s
}

func (s *GetTrackListRequest) SetOwnerId(v int64) *GetTrackListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetTrackListRequest) SetPageNumber(v string) *GetTrackListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetTrackListRequest) SetPageSize(v string) *GetTrackListRequest {
	s.PageSize = &v
	return s
}

func (s *GetTrackListRequest) SetResourceOwnerAccount(v string) *GetTrackListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetTrackListRequest) SetResourceOwnerId(v int64) *GetTrackListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetTrackListRequest) SetStartTime(v string) *GetTrackListRequest {
	s.StartTime = &v
	return s
}

func (s *GetTrackListRequest) SetTagName(v string) *GetTrackListRequest {
	s.TagName = &v
	return s
}

func (s *GetTrackListRequest) SetTotal(v string) *GetTrackListRequest {
	s.Total = &v
	return s
}

func (s *GetTrackListRequest) Validate() error {
	return dara.Validate(s)
}
