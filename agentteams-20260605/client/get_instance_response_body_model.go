// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceResponseBody
	GetCode() *string
	SetData(v *GetInstanceResponseBodyData) *GetInstanceResponseBody
	GetData() *GetInstanceResponseBodyData
	SetMessage(v string) *GetInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceResponseBody
	GetSuccess() *bool
}

type GetInstanceResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// request-1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceResponseBody) GetData() *GetInstanceResponseBodyData {
	return s.Data
}

func (s *GetInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceResponseBody) SetCode(v string) *GetInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceResponseBody) SetData(v *GetInstanceResponseBodyData) *GetInstanceResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceResponseBody) SetMessage(v string) *GetInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetSuccess(v bool) *GetInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceResponseBodyData struct {
	CreateTime    *string                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceId    *string                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName  *string                             `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceSpec  *string                             `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	NetworkType   *string                             `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OssBucketName *string                             `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	PaymentType   *string                             `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	RegionId      *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroup *string                             `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	Status        *string                             `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime    *string                             `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VpcId         *string                             `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Zones         []*GetInstanceResponseBodyDataZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s GetInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetInstanceResponseBodyData) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *GetInstanceResponseBodyData) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetInstanceResponseBodyData) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *GetInstanceResponseBodyData) GetPaymentType() *string {
	return s.PaymentType
}

func (s *GetInstanceResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceResponseBodyData) GetSecurityGroup() *string {
	return s.SecurityGroup
}

func (s *GetInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetInstanceResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *GetInstanceResponseBodyData) GetZones() []*GetInstanceResponseBodyDataZones {
	return s.Zones
}

func (s *GetInstanceResponseBodyData) SetCreateTime(v string) *GetInstanceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetInstanceId(v string) *GetInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetInstanceName(v string) *GetInstanceResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetInstanceSpec(v string) *GetInstanceResponseBodyData {
	s.InstanceSpec = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetNetworkType(v string) *GetInstanceResponseBodyData {
	s.NetworkType = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetOssBucketName(v string) *GetInstanceResponseBodyData {
	s.OssBucketName = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetPaymentType(v string) *GetInstanceResponseBodyData {
	s.PaymentType = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetRegionId(v string) *GetInstanceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetSecurityGroup(v string) *GetInstanceResponseBodyData {
	s.SecurityGroup = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetStatus(v string) *GetInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetUpdateTime(v string) *GetInstanceResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetVpcId(v string) *GetInstanceResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetZones(v []*GetInstanceResponseBodyDataZones) *GetInstanceResponseBodyData {
	s.Zones = v
	return s
}

func (s *GetInstanceResponseBodyData) Validate() error {
	if s.Zones != nil {
		for _, item := range s.Zones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceResponseBodyDataZones struct {
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetInstanceResponseBodyDataZones) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyDataZones) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataZones) GetVswitchId() *string {
	return s.VswitchId
}

func (s *GetInstanceResponseBodyDataZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetInstanceResponseBodyDataZones) SetVswitchId(v string) *GetInstanceResponseBodyDataZones {
	s.VswitchId = &v
	return s
}

func (s *GetInstanceResponseBodyDataZones) SetZoneId(v string) *GetInstanceResponseBodyDataZones {
	s.ZoneId = &v
	return s
}

func (s *GetInstanceResponseBodyDataZones) Validate() error {
	return dara.Validate(s)
}
