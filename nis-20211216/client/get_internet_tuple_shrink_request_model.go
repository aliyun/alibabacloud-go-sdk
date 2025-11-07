// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInternetTupleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*int64) *GetInternetTupleShrinkRequest
	GetAccountIds() []*int64
	SetBeginTime(v int64) *GetInternetTupleShrinkRequest
	GetBeginTime() *int64
	SetCloudIp(v string) *GetInternetTupleShrinkRequest
	GetCloudIp() *string
	SetCloudIpListShrink(v string) *GetInternetTupleShrinkRequest
	GetCloudIpListShrink() *string
	SetCloudIsp(v string) *GetInternetTupleShrinkRequest
	GetCloudIsp() *string
	SetCloudPort(v string) *GetInternetTupleShrinkRequest
	GetCloudPort() *string
	SetDirection(v string) *GetInternetTupleShrinkRequest
	GetDirection() *string
	SetEndTime(v int64) *GetInternetTupleShrinkRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetInternetTupleShrinkRequest
	GetInstanceId() *string
	SetInstanceListShrink(v string) *GetInternetTupleShrinkRequest
	GetInstanceListShrink() *string
	SetOrderBy(v string) *GetInternetTupleShrinkRequest
	GetOrderBy() *string
	SetOtherCity(v string) *GetInternetTupleShrinkRequest
	GetOtherCity() *string
	SetOtherCountry(v string) *GetInternetTupleShrinkRequest
	GetOtherCountry() *string
	SetOtherIp(v string) *GetInternetTupleShrinkRequest
	GetOtherIp() *string
	SetOtherIsp(v string) *GetInternetTupleShrinkRequest
	GetOtherIsp() *string
	SetOtherPort(v string) *GetInternetTupleShrinkRequest
	GetOtherPort() *string
	SetProtocol(v string) *GetInternetTupleShrinkRequest
	GetProtocol() *string
	SetRegionId(v string) *GetInternetTupleShrinkRequest
	GetRegionId() *string
	SetSort(v string) *GetInternetTupleShrinkRequest
	GetSort() *string
	SetTopN(v int32) *GetInternetTupleShrinkRequest
	GetTopN() *int32
	SetTupleType(v int32) *GetInternetTupleShrinkRequest
	GetTupleType() *int32
	SetUseMultiAccount(v bool) *GetInternetTupleShrinkRequest
	GetUseMultiAccount() *bool
}

type GetInternetTupleShrinkRequest struct {
	// The IDs of member accounts.
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The beginning of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1684373600099
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The local IP address.
	//
	// example:
	//
	// 112.74.XX.XX
	CloudIp *string `json:"CloudIp,omitempty" xml:"CloudIp,omitempty"`
	// The local IP addresses for filtering.
	CloudIpListShrink *string `json:"CloudIpList,omitempty" xml:"CloudIpList,omitempty"`
	// The local Internet service provider (ISP).
	//
	// >  In most cases, the value is Alibaba or Alibaba Cloud.
	//
	// example:
	//
	// Alibaba
	CloudIsp *string `json:"CloudIsp,omitempty" xml:"CloudIsp,omitempty"`
	// The local port.
	//
	// >  This parameter is required only if you set GroupBy to CloudPort.
	//
	// example:
	//
	// 443
	CloudPort *string `json:"CloudPort,omitempty" xml:"CloudPort,omitempty"`
	// The direction of the Internet traffic that you want to query. Valid values:
	//
	// 	- **in**: inbound
	//
	// 	- **out**: outbound
	//
	// This parameter is required.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1684373700099
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the Alibaba Cloud instance.
	//
	// example:
	//
	// eip-sample*
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance IDs for filtering.
	InstanceListShrink *string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty"`
	// The metric for data ranking. Default value: **ByteCount**. This value indicates that Internet traffic data is ranked by traffic volume.
	//
	// Valid values:
	//
	// 	- Rtt
	//
	// 	- ByteCount
	//
	// 	- PacketCount
	//
	// 	- RetransmitRate
	//
	// example:
	//
	// ByteCount
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The remote city.
	//
	// >  This parameter is required only if you set **TupleType*	- to **2*	- or **5**.
	//
	// example:
	//
	// Hangzhou
	OtherCity *string `json:"OtherCity,omitempty" xml:"OtherCity,omitempty"`
	// The remote country.
	//
	// >  This parameter is required only if you set **TupleType*	- to **2*	- or **5**.
	//
	// example:
	//
	// China
	OtherCountry *string `json:"OtherCountry,omitempty" xml:"OtherCountry,omitempty"`
	// The remote IP address.
	//
	// > This parameter is required only when you set **TupleType** to **2** or **5**.
	//
	// example:
	//
	// 122.112.XX.XX
	OtherIp *string `json:"OtherIp,omitempty" xml:"OtherIp,omitempty"`
	// The remote ISP.
	//
	// > This parameter is required if you want to view the information about the remote ISP.
	//
	// example:
	//
	// China Mobile
	OtherIsp *string `json:"OtherIsp,omitempty" xml:"OtherIsp,omitempty"`
	// The remote port.
	//
	// > This parameter is required only when you set **TupleType*	- to **5**.
	//
	// example:
	//
	// 40002
	OtherPort *string `json:"OtherPort,omitempty" xml:"OtherPort,omitempty"`
	// The protocol number.
	//
	// > All protocols are supported. This parameter is required only when you set **TupleType** to **5**.
	//
	// example:
	//
	// 6
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the region for which you want to query the Internet traffic.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The order in which instances are ranked by Internet traffic. Valid values:
	//
	// 	- **desc**: the descending order
	//
	// 	- **asc**: the ascending order
	//
	// example:
	//
	// desc
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// Specifies top-N traffic data to display. Default value: **10**. This value specifies to display top-10 traffic data by default. Max value: **100**.
	//
	// example:
	//
	// 10
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
	// The type of the tuple. Valid values:
	//
	// 	- **1**: 1-tuple
	//
	// 	- **2**: 2-tuple
	//
	// 	- **5**: 5-tuple
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TupleType *int32 `json:"TupleType,omitempty" xml:"TupleType,omitempty"`
	// Specifies whether to enable the multi-account management feature. Default value: **false**. This value specifies that the multi-account management feature is disabled.
	//
	// >  By default, the multi-account management feature is not available. If you want to use this feature, contact your account manager to apply for permissions.
	//
	// example:
	//
	// false
	UseMultiAccount *bool `json:"UseMultiAccount,omitempty" xml:"UseMultiAccount,omitempty"`
}

