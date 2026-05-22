// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKvNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCapacity(v int64) *GetKvNamespaceResponseBody
	GetCapacity() *int64
	SetCapacityString(v string) *GetKvNamespaceResponseBody
	GetCapacityString() *string
	SetCapacityUsed(v int64) *GetKvNamespaceResponseBody
	GetCapacityUsed() *int64
	SetCapacityUsedString(v string) *GetKvNamespaceResponseBody
	GetCapacityUsedString() *string
	SetDescription(v string) *GetKvNamespaceResponseBody
	GetDescription() *string
	SetNamespace(v string) *GetKvNamespaceResponseBody
	GetNamespace() *string
	SetNamespaceId(v string) *GetKvNamespaceResponseBody
	GetNamespaceId() *string
	SetRequestId(v string) *GetKvNamespaceResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetKvNamespaceResponseBody
	GetStatus() *string
}

type GetKvNamespaceResponseBody struct {
	Capacity           *int64  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	CapacityString     *string `json:"CapacityString,omitempty" xml:"CapacityString,omitempty"`
	CapacityUsed       *int64  `json:"CapacityUsed,omitempty" xml:"CapacityUsed,omitempty"`
	CapacityUsedString *string `json:"CapacityUsedString,omitempty" xml:"CapacityUsedString,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Namespace          *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceId        *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetKvNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKvNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetKvNamespaceResponseBody) GetCapacity() *int64 {
	return s.Capacity
}

func (s *GetKvNamespaceResponseBody) GetCapacityString() *string {
	return s.CapacityString
}

func (s *GetKvNamespaceResponseBody) GetCapacityUsed() *int64 {
	return s.CapacityUsed
}

func (s *GetKvNamespaceResponseBody) GetCapacityUsedString() *string {
	return s.CapacityUsedString
}

func (s *GetKvNamespaceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetKvNamespaceResponseBody) GetNamespace() *string {
	return s.Namespace
}

func (s *GetKvNamespaceResponseBody) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetKvNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKvNamespaceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetKvNamespaceResponseBody) SetCapacity(v int64) *GetKvNamespaceResponseBody {
	s.Capacity = &v
	return s
}

func (s *GetKvNamespaceResponseBody) SetCapacityString(v string) *GetKvNamespaceResponseBody {
	s.CapacityString = &v
	return s
}

func (s *GetKvNamespaceResponseBody) SetCapacityUsed(v int64) *GetKvNamespaceResponseBody {
	s.CapacityUsed = &v
	return s
}

func (s *GetKvNamespaceResponseBody) SetCapacityUsedString(v string) *GetKvNamespaceResponseBody {
	s.CapacityUsedString = &v
	return s
}

func (s *GetKvNamespaceResponseBody) SetDescription(v string) *GetKvNamespaceResponseBody {
	s.Description = &v
	return s
}

func (s *GetKvNamespaceResponseBody) SetNamespace(v string) *GetKvNamespaceResponseBody {
	s.Namespace = &v
	return s
}

func (s *GetKvNamespaceResponseBody) SetNamespaceId(v string) *GetKvNamespaceResponseBody {
	s.NamespaceId = &v
	return s
}

func (s *GetKvNamespaceResponseBody) SetRequestId(v string) *GetKvNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKvNamespaceResponseBody) SetStatus(v string) *GetKvNamespaceResponseBody {
	s.Status = &v
	return s
}

func (s *GetKvNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
