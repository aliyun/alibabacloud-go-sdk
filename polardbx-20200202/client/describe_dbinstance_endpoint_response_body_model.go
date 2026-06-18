// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDBInstanceEndpointResponseBodyData) *DescribeDBInstanceEndpointResponseBody
	GetData() *DescribeDBInstanceEndpointResponseBodyData
	SetMaxResults(v int32) *DescribeDBInstanceEndpointResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDBInstanceEndpointResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeDBInstanceEndpointResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceEndpointResponseBody struct {
	Data *DescribeDBInstanceEndpointResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1000
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// xxdds
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// A501A191-BD70-5E50-98A9-C2A486A82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEndpointResponseBody) GetData() *DescribeDBInstanceEndpointResponseBodyData {
	return s.Data
}

func (s *DescribeDBInstanceEndpointResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDBInstanceEndpointResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDBInstanceEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceEndpointResponseBody) SetData(v *DescribeDBInstanceEndpointResponseBodyData) *DescribeDBInstanceEndpointResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBody) SetMaxResults(v int32) *DescribeDBInstanceEndpointResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBody) SetNextToken(v string) *DescribeDBInstanceEndpointResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBody) SetRequestId(v string) *DescribeDBInstanceEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceEndpointResponseBodyData struct {
	Items []*DescribeDBInstanceEndpointResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceEndpointResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEndpointResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEndpointResponseBodyData) GetItems() []*DescribeDBInstanceEndpointResponseBodyDataItems {
	return s.Items
}

