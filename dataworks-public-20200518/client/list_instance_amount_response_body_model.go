// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceAmountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceCounts(v []*ListInstanceAmountResponseBodyInstanceCounts) *ListInstanceAmountResponseBody
	GetInstanceCounts() []*ListInstanceAmountResponseBodyInstanceCounts
	SetRequestId(v string) *ListInstanceAmountResponseBody
	GetRequestId() *string
}

type ListInstanceAmountResponseBody struct {
	// The trend of the number of auto triggered node instances within the specified period of time.
	InstanceCounts []*ListInstanceAmountResponseBodyInstanceCounts `json:"InstanceCounts,omitempty" xml:"InstanceCounts,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 95279527adhfj****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstanceAmountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAmountResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceAmountResponseBody) GetInstanceCounts() []*ListInstanceAmountResponseBodyInstanceCounts {
	return s.InstanceCounts
}

func (s *ListInstanceAmountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceAmountResponseBody) SetInstanceCounts(v []*ListInstanceAmountResponseBodyInstanceCounts) *ListInstanceAmountResponseBody {
	s.InstanceCounts = v
	return s
}

func (s *ListInstanceAmountResponseBody) SetRequestId(v string) *ListInstanceAmountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceAmountResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstanceAmountResponseBodyInstanceCounts struct {
	// The number of auto triggered node instances.
	//
	// example:
	//
	// 9527
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The data timestamp at which the number of auto triggered node instances was obtained. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1623772800000
	Date *int64 `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s ListInstanceAmountResponseBodyInstanceCounts) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAmountResponseBodyInstanceCounts) GoString() string {
	return s.String()
}

func (s *ListInstanceAmountResponseBodyInstanceCounts) GetCount() *int32 {
	return s.Count
}

func (s *ListInstanceAmountResponseBodyInstanceCounts) GetDate() *int64 {
	return s.Date
}

func (s *ListInstanceAmountResponseBodyInstanceCounts) SetCount(v int32) *ListInstanceAmountResponseBodyInstanceCounts {
	s.Count = &v
	return s
}

func (s *ListInstanceAmountResponseBodyInstanceCounts) SetDate(v int64) *ListInstanceAmountResponseBodyInstanceCounts {
	s.Date = &v
	return s
}

func (s *ListInstanceAmountResponseBodyInstanceCounts) Validate() error {
	return dara.Validate(s)
}
