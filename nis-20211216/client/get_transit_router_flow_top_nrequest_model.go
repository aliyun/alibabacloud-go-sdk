// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTransitRouterFlowTopNRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*int64) *GetTransitRouterFlowTopNRequest
	GetAccountIds() []*int64
	SetBandwithPackageId(v string) *GetTransitRouterFlowTopNRequest
	GetBandwithPackageId() *string
	SetBeginTime(v int64) *GetTransitRouterFlowTopNRequest
	GetBeginTime() *int64
	SetCenId(v string) *GetTransitRouterFlowTopNRequest
	GetCenId() *string
	SetDirection(v string) *GetTransitRouterFlowTopNRequest
	GetDirection() *string
	SetEndTime(v int64) *GetTransitRouterFlowTopNRequest
	GetEndTime() *int64
	SetGroupBy(v string) *GetTransitRouterFlowTopNRequest
	GetGroupBy() *string
	SetOrderBy(v string) *GetTransitRouterFlowTopNRequest
	GetOrderBy() *string
	SetOtherIp(v string) *GetTransitRouterFlowTopNRequest
	GetOtherIp() *string
	SetOtherPort(v string) *GetTransitRouterFlowTopNRequest
	GetOtherPort() *string
	SetOtherRegion(v string) *GetTransitRouterFlowTopNRequest
	GetOtherRegion() *string
	SetProtocol(v string) *GetTransitRouterFlowTopNRequest
	GetProtocol() *string
	SetSort(v string) *GetTransitRouterFlowTopNRequest
	GetSort() *string
	SetThisIp(v string) *GetTransitRouterFlowTopNRequest
	GetThisIp() *string
	SetThisPort(v string) *GetTransitRouterFlowTopNRequest
	GetThisPort() *string
	SetThisRegion(v string) *GetTransitRouterFlowTopNRequest
	GetThisRegion() *string
	SetTopN(v int32) *GetTransitRouterFlowTopNRequest
	GetTopN() *int32
	SetUseMultiAccount(v bool) *GetTransitRouterFlowTopNRequest
	GetUseMultiAccount() *bool
}

type GetTransitRouterFlowTopNRequest struct {
	// The IDs of the member accounts.
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
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

func (s GetTransitRouterFlowTopNRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTransitRouterFlowTopNRequest) GoString() string {
	return s.String()
}

func (s *GetTransitRouterFlowTopNRequest) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *GetTransitRouterFlowTopNRequest) GetBandwithPackageId() *string {
	return s.BandwithPackageId
}

func (s *GetTransitRouterFlowTopNRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetTransitRouterFlowTopNRequest) GetCenId() *string {
	return s.CenId
}

func (s *GetTransitRouterFlowTopNRequest) GetDirection() *string {
	return s.Direction
}

func (s *GetTransitRouterFlowTopNRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetTransitRouterFlowTopNRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *GetTransitRouterFlowTopNRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetTransitRouterFlowTopNRequest) GetOtherIp() *string {
	return s.OtherIp
}

func (s *GetTransitRouterFlowTopNRequest) GetOtherPort() *string {
	return s.OtherPort
}

func (s *GetTransitRouterFlowTopNRequest) GetOtherRegion() *string {
	return s.OtherRegion
}

func (s *GetTransitRouterFlowTopNRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *GetTransitRouterFlowTopNRequest) GetSort() *string {
	return s.Sort
}

func (s *GetTransitRouterFlowTopNRequest) GetThisIp() *string {
	return s.ThisIp
}

func (s *GetTransitRouterFlowTopNRequest) GetThisPort() *string {
	return s.ThisPort
}

func (s *GetTransitRouterFlowTopNRequest) GetThisRegion() *string {
	return s.ThisRegion
}

func (s *GetTransitRouterFlowTopNRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *GetTransitRouterFlowTopNRequest) GetUseMultiAccount() *bool {
	return s.UseMultiAccount
}

func (s *GetTransitRouterFlowTopNRequest) SetAccountIds(v []*int64) *GetTransitRouterFlowTopNRequest {
	s.AccountIds = v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetBandwithPackageId(v string) *GetTransitRouterFlowTopNRequest {
	s.BandwithPackageId = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetBeginTime(v int64) *GetTransitRouterFlowTopNRequest {
	s.BeginTime = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetCenId(v string) *GetTransitRouterFlowTopNRequest {
	s.CenId = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetDirection(v string) *GetTransitRouterFlowTopNRequest {
	s.Direction = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetEndTime(v int64) *GetTransitRouterFlowTopNRequest {
	s.EndTime = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetGroupBy(v string) *GetTransitRouterFlowTopNRequest {
	s.GroupBy = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetOrderBy(v string) *GetTransitRouterFlowTopNRequest {
	s.OrderBy = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetOtherIp(v string) *GetTransitRouterFlowTopNRequest {
	s.OtherIp = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetOtherPort(v string) *GetTransitRouterFlowTopNRequest {
	s.OtherPort = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetOtherRegion(v string) *GetTransitRouterFlowTopNRequest {
	s.OtherRegion = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetProtocol(v string) *GetTransitRouterFlowTopNRequest {
	s.Protocol = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetSort(v string) *GetTransitRouterFlowTopNRequest {
	s.Sort = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetThisIp(v string) *GetTransitRouterFlowTopNRequest {
	s.ThisIp = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetThisPort(v string) *GetTransitRouterFlowTopNRequest {
	s.ThisPort = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetThisRegion(v string) *GetTransitRouterFlowTopNRequest {
	s.ThisRegion = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetTopN(v int32) *GetTransitRouterFlowTopNRequest {
	s.TopN = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetUseMultiAccount(v bool) *GetTransitRouterFlowTopNRequest {
	s.UseMultiAccount = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) Validate() error {
	return dara.Validate(s)
}