func (s *DescribeDBInstanceEndpointResponseBodyData) SetItems(v []*DescribeDBInstanceEndpointResponseBodyDataItems) *DescribeDBInstanceEndpointResponseBodyData {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceEndpointResponseBodyDataItems struct {
	Endpoint   *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint     `json:"Endpoint,omitempty" xml:"Endpoint,omitempty" type:"Struct"`
	RealServer []*DescribeDBInstanceEndpointResponseBodyDataItemsRealServer `json:"RealServer,omitempty" xml:"RealServer,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceEndpointResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEndpointResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItems) GetEndpoint() *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint {
	return s.Endpoint
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItems) GetRealServer() []*DescribeDBInstanceEndpointResponseBodyDataItemsRealServer {
	return s.RealServer
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItems) SetEndpoint(v *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) *DescribeDBInstanceEndpointResponseBodyDataItems {
	s.Endpoint = v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItems) SetRealServer(v []*DescribeDBInstanceEndpointResponseBodyDataItemsRealServer) *DescribeDBInstanceEndpointResponseBodyDataItems {
	s.RealServer = v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItems) Validate() error {
	if s.Endpoint != nil {
		if err := s.Endpoint.Validate(); err != nil {
			return err
		}
	}
	if s.RealServer != nil {
		for _, item := range s.RealServer {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint struct {
	// example:
	//
	// 10.21.1.82
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// mdb.shard.4x.large.d
	Class *string `json:"Class,omitempty" xml:"Class,omitempty"`
	// example:
	//
	// epg-bp14wgzai7flglwdtkxfa
	EndpointGroupId *int64 `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// example:
	//
	// 2899
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// True
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// instance
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// example:
	//
	// VPC
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// example:
	//
	// 0
	ReadType *string `json:"ReadType,omitempty" xml:"ReadType,omitempty"`
	// example:
	//
	// sas-app
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// example:
	//
	// 3522367
	TunnelId *int64 `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	// example:
	//
	// TXT
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UserVisible *bool   `json:"UserVisible,omitempty" xml:"UserVisible,omitempty"`
	// example:
	//
	// vsw-2ze51hh6s8tsjgy19g5eu
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// http://100.100.100.200/latest/meta-data
	Vip *string `json:"Vip,omitempty" xml:"Vip,omitempty"`
	// example:
	//
	// vpc-bp1s9j8s4h4uqejp9k2z3
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vport
	Vport *int64 `json:"Vport,omitempty" xml:"Vport,omitempty"`
	// example:
	//
	// `curl uUys2ThR.popscan.xaliyun.com`
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) GetAddress() *string {
	return s.Address
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) GetClass() *string {
	return s.Class
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) GetEndpointGroupId() *int64 {
	return s.EndpointGroupId
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) GetId() *int64 {
	return s.Id
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) GetKind() *string {
	return s.Kind
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) GetNetType() *string {
	return s.NetType
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) GetReadType() *string {
	return s.ReadType
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) GetTargetName() *string {
	return s.TargetName
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) GetTunnelId() *int64 {
	return s.TunnelId
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) GetType() *string {
	return s.Type
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) GetUserVisible() *bool {
	return s.UserVisible
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) GetVip() *string {
	return s.Vip
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) GetVport() *int64 {
	return s.Vport
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) SetAddress(v string) *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint {
	s.Address = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) SetClass(v string) *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint {
	s.Class = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) SetEndpointGroupId(v int64) *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint {
	s.EndpointGroupId = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) SetId(v int64) *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint {
	s.Id = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) SetIsDefault(v bool) *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint {
	s.IsDefault = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) SetKind(v string) *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint {
	s.Kind = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) SetNetType(v string) *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint {
	s.NetType = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) SetReadType(v string) *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint {
	s.ReadType = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) SetTargetName(v string) *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint {
	s.TargetName = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) SetTunnelId(v int64) *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint {
	s.TunnelId = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) SetType(v string) *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint {
	s.Type = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) SetUserVisible(v bool) *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint {
	s.UserVisible = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) SetVSwitchId(v string) *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) SetVip(v string) *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint {
	s.Vip = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) SetVpcId(v string) *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) SetVport(v int64) *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint {
	s.Vport = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) SetZoneId(v string) *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceEndpointResponseBodyDataItemsRealServer struct {
	// example:
	//
	// False
	Activated *bool `json:"Activated,omitempty" xml:"Activated,omitempty"`
	// example:
	//
	// dds.cs.mid
	Class *string `json:"Class,omitempty" xml:"Class,omitempty"`
	// example:
	//
	// 172.29.32.166
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 节点id
	ReplicaId *int64 `json:"ReplicaId,omitempty" xml:"ReplicaId,omitempty"`
	// example:
	//
	// 255
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeDBInstanceEndpointResponseBodyDataItemsRealServer) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEndpointResponseBodyDataItemsRealServer) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsRealServer) GetActivated() *bool {
	return s.Activated
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsRealServer) GetClass() *string {
	return s.Class
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsRealServer) GetIp() *string {
	return s.Ip
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsRealServer) GetPort() *string {
	return s.Port
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsRealServer) GetReplicaId() *int64 {
	return s.ReplicaId
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsRealServer) GetWeight() *int64 {
	return s.Weight
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsRealServer) SetActivated(v bool) *DescribeDBInstanceEndpointResponseBodyDataItemsRealServer {
	s.Activated = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsRealServer) SetClass(v string) *DescribeDBInstanceEndpointResponseBodyDataItemsRealServer {
	s.Class = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsRealServer) SetIp(v string) *DescribeDBInstanceEndpointResponseBodyDataItemsRealServer {
	s.Ip = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsRealServer) SetPort(v string) *DescribeDBInstanceEndpointResponseBodyDataItemsRealServer {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsRealServer) SetReplicaId(v int64) *DescribeDBInstanceEndpointResponseBodyDataItemsRealServer {
	s.ReplicaId = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsRealServer) SetWeight(v int64) *DescribeDBInstanceEndpointResponseBodyDataItemsRealServer {
	s.Weight = &v
	return s
}

func (s *DescribeDBInstanceEndpointResponseBodyDataItemsRealServer) Validate() error {
	return dara.Validate(s)
}
