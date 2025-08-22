// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnKvAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCapacity(v int64) *DescribeDcdnKvAccountResponseBody
	GetCapacity() *int64
	SetCapacityString(v string) *DescribeDcdnKvAccountResponseBody
	GetCapacityString() *string
	SetCapacityUsed(v int64) *DescribeDcdnKvAccountResponseBody
	GetCapacityUsed() *int64
	SetCapacityUsedString(v string) *DescribeDcdnKvAccountResponseBody
	GetCapacityUsedString() *string
	SetNamespaceList(v []*DescribeDcdnKvAccountResponseBodyNamespaceList) *DescribeDcdnKvAccountResponseBody
	GetNamespaceList() []*DescribeDcdnKvAccountResponseBodyNamespaceList
	SetNamespaceQuota(v int32) *DescribeDcdnKvAccountResponseBody
	GetNamespaceQuota() *int32
	SetNamespaceUsed(v int32) *DescribeDcdnKvAccountResponseBody
	GetNamespaceUsed() *int32
	SetRequestId(v string) *DescribeDcdnKvAccountResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeDcdnKvAccountResponseBody
	GetStatus() *string
}

type DescribeDcdnKvAccountResponseBody struct {
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The available capacity of all namespaces.
	//
	// example:
	//
	// 2GB
	CapacityString *string `json:"CapacityString,omitempty" xml:"CapacityString,omitempty"`
	CapacityUsed   *int64  `json:"CapacityUsed,omitempty" xml:"CapacityUsed,omitempty"`
	// All namespaces have used capacity.
	//
	// example:
	//
	// 200 MB
	CapacityUsedString *string `json:"CapacityUsedString,omitempty" xml:"CapacityUsedString,omitempty"`
	// Details about the namespaces.
	NamespaceList []*DescribeDcdnKvAccountResponseBodyNamespaceList `json:"NamespaceList,omitempty" xml:"NamespaceList,omitempty" type:"Repeated"`
	// The maximum number of namespaces that you can apply for by using your account.
	//
	// example:
	//
	// 10
	NamespaceQuota *int32 `json:"NamespaceQuota,omitempty" xml:"NamespaceQuota,omitempty"`
	// The number of namespaces that you applied for by using your account.
	//
	// example:
	//
	// 1
	NamespaceUsed *int32 `json:"NamespaceUsed,omitempty" xml:"NamespaceUsed,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D61E4801-EAFF-4A63-AAE1-FBF6CE1CFD1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the account.
	//
	// 	- **online**: enabled
	//
	// 	- **offline**: disabled
	//
	// example:
	//
	// online,offline
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDcdnKvAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnKvAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnKvAccountResponseBody) GetCapacity() *int64 {
	return s.Capacity
}

func (s *DescribeDcdnKvAccountResponseBody) GetCapacityString() *string {
	return s.CapacityString
}

func (s *DescribeDcdnKvAccountResponseBody) GetCapacityUsed() *int64 {
	return s.CapacityUsed
}

func (s *DescribeDcdnKvAccountResponseBody) GetCapacityUsedString() *string {
	return s.CapacityUsedString
}

func (s *DescribeDcdnKvAccountResponseBody) GetNamespaceList() []*DescribeDcdnKvAccountResponseBodyNamespaceList {
	return s.NamespaceList
}

func (s *DescribeDcdnKvAccountResponseBody) GetNamespaceQuota() *int32 {
	return s.NamespaceQuota
}

func (s *DescribeDcdnKvAccountResponseBody) GetNamespaceUsed() *int32 {
	return s.NamespaceUsed
}

func (s *DescribeDcdnKvAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnKvAccountResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeDcdnKvAccountResponseBody) SetCapacity(v int64) *DescribeDcdnKvAccountResponseBody {
	s.Capacity = &v
	return s
}

func (s *DescribeDcdnKvAccountResponseBody) SetCapacityString(v string) *DescribeDcdnKvAccountResponseBody {
	s.CapacityString = &v
	return s
}

func (s *DescribeDcdnKvAccountResponseBody) SetCapacityUsed(v int64) *DescribeDcdnKvAccountResponseBody {
	s.CapacityUsed = &v
	return s
}

func (s *DescribeDcdnKvAccountResponseBody) SetCapacityUsedString(v string) *DescribeDcdnKvAccountResponseBody {
	s.CapacityUsedString = &v
	return s
}

func (s *DescribeDcdnKvAccountResponseBody) SetNamespaceList(v []*DescribeDcdnKvAccountResponseBodyNamespaceList) *DescribeDcdnKvAccountResponseBody {
	s.NamespaceList = v
	return s
}

func (s *DescribeDcdnKvAccountResponseBody) SetNamespaceQuota(v int32) *DescribeDcdnKvAccountResponseBody {
	s.NamespaceQuota = &v
	return s
}

func (s *DescribeDcdnKvAccountResponseBody) SetNamespaceUsed(v int32) *DescribeDcdnKvAccountResponseBody {
	s.NamespaceUsed = &v
	return s
}

func (s *DescribeDcdnKvAccountResponseBody) SetRequestId(v string) *DescribeDcdnKvAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnKvAccountResponseBody) SetStatus(v string) *DescribeDcdnKvAccountResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDcdnKvAccountResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnKvAccountResponseBodyNamespaceList struct {
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The available capacity of the namespace.
	//
	// example:
	//
	// 1 GB
	CapacityString *string `json:"CapacityString,omitempty" xml:"CapacityString,omitempty"`
	CapacityUsed   *int64  `json:"CapacityUsed,omitempty" xml:"CapacityUsed,omitempty"`
	// The namespace has used capacity.
	//
	// example:
	//
	// 100 MB
	CapacityUsedString *string `json:"CapacityUsedString,omitempty" xml:"CapacityUsedString,omitempty"`
	// The description of the namespace.
	//
	// example:
	//
	// the first namespace
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// ns1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// 12423131231****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The status of the namespace. Valid values:
	//
	// 	- **online**: normal
	//
	// 	- **delete**: pending delete
	//
	// 	- **deleting**: being deleted
	//
	// 	- **deleted**: deleted
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDcdnKvAccountResponseBodyNamespaceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnKvAccountResponseBodyNamespaceList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnKvAccountResponseBodyNamespaceList) GetCapacity() *int64 {
	return s.Capacity
}

func (s *DescribeDcdnKvAccountResponseBodyNamespaceList) GetCapacityString() *string {
	return s.CapacityString
}

func (s *DescribeDcdnKvAccountResponseBodyNamespaceList) GetCapacityUsed() *int64 {
	return s.CapacityUsed
}

func (s *DescribeDcdnKvAccountResponseBodyNamespaceList) GetCapacityUsedString() *string {
	return s.CapacityUsedString
}

func (s *DescribeDcdnKvAccountResponseBodyNamespaceList) GetDescription() *string {
	return s.Description
}

func (s *DescribeDcdnKvAccountResponseBodyNamespaceList) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeDcdnKvAccountResponseBodyNamespaceList) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeDcdnKvAccountResponseBodyNamespaceList) GetStatus() *string {
	return s.Status
}

func (s *DescribeDcdnKvAccountResponseBodyNamespaceList) SetCapacity(v int64) *DescribeDcdnKvAccountResponseBodyNamespaceList {
	s.Capacity = &v
	return s
}

func (s *DescribeDcdnKvAccountResponseBodyNamespaceList) SetCapacityString(v string) *DescribeDcdnKvAccountResponseBodyNamespaceList {
	s.CapacityString = &v
	return s
}

func (s *DescribeDcdnKvAccountResponseBodyNamespaceList) SetCapacityUsed(v int64) *DescribeDcdnKvAccountResponseBodyNamespaceList {
	s.CapacityUsed = &v
	return s
}

func (s *DescribeDcdnKvAccountResponseBodyNamespaceList) SetCapacityUsedString(v string) *DescribeDcdnKvAccountResponseBodyNamespaceList {
	s.CapacityUsedString = &v
	return s
}

func (s *DescribeDcdnKvAccountResponseBodyNamespaceList) SetDescription(v string) *DescribeDcdnKvAccountResponseBodyNamespaceList {
	s.Description = &v
	return s
}

func (s *DescribeDcdnKvAccountResponseBodyNamespaceList) SetNamespace(v string) *DescribeDcdnKvAccountResponseBodyNamespaceList {
	s.Namespace = &v
	return s
}

func (s *DescribeDcdnKvAccountResponseBodyNamespaceList) SetNamespaceId(v string) *DescribeDcdnKvAccountResponseBodyNamespaceList {
	s.NamespaceId = &v
	return s
}

func (s *DescribeDcdnKvAccountResponseBodyNamespaceList) SetStatus(v string) *DescribeDcdnKvAccountResponseBodyNamespaceList {
	s.Status = &v
	return s
}

func (s *DescribeDcdnKvAccountResponseBodyNamespaceList) Validate() error {
	return dara.Validate(s)
}
