// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInternetTupleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetInternetTupleResponseBodyData) *GetInternetTupleResponseBody
	GetData() []*GetInternetTupleResponseBodyData
	SetRequestId(v string) *GetInternetTupleResponseBody
	GetRequestId() *string
}

type GetInternetTupleResponseBody struct {
	// The ranking result of Internet traffic data.
	Data []*GetInternetTupleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D5E98683-355B-5867-8D3D-A24755F6895B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInternetTupleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInternetTupleResponseBody) GoString() string {
	return s.String()
}

func (s *GetInternetTupleResponseBody) GetData() []*GetInternetTupleResponseBodyData {
	return s.Data
}

func (s *GetInternetTupleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInternetTupleResponseBody) SetData(v []*GetInternetTupleResponseBodyData) *GetInternetTupleResponseBody {
	s.Data = v
	return s
}

func (s *GetInternetTupleResponseBody) SetRequestId(v string) *GetInternetTupleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInternetTupleResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInternetTupleResponseBodyData struct {
	// The access point of Alibaba Cloud.
	//
	// >  This parameter is valid only if you set **InstanceId*	- to the instance ID of an Anycast elastic IP address (EIP).
	//
	// example:
	//
	// cn-hongkong-pop
	AccessRegion *string `json:"AccessRegion,omitempty" xml:"AccessRegion,omitempty"`
	// The beginning of the time range that you queried. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1684373600099
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The traffic volume. Unit: bytes.
	//
	// example:
	//
	// 88
	ByteCount *float64 `json:"ByteCount,omitempty" xml:"ByteCount,omitempty"`
	// The local city.
	//
	// example:
	//
	// Nanjing
	CloudCity *string `json:"CloudCity,omitempty" xml:"CloudCity,omitempty"`
	// The local country or region.
	//
	// example:
	//
	// China
	CloudCountry *string `json:"CloudCountry,omitempty" xml:"CloudCountry,omitempty"`
	// The local IP address.
	//
	// example:
	//
	// 112.74.XX.XX
	CloudIp *string `json:"CloudIp,omitempty" xml:"CloudIp,omitempty"`
	// The local ISP.
	//
	// example:
	//
	// China Mobile
	CloudIsp *string `json:"CloudIsp,omitempty" xml:"CloudIsp,omitempty"`
	// The local port.
	//
	// example:
	//
	// 443
	CloudPort *string `json:"CloudPort,omitempty" xml:"CloudPort,omitempty"`
	// The service code of the instance to which the local IP address belongs.
	//
	// example:
	//
	// EIP
	CloudProduct *string `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	// The local province.
	//
	// example:
	//
	// Jiangsu
	CloudProvince *string `json:"CloudProvince,omitempty" xml:"CloudProvince,omitempty"`
	// The direction of Internet traffic. Valid values:
	//
	// 	- **in**: inbound
	//
	// 	- **out**: outbound
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The inbound traffic volume. Unit: bytes.
	//
	// example:
	//
	// 88
	InByteCount *float64 `json:"InByteCount,omitempty" xml:"InByteCount,omitempty"`
	// The number of inbound disordered packets.
	//
	// example:
	//
	// 2
	InOutOrderCount *float64 `json:"InOutOrderCount,omitempty" xml:"InOutOrderCount,omitempty"`
	// The number of inbound packets.
	//
	// example:
	//
	// 33
	InPacketCount *float64 `json:"InPacketCount,omitempty" xml:"InPacketCount,omitempty"`
	// The number of inbound repeated packets.
	//
	// example:
	//
	// 0
	InRetranCount *float64 `json:"InRetranCount,omitempty" xml:"InRetranCount,omitempty"`
	// The ID of the instance to which the local IP address belongs.
	//
	// example:
	//
	// eip-sample*
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The remote city. In most cases, this parameter is empty if you set **OtherCountry*	- to a country except China.
	//
	// example:
	//
	// Austin
	OtherCity *string `json:"OtherCity,omitempty" xml:"OtherCity,omitempty"`
	// The remote country or region.
	//
	// example:
	//
	// United States
	OtherCountry *string `json:"OtherCountry,omitempty" xml:"OtherCountry,omitempty"`
	// The remote IP address.
	//
	// example:
	//
	// 122.112.XX.XX
	OtherIp *string `json:"OtherIp,omitempty" xml:"OtherIp,omitempty"`
	// The remote ISP.
	//
	// example:
	//
	// amazon.com
	OtherIsp *string `json:"OtherIsp,omitempty" xml:"OtherIsp,omitempty"`
	// The remote port.
	//
	// example:
	//
	// 40002
	OtherPort *string `json:"OtherPort,omitempty" xml:"OtherPort,omitempty"`
	// The service code of the instance to which the remote IP address belongs. If the IP address is not on the cloud, this parameter is empty.
	//
	// example:
	//
	// ECS
	OtherProduct *string `json:"OtherProduct,omitempty" xml:"OtherProduct,omitempty"`
	// The remote province. In most cases, this parameter is empty if you set **OtherCountry*	- to a country except China.
	//
	// example:
	//
	// Texas
	OtherProvince *string `json:"OtherProvince,omitempty" xml:"OtherProvince,omitempty"`
	// The outbound traffic volume. Unit: bytes.
	//
	// example:
	//
	// 66
	OutByteCount *float64 `json:"OutByteCount,omitempty" xml:"OutByteCount,omitempty"`
	// The number of disordered packets.
	//
	// example:
	//
	// 1
	OutOrderCount *float64 `json:"OutOrderCount,omitempty" xml:"OutOrderCount,omitempty"`
	// The number of outbound disordered packets.
	//
	// example:
	//
	// 1
	OutOutOrderCount *float64 `json:"OutOutOrderCount,omitempty" xml:"OutOutOrderCount,omitempty"`
	// The number of outbound packets.
	//
	// example:
	//
	// 22
	OutPacketCount *float64 `json:"OutPacketCount,omitempty" xml:"OutPacketCount,omitempty"`
	// The number of outbound repeated packets.
	//
	// example:
	//
	// 1
	OutRetranCount *float64 `json:"OutRetranCount,omitempty" xml:"OutRetranCount,omitempty"`
	// The number of packets.
	//
	// example:
	//
	// 66
	PacketCount *float64 `json:"PacketCount,omitempty" xml:"PacketCount,omitempty"`
	// The protocol number.
	//
	// example:
	//
	// 6
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The retransmission rate of TCP packets.
	//
	// example:
	//
	// 0.1
	RetransmitRate *float64 `json:"RetransmitRate,omitempty" xml:"RetransmitRate,omitempty"`
	// The round-trip time (RTT). Unit: milliseconds.
	//
	// example:
	//
	// 10000
	Rtt *float64 `json:"Rtt,omitempty" xml:"Rtt,omitempty"`
}

func (s GetInternetTupleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInternetTupleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInternetTupleResponseBodyData) GetAccessRegion() *string {
	return s.AccessRegion
}

func (s *GetInternetTupleResponseBodyData) GetBeginTime() *string {
	return s.BeginTime
}

func (s *GetInternetTupleResponseBodyData) GetByteCount() *float64 {
	return s.ByteCount
}

func (s *GetInternetTupleResponseBodyData) GetCloudCity() *string {
	return s.CloudCity
}

func (s *GetInternetTupleResponseBodyData) GetCloudCountry() *string {
	return s.CloudCountry
}

func (s *GetInternetTupleResponseBodyData) GetCloudIp() *string {
	return s.CloudIp
}

func (s *GetInternetTupleResponseBodyData) GetCloudIsp() *string {
	return s.CloudIsp
}

func (s *GetInternetTupleResponseBodyData) GetCloudPort() *string {
	return s.CloudPort
}

func (s *GetInternetTupleResponseBodyData) GetCloudProduct() *string {
	return s.CloudProduct
}

func (s *GetInternetTupleResponseBodyData) GetCloudProvince() *string {
	return s.CloudProvince
}

func (s *GetInternetTupleResponseBodyData) GetDirection() *string {
	return s.Direction
}

func (s *GetInternetTupleResponseBodyData) GetInByteCount() *float64 {
	return s.InByteCount
}

func (s *GetInternetTupleResponseBodyData) GetInOutOrderCount() *float64 {
	return s.InOutOrderCount
}

func (s *GetInternetTupleResponseBodyData) GetInPacketCount() *float64 {
	return s.InPacketCount
}

func (s *GetInternetTupleResponseBodyData) GetInRetranCount() *float64 {
	return s.InRetranCount
}

func (s *GetInternetTupleResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInternetTupleResponseBodyData) GetOtherCity() *string {
	return s.OtherCity
}

func (s *GetInternetTupleResponseBodyData) GetOtherCountry() *string {
	return s.OtherCountry
}

func (s *GetInternetTupleResponseBodyData) GetOtherIp() *string {
	return s.OtherIp
}

func (s *GetInternetTupleResponseBodyData) GetOtherIsp() *string {
	return s.OtherIsp
}

func (s *GetInternetTupleResponseBodyData) GetOtherPort() *string {
	return s.OtherPort
}

func (s *GetInternetTupleResponseBodyData) GetOtherProduct() *string {
	return s.OtherProduct
}

func (s *GetInternetTupleResponseBodyData) GetOtherProvince() *string {
	return s.OtherProvince
}

func (s *GetInternetTupleResponseBodyData) GetOutByteCount() *float64 {
	return s.OutByteCount
}

func (s *GetInternetTupleResponseBodyData) GetOutOrderCount() *float64 {
	return s.OutOrderCount
}

func (s *GetInternetTupleResponseBodyData) GetOutOutOrderCount() *float64 {
	return s.OutOutOrderCount
}

func (s *GetInternetTupleResponseBodyData) GetOutPacketCount() *float64 {
	return s.OutPacketCount
}

func (s *GetInternetTupleResponseBodyData) GetOutRetranCount() *float64 {
	return s.OutRetranCount
}

func (s *GetInternetTupleResponseBodyData) GetPacketCount() *float64 {
	return s.PacketCount
}

func (s *GetInternetTupleResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *GetInternetTupleResponseBodyData) GetRetransmitRate() *float64 {
	return s.RetransmitRate
}

func (s *GetInternetTupleResponseBodyData) GetRtt() *float64 {
	return s.Rtt
}

func (s *GetInternetTupleResponseBodyData) SetAccessRegion(v string) *GetInternetTupleResponseBodyData {
	s.AccessRegion = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetBeginTime(v string) *GetInternetTupleResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetByteCount(v float64) *GetInternetTupleResponseBodyData {
	s.ByteCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetCloudCity(v string) *GetInternetTupleResponseBodyData {
	s.CloudCity = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetCloudCountry(v string) *GetInternetTupleResponseBodyData {
	s.CloudCountry = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetCloudIp(v string) *GetInternetTupleResponseBodyData {
	s.CloudIp = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetCloudIsp(v string) *GetInternetTupleResponseBodyData {
	s.CloudIsp = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetCloudPort(v string) *GetInternetTupleResponseBodyData {
	s.CloudPort = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetCloudProduct(v string) *GetInternetTupleResponseBodyData {
	s.CloudProduct = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetCloudProvince(v string) *GetInternetTupleResponseBodyData {
	s.CloudProvince = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetDirection(v string) *GetInternetTupleResponseBodyData {
	s.Direction = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetInByteCount(v float64) *GetInternetTupleResponseBodyData {
	s.InByteCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetInOutOrderCount(v float64) *GetInternetTupleResponseBodyData {
	s.InOutOrderCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetInPacketCount(v float64) *GetInternetTupleResponseBodyData {
	s.InPacketCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetInRetranCount(v float64) *GetInternetTupleResponseBodyData {
	s.InRetranCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetInstanceId(v string) *GetInternetTupleResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOtherCity(v string) *GetInternetTupleResponseBodyData {
	s.OtherCity = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOtherCountry(v string) *GetInternetTupleResponseBodyData {
	s.OtherCountry = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOtherIp(v string) *GetInternetTupleResponseBodyData {
	s.OtherIp = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOtherIsp(v string) *GetInternetTupleResponseBodyData {
	s.OtherIsp = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOtherPort(v string) *GetInternetTupleResponseBodyData {
	s.OtherPort = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOtherProduct(v string) *GetInternetTupleResponseBodyData {
	s.OtherProduct = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOtherProvince(v string) *GetInternetTupleResponseBodyData {
	s.OtherProvince = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOutByteCount(v float64) *GetInternetTupleResponseBodyData {
	s.OutByteCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOutOrderCount(v float64) *GetInternetTupleResponseBodyData {
	s.OutOrderCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOutOutOrderCount(v float64) *GetInternetTupleResponseBodyData {
	s.OutOutOrderCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOutPacketCount(v float64) *GetInternetTupleResponseBodyData {
	s.OutPacketCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOutRetranCount(v float64) *GetInternetTupleResponseBodyData {
	s.OutRetranCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetPacketCount(v float64) *GetInternetTupleResponseBodyData {
	s.PacketCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetProtocol(v string) *GetInternetTupleResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetRetransmitRate(v float64) *GetInternetTupleResponseBodyData {
	s.RetransmitRate = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetRtt(v float64) *GetInternetTupleResponseBodyData {
	s.Rtt = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
