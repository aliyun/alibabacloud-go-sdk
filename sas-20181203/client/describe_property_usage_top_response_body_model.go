// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyUsageTopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItemCount(v int32) *DescribePropertyUsageTopResponseBody
	GetItemCount() *int32
	SetRequestId(v string) *DescribePropertyUsageTopResponseBody
	GetRequestId() *string
	SetTopStatisticItems(v []*DescribePropertyUsageTopResponseBodyTopStatisticItems) *DescribePropertyUsageTopResponseBody
	GetTopStatisticItems() []*DescribePropertyUsageTopResponseBodyTopStatisticItems
	SetType(v string) *DescribePropertyUsageTopResponseBody
	GetType() *string
}

type DescribePropertyUsageTopResponseBody struct {
	// The number of fingerprints.
	//
	// example:
	//
	// 5
	ItemCount *int32 `json:"ItemCount,omitempty" xml:"ItemCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16AA5B62-A3C1-520B-B289-4BD971CC17AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistical results.
	TopStatisticItems []*DescribePropertyUsageTopResponseBodyTopStatisticItems `json:"TopStatisticItems,omitempty" xml:"TopStatisticItems,omitempty" type:"Repeated"`
	// The type of the asset fingerprint. Valid value:
	//
	// 	- **port**: port
	//
	// 	- **process**: process
	//
	// 	- **software**: software
	//
	// 	- **user**: account
	//
	// 	- **sca**: middleware
	//
	// example:
	//
	// sca
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePropertyUsageTopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyUsageTopResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyUsageTopResponseBody) GetItemCount() *int32 {
	return s.ItemCount
}

func (s *DescribePropertyUsageTopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePropertyUsageTopResponseBody) GetTopStatisticItems() []*DescribePropertyUsageTopResponseBodyTopStatisticItems {
	return s.TopStatisticItems
}

func (s *DescribePropertyUsageTopResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribePropertyUsageTopResponseBody) SetItemCount(v int32) *DescribePropertyUsageTopResponseBody {
	s.ItemCount = &v
	return s
}

func (s *DescribePropertyUsageTopResponseBody) SetRequestId(v string) *DescribePropertyUsageTopResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertyUsageTopResponseBody) SetTopStatisticItems(v []*DescribePropertyUsageTopResponseBodyTopStatisticItems) *DescribePropertyUsageTopResponseBody {
	s.TopStatisticItems = v
	return s
}

func (s *DescribePropertyUsageTopResponseBody) SetType(v string) *DescribePropertyUsageTopResponseBody {
	s.Type = &v
	return s
}

func (s *DescribePropertyUsageTopResponseBody) Validate() error {
	if s.TopStatisticItems != nil {
		for _, item := range s.TopStatisticItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePropertyUsageTopResponseBodyTopStatisticItems struct {
	// The quantity.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The statistical item.
	//
	// example:
	//
	// openssl
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribePropertyUsageTopResponseBodyTopStatisticItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyUsageTopResponseBodyTopStatisticItems) GoString() string {
	return s.String()
}

func (s *DescribePropertyUsageTopResponseBodyTopStatisticItems) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertyUsageTopResponseBodyTopStatisticItems) GetName() *string {
	return s.Name
}

func (s *DescribePropertyUsageTopResponseBodyTopStatisticItems) SetCount(v int32) *DescribePropertyUsageTopResponseBodyTopStatisticItems {
	s.Count = &v
	return s
}

func (s *DescribePropertyUsageTopResponseBodyTopStatisticItems) SetName(v string) *DescribePropertyUsageTopResponseBodyTopStatisticItems {
	s.Name = &v
	return s
}

func (s *DescribePropertyUsageTopResponseBodyTopStatisticItems) Validate() error {
	return dara.Validate(s)
}
