// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspiciousUUIDConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeSuspiciousUUIDConfigResponseBody
	GetCount() *int32
	SetRequestId(v string) *DescribeSuspiciousUUIDConfigResponseBody
	GetRequestId() *string
	SetUUIDList(v []*string) *DescribeSuspiciousUUIDConfigResponseBody
	GetUUIDList() []*string
}

type DescribeSuspiciousUUIDConfigResponseBody struct {
	// The total number of servers on which proactive defense of the specified type takes effect.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6044DC07-86F1-5DDA-A611-EC578EA4EEE6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The UUIDs of servers on which proactive defense of the specified type takes effect.
	UUIDList []*string `json:"UUIDList,omitempty" xml:"UUIDList,omitempty" type:"Repeated"`
}

func (s DescribeSuspiciousUUIDConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspiciousUUIDConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSuspiciousUUIDConfigResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeSuspiciousUUIDConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSuspiciousUUIDConfigResponseBody) GetUUIDList() []*string {
	return s.UUIDList
}

func (s *DescribeSuspiciousUUIDConfigResponseBody) SetCount(v int32) *DescribeSuspiciousUUIDConfigResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeSuspiciousUUIDConfigResponseBody) SetRequestId(v string) *DescribeSuspiciousUUIDConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSuspiciousUUIDConfigResponseBody) SetUUIDList(v []*string) *DescribeSuspiciousUUIDConfigResponseBody {
	s.UUIDList = v
	return s
}

func (s *DescribeSuspiciousUUIDConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
