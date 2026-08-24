// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKVCacheStoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKvcsId(v string) *GetKVCacheStoreRequest
	GetKvcsId() *string
	SetRegionId(v string) *GetKVCacheStoreRequest
	GetRegionId() *string
}

type GetKVCacheStoreRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// kvcs-your-id
	KvcsId *string `json:"KvcsId,omitempty" xml:"KvcsId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetKVCacheStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKVCacheStoreRequest) GoString() string {
	return s.String()
}

func (s *GetKVCacheStoreRequest) GetKvcsId() *string {
	return s.KvcsId
}

func (s *GetKVCacheStoreRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetKVCacheStoreRequest) SetKvcsId(v string) *GetKVCacheStoreRequest {
	s.KvcsId = &v
	return s
}

func (s *GetKVCacheStoreRequest) SetRegionId(v string) *GetKVCacheStoreRequest {
	s.RegionId = &v
	return s
}

func (s *GetKVCacheStoreRequest) Validate() error {
	return dara.Validate(s)
}
