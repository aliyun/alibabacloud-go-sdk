// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnGatewaysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVpnGatewaysResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpnGatewaysResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVpnGatewaysResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpnGatewaysResponseBody
	GetTotalCount() *int32
	SetVpnGateways(v *DescribeVpnGatewaysResponseBodyVpnGateways) *DescribeVpnGatewaysResponseBody
	GetVpnGateways() *DescribeVpnGatewaysResponseBodyVpnGateways
}

type DescribeVpnGatewaysResponseBody struct {
	// The page number of the list.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for paging queries.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DF11D6F6-E35A-41C3-9B20-6FC8A901FE65
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount  *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VpnGateways *DescribeVpnGatewaysResponseBodyVpnGateways `json:"VpnGateways,omitempty" xml:"VpnGateways,omitempty" type:"Struct"`
}

func (s DescribeVpnGatewaysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewaysResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpnGatewaysResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpnGatewaysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpnGatewaysResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpnGatewaysResponseBody) GetVpnGateways() *DescribeVpnGatewaysResponseBodyVpnGateways {
	return s.VpnGateways
}

func (s *DescribeVpnGatewaysResponseBody) SetPageNumber(v int32) *DescribeVpnGatewaysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBody) SetPageSize(v int32) *DescribeVpnGatewaysResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBody) SetRequestId(v string) *DescribeVpnGatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBody) SetTotalCount(v int32) *DescribeVpnGatewaysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBody) SetVpnGateways(v *DescribeVpnGatewaysResponseBodyVpnGateways) *DescribeVpnGatewaysResponseBody {
	s.VpnGateways = v
	return s
}

func (s *DescribeVpnGatewaysResponseBody) Validate() error {
	if s.VpnGateways != nil {
		if err := s.VpnGateways.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVpnGatewaysResponseBodyVpnGateways struct {
	VpnGateway []*DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway `json:"VpnGateway,omitempty" xml:"VpnGateway,omitempty" type:"Repeated"`
}

func (s DescribeVpnGatewaysResponseBodyVpnGateways) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewaysResponseBodyVpnGateways) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewaysResponseBodyVpnGateways) GetVpnGateway() []*DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	return s.VpnGateway
}

