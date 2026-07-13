// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNatGatewayStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetNatGatewayStatusResponseBody
	GetCode() *string
	SetData(v *GetNatGatewayStatusResponseBodyData) *GetNatGatewayStatusResponseBody
	GetData() *GetNatGatewayStatusResponseBodyData
	SetHttpStatusCode(v int32) *GetNatGatewayStatusResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetNatGatewayStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetNatGatewayStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetNatGatewayStatusResponseBody
	GetSuccess() *bool
}

type GetNatGatewayStatusResponseBody struct {
	// example:
	//
	// Success
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetNatGatewayStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// req-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNatGatewayStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNatGatewayStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetNatGatewayStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetNatGatewayStatusResponseBody) GetData() *GetNatGatewayStatusResponseBodyData {
	return s.Data
}

func (s *GetNatGatewayStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetNatGatewayStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetNatGatewayStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNatGatewayStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNatGatewayStatusResponseBody) SetCode(v string) *GetNatGatewayStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetNatGatewayStatusResponseBody) SetData(v *GetNatGatewayStatusResponseBodyData) *GetNatGatewayStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetNatGatewayStatusResponseBody) SetHttpStatusCode(v int32) *GetNatGatewayStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetNatGatewayStatusResponseBody) SetMessage(v string) *GetNatGatewayStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetNatGatewayStatusResponseBody) SetRequestId(v string) *GetNatGatewayStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNatGatewayStatusResponseBody) SetSuccess(v bool) *GetNatGatewayStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetNatGatewayStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNatGatewayStatusResponseBodyData struct {
	InstanceId           *string                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NatGatewayConfigured *bool                                             `json:"NatGatewayConfigured,omitempty" xml:"NatGatewayConfigured,omitempty"`
	NatGateways          []*GetNatGatewayStatusResponseBodyDataNatGateways `json:"NatGateways,omitempty" xml:"NatGateways,omitempty" type:"Repeated"`
	SnatConfigured       *bool                                             `json:"SnatConfigured,omitempty" xml:"SnatConfigured,omitempty"`
	Status               *string                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId                *string                                           `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneCidrCovered      *bool                                             `json:"ZoneCidrCovered,omitempty" xml:"ZoneCidrCovered,omitempty"`
	ZoneCidrs            []*GetNatGatewayStatusResponseBodyDataZoneCidrs   `json:"ZoneCidrs,omitempty" xml:"ZoneCidrs,omitempty" type:"Repeated"`
}

func (s GetNatGatewayStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetNatGatewayStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNatGatewayStatusResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetNatGatewayStatusResponseBodyData) GetNatGatewayConfigured() *bool {
	return s.NatGatewayConfigured
}

func (s *GetNatGatewayStatusResponseBodyData) GetNatGateways() []*GetNatGatewayStatusResponseBodyDataNatGateways {
	return s.NatGateways
}

func (s *GetNatGatewayStatusResponseBodyData) GetSnatConfigured() *bool {
	return s.SnatConfigured
}

func (s *GetNatGatewayStatusResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetNatGatewayStatusResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *GetNatGatewayStatusResponseBodyData) GetZoneCidrCovered() *bool {
	return s.ZoneCidrCovered
}

func (s *GetNatGatewayStatusResponseBodyData) GetZoneCidrs() []*GetNatGatewayStatusResponseBodyDataZoneCidrs {
	return s.ZoneCidrs
}

func (s *GetNatGatewayStatusResponseBodyData) SetInstanceId(v string) *GetNatGatewayStatusResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetNatGatewayStatusResponseBodyData) SetNatGatewayConfigured(v bool) *GetNatGatewayStatusResponseBodyData {
	s.NatGatewayConfigured = &v
	return s
}

func (s *GetNatGatewayStatusResponseBodyData) SetNatGateways(v []*GetNatGatewayStatusResponseBodyDataNatGateways) *GetNatGatewayStatusResponseBodyData {
	s.NatGateways = v
	return s
}

func (s *GetNatGatewayStatusResponseBodyData) SetSnatConfigured(v bool) *GetNatGatewayStatusResponseBodyData {
	s.SnatConfigured = &v
	return s
}

func (s *GetNatGatewayStatusResponseBodyData) SetStatus(v string) *GetNatGatewayStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetNatGatewayStatusResponseBodyData) SetVpcId(v string) *GetNatGatewayStatusResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *GetNatGatewayStatusResponseBodyData) SetZoneCidrCovered(v bool) *GetNatGatewayStatusResponseBodyData {
	s.ZoneCidrCovered = &v
	return s
}

func (s *GetNatGatewayStatusResponseBodyData) SetZoneCidrs(v []*GetNatGatewayStatusResponseBodyDataZoneCidrs) *GetNatGatewayStatusResponseBodyData {
	s.ZoneCidrs = v
	return s
}

func (s *GetNatGatewayStatusResponseBodyData) Validate() error {
	if s.NatGateways != nil {
		for _, item := range s.NatGateways {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ZoneCidrs != nil {
		for _, item := range s.ZoneCidrs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetNatGatewayStatusResponseBodyDataNatGateways struct {
	NatGatewayId   *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	SnatConfigured *bool   `json:"SnatConfigured,omitempty" xml:"SnatConfigured,omitempty"`
	SnatTableId    *string `json:"SnatTableId,omitempty" xml:"SnatTableId,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetNatGatewayStatusResponseBodyDataNatGateways) String() string {
	return dara.Prettify(s)
}

