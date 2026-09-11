// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKVCacheStoreAvailableHpnZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKvcsIds(v []*string) *ListKVCacheStoreAvailableHpnZonesRequest
	GetKvcsIds() []*string
	SetRegionId(v string) *ListKVCacheStoreAvailableHpnZonesRequest
	GetRegionId() *string
}

type ListKVCacheStoreAvailableHpnZonesRequest struct {
	// The list of KVCacheStore instance IDs, separated by commas. You can specify a minimum of 1 and a maximum of 100 instance IDs.
	//
	// This parameter is required.
	KvcsIds []*string `json:"KvcsIds,omitempty" xml:"KvcsIds,omitempty" type:"Repeated"`
	// The region ID, such as cn-hangzhou.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListKVCacheStoreAvailableHpnZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKVCacheStoreAvailableHpnZonesRequest) GoString() string {
	return s.String()
}

func (s *ListKVCacheStoreAvailableHpnZonesRequest) GetKvcsIds() []*string {
	return s.KvcsIds
}

func (s *ListKVCacheStoreAvailableHpnZonesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListKVCacheStoreAvailableHpnZonesRequest) SetKvcsIds(v []*string) *ListKVCacheStoreAvailableHpnZonesRequest {
	s.KvcsIds = v
	return s
}

func (s *ListKVCacheStoreAvailableHpnZonesRequest) SetRegionId(v string) *ListKVCacheStoreAvailableHpnZonesRequest {
	s.RegionId = &v
	return s
}

func (s *ListKVCacheStoreAvailableHpnZonesRequest) Validate() error {
	return dara.Validate(s)
}