func (s *DescribeVpnGatewaysResponseBodyVpnGateways) SetVpnGateway(v []*DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) *DescribeVpnGatewaysResponseBodyVpnGateways {
	s.VpnGateway = v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGateways) Validate() error {
	if s.VpnGateway != nil {
		for _, item := range s.VpnGateway {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway struct {
	AutoPropagate              *bool                                                                `json:"AutoPropagate,omitempty" xml:"AutoPropagate,omitempty"`
	BusinessStatus             *string                                                              `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	ChargeType                 *string                                                              `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreateTime                 *int64                                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description                *string                                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	DisasterRecoveryInternetIp *string                                                              `json:"DisasterRecoveryInternetIp,omitempty" xml:"DisasterRecoveryInternetIp,omitempty"`
	DisasterRecoveryVSwitchId  *string                                                              `json:"DisasterRecoveryVSwitchId,omitempty" xml:"DisasterRecoveryVSwitchId,omitempty"`
	EnableBgp                  *bool                                                                `json:"EnableBgp,omitempty" xml:"EnableBgp,omitempty"`
	EndTime                    *int64                                                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EniInstanceIds             *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds  `json:"EniInstanceIds,omitempty" xml:"EniInstanceIds,omitempty" type:"Struct"`
	GatewayType                *string                                                              `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	InternetIp                 *string                                                              `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	IpsecVpn                   *string                                                              `json:"IpsecVpn,omitempty" xml:"IpsecVpn,omitempty"`
	Name                       *string                                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	NetworkType                *string                                                              `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	ReservationData            *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData `json:"ReservationData,omitempty" xml:"ReservationData,omitempty" type:"Struct"`
	ResourceGroupId            *string                                                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Spec                       *string                                                              `json:"Spec,omitempty" xml:"Spec,omitempty"`
	SslMaxConnections          *int64                                                               `json:"SslMaxConnections,omitempty" xml:"SslMaxConnections,omitempty"`
	SslVpn                     *string                                                              `json:"SslVpn,omitempty" xml:"SslVpn,omitempty"`
	SslVpnInternetIp           *string                                                              `json:"SslVpnInternetIp,omitempty" xml:"SslVpnInternetIp,omitempty"`
	Status                     *string                                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag                        *string                                                              `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Tags                       *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags            `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VSwitchId                  *string                                                              `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                      *string                                                              `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpnGatewayId               *string                                                              `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
	VpnType                    *string                                                              `json:"VpnType,omitempty" xml:"VpnType,omitempty"`
}

func (s DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetAutoPropagate() *bool {
	return s.AutoPropagate
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetDescription() *string {
	return s.Description
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetDisasterRecoveryInternetIp() *string {
	return s.DisasterRecoveryInternetIp
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetDisasterRecoveryVSwitchId() *string {
	return s.DisasterRecoveryVSwitchId
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetEnableBgp() *bool {
	return s.EnableBgp
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetEniInstanceIds() *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds {
	return s.EniInstanceIds
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetGatewayType() *string {
	return s.GatewayType
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetIpsecVpn() *string {
	return s.IpsecVpn
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetName() *string {
	return s.Name
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetReservationData() *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData {
	return s.ReservationData
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetSpec() *string {
	return s.Spec
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetSslMaxConnections() *int64 {
	return s.SslMaxConnections
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetSslVpn() *string {
	return s.SslVpn
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetSslVpnInternetIp() *string {
	return s.SslVpnInternetIp
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetTag() *string {
	return s.Tag
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetTags() *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags {
	return s.Tags
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) GetVpnType() *string {
	return s.VpnType
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetAutoPropagate(v bool) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.AutoPropagate = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetBusinessStatus(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetChargeType(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.ChargeType = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetCreateTime(v int64) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.CreateTime = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetDescription(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.Description = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetDisasterRecoveryInternetIp(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.DisasterRecoveryInternetIp = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetDisasterRecoveryVSwitchId(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.DisasterRecoveryVSwitchId = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetEnableBgp(v bool) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.EnableBgp = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetEndTime(v int64) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.EndTime = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetEniInstanceIds(v *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.EniInstanceIds = v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetGatewayType(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.GatewayType = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetInternetIp(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.InternetIp = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetIpsecVpn(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.IpsecVpn = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetName(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.Name = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetNetworkType(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.NetworkType = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetReservationData(v *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.ReservationData = v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetResourceGroupId(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetSpec(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.Spec = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetSslMaxConnections(v int64) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.SslMaxConnections = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetSslVpn(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.SslVpn = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetSslVpnInternetIp(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.SslVpnInternetIp = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetStatus(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.Status = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetTag(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.Tag = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetTags(v *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.Tags = v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetVSwitchId(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetVpcId(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.VpcId = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetVpnGatewayId(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.VpnGatewayId = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) SetVpnType(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway {
	s.VpnType = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGateway) Validate() error {
	if s.EniInstanceIds != nil {
		if err := s.EniInstanceIds.Validate(); err != nil {
			return err
		}
	}
	if s.ReservationData != nil {
		if err := s.ReservationData.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds struct {
	EniInstanceId []*string `json:"EniInstanceId,omitempty" xml:"EniInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds) GetEniInstanceId() []*string {
	return s.EniInstanceId
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds) SetEniInstanceId(v []*string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds {
	s.EniInstanceId = v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayEniInstanceIds) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData struct {
	ReservationEndTime        *string `json:"ReservationEndTime,omitempty" xml:"ReservationEndTime,omitempty"`
	ReservationIpsec          *string `json:"ReservationIpsec,omitempty" xml:"ReservationIpsec,omitempty"`
	ReservationMaxConnections *int32  `json:"ReservationMaxConnections,omitempty" xml:"ReservationMaxConnections,omitempty"`
	ReservationOrderType      *string `json:"ReservationOrderType,omitempty" xml:"ReservationOrderType,omitempty"`
	ReservationSpec           *string `json:"ReservationSpec,omitempty" xml:"ReservationSpec,omitempty"`
	ReservationSsl            *string `json:"ReservationSsl,omitempty" xml:"ReservationSsl,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) GetReservationEndTime() *string {
	return s.ReservationEndTime
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) GetReservationIpsec() *string {
	return s.ReservationIpsec
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) GetReservationMaxConnections() *int32 {
	return s.ReservationMaxConnections
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) GetReservationOrderType() *string {
	return s.ReservationOrderType
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) GetReservationSpec() *string {
	return s.ReservationSpec
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) GetReservationSsl() *string {
	return s.ReservationSsl
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) SetReservationEndTime(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData {
	s.ReservationEndTime = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) SetReservationIpsec(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData {
	s.ReservationIpsec = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) SetReservationMaxConnections(v int32) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData {
	s.ReservationMaxConnections = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) SetReservationOrderType(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData {
	s.ReservationOrderType = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) SetReservationSpec(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData {
	s.ReservationSpec = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) SetReservationSsl(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData {
	s.ReservationSsl = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) SetStatus(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData {
	s.Status = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayReservationData) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags struct {
	Tag []*DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags) GetTag() []*DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag {
	return s.Tag
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags) SetTag(v []*DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags {
	s.Tag = v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag) SetKey(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag) SetValue(v string) *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeVpnGatewaysResponseBodyVpnGatewaysVpnGatewayTagsTag) Validate() error {
	return dara.Validate(s)
}