func (s GetNatGatewayStatusResponseBodyDataNatGateways) GoString() string {
	return s.String()
}

func (s *GetNatGatewayStatusResponseBodyDataNatGateways) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *GetNatGatewayStatusResponseBodyDataNatGateways) GetSnatConfigured() *bool {
	return s.SnatConfigured
}

func (s *GetNatGatewayStatusResponseBodyDataNatGateways) GetSnatTableId() *string {
	return s.SnatTableId
}

func (s *GetNatGatewayStatusResponseBodyDataNatGateways) GetStatus() *string {
	return s.Status
}

func (s *GetNatGatewayStatusResponseBodyDataNatGateways) SetNatGatewayId(v string) *GetNatGatewayStatusResponseBodyDataNatGateways {
	s.NatGatewayId = &v
	return s
}

func (s *GetNatGatewayStatusResponseBodyDataNatGateways) SetSnatConfigured(v bool) *GetNatGatewayStatusResponseBodyDataNatGateways {
	s.SnatConfigured = &v
	return s
}

func (s *GetNatGatewayStatusResponseBodyDataNatGateways) SetSnatTableId(v string) *GetNatGatewayStatusResponseBodyDataNatGateways {
	s.SnatTableId = &v
	return s
}

func (s *GetNatGatewayStatusResponseBodyDataNatGateways) SetStatus(v string) *GetNatGatewayStatusResponseBodyDataNatGateways {
	s.Status = &v
	return s
}

func (s *GetNatGatewayStatusResponseBodyDataNatGateways) Validate() error {
	return dara.Validate(s)
}

type GetNatGatewayStatusResponseBodyDataZoneCidrs struct {
	CidrBlock      *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	Covered        *bool   `json:"Covered,omitempty" xml:"Covered,omitempty"`
	NatGatewayId   *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	SnatEntryId    *string `json:"SnatEntryId,omitempty" xml:"SnatEntryId,omitempty"`
	SnatSourceCidr *string `json:"SnatSourceCidr,omitempty" xml:"SnatSourceCidr,omitempty"`
	VSwitchId      *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId         *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetNatGatewayStatusResponseBodyDataZoneCidrs) String() string {
	return dara.Prettify(s)
}

func (s GetNatGatewayStatusResponseBodyDataZoneCidrs) GoString() string {
	return s.String()
}

func (s *GetNatGatewayStatusResponseBodyDataZoneCidrs) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *GetNatGatewayStatusResponseBodyDataZoneCidrs) GetCovered() *bool {
	return s.Covered
}

func (s *GetNatGatewayStatusResponseBodyDataZoneCidrs) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *GetNatGatewayStatusResponseBodyDataZoneCidrs) GetSnatEntryId() *string {
	return s.SnatEntryId
}

func (s *GetNatGatewayStatusResponseBodyDataZoneCidrs) GetSnatSourceCidr() *string {
	return s.SnatSourceCidr
}

func (s *GetNatGatewayStatusResponseBodyDataZoneCidrs) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetNatGatewayStatusResponseBodyDataZoneCidrs) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetNatGatewayStatusResponseBodyDataZoneCidrs) SetCidrBlock(v string) *GetNatGatewayStatusResponseBodyDataZoneCidrs {
	s.CidrBlock = &v
	return s
}

func (s *GetNatGatewayStatusResponseBodyDataZoneCidrs) SetCovered(v bool) *GetNatGatewayStatusResponseBodyDataZoneCidrs {
	s.Covered = &v
	return s
}

func (s *GetNatGatewayStatusResponseBodyDataZoneCidrs) SetNatGatewayId(v string) *GetNatGatewayStatusResponseBodyDataZoneCidrs {
	s.NatGatewayId = &v
	return s
}

func (s *GetNatGatewayStatusResponseBodyDataZoneCidrs) SetSnatEntryId(v string) *GetNatGatewayStatusResponseBodyDataZoneCidrs {
	s.SnatEntryId = &v
	return s
}

func (s *GetNatGatewayStatusResponseBodyDataZoneCidrs) SetSnatSourceCidr(v string) *GetNatGatewayStatusResponseBodyDataZoneCidrs {
	s.SnatSourceCidr = &v
	return s
}

func (s *GetNatGatewayStatusResponseBodyDataZoneCidrs) SetVSwitchId(v string) *GetNatGatewayStatusResponseBodyDataZoneCidrs {
	s.VSwitchId = &v
	return s
}

func (s *GetNatGatewayStatusResponseBodyDataZoneCidrs) SetZoneId(v string) *GetNatGatewayStatusResponseBodyDataZoneCidrs {
	s.ZoneId = &v
	return s
}

func (s *GetNatGatewayStatusResponseBodyDataZoneCidrs) Validate() error {
	return dara.Validate(s)
}
