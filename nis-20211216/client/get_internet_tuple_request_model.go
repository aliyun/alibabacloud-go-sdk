// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInternetTupleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*int64) *GetInternetTupleRequest
	GetAccountIds() []*int64
	SetBeginTime(v int64) *GetInternetTupleRequest
	GetBeginTime() *int64
	SetCloudIp(v string) *GetInternetTupleRequest
	GetCloudIp() *string
	SetCloudIpList(v []*string) *GetInternetTupleRequest
	GetCloudIpList() []*string
	SetCloudIsp(v string) *GetInternetTupleRequest
	GetCloudIsp() *string
	SetCloudPort(v string) *GetInternetTupleRequest
	GetCloudPort() *string
	SetDirection(v string) *GetInternetTupleRequest
	GetDirection() *string
	SetEndTime(v int64) *GetInternetTupleRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetInternetTupleRequest
	GetInstanceId() *string
	SetInstanceList(v []*string) *GetInternetTupleRequest
	GetInstanceList() []*string
	SetOrderBy(v string) *GetInternetTupleRequest
	GetOrderBy() *string
	SetOtherCity(v string) *GetInternetTupleRequest
	GetOtherCity() *string
	SetOtherCountry(v string) *GetInternetTupleRequest
	GetOtherCountry() *string
	SetOtherIp(v string) *GetInternetTupleRequest
	GetOtherIp() *string
	SetOtherIsp(v string) *GetInternetTupleRequest
	GetOtherIsp() *string
	SetOtherPort(v string) *GetInternetTupleRequest
	GetOtherPort() *string
	SetProtocol(v string) *GetInternetTupleRequest
	GetProtocol() *string
	SetRegionId(v string) *GetInternetTupleRequest
	GetRegionId() *string
	SetSort(v string) *GetInternetTupleRequest
	GetSort() *string
	SetTopN(v int32) *GetInternetTupleRequest
	GetTopN() *int32
	SetTupleType(v int32) *GetInternetTupleRequest
	GetTupleType() *int32
	SetUseMultiAccount(v bool) *GetInternetTupleRequest
	GetUseMultiAccount() *bool
}

type GetInternetTupleRequest struct {
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
	CloudIpList []*string `json:"CloudIpList,omitempty" xml:"CloudIpList,omitempty" type:"Repeated"`
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
	InstanceList []*string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
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

func (s GetInternetTupleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInternetTupleRequest) GoString() string {
	return s.String()
}

func (s *GetInternetTupleRequest) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *GetInternetTupleRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetInternetTupleRequest) GetCloudIp() *string {
	return s.CloudIp
}

func (s *GetInternetTupleRequest) GetCloudIpList() []*string {
	return s.CloudIpList
}

func (s *GetInternetTupleRequest) GetCloudIsp() *string {
	return s.CloudIsp
}

func (s *GetInternetTupleRequest) GetCloudPort() *string {
	return s.CloudPort
}

func (s *GetInternetTupleRequest) GetDirection() *string {
	return s.Direction
}

func (s *GetInternetTupleRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetInternetTupleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInternetTupleRequest) GetInstanceList() []*string {
	return s.InstanceList
}

func (s *GetInternetTupleRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetInternetTupleRequest) GetOtherCity() *string {
	return s.OtherCity
}

func (s *GetInternetTupleRequest) GetOtherCountry() *string {
	return s.OtherCountry
}

func (s *GetInternetTupleRequest) GetOtherIp() *string {
	return s.OtherIp
}

func (s *GetInternetTupleRequest) GetOtherIsp() *string {
	return s.OtherIsp
}

func (s *GetInternetTupleRequest) GetOtherPort() *string {
	return s.OtherPort
}

func (s *GetInternetTupleRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *GetInternetTupleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInternetTupleRequest) GetSort() *string {
	return s.Sort
}

func (s *GetInternetTupleRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *GetInternetTupleRequest) GetTupleType() *int32 {
	return s.TupleType
}

func (s *GetInternetTupleRequest) GetUseMultiAccount() *bool {
	return s.UseMultiAccount
}

func (s *GetInternetTupleRequest) SetAccountIds(v []*int64) *GetInternetTupleRequest {
	s.AccountIds = v
	return s
}

func (s *GetInternetTupleRequest) SetBeginTime(v int64) *GetInternetTupleRequest {
	s.BeginTime = &v
	return s
}

func (s *GetInternetTupleRequest) SetCloudIp(v string) *GetInternetTupleRequest {
	s.CloudIp = &v
	return s
}

func (s *GetInternetTupleRequest) SetCloudIpList(v []*string) *GetInternetTupleRequest {
	s.CloudIpList = v
	return s
}

func (s *GetInternetTupleRequest) SetCloudIsp(v string) *GetInternetTupleRequest {
	s.CloudIsp = &v
	return s
}

func (s *GetInternetTupleRequest) SetCloudPort(v string) *GetInternetTupleRequest {
	s.CloudPort = &v
	return s
}

func (s *GetInternetTupleRequest) SetDirection(v string) *GetInternetTupleRequest {
	s.Direction = &v
	return s
}

func (s *GetInternetTupleRequest) SetEndTime(v int64) *GetInternetTupleRequest {
	s.EndTime = &v
	return s
}

func (s *GetInternetTupleRequest) SetInstanceId(v string) *GetInternetTupleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInternetTupleRequest) SetInstanceList(v []*string) *GetInternetTupleRequest {
	s.InstanceList = v
	return s
}

func (s *GetInternetTupleRequest) SetOrderBy(v string) *GetInternetTupleRequest {
	s.OrderBy = &v
	return s
}

func (s *GetInternetTupleRequest) SetOtherCity(v string) *GetInternetTupleRequest {
	s.OtherCity = &v
	return s
}

func (s *GetInternetTupleRequest) SetOtherCountry(v string) *GetInternetTupleRequest {
	s.OtherCountry = &v
	return s
}

func (s *GetInternetTupleRequest) SetOtherIp(v string) *GetInternetTupleRequest {
	s.OtherIp = &v
	return s
}

func (s *GetInternetTupleRequest) SetOtherIsp(v string) *GetInternetTupleRequest {
	s.OtherIsp = &v
	return s
}

func (s *GetInternetTupleRequest) SetOtherPort(v string) *GetInternetTupleRequest {
	s.OtherPort = &v
	return s
}

func (s *GetInternetTupleRequest) SetProtocol(v string) *GetInternetTupleRequest {
	s.Protocol = &v
	return s
}

func (s *GetInternetTupleRequest) SetRegionId(v string) *GetInternetTupleRequest {
	s.RegionId = &v
	return s
}

func (s *GetInternetTupleRequest) SetSort(v string) *GetInternetTupleRequest {
	s.Sort = &v
	return s
}

func (s *GetInternetTupleRequest) SetTopN(v int32) *GetInternetTupleRequest {
	s.TopN = &v
	return s
}

func (s *GetInternetTupleRequest) SetTupleType(v int32) *GetInternetTupleRequest {
	s.TupleType = &v
	return s
}

func (s *GetInternetTupleRequest) SetUseMultiAccount(v bool) *GetInternetTupleRequest {
	s.UseMultiAccount = &v
	return s
}

func (s *GetInternetTupleRequest) Validate() error {
	return dara.Validate(s)
}
