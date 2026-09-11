// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKVCacheStoreAvailableHpnZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceHpnZones(v []*ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZones) *ListKVCacheStoreAvailableHpnZonesResponseBody
	GetInstanceHpnZones() []*ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZones
	SetRequestId(v string) *ListKVCacheStoreAvailableHpnZonesResponseBody
	GetRequestId() *string
}

type ListKVCacheStoreAvailableHpnZonesResponseBody struct {
	// The list of available HpnZone information for instances.
	InstanceHpnZones []*ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZones `json:"InstanceHpnZones,omitempty" xml:"InstanceHpnZones,omitempty" type:"Repeated"`
	// The request ID. A request ID is returned regardless of whether the call is successful.
	//
	// example:
	//
	// 019FB5E9-F9E8-52F5-9C56-2CDF479CBEB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListKVCacheStoreAvailableHpnZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKVCacheStoreAvailableHpnZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListKVCacheStoreAvailableHpnZonesResponseBody) GetInstanceHpnZones() []*ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZones {
	return s.InstanceHpnZones
}

func (s *ListKVCacheStoreAvailableHpnZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKVCacheStoreAvailableHpnZonesResponseBody) SetInstanceHpnZones(v []*ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZones) *ListKVCacheStoreAvailableHpnZonesResponseBody {
	s.InstanceHpnZones = v
	return s
}

func (s *ListKVCacheStoreAvailableHpnZonesResponseBody) SetRequestId(v string) *ListKVCacheStoreAvailableHpnZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKVCacheStoreAvailableHpnZonesResponseBody) Validate() error {
	if s.InstanceHpnZones != nil {
		for _, item := range s.InstanceHpnZones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZones struct {
	// The list of available HpnZones for the instance.
	AvailableHpnZones []*ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZonesAvailableHpnZones `json:"AvailableHpnZones,omitempty" xml:"AvailableHpnZones,omitempty" type:"Repeated"`
	// The KVCacheStore instance ID.
	//
	// example:
	//
	// kvcs-xxxxx
	KvcsId *string `json:"KvcsId,omitempty" xml:"KvcsId,omitempty"`
	// The zone ID of the instance.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZones) String() string {
	return dara.Prettify(s)
}

func (s ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZones) GoString() string {
	return s.String()
}

func (s *ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZones) GetAvailableHpnZones() []*ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZonesAvailableHpnZones {
	return s.AvailableHpnZones
}

func (s *ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZones) GetKvcsId() *string {
	return s.KvcsId
}

func (s *ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZones) SetAvailableHpnZones(v []*ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZonesAvailableHpnZones) *ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZones {
	s.AvailableHpnZones = v
	return s
}

func (s *ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZones) SetKvcsId(v string) *ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZones {
	s.KvcsId = &v
	return s
}

func (s *ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZones) SetZoneId(v string) *ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZones {
	s.ZoneId = &v
	return s
}

func (s *ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZones) Validate() error {
	if s.AvailableHpnZones != nil {
		for _, item := range s.AvailableHpnZones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZonesAvailableHpnZones struct {
	// The cluster ID, such as default.
	//
	// example:
	//
	// default
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
}

func (s ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZonesAvailableHpnZones) String() string {
	return dara.Prettify(s)
}

func (s ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZonesAvailableHpnZones) GoString() string {
	return s.String()
}

func (s *ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZonesAvailableHpnZones) GetHpnZone() *string {
	return s.HpnZone
}

func (s *ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZonesAvailableHpnZones) SetHpnZone(v string) *ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZonesAvailableHpnZones {
	s.HpnZone = &v
	return s
}

func (s *ListKVCacheStoreAvailableHpnZonesResponseBodyInstanceHpnZonesAvailableHpnZones) Validate() error {
	return dara.Validate(s)
}
