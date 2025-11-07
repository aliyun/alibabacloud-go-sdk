// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTransitRouterFlowTopNShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIdsShrink(v string) *GetTransitRouterFlowTopNShrinkRequest
	GetAccountIdsShrink() *string
	SetBandwithPackageId(v string) *GetTransitRouterFlowTopNShrinkRequest
	GetBandwithPackageId() *string
	SetBeginTime(v int64) *GetTransitRouterFlowTopNShrinkRequest
	GetBeginTime() *int64
	SetCenId(v string) *GetTransitRouterFlowTopNShrinkRequest
	GetCenId() *string
	SetDirection(v string) *GetTransitRouterFlowTopNShrinkRequest
	GetDirection() *string
	SetEndTime(v int64) *GetTransitRouterFlowTopNShrinkRequest
	GetEndTime() *int64
	SetGroupBy(v string) *GetTransitRouterFlowTopNShrinkRequest
	GetGroupBy() *string
	SetOrderBy(v string) *GetTransitRouterFlowTopNShrinkRequest
	GetOrderBy() *string
	SetOtherIp(v string) *GetTransitRouterFlowTopNShrinkRequest
	GetOtherIp() *string
	SetOtherPort(v string) *GetTransitRouterFlowTopNShrinkRequest
	GetOtherPort() *string
	SetOtherRegion(v string) *GetTransitRouterFlowTopNShrinkRequest
	GetOtherRegion() *string
	SetProtocol(v string) *GetTransitRouterFlowTopNShrinkRequest
	GetProtocol() *string
	SetSort(v string) *GetTransitRouterFlowTopNShrinkRequest
	GetSort() *string
	SetThisIp(v string) *GetTransitRouterFlowTopNShrinkRequest
	GetThisIp() *string
	SetThisPort(v string) *GetTransitRouterFlowTopNShrinkRequest
	GetThisPort() *string
	SetThisRegion(v string) *GetTransitRouterFlowTopNShrinkRequest
	GetThisRegion() *string
	SetTopN(v int32) *GetTransitRouterFlowTopNShrinkRequest
	GetTopN() *int32
	SetUseMultiAccount(v bool) *GetTransitRouterFlowTopNShrinkRequest
	GetUseMultiAccount() *bool
}

