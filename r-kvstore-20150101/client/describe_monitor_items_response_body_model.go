// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMonitorItems(v *DescribeMonitorItemsResponseBodyMonitorItems) *DescribeMonitorItemsResponseBody
	GetMonitorItems() *DescribeMonitorItemsResponseBodyMonitorItems
	SetRequestId(v string) *DescribeMonitorItemsResponseBody
	GetRequestId() *string
}

type DescribeMonitorItemsResponseBody struct {
	// The returned metrics.
	//
	// > 	- **memoryUsage**, **GetQps**, and **PutQps*	- are supported only by Tair instances that use Redis 4.0 or later. **GetQps*	- and **PutQps*	- require the latest minor version. You can upgrade the major version or minor version of the instance as needed. For more information, see [Upgrade the major version](https://help.aliyun.com/document_detail/101764.html) and [Upgrade the minor version](https://help.aliyun.com/document_detail/56450.html).
	//
	// > 	- When you use instances of Redis 2.8, if the **hit_rate*	- metric is not displayed, you must upgrade the minor version of the instance. For more information, see [Upgrade the minor version](https://help.aliyun.com/document_detail/56450.html).
	MonitorItems *DescribeMonitorItemsResponseBodyMonitorItems `json:"MonitorItems,omitempty" xml:"MonitorItems,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 8BEB2618-9517-43F3-A233-E0B34512****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMonitorItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorItemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorItemsResponseBody) GetMonitorItems() *DescribeMonitorItemsResponseBodyMonitorItems {
	return s.MonitorItems
}

func (s *DescribeMonitorItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMonitorItemsResponseBody) SetMonitorItems(v *DescribeMonitorItemsResponseBodyMonitorItems) *DescribeMonitorItemsResponseBody {
	s.MonitorItems = v
	return s
}

func (s *DescribeMonitorItemsResponseBody) SetRequestId(v string) *DescribeMonitorItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitorItemsResponseBody) Validate() error {
	if s.MonitorItems != nil {
		if err := s.MonitorItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMonitorItemsResponseBodyMonitorItems struct {
	KVStoreMonitorItem []*DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem `json:"KVStoreMonitorItem,omitempty" xml:"KVStoreMonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeMonitorItemsResponseBodyMonitorItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorItemsResponseBodyMonitorItems) GoString() string {
	return s.String()
}

func (s *DescribeMonitorItemsResponseBodyMonitorItems) GetKVStoreMonitorItem() []*DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem {
	return s.KVStoreMonitorItem
}

func (s *DescribeMonitorItemsResponseBodyMonitorItems) SetKVStoreMonitorItem(v []*DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem) *DescribeMonitorItemsResponseBodyMonitorItems {
	s.KVStoreMonitorItem = v
	return s
}

func (s *DescribeMonitorItemsResponseBodyMonitorItems) Validate() error {
	if s.KVStoreMonitorItem != nil {
		for _, item := range s.KVStoreMonitorItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem struct {
	// The metric.
	//
	// example:
	//
	// select
	MonitorKey *string `json:"MonitorKey,omitempty" xml:"MonitorKey,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// Counts/s
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem) GetMonitorKey() *string {
	return s.MonitorKey
}

func (s *DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem) GetUnit() *string {
	return s.Unit
}

func (s *DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem) SetMonitorKey(v string) *DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem {
	s.MonitorKey = &v
	return s
}

func (s *DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem) SetUnit(v string) *DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem {
	s.Unit = &v
	return s
}

func (s *DescribeMonitorItemsResponseBodyMonitorItemsKVStoreMonitorItem) Validate() error {
	return dara.Validate(s)
}
