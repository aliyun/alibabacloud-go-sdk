// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceIndicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *ListInstanceIndicesRequest
	GetAll() *bool
	SetIsManaged(v bool) *ListInstanceIndicesRequest
	GetIsManaged() *bool
	SetIsOpenstore(v bool) *ListInstanceIndicesRequest
	GetIsOpenstore() *bool
	SetName(v string) *ListInstanceIndicesRequest
	GetName() *string
	SetPage(v int32) *ListInstanceIndicesRequest
	GetPage() *int32
	SetSize(v int32) *ListInstanceIndicesRequest
	GetSize() *int32
}

type ListInstanceIndicesRequest struct {
	// false
	//
	// example:
	//
	// false
	All *bool `json:"all,omitempty" xml:"all,omitempty"`
	// 15
	//
	// example:
	//
	// false
	IsManaged *bool `json:"isManaged,omitempty" xml:"isManaged,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// false
	IsOpenstore *bool `json:"isOpenstore,omitempty" xml:"isOpenstore,omitempty"`
	// 1
	//
	// example:
	//
	// log-0001
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The header of the response.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The total size of the index in Cloud Hosting. Unit: bytes.
	//
	// example:
	//
	// 15
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListInstanceIndicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceIndicesRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceIndicesRequest) GetAll() *bool {
	return s.All
}

func (s *ListInstanceIndicesRequest) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *ListInstanceIndicesRequest) GetIsOpenstore() *bool {
	return s.IsOpenstore
}

func (s *ListInstanceIndicesRequest) GetName() *string {
	return s.Name
}

func (s *ListInstanceIndicesRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListInstanceIndicesRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListInstanceIndicesRequest) SetAll(v bool) *ListInstanceIndicesRequest {
	s.All = &v
	return s
}

func (s *ListInstanceIndicesRequest) SetIsManaged(v bool) *ListInstanceIndicesRequest {
	s.IsManaged = &v
	return s
}

func (s *ListInstanceIndicesRequest) SetIsOpenstore(v bool) *ListInstanceIndicesRequest {
	s.IsOpenstore = &v
	return s
}

func (s *ListInstanceIndicesRequest) SetName(v string) *ListInstanceIndicesRequest {
	s.Name = &v
	return s
}

func (s *ListInstanceIndicesRequest) SetPage(v int32) *ListInstanceIndicesRequest {
	s.Page = &v
	return s
}

func (s *ListInstanceIndicesRequest) SetSize(v int32) *ListInstanceIndicesRequest {
	s.Size = &v
	return s
}

func (s *ListInstanceIndicesRequest) Validate() error {
	return dara.Validate(s)
}