type GetTransitRouterFlowTopNShrinkRequest struct {
	// The IDs of the member accounts.
	AccountIdsShrink *string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	// The ID of the CEN bandwidth plan.
	//
	// example:
	//
	// cenbwp-ia8kw1zjv4hyal*****
	BandwithPackageId *string `json:"BandwithPackageId,omitempty" xml:"BandwithPackageId,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. The maximum time range that you can query is 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1684373600099
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The CEN instance ID.
	//
	// example:
	//
	// cen-ia8kw1zjv4hyal****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The direction of the inter-region traffic in the local regions or for the local IP addresses. Valid values:
	//
	// 	- **in**: inbound traffic
	//
	// 	- **out**: outbound traffic
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
	// The dimension for ranking inter-region traffic data. The value of this parameter is case-sensitive. Valid values:
	//
	// 	- **1Tuple**: queries the rankings of inter-region traffic data for the local regions, Cloud Enterprise Network (CEN) instances, and IP addresses.
	//
	// 	- **2Tuple**: queries the rankings of inter-region traffic data for the local and remote regions, and the local and remote IP addresses.
	//
	// 	- **5Tuple**: queries the rankings of inter-region traffic data for the local and remote IP addresses, local and remote ports, and protocols.
	//
	// 	- **Cen**: queries the rankings of inter-region traffic data for CEN instances.
	//
	// 	- **RegionPair**: queries the rankings of inter-region traffic data for the local and remote regions.
	//
	// 	- **Port**: queries the rankings of inter-region traffic data for the local and remote ports.
	//
	// 	- **Protocol**: queries the rankings of inter-region traffic data for the protocols.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1Tuple
	GroupBy *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	// The metric for ranking inter-region traffic data. Default value: Bytes. This value specifies that inter-region traffic data is ranked by traffic volume.
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
	// example:
	//
	// 10869
	OtherPort *string `json:"OtherPort,omitempty" xml:"OtherPort,omitempty"`
	// The remote region.
	//
	// example:
	//
	// ap-southeast-1
	OtherRegion *string `json:"OtherRegion,omitempty" xml:"OtherRegion,omitempty"`
	// The protocol number.
	//
	// >  All protocols are supported. This parameter is required only if you set **GroupBy*	- to **5Tuple*	- or **Protocol**.
	//
	// example:
	//
	// 6
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The order for ranking inter-region traffic data. Valid values:
	//
	// 	- **desc**: descending order
	//
	// 	- **asc**: ascending order
	//
	// example:
	//
	// desc
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The local IP address.
	//
	// example:
	//
	// 1.8.XX.XX
	ThisIp *string `json:"ThisIp,omitempty" xml:"ThisIp,omitempty"`
	// The local port.
	//
	// example:
	//
	// 80
	ThisPort *string `json:"ThisPort,omitempty" xml:"ThisPort,omitempty"`
	// The local region where the **local IP address*	- resides.
	//
	// example:
	//
	// cn-shanghai
	ThisRegion *string `json:"ThisRegion,omitempty" xml:"ThisRegion,omitempty"`
	// Specifies the maximum number of data entries to display. Default value: **10**. Maximum value: 100.
	//
	// example:
	//
	// 20
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
	// Specifies whether to enable the multi-account management feature. Default value: **false**. This value specifies that the multi-account management feature is disabled.
	//
	// >  By default, the multi-account management feature is not available. If you want to use this feature, contact your account manager to apply for permissions.
	//
	// example:
	//
	// false
	UseMultiAccount *bool `json:"UseMultiAccount,omitempty" xml:"UseMultiAccount,omitempty"`
}

func (s GetTransitRouterFlowTopNShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTransitRouterFlowTopNShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetTransitRouterFlowTopNShrinkRequest) GetAccountIdsShrink() *string {
	return s.AccountIdsShrink
}

func (s *GetTransitRouterFlowTopNShrinkRequest) GetBandwithPackageId() *string {
	return s.BandwithPackageId
}

func (s *GetTransitRouterFlowTopNShrinkRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetTransitRouterFlowTopNShrinkRequest) GetCenId() *string {
	return s.CenId
}

func (s *GetTransitRouterFlowTopNShrinkRequest) GetDirection() *string {
	return s.Direction
}

func (s *GetTransitRouterFlowTopNShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetTransitRouterFlowTopNShrinkRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *GetTransitRouterFlowTopNShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetTransitRouterFlowTopNShrinkRequest) GetOtherIp() *string {
	return s.OtherIp
}

func (s *GetTransitRouterFlowTopNShrinkRequest) GetOtherPort() *string {
	return s.OtherPort
}

func (s *GetTransitRouterFlowTopNShrinkRequest) GetOtherRegion() *string {
	return s.OtherRegion
}

func (s *GetTransitRouterFlowTopNShrinkRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *GetTransitRouterFlowTopNShrinkRequest) GetSort() *string {
	return s.Sort
}

func (s *GetTransitRouterFlowTopNShrinkRequest) GetThisIp() *string {
	return s.ThisIp
}

func (s *GetTransitRouterFlowTopNShrinkRequest) GetThisPort() *string {
	return s.ThisPort
}

func (s *GetTransitRouterFlowTopNShrinkRequest) GetThisRegion() *string {
	return s.ThisRegion
}

func (s *GetTransitRouterFlowTopNShrinkRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *GetTransitRouterFlowTopNShrinkRequest) GetUseMultiAccount() *bool {
	return s.UseMultiAccount
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetAccountIdsShrink(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.AccountIdsShrink = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetBandwithPackageId(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.BandwithPackageId = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetBeginTime(v int64) *GetTransitRouterFlowTopNShrinkRequest {
	s.BeginTime = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetCenId(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.CenId = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetDirection(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.Direction = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetEndTime(v int64) *GetTransitRouterFlowTopNShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetGroupBy(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.GroupBy = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetOrderBy(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetOtherIp(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.OtherIp = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetOtherPort(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.OtherPort = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetOtherRegion(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.OtherRegion = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetProtocol(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetSort(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.Sort = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetThisIp(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.ThisIp = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetThisPort(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.ThisPort = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetThisRegion(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.ThisRegion = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetTopN(v int32) *GetTransitRouterFlowTopNShrinkRequest {
	s.TopN = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetUseMultiAccount(v bool) *GetTransitRouterFlowTopNShrinkRequest {
	s.UseMultiAccount = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) Validate() error {
	return dara.Validate(s)
}
