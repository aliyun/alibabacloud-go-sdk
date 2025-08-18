// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKvAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCapacity(v int64) *GetKvAccountResponseBody
	GetCapacity() *int64
	SetCapacityString(v string) *GetKvAccountResponseBody
	GetCapacityString() *string
	SetCapacityUsed(v int64) *GetKvAccountResponseBody
	GetCapacityUsed() *int64
	SetCapacityUsedString(v string) *GetKvAccountResponseBody
	GetCapacityUsedString() *string
	SetNamespaceList(v []*GetKvAccountResponseBodyNamespaceList) *GetKvAccountResponseBody
	GetNamespaceList() []*GetKvAccountResponseBodyNamespaceList
	SetNamespaceQuota(v int32) *GetKvAccountResponseBody
	GetNamespaceQuota() *int32
	SetNamespaceUsed(v int32) *GetKvAccountResponseBody
	GetNamespaceUsed() *int32
	SetRequestId(v string) *GetKvAccountResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetKvAccountResponseBody
	GetStatus() *string
}

type GetKvAccountResponseBody struct {
	// The available capacity of all namespaces in the account. Unit: bytes.
	//
	// example:
	//
	// 1073741824
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The available capacity of all namespaces in the account.
	//
	// example:
	//
	// 1 GB
	CapacityString *string `json:"CapacityString,omitempty" xml:"CapacityString,omitempty"`
	// The used capacity of all namespaces in the account. Unit: bytes.
	//
	// example:
	//
	// 10048576
	CapacityUsed *int64 `json:"CapacityUsed,omitempty" xml:"CapacityUsed,omitempty"`
	// The used capacity of all namespaces in the account.
	//
	// example:
	//
	// 100 MB
	CapacityUsedString *string `json:"CapacityUsedString,omitempty" xml:"CapacityUsedString,omitempty"`
	// The namespaces in the account.
	NamespaceList []*GetKvAccountResponseBodyNamespaceList `json:"NamespaceList,omitempty" xml:"NamespaceList,omitempty" type:"Repeated"`
	// The maximum number of namespaces that can be created in the account.
	//
	// example:
	//
	// 10
	NamespaceQuota *int32 `json:"NamespaceQuota,omitempty" xml:"NamespaceQuota,omitempty"`
	// The number of created namespaces that store key-value data in the account.
	//
	// example:
	//
	// 1
	NamespaceUsed *int32 `json:"NamespaceUsed,omitempty" xml:"NamespaceUsed,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether Edge KV is activated for the Alibaba Cloud account.
	//
	// 	- **online**
	//
	// 	- **offline**
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetKvAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKvAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetKvAccountResponseBody) GetCapacity() *int64 {
	return s.Capacity
}

func (s *GetKvAccountResponseBody) GetCapacityString() *string {
	return s.CapacityString
}

func (s *GetKvAccountResponseBody) GetCapacityUsed() *int64 {
	return s.CapacityUsed
}

func (s *GetKvAccountResponseBody) GetCapacityUsedString() *string {
	return s.CapacityUsedString
}

func (s *GetKvAccountResponseBody) GetNamespaceList() []*GetKvAccountResponseBodyNamespaceList {
	return s.NamespaceList
}

func (s *GetKvAccountResponseBody) GetNamespaceQuota() *int32 {
	return s.NamespaceQuota
}

func (s *GetKvAccountResponseBody) GetNamespaceUsed() *int32 {
	return s.NamespaceUsed
}

func (s *GetKvAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKvAccountResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetKvAccountResponseBody) SetCapacity(v int64) *GetKvAccountResponseBody {
	s.Capacity = &v
	return s
}

func (s *GetKvAccountResponseBody) SetCapacityString(v string) *GetKvAccountResponseBody {
	s.CapacityString = &v
	return s
}

func (s *GetKvAccountResponseBody) SetCapacityUsed(v int64) *GetKvAccountResponseBody {
	s.CapacityUsed = &v
	return s
}

func (s *GetKvAccountResponseBody) SetCapacityUsedString(v string) *GetKvAccountResponseBody {
	s.CapacityUsedString = &v
	return s
}

func (s *GetKvAccountResponseBody) SetNamespaceList(v []*GetKvAccountResponseBodyNamespaceList) *GetKvAccountResponseBody {
	s.NamespaceList = v
	return s
}

func (s *GetKvAccountResponseBody) SetNamespaceQuota(v int32) *GetKvAccountResponseBody {
	s.NamespaceQuota = &v
	return s
}

func (s *GetKvAccountResponseBody) SetNamespaceUsed(v int32) *GetKvAccountResponseBody {
	s.NamespaceUsed = &v
	return s
}

func (s *GetKvAccountResponseBody) SetRequestId(v string) *GetKvAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKvAccountResponseBody) SetStatus(v string) *GetKvAccountResponseBody {
	s.Status = &v
	return s
}

func (s *GetKvAccountResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetKvAccountResponseBodyNamespaceList struct {
	// The available capacity of the namespace. Unit: bytes.
	//
	// example:
	//
	// 1073741824
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The available capacity of the namespace.
	//
	// example:
	//
	// 1 GB
	CapacityString *string `json:"CapacityString,omitempty" xml:"CapacityString,omitempty"`
	// The used capacity of the namespace. Unit: bytes.
	//
	// example:
	//
	// 100048576
	CapacityUsed *int64 `json:"CapacityUsed,omitempty" xml:"CapacityUsed,omitempty"`
	// The used capacity of the namespace.
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
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// 643355322374688768
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The status of the namespace. Valid values:
	//
	// 	- **online**: normal.
	//
	// 	- **delete**: pending deletion.
	//
	// 	- **deleting**: being deleted.
	//
	// 	- **deleted**: deleted.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetKvAccountResponseBodyNamespaceList) String() string {
	return dara.Prettify(s)
}

func (s GetKvAccountResponseBodyNamespaceList) GoString() string {
	return s.String()
}

func (s *GetKvAccountResponseBodyNamespaceList) GetCapacity() *int64 {
	return s.Capacity
}

func (s *GetKvAccountResponseBodyNamespaceList) GetCapacityString() *string {
	return s.CapacityString
}

func (s *GetKvAccountResponseBodyNamespaceList) GetCapacityUsed() *int64 {
	return s.CapacityUsed
}

func (s *GetKvAccountResponseBodyNamespaceList) GetCapacityUsedString() *string {
	return s.CapacityUsedString
}

func (s *GetKvAccountResponseBodyNamespaceList) GetDescription() *string {
	return s.Description
}

func (s *GetKvAccountResponseBodyNamespaceList) GetNamespace() *string {
	return s.Namespace
}

func (s *GetKvAccountResponseBodyNamespaceList) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetKvAccountResponseBodyNamespaceList) GetStatus() *string {
	return s.Status
}

func (s *GetKvAccountResponseBodyNamespaceList) SetCapacity(v int64) *GetKvAccountResponseBodyNamespaceList {
	s.Capacity = &v
	return s
}

func (s *GetKvAccountResponseBodyNamespaceList) SetCapacityString(v string) *GetKvAccountResponseBodyNamespaceList {
	s.CapacityString = &v
	return s
}

func (s *GetKvAccountResponseBodyNamespaceList) SetCapacityUsed(v int64) *GetKvAccountResponseBodyNamespaceList {
	s.CapacityUsed = &v
	return s
}

func (s *GetKvAccountResponseBodyNamespaceList) SetCapacityUsedString(v string) *GetKvAccountResponseBodyNamespaceList {
	s.CapacityUsedString = &v
	return s
}

func (s *GetKvAccountResponseBodyNamespaceList) SetDescription(v string) *GetKvAccountResponseBodyNamespaceList {
	s.Description = &v
	return s
}

func (s *GetKvAccountResponseBodyNamespaceList) SetNamespace(v string) *GetKvAccountResponseBodyNamespaceList {
	s.Namespace = &v
	return s
}

func (s *GetKvAccountResponseBodyNamespaceList) SetNamespaceId(v string) *GetKvAccountResponseBodyNamespaceList {
	s.NamespaceId = &v
	return s
}

func (s *GetKvAccountResponseBodyNamespaceList) SetStatus(v string) *GetKvAccountResponseBodyNamespaceList {
	s.Status = &v
	return s
}

func (s *GetKvAccountResponseBodyNamespaceList) Validate() error {
	return dara.Validate(s)
}
