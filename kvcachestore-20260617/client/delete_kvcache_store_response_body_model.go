// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKVCacheStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKvcsId(v string) *DeleteKVCacheStoreResponseBody
	GetKvcsId() *string
	SetRequestId(v string) *DeleteKVCacheStoreResponseBody
	GetRequestId() *string
}

type DeleteKVCacheStoreResponseBody struct {
	// KVCacheStore KvcsId
	//
	// example:
	//
	// kvcs-your-id
	KvcsId *string `json:"KvcsId,omitempty" xml:"KvcsId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B127704C-ECB1-5B0A-AA9C-8F394A6F179F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteKVCacheStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKVCacheStoreResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKVCacheStoreResponseBody) GetKvcsId() *string {
	return s.KvcsId
}

func (s *DeleteKVCacheStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKVCacheStoreResponseBody) SetKvcsId(v string) *DeleteKVCacheStoreResponseBody {
	s.KvcsId = &v
	return s
}

func (s *DeleteKVCacheStoreResponseBody) SetRequestId(v string) *DeleteKVCacheStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKVCacheStoreResponseBody) Validate() error {
	return dara.Validate(s)
}
