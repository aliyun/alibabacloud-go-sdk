// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeTairKVCacheCustomInstanceDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ResizeTairKVCacheCustomInstanceDiskResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ResizeTairKVCacheCustomInstanceDiskResponseBody
	GetRequestId() *string
}

type ResizeTairKVCacheCustomInstanceDiskResponseBody struct {
	// example:
	//
	// 20905403119****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// ABAF95F6-35C1-4177-AF3A-70969EBD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeTairKVCacheCustomInstanceDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResizeTairKVCacheCustomInstanceDiskResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeTairKVCacheCustomInstanceDiskResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ResizeTairKVCacheCustomInstanceDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResizeTairKVCacheCustomInstanceDiskResponseBody) SetOrderId(v string) *ResizeTairKVCacheCustomInstanceDiskResponseBody {
	s.OrderId = &v
	return s
}

func (s *ResizeTairKVCacheCustomInstanceDiskResponseBody) SetRequestId(v string) *ResizeTairKVCacheCustomInstanceDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResizeTairKVCacheCustomInstanceDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
