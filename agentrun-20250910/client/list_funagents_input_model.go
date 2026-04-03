// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunagentsInput interface {
	dara.Model
	String() string
	GoString() string
	SetFunagentName(v string) *ListFunagentsInput
	GetFunagentName() *string
	SetPageNumber(v int) *ListFunagentsInput
	GetPageNumber() *int
	SetPageSize(v int) *ListFunagentsInput
	GetPageSize() *int
	SetStatus(v string) *ListFunagentsInput
	GetStatus() *string
}

type ListFunagentsInput struct {
	FunagentName *string `json:"funagentName,omitempty" xml:"funagentName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Status   *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListFunagentsInput) String() string {
	return dara.Prettify(s)
}

func (s ListFunagentsInput) GoString() string {
	return s.String()
}

func (s *ListFunagentsInput) GetFunagentName() *string {
	return s.FunagentName
}

func (s *ListFunagentsInput) GetPageNumber() *int {
	return s.PageNumber
}

func (s *ListFunagentsInput) GetPageSize() *int {
	return s.PageSize
}

func (s *ListFunagentsInput) GetStatus() *string {
	return s.Status
}

func (s *ListFunagentsInput) SetFunagentName(v string) *ListFunagentsInput {
	s.FunagentName = &v
	return s
}

func (s *ListFunagentsInput) SetPageNumber(v int) *ListFunagentsInput {
	s.PageNumber = &v
	return s
}

func (s *ListFunagentsInput) SetPageSize(v int) *ListFunagentsInput {
	s.PageSize = &v
	return s
}

func (s *ListFunagentsInput) SetStatus(v string) *ListFunagentsInput {
	s.Status = &v
	return s
}

func (s *ListFunagentsInput) Validate() error {
	return dara.Validate(s)
}
