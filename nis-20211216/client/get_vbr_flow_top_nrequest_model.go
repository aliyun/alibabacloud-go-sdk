// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVbrFlowTopNRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*int64) *GetVbrFlowTopNRequest
	GetAccountIds() []*int64
	SetAttachmentId(v string) *GetVbrFlowTopNRequest
	GetAttachmentId() *string
	SetBeginTime(v int64) *GetVbrFlowTopNRequest
	GetBeginTime() *int64
	SetCenId(v string) *GetVbrFlowTopNRequest
	GetCenId() *string
	SetCloudIp(v string) *GetVbrFlowTopNRequest
	GetCloudIp() *string
	SetCloudPort(v string) *GetVbrFlowTopNRequest
	GetCloudPort() *string
	SetDirection(v string) *GetVbrFlowTopNRequest
	GetDirection() *string
	SetEndTime(v int64) *GetVbrFlowTopNRequest
	GetEndTime() *int64
	SetGroupBy(v string) *GetVbrFlowTopNRequest
	GetGroupBy() *string
	SetOrderBy(v string) *GetVbrFlowTopNRequest
	GetOrderBy() *string
	SetOtherIp(v string) *GetVbrFlowTopNRequest
	GetOtherIp() *string
	SetOtherPort(v string) *GetVbrFlowTopNRequest
	GetOtherPort() *string
	SetProtocol(v string) *GetVbrFlowTopNRequest
	GetProtocol() *string
	SetRegionId(v string) *GetVbrFlowTopNRequest
	GetRegionId() *string
	SetSort(v string) *GetVbrFlowTopNRequest
	GetSort() *string
	SetTopN(v int32) *GetVbrFlowTopNRequest
	GetTopN() *int32
	SetUseMultiAccount(v bool) *GetVbrFlowTopNRequest
	GetUseMultiAccount() *bool
	SetVirtualBorderRouterId(v string) *GetVbrFlowTopNRequest
	GetVirtualBorderRouterId() *string
}

type GetVbrFlowTopNRequest struct {
	// The IDs of member accounts.
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
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

func (s GetVbrFlowTopNRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVbrFlowTopNRequest) GoString() string {
	return s.String()
}

func (s *GetVbrFlowTopNRequest) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *GetVbrFlowTopNRequest) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *GetVbrFlowTopNRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetVbrFlowTopNRequest) GetCenId() *string {
	return s.CenId
}

func (s *GetVbrFlowTopNRequest) GetCloudIp() *string {
	return s.CloudIp
}

func (s *GetVbrFlowTopNRequest) GetCloudPort() *string {
	return s.CloudPort
}

func (s *GetVbrFlowTopNRequest) GetDirection() *string {
	return s.Direction
}

func (s *GetVbrFlowTopNRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetVbrFlowTopNRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *GetVbrFlowTopNRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetVbrFlowTopNRequest) GetOtherIp() *string {
	return s.OtherIp
}

func (s *GetVbrFlowTopNRequest) GetOtherPort() *string {
	return s.OtherPort
}

func (s *GetVbrFlowTopNRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *GetVbrFlowTopNRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVbrFlowTopNRequest) GetSort() *string {
	return s.Sort
}

func (s *GetVbrFlowTopNRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *GetVbrFlowTopNRequest) GetUseMultiAccount() *bool {
	return s.UseMultiAccount
}

func (s *GetVbrFlowTopNRequest) GetVirtualBorderRouterId() *string {
	return s.VirtualBorderRouterId
}

func (s *GetVbrFlowTopNRequest) SetAccountIds(v []*int64) *GetVbrFlowTopNRequest {
	s.AccountIds = v
	return s
}

func (s *GetVbrFlowTopNRequest) SetAttachmentId(v string) *GetVbrFlowTopNRequest {
	s.AttachmentId = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetBeginTime(v int64) *GetVbrFlowTopNRequest {
	s.BeginTime = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetCenId(v string) *GetVbrFlowTopNRequest {
	s.CenId = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetCloudIp(v string) *GetVbrFlowTopNRequest {
	s.CloudIp = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetCloudPort(v string) *GetVbrFlowTopNRequest {
	s.CloudPort = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetDirection(v string) *GetVbrFlowTopNRequest {
	s.Direction = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetEndTime(v int64) *GetVbrFlowTopNRequest {
	s.EndTime = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetGroupBy(v string) *GetVbrFlowTopNRequest {
	s.GroupBy = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetOrderBy(v string) *GetVbrFlowTopNRequest {
	s.OrderBy = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetOtherIp(v string) *GetVbrFlowTopNRequest {
	s.OtherIp = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetOtherPort(v string) *GetVbrFlowTopNRequest {
	s.OtherPort = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetProtocol(v string) *GetVbrFlowTopNRequest {
	s.Protocol = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetRegionId(v string) *GetVbrFlowTopNRequest {
	s.RegionId = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetSort(v string) *GetVbrFlowTopNRequest {
	s.Sort = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetTopN(v int32) *GetVbrFlowTopNRequest {
	s.TopN = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetUseMultiAccount(v bool) *GetVbrFlowTopNRequest {
	s.UseMultiAccount = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetVirtualBorderRouterId(v string) *GetVbrFlowTopNRequest {
	s.VirtualBorderRouterId = &v
	return s
}

func (s *GetVbrFlowTopNRequest) Validate() error {
	return dara.Validate(s)
}
