// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRescaleDeviceServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceIds(v []*string) *RescaleDeviceServiceResponseBody
	GetDeviceIds() []*string
	SetOrderId(v string) *RescaleDeviceServiceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *RescaleDeviceServiceResponseBody
	GetRequestId() *string
	SetResourceDetailInfos(v []*RescaleDeviceServiceResponseBodyResourceDetailInfos) *RescaleDeviceServiceResponseBody
	GetResourceDetailInfos() []*RescaleDeviceServiceResponseBodyResourceDetailInfos
}

type RescaleDeviceServiceResponseBody struct {
	// The IDs of the devices.
	DeviceIds []*string `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty" type:"Repeated"`
	// The ID of the order.
	//
	// example:
	//
	// b3b5bb9a-4e0b-4cac-8ebf-e5e015726723
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3F3F3570-E721-53F6-853F-37B7725AC6CB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The key properties of the device.
	ResourceDetailInfos []*RescaleDeviceServiceResponseBodyResourceDetailInfos `json:"ResourceDetailInfos,omitempty" xml:"ResourceDetailInfos,omitempty" type:"Repeated"`
}

func (s RescaleDeviceServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RescaleDeviceServiceResponseBody) GoString() string {
	return s.String()
}

func (s *RescaleDeviceServiceResponseBody) GetDeviceIds() []*string {
	return s.DeviceIds
}

func (s *RescaleDeviceServiceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RescaleDeviceServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RescaleDeviceServiceResponseBody) GetResourceDetailInfos() []*RescaleDeviceServiceResponseBodyResourceDetailInfos {
	return s.ResourceDetailInfos
}

func (s *RescaleDeviceServiceResponseBody) SetDeviceIds(v []*string) *RescaleDeviceServiceResponseBody {
	s.DeviceIds = v
	return s
}

func (s *RescaleDeviceServiceResponseBody) SetOrderId(v string) *RescaleDeviceServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *RescaleDeviceServiceResponseBody) SetRequestId(v string) *RescaleDeviceServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RescaleDeviceServiceResponseBody) SetResourceDetailInfos(v []*RescaleDeviceServiceResponseBodyResourceDetailInfos) *RescaleDeviceServiceResponseBody {
	s.ResourceDetailInfos = v
	return s
}

func (s *RescaleDeviceServiceResponseBody) Validate() error {
	if s.ResourceDetailInfos != nil {
		for _, item := range s.ResourceDetailInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RescaleDeviceServiceResponseBodyResourceDetailInfos struct {
	// The ID of the device.
	//
	// example:
	//
	// h-uf6009zoa6hdbjyqxcn1
	ID *string `json:"ID,omitempty" xml:"ID,omitempty"`
	// The IP address of the device.
	//
	// example:
	//
	// 10.152.196.36
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The Internet service provider (ISP) to which the device belongs.
	//
	// example:
	//
	// telecom
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// The media access control (MAC) address of the device.
	//
	// example:
	//
	// 24:0B:88:04:71:E0
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// The ID of the edge node to which the device belongs.
	//
	// example:
	//
	// cn-chongqing-1
	RegionID *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	// The name of the server on which the device is deployed.
	//
	// example:
	//
	// ens-nc2
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// The status of the device.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the instance.
	//
	// example:
	//
	// ens.ac6.large
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RescaleDeviceServiceResponseBodyResourceDetailInfos) String() string {
	return dara.Prettify(s)
}

func (s RescaleDeviceServiceResponseBodyResourceDetailInfos) GoString() string {
	return s.String()
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) GetID() *string {
	return s.ID
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) GetIP() *string {
	return s.IP
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) GetISP() *string {
	return s.ISP
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) GetMac() *string {
	return s.Mac
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) GetRegionID() *string {
	return s.RegionID
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) GetServer() *string {
	return s.Server
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) GetStatus() *string {
	return s.Status
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) GetType() *string {
	return s.Type
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) SetID(v string) *RescaleDeviceServiceResponseBodyResourceDetailInfos {
	s.ID = &v
	return s
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) SetIP(v string) *RescaleDeviceServiceResponseBodyResourceDetailInfos {
	s.IP = &v
	return s
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) SetISP(v string) *RescaleDeviceServiceResponseBodyResourceDetailInfos {
	s.ISP = &v
	return s
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) SetMac(v string) *RescaleDeviceServiceResponseBodyResourceDetailInfos {
	s.Mac = &v
	return s
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) SetRegionID(v string) *RescaleDeviceServiceResponseBodyResourceDetailInfos {
	s.RegionID = &v
	return s
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) SetServer(v string) *RescaleDeviceServiceResponseBodyResourceDetailInfos {
	s.Server = &v
	return s
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) SetStatus(v string) *RescaleDeviceServiceResponseBodyResourceDetailInfos {
	s.Status = &v
	return s
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) SetType(v string) *RescaleDeviceServiceResponseBodyResourceDetailInfos {
	s.Type = &v
	return s
}

func (s *RescaleDeviceServiceResponseBodyResourceDetailInfos) Validate() error {
	return dara.Validate(s)
}
