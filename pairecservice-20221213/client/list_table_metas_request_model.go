// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableMetasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListTableMetasRequest
	GetInstanceId() *string
	SetModule(v string) *ListTableMetasRequest
	GetModule() *string
	SetName(v string) *ListTableMetasRequest
	GetName() *string
	SetPageNumber(v int32) *ListTableMetasRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTableMetasRequest
	GetPageSize() *int32
	SetType(v string) *ListTableMetasRequest
	GetType() *string
}

type ListTableMetasRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ABTest
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// example:
	//
	// test1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// MaxCompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTableMetasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTableMetasRequest) GoString() string {
	return s.String()
}

func (s *ListTableMetasRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTableMetasRequest) GetModule() *string {
	return s.Module
}

func (s *ListTableMetasRequest) GetName() *string {
	return s.Name
}

func (s *ListTableMetasRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTableMetasRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTableMetasRequest) GetType() *string {
	return s.Type
}

func (s *ListTableMetasRequest) SetInstanceId(v string) *ListTableMetasRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTableMetasRequest) SetModule(v string) *ListTableMetasRequest {
	s.Module = &v
	return s
}

func (s *ListTableMetasRequest) SetName(v string) *ListTableMetasRequest {
	s.Name = &v
	return s
}

func (s *ListTableMetasRequest) SetPageNumber(v int32) *ListTableMetasRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTableMetasRequest) SetPageSize(v int32) *ListTableMetasRequest {
	s.PageSize = &v
	return s
}

func (s *ListTableMetasRequest) SetType(v string) *ListTableMetasRequest {
	s.Type = &v
	return s
}

func (s *ListTableMetasRequest) Validate() error {
	return dara.Validate(s)
}
