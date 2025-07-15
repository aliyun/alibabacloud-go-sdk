// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnSslServerLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v int32) *DescribeVpnSslServerLogsRequest
	GetFrom() *int32
	SetMinutePeriod(v int32) *DescribeVpnSslServerLogsRequest
	GetMinutePeriod() *int32
	SetOwnerAccount(v string) *DescribeVpnSslServerLogsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVpnSslServerLogsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVpnSslServerLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpnSslServerLogsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeVpnSslServerLogsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVpnSslServerLogsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVpnSslServerLogsRequest
	GetResourceOwnerId() *int64
	SetSslVpnClientCertId(v string) *DescribeVpnSslServerLogsRequest
	GetSslVpnClientCertId() *string
	SetTo(v int32) *DescribeVpnSslServerLogsRequest
	GetTo() *int32
	SetVpnSslServerId(v string) *DescribeVpnSslServerLogsRequest
	GetVpnSslServerId() *string
}

type DescribeVpnSslServerLogsRequest struct {
	// The beginning of the time range to query. The value must be a unix timestamp. For example, 1600738962 specifies 09:42:42 (UTC+8) on September 22, 2020.
	//
	// >  If you specify **From**, you must also specify **To*	- or **MinutePeriod**.
	//
	// example:
	//
	// 1600738962
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// The interval at which log data is queried. Unit: minutes.
	//
	// >  If both **From*	- and **To*	- are not specified, you must specify **MinutePeriod**.
	//
	// example:
	//
	// 10
	MinutePeriod *int32  `json:"MinutePeriod,omitempty" xml:"MinutePeriod,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the SSL server is created. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SSL client certificate.
	//
	// example:
	//
	// vsc-m5euof6s5jy8vs5kd****
	SslVpnClientCertId *string `json:"SslVpnClientCertId,omitempty" xml:"SslVpnClientCertId,omitempty"`
	// The end of the time range to query. The value must be a unix timestamp. For example, 1600738962 specifies 09:42:42 (UTC+8) on September 22, 2020.
	//
	// >  If you specify **To**, you must also specify **From*	- or **MinutePeriod**.
	//
	// example:
	//
	// 1600738962
	To *int32 `json:"To,omitempty" xml:"To,omitempty"`
	// The ID of the SSL server.
	//
	// This parameter is required.
	//
	// example:
	//
	// vss-bp155e9yclsg1xgq4****
	VpnSslServerId *string `json:"VpnSslServerId,omitempty" xml:"VpnSslServerId,omitempty"`
}

func (s DescribeVpnSslServerLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnSslServerLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpnSslServerLogsRequest) GetFrom() *int32 {
	return s.From
}

func (s *DescribeVpnSslServerLogsRequest) GetMinutePeriod() *int32 {
	return s.MinutePeriod
}

func (s *DescribeVpnSslServerLogsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVpnSslServerLogsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpnSslServerLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpnSslServerLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpnSslServerLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpnSslServerLogsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVpnSslServerLogsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVpnSslServerLogsRequest) GetSslVpnClientCertId() *string {
	return s.SslVpnClientCertId
}

func (s *DescribeVpnSslServerLogsRequest) GetTo() *int32 {
	return s.To
}

func (s *DescribeVpnSslServerLogsRequest) GetVpnSslServerId() *string {
	return s.VpnSslServerId
}

func (s *DescribeVpnSslServerLogsRequest) SetFrom(v int32) *DescribeVpnSslServerLogsRequest {
	s.From = &v
	return s
}

func (s *DescribeVpnSslServerLogsRequest) SetMinutePeriod(v int32) *DescribeVpnSslServerLogsRequest {
	s.MinutePeriod = &v
	return s
}

func (s *DescribeVpnSslServerLogsRequest) SetOwnerAccount(v string) *DescribeVpnSslServerLogsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVpnSslServerLogsRequest) SetOwnerId(v int64) *DescribeVpnSslServerLogsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpnSslServerLogsRequest) SetPageNumber(v int32) *DescribeVpnSslServerLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpnSslServerLogsRequest) SetPageSize(v int32) *DescribeVpnSslServerLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpnSslServerLogsRequest) SetRegionId(v string) *DescribeVpnSslServerLogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVpnSslServerLogsRequest) SetResourceOwnerAccount(v string) *DescribeVpnSslServerLogsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVpnSslServerLogsRequest) SetResourceOwnerId(v int64) *DescribeVpnSslServerLogsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVpnSslServerLogsRequest) SetSslVpnClientCertId(v string) *DescribeVpnSslServerLogsRequest {
	s.SslVpnClientCertId = &v
	return s
}

func (s *DescribeVpnSslServerLogsRequest) SetTo(v int32) *DescribeVpnSslServerLogsRequest {
	s.To = &v
	return s
}

func (s *DescribeVpnSslServerLogsRequest) SetVpnSslServerId(v string) *DescribeVpnSslServerLogsRequest {
	s.VpnSslServerId = &v
	return s
}

func (s *DescribeVpnSslServerLogsRequest) Validate() error {
	return dara.Validate(s)
}
