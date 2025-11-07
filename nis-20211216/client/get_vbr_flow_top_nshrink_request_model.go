// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVbrFlowTopNShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIdsShrink(v string) *GetVbrFlowTopNShrinkRequest
	GetAccountIdsShrink() *string
	SetAttachmentId(v string) *GetVbrFlowTopNShrinkRequest
	GetAttachmentId() *string
	SetBeginTime(v int64) *GetVbrFlowTopNShrinkRequest
	GetBeginTime() *int64
	SetCenId(v string) *GetVbrFlowTopNShrinkRequest
	GetCenId() *string
	SetCloudIp(v string) *GetVbrFlowTopNShrinkRequest
	GetCloudIp() *string
	SetCloudPort(v string) *GetVbrFlowTopNShrinkRequest
	GetCloudPort() *string
	SetDirection(v string) *GetVbrFlowTopNShrinkRequest
	GetDirection() *string
	SetEndTime(v int64) *GetVbrFlowTopNShrinkRequest
	GetEndTime() *int64
	SetGroupBy(v string) *GetVbrFlowTopNShrinkRequest
	GetGroupBy() *string
	SetOrderBy(v string) *GetVbrFlowTopNShrinkRequest
	GetOrderBy() *string
	SetOtherIp(v string) *GetVbrFlowTopNShrinkRequest
	GetOtherIp() *string
	SetOtherPort(v string) *GetVbrFlowTopNShrinkRequest
	GetOtherPort() *string
	SetProtocol(v string) *GetVbrFlowTopNShrinkRequest
	GetProtocol() *string
	SetRegionId(v string) *GetVbrFlowTopNShrinkRequest
	GetRegionId() *string
	SetSort(v string) *GetVbrFlowTopNShrinkRequest
	GetSort() *string
	SetTopN(v int32) *GetVbrFlowTopNShrinkRequest
	GetTopN() *int32
	SetUseMultiAccount(v bool) *GetVbrFlowTopNShrinkRequest
	GetUseMultiAccount() *bool
	SetVirtualBorderRouterId(v string) *GetVbrFlowTopNShrinkRequest
	GetVirtualBorderRouterId() *string
}

