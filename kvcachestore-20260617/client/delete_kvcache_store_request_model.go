// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKVCacheStoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKvcsId(v string) *DeleteKVCacheStoreRequest
	GetKvcsId() *string
	SetRegionId(v string) *DeleteKVCacheStoreRequest
	GetRegionId() *string
}

type DeleteKVCacheStoreRequest struct {
	// The KvCacheStore instance ID. Only instances in the following states can be deleted: Available and Stopped. Instances in other states cannot be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// kvcs-your-id
	KvcsId *string `json:"KvcsId,omitempty" xml:"KvcsId,omitempty"`
	// The region ID, such as cn-hangzhou.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteKVCacheStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKVCacheStoreRequest) GoString() string {
	return s.String()
}

func (s *DeleteKVCacheStoreRequest) GetKvcsId() *string {
	return s.KvcsId
}

func (s *DeleteKVCacheStoreRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteKVCacheStoreRequest) SetKvcsId(v string) *DeleteKVCacheStoreRequest {
	s.KvcsId = &v
	return s
}

func (s *DeleteKVCacheStoreRequest) SetRegionId(v string) *DeleteKVCacheStoreRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteKVCacheStoreRequest) Validate() error {
	return dara.Validate(s)
}
