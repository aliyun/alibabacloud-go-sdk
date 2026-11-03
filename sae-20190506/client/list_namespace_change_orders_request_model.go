// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespaceChangeOrdersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoStatus(v string) *ListNamespaceChangeOrdersRequest
	GetCoStatus() *string
	SetCoType(v string) *ListNamespaceChangeOrdersRequest
	GetCoType() *string
	SetCurrentPage(v int32) *ListNamespaceChangeOrdersRequest
	GetCurrentPage() *int32
	SetKey(v string) *ListNamespaceChangeOrdersRequest
	GetKey() *string
	SetNamespaceId(v string) *ListNamespaceChangeOrdersRequest
	GetNamespaceId() *string
	SetPageSize(v int32) *ListNamespaceChangeOrdersRequest
	GetPageSize() *int32
}

type ListNamespaceChangeOrdersRequest struct {
	// The status of the change order. Valid values:
	//
	// - **0**: preparing
	//
	// - **1**: executing
	//
	// - **2**: succeeded
	//
	// - **3**: failed
	//
	// - **6**: terminated
	//
	// - **10**: failed due to a system exception
	//
	// example:
	//
	// 2
	CoStatus *string `json:"CoStatus,omitempty" xml:"CoStatus,omitempty"`
	// The type of the change order. Valid values:
	//
	// - **CoBatchStartApplication**: Starts applications in batches.
	//
	// - **CoBatchStopApplication**: Stops applications in batches.
	//
	// example:
	//
	// CoBatchStartApplication
	CoType *string `json:"CoType,omitempty" xml:"CoType,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The keyword used for fuzzy matching of the change order description. All change orders whose descriptions contain this **key*	- are returned.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListNamespaceChangeOrdersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNamespaceChangeOrdersRequest) GoString() string {
	return s.String()
}

func (s *ListNamespaceChangeOrdersRequest) GetCoStatus() *string {
	return s.CoStatus
}

func (s *ListNamespaceChangeOrdersRequest) GetCoType() *string {
	return s.CoType
}

func (s *ListNamespaceChangeOrdersRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListNamespaceChangeOrdersRequest) GetKey() *string {
	return s.Key
}

func (s *ListNamespaceChangeOrdersRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListNamespaceChangeOrdersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNamespaceChangeOrdersRequest) SetCoStatus(v string) *ListNamespaceChangeOrdersRequest {
	s.CoStatus = &v
	return s
}

func (s *ListNamespaceChangeOrdersRequest) SetCoType(v string) *ListNamespaceChangeOrdersRequest {
	s.CoType = &v
	return s
}

func (s *ListNamespaceChangeOrdersRequest) SetCurrentPage(v int32) *ListNamespaceChangeOrdersRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListNamespaceChangeOrdersRequest) SetKey(v string) *ListNamespaceChangeOrdersRequest {
	s.Key = &v
	return s
}

func (s *ListNamespaceChangeOrdersRequest) SetNamespaceId(v string) *ListNamespaceChangeOrdersRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListNamespaceChangeOrdersRequest) SetPageSize(v int32) *ListNamespaceChangeOrdersRequest {
	s.PageSize = &v
	return s
}

func (s *ListNamespaceChangeOrdersRequest) Validate() error {
	return dara.Validate(s)
}
