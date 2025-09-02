// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceCountTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceCounts(v []*GetInstanceCountTrendResponseBodyInstanceCounts) *GetInstanceCountTrendResponseBody
	GetInstanceCounts() []*GetInstanceCountTrendResponseBodyInstanceCounts
	SetRequestId(v string) *GetInstanceCountTrendResponseBody
	GetRequestId() *string
}

type GetInstanceCountTrendResponseBody struct {
	// The quantity trend of instances.
	InstanceCounts []*GetInstanceCountTrendResponseBodyInstanceCounts `json:"InstanceCounts,omitempty" xml:"InstanceCounts,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 95279527adhfj****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceCountTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceCountTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceCountTrendResponseBody) GetInstanceCounts() []*GetInstanceCountTrendResponseBodyInstanceCounts {
	return s.InstanceCounts
}

func (s *GetInstanceCountTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceCountTrendResponseBody) SetInstanceCounts(v []*GetInstanceCountTrendResponseBodyInstanceCounts) *GetInstanceCountTrendResponseBody {
	s.InstanceCounts = v
	return s
}

func (s *GetInstanceCountTrendResponseBody) SetRequestId(v string) *GetInstanceCountTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceCountTrendResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceCountTrendResponseBodyInstanceCounts struct {
	// The number of instances.
	//
	// example:
	//
	// 9527
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The data timestamp.
	//
	// example:
	//
	// 1600963200000
	Date *int64 `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s GetInstanceCountTrendResponseBodyInstanceCounts) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceCountTrendResponseBodyInstanceCounts) GoString() string {
	return s.String()
}

func (s *GetInstanceCountTrendResponseBodyInstanceCounts) GetCount() *int32 {
	return s.Count
}

func (s *GetInstanceCountTrendResponseBodyInstanceCounts) GetDate() *int64 {
	return s.Date
}

func (s *GetInstanceCountTrendResponseBodyInstanceCounts) SetCount(v int32) *GetInstanceCountTrendResponseBodyInstanceCounts {
	s.Count = &v
	return s
}

func (s *GetInstanceCountTrendResponseBodyInstanceCounts) SetDate(v int64) *GetInstanceCountTrendResponseBodyInstanceCounts {
	s.Date = &v
	return s
}

func (s *GetInstanceCountTrendResponseBodyInstanceCounts) Validate() error {
	return dara.Validate(s)
}