type GetVbrFlowTopNShrinkRequest struct {
	// The IDs of member accounts.
	AccountIdsShrink *string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	// The CEN connection ID.
	//
	// example:
	//
	// tr-attach-dnv870gmqzmb5u****
	AttachmentId *string `json:"AttachmentId,omitempty" xml:"AttachmentId,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. The maximum time range that you can query is 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1638239092000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The CEN instance ID.
	//
	// example:
	//
	// cen-ia8kw1zjv4hyal****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The local IP address.
	//
	// example:
	//
	// 112.74.XX.XX
	CloudIp *string `json:"CloudIp,omitempty" xml:"CloudIp,omitempty"`
	// The local port.
	//
	// >  This parameter is required only if you set **GroupBy*	- to **CloudPort**.
	//
	// example:
	//
	// 443
	CloudPort *string `json:"CloudPort,omitempty" xml:"CloudPort,omitempty"`
	// The direction of the hybrid cloud traffic in the local regions or for the local IP addresses. Valid values:
	//
	// 	- **in**: traffic from a data center to Alibaba Cloud
	//
	// 	- **out**: traffic from Alibaba Cloud to a data center
	//
	// This parameter is required.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. The maximum time range that you can query is 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1638239093000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The dimension for ranking hybrid cloud traffic data. The value of this parameter is case-sensitive. Valid values:
	//
	// 	- **1Tuple**: queries the rankings of hybrid cloud traffic data for the Cloud Enterprise Network (CEN) instances, CEN connections, virtual border routers (VBRs), and IP addresses.
	//
	// 	- **2Tuple**: queries the rankings of hybrid cloud traffic data for the local and remote IP addresses.
	//
	// 	- **5Tuple**: queries the rankings of hybrid cloud traffic data for the local and remote IP addresses, local and remote ports, and protocols.
	//
	// 	- **CloudPort**: queries the rankings of hybrid cloud traffic data for the local ports.
	//
	// 	- **OtherPort**: queries the rankings of hybrid cloud traffic data for the remote ports.
	//
	// 	- **Protocol**: queries the rankings of hybrid cloud traffic data for the protocols.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1Tuple
	GroupBy *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	// The metric for ranking hybrid cloud traffic data. Default value: Bytes. This value specifies that hybrid cloud traffic data is ranked by traffic volumes.
	//
	// example:
	//
	// Bytes
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The remote IP address.
	//
	// example:
	//
	// 122.112.XX.XX
	OtherIp *string `json:"OtherIp,omitempty" xml:"OtherIp,omitempty"`
	// The remote port.
	//
	// >  This parameter is required only if you set **GroupBy*	- to **OtherPort**.
	//
	// example:
	//
	// 40002
	OtherPort *string `json:"OtherPort,omitempty" xml:"OtherPort,omitempty"`
	// The protocol number.
	//
	// >  All protocols are supported. This parameter is required only if you set **GroupBy*	- to **5Tuple*	- or **Protocol**.
	//
	// example:
	//
	// 6
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The local region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The order for ranking hybrid cloud traffic data. Valid values:
	//
	// 	- **desc**: descending order
	//
	// 	- **asc**: ascending order
	//
	// example:
	//
	// desc
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// Specifies top-N traffic data to display. Default value: **10**. This value specifies that top-10 traffic data is displayed by default. Maximum value: **100**.
	//
	// example:
	//
	// 10
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
	// Specifies whether to enable the multi-account management feature. Default value: **false**. This value specifies that the multi-account management feature is disabled.
	//
	// >  By default, the multi-account management feature is not available. If you want to use this feature, contact your account manager to apply for permissions.
	//
	// example:
	//
	// false
	UseMultiAccount *bool `json:"UseMultiAccount,omitempty" xml:"UseMultiAccount,omitempty"`
	// The ID of the VBR that is associated with the Express Connect circuit.
	//
	// example:
	//
	// vbr-k1atj46citwuek42j****
	VirtualBorderRouterId *string `json:"VirtualBorderRouterId,omitempty" xml:"VirtualBorderRouterId,omitempty"`
}

func (s GetVbrFlowTopNShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVbrFlowTopNShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetVbrFlowTopNShrinkRequest) GetAccountIdsShrink() *string {
	return s.AccountIdsShrink
}

func (s *GetVbrFlowTopNShrinkRequest) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *GetVbrFlowTopNShrinkRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetVbrFlowTopNShrinkRequest) GetCenId() *string {
	return s.CenId
}

func (s *GetVbrFlowTopNShrinkRequest) GetCloudIp() *string {
	return s.CloudIp
}

func (s *GetVbrFlowTopNShrinkRequest) GetCloudPort() *string {
	return s.CloudPort
}

func (s *GetVbrFlowTopNShrinkRequest) GetDirection() *string {
	return s.Direction
}

func (s *GetVbrFlowTopNShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetVbrFlowTopNShrinkRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *GetVbrFlowTopNShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetVbrFlowTopNShrinkRequest) GetOtherIp() *string {
	return s.OtherIp
}

func (s *GetVbrFlowTopNShrinkRequest) GetOtherPort() *string {
	return s.OtherPort
}

func (s *GetVbrFlowTopNShrinkRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *GetVbrFlowTopNShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVbrFlowTopNShrinkRequest) GetSort() *string {
	return s.Sort
}

func (s *GetVbrFlowTopNShrinkRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *GetVbrFlowTopNShrinkRequest) GetUseMultiAccount() *bool {
	return s.UseMultiAccount
}

func (s *GetVbrFlowTopNShrinkRequest) GetVirtualBorderRouterId() *string {
	return s.VirtualBorderRouterId
}

func (s *GetVbrFlowTopNShrinkRequest) SetAccountIdsShrink(v string) *GetVbrFlowTopNShrinkRequest {
	s.AccountIdsShrink = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetAttachmentId(v string) *GetVbrFlowTopNShrinkRequest {
	s.AttachmentId = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetBeginTime(v int64) *GetVbrFlowTopNShrinkRequest {
	s.BeginTime = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetCenId(v string) *GetVbrFlowTopNShrinkRequest {
	s.CenId = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetCloudIp(v string) *GetVbrFlowTopNShrinkRequest {
	s.CloudIp = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetCloudPort(v string) *GetVbrFlowTopNShrinkRequest {
	s.CloudPort = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetDirection(v string) *GetVbrFlowTopNShrinkRequest {
	s.Direction = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetEndTime(v int64) *GetVbrFlowTopNShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetGroupBy(v string) *GetVbrFlowTopNShrinkRequest {
	s.GroupBy = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetOrderBy(v string) *GetVbrFlowTopNShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetOtherIp(v string) *GetVbrFlowTopNShrinkRequest {
	s.OtherIp = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetOtherPort(v string) *GetVbrFlowTopNShrinkRequest {
	s.OtherPort = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetProtocol(v string) *GetVbrFlowTopNShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetRegionId(v string) *GetVbrFlowTopNShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetSort(v string) *GetVbrFlowTopNShrinkRequest {
	s.Sort = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetTopN(v int32) *GetVbrFlowTopNShrinkRequest {
	s.TopN = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetUseMultiAccount(v bool) *GetVbrFlowTopNShrinkRequest {
	s.UseMultiAccount = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetVirtualBorderRouterId(v string) *GetVbrFlowTopNShrinkRequest {
	s.VirtualBorderRouterId = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) Validate() error {
	return dara.Validate(s)
}
