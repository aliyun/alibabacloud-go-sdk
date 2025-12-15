// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceNetExpireTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstanceNetExpireTimeResponseBody
	GetInstanceId() *string
	SetNetInfoItems(v *ModifyInstanceNetExpireTimeResponseBodyNetInfoItems) *ModifyInstanceNetExpireTimeResponseBody
	GetNetInfoItems() *ModifyInstanceNetExpireTimeResponseBodyNetInfoItems
	SetRequestId(v string) *ModifyInstanceNetExpireTimeResponseBody
	GetRequestId() *string
}

type ModifyInstanceNetExpireTimeResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Details about the extension period for which the classic network endpoint of the instance is retained.
	NetInfoItems *ModifyInstanceNetExpireTimeResponseBodyNetInfoItems `json:"NetInfoItems,omitempty" xml:"NetInfoItems,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9C4AF387-1EA3-4C84-8013-3F6B973EDDF5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceNetExpireTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceNetExpireTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNetExpireTimeResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceNetExpireTimeResponseBody) GetNetInfoItems() *ModifyInstanceNetExpireTimeResponseBodyNetInfoItems {
	return s.NetInfoItems
}

func (s *ModifyInstanceNetExpireTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceNetExpireTimeResponseBody) SetInstanceId(v string) *ModifyInstanceNetExpireTimeResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeResponseBody) SetNetInfoItems(v *ModifyInstanceNetExpireTimeResponseBodyNetInfoItems) *ModifyInstanceNetExpireTimeResponseBody {
	s.NetInfoItems = v
	return s
}

func (s *ModifyInstanceNetExpireTimeResponseBody) SetRequestId(v string) *ModifyInstanceNetExpireTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeResponseBody) Validate() error {
	if s.NetInfoItems != nil {
		if err := s.NetInfoItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyInstanceNetExpireTimeResponseBodyNetInfoItems struct {
	NetInfoItem []*ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem `json:"NetInfoItem,omitempty" xml:"NetInfoItem,omitempty" type:"Repeated"`
}

func (s ModifyInstanceNetExpireTimeResponseBodyNetInfoItems) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceNetExpireTimeResponseBodyNetInfoItems) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNetExpireTimeResponseBodyNetInfoItems) GetNetInfoItem() []*ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem {
	return s.NetInfoItem
}

func (s *ModifyInstanceNetExpireTimeResponseBodyNetInfoItems) SetNetInfoItem(v []*ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) *ModifyInstanceNetExpireTimeResponseBodyNetInfoItems {
	s.NetInfoItem = v
	return s
}

func (s *ModifyInstanceNetExpireTimeResponseBodyNetInfoItems) Validate() error {
	if s.NetInfoItem != nil {
		for _, item := range s.NetInfoItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem struct {
	// The endpoint of the classic network.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****.redis.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The network type of the instance. The returned value is **Classic**.
	//
	// example:
	//
	// Classic
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// The expiration time of the classic network endpoint.
	//
	// example:
	//
	// 2019-08-01T09:29:18Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The IP address of the instance in the classic network.
	//
	// example:
	//
	// 100.118.142.***
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The port number that is used to connect to the instance.
	//
	// example:
	//
	// 6379
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) GetDBInstanceNetType() *string {
	return s.DBInstanceNetType
}

func (s *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) GetIPAddress() *string {
	return s.IPAddress
}

func (s *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) GetPort() *string {
	return s.Port
}

func (s *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) SetConnectionString(v string) *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem {
	s.ConnectionString = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) SetDBInstanceNetType(v string) *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem {
	s.DBInstanceNetType = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) SetExpiredTime(v string) *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem {
	s.ExpiredTime = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) SetIPAddress(v string) *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem {
	s.IPAddress = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) SetPort(v string) *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem {
	s.Port = &v
	return s
}

func (s *ModifyInstanceNetExpireTimeResponseBodyNetInfoItemsNetInfoItem) Validate() error {
	return dara.Validate(s)
}
