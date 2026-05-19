// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectBindMachineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListFileProtectBindMachineRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListFileProtectBindMachineRequest
	GetPageSize() *int32
}

type ListFileProtectBindMachineRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFileProtectBindMachineRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectBindMachineRequest) GoString() string {
	return s.String()
}

func (s *ListFileProtectBindMachineRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListFileProtectBindMachineRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFileProtectBindMachineRequest) SetCurrentPage(v int32) *ListFileProtectBindMachineRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListFileProtectBindMachineRequest) SetPageSize(v int32) *ListFileProtectBindMachineRequest {
	s.PageSize = &v
	return s
}

func (s *ListFileProtectBindMachineRequest) Validate() error {
	return dara.Validate(s)
}
