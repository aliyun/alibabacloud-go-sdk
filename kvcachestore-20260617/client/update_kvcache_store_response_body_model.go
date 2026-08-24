// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKVCacheStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKvcsId(v string) *UpdateKVCacheStoreResponseBody
	GetKvcsId() *string
	SetOrderId(v string) *UpdateKVCacheStoreResponseBody
	GetOrderId() *string
	SetRequestId(v string) *UpdateKVCacheStoreResponseBody
	GetRequestId() *string
}

type UpdateKVCacheStoreResponseBody struct {
	// The ID of the modified KVCacheStore.
	//
	// example:
	//
	// kvs-xxxxx
	KvcsId *string `json:"KvcsId,omitempty" xml:"KvcsId,omitempty"`
	// The specification change order ID. Returned only when Capacity is modified.
	//
	// example:
	//
	// order-xxxxx
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID. The request ID is returned regardless of whether the operation is successful.
	//
	// example:
	//
	// 6AA27F1A-A62C-59C3-BCC7-D1DFA4E7EEA0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateKVCacheStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKVCacheStoreResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKVCacheStoreResponseBody) GetKvcsId() *string {
	return s.KvcsId
}

func (s *UpdateKVCacheStoreResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *UpdateKVCacheStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKVCacheStoreResponseBody) SetKvcsId(v string) *UpdateKVCacheStoreResponseBody {
	s.KvcsId = &v
	return s
}

func (s *UpdateKVCacheStoreResponseBody) SetOrderId(v string) *UpdateKVCacheStoreResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpdateKVCacheStoreResponseBody) SetRequestId(v string) *UpdateKVCacheStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKVCacheStoreResponseBody) Validate() error {
	return dara.Validate(s)
}
