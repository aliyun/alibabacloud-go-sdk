// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvId(v string) *ListEnvsRequest
	GetEnvId() *string
	SetEnvName(v string) *ListEnvsRequest
	GetEnvName() *string
	SetPageNumber(v int32) *ListEnvsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEnvsRequest
	GetPageSize() *int32
}

type ListEnvsRequest struct {
	// The ID of the environment. If you specify this parameter, the operation returns the information about the environment identified by the ID.
	//
	// example:
	//
	// 10YPA8H
	EnvId *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	// The keyword of the environment name. If you specify this parameter, the operation returns the information about the environments whose names contain the keyword.
	//
	// example:
	//
	// test-create
	EnvName *string `json:"EnvName,omitempty" xml:"EnvName,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of environments per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListEnvsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnvsRequest) GoString() string {
	return s.String()
}

func (s *ListEnvsRequest) GetEnvId() *string {
	return s.EnvId
}

func (s *ListEnvsRequest) GetEnvName() *string {
	return s.EnvName
}

func (s *ListEnvsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEnvsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEnvsRequest) SetEnvId(v string) *ListEnvsRequest {
	s.EnvId = &v
	return s
}

func (s *ListEnvsRequest) SetEnvName(v string) *ListEnvsRequest {
	s.EnvName = &v
	return s
}

func (s *ListEnvsRequest) SetPageNumber(v int32) *ListEnvsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEnvsRequest) SetPageSize(v int32) *ListEnvsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEnvsRequest) Validate() error {
	return dara.Validate(s)
}