func (s GetInternetTupleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInternetTupleShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetInternetTupleShrinkRequest) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *GetInternetTupleShrinkRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetInternetTupleShrinkRequest) GetCloudIp() *string {
	return s.CloudIp
}

func (s *GetInternetTupleShrinkRequest) GetCloudIpListShrink() *string {
	return s.CloudIpListShrink
}

func (s *GetInternetTupleShrinkRequest) GetCloudIsp() *string {
	return s.CloudIsp
}

func (s *GetInternetTupleShrinkRequest) GetCloudPort() *string {
	return s.CloudPort
}

func (s *GetInternetTupleShrinkRequest) GetDirection() *string {
	return s.Direction
}

func (s *GetInternetTupleShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetInternetTupleShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInternetTupleShrinkRequest) GetInstanceListShrink() *string {
	return s.InstanceListShrink
}

func (s *GetInternetTupleShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetInternetTupleShrinkRequest) GetOtherCity() *string {
	return s.OtherCity
}

func (s *GetInternetTupleShrinkRequest) GetOtherCountry() *string {
	return s.OtherCountry
}

func (s *GetInternetTupleShrinkRequest) GetOtherIp() *string {
	return s.OtherIp
}

func (s *GetInternetTupleShrinkRequest) GetOtherIsp() *string {
	return s.OtherIsp
}

func (s *GetInternetTupleShrinkRequest) GetOtherPort() *string {
	return s.OtherPort
}

func (s *GetInternetTupleShrinkRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *GetInternetTupleShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInternetTupleShrinkRequest) GetSort() *string {
	return s.Sort
}

func (s *GetInternetTupleShrinkRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *GetInternetTupleShrinkRequest) GetTupleType() *int32 {
	return s.TupleType
}

func (s *GetInternetTupleShrinkRequest) GetUseMultiAccount() *bool {
	return s.UseMultiAccount
}

func (s *GetInternetTupleShrinkRequest) SetAccountIds(v []*int64) *GetInternetTupleShrinkRequest {
	s.AccountIds = v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetBeginTime(v int64) *GetInternetTupleShrinkRequest {
	s.BeginTime = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetCloudIp(v string) *GetInternetTupleShrinkRequest {
	s.CloudIp = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetCloudIpListShrink(v string) *GetInternetTupleShrinkRequest {
	s.CloudIpListShrink = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetCloudIsp(v string) *GetInternetTupleShrinkRequest {
	s.CloudIsp = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetCloudPort(v string) *GetInternetTupleShrinkRequest {
	s.CloudPort = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetDirection(v string) *GetInternetTupleShrinkRequest {
	s.Direction = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetEndTime(v int64) *GetInternetTupleShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetInstanceId(v string) *GetInternetTupleShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetInstanceListShrink(v string) *GetInternetTupleShrinkRequest {
	s.InstanceListShrink = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetOrderBy(v string) *GetInternetTupleShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetOtherCity(v string) *GetInternetTupleShrinkRequest {
	s.OtherCity = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetOtherCountry(v string) *GetInternetTupleShrinkRequest {
	s.OtherCountry = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetOtherIp(v string) *GetInternetTupleShrinkRequest {
	s.OtherIp = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetOtherIsp(v string) *GetInternetTupleShrinkRequest {
	s.OtherIsp = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetOtherPort(v string) *GetInternetTupleShrinkRequest {
	s.OtherPort = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetProtocol(v string) *GetInternetTupleShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetRegionId(v string) *GetInternetTupleShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetSort(v string) *GetInternetTupleShrinkRequest {
	s.Sort = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetTopN(v int32) *GetInternetTupleShrinkRequest {
	s.TopN = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetTupleType(v int32) *GetInternetTupleShrinkRequest {
	s.TupleType = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetUseMultiAccount(v bool) *GetInternetTupleShrinkRequest {
	s.UseMultiAccount = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
