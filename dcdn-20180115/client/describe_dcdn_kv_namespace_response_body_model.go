// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnKvNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCapacity(v int64) *DescribeDcdnKvNamespaceResponseBody
	GetCapacity() *int64
	SetCapacityString(v string) *DescribeDcdnKvNamespaceResponseBody
	GetCapacityString() *string
	SetCapacityUsed(v int64) *DescribeDcdnKvNamespaceResponseBody
	GetCapacityUsed() *int64
	SetCapacityUsedString(v string) *DescribeDcdnKvNamespaceResponseBody
	GetCapacityUsedString() *string
	SetDescription(v string) *DescribeDcdnKvNamespaceResponseBody
	GetDescription() *string
	SetMode(v string) *DescribeDcdnKvNamespaceResponseBody
	GetMode() *string
	SetNamespace(v string) *DescribeDcdnKvNamespaceResponseBody
	GetNamespace() *string
	SetNamespaceId(v string) *DescribeDcdnKvNamespaceResponseBody
	GetNamespaceId() *string
	SetRequestId(v string) *DescribeDcdnKvNamespaceResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeDcdnKvNamespaceResponseBody
	GetStatus() *string
}

type DescribeDcdnKvNamespaceResponseBody struct {
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The available capacity of all namespaces in your account.
	//
	// example:
	//
	// 1 GB
	CapacityString *string `json:"CapacityString,omitempty" xml:"CapacityString,omitempty"`
	CapacityUsed   *int64  `json:"CapacityUsed,omitempty" xml:"CapacityUsed,omitempty"`
	// The used capacity of all namespaces in your account.
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
	// The system behavior when a key-value pair fails to be obtained at the edge. Valid values:
	//
	// 	- Normal (default): If a key-value pair fails to be obtained at the edge, DCDN attempts to query the key-value pair from the origin server to ensure global data consistency.
	//
	// 	- Rapid: If a key-value pair fails to be obtained at the edge, an error message indicating that the key does not exist is returned. This feature enhances key-value query performance but may decrease the hit rate of queries. To enable this feature, submit a ticket.
	//
	// example:
	//
	// Normal
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
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
	// The ID of the request.
	//
	// example:
	//
	// D61E4801-EAFF-4A63-AAE1-FBF6CE1CFD1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s DescribeDcdnKvNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnKvNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnKvNamespaceResponseBody) GetCapacity() *int64 {
	return s.Capacity
}

func (s *DescribeDcdnKvNamespaceResponseBody) GetCapacityString() *string {
	return s.CapacityString
}

func (s *DescribeDcdnKvNamespaceResponseBody) GetCapacityUsed() *int64 {
	return s.CapacityUsed
}

func (s *DescribeDcdnKvNamespaceResponseBody) GetCapacityUsedString() *string {
	return s.CapacityUsedString
}

func (s *DescribeDcdnKvNamespaceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeDcdnKvNamespaceResponseBody) GetMode() *string {
	return s.Mode
}

func (s *DescribeDcdnKvNamespaceResponseBody) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeDcdnKvNamespaceResponseBody) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeDcdnKvNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnKvNamespaceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeDcdnKvNamespaceResponseBody) SetCapacity(v int64) *DescribeDcdnKvNamespaceResponseBody {
	s.Capacity = &v
	return s
}

func (s *DescribeDcdnKvNamespaceResponseBody) SetCapacityString(v string) *DescribeDcdnKvNamespaceResponseBody {
	s.CapacityString = &v
	return s
}

func (s *DescribeDcdnKvNamespaceResponseBody) SetCapacityUsed(v int64) *DescribeDcdnKvNamespaceResponseBody {
	s.CapacityUsed = &v
	return s
}

func (s *DescribeDcdnKvNamespaceResponseBody) SetCapacityUsedString(v string) *DescribeDcdnKvNamespaceResponseBody {
	s.CapacityUsedString = &v
	return s
}

func (s *DescribeDcdnKvNamespaceResponseBody) SetDescription(v string) *DescribeDcdnKvNamespaceResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeDcdnKvNamespaceResponseBody) SetMode(v string) *DescribeDcdnKvNamespaceResponseBody {
	s.Mode = &v
	return s
}

func (s *DescribeDcdnKvNamespaceResponseBody) SetNamespace(v string) *DescribeDcdnKvNamespaceResponseBody {
	s.Namespace = &v
	return s
}

func (s *DescribeDcdnKvNamespaceResponseBody) SetNamespaceId(v string) *DescribeDcdnKvNamespaceResponseBody {
	s.NamespaceId = &v
	return s
}

func (s *DescribeDcdnKvNamespaceResponseBody) SetRequestId(v string) *DescribeDcdnKvNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnKvNamespaceResponseBody) SetStatus(v string) *DescribeDcdnKvNamespaceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDcdnKvNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
