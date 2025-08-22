// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDcdnKvNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *PutDcdnKvNamespaceResponseBody
	GetDescription() *string
	SetNamespace(v string) *PutDcdnKvNamespaceResponseBody
	GetNamespace() *string
	SetNamespaceId(v string) *PutDcdnKvNamespaceResponseBody
	GetNamespaceId() *string
	SetRequestId(v string) *PutDcdnKvNamespaceResponseBody
	GetRequestId() *string
	SetStatus(v string) *PutDcdnKvNamespaceResponseBody
	GetStatus() *string
}

type PutDcdnKvNamespaceResponseBody struct {
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

func (s PutDcdnKvNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutDcdnKvNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *PutDcdnKvNamespaceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *PutDcdnKvNamespaceResponseBody) GetNamespace() *string {
	return s.Namespace
}

func (s *PutDcdnKvNamespaceResponseBody) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *PutDcdnKvNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutDcdnKvNamespaceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *PutDcdnKvNamespaceResponseBody) SetDescription(v string) *PutDcdnKvNamespaceResponseBody {
	s.Description = &v
	return s
}

func (s *PutDcdnKvNamespaceResponseBody) SetNamespace(v string) *PutDcdnKvNamespaceResponseBody {
	s.Namespace = &v
	return s
}

func (s *PutDcdnKvNamespaceResponseBody) SetNamespaceId(v string) *PutDcdnKvNamespaceResponseBody {
	s.NamespaceId = &v
	return s
}

func (s *PutDcdnKvNamespaceResponseBody) SetRequestId(v string) *PutDcdnKvNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutDcdnKvNamespaceResponseBody) SetStatus(v string) *PutDcdnKvNamespaceResponseBody {
	s.Status = &v
	return s
}

func (s *PutDcdnKvNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
