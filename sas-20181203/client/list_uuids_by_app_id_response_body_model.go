// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUuidsByAppIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListUuidsByAppIdResponseBody
	GetCount() *int32
	SetRequestId(v string) *ListUuidsByAppIdResponseBody
	GetRequestId() *string
	SetUuids(v []*string) *ListUuidsByAppIdResponseBody
	GetUuids() []*string
}

type ListUuidsByAppIdResponseBody struct {
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1383B0DB-D5D6-4B0C-9E6B-75939C8E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The UUIDs of SAE instances.
	Uuids []*string `json:"Uuids,omitempty" xml:"Uuids,omitempty" type:"Repeated"`
}

func (s ListUuidsByAppIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUuidsByAppIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListUuidsByAppIdResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListUuidsByAppIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUuidsByAppIdResponseBody) GetUuids() []*string {
	return s.Uuids
}

func (s *ListUuidsByAppIdResponseBody) SetCount(v int32) *ListUuidsByAppIdResponseBody {
	s.Count = &v
	return s
}

func (s *ListUuidsByAppIdResponseBody) SetRequestId(v string) *ListUuidsByAppIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUuidsByAppIdResponseBody) SetUuids(v []*string) *ListUuidsByAppIdResponseBody {
	s.Uuids = v
	return s
}

func (s *ListUuidsByAppIdResponseBody) Validate() error {
	return dara.Validate(s)
}
